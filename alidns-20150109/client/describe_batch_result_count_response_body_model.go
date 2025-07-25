// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBatchResultCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBatchType(v string) *DescribeBatchResultCountResponseBody
	GetBatchType() *string
	SetFailedCount(v int32) *DescribeBatchResultCountResponseBody
	GetFailedCount() *int32
	SetReason(v string) *DescribeBatchResultCountResponseBody
	GetReason() *string
	SetRequestId(v string) *DescribeBatchResultCountResponseBody
	GetRequestId() *string
	SetStatus(v int32) *DescribeBatchResultCountResponseBody
	GetStatus() *int32
	SetSuccessCount(v int32) *DescribeBatchResultCountResponseBody
	GetSuccessCount() *int32
	SetTaskId(v int64) *DescribeBatchResultCountResponseBody
	GetTaskId() *int64
	SetTotalCount(v int32) *DescribeBatchResultCountResponseBody
	GetTotalCount() *int32
}

type DescribeBatchResultCountResponseBody struct {
	// The type of the batch operation.
	//
	// example:
	//
	// DOMAIN_ADD
	BatchType *string `json:"BatchType,omitempty" xml:"BatchType,omitempty"`
	// The total number of domain names or DNS records that failed to be processed.
	//
	// example:
	//
	// 2
	FailedCount *int32 `json:"FailedCount,omitempty" xml:"FailedCount,omitempty"`
	// The cause of the execution failure.
	//
	// example:
	//
	// failed_reason
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75446CC1-FC9A-4595-8D96-089D73D7A63D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the task. Valid values:
	//
	// 	- **-1**: No task for importing domain names or DNS records is submitted.
	//
	// 	- **0**: The task is being processed.
	//
	// 	- **1**: The task is complete.
	//
	// 	- **2**: The task failed.
	//
	// example:
	//
	// 0
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of domain names or DNS records that were processed.
	//
	// example:
	//
	// 2
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The ID of the last task.
	//
	// example:
	//
	// 123456
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The total number of DNS records that were processed in batches.
	//
	// example:
	//
	// 4
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeBatchResultCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeBatchResultCountResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeBatchResultCountResponseBody) GetBatchType() *string {
	return s.BatchType
}

func (s *DescribeBatchResultCountResponseBody) GetFailedCount() *int32 {
	return s.FailedCount
}

func (s *DescribeBatchResultCountResponseBody) GetReason() *string {
	return s.Reason
}

func (s *DescribeBatchResultCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeBatchResultCountResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeBatchResultCountResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *DescribeBatchResultCountResponseBody) GetTaskId() *int64 {
	return s.TaskId
}

func (s *DescribeBatchResultCountResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeBatchResultCountResponseBody) SetBatchType(v string) *DescribeBatchResultCountResponseBody {
	s.BatchType = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetFailedCount(v int32) *DescribeBatchResultCountResponseBody {
	s.FailedCount = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetReason(v string) *DescribeBatchResultCountResponseBody {
	s.Reason = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetRequestId(v string) *DescribeBatchResultCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetStatus(v int32) *DescribeBatchResultCountResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetSuccessCount(v int32) *DescribeBatchResultCountResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetTaskId(v int64) *DescribeBatchResultCountResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) SetTotalCount(v int32) *DescribeBatchResultCountResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeBatchResultCountResponseBody) Validate() error {
	return dara.Validate(s)
}
