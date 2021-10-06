package main

import (
        "bufio"
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
        "os"
        "strings"
        "time"
)

func main() {
        fmt.Println("Hello World, Hello Walls.io Recent posts")
        fmt.Println("Env URL value:", os.Getenv("URL"))

        for {

                resp, err := http.Get(os.Getenv("URL"))
                if err != nil {
                        fmt.Print(err)
                        os.Exit(3)
                }
                defer resp.Body.Close()

                bodyBytes, err := ioutil.ReadAll(resp.Body)
                bodyString := string(bodyBytes)

                if err != nil {
                        log.Fatal(err)
                }

                fmt.Print(resp.StatusCode, " ")
                scanner := bufio.NewScanner(strings.NewReader(bodyString))
                for scanner.Scan() {
                        fmt.Print(scanner.Text())
                }
                fmt.Print("\n")
                time.Sleep(2 * time.Second)
        }
}

