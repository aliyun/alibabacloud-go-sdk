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
	SetOwnerAccount(v string) *DescribeLoadBalancerUDPListenerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoadBalancerUDPListenerAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLoadBalancerUDPListenerAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLoadBalancerUDPListenerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoadBalancerUDPListenerAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeLoadBalancerUDPListenerAttributeRequest struct {
	// The frontend port used by the CLB instance.
	//
	// Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1rtfnodmywb43e*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the Classic Load Balancer (CLB) instance is deployed.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetRegionId(v string) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerUDPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadBalancerUDPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
