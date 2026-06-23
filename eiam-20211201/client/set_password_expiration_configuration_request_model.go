// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordExpirationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEffectiveAuthenticationSourceIds(v []*string) *SetPasswordExpirationConfigurationRequest
	GetEffectiveAuthenticationSourceIds() []*string
	SetInstanceId(v string) *SetPasswordExpirationConfigurationRequest
	GetInstanceId() *string
	SetPasswordExpirationAction(v string) *SetPasswordExpirationConfigurationRequest
	GetPasswordExpirationAction() *string
	SetPasswordExpirationNotificationChannels(v []*string) *SetPasswordExpirationConfigurationRequest
	GetPasswordExpirationNotificationChannels() []*string
	SetPasswordExpirationNotificationDuration(v int32) *SetPasswordExpirationConfigurationRequest
	GetPasswordExpirationNotificationDuration() *int32
	SetPasswordExpirationNotificationStatus(v string) *SetPasswordExpirationConfigurationRequest
	GetPasswordExpirationNotificationStatus() *string
	SetPasswordExpirationStatus(v string) *SetPasswordExpirationConfigurationRequest
	GetPasswordExpirationStatus() *string
	SetPasswordForcedUpdateDuration(v int32) *SetPasswordExpirationConfigurationRequest
	GetPasswordForcedUpdateDuration() *int32
	SetPasswordValidMaxDay(v int32) *SetPasswordExpirationConfigurationRequest
	GetPasswordValidMaxDay() *int32
}

type SetPasswordExpirationConfigurationRequest struct {
	// The list of effective authentication source IDs.
	EffectiveAuthenticationSourceIds []*string `json:"EffectiveAuthenticationSourceIds,omitempty" xml:"EffectiveAuthenticationSourceIds,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The action to take when a password expires. This parameter is required when PasswordExpirationStatus is set to enabled. Valid values:
	//
	// - forbid_login: Prohibit logon.
	//
	// - force_update_password: Force password change.
	//
	// - remind_update_password: Remind to change password.
	//
	// example:
	//
	// force_update_password
	PasswordExpirationAction *string `json:"PasswordExpirationAction,omitempty" xml:"PasswordExpirationAction,omitempty"`
	// The list of password expiration notification channels. This parameter is required when PasswordExpirationNotificationStatus is set to enabled.
	//
	// example:
	//
	// login
	PasswordExpirationNotificationChannels []*string `json:"PasswordExpirationNotificationChannels,omitempty" xml:"PasswordExpirationNotificationChannels,omitempty" type:"Repeated"`
	// The password expiration notification time, in days. This parameter is required when PasswordExpirationNotificationStatus is set to enabled.
	//
	// example:
	//
	// 7
	PasswordExpirationNotificationDuration *int32 `json:"PasswordExpirationNotificationDuration,omitempty" xml:"PasswordExpirationNotificationDuration,omitempty"`
	// The password expiration notification status. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enabled
	PasswordExpirationNotificationStatus *string `json:"PasswordExpirationNotificationStatus,omitempty" xml:"PasswordExpirationNotificationStatus,omitempty"`
	// The password expiration configuration status. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	PasswordExpirationStatus *string `json:"PasswordExpirationStatus,omitempty" xml:"PasswordExpirationStatus,omitempty"`
	// The forced password change time, in days. The value of this parameter must be greater than the value of PasswordExpirationNotificationDuration.
	//
	// example:
	//
	// 7
	PasswordForcedUpdateDuration *int32 `json:"PasswordForcedUpdateDuration,omitempty" xml:"PasswordForcedUpdateDuration,omitempty"`
	// The validity period of a password, in days. This parameter is required when PasswordExpirationStatus is set to enabled.
	//
	// example:
	//
	// 180
	PasswordValidMaxDay *int32 `json:"PasswordValidMaxDay,omitempty" xml:"PasswordValidMaxDay,omitempty"`
}

func (s SetPasswordExpirationConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordExpirationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordExpirationConfigurationRequest) GetEffectiveAuthenticationSourceIds() []*string {
	return s.EffectiveAuthenticationSourceIds
}

func (s *SetPasswordExpirationConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPasswordExpirationConfigurationRequest) GetPasswordExpirationAction() *string {
	return s.PasswordExpirationAction
}

func (s *SetPasswordExpirationConfigurationRequest) GetPasswordExpirationNotificationChannels() []*string {
	return s.PasswordExpirationNotificationChannels
}

func (s *SetPasswordExpirationConfigurationRequest) GetPasswordExpirationNotificationDuration() *int32 {
	return s.PasswordExpirationNotificationDuration
}

func (s *SetPasswordExpirationConfigurationRequest) GetPasswordExpirationNotificationStatus() *string {
	return s.PasswordExpirationNotificationStatus
}

func (s *SetPasswordExpirationConfigurationRequest) GetPasswordExpirationStatus() *string {
	return s.PasswordExpirationStatus
}

func (s *SetPasswordExpirationConfigurationRequest) GetPasswordForcedUpdateDuration() *int32 {
	return s.PasswordForcedUpdateDuration
}

func (s *SetPasswordExpirationConfigurationRequest) GetPasswordValidMaxDay() *int32 {
	return s.PasswordValidMaxDay
}

func (s *SetPasswordExpirationConfigurationRequest) SetEffectiveAuthenticationSourceIds(v []*string) *SetPasswordExpirationConfigurationRequest {
	s.EffectiveAuthenticationSourceIds = v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetInstanceId(v string) *SetPasswordExpirationConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordExpirationAction(v string) *SetPasswordExpirationConfigurationRequest {
	s.PasswordExpirationAction = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordExpirationNotificationChannels(v []*string) *SetPasswordExpirationConfigurationRequest {
	s.PasswordExpirationNotificationChannels = v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordExpirationNotificationDuration(v int32) *SetPasswordExpirationConfigurationRequest {
	s.PasswordExpirationNotificationDuration = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordExpirationNotificationStatus(v string) *SetPasswordExpirationConfigurationRequest {
	s.PasswordExpirationNotificationStatus = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordExpirationStatus(v string) *SetPasswordExpirationConfigurationRequest {
	s.PasswordExpirationStatus = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordForcedUpdateDuration(v int32) *SetPasswordExpirationConfigurationRequest {
	s.PasswordForcedUpdateDuration = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) SetPasswordValidMaxDay(v int32) *SetPasswordExpirationConfigurationRequest {
	s.PasswordValidMaxDay = &v
	return s
}

func (s *SetPasswordExpirationConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
