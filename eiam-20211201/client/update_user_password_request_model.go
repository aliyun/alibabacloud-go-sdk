// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserPasswordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UpdateUserPasswordRequest
	GetInstanceId() *string
	SetPassword(v string) *UpdateUserPasswordRequest
	GetPassword() *string
	SetPasswordForcedUpdateStatus(v string) *UpdateUserPasswordRequest
	GetPasswordForcedUpdateStatus() *string
	SetUserId(v string) *UpdateUserPasswordRequest
	GetUserId() *string
	SetUserNotificationChannels(v []*string) *UpdateUserPasswordRequest
	GetUserNotificationChannels() []*string
}

type UpdateUserPasswordRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new password of the account. For more information about the password format, see the "Password Policies" topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to enable forcible password change upon first logon. Default value: disabled. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordForcedUpdateStatus *string `json:"PasswordForcedUpdateStatus,omitempty" xml:"PasswordForcedUpdateStatus,omitempty"`
	// The account ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The methods for receiving password notifications.
	//
	// example:
	//
	// sms
	UserNotificationChannels []*string `json:"UserNotificationChannels,omitempty" xml:"UserNotificationChannels,omitempty" type:"Repeated"`
}

func (s UpdateUserPasswordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserPasswordRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserPasswordRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateUserPasswordRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateUserPasswordRequest) GetPasswordForcedUpdateStatus() *string {
	return s.PasswordForcedUpdateStatus
}

func (s *UpdateUserPasswordRequest) GetUserId() *string {
	return s.UserId
}

func (s *UpdateUserPasswordRequest) GetUserNotificationChannels() []*string {
	return s.UserNotificationChannels
}

func (s *UpdateUserPasswordRequest) SetInstanceId(v string) *UpdateUserPasswordRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserPasswordRequest) SetPassword(v string) *UpdateUserPasswordRequest {
	s.Password = &v
	return s
}

func (s *UpdateUserPasswordRequest) SetPasswordForcedUpdateStatus(v string) *UpdateUserPasswordRequest {
	s.PasswordForcedUpdateStatus = &v
	return s
}

func (s *UpdateUserPasswordRequest) SetUserId(v string) *UpdateUserPasswordRequest {
	s.UserId = &v
	return s
}

func (s *UpdateUserPasswordRequest) SetUserNotificationChannels(v []*string) *UpdateUserPasswordRequest {
	s.UserNotificationChannels = v
	return s
}

func (s *UpdateUserPasswordRequest) Validate() error {
	return dara.Validate(s)
}
