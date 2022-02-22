package main

import (
	"BracketValidation/utils"
	"fmt"
)

func main() {

	shortStrings := []string{"", "1", "[", "]", "[1", "1]", "[]", "][", "{]", "(}", "[1]", ")1)", "123", "12()[]3", "1{2()[]3}", "[(]", "[)]", "{[", "([])", "([]{}())", "(){}()", "([{}]{}(()))"}

	for _, item := range shortStrings {
		fmt.Printf("%s is %t\n", item, utils.BracketPairValidation(item))
	}

	fmt.Println("***** short string testing ends ******")

	longString1 := "{abcdefghi(jklmnoq([]prst{}uv)wxyz0)123456789abcdefghijkl[mno{pq(rs)tuv}wxy012]3456789}"
	fmt.Printf("long string 1 is %t\n", utils.BracketPairValidation(longString1))

	longString2 := "{abcdefghi(jklmnoq([]prst{}uv)wxyz0)1234[56)7]89abcdefghijkl[mno{pq(rs)tuv}wxy012]3456789}"
	fmt.Printf("long string 2 is %t\n", utils.BracketPairValidation(longString2))

	longString3 := "ab{cd9ef[ghijklmnoq([]prst{}uv)wxyz0123456789abcdefghi(jklmnoq([]{p{rst{}}uv}))wxyz0123456789abcdefghi(jklmnoq(prst{}uvwxyz0)123456789abcd)efghi(jklmnoq(prst{}uvwxyz0)12345)6789" +
		"abcdefghijklmnoqprstu(vwxyz0123456789abcdefghijklmnoqprstuvwxyz0123456789abcdefghijklmnoqprstuvwxyz0123456789abcdefghijklmnoqprstuvw(xyz[012]{34})56789)]}"
	fmt.Printf("long string 3 is %t\n", utils.BracketPairValidation(longString3))

	longString4 := "ab{cd9ef[ghijklmnoq([]prst{}uv)wxyz0123456789abcdefghi(jklmnoq([]{p{rst{}}uv}))wxyz0123456789abcdefghi(jklmnoq(prst{}uvwxyz0)123456789abcd)efghi(jklmnoq(prst{}uvwxyz0)12345)6789" +
		"abcdefghijklmnoqprstu(vwxyz0123456789abcdefghijklmnoqprstuvwxyz0123456789abcdefghijklmnoqprstuvwxyz0123456789abcdefghijklmnoqprstuvw(xyz[012]{34})56789])}"
	fmt.Printf("long string 4 is %t\n", utils.BracketPairValidation(longString4))

	longString5 := "{{{{{{{{{{[[[[[([(((([[[[[{{{{{{{{{{((((((((((((((((((({{{}}})))))))))))))))))))}}}}}}}}}}]]]]]))))])]]]]]}}}}}}}([[]])}}}"
	fmt.Printf("long string 5 is %t\n", utils.BracketPairValidation(longString5))
}
