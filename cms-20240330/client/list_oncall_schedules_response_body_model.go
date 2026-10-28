// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOncallSchedulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOncallSchedules(v []*ListOncallSchedulesResponseBodyOncallSchedules) *ListOncallSchedulesResponseBody
	GetOncallSchedules() []*ListOncallSchedulesResponseBodyOncallSchedules
	SetRequestId(v string) *ListOncallSchedulesResponseBody
	GetRequestId() *string
	SetTotal(v int32) *ListOncallSchedulesResponseBody
	GetTotal() *int32
}

type ListOncallSchedulesResponseBody struct {
	// The list of on-call schedules.
	OncallSchedules []*ListOncallSchedulesResponseBodyOncallSchedules `json:"oncallSchedules,omitempty" xml:"oncallSchedules,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 2
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListOncallSchedulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListOncallSchedulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListOncallSchedulesResponseBody) GetOncallSchedules() []*ListOncallSchedulesResponseBodyOncallSchedules {
	return s.OncallSchedules
}

func (s *ListOncallSchedulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListOncallSchedulesResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *ListOncallSchedulesResponseBody) SetOncallSchedules(v []*ListOncallSchedulesResponseBodyOncallSchedules) *ListOncallSchedulesResponseBody {
	s.OncallSchedules = v
	return s
}

func (s *ListOncallSchedulesResponseBody) SetRequestId(v string) *ListOncallSchedulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListOncallSchedulesResponseBody) SetTotal(v int32) *ListOncallSchedulesResponseBody {
	s.Total = &v
	return s
}

func (s *ListOncallSchedulesResponseBody) Validate() error {
	if s.OncallSchedules != nil {
		for _, item := range s.OncallSchedules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOncallSchedulesResponseBodyOncallSchedules struct {
	// The unique identifier of the on-call schedule.
	//
	// example:
	//
	// testId
	OncallScheduleId *string `json:"oncallScheduleId,omitempty" xml:"oncallScheduleId,omitempty"`
	// The name of the on-call schedule.
	//
	// example:
	//
	// testName
	OncallScheduleName *string `json:"oncallScheduleName,omitempty" xml:"oncallScheduleName,omitempty"`
	// The list of rotations.
	Rotations []*ListOncallSchedulesResponseBodyOncallSchedulesRotations `json:"rotations,omitempty" xml:"rotations,omitempty" type:"Repeated"`
	// The ID of the chatbot that sends shift reminders.
	//
	// example:
	//
	// testId
	ShiftRobotId *string `json:"shiftRobotId,omitempty" xml:"shiftRobotId,omitempty"`
	// The list of substitution rules.
	Substitudes []*ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes `json:"substitudes,omitempty" xml:"substitudes,omitempty" type:"Repeated"`
	// The name of the workspace.
	//
	// example:
	//
	// default-cms-5884500827389836-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s ListOncallSchedulesResponseBodyOncallSchedules) String() string {
	return dara.Prettify(s)
}

func (s ListOncallSchedulesResponseBodyOncallSchedules) GoString() string {
	return s.String()
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) GetOncallScheduleId() *string {
	return s.OncallScheduleId
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) GetOncallScheduleName() *string {
	return s.OncallScheduleName
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) GetRotations() []*ListOncallSchedulesResponseBodyOncallSchedulesRotations {
	return s.Rotations
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) GetShiftRobotId() *string {
	return s.ShiftRobotId
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) GetSubstitudes() []*ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes {
	return s.Substitudes
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) GetWorkspace() *string {
	return s.Workspace
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) SetOncallScheduleId(v string) *ListOncallSchedulesResponseBodyOncallSchedules {
	s.OncallScheduleId = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) SetOncallScheduleName(v string) *ListOncallSchedulesResponseBodyOncallSchedules {
	s.OncallScheduleName = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) SetRotations(v []*ListOncallSchedulesResponseBodyOncallSchedulesRotations) *ListOncallSchedulesResponseBodyOncallSchedules {
	s.Rotations = v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) SetShiftRobotId(v string) *ListOncallSchedulesResponseBodyOncallSchedules {
	s.ShiftRobotId = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) SetSubstitudes(v []*ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) *ListOncallSchedulesResponseBodyOncallSchedules {
	s.Substitudes = v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) SetWorkspace(v string) *ListOncallSchedulesResponseBodyOncallSchedules {
	s.Workspace = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedules) Validate() error {
	if s.Rotations != nil {
		for _, item := range s.Rotations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Substitudes != nil {
		for _, item := range s.Substitudes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListOncallSchedulesResponseBodyOncallSchedulesRotations struct {
	// The days of the week when the rotation is active.
	ActiveDays []*int32 `json:"activeDays,omitempty" xml:"activeDays,omitempty" type:"Repeated"`
	// The IDs of the contacts.
	Contacts []*string `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// The end time of the rotation. Format: HH:mm.
	//
	// example:
	//
	// 24:00
	RotationEndTime *string `json:"rotationEndTime,omitempty" xml:"rotationEndTime,omitempty"`
	// The name of the rotation.
	//
	// example:
	//
	// testName
	RotationName *string `json:"rotationName,omitempty" xml:"rotationName,omitempty"`
	// The start time of the rotation. Format: HH:mm.
	//
	// example:
	//
	// 00:00
	RotationStartTime *string `json:"rotationStartTime,omitempty" xml:"rotationStartTime,omitempty"`
	// The length of the shift interval. Used in conjunction with shiftRecurrenceFrequency. For example, a value of 3 indicates that the on-call personnel rotates every 3 days.
	//
	// example:
	//
	// 3
	ShiftLength *int32 `json:"shiftLength,omitempty" xml:"shiftLength,omitempty"`
	// The unit of the shift interval. Used in conjunction with shiftLength. For example, a value of Week indicates that the on-call personnel rotates every week.
	//
	// example:
	//
	// WEEKLY
	ShiftRecurrenceFrequency *string `json:"shiftRecurrenceFrequency,omitempty" xml:"shiftRecurrenceFrequency,omitempty"`
	// The start date of the rotation. Format: yyyy-MM-dd.
	//
	// example:
	//
	// 2024-01-01
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// The name of the time zone. The IANA Time Zone Database standard is used. Example: Asia/Shanghai.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s ListOncallSchedulesResponseBodyOncallSchedulesRotations) String() string {
	return dara.Prettify(s)
}

