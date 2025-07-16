// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEvents(v []*ListEventsResponseBodyEvents) *ListEventsResponseBody
	GetEvents() []*ListEventsResponseBodyEvents
	SetNextToken(v string) *ListEventsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListEventsResponseBody
	GetRequestId() *string
	SetSyncToken(v string) *ListEventsResponseBody
	GetSyncToken() *string
	SetVendorRequestId(v string) *ListEventsResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *ListEventsResponseBody
	GetVendorType() *string
}

type ListEventsResponseBody struct {
	Events []*ListEventsResponseBodyEvents `json:"events,omitempty" xml:"events,omitempty" type:"Repeated"`
	// example:
	//
	// cnNTbW1YbxxxxdlQrQT09
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// requestId
	//
	// example:
	//
	// 4248DCC9-785F-5A14-8BE0-830FD52E1261
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// zxcasdfvc000009
	SyncToken       *string `json:"syncToken,omitempty" xml:"syncToken,omitempty"`
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	VendorType      *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s ListEventsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBody) GetEvents() []*ListEventsResponseBodyEvents {
	return s.Events
}

func (s *ListEventsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventsResponseBody) GetSyncToken() *string {
	return s.SyncToken
}

func (s *ListEventsResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *ListEventsResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *ListEventsResponseBody) SetEvents(v []*ListEventsResponseBodyEvents) *ListEventsResponseBody {
	s.Events = v
	return s
}

func (s *ListEventsResponseBody) SetNextToken(v string) *ListEventsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListEventsResponseBody) SetRequestId(v string) *ListEventsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventsResponseBody) SetSyncToken(v string) *ListEventsResponseBody {
	s.SyncToken = &v
	return s
}

func (s *ListEventsResponseBody) SetVendorRequestId(v string) *ListEventsResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *ListEventsResponseBody) SetVendorType(v string) *ListEventsResponseBody {
	s.VendorType = &v
	return s
}

