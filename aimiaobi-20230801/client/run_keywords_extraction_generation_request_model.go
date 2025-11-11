// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunKeywordsExtractionGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *RunKeywordsExtractionGenerationRequest
	GetPrompt() *string
	SetReferenceData(v *RunKeywordsExtractionGenerationRequestReferenceData) *RunKeywordsExtractionGenerationRequest
	GetReferenceData() *RunKeywordsExtractionGenerationRequestReferenceData
	SetTaskId(v string) *RunKeywordsExtractionGenerationRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunKeywordsExtractionGenerationRequest
	GetWorkspaceId() *string
}

type RunKeywordsExtractionGenerationRequest struct {
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	ReferenceData *RunKeywordsExtractionGenerationRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunKeywordsExtractionGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunKeywordsExtractionGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunKeywordsExtractionGenerationRequest) GetReferenceData() *RunKeywordsExtractionGenerationRequestReferenceData {
	return s.ReferenceData
}

func (s *RunKeywordsExtractionGenerationRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunKeywordsExtractionGenerationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunKeywordsExtractionGenerationRequest) SetPrompt(v string) *RunKeywordsExtractionGenerationRequest {
	s.Prompt = &v
	return s
}

func (s *RunKeywordsExtractionGenerationRequest) SetReferenceData(v *RunKeywordsExtractionGenerationRequestReferenceData) *RunKeywordsExtractionGenerationRequest {
	s.ReferenceData = v
	return s
}

func (s *RunKeywordsExtractionGenerationRequest) SetTaskId(v string) *RunKeywordsExtractionGenerationRequest {
	s.TaskId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationRequest) SetWorkspaceId(v string) *RunKeywordsExtractionGenerationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationRequest) Validate() error {
	if s.ReferenceData != nil {
		if err := s.ReferenceData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunKeywordsExtractionGenerationRequestReferenceData struct {
	// This parameter is required.
	Contents []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s RunKeywordsExtractionGenerationRequestReferenceData) String() string {
	return dara.Prettify(s)
}

func (s RunKeywordsExtractionGenerationRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationRequestReferenceData) GetContents() []*string {
	return s.Contents
}

func (s *RunKeywordsExtractionGenerationRequestReferenceData) SetContents(v []*string) *RunKeywordsExtractionGenerationRequestReferenceData {
	s.Contents = v
	return s
}

func (s *RunKeywordsExtractionGenerationRequestReferenceData) Validate() error {
	return dara.Validate(s)
}
