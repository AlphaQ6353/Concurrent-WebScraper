// Website Links:
/*
	- https://example.com
	- https://news.ycombinator.com
	- https://golang.org
	- https://github.com
	- https://stackoverflow.com
	- https://www.wikipedia.org
	- https://www.reddit.com
	- https://www.bbc.com
	- https://www.cnn.com
	- https://www.nytimes.com
*/

package main

import (
	"fmt"
	"sync"
	"strings"
	"github.com/gocolly/colly"    // Web Scraper Module
)

var mutex sync.Mutex
var counter int

func scrap(web string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	mutex.Lock()                 // Mutex Lock for counter shared variable
	defer mutex.Unlock()         // Release Lock after changes 
	counter++
	fmt.Println(counter,":", web[8:])
	webScrape(web)
	wg.Done()
}

func webScrape(web string) {
	c := colly.NewCollector()
	msg:=make(chan string, 10)

	// Scrape Website Title
	c.OnHTML("title", func(e *colly.HTMLElement) {
		// title := e.Text
		msg<-e.Text
		// fmt.Printf("Title: %s\n", title)
		fmt.Printf("Title: %s\n", <-msg)
	})

	// Scrape MetaData Description
	c.OnHTML("meta[name=description]", func(e *colly.HTMLElement) {
		// description := e.Attr("content")
		msg<-e.Attr("content")
		// fmt.Printf("Meta Description: %s\n", description)
		fmt.Printf("Meta Description: %s\n", <-msg)
	})

	// Scrape Headings
	c.OnHTML("h1", func(e *colly.HTMLElement) {
		// heading := e.Text
		// lowerHeading := strings.ToLower(heading)
		msg<-strings.ToLower(e.Text)
		// fmt.Printf("Heading: %s\n", lowerHeading)
		fmt.Printf("Heading: %s\n", <-msg)
	})

	// Scrape Paragraphs
	c.OnHTML("p", func(e *colly.HTMLElement) {
		// paragraph := e.Text
		// lowerParagraph := strings.ToLower(paragraph)
		msg<-strings.ToLower(e.Text)
		// fmt.Printf("Paragraph: %s\n", lowerParagraph)
		fmt.Printf("Paragraph: %s\n", <-msg)
	})

	// Scrape Outgoing Links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// msg<-e.Attr(e.Attr("href"))
		fmt.Printf("Link: %s\n", link)
		// fmt.Printf("Link: %s\n", <-msg)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Request to %s failed: %v\n", r.Request.URL, err)
	})

	err := c.Visit(web)
	if err != nil {
		fmt.Printf("Failed to visit website: %v", err)
	}
}

func main() {
	var wg sync.WaitGroup
	var web string
	slice:=[]string{}
	for i:=0; true; i++ {
		fmt.Print("Enter website link (type `quit` to Exit): ")
		fmt.Scan(&web)
		if strings.ToLower(web)=="quit" {
			if len(slice)>0 {
				fmt.Printf("The following websites shall be Scraped:%v\n", slice)
			}
			break
		}
		slice=append(slice, web)
	}
	for j:=0; j<len(slice); j++ {
		wg.Add(1)
		go scrap(slice[j], &wg, &mutex)
		wg.Wait()
	}
}
