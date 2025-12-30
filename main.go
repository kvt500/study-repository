package main

import (
	"fmt"
	"study/feature1"
	"study/feature1/feture_postgres/simple_connection"
	"study/feature2"
)

func main() {
	fmt.Println("Hello, Git!")
	feature1.Feature1()
	feature2.Feature2()
	simple_connection.CheckConnection()
}
