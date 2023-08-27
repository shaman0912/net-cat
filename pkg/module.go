package pkg

import (
	"net"
	"sync"
)

// основные переменные
var (
	mutex       sync.Mutex
	clients     = make(map[string]Client)
	historytext = []string{}
	leaving     = make(chan Message)
	messages    = make(chan Message)
	join        = make(chan Message)
)

// cтруктура клиента
type Client struct {
	// имя пользователя
	name string
	// айпи адрес пользователя
	addr string
	// интерфейс конекшна
	conn net.Conn
}

// cтруктура сообщения
type Message struct {
	// текст сообщения
	text string
	// айпи адрес пользователя
	address string
	// имя пользователя
	name string
	// время
	time string
	// история сообщении
	history []string
}
