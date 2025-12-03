// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerDeleteProtectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteProtection(v string) *SetLoadBalancerDeleteProtectionRequest
	GetDeleteProtection() *string
	SetLoadBalancerId(v string) *SetLoadBalancerDeleteProtectionRequest
	GetLoadBalancerId() *string
	SetOwnerAccount(v string) *SetLoadBalancerDeleteProtectionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetLoadBalancerDeleteProtectionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetLoadBalancerDeleteProtectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetLoadBalancerDeleteProtectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetLoadBalancerDeleteProtectionRequest
	GetResourceOwnerId() *int64
}

type SetLoadBalancerDeleteProtectionRequest struct {
	// Specify whether to enable or disable deletion protection for the SLB instance.
	//
	// Valid values: **on and off**.
	//
	// This parameter is required.
	//
	// example:
	//
	// off
	DeleteProtection *string `json:"DeleteProtection,omitempty" xml:"DeleteProtection,omitempty"`
	// The ID of the SLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1b6c719dfa08e*****
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region where the SLB instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/27584.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s SetLoadBalancerDeleteProtectionRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerDeleteProtectionRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerDeleteProtectionRequest) GetDeleteProtection() *string {
	return s.DeleteProtection
}

func (s *SetLoadBalancerDeleteProtectionRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerDeleteProtectionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerDeleteProtectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLoadBalancerDeleteProtectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLoadBalancerDeleteProtectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetLoadBalancerDeleteProtectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetDeleteProtection(v string) *SetLoadBalancerDeleteProtectionRequest {
	s.DeleteProtection = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetLoadBalancerId(v string) *SetLoadBalancerDeleteProtectionRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetOwnerAccount(v string) *SetLoadBalancerDeleteProtectionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetOwnerId(v int64) *SetLoadBalancerDeleteProtectionRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetRegionId(v string) *SetLoadBalancerDeleteProtectionRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerDeleteProtectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) SetResourceOwnerId(v int64) *SetLoadBalancerDeleteProtectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerDeleteProtectionRequest) Validate() error {
	return dara.Validate(s)
}
