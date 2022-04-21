package main

import (
	"fmt"
	"sort"
	"strings"
)

const DelimeterChar = "|"

const LocationCollection = "CL2|CL11|CL|CL10|CL4|VCL|ML|CL5|VHOL_MSK"

func main()  {
	fmt.Println(positionLocation(LocationCollection, DelimeterChar, "CL10"))
	fmt.Println(nameLocation(LocationCollection,DelimeterChar, 3))
	fmt.Println(abcLocation(LocationCollection, DelimeterChar))
}

func positionLocation (locationCollection string, delimeterChar string, locationCode string) int {
	locs := strings.Split(locationCollection, delimeterChar)
	for nam, loc := range locs {
		if loc == locationCode {
			return nam + 1
		}
	}
	return 0
}

func nameLocation(locationCollection string, delimeterChar string, locationPosition int) string {
	locs := strings.Split(locationCollection, delimeterChar)
	for nam, loc := range locs {
		if nam + 1 == locationPosition {
			return loc
		}
	}
	return "Location not found."
}

func abcLocation(locationCollection string, delimeterChar string) []string {
	locs := strings.Split(locationCollection, delimeterChar)
	sort.Strings(locs)
	return locs
}