package main

import (
	"fmt"
	"time"
)

func main() {
	// Go channel은 goroutines끼리 데이터를 주고 받는 통로라 볼 수 있음
	// 상대편이 준비될 때까지 channel에서 대기함으로써 별도의 lock를 걸지 않고 데이터를 동기화 하는데 사용
	// make함수로 channel을 생성
	c := make(chan string)
	people := [3]string{"nick", "judy", "nana"}
	// 3개의 goroutines
	for _, person := range people {
		go isPerson(person, c)
	}
	// 함수가 실행되고 채널에 전달된 message를 출력함
	// goroutines의 경우 순차적 실행이 아닌 동시 실행이기 때문에 people의 순서와 message의 순서는 다를 수 있음
	for i := 0; i < len(people); i++ {
		// '<-channel' 를 통해 channel의 값을 받아옴
		fmt.Println(<-c)
	}
}

func isPerson(person string, c chan string) {
	time.Sleep(time.Second * 3)
	// 'channel <- data' 를 통해 channel에 값을 보냄
	c <- person + " is person"
}
