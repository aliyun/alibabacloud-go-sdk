// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecutions(v []*ListExecutionsResponseBodyExecutions) *ListExecutionsResponseBody
	GetExecutions() []*ListExecutionsResponseBodyExecutions
	SetNextToken(v string) *ListExecutionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListExecutionsResponseBody
	GetRequestId() *string
}

type ListExecutionsResponseBody struct {
	// The information about executions.
	Executions []*ListExecutionsResponseBodyExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	// The start key for the next query. This parameter is not returned if this is the last query.
	//
	// >  This parameter may not be displayed in the response because no next page exists.
	//
	// example:
	//
	// exec2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// testRequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBody) GetExecutions() []*ListExecutionsResponseBodyExecutions {
	return s.Executions
}

func (s *ListExecutionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExecutionsResponseBody) SetExecutions(v []*ListExecutionsResponseBodyExecutions) *ListExecutionsResponseBody {
	s.Executions = v
	return s
}

func (s *ListExecutionsResponseBody) SetNextToken(v string) *ListExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListExecutionsResponseBody) SetRequestId(v string) *ListExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExecutionsResponseBody) Validate() error {
	if s.Executions != nil {
		for _, item := range s.Executions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExecutionsResponseBodyExecutions struct {
	Environment *ListExecutionsResponseBodyExecutionsEnvironment `json:"Environment,omitempty" xml:"Environment,omitempty" type:"Struct"`
	// The definition of the flow.
	//
	// example:
	//
	// version: v1.0\\ntype: flow\\nname: test\\nsteps:\\n  - type: pass\\n    name: mypass
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
	// The output of the execution, which is in the JSON format
	//
	// example:
	//
	// {"key":"value"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The time when the execution started.
	//
	// example:
	//
	// 2019-01-01T01:01:01.001Z
	StartedTime *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// The status of the execution.
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

func (s ListExecutionsResponseBodyExecutions) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionsResponseBodyExecutions) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBodyExecutions) GetEnvironment() *ListExecutionsResponseBodyExecutionsEnvironment {
	return s.Environment
}

func (s *ListExecutionsResponseBodyExecutions) GetFlowDefinition() *string {
	return s.FlowDefinition
}

func (s *ListExecutionsResponseBodyExecutions) GetFlowName() *string {
	return s.FlowName
}

func (s *ListExecutionsResponseBodyExecutions) GetInput() *string {
	return s.Input
}

func (s *ListExecutionsResponseBodyExecutions) GetName() *string {
	return s.Name
}

func (s *ListExecutionsResponseBodyExecutions) GetOutput() *string {
	return s.Output
}

func (s *ListExecutionsResponseBodyExecutions) GetStartedTime() *string {
	return s.StartedTime
}

func (s *ListExecutionsResponseBodyExecutions) GetStatus() *string {
	return s.Status
}

func (s *ListExecutionsResponseBodyExecutions) GetStoppedTime() *string {
	return s.StoppedTime
}

func (s *ListExecutionsResponseBodyExecutions) SetEnvironment(v *ListExecutionsResponseBodyExecutionsEnvironment) *ListExecutionsResponseBodyExecutions {
	s.Environment = v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetFlowDefinition(v string) *ListExecutionsResponseBodyExecutions {
	s.FlowDefinition = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetFlowName(v string) *ListExecutionsResponseBodyExecutions {
	s.FlowName = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetInput(v string) *ListExecutionsResponseBodyExecutions {
	s.Input = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetName(v string) *ListExecutionsResponseBodyExecutions {
	s.Name = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetOutput(v string) *ListExecutionsResponseBodyExecutions {
	s.Output = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStartedTime(v string) *ListExecutionsResponseBodyExecutions {
	s.StartedTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStatus(v string) *ListExecutionsResponseBodyExecutions {
	s.Status = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) SetStoppedTime(v string) *ListExecutionsResponseBodyExecutions {
	s.StoppedTime = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutions) Validate() error {
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExecutionsResponseBodyExecutionsEnvironment struct {
	Variables []*ListExecutionsResponseBodyExecutionsEnvironmentVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s ListExecutionsResponseBodyExecutionsEnvironment) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionsResponseBodyExecutionsEnvironment) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBodyExecutionsEnvironment) GetVariables() []*ListExecutionsResponseBodyExecutionsEnvironmentVariables {
	return s.Variables
}

func (s *ListExecutionsResponseBodyExecutionsEnvironment) SetVariables(v []*ListExecutionsResponseBodyExecutionsEnvironmentVariables) *ListExecutionsResponseBodyExecutionsEnvironment {
	s.Variables = v
	return s
}

func (s *ListExecutionsResponseBodyExecutionsEnvironment) Validate() error {
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

type ListExecutionsResponseBodyExecutionsEnvironmentVariables struct {
	// example:
	//
	// key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListExecutionsResponseBodyExecutionsEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s ListExecutionsResponseBodyExecutionsEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *ListExecutionsResponseBodyExecutionsEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *ListExecutionsResponseBodyExecutionsEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *ListExecutionsResponseBodyExecutionsEnvironmentVariables) SetName(v string) *ListExecutionsResponseBodyExecutionsEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutionsEnvironmentVariables) SetValue(v string) *ListExecutionsResponseBodyExecutionsEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *ListExecutionsResponseBodyExecutionsEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}
