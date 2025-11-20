// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventsViewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*ListEventsViewResponseBodyEvents) *ListEventsViewResponseBody
	GetEvents() []*ListEventsViewResponseBodyEvents
	SetNextToken(v string) *ListEventsViewResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEventsViewResponseBody
	GetRequestId() *string
}

type ListEventsViewResponseBody struct {
	Events []*ListEventsViewResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// example:
	//
	// cnNTbWxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListEventsViewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBody) GetEvents() []*ListEventsViewResponseBodyEvents {
	return s.Events
}

func (s *ListEventsViewResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventsViewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventsViewResponseBody) SetEvents(v []*ListEventsViewResponseBodyEvents) *ListEventsViewResponseBody {
	s.Events = v
	return s
}

func (s *ListEventsViewResponseBody) SetNextToken(v string) *ListEventsViewResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEventsViewResponseBody) SetRequestId(v string) *ListEventsViewResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventsViewResponseBody) Validate() error {
	if s.Events != nil {
		for _, item := range s.Events {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEventsViewResponseBodyEvents struct {
	Attendees  []*ListEventsViewResponseBodyEventsAttendees  `json:"Attendees,omitempty" xml:"Attendees,omitempty" type:"Repeated"`
	Categories []*ListEventsViewResponseBodyEventsCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// something about this event
	Description        *string                                             `json:"Description,omitempty" xml:"Description,omitempty"`
	End                *ListEventsViewResponseBodyEventsEnd                `json:"End,omitempty" xml:"End,omitempty" type:"Struct"`
	ExtendedProperties *ListEventsViewResponseBodyEventsExtendedProperties `json:"ExtendedProperties,omitempty" xml:"ExtendedProperties,omitempty" type:"Struct"`
	// example:
	//
	// iiiP35xxx
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsAllDay            *bool                                                `json:"IsAllDay,omitempty" xml:"IsAllDay,omitempty"`
	Location            *ListEventsViewResponseBodyEventsLocation            `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	MeetingRooms        []*ListEventsViewResponseBodyEventsMeetingRooms      `json:"MeetingRooms,omitempty" xml:"MeetingRooms,omitempty" type:"Repeated"`
	OnlineMeetingInfo   *ListEventsViewResponseBodyEventsOnlineMeetingInfo   `json:"OnlineMeetingInfo,omitempty" xml:"OnlineMeetingInfo,omitempty" type:"Struct"`
	Organizer           *ListEventsViewResponseBodyEventsOrganizer           `json:"Organizer,omitempty" xml:"Organizer,omitempty" type:"Struct"`
	OriginStart         *ListEventsViewResponseBodyEventsOriginStart         `json:"OriginStart,omitempty" xml:"OriginStart,omitempty" type:"Struct"`
	Recurrence          *ListEventsViewResponseBodyEventsRecurrence          `json:"Recurrence,omitempty" xml:"Recurrence,omitempty" type:"Struct"`
	RichTextDescription *ListEventsViewResponseBodyEventsRichTextDescription `json:"RichTextDescription,omitempty" xml:"RichTextDescription,omitempty" type:"Struct"`
	// example:
	//
	// cnNTbxxx
	SeriesMasterId *string                                `json:"SeriesMasterId,omitempty" xml:"SeriesMasterId,omitempty"`
	Start          *ListEventsViewResponseBodyEventsStart `json:"Start,omitempty" xml:"Start,omitempty" type:"Struct"`
	// example:
	//
	// confirmed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// test event
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEventsViewResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEvents) GetAttendees() []*ListEventsViewResponseBodyEventsAttendees {
	return s.Attendees
}

func (s *ListEventsViewResponseBodyEvents) GetCategories() []*ListEventsViewResponseBodyEventsCategories {
	return s.Categories
}

func (s *ListEventsViewResponseBodyEvents) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEventsViewResponseBodyEvents) GetDescription() *string {
	return s.Description
}

func (s *ListEventsViewResponseBodyEvents) GetEnd() *ListEventsViewResponseBodyEventsEnd {
	return s.End
}

func (s *ListEventsViewResponseBodyEvents) GetExtendedProperties() *ListEventsViewResponseBodyEventsExtendedProperties {
	return s.ExtendedProperties
}

func (s *ListEventsViewResponseBodyEvents) GetId() *string {
	return s.Id
}

func (s *ListEventsViewResponseBodyEvents) GetIsAllDay() *bool {
	return s.IsAllDay
}

func (s *ListEventsViewResponseBodyEvents) GetLocation() *ListEventsViewResponseBodyEventsLocation {
	return s.Location
}

func (s *ListEventsViewResponseBodyEvents) GetMeetingRooms() []*ListEventsViewResponseBodyEventsMeetingRooms {
	return s.MeetingRooms
}

func (s *ListEventsViewResponseBodyEvents) GetOnlineMeetingInfo() *ListEventsViewResponseBodyEventsOnlineMeetingInfo {
	return s.OnlineMeetingInfo
}

func (s *ListEventsViewResponseBodyEvents) GetOrganizer() *ListEventsViewResponseBodyEventsOrganizer {
	return s.Organizer
}

func (s *ListEventsViewResponseBodyEvents) GetOriginStart() *ListEventsViewResponseBodyEventsOriginStart {
	return s.OriginStart
}

func (s *ListEventsViewResponseBodyEvents) GetRecurrence() *ListEventsViewResponseBodyEventsRecurrence {
	return s.Recurrence
}

func (s *ListEventsViewResponseBodyEvents) GetRichTextDescription() *ListEventsViewResponseBodyEventsRichTextDescription {
	return s.RichTextDescription
}

func (s *ListEventsViewResponseBodyEvents) GetSeriesMasterId() *string {
	return s.SeriesMasterId
}

func (s *ListEventsViewResponseBodyEvents) GetStart() *ListEventsViewResponseBodyEventsStart {
	return s.Start
}

func (s *ListEventsViewResponseBodyEvents) GetStatus() *string {
	return s.Status
}

func (s *ListEventsViewResponseBodyEvents) GetSummary() *string {
	return s.Summary
}

func (s *ListEventsViewResponseBodyEvents) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListEventsViewResponseBodyEvents) SetAttendees(v []*ListEventsViewResponseBodyEventsAttendees) *ListEventsViewResponseBodyEvents {
	s.Attendees = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetCategories(v []*ListEventsViewResponseBodyEventsCategories) *ListEventsViewResponseBodyEvents {
	s.Categories = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetCreateTime(v string) *ListEventsViewResponseBodyEvents {
	s.CreateTime = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetDescription(v string) *ListEventsViewResponseBodyEvents {
	s.Description = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetEnd(v *ListEventsViewResponseBodyEventsEnd) *ListEventsViewResponseBodyEvents {
	s.End = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetExtendedProperties(v *ListEventsViewResponseBodyEventsExtendedProperties) *ListEventsViewResponseBodyEvents {
	s.ExtendedProperties = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetId(v string) *ListEventsViewResponseBodyEvents {
	s.Id = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetIsAllDay(v bool) *ListEventsViewResponseBodyEvents {
	s.IsAllDay = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetLocation(v *ListEventsViewResponseBodyEventsLocation) *ListEventsViewResponseBodyEvents {
	s.Location = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetMeetingRooms(v []*ListEventsViewResponseBodyEventsMeetingRooms) *ListEventsViewResponseBodyEvents {
	s.MeetingRooms = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetOnlineMeetingInfo(v *ListEventsViewResponseBodyEventsOnlineMeetingInfo) *ListEventsViewResponseBodyEvents {
	s.OnlineMeetingInfo = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetOrganizer(v *ListEventsViewResponseBodyEventsOrganizer) *ListEventsViewResponseBodyEvents {
	s.Organizer = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetOriginStart(v *ListEventsViewResponseBodyEventsOriginStart) *ListEventsViewResponseBodyEvents {
	s.OriginStart = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetRecurrence(v *ListEventsViewResponseBodyEventsRecurrence) *ListEventsViewResponseBodyEvents {
	s.Recurrence = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetRichTextDescription(v *ListEventsViewResponseBodyEventsRichTextDescription) *ListEventsViewResponseBodyEvents {
	s.RichTextDescription = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetSeriesMasterId(v string) *ListEventsViewResponseBodyEvents {
	s.SeriesMasterId = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetStart(v *ListEventsViewResponseBodyEventsStart) *ListEventsViewResponseBodyEvents {
	s.Start = v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetStatus(v string) *ListEventsViewResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetSummary(v string) *ListEventsViewResponseBodyEvents {
	s.Summary = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) SetUpdateTime(v string) *ListEventsViewResponseBodyEvents {
	s.UpdateTime = &v
	return s
}

func (s *ListEventsViewResponseBodyEvents) Validate() error {
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

type ListEventsViewResponseBodyEventsAttendees struct {
	// example:
	//
	// tony
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 012345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// false
	IsOptional *bool `json:"IsOptional,omitempty" xml:"IsOptional,omitempty"`
	// example:
	//
	// accepted
	ResponseStatus *string `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty"`
	// example:
	//
	// false
	Self *bool `json:"Self,omitempty" xml:"Self,omitempty"`
}

func (s ListEventsViewResponseBodyEventsAttendees) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsAttendees) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsAttendees) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEventsViewResponseBodyEventsAttendees) GetId() *string {
	return s.Id
}

