package main

import (
	"fmt"

	linq "github.com/ahmetb/go-linq"
	"github.com/marvin5064/linq-json-filter/protobuf"
)

func main() {
	fmt.Println("Testing filtering by linq")
	items := initPayloads()
	fmt.Println("initialized items:", items)
	fmt.Println("------breaking line------")
	items = filterByConditions(items)
	fmt.Println("filtered items:", items)
}

func initPayloads() []*protobuf.SmallerP3Define {
	var result []*protobuf.SmallerP3Define
	for i := 1; i < 10; i++ {
		result = append(result, &protobuf.SmallerP3Define{
			Provider:        uint32(i),
			SpecialBetValue: fmt.Sprintf("Value: %v", i),
		})
	}
	return result
}

func filterByConditions(
	items []*protobuf.SmallerP3Define,
) []*protobuf.SmallerP3Define {

	linq.From(items).WhereT(func(item *protobuf.SmallerP3Define) bool {
		// filter schema comes here
		return item.GetProvider() < 6
	}).ToSlice(&items)

	return items
}
