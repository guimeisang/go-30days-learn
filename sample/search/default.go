package search

type defaultMatcher struct{}

func init(){
	var mathcer defaultMatcher
	Register("default", matcher)
	
}

func ( m defaultMatcher) Search(feed *Feed, searchTerm string)([]*Result, error){
	return nil, nil
}