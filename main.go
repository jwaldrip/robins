package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jwaldrip/odin/cli"
	"github.com/jwaldrip/robins/proxy"
)

var version = "0.0.1"
var app = cli.New(version, "A multi-host and multi-port random robin proxy", run)

func init() {
	app.DefineParams("hosts", "ports")
}

func main() {
	app.Start()
}

func run(c cli.Command) {
	hosts := strings.Split(c.Param("hosts").Get().(string), ",")
	ports := strings.Split(c.Param("ports").Get().(string), ",")

	// Create an independent proxy for each port, except the last, which will be blocking
	for _, port := range ports[:len(ports)-1] {
		go startProxy(hosts, port)
	}

	// Start the proxy for the last port
	startProxy(hosts, ports[len(ports)-1])
}

func startProxy(hosts []string, port string) {
	p := proxy.New(hosts, port)
	err := p.Listen()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
