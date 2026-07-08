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
	// A collection of titles to deduplicate against the newly generated titles. The total character count for all titles must not exceed 5K.
	DeduplicatedTitles []*string `json:"DeduplicatedTitles,omitempty" xml:"DeduplicatedTitles,omitempty" type:"Repeated"`
	// Data for title generation.
	//
	// This parameter is required.
	ReferenceData *RunTitleGenerationRequestReferenceData `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty" type:"Struct"`
	// The unique identifier for the associated creative article.
	//
	// > The system automatically generates the TaskId by default. You do not need to specify it. If subsequent tasks use the same TaskId, they belong to the same conversation group.
	//
	// example:
	//
	// xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Number of titles to generate, maximum 10.
	//
	// example:
	//
	// 10
	TitleCount *string `json:"TitleCount,omitempty" xml:"TitleCount,omitempty"`
	// The unique identifier for the Alibaba Cloud Model Studio workspace. For more information, see [Get the workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
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
	if s.ReferenceData != nil {
		if err := s.ReferenceData.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunTitleGenerationRequestReferenceData struct {
	// List of main content.
	//
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
