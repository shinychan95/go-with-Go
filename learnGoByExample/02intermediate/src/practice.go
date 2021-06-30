package main

import (
	"fmt"
)

/* 익명 함수 관련 내용 추가 */
// 함수 원형 정의
type calculatorNum func(int, int) int
type calculatorStr func(string, string) string

func calNum(f calculatorNum, a int, b int) int {
	result := f(a, b)
	return result
}

func calStr(f calculatorStr, a string, b string) string {
	sentence := f(a, b)
	return sentence
}

func main() {
	/* 1. Function */
	msg := "Hello World"
	say(msg)

	/* 2. Pass By Reference */
	addExclamation(&msg)
	say(msg)

	/* 3. Variadic Function */
	says("Hello", "New", "World")

	/* 4. Function Return */
	total := sum(1, 7, 3, 5, 9)
	println(total)

	count, total := sums(1, 7, 3, 5, 9)
	println(count, total)

	/* 5. Anonymous Function */
	// 함수 전체를 변수에 할당하거나 다른 함수의 파라미터에 직접 정의된다.
	// '변수명(파라미터들)' 형태로 사용된다.
	// 일급 함수 -> 기본 타입과 동일하게 취급 -> 함수의 파라미터 전달 혹은 리턴값으로 사용된다.
	// type문을 사용한 함수 원형 정의

	multi := func(i int, j int) int {
		return i * j
	}
	duple := func(i string, j string) string {
		return i + j + i + j
	}

	r1 := calNum(multi, 10, 20)
	fmt.Println(r1)

	r2 := calStr(duple, "Hello", " Golang ")
	fmt.Println(r2)

	/* 6. Closure */
	next := nextValue()

	println(next()) // 1
	println(next()) // 2
	println(next()) // 3

	anotherNext := nextValue()
	println(anotherNext()) // 1 다시 시작
	println(anotherNext()) // 2

	/* 7. Go Collection - Array */
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	println(arr1[2])

	var arr2 = [3]int{1, 2, 3}
	var arr3 = [...]int{4, 5, 6}
	fmt.Println(arr2)
	fmt.Println(arr3[2])

	var a = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, //끝에 콤마 추가
	}
	println(a[1][2])

	/* 7. Go Collection - Slice */
	// Go 배열은 크기를 동적으로 증가시키거나 부분 배열을 발췌하는 등의 기능이 없다.
	// Go Slice는 내부적으로 배열에 기초하여 만들어졌지만,
	// 이런 제약점들을 넘어 편리하고 유용한 기능을 제공한다.

	var s1 []int // 슬라이스 변수 선언 (크기를 지정하지 않는다)
	s1 = []int{1, 2, 3}
	s1[1] = 10
	fmt.Println(s1)

	// make를 활용하여 슬라이스 생성
	s2 := make([]int, 5, 10) // Slice type, Length, Capacity
	println(len(s2), cap(s2))

	// 부분 슬라이스
	s3 := []int{0, 1, 2, 3, 4, 5}
	s3 = s3[2:5]    // 2, 3, 4
	s3 = s3[1:]     // 3, 4
	fmt.Println(s3) // 3, 4 출력

	// 슬라이스 병합과 복사
	s4 := []int{0, 1}

	// 하나 확장
	s4 = append(s4, 2) // 0, 1, 2
	// 복수 요소들 확장
	s4 = append(s4, 3, 4, 5) // 0,1,2,3,4,5

	fmt.Println(s4)

	/* 8. Go Collection - Map */
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "Facebook",
	}

	println(tickers["GOOG"])

	// 추가 혹은 갱신
	tickers["NV"] = "NAVER"
	tickers["FB"] = "FaceBook"

	fmt.Println(tickers)

	noData := tickers["KKO"] // 값이 없으면 nil 혹은 zero 리턴
	println(noData)

	// 삭제
	delete(tickers, "FB")
	fmt.Println(tickers)

	// map 키 체크
	val, exists := tickers["MSFTS"]
	if !exists {
		println("No MSFT ticker")
	} else {
		fmt.Printf("Value is %s", val)
	}

	// for range 문을 사용하여 모든 맵 요소 출력
	// Map은 unordered 이므로 순서는 무작위
	for key, val := range tickers {
		fmt.Println(key, val)
	}
}

// Pass By Value
func say(msg string) {
	println(msg)
}

// Pass By Reference
func addExclamation(msg *string) {
	*msg += "!"
}

// Variadic Function
func says(msg ...string) { // 메소드 오버로딩이 지원되지 않는다?
	for _, s := range msg {
		println(s)
	}
}

// Function Return
func sum(nums ...int) int { // 파라미터 괄호 뒷 부분에 적어준다.
	s := 0
	for _, n := range nums {
		s += n
	}
	return s
}

// Named Return Parameter
func sums(nums ...int) (count int, total int) { // 명시적으로 리턴될 파라미터를 적는다.
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return // return만 적어준다.
}

// Closure
func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
