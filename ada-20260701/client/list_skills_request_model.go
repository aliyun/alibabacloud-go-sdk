// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreatorId(v string) *ListSkillsRequest
	GetCreatorId() *string
	SetMaxResults(v int32) *ListSkillsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSkillsRequest
	GetNextToken() *string
	SetPageNumber(v int64) *ListSkillsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListSkillsRequest
	GetPageSize() *int64
	SetQ(v string) *ListSkillsRequest
	GetQ() *string
	SetScope(v string) *ListSkillsRequest
	GetScope() *string
	SetVisibility(v string) *ListSkillsRequest
	GetVisibility() *string
}

type ListSkillsRequest struct {
	// Filters Skills by creator ID.
	//
	// example:
	//
	// example-user
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The number of entries per page for cursor-based pagination. Valid values: 1 to 100. Default value: `20`. If explicitly specified, cursor-based pagination takes precedence.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token returned by the server for the next page. Do not pass this parameter for the first query. For subsequent queries, use the value returned in the previous response.
	//
	// example:
	//
	// eyJzIjoiZDc3ZGRhYmE3MDMwYWM1NCIsInAiOjJ9.vkKVySx9G26993sTNLZqwGSmgciRsrRm2SgsjOMJoCQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number for compatible page-number-based pagination. Pages start from 1. Default value: `1`.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for compatible page-number-based pagination. Valid values: 1 to 100. Default value: `20`.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Performs a fuzzy match on the Skill name or description.
	//
	// example:
	//
	// review
	Q *string `json:"Q,omitempty" xml:"Q,omitempty"`
	// The query scope for Skills. Valid values: `SYSTEM` and `CUSTOM`. If omitted, both official and custom Skills are queried.
	//
	// example:
	//
	// CUSTOM
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// Filters Skills by visibility. Common values are `user` and `tenant`.
	//
	// example:
	//
	// user
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s ListSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsRequest) GoString() string {
	return s.String()
}

func (s *ListSkillsRequest) GetCreatorId() *string {
	return s.CreatorId
}

func (s *ListSkillsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListSkillsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListSkillsRequest) GetQ() *string {
	return s.Q
}

func (s *ListSkillsRequest) GetScope() *string {
	return s.Scope
}

func (s *ListSkillsRequest) GetVisibility() *string {
	return s.Visibility
}

func (s *ListSkillsRequest) SetCreatorId(v string) *ListSkillsRequest {
	s.CreatorId = &v
	return s
}

func (s *ListSkillsRequest) SetMaxResults(v int32) *ListSkillsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSkillsRequest) SetNextToken(v string) *ListSkillsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSkillsRequest) SetPageNumber(v int64) *ListSkillsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSkillsRequest) SetPageSize(v int64) *ListSkillsRequest {
	s.PageSize = &v
	return s
}

func (s *ListSkillsRequest) SetQ(v string) *ListSkillsRequest {
	s.Q = &v
	return s
}

func (s *ListSkillsRequest) SetScope(v string) *ListSkillsRequest {
	s.Scope = &v
	return s
}

func (s *ListSkillsRequest) SetVisibility(v string) *ListSkillsRequest {
	s.Visibility = &v
	return s
}

func (s *ListSkillsRequest) Validate() error {
	return dara.Validate(s)
}
