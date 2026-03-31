// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPasswordPolicy(v *SetPasswordPolicyResponseBodyPasswordPolicy) *SetPasswordPolicyResponseBody
	GetPasswordPolicy() *SetPasswordPolicyResponseBodyPasswordPolicy
	SetRequestId(v string) *SetPasswordPolicyResponseBody
	GetRequestId() *string
}

type SetPasswordPolicyResponseBody struct {
	// The password policy.
	PasswordPolicy *SetPasswordPolicyResponseBodyPasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetPasswordPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponseBody) GetPasswordPolicy() *SetPasswordPolicyResponseBodyPasswordPolicy {
	return s.PasswordPolicy
}

func (s *SetPasswordPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetPasswordPolicyResponseBody) SetPasswordPolicy(v *SetPasswordPolicyResponseBodyPasswordPolicy) *SetPasswordPolicyResponseBody {
	s.PasswordPolicy = v
	return s
}

func (s *SetPasswordPolicyResponseBody) SetRequestId(v string) *SetPasswordPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetPasswordPolicyResponseBody) Validate() error {
	if s.PasswordPolicy != nil {
		if err := s.PasswordPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SetPasswordPolicyResponseBodyPasswordPolicy struct {
	// Indicates whether a password expires. Valid values: `true` and `false`. Default value: `false`. If the parameter is unspecified, the default value false is returned.
	//
	// 	- If this parameter is set to `true`, the Alibaba Cloud account to which the RAM users belong must reset the password before the RAM users can log on to the Alibaba Cloud Management Console.
	//
	// 	- If this parameter is set to `false`, the RAM users can change the passwords after the passwords expire and then log on to the Alibaba Cloud Management Console.
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
	// Indicates whether a password must contain one or more lowercase letters.
	//
	// example:
	//
	// true
	RequireLowercaseCharacters *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	// Indicates whether a password must contain one or more digits.
	//
	// example:
	//
	// true
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Indicates whether a password must contain one or more special characters.
	//
	// example:
	//
	// true
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Indicates whether a password must contain one or more uppercase letters.
	//
	// example:
	//
	// true
	RequireUppercaseCharacters *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
}

func (s SetPasswordPolicyResponseBodyPasswordPolicy) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordPolicyResponseBodyPasswordPolicy) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetHardExpiry() *bool {
	return s.HardExpiry
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetMaxLoginAttemps() *int32 {
	return s.MaxLoginAttemps
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetMaxPasswordAge() *int32 {
	return s.MaxPasswordAge
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetMinimumPasswordLength() *int32 {
	return s.MinimumPasswordLength
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetPasswordReusePrevention() *int32 {
	return s.PasswordReusePrevention
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetRequireLowercaseCharacters() *bool {
	return s.RequireLowercaseCharacters
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetRequireNumbers() *bool {
	return s.RequireNumbers
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetRequireSymbols() *bool {
	return s.RequireSymbols
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetRequireUppercaseCharacters() *bool {
	return s.RequireUppercaseCharacters
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetHardExpiry(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.HardExpiry = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMaxLoginAttemps(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxLoginAttemps = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMaxPasswordAge(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxPasswordAge = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMinimumPasswordLength(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinimumPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordReusePrevention(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordReusePrevention = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireLowercaseCharacters(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireLowercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireNumbers(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireNumbers = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireSymbols(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireSymbols = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetRequireUppercaseCharacters(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireUppercaseCharacters = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) Validate() error {
	return dara.Validate(s)
}
