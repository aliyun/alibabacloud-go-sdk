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
	// A list of categories for filtering.
	Categories []*string `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
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
