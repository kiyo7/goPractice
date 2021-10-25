//一番初めに読み込まれる
package main

import (
	"fmt"
)

// "math/rand"
// "time"

// type User struct {
//     gender string
//     age int
// }

// type Student struct {
//     name string
// }

// func (s Student) avg(math, english float64) (avgResult float64) {
//   avgResult = (math + english) / 2
//    return
// }

type Student struct {
    name string
}

func (s Student) calAvg(data []float64) (avgResult float64) {
    sum := 0.0
    for i := 0; i < len(data); i++ {
        sum += data[i]
    }
    avgResult = sum / float64(len(data))
    return
}

func (s Student) judge(avg float64) (judgeResult string) {
    if avg >= 60 {
        judgeResult = "passed"
    } else {
        judgeResult = "failed"
    }
    return
}
 
type User struct {
    name string
}
   
 
func (s User) cal(height, weight float64) ( bmi float64) {
    bmi = weight / height / height * 10000
    return
}

func main() {

    a001 := Student{"sato"}
    data := []float64{40,90,50,60,70}
    var avg float64 = a001.calAvg(data)
    result := a001.judge(avg)
    fmt.Println(avg)
    fmt.Println(a001.name + " " + result)

    // s001 := User{name: "kiyo"}
    // fmt.Println(s001.name,s001.cal(171,53))
    
    // a001 := Student{"sato"}
    // fmt.Println(a001.avg(80, 70))
    // var s User 

    // s.gender = "male"
    // s.age = 20

    // s := User{gender: "male", age:20}

    // fmt.Println(s)

    // rand.Seed((time.Now().UnixNano()))
    // answer := rand.Intn(10) + 1
    // count := 0

    // for {
    //     var guess int
    //     fmt.Print("あなたの予想は？ ")
    //     fmt.Scanf("%v", &guess)
    //     count++
    
    //     if answer == guess {
    //         fmt.Printf("正解! %v回で正解できたよ\n", count)
    //         break
    //     } else if answer > guess {
    //         fmt.Println("小さいよ")
    //     } else {
    //         fmt.Println("大きいよ")
    //     }
    // }
    // fmt.Printf("答えは%vです\n",answer)
}
