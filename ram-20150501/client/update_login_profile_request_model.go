// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLoginProfileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMFABindRequired(v bool) *UpdateLoginProfileRequest
	GetMFABindRequired() *bool
	SetPassword(v string) *UpdateLoginProfileRequest
	GetPassword() *string
	SetPasswordResetRequired(v bool) *UpdateLoginProfileRequest
	GetPasswordResetRequired() *bool
	SetUserName(v string) *UpdateLoginProfileRequest
	GetUserName() *string
}

type UpdateLoginProfileRequest struct {
	// Specifies whether a multi-factor authentication (MFA) device must be bound to the RAM user upon logon.
	//
	// example:
	//
	// true
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// The logon password of the RAM user. The password must meet the password strength requirements.
	//
	// example:
	//
	// mypassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether the RAM user has to change the password upon logon.
	//
	// example:
	//
	// true
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// The name of the RAM user.
	//
	// example:
	//
	// zhangq****
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateLoginProfileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateLoginProfileRequest) GoString() string {
	return s.String()
}

func (s *UpdateLoginProfileRequest) GetMFABindRequired() *bool {
	return s.MFABindRequired
}

func (s *UpdateLoginProfileRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateLoginProfileRequest) GetPasswordResetRequired() *bool {
	return s.PasswordResetRequired
}

func (s *UpdateLoginProfileRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateLoginProfileRequest) SetMFABindRequired(v bool) *UpdateLoginProfileRequest {
	s.MFABindRequired = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetPassword(v string) *UpdateLoginProfileRequest {
	s.Password = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetPasswordResetRequired(v bool) *UpdateLoginProfileRequest {
	s.PasswordResetRequired = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetUserName(v string) *UpdateLoginProfileRequest {
	s.UserName = &v
	return s
}

func (s *UpdateLoginProfileRequest) Validate() error {
	return dara.Validate(s)
}
