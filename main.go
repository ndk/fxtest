package main

import (
	"fxtest/client"
	"fxtest/service"
)

func main() {
	s := service.NewService()
	c := client.NewClient(s)
	for i := 0; i < 3; i++ {
		c.DoSomething()
	}
}
