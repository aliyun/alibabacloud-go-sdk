// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataAgentAccuracyTestInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccuracyTestInsId(v string) *ListDataAgentAccuracyTestInstancesRequest
	GetAccuracyTestInsId() *string
	SetMaxResults(v int32) *ListDataAgentAccuracyTestInstancesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListDataAgentAccuracyTestInstancesRequest
	GetNextToken() *string
	SetPageNumber(v string) *ListDataAgentAccuracyTestInstancesRequest
	GetPageNumber() *string
	SetPageSize(v string) *ListDataAgentAccuracyTestInstancesRequest
	GetPageSize() *string
	SetWorkspaceId(v string) *ListDataAgentAccuracyTestInstancesRequest
	GetWorkspaceId() *string
}

type ListDataAgentAccuracyTestInstancesRequest struct {
	// The accuracy test instance ID.
	//
	// example:
	//
	// at-106n4rg17gv9fxxxxxxxxxx
	AccuracyTestInsId *string `json:"AccuracyTestInsId,omitempty" xml:"AccuracyTestInsId,omitempty"`
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
	// Nes****
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

func (s ListDataAgentAccuracyTestInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDataAgentAccuracyTestInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListDataAgentAccuracyTestInstancesRequest) GetAccuracyTestInsId() *string {
	return s.AccuracyTestInsId
}

func (s *ListDataAgentAccuracyTestInstancesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataAgentAccuracyTestInstancesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataAgentAccuracyTestInstancesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *ListDataAgentAccuracyTestInstancesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *ListDataAgentAccuracyTestInstancesRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDataAgentAccuracyTestInstancesRequest) SetAccuracyTestInsId(v string) *ListDataAgentAccuracyTestInstancesRequest {
	s.AccuracyTestInsId = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesRequest) SetMaxResults(v int32) *ListDataAgentAccuracyTestInstancesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesRequest) SetNextToken(v string) *ListDataAgentAccuracyTestInstancesRequest {
	s.NextToken = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesRequest) SetPageNumber(v string) *ListDataAgentAccuracyTestInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesRequest) SetPageSize(v string) *ListDataAgentAccuracyTestInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesRequest) SetWorkspaceId(v string) *ListDataAgentAccuracyTestInstancesRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListDataAgentAccuracyTestInstancesRequest) Validate() error {
	return dara.Validate(s)
}
