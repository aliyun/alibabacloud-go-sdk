// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMeetingRoomsScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMeetingRoomsScheduleResponseBody
	GetRequestId() *string
	SetScheduleInformation(v []*GetMeetingRoomsScheduleResponseBodyScheduleInformation) *GetMeetingRoomsScheduleResponseBody
	GetScheduleInformation() []*GetMeetingRoomsScheduleResponseBodyScheduleInformation
}

type GetMeetingRoomsScheduleResponseBody struct {
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId           *string                                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScheduleInformation []*GetMeetingRoomsScheduleResponseBodyScheduleInformation `json:"scheduleInformation,omitempty" xml:"scheduleInformation,omitempty" type:"Repeated"`
}

func (s GetMeetingRoomsScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMeetingRoomsScheduleResponseBody) GetScheduleInformation() []*GetMeetingRoomsScheduleResponseBodyScheduleInformation {
	return s.ScheduleInformation
}

func (s *GetMeetingRoomsScheduleResponseBody) SetRequestId(v string) *GetMeetingRoomsScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBody) SetScheduleInformation(v []*GetMeetingRoomsScheduleResponseBodyScheduleInformation) *GetMeetingRoomsScheduleResponseBody {
	s.ScheduleInformation = v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBody) Validate() error {
	if s.ScheduleInformation != nil {
		for _, item := range s.ScheduleInformation {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMeetingRoomsScheduleResponseBodyScheduleInformation struct {
	// example:
	//
	// 无权限
	Error *string `json:"Error,omitempty" xml:"Error,omitempty"`
	// example:
	//
	// 4002f89xxxxx
	RoomId        *string                                                                `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
	ScheduleItems []*GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems `json:"ScheduleItems,omitempty" xml:"ScheduleItems,omitempty" type:"Repeated"`
}

func (s GetMeetingRoomsScheduleResponseBodyScheduleInformation) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleResponseBodyScheduleInformation) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformation) GetError() *string {
	return s.Error
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformation) GetRoomId() *string {
	return s.RoomId
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformation) GetScheduleItems() []*GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems {
	return s.ScheduleItems
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformation) SetError(v string) *GetMeetingRoomsScheduleResponseBodyScheduleInformation {
	s.Error = &v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformation) SetRoomId(v string) *GetMeetingRoomsScheduleResponseBodyScheduleInformation {
	s.RoomId = &v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformation) SetScheduleItems(v []*GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) *GetMeetingRoomsScheduleResponseBodyScheduleInformation {
	s.ScheduleItems = v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformation) Validate() error {
	if s.ScheduleItems != nil {
		for _, item := range s.ScheduleItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems struct {
	End *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd `json:"End,omitempty" xml:"End,omitempty" type:"Struct"`
	// example:
	//
	// UzZvxxxxx
	EventId   *string                                                                       `json:"EventId,omitempty" xml:"EventId,omitempty"`
	Organizer *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsOrganizer `json:"Organizer,omitempty" xml:"Organizer,omitempty" type:"Struct"`
	Start     *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart     `json:"Start,omitempty" xml:"Start,omitempty" type:"Struct"`
	// example:
	//
	// BUSY
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) GetEnd() *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd {
	return s.End
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) GetEventId() *string {
	return s.EventId
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) GetOrganizer() *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsOrganizer {
	return s.Organizer
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) GetStart() *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart {
	return s.Start
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) GetStatus() *string {
	return s.Status
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) SetEnd(v *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd) *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems {
	s.End = v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) SetEventId(v string) *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems {
	s.EventId = &v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) SetOrganizer(v *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsOrganizer) *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems {
	s.Organizer = v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) SetStart(v *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart) *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems {
	s.Start = v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) SetStatus(v string) *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems {
	s.Status = &v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItems) Validate() error {
	if s.End != nil {
		if err := s.End.Validate(); err != nil {
			return err
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.Validate(); err != nil {
			return err
		}
	}
	if s.Start != nil {
		if err := s.Start.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd struct {
	// example:
	//
	// 2020-01-02T10:15:30+08:00
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd) GetDateTime() *string {
	return s.DateTime
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd) GetTimeZone() *string {
	return s.TimeZone
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd) SetDateTime(v string) *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd {
	s.DateTime = &v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd) SetTimeZone(v string) *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd {
	s.TimeZone = &v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsEnd) Validate() error {
	return dara.Validate(s)
}

type GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsOrganizer struct {
	// example:
	//
	// 012345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsOrganizer) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsOrganizer) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsOrganizer) GetId() *string {
	return s.Id
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsOrganizer) SetId(v string) *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsOrganizer {
	s.Id = &v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsOrganizer) Validate() error {
	return dara.Validate(s)
}

type GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart struct {
	// example:
	//
	// 2020-01-02T10:15:30+08:00
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart) String() string {
	return dara.Prettify(s)
}

func (s GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart) GoString() string {
	return s.String()
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart) GetDateTime() *string {
	return s.DateTime
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart) GetTimeZone() *string {
	return s.TimeZone
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart) SetDateTime(v string) *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart {
	s.DateTime = &v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart) SetTimeZone(v string) *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart {
	s.TimeZone = &v
	return s
}

func (s *GetMeetingRoomsScheduleResponseBodyScheduleInformationScheduleItemsStart) Validate() error {
	return dara.Validate(s)
}
