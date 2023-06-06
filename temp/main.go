package main
import "fmt"

type Student struct {
	name string
	rollNo int
	marks []int
	grades map[string]int
}

func main() {
	var s Student
	fmt.Printf("%v",s)
}