package main

import (
	"fileOrganisation/counters"
	"fmt"
)

func main()  {

	a := counters.admin{Rights: 12}
	a.Name = "Emmanuel"
	a.Email = "emma@gmail.com"

	fmt.Printf("%v\n",a)
}