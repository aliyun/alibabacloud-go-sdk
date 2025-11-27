// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *CloneParameterGroupRequest
	GetOwnerId() *int64
	SetParameterGroupDesc(v string) *CloneParameterGroupRequest
	GetParameterGroupDesc() *string
	SetParameterGroupId(v string) *CloneParameterGroupRequest
	GetParameterGroupId() *string
	SetParameterGroupName(v string) *CloneParameterGroupRequest
	GetParameterGroupName() *string
	SetRegionId(v string) *CloneParameterGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CloneParameterGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CloneParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CloneParameterGroupRequest
	GetResourceOwnerId() *int64
	SetTargetRegionId(v string) *CloneParameterGroupRequest
	GetTargetRegionId() *string
}

type CloneParameterGroupRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The description of the parameter template in the destination region.
	//
	// example:
	//
	// CloneGroup1
	ParameterGroupDesc *string `json:"ParameterGroupDesc,omitempty" xml:"ParameterGroupDesc,omitempty"`
	// The ID of the parameter template. You can call the DescribeParameterGroups operation to query the parameter template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rpg-13ppdh****
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The name of the parameter template in the destination region.
	//
	// This parameter is required.
	//
	// example:
	//
	// tartestgroup
	ParameterGroupName *string `json:"ParameterGroupName,omitempty" xml:"ParameterGroupName,omitempty"`
	// The ID of the source region to which the parameter template belongs. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. You can leave this parameter empty.
	//
	// example:
	//
	// rg-acfmy****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the destination region. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-qingdao
	TargetRegionId *string `json:"TargetRegionId,omitempty" xml:"TargetRegionId,omitempty"`
}

func (s CloneParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CloneParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *CloneParameterGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CloneParameterGroupRequest) GetParameterGroupDesc() *string {
	return s.ParameterGroupDesc
}

func (s *CloneParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *CloneParameterGroupRequest) GetParameterGroupName() *string {
	return s.ParameterGroupName
}

func (s *CloneParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CloneParameterGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CloneParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CloneParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CloneParameterGroupRequest) GetTargetRegionId() *string {
	return s.TargetRegionId
}

func (s *CloneParameterGroupRequest) SetOwnerId(v int64) *CloneParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *CloneParameterGroupRequest) SetParameterGroupDesc(v string) *CloneParameterGroupRequest {
	s.ParameterGroupDesc = &v
	return s
}

func (s *CloneParameterGroupRequest) SetParameterGroupId(v string) *CloneParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *CloneParameterGroupRequest) SetParameterGroupName(v string) *CloneParameterGroupRequest {
	s.ParameterGroupName = &v
	return s
}

func (s *CloneParameterGroupRequest) SetRegionId(v string) *CloneParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *CloneParameterGroupRequest) SetResourceGroupId(v string) *CloneParameterGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CloneParameterGroupRequest) SetResourceOwnerAccount(v string) *CloneParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CloneParameterGroupRequest) SetResourceOwnerId(v int64) *CloneParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CloneParameterGroupRequest) SetTargetRegionId(v string) *CloneParameterGroupRequest {
	s.TargetRegionId = &v
	return s
}

func (s *CloneParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
