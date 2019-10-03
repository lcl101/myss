package core

import (
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type SystemProxySettings interface {
	TurnOnGlobProxy()
	TurnOffGlobProxy()
}

var sigs = make(chan os.Signal, 1)

func resetProxySettings(proxySettings SystemProxySettings) {
	for {
		select {
		case <-sigs:
			log.Print("Flora-kit is shutdown now ...")
			if nil != proxySettings {
				proxySettings.TurnOffGlobProxy()
			}
			time.Sleep(time.Duration(2000))
			os.Exit(0)
		}
	}
}

func initProxySettings(bypass []string, addr string) {
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	var proxySettings SystemProxySettings
	if runtime.GOOS == "windows" {
		w := &windows{addr}
		proxySettings = w
	} else if runtime.GOOS == "darwin" {
		d := &darwin{bypass, addr}
		proxySettings = d
	}
	if nil != proxySettings {
		proxySettings.TurnOnGlobProxy()
	}
	go resetProxySettings(proxySettings)
}
