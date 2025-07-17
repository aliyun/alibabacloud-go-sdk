// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOnCallSchedulesDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetOnCallSchedulesDetailResponseBodyData) *GetOnCallSchedulesDetailResponseBody
	GetData() *GetOnCallSchedulesDetailResponseBodyData
	SetRequestId(v string) *GetOnCallSchedulesDetailResponseBody
	GetRequestId() *string
}

type GetOnCallSchedulesDetailResponseBody struct {
	// The information about the scheduling policy.
	Data *GetOnCallSchedulesDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 21E85B16-75A6-429A-9F65-8AAC9A54****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBody) GetData() *GetOnCallSchedulesDetailResponseBodyData {
	return s.Data
}

func (s *GetOnCallSchedulesDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOnCallSchedulesDetailResponseBody) SetData(v *GetOnCallSchedulesDetailResponseBodyData) *GetOnCallSchedulesDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBody) SetRequestId(v string) *GetOnCallSchedulesDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOnCallSchedulesDetailResponseBodyData struct {
	// The URL of the DingTalk chatbot, which is used to receive notifications about shift changes.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=69d4e009547e11069c6513309414937b7bf0482fb9284125b5******
	AlertRobotId *int64 `json:"AlertRobotId,omitempty" xml:"AlertRobotId,omitempty"`
	// The description of the scheduling policy.
	//
	// example:
	//
	// Test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the scheduling policy.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the scheduling policy.
	//
	// example:
	//
	// Scheduling policy test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The information about the final user on duty.
	RenderedFinnalEntries []*GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries `json:"RenderedFinnalEntries,omitempty" xml:"RenderedFinnalEntries,omitempty" type:"Repeated"`
	// The scheduled users on duty within a time range.
	RenderedLayerEntries [][]*GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries `json:"RenderedLayerEntries,omitempty" xml:"RenderedLayerEntries,omitempty" type:"Repeated"`
	// The information about the substitutes within a time range.
	RenderedSubstitudeEntries []*GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries `json:"RenderedSubstitudeEntries,omitempty" xml:"RenderedSubstitudeEntries,omitempty" type:"Repeated"`
	// The information about the shift.
	ScheduleLayers []*GetOnCallSchedulesDetailResponseBodyDataScheduleLayers `json:"ScheduleLayers,omitempty" xml:"ScheduleLayers,omitempty" type:"Repeated"`
}

func (s GetOnCallSchedulesDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyData) GetAlertRobotId() *int64 {
	return s.AlertRobotId
}

func (s *GetOnCallSchedulesDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetOnCallSchedulesDetailResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *GetOnCallSchedulesDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetOnCallSchedulesDetailResponseBodyData) GetRenderedFinnalEntries() []*GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries {
	return s.RenderedFinnalEntries
}

func (s *GetOnCallSchedulesDetailResponseBodyData) GetRenderedLayerEntries() [][]*GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries {
	return s.RenderedLayerEntries
}

func (s *GetOnCallSchedulesDetailResponseBodyData) GetRenderedSubstitudeEntries() []*GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries {
	return s.RenderedSubstitudeEntries
}

