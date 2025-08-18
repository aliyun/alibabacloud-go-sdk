// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledPreloadExecutionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedExecutions(v []*CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) *CreateScheduledPreloadExecutionsResponseBody
	GetFailedExecutions() []*CreateScheduledPreloadExecutionsResponseBodyFailedExecutions
	SetFailedMessages(v []*string) *CreateScheduledPreloadExecutionsResponseBody
	GetFailedMessages() []*string
	SetRequestId(v string) *CreateScheduledPreloadExecutionsResponseBody
	GetRequestId() *string
	SetSuccessCount(v int32) *CreateScheduledPreloadExecutionsResponseBody
	GetSuccessCount() *int32
	SetSuccessExecutions(v []*CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) *CreateScheduledPreloadExecutionsResponseBody
	GetSuccessExecutions() []*CreateScheduledPreloadExecutionsResponseBodySuccessExecutions
	SetTotalCount(v int32) *CreateScheduledPreloadExecutionsResponseBody
	GetTotalCount() *int32
}

type CreateScheduledPreloadExecutionsResponseBody struct {
	// The information about prefetch plans that failed to be created.
	FailedExecutions []*CreateScheduledPreloadExecutionsResponseBodyFailedExecutions `json:"FailedExecutions,omitempty" xml:"FailedExecutions,omitempty" type:"Repeated"`
	// The information about plan failures.
	FailedMessages []*string `json:"FailedMessages,omitempty" xml:"FailedMessages,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of prefetch plans that are created.
	//
	// example:
	//
	// 12
	SuccessCount *int32 `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	// The information about created prefetch plans.
	SuccessExecutions []*CreateScheduledPreloadExecutionsResponseBodySuccessExecutions `json:"SuccessExecutions,omitempty" xml:"SuccessExecutions,omitempty" type:"Repeated"`
	// The total number of new plans requested.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CreateScheduledPreloadExecutionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsResponseBody) GetFailedExecutions() []*CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	return s.FailedExecutions
}

func (s *CreateScheduledPreloadExecutionsResponseBody) GetFailedMessages() []*string {
	return s.FailedMessages
}

func (s *CreateScheduledPreloadExecutionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduledPreloadExecutionsResponseBody) GetSuccessCount() *int32 {
	return s.SuccessCount
}

func (s *CreateScheduledPreloadExecutionsResponseBody) GetSuccessExecutions() []*CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	return s.SuccessExecutions
}

func (s *CreateScheduledPreloadExecutionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetFailedExecutions(v []*CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) *CreateScheduledPreloadExecutionsResponseBody {
	s.FailedExecutions = v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetFailedMessages(v []*string) *CreateScheduledPreloadExecutionsResponseBody {
	s.FailedMessages = v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetRequestId(v string) *CreateScheduledPreloadExecutionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetSuccessCount(v int32) *CreateScheduledPreloadExecutionsResponseBody {
	s.SuccessCount = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetSuccessExecutions(v []*CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) *CreateScheduledPreloadExecutionsResponseBody {
	s.SuccessExecutions = v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBody) SetTotalCount(v int32) *CreateScheduledPreloadExecutionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateScheduledPreloadExecutionsResponseBodyFailedExecutions struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 15685865xxx14622
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The end time of the prefetch plans.
	//
	// example:
	//
	// 2024-06-03T02:43:35Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the prefetch plan.
	//
	// example:
	//
	// 66599bd7397885b43804901c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time interval between each batch execution. Unit: seconds.
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
	// The start time of the prefetch plans.
	//
	// example:
	//
	// 2024-06-02T02:43:35Z
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

func (s CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) GetAliUid() *string {
	return s.AliUid
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) GetId() *string {
	return s.Id
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) GetJobId() *string {
	return s.JobId
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) GetSliceLen() *int32 {
	return s.SliceLen
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) GetStatus() *string {
	return s.Status
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetAliUid(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.AliUid = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetEndTime(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.EndTime = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetId(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.Id = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetInterval(v int32) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.Interval = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetJobId(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.JobId = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetSliceLen(v int32) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.SliceLen = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetStartTime(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.StartTime = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) SetStatus(v string) *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions {
	s.Status = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodyFailedExecutions) Validate() error {
	return dara.Validate(s)
}

type CreateScheduledPreloadExecutionsResponseBodySuccessExecutions struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 15685865xxx14622
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The end time of the prefetch plans.
	//
	// example:
	//
	// 2024-06-03T02:43:35Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the prefetch plan.
	//
	// example:
	//
	// 66599bd7397885b43804901c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time interval between each batch execution. Unit: seconds.
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
	// The start time of the prefetch plans.
	//
	// example:
	//
	// 2024-06-02T02:43:35Z
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
	// failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) GetAliUid() *string {
	return s.AliUid
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) GetId() *string {
	return s.Id
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) GetInterval() *int32 {
	return s.Interval
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) GetJobId() *string {
	return s.JobId
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) GetSliceLen() *int32 {
	return s.SliceLen
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) GetStatus() *string {
	return s.Status
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetAliUid(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.AliUid = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetEndTime(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.EndTime = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetId(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.Id = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetInterval(v int32) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.Interval = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetJobId(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.JobId = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetSliceLen(v int32) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.SliceLen = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetStartTime(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.StartTime = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) SetStatus(v string) *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions {
	s.Status = &v
	return s
}

func (s *CreateScheduledPreloadExecutionsResponseBodySuccessExecutions) Validate() error {
	return dara.Validate(s)
}
