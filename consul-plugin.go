package consul_plugin

import (
	"github.com/jinzhu/configor"
	"github.com/micro/go-micro/registry"
	"github.com/pku-hit/consul"
)

var Registry = struct {
	Type string
	Host string
}{}

var Reg registry.Registry

func init() {
	configor.Load(&Registry, "config/registry.yml")

	Reg = consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			Registry.Host,
		}
	})
}