func (s *ListEventsViewResponseBodyEventsAttendees) GetIsOptional() *bool {
	return s.IsOptional
}

func (s *ListEventsViewResponseBodyEventsAttendees) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *ListEventsViewResponseBodyEventsAttendees) GetSelf() *bool {
	return s.Self
}

func (s *ListEventsViewResponseBodyEventsAttendees) SetDisplayName(v string) *ListEventsViewResponseBodyEventsAttendees {
	s.DisplayName = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsAttendees) SetId(v string) *ListEventsViewResponseBodyEventsAttendees {
	s.Id = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsAttendees) SetIsOptional(v bool) *ListEventsViewResponseBodyEventsAttendees {
	s.IsOptional = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsAttendees) SetResponseStatus(v string) *ListEventsViewResponseBodyEventsAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsAttendees) SetSelf(v bool) *ListEventsViewResponseBodyEventsAttendees {
	s.Self = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsAttendees) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsCategories struct {
	// example:
	//
	// tony
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
}

func (s ListEventsViewResponseBodyEventsCategories) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsCategories) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsCategories) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEventsViewResponseBodyEventsCategories) SetDisplayName(v string) *ListEventsViewResponseBodyEventsCategories {
	s.DisplayName = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsCategories) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsEnd struct {
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

func (s ListEventsViewResponseBodyEventsEnd) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsEnd) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsEnd) GetDate() *string {
	return s.Date
}

