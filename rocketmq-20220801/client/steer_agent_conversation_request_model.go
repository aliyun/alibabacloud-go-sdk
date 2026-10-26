// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSteerAgentConversationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *SteerAgentConversationRequest
	GetCommandId() *string
	SetMessage(v string) *SteerAgentConversationRequest
	GetMessage() *string
}

type SteerAgentConversationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cmd-20260819-0001
	CommandId *string `json:"commandId,omitempty" xml:"commandId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 请优先检查消费组在线状态
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
}

func (s SteerAgentConversationRequest) String() string {
	return dara.Prettify(s)
}

func (s SteerAgentConversationRequest) GoString() string {
	return s.String()
}

func (s *SteerAgentConversationRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *SteerAgentConversationRequest) GetMessage() *string {
	return s.Message
}

func (s *SteerAgentConversationRequest) SetCommandId(v string) *SteerAgentConversationRequest {
	s.CommandId = &v
	return s
}

func (s *SteerAgentConversationRequest) SetMessage(v string) *SteerAgentConversationRequest {
	s.Message = &v
	return s
}

func (s *SteerAgentConversationRequest) Validate() error {
	return dara.Validate(s)
}
