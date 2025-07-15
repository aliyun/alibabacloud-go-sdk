// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunGenerateQuestionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocId(v string) *RunGenerateQuestionsRequest
	GetDocId() *string
	SetModelName(v string) *RunGenerateQuestionsRequest
	GetModelName() *string
	SetReferenceContent(v string) *RunGenerateQuestionsRequest
	GetReferenceContent() *string
	SetSessionId(v string) *RunGenerateQuestionsRequest
	GetSessionId() *string
	SetWorkspaceId(v string) *RunGenerateQuestionsRequest
	GetWorkspaceId() *string
}

type RunGenerateQuestionsRequest struct {
	// example:
	//
	// oOgIwodFANW1u5MnqxysOh1rtld3xn
	DocId            *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	ModelName        *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	ReferenceContent *string `json:"ReferenceContent,omitempty" xml:"ReferenceContent,omitempty"`
	// example:
	//
	// f486c4e2-b773-4d65-88f8-2ba540610456
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-w335gauzlbba2vze
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunGenerateQuestionsRequest) String() string {
	return dara.Prettify(s)
}

func (s RunGenerateQuestionsRequest) GoString() string {
	return s.String()
}

func (s *RunGenerateQuestionsRequest) GetDocId() *string {
	return s.DocId
}

func (s *RunGenerateQuestionsRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RunGenerateQuestionsRequest) GetReferenceContent() *string {
	return s.ReferenceContent
}

func (s *RunGenerateQuestionsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunGenerateQuestionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunGenerateQuestionsRequest) SetDocId(v string) *RunGenerateQuestionsRequest {
	s.DocId = &v
	return s
}

func (s *RunGenerateQuestionsRequest) SetModelName(v string) *RunGenerateQuestionsRequest {
	s.ModelName = &v
	return s
}

func (s *RunGenerateQuestionsRequest) SetReferenceContent(v string) *RunGenerateQuestionsRequest {
	s.ReferenceContent = &v
	return s
}

func (s *RunGenerateQuestionsRequest) SetSessionId(v string) *RunGenerateQuestionsRequest {
	s.SessionId = &v
	return s
}

func (s *RunGenerateQuestionsRequest) SetWorkspaceId(v string) *RunGenerateQuestionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunGenerateQuestionsRequest) Validate() error {
	return dara.Validate(s)
}
