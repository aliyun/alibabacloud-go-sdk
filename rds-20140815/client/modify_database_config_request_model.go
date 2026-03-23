// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDatabaseConfigRequest
	GetDBInstanceId() *string
	SetDBName(v string) *ModifyDatabaseConfigRequest
	GetDBName() *string
	SetDatabasePropertyName(v string) *ModifyDatabaseConfigRequest
	GetDatabasePropertyName() *string
	SetDatabasePropertyValue(v string) *ModifyDatabaseConfigRequest
	GetDatabasePropertyValue() *string
	SetOwnerAccount(v string) *ModifyDatabaseConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDatabaseConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDatabaseConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDatabaseConfigRequest
	GetResourceOwnerId() *int64
}

type ModifyDatabaseConfigRequest struct {
	// This parameter is required.
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is required.
	DBName *string `json:"DBName,omitempty" xml:"DBName,omitempty"`
	// This parameter is required.
	DatabasePropertyName *string `json:"DatabasePropertyName,omitempty" xml:"DatabasePropertyName,omitempty"`
	// This parameter is required.
	DatabasePropertyValue *string `json:"DatabasePropertyValue,omitempty" xml:"DatabasePropertyValue,omitempty"`
	OwnerAccount          *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId               *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount  *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId       *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDatabaseConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDatabaseConfigRequest) GetDBName() *string {
	return s.DBName
}

func (s *ModifyDatabaseConfigRequest) GetDatabasePropertyName() *string {
	return s.DatabasePropertyName
}

func (s *ModifyDatabaseConfigRequest) GetDatabasePropertyValue() *string {
	return s.DatabasePropertyValue
}

func (s *ModifyDatabaseConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDatabaseConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDatabaseConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDatabaseConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDatabaseConfigRequest) SetDBInstanceId(v string) *ModifyDatabaseConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDatabaseConfigRequest) SetDBName(v string) *ModifyDatabaseConfigRequest {
	s.DBName = &v
	return s
}

func (s *ModifyDatabaseConfigRequest) SetDatabasePropertyName(v string) *ModifyDatabaseConfigRequest {
	s.DatabasePropertyName = &v
	return s
}

func (s *ModifyDatabaseConfigRequest) SetDatabasePropertyValue(v string) *ModifyDatabaseConfigRequest {
	s.DatabasePropertyValue = &v
	return s
}

func (s *ModifyDatabaseConfigRequest) SetOwnerAccount(v string) *ModifyDatabaseConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDatabaseConfigRequest) SetOwnerId(v int64) *ModifyDatabaseConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDatabaseConfigRequest) SetResourceOwnerAccount(v string) *ModifyDatabaseConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDatabaseConfigRequest) SetResourceOwnerId(v int64) *ModifyDatabaseConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDatabaseConfigRequest) Validate() error {
	return dara.Validate(s)
}
