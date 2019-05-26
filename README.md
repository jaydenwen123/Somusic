# Somusic
this is the command line music downloader,which contains lots of function such as search song and mv from kugou website、download song and mv with single or batch.、list search songs or mvs、show download songs or mvs .above functions have matching command.you can use help or h to find doc. finally there left a function that is  play music and play mv.i will complete this function in the near future.

# Getting Started #

## Installing ##

>  1. create a directory which is used to store the project.for example: `cd d:\golang\workspace\.`
>  2. you should execute this command `git clone https://github.com/wenxiaofeiCode/Somusic.git`
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
![help document](https://github.com/wenxiaofeiCode/Somusic/blob/master/usages/help1.png)
![help document](https://github.com/wenxiaofeiCode/Somusic/blob/master/usages/help2.png)
![help document](https://github.com/wenxiaofeiCode/Somusic/blob/master/usages/help3.png)

## Usages ##
In this section. I will use the search song keyword:`极致纯电音` and the search mv keyword :`三寸天堂` as the example to show how to use the somusic program.  

1. **search song with keyword.**


2. **search mv with keyword.**


3. **list the searched song information.**


4. **list the searched mv informtion.**


5. **download the searched song.**


6. **download the searched mv.**


7. **show the local downloaded songs.**


8. **show the local donwloaded mvs.**


9. **show the current saved download songs' directory.**


10. **show the current saved download mvs' directory.**


11. **change the saved download mvs' directory.**


12. **change the saved download mvs' directory.**


13. **change the program command line style.**


14. **change the program command line delimiter.**


15. **show or find the help document.**


16. **quit or exit the program.**


17. **clear the log information.**








