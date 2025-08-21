// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerSpec(v string) *DescribeLoadBalancerSpecRequest
	GetLoadBalancerSpec() *string
}

type DescribeLoadBalancerSpecRequest struct {
	// The specifications of the ELB instance.
	//
	// example:
	//
	// elb.s2.small
	LoadBalancerSpec *string `json:"LoadBalancerSpec,omitempty" xml:"LoadBalancerSpec,omitempty"`
}

func (s DescribeLoadBalancerSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerSpecRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerSpecRequest) GetLoadBalancerSpec() *string {
	return s.LoadBalancerSpec
}

func (s *DescribeLoadBalancerSpecRequest) SetLoadBalancerSpec(v string) *DescribeLoadBalancerSpecRequest {
	s.LoadBalancerSpec = &v
	return s
}

func (s *DescribeLoadBalancerSpecRequest) Validate() error {
	return dara.Validate(s)
}
