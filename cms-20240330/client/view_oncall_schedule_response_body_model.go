// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iViewOncallScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ViewOncallScheduleResponseBody
	GetRequestId() *string
	SetSchedule(v map[string][]*ScheduleValue) *ViewOncallScheduleResponseBody
	GetSchedule() map[string][]*ScheduleValue
}

type ViewOncallScheduleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The on-call calendar. The key is the shift name, and the value is a list of on-call shifts for the shift within the queried date range.
	//
	// example:
	//
	// {"testName":[{"contactId":"testId","contactName":"Zhang San","start":"2024-10-10 09:00","end":"2024-10-15 18:00"}]}
	Schedule map[string][]*ScheduleValue `json:"schedule,omitempty" xml:"schedule,omitempty"`
}

func (s ViewOncallScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ViewOncallScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *ViewOncallScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ViewOncallScheduleResponseBody) GetSchedule() map[string][]*ScheduleValue {
	return s.Schedule
}

func (s *ViewOncallScheduleResponseBody) SetRequestId(v string) *ViewOncallScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ViewOncallScheduleResponseBody) SetSchedule(v map[string][]*ScheduleValue) *ViewOncallScheduleResponseBody {
	s.Schedule = v
	return s
}

func (s *ViewOncallScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
