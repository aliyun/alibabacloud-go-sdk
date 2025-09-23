// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigName(v string) *ModifyDBNodeConfigRequest
	GetConfigName() *string
	SetConfigValue(v string) *ModifyDBNodeConfigRequest
	GetConfigValue() *string
	SetDBClusterId(v string) *ModifyDBNodeConfigRequest
	GetDBClusterId() *string
	SetDBNodeId(v string) *ModifyDBNodeConfigRequest
	GetDBNodeId() *string
	SetOwnerAccount(v string) *ModifyDBNodeConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDBNodeConfigRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyDBNodeConfigRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDBNodeConfigRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *ModifyDBNodeConfigRequest
	GetSecurityToken() *string
}

type ModifyDBNodeConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// FailoverPriority
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// perfdb
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-**************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pi-****************
	DBNodeId             *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s ModifyDBNodeConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeConfigRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeConfigRequest) GetConfigName() *string {
	return s.ConfigName
}

func (s *ModifyDBNodeConfigRequest) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *ModifyDBNodeConfigRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyDBNodeConfigRequest) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *ModifyDBNodeConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDBNodeConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBNodeConfigRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDBNodeConfigRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDBNodeConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *ModifyDBNodeConfigRequest) SetConfigName(v string) *ModifyDBNodeConfigRequest {
	s.ConfigName = &v
	return s
}

func (s *ModifyDBNodeConfigRequest) SetConfigValue(v string) *ModifyDBNodeConfigRequest {
	s.ConfigValue = &v
	return s
}

func (s *ModifyDBNodeConfigRequest) SetDBClusterId(v string) *ModifyDBNodeConfigRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyDBNodeConfigRequest) SetDBNodeId(v string) *ModifyDBNodeConfigRequest {
	s.DBNodeId = &v
	return s
}

func (s *ModifyDBNodeConfigRequest) SetOwnerAccount(v string) *ModifyDBNodeConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDBNodeConfigRequest) SetOwnerId(v int64) *ModifyDBNodeConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBNodeConfigRequest) SetResourceOwnerAccount(v string) *ModifyDBNodeConfigRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDBNodeConfigRequest) SetResourceOwnerId(v int64) *ModifyDBNodeConfigRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDBNodeConfigRequest) SetSecurityToken(v string) *ModifyDBNodeConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *ModifyDBNodeConfigRequest) Validate() error {
	return dara.Validate(s)
}
