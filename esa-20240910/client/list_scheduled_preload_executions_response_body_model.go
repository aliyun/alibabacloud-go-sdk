// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPreloadExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExecutions(v []*ListScheduledPreloadExecutionsResponseBodyExecutions) *ListScheduledPreloadExecutionsResponseBody
	GetExecutions() []*ListScheduledPreloadExecutionsResponseBodyExecutions
	SetRequestId(v string) *ListScheduledPreloadExecutionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListScheduledPreloadExecutionsResponseBody
	GetTotalCount() *int32
}

type ListScheduledPreloadExecutionsResponseBody struct {
	// The information about prefetch plans returned.
	Executions []*ListScheduledPreloadExecutionsResponseBodyExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// ET5BF670-09D5-4D0B-BEBY-D96A2A528000
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListScheduledPreloadExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPreloadExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadExecutionsResponseBody) GetExecutions() []*ListScheduledPreloadExecutionsResponseBodyExecutions {
	return s.Executions
}

func (s *ListScheduledPreloadExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListScheduledPreloadExecutionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListScheduledPreloadExecutionsResponseBody) SetExecutions(v []*ListScheduledPreloadExecutionsResponseBodyExecutions) *ListScheduledPreloadExecutionsResponseBody {
	s.Executions = v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBody) SetRequestId(v string) *ListScheduledPreloadExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBody) SetTotalCount(v int32) *ListScheduledPreloadExecutionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBody) Validate() error {
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

type ListScheduledPreloadExecutionsResponseBodyExecutions struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 15685865xxx14622
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The end time of the prefetch plan.
	//
	// example:
	//
	// 2024-05-31T18:10:48.849+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the prefetch plan.
	//
	// example:
	//
	// 66599bd7397885b43804901c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time interval between each batch execution in the plan. Unit: seconds.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The ID of the prefetch task.
	//
	// example:
	//
	// 665d3af3621bccf3fe29e1a4
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The number of URLs prefetched in each batch.
	//
	// example:
	//
	// 10
	SliceLen *int32 `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	// The start time of the prefetch plan.
	//
	// example:
	//
	// 2024-05-31T17:10:48.849+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the prefetch plan. Valid values:
	//
	// 	- **waiting**
	//
	// 	- **running**
	//
	// 	- **finished**
	//
	// 	- **failed**
	//
	// 	- **stopped**
	//
	// example:
	//
	// running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListScheduledPreloadExecutionsResponseBodyExecutions) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPreloadExecutionsResponseBodyExecutions) GoString() string {
	return s.String()
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) GetAliUid() *string {
	return s.AliUid
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) GetEndTime() *string {
	return s.EndTime
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) GetId() *string {
	return s.Id
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) GetInterval() *int32 {
	return s.Interval
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) GetJobId() *string {
	return s.JobId
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) GetSliceLen() *int32 {
	return s.SliceLen
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) GetStartTime() *string {
	return s.StartTime
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) GetStatus() *string {
	return s.Status
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetAliUid(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.AliUid = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetEndTime(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.EndTime = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetId(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.Id = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetInterval(v int32) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.Interval = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetJobId(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.JobId = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetSliceLen(v int32) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.SliceLen = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetStartTime(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.StartTime = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) SetStatus(v string) *ListScheduledPreloadExecutionsResponseBodyExecutions {
	s.Status = &v
	return s
}

func (s *ListScheduledPreloadExecutionsResponseBodyExecutions) Validate() error {
	return dara.Validate(s)
}
