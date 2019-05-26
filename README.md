# Somusic
this is the command line music downloader,which contains lots of function such as search song and mv from kugou website、download song and mv with single or batch.、list search songs or mvs、show download songs or mvs .above functions have matching command.you can use help or h to find doc. finally there left a function that is  play music and play mv.i will complete this function in the near future.

# Getting Started #

## Installing ##

>  1. create a directory which is used to store the project.for example: `cd d:\golang\workspace\.`
>  2. you should execute this command `git clone https://github.com/wenxiaofeiCode/Somusic.git`
>  3. if you want to move this project into %gopath%,then you can move it to your gopath's src directory.

Now this will retrieve the project to your local directory, you can start your travel.


## Help Document ##
the somusic support many functions which also can match  with  it's  command. there are all supported functions listed in the table.

| Command | Parameter |	Function |	Description |
|:---------:|:------------:|:------------|:--------------| 
| gboard |no paramter |download rank board | download the kugou rank board all song |
| lsong | [max count] or <start-end> | list searched songs | show the asc range musics list(note:the shown songs is not downloaded) |
| lmv | [max count] or <first-end> | list searched mvs |	show the asc range mv list,it's also like lsong. |
| gsong | [songid] or <first1-end1,first2-end2...> or <songid1,...,first1-end1,songid2,songid3...> | get or download songs from remote server | download one music or download according the range.it  support download single song、batch(first-end) songs 、 discontinuous(songid1,songid5,songid8,...) songs and mixed above all ways |
| gmv |[mvid] or <first1-end1,first2-end2...> or <mvid1,...,first1-end1,mvid5,mvid7...> |	download mv files with many ways |	gmv usage is also familiar with gsong.you can also see gsong's usage |
| psong |[songid] |	play the selected song | this function is just not completed.it will fill in the near future |
| pmv |[mvid] |	play the selected mv |	this function is just not completed.it will fill in the near future |
| qsong |[keyword] |	query song | query song with the  inputed keyword from the kugou music website |
| qmv |[keyword] |	query song | query mv with the inputed keyword from the kugou music website |
| ssong |no paramter |	Function |	Description |
| smv |no paramter |	Function |	Description |
| chstyle |no paramter |	Function |	Description |
| chdelimiter |no paramter |	Function |	Description |
| mvpath |no paramter |	Function |	Description |
| songpath |no paramter |	Function |	Description |
| chmvpath |no paramter |	Function |	Description |
| chsongpath |no paramter |	Function |	Description |
| help |no paramter |	Function |	Description |
| quit |no paramter |	Function |	Description |
| exit |no paramter |	Function |	Description |
| cls |	no paramter |Function |	Description |









| 一个普通标题 | 一个普通标题 | 一个普通标题 |
| ------ | ------ | ------ |
| 短文本 | 中等文本 | 稍微长一点的文本 |
| 稍微长一点的文本 | 短文本 | 中等文本 |

| 左对齐标题 | 右对齐标题 | 居中对齐标题 |
| :------| ------: | :------: |
| 短文本 | 中等文本 | 稍微长一点的文本 |
| 稍微长一点的文本 | 短文本 | 中等文本 |

