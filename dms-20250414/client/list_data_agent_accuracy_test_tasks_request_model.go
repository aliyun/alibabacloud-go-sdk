// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentAccuracyTestTasksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccuracyTestInsId(v string) *ListDataAgentAccuracyTestTasksRequest
	GetAccuracyTestInsId() *string
	SetAccuracyTestTaskId(v string) *ListDataAgentAccuracyTestTasksRequest
	GetAccuracyTestTaskId() *string
	SetMaxResults(v int32) *ListDataAgentAccuracyTestTasksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentAccuracyTestTasksRequest
	GetNextToken() *string
	SetPageNumber(v string) *ListDataAgentAccuracyTestTasksRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDataAgentAccuracyTestTasksRequest
	GetPageSize() *string
	SetWorkspaceId(v string) *ListDataAgentAccuracyTestTasksRequest
	GetWorkspaceId() *string
}

type ListDataAgentAccuracyTestTasksRequest struct {
	// The accuracy test instance ID.
	//
	// example:
	//
	// at-106n4rg17gv9fxxxxxxxxxx
	AccuracyTestInsId *string `json:"AccuracyTestInsId,omitempty" xml:"AccuracyTestInsId,omitempty"`
	// The task ID used for exact filtering.
	//
	// example:
	//
	// 692abb8f-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	AccuracyTestTaskId *string `json:"AccuracyTestTaskId,omitempty" xml:"AccuracyTestTaskId,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// NesLoK****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 8wfig6l33n4f4xxxxxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataAgentAccuracyTestTasksRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestTasksRequest) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestTasksRequest) GetAccuracyTestInsId() *string {
	return s.AccuracyTestInsId
}

func (s *ListDataAgentAccuracyTestTasksRequest) GetAccuracyTestTaskId() *string {
	return s.AccuracyTestTaskId
}

func (s *ListDataAgentAccuracyTestTasksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentAccuracyTestTasksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentAccuracyTestTasksRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataAgentAccuracyTestTasksRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataAgentAccuracyTestTasksRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataAgentAccuracyTestTasksRequest) SetAccuracyTestInsId(v string) *ListDataAgentAccuracyTestTasksRequest {
	s.AccuracyTestInsId = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksRequest) SetAccuracyTestTaskId(v string) *ListDataAgentAccuracyTestTasksRequest {
	s.AccuracyTestTaskId = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksRequest) SetMaxResults(v int32) *ListDataAgentAccuracyTestTasksRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksRequest) SetNextToken(v string) *ListDataAgentAccuracyTestTasksRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksRequest) SetPageNumber(v string) *ListDataAgentAccuracyTestTasksRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksRequest) SetPageSize(v string) *ListDataAgentAccuracyTestTasksRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksRequest) SetWorkspaceId(v string) *ListDataAgentAccuracyTestTasksRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataAgentAccuracyTestTasksRequest) Validate() error {
	return dara.Validate(s)
}
