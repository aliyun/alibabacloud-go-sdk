// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLoginProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMFABindRequired(v bool) *CreateLoginProfileRequest
	GetMFABindRequired() *bool
	SetPassword(v string) *CreateLoginProfileRequest
	GetPassword() *string
	SetPasswordResetRequired(v bool) *CreateLoginProfileRequest
	GetPasswordResetRequired() *bool
	SetUserName(v string) *CreateLoginProfileRequest
	GetUserName() *string
}

type CreateLoginProfileRequest struct {
	// Specifies whether the user must bind a multi-factor authentication (MFA) device upon next logon. Default value: `false`.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// The password. The password must meet the password strength requirements. For more information about password strength requirements, see [GetPasswordPolicy](https://help.aliyun.com/document_detail/2337691.html).
	//
	// example:
	//
	// mypassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether the user is required to reset the password upon logon. Default value: `false`.
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// The username.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s CreateLoginProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *CreateLoginProfileRequest) GetMFABindRequired() *bool {
	return s.MFABindRequired
}

func (s *CreateLoginProfileRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateLoginProfileRequest) GetPasswordResetRequired() *bool {
	return s.PasswordResetRequired
}

func (s *CreateLoginProfileRequest) GetUserName() *string {
	return s.UserName
}

func (s *CreateLoginProfileRequest) SetMFABindRequired(v bool) *CreateLoginProfileRequest {
	s.MFABindRequired = &v
	return s
}

func (s *CreateLoginProfileRequest) SetPassword(v string) *CreateLoginProfileRequest {
	s.Password = &v
	return s
}

func (s *CreateLoginProfileRequest) SetPasswordResetRequired(v bool) *CreateLoginProfileRequest {
	s.PasswordResetRequired = &v
	return s
}

func (s *CreateLoginProfileRequest) SetUserName(v string) *CreateLoginProfileRequest {
	s.UserName = &v
	return s
}

func (s *CreateLoginProfileRequest) Validate() error {
	return dara.Validate(s)
}
