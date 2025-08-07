// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBType(v string) *DescribeParameterGroupsRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeParameterGroupsRequest
	GetDBVersion() *string
	SetOwnerAccount(v string) *DescribeParameterGroupsRequest
	GetOwnerAccount() *string
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
	// The type of the database engine. Only **MySQL*	- is supported.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database engine. Valid values:
	//
	// 	- **5.6**
	//
	// 	- **5.7**
	//
	// 	- **8.0**
	//
	// example:
	//
	// 8.0
	DBVersion    *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the cluster.
	//
	// > You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query all regions that are available for your account, such as the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the virtual node belongs.
	//
	// example:
	//
	// rg-************
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

func (s *DescribeParameterGroupsRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeParameterGroupsRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeParameterGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
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

func (s *DescribeParameterGroupsRequest) SetDBType(v string) *DescribeParameterGroupsRequest {
	s.DBType = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetDBVersion(v string) *DescribeParameterGroupsRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeParameterGroupsRequest) SetOwnerAccount(v string) *DescribeParameterGroupsRequest {
	s.OwnerAccount = &v
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
