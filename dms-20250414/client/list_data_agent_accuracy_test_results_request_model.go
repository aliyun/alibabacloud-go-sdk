// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentAccuracyTestResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccuracyTestInsId(v string) *ListDataAgentAccuracyTestResultsRequest
	GetAccuracyTestInsId() *string
	SetAccuracyTestResultId(v string) *ListDataAgentAccuracyTestResultsRequest
	GetAccuracyTestResultId() *string
	SetAccuracyTestSubtaskId(v string) *ListDataAgentAccuracyTestResultsRequest
	GetAccuracyTestSubtaskId() *string
	SetAccuracyTestTaskId(v string) *ListDataAgentAccuracyTestResultsRequest
	GetAccuracyTestTaskId() *string
	SetMaxResults(v int32) *ListDataAgentAccuracyTestResultsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentAccuracyTestResultsRequest
	GetNextToken() *string
	SetPageNumber(v string) *ListDataAgentAccuracyTestResultsRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDataAgentAccuracyTestResultsRequest
	GetPageSize() *string
	SetRegionId(v string) *ListDataAgentAccuracyTestResultsRequest
	GetRegionId() *string
	SetWorkspaceId(v string) *ListDataAgentAccuracyTestResultsRequest
	GetWorkspaceId() *string
}

type ListDataAgentAccuracyTestResultsRequest struct {
	// The instance ID of the accuracy test.
	//
	// example:
	//
	// at-106n4rg17gv9fxxxxxxxxxx
	AccuracyTestInsId *string `json:"AccuracyTestInsId,omitempty" xml:"AccuracyTestInsId,omitempty"`
	// The result ID used to retrieve a single record.
	//
	// example:
	//
	// at-emhnbwewfngrxxxxxxxxxx
	AccuracyTestResultId *string `json:"AccuracyTestResultId,omitempty" xml:"AccuracyTestResultId,omitempty"`
	// The subtask ID used to filter results.
	//
	// example:
	//
	// f1eb8728-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	AccuracyTestSubtaskId *string `json:"AccuracyTestSubtaskId,omitempty" xml:"AccuracyTestSubtaskId,omitempty"`
	// The ID of the accuracy test task.
	//
	// example:
	//
	// 692abb8f-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	AccuracyTestTaskId *string `json:"AccuracyTestTaskId,omitempty" xml:"AccuracyTestTaskId,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// nu use
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// NesLoKLEdIZrKhDT7I2gS****
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
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 8wfig6l33n4f4xxxxxxxxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListDataAgentAccuracyTestResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestResultsRequest) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestResultsRequest) GetAccuracyTestInsId() *string {
	return s.AccuracyTestInsId
}

func (s *ListDataAgentAccuracyTestResultsRequest) GetAccuracyTestResultId() *string {
	return s.AccuracyTestResultId
}

func (s *ListDataAgentAccuracyTestResultsRequest) GetAccuracyTestSubtaskId() *string {
	return s.AccuracyTestSubtaskId
}

func (s *ListDataAgentAccuracyTestResultsRequest) GetAccuracyTestTaskId() *string {
	return s.AccuracyTestTaskId
}

func (s *ListDataAgentAccuracyTestResultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentAccuracyTestResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentAccuracyTestResultsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataAgentAccuracyTestResultsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataAgentAccuracyTestResultsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListDataAgentAccuracyTestResultsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataAgentAccuracyTestResultsRequest) SetAccuracyTestInsId(v string) *ListDataAgentAccuracyTestResultsRequest {
	s.AccuracyTestInsId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsRequest) SetAccuracyTestResultId(v string) *ListDataAgentAccuracyTestResultsRequest {
	s.AccuracyTestResultId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsRequest) SetAccuracyTestSubtaskId(v string) *ListDataAgentAccuracyTestResultsRequest {
	s.AccuracyTestSubtaskId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsRequest) SetAccuracyTestTaskId(v string) *ListDataAgentAccuracyTestResultsRequest {
	s.AccuracyTestTaskId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsRequest) SetMaxResults(v int32) *ListDataAgentAccuracyTestResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsRequest) SetNextToken(v string) *ListDataAgentAccuracyTestResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsRequest) SetPageNumber(v string) *ListDataAgentAccuracyTestResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsRequest) SetPageSize(v string) *ListDataAgentAccuracyTestResultsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsRequest) SetRegionId(v string) *ListDataAgentAccuracyTestResultsRequest {
	s.RegionId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsRequest) SetWorkspaceId(v string) *ListDataAgentAccuracyTestResultsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataAgentAccuracyTestResultsRequest) Validate() error {
	return dara.Validate(s)
}
