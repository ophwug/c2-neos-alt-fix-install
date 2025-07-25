package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

func connect(ipAddress string) (*ssh.Client, error) {
	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		return nil, fmt.Errorf("unable to parse private key: %v", err)
	}

	config := &ssh.ClientConfig{
		User: "comma",
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", ipAddress+":22", config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect: %v", err)
	}

	return client, nil
}

func executeCommand(client *ssh.Client, command string, streamOutput bool) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", fmt.Errorf("failed to create session: %v", err)
	}
	defer session.Close()

	if streamOutput {
		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		err = session.Run(command)
		if err != nil {
			return "", fmt.Errorf("failed to run command with streaming: %v", err)
		}
		return "", nil
	}

	output, err := session.CombinedOutput(command)
	if err != nil {
		return "", fmt.Errorf("failed to run command: %v\nOutput: %s", err, output)
	}

	return string(output), nil
}