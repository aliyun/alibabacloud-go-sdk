// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHiMarketConsumerSubscriptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListHiMarketConsumerSubscriptionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListHiMarketConsumerSubscriptionsRequest
	GetNextToken() *string
	SetPageNumber(v int32) *ListHiMarketConsumerSubscriptionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListHiMarketConsumerSubscriptionsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListHiMarketConsumerSubscriptionsRequest
	GetStatus() *string
}

type ListHiMarketConsumerSubscriptionsRequest struct {
	// The maximum number of entries per page. The default value and valid values are not specified in the API definition.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Pass the token returned in the previous response to retrieve the next page. Do not specify this parameter for the first request.
	//
	// example:
	//
	// token-xxxxxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The number of the page to return. Pages start from 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The subscription status, which is used to filter subscriptions. Valid values: PENDING and APPROVED.
	//
	// example:
	//
	// APPROVED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ListHiMarketConsumerSubscriptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListHiMarketConsumerSubscriptionsRequest) GoString() string {
	return s.String()
}

func (s *ListHiMarketConsumerSubscriptionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListHiMarketConsumerSubscriptionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListHiMarketConsumerSubscriptionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListHiMarketConsumerSubscriptionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListHiMarketConsumerSubscriptionsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListHiMarketConsumerSubscriptionsRequest) SetMaxResults(v int32) *ListHiMarketConsumerSubscriptionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsRequest) SetNextToken(v string) *ListHiMarketConsumerSubscriptionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsRequest) SetPageNumber(v int32) *ListHiMarketConsumerSubscriptionsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsRequest) SetPageSize(v int32) *ListHiMarketConsumerSubscriptionsRequest {
	s.PageSize = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsRequest) SetStatus(v string) *ListHiMarketConsumerSubscriptionsRequest {
	s.Status = &v
	return s
}

func (s *ListHiMarketConsumerSubscriptionsRequest) Validate() error {
	return dara.Validate(s)
}
