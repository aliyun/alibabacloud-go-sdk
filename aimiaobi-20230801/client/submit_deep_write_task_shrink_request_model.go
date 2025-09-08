// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDeepWriteTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgentOrchestrationShrink(v string) *SubmitDeepWriteTaskShrinkRequest
	GetAgentOrchestrationShrink() *string
	SetFilesShrink(v string) *SubmitDeepWriteTaskShrinkRequest
	GetFilesShrink() *string
	SetInput(v string) *SubmitDeepWriteTaskShrinkRequest
	GetInput() *string
	SetInstructions(v string) *SubmitDeepWriteTaskShrinkRequest
	GetInstructions() *string
	SetWorkspaceId(v string) *SubmitDeepWriteTaskShrinkRequest
	GetWorkspaceId() *string
}

type SubmitDeepWriteTaskShrinkRequest struct {
	AgentOrchestrationShrink *string `json:"AgentOrchestration,omitempty" xml:"AgentOrchestration,omitempty"`
	FilesShrink              *string `json:"Files,omitempty" xml:"Files,omitempty"`
	// This parameter is required.
	//
	// if can be null:
	// false
	Input        *string `json:"Input,omitempty" xml:"Input,omitempty"`
	Instructions *string `json:"Instructions,omitempty" xml:"Instructions,omitempty"`
	// example:
	//
	// llm-1setzb9xb8m11vrc
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitDeepWriteTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDeepWriteTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitDeepWriteTaskShrinkRequest) GetAgentOrchestrationShrink() *string {
	return s.AgentOrchestrationShrink
}

func (s *SubmitDeepWriteTaskShrinkRequest) GetFilesShrink() *string {
	return s.FilesShrink
}

func (s *SubmitDeepWriteTaskShrinkRequest) GetInput() *string {
	return s.Input
}

func (s *SubmitDeepWriteTaskShrinkRequest) GetInstructions() *string {
	return s.Instructions
}

func (s *SubmitDeepWriteTaskShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitDeepWriteTaskShrinkRequest) SetAgentOrchestrationShrink(v string) *SubmitDeepWriteTaskShrinkRequest {
	s.AgentOrchestrationShrink = &v
	return s
}

func (s *SubmitDeepWriteTaskShrinkRequest) SetFilesShrink(v string) *SubmitDeepWriteTaskShrinkRequest {
	s.FilesShrink = &v
	return s
}

func (s *SubmitDeepWriteTaskShrinkRequest) SetInput(v string) *SubmitDeepWriteTaskShrinkRequest {
	s.Input = &v
	return s
}

func (s *SubmitDeepWriteTaskShrinkRequest) SetInstructions(v string) *SubmitDeepWriteTaskShrinkRequest {
	s.Instructions = &v
	return s
}

func (s *SubmitDeepWriteTaskShrinkRequest) SetWorkspaceId(v string) *SubmitDeepWriteTaskShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitDeepWriteTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
