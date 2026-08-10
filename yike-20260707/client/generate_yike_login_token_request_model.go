// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateYikeLoginTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoCreateProduction(v string) *GenerateYikeLoginTokenRequest
	GetAutoCreateProduction() *string
	SetExpires(v string) *GenerateYikeLoginTokenRequest
	GetExpires() *string
	SetNickName(v string) *GenerateYikeLoginTokenRequest
	GetNickName() *string
	SetProductionAuth(v string) *GenerateYikeLoginTokenRequest
	GetProductionAuth() *string
	SetSubUserCredit(v string) *GenerateYikeLoginTokenRequest
	GetSubUserCredit() *string
	SetTenant(v string) *GenerateYikeLoginTokenRequest
	GetTenant() *string
	SetUserName(v string) *GenerateYikeLoginTokenRequest
	GetUserName() *string
	SetWorkspaceId(v string) *GenerateYikeLoginTokenRequest
	GetWorkspaceId() *string
}

type GenerateYikeLoginTokenRequest struct {
	// Specifies whether automatic creation of a project is enabled. Default value: false.
	//
	// example:
	//
	// false
	AutoCreateProduction *string `json:"AutoCreateProduction,omitempty" xml:"AutoCreateProduction,omitempty"`
	// The token expiration time, in seconds. Default value: 30 days.
	//
	// example:
	//
	// 0
	Expires *string `json:"Expires,omitempty" xml:"Expires,omitempty"`
	// The nickname.
	//
	// - Format check: The maximum length is 50 characters.
	//
	// - Special format validation: Chinese characters, English characters, digits, _ \\ / () ] [
	//
	// example:
	//
	// 冯凯
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The role of the user in the project.
	//
	// example:
	//
	// common
	ProductionAuth *string `json:"ProductionAuth,omitempty" xml:"ProductionAuth,omitempty"`
	// The default credits granted to the user.
	//
	// example:
	//
	// 1000
	SubUserCredit *string `json:"SubUserCredit,omitempty" xml:"SubUserCredit,omitempty"`
	// The tenant identifier.
	//
	// example:
	//
	// wanyou
	Tenant *string `json:"Tenant,omitempty" xml:"Tenant,omitempty"`
	// The username.
	//
	// example:
	//
	// userxxx
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 581236
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GenerateYikeLoginTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateYikeLoginTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateYikeLoginTokenRequest) GetAutoCreateProduction() *string {
	return s.AutoCreateProduction
}

func (s *GenerateYikeLoginTokenRequest) GetExpires() *string {
	return s.Expires
}

func (s *GenerateYikeLoginTokenRequest) GetNickName() *string {
	return s.NickName
}

func (s *GenerateYikeLoginTokenRequest) GetProductionAuth() *string {
	return s.ProductionAuth
}

func (s *GenerateYikeLoginTokenRequest) GetSubUserCredit() *string {
	return s.SubUserCredit
}

func (s *GenerateYikeLoginTokenRequest) GetTenant() *string {
	return s.Tenant
}

func (s *GenerateYikeLoginTokenRequest) GetUserName() *string {
	return s.UserName
}

func (s *GenerateYikeLoginTokenRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GenerateYikeLoginTokenRequest) SetAutoCreateProduction(v string) *GenerateYikeLoginTokenRequest {
	s.AutoCreateProduction = &v
	return s
}

func (s *GenerateYikeLoginTokenRequest) SetExpires(v string) *GenerateYikeLoginTokenRequest {
	s.Expires = &v
	return s
}

func (s *GenerateYikeLoginTokenRequest) SetNickName(v string) *GenerateYikeLoginTokenRequest {
	s.NickName = &v
	return s
}

func (s *GenerateYikeLoginTokenRequest) SetProductionAuth(v string) *GenerateYikeLoginTokenRequest {
	s.ProductionAuth = &v
	return s
}

func (s *GenerateYikeLoginTokenRequest) SetSubUserCredit(v string) *GenerateYikeLoginTokenRequest {
	s.SubUserCredit = &v
	return s
}

func (s *GenerateYikeLoginTokenRequest) SetTenant(v string) *GenerateYikeLoginTokenRequest {
	s.Tenant = &v
	return s
}

func (s *GenerateYikeLoginTokenRequest) SetUserName(v string) *GenerateYikeLoginTokenRequest {
	s.UserName = &v
	return s
}

func (s *GenerateYikeLoginTokenRequest) SetWorkspaceId(v string) *GenerateYikeLoginTokenRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GenerateYikeLoginTokenRequest) Validate() error {
	return dara.Validate(s)
}
