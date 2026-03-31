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
	SetNewUserName(v string) *UpdateUserRequest
	GetNewUserName() *string
	SetUserName(v string) *UpdateUserRequest
	GetUserName() *string
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
	// The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// xiaoq****
	NewDisplayName *string `json:"NewDisplayName,omitempty" xml:"NewDisplayName,omitempty"`
	// The new email address of the RAM user.
	//
	// >  This parameter applies only to the China site (aliyun.com).
	//
	// example:
	//
	// xiaoq****@example.com
	NewEmail *string `json:"NewEmail,omitempty" xml:"NewEmail,omitempty"`
	// The new mobile phone number of the RAM user.
	//
	// Format: \\<Country code>-\\<Mobile phone number>.
	//
	// >  This parameter applies only to the China site (aliyun.com).
	//
	// example:
	//
	// 86-1860000****
	NewMobilePhone *string `json:"NewMobilePhone,omitempty" xml:"NewMobilePhone,omitempty"`
	// The new username of the RAM user.
	//
	// The username must be 1 to 64 characters in length, and can contain letters, digits, periods (.), hyphens (-), and underscores (_).
	//
	// example:
	//
	// xiaoq****
	NewUserName *string `json:"NewUserName,omitempty" xml:"NewUserName,omitempty"`
	// The username of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
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

func (s *UpdateUserRequest) GetNewUserName() *string {
	return s.NewUserName
}

func (s *UpdateUserRequest) GetUserName() *string {
	return s.UserName
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

func (s *UpdateUserRequest) SetNewUserName(v string) *UpdateUserRequest {
	s.NewUserName = &v
	return s
}

func (s *UpdateUserRequest) SetUserName(v string) *UpdateUserRequest {
	s.UserName = &v
	return s
}

func (s *UpdateUserRequest) Validate() error {
	return dara.Validate(s)
}
