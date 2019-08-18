package ical

import (
	"log"
	"strings"
)

// Calendar struct follows the iCalendar object,
// defined as VCALENDAR in RFC 5545
type Calendar struct {
	// Required
	PRODID  string
	VERSION string
	// Optional
	CALSCALE string
	METHOD   string
	// Optional
	EVENT    []Event
	TODO     []string // TODO
	JOURNAL  []string // TODO
	FREEBUSY []string // TODO
	TIMEZONE []string // TODO
	ALARM    []string // TODO
}

// NewCalendar creates an instance of struct Calendar
func NewCalendar() *Calendar {
	c := new(Calendar)
	c.PRODID = "-//loozhengyuan//ical//EN"
	c.VERSION = "2.0"
	return c
}

func (c *Calendar) isReady() bool {
	if c.VERSION == "" {
		return false
	}
	if c.PRODID == "" {
		return false
	}
	return true
}

func (c *Calendar) generateCalendarProp() string {
	// Validate first
	status := c.isReady()
	if !status {
		log.Fatal("Event is not ready!")
	}

	// Create object
	var str strings.Builder

	// Write headers
	str.WriteString("BEGIN:VCALENDAR\r\n")

	// Write required params
	str.WriteString("VERSION:" + c.VERSION + "\r\n")
	str.WriteString("PRODID:" + c.PRODID + "\r\n")

	// Loop EVENT if exits
	for _, event := range c.EVENT {
		str.WriteString(event.generateEventProp())
	}

	// Write footers
	str.WriteString("END:VCALENDAR\r\n")

	return str.String()
}
