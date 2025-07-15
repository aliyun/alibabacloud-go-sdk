// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunHotwordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDocId(v string) *RunHotwordRequest
	GetDocId() *string
	SetModelName(v string) *RunHotwordRequest
	GetModelName() *string
	SetPrompt(v string) *RunHotwordRequest
	GetPrompt() *string
	SetReferenceContent(v string) *RunHotwordRequest
	GetReferenceContent() *string
	SetSessionId(v string) *RunHotwordRequest
	GetSessionId() *string
	SetWorkspaceId(v string) *RunHotwordRequest
	GetWorkspaceId() *string
}

type RunHotwordRequest struct {
	// example:
	//
	// 84ufBYEeLMZOjRFo84HJ7ySL3Efr55
	DocId            *string `json:"DocId,omitempty" xml:"DocId,omitempty"`
	ModelName        *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	Prompt           *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	ReferenceContent *string `json:"ReferenceContent,omitempty" xml:"ReferenceContent,omitempty"`
	// example:
	//
	// e32a1a3f-1f7e-41dd-b888-ef1d91b96d1e
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-baw8as25ll3wnzjr
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunHotwordRequest) String() string {
	return dara.Prettify(s)
}

func (s RunHotwordRequest) GoString() string {
	return s.String()
}

func (s *RunHotwordRequest) GetDocId() *string {
	return s.DocId
}

func (s *RunHotwordRequest) GetModelName() *string {
	return s.ModelName
}

func (s *RunHotwordRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunHotwordRequest) GetReferenceContent() *string {
	return s.ReferenceContent
}

func (s *RunHotwordRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *RunHotwordRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunHotwordRequest) SetDocId(v string) *RunHotwordRequest {
	s.DocId = &v
	return s
}

func (s *RunHotwordRequest) SetModelName(v string) *RunHotwordRequest {
	s.ModelName = &v
	return s
}

func (s *RunHotwordRequest) SetPrompt(v string) *RunHotwordRequest {
	s.Prompt = &v
	return s
}

func (s *RunHotwordRequest) SetReferenceContent(v string) *RunHotwordRequest {
	s.ReferenceContent = &v
	return s
}

func (s *RunHotwordRequest) SetSessionId(v string) *RunHotwordRequest {
	s.SessionId = &v
	return s
}

func (s *RunHotwordRequest) SetWorkspaceId(v string) *RunHotwordRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunHotwordRequest) Validate() error {
	return dara.Validate(s)
}
