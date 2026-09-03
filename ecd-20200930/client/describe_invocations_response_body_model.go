// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvocations(v []*DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody
	GetInvocations() []*DescribeInvocationsResponseBodyInvocations
	SetNextToken(v string) *DescribeInvocationsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeInvocationsResponseBody
	GetRequestId() *string
}

type DescribeInvocationsResponseBody struct {
	// The array of script execution records.
	Invocations []*DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Repeated"`
	// The pagination token returned in this call.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvocationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBody) GetInvocations() []*DescribeInvocationsResponseBodyInvocations {
	return s.Invocations
}

func (s *DescribeInvocationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInvocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvocationsResponseBody) SetInvocations(v []*DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetNextToken(v string) *DescribeInvocationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponseBody) Validate() error {
	if s.Invocations != nil {
		for _, item := range s.Invocations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInvocationsResponseBodyInvocations struct {
	// The script content, transmitted in Base64 encoding.
	//
	// example:
	//
	// cnBtIC1xYSB8IGdyZXAgdnNm****
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The script type.
	//
	// example:
	//
	// RunPowerShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The creation time of the task.
	//
	// example:
	//
	// 2020-12-19T09:15:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The cloud desktop scenario. Valid values:
	//
	// - Classic: the classic cloud desktop scenario.
	//
	// - JvsClaw: the JVS Claw cloud desktop scenario.
	//
	// example:
	//
	// Classic
	DesktopScenario *string `json:"DesktopScenario,omitempty" xml:"DesktopScenario,omitempty"`
	// The end user ID.
	//
	// example:
	//
	// User1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The overall execution status of the script. The overall execution status depends on the combined execution status of all cloud desktops in this call. Valid values:
	//
	// - Pending: The system is validating or sending the command. The overall execution status is Pending if at least one cloud desktop has a script execution status of Pending.
	//
	// - Running: The command is running on the cloud desktop. The overall execution status is Running if at least one cloud desktop has a script execution status of Running.
	//
	// - Success: The overall execution status is Success if the script execution status on each cloud desktop is Stopped or Success, and at least one cloud desktop has a script execution status of Success.
	//
	// - Failed: The overall execution status is Failed if the script execution status on each cloud desktop is Stopped or Failed. The return value is Failed when one or more of the following statuses occur on a cloud desktop:
	//
	//     - Command validation failed (Invalid)
	//
	//     - Command delivery failed (Aborted)
	//
	//     - Command execution completed with a non-zero exit code (Failed)
	//
	//     - Command execution timed out (Timeout)
	//
	//     - Command execution encountered an exception (Error)
	//
	// - Stopping: The task is being stopped. The overall execution status is Stopping if at least one instance has a script execution status of Stopping.
	//
	// - Stopped: The task has been stopped. The overall execution status is Stopped if the script execution status on all instances is Stopped. The return value is Stopped when the script execution status on an instance is one of the following:
	//
	//     - Task cancelled (Cancelled)
	//
	//     - Task terminated (Terminated)
	//
	// - PartialFailed: The overall execution status is PartialFailed if some instances succeeded and some instances failed. The overall execution status is PartialFailed if the script execution status on each instance is Success, Failed, or Stopped.
	//
	// example:
	//
	// Pending
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The total number of cloud desktops on which the script was executed.
	//
	// example:
	//
	// 1
	InvokeDesktopCount *int32 `json:"InvokeDesktopCount,omitempty" xml:"InvokeDesktopCount,omitempty"`
	// The total number of cloud desktops on which the script was executed successfully.
	//
	// example:
	//
	// 1
	InvokeDesktopSucceedCount *int32 `json:"InvokeDesktopSucceedCount,omitempty" xml:"InvokeDesktopSucceedCount,omitempty"`
	// The list of target cloud desktops for execution.
	InvokeDesktops []*DescribeInvocationsResponseBodyInvocationsInvokeDesktops `json:"InvokeDesktops,omitempty" xml:"InvokeDesktops,omitempty" type:"Repeated"`
	// The execution ID.
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocations) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocations) GetCommandContent() *string {
	return s.CommandContent
}

func (s *DescribeInvocationsResponseBodyInvocations) GetCommandType() *string {
	return s.CommandType
}

func (s *DescribeInvocationsResponseBodyInvocations) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInvocationsResponseBodyInvocations) GetDesktopScenario() *string {
	return s.DesktopScenario
}

func (s *DescribeInvocationsResponseBodyInvocations) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeInvocationsResponseBodyInvocations) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeInvocationsResponseBodyInvocations) GetInvokeDesktopCount() *int32 {
	return s.InvokeDesktopCount
}

func (s *DescribeInvocationsResponseBodyInvocations) GetInvokeDesktopSucceedCount() *int32 {
	return s.InvokeDesktopSucceedCount
}

