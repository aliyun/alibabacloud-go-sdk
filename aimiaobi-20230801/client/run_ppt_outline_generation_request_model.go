// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPptOutlineGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *RunPptOutlineGenerationRequest
	GetPrompt() *string
	SetWorkspaceId(v string) *RunPptOutlineGenerationRequest
	GetWorkspaceId() *string
}

type RunPptOutlineGenerationRequest struct {
	// This parameter is required.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-8v8a5mfpxffrj1pn
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunPptOutlineGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunPptOutlineGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunPptOutlineGenerationRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunPptOutlineGenerationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunPptOutlineGenerationRequest) SetPrompt(v string) *RunPptOutlineGenerationRequest {
	s.Prompt = &v
	return s
}

func (s *RunPptOutlineGenerationRequest) SetWorkspaceId(v string) *RunPptOutlineGenerationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunPptOutlineGenerationRequest) Validate() error {
	return dara.Validate(s)
}
