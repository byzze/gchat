package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"time"

	"gchat/server"

	"github.com/fatih/color"
)

var ccS = color.New(color.FgGreen, color.Bold) //绿色
var ccF = color.New(color.FgRed, color.Bold)   //红色
// go run .\server_main.go -port 8000
func main() {
	// flag 可用于接受由命令行传参数,多个参数可设置多个变量使用
	port := flag.Int("port", 7777, "设置服务器端口")
	// 解析，使用*port进行读取
	flag.Parse()
	fmt.Printf("◆ 服务器正在[:%s]启动...\n", strconv.Itoa(*port))
	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(*port))
	if err != nil {
		ccF.Println("服务器启动失败:", err)
		return
	}

	//启动一个新goroutine监听message
	go server.TransmitMsg()
	ccS.Println("◆ 已成功启动,正等待客户端连接……")
	for {
		//不停顿直接循环1000，集中并发量很大后，可能直接崩溃
		time.Sleep(20 * time.Millisecond)
		// 阻塞，监听链接
		conn, err := listener.Accept()
		if err != nil {
			ccF.Println("接受客户端连接失败:", err)
			return
		}
		ccS.Printf("[%v]的客户端已连接成功\n", conn.RemoteAddr())
		//并发聊天,一个客户端一个goroutine
		go server.HandleConnect(conn)
	}
}
