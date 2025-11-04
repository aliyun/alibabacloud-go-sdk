// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIAgentInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v []*ListAIAgentInstanceResponseBodyInstances) *ListAIAgentInstanceResponseBody
	GetInstances() []*ListAIAgentInstanceResponseBodyInstances
	SetRequestId(v string) *ListAIAgentInstanceResponseBody
	GetRequestId() *string
}

type ListAIAgentInstanceResponseBody struct {
	// The list of the AI agents.
	Instances []*ListAIAgentInstanceResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListAIAgentInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListAIAgentInstanceResponseBody) GetInstances() []*ListAIAgentInstanceResponseBodyInstances {
	return s.Instances
}

func (s *ListAIAgentInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAIAgentInstanceResponseBody) SetInstances(v []*ListAIAgentInstanceResponseBodyInstances) *ListAIAgentInstanceResponseBody {
	s.Instances = v
	return s
}

func (s *ListAIAgentInstanceResponseBody) SetRequestId(v string) *ListAIAgentInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAIAgentInstanceResponseBody) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAIAgentInstanceResponseBodyInstances struct {
	AgentConfig *AIAgentConfig `json:"AgentConfig,omitempty" xml:"AgentConfig,omitempty"`
	// The URL of the call log file for the AI agent. The structure of the file is CallLog in the JSON format.
	//
	// example:
	//
	// https://example.com/call_logs/12345.json
	CallLogUrl *string `json:"CallLogUrl,omitempty" xml:"CallLogUrl,omitempty"`
	// The runtime configurations of the AI agent.
	//
	// example:
	//
	// {"VoiceChat":{"AgentUserId":"voice_agent_001","ChannelId":"voice_channel_001","AuthToken":"your_voice_chat_auth_token"}}
	RuntimeConfig *AIAgentRuntimeConfig `json:"RuntimeConfig,omitempty" xml:"RuntimeConfig,omitempty"`
	// The state of the instance. Valid values:
	//
	// 	- Executing
	//
	// 	- Finished
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Deprecated
	//
	// The template configurations of the AI agent.
	//
	// example:
	//
	// {"VoiceChat": {"VoiceId": "zhixiaoxia"}}
	TemplateConfig *AIAgentTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The custom information.
	//
	// example:
	//
	// {"Email":"johndoe@example.com","Preferences":{"Language":"en"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ListAIAgentInstanceResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListAIAgentInstanceResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListAIAgentInstanceResponseBodyInstances) GetAgentConfig() *AIAgentConfig {
	return s.AgentConfig
}

func (s *ListAIAgentInstanceResponseBodyInstances) GetCallLogUrl() *string {
	return s.CallLogUrl
}

func (s *ListAIAgentInstanceResponseBodyInstances) GetRuntimeConfig() *AIAgentRuntimeConfig {
	return s.RuntimeConfig
}

func (s *ListAIAgentInstanceResponseBodyInstances) GetStatus() *string {
	return s.Status
}

func (s *ListAIAgentInstanceResponseBodyInstances) GetTemplateConfig() *AIAgentTemplateConfig {
	return s.TemplateConfig
}

func (s *ListAIAgentInstanceResponseBodyInstances) GetUserData() *string {
	return s.UserData
}

func (s *ListAIAgentInstanceResponseBodyInstances) SetAgentConfig(v *AIAgentConfig) *ListAIAgentInstanceResponseBodyInstances {
	s.AgentConfig = v
	return s
}

func (s *ListAIAgentInstanceResponseBodyInstances) SetCallLogUrl(v string) *ListAIAgentInstanceResponseBodyInstances {
	s.CallLogUrl = &v
	return s
}

func (s *ListAIAgentInstanceResponseBodyInstances) SetRuntimeConfig(v *AIAgentRuntimeConfig) *ListAIAgentInstanceResponseBodyInstances {
	s.RuntimeConfig = v
	return s
}

func (s *ListAIAgentInstanceResponseBodyInstances) SetStatus(v string) *ListAIAgentInstanceResponseBodyInstances {
	s.Status = &v
	return s
}

func (s *ListAIAgentInstanceResponseBodyInstances) SetTemplateConfig(v *AIAgentTemplateConfig) *ListAIAgentInstanceResponseBodyInstances {
	s.TemplateConfig = v
	return s
}

func (s *ListAIAgentInstanceResponseBodyInstances) SetUserData(v string) *ListAIAgentInstanceResponseBodyInstances {
	s.UserData = &v
	return s
}

func (s *ListAIAgentInstanceResponseBodyInstances) Validate() error {
	if s.AgentConfig != nil {
		if err := s.AgentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.RuntimeConfig != nil {
		if err := s.RuntimeConfig.Validate(); err != nil {
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
