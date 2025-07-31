// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAuditLogFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeAuditLogFilterRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *DescribeAuditLogFilterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAuditLogFilterRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeAuditLogFilterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAuditLogFilterRequest
	GetResourceOwnerId() *int64
	SetRoleType(v string) *DescribeAuditLogFilterRequest
	GetRoleType() *string
}

type DescribeAuditLogFilterRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp12c5b040dc****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role of the node in the instance. Valid values:
	//
	// 	- **mongos**: mongos node.
	//
	// 	- **db*	- : shard node.
	//
	// 	- **logic*	- : logical instance.
	//
	// example:
	//
	// logic
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DescribeAuditLogFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAuditLogFilterRequest) GoString() string {
	return s.String()
}

func (s *DescribeAuditLogFilterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeAuditLogFilterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAuditLogFilterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAuditLogFilterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAuditLogFilterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAuditLogFilterRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *DescribeAuditLogFilterRequest) SetDBInstanceId(v string) *DescribeAuditLogFilterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) SetOwnerAccount(v string) *DescribeAuditLogFilterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) SetOwnerId(v int64) *DescribeAuditLogFilterRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) SetResourceOwnerAccount(v string) *DescribeAuditLogFilterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) SetResourceOwnerId(v int64) *DescribeAuditLogFilterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) SetRoleType(v string) *DescribeAuditLogFilterRequest {
	s.RoleType = &v
	return s
}

func (s *DescribeAuditLogFilterRequest) Validate() error {
	return dara.Validate(s)
}
