package main

import (
	"fmt"       //출력을 위한 패키지
	"math/rand" // 랜덤함수 패키지
	"time"
)

// 난수 추출된 수의 소수 판정 프로그램 v0.3
// 소수 : 1과 자기자신외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
func main() {
	// seed 설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	isPrime := true
	number := rand.Intn(150) + 2
	fmt.Println("임의로 추출된 수 : ", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			//count = count + 1
		}
	}

	if isPrime == true {
		fmt.Println(number, "는(은) 소수입니다!")
	} else {
		fmt.Println(number, "는(은) 소수가 아닙니다~!")
	}

}
