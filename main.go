package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	fmt.Println("Hello World Eater Galactus")
	// Connect to the server
	conn, err := net.Dial("tcp", "localhost:3306")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Port 3306 is open.\n\n")

	// if err = conn.SetReadDeadline(time.Now().Add(timeout)); err != nil {
	// 	return
	// }

	lenBuf := make([]byte, 4)
	_, err = conn.Read(lenBuf)
	if err != nil {
		return
	}

	lenData, err := msgLength(lenBuf)
	if err != nil {
		return
	}

	buf := make([]byte, lenData)
	_, err = conn.Read(buf)
	if err == io.EOF {
		return
	}
	if err != nil {
		return
	}

	// result, err := parseHandshake(buf)
	// if err != nil {
	// 	return
	// }

	// NOTE: We have byte array and need to pull Handshake protocol from it
	currlen := len(buf)
	if currlen == 0 {
		return
	}

	// Pull out the handshake version
	handshakeVersion := int(buf[0])

	if handshakeVersion == 10 {
		//fmt.Println("Handshake Version TEN")

	} else if handshakeVersion == 9 {
		fmt.Println("Handshake Version NINE")
	} else {
		fmt.Println("Handshake Version UNSUPPORTED")
	}
	if err := conn.Close(); err != nil {
		return
	}

	return

	// Close the connection
	conn.Close()
}

func msgLength(b []byte) (int32, error) {
	buf := bytes.NewReader(b)
	var result int32
	err := binary.Read(buf, binary.LittleEndian, &result)

	return result, err
}
