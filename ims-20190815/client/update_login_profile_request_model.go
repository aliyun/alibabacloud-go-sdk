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
	SetStatus(v string) *UpdateLoginProfileRequest
	GetStatus() *string
	SetUserPrincipalName(v string) *UpdateLoginProfileRequest
	GetUserPrincipalName() *string
}

type UpdateLoginProfileRequest struct {
	// Specifies whether to enforce multi-factor authentication (MFA) for the RAM user. Valid values:
	//
	// - true: Enforce MFA. The RAM user must attach an MFA device at the next logon.
	//
	// - false: Do not enforce MFA.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// The new console logon password for the RAM user.
	//
	// The password must meet the password strength requirements.
	//
	// example:
	//
	// mypassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether the RAM user must reset the password at the next logon. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// Specifies whether to enable or disable password-based logon to the console. Valid values:
	//
	// - Active: Enabled.
	//
	// - Inactive: Disabled.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The logon name of the RAM user.
	//
	// This parameter is required.
	//
	// example:
	//
	// test@example.onaliyun.com
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
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

func (s *UpdateLoginProfileRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdateLoginProfileRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
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

func (s *UpdateLoginProfileRequest) SetStatus(v string) *UpdateLoginProfileRequest {
	s.Status = &v
	return s
}

func (s *UpdateLoginProfileRequest) SetUserPrincipalName(v string) *UpdateLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *UpdateLoginProfileRequest) Validate() error {
	return dara.Validate(s)
}
