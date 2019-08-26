package consul_plugin

import (
	"fmt"
	"github.com/jinzhu/configor"
	"github.com/micro/go-micro/registry"
	"github.com/pku-hit/consul"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Registry = struct {
	Type string
	Host string
}{}

var Reg registry.Registry

func init() {
	config, err := ioutil.ReadFile("config/registry.yml")
	if err != nil {
		fmt.Print(err)
	}
	yaml.Unmarshal(config, &Registry)

	Reg = consul.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			Registry.Host,
		}
	})
}
