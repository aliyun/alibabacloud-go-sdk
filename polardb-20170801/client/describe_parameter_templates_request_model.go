// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBType(v string) *DescribeParameterTemplatesRequest
	GetDBType() *string
	SetDBVersion(v string) *DescribeParameterTemplatesRequest
	GetDBVersion() *string
	SetOwnerAccount(v string) *DescribeParameterTemplatesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeParameterTemplatesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeParameterTemplatesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeParameterTemplatesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeParameterTemplatesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeParameterTemplatesRequest
	GetResourceOwnerId() *int64
}

type DescribeParameterTemplatesRequest struct {
	// The type of the database engine. Only **MySQL*	- is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The version of the database. Valid values:
	//
	// 	- **5.6**
	//
	// 	- **5.7**
	//
	// 	- **8.0**
	//
	// This parameter is required.
	//
	// example:
	//
	// 5.7
	DBVersion    *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// >  You can call the [DescribeRegions](https://help.aliyun.com/document_detail/98041.html) operation to query all regions that are available within your account, such as the region IDs.
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

func (s DescribeParameterTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeParameterTemplatesRequest) GetDBType() *string {
	return s.DBType
}

func (s *DescribeParameterTemplatesRequest) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeParameterTemplatesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeParameterTemplatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeParameterTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeParameterTemplatesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeParameterTemplatesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeParameterTemplatesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeParameterTemplatesRequest) SetDBType(v string) *DescribeParameterTemplatesRequest {
	s.DBType = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetDBVersion(v string) *DescribeParameterTemplatesRequest {
	s.DBVersion = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetRegionId(v string) *DescribeParameterTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceGroupId(v string) *DescribeParameterTemplatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerAccount(v string) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) SetResourceOwnerId(v int64) *DescribeParameterTemplatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeParameterTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
