// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicalModelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListLogicalModelsRequest
	GetGatewayId() *string
	SetKeyword(v string) *ListLogicalModelsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListLogicalModelsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListLogicalModelsRequest
	GetNextToken() *string
}

type ListLogicalModelsRequest struct {
	// The ID of the gateway instance. Specifies the gateway instance for which to query logical models.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-xxxxxxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The keyword used to filter logical models.
	//
	// example:
	//
	// chat
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of entries to return per page. Valid values: 1 to 100. This parameter is required. This operation uses token-based pagination with maxResults and nextToken instead of pageNumber and pageSize.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The nextToken value returned in the response of the previous page. Omit this parameter for the first page. The last page does not return a nextToken value.
	//
	// example:
	//
	// next-token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListLogicalModelsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLogicalModelsRequest) GoString() string {
	return s.String()
}

func (s *ListLogicalModelsRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListLogicalModelsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListLogicalModelsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListLogicalModelsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListLogicalModelsRequest) SetGatewayId(v string) *ListLogicalModelsRequest {
	s.GatewayId = &v
	return s
}

func (s *ListLogicalModelsRequest) SetKeyword(v string) *ListLogicalModelsRequest {
	s.Keyword = &v
	return s
}

func (s *ListLogicalModelsRequest) SetMaxResults(v int32) *ListLogicalModelsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListLogicalModelsRequest) SetNextToken(v string) *ListLogicalModelsRequest {
	s.NextToken = &v
	return s
}

func (s *ListLogicalModelsRequest) Validate() error {
	return dara.Validate(s)
}
