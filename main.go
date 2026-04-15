package main

import "fmt"

// Result holds retrieved data
type Result struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

func main() {
	fmt.Println("Hello there!")
}
