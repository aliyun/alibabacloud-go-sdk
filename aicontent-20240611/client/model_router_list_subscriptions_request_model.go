// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterListSubscriptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBalanceType(v string) *ModelRouterListSubscriptionsRequest
	GetBalanceType() *string
	SetMaxResults(v int32) *ModelRouterListSubscriptionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterListSubscriptionsRequest
	GetNextToken() *string
	SetStatus(v string) *ModelRouterListSubscriptionsRequest
	GetStatus() *string
}

type ModelRouterListSubscriptionsRequest struct {
	// Filters by balance type (permanent/monthly).
	//
	// example:
	//
	// permanent
	BalanceType *string `json:"balanceType,omitempty" xml:"balanceType,omitempty"`
	// The maximum number of results to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token. Do not specify this parameter for the first query. For subsequent queries, specify the value returned from the previous query. Set to "" when no more data is available. Set to "5" when there is a next page.
	//
	// example:
	//
	// "5" or ""
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Filters by status (active/stopped).
	//
	// example:
	//
	// active
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ModelRouterListSubscriptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterListSubscriptionsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterListSubscriptionsRequest) GetBalanceType() *string {
	return s.BalanceType
}

func (s *ModelRouterListSubscriptionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterListSubscriptionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterListSubscriptionsRequest) GetStatus() *string {
	return s.Status
}

func (s *ModelRouterListSubscriptionsRequest) SetBalanceType(v string) *ModelRouterListSubscriptionsRequest {
	s.BalanceType = &v
	return s
}

func (s *ModelRouterListSubscriptionsRequest) SetMaxResults(v int32) *ModelRouterListSubscriptionsRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterListSubscriptionsRequest) SetNextToken(v string) *ModelRouterListSubscriptionsRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterListSubscriptionsRequest) SetStatus(v string) *ModelRouterListSubscriptionsRequest {
	s.Status = &v
	return s
}

func (s *ModelRouterListSubscriptionsRequest) Validate() error {
	return dara.Validate(s)
}
