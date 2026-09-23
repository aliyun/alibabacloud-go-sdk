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
	// Specifies whether the Resource Access Management (RAM) user is required to enable multi-factor authentication (MFA). Valid values:
	//
	// - true: MFA is required. The RAM user must attach an MFA device at the next logon.
	//
	// - false: MFA is not required.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// The new console logon password of the Resource Access Management (RAM) user.
	//
	// The password must meet the password strength requirements.
	//
	// example:
	//
	// mypassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether the Resource Access Management (RAM) user is required to reset the password at the next logon. Valid values:
	//
	// - true
	//
	// - false
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// Enables or disables console password logon. Valid values:
	//
	// - Active: enables console password logon.
	//
	// - Inactive: disables console password logon.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The logon name of the Resource Access Management (RAM) user.
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
