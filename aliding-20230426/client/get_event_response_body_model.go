// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttendees(v []*GetEventResponseBodyAttendees) *GetEventResponseBody
	GetAttendees() []*GetEventResponseBodyAttendees
	SetCategories(v []*GetEventResponseBodyCategories) *GetEventResponseBody
	GetCategories() []*GetEventResponseBodyCategories
	SetCreateTime(v string) *GetEventResponseBody
	GetCreateTime() *string
	SetDescription(v string) *GetEventResponseBody
	GetDescription() *string
	SetEnd(v *GetEventResponseBodyEnd) *GetEventResponseBody
	GetEnd() *GetEventResponseBodyEnd
	SetExtendedProperties(v *GetEventResponseBodyExtendedProperties) *GetEventResponseBody
	GetExtendedProperties() *GetEventResponseBodyExtendedProperties
	SetId(v string) *GetEventResponseBody
	GetId() *string
	SetIsAllDay(v bool) *GetEventResponseBody
	GetIsAllDay() *bool
	SetLocation(v *GetEventResponseBodyLocation) *GetEventResponseBody
	GetLocation() *GetEventResponseBodyLocation
	SetMeetingRooms(v []*GetEventResponseBodyMeetingRooms) *GetEventResponseBody
	GetMeetingRooms() []*GetEventResponseBodyMeetingRooms
	SetOnlineMeetingInfo(v *GetEventResponseBodyOnlineMeetingInfo) *GetEventResponseBody
	GetOnlineMeetingInfo() *GetEventResponseBodyOnlineMeetingInfo
	SetOrganizer(v *GetEventResponseBodyOrganizer) *GetEventResponseBody
	GetOrganizer() *GetEventResponseBodyOrganizer
	SetOriginStart(v *GetEventResponseBodyOriginStart) *GetEventResponseBody
	GetOriginStart() *GetEventResponseBodyOriginStart
	SetRecurrence(v *GetEventResponseBodyRecurrence) *GetEventResponseBody
	GetRecurrence() *GetEventResponseBodyRecurrence
	SetReminders(v []*GetEventResponseBodyReminders) *GetEventResponseBody
	GetReminders() []*GetEventResponseBodyReminders
	SetRequestId(v string) *GetEventResponseBody
	GetRequestId() *string
	SetRichTextDescription(v *GetEventResponseBodyRichTextDescription) *GetEventResponseBody
	GetRichTextDescription() *GetEventResponseBodyRichTextDescription
	SetSeriesMasterId(v string) *GetEventResponseBody
	GetSeriesMasterId() *string
	SetStart(v *GetEventResponseBodyStart) *GetEventResponseBody
	GetStart() *GetEventResponseBodyStart
	SetStatus(v string) *GetEventResponseBody
	GetStatus() *string
	SetSummary(v string) *GetEventResponseBody
	GetSummary() *string
	SetUpdateTime(v string) *GetEventResponseBody
	GetUpdateTime() *string
}

