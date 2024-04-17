package client

import "fmt"

type service interface {
	GetSomething() string
}

type client struct {
	s service
}

func NewClient(s service) *client {
	return &client{
		s: s,
	}
}

func (c *client) DoSomething() {
	fmt.Println(c.s.GetSomething())
}
