// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMapRunResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConcurrency(v int64) *DescribeMapRunResponseBody
	GetConcurrency() *int64
	SetExecutionName(v string) *DescribeMapRunResponseBody
	GetExecutionName() *string
	SetItemCounts(v *DescribeMapRunResponseBodyItemCounts) *DescribeMapRunResponseBody
	GetItemCounts() *DescribeMapRunResponseBodyItemCounts
	SetMapRunName(v string) *DescribeMapRunResponseBody
	GetMapRunName() *string
	SetRequestId(v string) *DescribeMapRunResponseBody
	GetRequestId() *string
	SetStartedTime(v string) *DescribeMapRunResponseBody
	GetStartedTime() *string
	SetStatus(v string) *DescribeMapRunResponseBody
	GetStatus() *string
	SetStoppedTime(v string) *DescribeMapRunResponseBody
	GetStoppedTime() *string
	SetToleratedFailedCount(v int64) *DescribeMapRunResponseBody
	GetToleratedFailedCount() *int64
	SetToleratedFailedPercentage(v float32) *DescribeMapRunResponseBody
	GetToleratedFailedPercentage() *float32
}

type DescribeMapRunResponseBody struct {
	// example:
	//
	// 1
	Concurrency *int64 `json:"Concurrency,omitempty" xml:"Concurrency,omitempty"`
	// example:
	//
	// my_exec_name
	ExecutionName *string                               `json:"ExecutionName,omitempty" xml:"ExecutionName,omitempty"`
	ItemCounts    *DescribeMapRunResponseBodyItemCounts `json:"ItemCounts,omitempty" xml:"ItemCounts,omitempty" type:"Struct"`
	// example:
	//
	// c39142f1345b196d678333c41f113000
	MapRunName *string `json:"MapRunName,omitempty" xml:"MapRunName,omitempty"`
	// example:
	//
	// 3A44E113-9962-5B0B-AB92-14060EFE3164
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 2025-10-24T14:11:26+08:00
	StartedTime *string `json:"StartedTime,omitempty" xml:"StartedTime,omitempty"`
	// example:
	//
	// Succeeded
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 2025-10-24T14:11:28+08:00
	StoppedTime *string `json:"StoppedTime,omitempty" xml:"StoppedTime,omitempty"`
	// example:
	//
	// 100
	ToleratedFailedCount *int64 `json:"ToleratedFailedCount,omitempty" xml:"ToleratedFailedCount,omitempty"`
	// example:
	//
	// 20
	ToleratedFailedPercentage *float32 `json:"ToleratedFailedPercentage,omitempty" xml:"ToleratedFailedPercentage,omitempty"`
}

func (s DescribeMapRunResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMapRunResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMapRunResponseBody) GetConcurrency() *int64 {
	return s.Concurrency
}

func (s *DescribeMapRunResponseBody) GetExecutionName() *string {
	return s.ExecutionName
}

func (s *DescribeMapRunResponseBody) GetItemCounts() *DescribeMapRunResponseBodyItemCounts {
	return s.ItemCounts
}

func (s *DescribeMapRunResponseBody) GetMapRunName() *string {
	return s.MapRunName
}

func (s *DescribeMapRunResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMapRunResponseBody) GetStartedTime() *string {
	return s.StartedTime
}

func (s *DescribeMapRunResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeMapRunResponseBody) GetStoppedTime() *string {
	return s.StoppedTime
}

func (s *DescribeMapRunResponseBody) GetToleratedFailedCount() *int64 {
	return s.ToleratedFailedCount
}

func (s *DescribeMapRunResponseBody) GetToleratedFailedPercentage() *float32 {
	return s.ToleratedFailedPercentage
}

func (s *DescribeMapRunResponseBody) SetConcurrency(v int64) *DescribeMapRunResponseBody {
	s.Concurrency = &v
	return s
}

func (s *DescribeMapRunResponseBody) SetExecutionName(v string) *DescribeMapRunResponseBody {
	s.ExecutionName = &v
	return s
}

