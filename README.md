# Somusic #
This is the command line music downloader,which contains lots of functions such as search song and mv from kugou website、download song and mv with single or batch.、list search songs or mvs、show download songs or mvs .Above functions have matching commands.You can use help or h to find doc.Finally there left a function that is  play music and play mv.I will complete this function in the near future.

## Who Should See This Project? ##

- If you want to learn golang quickly  with a practical project;
- If you want to get a better understand about how  the command line software is runned;
- If you want to enter the Network Spider World quickly;
- If you are interested in the music website,and you want to build yourself music player;
- this project is absolutely nice for you.


## What Can You Get? ##
1. Spider the Kugou Music Website by Golang language.
2. Concurrnet download the file(including not only mp3 file and mp4 file,but also binary files and text files.) by Golang  goroutine and Channel.
3. Master skills of JSON and Golang struct、 interface、http、regexp and Golang basic knowledge.
4. Practice to use the famous librarys such as:goquery、gjson...
5. Understand how the command line tool is working. Such as windows cmd、Golang Web Framework  beego's bee tool.
6. Analysis the Http's Interface about Kugou Music Website.
7. Get the skill of Spider by Golang.

# Getting Started #

## Installing ##

>  1. create a directory which is used to store the project.for example: `cd d:\golang\workspace\.`
>  2. you should execute this command `git clone https://github.com/jaydenwen123/Somusic.git`
>  3. if you want to move this project into %gopath%,then you can move it to your gopath's src directory.

Now this will retrieve the project to your local directory, you can start your travel.


## Help Document ##
1.the somusic support many functions which also can match  with  it's  command. there are all supported functions listed in the table.

| Command | Parameter |		Function 	|	Description |
|:---------:|:------------|:------------|:--------------|
| gboard |no paramter | download rank board | download the kugou rank board all song |
| lsong | [max songid] or <start-end> | list searched songs | show the asc range musics list(note:the shown songs is not downloaded) |
| lmv | [max mvid] or <first-end> | list searched mvs |	show the asc range mv list,it's also like lsong. |
| gsong | [songid] or <first1-end1,first2-end2...> or <songid1,...,first1-end1,songid2,songid3...> | get or download songs from remote server | download one music or download according the range.it  support download single song、batch(first-end) songs 、 discontinuous(songid1,songid5,songid8,...) songs and mixed above all ways |
| gmv |[mvid] or <first1-end1,first2-end2...> or <mvid1,...,first1-end1,mvid5,mvid7...> |	download mv files with many ways |	gmv usage is also familiar with gsong.you can also see gsong's usage |
| psong |[songid] |	play the selected song | this function is just not completed.it will fill in the near future |
| pmv |[mvid] |	play the selected mv |	this function is just not completed.it will fill in the near future |
| qsong |[keyword] |	query song | query song with the  inputed keyword from the kugou music website |
| qmv |[keyword] |	query song | query mv with the inputed keyword from the kugou music website |
| ssong |no paramter |	show all downloaded local songs |	show the downloaded songs list,the songid you can use to play song in the list |
| smv |no paramter |	show all downloaded local mvs |	this comamnd's usage is similiar to the ssong.you can also see ssong. |
| chstyle |[new style string] |	change the style |change the style with the new style string.this's command is same to the command `style`.|
| style |[new style string] | change the style |change the style with new style string.|
| chdelimiter |[new delimiter chars] |  change the delimiter chars |it will change the old delimiter chars.|
| delimiter |[new delimiter chars] |  change the delimiter chars |it will change the old delimiter chars.this command function is same to the command `chdelimiter`|
| mvpath |no paramter |	show current downloaded mvpath |	show the save downloaded mvpath |
| songpath |no paramter |	show current downloaded songpath |	show the save downloaded songpath |
| chmvpath |[newmvpath]|	change the current mvpath |	change the save downloaded mv path. Use the `~` to recovery the default dirctory |
| chsongpath |[newsongpath] |	change the current songpath |	change the save downloaded song path. Use the `~`to recovery the default dirctory |
| help or h |no paramter |	show help doc |	show the help information |
| quit or CTRL+C |no paramter |	quit the program |	quit the program |
| exit or CTRL+C |no paramter |	quit the program |	quit the program |
| cls or clear | no paramter |clear the log information |	clear the log information.In current version,it is only support the windows clear log.next version will add the linux clear log function |


