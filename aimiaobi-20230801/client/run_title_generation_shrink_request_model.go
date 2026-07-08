// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTitleGenerationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeduplicatedTitlesShrink(v string) *RunTitleGenerationShrinkRequest
	GetDeduplicatedTitlesShrink() *string
	SetReferenceDataShrink(v string) *RunTitleGenerationShrinkRequest
	GetReferenceDataShrink() *string
	SetTaskId(v string) *RunTitleGenerationShrinkRequest
	GetTaskId() *string
	SetTitleCount(v string) *RunTitleGenerationShrinkRequest
	GetTitleCount() *string
	SetWorkspaceId(v string) *RunTitleGenerationShrinkRequest
	GetWorkspaceId() *string
}

type RunTitleGenerationShrinkRequest struct {
	// A collection of titles to deduplicate against the newly generated titles. The total character count for all titles must not exceed 5K.
	DeduplicatedTitlesShrink *string `json:"DeduplicatedTitles,omitempty" xml:"DeduplicatedTitles,omitempty"`
	// Data for title generation.
	//
	// This parameter is required.
	ReferenceDataShrink *string `json:"ReferenceData,omitempty" xml:"ReferenceData,omitempty"`
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

func (s RunTitleGenerationShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunTitleGenerationShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationShrinkRequest) GetDeduplicatedTitlesShrink() *string {
	return s.DeduplicatedTitlesShrink
}

func (s *RunTitleGenerationShrinkRequest) GetReferenceDataShrink() *string {
	return s.ReferenceDataShrink
}

func (s *RunTitleGenerationShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *RunTitleGenerationShrinkRequest) GetTitleCount() *string {
	return s.TitleCount
}

func (s *RunTitleGenerationShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *RunTitleGenerationShrinkRequest) SetDeduplicatedTitlesShrink(v string) *RunTitleGenerationShrinkRequest {
	s.DeduplicatedTitlesShrink = &v
	return s
}

func (s *RunTitleGenerationShrinkRequest) SetReferenceDataShrink(v string) *RunTitleGenerationShrinkRequest {
	s.ReferenceDataShrink = &v
	return s
}

func (s *RunTitleGenerationShrinkRequest) SetTaskId(v string) *RunTitleGenerationShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *RunTitleGenerationShrinkRequest) SetTitleCount(v string) *RunTitleGenerationShrinkRequest {
	s.TitleCount = &v
	return s
}

func (s *RunTitleGenerationShrinkRequest) SetWorkspaceId(v string) *RunTitleGenerationShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *RunTitleGenerationShrinkRequest) Validate() error {
	return dara.Validate(s)
}
