package main

import (
	"github.com/yangtizi/go/config/jsonconfig"
	log "github.com/yangtizi/go/log/zaplog"
)

// SConfig 配置
type SConfig struct {
	Redis string `json:"redis,omitempty"`
	MySQL string `json:"mysql,omitempty"`
}

var cfg *SConfig

func main() {
	loadConfigDemo()
}

func loadConfigDemo() {
	cfg = &SConfig{}
	jsonconfig.JSONParsing("./config.json", cfg)
	log.Infof("加载的配置是cfg = [%v]", cfg)
}
