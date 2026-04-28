// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBaseUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *BaseUserResponse
	GetAvatar() *string
	SetCreatedAt(v string) *BaseUserResponse
	GetCreatedAt() *string
	SetCreator(v string) *BaseUserResponse
	GetCreator() *string
	SetDefaultDriveId(v string) *BaseUserResponse
	GetDefaultDriveId() *string
	SetDefaultLocation(v string) *BaseUserResponse
	GetDefaultLocation() *string
	SetDenyChangePasswordBySelf(v bool) *BaseUserResponse
	GetDenyChangePasswordBySelf() *bool
	SetDescription(v string) *BaseUserResponse
	GetDescription() *string
	SetDomainId(v string) *BaseUserResponse
	GetDomainId() *string
	SetEmail(v string) *BaseUserResponse
	GetEmail() *string
	SetExpiredAt(v int64) *BaseUserResponse
	GetExpiredAt() *int64
	SetIsSync(v bool) *BaseUserResponse
	GetIsSync() *bool
	SetLastLoginTime(v int64) *BaseUserResponse
	GetLastLoginTime() *int64
	SetNeedChangePasswordNextLogin(v bool) *BaseUserResponse
	GetNeedChangePasswordNextLogin() *bool
	SetNickName(v string) *BaseUserResponse
	GetNickName() *string
	SetPathStatus(v string) *BaseUserResponse
	GetPathStatus() *string
	SetPermission(v map[string]*IDPermission) *BaseUserResponse
	GetPermission() map[string]*IDPermission
	SetPhone(v string) *BaseUserResponse
	GetPhone() *string
	SetPhoneRegion(v string) *BaseUserResponse
	GetPhoneRegion() *string
	SetRole(v string) *BaseUserResponse
	GetRole() *string
	SetStatus(v string) *BaseUserResponse
	GetStatus() *string
	SetUpdatedAt(v string) *BaseUserResponse
	GetUpdatedAt() *string
	SetUserData(v map[string]interface{}) *BaseUserResponse
	GetUserData() map[string]interface{}
	SetUserId(v string) *BaseUserResponse
	GetUserId() *string
	SetUserName(v string) *BaseUserResponse
	GetUserName() *string
}

type BaseUserResponse struct {
	// example:
	//
	// http://a.b.c/ccp.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// 1567407718386
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// example:
	//
	// system
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 123
	DefaultDriveId           *string `json:"default_drive_id,omitempty" xml:"default_drive_id,omitempty"`
	DefaultLocation          *string `json:"default_location,omitempty" xml:"default_location,omitempty"`
	DenyChangePasswordBySelf *bool   `json:"deny_change_password_by_self,omitempty" xml:"deny_change_password_by_self,omitempty"`
	// example:
	//
	// ccp team user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// hz999
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// example:
	//
	// 123@ccp.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 0
	ExpiredAt                   *int64 `json:"expired_at,omitempty" xml:"expired_at,omitempty"`
	IsSync                      *bool  `json:"is_sync,omitempty" xml:"is_sync,omitempty"`
	LastLoginTime               *int64 `json:"last_login_time,omitempty" xml:"last_login_time,omitempty"`
	NeedChangePasswordNextLogin *bool  `json:"need_change_password_next_login,omitempty" xml:"need_change_password_next_login,omitempty"`
	// example:
	//
	// abc
	NickName   *string                  `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	PathStatus *string                  `json:"path_status,omitempty" xml:"path_status,omitempty"`
	Permission map[string]*IDPermission `json:"permission,omitempty" xml:"permission,omitempty"`
	// example:
	//
	// 13700000000
	Phone       *string `json:"phone,omitempty" xml:"phone,omitempty"`
	PhoneRegion *string `json:"phone_region,omitempty" xml:"phone_region,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1567407718386
	UpdatedAt *string                `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UserData  map[string]interface{} `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// example:
	//
	// ccpuserid
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// example:
	//
	// name
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s BaseUserResponse) String() string {
	return dara.Prettify(s)
}

func (s BaseUserResponse) GoString() string {
	return s.String()
}

func (s *BaseUserResponse) GetAvatar() *string {
	return s.Avatar
}

func (s *BaseUserResponse) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *BaseUserResponse) GetCreator() *string {
	return s.Creator
}

func (s *BaseUserResponse) GetDefaultDriveId() *string {
	return s.DefaultDriveId
}

func (s *BaseUserResponse) GetDefaultLocation() *string {
	return s.DefaultLocation
}

func (s *BaseUserResponse) GetDenyChangePasswordBySelf() *bool {
	return s.DenyChangePasswordBySelf
}

func (s *BaseUserResponse) GetDescription() *string {
	return s.Description
}

func (s *BaseUserResponse) GetDomainId() *string {
	return s.DomainId
}

