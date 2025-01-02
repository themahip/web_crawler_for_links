Go Web Crawler
A simple Go-based web crawler that extracts links from a given website and resolves relative URLs to full URLs. It uses concurrency and a semaphore to limit the number of concurrent requests (5 at a time).

Features
Crawl a website and extract all the links.
Resolve relative URLs to full URLs.
Use 5 concurrent requests via goroutines and semaphore.
How to Use:
Run the program.
Enter the base URL when prompted (e.g., https://example.com).
The program will start crawling and display the found links.
Key Functions
getRequest(): Sends a GET request to the target URL.
discoverLinks(): Extracts all links from the page.
checkrelative(): Resolves relative URLs to absolute URLs.
ResolveRelativeUrl(): Checks if URLs are from the same domain and resolves them.
Crawl(): Crawls the target URL and adds found links to the worklist.
Notes
The program respects the site's structure and limits concurrent requests to 5.
Make sure to follow the website’s robots.txt policy when crawling.
