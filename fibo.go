package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func fibo(n int64) int64 {
    if n <= 1 {
        return n
    } else {
        return fibo(n-1) + fibo(n-2)
    }
}

func main() {
    start := time.Now()

    var nums int64 = 41
    arr := make([]int64, nums)

    // single threaded
    // for i := int64(0); i < nums; i++ {
    //     fmt.Print(i, " ")
    //     arr[i] = fibo(i)
    // }

    // multithreaded with channels
    // c := make(chan map[int64]int64, nums)
    // defer close(c)
    // for i := int64(0); i < nums; i++ {
    //     fmt.Print(i, " ")
    //     result := make(map[int64]int64, nums)
    //     go func ()  {
    //         result[i] = fibo(i)
    //         c <- result
    //     }()
    // }
    // for i := int64(0); i < nums; i++ {
    //     result, ok := <- c
    //     if !ok {
    //         fmt.Println("error reading from channel")
    //     }
    //     for k, v := range result {
    //         arr[k] = v
    //     }
    // }

    // multithreaded with waitgroups
    var wg sync.WaitGroup
    for i := int64(0); i < nums; i++ {
        fmt.Print(i, " ")
        wg.Add(1)
        go func ()  {
            defer wg.Done()
            arr[i] = fibo(i)
        }()
    }
    wg.Wait()


    fmt.Println()
    fmt.Printf("%#v\n", arr)

    timeElapsed := time.Since(start)
    log.Println("Total time: " + timeElapsed.String())
}