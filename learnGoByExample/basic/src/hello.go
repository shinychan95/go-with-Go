package main

import "fmt"

func main() {
	println("Hello World")

	// 1. Variable
	var i, j, k int = 1, 2, 3 // mismatch 주의
	var tmp int               // 기본값은 0
	v := 999                  // 타입 추론 및 func 내 var 생략

	println(i)
	println(j)
	println(k)
	println(tmp)
	println(v)

	// 2. Constant
	const (
		Visa   = "Visa"
		Master = "MasterCard"
		Amex   = "American Express"
	)

	const (
		Apple  = iota // 0
		Grape         // 1
		Orange        // 2
	)

	println(Visa)
	println(Apple)

	// 3. Data Type
	/*
		bool
		string (immutable)
		int, int8, int16, int 64
		uint uint8 uint16 uint32 uint64 uintptr
		float32 float64 complex64 complex128
		byte: uint8과 동일하며 바이트 코드에 사용
		rune: int32과 동일하며 유니코드 코드포인트에 사용한다
	*/

	// 4. String
	// Raw String Literal. 복수라인에 주로 사용된다.
	rawLiteral := `아리랑\n
아리랑\n
아라리요`

	// Interpreted String Literal
	interLiteral := "아리랑아리랑\n아리리요"
	// 아래와 같이 +를 사용하여 두 라인에 걸쳐 사용할 수도 있다.
	// interLiteral := "아리랑아리랑\n" +
	//                 "아리리요"

	fmt.Println(rawLiteral)
	fmt.Println()
	fmt.Println(interLiteral)

	// 5. Type Conversion
	var x5 int = 100
	var y5 uint = uint(x5)
	var z5 float32 = float32(y5)
	println(z5, y5)

	str5 := "ABC"
	bytes := []byte(str5)
	strCovt := string(bytes)
	println(bytes, strCovt)

	// 6. Operator
	// Pointer op
	var k6 int = 10
	var ptr = &k6 //k의 주소를 할당
	println(ptr)
	println(*ptr) //p가 가리키는 주소에 있는 실제 내용을 출력

	// 7. if statement
	i7 := 2
	if i7 == 1 { //같은 라인
		println("One")
	} else if i7 == 2 { //같은 라인
		println("Two")
	} else { //같은 라인
		println("Other")
	}

	// if 문에서 조건식 사용하기 전에 Optional Statement가 가능하다.
	max := 100
	if v7 := i7 * 2; v7 < max {
		println(v7)
	}

	// 8. Switch statement
	var name string
	var category = 2

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4: // 콤마를 통해 복수개의 case 값 적용 가능
		name = "Blog"
	default:
		name = "Other"
	}
	println(name)

	// Expression을 사용한 경우
	// switch x := category << 2; x - 1 {
	//    ...
	// }

	// Go만의 특별한 용법들
	// 1. switch 뒤에 expression이 없을 수 있음
	// 2. case문에 expression을 쓸 수 있음
	// 3. No default fall through (break 없어도 다른 case 문으로 가지 않는다)
	// 4. Type switch

	// 3번의 경우, fallthrough를 명시해주어 다음 case로 넘어갈 수 있고,
	// 4번의 경우, 아래와 같이 활용한다.
	explain := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		case string:
			fmt.Println("I'm a string")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	explain(name)
}
