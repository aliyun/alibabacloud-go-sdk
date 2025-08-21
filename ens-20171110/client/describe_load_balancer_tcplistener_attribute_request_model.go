// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerTCPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetLoadBalancerId() *string
}

type DescribeLoadBalancerTCPListenerAttributeRequest struct {
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
	// lb-5snthcyu1x10g7tywj7iu****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DescribeLoadBalancerTCPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerTCPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
