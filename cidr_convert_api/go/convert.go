package main

import (
	"fmt"
	"net"
	"strconv"
)

// These functions need to be implemented
func cidrToMask(value string) string {
	ci, _ := strconv.Atoi(value)
	mask := net.IP(net.CIDRMask(ci, 32)).String()
	return mask
}

func maskToCidr(value string) string {
	fmt.Println(value)
	mask := net.IPMask(net.ParseIP(value).To4())
	prefixSize, _ := mask.Size()
	fmt.Println(prefixSize)
	return strconv.Itoa(prefixSize)
}

func ipv4Validation(value string) bool {

	if net.ParseIP(value) == nil {
		return false
	}
	return true
}
