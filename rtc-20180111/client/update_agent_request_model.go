// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateAgentRequest
	GetAppId() *string
	SetChannelId(v string) *UpdateAgentRequest
	GetChannelId() *string
	SetTaskId(v string) *UpdateAgentRequest
	GetTaskId() *string
	SetVoiceChatConfig(v *UpdateAgentRequestVoiceChatConfig) *UpdateAgentRequest
	GetVoiceChatConfig() *UpdateAgentRequestVoiceChatConfig
}

type UpdateAgentRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4eah****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId          *string                            `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VoiceChatConfig *UpdateAgentRequestVoiceChatConfig `json:"VoiceChatConfig,omitempty" xml:"VoiceChatConfig,omitempty" type:"Struct"`
}

func (s UpdateAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAgentRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateAgentRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateAgentRequest) GetVoiceChatConfig() *UpdateAgentRequestVoiceChatConfig {
	return s.VoiceChatConfig
}

func (s *UpdateAgentRequest) SetAppId(v string) *UpdateAgentRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAgentRequest) SetChannelId(v string) *UpdateAgentRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateAgentRequest) SetTaskId(v string) *UpdateAgentRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateAgentRequest) SetVoiceChatConfig(v *UpdateAgentRequestVoiceChatConfig) *UpdateAgentRequest {
	s.VoiceChatConfig = v
	return s
}

func (s *UpdateAgentRequest) Validate() error {
	if s.VoiceChatConfig != nil {
		if err := s.VoiceChatConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateAgentRequestVoiceChatConfig struct {
	// example:
	//
	// 2
	ChatMode *int32 `json:"ChatMode,omitempty" xml:"ChatMode,omitempty"`
	// example:
	//
	// 2
	InterruptMode *int32 `json:"InterruptMode,omitempty" xml:"InterruptMode,omitempty"`
}

func (s UpdateAgentRequestVoiceChatConfig) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentRequestVoiceChatConfig) GoString() string {
	return s.String()
}

func (s *UpdateAgentRequestVoiceChatConfig) GetChatMode() *int32 {
	return s.ChatMode
}

func (s *UpdateAgentRequestVoiceChatConfig) GetInterruptMode() *int32 {
	return s.InterruptMode
}

func (s *UpdateAgentRequestVoiceChatConfig) SetChatMode(v int32) *UpdateAgentRequestVoiceChatConfig {
	s.ChatMode = &v
	return s
}

func (s *UpdateAgentRequestVoiceChatConfig) SetInterruptMode(v int32) *UpdateAgentRequestVoiceChatConfig {
	s.InterruptMode = &v
	return s
}

func (s *UpdateAgentRequestVoiceChatConfig) Validate() error {
	return dara.Validate(s)
}
