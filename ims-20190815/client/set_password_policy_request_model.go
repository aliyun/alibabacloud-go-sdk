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
	SetInterceptRiskPasswordOnApi(v bool) *SetPasswordPolicyRequest
	GetInterceptRiskPasswordOnApi() *bool
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
	// Specifies whether logon is blocked after a password expires.
	//
	// - true: After a password expires, the RAM user cannot log on to the console. An Alibaba Cloud account owner or a RAM administrator must reset the password before the RAM user can log on.
	//
	// - false (default): After a password expires, the RAM user can change the password and then log on.
	//
	// example:
	//
	// false
	HardExpire *bool `json:"HardExpire,omitempty" xml:"HardExpire,omitempty"`
	// Validity period for initial passwords. Initial passwords apply to newly created RAM users or users whose console logon settings are re-enabled.
	//
	// Valid values: 0 to 90. Unit: days.
	//
	// Default value: 14.
	//
	// A value of 0 disables this constraint.
	//
	// example:
	//
	// 14
	InitialPasswordAge *int32 `json:"InitialPasswordAge,omitempty" xml:"InitialPasswordAge,omitempty"`
	// Specifies whether threat passwords are blocked when set using APIs.
	//
	// Default value: false
	//
	// - true
	//
	// - false (default)
	//
	// example:
	//
	// false
	InterceptRiskPasswordOnApi *bool `json:"InterceptRiskPasswordOnApi,omitempty" xml:"InterceptRiskPasswordOnApi,omitempty"`
	// Maximum number of failed password attempts. After the specified number of consecutive incorrect password attempts, the account is locked for one hour.
	//
	// Valid values: 0 to 32.
	//
	// Default value: 0, which disables this constraint.
	//
	// example:
	//
	// 0
	MaxLoginAttemps *int32 `json:"MaxLoginAttemps,omitempty" xml:"MaxLoginAttemps,omitempty"`
	// Password validity period.
	//
	// Valid values: 0 to 1095. Unit: days.
	//
	// Default value: 0, which means passwords never expire.
	//
	// example:
	//
	// 0
	MaxPasswordAge *int32 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	// Minimum number of unique characters in a password.
	//
	// Valid values: 0 to 8.
	//
	// Default value: 0, which imposes no restriction.
	//
	// example:
	//
	// 0
	MinimumPasswordDifferentCharacter *int32 `json:"MinimumPasswordDifferentCharacter,omitempty" xml:"MinimumPasswordDifferentCharacter,omitempty"`
	// Minimum password length.
	//
	// Valid values: 8 to 32.
	//
	// Default value: 8.
	//
	// example:
	//
	// 8
	MinimumPasswordLength *int32 `json:"MinimumPasswordLength,omitempty" xml:"MinimumPasswordLength,omitempty"`
	// Specifies whether passwords must not contain the user name.
	//
	// - true
	//
	// - false (default)
	//
	// example:
	//
	// false
	PasswordNotContainUserName *bool `json:"PasswordNotContainUserName,omitempty" xml:"PasswordNotContainUserName,omitempty"`
	// Prevents reuse of previous passwords.
	//
	// Valid values: 0 to 24. This value specifies how many previous passwords are blocked from reuse.
	//
	// Default value: 0, which disables this constraint.
	//
	// example:
	//
	// 0
	PasswordReusePrevention *int32 `json:"PasswordReusePrevention,omitempty" xml:"PasswordReusePrevention,omitempty"`
	// Specifies whether passwords must contain lowercase letters.
	//
	// - true
	//
	// - false (default)
	//
	// example:
	//
	// false
	RequireLowercaseCharacters *bool `json:"RequireLowercaseCharacters,omitempty" xml:"RequireLowercaseCharacters,omitempty"`
	// Specifies whether passwords must contain numbers.
	//
	// - true
	//
	// - false (default)
	//
	// example:
	//
	// false
	RequireNumbers *bool `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	// Specifies whether passwords must contain special characters.
	//
	// - true
	//
	// - false (default)
	//
	// example:
	//
	// false
	RequireSymbols *bool `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	// Specifies whether passwords must contain uppercase letters.
	//
	// - true
	//
	// - false (default)
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

func (s *SetPasswordPolicyRequest) GetInterceptRiskPasswordOnApi() *bool {
	return s.InterceptRiskPasswordOnApi
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

func (s *SetPasswordPolicyRequest) SetInterceptRiskPasswordOnApi(v bool) *SetPasswordPolicyRequest {
	s.InterceptRiskPasswordOnApi = &v
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
