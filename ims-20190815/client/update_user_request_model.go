// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewComments(v string) *UpdateUserRequest
	GetNewComments() *string
	SetNewDisplayName(v string) *UpdateUserRequest
	GetNewDisplayName() *string
	SetNewEmail(v string) *UpdateUserRequest
	GetNewEmail() *string
	SetNewMobilePhone(v string) *UpdateUserRequest
	GetNewMobilePhone() *string
	SetNewUserPrincipalName(v string) *UpdateUserRequest
	GetNewUserPrincipalName() *string
	SetUserId(v string) *UpdateUserRequest
	GetUserId() *string
	SetUserPrincipalName(v string) *UpdateUserRequest
	GetUserPrincipalName() *string
}

type UpdateUserRequest struct {
	// The new description of the RAM user.
	//
	// The description must be 1 to 128 characters in length.
	//
	// example:
	//
	// This is a cloud computing engineer.
	NewComments *string `json:"NewComments,omitempty" xml:"NewComments,omitempty"`
	// The new display name of the RAM user.
	//
	// The name must be 1 to 24 characters in length.
	//
	// example:
	//
	// new
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// The new email address of the RAM user.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// alice@example.com
	NewEmail *string `json:"NewEmail,omitempty" xml:"NewEmail,omitempty"`
	// The new mobile phone number of the RAM user.
	//
	// Format: \\<Country code>-\\<Mobile phone number>.
	//
	// > This parameter is valid only on the China site (aliyun.com).
	//
	// example:
	//
	// 86-1868888****
	NewMobilePhone *string `json:"NewMobilePhone,omitempty" xml:"NewMobilePhone,omitempty"`
	// The new logon name of the RAM user.
	//
	// The name is in the format of `<username>@<AccountAlias>.onaliyun.com`. `<username>` indicates the name of the RAM user. `<AccountAlias>.onaliyun.com` indicates the default domain name.
	//
	// The value of `NewUserPrincipalName` must be `1 to 128` characters in length and can contain letters, digits, periods (.), hyphens (-), and underscores (_). The value of `<username>` must be `1 to 64` characters in length.
	//
	// example:
	//
	// new@example.onaliyun.com
	NewUserPrincipalName *string `json:"NewUserPrincipalName,omitempty" xml:"NewUserPrincipalName,omitempty"`
	// The ID of the RAM user.
	//
	// > You must specify only one of the following parameters: `UserPrincipalName` and `UserId`.
	//
	// example:
	//
	// 20732900249392****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The logon name of the RAM user.
	//
	// > You must specify only one of the following parameters: `UserPrincipalName` and `UserId`.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) GetNewComments() *string {
	return s.NewComments
}

func (s *UpdateUserRequest) GetNewDisplayName() *string {
	return s.NewDisplayName
}

func (s *UpdateUserRequest) GetNewEmail() *string {
	return s.NewEmail
}

func (s *UpdateUserRequest) GetNewMobilePhone() *string {
	return s.NewMobilePhone
}

func (s *UpdateUserRequest) GetNewUserPrincipalName() *string {
	return s.NewUserPrincipalName
}

func (s *UpdateUserRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *UpdateUserRequest) SetNewComments(v string) *UpdateUserRequest {
	s.NewComments = &v
	return s
}

func (s *UpdateUserRequest) SetNewDisplayName(v string) *UpdateUserRequest {
	s.NewDisplayName = &v
	return s
}

func (s *UpdateUserRequest) SetNewEmail(v string) *UpdateUserRequest {
	s.NewEmail = &v
	return s
}

func (s *UpdateUserRequest) SetNewMobilePhone(v string) *UpdateUserRequest {
	s.NewMobilePhone = &v
	return s
}

func (s *UpdateUserRequest) SetNewUserPrincipalName(v string) *UpdateUserRequest {
	s.NewUserPrincipalName = &v
	return s
}

func (s *UpdateUserRequest) SetUserId(v string) *UpdateUserRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserRequest) SetUserPrincipalName(v string) *UpdateUserRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}
