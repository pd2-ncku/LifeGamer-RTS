package main

import (
    "fmt"
    "log"
    "service"
    "protocal"
)

func main() {
    log.Println("Starting LifeGamer-RTS game server")

    listenDone := make(chan struct{})

    map_client := protocal.NewClient("map")
    map_client2 := protocal.NewClient("trade")

    if !map_client.Put("trade", "Generate") {
        fmt.Println("Put fail")
    } else {
        fmt.Printf("%s\n", map_client2.Get())
    }

    go func() {
        service.Start()
        close(listenDone)
    }()

    <-listenDone
}
