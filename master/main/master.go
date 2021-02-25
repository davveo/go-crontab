package main

import (
	"flag"
	"fmt"
	"github.com/davveo/go-crontab/master"
	"runtime"
	"time"
)

var confFile string

func initArgs()  {
	flag.StringVar(&confFile, "config", "./master.json", "指定master.json")
	flag.Parse()
}

func initEnv()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var (
		err error
	)

	initArgs()
	initEnv()

	if err = master.InitConfig(confFile); err != nil {
		goto ERR
	}
	// 初始化服务发现模块
	if err = master.InitWorkerMgr(); err != nil {
		goto ERR
	}

	// 日志管理器
	if err =master.InitLogMgr(); err != nil {
		goto ERR
	}

	//  任务管理器
	if err = master.InitJobMgr(); err != nil {
		goto ERR
	}

	// 启动Api HTTP服务
	if err = master.InitApiServer(); err != nil {
		goto ERR
	}

	// 正常退出
	for {
		time.Sleep(1 * time.Second)
	}

	return
ERR:
	fmt.Println(err)
}