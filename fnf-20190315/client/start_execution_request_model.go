// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallbackFnFTaskToken(v string) *StartExecutionRequest
	GetCallbackFnFTaskToken() *string
	SetExecutionName(v string) *StartExecutionRequest
	GetExecutionName() *string
	SetFlowName(v string) *StartExecutionRequest
	GetFlowName() *string
	SetInput(v string) *StartExecutionRequest
	GetInput() *string
	SetQualifier(v string) *StartExecutionRequest
	GetQualifier() *string
}

type StartExecutionRequest struct {
	// Specifies that the **TaskToken**-related tasks are called back after the execution in the flow ends.
	//
	// example:
	//
	// 12
	CallbackFnFTaskToken *string `json:"CallbackFnFTaskToken,omitempty" xml:"CallbackFnFTaskToken,omitempty"`
	// The name of the execution. The execution name is unique within a workflow. Configure this parameter based on the following rules:
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// example:
	//
	// exec
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// The name of the workflow to be executed.
	//
	// This parameter is required.
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
	// example:
	//
	// 1
	Qualifier *string `json:"Qualifier,omitempty" xml:"Qualifier,omitempty"`
}

func (s StartExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartExecutionRequest) GetCallbackFnFTaskToken() *string {
	return s.CallbackFnFTaskToken
}

func (s *StartExecutionRequest) GetExecutionName() *string {
	return s.ExecutionName
}

func (s *StartExecutionRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *StartExecutionRequest) GetInput() *string {
	return s.Input
}

func (s *StartExecutionRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *StartExecutionRequest) SetCallbackFnFTaskToken(v string) *StartExecutionRequest {
	s.CallbackFnFTaskToken = &v
	return s
}

func (s *StartExecutionRequest) SetExecutionName(v string) *StartExecutionRequest {
	s.ExecutionName = &v
	return s
}

func (s *StartExecutionRequest) SetFlowName(v string) *StartExecutionRequest {
	s.FlowName = &v
	return s
}

func (s *StartExecutionRequest) SetInput(v string) *StartExecutionRequest {
	s.Input = &v
	return s
}

func (s *StartExecutionRequest) SetQualifier(v string) *StartExecutionRequest {
	s.Qualifier = &v
	return s
}

func (s *StartExecutionRequest) Validate() error {
	return dara.Validate(s)
}
