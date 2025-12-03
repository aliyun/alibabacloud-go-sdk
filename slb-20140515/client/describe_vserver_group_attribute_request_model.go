// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVServerGroupAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeVServerGroupAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVServerGroupAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVServerGroupAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVServerGroupAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVServerGroupAttributeRequest
	GetResourceOwnerId() *int64
	SetVServerGroupId(v string) *DescribeVServerGroupAttributeRequest
	GetVServerGroupId() *string
}

type DescribeVServerGroupAttributeRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the Classic Load Balancer (CLB) instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the vServer group.
	//
	// This parameter is required.
	//
	// example:
	//
	// rsp-cige6****
	VServerGroupId *string `json:"VServerGroupId,omitempty" xml:"VServerGroupId,omitempty"`
}

func (s DescribeVServerGroupAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVServerGroupAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeVServerGroupAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVServerGroupAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVServerGroupAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVServerGroupAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVServerGroupAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVServerGroupAttributeRequest) GetVServerGroupId() *string {
	return s.VServerGroupId
}

func (s *DescribeVServerGroupAttributeRequest) SetOwnerAccount(v string) *DescribeVServerGroupAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVServerGroupAttributeRequest) SetOwnerId(v int64) *DescribeVServerGroupAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVServerGroupAttributeRequest) SetRegionId(v string) *DescribeVServerGroupAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVServerGroupAttributeRequest) SetResourceOwnerAccount(v string) *DescribeVServerGroupAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVServerGroupAttributeRequest) SetResourceOwnerId(v int64) *DescribeVServerGroupAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVServerGroupAttributeRequest) SetVServerGroupId(v string) *DescribeVServerGroupAttributeRequest {
	s.VServerGroupId = &v
	return s
}

func (s *DescribeVServerGroupAttributeRequest) Validate() error {
	return dara.Validate(s)
}
