// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitoringAgentProcessesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeMonitoringAgentProcessesResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeMonitoringAgentProcessesResponseBody
	GetMessage() *string
	SetNodeProcesses(v *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses) *DescribeMonitoringAgentProcessesResponseBody
	GetNodeProcesses() *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses
	SetRequestId(v string) *DescribeMonitoringAgentProcessesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeMonitoringAgentProcessesResponseBody
	GetSuccess() *bool
}

type DescribeMonitoringAgentProcessesResponseBody struct {
	// The HTTP status code.
	//
	// >  The value 200 indicates that the call was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The information about the processes.
	NodeProcesses *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses `json:"NodeProcesses,omitempty" xml:"NodeProcesses,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C11C0E85-6862-4F25-8D66-D6A5E0882984
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// 	- true: The call was successful.
	//
	// 	- false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeMonitoringAgentProcessesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentProcessesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentProcessesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeMonitoringAgentProcessesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeMonitoringAgentProcessesResponseBody) GetNodeProcesses() *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses {
	return s.NodeProcesses
}

func (s *DescribeMonitoringAgentProcessesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMonitoringAgentProcessesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeMonitoringAgentProcessesResponseBody) SetCode(v string) *DescribeMonitoringAgentProcessesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBody) SetMessage(v string) *DescribeMonitoringAgentProcessesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBody) SetNodeProcesses(v *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses) *DescribeMonitoringAgentProcessesResponseBody {
	s.NodeProcesses = v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBody) SetRequestId(v string) *DescribeMonitoringAgentProcessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBody) SetSuccess(v bool) *DescribeMonitoringAgentProcessesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBody) Validate() error {
	if s.NodeProcesses != nil {
		if err := s.NodeProcesses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMonitoringAgentProcessesResponseBodyNodeProcesses struct {
	NodeProcess []*DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess `json:"NodeProcess,omitempty" xml:"NodeProcess,omitempty" type:"Repeated"`
}

func (s DescribeMonitoringAgentProcessesResponseBodyNodeProcesses) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentProcessesResponseBodyNodeProcesses) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses) GetNodeProcess() []*DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	return s.NodeProcess
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses) SetNodeProcess(v []*DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses {
	s.NodeProcess = v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcesses) Validate() error {
	if s.NodeProcess != nil {
		for _, item := range s.NodeProcess {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess struct {
	// The command used to obtain the number of processes. Valid value: `number`.
	//
	// >  The `number` command obtains the number of processes that match the condition.
	//
	// example:
	//
	// number
	Command *string `json:"Command,omitempty" xml:"Command,omitempty"`
	// The ID of the application group.
	//
	// example:
	//
	// 3619****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-hp3hl3cx1pbahzy8****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the process.
	//
	// example:
	//
	// 234567
	ProcessId *int64 `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// The name of the process.
	//
	// example:
	//
	// Nginx
	ProcessName *string `json:"ProcessName,omitempty" xml:"ProcessName,omitempty"`
	// The user who launched the process.
	//
	// example:
	//
	// alice
	ProcessUser *string `json:"ProcessUser,omitempty" xml:"ProcessUser,omitempty"`
}

func (s DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) GoString() string {
	return s.String()
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) GetCommand() *string {
	return s.Command
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) GetProcessId() *int64 {
	return s.ProcessId
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) GetProcessName() *string {
	return s.ProcessName
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) GetProcessUser() *string {
	return s.ProcessUser
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetCommand(v string) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.Command = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetGroupId(v string) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.GroupId = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetInstanceId(v string) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.InstanceId = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetProcessId(v int64) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.ProcessId = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetProcessName(v string) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.ProcessName = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) SetProcessUser(v string) *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess {
	s.ProcessUser = &v
	return s
}

func (s *DescribeMonitoringAgentProcessesResponseBodyNodeProcessesNodeProcess) Validate() error {
	return dara.Validate(s)
}
