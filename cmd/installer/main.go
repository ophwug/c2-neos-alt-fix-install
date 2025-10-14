package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Fork struct {
	Name   string
	Owner  string
	Branch string
}

var forks = []Fork{
	{Name: "comma.ai Stock 0.8.13 (release2 branch)", Owner: "commaai", Branch: "release2"},
	{Name: "legacypilot (recommended for \"current\" use by Mr. One)", Owner: "dragonpilot-community", Branch: "d2"},
	{Name: "dragonpilot (not as stable/enthusiasts)", Owner: "dragonpilot-community", Branch: "r2"},
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the C2 NEOS Alternate Fix Install tool.")
	fmt.Println("This tool bypasses a bug in the NEOS setup screen where it always")
	fmt.Println("reports 'The network is not connected to the internet' even when it is.")
	fmt.Println("-----------------------------------------------------------------------")
	fmt.Println("First, ensure your comma two is on the same Wi-Fi network as this computer.")
	fmt.Println("On your device, go to More Options.")
	fmt.Println("Touch the triple-dot icon in the upper right corner and select Advanced.")
	fmt.Println("Scroll down and note the IPv4 address.")
	fmt.Println("It will likely start with 192.168.x.x, 10.x.x.x, or 172.16.x.x.")
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Print("Enter the device IP address: ")
	ipAddress, _ := reader.ReadString('\n')
	ipAddress = strings.TrimSpace(ipAddress)

	fmt.Println("\nAvailable forks:")
	for i, fork := range forks {
		fmt.Printf("%d. %s (%s/%s)\n", i+1, fork.Name, fork.Owner, fork.Branch)
	}
	fmt.Printf("%d. Custom\n", len(forks)+1)

	var githubOwner, githubBranch string
	for {
		fmt.Print("Select a fork to install: ")
		response, _ := reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(response))

		if err == nil && choice > 0 && choice <= len(forks)+1 {
			if choice <= len(forks) {
				selectedFork := forks[choice-1]
				githubOwner = selectedFork.Owner
				githubBranch = selectedFork.Branch
			} else {
				fmt.Println("\nNote: Custom option requires a valid branch name (not a tag or version number).")
				fmt.Println("For official comma.ai releases, use option 1 instead.")
				fmt.Print("Enter the GitHub repository owner: ")
				owner, _ := reader.ReadString('\n')
				githubOwner = strings.TrimSpace(owner)

				fmt.Print("Enter the GitHub branch name: ")
				branch, _ := reader.ReadString('\n')
				githubBranch = strings.TrimSpace(branch)
			}
			break
		}
		fmt.Println("Invalid selection. Please try again.")
	}

	fmt.Println("\n-----------------------------------------------------------------------")
	fmt.Printf("Device IP: %s\n", ipAddress)
	fmt.Printf("GitHub Repo: https://github.com/%s/openpilot.git\n", githubOwner)
	fmt.Printf("Branch: %s\n", githubBranch)
	fmt.Println("-----------------------------------------------------------------------")

	fmt.Println("\nConnecting to device...")
	client, err := connect(ipAddress)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer client.Close()

	if err := runSetup(client, githubOwner, githubBranch); err != nil {
		fmt.Printf("An error occurred during setup: %v\n", err)
	} else {
		fmt.Println("Installation script finished successfully.")
	}

	waitForExit()
}
