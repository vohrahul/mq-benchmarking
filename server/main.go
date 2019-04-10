package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/vohrahul/mq-benchmarking/server/service"
)

const defaultPort = 9500

func main() {
	var (
		port            = flag.Int("port", defaultPort, "service port")
	)
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	config := &daemon.Config{}

	d, err := daemon.NewDaemon(config)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Service started on port %d...\n", *port)
	if err := d.Start(*port); err != nil {
		fmt.Println(err)
		panic(err)
	}
}
