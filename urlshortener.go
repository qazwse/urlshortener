package main

import (
	"bufio"
	"log"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"time"
)

// ShortURL is a structure to hold information about a shortened URL
type ShortURL struct {
	ID           string
	URL          string
	Created      time.Time
	LastAccessed string
	NumVisits    uint
}

// Global var to hold our list of words in memory
// May move to non-global and just pass it around
// Not sure of the "correct" way to store permanent info yet
var dict []string
var dictLen int

// CreateDict reads a file, create an array containing each line
// Assumed that each line will be it's own word - e.g. /usr/share/dict/words
func CreateDict(filename string) {
	dict = make([]string, 1000)

	fp, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		dict = append(dict, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	dictLen = len(dict)
}

func randomString(numWords int) string {
	// Generates a string containing numWords random words from the global Dict array
	// This will be out url shortening scheme
	// E.g. server.com/HelloTestBeautifulScheme
	rand.Seed(time.Now().UnixNano())

	var str strings.Builder

	for i := 0; i < numWords; i++ {
		word := dict[rand.Intn(dictLen)]

		str.WriteString(strings.Title(word))
	}

	return str.String()
}

// NewShortURL creates a new short URL structure
func NewShortURL(userurl string) ShortURL {
	newurl := ShortURL{}

	// Check if the URL entered is a url
	_, err := url.Parse(userurl)
	if err != nil {
		log.Fatal(err)
	}

	newurl.ID = randomString(5)
	newurl.URL = userurl
	newurl.Created = time.Now()
	newurl.NumVisits = 0

	return newurl
}
