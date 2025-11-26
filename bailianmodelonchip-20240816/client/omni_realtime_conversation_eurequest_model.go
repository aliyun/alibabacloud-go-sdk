// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOmniRealtimeConversationEURequest interface {
	dara.Model
	String() string
	GoString() string
	SetInputAudio(v string) *OmniRealtimeConversationEURequest
	GetInputAudio() *string
	SetUserPrompt(v string) *OmniRealtimeConversationEURequest
	GetUserPrompt() *string
	SetVoice(v string) *OmniRealtimeConversationEURequest
	GetVoice() *string
}

type OmniRealtimeConversationEURequest struct {
	// This parameter is required.
	//
	// example:
	//
	// https://help-static-aliyun-doc.aliyuncs.com/file-manage-files/zh-CN/20250211/tixcef/cherry.wav
	InputAudio *string `json:"inputAudio,omitempty" xml:"inputAudio,omitempty"`
	UserPrompt *string `json:"userPrompt,omitempty" xml:"userPrompt,omitempty"`
	// example:
	//
	// Cherry
	Voice *string `json:"voice,omitempty" xml:"voice,omitempty"`
}

func (s OmniRealtimeConversationEURequest) String() string {
	return dara.Prettify(s)
}

func (s OmniRealtimeConversationEURequest) GoString() string {
	return s.String()
}

func (s *OmniRealtimeConversationEURequest) GetInputAudio() *string {
	return s.InputAudio
}

func (s *OmniRealtimeConversationEURequest) GetUserPrompt() *string {
	return s.UserPrompt
}

func (s *OmniRealtimeConversationEURequest) GetVoice() *string {
	return s.Voice
}

func (s *OmniRealtimeConversationEURequest) SetInputAudio(v string) *OmniRealtimeConversationEURequest {
	s.InputAudio = &v
	return s
}

func (s *OmniRealtimeConversationEURequest) SetUserPrompt(v string) *OmniRealtimeConversationEURequest {
	s.UserPrompt = &v
	return s
}

func (s *OmniRealtimeConversationEURequest) SetVoice(v string) *OmniRealtimeConversationEURequest {
	s.Voice = &v
	return s
}

func (s *OmniRealtimeConversationEURequest) Validate() error {
	return dara.Validate(s)
}
