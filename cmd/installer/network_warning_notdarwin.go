//go:build !darwin

package main

// printNetworkPermissionWarning is a no-op on non-macOS platforms.
// The network permission warning is only relevant for macOS.
func printNetworkPermissionWarning() {
	// Do nothing on non-macOS platforms
}
