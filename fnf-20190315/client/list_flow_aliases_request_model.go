// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowAliasesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowName(v string) *ListFlowAliasesRequest
	GetFlowName() *string
	SetLimit(v int32) *ListFlowAliasesRequest
	GetLimit() *int32
	SetNextToken(v string) *ListFlowAliasesRequest
	GetNextToken() *string
}

type ListFlowAliasesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example-flow-name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// list token
	//
	// example:
	//
	// token
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListFlowAliasesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFlowAliasesRequest) GoString() string {
	return s.String()
}

func (s *ListFlowAliasesRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *ListFlowAliasesRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListFlowAliasesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFlowAliasesRequest) SetFlowName(v string) *ListFlowAliasesRequest {
	s.FlowName = &v
	return s
}

func (s *ListFlowAliasesRequest) SetLimit(v int32) *ListFlowAliasesRequest {
	s.Limit = &v
	return s
}

func (s *ListFlowAliasesRequest) SetNextToken(v string) *ListFlowAliasesRequest {
	s.NextToken = &v
	return s
}

func (s *ListFlowAliasesRequest) Validate() error {
	return dara.Validate(s)
}