type GetEventResponseBody struct {
	Attendees  []*GetEventResponseBodyAttendees  `json:"attendees,omitempty" xml:"attendees,omitempty" type:"Repeated"`
	Categories []*GetEventResponseBodyCategories `json:"categories,omitempty" xml:"categories,omitempty" type:"Repeated"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	Description        *string                                 `json:"description,omitempty" xml:"description,omitempty"`
	End                *GetEventResponseBodyEnd                `json:"end,omitempty" xml:"end,omitempty" type:"Struct"`
	ExtendedProperties *GetEventResponseBodyExtendedProperties `json:"extendedProperties,omitempty" xml:"extendedProperties,omitempty" type:"Struct"`
	// example:
	//
	// iiiP35sJxxxxPRKgiEiF
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// true
	IsAllDay          *bool                                  `json:"isAllDay,omitempty" xml:"isAllDay,omitempty"`
	Location          *GetEventResponseBodyLocation          `json:"location,omitempty" xml:"location,omitempty" type:"Struct"`
	MeetingRooms      []*GetEventResponseBodyMeetingRooms    `json:"meetingRooms,omitempty" xml:"meetingRooms,omitempty" type:"Repeated"`
	OnlineMeetingInfo *GetEventResponseBodyOnlineMeetingInfo `json:"onlineMeetingInfo,omitempty" xml:"onlineMeetingInfo,omitempty" type:"Struct"`
	Organizer         *GetEventResponseBodyOrganizer         `json:"organizer,omitempty" xml:"organizer,omitempty" type:"Struct"`
	OriginStart       *GetEventResponseBodyOriginStart       `json:"originStart,omitempty" xml:"originStart,omitempty" type:"Struct"`
	Recurrence        *GetEventResponseBodyRecurrence        `json:"recurrence,omitempty" xml:"recurrence,omitempty" type:"Struct"`
	Reminders         []*GetEventResponseBodyReminders       `json:"reminders,omitempty" xml:"reminders,omitempty" type:"Repeated"`
	// requestId
	//
	// example:
	//
	// requestId
	RequestId           *string                                  `json:"requestId,omitempty" xml:"requestId,omitempty"`
	RichTextDescription *GetEventResponseBodyRichTextDescription `json:"richTextDescription,omitempty" xml:"richTextDescription,omitempty" type:"Struct"`
	// example:
	//
	// cnNTbW1YbxxxxvdlQrQT09
	SeriesMasterId *string                    `json:"seriesMasterId,omitempty" xml:"seriesMasterId,omitempty"`
	Start          *GetEventResponseBodyStart `json:"start,omitempty" xml:"start,omitempty" type:"Struct"`
	// example:
	//
	// confirmed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// test event
	Summary *string `json:"summary,omitempty" xml:"summary,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	UpdateTime *string `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s GetEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBody) GoString() string {
	return s.String()
}

func (s *GetEventResponseBody) GetAttendees() []*GetEventResponseBodyAttendees {
	return s.Attendees
}

func (s *GetEventResponseBody) GetCategories() []*GetEventResponseBodyCategories {
	return s.Categories
}

func (s *GetEventResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetEventResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetEventResponseBody) GetEnd() *GetEventResponseBodyEnd {
	return s.End
}

func (s *GetEventResponseBody) GetExtendedProperties() *GetEventResponseBodyExtendedProperties {
	return s.ExtendedProperties
}

func (s *GetEventResponseBody) GetId() *string {
	return s.Id
}

func (s *GetEventResponseBody) GetIsAllDay() *bool {
	return s.IsAllDay
}

func (s *GetEventResponseBody) GetLocation() *GetEventResponseBodyLocation {
	return s.Location
}

func (s *GetEventResponseBody) GetMeetingRooms() []*GetEventResponseBodyMeetingRooms {
	return s.MeetingRooms
}

func (s *GetEventResponseBody) GetOnlineMeetingInfo() *GetEventResponseBodyOnlineMeetingInfo {
	return s.OnlineMeetingInfo
}

func (s *GetEventResponseBody) GetOrganizer() *GetEventResponseBodyOrganizer {
	return s.Organizer
}

func (s *GetEventResponseBody) GetOriginStart() *GetEventResponseBodyOriginStart {
	return s.OriginStart
}

func (s *GetEventResponseBody) GetRecurrence() *GetEventResponseBodyRecurrence {
	return s.Recurrence
}

func (s *GetEventResponseBody) GetReminders() []*GetEventResponseBodyReminders {
	return s.Reminders
}

func (s *GetEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEventResponseBody) GetRichTextDescription() *GetEventResponseBodyRichTextDescription {
	return s.RichTextDescription
}

func (s *GetEventResponseBody) GetSeriesMasterId() *string {
	return s.SeriesMasterId
}

func (s *GetEventResponseBody) GetStart() *GetEventResponseBodyStart {
	return s.Start
}

func (s *GetEventResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetEventResponseBody) GetSummary() *string {
	return s.Summary
}

func (s *GetEventResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetEventResponseBody) SetAttendees(v []*GetEventResponseBodyAttendees) *GetEventResponseBody {
	s.Attendees = v
	return s
}

func (s *GetEventResponseBody) SetCategories(v []*GetEventResponseBodyCategories) *GetEventResponseBody {
	s.Categories = v
	return s
}

