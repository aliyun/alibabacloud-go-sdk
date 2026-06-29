// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopScheduledPreloadExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *StopScheduledPreloadExecutionResponseBody
	GetAliUid() *string
	SetEndTime(v string) *StopScheduledPreloadExecutionResponseBody
	GetEndTime() *string
	SetId(v string) *StopScheduledPreloadExecutionResponseBody
	GetId() *string
	SetInterval(v int32) *StopScheduledPreloadExecutionResponseBody
	GetInterval() *int32
	SetJobId(v string) *StopScheduledPreloadExecutionResponseBody
	GetJobId() *string
	SetRequestId(v string) *StopScheduledPreloadExecutionResponseBody
	GetRequestId() *string
	SetSliceLen(v int32) *StopScheduledPreloadExecutionResponseBody
	GetSliceLen() *int32
	SetStartTime(v string) *StopScheduledPreloadExecutionResponseBody
	GetStartTime() *string
	SetStatus(v string) *StopScheduledPreloadExecutionResponseBody
	GetStatus() *string
}

type StopScheduledPreloadExecutionResponseBody struct {
	// The Alibaba Cloud account ID.
	//
	// example:
	//
	// 15685865xxx14622
	AliUid *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The end time of the scheduled preload plan, in ISO 8601 format (e.g., 2024-01-01T00:00:00+Z).
	//
	// example:
	//
	// 2024-05-31T18:10:48.849+08:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The preload plan ID.
	//
	// example:
	//
	// 66599bd7397885b43804901c
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The execution interval between batches in the scheduled preload plan, in seconds.
	//
	// example:
	//
	// 60
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// The preload task ID.
	//
	// example:
	//
	// 665d3af3621bccf3fe29e1a4
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-A198-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of URLs per batch in the scheduled preload.
	//
	// example:
	//
	// 10
	SliceLen *int32 `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	// The start time of the scheduled preload plan, in ISO 8601 format (e.g., 2024-01-01T00:00:00+Z).
	//
	// example:
	//
	// 2024-05-31T17:10:48.849+08:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of the scheduled preload plan. Valid values:
	//
	// - **waiting**: Waiting to be executed.
	//
	// - **running**: Being executed.
	//
	// - **finished**: Execution completed.
	//
	// - **failed**: Execution failed.
	//
	// - **stopped**: Execution paused.
	//
	// example:
	//
	// stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s StopScheduledPreloadExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopScheduledPreloadExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *StopScheduledPreloadExecutionResponseBody) GetAliUid() *string {
	return s.AliUid
}

func (s *StopScheduledPreloadExecutionResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *StopScheduledPreloadExecutionResponseBody) GetId() *string {
	return s.Id
}

func (s *StopScheduledPreloadExecutionResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *StopScheduledPreloadExecutionResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *StopScheduledPreloadExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopScheduledPreloadExecutionResponseBody) GetSliceLen() *int32 {
	return s.SliceLen
}

func (s *StopScheduledPreloadExecutionResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *StopScheduledPreloadExecutionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *StopScheduledPreloadExecutionResponseBody) SetAliUid(v string) *StopScheduledPreloadExecutionResponseBody {
	s.AliUid = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetEndTime(v string) *StopScheduledPreloadExecutionResponseBody {
	s.EndTime = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetId(v string) *StopScheduledPreloadExecutionResponseBody {
	s.Id = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetInterval(v int32) *StopScheduledPreloadExecutionResponseBody {
	s.Interval = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetJobId(v string) *StopScheduledPreloadExecutionResponseBody {
	s.JobId = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetRequestId(v string) *StopScheduledPreloadExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetSliceLen(v int32) *StopScheduledPreloadExecutionResponseBody {
	s.SliceLen = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetStartTime(v string) *StopScheduledPreloadExecutionResponseBody {
	s.StartTime = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) SetStatus(v string) *StopScheduledPreloadExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
