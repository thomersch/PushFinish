package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
)

func concat_args(arglist []string) (string, string) {
	cmd := arglist[1]
	args := ""

	arglist = arglist[2:]

	for arg := range arglist {
		args += arglist[arg] + " "
	}

	return cmd, args
}

func get_hostname(hostname chan string) {
	output, err := exec.Command("hostname").Output()

	if err != nil {
		hostname <- "(unknown hostname)"
		return
	}

	hostname <- fmt.Sprintf("%s", output)
}

func get_tokens() (string, string) {
	token := os.Getenv("PUSHFINISH_TOKEN")
	user := os.Getenv("PUSHFINISH_USER")
	
	if(token == "" || user == "") {
		log.Fatal("Either application or user token are not set. Please make sure that the environmental variables PUSHFINISH_TOKEN and PUSHFINISH_USER are set.")
	}

	return token, user
}

func notify(hostname chan string, token string, user string) {
	message := fmt.Sprintf("%s has finished the operation.", <-hostname)

	postdata := url.Values{}
	postdata.Add("token", token)
	postdata.Add("user", user)
	postdata.Add("message", message)

	_, err := http.PostForm("https://api.pushover.net/1/messages.json", postdata)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) == 1 {
		log.Print("In order to be notified after an application has completed operation, just type: pushfinish <your command>")
		os.Exit(0)
	}

	hostname := make(chan string)
	go get_hostname(hostname)

	token, user := get_tokens()

	childcommand, childargs := concat_args(os.Args)

	childprocess := exec.Command(childcommand, childargs)
	childprocess.Stdout = os.Stdout
	childprocess.Stderr = os.Stderr
	childprocess.Run()

	notify(hostname, token, user)
}
