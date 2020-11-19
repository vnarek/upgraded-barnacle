package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
)

const NumberOfLogs = 1_000_000

func main() {
	fmt.Println(NumberOfLogs, " logs should be printed")
	var b []byte
	for i := 0; i < NumberOfLogs; i++ {
		binary.LittleEndian.PutUint32(b, uint32(i))
		fmt.Println("Hash for the number ", i, " is ", sha1.Sum(b))
	}
}
