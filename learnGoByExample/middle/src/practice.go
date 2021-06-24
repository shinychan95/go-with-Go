package main

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

	/* 6. Closure */
	next := nextValue()

	println(next()) // 1
	println(next()) // 2
	println(next()) // 3

	anotherNext := nextValue()
	println(anotherNext()) // 1 다시 시작
	println(anotherNext()) // 2

	/* 7. Go Collection - Array */

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
