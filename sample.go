package main

import (
  "fmt"
  "regexp"
  "strings"
  //"os"
)
func escape(name string) string {
	s := regexp.MustCompile(`[ <>:"/\\|?*%#]`).ReplaceAllString(name, "-")
	s = regexp.MustCompile(`--+`).ReplaceAllString(s, "-")
	return strings.Trim(strings.Replace(s, "--", "-", -1), "- ")
}
func main() {
    fmt.Println(escape("/Doc/use"))
}
