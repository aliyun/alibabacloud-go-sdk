// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommandInvocationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommandInvocations(v []*DescribeCommandInvocationsResponseBodyCommandInvocations) *DescribeCommandInvocationsResponseBody
	GetCommandInvocations() []*DescribeCommandInvocationsResponseBodyCommandInvocations
	SetPageNumber(v int32) *DescribeCommandInvocationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCommandInvocationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeCommandInvocationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCommandInvocationsResponseBody
	GetTotalCount() *int32
}

type DescribeCommandInvocationsResponseBody struct {
	// The list of command executions.
	CommandInvocations []*DescribeCommandInvocationsResponseBodyCommandInvocations `json:"CommandInvocations,omitempty" xml:"CommandInvocations,omitempty" type:"Repeated"`
	// The page number of the query result.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page settings for a paged query. This is the paging size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 30637AD6-D977-4833-A54C-CC89483E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCommandInvocationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCommandInvocationsResponseBody) GetCommandInvocations() []*DescribeCommandInvocationsResponseBodyCommandInvocations {
	return s.CommandInvocations
}

func (s *DescribeCommandInvocationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCommandInvocationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCommandInvocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCommandInvocationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCommandInvocationsResponseBody) SetCommandInvocations(v []*DescribeCommandInvocationsResponseBodyCommandInvocations) *DescribeCommandInvocationsResponseBody {
	s.CommandInvocations = v
	return s
}

