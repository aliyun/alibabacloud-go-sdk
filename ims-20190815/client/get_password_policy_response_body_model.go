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
	// The details of the password policy.
	PasswordPolicy *GetPasswordPolicyResponseBodyPasswordPolicy `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BDAA8408-E67C-428B-BFF0-1B2AC05C9610
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
	// Indicates whether to disable logon after the password expires.
	//
	// example:
	//
	// false
	HardExpire         *bool  `json:"HardExpire,omitempty" xml:"HardExpire,omitempty"`
	InitialPasswordAge *int32 `json:"InitialPasswordAge,omitempty" xml:"InitialPasswordAge,omitempty"`
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
	// The minimum required number of characters in a password.
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

func (s GetPasswordPolicyResponseBodyPasswordPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordPolicyResponseBodyPasswordPolicy) GoString() string {
	return s.String()
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetHardExpire() *bool {
	return s.HardExpire
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetInitialPasswordAge() *int32 {
	return s.InitialPasswordAge
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetMaxLoginAttemps() *int32 {
	return s.MaxLoginAttemps
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetMaxPasswordAge() *int32 {
	return s.MaxPasswordAge
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetMinimumPasswordDifferentCharacter() *int32 {
	return s.MinimumPasswordDifferentCharacter
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetMinimumPasswordLength() *int32 {
	return s.MinimumPasswordLength
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetPasswordNotContainUserName() *bool {
	return s.PasswordNotContainUserName
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetPasswordReusePrevention() *int32 {
	return s.PasswordReusePrevention
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetRequireLowercaseCharacters() *bool {
	return s.RequireLowercaseCharacters
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetRequireNumbers() *bool {
	return s.RequireNumbers
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetRequireSymbols() *bool {
	return s.RequireSymbols
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) GetRequireUppercaseCharacters() *bool {
	return s.RequireUppercaseCharacters
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetHardExpire(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.HardExpire = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetInitialPasswordAge(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.InitialPasswordAge = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxLoginAttemps(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxLoginAttemps = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMaxPasswordAge(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MaxPasswordAge = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMinimumPasswordDifferentCharacter(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinimumPasswordDifferentCharacter = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetMinimumPasswordLength(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.MinimumPasswordLength = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordNotContainUserName(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordNotContainUserName = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetPasswordReusePrevention(v int32) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.PasswordReusePrevention = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireLowercaseCharacters(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireLowercaseCharacters = &v
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

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) SetRequireUppercaseCharacters(v bool) *GetPasswordPolicyResponseBodyPasswordPolicy {
	s.RequireUppercaseCharacters = &v
	return s
}

func (s *GetPasswordPolicyResponseBodyPasswordPolicy) Validate() error {
	return dara.Validate(s)
}
