// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryClientDiscountLogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ModelRouterQueryClientDiscountLogsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryClientDiscountLogsRequest
	GetNextToken() *string
}

type ModelRouterQueryClientDiscountLogsRequest struct {
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

func (s ModelRouterQueryClientDiscountLogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientDiscountLogsRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientDiscountLogsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryClientDiscountLogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryClientDiscountLogsRequest) SetMaxResults(v int32) *ModelRouterQueryClientDiscountLogsRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsRequest) SetNextToken(v string) *ModelRouterQueryClientDiscountLogsRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryClientDiscountLogsRequest) Validate() error {
	return dara.Validate(s)
}
