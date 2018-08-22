package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	// 	"log"
	// 	"time"
	// 	"github.com/coreos/etcd/client"
	// 	"github.com/kataras/iris"
)

type EtcdConfig struct {
	Addr     string
	Port     int
	Protocol string
}
type WebConfig struct {
	Port int
}
type Config struct {
	Etcd EtcdConfig
	Web  WebConfig
}

// read a config
func ReadConfig(filepath string) (*Config, error) {
	fmt.Printf(filepath)
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	cfg := Config{}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, err
}

func (conf *Config) GetEndpoint() string {
	return conf.Etcd.Protocol + "://" + conf.Etcd.Addr + ":" + fmt.Sprintf("%d", conf.Etcd.Port)
}
func (conf *Config) GetWebPort() string {
	return fmt.Sprintf("%d", conf.Web.Port)
}
