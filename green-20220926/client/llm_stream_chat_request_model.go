// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLlmStreamChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannel(v string) *LlmStreamChatRequest
	GetChannel() *string
	SetMessages(v interface{}) *LlmStreamChatRequest
	GetMessages() interface{}
	SetTemperature(v float32) *LlmStreamChatRequest
	GetTemperature() *float32
	SetTopP(v float32) *LlmStreamChatRequest
	GetTopP() *float32
	SetType(v string) *LlmStreamChatRequest
	GetType() *string
}

type LlmStreamChatRequest struct {
	// example:
	//
	// guardCustomTest
	Channel *string `json:"Channel,omitempty" xml:"Channel,omitempty"`
	// Conversation information
	//
	// example:
	//
	// [{\\"content\\":\\"你好\\",\\"role\\":\\"user\\"}]
	Messages interface{} `json:"Messages,omitempty" xml:"Messages,omitempty"`
	// Temperature value for the large model
	//
	// example:
	//
	// 0.5
	Temperature *float32 `json:"Temperature,omitempty" xml:"Temperature,omitempty"`
	// Top p parameter controlling the randomness of the large model\\"s output.
	//
	// example:
	//
	// 0.5
	TopP *float32 `json:"TopP,omitempty" xml:"TopP,omitempty"`
	// Type of conversation
	//
	// example:
	//
	// image
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s LlmStreamChatRequest) String() string {
	return dara.Prettify(s)
}

func (s LlmStreamChatRequest) GoString() string {
	return s.String()
}

func (s *LlmStreamChatRequest) GetChannel() *string {
	return s.Channel
}

func (s *LlmStreamChatRequest) GetMessages() interface{} {
	return s.Messages
}

func (s *LlmStreamChatRequest) GetTemperature() *float32 {
	return s.Temperature
}

func (s *LlmStreamChatRequest) GetTopP() *float32 {
	return s.TopP
}

func (s *LlmStreamChatRequest) GetType() *string {
	return s.Type
}

func (s *LlmStreamChatRequest) SetChannel(v string) *LlmStreamChatRequest {
	s.Channel = &v
	return s
}

func (s *LlmStreamChatRequest) SetMessages(v interface{}) *LlmStreamChatRequest {
	s.Messages = v
	return s
}

func (s *LlmStreamChatRequest) SetTemperature(v float32) *LlmStreamChatRequest {
	s.Temperature = &v
	return s
}

func (s *LlmStreamChatRequest) SetTopP(v float32) *LlmStreamChatRequest {
	s.TopP = &v
	return s
}

func (s *LlmStreamChatRequest) SetType(v string) *LlmStreamChatRequest {
	s.Type = &v
	return s
}

func (s *LlmStreamChatRequest) Validate() error {
	return dara.Validate(s)
}
