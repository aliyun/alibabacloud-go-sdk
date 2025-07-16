// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryScheduleConferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *QueryScheduleConferenceResponseBody
	GetEndTime() *int64
	SetPhones(v []*string) *QueryScheduleConferenceResponseBody
	GetPhones() []*string
	SetRequestId(v string) *QueryScheduleConferenceResponseBody
	GetRequestId() *string
	SetRoomCode(v string) *QueryScheduleConferenceResponseBody
	GetRoomCode() *string
	SetScheduleConferenceId(v string) *QueryScheduleConferenceResponseBody
	GetScheduleConferenceId() *string
	SetStartTime(v int64) *QueryScheduleConferenceResponseBody
	GetStartTime() *int64
	SetTitle(v string) *QueryScheduleConferenceResponseBody
	GetTitle() *string
	SetUrl(v string) *QueryScheduleConferenceResponseBody
	GetUrl() *string
}

type QueryScheduleConferenceResponseBody struct {
	// example:
	//
	// 1687928400000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// [ "+86123xxxx" ]
	Phones []*string `json:"phones,omitempty" xml:"phones,omitempty" type:"Repeated"`
	// example:
	//
	// xxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 838 722 xxxxx
	RoomCode *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	// example:
	//
	// 2a489c68-xxxx-xxxx-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
	// example:
	//
	// 1687924800000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// example:
	//
	// 预约会议标题
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// example:
	//
	// https://meeting.dingtalk.com/j/Bsbp3ixxxxxUyJJ9
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s QueryScheduleConferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryScheduleConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryScheduleConferenceResponseBody) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryScheduleConferenceResponseBody) GetPhones() []*string {
	return s.Phones
}

func (s *QueryScheduleConferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryScheduleConferenceResponseBody) GetRoomCode() *string {
	return s.RoomCode
}

func (s *QueryScheduleConferenceResponseBody) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *QueryScheduleConferenceResponseBody) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryScheduleConferenceResponseBody) GetTitle() *string {
	return s.Title
}

func (s *QueryScheduleConferenceResponseBody) GetUrl() *string {
	return s.Url
}

func (s *QueryScheduleConferenceResponseBody) SetEndTime(v int64) *QueryScheduleConferenceResponseBody {
	s.EndTime = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetPhones(v []*string) *QueryScheduleConferenceResponseBody {
	s.Phones = v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetRequestId(v string) *QueryScheduleConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetRoomCode(v string) *QueryScheduleConferenceResponseBody {
	s.RoomCode = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetScheduleConferenceId(v string) *QueryScheduleConferenceResponseBody {
	s.ScheduleConferenceId = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetStartTime(v int64) *QueryScheduleConferenceResponseBody {
	s.StartTime = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetTitle(v string) *QueryScheduleConferenceResponseBody {
	s.Title = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) SetUrl(v string) *QueryScheduleConferenceResponseBody {
	s.Url = &v
	return s
}

func (s *QueryScheduleConferenceResponseBody) Validate() error {
	return dara.Validate(s)
}
