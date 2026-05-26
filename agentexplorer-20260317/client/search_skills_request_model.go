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
	SetSkip(v int32) *SearchSkillsRequest
	GetSkip() *int32
}

type SearchSkillsRequest struct {
	// example:
	//
	// compute.serverless,network
	CategoryCode *string `json:"categoryCode,omitempty" xml:"categoryCode,omitempty"`
	// example:
	//
	// ecs
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// AAAAAZjtYxxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
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

func (s *SearchSkillsRequest) SetSkip(v int32) *SearchSkillsRequest {
	s.Skip = &v
	return s
}

func (s *SearchSkillsRequest) Validate() error {
	return dara.Validate(s)
}
