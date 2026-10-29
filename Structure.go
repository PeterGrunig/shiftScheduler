package main

import (
	"time"
)


//Defines a user struct to hold the user's name, password and schedule
type User struct {
	Password string
	Role    string
	
}

type session struct {
	username string
	expiry time.Time
	role string
	Schedule Schedule
	StudentUsername string
	StudentSchedule Schedule
}

var sessions = map[string]session{} // Map to store session tokens and associated usernames



func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}


//Map that maps the username to the user struct, this will be used for authentication and to store the schedule for each user
//Student 1 has password "pass123" and role "student"
//Admin 1 has password "adminpass" and role "admin"
var users = map[string]User{
	"student1": {Password: "pass123", Role: "student"},
	"admin1":   {Password: "adminpass", Role: "admin"},
}
