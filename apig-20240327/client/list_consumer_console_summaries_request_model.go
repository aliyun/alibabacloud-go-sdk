// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerConsoleSummariesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v string) *ListConsumerConsoleSummariesRequest
	GetFilter() *string
	SetGatewayId(v string) *ListConsumerConsoleSummariesRequest
	GetGatewayId() *string
	SetMaxResults(v int32) *ListConsumerConsoleSummariesRequest
	GetMaxResults() *int32
	SetNameLike(v string) *ListConsumerConsoleSummariesRequest
	GetNameLike() *string
	SetNextToken(v string) *ListConsumerConsoleSummariesRequest
	GetNextToken() *string
}

type ListConsumerConsoleSummariesRequest struct {
	// The filter condition. Valid values: all, granted, and disabled. If you do not specify this parameter, all is used by default.
	//
	// example:
	//
	// all
	Filter *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The target AI gateway instance ID. Consumer identities are shared across the region, and authorizations are isolated by instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// gw-123456
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The maximum number of entries to return on each page. Valid values: 1 to 100. This parameter is required. Token-based pagination is used instead of pageNumber and pageSize.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The keyword for fuzzy search by consumer name. This parameter does not affect the total counts.
	//
	// example:
	//
	// beta
	NameLike *string `json:"nameLike,omitempty" xml:"nameLike,omitempty"`
	// The pagination token returned from the previous page. Omit this parameter when you query the first page. If the value is empty or omitted, it indicates the last page.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ==
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListConsumerConsoleSummariesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerConsoleSummariesRequest) GoString() string {
	return s.String()
}

func (s *ListConsumerConsoleSummariesRequest) GetFilter() *string {
	return s.Filter
}

func (s *ListConsumerConsoleSummariesRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListConsumerConsoleSummariesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListConsumerConsoleSummariesRequest) GetNameLike() *string {
	return s.NameLike
}

func (s *ListConsumerConsoleSummariesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListConsumerConsoleSummariesRequest) SetFilter(v string) *ListConsumerConsoleSummariesRequest {
	s.Filter = &v
	return s
}

func (s *ListConsumerConsoleSummariesRequest) SetGatewayId(v string) *ListConsumerConsoleSummariesRequest {
	s.GatewayId = &v
	return s
}

func (s *ListConsumerConsoleSummariesRequest) SetMaxResults(v int32) *ListConsumerConsoleSummariesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListConsumerConsoleSummariesRequest) SetNameLike(v string) *ListConsumerConsoleSummariesRequest {
	s.NameLike = &v
	return s
}

func (s *ListConsumerConsoleSummariesRequest) SetNextToken(v string) *ListConsumerConsoleSummariesRequest {
	s.NextToken = &v
	return s
}

func (s *ListConsumerConsoleSummariesRequest) Validate() error {
	return dara.Validate(s)
}
