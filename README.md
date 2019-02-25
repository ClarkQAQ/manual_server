# 人工服务器 manual_server
>每当你访问网页的时候就有一位程序员失去自由(雾

# 灵感来源
bilibili UP主
dimsmary 的一个视频:[【Dimsmary】人肉服务器？纯手动服务器](https://www.bilibili.com/video/av44212198 "【Dimsmary】人肉服务器？纯手动服务器")
一个十分有趣的UP主,但是UP用的是ESP8266然而我
> **没  有  E  S  P  8  2  6  6**

看着感觉挺好玩的但是他用的是大蛇蛇写的而且还没有开放下载...............然后就随随便便写了一下(写markdown的时间比写代码时间长系列)....(有不少bug但是懒得修了)



#运行方法
如果你下载的golang源码
```shell
go run manual_server.go
```
如果你下载的是编译好的版本
Linux(安卓也是一样):
```shell
chmod +x manual_server_linux_amd64
./manual_server_linux_amd64
```

Windows:
 ```
 直接双击
 ```
 
 
 # 使用方法
####  小白向教程,所以比较......大佬绕过....

![QWQ](https://i.loli.net/2019/02/26/5c746bfc8bd1b.png "QWQ")
![QAQ](https://i.loli.net/2019/02/26/5c746bfc7d615.png "QAQ")

```
 [Server Port(default:8080)]: #直接Enter默认使用8080端口,为了代码简介所以没有加任何过滤谨慎调皮
 [Path]: #客户端访问的路径根目录就是"/"
 [Head]: #客户端发送的请求头部信息,包含Cookie,UA,语言等信息
 [Get]: #客户端发送的get请求例:"127.0.0.1/?QAQ=QWQ"
 [HTTP Code]: #反馈给客户端的http状态码,比如200正常访问404页面不存在,301跳转,502 B....
 [Data]: #发送给客户端的信息,可以直接输入字符也可以输入html代码(html代码间不能有空格,一个BUG)
```
 `注意:在没有提示输入的时候千万不要输入任何内容,因为他会累计给下一次访问`
