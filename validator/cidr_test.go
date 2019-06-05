package validator

import "testing"

func TestNoOverlappingSubnets(t *testing.T) {
	subnets := []string{"10.0.0.0/24", "10.1.0.0/24", "10.2.0.0/24"}

	if err := CheckOverlapping(subnets); err != nil {
		t.Error("Subnet should not be overlapping: ", err.Error())
	}
}

func TestOverlappingSubnet(t *testing.T) {
	subnets := []string{"10.0.0.0/24", "10.0.0.0/24", "10.2.0.0/24"}

	if err := CheckOverlapping(subnets); err == nil {
		t.Error("10.0.0.0/24 overlaps with 10.0.0.0/24")
	}
}

