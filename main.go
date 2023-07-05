package main

import "fmt"

func main() {
	fmt.Println("Pointer")
	a := 2
	fmt.Println("Address :", &a) //0xc0000200c0
	fmt.Println("Value :", a)    //2
	b := &a
	fmt.Println(*b) //2
	fmt.Println(b)  //0xc0000200c0

	var x *int
	y := 2
	x = &y
	fmt.Println(x)  //0xc00010e018
	fmt.Println(*x) //2
	x = new(int)
	*x = 10
	fmt.Println(*x) //10

	yy := 20
	var xx *int = &yy
	fmt.Println("address of yy : ", &yy) //0xc0000a6048
	fmt.Println("value of yy : ", yy)    //20

	fmt.Println("Address of xx : ", xx) //0xc0000a6048
	fmt.Println("Value of xx : ", *xx)  //20
	*xx = *xx + 1
	fmt.Println("value of pointer xx : ", *xx) //21
	fmt.Println("value of yy : ", yy)          //21

	aa := 2
	bb := &aa        //bb = 2
	fmt.Println(aa)  // -> 2
	fmt.Println(*bb) // -> 2

	*bb = 3
	fmt.Println(aa)  // -> 3
	fmt.Println(*bb) // -> 3

	aaa := 2    // value = 2 ,address of aaa =0xc00
	bbb := &aaa // value bbb = address of aaa =0xc00 ,address of bbb = 0xc11
	ccc := &bbb // value ccc = address of bbb 0xc11

	fmt.Println("---------------")
	fmt.Printf("aaa : %d\n", aaa)
	fmt.Printf("*&aaa : %d\n", *&aaa)
	fmt.Printf("*bbb : %d\n", *bbb)
	fmt.Printf("**ccc : %d\n", **ccc)

	fmt.Println("---------------")

	fmt.Printf("&bbb : %d\n", &bbb)
	fmt.Printf("ccc : %d\n", ccc)
	fmt.Printf("*ccc : %d\n", *ccc)
	fmt.Printf("**ccc : %d\n", **ccc)
}
