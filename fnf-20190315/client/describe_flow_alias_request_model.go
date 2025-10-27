// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowAliasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowName(v string) *DescribeFlowAliasRequest
	GetFlowName() *string
	SetName(v string) *DescribeFlowAliasRequest
	GetName() *string
}

type DescribeFlowAliasRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// example-flow-name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example-alias-name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFlowAliasRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowAliasRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowAliasRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *DescribeFlowAliasRequest) GetName() *string {
	return s.Name
}

func (s *DescribeFlowAliasRequest) SetFlowName(v string) *DescribeFlowAliasRequest {
	s.FlowName = &v
	return s
}

func (s *DescribeFlowAliasRequest) SetName(v string) *DescribeFlowAliasRequest {
	s.Name = &v
	return s
}

func (s *DescribeFlowAliasRequest) Validate() error {
	return dara.Validate(s)
}
