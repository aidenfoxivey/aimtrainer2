package main

import (
    "fmt"
    "crypto/sha256"
)

func main() {
    s := []byte("hi, welcome from go")
    sum := sha256.Sum256(s)

    fmt.Printf("Hash sum is %x\n", sum)
}
