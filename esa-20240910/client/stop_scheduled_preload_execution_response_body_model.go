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
	AliUid    *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
