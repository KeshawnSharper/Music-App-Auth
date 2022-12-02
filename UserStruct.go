package main

type UserStruct struct {
	Username string `json:"Username"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	ID       int    `json:"ID"`
}

type AuthenticationStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenStruct struct {
	Role        string `json:"role"`
	Email       string `json:"email"`
	TokenString string `json:"token"`
}

type UserRequestStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"Username"`
}

var userList [0]UserStruct
