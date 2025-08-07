// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBType(v string) *DescribeParameterGroupRequest
	GetDBType() *string
	SetOwnerAccount(v string) *DescribeParameterGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParameterGroupRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *DescribeParameterGroupRequest
	GetParameterGroupId() *string
	SetRegionId(v string) *DescribeParameterGroupRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeParameterGroupRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeParameterGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterGroupRequest
	GetResourceOwnerId() *int64
}

type DescribeParameterGroupRequest struct {
	DBType       *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the parameter template.
	//
	// > You can call the [DescribeParameterGroups](https://help.aliyun.com/document_detail/207178.html) operation to query the details of all parameter templates of a specified region, such as the ID of a parameter template.
	//
	// This parameter is required.
	//
	// example:
	//
	// pcpg-**************
	ParameterGroupId *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	// The region ID.
	//
	// >You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query all regions that are available within your account, such as the region ID.
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

func (s DescribeParameterGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeParameterGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParameterGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterGroupRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *DescribeParameterGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParameterGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeParameterGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterGroupRequest) SetDBType(v string) *DescribeParameterGroupRequest {
	s.DBType = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetOwnerAccount(v string) *DescribeParameterGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetOwnerId(v int64) *DescribeParameterGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetParameterGroupId(v string) *DescribeParameterGroupRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetRegionId(v string) *DescribeParameterGroupRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetResourceGroupId(v string) *DescribeParameterGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetResourceOwnerAccount(v string) *DescribeParameterGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterGroupRequest) SetResourceOwnerId(v int64) *DescribeParameterGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterGroupRequest) Validate() error {
	return dara.Validate(s)
}
