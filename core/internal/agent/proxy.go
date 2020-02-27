package agent

import (
	"errors"
	"log"

	"github.com/txthinking/socks5"
)

// Socks5Proxy sock5 proxy server on agent, to use it, forward port 10800 to CC
func Socks5Proxy(op string) (err error) {
	switch op {
	case "on":
		go socks5Start()
	case "off":
		log.Print("Stopping Socks5Proxy")
		err = ProxyServer.Stop()
		if err != nil {
			log.Print(err)
		}
		ProxyServer = nil
	default:
		return errors.New("Operation not supported")
	}

	return err
}

func socks5Start() {
	var err error
	if ProxyServer == nil {
		socks5.Debug = true
		ProxyServer, err = socks5.NewClassicServer("127.0.0.1:10800", "127.0.0.1", "", "", 0, 0, 0, 60)
		if err != nil {
			log.Println(err)
			return
		}
	}

	log.Print("Socks5Proxy started")
	err = ProxyServer.Run(nil)
	if err != nil {
		log.Println(err)
	}
	log.Print("Socks5Proxy stopped")
}