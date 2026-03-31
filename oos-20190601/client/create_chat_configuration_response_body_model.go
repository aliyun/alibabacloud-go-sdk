// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateChatConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatConfiguration(v *CreateChatConfigurationResponseBodyChatConfiguration) *CreateChatConfigurationResponseBody
	GetChatConfiguration() *CreateChatConfigurationResponseBodyChatConfiguration
	SetRequestId(v string) *CreateChatConfigurationResponseBody
	GetRequestId() *string
}

type CreateChatConfigurationResponseBody struct {
	// example:
	//
	// {"Name": "TestChatConfig-9be97b40", "CreatedDate": "2025-12-11T13:49:10+00:00", "UpdatedDate": "2025-12-11T13:59:02+00:00"}
	ChatConfiguration *CreateChatConfigurationResponseBodyChatConfiguration `json:"ChatConfiguration,omitempty" xml:"ChatConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 4DB0****1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateChatConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateChatConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateChatConfigurationResponseBody) GetChatConfiguration() *CreateChatConfigurationResponseBodyChatConfiguration {
	return s.ChatConfiguration
}

func (s *CreateChatConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateChatConfigurationResponseBody) SetChatConfiguration(v *CreateChatConfigurationResponseBodyChatConfiguration) *CreateChatConfigurationResponseBody {
	s.ChatConfiguration = v
	return s
}

func (s *CreateChatConfigurationResponseBody) SetRequestId(v string) *CreateChatConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateChatConfigurationResponseBody) Validate() error {
	if s.ChatConfiguration != nil {
		if err := s.ChatConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateChatConfigurationResponseBodyChatConfiguration struct {
	// example:
	//
	// 2025-12-11T13:49:10+00:00
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// example:
	//
	// TestChatConfig-9be97b40
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2025-12-11T13:59:02+00:00
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s CreateChatConfigurationResponseBodyChatConfiguration) String() string {
	return dara.Prettify(s)
}

func (s CreateChatConfigurationResponseBodyChatConfiguration) GoString() string {
	return s.String()
}

func (s *CreateChatConfigurationResponseBodyChatConfiguration) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *CreateChatConfigurationResponseBodyChatConfiguration) GetName() *string {
	return s.Name
}

func (s *CreateChatConfigurationResponseBodyChatConfiguration) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *CreateChatConfigurationResponseBodyChatConfiguration) SetCreatedDate(v string) *CreateChatConfigurationResponseBodyChatConfiguration {
	s.CreatedDate = &v
	return s
}

func (s *CreateChatConfigurationResponseBodyChatConfiguration) SetName(v string) *CreateChatConfigurationResponseBodyChatConfiguration {
	s.Name = &v
	return s
}

func (s *CreateChatConfigurationResponseBodyChatConfiguration) SetUpdatedDate(v string) *CreateChatConfigurationResponseBodyChatConfiguration {
	s.UpdatedDate = &v
	return s
}

func (s *CreateChatConfigurationResponseBodyChatConfiguration) Validate() error {
	return dara.Validate(s)
}
