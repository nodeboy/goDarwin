package rtspserver

import (
	"net"
	"log"
)

type RtspServer struct{
	//config Config
	ip 	net.IP
	tcp net.TCPConn

}

func NewRtspServer() *RtspServer{
	return &RtspServer{}
}

func HandleConn(conn net.Conn){
	rtspsession:=NewRtspSession(conn)
	go rtspsession.HandleRtspSession()
}
func (server *RtspServer)StartServer(){

	listner,err:=net.Listen("tcp",server.ip.String())
	if err!=nil{
		log.Println(err)
	}
	for {
		tcpconn,err:=listner.Accept()
		if err!=nil{
			log.Println(err)
		}

		go HandleConn(tcpconn)
	}
}