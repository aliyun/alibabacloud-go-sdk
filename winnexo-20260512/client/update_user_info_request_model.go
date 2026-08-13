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
	// 用户头像 URL
	//
	// example:
	//
	// https://example.com/avatar.png
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 语言偏好: zh-CN, en-US
	//
	// example:
	//
	// string_value
	LanguagePreference *string `json:"languagePreference,omitempty" xml:"languagePreference,omitempty"`
	// 文件名
	//
	// example:
	//
	// 示例名称.pdf
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 用户服务描述，最多1000字符
	//
	// example:
	//
	// string_value
	Offering *string `json:"offering,omitempty" xml:"offering,omitempty"`
	// 用户角色描述（当profileRole为Others时使用），最多100字符
	//
	// example:
	//
	// string_value
	ProfileRoleInfo *string `json:"profileRoleInfo,omitempty" xml:"profileRoleInfo,omitempty"`
	// 用户自我介绍，最多1000字符
	//
	// example:
	//
	// string_value
	SelfIntroduction *string `json:"selfIntroduction,omitempty" xml:"selfIntroduction,omitempty"`
	// 租户ID，公共参数；winnexo-cli 通过 --tenant-id 显式传入
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
