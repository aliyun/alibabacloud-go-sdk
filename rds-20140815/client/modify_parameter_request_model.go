// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParameterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyParameterRequest
	GetClientToken() *string
	SetDBInstanceId(v string) *ModifyParameterRequest
	GetDBInstanceId() *string
	SetForcerestart(v bool) *ModifyParameterRequest
	GetForcerestart() *bool
	SetOwnerAccount(v string) *ModifyParameterRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyParameterRequest
	GetOwnerId() *int64
	SetParameterGroupId(v string) *ModifyParameterRequest
	GetParameterGroupId() *string
	SetParameters(v string) *ModifyParameterRequest
	GetParameters() *string
	SetResourceOwnerAccount(v string) *ModifyParameterRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyParameterRequest
	GetResourceOwnerId() *int64
	SetSwitchTime(v string) *ModifyParameterRequest
	GetSwitchTime() *string
	SetSwitchTimeMode(v string) *ModifyParameterRequest
	GetSwitchTimeMode() *string
}

type ModifyParameterRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	DBInstanceId         *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	Forcerestart         *bool   `json:"Forcerestart,omitempty" xml:"Forcerestart,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ParameterGroupId     *string `json:"ParameterGroupId,omitempty" xml:"ParameterGroupId,omitempty"`
	Parameters           *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	SwitchTime           *string `json:"SwitchTime,omitempty" xml:"SwitchTime,omitempty"`
	SwitchTimeMode       *string `json:"SwitchTimeMode,omitempty" xml:"SwitchTimeMode,omitempty"`
}

func (s ModifyParameterRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyParameterRequest) GoString() string {
	return s.String()
}

func (s *ModifyParameterRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyParameterRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyParameterRequest) GetForcerestart() *bool {
	return s.Forcerestart
}

func (s *ModifyParameterRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyParameterRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyParameterRequest) GetParameterGroupId() *string {
	return s.ParameterGroupId
}

func (s *ModifyParameterRequest) GetParameters() *string {
	return s.Parameters
}

func (s *ModifyParameterRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyParameterRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyParameterRequest) GetSwitchTime() *string {
	return s.SwitchTime
}

func (s *ModifyParameterRequest) GetSwitchTimeMode() *string {
	return s.SwitchTimeMode
}

func (s *ModifyParameterRequest) SetClientToken(v string) *ModifyParameterRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyParameterRequest) SetDBInstanceId(v string) *ModifyParameterRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyParameterRequest) SetForcerestart(v bool) *ModifyParameterRequest {
	s.Forcerestart = &v
	return s
}

func (s *ModifyParameterRequest) SetOwnerAccount(v string) *ModifyParameterRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyParameterRequest) SetOwnerId(v int64) *ModifyParameterRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyParameterRequest) SetParameterGroupId(v string) *ModifyParameterRequest {
	s.ParameterGroupId = &v
	return s
}

func (s *ModifyParameterRequest) SetParameters(v string) *ModifyParameterRequest {
	s.Parameters = &v
	return s
}

func (s *ModifyParameterRequest) SetResourceOwnerAccount(v string) *ModifyParameterRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyParameterRequest) SetResourceOwnerId(v int64) *ModifyParameterRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyParameterRequest) SetSwitchTime(v string) *ModifyParameterRequest {
	s.SwitchTime = &v
	return s
}

func (s *ModifyParameterRequest) SetSwitchTimeMode(v string) *ModifyParameterRequest {
	s.SwitchTimeMode = &v
	return s
}

func (s *ModifyParameterRequest) Validate() error {
	return dara.Validate(s)
}
