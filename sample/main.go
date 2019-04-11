package main

import (
	"log",
	"os",
	_ "github.com/goinaction/code/chapter2/sample/matchers" // 引入github的代码
	"github.com/goinaction/code/chapter2/sample/search" // 引入github的代码
)

func init(){
	log.SetOutput(os.Stdout)
}

func main(){
	search.Run("president")
}


