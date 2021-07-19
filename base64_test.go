package main

import (
	"encoding/base64"
	"log"
	"testing"
)

func TestBase64(t *testing.T) {
	data := `AAAAAA0AAAAAAAAAAQAAAAAJIDkBAAAAAgAAAAAAAAACAAAAAwAAAAAAAAADAAAABAAAAAAAAAAEAAAABQAAAAAAAAAFAAAABgAAAAB/AAAGAAAABwAAAAAhvzkHAAAACAAAAAB/AAAIAAAACQAAAAAAAAAJAAAACgAAAAB/AAAKAAAACwAAAABQvzkLAAAADAAAAAAAAAAMAAAADAAAAAEAAAAMAAAA5AcgOQAAAAABAAAAAAAAALoAAAAAAAAAAgAAAAAAAACMNxgCAAAAAAMAAAAAAAAA5TK/OQAAAAAEAAAAAAAAAI0xvzkAAAAABQAAAAAAAACBAQAAAAAAAAYAAAAAAAAA5QAAAAAAAAAHAAAAAAAAAIX///8AAAAACAAAAAAAAACtIE1pAAAAAAkAAAAAAAAA5QpUIAAAAAAKAAAAAAAAALogMiAAAAAACwAAAAAAAACmIDQgAQAAAAwAAAAAAAAAAgAAAA4AAAAA5LqM5Y2B5YWt5bqmAA==`
	dec, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		log.Printf("decode base64 that user input rules err:%v\n", err)
		return
	}

	t.Log(dec)
}
