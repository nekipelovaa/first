package main

import "fmt"

type SettlementId string

var cityId SettlementId = "228edf0c"
var strCityId string = "228edf0c"

func main() {
	fmt.Println(cityId == SettlementId(strCityId))
	fmt.Println("1 + 1 =", 1.01+1.01)
	test()
	test2()
	test3()
}

func test() {
	fmt.Printf("%d число", 3)
}

func test3() {
	fmt.Println("test3")
}

func test2() {
	fmt.Println("Test2")
}
