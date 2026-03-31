// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetChatConfiguration(v *GetChatConfigurationResponseBodyChatConfiguration) *GetChatConfigurationResponseBody
	GetChatConfiguration() *GetChatConfigurationResponseBodyChatConfiguration
	SetRequestId(v string) *GetChatConfigurationResponseBody
	GetRequestId() *string
}

type GetChatConfigurationResponseBody struct {
	// example:
	//
	// {
	//
	//   "Name": "chatops",
	//
	//   "Type": "DingTalk",
	//
	//   "RamRole": "OOSServiceRole",
	//
	//   "Configuration": "{\\"DingTalkClientId\\": \\"dingpxlbxp0rgs7uxmtb\\", \\"DingTalkClientSecret\\": \\"******\\", \\"DingTalkTemplateId\\": \\"381c5aee-f8af-48a4-94be-cce587e42ea4.schema\\"}",
	//
	//   "ResourceGroupId": "s",
	//
	//   "CreatedDate": "2025-12-26T09:53:02+00:00",
	//
	//   "UpdatedDate": "2026-01-07T02:14:55+00:00",
	//
	//   "Outputs": "{\\"URL\\": \\"https://1407907063606569.appflow.aliyunnest.com/webhook/\\"}",
	//
	//   "RequestId": "4DB0****1234"
	//
	// }
	ChatConfiguration *GetChatConfigurationResponseBodyChatConfiguration `json:"ChatConfiguration,omitempty" xml:"ChatConfiguration,omitempty" type:"Struct"`
	// example:
	//
	// 4DB0****1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetChatConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *GetChatConfigurationResponseBody) GetChatConfiguration() *GetChatConfigurationResponseBodyChatConfiguration {
	return s.ChatConfiguration
}

func (s *GetChatConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatConfigurationResponseBody) SetChatConfiguration(v *GetChatConfigurationResponseBodyChatConfiguration) *GetChatConfigurationResponseBody {
	s.ChatConfiguration = v
	return s
}

func (s *GetChatConfigurationResponseBody) SetRequestId(v string) *GetChatConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetChatConfigurationResponseBody) Validate() error {
	if s.ChatConfiguration != nil {
		if err := s.ChatConfiguration.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetChatConfigurationResponseBodyChatConfiguration struct {
	// example:
	//
	// {"DingTalkClientId": "dingpxlbxp0rgs7uxmtb", "DingTalkClientSecret": "******", "DingTalkTemplateId": "381c5aee-f8af-48a4-94be-cce587e42ea4.schema"}
	Configuration *string `json:"Configuration,omitempty" xml:"Configuration,omitempty"`
	// example:
	//
	// 2025-12-26T09:53:02+00:00
	CreatedDate *string `json:"CreatedDate,omitempty" xml:"CreatedDate,omitempty"`
	// example:
	//
	// chatops
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// {"URL": "https://1407907063606569.appflow.aliyunnest.com/webhook/"}
	Outputs *string `json:"Outputs,omitempty" xml:"Outputs,omitempty"`
	// example:
	//
	// OOSServiceRole
	RamRole *string `json:"RamRole,omitempty" xml:"RamRole,omitempty"`
	// example:
	//
	// s
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// DingTalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2026-01-07T02:14:55+00:00
	UpdatedDate *string `json:"UpdatedDate,omitempty" xml:"UpdatedDate,omitempty"`
}

func (s GetChatConfigurationResponseBodyChatConfiguration) String() string {
	return dara.Prettify(s)
}

func (s GetChatConfigurationResponseBodyChatConfiguration) GoString() string {
	return s.String()
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) GetConfiguration() *string {
	return s.Configuration
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) GetCreatedDate() *string {
	return s.CreatedDate
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) GetName() *string {
	return s.Name
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) GetOutputs() *string {
	return s.Outputs
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) GetRamRole() *string {
	return s.RamRole
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) GetType() *string {
	return s.Type
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) GetUpdatedDate() *string {
	return s.UpdatedDate
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) SetConfiguration(v string) *GetChatConfigurationResponseBodyChatConfiguration {
	s.Configuration = &v
	return s
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) SetCreatedDate(v string) *GetChatConfigurationResponseBodyChatConfiguration {
	s.CreatedDate = &v
	return s
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) SetName(v string) *GetChatConfigurationResponseBodyChatConfiguration {
	s.Name = &v
	return s
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) SetOutputs(v string) *GetChatConfigurationResponseBodyChatConfiguration {
	s.Outputs = &v
	return s
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) SetRamRole(v string) *GetChatConfigurationResponseBodyChatConfiguration {
	s.RamRole = &v
	return s
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) SetResourceGroupId(v string) *GetChatConfigurationResponseBodyChatConfiguration {
	s.ResourceGroupId = &v
	return s
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) SetType(v string) *GetChatConfigurationResponseBodyChatConfiguration {
	s.Type = &v
	return s
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) SetUpdatedDate(v string) *GetChatConfigurationResponseBodyChatConfiguration {
	s.UpdatedDate = &v
	return s
}

func (s *GetChatConfigurationResponseBodyChatConfiguration) Validate() error {
	return dara.Validate(s)
}
