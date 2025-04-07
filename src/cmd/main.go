package main

import (
	"fmt"
	"net/http"
	handlerHttp "weatherproxy/handler"

	godotenv "github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/weather", handlerHttp.Weather)

	fmt.Println("Starting local server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
