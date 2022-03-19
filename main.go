package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/valyala/fasthttp"
)

var (
	port *string
	path *string
	info = log.New(os.Stdout, "INFO: ", log.LstdFlags)
)

func init() {
	port = flag.String("port", "3030", "Define the port on which you want to run server")
	path = flag.String("path", ".", "Define the path which you want to server")
	flag.Parse()
}

func main() {
	info.Println(fmt.Sprintf(`Server getting started on port %s`, *port))
	if ok := fasthttp.ListenAndServe(fmt.Sprintf(":%s", *port), fasthttp.FSHandler(*path, 0)); ok != nil {
		log.Println(ok)
	}
}
