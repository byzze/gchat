package server

//userInfo 用于存储用户信息
type userInfo struct {
	name    string
	perC    chan []byte // 除了用户登录，推出时的消息，其他消息都由此传输
	AddUser chan []byte // 广播用户进入或退出
}
