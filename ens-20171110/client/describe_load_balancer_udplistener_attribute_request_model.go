// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerUDPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeLoadBalancerUDPListenerAttributeRequest
	GetLoadBalancerId() *string
}

type DescribeLoadBalancerUDPListenerAttributeRequest struct {
	// The listening port that you want to query. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5pzipr2fszqtl2xf64uy5****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DescribeLoadBalancerUDPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerUDPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
