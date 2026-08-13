// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *GetUserInfoResponseBody
	GetAvatar() *string
	SetCode(v string) *GetUserInfoResponseBody
	GetCode() *string
	SetCrmType(v string) *GetUserInfoResponseBody
	GetCrmType() *string
	SetIsAdmin(v bool) *GetUserInfoResponseBody
	GetIsAdmin() *bool
	SetIsSystemTenant(v bool) *GetUserInfoResponseBody
	GetIsSystemTenant() *bool
	SetLanguagePreference(v string) *GetUserInfoResponseBody
	GetLanguagePreference() *string
	SetMessage(v string) *GetUserInfoResponseBody
	GetMessage() *string
	SetName(v string) *GetUserInfoResponseBody
	GetName() *string
	SetOffering(v string) *GetUserInfoResponseBody
	GetOffering() *string
	SetParsedOffering(v string) *GetUserInfoResponseBody
	GetParsedOffering() *string
	SetProfileRole(v string) *GetUserInfoResponseBody
	GetProfileRole() *string
	SetProfileRoleInfo(v string) *GetUserInfoResponseBody
	GetProfileRoleInfo() *string
	SetRequestId(v string) *GetUserInfoResponseBody
	GetRequestId() *string
	SetSelfIntroduction(v string) *GetUserInfoResponseBody
	GetSelfIntroduction() *string
	SetTenantId(v int64) *GetUserInfoResponseBody
	GetTenantId() *int64
	SetTenantList(v []*GetUserInfoResponseBodyTenantList) *GetUserInfoResponseBody
	GetTenantList() []*GetUserInfoResponseBodyTenantList
	SetTenantName(v string) *GetUserInfoResponseBody
	GetTenantName() *string
	SetTenantType(v string) *GetUserInfoResponseBody
	GetTenantType() *string
	SetUserCode(v string) *GetUserInfoResponseBody
	GetUserCode() *string
	SetUserId(v int64) *GetUserInfoResponseBody
	GetUserId() *int64
}

