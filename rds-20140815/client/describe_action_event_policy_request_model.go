// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActionEventPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeActionEventPolicyRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeActionEventPolicyRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeActionEventPolicyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeActionEventPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActionEventPolicyRequest
	GetResourceOwnerId() *int64
}

type DescribeActionEventPolicyRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeActionEventPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActionEventPolicyRequest) GoString() string {
	return s.String()
}

func (s *DescribeActionEventPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActionEventPolicyRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeActionEventPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeActionEventPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActionEventPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActionEventPolicyRequest) SetOwnerId(v int64) *DescribeActionEventPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActionEventPolicyRequest) SetRegionId(v string) *DescribeActionEventPolicyRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeActionEventPolicyRequest) SetResourceGroupId(v string) *DescribeActionEventPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeActionEventPolicyRequest) SetResourceOwnerAccount(v string) *DescribeActionEventPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActionEventPolicyRequest) SetResourceOwnerId(v int64) *DescribeActionEventPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActionEventPolicyRequest) Validate() error {
	return dara.Validate(s)
}
