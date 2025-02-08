package main

func addOne(v *int){
	*v++
}
// thread
// process

func main() {
	value := 0
	go addOne(&value)
	go addOne(&value)
	// thêm go ở đây để chạy 2 goroutine cùng lúc
	// nếu không thêm go thì chỉ chạy 1 goroutine
	// và kết quả sẽ là 1
	// nếu thêm go thì kết quả sẽ là 2
}