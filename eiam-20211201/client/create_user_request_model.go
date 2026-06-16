// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateUserRequest
	GetClientToken() *string
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
	SetInstanceId(v string) *CreateUserRequest
	GetInstanceId() *string
	SetOrganizationalUnitIds(v []*string) *CreateUserRequest
	GetOrganizationalUnitIds() []*string
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
	// A client-provided token to ensure the idempotence of the request. This value must be unique for each request. The token can contain only ASCII characters and must be no more than 64 characters long. For more information, see [How to ensure idempotence](~~~/doc-en/6a938a5b-2402-4c9d-b235-3733a1f813c9.dita).
	//
	// example:
	//
	// client-token-example
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The list of custom fields.
	//
	// example:
	//
	// description
	CustomFields []*CreateUserRequestCustomFields `json:"CustomFields,omitempty" xml:"CustomFields,omitempty" type:"Repeated"`
	// The description. The maximum length is 256 characters.
	//
	// example:
	//
	// description text
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name. The maximum length is 128 characters.
	//
	// example:
	//
	// name_001
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The email address. The local part of the address can contain uppercase letters, lowercase letters, digits, periods (.), underscores (_), or hyphens (-). The maximum length is 128 characters.
	//
	// example:
	//
	// example@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// Indicates whether the email address is verified. A verified address is considered trusted. This parameter is required if you specify the `Email` parameter. For typical use, set this to `true`.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"EmailVerified,omitempty" xml:"EmailVerified,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// A list of subordinate organizational unit IDs. A user can belong to multiple organizational units.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The password. The format must comply with the password policy.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The settings for the password initialization policy.
	PasswordInitializationConfig *CreateUserRequestPasswordInitializationConfig `json:"PasswordInitializationConfig,omitempty" xml:"PasswordInitializationConfig,omitempty" type:"Struct"`
	// The phone number. It must be 6 to 15 digits long.
	//
	// example:
	//
	// 12345678901
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// Indicates whether the phone number is verified. A verified number is considered trusted. This parameter is required if you specify the `PhoneNumber` parameter. For typical use, set this to `true`.
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"PhoneNumberVerified,omitempty" xml:"PhoneNumberVerified,omitempty"`
	// The country code. It must contain 1 to 6 digits and must not include the plus sign (+).
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// The ID of the primary organizational unit.
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	PrimaryOrganizationalUnitId *string `json:"PrimaryOrganizationalUnitId,omitempty" xml:"PrimaryOrganizationalUnitId,omitempty"`
	// The external ID for associating the user with an external system. The maximum length is 128 characters. If this parameter is not specified, its value defaults to the system-generated user ID.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The username. It can contain letters, digits, and the following special characters: underscores (_), periods (.), at signs (@), and hyphens (-). The maximum length is 256 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_001
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s CreateUserRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) GetClientToken() *string {
	return s.ClientToken
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

func (s *CreateUserRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserRequest) GetOrganizationalUnitIds() []*string {
	return s.OrganizationalUnitIds
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

func (s *CreateUserRequest) SetClientToken(v string) *CreateUserRequest {
	s.ClientToken = &v
	return s
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

func (s *CreateUserRequest) SetInstanceId(v string) *CreateUserRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateUserRequest) SetOrganizationalUnitIds(v []*string) *CreateUserRequest {
	s.OrganizationalUnitIds = v
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
	// The identifier of the custom field. You must create this field in the console before you can use it in a request. For more information, see the Custom Fields module in the console.
	//
	// example:
	//
	// age
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The value of the custom field. This value must comply with the constraints defined for the corresponding custom field.
	//
	// example:
	//
	// 10
	FieldValue *string `json:"FieldValue,omitempty" xml:"FieldValue,omitempty"`
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
	// The status of forced password update. This setting is disabled by default. Valid values:
	//
	// - `enabled`: Enables forced password updates.
	//
	// - `disabled`: Disables forced password updates.
	//
	// example:
	//
	// enabled
	PasswordForcedUpdateStatus *string `json:"PasswordForcedUpdateStatus,omitempty" xml:"PasswordForcedUpdateStatus,omitempty"`
	// The priority of the password initialization policy. This setting is disabled by default. Valid values:
	//
	// - `global`: The service uses the instance-level password initialization policy and ignores the policy specified in this request. For more information, see the password initialization policy settings.
	//
	// - `custom`: The service uses the password initialization policy defined in this request. This includes the forced password update setting, the password initialization method, and the notification channels.
	//
	// example:
	//
	// global
	PasswordInitializationPolicyPriority *string `json:"PasswordInitializationPolicyPriority,omitempty" xml:"PasswordInitializationPolicyPriority,omitempty"`
	// The password initialization type. Valid values:
	//
	// - `random`: The system generates a random password.
	//
	// example:
	//
	// random
	PasswordInitializationType *string `json:"PasswordInitializationType,omitempty" xml:"PasswordInitializationType,omitempty"`
	// The list of password notification channels.
	//
	// example:
	//
	// sms
	UserNotificationChannels []*string `json:"UserNotificationChannels,omitempty" xml:"UserNotificationChannels,omitempty" type:"Repeated"`
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
