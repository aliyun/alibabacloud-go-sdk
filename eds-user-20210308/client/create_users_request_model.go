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
	// The date on which the convenience users are automatically locked.
	//
	// example:
	//
	// 2023-03-03
	AutoLockTime *string `json:"AutoLockTime,omitempty" xml:"AutoLockTime,omitempty"`
	IsLocalAdmin *bool   `json:"IsLocalAdmin,omitempty" xml:"IsLocalAdmin,omitempty"`
	// The initial password. If this parameter is left empty, an email for password reset is sent to the specified email address.
	//
	// example:
	//
	// Test123****
	Password           *string `json:"Password,omitempty" xml:"Password,omitempty"`
	PasswordExpireDays *string `json:"PasswordExpireDays,omitempty" xml:"PasswordExpireDays,omitempty"`
	// The information about the convenience user.
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
	// The email address of the convenience user. The email address is used to receive notifications about events such as desktop assignment. You must specify an email address or a mobile number to receive notifications.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The username of the convenience user. The name can contain lowercase letters, digits, and underscores (_), and must be 3 to 24 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test1
	EndUserId   *string   `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	GroupIdList []*string `json:"GroupIdList,omitempty" xml:"GroupIdList,omitempty" type:"Repeated"`
	// The organization to which the convenience user belongs.
	//
	// example:
	//
	// 1111****
	OrgId *string `json:"OrgId,omitempty" xml:"OrgId,omitempty"`
	// The type of the account ownership.
	//
	// Valid values:
	//
	// 	- CreateFromManager: administrator-activated
	//
	// 	- Normal: user-activated
	//
	// example:
	//
	// Normal
	OwnerType *string `json:"OwnerType,omitempty" xml:"OwnerType,omitempty"`
	// The user password.
	//
	// >  The password must be at least 10 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters (excluding spaces).
	//
	// example:
	//
	// password1
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Mobile numbers are not supported on the international site (alibabacloud.com).
	//
	// example:
	//
	// 1381111****
	Phone *string `json:"Phone,omitempty" xml:"Phone,omitempty"`
	// The display name of the end user.
	//
	// example:
	//
	// Bean
	RealNickName *string `json:"RealNickName,omitempty" xml:"RealNickName,omitempty"`
	// The remarks on the convenience user.
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
