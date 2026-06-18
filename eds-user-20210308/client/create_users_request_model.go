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
	// The date and time when the system automatically locks the convenience user\\"s account. The value must be in the `yyyy-MM-dd HH:mm:ss` format.
	//
	// example:
	//
	// 2025-11-28 00:00:00
	AutoLockTime *string `json:"AutoLockTime,omitempty" xml:"AutoLockTime,omitempty"`
	// The business channel.
	//
	// example:
	//
	// ENTERPRISE
	BusinessChannel *string `json:"BusinessChannel,omitempty" xml:"BusinessChannel,omitempty"`
	// Specifies whether to set the convenience user as a local administrator.
	//
	// example:
	//
	// true
	IsLocalAdmin *bool `json:"IsLocalAdmin,omitempty" xml:"IsLocalAdmin,omitempty"`
	// The initial password. If you do not specify this parameter, the system sends a password reset email to the convenience user\\"s email address.
	//
	// example:
	//
	// Test123****
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// By default, a convenience user\\"s password does not expire. You can use this parameter to specify a password validity period of 30 to 365 days. After the password expires, the user must reset it to log in again.
	//
	// > This feature is in invited preview. To use this feature, submit a ticket.
	//
	// example:
	//
	// 30
	PasswordExpireDays *string `json:"PasswordExpireDays,omitempty" xml:"PasswordExpireDays,omitempty"`
	// Details about the convenience users.
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
	// The email address of the convenience user. This email address is used for notifications, such as an alert when a cloud computer is assigned. You must specify either this parameter or the `Phone` parameter.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The user name. The user name must be 3 to 24 characters long and can contain lowercase letters, digits, and underscores (_).
	//
	// This parameter is required.
	//
	// example:
	//
	// alice
	EndUserId   *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	GroupIdList []*string `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty" type:"Repeated"`
	// The ID of the organization to which the convenience user belongs.
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
	// The password for the convenience user.
	//
	// > The password must be at least 10 characters long and contain characters from at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters (excluding spaces).
	//
	// example:
	//
	// Wuying1234
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// <props="china">
	//
	// The phone number of the convenience user. This phone number is used for notifications, such as a text message when a cloud computer is assigned. You must specify either this parameter or the `Email` parameter.
	//
	//
	//
	// <props="intl">
	//
	// Phone numbers are not supported on the international site.
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The display name of the convenience user.
	//
	// example:
	//
	// Bean
	RealNickName *string `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// A remark for the convenience user.
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
