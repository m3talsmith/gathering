package server

import (
	"fmt"
	"log"
	"net/http"
)

type httpserver struct {
	url      string
	cmd      chan string
	response chan string
	running  bool
}

func (s *httpserver) Start() error {
	go func(response, cmd chan string) {
		s.running = true
		log.Fatal(http.ListenAndServe(s.url, nil))
	}(s.response, s.cmd)
	return nil
}

func (s *httpserver) Stop() error {
	return nil
}

var instance Serveable

// Open creates a server connection to a url, returning back a pointer
// to the server so that you can stop it later or an error.
//
// Example:
//   s, _ := server.Open("localhost:3000") // Opens the server on localhost port 3000
//   s.Stop() // stops the server
func Open(url string) (Serveable, error) {
	instance = &httpserver{
		url:      url,
		cmd:      make(chan string, 1),
		response: make(chan string, 1),
		running:  false,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello http")
	})
	return instance, http.ListenAndServe(url, nil)
}
