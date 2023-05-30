package main

import (
	"github.com/urfave/cli"
	"sync"
	"time"
)

// VERSION indicates which version of the binary is running.
var VERSION string

var (
	fileName string
	url      string
	port     string
	protocol = []string{"tcp"}
	timeout  = time.Microsecond * 2000
	wg       sync.WaitGroup
)

func main() {

	a := cli.NewApp()
	a.Name = "Video Streaming Client /Files"
	a.Usage = "Client part of audio and video streamer / stream files"
	a.Author = "Valentyn Ponomarenko"
	a.Version = VERSION
	a.Email = "ValentynPonomarenko@gmail.com"

	a.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url",
			Usage: "Streaming Server end-point, e.g. --url https://1270.0.0.1:8080",
		},
		cli.StringFlag{
			Name:  "fileName, f",
			Usage: "file name path",
		},
	}
}
