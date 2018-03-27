package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var urls = []string{
		"http://www.arthomson.com",
		"http://artgweb.arthomson.com",
		"http://support.arthomson.com",
		"http://icp-edm.arthomson.com",
		"http://icp-red.arthomson.com",
		"http://icp-sur.arthomson.com",
		"http://doc.arthomson.local/alfresco/",
		"http://reports.arthomson.local/Reports/Pages/Folder.aspx",
		"http://un.org",
		"http://jw.org",
		"http://www.preeper.org",
		"http://web.arthomson.com"}
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
