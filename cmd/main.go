package main

import (
	"fmt"

	"github.com/Freyzan2006/more-struct-go/internal/set"
)

func main() {
	mySet := set.NewSet([]int{1, 2, 3})
	mySet.Add(1)
	mySet.Add(2)
	mySet.Add(3)
	mySet.Add(2)
	fmt.Println(mySet.Has(1));
	fmt.Println(mySet.Len());
	mySet.Delete(1);
	fmt.Println(mySet.Len());
	mySet.Clear();
	fmt.Println(mySet.Len());
}