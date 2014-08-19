package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func getLines(){
	file, err := os.Open("ocr.txt")
	if(err != nil){
		fmt.Println(err)
	}
	bufferedReader := bufio.NewReader(file)
	count := 0
	var lines []string
	for{
		str, bferror := bufferedReader.ReadString('\n')
		if bferror != nil{
			if bferror == io.EOF {
				break
			}
			fmt.Println(bferror)
		}
		if count == 3 {
			break
		}
		lines = append(lines, str)
		count = count+1
	}
	init := 0
	for i:=0 ; i<9 ; i++{
		for _, val := range lines{
			fmt.Println(val[init:init+3])
		}
		init = init+3
	}
	file.Close()
}

func main(){
	getLines()
}