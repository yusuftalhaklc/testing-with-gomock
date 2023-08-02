package handlers

type Handler interface {
	CreateUser(string) error
}
