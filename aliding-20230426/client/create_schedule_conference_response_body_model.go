// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduleConferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPhones(v []*string) *CreateScheduleConferenceResponseBody
	GetPhones() []*string
	SetRequestId(v string) *CreateScheduleConferenceResponseBody
	GetRequestId() *string
	SetRoomCode(v string) *CreateScheduleConferenceResponseBody
	GetRoomCode() *string
	SetScheduleConferenceId(v string) *CreateScheduleConferenceResponseBody
	GetScheduleConferenceId() *string
	SetUrl(v string) *CreateScheduleConferenceResponseBody
	GetUrl() *string
}

type CreateScheduleConferenceResponseBody struct {
	// example:
	//
	// +861234567
	Phones []*string `json:"phones,omitempty" xml:"phones,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// 1234567
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 83150xxxxxx
	RoomCode *string `json:"roomCode,omitempty" xml:"roomCode,omitempty"`
	// example:
	//
	// 5c7c9bb1-b256-4dc5-xxxx-xxxxxxxxxxxx
	ScheduleConferenceId *string `json:"scheduleConferenceId,omitempty" xml:"scheduleConferenceId,omitempty"`
	// example:
	//
	// https://meeting.dingtalk.com/j/knvMq1ixxxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CreateScheduleConferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduleConferenceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScheduleConferenceResponseBody) GetPhones() []*string {
	return s.Phones
}

func (s *CreateScheduleConferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScheduleConferenceResponseBody) GetRoomCode() *string {
	return s.RoomCode
}

func (s *CreateScheduleConferenceResponseBody) GetScheduleConferenceId() *string {
	return s.ScheduleConferenceId
}

func (s *CreateScheduleConferenceResponseBody) GetUrl() *string {
	return s.Url
}

func (s *CreateScheduleConferenceResponseBody) SetPhones(v []*string) *CreateScheduleConferenceResponseBody {
	s.Phones = v
	return s
}

func (s *CreateScheduleConferenceResponseBody) SetRequestId(v string) *CreateScheduleConferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScheduleConferenceResponseBody) SetRoomCode(v string) *CreateScheduleConferenceResponseBody {
	s.RoomCode = &v
	return s
}

func (s *CreateScheduleConferenceResponseBody) SetScheduleConferenceId(v string) *CreateScheduleConferenceResponseBody {
	s.ScheduleConferenceId = &v
	return s
}

func (s *CreateScheduleConferenceResponseBody) SetUrl(v string) *CreateScheduleConferenceResponseBody {
	s.Url = &v
	return s
}

func (s *CreateScheduleConferenceResponseBody) Validate() error {
	return dara.Validate(s)
}
