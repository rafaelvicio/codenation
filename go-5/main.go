package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	_ = parseHTML("https://github.com/topics/go?l=go")
}

func parseHTML(url string) error {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".text-normal mb-1").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		langague := s.Find("h1").Text()
		title := s.Find("i").Text()
		fmt.Printf("Review %d: %s - %s\n", i, langague, title)
	})

	return fmt.Errorf("Erro ao ler arquivo")
}
