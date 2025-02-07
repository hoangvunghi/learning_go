package main

import (
	"fmt"
	"be-ep/internal/logic"
)

type Student struct {
	Name string
	Age int
	// Name string là biến Name có kiểu dữ liệu là string
	// Age int là biến Age có kiểu dữ liệu là int
}

func (s Student) GetAge() int {
	return s.Age
}

func (s *Student) SetAge(age int) {
	s.Age = age
}



var (
	count = 0
	Count = 0
	// viết hoa thì là public
	// viết thường thì là private

	color = [3]string{"red", "green", "blue"}
	// [3] là số phần tử của mảng
	// string là kiểu dữ liệu của mảng
	// {"red", "green", "blue"} là giá trị của mảng
)

const (
	PI = 3.14
	number = 100
	// const là hằng số
	// const <tên hằng số> <kiểu dữ liệu> = <giá trị>
)

func counting() int {
	count := 0
	const PI = 3.14
	// count := 0 là khai báo biến count với giá trị ban đầu là 0
	// int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64, uintptr
	// float32, float64
	// string
	// byte // uint8
	// rune // int32
	// bool
	// complex64, complex128
	count++
	return count
}

func main() {
	// colors := make([]string, 0, 5)
	// make là hàm tạo mảng
	// make([]<kiểu dữ liệu>, <số phần tử>, <số phần tử tối đa>)
	// []string là kiểu dữ liệu của mảng
	colors1 := []string{"red", "green", "blue"}
	// []string là kiểu dữ liệu của mảng
	colors1 = append(colors1, "yellow")
	// append là hàm thêm phần tử vào mảng
	// append(<tên mảng>, <giá trị cần thêm>)
	s := Student{Name: "John", Age: 25}
	// Student{Name: "John", Age: 25} là khởi tạo biến s có kiểu dữ liệu là Student
	// Name: "John" là khởi tạo biến Name có giá trị là "John"
	// Age: 25 là khởi tạo biến Age có giá trị là 25
	fmt.Println(s.GetAge())
	fmt.Println(logic.SayHello("world"))
	fmt.Println(counting())
}