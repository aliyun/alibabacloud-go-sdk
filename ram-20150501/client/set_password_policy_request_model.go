// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHardExpiry(v bool) *SetPasswordPolicyRequest
	GetHardExpiry() *bool
	SetMaxLoginAttemps(v int32) *SetPasswordPolicyRequest
	GetMaxLoginAttemps() *int32
	SetMaxPasswordAge(v int32) *SetPasswordPolicyRequest
	GetMaxPasswordAge() *int32
	SetMinimumPasswordLength(v int32) *SetPasswordPolicyRequest
	GetMinimumPasswordLength() *int32
	SetPasswordReusePrevention(v int32) *SetPasswordPolicyRequest
	GetPasswordReusePrevention() *int32
	SetRequireLowercaseCharacters(v bool) *SetPasswordPolicyRequest
	GetRequireLowercaseCharacters() *bool
	SetRequireNumbers(v bool) *SetPasswordPolicyRequest
	GetRequireNumbers() *bool
	SetRequireSymbols(v bool) *SetPasswordPolicyRequest
	GetRequireSymbols() *bool
	SetRequireUppercaseCharacters(v bool) *SetPasswordPolicyRequest
	GetRequireUppercaseCharacters() *bool
}

type SetPasswordPolicyRequest struct {
	// Specifies whether a password will expire. Valid values: `true` and `false`. Default value: `false`. If you leave this parameter unspecified, the default value false is used.
	//
	// 	- If you set this parameter to `true`, the Alibaba Cloud account to which the RAM users belong must reset the passwords before the RAM users can log on to the Alibaba Cloud Management Console.
	//
	// 	- If you set this parameter to `false`, the RAM users can change the passwords after the passwords expire and then log on to the Alibaba Cloud Management Console.
	//
	// example:
	//
	// false
	HardExpiry *bool `json:"HardExpiry,omitempty" xml:"HardExpiry,omitempty"`
	// The maximum number of permitted logon attempts within one hour. The number of logon attempts is reset to zero if a RAM user changes the password.
	//
	// example:
	//
	// 5
	MaxLoginAttemps *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	// The number of days for which a password is valid. If you reset a password, the password validity period restarts. Default value: 0. The default value indicates that the password never expires.
	//
	// example:
	//
	// 0
	MaxPasswordAge *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	// The minimum number of characters in a password.
	//
	// Valid values: 8 to 32. Default value: 8.
	//
	// example:
	//
	// 12
	MinimumPasswordLength *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	// The number of previous passwords that a RAM user is prevented from reusing. Default value: 0. The default value indicates that the RAM user can reuse previous passwords.
	//
	// example:
	//
	// 0
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	// Specifies whether a password must contain one or more lowercase letters.
	//
	// example:
	//
	// true
	RequireLowercaseCharacters *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	// Specifies whether a password must contain one or more digits.
	//
	// example:
	//
	// true
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Specifies whether a password must contain one or more special characters.
	//
	// example:
	//
	// true
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Specifies whether a password must contain one or more uppercase letters.
	//
	// example:
	//
	// true
	RequireUppercaseCharacters *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
}

func (s SetPasswordPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyRequest) GetHardExpiry() *bool {
	return s.HardExpiry
}

func (s *SetPasswordPolicyRequest) GetMaxLoginAttemps() *int32 {
	return s.MaxLoginAttemps
}

func (s *SetPasswordPolicyRequest) GetMaxPasswordAge() *int32 {
	return s.MaxPasswordAge
}

func (s *SetPasswordPolicyRequest) GetMinimumPasswordLength() *int32 {
	return s.MinimumPasswordLength
}

func (s *SetPasswordPolicyRequest) GetPasswordReusePrevention() *int32 {
	return s.PasswordReusePrevention
}

func (s *SetPasswordPolicyRequest) GetRequireLowercaseCharacters() *bool {
	return s.RequireLowercaseCharacters
}

func (s *SetPasswordPolicyRequest) GetRequireNumbers() *bool {
	return s.RequireNumbers
}

func (s *SetPasswordPolicyRequest) GetRequireSymbols() *bool {
	return s.RequireSymbols
}

func (s *SetPasswordPolicyRequest) GetRequireUppercaseCharacters() *bool {
	return s.RequireUppercaseCharacters
}

func (s *SetPasswordPolicyRequest) SetHardExpiry(v bool) *SetPasswordPolicyRequest {
	s.HardExpiry = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxLoginAttemps(v int32) *SetPasswordPolicyRequest {
	s.MaxLoginAttemps = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMaxPasswordAge(v int32) *SetPasswordPolicyRequest {
	s.MaxPasswordAge = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMinimumPasswordLength(v int32) *SetPasswordPolicyRequest {
	s.MinimumPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetPasswordReusePrevention(v int32) *SetPasswordPolicyRequest {
	s.PasswordReusePrevention = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireLowercaseCharacters(v bool) *SetPasswordPolicyRequest {
	s.RequireLowercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireNumbers(v bool) *SetPasswordPolicyRequest {
	s.RequireNumbers = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireSymbols(v bool) *SetPasswordPolicyRequest {
	s.RequireSymbols = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetRequireUppercaseCharacters(v bool) *SetPasswordPolicyRequest {
	s.RequireUppercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyRequest) Validate() error {
	return dara.Validate(s)
}
