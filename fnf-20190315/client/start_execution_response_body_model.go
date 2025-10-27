// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlowDefinition(v string) *StartExecutionResponseBody
	GetFlowDefinition() *string
	SetFlowName(v string) *StartExecutionResponseBody
	GetFlowName() *string
	SetInput(v string) *StartExecutionResponseBody
	GetInput() *string
	SetName(v string) *StartExecutionResponseBody
	GetName() *string
	SetOutput(v string) *StartExecutionResponseBody
	GetOutput() *string
	SetRequestId(v string) *StartExecutionResponseBody
	GetRequestId() *string
	SetStartedTime(v string) *StartExecutionResponseBody
	GetStartedTime() *string
	SetStatus(v string) *StartExecutionResponseBody
	GetStatus() *string
	SetStoppedTime(v string) *StartExecutionResponseBody
	GetStoppedTime() *string
}

type StartExecutionResponseBody struct {
	// The definition of the flow.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n - type: pass\\n name: mypass
	FlowDefinition *string `json:"FlowDefinition,omitempty" xml:"FlowDefinition,omitempty"`
	// The name of the workflow.
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
	// exec1
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
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The time when the execution stopped.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StoppedTime *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
}

func (s StartExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartExecutionResponseBody) GetFlowDefinition() *string {
	return s.FlowDefinition
}

func (s *StartExecutionResponseBody) GetFlowName() *string {
	return s.FlowName
}

func (s *StartExecutionResponseBody) GetInput() *string {
	return s.Input
}

func (s *StartExecutionResponseBody) GetName() *string {
	return s.Name
}

func (s *StartExecutionResponseBody) GetOutput() *string {
	return s.Output
}

func (s *StartExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartExecutionResponseBody) GetStartedTime() *string {
	return s.StartedTime
}

func (s *StartExecutionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *StartExecutionResponseBody) GetStoppedTime() *string {
	return s.StoppedTime
}

func (s *StartExecutionResponseBody) SetFlowDefinition(v string) *StartExecutionResponseBody {
	s.FlowDefinition = &v
	return s
}

func (s *StartExecutionResponseBody) SetFlowName(v string) *StartExecutionResponseBody {
	s.FlowName = &v
	return s
}

func (s *StartExecutionResponseBody) SetInput(v string) *StartExecutionResponseBody {
	s.Input = &v
	return s
}

func (s *StartExecutionResponseBody) SetName(v string) *StartExecutionResponseBody {
	s.Name = &v
	return s
}

func (s *StartExecutionResponseBody) SetOutput(v string) *StartExecutionResponseBody {
	s.Output = &v
	return s
}

func (s *StartExecutionResponseBody) SetRequestId(v string) *StartExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartExecutionResponseBody) SetStartedTime(v string) *StartExecutionResponseBody {
	s.StartedTime = &v
	return s
}

func (s *StartExecutionResponseBody) SetStatus(v string) *StartExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *StartExecutionResponseBody) SetStoppedTime(v string) *StartExecutionResponseBody {
	s.StoppedTime = &v
	return s
}

func (s *StartExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
