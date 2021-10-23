//一番初めに読み込まれる
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    rand.Seed((time.Now().UnixNano()))
    answer := rand.Intn(10) + 1
    count := 0

    for {
        var guess int
        fmt.Print("あなたの予想は？ ")
        fmt.Scanf("%v", &guess)
        count++
    
        if answer == guess {
            fmt.Printf("正解! %v回で正解できたよ\n", count)
            break
        } else if answer > guess {
            fmt.Println("小さいよ")
        } else {
            fmt.Println("大きいよ")
        }
    }
    fmt.Printf("答えは%vです\n",answer)
}
