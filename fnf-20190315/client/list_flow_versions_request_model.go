// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowVersionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowName(v string) *ListFlowVersionsRequest
	GetFlowName() *string
	SetLimit(v string) *ListFlowVersionsRequest
	GetLimit() *string
	SetNextToken(v string) *ListFlowVersionsRequest
	GetNextToken() *string
}

type ListFlowVersionsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example-flow-name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// example:
	//
	// 10
	Limit *string `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// list token
	//
	// example:
	//
	// token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListFlowVersionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowVersionsRequest) GoString() string {
	return s.String()
}

func (s *ListFlowVersionsRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *ListFlowVersionsRequest) GetLimit() *string {
	return s.Limit
}

func (s *ListFlowVersionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFlowVersionsRequest) SetFlowName(v string) *ListFlowVersionsRequest {
	s.FlowName = &v
	return s
}

func (s *ListFlowVersionsRequest) SetLimit(v string) *ListFlowVersionsRequest {
	s.Limit = &v
	return s
}

func (s *ListFlowVersionsRequest) SetNextToken(v string) *ListFlowVersionsRequest {
	s.NextToken = &v
	return s
}

func (s *ListFlowVersionsRequest) Validate() error {
	return dara.Validate(s)
}