func (s *ListEventsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEvents struct {
	Attendees  []*ListEventsResponseBodyEventsAttendees  `json:"Attendees,omitempty" xml:"Attendees,omitempty" type:"Repeated"`
	Categories []*ListEventsResponseBodyEventsCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// something about this event
	Description        *string                                         `json:"Description,omitempty" xml:"Description,omitempty"`
	End                *ListEventsResponseBodyEventsEnd                `json:"End,omitempty" xml:"End,omitempty" type:"Struct"`
	ExtendedProperties *ListEventsResponseBodyEventsExtendedProperties `json:"ExtendedProperties,omitempty" xml:"ExtendedProperties,omitempty" type:"Struct"`
	// example:
	//
	// cnNTbW1YbxxxxdEgvdlQrQT09
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// true
	IsAllDay            *bool                                            `json:"IsAllDay,omitempty" xml:"IsAllDay,omitempty"`
	Location            *ListEventsResponseBodyEventsLocation            `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	MeetingRooms        []*ListEventsResponseBodyEventsMeetingRooms      `json:"MeetingRooms,omitempty" xml:"MeetingRooms,omitempty" type:"Repeated"`
	OnlineMeetingInfo   *ListEventsResponseBodyEventsOnlineMeetingInfo   `json:"OnlineMeetingInfo,omitempty" xml:"OnlineMeetingInfo,omitempty" type:"Struct"`
	Organizer           *ListEventsResponseBodyEventsOrganizer           `json:"Organizer,omitempty" xml:"Organizer,omitempty" type:"Struct"`
	OriginStart         *ListEventsResponseBodyEventsOriginStart         `json:"OriginStart,omitempty" xml:"OriginStart,omitempty" type:"Struct"`
	Recurrence          *ListEventsResponseBodyEventsRecurrence          `json:"Recurrence,omitempty" xml:"Recurrence,omitempty" type:"Struct"`
	Reminders           []*ListEventsResponseBodyEventsReminders         `json:"Reminders,omitempty" xml:"Reminders,omitempty" type:"Repeated"`
	RichTextDescription *ListEventsResponseBodyEventsRichTextDescription `json:"RichTextDescription,omitempty" xml:"RichTextDescription,omitempty" type:"Struct"`
	// example:
	//
	// cnNTbWxxxxaFJZdEgvdlQrQT09
	SeriesMasterId *string                            `json:"SeriesMasterId,omitempty" xml:"SeriesMasterId,omitempty"`
	Start          *ListEventsResponseBodyEventsStart `json:"Start,omitempty" xml:"Start,omitempty" type:"Struct"`
	// example:
	//
	// confirmed
	Status  *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListEventsResponseBodyEvents) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEvents) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEvents) GetAttendees() []*ListEventsResponseBodyEventsAttendees {
	return s.Attendees
}

func (s *ListEventsResponseBodyEvents) GetCategories() []*ListEventsResponseBodyEventsCategories {
	return s.Categories
}

func (s *ListEventsResponseBodyEvents) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListEventsResponseBodyEvents) GetDescription() *string {
	return s.Description
}

func (s *ListEventsResponseBodyEvents) GetEnd() *ListEventsResponseBodyEventsEnd {
	return s.End
}

func (s *ListEventsResponseBodyEvents) GetExtendedProperties() *ListEventsResponseBodyEventsExtendedProperties {
	return s.ExtendedProperties
}

func (s *ListEventsResponseBodyEvents) GetId() *string {
	return s.Id
}

func (s *ListEventsResponseBodyEvents) GetIsAllDay() *bool {
	return s.IsAllDay
}

func (s *ListEventsResponseBodyEvents) GetLocation() *ListEventsResponseBodyEventsLocation {
	return s.Location
}

func (s *ListEventsResponseBodyEvents) GetMeetingRooms() []*ListEventsResponseBodyEventsMeetingRooms {
	return s.MeetingRooms
}

func (s *ListEventsResponseBodyEvents) GetOnlineMeetingInfo() *ListEventsResponseBodyEventsOnlineMeetingInfo {
	return s.OnlineMeetingInfo
}

func (s *ListEventsResponseBodyEvents) GetOrganizer() *ListEventsResponseBodyEventsOrganizer {
	return s.Organizer
}

func (s *ListEventsResponseBodyEvents) GetOriginStart() *ListEventsResponseBodyEventsOriginStart {
	return s.OriginStart
}

func (s *ListEventsResponseBodyEvents) GetRecurrence() *ListEventsResponseBodyEventsRecurrence {
	return s.Recurrence
}

func (s *ListEventsResponseBodyEvents) GetReminders() []*ListEventsResponseBodyEventsReminders {
	return s.Reminders
}

func (s *ListEventsResponseBodyEvents) GetRichTextDescription() *ListEventsResponseBodyEventsRichTextDescription {
	return s.RichTextDescription
}

func (s *ListEventsResponseBodyEvents) GetSeriesMasterId() *string {
	return s.SeriesMasterId
}

func (s *ListEventsResponseBodyEvents) GetStart() *ListEventsResponseBodyEventsStart {
	return s.Start
}

func (s *ListEventsResponseBodyEvents) GetStatus() *string {
	return s.Status
}

func (s *ListEventsResponseBodyEvents) GetSummary() *string {
	return s.Summary
}

func (s *ListEventsResponseBodyEvents) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListEventsResponseBodyEvents) SetAttendees(v []*ListEventsResponseBodyEventsAttendees) *ListEventsResponseBodyEvents {
	s.Attendees = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetCategories(v []*ListEventsResponseBodyEventsCategories) *ListEventsResponseBodyEvents {
	s.Categories = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetCreateTime(v string) *ListEventsResponseBodyEvents {
	s.CreateTime = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetDescription(v string) *ListEventsResponseBodyEvents {
	s.Description = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetEnd(v *ListEventsResponseBodyEventsEnd) *ListEventsResponseBodyEvents {
	s.End = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetExtendedProperties(v *ListEventsResponseBodyEventsExtendedProperties) *ListEventsResponseBodyEvents {
	s.ExtendedProperties = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetId(v string) *ListEventsResponseBodyEvents {
	s.Id = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetIsAllDay(v bool) *ListEventsResponseBodyEvents {
	s.IsAllDay = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetLocation(v *ListEventsResponseBodyEventsLocation) *ListEventsResponseBodyEvents {
	s.Location = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetMeetingRooms(v []*ListEventsResponseBodyEventsMeetingRooms) *ListEventsResponseBodyEvents {
	s.MeetingRooms = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetOnlineMeetingInfo(v *ListEventsResponseBodyEventsOnlineMeetingInfo) *ListEventsResponseBodyEvents {
	s.OnlineMeetingInfo = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetOrganizer(v *ListEventsResponseBodyEventsOrganizer) *ListEventsResponseBodyEvents {
	s.Organizer = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetOriginStart(v *ListEventsResponseBodyEventsOriginStart) *ListEventsResponseBodyEvents {
	s.OriginStart = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetRecurrence(v *ListEventsResponseBodyEventsRecurrence) *ListEventsResponseBodyEvents {
	s.Recurrence = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetReminders(v []*ListEventsResponseBodyEventsReminders) *ListEventsResponseBodyEvents {
	s.Reminders = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetRichTextDescription(v *ListEventsResponseBodyEventsRichTextDescription) *ListEventsResponseBodyEvents {
	s.RichTextDescription = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetSeriesMasterId(v string) *ListEventsResponseBodyEvents {
	s.SeriesMasterId = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetStart(v *ListEventsResponseBodyEventsStart) *ListEventsResponseBodyEvents {
	s.Start = v
	return s
}

func (s *ListEventsResponseBodyEvents) SetStatus(v string) *ListEventsResponseBodyEvents {
	s.Status = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetSummary(v string) *ListEventsResponseBodyEvents {
	s.Summary = &v
	return s
}

func (s *ListEventsResponseBodyEvents) SetUpdateTime(v string) *ListEventsResponseBodyEvents {
	s.UpdateTime = &v
	return s
}

func (s *ListEventsResponseBodyEvents) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsAttendees struct {
	// example:
	//
	// tony
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
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

func (s ListEventsResponseBodyEventsAttendees) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsAttendees) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsAttendees) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEventsResponseBodyEventsAttendees) GetIsOptional() *bool {
	return s.IsOptional
}

func (s *ListEventsResponseBodyEventsAttendees) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *ListEventsResponseBodyEventsAttendees) GetSelf() *bool {
	return s.Self
}

func (s *ListEventsResponseBodyEventsAttendees) SetDisplayName(v string) *ListEventsResponseBodyEventsAttendees {
	s.DisplayName = &v
	return s
}

func (s *ListEventsResponseBodyEventsAttendees) SetIsOptional(v bool) *ListEventsResponseBodyEventsAttendees {
	s.IsOptional = &v
	return s
}

func (s *ListEventsResponseBodyEventsAttendees) SetResponseStatus(v string) *ListEventsResponseBodyEventsAttendees {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsResponseBodyEventsAttendees) SetSelf(v bool) *ListEventsResponseBodyEventsAttendees {
	s.Self = &v
	return s
}

func (s *ListEventsResponseBodyEventsAttendees) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsCategories struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
}

func (s ListEventsResponseBodyEventsCategories) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsCategories) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsCategories) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEventsResponseBodyEventsCategories) SetDisplayName(v string) *ListEventsResponseBodyEventsCategories {
	s.DisplayName = &v
	return s
}

func (s *ListEventsResponseBodyEventsCategories) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsEnd struct {
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

func (s ListEventsResponseBodyEventsEnd) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsEnd) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsEnd) GetDate() *string {
	return s.Date
}

func (s *ListEventsResponseBodyEventsEnd) GetDateTime() *string {
	return s.DateTime
}

func (s *ListEventsResponseBodyEventsEnd) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListEventsResponseBodyEventsEnd) SetDate(v string) *ListEventsResponseBodyEventsEnd {
	s.Date = &v
	return s
}

func (s *ListEventsResponseBodyEventsEnd) SetDateTime(v string) *ListEventsResponseBodyEventsEnd {
	s.DateTime = &v
	return s
}

func (s *ListEventsResponseBodyEventsEnd) SetTimeZone(v string) *ListEventsResponseBodyEventsEnd {
	s.TimeZone = &v
	return s
}

func (s *ListEventsResponseBodyEventsEnd) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsExtendedProperties struct {
	SharedProperties *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties `json:"SharedProperties,omitempty" xml:"SharedProperties,omitempty" type:"Struct"`
}

func (s ListEventsResponseBodyEventsExtendedProperties) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsExtendedProperties) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsExtendedProperties) GetSharedProperties() *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties {
	return s.SharedProperties
}

func (s *ListEventsResponseBodyEventsExtendedProperties) SetSharedProperties(v *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) *ListEventsResponseBodyEventsExtendedProperties {
	s.SharedProperties = v
	return s
}

func (s *ListEventsResponseBodyEventsExtendedProperties) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsExtendedPropertiesSharedProperties struct {
	// example:
	//
	// ding*********31
	BelongCorpId *string `json:"BelongCorpId,omitempty" xml:"BelongCorpId,omitempty"`
	// example:
	//
	// zxcv90asdf123===
	SourceOpenCid *string `json:"SourceOpenCid,omitempty" xml:"SourceOpenCid,omitempty"`
}

func (s ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) GetBelongCorpId() *string {
	return s.BelongCorpId
}

func (s *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) GetSourceOpenCid() *string {
	return s.SourceOpenCid
}

func (s *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) SetBelongCorpId(v string) *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties {
	s.BelongCorpId = &v
	return s
}

func (s *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) SetSourceOpenCid(v string) *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties {
	s.SourceOpenCid = &v
	return s
}

func (s *ListEventsResponseBodyEventsExtendedPropertiesSharedProperties) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsLocation struct {
	// example:
	//
	// room 1-2-3
	DisplayName  *string   `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	MeetingRooms []*string `json:"MeetingRooms,omitempty" xml:"MeetingRooms,omitempty" type:"Repeated"`
}

func (s ListEventsResponseBodyEventsLocation) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsLocation) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsLocation) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEventsResponseBodyEventsLocation) GetMeetingRooms() []*string {
	return s.MeetingRooms
}

