// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKopilotChatStreamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *KopilotChatStreamRequest
	GetAcceptLanguage() *string
	SetMessage(v string) *KopilotChatStreamRequest
	GetMessage() *string
	SetRegionId(v string) *KopilotChatStreamRequest
	GetRegionId() *string
	SetSessionId(v string) *KopilotChatStreamRequest
	GetSessionId() *string
}

type KopilotChatStreamRequest struct {
	// The language for system operation prompts. Valid values: zh-CN, en-US, and ja-JP. If not specified, the compatible language parameter is read. If neither is specified, the request language is used. This parameter does not guarantee that the language of freely generated model content will change.
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The chat message content entered by the user.
	//
	// example:
	//
	// hello
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// efff104e-c0b1-4005-8c73-*********
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s KopilotChatStreamRequest) String() string {
	return dara.Prettify(s)
}

func (s KopilotChatStreamRequest) GoString() string {
	return s.String()
}

func (s *KopilotChatStreamRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *KopilotChatStreamRequest) GetMessage() *string {
	return s.Message
}

func (s *KopilotChatStreamRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *KopilotChatStreamRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *KopilotChatStreamRequest) SetAcceptLanguage(v string) *KopilotChatStreamRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *KopilotChatStreamRequest) SetMessage(v string) *KopilotChatStreamRequest {
	s.Message = &v
	return s
}

func (s *KopilotChatStreamRequest) SetRegionId(v string) *KopilotChatStreamRequest {
	s.RegionId = &v
	return s
}

func (s *KopilotChatStreamRequest) SetSessionId(v string) *KopilotChatStreamRequest {
	s.SessionId = &v
	return s
}

func (s *KopilotChatStreamRequest) Validate() error {
	return dara.Validate(s)
}
