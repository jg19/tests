package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	cmd := exec.Command("./nlpnetTest.py")
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdin: %s", err.Error())
	}

	stdout, err := cmd.StdoutPipe()
	if nil != err {
		log.Fatalf("Error obtaining stdout: %s", err.Error())
	}

	reader := bufio.NewReader(stdout)
	go func(reader io.Reader) {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {

			log.Printf("Reading from subprocess: %s", scanner.Text())
			stdin.Write([]byte("O rato roeu a roupa do rei de roma.\n"))
			time.Sleep(1 * time.Second)
		}
	}(reader)

	if err := cmd.Start(); nil != err {
		log.Fatalf("Error starting program: %s, %s", cmd.Path, err.Error())
	}
	cmd.Wait()

}
