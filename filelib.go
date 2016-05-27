package filelib

import (
	"io/ioutil"
	"os"
	"path"
)

// Match defines which files are valid
type Match func(info os.FileInfo, filename string) bool

// Search returns a channel where all the valid files will be sent
func Search(match Match, file ...string) <-chan string {
	matchFiles := make(chan string)
	go search(file, match, matchFiles)
	return matchFiles
}

func search(files []string, match Match, matchFiles chan<- string) {
	defer close(matchFiles)
	for _, file := range files {
		searchWrapped(file, match, matchFiles)
	}
}

func searchWrapped(file string, match Match, matchFiles chan<- string) {
	info, err := os.Stat(file)
	if err != nil {
		return
	}
	if match(info, file) {
		matchFiles <- file
	}
	if info.IsDir() {
		dirFiles, _ := ioutil.ReadDir(file)
		for _, f := range dirFiles {
			searchWrapped(path.Join(file, f.Name()), match, matchFiles)
		}
	}
}
