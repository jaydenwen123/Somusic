package kugou

import (
	"Somusic/common"
	"Somusic/util"
	"bufio"
	"flag"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/tidwall/gjson"
	"os"
	"strings"
	"time"
)

var (
	//歌曲信息
	songInfos  = make([]*SongInfo, 0)
	mvInfos  = make([]*MVInfo, 0)
	downloadSongInfos =make([]*SongInfo,0)
	downloadMVInfos =make([]*MVInfo,0)
	argMessage = `命令行版的音乐播放器使用说明：
Usage of spider:
  -keyword string
        this is your search keyword!!
`
)

var keyword string

func init() {
	//1.接收关键词
	//keyword=AcceptFromConsole()
}

//下载关键词搜索到的歌曲
func DownloadSearchMusicInfo() {
	//记录下载时间
	//start := time.Now()
	//2.发送请求，接收json数据
	search := strings.Replace(searchUrl, "{}", keyword, -1)
	search = strings.Replace(search, "$", "1", -1)
	//获得总共的条数
	searchInfos := common.RequestJson(search, HEADER)
	total := GetTotal(searchInfos)
	search = strings.Replace(searchUrl, "{}", keyword, -1)
	search = strings.Replace(search, "$", total, -1)
	logs.Info("恭喜你，总共搜索到", total, "首歌曲！！！")
	logs.Info("搜索歌曲的链接：", search)
	logs.Info("正在搜索数据中，请耐心等待.....")
	searchInfos = common.RequestJsonWithRetry(search, HEADER)
	//初始化保存歌曲目录
	//downloadSaveMVDir
	saveBasePath := downloadSaveSongDir+ keyword
	logs.Info("正在初始化目录,请等待......")
	util.InitDir(saveBasePath)
	//initSaveDir(downloadSaveMVDir)
	util.InitDir(downloadSaveSongDir)
	logs.Info("初始化目录完毕.....")
	//解析json数据放在保存前面，采用go协程去解析，解约时间
	//解析json数据，并得到hash
	if gjson.Valid(searchInfos) {
		logs.Info("正在解析歌曲数据，请等待........")
		parsed := make(chan bool)
		go func(done chan bool) {
			songInfos = ParseSearchSongsInfos(searchInfos)
			util.Save2JsonFile(songInfos, saveBasePath+string(os.PathSeparator)+"songs.json")
			done <- true
		}(parsed)

		//保存json数据到文件中
		filePath := saveBasePath + string(os.PathSeparator)+"data.json"
		util.SaveJsonStr2File(searchInfos, filePath)
		<-parsed
		logs.Info("解析歌曲数据完毕.......")
		logs.Info("保存歌曲的目录：", saveBasePath)
	}else{
		logs.Error("由于服务器原因，数据获取失败，请重试...")
	}
	//4.下载歌曲
	//DownloadByRank(hashs, saveBasePath)
	//完成时间
	//end := time.Now()
	//logs.Info("总共下载", downloadMaxCount, "个文件!总耗时为", fmt.Sprintf("%v", end.Sub(start)))
}

//根据歌曲排名，下载前downMaxCount首歌曲
func DownloadByRank(hashs []gjson.Result, saveBasePath string) {
	if downloadMaxCount > len(hashs) {
		downloadMaxCount = len(hashs)
	}
	done := make(chan DownloadMsg, downloadMaxCount)
	finish := make(chan bool)
	for index, hash := range hashs {
		if index >= downloadMaxCount {
			break
		}
		// 并发下载
		go DownloadMusic(hash.String(), saveBasePath, ".mp3", index+1, done)
		if index == 0 {
			logs.Info("---------------正在多线程下载，请耐心等待--------------")
			go func() {
				for i := 0; i < downloadMaxCount; i++ {
					downloadInfo := <-done
					if downloadInfo.Success {
						logs.Info("第  (", downloadInfo.FileId, ")  个文件  [", downloadInfo.FileName, "]  ", "下载成功")
					} else {
						logs.Error("第  (", downloadInfo.FileId, ")  个文件  [", downloadInfo.FileName, "]  ", "下载失败")
					}
				}
				finish <- true
			}()

		}
	}
	<-finish
}

