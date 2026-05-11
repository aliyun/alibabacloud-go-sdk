// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollectedNumberRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalContext(v string) *CollectedNumberRequest
	GetAdditionalContext() *string
	SetConversationId(v string) *CollectedNumberRequest
	GetConversationId() *string
	SetInstanceId(v string) *CollectedNumberRequest
	GetInstanceId() *string
	SetInstanceOwnerId(v int64) *CollectedNumberRequest
	GetInstanceOwnerId() *int64
	SetNumber(v string) *CollectedNumberRequest
	GetNumber() *string
}

type CollectedNumberRequest struct {
	AdditionalContext *string `json:"AdditionalContext,omitempty" xml:"AdditionalContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0099b75d-60fd-4c63-8541-7fbba0ae6bb0
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0099b75d-60fd-4c63-8541-7fbba0ae6bb0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1426738157626835
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
	// example:
	//
	// 1500060224
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
}

func (s CollectedNumberRequest) String() string {
	return dara.Prettify(s)
}

func (s CollectedNumberRequest) GoString() string {
	return s.String()
}

func (s *CollectedNumberRequest) GetAdditionalContext() *string {
	return s.AdditionalContext
}

func (s *CollectedNumberRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *CollectedNumberRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CollectedNumberRequest) GetInstanceOwnerId() *int64 {
	return s.InstanceOwnerId
}

func (s *CollectedNumberRequest) GetNumber() *string {
	return s.Number
}

func (s *CollectedNumberRequest) SetAdditionalContext(v string) *CollectedNumberRequest {
	s.AdditionalContext = &v
	return s
}

func (s *CollectedNumberRequest) SetConversationId(v string) *CollectedNumberRequest {
	s.ConversationId = &v
	return s
}

func (s *CollectedNumberRequest) SetInstanceId(v string) *CollectedNumberRequest {
	s.InstanceId = &v
	return s
}

func (s *CollectedNumberRequest) SetInstanceOwnerId(v int64) *CollectedNumberRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *CollectedNumberRequest) SetNumber(v string) *CollectedNumberRequest {
	s.Number = &v
	return s
}

func (s *CollectedNumberRequest) Validate() error {
	return dara.Validate(s)
}
