package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("http://localhost:8080")

	http.ListenAndServe(port, nil)
}
