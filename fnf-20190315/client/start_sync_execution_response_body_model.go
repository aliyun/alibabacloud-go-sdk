// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSyncExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnvironment(v *StartSyncExecutionResponseBodyEnvironment) *StartSyncExecutionResponseBody
	GetEnvironment() *StartSyncExecutionResponseBodyEnvironment
	SetErrorCode(v string) *StartSyncExecutionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *StartSyncExecutionResponseBody
	GetErrorMessage() *string
	SetFlowName(v string) *StartSyncExecutionResponseBody
	GetFlowName() *string
	SetName(v string) *StartSyncExecutionResponseBody
	GetName() *string
	SetOutput(v string) *StartSyncExecutionResponseBody
	GetOutput() *string
	SetRequestId(v string) *StartSyncExecutionResponseBody
	GetRequestId() *string
	SetStartedTime(v string) *StartSyncExecutionResponseBody
	GetStartedTime() *string
	SetStatus(v string) *StartSyncExecutionResponseBody
	GetStatus() *string
	SetStoppedTime(v string) *StartSyncExecutionResponseBody
	GetStoppedTime() *string
}

type StartSyncExecutionResponseBody struct {
	Environment *StartSyncExecutionResponseBodyEnvironment `json:"Environment,omitempty" xml:"Environment,omitempty" type:"Struct"`
	// The error code that is returned if the execution failed.
	//
	// example:
	//
	// ActionNotSupported
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message that indicates the execution timed out.
	//
	// example:
	//
	// Standard execution is not supported
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The name of the flow.
	//
	// example:
	//
	// my_flow_name
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	// The name of the execution.
	//
	// example:
	//
	// my_exec_name:{UUID}
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The output of the execution, which is in the JSON format.
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
	// The status of the execution. Valid values:
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

func (s StartSyncExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSyncExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartSyncExecutionResponseBody) GetEnvironment() *StartSyncExecutionResponseBodyEnvironment {
	return s.Environment
}

func (s *StartSyncExecutionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *StartSyncExecutionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StartSyncExecutionResponseBody) GetFlowName() *string {
	return s.FlowName
}

func (s *StartSyncExecutionResponseBody) GetName() *string {
	return s.Name
}

func (s *StartSyncExecutionResponseBody) GetOutput() *string {
	return s.Output
}

func (s *StartSyncExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSyncExecutionResponseBody) GetStartedTime() *string {
	return s.StartedTime
}

func (s *StartSyncExecutionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *StartSyncExecutionResponseBody) GetStoppedTime() *string {
	return s.StoppedTime
}

func (s *StartSyncExecutionResponseBody) SetEnvironment(v *StartSyncExecutionResponseBodyEnvironment) *StartSyncExecutionResponseBody {
	s.Environment = v
	return s
}

func (s *StartSyncExecutionResponseBody) SetErrorCode(v string) *StartSyncExecutionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetErrorMessage(v string) *StartSyncExecutionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetFlowName(v string) *StartSyncExecutionResponseBody {
	s.FlowName = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetName(v string) *StartSyncExecutionResponseBody {
	s.Name = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetOutput(v string) *StartSyncExecutionResponseBody {
	s.Output = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetRequestId(v string) *StartSyncExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetStartedTime(v string) *StartSyncExecutionResponseBody {
	s.StartedTime = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetStatus(v string) *StartSyncExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *StartSyncExecutionResponseBody) SetStoppedTime(v string) *StartSyncExecutionResponseBody {
	s.StoppedTime = &v
	return s
}

func (s *StartSyncExecutionResponseBody) Validate() error {
	if s.Environment != nil {
		if err := s.Environment.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartSyncExecutionResponseBodyEnvironment struct {
	Variables []*StartSyncExecutionResponseBodyEnvironmentVariables `json:"Variables,omitempty" xml:"Variables,omitempty" type:"Repeated"`
}

func (s StartSyncExecutionResponseBodyEnvironment) String() string {
	return dara.Prettify(s)
}

func (s StartSyncExecutionResponseBodyEnvironment) GoString() string {
	return s.String()
}

func (s *StartSyncExecutionResponseBodyEnvironment) GetVariables() []*StartSyncExecutionResponseBodyEnvironmentVariables {
	return s.Variables
}

func (s *StartSyncExecutionResponseBodyEnvironment) SetVariables(v []*StartSyncExecutionResponseBodyEnvironmentVariables) *StartSyncExecutionResponseBodyEnvironment {
	s.Variables = v
	return s
}

func (s *StartSyncExecutionResponseBodyEnvironment) Validate() error {
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

type StartSyncExecutionResponseBodyEnvironmentVariables struct {
	// example:
	//
	// key
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s StartSyncExecutionResponseBodyEnvironmentVariables) String() string {
	return dara.Prettify(s)
}

func (s StartSyncExecutionResponseBodyEnvironmentVariables) GoString() string {
	return s.String()
}

func (s *StartSyncExecutionResponseBodyEnvironmentVariables) GetName() *string {
	return s.Name
}

func (s *StartSyncExecutionResponseBodyEnvironmentVariables) GetValue() *string {
	return s.Value
}

func (s *StartSyncExecutionResponseBodyEnvironmentVariables) SetName(v string) *StartSyncExecutionResponseBodyEnvironmentVariables {
	s.Name = &v
	return s
}

func (s *StartSyncExecutionResponseBodyEnvironmentVariables) SetValue(v string) *StartSyncExecutionResponseBodyEnvironmentVariables {
	s.Value = &v
	return s
}

func (s *StartSyncExecutionResponseBodyEnvironmentVariables) Validate() error {
	return dara.Validate(s)
}
