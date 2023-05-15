package main

import "fmt"

type SettlementId string

var cityId SettlementId = "228edf0c-0c0d-4db6-bf0e-d508e68270a3"
var strCityId string = "228edf0c-0c0d-4db6-bf0e-d508e68270a3"

func main() {
	fmt.Println(cityId == SettlementId(strCityId))
	fmt.Println("1 + 1 =", 1.01+1.01)
	test()
}

func test() {
	fmt.Printf("%d число", 3)
}
