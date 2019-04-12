package search

import(
	"log"
)

type Result struct{
	Field string
	Content string
}

type Matcher interface{
	search(feed *Feed, searchTerm string)([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result){
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil{
		log.PrintIn(err)
		return
	}

	for _, result := range searchResults{
		results <- result
	}
}

func Display(results chan *Result){
	for result := range results{
		log.Printf("%s: \n%s\n\n", result.Field, result.Content)
	}
}