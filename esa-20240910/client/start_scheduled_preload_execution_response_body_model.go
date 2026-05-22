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
