package event_default

import "github.com/chefsgo/event"

func Driver() event.Driver {
	return &defaultDriver{}
}

func init() {
	event.Register("default", Driver())
}
