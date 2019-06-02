package main

import(
	"fmt"
	"net"
	"bufio"
)

func main() {
	li, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		{
			if err != nil {
				fmt.Println(err)
			}
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "i heard you say: %s\n", ln)
	}
	defer conn.Close()
}