func (s *GetEventResponseBody) SetCreateTime(v string) *GetEventResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetEventResponseBody) SetDescription(v string) *GetEventResponseBody {
	s.Description = &v
	return s
}

func (s *GetEventResponseBody) SetEnd(v *GetEventResponseBodyEnd) *GetEventResponseBody {
	s.End = v
	return s
}

func (s *GetEventResponseBody) SetExtendedProperties(v *GetEventResponseBodyExtendedProperties) *GetEventResponseBody {
	s.ExtendedProperties = v
	return s
}

func (s *GetEventResponseBody) SetId(v string) *GetEventResponseBody {
	s.Id = &v
	return s
}

func (s *GetEventResponseBody) SetIsAllDay(v bool) *GetEventResponseBody {
	s.IsAllDay = &v
	return s
}

func (s *GetEventResponseBody) SetLocation(v *GetEventResponseBodyLocation) *GetEventResponseBody {
	s.Location = v
	return s
}

func (s *GetEventResponseBody) SetMeetingRooms(v []*GetEventResponseBodyMeetingRooms) *GetEventResponseBody {
	s.MeetingRooms = v
	return s
}

func (s *GetEventResponseBody) SetOnlineMeetingInfo(v *GetEventResponseBodyOnlineMeetingInfo) *GetEventResponseBody {
	s.OnlineMeetingInfo = v
	return s
}

func (s *GetEventResponseBody) SetOrganizer(v *GetEventResponseBodyOrganizer) *GetEventResponseBody {
	s.Organizer = v
	return s
}

func (s *GetEventResponseBody) SetOriginStart(v *GetEventResponseBodyOriginStart) *GetEventResponseBody {
	s.OriginStart = v
	return s
}

func (s *GetEventResponseBody) SetRecurrence(v *GetEventResponseBodyRecurrence) *GetEventResponseBody {
	s.Recurrence = v
	return s
}

func (s *GetEventResponseBody) SetReminders(v []*GetEventResponseBodyReminders) *GetEventResponseBody {
	s.Reminders = v
	return s
}

func (s *GetEventResponseBody) SetRequestId(v string) *GetEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEventResponseBody) SetRichTextDescription(v *GetEventResponseBodyRichTextDescription) *GetEventResponseBody {
	s.RichTextDescription = v
	return s
}

func (s *GetEventResponseBody) SetSeriesMasterId(v string) *GetEventResponseBody {
	s.SeriesMasterId = &v
	return s
}

func (s *GetEventResponseBody) SetStart(v *GetEventResponseBodyStart) *GetEventResponseBody {
	s.Start = v
	return s
}

func (s *GetEventResponseBody) SetStatus(v string) *GetEventResponseBody {
	s.Status = &v
	return s
}

func (s *GetEventResponseBody) SetSummary(v string) *GetEventResponseBody {
	s.Summary = &v
	return s
}

