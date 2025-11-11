// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTranslateGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *RunTranslateGenerationRequest
	GetPrompt() *string
	SetReferenceData(v *RunTranslateGenerationRequestReferenceData) *RunTranslateGenerationRequest
	GetReferenceData() *RunTranslateGenerationRequestReferenceData
	SetTaskId(v string) *RunTranslateGenerationRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunTranslateGenerationRequest
	GetWorkspaceId() *string
}

type RunTranslateGenerationRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// toEnglish
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	ReferenceData *RunTranslateGenerationRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
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

func (s RunTranslateGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunTranslateGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunTranslateGenerationRequest) GetReferenceData() *RunTranslateGenerationRequestReferenceData {
	return s.ReferenceData
}

func (s *RunTranslateGenerationRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunTranslateGenerationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunTranslateGenerationRequest) SetPrompt(v string) *RunTranslateGenerationRequest {
	s.Prompt = &v
	return s
}

func (s *RunTranslateGenerationRequest) SetReferenceData(v *RunTranslateGenerationRequestReferenceData) *RunTranslateGenerationRequest {
	s.ReferenceData = v
	return s
}

func (s *RunTranslateGenerationRequest) SetTaskId(v string) *RunTranslateGenerationRequest {
	s.TaskId = &v
	return s
}

func (s *RunTranslateGenerationRequest) SetWorkspaceId(v string) *RunTranslateGenerationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunTranslateGenerationRequest) Validate() error {
	if s.ReferenceData != nil {
		if err := s.ReferenceData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunTranslateGenerationRequestReferenceData struct {
	// This parameter is required.
	Contents []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s RunTranslateGenerationRequestReferenceData) String() string {
	return dara.Prettify(s)
}

func (s RunTranslateGenerationRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunTranslateGenerationRequestReferenceData) GetContents() []*string {
	return s.Contents
}

func (s *RunTranslateGenerationRequestReferenceData) SetContents(v []*string) *RunTranslateGenerationRequestReferenceData {
	s.Contents = v
	return s
}

func (s *RunTranslateGenerationRequestReferenceData) Validate() error {
	return dara.Validate(s)
}
