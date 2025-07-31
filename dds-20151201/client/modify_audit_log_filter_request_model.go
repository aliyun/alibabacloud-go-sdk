// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAuditLogFilterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyAuditLogFilterRequest
	GetDBInstanceId() *string
	SetFilter(v string) *ModifyAuditLogFilterRequest
	GetFilter() *string
	SetOwnerAccount(v string) *ModifyAuditLogFilterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyAuditLogFilterRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyAuditLogFilterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyAuditLogFilterRequest
	GetResourceOwnerId() *int64
	SetRoleType(v string) *ModifyAuditLogFilterRequest
	GetRoleType() *string
}

type ModifyAuditLogFilterRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dds-bp12c5b040dc****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The type of logs collected by the audit log feature of the instance. Separate multiple types with commas (,). Valid values:
	//
	// 	- **admin**: O\\&M and management operations
	//
	// 	- **slow**: slow query logs
	//
	// 	- **query**: query operations
	//
	// 	- **insert**: insert operations
	//
	// 	- **update**: update operations
	//
	// 	- **delete**: delete operations
	//
	// 	- **command**: protocol commands such as the aggregate method
	//
	// This parameter is required.
	//
	// example:
	//
	// insert,query,update,delete
	Filter               *string `json:"Filter,omitempty" xml:"Filter,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The role of the node in the instance. Valid values:
	//
	// 	- **primary**
	//
	// 	- **secondary**
	//
	// example:
	//
	// primary
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s ModifyAuditLogFilterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAuditLogFilterRequest) GoString() string {
	return s.String()
}

func (s *ModifyAuditLogFilterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyAuditLogFilterRequest) GetFilter() *string {
	return s.Filter
}

func (s *ModifyAuditLogFilterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyAuditLogFilterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyAuditLogFilterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyAuditLogFilterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyAuditLogFilterRequest) GetRoleType() *string {
	return s.RoleType
}

func (s *ModifyAuditLogFilterRequest) SetDBInstanceId(v string) *ModifyAuditLogFilterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetFilter(v string) *ModifyAuditLogFilterRequest {
	s.Filter = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetOwnerAccount(v string) *ModifyAuditLogFilterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetOwnerId(v int64) *ModifyAuditLogFilterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetResourceOwnerAccount(v string) *ModifyAuditLogFilterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetResourceOwnerId(v int64) *ModifyAuditLogFilterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) SetRoleType(v string) *ModifyAuditLogFilterRequest {
	s.RoleType = &v
	return s
}

func (s *ModifyAuditLogFilterRequest) Validate() error {
	return dara.Validate(s)
}
