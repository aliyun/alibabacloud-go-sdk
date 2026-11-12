// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvocationResult(v *DescribeInvocationResultResponseBodyInvocationResult) *DescribeInvocationResultResponseBody
	GetInvocationResult() *DescribeInvocationResultResponseBodyInvocationResult
	SetRequestId(v string) *DescribeInvocationResultResponseBody
	GetRequestId() *string
}

type DescribeInvocationResultResponseBody struct {
	// The command execution result.
	InvocationResult *DescribeInvocationResultResponseBodyInvocationResult `json:"InvocationResult,omitempty" xml:"InvocationResult,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvocationResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultResponseBody) GetInvocationResult() *DescribeInvocationResultResponseBodyInvocationResult {
	return s.InvocationResult
}

func (s *DescribeInvocationResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvocationResultResponseBody) SetInvocationResult(v *DescribeInvocationResultResponseBodyInvocationResult) *DescribeInvocationResultResponseBody {
	s.InvocationResult = v
	return s
}

func (s *DescribeInvocationResultResponseBody) SetRequestId(v string) *DescribeInvocationResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationResultResponseBody) Validate() error {
	if s.InvocationResult != nil {
		if err := s.InvocationResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInvocationResultResponseBodyInvocationResult struct {
	// The error code for the reason why the command failed to be sent or run. Valid values:
	//
	// - Empty: The command was run normally.
	//
	// - InstanceNotExists: The specified instance does not exist or has been released.
	//
	// - InstanceReleased: The instance was released during command execution.
	//
	// - InstanceNotRunning: The instance was not in the Running state when the command was run.
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
	// - DeliveryTimeout: The command sending timed out.
	//
	// - ExecutionTimeout: The command execution timed out.
	//
	// - ExecutionException: An exception occurred during command execution.
	//
	// - ExecutionInterrupted: The command execution was interrupted.
	//
	// - ExitCodeNonzero: The command finished running, and the exit code is not 0.
	//
	// example:
	//
	// InstanceNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The details of the reason why the command failed to be sent or run. Valid values:
	//
	// - Empty: The command was run normally.
	//
	// - the specified instance does not exists: The specified instance does not exist or has been released.
	//
	// - the instance has released when create task: The instance was released during command execution.
	//
	// - the instance is not running when create task: The instance was not in the Running state when the command was run.
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
	// - the command delivery has been timeout: The command sending timed out.
	//
	// - the command execution has been timeout: The command execution timed out.
	//
	// - the command execution got an exception: An exception occurred during command execution.
	//
	// - the command execution has been interrupted: The command execution was interrupted.
	//
	// - the command execution exit code is not zero: The command finished running, and the exit code is not 0.
	//
	// example:
	//
	// the specified instance does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the command.
	//
	// - For Linux instances, the exit code is the exit code of the shell command.
	//
	// - For Windows instances, the exit code is the exit code of the Bat or PowerShell command.
	//
	// example:
	//
	// 0
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the command finished running on the instance. The time is displayed in UTC (Z) in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2022-07-11T06:37:17Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// The instance ID of the simple application server.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The execution status of the command on the instance. Valid values:
	//
	// - Pending: The system is verifying or sending the command.
	//
	// - Invalid: The specified command type or parameter is invalid.
	//
	// - Aborted: Failed to send the command to the instance. The instance must be in the Running state, and the command must be sent within 1 minute.
	//
	// - Running: The command is being run on the instance.
	//
	// - Success: The command finished running, and the exit code is 0.
	//
	// - Failed: The command finished running, and the exit code is not 0.
	//
	// - Error: An exception occurred during command execution and the command cannot continue.
	//
	// - Timeout: The command execution timed out.
	//
	// - Cancelled: The command execution has been canceled, and the command was never started.
	//
	// - Stopping: The running command is being stopped.
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
	// t-bj02prjje65****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The command execution status. Valid values:
	//
	// - Running: Running.
	//
	// - Finished: Finished.
	//
	// - Failed: Failed.
	//
	// - Stopped: Stopped.
	//
	// example:
	//
	// Finished
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// The username used to run the command on the simple application server.
	//
	// example:
	//
	// root
	InvokeUser *string `json:"InvokeUser,omitempty" xml:"InvokeUser,omitempty"`
	// The output of the command execution. The standard output (stdout) and standard error output (stderr) are merged into a single Output. The capacity is calculated separately for each execution and each instance. The maximum capacity of a single Output is 24 KB. Content that exceeds this limit is truncated and discarded. Even if the ExitCode is 0 or the execution status is Success, the output may still be truncated. The returned output is in UTF-8 plaintext and does not require Base64 decoding.
	//
	// example:
	//
	// admin
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The time when the command started to be run on the instance. The time is displayed in UTC (Z) in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2022-07-11T06:37:16Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInvocationResultResponseBodyInvocationResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultResponseBodyInvocationResult) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) GetExitCode() *int64 {
	return s.ExitCode
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) GetInvokeRecordStatus() *string {
	return s.InvokeRecordStatus
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) GetInvokeUser() *string {
	return s.InvokeUser
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) GetOutput() *string {
	return s.Output
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetErrorCode(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetErrorInfo(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetExitCode(v int64) *DescribeInvocationResultResponseBodyInvocationResult {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetFinishedTime(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.FinishedTime = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInstanceId(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvocationStatus(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvokeId(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvokeRecordStatus(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetInvokeUser(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.InvokeUser = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetOutput(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.Output = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) SetStartTime(v string) *DescribeInvocationResultResponseBodyInvocationResult {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationResultResponseBodyInvocationResult) Validate() error {
	return dara.Validate(s)
}