func (s *ListEventsViewResponseBodyEventsEnd) GetDateTime() *string {
	return s.DateTime
}

func (s *ListEventsViewResponseBodyEventsEnd) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListEventsViewResponseBodyEventsEnd) SetDate(v string) *ListEventsViewResponseBodyEventsEnd {
	s.Date = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsEnd) SetDateTime(v string) *ListEventsViewResponseBodyEventsEnd {
	s.DateTime = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsEnd) SetTimeZone(v string) *ListEventsViewResponseBodyEventsEnd {
	s.TimeZone = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsEnd) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsExtendedProperties struct {
	SharedProperties *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties `json:"SharedProperties,omitempty" xml:"SharedProperties,omitempty" type:"Struct"`
}

func (s ListEventsViewResponseBodyEventsExtendedProperties) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsExtendedProperties) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsExtendedProperties) GetSharedProperties() *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties {
	return s.SharedProperties
}

func (s *ListEventsViewResponseBodyEventsExtendedProperties) SetSharedProperties(v *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) *ListEventsViewResponseBodyEventsExtendedProperties {
	s.SharedProperties = v
	return s
}

func (s *ListEventsViewResponseBodyEventsExtendedProperties) Validate() error {
	if s.SharedProperties != nil {
		if err := s.SharedProperties.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties struct {
	// example:
	//
	// dingd8*****1231
	BelongCorpId *string `json:"BelongCorpId,omitempty" xml:"BelongCorpId,omitempty"`
	// example:
	//
	// zxcvasdfa123===
	SourceOpenCid *string `json:"SourceOpenCid,omitempty" xml:"SourceOpenCid,omitempty"`
}

func (s ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) GetBelongCorpId() *string {
	return s.BelongCorpId
}

func (s *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) GetSourceOpenCid() *string {
	return s.SourceOpenCid
}

func (s *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) SetBelongCorpId(v string) *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties {
	s.BelongCorpId = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) SetSourceOpenCid(v string) *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties {
	s.SourceOpenCid = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsExtendedPropertiesSharedProperties) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsLocation struct {
	// example:
	//
	// tony
	DisplayName  *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	MeetingRooms []*string `json:"MeetingRooms,omitempty" xml:"MeetingRooms,omitempty" type:"Repeated"`
}

func (s ListEventsViewResponseBodyEventsLocation) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsLocation) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsLocation) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEventsViewResponseBodyEventsLocation) GetMeetingRooms() []*string {
	return s.MeetingRooms
}

