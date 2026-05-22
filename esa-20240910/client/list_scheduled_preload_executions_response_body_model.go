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
	Executions []*ListScheduledPreloadExecutionsResponseBodyExecutions `json:"Executions,omitempty" xml:"Executions,omitempty" type:"Repeated"`
	RequestId  *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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
	AliUid    *string `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Id        *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Interval  *int32  `json:"Interval,omitempty" xml:"Interval,omitempty"`
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	SliceLen  *int32  `json:"SliceLen,omitempty" xml:"SliceLen,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
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
