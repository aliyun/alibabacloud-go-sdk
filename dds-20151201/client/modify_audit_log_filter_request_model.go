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
	// The collection types of audit logs. Separate multiple collection types with commas (,).
	//
	// - **admin**: O&M and management operations.
	//
	// - **slow**: Slow queries.
	//
	// - **query**: Query operations.
	//
	// - **insert**: Insert operations.
	//
	// - **update**: Update operations.
	//
	// - **delete**: Delete operations.
	//
	// - **command**: Protocol commands, such as the aggregate method.
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
	// - **db**: shard node
	//
	// - **mongos**: mongos node
	//
	//
	//
	//
	//
	//
	// > Metric description
	//
	// > - This parameter applies only to sharded cluster instances. If this parameter is left empty, the default value db is used. You do not need to specify this parameter for replica set instances.
	//
	// example:
	//
	// db
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
