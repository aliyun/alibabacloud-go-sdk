// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnalysisTagDetailByTaskIdShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoriesShrink(v string) *ListAnalysisTagDetailByTaskIdShrinkRequest
	GetCategoriesShrink() *string
	SetCurrent(v int32) *ListAnalysisTagDetailByTaskIdShrinkRequest
	GetCurrent() *int32
	SetMaxResults(v int32) *ListAnalysisTagDetailByTaskIdShrinkRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAnalysisTagDetailByTaskIdShrinkRequest
	GetNextToken() *string
	SetSize(v int32) *ListAnalysisTagDetailByTaskIdShrinkRequest
	GetSize() *int32
	SetTaskId(v string) *ListAnalysisTagDetailByTaskIdShrinkRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *ListAnalysisTagDetailByTaskIdShrinkRequest
	GetWorkspaceId() *string
}

type ListAnalysisTagDetailByTaskIdShrinkRequest struct {
	// A list of categories for filtering.
	CategoriesShrink *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// The maximum number of entries to return on each page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token for the next page of results.
	//
	// example:
	//
	// token-xxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The number of records to request.
	//
	// example:
	//
	// 3
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The unique ID of the task.
	//
	// > This parameter is optional. The system automatically generates a task ID. If you specify the same TaskId for multiple tasks, the tasks are considered part of the same conversation.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The unique ID of the Alibaba Cloud Model Studio workspace. For more information, see [Get a Workspace ID](https://help.aliyun.com/document_detail/2782167.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListAnalysisTagDetailByTaskIdShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisTagDetailByTaskIdShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) GetCategoriesShrink() *string {
	return s.CategoriesShrink
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) SetCategoriesShrink(v string) *ListAnalysisTagDetailByTaskIdShrinkRequest {
	s.CategoriesShrink = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) SetCurrent(v int32) *ListAnalysisTagDetailByTaskIdShrinkRequest {
	s.Current = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) SetMaxResults(v int32) *ListAnalysisTagDetailByTaskIdShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) SetNextToken(v string) *ListAnalysisTagDetailByTaskIdShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) SetSize(v int32) *ListAnalysisTagDetailByTaskIdShrinkRequest {
	s.Size = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) SetTaskId(v string) *ListAnalysisTagDetailByTaskIdShrinkRequest {
	s.TaskId = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) SetWorkspaceId(v string) *ListAnalysisTagDetailByTaskIdShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdShrinkRequest) Validate() error {
	return dara.Validate(s)
}
