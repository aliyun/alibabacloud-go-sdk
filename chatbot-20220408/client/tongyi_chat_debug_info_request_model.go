// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTongyiChatDebugInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *TongyiChatDebugInfoRequest
	GetAgentKey() *string
	SetInstanceId(v string) *TongyiChatDebugInfoRequest
	GetInstanceId() *string
	SetMessageId(v string) *TongyiChatDebugInfoRequest
	GetMessageId() *string
}

type TongyiChatDebugInfoRequest struct {
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// chatbot-cn-7QuUfaqMQe
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 396E74B3-D84B-46CE-9E51-91C06AD22E31
	MessageId *string `json:"MessageId,omitempty" xml:"MessageId,omitempty"`
}

func (s TongyiChatDebugInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s TongyiChatDebugInfoRequest) GoString() string {
	return s.String()
}

func (s *TongyiChatDebugInfoRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *TongyiChatDebugInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TongyiChatDebugInfoRequest) GetMessageId() *string {
	return s.MessageId
}

func (s *TongyiChatDebugInfoRequest) SetAgentKey(v string) *TongyiChatDebugInfoRequest {
	s.AgentKey = &v
	return s
}

func (s *TongyiChatDebugInfoRequest) SetInstanceId(v string) *TongyiChatDebugInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *TongyiChatDebugInfoRequest) SetMessageId(v string) *TongyiChatDebugInfoRequest {
	s.MessageId = &v
	return s
}

func (s *TongyiChatDebugInfoRequest) Validate() error {
	return dara.Validate(s)
}
