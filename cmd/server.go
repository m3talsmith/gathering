package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/m3talsmith/gathering/pkg/server"
)

var url, host, port string

func init() {
	flag.StringVar(&host, "host", "localhost", "Host to run the server to run on")
	flag.StringVar(&port, "port", "3000", "Port to run the server on")
	flag.Parse()
	url = fmt.Sprintf("%s:%s", host, port)
}

func main() {
	log.Printf("Starting server at %s", url)
	s, err := server.Open(url)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(s)
}
