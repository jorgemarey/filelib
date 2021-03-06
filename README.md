# filelib

Simple Golang library to search for files that match some criteria.

Just:

```go
import (
  // Import the library
  "filelib"
  "fmt"
  "os"
  "strings"
)

func main() {
  // Create a function. This will be used against the files and
  // only those who match will be returned.
  match := func(fi os.FileInfo, fn string) bool {
    return !fi.IsDir() && strings.Contains(fn, "awesome")
  }
  // Search for them!
  files := filelib.Search(match, os.Args[1:]...)
  // filelib.Search return a channel
  for f := range files {
    fmt.Println(f)
  }
}
```
