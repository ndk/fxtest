package service

import "fmt"

type service struct {
	i int
}

func NewService() *service {
	return &service{}
}

func (s *service) GetSomething() string {
	s.i++
	return fmt.Sprintf("something %d", s.i)
}
