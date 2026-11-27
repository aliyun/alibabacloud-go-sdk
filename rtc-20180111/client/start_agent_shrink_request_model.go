// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *StartAgentShrinkRequest
	GetAppId() *string
	SetChannelId(v string) *StartAgentShrinkRequest
	GetChannelId() *string
	SetRtcConfigShrink(v string) *StartAgentShrinkRequest
	GetRtcConfigShrink() *string
	SetTaskId(v string) *StartAgentShrinkRequest
	GetTaskId() *string
	SetTemplateId(v string) *StartAgentShrinkRequest
	GetTemplateId() *string
	SetVoiceChatConfigShrink(v string) *StartAgentShrinkRequest
	GetVoiceChatConfigShrink() *string
}

type StartAgentShrinkRequest struct {
	// The application ID. You can create and query the ID in the console. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// aoe****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The ID of the channel to join. Only a single ID is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// The configuration for the agent to join a Real-Time Communication (RTC) channel.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	RtcConfigShrink *string `json:"RtcConfig,omitempty" xml:"RtcConfig,omitempty"`
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 76dasgb****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// The configuration for AI voice chat.
	VoiceChatConfigShrink *string `json:"VoiceChatConfig,omitempty" xml:"VoiceChatConfig,omitempty"`
}

func (s StartAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartAgentShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *StartAgentShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *StartAgentShrinkRequest) GetRtcConfigShrink() *string {
	return s.RtcConfigShrink
}

func (s *StartAgentShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *StartAgentShrinkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *StartAgentShrinkRequest) GetVoiceChatConfigShrink() *string {
	return s.VoiceChatConfigShrink
}

func (s *StartAgentShrinkRequest) SetAppId(v string) *StartAgentShrinkRequest {
	s.AppId = &v
	return s
}

func (s *StartAgentShrinkRequest) SetChannelId(v string) *StartAgentShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *StartAgentShrinkRequest) SetRtcConfigShrink(v string) *StartAgentShrinkRequest {
	s.RtcConfigShrink = &v
	return s
}

func (s *StartAgentShrinkRequest) SetTaskId(v string) *StartAgentShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *StartAgentShrinkRequest) SetTemplateId(v string) *StartAgentShrinkRequest {
	s.TemplateId = &v
	return s
}

func (s *StartAgentShrinkRequest) SetVoiceChatConfigShrink(v string) *StartAgentShrinkRequest {
	s.VoiceChatConfigShrink = &v
	return s
}

func (s *StartAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
