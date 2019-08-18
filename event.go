package ical

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Event struct type is a iCalendar component, defined
// as VEVENT in RFC 5545
type Event struct {
	// Required
	UID     string
	DTSTAMP *time.Time
	// Required only if Calendar object does not
	// specify the METHOD property
	DTSTART *time.Time
	// Optional
	CLASS       string
	CREATED     *time.Time
	DESCRIPTION string
	GEO         float64
	LASTMOD     *time.Time // LAST-MOD
	LOCATION    string
	ORGANIZER   string // TODO: SPECIAL TYPE: CAL_ADDRESS
	PRIORITY    uint8
	SEQ         uint8
	STATUS      string
	SUMMARY     string
	TRANSP      string
	URL         string
	RECURID     string // TODO: SPECIAL TYPE!
	RRULE       string // TODO: SPECIAL TYPE!
	// Optional but should not be declared together
	DTEND    *time.Time
	DURATION string // TODO: SPECIAL TYPE PT0H0M0S
	// Optional and supports multiple declarations
	ATTACH     []string
	ATTENDEE   []string // TODO: SPECIAL TYPE: CAL_ADDRESS
	CATEGORIES []string
	COMMENT    []string
	CONTACT    []string
	EXDATE     []*time.Time
	RSTATUS    []string
	RELATED    []string
	RESOURCES  []string
	RDATE      []*time.Time
	XPROP      []string // X-PROP
	IANAPROP   []string // IANA-PROP
}

// NewEvent creates an instance of struct Event
func NewEvent() *Event {

	// Creates a new instance
	e := new(Event)

	// Get timestamp
	currentTimestamp := time.Now()

	// Assign struct values
	e.UID = fmt.Sprintf("%s@example.com", FormatDateTime(currentTimestamp))
	e.DTSTAMP = &currentTimestamp

	return e
}

func (e *Event) isReady() bool {
	if e.UID == "" {
		return false
	}
	if e.DTSTAMP == nil {
		return false
	}
	if e.DTSTART == nil {
		return false
	}
	if e.DTEND != nil && e.DURATION != "" {
		return false
	}
	return true
}

func (e *Event) generateEventProp() string {
	// Validate first
	status := e.isReady()
	if !status {
		log.Fatal("Event is not ready!")
	}

	// Create object
	var str strings.Builder

	// Write headers
	str.WriteString("BEGIN:VEVENT\r\n")

	// Write required params
	str.WriteString("UID:" + e.UID + "\r\n")
	str.WriteString("DTSTAMP:" + FormatDateTime(*e.DTSTAMP) + "\r\n")
	str.WriteString("DTSTART:" + FormatDateTime(*e.DTSTART) + "\r\n")

	// Write optional params
	str.WriteString("SUMMARY:" + e.SUMMARY + "\r\n")

	// Write footers
	str.WriteString("END:VEVENT\r\n")

	return str.String()
}