func (s *ListEventsResponseBodyEventsLocation) SetDisplayName(v string) *ListEventsResponseBodyEventsLocation {
	s.DisplayName = &v
	return s
}

func (s *ListEventsResponseBodyEventsLocation) SetMeetingRooms(v []*string) *ListEventsResponseBodyEventsLocation {
	s.MeetingRooms = v
	return s
}

func (s *ListEventsResponseBodyEventsLocation) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsMeetingRooms struct {
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

func (s ListEventsResponseBodyEventsMeetingRooms) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsMeetingRooms) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsMeetingRooms) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEventsResponseBodyEventsMeetingRooms) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *ListEventsResponseBodyEventsMeetingRooms) GetRoomId() *string {
	return s.RoomId
}

func (s *ListEventsResponseBodyEventsMeetingRooms) SetDisplayName(v string) *ListEventsResponseBodyEventsMeetingRooms {
	s.DisplayName = &v
	return s
}

func (s *ListEventsResponseBodyEventsMeetingRooms) SetResponseStatus(v string) *ListEventsResponseBodyEventsMeetingRooms {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsResponseBodyEventsMeetingRooms) SetRoomId(v string) *ListEventsResponseBodyEventsMeetingRooms {
	s.RoomId = &v
	return s
}

func (s *ListEventsResponseBodyEventsMeetingRooms) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsOnlineMeetingInfo struct {
	// example:
	//
	// 5c4df21dxxxx-a6db402b9f3a
	ConferenceId *string                `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	ExtraInfo    map[string]interface{} `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// example:
	//
	// dingtalk
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// dingtalk://dingtalkclient/page/videoCoxxxxndar?confId=5c4df21dxxxx2b9f3a&calendarId=92xxxx36
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s ListEventsResponseBodyEventsOnlineMeetingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsOnlineMeetingInfo) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) GetExtraInfo() map[string]interface{} {
	return s.ExtraInfo
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) GetType() *string {
	return s.Type
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) GetUrl() *string {
	return s.Url
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) SetConferenceId(v string) *ListEventsResponseBodyEventsOnlineMeetingInfo {
	s.ConferenceId = &v
	return s
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) SetExtraInfo(v map[string]interface{}) *ListEventsResponseBodyEventsOnlineMeetingInfo {
	s.ExtraInfo = v
	return s
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) SetType(v string) *ListEventsResponseBodyEventsOnlineMeetingInfo {
	s.Type = &v
	return s
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) SetUrl(v string) *ListEventsResponseBodyEventsOnlineMeetingInfo {
	s.Url = &v
	return s
}

