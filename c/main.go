package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func CreateRequest(url string, wg *sync.WaitGroup, responses chan string) {

	defer wg.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println("error while create get request for url:", url, "error:", err)

	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error while sending get request for url:", url, "error:", err)
	}

	if resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusOK {
			responses <- string(resp.Status)
		}
	}

}

func GetResponses(urls []string) {

	var wg sync.WaitGroup
	responses := make(chan string, len(urls))

	for _, url := range urls {
		wg.Add(1)
		go CreateRequest(url, &wg, responses)
	}

	wg.Wait()
	close(responses)

	if len(responses) == len(urls) {
		for response := range responses {
			fmt.Println(response)
		}
	} else {
		fmt.Println("one or more than one of request failed")
	}
}

func main() {
	urls := []string{"http://www.google.com", "http://www.stackoverflow.com", "http://www.github.com", "http://www.twitter.com"}
	GetResponses(urls)
}
