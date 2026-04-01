// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountMaskingPrivilegeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceName(v string) *DescribeAccountMaskingPrivilegeRequest
	GetDBInstanceName() *string
	SetDBName(v string) *DescribeAccountMaskingPrivilegeRequest
	GetDBName() *string
	SetOwnerId(v string) *DescribeAccountMaskingPrivilegeRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeAccountMaskingPrivilegeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAccountMaskingPrivilegeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccountMaskingPrivilegeRequest
	GetResourceOwnerId() *int64
	SetUserName(v string) *DescribeAccountMaskingPrivilegeRequest
	GetUserName() *string
}

type DescribeAccountMaskingPrivilegeRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// rm-t4n8t18o******6d5
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// myDB
	DBName  *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// ap-southeast-1
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// rds
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeAccountMaskingPrivilegeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountMaskingPrivilegeRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountMaskingPrivilegeRequest) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeAccountMaskingPrivilegeRequest) GetDBName() *string {
	return s.DBName
}

func (s *DescribeAccountMaskingPrivilegeRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeAccountMaskingPrivilegeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccountMaskingPrivilegeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAccountMaskingPrivilegeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAccountMaskingPrivilegeRequest) GetUserName() *string {
	return s.UserName
}

func (s *DescribeAccountMaskingPrivilegeRequest) SetDBInstanceName(v string) *DescribeAccountMaskingPrivilegeRequest {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeRequest) SetDBName(v string) *DescribeAccountMaskingPrivilegeRequest {
	s.DBName = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeRequest) SetOwnerId(v string) *DescribeAccountMaskingPrivilegeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeRequest) SetRegionId(v string) *DescribeAccountMaskingPrivilegeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeRequest) SetResourceOwnerAccount(v string) *DescribeAccountMaskingPrivilegeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeRequest) SetResourceOwnerId(v int64) *DescribeAccountMaskingPrivilegeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeRequest) SetUserName(v string) *DescribeAccountMaskingPrivilegeRequest {
	s.UserName = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeRequest) Validate() error {
	return dara.Validate(s)
}
