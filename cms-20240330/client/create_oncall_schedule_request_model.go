// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOncallScheduleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOncallScheduleId(v string) *CreateOncallScheduleRequest
	GetOncallScheduleId() *string
	SetOncallScheduleName(v string) *CreateOncallScheduleRequest
	GetOncallScheduleName() *string
	SetRotations(v []*CreateOncallScheduleRequestRotations) *CreateOncallScheduleRequest
	GetRotations() []*CreateOncallScheduleRequestRotations
	SetShiftRobotId(v string) *CreateOncallScheduleRequest
	GetShiftRobotId() *string
	SetSubstitudes(v []*CreateOncallScheduleRequestSubstitudes) *CreateOncallScheduleRequest
	GetSubstitudes() []*CreateOncallScheduleRequestSubstitudes
	SetWorkspace(v string) *CreateOncallScheduleRequest
	GetWorkspace() *string
}

type CreateOncallScheduleRequest struct {
	// The unique identifier of the on-call schedule.
	//
	// example:
	//
	// testOncallScheduleId
	OncallScheduleId *string `json:"oncallScheduleId,omitempty" xml:"oncallScheduleId,omitempty"`
	// The name of the on-call schedule.
	//
	// example:
	//
	// testOncallScheduleName
	OncallScheduleName *string `json:"oncallScheduleName,omitempty" xml:"oncallScheduleName,omitempty"`
	// The list of rotations.
	Rotations []*CreateOncallScheduleRequestRotations `json:"rotations,omitempty" xml:"rotations,omitempty" type:"Repeated"`
	// The ID of the chatbot for shift reminders.
	//
	// example:
	//
	// testRobotId
	ShiftRobotId *string `json:"shiftRobotId,omitempty" xml:"shiftRobotId,omitempty"`
	// The list of substitution rules.
	Substitudes []*CreateOncallScheduleRequestSubstitudes `json:"substitudes,omitempty" xml:"substitudes,omitempty" type:"Repeated"`
	// The name of the workspace.
	//
	// example:
	//
	// default-cms-5884500827389836-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreateOncallScheduleRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOncallScheduleRequest) GoString() string {
	return s.String()
}

func (s *CreateOncallScheduleRequest) GetOncallScheduleId() *string {
	return s.OncallScheduleId
}

func (s *CreateOncallScheduleRequest) GetOncallScheduleName() *string {
	return s.OncallScheduleName
}

func (s *CreateOncallScheduleRequest) GetRotations() []*CreateOncallScheduleRequestRotations {
	return s.Rotations
}

func (s *CreateOncallScheduleRequest) GetShiftRobotId() *string {
	return s.ShiftRobotId
}

func (s *CreateOncallScheduleRequest) GetSubstitudes() []*CreateOncallScheduleRequestSubstitudes {
	return s.Substitudes
}

func (s *CreateOncallScheduleRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreateOncallScheduleRequest) SetOncallScheduleId(v string) *CreateOncallScheduleRequest {
	s.OncallScheduleId = &v
	return s
}

func (s *CreateOncallScheduleRequest) SetOncallScheduleName(v string) *CreateOncallScheduleRequest {
	s.OncallScheduleName = &v
	return s
}

func (s *CreateOncallScheduleRequest) SetRotations(v []*CreateOncallScheduleRequestRotations) *CreateOncallScheduleRequest {
	s.Rotations = v
	return s
}

func (s *CreateOncallScheduleRequest) SetShiftRobotId(v string) *CreateOncallScheduleRequest {
	s.ShiftRobotId = &v
	return s
}

func (s *CreateOncallScheduleRequest) SetSubstitudes(v []*CreateOncallScheduleRequestSubstitudes) *CreateOncallScheduleRequest {
	s.Substitudes = v
	return s
}

func (s *CreateOncallScheduleRequest) SetWorkspace(v string) *CreateOncallScheduleRequest {
	s.Workspace = &v
	return s
}

