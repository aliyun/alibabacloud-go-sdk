// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordExpirationConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPasswordExpirationConfiguration(v *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) *GetPasswordExpirationConfigurationResponseBody
	GetPasswordExpirationConfiguration() *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration
	SetRequestId(v string) *GetPasswordExpirationConfigurationResponseBody
	GetRequestId() *string
}

type GetPasswordExpirationConfigurationResponseBody struct {
	// The password expiration configurations.
	PasswordExpirationConfiguration *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration `json:"PasswordExpirationConfiguration,omitempty" xml:"PasswordExpirationConfiguration,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPasswordExpirationConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordExpirationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetPasswordExpirationConfigurationResponseBody) GetPasswordExpirationConfiguration() *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	return s.PasswordExpirationConfiguration
}

func (s *GetPasswordExpirationConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPasswordExpirationConfigurationResponseBody) SetPasswordExpirationConfiguration(v *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) *GetPasswordExpirationConfigurationResponseBody {
	s.PasswordExpirationConfiguration = v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBody) SetRequestId(v string) *GetPasswordExpirationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBody) Validate() error {
	if s.PasswordExpirationConfiguration != nil {
		if err := s.PasswordExpirationConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration struct {
	// The list of valid authentication IDs. The default is all ["ia_all"]
	//
	// ia_all: All. If you fill in this value, you cannot fill in other values
	//
	// ia_password: Account password login
	//
	// ia_otp_sms: SMS verification code login method
	//
	// ia_webauthn: WebAuthn authenticator login method
	//
	// idp_xxx: Specific identity provider authentication method
	EffectiveAuthenticationSourceIds []*string `json:"EffectiveAuthenticationSourceIds,omitempty" xml:"EffectiveAuthenticationSourceIds,omitempty" type:"Repeated"`
	// The action to take when a password expires. Valid values:
	//
	// 	- forbid_login: Prohibit the user from using the password to log on to IDaaS.
	//
	// 	- force_update_password: Force the user to change the password.
	//
	// 	- remind_update_password: Remind the user to change the password.
	//
	// example:
	//
	// forbid_login
	PasswordExpirationAction *string `json:"PasswordExpirationAction,omitempty" xml:"PasswordExpirationAction,omitempty"`
	// The methods for receiving password expiration notifications.
	//
	// example:
	//
	// login
	PasswordExpirationNotificationChannels []*string `json:"PasswordExpirationNotificationChannels,omitempty" xml:"PasswordExpirationNotificationChannels,omitempty" type:"Repeated"`
	// The number of days before the expiration date during which password expiration notifications are sent. Unit: day.
	//
	// example:
	//
	// 7
	PasswordExpirationNotificationDuration *int32 `json:"PasswordExpirationNotificationDuration,omitempty" xml:"PasswordExpirationNotificationDuration,omitempty"`
	// Indicates whether the password expiration notification feature is enabled. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordExpirationNotificationStatus *string `json:"PasswordExpirationNotificationStatus,omitempty" xml:"PasswordExpirationNotificationStatus,omitempty"`
	// Indicates whether the password expiration feature is enabled. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordExpirationStatus *string `json:"PasswordExpirationStatus,omitempty" xml:"PasswordExpirationStatus,omitempty"`
	// The number of days before which users must change the password to prevent password expiration. Unit: day.
	//
	// example:
	//
	// 3
	PasswordForcedUpdateDuration *int32 `json:"PasswordForcedUpdateDuration,omitempty" xml:"PasswordForcedUpdateDuration,omitempty"`
	// The validity period of a password. Unit: day.
	//
	// example:
	//
	// 180
	PasswordValidMaxDay *int32 `json:"PasswordValidMaxDay,omitempty" xml:"PasswordValidMaxDay,omitempty"`
}

func (s GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) GoString() string {
	return s.String()
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) GetEffectiveAuthenticationSourceIds() []*string {
	return s.EffectiveAuthenticationSourceIds
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) GetPasswordExpirationAction() *string {
	return s.PasswordExpirationAction
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) GetPasswordExpirationNotificationChannels() []*string {
	return s.PasswordExpirationNotificationChannels
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) GetPasswordExpirationNotificationDuration() *int32 {
	return s.PasswordExpirationNotificationDuration
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) GetPasswordExpirationNotificationStatus() *string {
	return s.PasswordExpirationNotificationStatus
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) GetPasswordExpirationStatus() *string {
	return s.PasswordExpirationStatus
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) GetPasswordForcedUpdateDuration() *int32 {
	return s.PasswordForcedUpdateDuration
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) GetPasswordValidMaxDay() *int32 {
	return s.PasswordValidMaxDay
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetEffectiveAuthenticationSourceIds(v []*string) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.EffectiveAuthenticationSourceIds = v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordExpirationAction(v string) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordExpirationAction = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordExpirationNotificationChannels(v []*string) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordExpirationNotificationChannels = v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordExpirationNotificationDuration(v int32) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordExpirationNotificationDuration = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordExpirationNotificationStatus(v string) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordExpirationNotificationStatus = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordExpirationStatus(v string) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordExpirationStatus = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordForcedUpdateDuration(v int32) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordForcedUpdateDuration = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) SetPasswordValidMaxDay(v int32) *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration {
	s.PasswordValidMaxDay = &v
	return s
}

func (s *GetPasswordExpirationConfigurationResponseBodyPasswordExpirationConfiguration) Validate() error {
	return dara.Validate(s)
}
