package main

import (
	"fmt"
	"log"
	"net/http"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Главная</h1>")
}

func redirectTo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Редирект на другую страницу</h1>")
}

func main() {
	http.HandleFunc("/", indexPage)
	http.HandleFunc("/redirect/to/", redirectTo)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

//func main() {
//	Shortener()
//}
//
//type Shortener interface {
//	Shorten(url string) string
//	Resolve(url string) string
//}
//
//func Shorten() string {
//
//}