func (s *CreateOncallScheduleRequest) Validate() error {
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

type CreateOncallScheduleRequestRotations struct {
	// The days of the week when the rotation is active.
	//
	// This parameter is required.
	ActiveDays []*int32 `json:"activeDays,omitempty" xml:"activeDays,omitempty" type:"Repeated"`
	// The list of contacts.
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
	// testRotationName
	RotationName *string `json:"rotationName,omitempty" xml:"rotationName,omitempty"`
	// The start time of the rotation. Format: HH:mm.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00:00
	RotationStartTime *string `json:"rotationStartTime,omitempty" xml:"rotationStartTime,omitempty"`
	// The length of the shift interval. Use this parameter in conjunction with shiftRecurrenceFrequency. For example, set this parameter to 3 to rotate to the next on-call contact every 3 days.
	//
	// example:
	//
	// 1
	ShiftLength *int32 `json:"shiftLength,omitempty" xml:"shiftLength,omitempty"`
	// The unit of the shift interval. Use this parameter in conjunction with shiftRecurrenceFrequency. For example, set this parameter to week to rotate to the next on-call contact every week.
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
	// 2024-08-14
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// The name of the time zone. The time zone must follow the **IANA Time Zone Database*	- standard, such as `Asia/Shanghai`.
	//
	// This parameter is required.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s CreateOncallScheduleRequestRotations) String() string {
	return dara.Prettify(s)
}

func (s CreateOncallScheduleRequestRotations) GoString() string {
	return s.String()
}

func (s *CreateOncallScheduleRequestRotations) GetActiveDays() []*int32 {
	return s.ActiveDays
}

func (s *CreateOncallScheduleRequestRotations) GetContacts() []*string {
	return s.Contacts
}

func (s *CreateOncallScheduleRequestRotations) GetRotationEndTime() *string {
	return s.RotationEndTime
}

func (s *CreateOncallScheduleRequestRotations) GetRotationName() *string {
	return s.RotationName
}

func (s *CreateOncallScheduleRequestRotations) GetRotationStartTime() *string {
	return s.RotationStartTime
}

func (s *CreateOncallScheduleRequestRotations) GetShiftLength() *int32 {
	return s.ShiftLength
}

func (s *CreateOncallScheduleRequestRotations) GetShiftRecurrenceFrequency() *string {
	return s.ShiftRecurrenceFrequency
}

func (s *CreateOncallScheduleRequestRotations) GetStartDate() *string {
	return s.StartDate
}

func (s *CreateOncallScheduleRequestRotations) GetTimeZone() *string {
	return s.TimeZone
}

func (s *CreateOncallScheduleRequestRotations) SetActiveDays(v []*int32) *CreateOncallScheduleRequestRotations {
	s.ActiveDays = v
	return s
}

func (s *CreateOncallScheduleRequestRotations) SetContacts(v []*string) *CreateOncallScheduleRequestRotations {
	s.Contacts = v
	return s
}

func (s *CreateOncallScheduleRequestRotations) SetRotationEndTime(v string) *CreateOncallScheduleRequestRotations {
	s.RotationEndTime = &v
	return s
}

func (s *CreateOncallScheduleRequestRotations) SetRotationName(v string) *CreateOncallScheduleRequestRotations {
	s.RotationName = &v
	return s
}

func (s *CreateOncallScheduleRequestRotations) SetRotationStartTime(v string) *CreateOncallScheduleRequestRotations {
	s.RotationStartTime = &v
	return s
}

func (s *CreateOncallScheduleRequestRotations) SetShiftLength(v int32) *CreateOncallScheduleRequestRotations {
	s.ShiftLength = &v
	return s
}

func (s *CreateOncallScheduleRequestRotations) SetShiftRecurrenceFrequency(v string) *CreateOncallScheduleRequestRotations {
	s.ShiftRecurrenceFrequency = &v
	return s
}

func (s *CreateOncallScheduleRequestRotations) SetStartDate(v string) *CreateOncallScheduleRequestRotations {
	s.StartDate = &v
	return s
}

func (s *CreateOncallScheduleRequestRotations) SetTimeZone(v string) *CreateOncallScheduleRequestRotations {
	s.TimeZone = &v
	return s
}

func (s *CreateOncallScheduleRequestRotations) Validate() error {
	return dara.Validate(s)
}

type CreateOncallScheduleRequestSubstitudes struct {
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

func (s CreateOncallScheduleRequestSubstitudes) String() string {
	return dara.Prettify(s)
}

func (s CreateOncallScheduleRequestSubstitudes) GoString() string {
	return s.String()
}

func (s *CreateOncallScheduleRequestSubstitudes) GetContactId() *string {
	return s.ContactId
}

func (s *CreateOncallScheduleRequestSubstitudes) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateOncallScheduleRequestSubstitudes) GetReplacedContactId() *string {
	return s.ReplacedContactId
}

func (s *CreateOncallScheduleRequestSubstitudes) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateOncallScheduleRequestSubstitudes) SetContactId(v string) *CreateOncallScheduleRequestSubstitudes {
	s.ContactId = &v
	return s
}

func (s *CreateOncallScheduleRequestSubstitudes) SetEndTime(v string) *CreateOncallScheduleRequestSubstitudes {
	s.EndTime = &v
	return s
}

func (s *CreateOncallScheduleRequestSubstitudes) SetReplacedContactId(v string) *CreateOncallScheduleRequestSubstitudes {
	s.ReplacedContactId = &v
	return s
}

func (s *CreateOncallScheduleRequestSubstitudes) SetStartTime(v string) *CreateOncallScheduleRequestSubstitudes {
	s.StartTime = &v
	return s
}

func (s *CreateOncallScheduleRequestSubstitudes) Validate() error {
	return dara.Validate(s)
}
