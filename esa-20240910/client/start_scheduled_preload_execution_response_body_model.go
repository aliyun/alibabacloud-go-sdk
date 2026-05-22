// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartScheduledPreloadExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *StartScheduledPreloadExecutionResponseBody
	GetAliUid() *string
	SetEndTime(v string) *StartScheduledPreloadExecutionResponseBody
	GetEndTime() *string
	SetId(v string) *StartScheduledPreloadExecutionResponseBody
	GetId() *string
	SetInterval(v int32) *StartScheduledPreloadExecutionResponseBody
	GetInterval() *int32
	SetJobId(v string) *StartScheduledPreloadExecutionResponseBody
	GetJobId() *string
	SetRequestId(v string) *StartScheduledPreloadExecutionResponseBody
	GetRequestId() *string
	SetSliceLen(v int32) *StartScheduledPreloadExecutionResponseBody
	GetSliceLen() *int32
	SetStartTime(v string) *StartScheduledPreloadExecutionResponseBody
	GetStartTime() *string
	SetStatus(v string) *StartScheduledPreloadExecutionResponseBody
	GetStatus() *string
}

type StartScheduledPreloadExecutionResponseBody struct {
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
	// 665d3b48621bccf3fe29e1a7
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
	// The request ID.
	//
	// example:
	//
	// 65C66B7B-671A-8297-9187-2R5477247B76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// waiting
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StartScheduledPreloadExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartScheduledPreloadExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StartScheduledPreloadExecutionResponseBody) GetAliUid() *string {
	return s.AliUid
}

func (s *StartScheduledPreloadExecutionResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *StartScheduledPreloadExecutionResponseBody) GetId() *string {
	return s.Id
}

func (s *StartScheduledPreloadExecutionResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *StartScheduledPreloadExecutionResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *StartScheduledPreloadExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartScheduledPreloadExecutionResponseBody) GetSliceLen() *int32 {
	return s.SliceLen
}

func (s *StartScheduledPreloadExecutionResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *StartScheduledPreloadExecutionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *StartScheduledPreloadExecutionResponseBody) SetAliUid(v string) *StartScheduledPreloadExecutionResponseBody {
	s.AliUid = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetEndTime(v string) *StartScheduledPreloadExecutionResponseBody {
	s.EndTime = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetId(v string) *StartScheduledPreloadExecutionResponseBody {
	s.Id = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetInterval(v int32) *StartScheduledPreloadExecutionResponseBody {
	s.Interval = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetJobId(v string) *StartScheduledPreloadExecutionResponseBody {
	s.JobId = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetRequestId(v string) *StartScheduledPreloadExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetSliceLen(v int32) *StartScheduledPreloadExecutionResponseBody {
	s.SliceLen = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetStartTime(v string) *StartScheduledPreloadExecutionResponseBody {
	s.StartTime = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) SetStatus(v string) *StartScheduledPreloadExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
