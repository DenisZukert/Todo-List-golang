package main

import (
	"Todo/internal/routers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", routers.RegisterRoute)

	fmt.Println("Server is running on localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
