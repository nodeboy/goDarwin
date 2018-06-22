package rtspserver

import (
	"net"
	"github.com/gortc/sdp"
)

type RtspSession struct{
	rtppaireds map[string]RtpPaired
	streamnum int;
	url string
	conn net.Conn
	playsession map[string]*RtspSession
	pushsession *RtspSession
	sdpmessage        sdp.Message
}

func NewRtspSession(conn net.Conn) *RtspSession{

	return &RtspSession{conn:conn}
}

func (s * RtspSession)HandleRtspSession(){
	//handle rtspsession

}

