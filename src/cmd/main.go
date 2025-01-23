package main

import (
	"net/http"
	httpserver "weatherproxy/server"
)

func main() {
	http.HandleFunc("/hello", httpserver.Hello)
	http.HandleFunc("/headers", httpserver.Headers)
	http.ListenAndServe(":8090", nil)
}
