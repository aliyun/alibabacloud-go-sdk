// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAnycastServerRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetServiceLocation(v string) *DescribeAnycastServerRegionsRequest
	GetServiceLocation() *string
}

type DescribeAnycastServerRegionsRequest struct {
	// The access area from which you use the Anycast EIP to communicate with the Internet.
	//
	// Set the value to **international**, which specifies the areas outside the Chinese mainland.
	//
	// This parameter is required.
	//
	// example:
	//
	// international
	ServiceLocation *string `json:"ServiceLocation,omitempty" xml:"ServiceLocation,omitempty"`
}

func (s DescribeAnycastServerRegionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAnycastServerRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAnycastServerRegionsRequest) GetServiceLocation() *string {
	return s.ServiceLocation
}

func (s *DescribeAnycastServerRegionsRequest) SetServiceLocation(v string) *DescribeAnycastServerRegionsRequest {
	s.ServiceLocation = &v
	return s
}

func (s *DescribeAnycastServerRegionsRequest) Validate() error {
	return dara.Validate(s)
}