func (s *DescribeCommandInvocationsResponseBody) SetPageNumber(v int32) *DescribeCommandInvocationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBody) SetPageSize(v int32) *DescribeCommandInvocationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBody) SetRequestId(v string) *DescribeCommandInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBody) SetTotalCount(v int32) *DescribeCommandInvocationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBody) Validate() error {
	if s.CommandInvocations != nil {
		for _, item := range s.CommandInvocations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommandInvocationsResponseBodyCommandInvocations struct {
	// The command content.
	//
	// example:
	//
	// echo 123
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command description.
	//
	// example:
	//
	// testDescription
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// The command ID.
	//
	// example:
	//
	// c-hy0338xh28r****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The command name.
	//
	// example:
	//
	// testCommandName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The command type.
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The time when the command was created, in the format of YYYY-MM-DDThh:mm:ss.
	//
	// example:
	//
	// 2023-04-27T10:11:58
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The execution status of the command. Valid values:
	//
	// - Pending: The system is validating or sending the command.
	//
	// - Invalid: The specified command type or parameter is invalid.
	//
	// - Aborted: Failed to send the command to the instance. The instance must be running and the command must be delivered within 1 minute.
	//
	// - Running: The command is running on the instance.
	//
	// - Success: The command execution is complete and the exit code is 0.
	//
	// - Failed: The command execution is complete and the exit code is not 0.
	//
	// - Error: An exception occurred during command execution and the command cannot continue.
	//
	// - Timeout: The command execution timed out.
	//
	// - Cancelled: The command execution action has been canceled and the command was never started.
	//
	// - Stopping: The command execution is being stopped.
	//
	// - Terminated: The command was terminated during execution.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The command execution ID.
	//
	// example:
	//
	// t-hz0373jyzxt****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The list of target instances for execution.
	InvokeInstances []*DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances `json:"InvokeInstances,omitempty" xml:"InvokeInstances,omitempty" type:"Repeated"`
	// The custom parameters included in the command. If no parameters exist, the default value is {}.
	//
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The timeout period, in seconds.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The username used to run the command.
	//
	// example:
	//
	// root
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The path in which the command was executed.
	//
	// example:
	//
	// c:\\wwwroot
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeCommandInvocationsResponseBodyCommandInvocations) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandInvocationsResponseBodyCommandInvocations) GoString() string {
	return s.String()
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetCommandContent() *string {
	return s.CommandContent
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetCommandDescription() *string {
	return s.CommandDescription
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetCommandName() *string {
	return s.CommandName
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetCommandType() *string {
	return s.CommandType
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetInvokeInstances() []*DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	return s.InvokeInstances
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetParameters() *string {
	return s.Parameters
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetTimeout() *int64 {
	return s.Timeout
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetUsername() *string {
	return s.Username
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCommandContent(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CommandContent = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCommandDescription(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CommandDescription = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCommandId(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CommandId = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCommandName(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CommandName = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCommandType(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CommandType = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetCreationTime(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.CreationTime = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetInvocationStatus(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetInvokeId(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.InvokeId = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetInvokeInstances(v []*DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.InvokeInstances = v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetParameters(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.Parameters = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetTimeout(v int64) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.Timeout = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetUsername(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.Username = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) SetWorkingDir(v string) *DescribeCommandInvocationsResponseBodyCommandInvocations {
	s.WorkingDir = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocations) Validate() error {
	if s.InvokeInstances != nil {
		for _, item := range s.InvokeInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances struct {
	// The error code for the command delivery failure or execution failure. Valid values:
	//
	// - Empty: The command executed normally.
	//
	// - InstanceNotExists: The specified instance does not exist or has been released.
	//
	// - InstanceReleased: The instance was released during command execution.
	//
	// - InstanceNotRunning: The instance was not running when the command was executed.
	//
	// - CommandNotApplicable: The command is not applicable to the specified instance.
	//
	// - AccountNotExists: The specified account does not exist.
	//
	// - DirectoryNotExists: The specified directory does not exist.
	//
	// - BadCronExpression: The specified cron expression is invalid.
	//
	// - ClientNotRunning: The Cloud Assistant Agent is not running.
	//
	// - ClientNotResponse: The Cloud Assistant Agent is not responding.
	//
	// - ClientIsUpgrading: The Cloud Assistant Agent is being upgraded.
	//
	// - ClientNeedUpgrade: The Cloud Assistant Agent needs to be upgraded.
	//
	// - DeliveryTimeout: Command delivery timed out.
	//
	// - ExecutionTimeout: Command execution timed out.
	//
	// - ExecutionException: An exception occurred during command execution.
	//
	// - ExecutionInterrupted: Command execution was interrupted.
	//
	// - ExitCodeNonzero: Command execution completed with a non-zero exit code.
	//
	// example:
	//
	// InstanceNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The details of the command delivery failure or execution failure. Valid values:
	//
	// - Empty: The command executed normally.
	//
	// - the specified instance does not exists: The specified instance does not exist or has been released.
	//
	// - the instance has released when create task: The instance was released during command execution.
	//
	// - the instance is not running when create task: The instance was not running when the command was executed.
	//
	// - the command is not applicable: The command is not applicable to the specified instance.
	//
	// - the specified account does not exists: The specified account does not exist.
	//
	// - the specified directory does not exists: The specified directory does not exist.
	//
	// - the cron job expression is invalid: The specified cron expression is invalid.
	//
	// - the aliyun service is not running on the instance: The Cloud Assistant Agent is not running.
	//
	// - the aliyun service in the instance does not response: The Cloud Assistant Agent is not responding.
	//
	// - the aliyun service in the instance is upgrading now: The Cloud Assistant Agent is being upgraded.
	//
	// - the aliyun service in the instance need upgrade: The Cloud Assistant Agent needs to be upgraded.
	//
	// - the command delivery has been timeout: Command delivery timed out.
	//
	// - the command execution has been timeout: Command execution timed out.
	//
	// - the command execution got an exception: An exception occurred during command execution.
	//
	// - the command execution has been interrupted: Command execution was interrupted.
	//
	// - the command execution exit code is not zero: Command execution completed with a non-zero exit code.
	//
	// example:
	//
	// the specified instance does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the command.
	//
	// - For Linux instances, this is the exit code of the Shell command.
	//
	// - For Windows instances, this is the exit code of the Bat or PowerShell command.
	//
	// example:
	//
	// 0
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the command process ended. The time is displayed in UTC (Z) in the format of YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2023-04-03T02:42:29Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The instance ID of the simple application server.
	//
	// example:
	//
	// 2445f4aecdac4b71ba2c7e3a7ccf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The command execution status on a single instance. Valid values:
	//
	// - Pending: The system is validating or sending the command.
	//
	// - Invalid: The specified command type or parameter is invalid.
	//
	// - Aborted: Failed to send the command to the instance. The instance must be running and the command must be delivered within 1 minute.
	//
	// - Running: The command is running on the instance.
	//
	// - Success:
	//
	//     - One-time task: The command execution is complete and the exit code is 0.
	//
	//     - Scheduled task: The last execution succeeded with an exit code of 0, and the specified schedule has ended.
	//
	// - Failed:
	//
	//     -  One-time task: The command execution is complete and the exit code is not 0.
	//
	//     -  Scheduled task: The last execution succeeded with an exit code that is not 0, and the specified schedule will be aborted.
	//
	// - Error: An exception occurred during command execution and the command cannot continue.
	//
	// - Timeout: The command execution timed out.
	//
	// - Cancelled: The command execution action has been canceled and the command was never started.
	//
	// - Stopping: The task is being stopped.
	//
	// - Terminated: The command was terminated during execution.
	//
	// example:
	//
	// Running
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The output information of the command execution.
	//
	// example:
	//
	// OutputMsg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The time when the command started running on the instance. The time is displayed in UTC (Z) in the format of YYYY-MM-DDThh:mm:ssZ.
	//
	// example:
	//
	// 2023-05-09T03:32:24Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) GoString() string {
	return s.String()
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) GetExitCode() *int64 {
	return s.ExitCode
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) GetOutput() *string {
	return s.Output
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetErrorCode(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.ErrorCode = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetErrorInfo(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetExitCode(v int64) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.ExitCode = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetFinishTime(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.FinishTime = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetInstanceId(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetInvocationStatus(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetOutput(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.Output = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) SetStartTime(v string) *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances {
	s.StartTime = &v
	return s
}

func (s *DescribeCommandInvocationsResponseBodyCommandInvocationsInvokeInstances) Validate() error {
	return dara.Validate(s)
}
