package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Repositorie struct
type Repositorie struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Stars       int64  `json:"stargazers_count"`
}

func main() {
	_ = githubStars("go")
}

func githubStars(lang string) error {

	resp, err := http.Get("http://www.mocky.io/v2/5b65dc7e330000d60df6aa8d")
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var repositories []Repositorie

	err = json.Unmarshal(body, &repositories)

	if err != nil {
		log.Panicln(err)
	}

	r, err := json.Marshal(repositories)

	err = ioutil.WriteFile("stars.json", r, 0644)
	if err != nil {
		log.Panicln(err)
	}

	return nil
}
