// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommandInvocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *DescribeCommandInvocationsRequest
	GetCommandId() *string
	SetCommandName(v string) *DescribeCommandInvocationsRequest
	GetCommandName() *string
	SetCommandType(v string) *DescribeCommandInvocationsRequest
	GetCommandType() *string
	SetInstanceId(v string) *DescribeCommandInvocationsRequest
	GetInstanceId() *string
	SetInvocationStatus(v string) *DescribeCommandInvocationsRequest
	GetInvocationStatus() *string
	SetInvokeId(v string) *DescribeCommandInvocationsRequest
	GetInvokeId() *string
	SetPageNumber(v string) *DescribeCommandInvocationsRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeCommandInvocationsRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeCommandInvocationsRequest
	GetRegionId() *string
}

type DescribeCommandInvocationsRequest struct {
	// The command ID. You can call [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) to query all available command IDs.
	//
	// example:
	//
	// c-sh02yh0932w****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The command name. If the InstanceId parameter is also specified, this parameter does not take effect.
	//
	// example:
	//
	// testName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The command type. Valid values:
	//
	// - RunBatScript: Bat command for Windows instances.
	//
	// - RunPowerShellScript: PowerShell command for Windows instances.
	//
	// - RunShellScript: Shell command for Linux instances.
	//
	// example:
	//
	// RunPowerShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The instance ID of the simple application server.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The overall execution status of the command. The overall execution status depends on the common execution status across all instances in this invocation. Valid values:
	//
	// -  Pending: The system is validating or sending the command. If the command execution status on at least one instance is Pending, the overall execution status is Pending.
	//
	// -  Running: The command is running on instances. If the command execution status on at least one instance is Running, the overall execution status is Running.
	//
	// -  Success: The command execution status on each instance is Stopped or Success, and the command execution status on at least one instance is Success. In this case, the overall execution status is Success.
	//
	//     -  One-time task: The command execution is complete and the exit code is 0.
	//
	// -  Failed: The command execution status on each instance is Stopped or Failed. In this case, the overall execution status is Failed. The overall execution status is Failed when the command execution status on an instance is one or more of the following values:
	//
	//     - Command validation failed (Invalid).
	//
	//     - Command sending failed (Aborted).
	//
	//     - Command execution is complete but the exit code is not 0 (Failed).
	//
	//     - Command execution timed out (Timeout).
	//
	//     - Command execution encountered an exception (Error).
	//
	// -  Stopping: The task is being stopped. If the command execution status on at least one instance is Stopping, the overall execution status is Stopping.
	//
	// -  Stopped: The task is stopped. If the command execution status on all instances is Stopped, the overall execution status is Stopped. The overall execution status is Stopped when the command execution status on an instance is one of the following values:
	//
	//     - Task canceled (Cancelled).
	//
	//     - Task terminated (Terminated).
	//
	// -  PartialFailed: The command execution succeeded on some instances and failed on others. If the command execution status on each instance is Success, Failed, or Stopped, the overall execution status is PartialFailed.
	//
	// > The `InvokeStatus` response parameter has a similar meaning, but we recommend that you check the value of this parameter.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The command execution ID.
	//
	// example:
	//
	// t-hz02p9545t6****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The page number. Minimum value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page for a paged query.
	//
	// Maximum value: 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call [ListRegions](https://help.aliyun.com/document_detail/189315.html) to query available regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCommandInvocationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCommandInvocationsRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeCommandInvocationsRequest) GetCommandName() *string {
	return s.CommandName
}

func (s *DescribeCommandInvocationsRequest) GetCommandType() *string {
	return s.CommandType
}

func (s *DescribeCommandInvocationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCommandInvocationsRequest) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeCommandInvocationsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeCommandInvocationsRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeCommandInvocationsRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeCommandInvocationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCommandInvocationsRequest) SetCommandId(v string) *DescribeCommandInvocationsRequest {
	s.CommandId = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetCommandName(v string) *DescribeCommandInvocationsRequest {
	s.CommandName = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetCommandType(v string) *DescribeCommandInvocationsRequest {
	s.CommandType = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetInstanceId(v string) *DescribeCommandInvocationsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetInvocationStatus(v string) *DescribeCommandInvocationsRequest {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetInvokeId(v string) *DescribeCommandInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetPageNumber(v string) *DescribeCommandInvocationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetPageSize(v string) *DescribeCommandInvocationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) SetRegionId(v string) *DescribeCommandInvocationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCommandInvocationsRequest) Validate() error {
	return dara.Validate(s)
}