2.the following is the  picture of help document,which is runned int the goland ide.
![help document](https://github.com/jaydenwen123/Somusic/blob/master/usages/help1.png)
![help document](https://github.com/jaydenwen123/Somusic/blob/master/usages/help2.png)
![help document](https://github.com/jaydenwen123/Somusic/blob/master/usages/help3.png)

## Usages ##
In this section. I will use the search song keyword:`bigbig` and `天使的翅膀`, and the search mv keyword :`Falling Down` and `小幸运` as the example to show how to use the somusic program.

1.**search song with keyword.**

>command: `qsong bigbig`(bigbig)
>![qsong](https://github.com/jaydenwen123/Somusic/blob/master/usages/qsong.png)
>command: `qsong 天使的翅膀`(天使的翅膀)
>![qsong](https://github.com/jaydenwen123/Somusic/blob/master/usages/qsong天使的翅膀.png)

2.**search mv with keyword.**

>command: `qmv falling down`(falling down)
>![qmv](https://github.com/jaydenwen123/Somusic/blob/master/usages/qmv.png)
>command: `qmv 小幸运`(小幸运)
>![qmv](https://github.com/jaydenwen123/Somusic/blob/master/usages/qmv小幸运.png)


3.**list the searched song information.**

>command: `lsong`(小幸运)
>![lsong](https://github.com/jaydenwen123/Somusic/blob/master/usages/lsong.png)
>command: `lsong 11`(big big world)
>![lsong](https://github.com/jaydenwen123/Somusic/blob/master/usages/lsong2.png)
>command: `lsong `(天使的翅膀)
>![lsong](https://github.com/jaydenwen123/Somusic/blob/master/usages/lsong天使的翅膀.png)

4.**list the searched mv informtion.**

>command: `lmv`(Falling Down)
>![lmv](https://github.com/jaydenwen123/Somusic/blob/master/usages/lmv.png)
>![lmv](https://github.com/jaydenwen123/Somusic/blob/master/usages/lmv2.png)
>command:`lmv`(小幸运)
>![lmv](https://github.com/jaydenwen123/Somusic/blob/master/usages/lmv小幸运.png)

5.**download the searched song.**

>command: `gsong 3,6`(big big world)
>![gsong](https://github.com/jaydenwen123/Somusic/blob/master/usages/gsong.png)

6.**download the searched mv.**

>command: `gmv 1-10`(Falling Down)
>![gmv](https://github.com/jaydenwen123/Somusic/blob/master/usages/gmv.png)
>![gmv](https://github.com/jaydenwen123/Somusic/blob/master/usages/gmv2.png)
>command:`gmv 1-5`(小幸运)
>![gmv](https://github.com/jaydenwen123/Somusic/blob/master/usages/gmv小幸运.png)
>![gmv](https://github.com/jaydenwen123/Somusic/blob/master/usages/gmv小幸运2.png)

7.**show the local downloaded songs.**

>command: `ssong`
>![ssong](https://github.com/jaydenwen123/Somusic/blob/master/usages/ssong .png)
>![ssong](https://github.com/jaydenwen123/Somusic/blob/master/usages/ssong.png)

8.**show the local donwloaded mvs.**

>command: `smv`
>![smv](https://github.com/jaydenwen123/Somusic/blob/master/usages/smv.png)

9.**show the current saved download songs' directory.**

>command: `songpath`
>![songpath](https://github.com/jaydenwen123/Somusic/blob/master/usages/songpath.png)

10.**show the current saved download mvs' directory.**

>command: `mvpath`
>![mvpath](https://github.com/jaydenwen123/Somusic/blob/master/usages/mvpath.png)
>

11.**change the saved download mvs' directory.**

>command: `chmvpath D:\歌曲`
>![chmvpath](https://github.com/jaydenwen123/Somusic/blob/master/usages/chmvpath.png)

12.**change the saved download mvs' directory.**

>command: `chsongpath D:\歌曲`
>![chsongpath](https://github.com/jaydenwen123/Somusic/blob/master/usages/chsongpath.png)

13.**change the program command line style.**

>command: `style mimusic`
>![style](https://github.com/jaydenwen123/Somusic/blob/master/usages/style.png)

14.**change the program command line delimiter.**

>command: `delimiter #`
>![delimiter](https://github.com/jaydenwen123/Somusic/blob/master/usages/delimiter.png)

15.**show or find the help document.**

>command: `help`
>![help](https://github.com/jaydenwen123/Somusic/blob/master/usages/help1.png)
>![help](https://github.com/jaydenwen123/Somusic/blob/master/usages/help2.png)
>![help](https://github.com/jaydenwen123/Somusic/blob/master/usages/help3.png)

16.**quit or exit the program.**

>command: `exit`
>![quit](https://github.com/jaydenwen123/Somusic/blob/master/usages/exit.png)

17.**clear the log information.**

>command: `cls`
>![cls](https://github.com/jaydenwen123/Somusic/blob/master/usages/cls.png)


# Reference #
1. gjson(https://github.com/tidwall/gjson)
2. goquery(https://github.com/PuerkitoBio/goquery)
3. gorm(https://github.com/jinzhu/gorm)
4. beego orm(https://github.com/astaxie/beego/orm)
5. beego logs(https://github.com/astaxie/beego/logs)
6. regexp standard library(https://studygolang.com/pkgdoc)
7. net/http standard library(https://studygolang.com/pkgdoc)
8. channel&goroutine(https://gobyexample.com)

# What's need to Improve #

- 1.play song or play mv in reality.
- 2.config the variables into the file.such as save download song directory and mv directory,software command line style and delimiter.
- 3.add the cache module.which can improve the somusic's performance.

# Contace Me #
If you are interested in this project、 like coding or any questions,you can contact with me by following ways.
> QQ:2282186474
>
> WeChat:wen2282186474
>
> Eamil: 2282186474@qq.com



