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
	SetOwnerAccount(v string) *DescribeLoadBalancerAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLoadBalancerAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLoadBalancerAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLoadBalancerAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLoadBalancerAttributeRequest
	GetResourceOwnerId() *int64
}

type DescribeLoadBalancerAttributeRequest struct {
	// The CLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1b6c719dfa08ex****
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

func (s DescribeLoadBalancerAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLoadBalancerAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeLoadBalancerAttributeRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DescribeLoadBalancerAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLoadBalancerAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLoadBalancerAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLoadBalancerAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLoadBalancerAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLoadBalancerAttributeRequest) SetLoadBalancerId(v string) *DescribeLoadBalancerAttributeRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetOwnerAccount(v string) *DescribeLoadBalancerAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetOwnerId(v int64) *DescribeLoadBalancerAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetRegionId(v string) *DescribeLoadBalancerAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetResourceOwnerAccount(v string) *DescribeLoadBalancerAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) SetResourceOwnerId(v int64) *DescribeLoadBalancerAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLoadBalancerAttributeRequest) Validate() error {
	return dara.Validate(s)
}
