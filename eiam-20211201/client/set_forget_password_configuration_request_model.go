// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetForgetPasswordConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthenticationChannels(v []*string) *SetForgetPasswordConfigurationRequest
	GetAuthenticationChannels() []*string
	SetForgetPasswordStatus(v string) *SetForgetPasswordConfigurationRequest
	GetForgetPasswordStatus() *string
	SetInstanceId(v string) *SetForgetPasswordConfigurationRequest
	GetInstanceId() *string
}

type SetForgetPasswordConfigurationRequest struct {
	// The authentication channels. Valid values: email and sms.
	//
	// example:
	//
	// email
	AuthenticationChannels []*string `json:"AuthenticationChannels,omitempty" xml:"AuthenticationChannels,omitempty" type:"Repeated"`
	// The status of the forgot password feature. Valid values: enabled and disabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	ForgetPasswordStatus *string `json:"ForgetPasswordStatus,omitempty" xml:"ForgetPasswordStatus,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// eiam-111ccc1111
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s SetForgetPasswordConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s SetForgetPasswordConfigurationRequest) GoString() string {
	return s.String()
}

func (s *SetForgetPasswordConfigurationRequest) GetAuthenticationChannels() []*string {
	return s.AuthenticationChannels
}

func (s *SetForgetPasswordConfigurationRequest) GetForgetPasswordStatus() *string {
	return s.ForgetPasswordStatus
}

func (s *SetForgetPasswordConfigurationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SetForgetPasswordConfigurationRequest) SetAuthenticationChannels(v []*string) *SetForgetPasswordConfigurationRequest {
	s.AuthenticationChannels = v
	return s
}

func (s *SetForgetPasswordConfigurationRequest) SetForgetPasswordStatus(v string) *SetForgetPasswordConfigurationRequest {
	s.ForgetPasswordStatus = &v
	return s
}

func (s *SetForgetPasswordConfigurationRequest) SetInstanceId(v string) *SetForgetPasswordConfigurationRequest {
	s.InstanceId = &v
	return s
}

func (s *SetForgetPasswordConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
