// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryClientTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ModelRouterQueryClientTreeRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryClientTreeRequest
	GetNextToken() *string
}

type ModelRouterQueryClientTreeRequest struct {
	// maxResults
	//
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

func (s ModelRouterQueryClientTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientTreeRequest) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientTreeRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryClientTreeRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryClientTreeRequest) SetMaxResults(v int32) *ModelRouterQueryClientTreeRequest {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryClientTreeRequest) SetNextToken(v string) *ModelRouterQueryClientTreeRequest {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryClientTreeRequest) Validate() error {
	return dara.Validate(s)
}
