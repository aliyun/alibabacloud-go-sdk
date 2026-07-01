// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAIAgentInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentConfig(v *AIAgentConfig) *UpdateAIAgentInstanceRequest
	GetAgentConfig() *AIAgentConfig
	SetInstanceId(v string) *UpdateAIAgentInstanceRequest
	GetInstanceId() *string
	SetTemplateConfig(v *AIAgentTemplateConfig) *UpdateAIAgentInstanceRequest
	GetTemplateConfig() *AIAgentTemplateConfig
	SetUserData(v string) *UpdateAIAgentInstanceRequest
	GetUserData() *string
}

type UpdateAIAgentInstanceRequest struct {
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
	AgentConfig *AIAgentConfig `json:"AgentConfig,omitempty" xml:"AgentConfig,omitempty"`
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
	TemplateConfig *AIAgentTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// Custom user data.
	//
	// example:
	//
	// {"VoiceId":"xiaoxia"}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UpdateAIAgentInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAIAgentInstanceRequest) GoString() string {
	return s.String()
}

func (s *UpdateAIAgentInstanceRequest) GetAgentConfig() *AIAgentConfig {
	return s.AgentConfig
}

func (s *UpdateAIAgentInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateAIAgentInstanceRequest) GetTemplateConfig() *AIAgentTemplateConfig {
	return s.TemplateConfig
}

func (s *UpdateAIAgentInstanceRequest) GetUserData() *string {
	return s.UserData
}

func (s *UpdateAIAgentInstanceRequest) SetAgentConfig(v *AIAgentConfig) *UpdateAIAgentInstanceRequest {
	s.AgentConfig = v
	return s
}

func (s *UpdateAIAgentInstanceRequest) SetInstanceId(v string) *UpdateAIAgentInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateAIAgentInstanceRequest) SetTemplateConfig(v *AIAgentTemplateConfig) *UpdateAIAgentInstanceRequest {
	s.TemplateConfig = v
	return s
}

func (s *UpdateAIAgentInstanceRequest) SetUserData(v string) *UpdateAIAgentInstanceRequest {
	s.UserData = &v
	return s
}

func (s *UpdateAIAgentInstanceRequest) Validate() error {
	if s.AgentConfig != nil {
		if err := s.AgentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.TemplateConfig != nil {
		if err := s.TemplateConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}
