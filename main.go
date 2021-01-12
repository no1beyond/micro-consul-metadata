package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	sname := flag.String("sname", "", "service name")
	addr := flag.String("addr", "127.0.0.1:8500", "consul addr")
	flag.Parse()

	r := consul.NewRegistry(consul.Config(&api.Config{Address: *addr}))
	srvs, err := r.GetService(*sname)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, srv := range srvs {
		fmt.Printf("%+v\n", *srv)
		for _, node := range srv.Nodes {
			fmt.Printf("\t%+v\n", *node)
		}
	}
}
