// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirectoryId(v string) *SetPasswordPolicyRequest
	GetDirectoryId() *string
	SetMaxLoginAttempts(v int32) *SetPasswordPolicyRequest
	GetMaxLoginAttempts() *int32
	SetMaxPasswordAge(v int32) *SetPasswordPolicyRequest
	GetMaxPasswordAge() *int32
	SetMinPasswordDifferentChars(v int32) *SetPasswordPolicyRequest
	GetMinPasswordDifferentChars() *int32
	SetMinPasswordLength(v int32) *SetPasswordPolicyRequest
	GetMinPasswordLength() *int32
	SetPasswordNotContainUsername(v bool) *SetPasswordPolicyRequest
	GetPasswordNotContainUsername() *bool
	SetPasswordReusePrevention(v int32) *SetPasswordPolicyRequest
	GetPasswordReusePrevention() *int32
}

type SetPasswordPolicyRequest struct {
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The number of password retries.
	//
	// If you enter wrong passwords for the specified consecutive times, the account is locked for 1 hour.
	//
	// Valid values: 0 to 32. The value 0 specifies that the number of password retries is not limited.
	//
	// example:
	//
	// 5
	MaxLoginAttempts *int32 `json:"MaxLoginAttempts,omitempty" xml:"MaxLoginAttempts,omitempty"`
	// The validity period of a password.
	//
	// Valid values: 1 to 120. Unit: days.
	//
	// example:
	//
	// 90
	MaxPasswordAge *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	// The minimum number of unique characters in a password.
	//
	// The minimum value is 0, which specifies that the minimum number of unique characters in a password is not limited. The maximum value is the value of the `MinPasswordLength` parameter.
	//
	// example:
	//
	// 8
	MinPasswordDifferentChars *int32 `json:"MinPasswordDifferentChars,omitempty" xml:"MinPasswordDifferentChars,omitempty"`
	// The minimum password length.
	//
	// Valid values: 8 to 32 characters.
	//
	// example:
	//
	// 8
	MinPasswordLength *int32 `json:"MinPasswordLength,omitempty" xml:"MinPasswordLength,omitempty"`
	// Specifies whether a password can contain the username. Valid value:
	//
	// 	- true: A password cannot contain the username.
	//
	// 	- false: A password can contain the username.
	//
	// example:
	//
	// true
	PasswordNotContainUsername *bool `json:"PasswordNotContainUsername,omitempty" xml:"PasswordNotContainUsername,omitempty"`
	// The policy for password history check.
	//
	// The previous N passwords cannot be reused. Valid values of N: 0 to 24. The value 0 specifies that all historical passwords can be reused.
	//
	// >  Passwords that are generated before January 5, 2024 are not counted as historical passwords.
	//
	// example:
	//
	// 1
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
}

func (s SetPasswordPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *SetPasswordPolicyRequest) GetMaxLoginAttempts() *int32 {
	return s.MaxLoginAttempts
}

func (s *SetPasswordPolicyRequest) GetMaxPasswordAge() *int32 {
	return s.MaxPasswordAge
}

func (s *SetPasswordPolicyRequest) GetMinPasswordDifferentChars() *int32 {
	return s.MinPasswordDifferentChars
}

func (s *SetPasswordPolicyRequest) GetMinPasswordLength() *int32 {
	return s.MinPasswordLength
}

func (s *SetPasswordPolicyRequest) GetPasswordNotContainUsername() *bool {
	return s.PasswordNotContainUsername
}

func (s *SetPasswordPolicyRequest) GetPasswordReusePrevention() *int32 {
	return s.PasswordReusePrevention
}

func (s *SetPasswordPolicyRequest) SetDirectoryId(v string) *SetPasswordPolicyRequest {
	s.DirectoryId = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxLoginAttempts(v int32) *SetPasswordPolicyRequest {
	s.MaxLoginAttempts = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxPasswordAge(v int32) *SetPasswordPolicyRequest {
	s.MaxPasswordAge = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMinPasswordDifferentChars(v int32) *SetPasswordPolicyRequest {
	s.MinPasswordDifferentChars = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMinPasswordLength(v int32) *SetPasswordPolicyRequest {
	s.MinPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetPasswordNotContainUsername(v bool) *SetPasswordPolicyRequest {
	s.PasswordNotContainUsername = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetPasswordReusePrevention(v int32) *SetPasswordPolicyRequest {
	s.PasswordReusePrevention = &v
	return s
}

func (s *SetPasswordPolicyRequest) Validate() error {
	return dara.Validate(s)
}
