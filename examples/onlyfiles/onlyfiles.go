package main

import (
	"filelib"
	"fmt"
	"os"
)

func main() {
	match := func(fi os.FileInfo) bool { return !fi.IsDir() }
	files := filelib.Search(match, os.Args[1:]...)
	for f := range files {
		fmt.Println(f)
	}
}
