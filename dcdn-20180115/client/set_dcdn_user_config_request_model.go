// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDcdnUserConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v string) *SetDcdnUserConfigRequest
	GetConfigs() *string
	SetFunctionId(v int32) *SetDcdnUserConfigRequest
	GetFunctionId() *int32
	SetOwnerAccount(v string) *SetDcdnUserConfigRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *SetDcdnUserConfigRequest
	GetOwnerId() *int64
	SetSecurityToken(v string) *SetDcdnUserConfigRequest
	GetSecurityToken() *string
}

type SetDcdnUserConfigRequest struct {
	// The configuration parameters of the feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// \\"argId\\":12,\\"argValue\\":\\"on\\"
	Configs *string `json:"Configs,omitempty" xml:"Configs,omitempty"`
	// The ID of the feature.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6
	FunctionId    *int32  `json:"FunctionId,omitempty" xml:"FunctionId,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s SetDcdnUserConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDcdnUserConfigRequest) GoString() string {
	return s.String()
}

func (s *SetDcdnUserConfigRequest) GetConfigs() *string {
	return s.Configs
}

func (s *SetDcdnUserConfigRequest) GetFunctionId() *int32 {
	return s.FunctionId
}

func (s *SetDcdnUserConfigRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *SetDcdnUserConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *SetDcdnUserConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *SetDcdnUserConfigRequest) SetConfigs(v string) *SetDcdnUserConfigRequest {
	s.Configs = &v
	return s
}

func (s *SetDcdnUserConfigRequest) SetFunctionId(v int32) *SetDcdnUserConfigRequest {
	s.FunctionId = &v
	return s
}

func (s *SetDcdnUserConfigRequest) SetOwnerAccount(v string) *SetDcdnUserConfigRequest {
	s.OwnerAccount = &v
	return s
}

func (s *SetDcdnUserConfigRequest) SetOwnerId(v int64) *SetDcdnUserConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *SetDcdnUserConfigRequest) SetSecurityToken(v string) *SetDcdnUserConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *SetDcdnUserConfigRequest) Validate() error {
	return dara.Validate(s)
}
