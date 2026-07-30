// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomFields(v []*CreateUserRequestCustomFields) *CreateUserRequest
	GetCustomFields() []*CreateUserRequestCustomFields
	SetDescription(v string) *CreateUserRequest
	GetDescription() *string
	SetDisplayName(v string) *CreateUserRequest
	GetDisplayName() *string
	SetEmail(v string) *CreateUserRequest
	GetEmail() *string
	SetEmailVerified(v bool) *CreateUserRequest
	GetEmailVerified() *bool
	SetPassword(v string) *CreateUserRequest
	GetPassword() *string
	SetPasswordInitializationConfig(v *CreateUserRequestPasswordInitializationConfig) *CreateUserRequest
	GetPasswordInitializationConfig() *CreateUserRequestPasswordInitializationConfig
	SetPhoneNumber(v string) *CreateUserRequest
	GetPhoneNumber() *string
	SetPhoneNumberVerified(v bool) *CreateUserRequest
	GetPhoneNumberVerified() *bool
	SetPhoneRegion(v string) *CreateUserRequest
	GetPhoneRegion() *string
	SetPrimaryOrganizationalUnitId(v string) *CreateUserRequest
	GetPrimaryOrganizationalUnitId() *string
	SetUserExternalId(v string) *CreateUserRequest
	GetUserExternalId() *string
	SetUsername(v string) *CreateUserRequest
	GetUsername() *string
}

