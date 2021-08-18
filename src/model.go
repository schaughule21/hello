package main

var users = map[string]User{
	"1": {
		ID:        1,
		FirstName: "Jennifer",
		LastName:  "Watson",
	},
	"2": {
		ID:        2,
		FirstName: "abc",
		LastName:  "xyz",
	},
	"3": {
		ID:        3,
		FirstName: "pqr",
		LastName:  "lmn",
	},
	"4": {
		ID:        4,
		FirstName: "Siddhesh",
		LastName:  "Chaughule",
	},
}

type User struct {
	ID        uint
	FirstName string `form:"fname"`
	LastName  string `form:"lname"`
}
