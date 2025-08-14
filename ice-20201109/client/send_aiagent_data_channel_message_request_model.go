// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendAIAgentDataChannelMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *SendAIAgentDataChannelMessageRequest
	GetInstanceId() *string
	SetMessage(v string) *SendAIAgentDataChannelMessageRequest
	GetMessage() *string
}

type SendAIAgentDataChannelMessageRequest struct {
	// The ID of the AI agent in the conversation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 39f8e0bc005e4f309379701645f4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The DataChannel message you want to send. You must specify a JSON string. The value can be up to 8,192 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// {"key":"value"}
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s SendAIAgentDataChannelMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s SendAIAgentDataChannelMessageRequest) GoString() string {
	return s.String()
}

func (s *SendAIAgentDataChannelMessageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SendAIAgentDataChannelMessageRequest) GetMessage() *string {
	return s.Message
}

func (s *SendAIAgentDataChannelMessageRequest) SetInstanceId(v string) *SendAIAgentDataChannelMessageRequest {
	s.InstanceId = &v
	return s
}

func (s *SendAIAgentDataChannelMessageRequest) SetMessage(v string) *SendAIAgentDataChannelMessageRequest {
	s.Message = &v
	return s
}

func (s *SendAIAgentDataChannelMessageRequest) Validate() error {
	return dara.Validate(s)
}
