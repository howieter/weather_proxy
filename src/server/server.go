package server

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("zdarova zaibal")
}

func Headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
