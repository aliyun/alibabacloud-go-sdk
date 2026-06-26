// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDialogueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalContext(v string) *DialogueRequest
	GetAdditionalContext() *string
	SetCalledNumber(v string) *DialogueRequest
	GetCalledNumber() *string
	SetCallingNumber(v string) *DialogueRequest
	GetCallingNumber() *string
	SetConversationId(v string) *DialogueRequest
	GetConversationId() *string
	SetEmotion(v string) *DialogueRequest
	GetEmotion() *string
	SetInstanceId(v string) *DialogueRequest
	GetInstanceId() *string
	SetInstanceOwnerId(v int64) *DialogueRequest
	GetInstanceOwnerId() *int64
	SetUtterance(v string) *DialogueRequest
	GetUtterance() *string
}

type DialogueRequest struct {
	// The conversation context.
	//
	// example:
	//
	// {}
	AdditionalContext *string `json:"AdditionalContext,omitempty" xml:"AdditionalContext,omitempty"`
	// The called number.
	//
	// example:
	//
	// 10086
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// The calling number.
	//
	// example:
	//
	// 18851708605
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// The ID of the conversation.
	//
	// This parameter is required.
	//
	// example:
	//
	// da37319b-6c83-4268-9f19-814aed62e401
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	Emotion        *string `json:"Emotion,omitempty" xml:"Emotion,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21e0b2a3-168d-4ba7-9009-afc42666eb54
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the Alibaba Cloud account that owns the instance.
	//
	// example:
	//
	// 1426738157626835
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	// The user\\"s input.
	//
	// This parameter is required.
	//
	// example:
	//
	// 行吧，那我就不打扰您了，再见。
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s DialogueRequest) String() string {
	return dara.Prettify(s)
}

func (s DialogueRequest) GoString() string {
	return s.String()
}

func (s *DialogueRequest) GetAdditionalContext() *string {
	return s.AdditionalContext
}

func (s *DialogueRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *DialogueRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *DialogueRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *DialogueRequest) GetEmotion() *string {
	return s.Emotion
}

func (s *DialogueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DialogueRequest) GetInstanceOwnerId() *int64 {
	return s.InstanceOwnerId
}

func (s *DialogueRequest) GetUtterance() *string {
	return s.Utterance
}

func (s *DialogueRequest) SetAdditionalContext(v string) *DialogueRequest {
	s.AdditionalContext = &v
	return s
}

func (s *DialogueRequest) SetCalledNumber(v string) *DialogueRequest {
	s.CalledNumber = &v
	return s
}

func (s *DialogueRequest) SetCallingNumber(v string) *DialogueRequest {
	s.CallingNumber = &v
	return s
}

func (s *DialogueRequest) SetConversationId(v string) *DialogueRequest {
	s.ConversationId = &v
	return s
}

func (s *DialogueRequest) SetEmotion(v string) *DialogueRequest {
	s.Emotion = &v
	return s
}

func (s *DialogueRequest) SetInstanceId(v string) *DialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *DialogueRequest) SetInstanceOwnerId(v int64) *DialogueRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *DialogueRequest) SetUtterance(v string) *DialogueRequest {
	s.Utterance = &v
	return s
}

func (s *DialogueRequest) Validate() error {
	return dara.Validate(s)
}
