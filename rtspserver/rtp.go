package rtspserver

import "net"

type  Rtp struct{
	udp *net.UDPConn
}

type RtpPaired struct{
	rtps [2]Rtp
	rtpid string
}