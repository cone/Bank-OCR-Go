package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

var m = make(map[string]string)

func getAccountNumbers(path string) string{
	file, err := os.Open(path)
	var numbers string
	if(err != nil){
		fmt.Println(err)
	}
	bufferedReader := bufio.NewReader(file)
	count := 0
	var lines []string
	for{
		str, bferror := bufferedReader.ReadString('\n')
		if count == 3 {
			numbers = numbers+getAccountNumber(lines)+"\n"
			lines = []string{}
			count = 0
			continue
		}
		if bferror != nil{
			if bferror == io.EOF {
				break
			}
			fmt.Println(bferror)
		}
		lines = append(lines, str)
		count = count+1
	}
	file.Close()
	return numbers
}

func getAccountNumber(lines []string) string{
	init := 0
	ocrNum := ""
	accountNo := ""
	for i:=0 ; i<9 ; i++{
		for _, val := range lines{
			ocrNum = ocrNum + val[init:init+3]
		}
		init = init+3
		accountNo = accountNo+getNumber(ocrNum)
		ocrNum=""
	}
	return accountNo
}

func getNumber(key string) string{
	return m[key]
}

func initialize(){
	m[" _ | ||_|"] = "0"
	m["     |  |"] = "1"
	m[" _  _||_ "] = "2"
	m[" _  _| _|"] = "3"
	m["   |_|  |"] = "4"
	m[" _ |_  _|"] = "5"
	m[" _ |_ |_|"] = "6"
	m[" _   |  |"] = "7"
	m[" _ |_||_|"] = "8"
	m[" _ |_| _|"] = "9"
}

func main(){
	initialize()
	fmt.Println(getAccountNumbers("ocr.txt"))
}