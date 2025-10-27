// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListFlowsRequest
	GetLimit() *int32
	SetNextToken(v string) *ListFlowsRequest
	GetNextToken() *string
}

type ListFlowsRequest struct {
	// The number of workflows that you want to query. Valid values: 1 - 999. Default value: 60.
	//
	// example:
	//
	// 1
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The token to start the query.
	//
	// example:
	//
	// flow_nextxxx
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListFlowsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListFlowsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFlowsRequest) SetLimit(v int32) *ListFlowsRequest {
	s.Limit = &v
	return s
}

func (s *ListFlowsRequest) SetNextToken(v string) *ListFlowsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFlowsRequest) Validate() error {
	return dara.Validate(s)
}
