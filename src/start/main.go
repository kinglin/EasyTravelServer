// start project main.go
package main

import (
	"controller"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

//收到请求结构体
type Request struct {
	What string `json:"what"`
}

//定义网络常量
var (
	host   = "115.156.249.22"
	port   = "12345"
	remote = host + ":" + port
	data   = make([]byte, 1000000)
)

func main() {

	//服务器启动
	fmt.Println("server start")

	lis, err := net.Listen("tcp", remote)
	defer lis.Close()

	if err != nil {
		fmt.Println("Error when listen: ", remote)
		os.Exit(-1)
	}

	//主线程开始for循环，用于接收不同客户端请求，并分发处理
	for {
		var requestFromClient string
		connect, err := lis.Accept()
		if err != nil {
			fmt.Println("Error accepting client: ", err.Error())
			os.Exit(0)
		}

		//显示新连接的客户端
		fmt.Println("New connection: ", connect.RemoteAddr())

		length, _ := connect.Read(data)

		//拿到用户请求数据
		requestFromClient = string(data[0:length])

		fmt.Printf("%s said: %s", connect.RemoteAddr(), requestFromClient)

		//将用户请求的json解析，然后分发到线程
		var jsonRequest Request
		if err := json.Unmarshal([]byte(requestFromClient), &jsonRequest); err != nil {
			fmt.Println(err)
		}

		switch jsonRequest.What {
		//注册
		case "register":
			go controller.HandleRegister(requestFromClient, connect)
		default:
			fmt.Println("no such request match")
		}
	}

}
