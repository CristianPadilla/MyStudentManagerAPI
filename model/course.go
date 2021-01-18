package model

import "time"

type Course struct {
	ID         uint
	Name       string
	Code       string
	Subject    string
	Group      string
	time       time.Time
	Day        string
	Professor  string
	Enrollment Enrollment
}
