// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFlowVersion(v string) *DescribeFlowRequest
	GetFlowVersion() *string
	SetName(v string) *DescribeFlowRequest
	GetName() *string
}

type DescribeFlowRequest struct {
	// example:
	//
	// 1
	FlowVersion *string `json:"FlowVersion,omitempty" xml:"FlowVersion,omitempty"`
	// The name of the flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowRequest) GoString() string {
	return s.String()
}

func (s *DescribeFlowRequest) GetFlowVersion() *string {
	return s.FlowVersion
}

func (s *DescribeFlowRequest) GetName() *string {
	return s.Name
}

func (s *DescribeFlowRequest) SetFlowVersion(v string) *DescribeFlowRequest {
	s.FlowVersion = &v
	return s
}

func (s *DescribeFlowRequest) SetName(v string) *DescribeFlowRequest {
	s.Name = &v
	return s
}

func (s *DescribeFlowRequest) Validate() error {
	return dara.Validate(s)
}
