package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var urls = []string{
		"http://un.org",
	}
	var wg sync.WaitGroup
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go httpCheck(urls[i], &wg)
	}
	wg.Wait()
}

func httpCheck(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	start := time.Now()
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(resp)
	fmt.Println(resp.StatusCode, time.Since(start), url)
}
