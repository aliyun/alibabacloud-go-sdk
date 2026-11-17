// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiModelSourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListAiModelSourcesRequest
	GetGatewayId() *string
	SetKeyword(v string) *ListAiModelSourcesRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListAiModelSourcesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListAiModelSourcesRequest
	GetNextToken() *string
	SetProvider(v string) *ListAiModelSourcesRequest
	GetProvider() *string
}

type ListAiModelSourcesRequest struct {
	// The ID of the AI gateway instance to which the resource belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-d00000000000000000000
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The search keyword used to match resources by name or identifier.
	//
	// example:
	//
	// beta
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of results to return per page. Valid values: 1 to 100. This parameter is required. Do not use pageNumber or pageSize for token-based pagination.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token returned from the previous page. Omit this parameter for the first page. The token is empty or omitted on the last page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The stable identifier of the model provider.
	//
	// example:
	//
	// qwen
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s ListAiModelSourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListAiModelSourcesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListAiModelSourcesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAiModelSourcesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAiModelSourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAiModelSourcesRequest) GetProvider() *string {
	return s.Provider
}

func (s *ListAiModelSourcesRequest) SetGatewayId(v string) *ListAiModelSourcesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListAiModelSourcesRequest) SetKeyword(v string) *ListAiModelSourcesRequest {
	s.Keyword = &v
	return s
}

func (s *ListAiModelSourcesRequest) SetMaxResults(v int32) *ListAiModelSourcesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAiModelSourcesRequest) SetNextToken(v string) *ListAiModelSourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListAiModelSourcesRequest) SetProvider(v string) *ListAiModelSourcesRequest {
	s.Provider = &v
	return s
}

func (s *ListAiModelSourcesRequest) Validate() error {
	return dara.Validate(s)
}
