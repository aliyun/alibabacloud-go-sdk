// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateUserInfoResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateUserInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateUserInfoResponseBody
	GetRequestId() *string
	SetUser(v *UpdateUserInfoResponseBodyUser) *UpdateUserInfoResponseBody
	GetUser() *UpdateUserInfoResponseBodyUser
}

type UpdateUserInfoResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string                         `json:"requestId,omitempty" xml:"requestId,omitempty"`
	User      *UpdateUserInfoResponseBodyUser `json:"user,omitempty" xml:"user,omitempty" type:"Struct"`
}

func (s UpdateUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateUserInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserInfoResponseBody) GetUser() *UpdateUserInfoResponseBodyUser {
	return s.User
}

func (s *UpdateUserInfoResponseBody) SetCode(v string) *UpdateUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserInfoResponseBody) SetMessage(v string) *UpdateUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserInfoResponseBody) SetRequestId(v string) *UpdateUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserInfoResponseBody) SetUser(v *UpdateUserInfoResponseBodyUser) *UpdateUserInfoResponseBody {
	s.User = v
	return s
}

func (s *UpdateUserInfoResponseBody) Validate() error {
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateUserInfoResponseBodyUser struct {
	// 用户头像URL
	//
	// example:
	//
	// https://example.com/avatar.png
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 是否为超级管理员
	//
	// example:
	//
	// true
	IsAdmin *bool `json:"isAdmin,omitempty" xml:"isAdmin,omitempty"`
	// 当前登录租户是否为系统租户
	//
	// example:
	//
	// true
	IsSystemTenant *bool `json:"isSystemTenant,omitempty" xml:"isSystemTenant,omitempty"`
	// 用户语言偏好
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
	// 用户服务描述
	//
	// example:
	//
	// string_value
	Offering *string `json:"offering,omitempty" xml:"offering,omitempty"`
	// 用户服务解析结果（JSON格式）
	//
	// example:
	//
	// string_value
	ParsedOffering *string `json:"parsedOffering,omitempty" xml:"parsedOffering,omitempty"`
	// 用户角色
	//
	// example:
	//
	// string_value
	ProfileRole *string `json:"profileRole,omitempty" xml:"profileRole,omitempty"`
	// 用户角色描述
	//
	// example:
	//
	// string_value
	ProfileRoleInfo *string `json:"profileRoleInfo,omitempty" xml:"profileRoleInfo,omitempty"`
	// 用户自我介绍
	//
	// example:
	//
	// string_value
	SelfIntroduction *string `json:"selfIntroduction,omitempty" xml:"selfIntroduction,omitempty"`
	// 当前租户ID
	//
	// example:
	//
	// 10000
	TenantId   *int64                                      `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TenantList []*UpdateUserInfoResponseBodyUserTenantList `json:"tenantList,omitempty" xml:"tenantList,omitempty" type:"Repeated"`
	// 当前租户名称
	//
	// example:
	//
	// string_value
	TenantName *string `json:"tenantName,omitempty" xml:"tenantName,omitempty"`
	// 用户代码
	//
	// example:
	//
	// string_value
	UserCode *string `json:"userCode,omitempty" xml:"userCode,omitempty"`
	// 用户ID
	//
	// example:
	//
	// 1
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdateUserInfoResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserInfoResponseBodyUser) GoString() string {
	return s.String()
}

func (s *UpdateUserInfoResponseBodyUser) GetAvatar() *string {
	return s.Avatar
}

func (s *UpdateUserInfoResponseBodyUser) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *UpdateUserInfoResponseBodyUser) GetIsSystemTenant() *bool {
	return s.IsSystemTenant
}

func (s *UpdateUserInfoResponseBodyUser) GetLanguagePreference() *string {
	return s.LanguagePreference
}

func (s *UpdateUserInfoResponseBodyUser) GetName() *string {
	return s.Name
}

func (s *UpdateUserInfoResponseBodyUser) GetOffering() *string {
	return s.Offering
}

func (s *UpdateUserInfoResponseBodyUser) GetParsedOffering() *string {
	return s.ParsedOffering
}

func (s *UpdateUserInfoResponseBodyUser) GetProfileRole() *string {
	return s.ProfileRole
}

func (s *UpdateUserInfoResponseBodyUser) GetProfileRoleInfo() *string {
	return s.ProfileRoleInfo
}

func (s *UpdateUserInfoResponseBodyUser) GetSelfIntroduction() *string {
	return s.SelfIntroduction
}

func (s *UpdateUserInfoResponseBodyUser) GetTenantId() *int64 {
	return s.TenantId
}

func (s *UpdateUserInfoResponseBodyUser) GetTenantList() []*UpdateUserInfoResponseBodyUserTenantList {
	return s.TenantList
}

func (s *UpdateUserInfoResponseBodyUser) GetTenantName() *string {
	return s.TenantName
}

func (s *UpdateUserInfoResponseBodyUser) GetUserCode() *string {
	return s.UserCode
}

func (s *UpdateUserInfoResponseBodyUser) GetUserId() *int64 {
	return s.UserId
}

func (s *UpdateUserInfoResponseBodyUser) SetAvatar(v string) *UpdateUserInfoResponseBodyUser {
	s.Avatar = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetIsAdmin(v bool) *UpdateUserInfoResponseBodyUser {
	s.IsAdmin = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetIsSystemTenant(v bool) *UpdateUserInfoResponseBodyUser {
	s.IsSystemTenant = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetLanguagePreference(v string) *UpdateUserInfoResponseBodyUser {
	s.LanguagePreference = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetName(v string) *UpdateUserInfoResponseBodyUser {
	s.Name = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetOffering(v string) *UpdateUserInfoResponseBodyUser {
	s.Offering = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetParsedOffering(v string) *UpdateUserInfoResponseBodyUser {
	s.ParsedOffering = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetProfileRole(v string) *UpdateUserInfoResponseBodyUser {
	s.ProfileRole = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetProfileRoleInfo(v string) *UpdateUserInfoResponseBodyUser {
	s.ProfileRoleInfo = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetSelfIntroduction(v string) *UpdateUserInfoResponseBodyUser {
	s.SelfIntroduction = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetTenantId(v int64) *UpdateUserInfoResponseBodyUser {
	s.TenantId = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetTenantList(v []*UpdateUserInfoResponseBodyUserTenantList) *UpdateUserInfoResponseBodyUser {
	s.TenantList = v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetTenantName(v string) *UpdateUserInfoResponseBodyUser {
	s.TenantName = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetUserCode(v string) *UpdateUserInfoResponseBodyUser {
	s.UserCode = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) SetUserId(v int64) *UpdateUserInfoResponseBodyUser {
	s.UserId = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUser) Validate() error {
	if s.TenantList != nil {
		for _, item := range s.TenantList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateUserInfoResponseBodyUserTenantList struct {
	// 租户ID
	//
	// example:
	//
	// 10000
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// 租户名称
	//
	// example:
	//
	// string_value
	TenantName *string `json:"tenantName,omitempty" xml:"tenantName,omitempty"`
}

func (s UpdateUserInfoResponseBodyUserTenantList) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserInfoResponseBodyUserTenantList) GoString() string {
	return s.String()
}

func (s *UpdateUserInfoResponseBodyUserTenantList) GetTenantId() *int64 {
	return s.TenantId
}

func (s *UpdateUserInfoResponseBodyUserTenantList) GetTenantName() *string {
	return s.TenantName
}

func (s *UpdateUserInfoResponseBodyUserTenantList) SetTenantId(v int64) *UpdateUserInfoResponseBodyUserTenantList {
	s.TenantId = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUserTenantList) SetTenantName(v string) *UpdateUserInfoResponseBodyUserTenantList {
	s.TenantName = &v
	return s
}

func (s *UpdateUserInfoResponseBodyUserTenantList) Validate() error {
	return dara.Validate(s)
}
