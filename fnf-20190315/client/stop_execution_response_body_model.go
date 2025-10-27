// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowDefinition(v string) *StopExecutionResponseBody
	GetFlowDefinition() *string
	SetFlowName(v string) *StopExecutionResponseBody
	GetFlowName() *string
	SetInput(v string) *StopExecutionResponseBody
	GetInput() *string
	SetName(v string) *StopExecutionResponseBody
	GetName() *string
	SetOutput(v string) *StopExecutionResponseBody
	GetOutput() *string
	SetRequestId(v string) *StopExecutionResponseBody
	GetRequestId() *string
	SetRoleArn(v string) *StopExecutionResponseBody
	GetRoleArn() *string
	SetStartedTime(v string) *StopExecutionResponseBody
	GetStartedTime() *string
	SetStatus(v string) *StopExecutionResponseBody
	GetStatus() *string
	SetStoppedTime(v string) *StopExecutionResponseBody
	GetStoppedTime() *string
}

type StopExecutionResponseBody struct {
	// The definition of the flow.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n - type: pass\\n name: mypass
	FlowDefinition *string `json:"FlowDefinition,omitempty" xml:"FlowDefinition,omitempty"`
	// The name of the flow.
	//
	// example:
	//
	// flow
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The input of the execution, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Input *string `json:"Input,omitempty" xml:"Input,omitempty"`
	// The name of the execution.
	//
	// example:
	//
	// exec
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The execution result, which is in the JSON format.
	//
	// example:
	//
	// {"key":"value"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud resource name (ARN) of the role that executed the flow. If the RoleArn in the flow definition is changed during the execution of the flow, the system records and returns a snapshot of the original RoleArn.
	//
	// >  If you do not specify the RoleArn parameter in the request parameters, the response parameters do not contain the RoleArn parameter.
	//
	// example:
	//
	// acs:ram:${region}:${accountID}:${role}
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// The time when the execution started.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StartedTime *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The execution status. Valid values:
	//
	// 	- **Starting**
	//
	// 	- **Running**
	//
	// 	- **Stopped**
	//
	// 	- **Succeeded**
	//
	// 	- **Failed**
	//
	// 	- **TimedOut**
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the execution stopped.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StoppedTime *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
}

func (s StopExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StopExecutionResponseBody) GetFlowDefinition() *string {
	return s.FlowDefinition
}

func (s *StopExecutionResponseBody) GetFlowName() *string {
	return s.FlowName
}

func (s *StopExecutionResponseBody) GetInput() *string {
	return s.Input
}

func (s *StopExecutionResponseBody) GetName() *string {
	return s.Name
}

func (s *StopExecutionResponseBody) GetOutput() *string {
	return s.Output
}

func (s *StopExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopExecutionResponseBody) GetRoleArn() *string {
	return s.RoleArn
}

func (s *StopExecutionResponseBody) GetStartedTime() *string {
	return s.StartedTime
}

func (s *StopExecutionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *StopExecutionResponseBody) GetStoppedTime() *string {
	return s.StoppedTime
}

func (s *StopExecutionResponseBody) SetFlowDefinition(v string) *StopExecutionResponseBody {
	s.FlowDefinition = &v
	return s
}

func (s *StopExecutionResponseBody) SetFlowName(v string) *StopExecutionResponseBody {
	s.FlowName = &v
	return s
}

func (s *StopExecutionResponseBody) SetInput(v string) *StopExecutionResponseBody {
	s.Input = &v
	return s
}

func (s *StopExecutionResponseBody) SetName(v string) *StopExecutionResponseBody {
	s.Name = &v
	return s
}

func (s *StopExecutionResponseBody) SetOutput(v string) *StopExecutionResponseBody {
	s.Output = &v
	return s
}

func (s *StopExecutionResponseBody) SetRequestId(v string) *StopExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopExecutionResponseBody) SetRoleArn(v string) *StopExecutionResponseBody {
	s.RoleArn = &v
	return s
}

func (s *StopExecutionResponseBody) SetStartedTime(v string) *StopExecutionResponseBody {
	s.StartedTime = &v
	return s
}

func (s *StopExecutionResponseBody) SetStatus(v string) *StopExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *StopExecutionResponseBody) SetStoppedTime(v string) *StopExecutionResponseBody {
	s.StoppedTime = &v
	return s
}

func (s *StopExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
