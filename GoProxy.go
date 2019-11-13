package main

import (
	"fmt"
	"flag"
)

func main()  {
	fmt.Println("start server of redirect http flow to the socks5 channel")
	httpPort := flag.Int("httpPort", 12333, "http listen port")
	var name string
	flag.StringVar(&name, "name", "123", "name")

	flag.Parse()

	fmt.Println("ok:", *ok)
	fmt.Println("id:", *id)
	fmt.Println("port:", *port)
	fmt.Println("name:", name)
}
