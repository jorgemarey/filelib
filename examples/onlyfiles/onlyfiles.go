package main

import (
	"fmt"
	"os"

	"github.com/jorgemarey/filelib"
)

func main() {
	match := func(fi os.FileInfo, fn string) bool { return !fi.IsDir() }
	files := filelib.Search(match, os.Args[1:]...)
	for f := range files {
		fmt.Println(f)
	}
}
