// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvatar(v string) *CreateUserResponseBody
	GetAvatar() *string
	SetCreatedAt(v int64) *CreateUserResponseBody
	GetCreatedAt() *int64
	SetCreator(v string) *CreateUserResponseBody
	GetCreator() *string
	SetDefaultDriveId(v string) *CreateUserResponseBody
	GetDefaultDriveId() *string
	SetDescription(v string) *CreateUserResponseBody
	GetDescription() *string
	SetDomainId(v string) *CreateUserResponseBody
	GetDomainId() *string
	SetEmail(v string) *CreateUserResponseBody
	GetEmail() *string
	SetNickName(v string) *CreateUserResponseBody
	GetNickName() *string
	SetPhone(v string) *CreateUserResponseBody
	GetPhone() *string
	SetRole(v string) *CreateUserResponseBody
	GetRole() *string
	SetStatus(v string) *CreateUserResponseBody
	GetStatus() *string
	SetUpdatedAt(v int64) *CreateUserResponseBody
	GetUpdatedAt() *int64
	SetUserData(v map[string]interface{}) *CreateUserResponseBody
	GetUserData() map[string]interface{}
	SetUserId(v string) *CreateUserResponseBody
	GetUserId() *string
	SetUserName(v string) *CreateUserResponseBody
	GetUserName() *string
}

type CreateUserResponseBody struct {
	// The URL of the profile picture.
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
	// The ID of the default drive.
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
	// The email address.
	//
	// example:
	//
	// a@a.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// 001
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// The phone number.
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
	// admin
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The state of the user. Valid values:
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
	// The custom data.
	UserData map[string]interface{} `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The user ID.
	//
	// example:
	//
	// dingding_abc001
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// The username.
	//
	// example:
	//
	// pds
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) GetAvatar() *string {
	return s.Avatar
}

func (s *CreateUserResponseBody) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *CreateUserResponseBody) GetCreator() *string {
	return s.Creator
}

func (s *CreateUserResponseBody) GetDefaultDriveId() *string {
	return s.DefaultDriveId
}

func (s *CreateUserResponseBody) GetDescription() *string {
	return s.Description
}

func (s *CreateUserResponseBody) GetDomainId() *string {
	return s.DomainId
}

func (s *CreateUserResponseBody) GetEmail() *string {
	return s.Email
}

func (s *CreateUserResponseBody) GetNickName() *string {
	return s.NickName
}

func (s *CreateUserResponseBody) GetPhone() *string {
	return s.Phone
}

func (s *CreateUserResponseBody) GetRole() *string {
	return s.Role
}

func (s *CreateUserResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateUserResponseBody) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *CreateUserResponseBody) GetUserData() map[string]interface{} {
	return s.UserData
}

func (s *CreateUserResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *CreateUserResponseBody) GetUserName() *string {
	return s.UserName
}

func (s *CreateUserResponseBody) SetAvatar(v string) *CreateUserResponseBody {
	s.Avatar = &v
	return s
}

func (s *CreateUserResponseBody) SetCreatedAt(v int64) *CreateUserResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateUserResponseBody) SetCreator(v string) *CreateUserResponseBody {
	s.Creator = &v
	return s
}

func (s *CreateUserResponseBody) SetDefaultDriveId(v string) *CreateUserResponseBody {
	s.DefaultDriveId = &v
	return s
}

func (s *CreateUserResponseBody) SetDescription(v string) *CreateUserResponseBody {
	s.Description = &v
	return s
}

func (s *CreateUserResponseBody) SetDomainId(v string) *CreateUserResponseBody {
	s.DomainId = &v
	return s
}

func (s *CreateUserResponseBody) SetEmail(v string) *CreateUserResponseBody {
	s.Email = &v
	return s
}

func (s *CreateUserResponseBody) SetNickName(v string) *CreateUserResponseBody {
	s.NickName = &v
	return s
}

func (s *CreateUserResponseBody) SetPhone(v string) *CreateUserResponseBody {
	s.Phone = &v
	return s
}

func (s *CreateUserResponseBody) SetRole(v string) *CreateUserResponseBody {
	s.Role = &v
	return s
}

func (s *CreateUserResponseBody) SetStatus(v string) *CreateUserResponseBody {
	s.Status = &v
	return s
}

func (s *CreateUserResponseBody) SetUpdatedAt(v int64) *CreateUserResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *CreateUserResponseBody) SetUserData(v map[string]interface{}) *CreateUserResponseBody {
	s.UserData = v
	return s
}

func (s *CreateUserResponseBody) SetUserId(v string) *CreateUserResponseBody {
	s.UserId = &v
	return s
}

func (s *CreateUserResponseBody) SetUserName(v string) *CreateUserResponseBody {
	s.UserName = &v
	return s
}

func (s *CreateUserResponseBody) Validate() error {
	return dara.Validate(s)
}
