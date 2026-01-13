package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/xtls/xray-core/core"
	_ "github.com/xtls/xray-core/main/distro/all"
)

func main() {
	// 1. 定义命令行参数
	configFile := flag.String("config", "./config/config.json", "Path to the config file")
	flag.Parse()

	// 2. 读取配置文件
	configData, err := os.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	// 3. 解析配置并启动内核
	serverConfig, err := core.LoadConfig("json", configData)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	instance, err := core.New(serverConfig)
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}

	if err := instance.Start(); err != nil {
		log.Fatalf("Failed to start instance: %v", err)
	}

	log.Printf("V2Ray Reality Tunnel started with config: %s", *configFile)

	// 4. 监听退出信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	instance.Close()
}
