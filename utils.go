package ical

import (
	"io/ioutil"
	"log"
	"os"
	"time"
)

// FormatDateTime function formats a time.Time object
// in accordance to RFC 5545
func FormatDateTime(t time.Time) string {
	layout := "20060102T150405Z"
	timestamp := t.Format(layout)
	return timestamp
}

// OutputToFile function outputs the formatted calendar string
// to an ics file
func OutputToFile(filename string, content []byte, mode os.FileMode) {
	// Write content to file
	err := ioutil.WriteFile(filename, content, mode)
	if err != nil {
		log.Fatal(err)
	}
}
