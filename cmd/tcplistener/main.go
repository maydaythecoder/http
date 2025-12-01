package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func getLinesReader(f io.ReadCloser) <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)
		defer f.Close()

		buf := make([]byte, 8)
		var line []byte

		for {
			n, err := f.Read(buf)
			if n > 0 {
				chunk := buf[:n]

				for {
					i := bytes.IndexByte(chunk, '\n')
					if i == -1 {
						line = append(line, chunk...)
						break
					}

					line = append(line, chunk[:i]...)
					out <- string(line)
					line = line[:0]

					chunk = chunk[i+1:]
					if len(chunk) == 0 {
						break
					}
				}
			}

			if err == io.EOF {
				if len(line) > 0 {
					out <- string(line)
				}
				return
			}
			if err != nil {
				log.Println("read error:", err)
				return
			}
		}
	}()

	return out
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}
		for line := range getLinesReader(conn) {
			fmt.Println("read:", line)
		}
	}
}