func (s *GetEventResponseBody) SetUpdateTime(v string) *GetEventResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *GetEventResponseBody) Validate() error {
	if s.Attendees != nil {
		for _, item := range s.Attendees {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Categories != nil {
		for _, item := range s.Categories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.End != nil {
		if err := s.End.Validate(); err != nil {
			return err
		}
	}
	if s.ExtendedProperties != nil {
		if err := s.ExtendedProperties.Validate(); err != nil {
			return err
		}
	}
	if s.Location != nil {
		if err := s.Location.Validate(); err != nil {
			return err
		}
	}
	if s.MeetingRooms != nil {
		for _, item := range s.MeetingRooms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OnlineMeetingInfo != nil {
		if err := s.OnlineMeetingInfo.Validate(); err != nil {
			return err
		}
	}
	if s.Organizer != nil {
		if err := s.Organizer.Validate(); err != nil {
			return err
		}
	}
	if s.OriginStart != nil {
		if err := s.OriginStart.Validate(); err != nil {
			return err
		}
	}
	if s.Recurrence != nil {
		if err := s.Recurrence.Validate(); err != nil {
			return err
		}
	}
	if s.Reminders != nil {
		for _, item := range s.Reminders {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RichTextDescription != nil {
		if err := s.RichTextDescription.Validate(); err != nil {
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

type GetEventResponseBodyAttendees struct {
	// example:
	//
	// jack
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsOptional *bool `json:"IsOptional,omitempty" xml:"IsOptional,omitempty"`
	// example:
	//
	// accepted
	ResponseStatus *string `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty"`
	// example:
	//
	// true
	Self *bool `json:"Self,omitempty" xml:"Self,omitempty"`
}

func (s GetEventResponseBodyAttendees) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyAttendees) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyAttendees) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetEventResponseBodyAttendees) GetId() *string {
	return s.Id
}

func (s *GetEventResponseBodyAttendees) GetIsOptional() *bool {
	return s.IsOptional
}

func (s *GetEventResponseBodyAttendees) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *GetEventResponseBodyAttendees) GetSelf() *bool {
	return s.Self
}

func (s *GetEventResponseBodyAttendees) SetDisplayName(v string) *GetEventResponseBodyAttendees {
	s.DisplayName = &v
	return s
}

func (s *GetEventResponseBodyAttendees) SetId(v string) *GetEventResponseBodyAttendees {
	s.Id = &v
	return s
}

func (s *GetEventResponseBodyAttendees) SetIsOptional(v bool) *GetEventResponseBodyAttendees {
	s.IsOptional = &v
	return s
}

func (s *GetEventResponseBodyAttendees) SetResponseStatus(v string) *GetEventResponseBodyAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *GetEventResponseBodyAttendees) SetSelf(v bool) *GetEventResponseBodyAttendees {
	s.Self = &v
	return s
}

func (s *GetEventResponseBodyAttendees) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyCategories struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
}

func (s GetEventResponseBodyCategories) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyCategories) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyCategories) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetEventResponseBodyCategories) SetDisplayName(v string) *GetEventResponseBodyCategories {
	s.DisplayName = &v
	return s
}

func (s *GetEventResponseBodyCategories) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyEnd struct {
	// example:
	//
	// 2020-01-01
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s GetEventResponseBodyEnd) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyEnd) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyEnd) GetDate() *string {
	return s.Date
}

func (s *GetEventResponseBodyEnd) GetDateTime() *string {
	return s.DateTime
}

func (s *GetEventResponseBodyEnd) GetTimeZone() *string {
	return s.TimeZone
}

func (s *GetEventResponseBodyEnd) SetDate(v string) *GetEventResponseBodyEnd {
	s.Date = &v
	return s
}

func (s *GetEventResponseBodyEnd) SetDateTime(v string) *GetEventResponseBodyEnd {
	s.DateTime = &v
	return s
}

func (s *GetEventResponseBodyEnd) SetTimeZone(v string) *GetEventResponseBodyEnd {
	s.TimeZone = &v
	return s
}

func (s *GetEventResponseBodyEnd) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyExtendedProperties struct {
	SharedProperties *GetEventResponseBodyExtendedPropertiesSharedProperties `json:"SharedProperties,omitempty" xml:"SharedProperties,omitempty" type:"Struct"`
}

func (s GetEventResponseBodyExtendedProperties) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyExtendedProperties) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyExtendedProperties) GetSharedProperties() *GetEventResponseBodyExtendedPropertiesSharedProperties {
	return s.SharedProperties
}

func (s *GetEventResponseBodyExtendedProperties) SetSharedProperties(v *GetEventResponseBodyExtendedPropertiesSharedProperties) *GetEventResponseBodyExtendedProperties {
	s.SharedProperties = v
	return s
}

func (s *GetEventResponseBodyExtendedProperties) Validate() error {
	if s.SharedProperties != nil {
		if err := s.SharedProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEventResponseBodyExtendedPropertiesSharedProperties struct {
	// example:
	//
	// dingd*****1231231
	BelongCorpId *string `json:"BelongCorpId,omitempty" xml:"BelongCorpId,omitempty"`
	// example:
	//
	// zxcvasdfvb123====
	SourceOpenCid *string `json:"SourceOpenCid,omitempty" xml:"SourceOpenCid,omitempty"`
}

func (s GetEventResponseBodyExtendedPropertiesSharedProperties) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyExtendedPropertiesSharedProperties) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyExtendedPropertiesSharedProperties) GetBelongCorpId() *string {
	return s.BelongCorpId
}

