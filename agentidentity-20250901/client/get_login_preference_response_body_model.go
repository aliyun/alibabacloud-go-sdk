// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginPreferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLoginPreference(v *GetLoginPreferenceResponseBodyLoginPreference) *GetLoginPreferenceResponseBody
	GetLoginPreference() *GetLoginPreferenceResponseBodyLoginPreference
	SetPasswordPolicy(v *GetLoginPreferenceResponseBodyPasswordPolicy) *GetLoginPreferenceResponseBody
	GetPasswordPolicy() *GetLoginPreferenceResponseBodyPasswordPolicy
	SetRequestId(v string) *GetLoginPreferenceResponseBody
	GetRequestId() *string
}

type GetLoginPreferenceResponseBody struct {
	LoginPreference *GetLoginPreferenceResponseBodyLoginPreference `json:"LoginPreference,omitempty" xml:"LoginPreference,omitempty" type:"Struct"`
	PasswordPolicy  *GetLoginPreferenceResponseBodyPasswordPolicy  `json:"PasswordPolicy,omitempty" xml:"PasswordPolicy,omitempty" type:"Struct"`
	RequestId       *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetLoginPreferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLoginPreferenceResponseBody) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceResponseBody) GetLoginPreference() *GetLoginPreferenceResponseBodyLoginPreference {
	return s.LoginPreference
}

func (s *GetLoginPreferenceResponseBody) GetPasswordPolicy() *GetLoginPreferenceResponseBodyPasswordPolicy {
	return s.PasswordPolicy
}

func (s *GetLoginPreferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLoginPreferenceResponseBody) SetLoginPreference(v *GetLoginPreferenceResponseBodyLoginPreference) *GetLoginPreferenceResponseBody {
	s.LoginPreference = v
	return s
}

func (s *GetLoginPreferenceResponseBody) SetPasswordPolicy(v *GetLoginPreferenceResponseBodyPasswordPolicy) *GetLoginPreferenceResponseBody {
	s.PasswordPolicy = v
	return s
}

func (s *GetLoginPreferenceResponseBody) SetRequestId(v string) *GetLoginPreferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLoginPreferenceResponseBody) Validate() error {
	if s.LoginPreference != nil {
		if err := s.LoginPreference.Validate(); err != nil {
			return err
		}
	}
	if s.PasswordPolicy != nil {
		if err := s.PasswordPolicy.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLoginPreferenceResponseBodyLoginPreference struct {
	EnablePasswordLogin *bool `json:"EnablePasswordLogin,omitempty" xml:"EnablePasswordLogin,omitempty"`
}

func (s GetLoginPreferenceResponseBodyLoginPreference) String() string {
	return dara.Prettify(s)
}

func (s GetLoginPreferenceResponseBodyLoginPreference) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceResponseBodyLoginPreference) GetEnablePasswordLogin() *bool {
	return s.EnablePasswordLogin
}

func (s *GetLoginPreferenceResponseBodyLoginPreference) SetEnablePasswordLogin(v bool) *GetLoginPreferenceResponseBodyLoginPreference {
	s.EnablePasswordLogin = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyLoginPreference) Validate() error {
	return dara.Validate(s)
}

type GetLoginPreferenceResponseBodyPasswordPolicy struct {
	HardExpire                 *bool  `json:"HardExpire,omitempty" xml:"HardExpire,omitempty"`
	MaxLoginAttempts           *int64 `json:"MaxLoginAttempts,omitempty" xml:"MaxLoginAttempts,omitempty"`
	MaxPasswordAge             *int64 `json:"MaxPasswordAge,omitempty" xml:"MaxPasswordAge,omitempty"`
	MaxPasswordLength          *int32 `json:"MaxPasswordLength,omitempty" xml:"MaxPasswordLength,omitempty"`
	MinPasswordLength          *int32 `json:"MinPasswordLength,omitempty" xml:"MinPasswordLength,omitempty"`
	PasswordNotContainUserName *bool  `json:"PasswordNotContainUserName,omitempty" xml:"PasswordNotContainUserName,omitempty"`
	RequireLowerCaseChars      *bool  `json:"RequireLowerCaseChars,omitempty" xml:"RequireLowerCaseChars,omitempty"`
	RequireNumbers             *bool  `json:"RequireNumbers,omitempty" xml:"RequireNumbers,omitempty"`
	RequireSymbols             *bool  `json:"RequireSymbols,omitempty" xml:"RequireSymbols,omitempty"`
	RequireUpperCaseChars      *bool  `json:"RequireUpperCaseChars,omitempty" xml:"RequireUpperCaseChars,omitempty"`
}

func (s GetLoginPreferenceResponseBodyPasswordPolicy) String() string {
	return dara.Prettify(s)
}

func (s GetLoginPreferenceResponseBodyPasswordPolicy) GoString() string {
	return s.String()
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) GetHardExpire() *bool {
	return s.HardExpire
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) GetMaxLoginAttempts() *int64 {
	return s.MaxLoginAttempts
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) GetMaxPasswordAge() *int64 {
	return s.MaxPasswordAge
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) GetMaxPasswordLength() *int32 {
	return s.MaxPasswordLength
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) GetMinPasswordLength() *int32 {
	return s.MinPasswordLength
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) GetPasswordNotContainUserName() *bool {
	return s.PasswordNotContainUserName
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) GetRequireLowerCaseChars() *bool {
	return s.RequireLowerCaseChars
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) GetRequireNumbers() *bool {
	return s.RequireNumbers
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) GetRequireSymbols() *bool {
	return s.RequireSymbols
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) GetRequireUpperCaseChars() *bool {
	return s.RequireUpperCaseChars
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) SetHardExpire(v bool) *GetLoginPreferenceResponseBodyPasswordPolicy {
	s.HardExpire = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) SetMaxLoginAttempts(v int64) *GetLoginPreferenceResponseBodyPasswordPolicy {
	s.MaxLoginAttempts = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) SetMaxPasswordAge(v int64) *GetLoginPreferenceResponseBodyPasswordPolicy {
	s.MaxPasswordAge = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) SetMaxPasswordLength(v int32) *GetLoginPreferenceResponseBodyPasswordPolicy {
	s.MaxPasswordLength = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) SetMinPasswordLength(v int32) *GetLoginPreferenceResponseBodyPasswordPolicy {
	s.MinPasswordLength = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) SetPasswordNotContainUserName(v bool) *GetLoginPreferenceResponseBodyPasswordPolicy {
	s.PasswordNotContainUserName = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) SetRequireLowerCaseChars(v bool) *GetLoginPreferenceResponseBodyPasswordPolicy {
	s.RequireLowerCaseChars = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) SetRequireNumbers(v bool) *GetLoginPreferenceResponseBodyPasswordPolicy {
	s.RequireNumbers = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) SetRequireSymbols(v bool) *GetLoginPreferenceResponseBodyPasswordPolicy {
	s.RequireSymbols = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) SetRequireUpperCaseChars(v bool) *GetLoginPreferenceResponseBodyPasswordPolicy {
	s.RequireUpperCaseChars = &v
	return s
}

func (s *GetLoginPreferenceResponseBodyPasswordPolicy) Validate() error {
	return dara.Validate(s)
}
