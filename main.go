package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	ics "github.com/arran4/golang-ical"
)

func init() {
	functions.HTTP("HandleRequest", HandleCalendar)
}

func main() {
	// Parse command line flags
	port := flag.Int("port", 8080, "Port to serve on")
	flag.Parse()

	// Create a handler for the calendar endpoint
	http.HandleFunc("/", HandleCalendar)

	// Start the server
	serverAddr := fmt.Sprintf(":%d", *port)
	log.Printf("Starting iCal server on http://localhost%s", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}

func HandleCalendar(w http.ResponseWriter, r *http.Request) {
	// Create a new calendar
	cal := ics.NewCalendar()
	cal.SetMethod(ics.MethodPublish)
	cal.SetProductId("-//go-mergecal//NONSGML v1.0//EN")
	cal.SetName("Merged Calendar")
	cal.SetDescription("Calendar served by go-mergecal")

	// Add a sample event (in a real application, you would fetch and merge events from multiple sources)
	event := cal.AddEvent(fmt.Sprintf("event-%d@go-mergecal", time.Now().Unix()))
	event.SetCreatedTime(time.Now())
	event.SetDtStampTime(time.Now())
	event.SetModifiedAt(time.Now())
	event.SetStartAt(time.Now().Add(24 * time.Hour))
	event.SetEndAt(time.Now().Add(25 * time.Hour))
	event.SetSummary("Sample Event")
	event.SetDescription("This is a sample event from go-mergecal")
	event.SetLocation("Virtual")

	// Set appropriate content type and write the calendar
	w.Header().Set("Content-Type", "text/calendar; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename=calendar.ics")
	w.WriteHeader(http.StatusOK)
	cal.SerializeTo(w)

	log.Printf("Served calendar to %s", r.RemoteAddr)
}
