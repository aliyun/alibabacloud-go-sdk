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
	SetPageNumber(v int32) *DescribeInvocationsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInvocationsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInvocationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInvocationsResponseBody
	GetTotalCount() *int32
}

type DescribeInvocationsResponseBody struct {
	// The command execution records.
	Invocations []*DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Settings for paging in a paged query.
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
	// The total number of commands.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *DescribeInvocationsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInvocationsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInvocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvocationsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInvocationsResponseBody) SetInvocations(v []*DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetPageNumber(v int32) *DescribeInvocationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetPageSize(v int32) *DescribeInvocationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponseBody) SetTotalCount(v int32) *DescribeInvocationsResponseBody {
	s.TotalCount = &v
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
	// The command content, which is transmitted in Base64 encoding.
	//
	// example:
	//
	// bHM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command name.
	//
	// example:
	//
	// testname
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The type of the command. Valid values:
	//
	// - RunBatScript: a Bat command for Windows instances.
	//
	// - RunPowerShellScript: a PowerShell command for Windows instances.
	//
	// - RunShellScript: a shell command for Linux instances.
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// The time when the command was created. The time is displayed in UTC (Z) in the YYYY-MM-DDThh:mm:ssZ format.
	//
	// example:
	//
	// 2022-07-11T06:37:16Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The execution status of the command. Valid values:
	//
	// - Pending: The system is verifying or sending the command.
	//
	// - Invalid: The specified command type or parameter is invalid.
	//
	// - Aborted: The command failed to be sent to the instance. The instance must be in the Running state, and the command must be delivered within one minute.
	//
	// - Running: The command is being run on the instance.
	//
	// - Success: The command execution is complete, and the exit code is 0.
	//
	// - Failed: The command execution is complete, and the exit code is non-zero.
	//
	// - Error: An exception occurred during command execution, and the command could not continue.
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
	// t-hz02p9545t6****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// The status of the command. Valid values:
	//
	// - Running: The command is running.
	//
	// - Finished: The command execution is complete.
	//
	// - Failed: The command execution failed.
	//
	// - Stopped: The command has been stopped.
	//
	// example:
	//
	// Finished
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// The custom parameters included in the command. If no parameters exist, the default value is {}.
	//
	// example:
	//
	// {}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
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

func (s *DescribeInvocationsResponseBodyInvocations) GetCommandName() *string {
	return s.CommandName
}

func (s *DescribeInvocationsResponseBodyInvocations) GetCommandType() *string {
	return s.CommandType
}

func (s *DescribeInvocationsResponseBodyInvocations) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInvocationsResponseBodyInvocations) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeInvocationsResponseBodyInvocations) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationsResponseBodyInvocations) GetInvokeStatus() *string {
	return s.InvokeStatus
}

func (s *DescribeInvocationsResponseBodyInvocations) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandContent(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandContent = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetCommandName(v string) *DescribeInvocationsResponseBodyInvocations {
	s.CommandName = &v
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

func (s *DescribeInvocationsResponseBodyInvocations) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeId(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvokeStatus(v string) *DescribeInvocationsResponseBodyInvocations {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) SetParameters(v map[string]interface{}) *DescribeInvocationsResponseBodyInvocations {
	s.Parameters = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) Validate() error {
	return dara.Validate(s)
}
