package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func requestHandlers() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	CreateDict("/usr/share/dict/words")

	for index := 0; index < 100; index++ {
		a := NewShortURL("test")
		fmt.Println(a.ID)
	}

}
