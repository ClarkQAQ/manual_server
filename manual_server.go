 //习惯开头空行
package main

import (
	"fmt"
	"net/http"
)
var (
	data string
	httpsc int
	port = "8080"
	Y = "-------------------------"
)

// 处理Web请求
func web(w http.ResponseWriter, r *http.Request) { 
//习惯开头空行
	
	//打印请求信息
	fmt.Print("[Path]:",r.URL.Path,"\n")
	fmt.Print("[Head]:",r.Header,"\n")
	fmt.Print("[Get]:",r.URL.RawQuery,"\n")
	//提示输入和读取输入
	fmt.Print("\n",Y,"\n[HTTP Code]:")
	fmt.Scanln(&httpsc)
	fmt.Print(Y,"\n[Data]:")
	fmt.Scanln(&data)
	fmt.Print(Y,"\n")
	//把内容发给客户端
	w.WriteHeader(httpsc)
	fmt.Fprintf(w,data)
}

func main() {
//习惯开头空行
	fmt.Println("A simple manual server")
	//获取自定义端口信息
	fmt.Print("[Server Port(default:8080)]:")
	fmt.Scanln(&port)
	//打印服务器端口
	fmt.Println("[HTTP Port]:", port)
	fmt.Println("ClarkQAQ")
	//启动服务器
	http.HandleFunc("/", web)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
