package main

import (
     "bufio"
     "fmt"
     "os"
     "strings"
)

func main() {

    reader := bufio.NewReader(os.Stdin)
     X, _ := reader.ReadString('\n')
          // convert CRLF to LF
     X = strings.Replace(X, "\n", "", -1)

    fmt.Println("Hello, World.")
    fmt.Println(X)
}
