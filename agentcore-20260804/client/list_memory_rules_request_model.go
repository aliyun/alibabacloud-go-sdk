// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMemoryRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *ListMemoryRulesRequest
	GetEnabled() *bool
	SetMaxResults(v int32) *ListMemoryRulesRequest
	GetMaxResults() *int32
	SetNameLike(v string) *ListMemoryRulesRequest
	GetNameLike() *string
	SetNextToken(v string) *ListMemoryRulesRequest
	GetNextToken() *string
}

type ListMemoryRulesRequest struct {
	// Filters by the enabled status.
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nameLike
	//
	// example:
	//
	// Preference
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// The pagination token returned for the same resource and filter conditions.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwMH0=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListMemoryRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMemoryRulesRequest) GoString() string {
	return s.String()
}

func (s *ListMemoryRulesRequest) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListMemoryRulesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMemoryRulesRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListMemoryRulesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMemoryRulesRequest) SetEnabled(v bool) *ListMemoryRulesRequest {
	s.Enabled = &v
	return s
}

func (s *ListMemoryRulesRequest) SetMaxResults(v int32) *ListMemoryRulesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMemoryRulesRequest) SetNameLike(v string) *ListMemoryRulesRequest {
	s.NameLike = &v
	return s
}

func (s *ListMemoryRulesRequest) SetNextToken(v string) *ListMemoryRulesRequest {
	s.NextToken = &v
	return s
}

func (s *ListMemoryRulesRequest) Validate() error {
	return dara.Validate(s)
}
