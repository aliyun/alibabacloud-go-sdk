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
	// The details of the on-call schedule.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOnCallSchedulesDetailResponseBodyData struct {
	// The webhook URL of the DingTalk bot for rotation notifications.
	//
	// example:
	//
	// https://oapi.dingtalk.com/robot/send?access_token=69d4e009547e11069c6513309414937b7bf0482fb9284125b5******
	AlertRobotId *int64 `json:"AlertRobotId,omitempty" xml:"AlertRobotId,omitempty"`
	// The description of the on-call schedule.
	//
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the on-call schedule.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the on-call schedule.
	//
	// example:
	//
	// 排班策略测试
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The final list of on-call contacts, after accounting for all rotations and substitutions.
	RenderedFinnalEntries []*GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries `json:"RenderedFinnalEntries,omitempty" xml:"RenderedFinnalEntries,omitempty" type:"Repeated"`
	// A list of contacts on duty within the specified time range, as defined by the schedule layers.
	RenderedLayerEntries [][]*GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntries `json:"RenderedLayerEntries,omitempty" xml:"RenderedLayerEntries,omitempty" type:"Repeated"`
	// A list of substitutes scheduled within the specified time range.
	RenderedSubstitudeEntries []*GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntries `json:"RenderedSubstitudeEntries,omitempty" xml:"RenderedSubstitudeEntries,omitempty" type:"Repeated"`
	// A list of schedule layers.
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
	if s.RenderedFinnalEntries != nil {
		for _, item := range s.RenderedFinnalEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RenderedSubstitudeEntries != nil {
		for _, item := range s.RenderedSubstitudeEntries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ScheduleLayers != nil {
		for _, item := range s.ScheduleLayers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntries struct {
	// The end time of the on-call duty for the contact.
	//
	// example:
	//
	// 2022-10-30
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// Details of the final on-call contact.
	SimpleContact *GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact `json:"SimpleContact,omitempty" xml:"SimpleContact,omitempty" type:"Struct"`
	// The start time of the on-call duty for the contact.
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
	if s.SimpleContact != nil {
		if err := s.SimpleContact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedFinnalEntriesSimpleContact struct {
	// The contact ID.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The contact name.
	//
	// example:
	//
	// 员工1
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
	// The start time of the on-call duty for the contact.
	//
	// example:
	//
	// 2022-10-01
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
	// The end time of the on-call duty for the contact.
	//
	// example:
	//
	// 2022-10-30
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// Details of the on-duty contact.
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
	if s.SimpleContact != nil {
		if err := s.SimpleContact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedLayerEntriesSimpleContact struct {
	// The contact ID.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The contact name.
	//
	// example:
	//
	// 员工1
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
	// The end time of the on-call duty for the substitute.
	//
	// example:
	//
	// 2022-10-30
	End *string `json:"End,omitempty" xml:"End,omitempty"`
	// Details of the substitute.
	SimpleContact *GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact `json:"SimpleContact,omitempty" xml:"SimpleContact,omitempty" type:"Struct"`
	// The start time of the on-call duty for the substitute.
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
	if s.SimpleContact != nil {
		if err := s.SimpleContact.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetOnCallSchedulesDetailResponseBodyDataRenderedSubstitudeEntriesSimpleContact struct {
	// The substitute ID.
	//
	// example:
	//
	// 234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The substitute name.
	//
	// example:
	//
	// 员工2
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
	// A list of contact IDs for the schedule layer.
	ContactIds []*int64 `json:"ContactIds,omitempty" xml:"ContactIds,omitempty" type:"Repeated"`
	// A list of restrictions for the schedule layer.
	Restrictions []*GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions `json:"Restrictions,omitempty" xml:"Restrictions,omitempty" type:"Repeated"`
	// The rotation type. Valid values:
	//
	// - `DAY`: Rotates every day.
	//
	// - `WEEK`: Rotates every week.
	//
	// - `CUSTOM`: Rotates based on a custom schedule.
	//
	// example:
	//
	// DAY
	RotationType *string `json:"RotationType,omitempty" xml:"RotationType,omitempty"`
	// The shift length for the rotation, in hours.
	//
	// example:
	//
	// 8
	ShiftLength *int64 `json:"ShiftLength,omitempty" xml:"ShiftLength,omitempty"`
	// The start time for the rotation.
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
	if s.Restrictions != nil {
		for _, item := range s.Restrictions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetOnCallSchedulesDetailResponseBodyDataScheduleLayersRestrictions struct {
	// The end time for on-call duty each day.
	//
	// example:
	//
	// 18:00
	EndTimeOfDay *string `json:"EndTimeOfDay,omitempty" xml:"EndTimeOfDay,omitempty"`
	// The type of restriction. Valid values:
	//
	// - `daily_restriction`: A daily time-based restriction.
	//
	// - `weekly_restriction`: A weekly time-based restriction.
	//
	// example:
	//
	// daily_restriction
	RestrictionType *string `json:"RestrictionType,omitempty" xml:"RestrictionType,omitempty"`
	// The start time for on-call duty each day.
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
