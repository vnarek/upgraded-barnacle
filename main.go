package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const NumberOfLogs = 2000

func main() {
	fmt.Println(NumberOfLogs-1, " logs should be printed")
	for i := 0; i < NumberOfLogs; i++ {
		fmt.Println("Hash for the number ", i, " is ", i%6151)
	}
	if len(os.Args) > 1 {
		if os.Args[1] == "build" {
			time.Sleep(10*time.Second)
			panic("some horrible stuff happened")
			return
		}
	}
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("OK"))
	}))

	http.ListenAndServe(":8080", nil)
}
