package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type URL struct {
	ID           string    `json:"id`
	OriginalURL  string    `json:"originl_url"`
	ShortUrl     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShort(OriginalURL string) string {
	hasher := md5.New()
	hasher.Write([]byte(OriginalURL))
	fmt.Println("hasher:", hasher)
	data := hasher.Sum(nil)
	fmt.Println("data:", data)
	hash := hex.EncodeToString(data)
	fmt.Println("data:", hash)
	fmt.Println("data:", hash[:8])

	return hash[:8]
}

func createURL(originalURL string) string {
	shortURL := generateShort(originalURL)
	id := shortURL
	urlDB[id] = URL{
		ID:           id,
		OriginalURL:  originalURL,
		ShortUrl:     shortURL,
		CreationDate: time.Now(),
	}
	return shortURL
}

func getURL(id string) (URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return URL{}, errors.New("URL not found!")
	}

	return url, nil
}

func RootPageURL(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "GET Method")
}

func ShortURLHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	shortUrl_ := createURL(data.URL)
	// fmt.Fprintf(w, shortUrl)

	response := struct {
		ShortUrl string `json:"short_url"`
	}{ShortUrl: shortUrl_}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func redirectURLHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/redirect/"):]
	url, err := getURL(id)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusNotFound)
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func main() {
	fmt.Println("Startin url shortner!")

	OriginlURL := "https://github.com/vivekparmar73"
	generateShort(OriginlURL)
	// createURL(OriginlURL)

	//handler function for handle all the requests
	http.HandleFunc("/", RootPageURL)
	http.HandleFunc("/shorten", ShortURLHandler)
	http.HandleFunc("/redirect/", redirectURLHandler)

	//start HTTP serevr on port 3000
	fmt.Println("Server start on 3000")
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		fmt.Println("Error occured on server start", err)
	}

}
