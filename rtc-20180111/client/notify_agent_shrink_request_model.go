// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyAgentShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *NotifyAgentShrinkRequest
	GetAppId() *string
	SetBackgroundMusicShrink(v string) *NotifyAgentShrinkRequest
	GetBackgroundMusicShrink() *string
	SetChannelId(v string) *NotifyAgentShrinkRequest
	GetChannelId() *string
	SetCustomAttribute(v string) *NotifyAgentShrinkRequest
	GetCustomAttribute() *string
	SetInterruptable(v bool) *NotifyAgentShrinkRequest
	GetInterruptable() *bool
	SetMessage(v string) *NotifyAgentShrinkRequest
	GetMessage() *string
	SetPriority(v int32) *NotifyAgentShrinkRequest
	GetPriority() *int32
	SetTaskId(v string) *NotifyAgentShrinkRequest
	GetTaskId() *string
}

type NotifyAgentShrinkRequest struct {
	// The application ID. Create and query it in the console. Only a single ID is supported.
	//
	// example:
	//
	// aec****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Background audio.
	BackgroundMusicShrink *string `json:"BackgroundMusic,omitempty" xml:"BackgroundMusic,omitempty"`
	// The channel to join. Only a single ID is supported.
	//
	// example:
	//
	// yourChannelId
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
	// Custom parameters passed through to the client.
	//
	// example:
	//
	// {\"color\":\"blue\"}
	CustomAttribute *string `json:"CustomAttribute,omitempty" xml:"CustomAttribute,omitempty"`
	// Indicates whether the announcement can be interrupted. If false, the message cannot be interrupted by voice input, key presses, or other methods.
	//
	// example:
	//
	// true
	Interruptable *bool `json:"Interruptable,omitempty" xml:"Interruptable,omitempty"`
	// The text content to announce.
	//
	// example:
	//
	// 您好，现在是2025年4月16日，请问需要什么帮助
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The priority level for inserting the announcement:
	//
	// 1 — Interrupt the current interaction and announce immediately.
	//
	// 2 — Announce after the current interaction ends.
	//
	// 3 — Announce only if no interaction is active.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The task ID.
	//
	// example:
	//
	// yourTaskId
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s NotifyAgentShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s NotifyAgentShrinkRequest) GoString() string {
	return s.String()
}

func (s *NotifyAgentShrinkRequest) GetAppId() *string {
	return s.AppId
}

func (s *NotifyAgentShrinkRequest) GetBackgroundMusicShrink() *string {
	return s.BackgroundMusicShrink
}

func (s *NotifyAgentShrinkRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *NotifyAgentShrinkRequest) GetCustomAttribute() *string {
	return s.CustomAttribute
}

func (s *NotifyAgentShrinkRequest) GetInterruptable() *bool {
	return s.Interruptable
}

func (s *NotifyAgentShrinkRequest) GetMessage() *string {
	return s.Message
}

func (s *NotifyAgentShrinkRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *NotifyAgentShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *NotifyAgentShrinkRequest) SetAppId(v string) *NotifyAgentShrinkRequest {
	s.AppId = &v
	return s
}

func (s *NotifyAgentShrinkRequest) SetBackgroundMusicShrink(v string) *NotifyAgentShrinkRequest {
	s.BackgroundMusicShrink = &v
	return s
}

func (s *NotifyAgentShrinkRequest) SetChannelId(v string) *NotifyAgentShrinkRequest {
	s.ChannelId = &v
	return s
}

func (s *NotifyAgentShrinkRequest) SetCustomAttribute(v string) *NotifyAgentShrinkRequest {
	s.CustomAttribute = &v
	return s
}

func (s *NotifyAgentShrinkRequest) SetInterruptable(v bool) *NotifyAgentShrinkRequest {
	s.Interruptable = &v
	return s
}

func (s *NotifyAgentShrinkRequest) SetMessage(v string) *NotifyAgentShrinkRequest {
	s.Message = &v
	return s
}

func (s *NotifyAgentShrinkRequest) SetPriority(v int32) *NotifyAgentShrinkRequest {
	s.Priority = &v
	return s
}

func (s *NotifyAgentShrinkRequest) SetTaskId(v string) *NotifyAgentShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *NotifyAgentShrinkRequest) Validate() error {
	return dara.Validate(s)
}
