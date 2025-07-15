// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTitleGenerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeduplicatedTitles(v []*string) *RunTitleGenerationRequest
	GetDeduplicatedTitles() []*string
	SetReferenceData(v *RunTitleGenerationRequestReferenceData) *RunTitleGenerationRequest
	GetReferenceData() *RunTitleGenerationRequestReferenceData
	SetTaskId(v string) *RunTitleGenerationRequest
	GetTaskId() *string
	SetTitleCount(v string) *RunTitleGenerationRequest
	GetTitleCount() *string
	SetWorkspaceId(v string) *RunTitleGenerationRequest
	GetWorkspaceId() *string
}

type RunTitleGenerationRequest struct {
	DeduplicatedTitles []*string `json:"DeduplicatedTitles,omitempty" xml:"DeduplicatedTitles,omitempty" type:"Repeated"`
	// This parameter is required.
	ReferenceData *RunTitleGenerationRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
	// example:
	//
	// xxxx
	TaskId     *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	TitleCount *string `json:"TitleCount,omitempty" xml:"TitleCount,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s RunTitleGenerationRequest) String() string {
	return dara.Prettify(s)
}

func (s RunTitleGenerationRequest) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationRequest) GetDeduplicatedTitles() []*string {
	return s.DeduplicatedTitles
}

func (s *RunTitleGenerationRequest) GetReferenceData() *RunTitleGenerationRequestReferenceData {
	return s.ReferenceData
}

func (s *RunTitleGenerationRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunTitleGenerationRequest) GetTitleCount() *string {
	return s.TitleCount
}

func (s *RunTitleGenerationRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunTitleGenerationRequest) SetDeduplicatedTitles(v []*string) *RunTitleGenerationRequest {
	s.DeduplicatedTitles = v
	return s
}

func (s *RunTitleGenerationRequest) SetReferenceData(v *RunTitleGenerationRequestReferenceData) *RunTitleGenerationRequest {
	s.ReferenceData = v
	return s
}

func (s *RunTitleGenerationRequest) SetTaskId(v string) *RunTitleGenerationRequest {
	s.TaskId = &v
	return s
}

func (s *RunTitleGenerationRequest) SetTitleCount(v string) *RunTitleGenerationRequest {
	s.TitleCount = &v
	return s
}

func (s *RunTitleGenerationRequest) SetWorkspaceId(v string) *RunTitleGenerationRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunTitleGenerationRequest) Validate() error {
	return dara.Validate(s)
}

type RunTitleGenerationRequestReferenceData struct {
	// This parameter is required.
	Contents []*string `json:"Contents,omitempty" xml:"Contents,omitempty" type:"Repeated"`
}

func (s RunTitleGenerationRequestReferenceData) String() string {
	return dara.Prettify(s)
}

func (s RunTitleGenerationRequestReferenceData) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationRequestReferenceData) GetContents() []*string {
	return s.Contents
}

func (s *RunTitleGenerationRequestReferenceData) SetContents(v []*string) *RunTitleGenerationRequestReferenceData {
	s.Contents = v
	return s
}

func (s *RunTitleGenerationRequestReferenceData) Validate() error {
	return dara.Validate(s)
}
