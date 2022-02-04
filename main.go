package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	defaultPath     string
	defaultPort     = 8080
	defaultRoute, _ = os.Executable()
)

var config struct {
	port   int
	routes string
}

func init() {
	flag.IntVar(&config.port, "port", defaultPort, "address to listen on (environment variable PORT)")
	flag.IntVar(&config.port, "p", defaultPort, "(alias for -port)")
	flag.StringVar(&config.routes, "d", defaultRoute, "define route directory")
	flag.Parse()
}
func main() {
	addr := fmt.Sprintf(":%v", config.port)
	err := InitServe(addr)
	if err != nil {
		log.Fatalf("start server: %v", err)
	}
}
