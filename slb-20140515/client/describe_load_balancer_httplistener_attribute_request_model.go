// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLoadBalancerHTTPListenerAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetListenerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeRequest
	GetListenerPort() *int32
	SetLoadBalancerId(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoadBalancerHTTPListenerAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoadBalancerHTTPListenerAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeLoadBalancerHTTPListenerAttributeRequest struct {
	// The frontend port that is used by the CLB instance.
	//
	// Valid values: **1*	- to **65535**.
	//
	// This parameter is required.
	//
	// example:
	//
	// 80
	ListenerPort *int32 `json:"ListenerPort,omitempty" xml:"ListenerPort,omitempty"`
	// The CLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1uaunez0uho0zf0****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the CLB instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLoadBalancerHTTPListenerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerHTTPListenerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) GetListenerPort() *int32 {
	return s.ListenerPort
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetRegionId(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerHTTPListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
