// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAIAgentInstanceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentConfigShrink(v string) *UpdateAIAgentInstanceShrinkRequest
	GetAgentConfigShrink() *string
	SetInstanceId(v string) *UpdateAIAgentInstanceShrinkRequest
	GetInstanceId() *string
	SetTemplateConfigShrink(v string) *UpdateAIAgentInstanceShrinkRequest
	GetTemplateConfigShrink() *string
	SetUserData(v string) *UpdateAIAgentInstanceShrinkRequest
	GetUserData() *string
}

type UpdateAIAgentInstanceShrinkRequest struct {
	AgentConfigShrink *string `json:"AgentConfig,omitempty" xml:"AgentConfig,omitempty"`
	// The ID of the AI agent that you want to update.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Deprecated
	//
	// The template configurations of the AI agent. The configurations are merged with the template configurations that are used to start the AI agent. For more information, see the definition of TemplateConfig.
	TemplateConfigShrink *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// example:
	//
	// {"VoiceId":"xiaoxia"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UpdateAIAgentInstanceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAIAgentInstanceShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAIAgentInstanceShrinkRequest) GetAgentConfigShrink() *string {
	return s.AgentConfigShrink
}

func (s *UpdateAIAgentInstanceShrinkRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAIAgentInstanceShrinkRequest) GetTemplateConfigShrink() *string {
	return s.TemplateConfigShrink
}

func (s *UpdateAIAgentInstanceShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *UpdateAIAgentInstanceShrinkRequest) SetAgentConfigShrink(v string) *UpdateAIAgentInstanceShrinkRequest {
	s.AgentConfigShrink = &v
	return s
}

func (s *UpdateAIAgentInstanceShrinkRequest) SetInstanceId(v string) *UpdateAIAgentInstanceShrinkRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAIAgentInstanceShrinkRequest) SetTemplateConfigShrink(v string) *UpdateAIAgentInstanceShrinkRequest {
	s.TemplateConfigShrink = &v
	return s
}

func (s *UpdateAIAgentInstanceShrinkRequest) SetUserData(v string) *UpdateAIAgentInstanceShrinkRequest {
	s.UserData = &v
	return s
}

func (s *UpdateAIAgentInstanceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
