package main

import (
	"net"
	"strconv"
)

// These functions need to be implemented
func cidrToMask(value string) string {
	ci, _ := strconv.Atoi(value)
	mask := net.IP(net.CIDRMask(ci, 32)).String()
	if mask == "0.0.0.0" {
		return "invalid"
	}
	return mask
}

func maskToCidr(value string) string {
	mask := net.IPMask(net.ParseIP(value).To4())
	prefixSize, _ := mask.Size()
	if prefixSize == 0 {
		return "invalid"
	}
	return strconv.Itoa(prefixSize)
}

func ipv4Validation(value string) bool {
	if net.ParseIP(value) == nil {
		return false
	}
	return true
}
