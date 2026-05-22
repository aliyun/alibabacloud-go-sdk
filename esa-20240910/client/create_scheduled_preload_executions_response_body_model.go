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
	FailedExecutions  []*CreateScheduledPreloadExecutionsResponseBodyFailedExecutions  `json:"FailedExecutions,omitempty" xml:"FailedExecutions,omitempty" type:"Repeated"`
	FailedMessages    []*string                                                        `json:"FailedMessages,omitempty" xml:"FailedMessages,omitempty" type:"Repeated"`
	RequestId         *string                                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SuccessCount      *int32                                                           `json:"SuccessCount,omitempty" xml:"SuccessCount,omitempty"`
	SuccessExecutions []*CreateScheduledPreloadExecutionsResponseBodySuccessExecutions `json:"SuccessExecutions,omitempty" xml:"SuccessExecutions,omitempty" type:"Repeated"`
	TotalCount        *int32                                                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	if s.FailedExecutions != nil {
		for _, item := range s.FailedExecutions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SuccessExecutions != nil {
		for _, item := range s.SuccessExecutions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateScheduledPreloadExecutionsResponseBodyFailedExecutions struct {
	AliUid    *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
	AliUid    *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
