// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordComplexityConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPasswordComplexityConfiguration(v *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) *GetPasswordComplexityConfigurationResponseBody
	GetPasswordComplexityConfiguration() *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration
	SetRequestId(v string) *GetPasswordComplexityConfigurationResponseBody
	GetRequestId() *string
}

type GetPasswordComplexityConfigurationResponseBody struct {
	// The password complexity policy configuration.
	PasswordComplexityConfiguration *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration `json:"PasswordComplexityConfiguration,omitempty" xml:"PasswordComplexityConfiguration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPasswordComplexityConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordComplexityConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetPasswordComplexityConfigurationResponseBody) GetPasswordComplexityConfiguration() *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration {
	return s.PasswordComplexityConfiguration
}

func (s *GetPasswordComplexityConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPasswordComplexityConfigurationResponseBody) SetPasswordComplexityConfiguration(v *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) *GetPasswordComplexityConfigurationResponseBody {
	s.PasswordComplexityConfiguration = v
	return s
}

func (s *GetPasswordComplexityConfigurationResponseBody) SetRequestId(v string) *GetPasswordComplexityConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPasswordComplexityConfigurationResponseBody) Validate() error {
	if s.PasswordComplexityConfiguration != nil {
		if err := s.PasswordComplexityConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration struct {
	// Indicates whether logon with a weak password is disabled.
	//
	// example:
	//
	// false
	DisabledWeakPasswordLogin *bool `json:"DisabledWeakPasswordLogin,omitempty" xml:"DisabledWeakPasswordLogin,omitempty"`
	// The time when the weak password logon restriction takes effect.
	//
	// example:
	//
	// 1773383634936
	DisabledWeakPasswordLoginStartedAt *int64 `json:"DisabledWeakPasswordLoginStartedAt,omitempty" xml:"DisabledWeakPasswordLoginStartedAt,omitempty"`
	// The list of password complexity rules.
	PasswordComplexityRules []*GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules `json:"PasswordComplexityRules,omitempty" xml:"PasswordComplexityRules,omitempty" type:"Repeated"`
	// The minimum password length.
	//
	// example:
	//
	// 3
	PasswordMinLength *int32 `json:"PasswordMinLength,omitempty" xml:"PasswordMinLength,omitempty"`
}

func (s GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) GoString() string {
	return s.String()
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) GetDisabledWeakPasswordLogin() *bool {
	return s.DisabledWeakPasswordLogin
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) GetDisabledWeakPasswordLoginStartedAt() *int64 {
	return s.DisabledWeakPasswordLoginStartedAt
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) GetPasswordComplexityRules() []*GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules {
	return s.PasswordComplexityRules
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) GetPasswordMinLength() *int32 {
	return s.PasswordMinLength
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) SetDisabledWeakPasswordLogin(v bool) *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration {
	s.DisabledWeakPasswordLogin = &v
	return s
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) SetDisabledWeakPasswordLoginStartedAt(v int64) *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration {
	s.DisabledWeakPasswordLoginStartedAt = &v
	return s
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) SetPasswordComplexityRules(v []*GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules) *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration {
	s.PasswordComplexityRules = v
	return s
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) SetPasswordMinLength(v int32) *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration {
	s.PasswordMinLength = &v
	return s
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfiguration) Validate() error {
	if s.PasswordComplexityRules != nil {
		for _, item := range s.PasswordComplexityRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules struct {
	// The password check type. Valid values:
	//
	// - inclusion_upper_case: The password must contain uppercase letters.
	//
	// - inclusion_lower_case: The password must contain lowercase letters.
	//
	// - inclusion_special_case: The password must contain at least one of the following special characters: \\` @ % + \\ / \\" ! # $ ^ ? : , ( ) { } [ ] \\~ - _ . \\`.
	//
	// - inclusion_number: The password must contain digits.
	//
	// - exclusion_username: The password cannot contain the username.
	//
	// - exclusion_email: The password cannot contain the email address.
	//
	// - exclusion_phone_number: The password cannot contain the mobile number.
	//
	// - exclusion_display_name: The password cannot contain the display name.
	//
	// example:
	//
	// inclusion_upper_case
	PasswordCheckType *string `json:"PasswordCheckType,omitempty" xml:"PasswordCheckType,omitempty"`
}

func (s GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules) GoString() string {
	return s.String()
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules) GetPasswordCheckType() *string {
	return s.PasswordCheckType
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules) SetPasswordCheckType(v string) *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules {
	s.PasswordCheckType = &v
	return s
}

func (s *GetPasswordComplexityConfigurationResponseBodyPasswordComplexityConfigurationPasswordComplexityRules) Validate() error {
	return dara.Validate(s)
}
