// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeRequest
	GetLoadBalancerId() *string
}

type DescribeLoadBalancerAttributeRequest struct {
	// The ID of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5rcvo1n1t3hykfhhjwjgq****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DescribeLoadBalancerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
