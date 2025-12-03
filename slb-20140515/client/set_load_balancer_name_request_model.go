// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLoadBalancerNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLoadBalancerId(v string) *SetLoadBalancerNameRequest
	GetLoadBalancerId() *string
	SetLoadBalancerName(v string) *SetLoadBalancerNameRequest
	GetLoadBalancerName() *string
	SetOwnerAccount(v string) *SetLoadBalancerNameRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetLoadBalancerNameRequest
	GetOwnerId() *int64
	SetRegionId(v string) *SetLoadBalancerNameRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *SetLoadBalancerNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SetLoadBalancerNameRequest
	GetResourceOwnerId() *int64
}

type SetLoadBalancerNameRequest struct {
	// The region ID of the CLB instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// lb-bp1b6c719dfa08e******
	LoadBalancerId *string `json:"LoadBalancerId,omitempty" xml:"LoadBalancerId,omitempty"`
	// The new name of the CLB instance.
	//
	// The name must be 1 to 80 characters in length, and can contain letters, digits, periods (.), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	LoadBalancerName *string `json:"LoadBalancerName,omitempty" xml:"LoadBalancerName,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s SetLoadBalancerNameRequest) String() string {
	return dara.Prettify(s)
}

func (s SetLoadBalancerNameRequest) GoString() string {
	return s.String()
}

func (s *SetLoadBalancerNameRequest) GetLoadBalancerId() *string {
	return s.LoadBalancerId
}

func (s *SetLoadBalancerNameRequest) GetLoadBalancerName() *string {
	return s.LoadBalancerName
}

func (s *SetLoadBalancerNameRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetLoadBalancerNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetLoadBalancerNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetLoadBalancerNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SetLoadBalancerNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SetLoadBalancerNameRequest) SetLoadBalancerId(v string) *SetLoadBalancerNameRequest {
	s.LoadBalancerId = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetLoadBalancerName(v string) *SetLoadBalancerNameRequest {
	s.LoadBalancerName = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetOwnerAccount(v string) *SetLoadBalancerNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetOwnerId(v int64) *SetLoadBalancerNameRequest {
	s.OwnerId = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetRegionId(v string) *SetLoadBalancerNameRequest {
	s.RegionId = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetResourceOwnerAccount(v string) *SetLoadBalancerNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SetLoadBalancerNameRequest) SetResourceOwnerId(v int64) *SetLoadBalancerNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SetLoadBalancerNameRequest) Validate() error {
	return dara.Validate(s)
}
