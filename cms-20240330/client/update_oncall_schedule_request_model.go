// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOncallScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOncallScheduleName(v string) *UpdateOncallScheduleRequest
	GetOncallScheduleName() *string
	SetRotations(v []*UpdateOncallScheduleRequestRotations) *UpdateOncallScheduleRequest
	GetRotations() []*UpdateOncallScheduleRequestRotations
	SetShiftRobotId(v string) *UpdateOncallScheduleRequest
	GetShiftRobotId() *string
	SetSubstitudes(v []*UpdateOncallScheduleRequestSubstitudes) *UpdateOncallScheduleRequest
	GetSubstitudes() []*UpdateOncallScheduleRequestSubstitudes
}

type UpdateOncallScheduleRequest struct {
	// The name of the on-call schedule.
	//
	// example:
	//
	// testName
	OncallScheduleName *string `json:"oncallScheduleName,omitempty" xml:"oncallScheduleName,omitempty"`
	// The list of rotations.
	Rotations []*UpdateOncallScheduleRequestRotations `json:"rotations,omitempty" xml:"rotations,omitempty" type:"Repeated"`
	// The ID of the shift reminder chatbot.
	//
	// example:
	//
	// testId
	ShiftRobotId *string `json:"shiftRobotId,omitempty" xml:"shiftRobotId,omitempty"`
	// The list of substitution rules.
	Substitudes []*UpdateOncallScheduleRequestSubstitudes `json:"substitudes,omitempty" xml:"substitudes,omitempty" type:"Repeated"`
}

func (s UpdateOncallScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateOncallScheduleRequest) GoString() string {
	return s.String()
}

func (s *UpdateOncallScheduleRequest) GetOncallScheduleName() *string {
	return s.OncallScheduleName
}

func (s *UpdateOncallScheduleRequest) GetRotations() []*UpdateOncallScheduleRequestRotations {
	return s.Rotations
}

func (s *UpdateOncallScheduleRequest) GetShiftRobotId() *string {
	return s.ShiftRobotId
}

func (s *UpdateOncallScheduleRequest) GetSubstitudes() []*UpdateOncallScheduleRequestSubstitudes {
	return s.Substitudes
}

func (s *UpdateOncallScheduleRequest) SetOncallScheduleName(v string) *UpdateOncallScheduleRequest {
	s.OncallScheduleName = &v
	return s
}

func (s *UpdateOncallScheduleRequest) SetRotations(v []*UpdateOncallScheduleRequestRotations) *UpdateOncallScheduleRequest {
	s.Rotations = v
	return s
}

func (s *UpdateOncallScheduleRequest) SetShiftRobotId(v string) *UpdateOncallScheduleRequest {
	s.ShiftRobotId = &v
	return s
}

func (s *UpdateOncallScheduleRequest) SetSubstitudes(v []*UpdateOncallScheduleRequestSubstitudes) *UpdateOncallScheduleRequest {
	s.Substitudes = v
	return s
}

