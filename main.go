package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, _ := url.Parse("http://localhost:8080/?page=1&page=2&a=1&")

	queryMap := u.Query()
	fmt.Println(queryMap["page"])
	fmt.Println(queryMap["a"])
	fmt.Println(queryMap["b"])
}
