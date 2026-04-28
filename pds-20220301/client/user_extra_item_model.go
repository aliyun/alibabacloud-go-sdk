// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUserExtraItem interface {
	dara.Model
	String() string
	GoString() string
	SetAccount(v []*AccountLinkInfo) *UserExtraItem
	GetAccount() []*AccountLinkInfo
	SetAvatar(v string) *UserExtraItem
	GetAvatar() *string
	SetCreatedAt(v string) *UserExtraItem
	GetCreatedAt() *string
	SetCreator(v string) *UserExtraItem
	GetCreator() *string
	SetDefaultDrive(v *BaseDriveResponse) *UserExtraItem
	GetDefaultDrive() *BaseDriveResponse
	SetDefaultDriveId(v string) *UserExtraItem
	GetDefaultDriveId() *string
	SetDefaultLocation(v string) *UserExtraItem
	GetDefaultLocation() *string
	SetDenyChangePasswordBySelf(v bool) *UserExtraItem
	GetDenyChangePasswordBySelf() *bool
	SetDescription(v string) *UserExtraItem
	GetDescription() *string
	SetDomainId(v string) *UserExtraItem
	GetDomainId() *string
	SetEmail(v string) *UserExtraItem
	GetEmail() *string
	SetExpiredAt(v int64) *UserExtraItem
	GetExpiredAt() *int64
	SetIsSync(v bool) *UserExtraItem
	GetIsSync() *bool
	SetLastLoginTime(v int64) *UserExtraItem
	GetLastLoginTime() *int64
	SetNeedChangePasswordNextLogin(v bool) *UserExtraItem
	GetNeedChangePasswordNextLogin() *bool
	SetNickName(v string) *UserExtraItem
	GetNickName() *string
	SetParentGroup(v []*BaseDriveResponse) *UserExtraItem
	GetParentGroup() []*BaseDriveResponse
	SetPathStatus(v string) *UserExtraItem
	GetPathStatus() *string
	SetPermission(v map[string]*IDPermission) *UserExtraItem
	GetPermission() map[string]*IDPermission
	SetPhone(v string) *UserExtraItem
	GetPhone() *string
	SetPhoneRegion(v string) *UserExtraItem
	GetPhoneRegion() *string
	SetRole(v string) *UserExtraItem
	GetRole() *string
	SetStatus(v string) *UserExtraItem
	GetStatus() *string
	SetUpdatedAt(v string) *UserExtraItem
	GetUpdatedAt() *string
	SetUserData(v map[string]interface{}) *UserExtraItem
	GetUserData() map[string]interface{}
	SetUserId(v string) *UserExtraItem
	GetUserId() *string
	SetUserName(v string) *UserExtraItem
	GetUserName() *string
}

