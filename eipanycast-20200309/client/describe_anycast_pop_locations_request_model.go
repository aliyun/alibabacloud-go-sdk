// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnycastPopLocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceLocation(v string) *DescribeAnycastPopLocationsRequest
	GetServiceLocation() *string
}

type DescribeAnycastPopLocationsRequest struct {
	// The access area of the Anycast elastic IP address (EIP).
	//
	// Set the value to **international**, which specifies the areas outside the Chinese mainland.
	//
	// example:
	//
	// international
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
}

func (s DescribeAnycastPopLocationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastPopLocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnycastPopLocationsRequest) GetServiceLocation() *string {
	return s.ServiceLocation
}

func (s *DescribeAnycastPopLocationsRequest) SetServiceLocation(v string) *DescribeAnycastPopLocationsRequest {
	s.ServiceLocation = &v
	return s
}

func (s *DescribeAnycastPopLocationsRequest) Validate() error {
	return dara.Validate(s)
}
