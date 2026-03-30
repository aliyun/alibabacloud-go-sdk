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
	// Specifies whether to forcefully enable multi-factor authentication (MFA) for the RAM user. Valid values:
	//
	// 	- true: forcefully enables MFA for the RAM user. The RAM user must bind an MFA device upon the next logon.
	//
	// 	- false (default): does not forcefully enable MFA for the RAM user.
	//
	// example:
	//
	// false
	MFABindRequired *bool `json:"MFABindRequired,omitempty" xml:"MFABindRequired,omitempty"`
	// The password that the RAM user uses to log on to the console.
	//
	// The password must meet the complexity requirements.
	//
	// example:
	//
	// mypassword
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether the RAM user is required to reset the password upon the next logon. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	PasswordResetRequired *bool `json:"PasswordResetRequired,omitempty" xml:"PasswordResetRequired,omitempty"`
	// Specifies whether to enable password-based logons to the console. Valid values:
	//
	// 	- Active: Password-based logon is enabled. This is the default value.
	//
	// 	- Inactive: Password-based logon is disabled.
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
