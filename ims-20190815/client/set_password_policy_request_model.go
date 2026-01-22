// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHardExpire(v bool) *SetPasswordPolicyRequest
	GetHardExpire() *bool
	SetInitialPasswordAge(v int32) *SetPasswordPolicyRequest
	GetInitialPasswordAge() *int32
	SetMaxLoginAttemps(v int32) *SetPasswordPolicyRequest
	GetMaxLoginAttemps() *int32
	SetMaxPasswordAge(v int32) *SetPasswordPolicyRequest
	GetMaxPasswordAge() *int32
	SetMinimumPasswordDifferentCharacter(v int32) *SetPasswordPolicyRequest
	GetMinimumPasswordDifferentCharacter() *int32
	SetMinimumPasswordLength(v int32) *SetPasswordPolicyRequest
	GetMinimumPasswordLength() *int32
	SetPasswordNotContainUserName(v bool) *SetPasswordPolicyRequest
	GetPasswordNotContainUserName() *bool
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
	// Specifies whether to disable logon after the password expires. Valid values:
	//
	// 	- true: After the password expires, you cannot use the password to log on to the console. You can log on to the console only after you reset the password by using your Alibaba Cloud account or as a RAM user that has administrative rights.
	//
	// 	- false: After the password expires, you can change the password to log on to the console. This is the default value.
	//
	// example:
	//
	// false
	HardExpire         *bool  `json:"HardExpire,omitempty" xml:"HardExpire,omitempty"`
	InitialPasswordAge *int32 `json:"InitialPasswordAge,omitempty" xml:"InitialPasswordAge,omitempty"`
	// The maximum number of password retries. If you enter the wrong passwords for the specified consecutive times, the account is locked for one hour.
	//
	// Valid values: 0 to 32.
	//
	// The default value is 0, which indicates that the password retries are not limited.
	//
	// example:
	//
	// 0
	MaxLoginAttemps *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	// The validity period of the password.
	//
	// Valid values: 0 to 1095. Unit: days.
	//
	// The default value is 0, which indicates that the password never expires.
	//
	// example:
	//
	// 0
	MaxPasswordAge *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	// The minimum number of unique characters in the password.
	//
	// Valid values: 0 to 8.
	//
	// The default value is 0, which indicates that no limits are imposed on the number of unique characters in a password.
	//
	// example:
	//
	// 0
	MinimumPasswordDifferentCharacter *int32 `json:"MinimumPasswordDifferentCharacter,omitempty" xml:"MinimumPasswordDifferentCharacter,omitempty"`
	// The minimum number of characters in the password.
	//
	// Valid values: 8 to 32. Default value: 8.
	//
	// example:
	//
	// 8
	MinimumPasswordLength *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	// Specifies whether to exclude the username from the password. Valid values:
	//
	// 	- true: A password cannot contain the username.
	//
	// 	- false: A password can contain the username. This is the default value.
	//
	// example:
	//
	// false
	PasswordNotContainUserName *bool `json:"PasswordNotContainUserName,omitempty" xml:"PasswordNotContainUserName,omitempty"`
	// The policy for password history check.
	//
	// The previous N passwords cannot be reused. Valid values of N: 0 to 24.
	//
	// The default value is 0, which indicates that RAM users can reuse previous passwords.
	//
	// example:
	//
	// 0
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	// Specifies whether the password must contain lowercase letters. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	RequireLowercaseCharacters *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	// Specifies whether the password must contain digits. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Specifies whether the password must contain special characters. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Specifies whether the password must contain uppercase letters. Valid values:
	//
	// 	- true
	//
	// 	- false (default)
	//
	// example:
	//
	// false
	RequireUppercaseCharacters *bool `json:"RequireUppercaseCharacters,omitempty" xml:"RequireUppercaseCharacters,omitempty"`
}

func (s SetPasswordPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordPolicyRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordPolicyRequest) GetHardExpire() *bool {
	return s.HardExpire
}

func (s *SetPasswordPolicyRequest) GetInitialPasswordAge() *int32 {
	return s.InitialPasswordAge
}

func (s *SetPasswordPolicyRequest) GetMaxLoginAttemps() *int32 {
	return s.MaxLoginAttemps
}

func (s *SetPasswordPolicyRequest) GetMaxPasswordAge() *int32 {
	return s.MaxPasswordAge
}

func (s *SetPasswordPolicyRequest) GetMinimumPasswordDifferentCharacter() *int32 {
	return s.MinimumPasswordDifferentCharacter
}

func (s *SetPasswordPolicyRequest) GetMinimumPasswordLength() *int32 {
	return s.MinimumPasswordLength
}

func (s *SetPasswordPolicyRequest) GetPasswordNotContainUserName() *bool {
	return s.PasswordNotContainUserName
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

func (s *SetPasswordPolicyRequest) SetHardExpire(v bool) *SetPasswordPolicyRequest {
	s.HardExpire = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetInitialPasswordAge(v int32) *SetPasswordPolicyRequest {
	s.InitialPasswordAge = &v
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

func (s *SetPasswordPolicyRequest) SetMinimumPasswordDifferentCharacter(v int32) *SetPasswordPolicyRequest {
	s.MinimumPasswordDifferentCharacter = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetMinimumPasswordLength(v int32) *SetPasswordPolicyRequest {
	s.MinimumPasswordLength = &v
	return s
}

func (s *SetPasswordPolicyRequest) SetPasswordNotContainUserName(v bool) *SetPasswordPolicyRequest {
	s.PasswordNotContainUserName = &v
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