//从所有的歌曲json数据中获取所有的hash值
func ParseSearchSongsInfos(songInfos string) []*SongInfo {
	names := gjson.Get(songInfos, "data.lists.#.FileName").Array()
	albumNames := gjson.Get(songInfos, "data.lists.#.AlbumName").Array()
	durations := gjson.Get(songInfos, "data.lists.#.Duration").Array()
	fileSizes := gjson.Get(songInfos, "data.lists.#.FileSize").Array()
	hashs := gjson.Get(songInfos, "data.lists.#.FileHash").Array()
	songs := make([]*SongInfo, len(names))
	for index, name := range names {
		songs[index] = NewSongInfo(
			fmt.Sprintf("%d", index+1),
			 name.String(),
			albumNames[index].String(),
			durations[index].String(),
			fileSizes[index].String(),
			hashs[index].String())
	}
	return songs
}

//从接送文件中获取总的条数
func GetTotal(info string) string {
	return gjson.Get(info, "data.total").String()
}


//根据每首歌曲的hash值下载歌曲
func DownloadMusic(hash,savePath string,fileSuffix string,fileIndex int,done chan DownloadMsg) {
	//1.首先根据hash值获取歌曲的json数据
	songUrl:=strings.Replace(songInfoTemplateUrl,"{}",hash,-1)
	songJson := common.RequestJson(songUrl, HEADER)
	//估计是cookie变化了，所以需要重新设置一次cookie
	//time.Sleep(time.Millisecond*500)
	//2.提取歌曲下载链接
	//audio_name,img,play_url
	song := ExactSongInfo(songJson)
	song.SourceUrl=songUrl
	//打印歌曲信息
	fmt.Println("##########################第"+fmt.Sprintf("%d",fileIndex)+"首歌曲信息#######################")
	fmt.Println(song.ToString())
	//fmt.Println()
	if song.Url==""{
		logs.Error("歌曲没有下载链接！！！")
		done<-DownloadMsg{FileName:song.Name,FileId:fileIndex,Success:false}
		return
		//time.Sleep(time.Millisecond*20)
		//DownloadMusic(hash,savePath,done)
		//return
	}
	//3.正式下载 为了提升用户体验，先不下载重试
	err := util.Download(song.Url, savePath+"/"+song.Name+fileSuffix)
	if err!=nil{
		//logs.Error("歌曲：",song.Name,"下载失败...")
		done<-DownloadMsg{FileName:song.Name,Success:false,FileId:fileIndex}
	}else{
		//logs.Info("歌曲：",song.Name,"下载成功...")
		done<-DownloadMsg{FileName:song.Name,Success:true,FileId:fileIndex}
	}
	time.Sleep(time.Millisecond*50)
}

//提取歌曲的重要信息
func ExactSongInfo(songJson string) *SongInfo {
	song:=&SongInfo{}
	song.Name = gjson.Get(songJson, SONG_NAME_PATH).String()
	song.Img = gjson.Get(songJson, IMG_URL_PATH).String()
	song.Url = gjson.Get(songJson, PLAY_URL_PATH).String()
	return song
}



//从控制台接收参数
func AcceptFromConsole() string {
	var keyword string
	flag.StringVar(&keyword, "keyword", "", "this is your search keyword!!")
	flag.Parse()
	//用户没有输入参数
	if keyword == "" || len(keyword) == 0 {
		keyword = AcceptInputKeyWord()
	}
	return keyword
}

//接收用户从控制台输入的关键词
func AcceptInputKeyWord() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s", "please input your keyword:")
	//_, err := fmt.Scanf("%s\n", &keyword)
	line, _, err := reader.ReadLine()
	keyword = string(line)
	if err != nil {
		logs.Error("accept the keyword error!!!,please try again.", err.Error())
	}
	return keyword
}