func (s *GetOnCallSchedulesDetailResponseBodyData) GetScheduleLayers() []*GetOnCallSchedulesDetailResponseBodyDataScheduleLayers {
	return s.ScheduleLayers
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetAlertRobotId(v int64) *GetOnCallSchedulesDetailResponseBodyData {
	s.AlertRobotId = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetDescription(v string) *GetOnCallSchedulesDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetId(v int64) *GetOnCallSchedulesDetailResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetName(v string) *GetOnCallSchedulesDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetRenderedFinnalEntries(v []*GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) *GetOnCallSchedulesDetailResponseBodyData {
	s.RenderedFinnalEntries = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetRenderedLayerEntries(v [][]*GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) *GetOnCallSchedulesDetailResponseBodyData {
	s.RenderedLayerEntries = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetRenderedSubstitudeEntries(v []*GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) *GetOnCallSchedulesDetailResponseBodyData {
	s.RenderedSubstitudeEntries = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) SetScheduleLayers(v []*GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) *GetOnCallSchedulesDetailResponseBodyData {
	s.ScheduleLayers = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries struct {
	// The date on which the user completed shift work.
	//
	// example:
	//
	// 2022-10-30
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The information about the user on duty.
	SimpleContact *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact `json:"SimpleContact,omitempty" xml:"SimpleContact,omitempty" type:"Struct"`
	// The date from which the user started shift work.
	//
	// example:
	//
	// 2022-10-01
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) GetEnd() *string {
	return s.End
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) GetSimpleContact() *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact {
	return s.SimpleContact
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) GetStart() *string {
	return s.Start
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) SetEnd(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries {
	s.End = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) SetSimpleContact(v *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries {
	s.SimpleContact = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) SetStart(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries {
	s.Start = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries) Validate() error {
	return dara.Validate(s)
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact struct {
	// The ID of the user on duty.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the user on duty.
	//
	// example:
	//
	// Employee 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) GetId() *int64 {
	return s.Id
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) GetName() *string {
	return s.Name
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) SetId(v int64) *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact {
	s.Id = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) SetName(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact {
	s.Name = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact) Validate() error {
	return dara.Validate(s)
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries struct {
	// The date from which the scheduled user was supposed to start shift work.
	//
	// example:
	//
	// 2022-10-01
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
	// The date on which the scheduled user was supposed to complete shift work.
	//
	// example:
	//
	// 2022-10-30
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The information about the scheduled user.
	SimpleContact *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact `json:"SimpleContact,omitempty" xml:"SimpleContact,omitempty" type:"Struct"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) GetStart() *string {
	return s.Start
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) GetEnd() *string {
	return s.End
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) GetSimpleContact() *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact {
	return s.SimpleContact
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) SetStart(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries {
	s.Start = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) SetEnd(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries {
	s.End = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) SetSimpleContact(v *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries {
	s.SimpleContact = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries) Validate() error {
	return dara.Validate(s)
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact struct {
	// The ID of the scheduled user.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the scheduled user.
	//
	// example:
	//
	// Employee 1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) GetId() *int64 {
	return s.Id
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) GetName() *string {
	return s.Name
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) SetId(v int64) *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact {
	s.Id = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) SetName(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact {
	s.Name = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact) Validate() error {
	return dara.Validate(s)
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries struct {
	// The date on which the substitute was supposed to complete shift work.
	//
	// example:
	//
	// 2022-10-30
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// The information about the substitute.
	SimpleContact *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact `json:"SimpleContact,omitempty" xml:"SimpleContact,omitempty" type:"Struct"`
	// The date from which the substitute was supposed to start shift work.
	//
	// example:
	//
	// 2022-10-01
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) GetEnd() *string {
	return s.End
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) GetSimpleContact() *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact {
	return s.SimpleContact
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) GetStart() *string {
	return s.Start
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) SetEnd(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries {
	s.End = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) SetSimpleContact(v *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries {
	s.SimpleContact = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) SetStart(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries {
	s.Start = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries) Validate() error {
	return dara.Validate(s)
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact struct {
	// The ID of the substitute.
	//
	// example:
	//
	// 234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the substitute.
	//
	// example:
	//
	// Employee 2
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) GetId() *int64 {
	return s.Id
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) GetName() *string {
	return s.Name
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) SetId(v int64) *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact {
	s.Id = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) SetName(v string) *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact {
	s.Name = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact) Validate() error {
	return dara.Validate(s)
}

type GetOnCallSchedulesDetailResponseBodyDataScheduleLayers struct {
	// The ID list of users on duty.
	ContactIds []*int64 `json:"ContactIds,omitempty" xml:"ContactIds,omitempty" type:"Repeated"`
	// The limit on the time of the shift.
	Restrictions []*GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions `json:"Restrictions,omitempty" xml:"Restrictions,omitempty" type:"Repeated"`
	// The type of the shift. Valid values:
	//
	// 	- DAY
	//
	// 	- WEEK
	//
	// 	- CUSTOM
	//
	// example:
	//
	// DAY
	RotationType *string `json:"RotationType,omitempty" xml:"RotationType,omitempty"`
	// The shift cycle. Unit: hours.
	//
	// example:
	//
	// 8
	ShiftLength *int64 `json:"ShiftLength,omitempty" xml:"ShiftLength,omitempty"`
	// The date on which the shift change took effect.
	//
	// example:
	//
	// 2022-10-01
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) GetContactIds() []*int64 {
	return s.ContactIds
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) GetRestrictions() []*GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions {
	return s.Restrictions
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) GetRotationType() *string {
	return s.RotationType
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) GetShiftLength() *int64 {
	return s.ShiftLength
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) GetStartTime() *string {
	return s.StartTime
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) SetContactIds(v []*int64) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers {
	s.ContactIds = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) SetRestrictions(v []*GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers {
	s.Restrictions = v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) SetRotationType(v string) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers {
	s.RotationType = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) SetShiftLength(v int64) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers {
	s.ShiftLength = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) SetStartTime(v string) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers {
	s.StartTime = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayers) Validate() error {
	return dara.Validate(s)
}

type GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions struct {
	// The end time of the shift per day.
	//
	// example:
	//
	// 18:00
	EndTimeOfDay *string `json:"EndTimeOfDay,omitempty" xml:"EndTimeOfDay,omitempty"`
	// The type of the limit. Valid values:
	//
	// 	- daily_restriction
	//
	// 	- weekly_restriction
	//
	// example:
	//
	// daily_restriction
	RestrictionType *string `json:"RestrictionType,omitempty" xml:"RestrictionType,omitempty"`
	// The start time of the shift per day.
	//
	// example:
	//
	// 09:00
	StartTimeOfDay *string `json:"StartTimeOfDay,omitempty" xml:"StartTimeOfDay,omitempty"`
}

func (s GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) String() string {
	return dara.Prettify(s)
}

func (s GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) GoString() string {
	return s.String()
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) GetEndTimeOfDay() *string {
	return s.EndTimeOfDay
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) GetRestrictionType() *string {
	return s.RestrictionType
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) GetStartTimeOfDay() *string {
	return s.StartTimeOfDay
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) SetEndTimeOfDay(v string) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions {
	s.EndTimeOfDay = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) SetRestrictionType(v string) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions {
	s.RestrictionType = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) SetStartTimeOfDay(v string) *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions {
	s.StartTimeOfDay = &v
	return s
}

func (s *GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions) Validate() error {
	return dara.Validate(s)
}
