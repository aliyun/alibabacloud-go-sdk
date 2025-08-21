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
	// The details of the password policy.
	PasswordPolicy *SetPasswordPolicyResponseBodyPasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 3FB5551F-B2ED-40D4-8392-1E4AC2384EFD
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
	return dara.Validate(s)
}

type SetPasswordPolicyResponseBodyPasswordPolicy struct {
	// Indicates whether to disable logon after the password expires.
	//
	// example:
	//
	// false
	HardExpire *bool `json:"HardExpire,omitempty" xml:"HardExpire,omitempty"`
	// The maximum number of password retries.
	//
	// example:
	//
	// 0
	MaxLoginAttemps *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	// The validity period of the password.
	//
	// example:
	//
	// 0
	MaxPasswordAge *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	// The minimum number of unique characters in the password.
	//
	// example:
	//
	// 0
	MinimumPasswordDifferentCharacter *int32 `json:"MinimumPasswordDifferentCharacter,omitempty" xml:"MinimumPasswordDifferentCharacter,omitempty"`
	// The minimum number of characters in the password.
	//
	// example:
	//
	// 8
	MinimumPasswordLength *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	// Indicates whether to exclude the username from the password.
	//
	// example:
	//
	// false
	PasswordNotContainUserName *bool `json:"PasswordNotContainUserName,omitempty" xml:"PasswordNotContainUserName,omitempty"`
	// The policy for password history check.
	//
	// example:
	//
	// 0
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	// Indicates whether the password must contain lowercase letters.
	//
	// example:
	//
	// false
	RequireLowercaseCharacters *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	// Indicates whether the password must contain digits.
	//
	// example:
	//
	// false
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Indicates whether the password must contain special characters.
	//
	// example:
	//
	// false
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Indicates whether the password must contain uppercase letters.
	//
	// example:
	//
	// false
	RequireUppercaseCharacters *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
}

func (s SetPasswordPolicyResponseBodyPasswordPolicy) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordPolicyResponseBodyPasswordPolicy) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetHardExpire() *bool {
	return s.HardExpire
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetMaxLoginAttemps() *int32 {
	return s.MaxLoginAttemps
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetMaxPasswordAge() *int32 {
	return s.MaxPasswordAge
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetMinimumPasswordDifferentCharacter() *int32 {
	return s.MinimumPasswordDifferentCharacter
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetMinimumPasswordLength() *int32 {
	return s.MinimumPasswordLength
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) GetPasswordNotContainUserName() *bool {
	return s.PasswordNotContainUserName
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

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetHardExpire(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.HardExpire = &v
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

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMinimumPasswordDifferentCharacter(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinimumPasswordDifferentCharacter = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetMinimumPasswordLength(v int32) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinimumPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordNotContainUserName(v bool) *SetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordNotContainUserName = &v
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
