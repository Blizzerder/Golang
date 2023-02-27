//strings包与strconv包相关代码


package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//前缀与后缀
	var str string = "This is an example of a string"
	//前缀判断
	fmt.Printf("Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str , "Th"))
	//后缀判断
	fmt.Printf("Does the string \"%s\" have suffix %s? ", str, "ing")
	fmt.Printf("%t\n", strings.HasSuffix(str, "ing"))

	//字符串包含关系
	var s1 string = "abcd"
	var s2 string = "bc"
	fmt.Printf("Does the string \"%s\" contain \"%s\"? ", s1, s2)
	fmt.Printf("%t\n",strings.Contains(s1, s2))

	//子字符串在父字符串中出现的位置
	var s3 string = "Hi, I'm Marc, Hi."
	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(s3, "Marc"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(s3, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(s3, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(s3, "Burger"))

	//字符串替换
	var s4 string = "abcabcabc"
	s4 = strings.Replace(s4, "a", "0", -1)
	fmt.Printf("%s\n", s4)

	//统计字符串出现次数
	var s5 string = "Hello, how is it going, Hugo?"
	var s6 = "gggggggggg"

	fmt.Printf("Number of H's in %s is: ", s5)
	fmt.Printf("%d\n", strings.Count(s5, "H"))

	fmt.Printf("Number of double g's in %s is: ", s6)
	fmt.Printf("%d\n", strings.Count(s6, "gg"))

	//重复字符串
	var s7 string = "Hi there"
	var s8 string
	s8 = strings.Repeat(s7, 3)
	fmt.Printf("%s\n",s8)

	//将字符串大小写
	var s9 string = "Hey, how are you George?"
	var s10 string
	var s11 string

	fmt.Printf("The original string is: %s\n", s9)
	s10 = strings.ToLower(s9)
	fmt.Printf("The lowercase string is: %s\n", s10)
	s11 = strings.ToUpper(s9)
	fmt.Printf("The uppercase string is: %s\n", s11)

	//字符串的分割与拼接
	s12 := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(s12)						//分割
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	s13 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(s13, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2,";")				//拼接
	fmt.Printf("sl2 joined by ;: %s\n", str3)

	//字符串与其他类型转换：strconv包
	var s14 string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(s14)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
