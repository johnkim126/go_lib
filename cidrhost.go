package main

import (
	"fmt"
	"net"
	"flag"
	goCidr "github.com/johnkim126/go-cidr/cidr"
)

func cidrhost(cidr string, hostIndex int) (string, error) {
	// adapted from https://github.com/hashicorp/terraform/blob/fe0cc3b0db0d1a5676c3d1a92ea8c5ff829b4233/config/interpolate_funcs.go#L253-L264
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return "", fmt.Errorf("invalid CIDR expression: %s", err)
	}

	ip, err := goCidr.Host(network, hostIndex)
	if err != nil {
		return "", err
	}

	return ip.String(), nil
}

func main() {
	cidr := flag.String("cidr", "10.0.0.0/16", "cidr")
	hostIndex := flag.Int("hostindex", 1, "host index")
	flag.Parse()
	ip, err := cidrhost(*cidr, *hostIndex)
	if err == nil {
		fmt.Print(ip)
	}
}
