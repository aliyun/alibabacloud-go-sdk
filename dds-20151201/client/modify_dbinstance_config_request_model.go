// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigName(v string) *ModifyDBInstanceConfigRequest
	GetConfigName() *string
	SetConfigValue(v string) *ModifyDBInstanceConfigRequest
	GetConfigValue() *string
	SetDBInstanceId(v string) *ModifyDBInstanceConfigRequest
	GetDBInstanceId() *string
	SetOwnerAccount(v string) *ModifyDBInstanceConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBInstanceConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBInstanceConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBInstanceConfigRequest
	GetResourceOwnerId() *int64
}

type ModifyDBInstanceConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// manual_check
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dds-bp1ea17b41ab****
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDBInstanceConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceConfigRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *ModifyDBInstanceConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ModifyDBInstanceConfigRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBInstanceConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBInstanceConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBInstanceConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBInstanceConfigRequest) SetConfigName(v string) *ModifyDBInstanceConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetConfigValue(v string) *ModifyDBInstanceConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetDBInstanceId(v string) *ModifyDBInstanceConfigRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetOwnerAccount(v string) *ModifyDBInstanceConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetOwnerId(v int64) *ModifyDBInstanceConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetResourceOwnerAccount(v string) *ModifyDBInstanceConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) SetResourceOwnerId(v int64) *ModifyDBInstanceConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBInstanceConfigRequest) Validate() error {
	return dara.Validate(s)
}
