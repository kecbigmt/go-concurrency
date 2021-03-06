package main

import (
  "fmt"
  "sync"
  "time"
)

const runtime = time.Second * 1

func main() {
  var wg sync.WaitGroup
  var sharedLock sync.Mutex
  
  greedyWorker := func() {
    defer wg.Done()
    
    var count int
    for begin := time.Now(); time.Since(begin) <= runtime; {
      sharedLock.Lock()
      time.Sleep(time.Nanosecond * 3)
      sharedLock.Unlock()
      count++
    }

    fmt.Printf("Greedy worker was able to execute %v work loops\n", count)
  }

  politeWorker := func() {
    defer wg.Done()

    var count int
    for begin := time.Now(); time.Since(begin) <= runtime; {
      sharedLock.Lock()
      time.Sleep(time.Nanosecond * 1)
      sharedLock.Unlock()

      sharedLock.Lock()
      time.Sleep(time.Nanosecond * 1)
      sharedLock.Unlock()

      sharedLock.Lock()
      time.Sleep(time.Nanosecond * 1)
      sharedLock.Unlock()

      count++
    }

    fmt.Printf("Polite worker was able to execute %v work loops\n", count)
  }

  wg.Add(2)
  go greedyWorker()
  go politeWorker()

  wg.Wait()
}
