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
	CategoriesShrink *string `json:"Categories,omitempty" xml:"Categories,omitempty"`
	// example:
	//
	// 1
	Current *int32 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// token-xxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 3
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
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
