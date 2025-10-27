// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionName(v string) *DescribeExecutionRequest
	GetExecutionName() *string
	SetFlowName(v string) *DescribeExecutionRequest
	GetFlowName() *string
	SetWaitTimeSeconds(v int32) *DescribeExecutionRequest
	GetWaitTimeSeconds() *int32
}

type DescribeExecutionRequest struct {
	// The name of the execution.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// The name of the workflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The maximum period of time for long polling waits. Valid values: 0 to 60. Unit: seconds. Configure this parameter based on the following rules:
	//
	// 	- If the value is 0, the system immediately returns the current execution status.
	//
	// 	- If the value is greater than 0, the long polling request waits until the execution ends or the specified period elapses.
	//
	// example:
	//
	// 20
	WaitTimeSeconds *int32 `json:"WaitTimeSeconds,omitempty" xml:"WaitTimeSeconds,omitempty"`
}

func (s DescribeExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutionRequest) GoString() string {
	return s.String()
}

func (s *DescribeExecutionRequest) GetExecutionName() *string {
	return s.ExecutionName
}

func (s *DescribeExecutionRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *DescribeExecutionRequest) GetWaitTimeSeconds() *int32 {
	return s.WaitTimeSeconds
}

func (s *DescribeExecutionRequest) SetExecutionName(v string) *DescribeExecutionRequest {
	s.ExecutionName = &v
	return s
}

func (s *DescribeExecutionRequest) SetFlowName(v string) *DescribeExecutionRequest {
	s.FlowName = &v
	return s
}

func (s *DescribeExecutionRequest) SetWaitTimeSeconds(v int32) *DescribeExecutionRequest {
	s.WaitTimeSeconds = &v
	return s
}

func (s *DescribeExecutionRequest) Validate() error {
	return dara.Validate(s)
}
