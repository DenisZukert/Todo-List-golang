package main

import (
	"fmt"
	"net/http"
)

func taskHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "<h1>Hello from GO-server!</h1>")
}

func main() {
	http.HandleFunc("/", taskHandler)
	fmt.Println("Server is running on localhost:3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Server start error:", err)
		return
	}
}
