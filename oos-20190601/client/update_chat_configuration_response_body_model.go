// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateChatConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatConfiguration(v *UpdateChatConfigurationResponseBodyChatConfiguration) *UpdateChatConfigurationResponseBody
	GetChatConfiguration() *UpdateChatConfigurationResponseBodyChatConfiguration
	SetRequestId(v string) *UpdateChatConfigurationResponseBody
	GetRequestId() *string
}

type UpdateChatConfigurationResponseBody struct {
	// example:
	//
	// {"Name": "dingtalk", "CreatedDate": "2025-12-23T10:14:28+00:00", "UpdatedDate": "2025-12-23T10:16:37.178097"}
	ChatConfiguration *UpdateChatConfigurationResponseBodyChatConfiguration `json:"ChatConfiguration,omitempty" xml:"ChatConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 4DB0****1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateChatConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateChatConfigurationResponseBody) GetChatConfiguration() *UpdateChatConfigurationResponseBodyChatConfiguration {
	return s.ChatConfiguration
}

func (s *UpdateChatConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateChatConfigurationResponseBody) SetChatConfiguration(v *UpdateChatConfigurationResponseBodyChatConfiguration) *UpdateChatConfigurationResponseBody {
	s.ChatConfiguration = v
	return s
}

func (s *UpdateChatConfigurationResponseBody) SetRequestId(v string) *UpdateChatConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateChatConfigurationResponseBody) Validate() error {
	if s.ChatConfiguration != nil {
		if err := s.ChatConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateChatConfigurationResponseBodyChatConfiguration struct {
	// example:
	//
	// 2025-12-23T10:14:28+00:00
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// example:
	//
	// dingtalk
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2025-12-23T10:16:37.178097
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s UpdateChatConfigurationResponseBodyChatConfiguration) String() string {
	return dara.Prettify(s)
}

func (s UpdateChatConfigurationResponseBodyChatConfiguration) GoString() string {
	return s.String()
}

func (s *UpdateChatConfigurationResponseBodyChatConfiguration) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *UpdateChatConfigurationResponseBodyChatConfiguration) GetName() *string {
	return s.Name
}

func (s *UpdateChatConfigurationResponseBodyChatConfiguration) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *UpdateChatConfigurationResponseBodyChatConfiguration) SetCreatedDate(v string) *UpdateChatConfigurationResponseBodyChatConfiguration {
	s.CreatedDate = &v
	return s
}

func (s *UpdateChatConfigurationResponseBodyChatConfiguration) SetName(v string) *UpdateChatConfigurationResponseBodyChatConfiguration {
	s.Name = &v
	return s
}

func (s *UpdateChatConfigurationResponseBodyChatConfiguration) SetUpdatedDate(v string) *UpdateChatConfigurationResponseBodyChatConfiguration {
	s.UpdatedDate = &v
	return s
}

func (s *UpdateChatConfigurationResponseBodyChatConfiguration) Validate() error {
	return dara.Validate(s)
}
