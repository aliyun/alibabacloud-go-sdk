// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetForgetPasswordConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpenForgetPasswordConfiguration(v *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) *GetForgetPasswordConfigurationResponseBody
	GetOpenForgetPasswordConfiguration() *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration
	SetRequestId(v string) *GetForgetPasswordConfigurationResponseBody
	GetRequestId() *string
}

type GetForgetPasswordConfigurationResponseBody struct {
	// The forgot password configurations.
	OpenForgetPasswordConfiguration *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration `json:"OpenForgetPasswordConfiguration,omitempty" xml:"OpenForgetPasswordConfiguration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetForgetPasswordConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetForgetPasswordConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetForgetPasswordConfigurationResponseBody) GetOpenForgetPasswordConfiguration() *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration {
	return s.OpenForgetPasswordConfiguration
}

func (s *GetForgetPasswordConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetForgetPasswordConfigurationResponseBody) SetOpenForgetPasswordConfiguration(v *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) *GetForgetPasswordConfigurationResponseBody {
	s.OpenForgetPasswordConfiguration = v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBody) SetRequestId(v string) *GetForgetPasswordConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration struct {
	// The authentication channels. Valid values:
	//
	// email
	//
	// sms
	//
	// totp
	//
	// web_authn
	AuthenticationChannels []*string `json:"AuthenticationChannels,omitempty" xml:"AuthenticationChannels,omitempty" type:"Repeated"`
	// Indicates whether the forgot password feature is enabled.
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Indicates whether email authentication is enabled for the forgot password feature.
	//
	// example:
	//
	// true
	EnableEmail *bool `json:"EnableEmail,omitempty" xml:"EnableEmail,omitempty"`
	// Indicates whether Short Message Service (SMS) authentication is enabled for the forgot password feature.
	//
	// example:
	//
	// true
	EnableSms *bool `json:"EnableSms,omitempty" xml:"EnableSms,omitempty"`
	// The status of the forgot password feature. Valid values: enabled and disabled.
	//
	// example:
	//
	// enabled
	ForgetPasswordStatus *string `json:"ForgetPasswordStatus,omitempty" xml:"ForgetPasswordStatus,omitempty"`
}

func (s GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) GoString() string {
	return s.String()
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) GetAuthenticationChannels() []*string {
	return s.AuthenticationChannels
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) GetEnable() *bool {
	return s.Enable
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) GetEnableEmail() *bool {
	return s.EnableEmail
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) GetEnableSms() *bool {
	return s.EnableSms
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) GetForgetPasswordStatus() *string {
	return s.ForgetPasswordStatus
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) SetAuthenticationChannels(v []*string) *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration {
	s.AuthenticationChannels = v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) SetEnable(v bool) *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration {
	s.Enable = &v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) SetEnableEmail(v bool) *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration {
	s.EnableEmail = &v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) SetEnableSms(v bool) *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration {
	s.EnableSms = &v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) SetForgetPasswordStatus(v string) *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration {
	s.ForgetPasswordStatus = &v
	return s
}

func (s *GetForgetPasswordConfigurationResponseBodyOpenForgetPasswordConfiguration) Validate() error {
	return dara.Validate(s)
}
