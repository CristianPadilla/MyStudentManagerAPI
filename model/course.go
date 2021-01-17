package model

import "time"

var (
	days = []string{"Sunday", "Monday", "Tuesday", "s"}
)

type Course struct {
	ID      uint
	Name    string
	Code    string
	Subject string
	Group   string
	time    time.Time
}
