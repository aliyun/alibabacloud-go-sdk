// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAIAgentInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstance(v *DescribeAIAgentInstanceResponseBodyInstance) *DescribeAIAgentInstanceResponseBody
	GetInstance() *DescribeAIAgentInstanceResponseBodyInstance
	SetRequestId(v string) *DescribeAIAgentInstanceResponseBody
	GetRequestId() *string
}

type DescribeAIAgentInstanceResponseBody struct {
	// Information about the AI agent instance.
	Instance *DescribeAIAgentInstanceResponseBodyInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7B117AF5-2A16-412C-B127-FA6175ED1AD0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAIAgentInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIAgentInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAIAgentInstanceResponseBody) GetInstance() *DescribeAIAgentInstanceResponseBodyInstance {
	return s.Instance
}

func (s *DescribeAIAgentInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAIAgentInstanceResponseBody) SetInstance(v *DescribeAIAgentInstanceResponseBodyInstance) *DescribeAIAgentInstanceResponseBody {
	s.Instance = v
	return s
}

func (s *DescribeAIAgentInstanceResponseBody) SetRequestId(v string) *DescribeAIAgentInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAIAgentInstanceResponseBody) Validate() error {
	if s.Instance != nil {
		if err := s.Instance.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAIAgentInstanceResponseBodyInstance struct {
	// The configuration of the AI agent.
	AgentConfig *AIAgentConfig `json:"AgentConfig,omitempty" xml:"AgentConfig,omitempty"`
	// Information about the call.
	CallInfo *AIAgentCallInfo `json:"CallInfo,omitempty" xml:"CallInfo,omitempty"`
	// The URL of the call log.
	//
	// example:
	//
	// https://example.com/call_logs/12345
	CallLogUrl *string `json:"CallLogUrl,omitempty" xml:"CallLogUrl,omitempty"`
	// The creation time (UTC).
	//
	// example:
	//
	// 2025-07-18T06:39:08.000+00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The modification time (UTC).
	//
	// example:
	//
	// 2025-07-18T06:40:12.000+00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The runtime configuration of the AI agent.
	//
	// example:
	//
	// {"VoiceChat":{"AgentUserId":"voice_agent_001","ChannelId":"voice_channel_001","AuthToken":"your_voice_chat_auth_token"}}
	RuntimeConfig *AIAgentRuntimeConfig `json:"RuntimeConfig,omitempty" xml:"RuntimeConfig,omitempty"`
	// The session ID for the conversation. This parameter is empty by default.
	//
	// example:
	//
	// 955535**************
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The status of the AI agent instance, such as `Finished` or `Executing`.
	//
	// example:
	//
	// Finished
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Deprecated
	//
	// The AI agent template configuration.
	//
	// example:
	//
	// {"VoiceChat": {"AppId": "your_voice_chat_app_id"}}
	TemplateConfig *AIAgentTemplateConfig `json:"TemplateConfig,omitempty" xml:"TemplateConfig,omitempty"`
	// The user data.
	//
	// example:
	//
	// {"Email":"johndoe@example.com","Preferences":{"Language":"en"}}
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s DescribeAIAgentInstanceResponseBodyInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeAIAgentInstanceResponseBodyInstance) GoString() string {
	return s.String()
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) GetAgentConfig() *AIAgentConfig {
	return s.AgentConfig
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) GetCallInfo() *AIAgentCallInfo {
	return s.CallInfo
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) GetCallLogUrl() *string {
	return s.CallLogUrl
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) GetRuntimeConfig() *AIAgentRuntimeConfig {
	return s.RuntimeConfig
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) GetTemplateConfig() *AIAgentTemplateConfig {
	return s.TemplateConfig
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) GetUserData() *string {
	return s.UserData
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) SetAgentConfig(v *AIAgentConfig) *DescribeAIAgentInstanceResponseBodyInstance {
	s.AgentConfig = v
	return s
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) SetCallInfo(v *AIAgentCallInfo) *DescribeAIAgentInstanceResponseBodyInstance {
	s.CallInfo = v
	return s
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) SetCallLogUrl(v string) *DescribeAIAgentInstanceResponseBodyInstance {
	s.CallLogUrl = &v
	return s
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) SetGmtCreate(v string) *DescribeAIAgentInstanceResponseBodyInstance {
	s.GmtCreate = &v
	return s
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) SetGmtModified(v string) *DescribeAIAgentInstanceResponseBodyInstance {
	s.GmtModified = &v
	return s
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) SetRuntimeConfig(v *AIAgentRuntimeConfig) *DescribeAIAgentInstanceResponseBodyInstance {
	s.RuntimeConfig = v
	return s
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) SetSessionId(v string) *DescribeAIAgentInstanceResponseBodyInstance {
	s.SessionId = &v
	return s
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) SetStatus(v string) *DescribeAIAgentInstanceResponseBodyInstance {
	s.Status = &v
	return s
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) SetTemplateConfig(v *AIAgentTemplateConfig) *DescribeAIAgentInstanceResponseBodyInstance {
	s.TemplateConfig = v
	return s
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) SetUserData(v string) *DescribeAIAgentInstanceResponseBodyInstance {
	s.UserData = &v
	return s
}

func (s *DescribeAIAgentInstanceResponseBodyInstance) Validate() error {
	if s.AgentConfig != nil {
		if err := s.AgentConfig.Validate(); err != nil {
			return err
		}
	}
	if s.CallInfo != nil {
		if err := s.CallInfo.Validate(); err != nil {
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
