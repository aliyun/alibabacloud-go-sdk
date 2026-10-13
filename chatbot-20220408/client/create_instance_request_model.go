// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CreateInstanceRequest
	GetAgentKey() *string
	SetIntroduction(v string) *CreateInstanceRequest
	GetIntroduction() *string
	SetLanguageCode(v string) *CreateInstanceRequest
	GetLanguageCode() *string
	SetName(v string) *CreateInstanceRequest
	GetName() *string
	SetRobotType(v string) *CreateInstanceRequest
	GetRobotType() *string
}

type CreateInstanceRequest struct {
	// The workspace key. If this parameter is not specified, the default workspace is accessed. Obtain the key value from the business management page of the Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The description of the chatbot. The description can be up to 50 characters in length.
	//
	// example:
	//
	// A sample chatbot for the mobile app
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// The language of the chatbot service, such as zh-cn or en-us. Default value: the default language of the workspace. Valid values: the languages supported by the workspace.
	//
	// example:
	//
	// zh-cn
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	// The name of the chatbot. The name can be up to 50 characters in length.
	//
	// example:
	//
	// SampleChatbot
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the chatbot. Default value: scenario_im.
	//
	// example:
	//
	// scenario_im
	RobotType *string `json:"RobotType,omitempty" xml:"RobotType,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CreateInstanceRequest) GetIntroduction() *string {
	return s.Introduction
}

func (s *CreateInstanceRequest) GetLanguageCode() *string {
	return s.LanguageCode
}

func (s *CreateInstanceRequest) GetName() *string {
	return s.Name
}

func (s *CreateInstanceRequest) GetRobotType() *string {
	return s.RobotType
}

func (s *CreateInstanceRequest) SetAgentKey(v string) *CreateInstanceRequest {
	s.AgentKey = &v
	return s
}

func (s *CreateInstanceRequest) SetIntroduction(v string) *CreateInstanceRequest {
	s.Introduction = &v
	return s
}

func (s *CreateInstanceRequest) SetLanguageCode(v string) *CreateInstanceRequest {
	s.LanguageCode = &v
	return s
}

func (s *CreateInstanceRequest) SetName(v string) *CreateInstanceRequest {
	s.Name = &v
	return s
}

func (s *CreateInstanceRequest) SetRobotType(v string) *CreateInstanceRequest {
	s.RobotType = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