func (s *GetEventResponseBodyExtendedPropertiesSharedProperties) GetSourceOpenCid() *string {
	return s.SourceOpenCid
}

func (s *GetEventResponseBodyExtendedPropertiesSharedProperties) SetBelongCorpId(v string) *GetEventResponseBodyExtendedPropertiesSharedProperties {
	s.BelongCorpId = &v
	return s
}

func (s *GetEventResponseBodyExtendedPropertiesSharedProperties) SetSourceOpenCid(v string) *GetEventResponseBodyExtendedPropertiesSharedProperties {
	s.SourceOpenCid = &v
	return s
}

func (s *GetEventResponseBodyExtendedPropertiesSharedProperties) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyLocation struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// --
	MeetingRooms []*string `json:"MeetingRooms,omitempty" xml:"MeetingRooms,omitempty" type:"Repeated"`
}

func (s GetEventResponseBodyLocation) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyLocation) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyLocation) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetEventResponseBodyLocation) GetMeetingRooms() []*string {
	return s.MeetingRooms
}

func (s *GetEventResponseBodyLocation) SetDisplayName(v string) *GetEventResponseBodyLocation {
	s.DisplayName = &v
	return s
}

func (s *GetEventResponseBodyLocation) SetMeetingRooms(v []*string) *GetEventResponseBodyLocation {
	s.MeetingRooms = v
	return s
}

func (s *GetEventResponseBodyLocation) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyMeetingRooms struct {
	// example:
	//
	// room 1-2-3
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// accepted
	ResponseStatus *string `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty"`
	// example:
	//
	// c10315a8b4e740a317813ab6fxxxxxx
	RoomId *string `json:"RoomId,omitempty" xml:"RoomId,omitempty"`
}

func (s GetEventResponseBodyMeetingRooms) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyMeetingRooms) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyMeetingRooms) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetEventResponseBodyMeetingRooms) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *GetEventResponseBodyMeetingRooms) GetRoomId() *string {
	return s.RoomId
}

func (s *GetEventResponseBodyMeetingRooms) SetDisplayName(v string) *GetEventResponseBodyMeetingRooms {
	s.DisplayName = &v
	return s
}

func (s *GetEventResponseBodyMeetingRooms) SetResponseStatus(v string) *GetEventResponseBodyMeetingRooms {
	s.ResponseStatus = &v
	return s
}

func (s *GetEventResponseBodyMeetingRooms) SetRoomId(v string) *GetEventResponseBodyMeetingRooms {
	s.RoomId = &v
	return s
}

func (s *GetEventResponseBodyMeetingRooms) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyOnlineMeetingInfo struct {
	// example:
	//
	// 5c4df21d-xxxx-a6db402b9f3a
	ConferenceId *string                `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	ExtraInfo    map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// dingtalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// dingtalk://dingtalkclient/page/videoxxxxalendar?confId=5c4df21d-xxxx9f3f&calendarId=127xxxx124
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetEventResponseBodyOnlineMeetingInfo) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyOnlineMeetingInfo) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyOnlineMeetingInfo) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *GetEventResponseBodyOnlineMeetingInfo) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *GetEventResponseBodyOnlineMeetingInfo) GetType() *string {
	return s.Type
}

func (s *GetEventResponseBodyOnlineMeetingInfo) GetUrl() *string {
	return s.Url
}

func (s *GetEventResponseBodyOnlineMeetingInfo) SetConferenceId(v string) *GetEventResponseBodyOnlineMeetingInfo {
	s.ConferenceId = &v
	return s
}

func (s *GetEventResponseBodyOnlineMeetingInfo) SetExtraInfo(v map[string]interface{}) *GetEventResponseBodyOnlineMeetingInfo {
	s.ExtraInfo = v
	return s
}

func (s *GetEventResponseBodyOnlineMeetingInfo) SetType(v string) *GetEventResponseBodyOnlineMeetingInfo {
	s.Type = &v
	return s
}

func (s *GetEventResponseBodyOnlineMeetingInfo) SetUrl(v string) *GetEventResponseBodyOnlineMeetingInfo {
	s.Url = &v
	return s
}

func (s *GetEventResponseBodyOnlineMeetingInfo) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyOrganizer struct {
	// example:
	//
	// tony
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 123456
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// accepted
	ResponseStatus *string `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty"`
	// example:
	//
	// true
	Self *bool `json:"Self,omitempty" xml:"Self,omitempty"`
}

