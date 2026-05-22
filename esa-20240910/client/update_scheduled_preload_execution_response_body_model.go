// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledPreloadExecutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v string) *UpdateScheduledPreloadExecutionResponseBody
	GetAliUid() *string
	SetEndTime(v string) *UpdateScheduledPreloadExecutionResponseBody
	GetEndTime() *string
	SetId(v string) *UpdateScheduledPreloadExecutionResponseBody
	GetId() *string
	SetInterval(v int32) *UpdateScheduledPreloadExecutionResponseBody
	GetInterval() *int32
	SetJobId(v string) *UpdateScheduledPreloadExecutionResponseBody
	GetJobId() *string
	SetRequestId(v string) *UpdateScheduledPreloadExecutionResponseBody
	GetRequestId() *string
	SetSliceLen(v int32) *UpdateScheduledPreloadExecutionResponseBody
	GetSliceLen() *int32
	SetStartTime(v string) *UpdateScheduledPreloadExecutionResponseBody
	GetStartTime() *string
	SetStatus(v string) *UpdateScheduledPreloadExecutionResponseBody
	GetStatus() *string
}

type UpdateScheduledPreloadExecutionResponseBody struct {
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

func (s UpdateScheduledPreloadExecutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledPreloadExecutionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPreloadExecutionResponseBody) GetAliUid() *string {
	return s.AliUid
}

func (s *UpdateScheduledPreloadExecutionResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateScheduledPreloadExecutionResponseBody) GetId() *string {
	return s.Id
}

func (s *UpdateScheduledPreloadExecutionResponseBody) GetInterval() *int32 {
	return s.Interval
}

func (s *UpdateScheduledPreloadExecutionResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *UpdateScheduledPreloadExecutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScheduledPreloadExecutionResponseBody) GetSliceLen() *int32 {
	return s.SliceLen
}

func (s *UpdateScheduledPreloadExecutionResponseBody) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateScheduledPreloadExecutionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetAliUid(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.AliUid = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetEndTime(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.EndTime = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetId(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.Id = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetInterval(v int32) *UpdateScheduledPreloadExecutionResponseBody {
	s.Interval = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetJobId(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetRequestId(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetSliceLen(v int32) *UpdateScheduledPreloadExecutionResponseBody {
	s.SliceLen = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetStartTime(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.StartTime = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) SetStatus(v string) *UpdateScheduledPreloadExecutionResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponseBody) Validate() error {
	return dara.Validate(s)
}