func (s *ListEventsViewResponseBodyEventsLocation) SetDisplayName(v string) *ListEventsViewResponseBodyEventsLocation {
	s.DisplayName = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsLocation) SetMeetingRooms(v []*string) *ListEventsViewResponseBodyEventsLocation {
	s.MeetingRooms = v
	return s
}

func (s *ListEventsViewResponseBodyEventsLocation) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsMeetingRooms struct {
	// example:
	//
	// tony
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

func (s ListEventsViewResponseBodyEventsMeetingRooms) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsMeetingRooms) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsMeetingRooms) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEventsViewResponseBodyEventsMeetingRooms) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *ListEventsViewResponseBodyEventsMeetingRooms) GetRoomId() *string {
	return s.RoomId
}

func (s *ListEventsViewResponseBodyEventsMeetingRooms) SetDisplayName(v string) *ListEventsViewResponseBodyEventsMeetingRooms {
	s.DisplayName = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsMeetingRooms) SetResponseStatus(v string) *ListEventsViewResponseBodyEventsMeetingRooms {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsMeetingRooms) SetRoomId(v string) *ListEventsViewResponseBodyEventsMeetingRooms {
	s.RoomId = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsMeetingRooms) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsOnlineMeetingInfo struct {
	// example:
	//
	// 5c4df2xxx
	ConferenceId *string                `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	ExtraInfo    map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// dingtalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// dingtalk://dingtalkclient/page/xxx?confId=xxx&calendarId=xxx
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListEventsViewResponseBodyEventsOnlineMeetingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsOnlineMeetingInfo) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) GetType() *string {
	return s.Type
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) GetUrl() *string {
	return s.Url
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) SetConferenceId(v string) *ListEventsViewResponseBodyEventsOnlineMeetingInfo {
	s.ConferenceId = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) SetExtraInfo(v map[string]interface{}) *ListEventsViewResponseBodyEventsOnlineMeetingInfo {
	s.ExtraInfo = v
	return s
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) SetType(v string) *ListEventsViewResponseBodyEventsOnlineMeetingInfo {
	s.Type = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) SetUrl(v string) *ListEventsViewResponseBodyEventsOnlineMeetingInfo {
	s.Url = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOnlineMeetingInfo) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsOrganizer struct {
	// example:
	//
	// tony
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 012345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// accepted
	ResponseStatus *string `json:"ResponseStatus,omitempty" xml:"ResponseStatus,omitempty"`
	// example:
	//
	// false
	Self *bool `json:"Self,omitempty" xml:"Self,omitempty"`
}

func (s ListEventsViewResponseBodyEventsOrganizer) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsOrganizer) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsOrganizer) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEventsViewResponseBodyEventsOrganizer) GetId() *string {
	return s.Id
}

func (s *ListEventsViewResponseBodyEventsOrganizer) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *ListEventsViewResponseBodyEventsOrganizer) GetSelf() *bool {
	return s.Self
}

func (s *ListEventsViewResponseBodyEventsOrganizer) SetDisplayName(v string) *ListEventsViewResponseBodyEventsOrganizer {
	s.DisplayName = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOrganizer) SetId(v string) *ListEventsViewResponseBodyEventsOrganizer {
	s.Id = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOrganizer) SetResponseStatus(v string) *ListEventsViewResponseBodyEventsOrganizer {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOrganizer) SetSelf(v bool) *ListEventsViewResponseBodyEventsOrganizer {
	s.Self = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOrganizer) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsOriginStart struct {
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
}

func (s ListEventsViewResponseBodyEventsOriginStart) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsOriginStart) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsOriginStart) GetDateTime() *string {
	return s.DateTime
}

func (s *ListEventsViewResponseBodyEventsOriginStart) SetDateTime(v string) *ListEventsViewResponseBodyEventsOriginStart {
	s.DateTime = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsOriginStart) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsRecurrence struct {
	Pattern *ListEventsViewResponseBodyEventsRecurrencePattern `json:"Pattern,omitempty" xml:"Pattern,omitempty" type:"Struct"`
	Range   *ListEventsViewResponseBodyEventsRecurrenceRange   `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
}

