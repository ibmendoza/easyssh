package main

import (
	"flag"
	"fmt"
	"github.com/ibmendoza/easyssh"
	"log"
)

var user = flag.String("user", "", "username")
var server = flag.String("server", "", "server name or IP address")
var pswd = flag.String("pswd", "", "password")
var scp = flag.String("scp", "", "optional file to upload")
var cmd = flag.String("cmd", "", "optional command to run")

func main() {

	flag.Parse()

	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		User:     *user,
		Server:   *server,
		Password: *pswd,
		Port:     "22",
	}

	// Call Run method with command you want to run on remote server.
	response, err := ssh.Run(*cmd)
	// Handle errors
	if err != nil {
		//panic("Can't run remote command: " + err.Error())
		log.Fatal("Can't run remote command: " + err.Error())
	} else {
		fmt.Println(response)
	}

	if *scp != "" {
		// Call Scp method with file you want to upload to remote server.
		err = ssh.Scp(*scp)

		// Handle errors
		if err != nil {
			log.Fatal("Can't run remote command: " + err.Error())
		} else {
			fmt.Println("success")

			response, _ = ssh.Run("ls -al " + *scp)

			fmt.Println(response)
		}
	}
}
