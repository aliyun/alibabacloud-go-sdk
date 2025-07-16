// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMessageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *ListMessageRequest
	GetAssistantId() *string
	SetLimit(v int32) *ListMessageRequest
	GetLimit() *int32
	SetOrder(v string) *ListMessageRequest
	GetOrder() *string
	SetOriginalAssistantId(v string) *ListMessageRequest
	GetOriginalAssistantId() *string
	SetRunId(v string) *ListMessageRequest
	GetRunId() *string
	SetSourceIdOfOriginalAssistantId(v string) *ListMessageRequest
	GetSourceIdOfOriginalAssistantId() *string
	SetSourceTypeOfOriginalAssistantId(v string) *ListMessageRequest
	GetSourceTypeOfOriginalAssistantId() *string
	SetThreadId(v string) *ListMessageRequest
	GetThreadId() *string
}

type ListMessageRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// assistantId1
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	// example:
	//
	// 20
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// assistantId2
	OriginalAssistantId *string `json:"originalAssistantId,omitempty" xml:"originalAssistantId,omitempty"`
	// example:
	//
	// runId123
	RunId *string `json:"runId,omitempty" xml:"runId,omitempty"`
	// example:
	//
	// agentKey1
	SourceIdOfOriginalAssistantId *string `json:"sourceIdOfOriginalAssistantId,omitempty" xml:"sourceIdOfOriginalAssistantId,omitempty"`
	// example:
	//
	// 1
	SourceTypeOfOriginalAssistantId *string `json:"sourceTypeOfOriginalAssistantId,omitempty" xml:"sourceTypeOfOriginalAssistantId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// threadId123
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s ListMessageRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMessageRequest) GoString() string {
	return s.String()
}

func (s *ListMessageRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *ListMessageRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListMessageRequest) GetOrder() *string {
	return s.Order
}

func (s *ListMessageRequest) GetOriginalAssistantId() *string {
	return s.OriginalAssistantId
}

func (s *ListMessageRequest) GetRunId() *string {
	return s.RunId
}

func (s *ListMessageRequest) GetSourceIdOfOriginalAssistantId() *string {
	return s.SourceIdOfOriginalAssistantId
}

func (s *ListMessageRequest) GetSourceTypeOfOriginalAssistantId() *string {
	return s.SourceTypeOfOriginalAssistantId
}

func (s *ListMessageRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *ListMessageRequest) SetAssistantId(v string) *ListMessageRequest {
	s.AssistantId = &v
	return s
}

func (s *ListMessageRequest) SetLimit(v int32) *ListMessageRequest {
	s.Limit = &v
	return s
}

func (s *ListMessageRequest) SetOrder(v string) *ListMessageRequest {
	s.Order = &v
	return s
}

func (s *ListMessageRequest) SetOriginalAssistantId(v string) *ListMessageRequest {
	s.OriginalAssistantId = &v
	return s
}

func (s *ListMessageRequest) SetRunId(v string) *ListMessageRequest {
	s.RunId = &v
	return s
}

func (s *ListMessageRequest) SetSourceIdOfOriginalAssistantId(v string) *ListMessageRequest {
	s.SourceIdOfOriginalAssistantId = &v
	return s
}

func (s *ListMessageRequest) SetSourceTypeOfOriginalAssistantId(v string) *ListMessageRequest {
	s.SourceTypeOfOriginalAssistantId = &v
	return s
}

func (s *ListMessageRequest) SetThreadId(v string) *ListMessageRequest {
	s.ThreadId = &v
	return s
}

func (s *ListMessageRequest) Validate() error {
	return dara.Validate(s)
}