func (s *UpdateOncallScheduleRequest) Validate() error {
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

type UpdateOncallScheduleRequestRotations struct {
	// The days of the week when the rotation takes effect.
	//
	// This parameter is required.
	ActiveDays []*int32 `json:"activeDays,omitempty" xml:"activeDays,omitempty" type:"Repeated"`
	// The IDs of the contacts.
	//
	// This parameter is required.
	Contacts []*string `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// The end time of the rotation. Format: HH:mm.
	//
	// This parameter is required.
	//
	// example:
	//
	// 24:00
	RotationEndTime *string `json:"rotationEndTime,omitempty" xml:"rotationEndTime,omitempty"`
	// The name of the rotation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	RotationName *string `json:"rotationName,omitempty" xml:"rotationName,omitempty"`
	// The start time of the rotation. Format: HH:mm.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00:00
	RotationStartTime *string `json:"rotationStartTime,omitempty" xml:"rotationStartTime,omitempty"`
	// The length of the shift cycle. Used in conjunction with shiftRecurrenceFrequency, such as switching to the next on-call person every 3 days.
	//
	// example:
	//
	// 3
	ShiftLength *int32 `json:"shiftLength,omitempty" xml:"shiftLength,omitempty"`
	// The unit of the shift cycle. Used in conjunction with shiftRecurrenceFrequency, such as switching to the next on-call person every week.
	//
	// This parameter is required.
	//
	// example:
	//
	// WEEKLY
	ShiftRecurrenceFrequency *string `json:"shiftRecurrenceFrequency,omitempty" xml:"shiftRecurrenceFrequency,omitempty"`
	// The start date of the rotation. Format: yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2024-04-04
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// The time zone name. Uses the IANA Time Zone Database standard, such as Asia/Shanghai.
	//
	// This parameter is required.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s UpdateOncallScheduleRequestRotations) String() string {
	return dara.Prettify(s)
}

func (s UpdateOncallScheduleRequestRotations) GoString() string {
	return s.String()
}

func (s *UpdateOncallScheduleRequestRotations) GetActiveDays() []*int32 {
	return s.ActiveDays
}

func (s *UpdateOncallScheduleRequestRotations) GetContacts() []*string {
	return s.Contacts
}

func (s *UpdateOncallScheduleRequestRotations) GetRotationEndTime() *string {
	return s.RotationEndTime
}

func (s *UpdateOncallScheduleRequestRotations) GetRotationName() *string {
	return s.RotationName
}

func (s *UpdateOncallScheduleRequestRotations) GetRotationStartTime() *string {
	return s.RotationStartTime
}

func (s *UpdateOncallScheduleRequestRotations) GetShiftLength() *int32 {
	return s.ShiftLength
}

func (s *UpdateOncallScheduleRequestRotations) GetShiftRecurrenceFrequency() *string {
	return s.ShiftRecurrenceFrequency
}

func (s *UpdateOncallScheduleRequestRotations) GetStartDate() *string {
	return s.StartDate
}

func (s *UpdateOncallScheduleRequestRotations) GetTimeZone() *string {
	return s.TimeZone
}

func (s *UpdateOncallScheduleRequestRotations) SetActiveDays(v []*int32) *UpdateOncallScheduleRequestRotations {
	s.ActiveDays = v
	return s
}

func (s *UpdateOncallScheduleRequestRotations) SetContacts(v []*string) *UpdateOncallScheduleRequestRotations {
	s.Contacts = v
	return s
}

func (s *UpdateOncallScheduleRequestRotations) SetRotationEndTime(v string) *UpdateOncallScheduleRequestRotations {
	s.RotationEndTime = &v
	return s
}

func (s *UpdateOncallScheduleRequestRotations) SetRotationName(v string) *UpdateOncallScheduleRequestRotations {
	s.RotationName = &v
	return s
}

func (s *UpdateOncallScheduleRequestRotations) SetRotationStartTime(v string) *UpdateOncallScheduleRequestRotations {
	s.RotationStartTime = &v
	return s
}

func (s *UpdateOncallScheduleRequestRotations) SetShiftLength(v int32) *UpdateOncallScheduleRequestRotations {
	s.ShiftLength = &v
	return s
}

func (s *UpdateOncallScheduleRequestRotations) SetShiftRecurrenceFrequency(v string) *UpdateOncallScheduleRequestRotations {
	s.ShiftRecurrenceFrequency = &v
	return s
}

func (s *UpdateOncallScheduleRequestRotations) SetStartDate(v string) *UpdateOncallScheduleRequestRotations {
	s.StartDate = &v
	return s
}

func (s *UpdateOncallScheduleRequestRotations) SetTimeZone(v string) *UpdateOncallScheduleRequestRotations {
	s.TimeZone = &v
	return s
}

func (s *UpdateOncallScheduleRequestRotations) Validate() error {
	return dara.Validate(s)
}

type UpdateOncallScheduleRequestSubstitudes struct {
	// The contact ID of the substitute.
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
	// The contact ID of the replaced person.
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

func (s UpdateOncallScheduleRequestSubstitudes) String() string {
	return dara.Prettify(s)
}

func (s UpdateOncallScheduleRequestSubstitudes) GoString() string {
	return s.String()
}

func (s *UpdateOncallScheduleRequestSubstitudes) GetContactId() *string {
	return s.ContactId
}

func (s *UpdateOncallScheduleRequestSubstitudes) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateOncallScheduleRequestSubstitudes) GetReplacedContactId() *string {
	return s.ReplacedContactId
}

func (s *UpdateOncallScheduleRequestSubstitudes) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateOncallScheduleRequestSubstitudes) SetContactId(v string) *UpdateOncallScheduleRequestSubstitudes {
	s.ContactId = &v
	return s
}

func (s *UpdateOncallScheduleRequestSubstitudes) SetEndTime(v string) *UpdateOncallScheduleRequestSubstitudes {
	s.EndTime = &v
	return s
}

func (s *UpdateOncallScheduleRequestSubstitudes) SetReplacedContactId(v string) *UpdateOncallScheduleRequestSubstitudes {
	s.ReplacedContactId = &v
	return s
}

func (s *UpdateOncallScheduleRequestSubstitudes) SetStartTime(v string) *UpdateOncallScheduleRequestSubstitudes {
	s.StartTime = &v
	return s
}

func (s *UpdateOncallScheduleRequestSubstitudes) Validate() error {
	return dara.Validate(s)
}