func (s GetEventResponseBodyOrganizer) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyOrganizer) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyOrganizer) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetEventResponseBodyOrganizer) GetId() *string {
	return s.Id
}

func (s *GetEventResponseBodyOrganizer) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *GetEventResponseBodyOrganizer) GetSelf() *bool {
	return s.Self
}

func (s *GetEventResponseBodyOrganizer) SetDisplayName(v string) *GetEventResponseBodyOrganizer {
	s.DisplayName = &v
	return s
}

func (s *GetEventResponseBodyOrganizer) SetId(v string) *GetEventResponseBodyOrganizer {
	s.Id = &v
	return s
}

func (s *GetEventResponseBodyOrganizer) SetResponseStatus(v string) *GetEventResponseBodyOrganizer {
	s.ResponseStatus = &v
	return s
}

func (s *GetEventResponseBodyOrganizer) SetSelf(v bool) *GetEventResponseBodyOrganizer {
	s.Self = &v
	return s
}

func (s *GetEventResponseBodyOrganizer) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyOriginStart struct {
	// example:
	//
	// 2023-01-01T00:00:00Z
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
}

func (s GetEventResponseBodyOriginStart) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyOriginStart) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyOriginStart) GetDateTime() *string {
	return s.DateTime
}

func (s *GetEventResponseBodyOriginStart) SetDateTime(v string) *GetEventResponseBodyOriginStart {
	s.DateTime = &v
	return s
}

