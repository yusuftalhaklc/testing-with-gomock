package models

import "github.com/yusuftalhaklc/testing-with-gomock/handlers"

type User struct {
	Handler handlers.Handler
}

func (u *User) Create() error {
	return u.Handler.CreateUser("User created")
}

//go:generate mockgen -destination=../mocks/mock_handler.go -package=mocks github.com/yusuftalhaklc/testing-with-gomock/handlers Handler
