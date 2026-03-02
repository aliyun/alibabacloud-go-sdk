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
	SetStatus(v string) *CreateLoginProfileRequest
	GetStatus() *string
	SetUserPrincipalName(v string) *CreateLoginProfileRequest
	GetUserPrincipalName() *string
}

type CreateLoginProfileRequest struct {
	// Specifies whether the RAM user must enable multi-factor authentication (MFA). Valid values:
	//
	// - true: MFA is required. The RAM user must bind an MFA device at the next logon.
	//
	// - false (default): MFA is not required.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// The logon password for the RAM user.
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
	// - false (default)
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// Specifies whether to enable password-based logon for the console. Valid values:
	//
	// - Active (default): Enables logon.
	//
	// - Inactive: Disables logon.
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

func (s *CreateLoginProfileRequest) GetStatus() *string {
	return s.Status
}

func (s *CreateLoginProfileRequest) GetUserPrincipalName() *string {
	return s.UserPrincipalName
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

func (s *CreateLoginProfileRequest) SetStatus(v string) *CreateLoginProfileRequest {
	s.Status = &v
	return s
}

func (s *CreateLoginProfileRequest) SetUserPrincipalName(v string) *CreateLoginProfileRequest {
	s.UserPrincipalName = &v
	return s
}

func (s *CreateLoginProfileRequest) Validate() error {
	return dara.Validate(s)
}