func (s *BaseUserResponse) GetEmail() *string {
	return s.Email
}

func (s *BaseUserResponse) GetExpiredAt() *int64 {
	return s.ExpiredAt
}

func (s *BaseUserResponse) GetIsSync() *bool {
	return s.IsSync
}

func (s *BaseUserResponse) GetLastLoginTime() *int64 {
	return s.LastLoginTime
}

func (s *BaseUserResponse) GetNeedChangePasswordNextLogin() *bool {
	return s.NeedChangePasswordNextLogin
}

func (s *BaseUserResponse) GetNickName() *string {
	return s.NickName
}

func (s *BaseUserResponse) GetPathStatus() *string {
	return s.PathStatus
}

func (s *BaseUserResponse) GetPermission() map[string]*IDPermission {
	return s.Permission
}

func (s *BaseUserResponse) GetPhone() *string {
	return s.Phone
}

func (s *BaseUserResponse) GetPhoneRegion() *string {
	return s.PhoneRegion
}

func (s *BaseUserResponse) GetRole() *string {
	return s.Role
}

func (s *BaseUserResponse) GetStatus() *string {
	return s.Status
}

func (s *BaseUserResponse) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *BaseUserResponse) GetUserData() map[string]interface{} {
	return s.UserData
}

func (s *BaseUserResponse) GetUserId() *string {
	return s.UserId
}

func (s *BaseUserResponse) GetUserName() *string {
	return s.UserName
}

func (s *BaseUserResponse) SetAvatar(v string) *BaseUserResponse {
	s.Avatar = &v
	return s
}

func (s *BaseUserResponse) SetCreatedAt(v string) *BaseUserResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseUserResponse) SetCreator(v string) *BaseUserResponse {
	s.Creator = &v
	return s
}

func (s *BaseUserResponse) SetDefaultDriveId(v string) *BaseUserResponse {
	s.DefaultDriveId = &v
	return s
}

func (s *BaseUserResponse) SetDefaultLocation(v string) *BaseUserResponse {
	s.DefaultLocation = &v
	return s
}

func (s *BaseUserResponse) SetDenyChangePasswordBySelf(v bool) *BaseUserResponse {
	s.DenyChangePasswordBySelf = &v
	return s
}

func (s *BaseUserResponse) SetDescription(v string) *BaseUserResponse {
	s.Description = &v
	return s
}

func (s *BaseUserResponse) SetDomainId(v string) *BaseUserResponse {
	s.DomainId = &v
	return s
}

func (s *BaseUserResponse) SetEmail(v string) *BaseUserResponse {
	s.Email = &v
	return s
}

func (s *BaseUserResponse) SetExpiredAt(v int64) *BaseUserResponse {
	s.ExpiredAt = &v
	return s
}

func (s *BaseUserResponse) SetIsSync(v bool) *BaseUserResponse {
	s.IsSync = &v
	return s
}

func (s *BaseUserResponse) SetLastLoginTime(v int64) *BaseUserResponse {
	s.LastLoginTime = &v
	return s
}

func (s *BaseUserResponse) SetNeedChangePasswordNextLogin(v bool) *BaseUserResponse {
	s.NeedChangePasswordNextLogin = &v
	return s
}

func (s *BaseUserResponse) SetNickName(v string) *BaseUserResponse {
	s.NickName = &v
	return s
}

func (s *BaseUserResponse) SetPathStatus(v string) *BaseUserResponse {
	s.PathStatus = &v
	return s
}

func (s *BaseUserResponse) SetPermission(v map[string]*IDPermission) *BaseUserResponse {
	s.Permission = v
	return s
}

func (s *BaseUserResponse) SetPhone(v string) *BaseUserResponse {
	s.Phone = &v
	return s
}

func (s *BaseUserResponse) SetPhoneRegion(v string) *BaseUserResponse {
	s.PhoneRegion = &v
	return s
}

func (s *BaseUserResponse) SetRole(v string) *BaseUserResponse {
	s.Role = &v
	return s
}

func (s *BaseUserResponse) SetStatus(v string) *BaseUserResponse {
	s.Status = &v
	return s
}

func (s *BaseUserResponse) SetUpdatedAt(v string) *BaseUserResponse {
	s.UpdatedAt = &v
	return s
}

func (s *BaseUserResponse) SetUserData(v map[string]interface{}) *BaseUserResponse {
	s.UserData = v
	return s
}

func (s *BaseUserResponse) SetUserId(v string) *BaseUserResponse {
	s.UserId = &v
	return s
}

func (s *BaseUserResponse) SetUserName(v string) *BaseUserResponse {
	s.UserName = &v
	return s
}

func (s *BaseUserResponse) Validate() error {
	return dara.Validate(s)
}
