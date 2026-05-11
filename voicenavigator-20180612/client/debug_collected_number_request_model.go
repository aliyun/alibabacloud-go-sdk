// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDebugCollectedNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *DebugCollectedNumberRequest
	GetConversationId() *string
	SetInstanceId(v string) *DebugCollectedNumberRequest
	GetInstanceId() *string
	SetNumber(v string) *DebugCollectedNumberRequest
	GetNumber() *string
}

type DebugCollectedNumberRequest struct {
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
	// 7cefbff0-8d50-4d6f-b93c-73cee23c1555
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 123
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s DebugCollectedNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s DebugCollectedNumberRequest) GoString() string {
	return s.String()
}

func (s *DebugCollectedNumberRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *DebugCollectedNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DebugCollectedNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *DebugCollectedNumberRequest) SetConversationId(v string) *DebugCollectedNumberRequest {
	s.ConversationId = &v
	return s
}

func (s *DebugCollectedNumberRequest) SetInstanceId(v string) *DebugCollectedNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *DebugCollectedNumberRequest) SetNumber(v string) *DebugCollectedNumberRequest {
	s.Number = &v
	return s
}

func (s *DebugCollectedNumberRequest) Validate() error {
	return dara.Validate(s)
}
