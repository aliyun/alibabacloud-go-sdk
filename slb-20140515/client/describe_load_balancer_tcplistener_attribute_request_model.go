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
	SetOwnerAccount(v string) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoadBalancerTCPListenerAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeLoadBalancerTCPListenerAttributeRequest struct {
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
	// lb-bp1ygod3yctvg1y****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can query the region ID from the [Regions and zones](https://help.aliyun.com/document_detail/40654.html) list or by calling the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetRegionId(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerTCPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadBalancerTCPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