func (s ListOncallSchedulesResponseBodyOncallSchedulesRotations) GoString() string {
	return s.String()
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) GetActiveDays() []*int32 {
	return s.ActiveDays
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) GetContacts() []*string {
	return s.Contacts
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) GetRotationEndTime() *string {
	return s.RotationEndTime
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) GetRotationName() *string {
	return s.RotationName
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) GetRotationStartTime() *string {
	return s.RotationStartTime
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) GetShiftLength() *int32 {
	return s.ShiftLength
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) GetShiftRecurrenceFrequency() *string {
	return s.ShiftRecurrenceFrequency
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) GetStartDate() *string {
	return s.StartDate
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) GetTimeZone() *string {
	return s.TimeZone
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) SetActiveDays(v []*int32) *ListOncallSchedulesResponseBodyOncallSchedulesRotations {
	s.ActiveDays = v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) SetContacts(v []*string) *ListOncallSchedulesResponseBodyOncallSchedulesRotations {
	s.Contacts = v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) SetRotationEndTime(v string) *ListOncallSchedulesResponseBodyOncallSchedulesRotations {
	s.RotationEndTime = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) SetRotationName(v string) *ListOncallSchedulesResponseBodyOncallSchedulesRotations {
	s.RotationName = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) SetRotationStartTime(v string) *ListOncallSchedulesResponseBodyOncallSchedulesRotations {
	s.RotationStartTime = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) SetShiftLength(v int32) *ListOncallSchedulesResponseBodyOncallSchedulesRotations {
	s.ShiftLength = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) SetShiftRecurrenceFrequency(v string) *ListOncallSchedulesResponseBodyOncallSchedulesRotations {
	s.ShiftRecurrenceFrequency = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) SetStartDate(v string) *ListOncallSchedulesResponseBodyOncallSchedulesRotations {
	s.StartDate = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) SetTimeZone(v string) *ListOncallSchedulesResponseBodyOncallSchedulesRotations {
	s.TimeZone = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesRotations) Validate() error {
	return dara.Validate(s)
}

type ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes struct {
	// The ID of the substitute contact.
	//
	// example:
	//
	// testId
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	// The end time of the substitution. Format: YYYY-MM-DD HH:mm.
	//
	// example:
	//
	// 2024-10-15 18:00
	EndTime *string `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// The ID of the replaced contact.
	//
	// example:
	//
	// testId
	ReplacedContactId *string `json:"replacedContactId,omitempty" xml:"replacedContactId,omitempty"`
	// The start time of the substitution. Format: YYYY-MM-DD HH:mm.
	//
	// example:
	//
	// 2024-10-10 09:00
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) String() string {
	return dara.Prettify(s)
}

func (s ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) GoString() string {
	return s.String()
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) GetContactId() *string {
	return s.ContactId
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) GetEndTime() *string {
	return s.EndTime
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) GetReplacedContactId() *string {
	return s.ReplacedContactId
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) GetStartTime() *string {
	return s.StartTime
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) SetContactId(v string) *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes {
	s.ContactId = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) SetEndTime(v string) *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes {
	s.EndTime = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) SetReplacedContactId(v string) *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes {
	s.ReplacedContactId = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) SetStartTime(v string) *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes {
	s.StartTime = &v
	return s
}

func (s *ListOncallSchedulesResponseBodyOncallSchedulesSubstitudes) Validate() error {
	return dara.Validate(s)
}