func (s *GetEventResponseBodyOriginStart) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyRecurrence struct {
	Pattern *GetEventResponseBodyRecurrencePattern `json:"Pattern,omitempty" xml:"Pattern,omitempty" type:"Struct"`
	Range   *GetEventResponseBodyRecurrenceRange   `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
}

func (s GetEventResponseBodyRecurrence) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyRecurrence) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyRecurrence) GetPattern() *GetEventResponseBodyRecurrencePattern {
	return s.Pattern
}

func (s *GetEventResponseBodyRecurrence) GetRange() *GetEventResponseBodyRecurrenceRange {
	return s.Range
}

func (s *GetEventResponseBodyRecurrence) SetPattern(v *GetEventResponseBodyRecurrencePattern) *GetEventResponseBodyRecurrence {
	s.Pattern = v
	return s
}

func (s *GetEventResponseBodyRecurrence) SetRange(v *GetEventResponseBodyRecurrenceRange) *GetEventResponseBodyRecurrence {
	s.Range = v
	return s
}

func (s *GetEventResponseBodyRecurrence) Validate() error {
	if s.Pattern != nil {
		if err := s.Pattern.Validate(); err != nil {
			return err
		}
	}
	if s.Range != nil {
		if err := s.Range.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEventResponseBodyRecurrencePattern struct {
	// example:
	//
	// 14
	DayOfMonth *int32 `json:"DayOfMonth,omitempty" xml:"DayOfMonth,omitempty"`
	// example:
	//
	// monday
	DaysOfWeek *string `json:"DaysOfWeek,omitempty" xml:"DaysOfWeek,omitempty"`
	// example:
	//
	// first
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// example:
	//
	// 1
	Interval *int32 `json:"Interval,omitempty" xml:"Interval,omitempty"`
	// example:
	//
	// daily
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetEventResponseBodyRecurrencePattern) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyRecurrencePattern) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyRecurrencePattern) GetDayOfMonth() *int32 {
	return s.DayOfMonth
}

func (s *GetEventResponseBodyRecurrencePattern) GetDaysOfWeek() *string {
	return s.DaysOfWeek
}

func (s *GetEventResponseBodyRecurrencePattern) GetIndex() *string {
	return s.Index
}

func (s *GetEventResponseBodyRecurrencePattern) GetInterval() *int32 {
	return s.Interval
}

func (s *GetEventResponseBodyRecurrencePattern) GetType() *string {
	return s.Type
}

func (s *GetEventResponseBodyRecurrencePattern) SetDayOfMonth(v int32) *GetEventResponseBodyRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *GetEventResponseBodyRecurrencePattern) SetDaysOfWeek(v string) *GetEventResponseBodyRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *GetEventResponseBodyRecurrencePattern) SetIndex(v string) *GetEventResponseBodyRecurrencePattern {
	s.Index = &v
	return s
}

func (s *GetEventResponseBodyRecurrencePattern) SetInterval(v int32) *GetEventResponseBodyRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *GetEventResponseBodyRecurrencePattern) SetType(v string) *GetEventResponseBodyRecurrencePattern {
	s.Type = &v
	return s
}

func (s *GetEventResponseBodyRecurrencePattern) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyRecurrenceRange struct {
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// example:
	//
	// 5
	NumberOfOccurrences *int32 `json:"NumberOfOccurrences,omitempty" xml:"NumberOfOccurrences,omitempty"`
	// example:
	//
	// noEnd
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetEventResponseBodyRecurrenceRange) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyRecurrenceRange) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyRecurrenceRange) GetEndDate() *string {
	return s.EndDate
}

func (s *GetEventResponseBodyRecurrenceRange) GetNumberOfOccurrences() *int32 {
	return s.NumberOfOccurrences
}

func (s *GetEventResponseBodyRecurrenceRange) GetType() *string {
	return s.Type
}

func (s *GetEventResponseBodyRecurrenceRange) SetEndDate(v string) *GetEventResponseBodyRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *GetEventResponseBodyRecurrenceRange) SetNumberOfOccurrences(v int32) *GetEventResponseBodyRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *GetEventResponseBodyRecurrenceRange) SetType(v string) *GetEventResponseBodyRecurrenceRange {
	s.Type = &v
	return s
}

func (s *GetEventResponseBodyRecurrenceRange) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyReminders struct {
	// example:
	//
	// dingtalk
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// example:
	//
	// 15
	Minutes *string `json:"Minutes,omitempty" xml:"Minutes,omitempty"`
}

func (s GetEventResponseBodyReminders) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyReminders) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyReminders) GetMethod() *string {
	return s.Method
}

func (s *GetEventResponseBodyReminders) GetMinutes() *string {
	return s.Minutes
}

func (s *GetEventResponseBodyReminders) SetMethod(v string) *GetEventResponseBodyReminders {
	s.Method = &v
	return s
}

func (s *GetEventResponseBodyReminders) SetMinutes(v string) *GetEventResponseBodyReminders {
	s.Minutes = &v
	return s
}

func (s *GetEventResponseBodyReminders) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyRichTextDescription struct {
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s GetEventResponseBodyRichTextDescription) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyRichTextDescription) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyRichTextDescription) GetText() *string {
	return s.Text
}

func (s *GetEventResponseBodyRichTextDescription) SetText(v string) *GetEventResponseBodyRichTextDescription {
	s.Text = &v
	return s
}

func (s *GetEventResponseBodyRichTextDescription) Validate() error {
	return dara.Validate(s)
}

type GetEventResponseBodyStart struct {
	// example:
	//
	// 2020-01-01
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s GetEventResponseBodyStart) String() string {
	return dara.Prettify(s)
}

func (s GetEventResponseBodyStart) GoString() string {
	return s.String()
}

func (s *GetEventResponseBodyStart) GetDate() *string {
	return s.Date
}

func (s *GetEventResponseBodyStart) GetDateTime() *string {
	return s.DateTime
}

func (s *GetEventResponseBodyStart) GetTimeZone() *string {
	return s.TimeZone
}

func (s *GetEventResponseBodyStart) SetDate(v string) *GetEventResponseBodyStart {
	s.Date = &v
	return s
}

func (s *GetEventResponseBodyStart) SetDateTime(v string) *GetEventResponseBodyStart {
	s.DateTime = &v
	return s
}

func (s *GetEventResponseBodyStart) SetTimeZone(v string) *GetEventResponseBodyStart {
	s.TimeZone = &v
	return s
}

func (s *GetEventResponseBodyStart) Validate() error {
	return dara.Validate(s)
}
