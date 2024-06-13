# Web Scraper with Go and Colly

This Go program demonstrates a basic web scraper using the Colly library. It allows you to input multiple website URLs and scrapes information such as title, meta description, headings, paragraphs, and outgoing links from each website.

## Requirements

- Go programming language (Golang)
- Colly library (github.com/gocolly/colly)

## Installation

1. **Install Go Language** if you haven't already by clicking on the link provided below: 

    [Go Installation](https://go.dev/doc/install)

2. **Install Colly**

    ```bash
    go get -u github.com/gocolly/colly/...
    ```

3. Download **main.go** file from this repository.

4. Run the program
    ```bash
    go run main.go
    ````

## Usage

- When prompted, enter the URLs of the websites you want to scrape. Type quit to stop entering URLs and start scraping.

- The program will output the following information for each website:

    - Website title
    - Meta description
    - Headings (h1)
    - Paragraphs (p)
    - Outgoing links (a[href])

- Make sure you have a stable internet connection as the program fetches data from the live websites.

## Notes

- This program uses a mutex for the counter and to synchronize access to shared variables across goroutines.
- Error handling is implemented for failed requests to websites.