// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAIAgentSpeechRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnableInterrupt(v bool) *SendAIAgentSpeechRequest
	GetEnableInterrupt() *bool
	SetInstanceId(v string) *SendAIAgentSpeechRequest
	GetInstanceId() *string
	SetText(v string) *SendAIAgentSpeechRequest
	GetText() *string
	SetType(v string) *SendAIAgentSpeechRequest
	GetType() *string
}

type SendAIAgentSpeechRequest struct {
	// Specifies whether the broadcast can interrupt the ongoing speech. Default value: true
	//
	// example:
	//
	// true
	EnableInterrupt *bool `json:"EnableInterrupt,omitempty" xml:"EnableInterrupt,omitempty"`
	// Agent instance ID.
	//
	// > The InstanceId is the unique ID returned after successfully starting an agent instance. For details about starting an agent instance, see [StartAIAgentInstance](https://help.aliyun.com/document_detail/2846201.html) and [GenerateAIAgentCall](https://help.aliyun.com/document_detail/2846209.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The text content to be played back. The supported input format varies based on the Type parameter. The length cannot exceed 1024 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hello, welcome to our service.
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
	// Input type. Valid values:
	//
	// - Text: Input is plain text.
	//
	// - AudioUrl: Input is an audio URL.
	//
	// Default value: Text.
	//
	// example:
	//
	// Text
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s SendAIAgentSpeechRequest) String() string {
	return dara.Prettify(s)
}

func (s SendAIAgentSpeechRequest) GoString() string {
	return s.String()
}

func (s *SendAIAgentSpeechRequest) GetEnableInterrupt() *bool {
	return s.EnableInterrupt
}

func (s *SendAIAgentSpeechRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendAIAgentSpeechRequest) GetText() *string {
	return s.Text
}

func (s *SendAIAgentSpeechRequest) GetType() *string {
	return s.Type
}

func (s *SendAIAgentSpeechRequest) SetEnableInterrupt(v bool) *SendAIAgentSpeechRequest {
	s.EnableInterrupt = &v
	return s
}

func (s *SendAIAgentSpeechRequest) SetInstanceId(v string) *SendAIAgentSpeechRequest {
	s.InstanceId = &v
	return s
}

func (s *SendAIAgentSpeechRequest) SetText(v string) *SendAIAgentSpeechRequest {
	s.Text = &v
	return s
}

func (s *SendAIAgentSpeechRequest) SetType(v string) *SendAIAgentSpeechRequest {
	s.Type = &v
	return s
}

func (s *SendAIAgentSpeechRequest) Validate() error {
	return dara.Validate(s)
}
