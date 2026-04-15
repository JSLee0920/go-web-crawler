package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// Result holds retrieved data
type Result struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}

func fetchURL(url string, wg *sync.WaitGroup, ch chan<- Result) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("Error fetching %s: %v", url, err)
		return
	}
	title := doc.Find("title").Text()
	ch <- Result{URL: url, Title: title}
}

func main() {
	fmt.Println("Hello there!")
}
