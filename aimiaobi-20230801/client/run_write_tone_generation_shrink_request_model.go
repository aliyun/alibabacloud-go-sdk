// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWriteToneGenerationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *RunWriteToneGenerationShrinkRequest
	GetPrompt() *string
	SetReferenceDataShrink(v string) *RunWriteToneGenerationShrinkRequest
	GetReferenceDataShrink() *string
	SetTaskId(v string) *RunWriteToneGenerationShrinkRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunWriteToneGenerationShrinkRequest
	GetWorkspaceId() *string
}

type RunWriteToneGenerationShrinkRequest struct {
	// This parameter is required.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
	// example:
	//
	// 7AA2AE16-D873-5C5F-9708-15396C382EB1
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunWriteToneGenerationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunWriteToneGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationShrinkRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunWriteToneGenerationShrinkRequest) GetReferenceDataShrink() *string {
	return s.ReferenceDataShrink
}

func (s *RunWriteToneGenerationShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunWriteToneGenerationShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunWriteToneGenerationShrinkRequest) SetPrompt(v string) *RunWriteToneGenerationShrinkRequest {
	s.Prompt = &v
	return s
}

func (s *RunWriteToneGenerationShrinkRequest) SetReferenceDataShrink(v string) *RunWriteToneGenerationShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunWriteToneGenerationShrinkRequest) SetTaskId(v string) *RunWriteToneGenerationShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunWriteToneGenerationShrinkRequest) SetWorkspaceId(v string) *RunWriteToneGenerationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunWriteToneGenerationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
