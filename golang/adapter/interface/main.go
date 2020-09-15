package main

import "fmt"

type Handler interface {
	Say()
}

func (s *SayHandler) Say() {
	fmt.Printf("%s", s.message)
}

type SayHandler struct {
	message string
	Handler Handler
}

func NewSayHandler(message string) *SayHandler {
	sayHandler := &SayHandler{
		message: message,
	}
	sayHandler.Handler = sayHandler
	return sayHandler
}

func (s *SayHandler) HandleSay() {
	s.Handler.Say()
}

func main() {
	s := NewSayHandler("Hi")
	s.Say()
}
