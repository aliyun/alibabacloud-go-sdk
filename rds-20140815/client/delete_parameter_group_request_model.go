// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteParameterGroupRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *DeleteParameterGroupRequest
	GetParameterGroupId() *string
	SetRegionId(v string) *DeleteParameterGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DeleteParameterGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DeleteParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteParameterGroupRequest
	GetResourceOwnerId() *int64
}

type DeleteParameterGroupRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameter template ID. You can call the DescribeParameterGroups operation to query the parameter template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rpg-gfs****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute to obtain the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteParameterGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DeleteParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteParameterGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DeleteParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteParameterGroupRequest) SetOwnerId(v int64) *DeleteParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetParameterGroupId(v string) *DeleteParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetRegionId(v string) *DeleteParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetResourceGroupId(v string) *DeleteParameterGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetResourceOwnerAccount(v string) *DeleteParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteParameterGroupRequest) SetResourceOwnerId(v int64) *DeleteParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
