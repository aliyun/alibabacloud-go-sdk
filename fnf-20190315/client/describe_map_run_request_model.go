// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMapRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionName(v string) *DescribeMapRunRequest
	GetExecutionName() *string
	SetFlowName(v string) *DescribeMapRunRequest
	GetFlowName() *string
	SetMapRunName(v string) *DescribeMapRunRequest
	GetMapRunName() *string
	SetRequestId(v string) *DescribeMapRunRequest
	GetRequestId() *string
}

type DescribeMapRunRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// my_exec_name
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// my_flow_name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// c39142f1345b196d678333c41f113200
	MapRunName *string `json:"MapRunName,omitempty" xml:"MapRunName,omitempty"`
	// example:
	//
	// 3A44E113-9962-5B0B-AB92-14060EFE3164
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMapRunRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMapRunRequest) GoString() string {
	return s.String()
}

func (s *DescribeMapRunRequest) GetExecutionName() *string {
	return s.ExecutionName
}

func (s *DescribeMapRunRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *DescribeMapRunRequest) GetMapRunName() *string {
	return s.MapRunName
}

func (s *DescribeMapRunRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMapRunRequest) SetExecutionName(v string) *DescribeMapRunRequest {
	s.ExecutionName = &v
	return s
}

func (s *DescribeMapRunRequest) SetFlowName(v string) *DescribeMapRunRequest {
	s.FlowName = &v
	return s
}

func (s *DescribeMapRunRequest) SetMapRunName(v string) *DescribeMapRunRequest {
	s.MapRunName = &v
	return s
}

func (s *DescribeMapRunRequest) SetRequestId(v string) *DescribeMapRunRequest {
	s.RequestId = &v
	return s
}

func (s *DescribeMapRunRequest) Validate() error {
	return dara.Validate(s)
}
