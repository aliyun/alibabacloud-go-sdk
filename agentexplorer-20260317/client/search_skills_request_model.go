// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchSkillsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryCode(v string) *SearchSkillsRequest
	GetCategoryCode() *string
	SetKeyword(v string) *SearchSkillsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *SearchSkillsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *SearchSkillsRequest
	GetNextToken() *string
	SetSearchMode(v string) *SearchSkillsRequest
	GetSearchMode() *string
	SetSkip(v int32) *SearchSkillsRequest
	GetSkip() *int32
}

type SearchSkillsRequest struct {
	// The skill category code. Separate multiple codes with commas. For a second-level category, use the format: first-level category.second-level category.
	//
	// example:
	//
	// compute.serverless,network
	CategoryCode *string `json:"categoryCode,omitempty" xml:"categoryCode,omitempty"`
	// The search keyword.
	//
	// example:
	//
	// ecs
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of entries per page for a paged query. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next query. Set this to the NextToken value returned by the previous API call.
	//
	// example:
	//
	// AAAAAZjtYxxxxxxxx
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	SearchMode *string `json:"searchMode,omitempty" xml:"searchMode,omitempty"`
	// The number of entries to skip for pagination.
	//
	// example:
	//
	// 10
	Skip *int32 `json:"skip,omitempty" xml:"skip,omitempty"`
}

func (s SearchSkillsRequest) String() string {
	return dara.Prettify(s)
}

func (s SearchSkillsRequest) GoString() string {
	return s.String()
}

func (s *SearchSkillsRequest) GetCategoryCode() *string {
	return s.CategoryCode
}

func (s *SearchSkillsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *SearchSkillsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *SearchSkillsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *SearchSkillsRequest) GetSearchMode() *string {
	return s.SearchMode
}

func (s *SearchSkillsRequest) GetSkip() *int32 {
	return s.Skip
}

func (s *SearchSkillsRequest) SetCategoryCode(v string) *SearchSkillsRequest {
	s.CategoryCode = &v
	return s
}

func (s *SearchSkillsRequest) SetKeyword(v string) *SearchSkillsRequest {
	s.Keyword = &v
	return s
}

func (s *SearchSkillsRequest) SetMaxResults(v int32) *SearchSkillsRequest {
	s.MaxResults = &v
	return s
}

func (s *SearchSkillsRequest) SetNextToken(v string) *SearchSkillsRequest {
	s.NextToken = &v
	return s
}

func (s *SearchSkillsRequest) SetSearchMode(v string) *SearchSkillsRequest {
	s.SearchMode = &v
	return s
}

func (s *SearchSkillsRequest) SetSkip(v int32) *SearchSkillsRequest {
	s.Skip = &v
	return s
}

func (s *SearchSkillsRequest) Validate() error {
	return dara.Validate(s)
}
