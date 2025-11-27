// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCharacterSetNameRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEngine(v string) *DescribeCharacterSetNameRequest
	GetEngine() *string
	SetOwnerAccount(v string) *DescribeCharacterSetNameRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCharacterSetNameRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCharacterSetNameRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeCharacterSetNameRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeCharacterSetNameRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCharacterSetNameRequest
	GetResourceOwnerId() *int64
}

type DescribeCharacterSetNameRequest struct {
	// The type of the database engine. Valid values:
	//
	// 	- **mysql**
	//
	// 	- **mssql**
	//
	// 	- **PostgreSQL**
	//
	// 	- **MariaDB**
	//
	// This parameter is required.
	//
	// example:
	//
	// mysql
	Engine       *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the DescribeRegions operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID. You can call the DescribeDBInstanceAttribute operation to query the resource group ID.
	//
	// example:
	//
	// rg-acfmy*****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeCharacterSetNameRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCharacterSetNameRequest) GoString() string {
	return s.String()
}

func (s *DescribeCharacterSetNameRequest) GetEngine() *string {
	return s.Engine
}

func (s *DescribeCharacterSetNameRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCharacterSetNameRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCharacterSetNameRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCharacterSetNameRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCharacterSetNameRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCharacterSetNameRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCharacterSetNameRequest) SetEngine(v string) *DescribeCharacterSetNameRequest {
	s.Engine = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetOwnerAccount(v string) *DescribeCharacterSetNameRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetOwnerId(v int64) *DescribeCharacterSetNameRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetRegionId(v string) *DescribeCharacterSetNameRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetResourceGroupId(v string) *DescribeCharacterSetNameRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetResourceOwnerAccount(v string) *DescribeCharacterSetNameRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) SetResourceOwnerId(v int64) *DescribeCharacterSetNameRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCharacterSetNameRequest) Validate() error {
	return dara.Validate(s)
}
