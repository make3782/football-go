// registrator
package services

import (
	etcdclient "github.com/coreos/etcd/client"
	// "golang.org/x/net/context"
	// "google.golang.org/grpc"
	"log"
	"os"
	"strings"
	"sync"
)

var (
	DEFAULT_ETCD_HOST = "http://127.0.0.1:2379"
)

type service_pool struct {
}

func (p *service_pool) init(names ...string) {
	machines := []string{DEFAULT_ETCD_HOST}
	// 如果设置了环境变量，则优先使用环境变量的地址
	if etcd_host := os.Getenv("ETCD_HOST"); etcd_host != "" {
		machines = strings.Split(etcd_host, ";")
	}
	cfg := etcdclient.Config{Endpoints: machines, Transport: etcdclient.DefaultTransport}
	client, err := etcdclient.New(cfg)
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}
	log.Printf("%v", client)
}

var (
	_default_pool service_pool
	once          sync.Once
)

func Init(names ...string) {
	once.Do(func() { _default_pool.init(names...) })
}
