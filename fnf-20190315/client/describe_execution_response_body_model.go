// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v *DescribeExecutionResponseBodyEnvironment) *DescribeExecutionResponseBody
	GetEnvironment() *DescribeExecutionResponseBodyEnvironment
	SetFlowDefinition(v string) *DescribeExecutionResponseBody
	GetFlowDefinition() *string
	SetFlowName(v string) *DescribeExecutionResponseBody
	GetFlowName() *string
	SetInput(v string) *DescribeExecutionResponseBody
	GetInput() *string
	SetName(v string) *DescribeExecutionResponseBody
	GetName() *string
	SetOutput(v string) *DescribeExecutionResponseBody
	GetOutput() *string
	SetRequestId(v string) *DescribeExecutionResponseBody
	GetRequestId() *string
	SetStartedTime(v string) *DescribeExecutionResponseBody
	GetStartedTime() *string
	SetStatus(v string) *DescribeExecutionResponseBody
	GetStatus() *string
	SetStoppedTime(v string) *DescribeExecutionResponseBody
	GetStoppedTime() *string
}

type DescribeExecutionResponseBody struct {
	Environment *DescribeExecutionResponseBodyEnvironment `json:"Environment,omitempty" xml:"Environment,omitempty" type:"Struct"`
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

func (s DescribeExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExecutionResponseBody) GetEnvironment() *DescribeExecutionResponseBodyEnvironment {
	return s.Environment
}

func (s *DescribeExecutionResponseBody) GetFlowDefinition() *string {
	return s.FlowDefinition
}

func (s *DescribeExecutionResponseBody) GetFlowName() *string {
	return s.FlowName
}

func (s *DescribeExecutionResponseBody) GetInput() *string {
	return s.Input
}

func (s *DescribeExecutionResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeExecutionResponseBody) GetOutput() *string {
	return s.Output
}

func (s *DescribeExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExecutionResponseBody) GetStartedTime() *string {
	return s.StartedTime
}

func (s *DescribeExecutionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeExecutionResponseBody) GetStoppedTime() *string {
	return s.StoppedTime
}

func (s *DescribeExecutionResponseBody) SetEnvironment(v *DescribeExecutionResponseBodyEnvironment) *DescribeExecutionResponseBody {
	s.Environment = v
	return s
}

func (s *DescribeExecutionResponseBody) SetFlowDefinition(v string) *DescribeExecutionResponseBody {
	s.FlowDefinition = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetFlowName(v string) *DescribeExecutionResponseBody {
	s.FlowName = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetInput(v string) *DescribeExecutionResponseBody {
	s.Input = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetName(v string) *DescribeExecutionResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetOutput(v string) *DescribeExecutionResponseBody {
	s.Output = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetRequestId(v string) *DescribeExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetStartedTime(v string) *DescribeExecutionResponseBody {
	s.StartedTime = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetStatus(v string) *DescribeExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeExecutionResponseBody) SetStoppedTime(v string) *DescribeExecutionResponseBody {
	s.StoppedTime = &v
	return s
}

func (s *DescribeExecutionResponseBody) Validate() error {
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeExecutionResponseBodyEnvironment struct {
	Variables []*DescribeExecutionResponseBodyEnvironmentVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s DescribeExecutionResponseBodyEnvironment) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutionResponseBodyEnvironment) GoString() string {
	return s.String()
}

func (s *DescribeExecutionResponseBodyEnvironment) GetVariables() []*DescribeExecutionResponseBodyEnvironmentVariables {
	return s.Variables
}

func (s *DescribeExecutionResponseBodyEnvironment) SetVariables(v []*DescribeExecutionResponseBodyEnvironmentVariables) *DescribeExecutionResponseBodyEnvironment {
	s.Variables = v
	return s
}

func (s *DescribeExecutionResponseBodyEnvironment) Validate() error {
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeExecutionResponseBodyEnvironmentVariables struct {
	// example:
	//
	// key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeExecutionResponseBodyEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s DescribeExecutionResponseBodyEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *DescribeExecutionResponseBodyEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *DescribeExecutionResponseBodyEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *DescribeExecutionResponseBodyEnvironmentVariables) SetName(v string) *DescribeExecutionResponseBodyEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *DescribeExecutionResponseBodyEnvironmentVariables) SetValue(v string) *DescribeExecutionResponseBodyEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *DescribeExecutionResponseBodyEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}
