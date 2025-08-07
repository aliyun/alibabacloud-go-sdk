// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeleteParameterGroupRequest
	GetOwnerAccount() *string
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
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The parameter template ID.
	//
	//
	//
	// >  You can call the [DescribeParameterGroups](https://help.aliyun.com/document_detail/207178.html) operation to query the parameter template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcpg-**************
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The region ID.
	//
	//
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query available regions.
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
	// rg-************
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

func (s *DeleteParameterGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
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

func (s *DeleteParameterGroupRequest) SetOwnerAccount(v string) *DeleteParameterGroupRequest {
	s.OwnerAccount = &v
	return s
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
