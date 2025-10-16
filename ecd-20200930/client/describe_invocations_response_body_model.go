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
	// The command execution records.
	Invocations []*DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Repeated"`
	// The query token that is returned from this call.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
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
	// The Base64-encoded command content.
	//
	// example:
	//
	// cnBtIC1xYSB8IGdyZXAgdnNm****
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The type of the command.
	//
	// example:
	//
	// RunPowerShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The time when the execution task is created.
	//
	// example:
	//
	// 2020-12-19T09:15:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The ID of the end user.
	//
	// example:
	//
	// User1
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The overall execution status of the command. The value of this parameter depends on the execution status of the command on all the involved cloud computers. Valid values:
	//
	// 	- Pending: The command is being verified or sent. If the execution status is Pending on at least one cloud computer, the overall status is considered Pending.
	//
	// 	- Running: The command is being executed on cloud computers. If the execution status is Running on at least one cloud computer, the overall status is considered Running.
	//
	// 	- Success: If the execution status is Success on at least one cloud computer and either Success or Stopped on all other cloud computers, the overall status is considered Success.
	//
	// 	- Failed: If the execution status is Stopped or Failed on all cloud computers, the overall status is considered Failed. If any execution status on cloud computers matches one of the following values, Failed is returned.
	//
	//     	- Invalid: The command is invalid.
	//
	//     	- Aborted: The command failed to be sent.
	//
	//     	- Failed: The command is executed, but the exit code is not 0.
	//
	//     	- Timeout: The command execution timed out.
	//
	//     	- Error: An error occurred when the command is being executed.
	//
	// 	- Stopping: The command execution is being stopped. If the execution status is Stopping on at least one cloud computer, the overall status is considered Stopping.
	//
	// 	- Stopped: The command execution stops. If the execution status is Stopped on at least one cloud computer, the overall status is considered Stopped. If any execution status on cloud computers matches one of the following values, Stopped is returned.
	//
	//     	- Cancelled: The command execution is canceled.
	//
	//     	- Terminated: The command execution is terminated.
	//
	// 	- PartialFailed: The command execution succeeded on some cloud computers but failed on others. If the execution status on any cloud computer is Success, Failed, or Stopped, the overall status is considered PartialFailed.
	//
	// example:
	//
	// Pending
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The total number of cloud computers on which the command is executed.
	//
	// example:
	//
	// 1
	InvokeDesktopCount *int32 `json:"InvokeDesktopCount,omitempty" xml:"InvokeDesktopCount,omitempty"`
	// The total number of cloud computers on which the command execution succeeds.
	//
	// example:
	//
	// 1
	InvokeDesktopSucceedCount *int32 `json:"InvokeDesktopSucceedCount,omitempty" xml:"InvokeDesktopSucceedCount,omitempty"`
	// The cloud computers on which the command is executed.
	InvokeDesktops []*DescribeInvocationsResponseBodyInvocationsInvokeDesktops `json:"InvokeDesktops,omitempty" xml:"InvokeDesktops,omitempty" type:"Repeated"`
	// The ID of the execution.
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
	// The time when the command execution was performed.
	//
	// example:
	//
	// 2020-12-20T06:15:54Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-7w78ozhjcwa3u****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The cloud computer name.
	//
	// example:
	//
	// demo1234
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The size of the text that is truncated and discarded when the Output value exceeds 24 KB in size.
	//
	// example:
	//
	// 0
	Dropped *int32 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	// The code explaining why the command failed to be sent or executed. Valid values:
	//
	// 	- Null: The command is executed successfully.
	//
	// 	- InstanceNotExists: The specified cloud computer does not exist or is released.
	//
	// 	- InstanceReleased: The cloud computer is released during the execution.
	//
	// 	- InstanceNotRunning: The cloud computer is not running during the execution.
	//
	// 	- CommandNotApplicable: The command cannot be executed on the specified cloud computer.
	//
	// 	- ClientNotRunning: The Cloud Assistant agent is not running.
	//
	// 	- ClientNotResponse: The Cloud Assistant agent does not respond.
	//
	// 	- ClientIsUpgrading: The Cloud Assistant agent is being updated.
	//
	// 	- ClientNeedUpgrade: The Cloud Assistant agent needs to be updated.
	//
	// 	- DeliveryTimeout: The command sending times out.
	//
	// 	- ExecutionTimeout: The command execution times out.
	//
	// 	- ExecutionException: An exception occurs when the command is being executed.
	//
	// 	- ExecutionInterrupted: The command execution is interrupted.
	//
	// 	- ExitCodeNonzero: The command execution completes, but the exit code is not 0.
	//
	// example:
	//
	// InstanceNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The message explaining why the command failed to be sent or executed. Valid values:
	//
	// 	- Null: The command is executed successfully.
	//
	// 	- the specified instance does not exists: The specified cloud computer does not exist or is released.
	//
	// 	- the instance has released when create task: The cloud computer is released during the execution.
	//
	// 	- the instance is not running when create task: The cloud computer is not running during the execution.
	//
	// 	- the command is not applicable: The command cannot be executed on the specified cloud computer.
	//
	// 	- the aliyun service is not running on the instance: The Cloud Assistant agent is not running.
	//
	// 	- the aliyun service in the instance does not response: The Cloud Assistant agent does not respond.
	//
	// 	- the aliyun service in the instance is upgrading now: The Cloud Assistant agent is being updated.
	//
	// 	- the aliyun service in the instance need upgrade: The Cloud Assistant agent needs to be updated.
	//
	// 	- the command delivery has been timeout: The command sending times out.
	//
	// 	- the command execution has been timeout: The command execution times out.
	//
	// 	- the command execution got an exception: An exception occurs when the command is being executed.
	//
	// 	- the command execution has been interrupted: The command execution is interrupted.
	//
	// 	- the command execution exit code is not zero: The command execution completes, but the exit code is not 0.
	//
	// example:
	//
	// The specified instance does not exist.
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the execution.
	//
	// example:
	//
	// 0
	ExitCode *int64 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// The time when the command execution ended.
	//
	// example:
	//
	// 2020-12-20T06:15:56Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The execution progress of the command on a single cloud computer.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The command output.
	//
	// 	- When the `IncludeOutput` parameter is set to false, the output is not returned.
	//
	// 	- When the `ContentEncoding` parameter is set to Base64, the output is returned as a Base64-encoded string.
	//
	// example:
	//
	// OutPutTestmsg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The number of times the command has been executed on the cloud computer.
	//
	// example:
	//
	// 0
	Repeats *int32 `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	// The start time of the command execution.
	//
	// example:
	//
	// 2020-12-20T06:15:55Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The stop time of the command execution (StopInvocatio).
	//
	// example:
	//
	// 2020-12-25T09:15:47Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// The time when the execution status was updated.
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
