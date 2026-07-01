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
	// The AI agent configuration to update. This configuration is merged with the existing configuration of the instance. For more information, see the AIAgentConfig definition. The following parameters in AIAgentConfig can be updated:
	//
	// - VoiceId
	//
	// - EnableVoiceInterrupt
	//
	// - Greeting
	//
	// - Volume
	//
	// - EnablePushToTalk
	//
	// - UseVoiceprint
	//
	// - BailianAppParams
	AgentConfigShrink *string `json:"AgentConfig,omitempty" xml:"AgentConfig,omitempty"`
	// The ID of the AI agent instance.
	//
	// > This unique ID is returned after the AI agent instance starts successfully. For more information about starting an agent, see [StartAIAgentInstance](https://help.aliyun.com/document_detail/2846201.html) and [GenerateAIAgentCall](https://help.aliyun.com/document_detail/2846209.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Deprecated
	//
	// > The AI agent template configuration. This parameter is deprecated. Use the AgentConfig parameter instead.
	//
	// The AI agent configuration to update. This configuration is merged with the existing configuration of the instance. For more information, see the [AIAgentTemplateConfig](https://help.aliyun.com/document_detail/2846193.html) definition.
	//
	// The following parameters in AIAgentTemplateConfig can be updated:
	//
	// - VoiceId (Voice ID)
	//
	// - EnableVoiceInterrupt (Enable voice interruption)
	//
	// - Greeting (Greeting)
	//
	// - Volume (Volume)
	//
	// - EnablePushToTalk (Enable push-to-talk)
	//
	// - UseVoiceprint (Use voiceprint)
	//
	// - AsrMaxSilence (ASR maximum silence duration)
	TemplateConfigShrink *string `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// Custom user data.
	//
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
