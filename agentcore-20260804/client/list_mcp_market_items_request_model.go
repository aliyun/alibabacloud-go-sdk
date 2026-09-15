// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpMarketItemsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyword(v string) *ListMcpMarketItemsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListMcpMarketItemsRequest
	GetMaxResults() *int32
	SetMcpType(v string) *ListMcpMarketItemsRequest
	GetMcpType() *string
	SetNextToken(v string) *ListMcpMarketItemsRequest
	GetNextToken() *string
	SetOfficialTag(v string) *ListMcpMarketItemsRequest
	GetOfficialTag() *string
}

type ListMcpMarketItemsRequest struct {
	// The keyword used to filter MCP marketplace templates.
	//
	// example:
	//
	// Knowledge
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of records to return in this query.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The MCP type.
	//
	// example:
	//
	// CODE_PACKAGE
	McpType *string `json:"mcpType,omitempty" xml:"mcpType,omitempty"`
	// The pagination token used to retrieve the next page of results.
	//
	// example:
	//
	// 20
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The official usage tag.
	//
	// example:
	//
	// KNOWLEDGE_BASE
	OfficialTag *string `json:"officialTag,omitempty" xml:"officialTag,omitempty"`
}

func (s ListMcpMarketItemsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMcpMarketItemsRequest) GoString() string {
	return s.String()
}

func (s *ListMcpMarketItemsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListMcpMarketItemsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMcpMarketItemsRequest) GetMcpType() *string {
	return s.McpType
}

func (s *ListMcpMarketItemsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMcpMarketItemsRequest) GetOfficialTag() *string {
	return s.OfficialTag
}

func (s *ListMcpMarketItemsRequest) SetKeyword(v string) *ListMcpMarketItemsRequest {
	s.Keyword = &v
	return s
}

func (s *ListMcpMarketItemsRequest) SetMaxResults(v int32) *ListMcpMarketItemsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListMcpMarketItemsRequest) SetMcpType(v string) *ListMcpMarketItemsRequest {
	s.McpType = &v
	return s
}

func (s *ListMcpMarketItemsRequest) SetNextToken(v string) *ListMcpMarketItemsRequest {
	s.NextToken = &v
	return s
}

func (s *ListMcpMarketItemsRequest) SetOfficialTag(v string) *ListMcpMarketItemsRequest {
	s.OfficialTag = &v
	return s
}

func (s *ListMcpMarketItemsRequest) Validate() error {
	return dara.Validate(s)
}
