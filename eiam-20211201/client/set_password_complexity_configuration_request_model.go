// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordComplexityConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SetPasswordComplexityConfigurationRequest
	GetInstanceId() *string
	SetPasswordComplexityRules(v []*SetPasswordComplexityConfigurationRequestPasswordComplexityRules) *SetPasswordComplexityConfigurationRequest
	GetPasswordComplexityRules() []*SetPasswordComplexityConfigurationRequestPasswordComplexityRules
	SetPasswordMinLength(v int32) *SetPasswordComplexityConfigurationRequest
	GetPasswordMinLength() *int32
}

type SetPasswordComplexityConfigurationRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The password complexity rules.
	PasswordComplexityRules []*SetPasswordComplexityConfigurationRequestPasswordComplexityRules `json:"PasswordComplexityRules,omitempty" xml:"PasswordComplexityRules,omitempty" type:"Repeated"`
	// The minimum number of characters in a password.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PasswordMinLength *int32 `json:"PasswordMinLength,omitempty" xml:"PasswordMinLength,omitempty"`
}

func (s SetPasswordComplexityConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordComplexityConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordComplexityConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPasswordComplexityConfigurationRequest) GetPasswordComplexityRules() []*SetPasswordComplexityConfigurationRequestPasswordComplexityRules {
	return s.PasswordComplexityRules
}

func (s *SetPasswordComplexityConfigurationRequest) GetPasswordMinLength() *int32 {
	return s.PasswordMinLength
}

func (s *SetPasswordComplexityConfigurationRequest) SetInstanceId(v string) *SetPasswordComplexityConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPasswordComplexityConfigurationRequest) SetPasswordComplexityRules(v []*SetPasswordComplexityConfigurationRequestPasswordComplexityRules) *SetPasswordComplexityConfigurationRequest {
	s.PasswordComplexityRules = v
	return s
}

func (s *SetPasswordComplexityConfigurationRequest) SetPasswordMinLength(v int32) *SetPasswordComplexityConfigurationRequest {
	s.PasswordMinLength = &v
	return s
}

func (s *SetPasswordComplexityConfigurationRequest) Validate() error {
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

type SetPasswordComplexityConfigurationRequestPasswordComplexityRules struct {
	// The type of the password check. Valid values:
	//
	// 	- inclusion_upper_case: The password must contain uppercase letters.
	//
	// 	- inclusion_lower_case: The password must contain lowercase letters.
	//
	// 	- inclusion_special_case: The password must contain one or more of the following special characters: @ % + \\ / \\" ! # $ ^ ? : , ( ) { } [ ] ~ - _ .
	//
	// 	- inclusion_number: The password must contain digits.
	//
	// 	- exclusion_username: The password cannot contain a username.
	//
	// 	- exclusion_email: The password cannot contain an email prefix.
	//
	// 	- exclusion_phone_number: The password cannot contain a mobile number.
	//
	// 	- exclusion_display_name: The password cannot contain a display name.
	//
	// example:
	//
	// inclusion_upper_case
	PasswordCheckType *string `json:"PasswordCheckType,omitempty" xml:"PasswordCheckType,omitempty"`
}

func (s SetPasswordComplexityConfigurationRequestPasswordComplexityRules) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordComplexityConfigurationRequestPasswordComplexityRules) GoString() string {
	return s.String()
}

func (s *SetPasswordComplexityConfigurationRequestPasswordComplexityRules) GetPasswordCheckType() *string {
	return s.PasswordCheckType
}

func (s *SetPasswordComplexityConfigurationRequestPasswordComplexityRules) SetPasswordCheckType(v string) *SetPasswordComplexityConfigurationRequestPasswordComplexityRules {
	s.PasswordCheckType = &v
	return s
}

func (s *SetPasswordComplexityConfigurationRequestPasswordComplexityRules) Validate() error {
	return dara.Validate(s)
}
