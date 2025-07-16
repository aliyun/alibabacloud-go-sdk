// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateThreadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssistantId(v string) *CreateThreadRequest
	GetAssistantId() *string
	SetOriginalAssistantId(v string) *CreateThreadRequest
	GetOriginalAssistantId() *string
	SetSourceIdOfOriginalAssistantId(v string) *CreateThreadRequest
	GetSourceIdOfOriginalAssistantId() *string
	SetSourceTypeOfOriginalAssistantId(v int32) *CreateThreadRequest
	GetSourceTypeOfOriginalAssistantId() *int32
}

type CreateThreadRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// assistantId1
	AssistantId *string `json:"assistantId,omitempty" xml:"assistantId,omitempty"`
	// example:
	//
	// assistantId
	OriginalAssistantId *string `json:"originalAssistantId,omitempty" xml:"originalAssistantId,omitempty"`
	// example:
	//
	// agentKey1
	SourceIdOfOriginalAssistantId   *string `json:"sourceIdOfOriginalAssistantId,omitempty" xml:"sourceIdOfOriginalAssistantId,omitempty"`
	SourceTypeOfOriginalAssistantId *int32  `json:"sourceTypeOfOriginalAssistantId,omitempty" xml:"sourceTypeOfOriginalAssistantId,omitempty"`
}

func (s CreateThreadRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateThreadRequest) GoString() string {
	return s.String()
}

func (s *CreateThreadRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *CreateThreadRequest) GetOriginalAssistantId() *string {
	return s.OriginalAssistantId
}

func (s *CreateThreadRequest) GetSourceIdOfOriginalAssistantId() *string {
	return s.SourceIdOfOriginalAssistantId
}

func (s *CreateThreadRequest) GetSourceTypeOfOriginalAssistantId() *int32 {
	return s.SourceTypeOfOriginalAssistantId
}

func (s *CreateThreadRequest) SetAssistantId(v string) *CreateThreadRequest {
	s.AssistantId = &v
	return s
}

func (s *CreateThreadRequest) SetOriginalAssistantId(v string) *CreateThreadRequest {
	s.OriginalAssistantId = &v
	return s
}

func (s *CreateThreadRequest) SetSourceIdOfOriginalAssistantId(v string) *CreateThreadRequest {
	s.SourceIdOfOriginalAssistantId = &v
	return s
}

func (s *CreateThreadRequest) SetSourceTypeOfOriginalAssistantId(v int32) *CreateThreadRequest {
	s.SourceTypeOfOriginalAssistantId = &v
	return s
}

func (s *CreateThreadRequest) Validate() error {
	return dara.Validate(s)
}
