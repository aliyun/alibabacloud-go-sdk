// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSupportedPricingApisRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListSupportedPricingApisRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListSupportedPricingApisRequest
	GetNextToken() *string
}

type ListSupportedPricingApisRequest struct {
	MaxResults *int32  `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	NextToken  *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
}

func (s ListSupportedPricingApisRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSupportedPricingApisRequest) GoString() string {
	return s.String()
}

func (s *ListSupportedPricingApisRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSupportedPricingApisRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSupportedPricingApisRequest) SetMaxResults(v int32) *ListSupportedPricingApisRequest {
	s.MaxResults = &v
	return s
}

func (s *ListSupportedPricingApisRequest) SetNextToken(v string) *ListSupportedPricingApisRequest {
	s.NextToken = &v
	return s
}

func (s *ListSupportedPricingApisRequest) Validate() error {
	return dara.Validate(s)
}
