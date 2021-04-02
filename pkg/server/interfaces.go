package server

type Serveable interface {
	Stop() error
}
