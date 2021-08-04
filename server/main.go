package main

import log "github.com/sirupsen/logrus"

const (
	Address string = "localhost:8123"
	Method  string = "tcp"
)

func init() {
	InitLog()
}

func main() {

}

func InitLog() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
	})
}
