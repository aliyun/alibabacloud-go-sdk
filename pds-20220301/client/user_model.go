// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUser interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *User
	GetAvatar() *string
	SetCreatedAt(v int64) *User
	GetCreatedAt() *int64
	SetCreator(v string) *User
	GetCreator() *string
	SetDefaultDriveId(v string) *User
	GetDefaultDriveId() *string
	SetDescription(v string) *User
	GetDescription() *string
	SetDomainId(v string) *User
	GetDomainId() *string
	SetEmail(v string) *User
	GetEmail() *string
	SetNickName(v string) *User
	GetNickName() *string
	SetPhone(v string) *User
	GetPhone() *string
	SetRole(v string) *User
	GetRole() *string
	SetStatus(v string) *User
	GetStatus() *string
	SetUpdatedAt(v int64) *User
	GetUpdatedAt() *int64
	SetUserData(v map[string]*string) *User
	GetUserData() map[string]*string
	SetUserId(v string) *User
	GetUserId() *string
	SetUserName(v string) *User
	GetUserName() *string
}

type User struct {
	// The profile picture.
	//
	// example:
	//
	// http://aa.com/1.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The time when the user was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1639762579768
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The user who created the user.
	//
	// example:
	//
	// user1
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The default drive ID.
	//
	// example:
	//
	// 1
	DefaultDriveId *string `json:"default_drive_id,omitempty" xml:"default_drive_id,omitempty"`
	// The description of the user.
	//
	// example:
	//
	// vipuser
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// a@aliyunpds.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// 001
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// The mobile number of the user.
	//
	// example:
	//
	// 13900001111
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The role of the user. Valid values:
	//
	// 	- superadmin
	//
	// 	- admin
	//
	// 	- user
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The status of the user. Valid values:
	//
	// 	- disabled: The user is prohibited from logon.
	//
	// 	- enabled: The user is in a normal state.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the user was modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1639762579768
	UpdatedAt *int64 `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The custom data of the user.
	UserData map[string]*string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The user ID.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// The name of the user.
	//
	// example:
	//
	// pds
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s User) String() string {
	return dara.Prettify(s)
}

func (s User) GoString() string {
	return s.String()
}

func (s *User) GetAvatar() *string {
	return s.Avatar
}

func (s *User) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *User) GetCreator() *string {
	return s.Creator
}

func (s *User) GetDefaultDriveId() *string {
	return s.DefaultDriveId
}

func (s *User) GetDescription() *string {
	return s.Description
}

func (s *User) GetDomainId() *string {
	return s.DomainId
}

func (s *User) GetEmail() *string {
	return s.Email
}

func (s *User) GetNickName() *string {
	return s.NickName
}

func (s *User) GetPhone() *string {
	return s.Phone
}

func (s *User) GetRole() *string {
	return s.Role
}

func (s *User) GetStatus() *string {
	return s.Status
}

func (s *User) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *User) GetUserData() map[string]*string {
	return s.UserData
}

func (s *User) GetUserId() *string {
	return s.UserId
}

func (s *User) GetUserName() *string {
	return s.UserName
}

func (s *User) SetAvatar(v string) *User {
	s.Avatar = &v
	return s
}

func (s *User) SetCreatedAt(v int64) *User {
	s.CreatedAt = &v
	return s
}

func (s *User) SetCreator(v string) *User {
	s.Creator = &v
	return s
}

func (s *User) SetDefaultDriveId(v string) *User {
	s.DefaultDriveId = &v
	return s
}

func (s *User) SetDescription(v string) *User {
	s.Description = &v
	return s
}

func (s *User) SetDomainId(v string) *User {
	s.DomainId = &v
	return s
}

func (s *User) SetEmail(v string) *User {
	s.Email = &v
	return s
}

func (s *User) SetNickName(v string) *User {
	s.NickName = &v
	return s
}

func (s *User) SetPhone(v string) *User {
	s.Phone = &v
	return s
}

func (s *User) SetRole(v string) *User {
	s.Role = &v
	return s
}

func (s *User) SetStatus(v string) *User {
	s.Status = &v
	return s
}

func (s *User) SetUpdatedAt(v int64) *User {
	s.UpdatedAt = &v
	return s
}

func (s *User) SetUserData(v map[string]*string) *User {
	s.UserData = v
	return s
}

func (s *User) SetUserId(v string) *User {
	s.UserId = &v
	return s
}

func (s *User) SetUserName(v string) *User {
	s.UserName = &v
	return s
}

func (s *User) Validate() error {
	return dara.Validate(s)
}
