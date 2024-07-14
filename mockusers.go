package main

/**
just creating a basic struct and map of users for the purpose of this sample
*/

type User struct {
	Id        string `json:"Id"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}

var UserList = map[string]User{
	"1": {Id: "1", FirstName: "George", LastName: "Jetson", Email: "gjetson@spacely.com"},
	"2": {Id: "2", FirstName: "Fred", LastName: "Flintstone", Email: "flintstone@slaterockandgravel.com"},
	"3": {Id: "3", FirstName: "Scoobert", LastName: "Doo", Email: "scooby@mysteryinc.com"},
	"4": {Id: "4", FirstName: "Yogi", LastName: "Bear", Email: "yogi@jellystonepark.net"},
}
