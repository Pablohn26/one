package goca

import (
	"fmt"

	"github.com/OpenNebula/one/src/oca/go/src/goca/flow/httpclient"
	"github.com/OpenNebula/one/src/oca/go/src/goca/flow/service"
)

var client httpclient.Client

func init() {
	client = httpclient.New("oneadmin", "opennebula", "http://10.10.0.56:2474")
}

func TestShow() {
	serv := service.Show(client, 4)

	fmt.Println(serv)
}

func TestShutdown() {
	result, body := service.TestShutdown(client, 3)

	fmt.Println(result)
	fmt.Println(body)
}

func TestDelete() {
	result, body := service.Delete(client, 3)

	fmt.Println(result)
	fmt.Println(body)
}