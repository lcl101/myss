package core

import (
	"errors"
	"fmt"
	"net"
)

const (
	ServerTypeShadowSocks = "shadowsocks"
	ServerTypeCustom      = "custom"
	ServerTypeHttp        = "http"
	ServerTypeHttps       = "https"
	ServerTypeDirect      = "direct"
	ServerTypeReject      = "direct"

	LocalServerSocksV5 = "localSocksv5"
	LocalServerHttp    = "localHttp"

	socksVer5       = 5
	socksVer4       = 4
	httpProxy       = 71
	socksCmdConnect = 1

	typeIPv4 = 1 // type is ipv4 address
	typeDm   = 3 // type is domain address
	typeIPv6 = 4 // type is ipv6 address
)

type ProxyServer interface {
	//proxy type
	ProxyType() string
	//dial
	DialWithRawAddr(raw []byte, host string) (remote net.Conn, err error)
	//
	FailCount() int

	AddFail()
	//
	ResetFailCount()
}

var (
	errAddrType      = errors.New("socks addr type not supported")
	errVer           = errors.New("socks version not supported")
	errAuthExtraData = errors.New("socks authentication get extra data")
	errReqExtraData  = errors.New("socks request get extra data")
	errCmd           = errors.New("socks command not supported")
	errReject        = errors.New("socks reject this request")
	errSupported     = errors.New("proxy type not supported")
	errConnect       = errors.New("connection remote shadowsocks fail")
	errProxy         = errors.New("error proxy action")
)

var proxyConfig *ProxyConfig

//Run 运行代码
func Run(surgeCfg, geoipCfg string) {
	proxyConfig = LoadConfig(surgeCfg, geoipCfg)
	listenAddr := fmt.Sprintf("%s:%d", proxyConfig.LocalHost, proxyConfig.LocalSocksPort)
}