func (s *DescribeMapRunResponseBody) SetItemCounts(v *DescribeMapRunResponseBodyItemCounts) *DescribeMapRunResponseBody {
	s.ItemCounts = v
	return s
}

func (s *DescribeMapRunResponseBody) SetMapRunName(v string) *DescribeMapRunResponseBody {
	s.MapRunName = &v
	return s
}

func (s *DescribeMapRunResponseBody) SetRequestId(v string) *DescribeMapRunResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMapRunResponseBody) SetStartedTime(v string) *DescribeMapRunResponseBody {
	s.StartedTime = &v
	return s
}

func (s *DescribeMapRunResponseBody) SetStatus(v string) *DescribeMapRunResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeMapRunResponseBody) SetStoppedTime(v string) *DescribeMapRunResponseBody {
	s.StoppedTime = &v
	return s
}

func (s *DescribeMapRunResponseBody) SetToleratedFailedCount(v int64) *DescribeMapRunResponseBody {
	s.ToleratedFailedCount = &v
	return s
}

func (s *DescribeMapRunResponseBody) SetToleratedFailedPercentage(v float32) *DescribeMapRunResponseBody {
	s.ToleratedFailedPercentage = &v
	return s
}

func (s *DescribeMapRunResponseBody) Validate() error {
	if s.ItemCounts != nil {
		if err := s.ItemCounts.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeMapRunResponseBodyItemCounts struct {
	// example:
	//
	// 100
	Aborted *int64 `json:"Aborted,omitempty" xml:"Aborted,omitempty"`
	// example:
	//
	// 100
	Failed *int64 `json:"Failed,omitempty" xml:"Failed,omitempty"`
	// example:
	//
	// 100
	Pending *int64 `json:"Pending,omitempty" xml:"Pending,omitempty"`
	// example:
	//
	// 100
	Running *int64 `json:"Running,omitempty" xml:"Running,omitempty"`
	// example:
	//
	// 100
	Succeed *int64 `json:"Succeed,omitempty" xml:"Succeed,omitempty"`
	// example:
	//
	// 500
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s DescribeMapRunResponseBodyItemCounts) String() string {
	return dara.Prettify(s)
}

func (s DescribeMapRunResponseBodyItemCounts) GoString() string {
	return s.String()
}

func (s *DescribeMapRunResponseBodyItemCounts) GetAborted() *int64 {
	return s.Aborted
}

func (s *DescribeMapRunResponseBodyItemCounts) GetFailed() *int64 {
	return s.Failed
}

func (s *DescribeMapRunResponseBodyItemCounts) GetPending() *int64 {
	return s.Pending
}

func (s *DescribeMapRunResponseBodyItemCounts) GetRunning() *int64 {
	return s.Running
}

func (s *DescribeMapRunResponseBodyItemCounts) GetSucceed() *int64 {
	return s.Succeed
}

func (s *DescribeMapRunResponseBodyItemCounts) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeMapRunResponseBodyItemCounts) SetAborted(v int64) *DescribeMapRunResponseBodyItemCounts {
	s.Aborted = &v
	return s
}

func (s *DescribeMapRunResponseBodyItemCounts) SetFailed(v int64) *DescribeMapRunResponseBodyItemCounts {
	s.Failed = &v
	return s
}

func (s *DescribeMapRunResponseBodyItemCounts) SetPending(v int64) *DescribeMapRunResponseBodyItemCounts {
	s.Pending = &v
	return s
}

func (s *DescribeMapRunResponseBodyItemCounts) SetRunning(v int64) *DescribeMapRunResponseBodyItemCounts {
	s.Running = &v
	return s
}

func (s *DescribeMapRunResponseBodyItemCounts) SetSucceed(v int64) *DescribeMapRunResponseBodyItemCounts {
	s.Succeed = &v
	return s
}

func (s *DescribeMapRunResponseBodyItemCounts) SetTotal(v int64) *DescribeMapRunResponseBodyItemCounts {
	s.Total = &v
	return s
}

func (s *DescribeMapRunResponseBodyItemCounts) Validate() error {
	return dara.Validate(s)
}
