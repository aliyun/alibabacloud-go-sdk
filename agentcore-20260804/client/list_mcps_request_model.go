// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMcpsRequest
	GetMaxResults() *int32
	SetName(v string) *ListMcpsRequest
	GetName() *string
	SetNextToken(v string) *ListMcpsRequest
	GetNextToken() *string
	SetOfficialTag(v string) *ListMcpsRequest
	GetOfficialTag() *string
	SetSearchType(v string) *ListMcpsRequest
	GetSearchType() *string
	SetUsageActive(v bool) *ListMcpsRequest
	GetUsageActive() *bool
}

type ListMcpsRequest struct {
	// The maximum number of entries per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The MCP service name or service ID. Used together with SearchType.
	//
	// example:
	//
	// my-mcp-server
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// next-page-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Filters results by official usage tag.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	OfficialTag *string `json:"officialTag,omitempty" xml:"officialTag,omitempty"`
	// The name matching method. Takes effect only when Name is specified. Valid values:
	//
	// - accurate: exact match.
	//
	// - blur: fuzzy match.
	//
	// Default value: blur.
	//
	// example:
	//
	// blur
	SearchType *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
	// Specifies whether the service is still bound by the official template usage constraint.
	UsageActive *bool `json:"usageActive,omitempty" xml:"usageActive,omitempty"`
}

func (s ListMcpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsRequest) GoString() string {
	return s.String()
}

func (s *ListMcpsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpsRequest) GetName() *string {
	return s.Name
}

func (s *ListMcpsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpsRequest) GetOfficialTag() *string {
	return s.OfficialTag
}

func (s *ListMcpsRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *ListMcpsRequest) GetUsageActive() *bool {
	return s.UsageActive
}

func (s *ListMcpsRequest) SetMaxResults(v int32) *ListMcpsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMcpsRequest) SetName(v string) *ListMcpsRequest {
	s.Name = &v
	return s
}

func (s *ListMcpsRequest) SetNextToken(v string) *ListMcpsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMcpsRequest) SetOfficialTag(v string) *ListMcpsRequest {
	s.OfficialTag = &v
	return s
}

func (s *ListMcpsRequest) SetSearchType(v string) *ListMcpsRequest {
	s.SearchType = &v
	return s
}

func (s *ListMcpsRequest) SetUsageActive(v bool) *ListMcpsRequest {
	s.UsageActive = &v
	return s
}

func (s *ListMcpsRequest) Validate() error {
	return dara.Validate(s)
}
