// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTag(v string) *ListMcpsRequest
	GetCustomTag() *string
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
	// The custom tag used for filtering. An exact match is required.
	//
	// example:
	//
	// Knowledge base
	CustomTag *string `json:"customTag,omitempty" xml:"customTag,omitempty"`
	// The maximum number of entries to return per page. Valid values: 1 to 100. Default value: 10.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The MCP service name or service ID. This parameter is used in conjunction with SearchType.
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
	// The official usage tag used for filtering.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	OfficialTag *string `json:"officialTag,omitempty" xml:"officialTag,omitempty"`
	// The name match type. This parameter takes effect only when Name is specified. Valid values: accurate (exact match) and blur (fuzzy match). Default value: blur.
	//
	// example:
	//
	// blur
	SearchType *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
	// Specifies whether the service is still subject to official template usage constraints.
	UsageActive *bool `json:"usageActive,omitempty" xml:"usageActive,omitempty"`
}

func (s ListMcpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpsRequest) GoString() string {
	return s.String()
}

func (s *ListMcpsRequest) GetCustomTag() *string {
	return s.CustomTag
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

func (s *ListMcpsRequest) SetCustomTag(v string) *ListMcpsRequest {
	s.CustomTag = &v
	return s
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
