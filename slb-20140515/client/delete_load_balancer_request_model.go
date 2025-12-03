// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLoadBalancerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *DeleteLoadBalancerRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *DeleteLoadBalancerRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteLoadBalancerRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLoadBalancerRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteLoadBalancerRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteLoadBalancerRequest
	GetResourceOwnerId() *int64
}

type DeleteLoadBalancerRequest struct {
	// The SLB instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1h66tp5uat8********
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SLB instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteLoadBalancerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLoadBalancerRequest) GoString() string {
	return s.String()
}

func (s *DeleteLoadBalancerRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *DeleteLoadBalancerRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteLoadBalancerRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLoadBalancerRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLoadBalancerRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteLoadBalancerRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteLoadBalancerRequest) SetLoadBalancerId(v string) *DeleteLoadBalancerRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetOwnerAccount(v string) *DeleteLoadBalancerRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetOwnerId(v int64) *DeleteLoadBalancerRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetRegionId(v string) *DeleteLoadBalancerRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetResourceOwnerAccount(v string) *DeleteLoadBalancerRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteLoadBalancerRequest) SetResourceOwnerId(v int64) *DeleteLoadBalancerRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteLoadBalancerRequest) Validate() error {
	return dara.Validate(s)
}
