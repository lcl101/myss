package main

import (
	"flag"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/lcl101/myss/core"
)

func main() {
	var configFile, geoipdb string
	execPath := getCurrPath()
	flag.StringVar(&configFile, "s", strings.Join([]string{execPath, "liclss.conf"}, ""), "specify surge config file")
	flag.StringVar(&geoipdb, "d", strings.Join([]string{execPath, "geoip.mmdb"}, ""), "specify geoip db file")
	flag.Parse()
	// configFile = "/home/aoki/work/bin/liclss.conf"
	// geoipdb = "/home/aoki/work/bin/geoip.mmdb"
	core.Run(configFile, geoipdb)
}

/*获取当前文件执行的路径*/
func getCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	sep := string(os.PathSeparator)
	splitstring := strings.Split(path, sep)
	size := len(splitstring)
	splitstring = strings.Split(path, splitstring[size-1])
	// ret := strings.Replace(splitstring[0], "\\", "/", size-1)
	return splitstring[0]
}
