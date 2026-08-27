// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *UpdateUserInfoRequest
	GetAvatar() *string
	SetLanguagePreference(v string) *UpdateUserInfoRequest
	GetLanguagePreference() *string
	SetName(v string) *UpdateUserInfoRequest
	GetName() *string
	SetOffering(v string) *UpdateUserInfoRequest
	GetOffering() *string
	SetProfileRoleInfo(v string) *UpdateUserInfoRequest
	GetProfileRoleInfo() *string
	SetSelfIntroduction(v string) *UpdateUserInfoRequest
	GetSelfIntroduction() *string
	SetTenantId(v string) *UpdateUserInfoRequest
	GetTenantId() *string
}

type UpdateUserInfoRequest struct {
	// The profile picture URL.
	//
	// example:
	//
	// https://example.com/avatar.png
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The language preference: zh-CN, en-US.
	//
	// example:
	//
	// string_value
	LanguagePreference *string `json:"languagePreference,omitempty" xml:"languagePreference,omitempty"`
	// The name.
	//
	// example:
	//
	// SampleName.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The user service description. Maximum length: 1000 characters.
	//
	// example:
	//
	// string_value
	Offering *string `json:"offering,omitempty" xml:"offering,omitempty"`
	// The user role description (used when profileRole is set to Others). Maximum length: 100 characters.
	//
	// example:
	//
	// string_value
	ProfileRoleInfo *string `json:"profileRoleInfo,omitempty" xml:"profileRoleInfo,omitempty"`
	// The user self-introduction. Maximum length: 1000 characters.
	//
	// example:
	//
	// string_value
	SelfIntroduction *string `json:"selfIntroduction,omitempty" xml:"selfIntroduction,omitempty"`
	// The effective tenant ID.
	//
	// example:
	//
	// 10000
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s UpdateUserInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserInfoRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserInfoRequest) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdateUserInfoRequest) GetLanguagePreference() *string {
	return s.LanguagePreference
}

func (s *UpdateUserInfoRequest) GetName() *string {
	return s.Name
}

func (s *UpdateUserInfoRequest) GetOffering() *string {
	return s.Offering
}

func (s *UpdateUserInfoRequest) GetProfileRoleInfo() *string {
	return s.ProfileRoleInfo
}

func (s *UpdateUserInfoRequest) GetSelfIntroduction() *string {
	return s.SelfIntroduction
}

func (s *UpdateUserInfoRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *UpdateUserInfoRequest) SetAvatar(v string) *UpdateUserInfoRequest {
	s.Avatar = &v
	return s
}

func (s *UpdateUserInfoRequest) SetLanguagePreference(v string) *UpdateUserInfoRequest {
	s.LanguagePreference = &v
	return s
}

func (s *UpdateUserInfoRequest) SetName(v string) *UpdateUserInfoRequest {
	s.Name = &v
	return s
}

func (s *UpdateUserInfoRequest) SetOffering(v string) *UpdateUserInfoRequest {
	s.Offering = &v
	return s
}

func (s *UpdateUserInfoRequest) SetProfileRoleInfo(v string) *UpdateUserInfoRequest {
	s.ProfileRoleInfo = &v
	return s
}

func (s *UpdateUserInfoRequest) SetSelfIntroduction(v string) *UpdateUserInfoRequest {
	s.SelfIntroduction = &v
	return s
}

func (s *UpdateUserInfoRequest) SetTenantId(v string) *UpdateUserInfoRequest {
	s.TenantId = &v
	return s
}

func (s *UpdateUserInfoRequest) Validate() error {
	return dara.Validate(s)
}
