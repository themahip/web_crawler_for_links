
# webcrawler

A simple Go-based web crawler that extracts links from a given website and resolves relative URLs to full URLs. It uses concurrency and a semaphore to limit the number of concurrent requests (5 at a time).


# features
    1. Crawl a website and extract all the links.
    2. Resolve relative URLs to full URLs.
    3. Use 5 concurrent requests via goroutines and semaphore.
# How To Use
    1. go run main.go
    2. Enter the base URL when prompted (e.g., https://example.com).
    3. The program will start crawling and display the found links.