type UserExtraItem struct {
	Account []*AccountLinkInfo `json:"account,omitempty" xml:"account,omitempty" type:"Repeated"`
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
	// if can be null:
	// true
	DefaultDrive *BaseDriveResponse `json:"default_drive,omitempty" xml:"default_drive,omitempty"`
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
	NickName    *string                  `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	ParentGroup []*BaseDriveResponse     `json:"parent_group,omitempty" xml:"parent_group,omitempty" type:"Repeated"`
	PathStatus  *string                  `json:"path_status,omitempty" xml:"path_status,omitempty"`
	Permission  map[string]*IDPermission `json:"permission,omitempty" xml:"permission,omitempty"`
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

func (s UserExtraItem) String() string {
	return dara.Prettify(s)
}

func (s UserExtraItem) GoString() string {
	return s.String()
}

func (s *UserExtraItem) GetAccount() []*AccountLinkInfo {
	return s.Account
}

func (s *UserExtraItem) GetAvatar() *string {
	return s.Avatar
}

func (s *UserExtraItem) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UserExtraItem) GetCreator() *string {
	return s.Creator
}

func (s *UserExtraItem) GetDefaultDrive() *BaseDriveResponse {
	return s.DefaultDrive
}

func (s *UserExtraItem) GetDefaultDriveId() *string {
	return s.DefaultDriveId
}

func (s *UserExtraItem) GetDefaultLocation() *string {
	return s.DefaultLocation
}

func (s *UserExtraItem) GetDenyChangePasswordBySelf() *bool {
	return s.DenyChangePasswordBySelf
}

func (s *UserExtraItem) GetDescription() *string {
	return s.Description
}

func (s *UserExtraItem) GetDomainId() *string {
	return s.DomainId
}

func (s *UserExtraItem) GetEmail() *string {
	return s.Email
}

func (s *UserExtraItem) GetExpiredAt() *int64 {
	return s.ExpiredAt
}

func (s *UserExtraItem) GetIsSync() *bool {
	return s.IsSync
}

func (s *UserExtraItem) GetLastLoginTime() *int64 {
	return s.LastLoginTime
}

func (s *UserExtraItem) GetNeedChangePasswordNextLogin() *bool {
	return s.NeedChangePasswordNextLogin
}

func (s *UserExtraItem) GetNickName() *string {
	return s.NickName
}

func (s *UserExtraItem) GetParentGroup() []*BaseDriveResponse {
	return s.ParentGroup
}

func (s *UserExtraItem) GetPathStatus() *string {
	return s.PathStatus
}

func (s *UserExtraItem) GetPermission() map[string]*IDPermission {
	return s.Permission
}

func (s *UserExtraItem) GetPhone() *string {
	return s.Phone
}

func (s *UserExtraItem) GetPhoneRegion() *string {
	return s.PhoneRegion
}

func (s *UserExtraItem) GetRole() *string {
	return s.Role
}

func (s *UserExtraItem) GetStatus() *string {
	return s.Status
}

func (s *UserExtraItem) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *UserExtraItem) GetUserData() map[string]interface{} {
	return s.UserData
}

func (s *UserExtraItem) GetUserId() *string {
	return s.UserId
}

func (s *UserExtraItem) GetUserName() *string {
	return s.UserName
}

func (s *UserExtraItem) SetAccount(v []*AccountLinkInfo) *UserExtraItem {
	s.Account = v
	return s
}

func (s *UserExtraItem) SetAvatar(v string) *UserExtraItem {
	s.Avatar = &v
	return s
}

func (s *UserExtraItem) SetCreatedAt(v string) *UserExtraItem {
	s.CreatedAt = &v
	return s
}

func (s *UserExtraItem) SetCreator(v string) *UserExtraItem {
	s.Creator = &v
	return s
}

func (s *UserExtraItem) SetDefaultDrive(v *BaseDriveResponse) *UserExtraItem {
	s.DefaultDrive = v
	return s
}

func (s *UserExtraItem) SetDefaultDriveId(v string) *UserExtraItem {
	s.DefaultDriveId = &v
	return s
}

func (s *UserExtraItem) SetDefaultLocation(v string) *UserExtraItem {
	s.DefaultLocation = &v
	return s
}

func (s *UserExtraItem) SetDenyChangePasswordBySelf(v bool) *UserExtraItem {
	s.DenyChangePasswordBySelf = &v
	return s
}

func (s *UserExtraItem) SetDescription(v string) *UserExtraItem {
	s.Description = &v
	return s
}

func (s *UserExtraItem) SetDomainId(v string) *UserExtraItem {
	s.DomainId = &v
	return s
}

func (s *UserExtraItem) SetEmail(v string) *UserExtraItem {
	s.Email = &v
	return s
}

func (s *UserExtraItem) SetExpiredAt(v int64) *UserExtraItem {
	s.ExpiredAt = &v
	return s
}

func (s *UserExtraItem) SetIsSync(v bool) *UserExtraItem {
	s.IsSync = &v
	return s
}

func (s *UserExtraItem) SetLastLoginTime(v int64) *UserExtraItem {
	s.LastLoginTime = &v
	return s
}

func (s *UserExtraItem) SetNeedChangePasswordNextLogin(v bool) *UserExtraItem {
	s.NeedChangePasswordNextLogin = &v
	return s
}

func (s *UserExtraItem) SetNickName(v string) *UserExtraItem {
	s.NickName = &v
	return s
}

func (s *UserExtraItem) SetParentGroup(v []*BaseDriveResponse) *UserExtraItem {
	s.ParentGroup = v
	return s
}

func (s *UserExtraItem) SetPathStatus(v string) *UserExtraItem {
	s.PathStatus = &v
	return s
}

func (s *UserExtraItem) SetPermission(v map[string]*IDPermission) *UserExtraItem {
	s.Permission = v
	return s
}

func (s *UserExtraItem) SetPhone(v string) *UserExtraItem {
	s.Phone = &v
	return s
}

func (s *UserExtraItem) SetPhoneRegion(v string) *UserExtraItem {
	s.PhoneRegion = &v
	return s
}

func (s *UserExtraItem) SetRole(v string) *UserExtraItem {
	s.Role = &v
	return s
}

func (s *UserExtraItem) SetStatus(v string) *UserExtraItem {
	s.Status = &v
	return s
}

func (s *UserExtraItem) SetUpdatedAt(v string) *UserExtraItem {
	s.UpdatedAt = &v
	return s
}

func (s *UserExtraItem) SetUserData(v map[string]interface{}) *UserExtraItem {
	s.UserData = v
	return s
}

func (s *UserExtraItem) SetUserId(v string) *UserExtraItem {
	s.UserId = &v
	return s
}

func (s *UserExtraItem) SetUserName(v string) *UserExtraItem {
	s.UserName = &v
	return s
}

func (s *UserExtraItem) Validate() error {
	if s.Account != nil {
		for _, item := range s.Account {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DefaultDrive != nil {
		if err := s.DefaultDrive.Validate(); err != nil {
			return err
		}
	}
	if s.ParentGroup != nil {
		for _, item := range s.ParentGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
