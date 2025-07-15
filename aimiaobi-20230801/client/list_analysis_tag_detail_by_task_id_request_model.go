// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnalysisTagDetailByTaskIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategories(v []*string) *ListAnalysisTagDetailByTaskIdRequest
	GetCategories() []*string
	SetCurrent(v int32) *ListAnalysisTagDetailByTaskIdRequest
	GetCurrent() *int32
	SetMaxResults(v int32) *ListAnalysisTagDetailByTaskIdRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAnalysisTagDetailByTaskIdRequest
	GetNextToken() *string
	SetSize(v int32) *ListAnalysisTagDetailByTaskIdRequest
	GetSize() *int32
	SetTaskId(v string) *ListAnalysisTagDetailByTaskIdRequest
	GetTaskId() *string
	SetWorkspaceId(v string) *ListAnalysisTagDetailByTaskIdRequest
	GetWorkspaceId() *string
}

type ListAnalysisTagDetailByTaskIdRequest struct {
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
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

func (s ListAnalysisTagDetailByTaskIdRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisTagDetailByTaskIdRequest) GoString() string {
	return s.String()
}

func (s *ListAnalysisTagDetailByTaskIdRequest) GetCategories() []*string {
	return s.Categories
}

func (s *ListAnalysisTagDetailByTaskIdRequest) GetCurrent() *int32 {
	return s.Current
}

func (s *ListAnalysisTagDetailByTaskIdRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAnalysisTagDetailByTaskIdRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAnalysisTagDetailByTaskIdRequest) GetSize() *int32 {
	return s.Size
}

func (s *ListAnalysisTagDetailByTaskIdRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *ListAnalysisTagDetailByTaskIdRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListAnalysisTagDetailByTaskIdRequest) SetCategories(v []*string) *ListAnalysisTagDetailByTaskIdRequest {
	s.Categories = v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdRequest) SetCurrent(v int32) *ListAnalysisTagDetailByTaskIdRequest {
	s.Current = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdRequest) SetMaxResults(v int32) *ListAnalysisTagDetailByTaskIdRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdRequest) SetNextToken(v string) *ListAnalysisTagDetailByTaskIdRequest {
	s.NextToken = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdRequest) SetSize(v int32) *ListAnalysisTagDetailByTaskIdRequest {
	s.Size = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdRequest) SetTaskId(v string) *ListAnalysisTagDetailByTaskIdRequest {
	s.TaskId = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdRequest) SetWorkspaceId(v string) *ListAnalysisTagDetailByTaskIdRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdRequest) Validate() error {
	return dara.Validate(s)
}
