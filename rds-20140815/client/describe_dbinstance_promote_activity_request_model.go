// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstancePromoteActivityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *DescribeDBInstancePromoteActivityRequest
	GetAliUid() *string
	SetDbInstanceName(v string) *DescribeDBInstancePromoteActivityRequest
	GetDbInstanceName() *string
	SetOwnerId(v int64) *DescribeDBInstancePromoteActivityRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *DescribeDBInstancePromoteActivityRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDBInstancePromoteActivityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDBInstancePromoteActivityRequest
	GetResourceOwnerId() *int64
}

type DescribeDBInstancePromoteActivityRequest struct {
	// The ID of the Alibaba Cloud account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22973492**********
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// 111
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeDBInstancePromoteActivityRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstancePromoteActivityRequest) GoString() string {
	return s.String()
}

func (s *DescribeDBInstancePromoteActivityRequest) GetAliUid() *string {
	return s.AliUid
}

func (s *DescribeDBInstancePromoteActivityRequest) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeDBInstancePromoteActivityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDBInstancePromoteActivityRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDBInstancePromoteActivityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDBInstancePromoteActivityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDBInstancePromoteActivityRequest) SetAliUid(v string) *DescribeDBInstancePromoteActivityRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityRequest) SetDbInstanceName(v string) *DescribeDBInstancePromoteActivityRequest {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityRequest) SetOwnerId(v int64) *DescribeDBInstancePromoteActivityRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityRequest) SetResourceGroupId(v string) *DescribeDBInstancePromoteActivityRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityRequest) SetResourceOwnerAccount(v string) *DescribeDBInstancePromoteActivityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityRequest) SetResourceOwnerId(v int64) *DescribeDBInstancePromoteActivityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDBInstancePromoteActivityRequest) Validate() error {
	return dara.Validate(s)
}
