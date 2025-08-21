// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *UpdateInstanceRequest
	GetAgentKey() *string
	SetInstanceId(v string) *UpdateInstanceRequest
	GetInstanceId() *string
	SetIntroduction(v string) *UpdateInstanceRequest
	GetIntroduction() *string
	SetName(v string) *UpdateInstanceRequest
	GetName() *string
}

type UpdateInstanceRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 用于C端问答的机器人
	Introduction *string `json:"Introduction,omitempty" xml:"Introduction,omitempty"`
	// example:
	//
	// 智能客服-小C
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UpdateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateInstanceRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *UpdateInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateInstanceRequest) GetIntroduction() *string {
	return s.Introduction
}

func (s *UpdateInstanceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateInstanceRequest) SetAgentKey(v string) *UpdateInstanceRequest {
	s.AgentKey = &v
	return s
}

func (s *UpdateInstanceRequest) SetInstanceId(v string) *UpdateInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateInstanceRequest) SetIntroduction(v string) *UpdateInstanceRequest {
	s.Introduction = &v
	return s
}

func (s *UpdateInstanceRequest) SetName(v string) *UpdateInstanceRequest {
	s.Name = &v
	return s
}

func (s *UpdateInstanceRequest) Validate() error {
	return dara.Validate(s)
}