func (s *ListEventsResponseBodyEventsOnlineMeetingInfo) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsOrganizer struct {
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
	// true
	Self *bool `json:"Self,omitempty" xml:"Self,omitempty"`
}

func (s ListEventsResponseBodyEventsOrganizer) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsOrganizer) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsOrganizer) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListEventsResponseBodyEventsOrganizer) GetResponseStatus() *string {
	return s.ResponseStatus
}

func (s *ListEventsResponseBodyEventsOrganizer) GetSelf() *bool {
	return s.Self
}

func (s *ListEventsResponseBodyEventsOrganizer) SetDisplayName(v string) *ListEventsResponseBodyEventsOrganizer {
	s.DisplayName = &v
	return s
}

func (s *ListEventsResponseBodyEventsOrganizer) SetResponseStatus(v string) *ListEventsResponseBodyEventsOrganizer {
	s.ResponseStatus = &v
	return s
}

func (s *ListEventsResponseBodyEventsOrganizer) SetSelf(v bool) *ListEventsResponseBodyEventsOrganizer {
	s.Self = &v
	return s
}

func (s *ListEventsResponseBodyEventsOrganizer) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsOriginStart struct {
	// example:
	//
	// 2023-01-01T00:00:00Z
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
}

func (s ListEventsResponseBodyEventsOriginStart) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsOriginStart) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsOriginStart) GetDateTime() *string {
	return s.DateTime
}

