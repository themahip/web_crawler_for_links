package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var userAgents = []string{ //can store multiple user agent so that request will be send through multiple users.
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36",
}

func getRequest(targetUrl string) (*http.Response, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", targetUrl, nil)
	req.Header.Set("User-Agent", userAgents[0])
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func discoverLinks(res *http.Response, baseDomain string) []string {
	if res != nil {
		doc, _ := goquery.NewDocumentFromResponse(res)
		foundUrls := []string{}
		if doc != nil {

			doc.Find("a").Each(func(i int, s *goquery.Selection) {
				res, _ := s.Attr("href")
				foundUrls = append(foundUrls, res)
			})
		}
		return foundUrls
	}
	return nil
}

// checking relative url and adding to the base url
func checkrelative(baseUrl string, foundUrl string) string {
	if strings.HasPrefix(foundUrl, "/") {
		return fmt.Sprintf("%s%s", baseUrl, foundUrl)
	} else {
		return foundUrl
	}
}

func ResolveRelativeUrl(baseUrl string, foundUrl string) (bool, string) {
	resultHref := checkrelative(baseUrl, foundUrl)
	baseParse, _ := url.Parse(baseUrl)
	resultParse, _ := url.Parse(resultHref)
	if baseParse != nil && resultParse != nil {
		if baseParse.Host == resultParse.Host {
			return true, resultHref
		} else {
			return false, ""
		}
	}
	return false, ""
}

var token = make(chan struct{}, 5) // implementing symaphore of 5

func Crawl(targetUrl string, baseDomain string) []string {
	token <- struct{}{} // adding to symaphore
	res, _ := getRequest(targetUrl)
	<-token // taking out from symaphore
	links := discoverLinks(res, baseDomain)
	foundUrl := []string{}
	for _, link := range links {
		ok, url := ResolveRelativeUrl(baseDomain, link)
		if ok {
			if url != "" {
				foundUrl = append(foundUrl, url)
				fmt.Println(url)
			}
		}
	}

	return foundUrl
}
func main() {
	var baseDomain string
	worklist := make(chan []string) // stores all the link
	fmt.Println("enter the base url")
	fmt.Scanln(&baseDomain)
	go func() {
		worklist <- []string{baseDomain}

	}()

	seen := make(map[string]bool) // to check if we have gone through every link or not
	// each iteration decreasing n so that crawl is made done
	for n := 1; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func() {
					foundList := Crawl(link, baseDomain)
					if foundList != nil {
						worklist <- foundList
					}
				}()
			}
		}
	}
}
