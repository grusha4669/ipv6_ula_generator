package main

import (
	"crypto/rand"
	"fmt"
	"net/netip"
)

func main() {
	// Create a 16-byte array for the IPv6 address
	var bytes [16]byte

	// Fill the array with random bytes
	_, err := rand.Read(bytes[:])
	if err != nil {
		panic(err)
	}

	// Forcibly set the first byte to 0xfd (according to the fd00::/8 network)
	bytes[0] = 0xfd

	// Form a valid IPv6 address
	ipv6 := netip.AddrFrom16(bytes)

	fmt.Println(ipv6.String())
	fmt.Println(netip.PrefixFrom(ipv6, 64).Masked())
	fmt.Println(netip.PrefixFrom(ipv6, 48).Masked())
}
