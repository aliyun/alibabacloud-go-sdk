// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBeginDialogueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalledNumber(v string) *BeginDialogueRequest
	GetCalledNumber() *string
	SetCallingNumber(v string) *BeginDialogueRequest
	GetCallingNumber() *string
	SetConversationId(v string) *BeginDialogueRequest
	GetConversationId() *string
	SetInitialContext(v string) *BeginDialogueRequest
	GetInitialContext() *string
	SetInstanceId(v string) *BeginDialogueRequest
	GetInstanceId() *string
	SetInstanceOwnerId(v int64) *BeginDialogueRequest
	GetInstanceOwnerId() *int64
}

type BeginDialogueRequest struct {
	// example:
	//
	// 10086
	CalledNumber *string `json:"CalledNumber,omitempty" xml:"CalledNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1358158****
	CallingNumber *string `json:"CallingNumber,omitempty" xml:"CallingNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c28fc549-d88f-4f6e-85ad-a0806e5e39c0
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// {\\"channelId\\":\\"fe2559d3-5fc9-4fa5-8314-32b9f762791d\\"}
	InitialContext *string `json:"InitialContext,omitempty" xml:"InitialContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 4d7db6670b8e41b5adb1f21560ea9272
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1231639035307976
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
}

func (s BeginDialogueRequest) String() string {
	return dara.Prettify(s)
}

func (s BeginDialogueRequest) GoString() string {
	return s.String()
}

func (s *BeginDialogueRequest) GetCalledNumber() *string {
	return s.CalledNumber
}

func (s *BeginDialogueRequest) GetCallingNumber() *string {
	return s.CallingNumber
}

func (s *BeginDialogueRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *BeginDialogueRequest) GetInitialContext() *string {
	return s.InitialContext
}

func (s *BeginDialogueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BeginDialogueRequest) GetInstanceOwnerId() *int64 {
	return s.InstanceOwnerId
}

func (s *BeginDialogueRequest) SetCalledNumber(v string) *BeginDialogueRequest {
	s.CalledNumber = &v
	return s
}

func (s *BeginDialogueRequest) SetCallingNumber(v string) *BeginDialogueRequest {
	s.CallingNumber = &v
	return s
}

func (s *BeginDialogueRequest) SetConversationId(v string) *BeginDialogueRequest {
	s.ConversationId = &v
	return s
}

func (s *BeginDialogueRequest) SetInitialContext(v string) *BeginDialogueRequest {
	s.InitialContext = &v
	return s
}

func (s *BeginDialogueRequest) SetInstanceId(v string) *BeginDialogueRequest {
	s.InstanceId = &v
	return s
}

func (s *BeginDialogueRequest) SetInstanceOwnerId(v int64) *BeginDialogueRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *BeginDialogueRequest) Validate() error {
	return dara.Validate(s)
}
