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
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// 用于手机app的小蜜机器人
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// zh-cn
	LanguageCode *string `json:"LanguageCode,omitempty" xml:"LanguageCode,omitempty"`
	// example:
	//
	// 小蜜机器人
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
