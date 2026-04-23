// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterBillingCostTabsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ModelRouterBillingCostTabsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterBillingCostTabsRequest
	GetNextToken() *string
}

type ModelRouterBillingCostTabsRequest struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextToken
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ModelRouterBillingCostTabsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterBillingCostTabsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterBillingCostTabsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterBillingCostTabsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterBillingCostTabsRequest) SetMaxResults(v int32) *ModelRouterBillingCostTabsRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterBillingCostTabsRequest) SetNextToken(v string) *ModelRouterBillingCostTabsRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterBillingCostTabsRequest) Validate() error {
	return dara.Validate(s)
}
