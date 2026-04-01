// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableDetail(v bool) *DescribeParameterGroupsRequest
	GetEnableDetail() *bool
	SetOwnerId(v int64) *DescribeParameterGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeParameterGroupsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeParameterGroupsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeParameterGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterGroupsRequest
	GetResourceOwnerId() *int64
}

type DescribeParameterGroupsRequest struct {
	// Specifies whether to return the parameter overview.
	//
	// 	- **false*	- (default): The parameter overview is returned.
	//
	// 	- **true**: The parameter overview is not returned.
	//
	// example:
	//
	// false
	EnableDetail *bool  `json:"EnableDetail,omitempty" xml:"EnableDetail,omitempty"`
	OwnerId      *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeParameterGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupsRequest) GetEnableDetail() *bool {
	return s.EnableDetail
}

func (s *DescribeParameterGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParameterGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeParameterGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterGroupsRequest) SetEnableDetail(v bool) *DescribeParameterGroupsRequest {
	s.EnableDetail = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetOwnerId(v int64) *DescribeParameterGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetRegionId(v string) *DescribeParameterGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetResourceGroupId(v string) *DescribeParameterGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetResourceOwnerAccount(v string) *DescribeParameterGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetResourceOwnerId(v int64) *DescribeParameterGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterGroupsRequest) Validate() error {
	return dara.Validate(s)
}
