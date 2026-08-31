// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserByAccessKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserByAccessKeyResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetUserByAccessKeyResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserByAccessKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserByAccessKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserByAccessKeyResponseBody
	GetSuccess() *bool
	SetUserInfo(v *GetUserByAccessKeyResponseBodyUserInfo) *GetUserByAccessKeyResponseBody
	GetUserInfo() *GetUserByAccessKeyResponseBodyUserInfo
}

type GetUserByAccessKeyResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned for the request.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The user information.
	UserInfo *GetUserByAccessKeyResponseBodyUserInfo `json:"UserInfo,omitempty" xml:"UserInfo,omitempty" type:"Struct"`
}

func (s GetUserByAccessKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserByAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserByAccessKeyResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserByAccessKeyResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserByAccessKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserByAccessKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserByAccessKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserByAccessKeyResponseBody) GetUserInfo() *GetUserByAccessKeyResponseBodyUserInfo {
	return s.UserInfo
}

func (s *GetUserByAccessKeyResponseBody) SetCode(v string) *GetUserByAccessKeyResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserByAccessKeyResponseBody) SetHttpStatusCode(v int32) *GetUserByAccessKeyResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserByAccessKeyResponseBody) SetMessage(v string) *GetUserByAccessKeyResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserByAccessKeyResponseBody) SetRequestId(v string) *GetUserByAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserByAccessKeyResponseBody) SetSuccess(v bool) *GetUserByAccessKeyResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserByAccessKeyResponseBody) SetUserInfo(v *GetUserByAccessKeyResponseBodyUserInfo) *GetUserByAccessKeyResponseBody {
	s.UserInfo = v
	return s
}

func (s *GetUserByAccessKeyResponseBody) Validate() error {
	if s.UserInfo != nil {
		if err := s.UserInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserByAccessKeyResponseBodyUserInfo struct {
	// The display name of the user.
	//
	// example:
	//
	// zhangsan
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The Dataphin user ID.
	//
	// example:
	//
	// 30011210
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// xx@aliyun.com
	Mail *string `json:"Mail,omitempty" xml:"Mail,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// zhangsan
	NickName *string `json:"NickName,omitempty" xml:"NickName,omitempty"`
	// The account source type, such as ALIYUN_OAUTH2, PUBLICCLOUD_OAUTH2, BUC, or APSARA.
	//
	// example:
	//
	// ALIYUN_OAUTH2
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The source account ID of the user during SSO integration.
	//
	// example:
	//
	// 12345
	SourceUserId *string `json:"SourceUserId,omitempty" xml:"SourceUserId,omitempty"`
	// The tenant member status. Valid values:
	//
	// - NORMAL: Normal.
	//
	// - DEACTIVATE: Deactivated.
	//
	// - DELETE: Deleted.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tenant-level roles assigned to the AK/SK owner in the current tenant.
	TenantRoles []*GetUserByAccessKeyResponseBodyUserInfoTenantRoles `json:"TenantRoles,omitempty" xml:"TenantRoles,omitempty" type:"Repeated"`
	// The username of the account.
	//
	// example:
	//
	// zhangsan
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s GetUserByAccessKeyResponseBodyUserInfo) String() string {
	return dara.Prettify(s)
}

func (s GetUserByAccessKeyResponseBodyUserInfo) GoString() string {
	return s.String()
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) GetId() *string {
	return s.Id
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) GetMail() *string {
	return s.Mail
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) GetNickName() *string {
	return s.NickName
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) GetSourceType() *string {
	return s.SourceType
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) GetSourceUserId() *string {
	return s.SourceUserId
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) GetStatus() *string {
	return s.Status
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) GetTenantRoles() []*GetUserByAccessKeyResponseBodyUserInfoTenantRoles {
	return s.TenantRoles
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) GetUserName() *string {
	return s.UserName
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) SetDisplayName(v string) *GetUserByAccessKeyResponseBodyUserInfo {
	s.DisplayName = &v
	return s
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) SetId(v string) *GetUserByAccessKeyResponseBodyUserInfo {
	s.Id = &v
	return s
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) SetMail(v string) *GetUserByAccessKeyResponseBodyUserInfo {
	s.Mail = &v
	return s
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) SetNickName(v string) *GetUserByAccessKeyResponseBodyUserInfo {
	s.NickName = &v
	return s
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) SetSourceType(v string) *GetUserByAccessKeyResponseBodyUserInfo {
	s.SourceType = &v
	return s
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) SetSourceUserId(v string) *GetUserByAccessKeyResponseBodyUserInfo {
	s.SourceUserId = &v
	return s
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) SetStatus(v string) *GetUserByAccessKeyResponseBodyUserInfo {
	s.Status = &v
	return s
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) SetTenantRoles(v []*GetUserByAccessKeyResponseBodyUserInfoTenantRoles) *GetUserByAccessKeyResponseBodyUserInfo {
	s.TenantRoles = v
	return s
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) SetUserName(v string) *GetUserByAccessKeyResponseBodyUserInfo {
	s.UserName = &v
	return s
}

func (s *GetUserByAccessKeyResponseBodyUserInfo) Validate() error {
	if s.TenantRoles != nil {
		for _, item := range s.TenantRoles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUserByAccessKeyResponseBodyUserInfoTenantRoles struct {
	// The role identifier, such as SUPER_ADMIN or COMMON_USER.
	//
	// example:
	//
	// SUPER_ADMIN
	RoleKey *string `json:"RoleKey,omitempty" xml:"RoleKey,omitempty"`
	// The role name.
	//
	// example:
	//
	// Tenant Administrator
	RoleName *string `json:"RoleName,omitempty" xml:"RoleName,omitempty"`
}

func (s GetUserByAccessKeyResponseBodyUserInfoTenantRoles) String() string {
	return dara.Prettify(s)
}

func (s GetUserByAccessKeyResponseBodyUserInfoTenantRoles) GoString() string {
	return s.String()
}

func (s *GetUserByAccessKeyResponseBodyUserInfoTenantRoles) GetRoleKey() *string {
	return s.RoleKey
}

func (s *GetUserByAccessKeyResponseBodyUserInfoTenantRoles) GetRoleName() *string {
	return s.RoleName
}

func (s *GetUserByAccessKeyResponseBodyUserInfoTenantRoles) SetRoleKey(v string) *GetUserByAccessKeyResponseBodyUserInfoTenantRoles {
	s.RoleKey = &v
	return s
}

func (s *GetUserByAccessKeyResponseBodyUserInfoTenantRoles) SetRoleName(v string) *GetUserByAccessKeyResponseBodyUserInfoTenantRoles {
	s.RoleName = &v
	return s
}

func (s *GetUserByAccessKeyResponseBodyUserInfoTenantRoles) Validate() error {
	return dara.Validate(s)
}
