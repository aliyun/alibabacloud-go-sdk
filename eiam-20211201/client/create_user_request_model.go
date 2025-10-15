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
	// The password initialization method. Set the value to random,
	//
	// 	- which indicates that the password is randomly generated.
	//
	// example:
	//
	// client-token-demo
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The custom extended fields.
	//
	// example:
	//
	// description
	CustomFields []*CreateUserRequestCustomFields `json:"CustomFields,omitempty" xml:"CustomFields,omitempty" type:"Repeated"`
	// The description of the organization. The value can be up to 256 characters in length.
	//
	// example:
	//
	// description text
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the account. The display name can be up to 64 characters in length.
	//
	// example:
	//
	// name_001
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The name of the account. The name can be up to 64 characters in length and can contain letters, digits, underscores (_), periods (.), at signs (@), and hyphens (-).
	//
	// example:
	//
	// example@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The description of the account. The description can be up to 256 characters in length.
	//
	// example:
	//
	// true
	EmailVerified *bool `json:"EmailVerified,omitempty" xml:"EmailVerified,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether the mobile phone number is a trusted mobile phone number. This parameter is required if the PhoneNumber parameter is specified. If you have no special business requirements, set this parameter to true.
	OrganizationalUnitIds []*string `json:"OrganizationalUnitIds,omitempty" xml:"OrganizationalUnitIds,omitempty" type:"Repeated"`
	// The password of the account. For more information,view the password policyof the instanceinthe IDaaS console.
	//
	// example:
	//
	// 123456xxx
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The configurations for password initialization.
	PasswordInitializationConfig *CreateUserRequestPasswordInitializationConfig `json:"PasswordInitializationConfig,omitempty" xml:"PasswordInitializationConfig,omitempty" type:"Struct"`
	// The ID of the account.
	//
	// example:
	//
	// 12345678901xxx
	PhoneNumber *string `json:"PhoneNumber,omitempty" xml:"PhoneNumber,omitempty"`
	// The email address of the user who owns the account. The email address prefix can contain letters, digits, underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// true
	PhoneNumberVerified *bool `json:"PhoneNumberVerified,omitempty" xml:"PhoneNumberVerified,omitempty"`
	// The IDs of organizational units to which the account belongs. An account can belong to multiple organizational units.
	//
	// example:
	//
	// 86
	PhoneRegion *string `json:"PhoneRegion,omitempty" xml:"PhoneRegion,omitempty"`
	// 主组织ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// ou_wovwffm62xifdziem7an7xxxxx
	PrimaryOrganizationalUnitId *string `json:"PrimaryOrganizationalUnitId,omitempty" xml:"PrimaryOrganizationalUnitId,omitempty"`
	// The display name of the account. The display name can be up to 64 characters in length.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserExternalId *string `json:"UserExternalId,omitempty" xml:"UserExternalId,omitempty"`
	// The name of the extended field. You must create the extended field in advance. To create an extended field, log on to the IDaaS console. In the left-side navigation pane, choose Accounts > Extended Fields, and then click Create Field on the Extended Fields page.
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
	// The name of the extended field. You must create the extended field in advance. To create an extended field, log on to the IDaaS console. In the left-side navigation pane, choose Accounts > Extended Fields, and then click Create Field on the Extended Fields page.
	//
	// example:
	//
	// age
	FieldName *string `json:"FieldName,omitempty" xml:"FieldName,omitempty"`
	// The value of the extended field. The value follows the limits on the properties of the extended field.
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
	// Specifies whether to forcibly change the password status. Default value: disabled. Valid values:
	//
	// 	- enabled:forcibly changes the password status.
	//
	// 	- disabled: does not forcibly change the password status.
	//
	// example:
	//
	// enabled
	PasswordForcedUpdateStatus *string `json:"PasswordForcedUpdateStatus,omitempty" xml:"PasswordForcedUpdateStatus,omitempty"`
	// The priority of the password initialization policy. By default, this parameter does not take effect. Valid values:
	//
	// 	- global:The password initialization policy globally takes effect.
	//
	// 	- custom: The password initialization policy takes effect based on custom settings.
	//
	// example:
	//
	// global
	PasswordInitializationPolicyPriority *string `json:"PasswordInitializationPolicyPriority,omitempty" xml:"PasswordInitializationPolicyPriority,omitempty"`
	// The password initialization method. Set the value to random,
	//
	// 	- whichindicates that the password is randomly generated.
	//
	// example:
	//
	// random
	PasswordInitializationType *string `json:"PasswordInitializationType,omitempty" xml:"PasswordInitializationType,omitempty"`
	// The value of the extended field. The value follows the limits on the properties of the extended field.
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
