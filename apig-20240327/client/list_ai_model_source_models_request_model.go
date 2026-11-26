// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAiModelSourceModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListAiModelSourceModelsRequest
	GetGatewayId() *string
	SetKeyword(v string) *ListAiModelSourceModelsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListAiModelSourceModelsRequest
	GetMaxResults() *int32
	SetModelSourceId(v string) *ListAiModelSourceModelsRequest
	GetModelSourceId() *string
	SetNextToken(v string) *ListAiModelSourceModelsRequest
	GetNextToken() *string
}

type ListAiModelSourceModelsRequest struct {
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
	// The maximum number of results to return per page. Valid values: 1 to 100. This parameter is required. Token-based pagination does not use pageNumber or pageSize.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The ID of the model source. Use the modelSourceId returned by CreateAiModelSource.
	//
	// example:
	//
	// ms-d00000000000000000000
	ModelSourceId *string `json:"modelSourceId,omitempty" xml:"modelSourceId,omitempty"`
	// The pagination token returned from the previous page. Omit this parameter for the first page. The token is empty or omitted for the last page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListAiModelSourceModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAiModelSourceModelsRequest) GoString() string {
	return s.String()
}

func (s *ListAiModelSourceModelsRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListAiModelSourceModelsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListAiModelSourceModelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListAiModelSourceModelsRequest) GetModelSourceId() *string {
	return s.ModelSourceId
}

func (s *ListAiModelSourceModelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAiModelSourceModelsRequest) SetGatewayId(v string) *ListAiModelSourceModelsRequest {
	s.GatewayId = &v
	return s
}

func (s *ListAiModelSourceModelsRequest) SetKeyword(v string) *ListAiModelSourceModelsRequest {
	s.Keyword = &v
	return s
}

func (s *ListAiModelSourceModelsRequest) SetMaxResults(v int32) *ListAiModelSourceModelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListAiModelSourceModelsRequest) SetModelSourceId(v string) *ListAiModelSourceModelsRequest {
	s.ModelSourceId = &v
	return s
}

func (s *ListAiModelSourceModelsRequest) SetNextToken(v string) *ListAiModelSourceModelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListAiModelSourceModelsRequest) Validate() error {
	return dara.Validate(s)
}
