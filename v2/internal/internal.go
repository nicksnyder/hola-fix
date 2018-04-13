package internal

type Message struct {
	Msg string
}

func (m *Message) String() string {
	return m.Msg
}
