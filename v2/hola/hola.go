package hola

import "github.com/nicksnyder/hola/v2/internal"

func Hola2() string {
	m := &Message{Msg: "Hola2"}
	return m.String()
}

type Message = internal.Message



