// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSyncExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionName(v string) *StartSyncExecutionRequest
	GetExecutionName() *string
	SetFlowName(v string) *StartSyncExecutionRequest
	GetFlowName() *string
	SetInput(v string) *StartSyncExecutionRequest
	GetInput() *string
	SetQualifier(v string) *StartSyncExecutionRequest
	GetQualifier() *string
}

type StartSyncExecutionRequest struct {
	// The name of the execution that you want to start. The name must meet the following conventions:
	//
	// 	- The name can contain letters, digits, underscores (_), and hyphens (-).
	//
	// 	- The name must start with a letter or an underscore (_).
	//
	// 	- The name is case-sensitive.
	//
	// 	- The name must be 1 to 128 characters in length.
	//
	// Different from the StartExecution operation, in the synchronous execution mode, the execution name is no longer required to be unique within a flow. You can choose to provide an execution name to identify the current execution. In this case, the system adds a UUID to the current execution name. The used format is {ExecutionName}:{UUID}. If you do not specify the execution name, the system automatically generates an execution name.
	//
	// example:
	//
	// my_exec_name
	ExecutionName *string `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	// The name of the workflow to be executed.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_flow_name
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

func (s StartSyncExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartSyncExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartSyncExecutionRequest) GetExecutionName() *string {
	return s.ExecutionName
}

func (s *StartSyncExecutionRequest) GetFlowName() *string {
	return s.FlowName
}

func (s *StartSyncExecutionRequest) GetInput() *string {
	return s.Input
}

func (s *StartSyncExecutionRequest) GetQualifier() *string {
	return s.Qualifier
}

func (s *StartSyncExecutionRequest) SetExecutionName(v string) *StartSyncExecutionRequest {
	s.ExecutionName = &v
	return s
}

func (s *StartSyncExecutionRequest) SetFlowName(v string) *StartSyncExecutionRequest {
	s.FlowName = &v
	return s
}

func (s *StartSyncExecutionRequest) SetInput(v string) *StartSyncExecutionRequest {
	s.Input = &v
	return s
}

func (s *StartSyncExecutionRequest) SetQualifier(v string) *StartSyncExecutionRequest {
	s.Qualifier = &v
	return s
}

func (s *StartSyncExecutionRequest) Validate() error {
	return dara.Validate(s)
}
