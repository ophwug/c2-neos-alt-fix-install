//go:build darwin

package main

import "fmt"

// printNetworkPermissionWarning prints a warning about macOS network permissions
// when connection to the device fails. This is only compiled on macOS builds.
func printNetworkPermissionWarning() {
	fmt.Println("\n⚠️  NETWORK CONNECTION TROUBLESHOOTING")
	fmt.Println("═══════════════════════════════════════════════════════════════")
	fmt.Println("If you recently installed or reinstalled your terminal application")
	fmt.Println("(Terminal, iTerm2, etc.), macOS may have asked you to grant network")
	fmt.Println("access permission. If you accidentally denied this permission, all")
	fmt.Println("network commands will fail.")
	fmt.Println()
	fmt.Println("TO FIX THIS ISSUE:")
	fmt.Println("1. Open System Settings")
	fmt.Println("2. Go to: Privacy & Security → Local Network")
	fmt.Println("3. Find your terminal app in the list (Terminal, iTerm2, etc.)")
	fmt.Println("4. Make sure the toggle next to it is ENABLED")
	fmt.Println("5. If not listed, restart your terminal - macOS should prompt you")
	fmt.Println("6. After granting permission, restart this installer")
	fmt.Println()
	fmt.Println("For more details, see:")
	fmt.Println("https://gitlab.com/gnachman/iterm2/-/issues/12071")
	fmt.Println("═══════════════════════════════════════════════════════════════\n")
}
