// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoLockTime(v string) *CreateUsersRequest
	GetAutoLockTime() *string
	SetBusinessChannel(v string) *CreateUsersRequest
	GetBusinessChannel() *string
	SetIsLocalAdmin(v bool) *CreateUsersRequest
	GetIsLocalAdmin() *bool
	SetPassword(v string) *CreateUsersRequest
	GetPassword() *string
	SetPasswordExpireDays(v string) *CreateUsersRequest
	GetPasswordExpireDays() *string
	SetUsers(v []*CreateUsersRequestUsers) *CreateUsersRequest
	GetUsers() []*CreateUsersRequestUsers
}

type CreateUsersRequest struct {
	// The date when the account is automatically locked. Specify the date in the yyyy-MM-dd HH:mm:ss format.
	//
	// example:
	//
	// 2025-11-28 00:00:00
	AutoLockTime *string `json:"AutoLockTime,omitempty" xml:"AutoLockTime,omitempty"`
	// The channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// Specifies whether to set the convenience account as a local administrator.
	//
	// example:
	//
	// true
	IsLocalAdmin *bool `json:"IsLocalAdmin,omitempty" xml:"IsLocalAdmin,omitempty"`
	// The initial password. If you do not specify this parameter, a password reset email is sent to the user\\"s email address.
	//
	// example:
	//
	// Test123****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The validity period of the password for the user account. By default, the password never expires. You can set this parameter to a value from 30 to 365 to specify the validity period in days. After the password expires, the end user must change the password before logging on again.
	//
	// > This feature is in invitational preview. To use this feature, submit a ticket.
	//
	// example:
	//
	// 30
	PasswordExpireDays *string `json:"PasswordExpireDays,omitempty" xml:"PasswordExpireDays,omitempty"`
	// The convenience account information.
	//
	// This parameter is required.
	//
	// example:
	//
	// CreateUsers
	Users []*CreateUsersRequestUsers `json:"Users,omitempty" xml:"Users,omitempty" type:"Repeated"`
}

func (s CreateUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersRequest) GoString() string {
	return s.String()
}

func (s *CreateUsersRequest) GetAutoLockTime() *string {
	return s.AutoLockTime
}

func (s *CreateUsersRequest) GetBusinessChannel() *string {
	return s.BusinessChannel
}

func (s *CreateUsersRequest) GetIsLocalAdmin() *bool {
	return s.IsLocalAdmin
}

func (s *CreateUsersRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateUsersRequest) GetPasswordExpireDays() *string {
	return s.PasswordExpireDays
}

func (s *CreateUsersRequest) GetUsers() []*CreateUsersRequestUsers {
	return s.Users
}

func (s *CreateUsersRequest) SetAutoLockTime(v string) *CreateUsersRequest {
	s.AutoLockTime = &v
	return s
}

func (s *CreateUsersRequest) SetBusinessChannel(v string) *CreateUsersRequest {
	s.BusinessChannel = &v
	return s
}

func (s *CreateUsersRequest) SetIsLocalAdmin(v bool) *CreateUsersRequest {
	s.IsLocalAdmin = &v
	return s
}

func (s *CreateUsersRequest) SetPassword(v string) *CreateUsersRequest {
	s.Password = &v
	return s
}

func (s *CreateUsersRequest) SetPasswordExpireDays(v string) *CreateUsersRequest {
	s.PasswordExpireDays = &v
	return s
}

func (s *CreateUsersRequest) SetUsers(v []*CreateUsersRequestUsers) *CreateUsersRequest {
	s.Users = v
	return s
}

func (s *CreateUsersRequest) Validate() error {
	if s.Users != nil {
		for _, item := range s.Users {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateUsersRequestUsers struct {
	// The email address. The email address is used to receive relevant emails, such as notification emails for cloud computer allocation. You must specify either the email address or the phone number.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username. The username is custom and must be 3 to 24 characters in length. It can contain only lowercase letters, digits, and underscores.
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The IDs of the user groups to join.
	GroupIdList []*string `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty" type:"Repeated"`
	// The organization to which the user belongs.
	//
	// example:
	//
	// design
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The account activation type.
	//
	// example:
	//
	// Normal
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The user password.
	//
	// > The password must be at least 10 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters (excluding spaces).
	//
	// example:
	//
	// Wuying1234
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// <props="china">The phone number. The phone number is used to receive relevant text messages, such as notification text messages for cloud computer allocation. You must specify either the email address or the phone number.
	//
	// <props="intl">Phone numbers are not supported on the international site.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The display name of the user.
	//
	// example:
	//
	// Bean
	RealNickName *string `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The user remarks.
	//
	// example:
	//
	// remark1
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreateUsersRequestUsers) String() string {
	return dara.Prettify(s)
}

func (s CreateUsersRequestUsers) GoString() string {
	return s.String()
}

func (s *CreateUsersRequestUsers) GetEmail() *string {
	return s.Email
}

func (s *CreateUsersRequestUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *CreateUsersRequestUsers) GetGroupIdList() []*string {
	return s.GroupIdList
}

func (s *CreateUsersRequestUsers) GetOrgId() *string {
	return s.OrgId
}

func (s *CreateUsersRequestUsers) GetOwnerType() *string {
	return s.OwnerType
}

func (s *CreateUsersRequestUsers) GetPassword() *string {
	return s.Password
}

func (s *CreateUsersRequestUsers) GetPhone() *string {
	return s.Phone
}

func (s *CreateUsersRequestUsers) GetRealNickName() *string {
	return s.RealNickName
}

func (s *CreateUsersRequestUsers) GetRemark() *string {
	return s.Remark
}

func (s *CreateUsersRequestUsers) SetEmail(v string) *CreateUsersRequestUsers {
	s.Email = &v
	return s
}

func (s *CreateUsersRequestUsers) SetEndUserId(v string) *CreateUsersRequestUsers {
	s.EndUserId = &v
	return s
}

func (s *CreateUsersRequestUsers) SetGroupIdList(v []*string) *CreateUsersRequestUsers {
	s.GroupIdList = v
	return s
}

func (s *CreateUsersRequestUsers) SetOrgId(v string) *CreateUsersRequestUsers {
	s.OrgId = &v
	return s
}

func (s *CreateUsersRequestUsers) SetOwnerType(v string) *CreateUsersRequestUsers {
	s.OwnerType = &v
	return s
}

func (s *CreateUsersRequestUsers) SetPassword(v string) *CreateUsersRequestUsers {
	s.Password = &v
	return s
}

func (s *CreateUsersRequestUsers) SetPhone(v string) *CreateUsersRequestUsers {
	s.Phone = &v
	return s
}

func (s *CreateUsersRequestUsers) SetRealNickName(v string) *CreateUsersRequestUsers {
	s.RealNickName = &v
	return s
}

func (s *CreateUsersRequestUsers) SetRemark(v string) *CreateUsersRequestUsers {
	s.Remark = &v
	return s
}

func (s *CreateUsersRequestUsers) Validate() error {
	return dara.Validate(s)
}
