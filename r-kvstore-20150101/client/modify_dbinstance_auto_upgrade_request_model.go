// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceAutoUpgradeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBInstanceAutoUpgradeRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyDBInstanceAutoUpgradeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceAutoUpgradeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceAutoUpgradeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceAutoUpgradeRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyDBInstanceAutoUpgradeRequest
	GetSecurityToken() *string
	SetValue(v string) *ModifyDBInstanceAutoUpgradeRequest
	GetValue() *string
}

type ModifyDBInstanceAutoUpgradeRequest struct {
	// The instance ID. You can call the DescribeDBInstances operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// Specifies whether to enable automatic minor version update. Valid values:
	//
	// 	- **1**: enables automatic minor version update.
	//
	// 	- **0**: disables automatic minor version update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ModifyDBInstanceAutoUpgradeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceAutoUpgradeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAutoUpgradeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceAutoUpgradeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceAutoUpgradeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceAutoUpgradeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceAutoUpgradeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceAutoUpgradeRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyDBInstanceAutoUpgradeRequest) GetValue() *string {
	return s.Value
}

func (s *ModifyDBInstanceAutoUpgradeRequest) SetDBInstanceId(v string) *ModifyDBInstanceAutoUpgradeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeRequest) SetOwnerAccount(v string) *ModifyDBInstanceAutoUpgradeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeRequest) SetOwnerId(v int64) *ModifyDBInstanceAutoUpgradeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceAutoUpgradeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceAutoUpgradeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeRequest) SetSecurityToken(v string) *ModifyDBInstanceAutoUpgradeRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeRequest) SetValue(v string) *ModifyDBInstanceAutoUpgradeRequest {
	s.Value = &v
	return s
}

func (s *ModifyDBInstanceAutoUpgradeRequest) Validate() error {
	return dara.Validate(s)
}
