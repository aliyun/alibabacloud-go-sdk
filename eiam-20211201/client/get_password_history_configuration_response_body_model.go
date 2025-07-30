// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPasswordHistoryConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPasswordHistoryConfiguration(v *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) *GetPasswordHistoryConfigurationResponseBody
	GetPasswordHistoryConfiguration() *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration
	SetRequestId(v string) *GetPasswordHistoryConfigurationResponseBody
	GetRequestId() *string
}

type GetPasswordHistoryConfigurationResponseBody struct {
	// The password history configurations.
	PasswordHistoryConfiguration *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration `json:"PasswordHistoryConfiguration,omitempty" xml:"PasswordHistoryConfiguration,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPasswordHistoryConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordHistoryConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetPasswordHistoryConfigurationResponseBody) GetPasswordHistoryConfiguration() *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration {
	return s.PasswordHistoryConfiguration
}

func (s *GetPasswordHistoryConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPasswordHistoryConfigurationResponseBody) SetPasswordHistoryConfiguration(v *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) *GetPasswordHistoryConfigurationResponseBody {
	s.PasswordHistoryConfiguration = v
	return s
}

func (s *GetPasswordHistoryConfigurationResponseBody) SetRequestId(v string) *GetPasswordHistoryConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPasswordHistoryConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration struct {
	// The maximum number of recent passwords that are retained.
	//
	// example:
	//
	// 3
	PasswordHistoryMaxRetention *int32 `json:"PasswordHistoryMaxRetention,omitempty" xml:"PasswordHistoryMaxRetention,omitempty"`
	// Indicates whether the password history feature is enabled. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// example:
	//
	// enabled
	PasswordHistoryStatus *string `json:"PasswordHistoryStatus,omitempty" xml:"PasswordHistoryStatus,omitempty"`
}

func (s GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) GoString() string {
	return s.String()
}

func (s *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) GetPasswordHistoryMaxRetention() *int32 {
	return s.PasswordHistoryMaxRetention
}

func (s *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) GetPasswordHistoryStatus() *string {
	return s.PasswordHistoryStatus
}

func (s *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) SetPasswordHistoryMaxRetention(v int32) *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration {
	s.PasswordHistoryMaxRetention = &v
	return s
}

func (s *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) SetPasswordHistoryStatus(v string) *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration {
	s.PasswordHistoryStatus = &v
	return s
}

func (s *GetPasswordHistoryConfigurationResponseBodyPasswordHistoryConfiguration) Validate() error {
	return dara.Validate(s)
}
