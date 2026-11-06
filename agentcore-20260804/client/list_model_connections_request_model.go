// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIncludeModels(v bool) *ListModelConnectionsRequest
	GetIncludeModels() *bool
	SetMaxResults(v int32) *ListModelConnectionsRequest
	GetMaxResults() *int32
	SetName(v string) *ListModelConnectionsRequest
	GetName() *string
	SetNextToken(v string) *ListModelConnectionsRequest
	GetNextToken() *string
	SetProtocol(v string) *ListModelConnectionsRequest
	GetProtocol() *string
	SetProviderType(v string) *ListModelConnectionsRequest
	GetProviderType() *string
	SetSearchType(v string) *ListModelConnectionsRequest
	GetSearchType() *string
}

type ListModelConnectionsRequest struct {
	// Specifies whether to return the summary of associated models in each model connection. Default value: false.
	//
	// example:
	//
	// false
	IncludeModels *bool `json:"includeModels,omitempty" xml:"includeModels,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 0 to 100. Default value: 10. If this parameter is not specified or is set to 0, the default value is used.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The name of the model connection. The name must be 1 to 128 characters in length and cannot contain whitespace characters.
	//
	// example:
	//
	// name
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The pagination token. Pass the token returned in the previous query for the current request. An empty response indicates that no more pages are available.
	//
	// example:
	//
	// bW9kZWwtbWFuYWdlbWVudC1vZmZzZXQ6bW9kZWwtY29ubmVjdGlvbjoxMA
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The invocation protocol used to filter model connections.
	//
	// example:
	//
	// OpenAI/v1
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// The model provider type used to filter model connections.
	//
	// example:
	//
	// qwen
	ProviderType *string `json:"providerType,omitempty" xml:"providerType,omitempty"`
	// The name matching mode. This parameter takes effect only when Name is specified. Valid values:
	//
	// - accurate: exact match
	//
	// - blur: fuzzy match
	//
	// Default value: blur.
	//
	// example:
	//
	// blur
	SearchType *string `json:"searchType,omitempty" xml:"searchType,omitempty"`
}

func (s ListModelConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelConnectionsRequest) GoString() string {
	return s.String()
}

func (s *ListModelConnectionsRequest) GetIncludeModels() *bool {
	return s.IncludeModels
}

func (s *ListModelConnectionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelConnectionsRequest) GetName() *string {
	return s.Name
}

func (s *ListModelConnectionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelConnectionsRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *ListModelConnectionsRequest) GetProviderType() *string {
	return s.ProviderType
}

func (s *ListModelConnectionsRequest) GetSearchType() *string {
	return s.SearchType
}

func (s *ListModelConnectionsRequest) SetIncludeModels(v bool) *ListModelConnectionsRequest {
	s.IncludeModels = &v
	return s
}

func (s *ListModelConnectionsRequest) SetMaxResults(v int32) *ListModelConnectionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelConnectionsRequest) SetName(v string) *ListModelConnectionsRequest {
	s.Name = &v
	return s
}

func (s *ListModelConnectionsRequest) SetNextToken(v string) *ListModelConnectionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelConnectionsRequest) SetProtocol(v string) *ListModelConnectionsRequest {
	s.Protocol = &v
	return s
}

func (s *ListModelConnectionsRequest) SetProviderType(v string) *ListModelConnectionsRequest {
	s.ProviderType = &v
	return s
}

func (s *ListModelConnectionsRequest) SetSearchType(v string) *ListModelConnectionsRequest {
	s.SearchType = &v
	return s
}

func (s *ListModelConnectionsRequest) Validate() error {
	return dara.Validate(s)
}
