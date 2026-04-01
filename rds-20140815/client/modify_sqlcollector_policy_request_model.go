// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySQLCollectorPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifySQLCollectorPolicyRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifySQLCollectorPolicyRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySQLCollectorPolicyRequest
	GetOwnerId() *int64
	SetResourceGroupId(v string) *ModifySQLCollectorPolicyRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ModifySQLCollectorPolicyRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySQLCollectorPolicyRequest
	GetResourceOwnerId() *int64
	SetSQLCollectorStatus(v string) *ModifySQLCollectorPolicyRequest
	GetSQLCollectorStatus() *string
}

type ModifySQLCollectorPolicyRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to query the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The resource group ID. You can call theDescribeDBInstanceAttribute operation to query the most recent region list.
	//
	// example:
	//
	// rg-acfmx**********
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable the SQL Explorer (SQL Audit) feature. Valid values:
	//
	// 	- **Enable**
	//
	// 	- **Disabled**
	//
	// This parameter is required.
	//
	// example:
	//
	// Enable
	SQLCollectorStatus *string `json:"SQLCollectorStatus,omitempty" xml:"SQLCollectorStatus,omitempty"`
}

func (s ModifySQLCollectorPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySQLCollectorPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifySQLCollectorPolicyRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifySQLCollectorPolicyRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySQLCollectorPolicyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySQLCollectorPolicyRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifySQLCollectorPolicyRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySQLCollectorPolicyRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySQLCollectorPolicyRequest) GetSQLCollectorStatus() *string {
	return s.SQLCollectorStatus
}

func (s *ModifySQLCollectorPolicyRequest) SetDBInstanceId(v string) *ModifySQLCollectorPolicyRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) SetOwnerAccount(v string) *ModifySQLCollectorPolicyRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) SetOwnerId(v int64) *ModifySQLCollectorPolicyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) SetResourceGroupId(v string) *ModifySQLCollectorPolicyRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) SetResourceOwnerAccount(v string) *ModifySQLCollectorPolicyRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) SetResourceOwnerId(v int64) *ModifySQLCollectorPolicyRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) SetSQLCollectorStatus(v string) *ModifySQLCollectorPolicyRequest {
	s.SQLCollectorStatus = &v
	return s
}

func (s *ModifySQLCollectorPolicyRequest) Validate() error {
	return dara.Validate(s)
}
