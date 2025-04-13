package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func removeEmptyStrings(slice []string) []string {
	var result []string
	for _, str := range slice {
		if str != "" {
			if !strings.HasPrefix(str, "with") {
				result = append(result, str)
			}
		}

	}
	return result
}

type Book_format struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("thestorygraph.com", "app.thestorygraph.com"),
	)

	// On every a element which has href attribute call callback
	c.OnHTML(".book-title-author-and-series", func(e *colly.HTMLElement) {
		// Splits the text by spaces
		var books []string = strings.Split(e.Text, "              ")

		// Remove all the extra spaces
		for i, book := range books {
			books[i] = strings.TrimSpace(strings.ReplaceAll(book, "   ", ""))
		}

		// Remove empty strings and the ones that start with "with"
		cleaned_books := removeEmptyStrings(books)

		// Assign the first element to Title and the last element to Author
		book := Book_format{Title: cleaned_books[0], Author: cleaned_books[len(cleaned_books)-1]}

		// Print the book details
		fmt.Printf("%s\n", book)
		fmt.Printf("Title: %s\n", book.Title)
		fmt.Printf("Author: %s\n", book.Author)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://app.thestorygraph.com/to-read/camfurbush")

	// ToDo Remove duplicate responses

}
