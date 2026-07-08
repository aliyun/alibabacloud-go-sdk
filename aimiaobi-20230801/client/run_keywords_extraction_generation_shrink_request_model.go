// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunKeywordsExtractionGenerationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *RunKeywordsExtractionGenerationShrinkRequest
	GetPrompt() *string
	SetReferenceDataShrink(v string) *RunKeywordsExtractionGenerationShrinkRequest
	GetReferenceDataShrink() *string
	SetTaskId(v string) *RunKeywordsExtractionGenerationShrinkRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunKeywordsExtractionGenerationShrinkRequest
	GetWorkspaceId() *string
}

type RunKeywordsExtractionGenerationShrinkRequest struct {
	// Custom prompt.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// Data required for generation.
	//
	// This parameter is required.
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
	// The unique identifier for the associated creation article.
	//
	// > TaskId is not required by default; the system automatically generates it. If subsequent tasks use the same TaskId, they belong to the same conversation group.
	//
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The unique identifier for the Alibaba Cloud Model Studio workspace. Obtain the [Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunKeywordsExtractionGenerationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunKeywordsExtractionGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) GetReferenceDataShrink() *string {
	return s.ReferenceDataShrink
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) SetPrompt(v string) *RunKeywordsExtractionGenerationShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) SetReferenceDataShrink(v string) *RunKeywordsExtractionGenerationShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) SetTaskId(v string) *RunKeywordsExtractionGenerationShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) SetWorkspaceId(v string) *RunKeywordsExtractionGenerationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunKeywordsExtractionGenerationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
