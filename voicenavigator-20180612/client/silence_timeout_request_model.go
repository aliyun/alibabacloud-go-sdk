// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSilenceTimeoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConversationId(v string) *SilenceTimeoutRequest
	GetConversationId() *string
	SetInitialContext(v string) *SilenceTimeoutRequest
	GetInitialContext() *string
	SetInstanceId(v string) *SilenceTimeoutRequest
	GetInstanceId() *string
	SetInstanceOwnerId(v int64) *SilenceTimeoutRequest
	GetInstanceOwnerId() *int64
}

type SilenceTimeoutRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 0099b75d-60fd-4c63-8541-7fbba0ae6bb0
	ConversationId *string `json:"ConversationId,omitempty" xml:"ConversationId,omitempty"`
	// example:
	//
	// {}
	InitialContext *string `json:"InitialContext,omitempty" xml:"InitialContext,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0099b75d-60fd-4c63-8541-7fbba0ae6bb0
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1231639035307976
	InstanceOwnerId *int64 `json:"InstanceOwnerId,omitempty" xml:"InstanceOwnerId,omitempty"`
}

func (s SilenceTimeoutRequest) String() string {
	return dara.Prettify(s)
}

func (s SilenceTimeoutRequest) GoString() string {
	return s.String()
}

func (s *SilenceTimeoutRequest) GetConversationId() *string {
	return s.ConversationId
}

func (s *SilenceTimeoutRequest) GetInitialContext() *string {
	return s.InitialContext
}

func (s *SilenceTimeoutRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SilenceTimeoutRequest) GetInstanceOwnerId() *int64 {
	return s.InstanceOwnerId
}

func (s *SilenceTimeoutRequest) SetConversationId(v string) *SilenceTimeoutRequest {
	s.ConversationId = &v
	return s
}

func (s *SilenceTimeoutRequest) SetInitialContext(v string) *SilenceTimeoutRequest {
	s.InitialContext = &v
	return s
}

func (s *SilenceTimeoutRequest) SetInstanceId(v string) *SilenceTimeoutRequest {
	s.InstanceId = &v
	return s
}

func (s *SilenceTimeoutRequest) SetInstanceOwnerId(v int64) *SilenceTimeoutRequest {
	s.InstanceOwnerId = &v
	return s
}

func (s *SilenceTimeoutRequest) Validate() error {
	return dara.Validate(s)
}
