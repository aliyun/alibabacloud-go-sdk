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
	SetOwnerAccount(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoadBalancerHTTPSListenerAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoadBalancerHTTPSListenerAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeLoadBalancerHTTPSListenerAttributeRequest struct {
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
	// lb-bp1mxu5r8lau****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the CLB instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
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

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetListenerPort(v int32) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.ListenerPort = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetRegionId(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerHTTPSListenerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadBalancerHTTPSListenerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
