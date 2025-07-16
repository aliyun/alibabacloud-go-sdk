// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowStructViewContent(v bool) *CreateRunRequest
	GetAllowStructViewContent() *bool
	SetAssistantId(v string) *CreateRunRequest
	GetAssistantId() *string
	SetOriginalAssistantId(v string) *CreateRunRequest
	GetOriginalAssistantId() *string
	SetSourceIdOfOriginalAssistantId(v string) *CreateRunRequest
	GetSourceIdOfOriginalAssistantId() *string
	SetSourceTypeOfOriginalAssistantId(v string) *CreateRunRequest
	GetSourceTypeOfOriginalAssistantId() *string
	SetStream(v bool) *CreateRunRequest
	GetStream() *bool
	SetThreadId(v string) *CreateRunRequest
	GetThreadId() *string
}

type CreateRunRequest struct {
	AllowStructViewContent *bool `json:"allowStructViewContent,omitempty" xml:"allowStructViewContent,omitempty"`
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
	SourceIdOfOriginalAssistantId *string `json:"sourceIdOfOriginalAssistantId,omitempty" xml:"sourceIdOfOriginalAssistantId,omitempty"`
	// example:
	//
	// 1
	SourceTypeOfOriginalAssistantId *string `json:"sourceTypeOfOriginalAssistantId,omitempty" xml:"sourceTypeOfOriginalAssistantId,omitempty"`
	// example:
	//
	// false
	Stream *bool `json:"stream,omitempty" xml:"stream,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// threadId123
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
}

func (s CreateRunRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRunRequest) GoString() string {
	return s.String()
}

func (s *CreateRunRequest) GetAllowStructViewContent() *bool {
	return s.AllowStructViewContent
}

func (s *CreateRunRequest) GetAssistantId() *string {
	return s.AssistantId
}

func (s *CreateRunRequest) GetOriginalAssistantId() *string {
	return s.OriginalAssistantId
}

func (s *CreateRunRequest) GetSourceIdOfOriginalAssistantId() *string {
	return s.SourceIdOfOriginalAssistantId
}

func (s *CreateRunRequest) GetSourceTypeOfOriginalAssistantId() *string {
	return s.SourceTypeOfOriginalAssistantId
}

func (s *CreateRunRequest) GetStream() *bool {
	return s.Stream
}

func (s *CreateRunRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *CreateRunRequest) SetAllowStructViewContent(v bool) *CreateRunRequest {
	s.AllowStructViewContent = &v
	return s
}

func (s *CreateRunRequest) SetAssistantId(v string) *CreateRunRequest {
	s.AssistantId = &v
	return s
}

func (s *CreateRunRequest) SetOriginalAssistantId(v string) *CreateRunRequest {
	s.OriginalAssistantId = &v
	return s
}

func (s *CreateRunRequest) SetSourceIdOfOriginalAssistantId(v string) *CreateRunRequest {
	s.SourceIdOfOriginalAssistantId = &v
	return s
}

func (s *CreateRunRequest) SetSourceTypeOfOriginalAssistantId(v string) *CreateRunRequest {
	s.SourceTypeOfOriginalAssistantId = &v
	return s
}

func (s *CreateRunRequest) SetStream(v bool) *CreateRunRequest {
	s.Stream = &v
	return s
}

func (s *CreateRunRequest) SetThreadId(v string) *CreateRunRequest {
	s.ThreadId = &v
	return s
}

func (s *CreateRunRequest) Validate() error {
	return dara.Validate(s)
}