type CreateUserRequest struct {
	// A list of custom fields for the account.
	CustomFields []*CreateUserRequestCustomFields `json:"customFields,omitempty" xml:"customFields,omitempty" type:"Repeated"`
	// The account description. The maximum length is 256 characters.
	//
	// example:
	//
	// 测试账户
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The display name. The maximum length is 128 characters.
	//
	// example:
	//
	// display_name001
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The email address. The local-part of the address can contain uppercase and lowercase letters, digits, periods (`.`), underscores (`_`), and hyphens (`-`). The maximum length is 128 characters.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// Specifies whether the email is verified. This parameter is required if `email` is set. Typically, set this to `true`.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"emailVerified,omitempty" xml:"emailVerified,omitempty"`
	// The account password. For password complexity rules, see the password policy in the IDaaS console.
	//
	// example:
	//
	// xxxxx
	Password *string `json:"password,omitempty" xml:"password,omitempty"`
	// The password initialization configuration.
	PasswordInitializationConfig *CreateUserRequestPasswordInitializationConfig `json:"passwordInitializationConfig,omitempty" xml:"passwordInitializationConfig,omitempty" type:"Struct"`
	// The account phone number. It must be 6 to 15 digits long.
	//
	// example:
	//
	// 156xxxxxxx
	PhoneNumber *string `json:"phoneNumber,omitempty" xml:"phoneNumber,omitempty"`
	// Specifies whether the phone number is verified. This parameter is required if `phoneNumber` is set. Typically, set this to `true`.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"phoneNumberVerified,omitempty" xml:"phoneNumberVerified,omitempty"`
	// The phone region code. For example, the code for the Chinese mainland is `86`. Do not include a `00` prefix or a plus sign (`+`). This parameter is required if `phoneNumber` is set.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"phoneRegion,omitempty" xml:"phoneRegion,omitempty"`
	// The ID of the primary organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	PrimaryOrganizationalUnitId *string `json:"primaryOrganizationalUnitId,omitempty" xml:"primaryOrganizationalUnitId,omitempty"`
	// The external user ID, used to associate the account with an external system. The maximum length is 128 characters. If unspecified, it defaults to the account ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"userExternalId,omitempty" xml:"userExternalId,omitempty"`
	// The username. It can contain letters, digits, and the following special characters: underscore (`_`), period (`.`), at sign (`@`), and hyphen (`-`). The maximum length is 256 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// name001
	Username *string `json:"username,omitempty" xml:"username,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetCustomFields() []*CreateUserRequestCustomFields {
	return s.CustomFields
}

func (s *CreateUserRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateUserRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateUserRequest) GetEmailVerified() *bool {
	return s.EmailVerified
}

func (s *CreateUserRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateUserRequest) GetPasswordInitializationConfig() *CreateUserRequestPasswordInitializationConfig {
	return s.PasswordInitializationConfig
}

func (s *CreateUserRequest) GetPhoneNumber() *string {
	return s.PhoneNumber
}

func (s *CreateUserRequest) GetPhoneNumberVerified() *bool {
	return s.PhoneNumberVerified
}

func (s *CreateUserRequest) GetPhoneRegion() *string {
	return s.PhoneRegion
}

func (s *CreateUserRequest) GetPrimaryOrganizationalUnitId() *string {
	return s.PrimaryOrganizationalUnitId
}

func (s *CreateUserRequest) GetUserExternalId() *string {
	return s.UserExternalId
}

func (s *CreateUserRequest) GetUsername() *string {
	return s.Username
}

func (s *CreateUserRequest) SetCustomFields(v []*CreateUserRequestCustomFields) *CreateUserRequest {
	s.CustomFields = v
	return s
}

func (s *CreateUserRequest) SetDescription(v string) *CreateUserRequest {
	s.Description = &v
	return s
}

func (s *CreateUserRequest) SetDisplayName(v string) *CreateUserRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetEmailVerified(v bool) *CreateUserRequest {
	s.EmailVerified = &v
	return s
}

func (s *CreateUserRequest) SetPassword(v string) *CreateUserRequest {
	s.Password = &v
	return s
}

func (s *CreateUserRequest) SetPasswordInitializationConfig(v *CreateUserRequestPasswordInitializationConfig) *CreateUserRequest {
	s.PasswordInitializationConfig = v
	return s
}

func (s *CreateUserRequest) SetPhoneNumber(v string) *CreateUserRequest {
	s.PhoneNumber = &v
	return s
}

func (s *CreateUserRequest) SetPhoneNumberVerified(v bool) *CreateUserRequest {
	s.PhoneNumberVerified = &v
	return s
}

func (s *CreateUserRequest) SetPhoneRegion(v string) *CreateUserRequest {
	s.PhoneRegion = &v
	return s
}

func (s *CreateUserRequest) SetPrimaryOrganizationalUnitId(v string) *CreateUserRequest {
	s.PrimaryOrganizationalUnitId = &v
	return s
}

func (s *CreateUserRequest) SetUserExternalId(v string) *CreateUserRequest {
	s.UserExternalId = &v
	return s
}

func (s *CreateUserRequest) SetUsername(v string) *CreateUserRequest {
	s.Username = &v
	return s
}

func (s *CreateUserRequest) Validate() error {
	if s.CustomFields != nil {
		for _, item := range s.CustomFields {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PasswordInitializationConfig != nil {
		if err := s.PasswordInitializationConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserRequestCustomFields struct {
	// The name of the custom field. You can view the field\\"s data type and value range in the IDaaS console.
	//
	// example:
	//
	// age
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// The value of the custom field.
	//
	// example:
	//
	// fieldValue_001
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
}

func (s CreateUserRequestCustomFields) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequestCustomFields) GoString() string {
	return s.String()
}

func (s *CreateUserRequestCustomFields) GetFieldName() *string {
	return s.FieldName
}

func (s *CreateUserRequestCustomFields) GetFieldValue() *string {
	return s.FieldValue
}

func (s *CreateUserRequestCustomFields) SetFieldName(v string) *CreateUserRequestCustomFields {
	s.FieldName = &v
	return s
}

func (s *CreateUserRequestCustomFields) SetFieldValue(v string) *CreateUserRequestCustomFields {
	s.FieldValue = &v
	return s
}

func (s *CreateUserRequestCustomFields) Validate() error {
	return dara.Validate(s)
}

type CreateUserRequestPasswordInitializationConfig struct {
	// The password forced update status. By default, this feature is disabled. Valid values:
	//
	// - `enabled`: Enables the feature.
	//
	// - `disabled`: Disables the feature.
	//
	// example:
	//
	// enabled
	PasswordForcedUpdateStatus *string `json:"passwordForcedUpdateStatus,omitempty" xml:"passwordForcedUpdateStatus,omitempty"`
	// The priority of the password initialization policy. Valid values:
	//
	// - `global`: Uses the instance-level password initialization policy and ignores the custom settings in this request. For more information, see the password initialization policy configuration in the IDaaS console.
	//
	// - `custom`: Uses the custom password initialization policy defined in this request. This includes settings for forced password updates, the initialization type, and notification channels.
	//
	// example:
	//
	// global
	PasswordInitializationPolicyPriority *string `json:"passwordInitializationPolicyPriority,omitempty" xml:"passwordInitializationPolicyPriority,omitempty"`
	// The password initialization type. Valid values:
	//
	// - `random`: A randomly generated password.
	//
	// example:
	//
	// random
	PasswordInitializationType *string `json:"passwordInitializationType,omitempty" xml:"passwordInitializationType,omitempty"`
	// The user notification channels. Valid values:
	//
	// - `email`: Email
	//
	// - `sms`: SMS
	//
	// example:
	//
	// sms
	UserNotificationChannels []*string `json:"userNotificationChannels,omitempty" xml:"userNotificationChannels,omitempty" type:"Repeated"`
}

func (s CreateUserRequestPasswordInitializationConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequestPasswordInitializationConfig) GoString() string {
	return s.String()
}

func (s *CreateUserRequestPasswordInitializationConfig) GetPasswordForcedUpdateStatus() *string {
	return s.PasswordForcedUpdateStatus
}

func (s *CreateUserRequestPasswordInitializationConfig) GetPasswordInitializationPolicyPriority() *string {
	return s.PasswordInitializationPolicyPriority
}

func (s *CreateUserRequestPasswordInitializationConfig) GetPasswordInitializationType() *string {
	return s.PasswordInitializationType
}

func (s *CreateUserRequestPasswordInitializationConfig) GetUserNotificationChannels() []*string {
	return s.UserNotificationChannels
}

func (s *CreateUserRequestPasswordInitializationConfig) SetPasswordForcedUpdateStatus(v string) *CreateUserRequestPasswordInitializationConfig {
	s.PasswordForcedUpdateStatus = &v
	return s
}

func (s *CreateUserRequestPasswordInitializationConfig) SetPasswordInitializationPolicyPriority(v string) *CreateUserRequestPasswordInitializationConfig {
	s.PasswordInitializationPolicyPriority = &v
	return s
}

func (s *CreateUserRequestPasswordInitializationConfig) SetPasswordInitializationType(v string) *CreateUserRequestPasswordInitializationConfig {
	s.PasswordInitializationType = &v
	return s
}

func (s *CreateUserRequestPasswordInitializationConfig) SetUserNotificationChannels(v []*string) *CreateUserRequestPasswordInitializationConfig {
	s.UserNotificationChannels = v
	return s
}

func (s *CreateUserRequestPasswordInitializationConfig) Validate() error {
	return dara.Validate(s)
}