func (s *ListEventsResponseBodyEventsOriginStart) SetDateTime(v string) *ListEventsResponseBodyEventsOriginStart {
	s.DateTime = &v
	return s
}

func (s *ListEventsResponseBodyEventsOriginStart) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsRecurrence struct {
	Pattern *ListEventsResponseBodyEventsRecurrencePattern `json:"Pattern,omitempty" xml:"Pattern,omitempty" type:"Struct"`
	Range   *ListEventsResponseBodyEventsRecurrenceRange   `json:"Range,omitempty" xml:"Range,omitempty" type:"Struct"`
}

func (s ListEventsResponseBodyEventsRecurrence) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsRecurrence) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsRecurrence) GetPattern() *ListEventsResponseBodyEventsRecurrencePattern {
	return s.Pattern
}

func (s *ListEventsResponseBodyEventsRecurrence) GetRange() *ListEventsResponseBodyEventsRecurrenceRange {
	return s.Range
}

func (s *ListEventsResponseBodyEventsRecurrence) SetPattern(v *ListEventsResponseBodyEventsRecurrencePattern) *ListEventsResponseBodyEventsRecurrence {
	s.Pattern = v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrence) SetRange(v *ListEventsResponseBodyEventsRecurrenceRange) *ListEventsResponseBodyEventsRecurrence {
	s.Range = v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrence) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsRecurrencePattern struct {
	// example:
	//
	// 1
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
	// daily
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListEventsResponseBodyEventsRecurrencePattern) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsRecurrencePattern) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) GetDayOfMonth() *int32 {
	return s.DayOfMonth
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) GetDaysOfWeek() *string {
	return s.DaysOfWeek
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) GetIndex() *string {
	return s.Index
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) GetInterval() *int32 {
	return s.Interval
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) GetType() *string {
	return s.Type
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) SetDayOfMonth(v int32) *ListEventsResponseBodyEventsRecurrencePattern {
	s.DayOfMonth = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) SetDaysOfWeek(v string) *ListEventsResponseBodyEventsRecurrencePattern {
	s.DaysOfWeek = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) SetIndex(v string) *ListEventsResponseBodyEventsRecurrencePattern {
	s.Index = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) SetInterval(v int32) *ListEventsResponseBodyEventsRecurrencePattern {
	s.Interval = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) SetType(v string) *ListEventsResponseBodyEventsRecurrencePattern {
	s.Type = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrencePattern) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsRecurrenceRange struct {
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

func (s ListEventsResponseBodyEventsRecurrenceRange) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsRecurrenceRange) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsRecurrenceRange) GetEndDate() *string {
	return s.EndDate
}

