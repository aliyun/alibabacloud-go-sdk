// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunWriteToneGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrompt(v string) *RunWriteToneGenerationRequest
	GetPrompt() *string
	SetReferenceData(v *RunWriteToneGenerationRequestReferenceData) *RunWriteToneGenerationRequest
	GetReferenceData() *RunWriteToneGenerationRequestReferenceData
	SetTaskId(v string) *RunWriteToneGenerationRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *RunWriteToneGenerationRequest
	GetWorkspaceId() *string
}

type RunWriteToneGenerationRequest struct {
	// This parameter is required.
	Prompt *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// This parameter is required.
	ReferenceData *RunWriteToneGenerationRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
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

func (s RunWriteToneGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunWriteToneGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *RunWriteToneGenerationRequest) GetReferenceData() *RunWriteToneGenerationRequestReferenceData {
	return s.ReferenceData
}

func (s *RunWriteToneGenerationRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunWriteToneGenerationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunWriteToneGenerationRequest) SetPrompt(v string) *RunWriteToneGenerationRequest {
	s.Prompt = &v
	return s
}

func (s *RunWriteToneGenerationRequest) SetReferenceData(v *RunWriteToneGenerationRequestReferenceData) *RunWriteToneGenerationRequest {
	s.ReferenceData = v
	return s
}

func (s *RunWriteToneGenerationRequest) SetTaskId(v string) *RunWriteToneGenerationRequest {
	s.TaskId = &v
	return s
}

func (s *RunWriteToneGenerationRequest) SetWorkspaceId(v string) *RunWriteToneGenerationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunWriteToneGenerationRequest) Validate() error {
	return dara.Validate(s)
}

type RunWriteToneGenerationRequestReferenceData struct {
	// This parameter is required.
	Contents []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s RunWriteToneGenerationRequestReferenceData) String() string {
	return dara.Prettify(s)
}

func (s RunWriteToneGenerationRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunWriteToneGenerationRequestReferenceData) GetContents() []*string {
	return s.Contents
}

func (s *RunWriteToneGenerationRequestReferenceData) SetContents(v []*string) *RunWriteToneGenerationRequestReferenceData {
	s.Contents = v
	return s
}

func (s *RunWriteToneGenerationRequestReferenceData) Validate() error {
	return dara.Validate(s)
}
