package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing() {
    for j := 0; j < 1000000
    {
        i += 1
    }
}

func decrementing() {
    for j := 0; j < 1000000
    {
        i -= 1
    }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    time.Sleep(100*time.Millisecond)
    Println("The magic number is:", i)
}
