package pkg

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// функция возвращения нового сообщения от пользователя к к серверу
func NewMessage(msg string, conn net.Conn, cl Client, time string) Message {
	return Message{
		text:    msg,
		address: cl.addr,
		name:    cl.name,
		time:    time,
		history: historytext,
	}
}

// приветственное окно для нового подключенного пользователя
func Welcome(conn net.Conn) {
	file, err := os.ReadFile("image.txt")
	if err != nil {
		fmt.Printf("couldn't read this file")
	}
	strWelcome := (string(file))
	conn.Write([]byte("Welcome to TCP-Chat!\n" + strWelcome + "\n"))
}

// получение имени каждого подключенного пользвоателя
func GetName(conn net.Conn) string {
	conn.Write([]byte("[Enter your name]:"))
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	}

	temp := strings.TrimSpace(string(data))
	// проверка на наличие пустоты
	if temp == "" || len(temp) == 0 {
		fmt.Fprintln(conn, "Incorrect input")
		return GetName(conn)
	}
	// проверка на символы ASCII
	for _, w := range temp {
		if w < 32 || w > 127 {
			fmt.Fprintln(conn, "Incorrect input")
			return GetName(conn)
		}
	}
	// проверка имени в существующей мапе
	for i := range clients {
		if i == temp {
			fmt.Fprintf(conn, "User already exist\n")
			return GetName(conn)
		}
	}
	return temp
}

// функция правильности ввода каждого подключенного к серверу пользователя
func IsCorrect(s string, conn net.Conn, time string, username string) bool {
	for _, w := range s {
		if w < 32 || w > 127 {
			return false
		}
	}
	return true
}
