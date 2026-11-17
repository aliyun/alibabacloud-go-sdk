// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListModelGroupsRequest
	GetGatewayId() *string
	SetKeyword(v string) *ListModelGroupsRequest
	GetKeyword() *string
	SetMaxResults(v int32) *ListModelGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListModelGroupsRequest
	GetNextToken() *string
}

type ListModelGroupsRequest struct {
	// The ID of the AI gateway instance to which the resource belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-d00000000000000000000
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The keyword used to match model groups by name, description, member identity, or invocation name.
	//
	// example:
	//
	// beta
	Keyword *string `json:"keyword,omitempty" xml:"keyword,omitempty"`
	// The maximum number of results to return per page. Valid values: 1 to 100. This parameter is required. Token-based pagination is used instead of pageNumber and pageSize.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token returned from the previous page. Omit this parameter for the first page. The value is empty or omitted for the last page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListModelGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListModelGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListModelGroupsRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListModelGroupsRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListModelGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelGroupsRequest) SetGatewayId(v string) *ListModelGroupsRequest {
	s.GatewayId = &v
	return s
}

func (s *ListModelGroupsRequest) SetKeyword(v string) *ListModelGroupsRequest {
	s.Keyword = &v
	return s
}

func (s *ListModelGroupsRequest) SetMaxResults(v int32) *ListModelGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListModelGroupsRequest) SetNextToken(v string) *ListModelGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListModelGroupsRequest) Validate() error {
	return dara.Validate(s)
}
