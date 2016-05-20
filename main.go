package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/operable/cogexec/messages"
	"os"
	"os/exec"
	"time"
)

func ackDeath(success bool, encoder *gob.Encoder) {
	encoder.Encode(&messages.ExecCommandResponse{
		Success: success,
		Dead:    true,
	})
}

func runCommand(req *messages.ExecCommandRequest, encoder *gob.Encoder) {
	command := exec.Command(req.Executable)
	command.Env = req.Env
	input := bytes.NewBuffer(req.CogEnv)
	stdout := bytes.NewBuffer([]byte{})
	stderr := bytes.NewBuffer([]byte{})
	command.Stdin = input
	command.Stdout = stdout
	command.Stderr = stderr
	start := time.Now()
	err := command.Run()
	finish := time.Now()
	stderrFinal := stderr.Bytes()
	if len(stderrFinal) == 0 && err != nil {
		stderrFinal = []byte(fmt.Sprintf("%s\n", err))
	}
	resp := &messages.ExecCommandResponse{
		Executable: req.Executable,
		Success:    err == nil,
		Stdout:     stdout.Bytes(),
		Stderr:     stderrFinal,
		Elapsed:    finish.Sub(start),
	}
	encoder.Encode(resp)
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	decoder := gob.NewDecoder(stdin)
	encoder := gob.NewEncoder(os.Stdout)
	for {
		req := &messages.ExecCommandRequest{}
		err := decoder.Decode(req)
		if err != nil {
			ackDeath(false, encoder)
		}
		if req.Die {
			ackDeath(true, encoder)
			break
		}
		runCommand(req, encoder)
	}
}
