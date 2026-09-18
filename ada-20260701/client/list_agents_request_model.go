// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorId(v string) *ListAgentsRequest
	GetCreatorId() *string
	SetMaxResults(v int32) *ListAgentsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAgentsRequest
	GetNextToken() *string
	SetPageNumber(v int64) *ListAgentsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAgentsRequest
	GetPageSize() *int64
	SetQ(v string) *ListAgentsRequest
	GetQ() *string
	SetRequiredRuntime(v string) *ListAgentsRequest
	GetRequiredRuntime() *string
	SetScope(v string) *ListAgentsRequest
	GetScope() *string
	SetVisibility(v string) *ListAgentsRequest
	GetVisibility() *string
}

type ListAgentsRequest struct {
	// Filters agents by the exact creator ID.
	//
	// example:
	//
	// example-user
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The number of entries per page for cursor-based pagination. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The next page token returned in the previous response.
	//
	// example:
	//
	// eyJwYWdlIjoyfQ.example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number for page number-based pagination. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for page number-based pagination. Valid values: 1 to 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The search keyword. Matches the name, display name, or description.
	//
	// example:
	//
	// code review
	Q *string `json:"Q,omitempty" xml:"Q,omitempty"`
	// Filters agents by runtime label.
	//
	// example:
	//
	// qwen
	RequiredRuntime *string `json:"RequiredRuntime,omitempty" xml:"RequiredRuntime,omitempty"`
	// The query scope. Valid values: `SYSTEM` and `CUSTOM`.
	//
	// example:
	//
	// CUSTOM
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Filters agents by visibility scope. Valid values: `user` and `tenant`.
	//
	// example:
	//
	// user
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListAgentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsRequest) GoString() string {
	return s.String()
}

func (s *ListAgentsRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListAgentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAgentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAgentsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAgentsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAgentsRequest) GetQ() *string {
	return s.Q
}

func (s *ListAgentsRequest) GetRequiredRuntime() *string {
	return s.RequiredRuntime
}

func (s *ListAgentsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListAgentsRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *ListAgentsRequest) SetCreatorId(v string) *ListAgentsRequest {
	s.CreatorId = &v
	return s
}

func (s *ListAgentsRequest) SetMaxResults(v int32) *ListAgentsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAgentsRequest) SetNextToken(v string) *ListAgentsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAgentsRequest) SetPageNumber(v int64) *ListAgentsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAgentsRequest) SetPageSize(v int64) *ListAgentsRequest {
	s.PageSize = &v
	return s
}

func (s *ListAgentsRequest) SetQ(v string) *ListAgentsRequest {
	s.Q = &v
	return s
}

func (s *ListAgentsRequest) SetRequiredRuntime(v string) *ListAgentsRequest {
	s.RequiredRuntime = &v
	return s
}

func (s *ListAgentsRequest) SetScope(v string) *ListAgentsRequest {
	s.Scope = &v
	return s
}

func (s *ListAgentsRequest) SetVisibility(v string) *ListAgentsRequest {
	s.Visibility = &v
	return s
}

func (s *ListAgentsRequest) Validate() error {
	return dara.Validate(s)
}
