// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunDocWashingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelId(v string) *RunDocWashingRequest
	GetModelId() *string
	SetPrompt(v string) *RunDocWashingRequest
	GetPrompt() *string
	SetReferenceContent(v string) *RunDocWashingRequest
	GetReferenceContent() *string
	SetSessionId(v string) *RunDocWashingRequest
	GetSessionId() *string
	SetTopic(v string) *RunDocWashingRequest
	GetTopic() *string
	SetWordNumber(v int32) *RunDocWashingRequest
	GetWordNumber() *int32
	SetWorkspaceId(v string) *RunDocWashingRequest
	GetWorkspaceId() *string
	SetWritingTypeName(v string) *RunDocWashingRequest
	GetWritingTypeName() *string
	SetWritingTypeRefDoc(v string) *RunDocWashingRequest
	GetWritingTypeRefDoc() *string
}

type RunDocWashingRequest struct {
	ModelId *string `json:"ModelId,omitempty" xml:"ModelId,omitempty"`
	Prompt  *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	ReferenceContent *string `json:"ReferenceContent,omitempty" xml:"ReferenceContent,omitempty"`
	SessionId        *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	Topic            *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
	// example:
	//
	// 500
	WordNumber *int32 `json:"WordNumber,omitempty" xml:"WordNumber,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-2setzb9x4ewsd
	WorkspaceId       *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	WritingTypeName   *string `json:"WritingTypeName,omitempty" xml:"WritingTypeName,omitempty"`
	WritingTypeRefDoc *string `json:"WritingTypeRefDoc,omitempty" xml:"WritingTypeRefDoc,omitempty"`
}

func (s RunDocWashingRequest) String() string {
	return dara.Prettify(s)
}

func (s RunDocWashingRequest) GoString() string {
	return s.String()
}

func (s *RunDocWashingRequest) GetModelId() *string {
	return s.ModelId
}

func (s *RunDocWashingRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunDocWashingRequest) GetReferenceContent() *string {
	return s.ReferenceContent
}

func (s *RunDocWashingRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunDocWashingRequest) GetTopic() *string {
	return s.Topic
}

func (s *RunDocWashingRequest) GetWordNumber() *int32 {
	return s.WordNumber
}

func (s *RunDocWashingRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunDocWashingRequest) GetWritingTypeName() *string {
	return s.WritingTypeName
}

func (s *RunDocWashingRequest) GetWritingTypeRefDoc() *string {
	return s.WritingTypeRefDoc
}

func (s *RunDocWashingRequest) SetModelId(v string) *RunDocWashingRequest {
	s.ModelId = &v
	return s
}

func (s *RunDocWashingRequest) SetPrompt(v string) *RunDocWashingRequest {
	s.Prompt = &v
	return s
}

func (s *RunDocWashingRequest) SetReferenceContent(v string) *RunDocWashingRequest {
	s.ReferenceContent = &v
	return s
}

func (s *RunDocWashingRequest) SetSessionId(v string) *RunDocWashingRequest {
	s.SessionId = &v
	return s
}

func (s *RunDocWashingRequest) SetTopic(v string) *RunDocWashingRequest {
	s.Topic = &v
	return s
}

func (s *RunDocWashingRequest) SetWordNumber(v int32) *RunDocWashingRequest {
	s.WordNumber = &v
	return s
}

func (s *RunDocWashingRequest) SetWorkspaceId(v string) *RunDocWashingRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunDocWashingRequest) SetWritingTypeName(v string) *RunDocWashingRequest {
	s.WritingTypeName = &v
	return s
}

func (s *RunDocWashingRequest) SetWritingTypeRefDoc(v string) *RunDocWashingRequest {
	s.WritingTypeRefDoc = &v
	return s
}

func (s *RunDocWashingRequest) Validate() error {
	return dara.Validate(s)
}
