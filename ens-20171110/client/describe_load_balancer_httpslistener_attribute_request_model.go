// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerHTTPSListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest
	GetLoadBalancerId() *string
}

type DescribeLoadBalancerHTTPSListenerAttributeRequest struct {
	// The listening port that you want to query. Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 443
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the ELB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-5qoxu2rsr0ytanpn4r3i****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
}

func (s DescribeLoadBalancerHTTPSListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPSListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
