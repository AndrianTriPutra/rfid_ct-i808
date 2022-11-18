package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed connect to socket:%v", err)
	}
	defer connection.Close()

	tag := ""
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		<-ticker.C
		buffer := make([]byte, 1024)
		mLen, err := connection.Read(buffer)
		if err != nil {
			log.Printf("Error reading:%s", err.Error())
		}

		str := fmt.Sprintf("%x", buffer[:mLen])
		if strings.Contains(str, "0700") && len(str) > 16 {
			id := str[strings.Index(str, "0700") : strings.Index(str, "0700")+16]
			log.Printf("id : %s", id)
			if id != tag {
				tag = id
				log.Printf("tag: %s", tag)
			}
		}
	}

}
