// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugDialogueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalContext(v string) *DebugDialogueRequest
	GetAdditionalContext() *string
	SetConversationId(v string) *DebugDialogueRequest
	GetConversationId() *string
	SetInstanceId(v string) *DebugDialogueRequest
	GetInstanceId() *string
	SetUtterance(v string) *DebugDialogueRequest
	GetUtterance() *string
}

type DebugDialogueRequest struct {
	// example:
	//
	// {}
	AdditionalContext *string `json:"AdditionalContext,omitempty" xml:"AdditionalContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 7cefbff0-8d50-4d6f-b93c-73cee23c1555
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abb4aa26-3a8e-43dd-82f8-0c3898c9c67f
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Utterance *string `json:"Utterance,omitempty" xml:"Utterance,omitempty"`
}

func (s DebugDialogueRequest) String() string {
	return dara.Prettify(s)
}

func (s DebugDialogueRequest) GoString() string {
	return s.String()
}

func (s *DebugDialogueRequest) GetAdditionalContext() *string {
	return s.AdditionalContext
}

func (s *DebugDialogueRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *DebugDialogueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DebugDialogueRequest) GetUtterance() *string {
	return s.Utterance
}

func (s *DebugDialogueRequest) SetAdditionalContext(v string) *DebugDialogueRequest {
	s.AdditionalContext = &v
	return s
}

func (s *DebugDialogueRequest) SetConversationId(v string) *DebugDialogueRequest {
	s.ConversationId = &v
	return s
}

func (s *DebugDialogueRequest) SetInstanceId(v string) *DebugDialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *DebugDialogueRequest) SetUtterance(v string) *DebugDialogueRequest {
	s.Utterance = &v
	return s
}

func (s *DebugDialogueRequest) Validate() error {
	return dara.Validate(s)
}
