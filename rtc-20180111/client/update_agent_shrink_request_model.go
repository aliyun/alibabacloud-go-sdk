// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *UpdateAgentShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *UpdateAgentShrinkRequest
	GetChannelId() *string
	SetTaskId(v string) *UpdateAgentShrinkRequest
	GetTaskId() *string
	SetVoiceChatConfigShrink(v string) *UpdateAgentShrinkRequest
	GetVoiceChatConfigShrink() *string
}

type UpdateAgentShrinkRequest struct {
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
	TaskId                *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	VoiceChatConfigShrink *string `json:"VoiceChatConfig,omitempty" xml:"VoiceChatConfig,omitempty"`
}

func (s UpdateAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateAgentShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *UpdateAgentShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *UpdateAgentShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateAgentShrinkRequest) GetVoiceChatConfigShrink() *string {
	return s.VoiceChatConfigShrink
}

func (s *UpdateAgentShrinkRequest) SetAppId(v string) *UpdateAgentShrinkRequest {
	s.AppId = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetChannelId(v string) *UpdateAgentShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetTaskId(v string) *UpdateAgentShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *UpdateAgentShrinkRequest) SetVoiceChatConfigShrink(v string) *UpdateAgentShrinkRequest {
	s.VoiceChatConfigShrink = &v
	return s
}

func (s *UpdateAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
