// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPasswordPolicy(v *GetPasswordPolicyResponseBodyPasswordPolicy) *GetPasswordPolicyResponseBody
	GetPasswordPolicy() *GetPasswordPolicyResponseBodyPasswordPolicy
	SetRequestId(v string) *GetPasswordPolicyResponseBody
	GetRequestId() *string
}

type GetPasswordPolicyResponseBody struct {
	// The password policy.
	PasswordPolicy *GetPasswordPolicyResponseBodyPasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B7C6E839-FB65-59BE-B753-003AA8AF7DF7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPasswordPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponseBody) GetPasswordPolicy() *GetPasswordPolicyResponseBodyPasswordPolicy {
	return s.PasswordPolicy
}

func (s *GetPasswordPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPasswordPolicyResponseBody) SetPasswordPolicy(v *GetPasswordPolicyResponseBodyPasswordPolicy) *GetPasswordPolicyResponseBody {
	s.PasswordPolicy = v
	return s
}

func (s *GetPasswordPolicyResponseBody) SetRequestId(v string) *GetPasswordPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPasswordPolicyResponseBody) Validate() error {
	if s.PasswordPolicy != nil {
		if err := s.PasswordPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPasswordPolicyResponseBodyPasswordPolicy struct {
	// Indicates whether to disable logon after a password expires. Valid values:
	//
	// 	- true: disables logon after a password expires.
	//
	// 	- false: does not disable logon after a password expires.
	//
	// example:
	//
	// true
	HardExpire *bool `json:"HardExpire,omitempty" xml:"HardExpire,omitempty"`
	// The number of password retries.
	//
	// If wrong passwords are entered for the specified consecutive times, the account is locked for 1 hour.
	//
	// Valid values: 0 to 32. The value 0 indicates that the number of password retries is not limited.
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
	// The maximum password length.
	//
	// example:
	//
	// 32
	MaxPasswordLength *int32 `json:"MaxPasswordLength,omitempty" xml:"MaxPasswordLength,omitempty"`
	// The minimum number of different characters in a password.
	//
	// The minimum value is 0, which indicates that the minimum number of different characters in a password is not limited. The maximum value is the value of the `MinPasswordLength` parameter.
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
	// Indicates whether to exclude the username from the password. Valid values:
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
	// The previous N passwords cannot be reused. Valid values of N: 0 to 24. The value 0 indicates that all historical passwords can be reused.
	//
	// >  Passwords that are generated before January 5, 2024 are not counted as historical passwords.
	//
	// example:
	//
	// 1
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	// Indicates whether lowercase letters are required in a password. Valid values:
	//
	// 	- true: Lowercase letters are required in a password.
	//
	// 	- false: Lowercase letters are not required in a password.
	//
	// example:
	//
	// true
	RequireLowerCaseChars *bool `json:"RequireLowerCaseChars,omitempty" xml:"RequireLowerCaseChars,omitempty"`
	// Indicates whether digits are required in a password. Valid values:
	//
	// 	- true: Digits are required in a password.
	//
	// 	- false: Digits are not required in a password.
	//
	// example:
	//
	// true
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Indicates whether special characters are required in a password. Valid values:
	//
	// 	- true: Special characters are required in a password.
	//
	// 	- false: Special characters are not required in a password.
	//
	// example:
	//
	// true
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Indicates whether uppercase letters are required in a password. Valid values:
	//
	// 	- true: Uppercase letters are required in a password.
	//
	// 	- false: Uppercase letters are not required in a password.
	//
	// example:
	//
	// true
	RequireUpperCaseChars *bool `json:"RequireUpperCaseChars,omitempty" xml:"RequireUpperCaseChars,omitempty"`
}

func (s GetPasswordPolicyResponseBodyPasswordPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordPolicyResponseBodyPasswordPolicy) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetHardExpire() *bool {
	return s.HardExpire
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetMaxLoginAttempts() *int32 {
	return s.MaxLoginAttempts
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetMaxPasswordAge() *int32 {
	return s.MaxPasswordAge
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetMaxPasswordLength() *int32 {
	return s.MaxPasswordLength
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetMinPasswordDifferentChars() *int32 {
	return s.MinPasswordDifferentChars
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetMinPasswordLength() *int32 {
	return s.MinPasswordLength
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetPasswordNotContainUsername() *bool {
	return s.PasswordNotContainUsername
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetPasswordReusePrevention() *int32 {
	return s.PasswordReusePrevention
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetRequireLowerCaseChars() *bool {
	return s.RequireLowerCaseChars
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetRequireNumbers() *bool {
	return s.RequireNumbers
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetRequireSymbols() *bool {
	return s.RequireSymbols
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetRequireUpperCaseChars() *bool {
	return s.RequireUpperCaseChars
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetHardExpire(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.HardExpire = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxLoginAttempts(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxLoginAttempts = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxPasswordAge(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxPasswordAge = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxPasswordLength(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxPasswordLength = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMinPasswordDifferentChars(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinPasswordDifferentChars = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMinPasswordLength(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinPasswordLength = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordNotContainUsername(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordNotContainUsername = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordReusePrevention(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordReusePrevention = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireLowerCaseChars(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireLowerCaseChars = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireNumbers(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireNumbers = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireSymbols(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireSymbols = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireUpperCaseChars(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireUpperCaseChars = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) Validate() error {
	return dara.Validate(s)
}
