// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchLSQLV3MySQLServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v int32) *SwitchLSQLV3MySQLServiceRequest
	GetActionType() *int32
	SetInstanceId(v string) *SwitchLSQLV3MySQLServiceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *SwitchLSQLV3MySQLServiceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SwitchLSQLV3MySQLServiceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *SwitchLSQLV3MySQLServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *SwitchLSQLV3MySQLServiceRequest
	GetResourceOwnerId() *int64
	SetSecurityToken(v string) *SwitchLSQLV3MySQLServiceRequest
	GetSecurityToken() *string
}

type SwitchLSQLV3MySQLServiceRequest struct {
	// The type of the operation. Valid value:
	//
	// 	- 1: enables the MySQL compatibility feature.
	//
	// 	- 0: disables the MySQL compatibility feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	ActionType *int32 `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ld-bp1o3y0yme2i2****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SecurityToken        *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SwitchLSQLV3MySQLServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s SwitchLSQLV3MySQLServiceRequest) GoString() string {
	return s.String()
}

func (s *SwitchLSQLV3MySQLServiceRequest) GetActionType() *int32 {
	return s.ActionType
}

func (s *SwitchLSQLV3MySQLServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SwitchLSQLV3MySQLServiceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SwitchLSQLV3MySQLServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SwitchLSQLV3MySQLServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *SwitchLSQLV3MySQLServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *SwitchLSQLV3MySQLServiceRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetActionType(v int32) *SwitchLSQLV3MySQLServiceRequest {
	s.ActionType = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetInstanceId(v string) *SwitchLSQLV3MySQLServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetOwnerAccount(v string) *SwitchLSQLV3MySQLServiceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetOwnerId(v int64) *SwitchLSQLV3MySQLServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetResourceOwnerAccount(v string) *SwitchLSQLV3MySQLServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetResourceOwnerId(v int64) *SwitchLSQLV3MySQLServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) SetSecurityToken(v string) *SwitchLSQLV3MySQLServiceRequest {
	s.SecurityToken = &v
	return s
}

func (s *SwitchLSQLV3MySQLServiceRequest) Validate() error {
	return dara.Validate(s)
}
