package validator

import (
	"log"
	"net"
	"fmt"
)

// Compare to subnets to see if they overlap
func isOverlapping(subnetA *net.IPNet, subnetB *net.IPNet) bool {
	return subnetA.Contains(subnetB.IP) || subnetB.Contains(subnetA.IP)
}

// Check to see if any of the subnets overlap
func CheckOverlapping(userSubnets []string) (error) {
	subnets := make([]*net.IPNet, 0)

	for _, cidrAddress := range userSubnets {
		_, subnet, err := net.ParseCIDR(cidrAddress)

		if err != nil {
			log.Fatalf("%s is not a valid cidr notation", cidrAddress)
		}

		subnets = append(subnets, subnet)
	}

	for i := range subnets {
		for _, subnet := range subnets[i + 1:] {
			if isOverlapping(subnet, subnets[i]) {
				return fmt.Errorf("The subnet %v overlaps with %v", subnet.String(), subnets[i].String())
			}
		}
	}

	return nil
}
