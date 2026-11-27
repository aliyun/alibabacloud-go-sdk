// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iNotifyAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *NotifyAgentRequest
	GetAppId() *string
	SetBackgroundMusic(v *NotifyAgentRequestBackgroundMusic) *NotifyAgentRequest
	GetBackgroundMusic() *NotifyAgentRequestBackgroundMusic
	SetChannelId(v string) *NotifyAgentRequest
	GetChannelId() *string
	SetCustomAttribute(v string) *NotifyAgentRequest
	GetCustomAttribute() *string
	SetInterruptable(v bool) *NotifyAgentRequest
	GetInterruptable() *bool
	SetMessage(v string) *NotifyAgentRequest
	GetMessage() *string
	SetPriority(v int32) *NotifyAgentRequest
	GetPriority() *int32
	SetTaskId(v string) *NotifyAgentRequest
	GetTaskId() *string
}

type NotifyAgentRequest struct {
	// The application ID. Create and query it in the console. Only a single ID is supported.
	//
	// example:
	//
	// aec****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Background audio.
	BackgroundMusic *NotifyAgentRequestBackgroundMusic `json:"BackgroundMusic,omitempty" xml:"BackgroundMusic,omitempty" type:"Struct"`
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

func (s NotifyAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s NotifyAgentRequest) GoString() string {
	return s.String()
}

func (s *NotifyAgentRequest) GetAppId() *string {
	return s.AppId
}

func (s *NotifyAgentRequest) GetBackgroundMusic() *NotifyAgentRequestBackgroundMusic {
	return s.BackgroundMusic
}

func (s *NotifyAgentRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *NotifyAgentRequest) GetCustomAttribute() *string {
	return s.CustomAttribute
}

func (s *NotifyAgentRequest) GetInterruptable() *bool {
	return s.Interruptable
}

func (s *NotifyAgentRequest) GetMessage() *string {
	return s.Message
}

func (s *NotifyAgentRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *NotifyAgentRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *NotifyAgentRequest) SetAppId(v string) *NotifyAgentRequest {
	s.AppId = &v
	return s
}

func (s *NotifyAgentRequest) SetBackgroundMusic(v *NotifyAgentRequestBackgroundMusic) *NotifyAgentRequest {
	s.BackgroundMusic = v
	return s
}

func (s *NotifyAgentRequest) SetChannelId(v string) *NotifyAgentRequest {
	s.ChannelId = &v
	return s
}

func (s *NotifyAgentRequest) SetCustomAttribute(v string) *NotifyAgentRequest {
	s.CustomAttribute = &v
	return s
}

func (s *NotifyAgentRequest) SetInterruptable(v bool) *NotifyAgentRequest {
	s.Interruptable = &v
	return s
}

func (s *NotifyAgentRequest) SetMessage(v string) *NotifyAgentRequest {
	s.Message = &v
	return s
}

func (s *NotifyAgentRequest) SetPriority(v int32) *NotifyAgentRequest {
	s.Priority = &v
	return s
}

func (s *NotifyAgentRequest) SetTaskId(v string) *NotifyAgentRequest {
	s.TaskId = &v
	return s
}

func (s *NotifyAgentRequest) Validate() error {
	if s.BackgroundMusic != nil {
		if err := s.BackgroundMusic.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type NotifyAgentRequestBackgroundMusic struct {
	// The audio format.
	//
	// - wav
	//
	// example:
	//
	// wav
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// The resource URL.
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s NotifyAgentRequestBackgroundMusic) String() string {
	return dara.Prettify(s)
}

func (s NotifyAgentRequestBackgroundMusic) GoString() string {
	return s.String()
}

func (s *NotifyAgentRequestBackgroundMusic) GetFormat() *string {
	return s.Format
}

func (s *NotifyAgentRequestBackgroundMusic) GetUrl() *string {
	return s.Url
}

func (s *NotifyAgentRequestBackgroundMusic) SetFormat(v string) *NotifyAgentRequestBackgroundMusic {
	s.Format = &v
	return s
}

func (s *NotifyAgentRequestBackgroundMusic) SetUrl(v string) *NotifyAgentRequestBackgroundMusic {
	s.Url = &v
	return s
}

func (s *NotifyAgentRequestBackgroundMusic) Validate() error {
	return dara.Validate(s)
}
