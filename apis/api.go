package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/swanson", swansonHandler)
	http.HandleFunc("/cat", catHandler)
	http.HandleFunc("/taco", tacoHandler)
	http.HandleFunc("/xkcd", xkcdHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func swansonHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://ron-swanson-quotes.herokuapp.com/v2/quotes")
	if err != nil {
		log.Println("Error getting Swanson quote")
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body of request")
	}
	w.Write(body)
}

func catHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://thecatapi.com/api/images/get.php")
	if err != nil {
		log.Println("Error getting cat image")
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body of request")
	}
	w.Write(body)
}

func tacoHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://taco-randomizer.herokuapp.com/")
	if err != nil {
		log.Println("Error getting taco recipe")
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body of request")
	}
	w.Write(body)
}

func xkcdHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		log.Println("Error getting the xkcd")
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading body of request")
	}

	w.Write(body)
}