func (s ListEventsViewResponseBodyEventsRecurrence) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsRecurrence) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsRecurrence) GetPattern() *ListEventsViewResponseBodyEventsRecurrencePattern {
	return s.Pattern
}

func (s *ListEventsViewResponseBodyEventsRecurrence) GetRange() *ListEventsViewResponseBodyEventsRecurrenceRange {
	return s.Range
}

func (s *ListEventsViewResponseBodyEventsRecurrence) SetPattern(v *ListEventsViewResponseBodyEventsRecurrencePattern) *ListEventsViewResponseBodyEventsRecurrence {
	s.Pattern = v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrence) SetRange(v *ListEventsViewResponseBodyEventsRecurrenceRange) *ListEventsViewResponseBodyEventsRecurrence {
	s.Range = v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrence) Validate() error {
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

type ListEventsViewResponseBodyEventsRecurrencePattern struct {
	// example:
	//
	// 14
	DayOfMonth *int32 `json:"DayOfMonth,omitempty" xml:"DayOfMonth,omitempty"`
	// example:
	//
	// sunday
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
	// dingtalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEventsViewResponseBodyEventsRecurrencePattern) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsRecurrencePattern) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) GetDayOfMonth() *int32 {
	return s.DayOfMonth
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) GetDaysOfWeek() *string {
	return s.DaysOfWeek
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) GetIndex() *string {
	return s.Index
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) GetInterval() *int32 {
	return s.Interval
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) GetType() *string {
	return s.Type
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) SetDayOfMonth(v int32) *ListEventsViewResponseBodyEventsRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) SetDaysOfWeek(v string) *ListEventsViewResponseBodyEventsRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) SetIndex(v string) *ListEventsViewResponseBodyEventsRecurrencePattern {
	s.Index = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) SetInterval(v int32) *ListEventsViewResponseBodyEventsRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) SetType(v string) *ListEventsViewResponseBodyEventsRecurrencePattern {
	s.Type = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrencePattern) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsRecurrenceRange struct {
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
	// dingtalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEventsViewResponseBodyEventsRecurrenceRange) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsRecurrenceRange) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsRecurrenceRange) GetEndDate() *string {
	return s.EndDate
}

func (s *ListEventsViewResponseBodyEventsRecurrenceRange) GetNumberOfOccurrences() *int32 {
	return s.NumberOfOccurrences
}

func (s *ListEventsViewResponseBodyEventsRecurrenceRange) GetType() *string {
	return s.Type
}

func (s *ListEventsViewResponseBodyEventsRecurrenceRange) SetEndDate(v string) *ListEventsViewResponseBodyEventsRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrenceRange) SetNumberOfOccurrences(v int32) *ListEventsViewResponseBodyEventsRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrenceRange) SetType(v string) *ListEventsViewResponseBodyEventsRecurrenceRange {
	s.Type = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRecurrenceRange) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsRichTextDescription struct {
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ListEventsViewResponseBodyEventsRichTextDescription) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsRichTextDescription) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsRichTextDescription) GetText() *string {
	return s.Text
}

func (s *ListEventsViewResponseBodyEventsRichTextDescription) SetText(v string) *ListEventsViewResponseBodyEventsRichTextDescription {
	s.Text = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsRichTextDescription) Validate() error {
	return dara.Validate(s)
}

type ListEventsViewResponseBodyEventsStart struct {
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

func (s ListEventsViewResponseBodyEventsStart) String() string {
	return dara.Prettify(s)
}

func (s ListEventsViewResponseBodyEventsStart) GoString() string {
	return s.String()
}

func (s *ListEventsViewResponseBodyEventsStart) GetDate() *string {
	return s.Date
}

func (s *ListEventsViewResponseBodyEventsStart) GetDateTime() *string {
	return s.DateTime
}

func (s *ListEventsViewResponseBodyEventsStart) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListEventsViewResponseBodyEventsStart) SetDate(v string) *ListEventsViewResponseBodyEventsStart {
	s.Date = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsStart) SetDateTime(v string) *ListEventsViewResponseBodyEventsStart {
	s.DateTime = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsStart) SetTimeZone(v string) *ListEventsViewResponseBodyEventsStart {
	s.TimeZone = &v
	return s
}

func (s *ListEventsViewResponseBodyEventsStart) Validate() error {
	return dara.Validate(s)
}
