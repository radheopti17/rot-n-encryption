package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {
	li,err :=net.Listen("tcp",":8080")
	if err !=nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}
func handle(conn net.Conn)  {

	fmt.Fprintf(conn,"Value of n (ex. rot-13/rot-15 n must be between 1 to 26)")
	var n =-1
	scanner := bufio.NewScanner(conn)
	for scanner.Scan(){
		line:=scanner.Text()
		if n<=0 {
			val, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}
			if val<=0 && val>26 {
				fmt.Fprintf(conn,"Value of n (ex. rot-13/rot-15 n must be between 1 to 26)")
				continue
			}
			n=val
			continue
		}
		temp:=[]byte(line)
		out:= make([]byte,len(temp))
		for i,j := range temp {
			if j<=90 && j>=65 {
				if j<=byte(90-n) {
					out[i]=byte(int(j)+n)
				} else {
					out[i]=byte(int(j)-26+n)
				}
			} else {
				if j<= byte(109-n) {
					out[i]=byte(int(j)+n)
				} else {
					out[i]=byte(int(j)-26+n)
				}
			}
		}
		fmt.Println(out)
		fmt.Fprintf(conn,"%s\n",out)
	}
	defer conn.Close()
	fmt.Println("Connection Closed")
}