type GetUserInfoResponseBody struct {
	// 用户头像URL
	//
	// example:
	//
	// https://example.com/avatar.png
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// CRM 类型
	//
	// example:
	//
	// standard
	CrmType *string `json:"crmType,omitempty" xml:"crmType,omitempty"`
	// 是否为超级管理员
	//
	// example:
	//
	// true
	IsAdmin *bool `json:"isAdmin,omitempty" xml:"isAdmin,omitempty"`
	// 当前登录租户是否为系统租户（tenantId=10000）
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
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
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
	// 用户角色描述（当profileRole为Others时使用）
	//
	// example:
	//
	// string_value
	ProfileRoleInfo *string `json:"profileRoleInfo,omitempty" xml:"profileRoleInfo,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	TenantId   *int64                               `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	TenantList []*GetUserInfoResponseBodyTenantList `json:"tenantList,omitempty" xml:"tenantList,omitempty" type:"Repeated"`
	// 当前租户名称
	//
	// example:
	//
	// string_value
	TenantName *string `json:"tenantName,omitempty" xml:"tenantName,omitempty"`
	// 租户类型
	//
	// example:
	//
	// normal
	TenantType *string `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
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

func (s GetUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBody) GetAvatar() *string {
	return s.Avatar
}

func (s *GetUserInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserInfoResponseBody) GetCrmType() *string {
	return s.CrmType
}

func (s *GetUserInfoResponseBody) GetIsAdmin() *bool {
	return s.IsAdmin
}

func (s *GetUserInfoResponseBody) GetIsSystemTenant() *bool {
	return s.IsSystemTenant
}

func (s *GetUserInfoResponseBody) GetLanguagePreference() *string {
	return s.LanguagePreference
}

func (s *GetUserInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserInfoResponseBody) GetName() *string {
	return s.Name
}

func (s *GetUserInfoResponseBody) GetOffering() *string {
	return s.Offering
}

func (s *GetUserInfoResponseBody) GetParsedOffering() *string {
	return s.ParsedOffering
}

func (s *GetUserInfoResponseBody) GetProfileRole() *string {
	return s.ProfileRole
}

func (s *GetUserInfoResponseBody) GetProfileRoleInfo() *string {
	return s.ProfileRoleInfo
}

func (s *GetUserInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserInfoResponseBody) GetSelfIntroduction() *string {
	return s.SelfIntroduction
}

func (s *GetUserInfoResponseBody) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetUserInfoResponseBody) GetTenantList() []*GetUserInfoResponseBodyTenantList {
	return s.TenantList
}

func (s *GetUserInfoResponseBody) GetTenantName() *string {
	return s.TenantName
}

func (s *GetUserInfoResponseBody) GetTenantType() *string {
	return s.TenantType
}

func (s *GetUserInfoResponseBody) GetUserCode() *string {
	return s.UserCode
}

func (s *GetUserInfoResponseBody) GetUserId() *int64 {
	return s.UserId
}

func (s *GetUserInfoResponseBody) SetAvatar(v string) *GetUserInfoResponseBody {
	s.Avatar = &v
	return s
}

func (s *GetUserInfoResponseBody) SetCode(v string) *GetUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserInfoResponseBody) SetCrmType(v string) *GetUserInfoResponseBody {
	s.CrmType = &v
	return s
}

func (s *GetUserInfoResponseBody) SetIsAdmin(v bool) *GetUserInfoResponseBody {
	s.IsAdmin = &v
	return s
}

func (s *GetUserInfoResponseBody) SetIsSystemTenant(v bool) *GetUserInfoResponseBody {
	s.IsSystemTenant = &v
	return s
}

func (s *GetUserInfoResponseBody) SetLanguagePreference(v string) *GetUserInfoResponseBody {
	s.LanguagePreference = &v
	return s
}

func (s *GetUserInfoResponseBody) SetMessage(v string) *GetUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserInfoResponseBody) SetName(v string) *GetUserInfoResponseBody {
	s.Name = &v
	return s
}

func (s *GetUserInfoResponseBody) SetOffering(v string) *GetUserInfoResponseBody {
	s.Offering = &v
	return s
}

func (s *GetUserInfoResponseBody) SetParsedOffering(v string) *GetUserInfoResponseBody {
	s.ParsedOffering = &v
	return s
}

func (s *GetUserInfoResponseBody) SetProfileRole(v string) *GetUserInfoResponseBody {
	s.ProfileRole = &v
	return s
}

func (s *GetUserInfoResponseBody) SetProfileRoleInfo(v string) *GetUserInfoResponseBody {
	s.ProfileRoleInfo = &v
	return s
}

func (s *GetUserInfoResponseBody) SetRequestId(v string) *GetUserInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserInfoResponseBody) SetSelfIntroduction(v string) *GetUserInfoResponseBody {
	s.SelfIntroduction = &v
	return s
}

func (s *GetUserInfoResponseBody) SetTenantId(v int64) *GetUserInfoResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetUserInfoResponseBody) SetTenantList(v []*GetUserInfoResponseBodyTenantList) *GetUserInfoResponseBody {
	s.TenantList = v
	return s
}

func (s *GetUserInfoResponseBody) SetTenantName(v string) *GetUserInfoResponseBody {
	s.TenantName = &v
	return s
}

func (s *GetUserInfoResponseBody) SetTenantType(v string) *GetUserInfoResponseBody {
	s.TenantType = &v
	return s
}

func (s *GetUserInfoResponseBody) SetUserCode(v string) *GetUserInfoResponseBody {
	s.UserCode = &v
	return s
}

func (s *GetUserInfoResponseBody) SetUserId(v int64) *GetUserInfoResponseBody {
	s.UserId = &v
	return s
}

func (s *GetUserInfoResponseBody) Validate() error {
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

type GetUserInfoResponseBodyTenantList struct {
	// CRM 类型
	//
	// example:
	//
	// standard
	CrmType *string `json:"crmType,omitempty" xml:"crmType,omitempty"`
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
	// 租户类型
	//
	// example:
	//
	// normal
	TenantType *string `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
}

func (s GetUserInfoResponseBodyTenantList) String() string {
	return dara.Prettify(s)
}

func (s GetUserInfoResponseBodyTenantList) GoString() string {
	return s.String()
}

func (s *GetUserInfoResponseBodyTenantList) GetCrmType() *string {
	return s.CrmType
}

func (s *GetUserInfoResponseBodyTenantList) GetTenantId() *int64 {
	return s.TenantId
}

func (s *GetUserInfoResponseBodyTenantList) GetTenantName() *string {
	return s.TenantName
}

func (s *GetUserInfoResponseBodyTenantList) GetTenantType() *string {
	return s.TenantType
}

func (s *GetUserInfoResponseBodyTenantList) SetCrmType(v string) *GetUserInfoResponseBodyTenantList {
	s.CrmType = &v
	return s
}

func (s *GetUserInfoResponseBodyTenantList) SetTenantId(v int64) *GetUserInfoResponseBodyTenantList {
	s.TenantId = &v
	return s
}

func (s *GetUserInfoResponseBodyTenantList) SetTenantName(v string) *GetUserInfoResponseBodyTenantList {
	s.TenantName = &v
	return s
}

func (s *GetUserInfoResponseBodyTenantList) SetTenantType(v string) *GetUserInfoResponseBodyTenantList {
	s.TenantType = &v
	return s
}

func (s *GetUserInfoResponseBodyTenantList) Validate() error {
	return dara.Validate(s)
}