func (s *ListEventsResponseBodyEventsRecurrenceRange) GetNumberOfOccurrences() *int32 {
	return s.NumberOfOccurrences
}

func (s *ListEventsResponseBodyEventsRecurrenceRange) GetType() *string {
	return s.Type
}

func (s *ListEventsResponseBodyEventsRecurrenceRange) SetEndDate(v string) *ListEventsResponseBodyEventsRecurrenceRange {
	s.EndDate = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrenceRange) SetNumberOfOccurrences(v int32) *ListEventsResponseBodyEventsRecurrenceRange {
	s.NumberOfOccurrences = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrenceRange) SetType(v string) *ListEventsResponseBodyEventsRecurrenceRange {
	s.Type = &v
	return s
}

func (s *ListEventsResponseBodyEventsRecurrenceRange) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsReminders struct {
	// example:
	//
	// dingtalk
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// example:
	//
	// 15
	Minutes *string `json:"Minutes,omitempty" xml:"Minutes,omitempty"`
}

func (s ListEventsResponseBodyEventsReminders) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsReminders) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsReminders) GetMethod() *string {
	return s.Method
}

func (s *ListEventsResponseBodyEventsReminders) GetMinutes() *string {
	return s.Minutes
}

func (s *ListEventsResponseBodyEventsReminders) SetMethod(v string) *ListEventsResponseBodyEventsReminders {
	s.Method = &v
	return s
}

func (s *ListEventsResponseBodyEventsReminders) SetMinutes(v string) *ListEventsResponseBodyEventsReminders {
	s.Minutes = &v
	return s
}

func (s *ListEventsResponseBodyEventsReminders) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsRichTextDescription struct {
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ListEventsResponseBodyEventsRichTextDescription) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsRichTextDescription) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsRichTextDescription) GetText() *string {
	return s.Text
}

func (s *ListEventsResponseBodyEventsRichTextDescription) SetText(v string) *ListEventsResponseBodyEventsRichTextDescription {
	s.Text = &v
	return s
}

func (s *ListEventsResponseBodyEventsRichTextDescription) Validate() error {
	return dara.Validate(s)
}

type ListEventsResponseBodyEventsStart struct {
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

func (s ListEventsResponseBodyEventsStart) String() string {
	return dara.Prettify(s)
}

func (s ListEventsResponseBodyEventsStart) GoString() string {
	return s.String()
}

func (s *ListEventsResponseBodyEventsStart) GetDate() *string {
	return s.Date
}

func (s *ListEventsResponseBodyEventsStart) GetDateTime() *string {
	return s.DateTime
}

func (s *ListEventsResponseBodyEventsStart) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListEventsResponseBodyEventsStart) SetDate(v string) *ListEventsResponseBodyEventsStart {
	s.Date = &v
	return s
}

func (s *ListEventsResponseBodyEventsStart) SetDateTime(v string) *ListEventsResponseBodyEventsStart {
	s.DateTime = &v
	return s
}

func (s *ListEventsResponseBodyEventsStart) SetTimeZone(v string) *ListEventsResponseBodyEventsStart {
	s.TimeZone = &v
	return s
}

func (s *ListEventsResponseBodyEventsStart) Validate() error {
	return dara.Validate(s)
}
