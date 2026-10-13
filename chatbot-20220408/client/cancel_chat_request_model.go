// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelChatRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentKey(v string) *CancelChatRequest
	GetAgentKey() *string
	SetAnswer(v string) *CancelChatRequest
	GetAnswer() *string
	SetChatId(v string) *CancelChatRequest
	GetChatId() *string
	SetInstanceId(v string) *CancelChatRequest
	GetInstanceId() *string
	SetSessionId(v string) *CancelChatRequest
	GetSessionId() *string
	SetType(v string) *CancelChatRequest
	GetType() *string
}

type CancelChatRequest struct {
	// The business space key. If this parameter is not specified, the default business space is accessed. You can obtain the key value on the business management page of the Alibaba Cloud account.
	//
	// example:
	//
	// ac627989eb4f8a98ed05fd098bbae5_p_beebot_public
	AgentKey *string `json:"AgentKey,omitempty" xml:"AgentKey,omitempty"`
	// The answer content.
	//
	// example:
	//
	// <This answer was inter>
	Answer *string `json:"Answer,omitempty" xml:"Answer,omitempty"`
	// The ID that identifies a single conversation.
	//
	// example:
	//
	// 7105a351-b2e7-4c9e-8437-c43a043c0a4e
	ChatId *string `json:"ChatId,omitempty" xml:"ChatId,omitempty"`
	// The chatbot ID.
	//
	// example:
	//
	// chatbot-cn-yjzbyrEvqd
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The session ID, which is used to identify a visitor\\"s session and maintain context information. For a new visitor, you do not need to specify this parameter when calling the Chat operation for the first time. The chatbot starts a session and returns the SessionId in the response of the Chat operation. For subsequent conversation turns of the visitor, pass the SessionId of the current session when calling the Chat operation so that the chatbot can continue the conversation based on the SessionId. Maximum length: 64 characters.
	//
	// example:
	//
	// c1187530338311ebade7cf3eaeb3668a
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// The interruption type.
	//
	// example:
	//
	// Canceled
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CancelChatRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelChatRequest) GoString() string {
	return s.String()
}

func (s *CancelChatRequest) GetAgentKey() *string {
	return s.AgentKey
}

func (s *CancelChatRequest) GetAnswer() *string {
	return s.Answer
}

func (s *CancelChatRequest) GetChatId() *string {
	return s.ChatId
}

func (s *CancelChatRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CancelChatRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CancelChatRequest) GetType() *string {
	return s.Type
}

func (s *CancelChatRequest) SetAgentKey(v string) *CancelChatRequest {
	s.AgentKey = &v
	return s
}

func (s *CancelChatRequest) SetAnswer(v string) *CancelChatRequest {
	s.Answer = &v
	return s
}

func (s *CancelChatRequest) SetChatId(v string) *CancelChatRequest {
	s.ChatId = &v
	return s
}

func (s *CancelChatRequest) SetInstanceId(v string) *CancelChatRequest {
	s.InstanceId = &v
	return s
}

func (s *CancelChatRequest) SetSessionId(v string) *CancelChatRequest {
	s.SessionId = &v
	return s
}

func (s *CancelChatRequest) SetType(v string) *CancelChatRequest {
	s.Type = &v
	return s
}

func (s *CancelChatRequest) Validate() error {
	return dara.Validate(s)
}