func (s *DescribeInvocationsResponseBodyInvocations) GetInvokeDesktops() []*DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	return s.InvokeDesktops
}

func (s *DescribeInvocationsResponseBodyInvocations) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandContent(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandContent = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandType(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetDesktopScenario(v string) *DescribeInvocationsResponseBodyInvocations {
	s.DesktopScenario = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetEndUserId(v string) *DescribeInvocationsResponseBodyInvocations {
	s.EndUserId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeDesktopCount(v int32) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeDesktopCount = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeDesktopSucceedCount(v int32) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeDesktopSucceedCount = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeDesktops(v []*DescribeInvocationsResponseBodyInvocationsInvokeDesktops) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeDesktops = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeId(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) Validate() error {
	if s.InvokeDesktops != nil {
		for _, item := range s.InvokeDesktops {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInvocationsResponseBodyInvocationsInvokeDesktops struct {
	// The creation time of the script process.
	//
	// example:
	//
	// 2020-12-20T06:15:54Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The cloud desktop ID.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The cloud desktop name.
	//
	// example:
	//
	// demo1234
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The length of the truncated and discarded text after the text length in the Output field exceeded 24 KB.
	//
	// example:
	//
	// 0
	Dropped *int32 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	// The error code indicating the reason for command delivery failure or execution failure. Valid values:
	//
	// - Empty: The command ran normally.
	//
	// - InstanceNotExists: The specified cloud desktop does not exist or has been released.
	//
	// - InstanceReleased: The cloud desktop was released during task execution.
	//
	// - InstanceNotRunning: The cloud desktop was not running when the task was created.
	//
	// - CommandNotApplicable: The command is not applicable to the specified cloud desktop.
	//
	// - ClientNotRunning: The Cloud Assistant client is not running.
	//
	// - ClientNotResponse: The Cloud Assistant client is not responding.
	//
	// - ClientIsUpgrading: The Cloud Assistant client is being upgraded.
	//
	// - ClientNeedUpgrade: The Cloud Assistant client needs to be upgraded.
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
	// The detailed reason for command delivery failure or execution failure. Valid values:
	//
	// - Empty: The command ran normally.
	//
	// - the specified instance does not exists: The specified cloud desktop does not exist or has been released.
	//
	// - the instance has released when create task: The cloud desktop was released during task execution.
	//
	// - the instance is not running when create task: The cloud desktop was not running when the task was created.
	//
	// - the command is not applicable: The command is not applicable to the specified cloud desktop.
	//
	// - the aliyun service is not running on the instance: The Cloud Assistant client is not running.
	//
	// - the aliyun service in the instance does not response: The Cloud Assistant client is not responding.
	//
	// - the aliyun service in the instance is upgrading now: The Cloud Assistant client is being upgraded.
	//
	// - the aliyun service in the instance need upgrade: The Cloud Assistant client needs to be upgraded.
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
	// The specified instance does not exist.
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the script process.
	//
	// example:
	//
	// 0
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The end time of the script process.
	//
	// example:
	//
	// 2020-12-20T06:15:56Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The script process status on a single cloud desktop.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// jvs agent id。
	//
	// example:
	//
	// jvs-7xjos2l****
	JvsAgentId *string `json:"JvsAgentId,omitempty" xml:"JvsAgentId,omitempty"`
	// The output information of the script process.
	//
	// - If the request parameter `IncludeOutput` is set to false, Output is not returned.
	//
	// - If the request parameter `ContentEncoding` is set to Base64, Output is the Base64-encoded output information.
	//
	// example:
	//
	// OutPutTestmsg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The number of times the command was executed on the cloud desktop.
	//
	// example:
	//
	// 0
	Repeats *int32 `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	// The time when the script process started running on the cloud desktop.
	//
	// example:
	//
	// 2020-12-20T06:15:55Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when the execution was stopped, if StopInvocation was called.
	//
	// example:
	//
	// 2020-12-25T09:15:47Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// The update time of the task status.
	//
	// example:
	//
	// 2020-12-25T06:15:56Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvokeDesktops) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetDropped() *int32 {
	return s.Dropped
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetExitCode() *int64 {
	return s.ExitCode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetJvsAgentId() *string {
	return s.JvsAgentId
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetOutput() *string {
	return s.Output
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetRepeats() *int32 {
	return s.Repeats
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetDesktopId(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.DesktopId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetDesktopName(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.DesktopName = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetDropped(v int32) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.Dropped = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetErrorCode(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetErrorInfo(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetExitCode(v int64) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetFinishTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetJvsAgentId(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.JvsAgentId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetOutput(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetRepeats(v int32) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.Repeats = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetStartTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetStopTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.StopTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) SetUpdateTime(v string) *DescribeInvocationsResponseBodyInvocationsInvokeDesktops {
	s.UpdateTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvokeDesktops) Validate() error {
	return dara.Validate(s)
}
