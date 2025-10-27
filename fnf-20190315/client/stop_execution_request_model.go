// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCause(v string) *StopExecutionRequest
	GetCause() *string
	SetError(v string) *StopExecutionRequest
	GetError() *string
	SetExecutionName(v string) *StopExecutionRequest
	GetExecutionName() *string
	SetFlowName(v string) *StopExecutionRequest
	GetFlowName() *string
}

type StopExecutionRequest struct {
	// The reason for stopping the execution. The value must be 1 to 4,096 characters in length.
	//
	// example:
	//
	// for test
	Cause *string `json:"Cause,omitempty" xml:"Cause,omitempty"`
	// The error code for stopping the execution. The error code must be 1 to 128 characters in length.
	//
	// example:
	//
	// nill
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// The name of the execution to be stopped. You can call the **ListExecutions*	- operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// The name of the workflow to be stopped. You can call the **ListFlows*	- operation to obtain the value of this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s StopExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StopExecutionRequest) GoString() string {
	return s.String()
}

func (s *StopExecutionRequest) GetCause() *string {
	return s.Cause
}

func (s *StopExecutionRequest) GetError() *string {
	return s.Error
}

func (s *StopExecutionRequest) GetExecutionName() *string {
	return s.ExecutionName
}

func (s *StopExecutionRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *StopExecutionRequest) SetCause(v string) *StopExecutionRequest {
	s.Cause = &v
	return s
}

func (s *StopExecutionRequest) SetError(v string) *StopExecutionRequest {
	s.Error = &v
	return s
}

func (s *StopExecutionRequest) SetExecutionName(v string) *StopExecutionRequest {
	s.ExecutionName = &v
	return s
}

func (s *StopExecutionRequest) SetFlowName(v string) *StopExecutionRequest {
	s.FlowName = &v
	return s
}

func (s *StopExecutionRequest) Validate() error {
	return dara.Validate(s)
}
