// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScheduledTaskExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeScheduledTaskExecutionsResponseBody
	GetCode() *string
	SetExecutions(v []*DescribeScheduledTaskExecutionsResponseBodyExecutions) *DescribeScheduledTaskExecutionsResponseBody
	GetExecutions() []*DescribeScheduledTaskExecutionsResponseBodyExecutions
	SetMaxResults(v int32) *DescribeScheduledTaskExecutionsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *DescribeScheduledTaskExecutionsResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeScheduledTaskExecutionsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeScheduledTaskExecutionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeScheduledTaskExecutionsResponseBody
	GetTotalCount() *int32
}

type DescribeScheduledTaskExecutionsResponseBody struct {
	// The status code of the operation.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The list of task execution records.
	Executions []*DescribeScheduledTaskExecutionsResponseBodyExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	// The maximum number of entries returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6kU****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9A51B1DF-96FF-3BCC-B08C-783161D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeScheduledTaskExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTaskExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTaskExecutionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeScheduledTaskExecutionsResponseBody) GetExecutions() []*DescribeScheduledTaskExecutionsResponseBodyExecutions {
	return s.Executions
}

func (s *DescribeScheduledTaskExecutionsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeScheduledTaskExecutionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeScheduledTaskExecutionsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeScheduledTaskExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeScheduledTaskExecutionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeScheduledTaskExecutionsResponseBody) SetCode(v string) *DescribeScheduledTaskExecutionsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBody) SetExecutions(v []*DescribeScheduledTaskExecutionsResponseBodyExecutions) *DescribeScheduledTaskExecutionsResponseBody {
	s.Executions = v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBody) SetMaxResults(v int32) *DescribeScheduledTaskExecutionsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBody) SetMessage(v string) *DescribeScheduledTaskExecutionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBody) SetNextToken(v string) *DescribeScheduledTaskExecutionsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBody) SetRequestId(v string) *DescribeScheduledTaskExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBody) SetTotalCount(v int32) *DescribeScheduledTaskExecutionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBody) Validate() error {
	if s.Executions != nil {
		for _, item := range s.Executions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeScheduledTaskExecutionsResponseBodyExecutions struct {
	// The end time.
	//
	// example:
	//
	// 2026-06-12T00:05:30
	CompletedAt *string `json:"CompletedAt,omitempty" xml:"CompletedAt,omitempty"`
	// The configuration snapshot in JSON format.
	//
	// example:
	//
	// {"maxSteps":10}
	ConfigSnapshot *string `json:"ConfigSnapshot,omitempty" xml:"ConfigSnapshot,omitempty"`
	// The execution duration in milliseconds.
	//
	// example:
	//
	// 330000
	DurationMs *int64 `json:"DurationMs,omitempty" xml:"DurationMs,omitempty"`
	// The error code.
	//
	// example:
	//
	// TaskTimeout
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// task timeout after 600s
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// acp-axxkuuxahbu1*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The execution output in JSON format.
	//
	// example:
	//
	// {"result":"ok"}
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The ID of the scheduled task.
	//
	// example:
	//
	// sch-260705-agb*****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 2026-06-12T00:00:00
	StartedAt *string `json:"StartedAt,omitempty" xml:"StartedAt,omitempty"`
	// The execution status.
	//
	// example:
	//
	// COMPLETED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the scheduled task execution record.
	//
	// example:
	//
	// t-260703-gby*****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeScheduledTaskExecutionsResponseBodyExecutions) String() string {
	return dara.Prettify(s)
}

func (s DescribeScheduledTaskExecutionsResponseBodyExecutions) GoString() string {
	return s.String()
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) GetCompletedAt() *string {
	return s.CompletedAt
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) GetConfigSnapshot() *string {
	return s.ConfigSnapshot
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) GetDurationMs() *int64 {
	return s.DurationMs
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) GetOutput() *string {
	return s.Output
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) GetStartedAt() *string {
	return s.StartedAt
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) GetStatus() *string {
	return s.Status
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) SetCompletedAt(v string) *DescribeScheduledTaskExecutionsResponseBodyExecutions {
	s.CompletedAt = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) SetConfigSnapshot(v string) *DescribeScheduledTaskExecutionsResponseBodyExecutions {
	s.ConfigSnapshot = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) SetDurationMs(v int64) *DescribeScheduledTaskExecutionsResponseBodyExecutions {
	s.DurationMs = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) SetErrorCode(v string) *DescribeScheduledTaskExecutionsResponseBodyExecutions {
	s.ErrorCode = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) SetErrorMessage(v string) *DescribeScheduledTaskExecutionsResponseBodyExecutions {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) SetInstanceId(v string) *DescribeScheduledTaskExecutionsResponseBodyExecutions {
	s.InstanceId = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) SetOutput(v string) *DescribeScheduledTaskExecutionsResponseBodyExecutions {
	s.Output = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) SetScheduledId(v string) *DescribeScheduledTaskExecutionsResponseBodyExecutions {
	s.ScheduledId = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) SetStartedAt(v string) *DescribeScheduledTaskExecutionsResponseBodyExecutions {
	s.StartedAt = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) SetStatus(v string) *DescribeScheduledTaskExecutionsResponseBodyExecutions {
	s.Status = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) SetTaskId(v string) *DescribeScheduledTaskExecutionsResponseBodyExecutions {
	s.TaskId = &v
	return s
}

func (s *DescribeScheduledTaskExecutionsResponseBodyExecutions) Validate() error {
	return dara.Validate(s)
}
