// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetPasswordInitializationConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SetPasswordInitializationConfigurationRequest
	GetInstanceId() *string
	SetPasswordForcedUpdateStatus(v string) *SetPasswordInitializationConfigurationRequest
	GetPasswordForcedUpdateStatus() *string
	SetPasswordInitializationNotificationChannels(v []*string) *SetPasswordInitializationConfigurationRequest
	GetPasswordInitializationNotificationChannels() []*string
	SetPasswordInitializationStatus(v string) *SetPasswordInitializationConfigurationRequest
	GetPasswordInitializationStatus() *string
	SetPasswordInitializationType(v string) *SetPasswordInitializationConfigurationRequest
	GetPasswordInitializationType() *string
}

type SetPasswordInitializationConfigurationRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable forcible password change upon first logon. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordForcedUpdateStatus *string `json:"PasswordForcedUpdateStatus,omitempty" xml:"PasswordForcedUpdateStatus,omitempty"`
	// The methods for receiving password initialization notifications.
	//
	// example:
	//
	// email
	PasswordInitializationNotificationChannels []*string `json:"PasswordInitializationNotificationChannels,omitempty" xml:"PasswordInitializationNotificationChannels,omitempty" type:"Repeated"`
	// Specifies whether to enable password initialization. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	PasswordInitializationStatus *string `json:"PasswordInitializationStatus,omitempty" xml:"PasswordInitializationStatus,omitempty"`
	// The password initialization method. This parameter is required when PasswordInitializationStatus is set to enabled. Set the value to random.
	//
	// 	- random: A randomly generated password is used.
	//
	// example:
	//
	// random
	PasswordInitializationType *string `json:"PasswordInitializationType,omitempty" xml:"PasswordInitializationType,omitempty"`
}

func (s SetPasswordInitializationConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s SetPasswordInitializationConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetPasswordInitializationConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetPasswordInitializationConfigurationRequest) GetPasswordForcedUpdateStatus() *string {
	return s.PasswordForcedUpdateStatus
}

func (s *SetPasswordInitializationConfigurationRequest) GetPasswordInitializationNotificationChannels() []*string {
	return s.PasswordInitializationNotificationChannels
}

func (s *SetPasswordInitializationConfigurationRequest) GetPasswordInitializationStatus() *string {
	return s.PasswordInitializationStatus
}

func (s *SetPasswordInitializationConfigurationRequest) GetPasswordInitializationType() *string {
	return s.PasswordInitializationType
}

func (s *SetPasswordInitializationConfigurationRequest) SetInstanceId(v string) *SetPasswordInitializationConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetPasswordInitializationConfigurationRequest) SetPasswordForcedUpdateStatus(v string) *SetPasswordInitializationConfigurationRequest {
	s.PasswordForcedUpdateStatus = &v
	return s
}

func (s *SetPasswordInitializationConfigurationRequest) SetPasswordInitializationNotificationChannels(v []*string) *SetPasswordInitializationConfigurationRequest {
	s.PasswordInitializationNotificationChannels = v
	return s
}

func (s *SetPasswordInitializationConfigurationRequest) SetPasswordInitializationStatus(v string) *SetPasswordInitializationConfigurationRequest {
	s.PasswordInitializationStatus = &v
	return s
}

func (s *SetPasswordInitializationConfigurationRequest) SetPasswordInitializationType(v string) *SetPasswordInitializationConfigurationRequest {
	s.PasswordInitializationType = &v
	return s
}

func (s *SetPasswordInitializationConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
