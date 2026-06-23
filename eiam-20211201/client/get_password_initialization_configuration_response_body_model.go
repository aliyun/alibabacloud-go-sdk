// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordInitializationConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPasswordInitializationConfiguration(v *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) *GetPasswordInitializationConfigurationResponseBody
	GetPasswordInitializationConfiguration() *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration
	SetRequestId(v string) *GetPasswordInitializationConfigurationResponseBody
	GetRequestId() *string
}

type GetPasswordInitializationConfigurationResponseBody struct {
	// The password initialization configuration.
	PasswordInitializationConfiguration *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration `json:"PasswordInitializationConfiguration,omitempty" xml:"PasswordInitializationConfiguration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPasswordInitializationConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordInitializationConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetPasswordInitializationConfigurationResponseBody) GetPasswordInitializationConfiguration() *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration {
	return s.PasswordInitializationConfiguration
}

func (s *GetPasswordInitializationConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPasswordInitializationConfigurationResponseBody) SetPasswordInitializationConfiguration(v *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) *GetPasswordInitializationConfigurationResponseBody {
	s.PasswordInitializationConfiguration = v
	return s
}

func (s *GetPasswordInitializationConfigurationResponseBody) SetRequestId(v string) *GetPasswordInitializationConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPasswordInitializationConfigurationResponseBody) Validate() error {
	if s.PasswordInitializationConfiguration != nil {
		if err := s.PasswordInitializationConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration struct {
	// The status of forced password update. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enabled
	PasswordForcedUpdateStatus *string `json:"PasswordForcedUpdateStatus,omitempty" xml:"PasswordForcedUpdateStatus,omitempty"`
	// The list of password initialization notification channels.
	//
	// example:
	//
	// email
	PasswordInitializationNotificationChannels []*string `json:"PasswordInitializationNotificationChannels,omitempty" xml:"PasswordInitializationNotificationChannels,omitempty" type:"Repeated"`
	// The status of the password initialization configuration. Valid values:
	//
	// - enabled: Enabled.
	//
	// - disabled: Disabled.
	//
	// example:
	//
	// enabled
	PasswordInitializationStatus *string `json:"PasswordInitializationStatus,omitempty" xml:"PasswordInitializationStatus,omitempty"`
	// The password initialization method. Valid values:
	//
	// - random: Random.
	//
	// example:
	//
	// random
	PasswordInitializationType *string `json:"PasswordInitializationType,omitempty" xml:"PasswordInitializationType,omitempty"`
}

func (s GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) GoString() string {
	return s.String()
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) GetPasswordForcedUpdateStatus() *string {
	return s.PasswordForcedUpdateStatus
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) GetPasswordInitializationNotificationChannels() []*string {
	return s.PasswordInitializationNotificationChannels
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) GetPasswordInitializationStatus() *string {
	return s.PasswordInitializationStatus
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) GetPasswordInitializationType() *string {
	return s.PasswordInitializationType
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) SetPasswordForcedUpdateStatus(v string) *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration {
	s.PasswordForcedUpdateStatus = &v
	return s
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) SetPasswordInitializationNotificationChannels(v []*string) *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration {
	s.PasswordInitializationNotificationChannels = v
	return s
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) SetPasswordInitializationStatus(v string) *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration {
	s.PasswordInitializationStatus = &v
	return s
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) SetPasswordInitializationType(v string) *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration {
	s.PasswordInitializationType = &v
	return s
}

func (s *GetPasswordInitializationConfigurationResponseBodyPasswordInitializationConfiguration) Validate() error {
	return dara.Validate(s)
}
