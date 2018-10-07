// +skip_license_check

/*
This file contains portions of code directly taken from the 'xenolf/lego' project.
A copy of the license for this code can be found in the file named LICENSE in
this directory.
*/

package util

import "fmt"

// DNS01Record returns a DNS record which will fulfill the `dns-01` challenge
// TODO: move this into a non-generic place by resolving import cycle in dns package
func DNS01Record(domain, value string, nameservers []string) (string, string, int, error) {
	return fmt.Sprintf("_acme-challenge.%s.", domain), value, 60, nil
}
