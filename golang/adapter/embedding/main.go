package main

import "fmt"

type Handler struct {
	message string
}

func (h *Handler) Say() {
	fmt.Printf("%s", h.message)
}

type SayHandler struct {
	*Handler // embedding
}

func NewSayHandler(message string) *SayHandler {
	return &SayHandler{
		Handler: &Handler{message: message},
	}
}

func (s *SayHandler) HandleSay() {
	s.Say()
}

func main() {
	s := NewSayHandler("Hi")
	s.Say()
}
