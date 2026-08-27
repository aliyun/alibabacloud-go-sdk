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
	// The profile picture URL.
	//
	// example:
	//
	// https://example.com/avatar.png
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The CRM type.
	//
	// example:
	//
	// standard
	CrmType *string `json:"crmType,omitempty" xml:"crmType,omitempty"`
	// Indicates whether the user is an enterprise administrator.
	//
	// example:
	//
	// true
	IsAdmin *bool `json:"isAdmin,omitempty" xml:"isAdmin,omitempty"`
	// Indicates whether the current logon tenant is the system tenant (tenantId=10000).
	//
	// example:
	//
	// true
	IsSystemTenant *bool `json:"isSystemTenant,omitempty" xml:"isSystemTenant,omitempty"`
	// The language preference.
	//
	// example:
	//
	// string_value
	LanguagePreference *string `json:"languagePreference,omitempty" xml:"languagePreference,omitempty"`
	// The status code description.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The username.
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
	// The parsed result of the user service (JSON format).
	//
	// example:
	//
	// string_value
	ParsedOffering *string `json:"parsedOffering,omitempty" xml:"parsedOffering,omitempty"`
	// The user role.
	//
	// example:
	//
	// string_value
	ProfileRole *string `json:"profileRole,omitempty" xml:"profileRole,omitempty"`
	// The personal profile.
	//
	// example:
	//
	// string_value
	ProfileRoleInfo *string `json:"profileRoleInfo,omitempty" xml:"profileRoleInfo,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The tenant list.
	TenantList []*GetUserInfoResponseBodyTenantList `json:"tenantList,omitempty" xml:"tenantList,omitempty" type:"Repeated"`
	// The current tenant name.
	//
	// example:
	//
	// string_value
	TenantName *string `json:"tenantName,omitempty" xml:"tenantName,omitempty"`
	// The tenant type. Valid values:
	//
	// - user: individual.
	//
	// - org: enterprise.
	//
	// - group: group.
	//
	// example:
	//
	// normal
	TenantType *string `json:"tenantType,omitempty" xml:"tenantType,omitempty"`
	// The user code.
	//
	// example:
	//
	// string_value
	UserCode *string `json:"userCode,omitempty" xml:"userCode,omitempty"`
	// The user ID.
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
	// The CRM type.
	//
	// example:
	//
	// standard
	CrmType *string `json:"crmType,omitempty" xml:"crmType,omitempty"`
	// The ID of the tenant to which the task belongs.
	//
	// example:
	//
	// 10000
	TenantId *int64 `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
	// The tenant name.
	//
	// example:
	//
	// string_value
	TenantName *string `json:"tenantName,omitempty" xml:"tenantName,omitempty"`
	// The tenant type. Valid values:
	//
	// - user: individual.
	//
	// - org: enterprise.
	//
	// - group: group.
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
