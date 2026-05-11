// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugBeginDialogueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *DebugBeginDialogueRequest
	GetCalledNumber() *string
	SetCallingNumber(v string) *DebugBeginDialogueRequest
	GetCallingNumber() *string
	SetConversationId(v string) *DebugBeginDialogueRequest
	GetConversationId() *string
	SetInitialContext(v string) *DebugBeginDialogueRequest
	GetInitialContext() *string
	SetInstanceId(v string) *DebugBeginDialogueRequest
	GetInstanceId() *string
}

type DebugBeginDialogueRequest struct {
	// example:
	//
	// 10086
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 135815*****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8a503680-815d-473e-a9b0-e010f47a64d2
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// {}
	InitialContext *string `json:"InitialContext,omitempty" xml:"InitialContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 8a503680-815d-473e-a9b0-e010f47a64d2
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DebugBeginDialogueRequest) String() string {
	return dara.Prettify(s)
}

func (s DebugBeginDialogueRequest) GoString() string {
	return s.String()
}

func (s *DebugBeginDialogueRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *DebugBeginDialogueRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *DebugBeginDialogueRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *DebugBeginDialogueRequest) GetInitialContext() *string {
	return s.InitialContext
}

func (s *DebugBeginDialogueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DebugBeginDialogueRequest) SetCalledNumber(v string) *DebugBeginDialogueRequest {
	s.CalledNumber = &v
	return s
}

func (s *DebugBeginDialogueRequest) SetCallingNumber(v string) *DebugBeginDialogueRequest {
	s.CallingNumber = &v
	return s
}

func (s *DebugBeginDialogueRequest) SetConversationId(v string) *DebugBeginDialogueRequest {
	s.ConversationId = &v
	return s
}

func (s *DebugBeginDialogueRequest) SetInitialContext(v string) *DebugBeginDialogueRequest {
	s.InitialContext = &v
	return s
}

func (s *DebugBeginDialogueRequest) SetInstanceId(v string) *DebugBeginDialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *DebugBeginDialogueRequest) Validate() error {
	return dara.Validate(s)
}
