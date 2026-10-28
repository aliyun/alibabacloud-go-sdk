// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOncallScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOncallScheduleId(v string) *GetOncallScheduleResponseBody
	GetOncallScheduleId() *string
	SetOncallScheduleName(v string) *GetOncallScheduleResponseBody
	GetOncallScheduleName() *string
	SetRequestId(v string) *GetOncallScheduleResponseBody
	GetRequestId() *string
	SetRotations(v []*GetOncallScheduleResponseBodyRotations) *GetOncallScheduleResponseBody
	GetRotations() []*GetOncallScheduleResponseBodyRotations
	SetShiftRobotId(v string) *GetOncallScheduleResponseBody
	GetShiftRobotId() *string
	SetSubstitudes(v []*GetOncallScheduleResponseBodySubstitudes) *GetOncallScheduleResponseBody
	GetSubstitudes() []*GetOncallScheduleResponseBodySubstitudes
	SetWorkspace(v string) *GetOncallScheduleResponseBody
	GetWorkspace() *string
}

type GetOncallScheduleResponseBody struct {
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
	// testOncallScheduleName
	OncallScheduleName *string `json:"oncallScheduleName,omitempty" xml:"oncallScheduleName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of rotations.
	Rotations []*GetOncallScheduleResponseBodyRotations `json:"rotations,omitempty" xml:"rotations,omitempty" type:"Repeated"`
	// The ID of the shift reminder chatbot.
	//
	// example:
	//
	// testId
	ShiftRobotId *string `json:"shiftRobotId,omitempty" xml:"shiftRobotId,omitempty"`
	// The list of substitution rules.
	Substitudes []*GetOncallScheduleResponseBodySubstitudes `json:"substitudes,omitempty" xml:"substitudes,omitempty" type:"Repeated"`
	// The name of the workspace.
	//
	// example:
	//
	// default-cms-5884500827389836-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetOncallScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOncallScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetOncallScheduleResponseBody) GetOncallScheduleId() *string {
	return s.OncallScheduleId
}

func (s *GetOncallScheduleResponseBody) GetOncallScheduleName() *string {
	return s.OncallScheduleName
}

func (s *GetOncallScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOncallScheduleResponseBody) GetRotations() []*GetOncallScheduleResponseBodyRotations {
	return s.Rotations
}

func (s *GetOncallScheduleResponseBody) GetShiftRobotId() *string {
	return s.ShiftRobotId
}

func (s *GetOncallScheduleResponseBody) GetSubstitudes() []*GetOncallScheduleResponseBodySubstitudes {
	return s.Substitudes
}

func (s *GetOncallScheduleResponseBody) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetOncallScheduleResponseBody) SetOncallScheduleId(v string) *GetOncallScheduleResponseBody {
	s.OncallScheduleId = &v
	return s
}

func (s *GetOncallScheduleResponseBody) SetOncallScheduleName(v string) *GetOncallScheduleResponseBody {
	s.OncallScheduleName = &v
	return s
}

func (s *GetOncallScheduleResponseBody) SetRequestId(v string) *GetOncallScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOncallScheduleResponseBody) SetRotations(v []*GetOncallScheduleResponseBodyRotations) *GetOncallScheduleResponseBody {
	s.Rotations = v
	return s
}

func (s *GetOncallScheduleResponseBody) SetShiftRobotId(v string) *GetOncallScheduleResponseBody {
	s.ShiftRobotId = &v
	return s
}

func (s *GetOncallScheduleResponseBody) SetSubstitudes(v []*GetOncallScheduleResponseBodySubstitudes) *GetOncallScheduleResponseBody {
	s.Substitudes = v
	return s
}

func (s *GetOncallScheduleResponseBody) SetWorkspace(v string) *GetOncallScheduleResponseBody {
	s.Workspace = &v
	return s
}

func (s *GetOncallScheduleResponseBody) Validate() error {
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

type GetOncallScheduleResponseBodyRotations struct {
	// The days of the week when the rotation is active.
	ActiveDays []*int32 `json:"activeDays,omitempty" xml:"activeDays,omitempty" type:"Repeated"`
	// The contact ID.
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
	// The length of the shift rotation cycle. Used in combination with shiftRecurrenceFrequency. For example, set this parameter to 3 to rotate to the next on-call contact every 3 days.
	//
	// example:
	//
	// 4
	ShiftLength *int32 `json:"shiftLength,omitempty" xml:"shiftLength,omitempty"`
	// The unit of the shift rotation cycle. Used in combination with shiftRecurrenceFrequency. For example, set this parameter to Week to rotate to the next on-call contact every week.
	//
	// example:
	//
	// WEEKLY
	ShiftRecurrenceFrequency *string `json:"shiftRecurrenceFrequency,omitempty" xml:"shiftRecurrenceFrequency,omitempty"`
	// The start date of the rotation. Format: yyyy-MM-dd.
	//
	// example:
	//
	// 2024-10-10
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// The time zone name based on the IANA Time Zone Database standard. Example: Asia/Shanghai.
	//
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"timeZone,omitempty" xml:"timeZone,omitempty"`
}

func (s GetOncallScheduleResponseBodyRotations) String() string {
	return dara.Prettify(s)
}

func (s GetOncallScheduleResponseBodyRotations) GoString() string {
	return s.String()
}

func (s *GetOncallScheduleResponseBodyRotations) GetActiveDays() []*int32 {
	return s.ActiveDays
}

func (s *GetOncallScheduleResponseBodyRotations) GetContacts() []*string {
	return s.Contacts
}

func (s *GetOncallScheduleResponseBodyRotations) GetRotationEndTime() *string {
	return s.RotationEndTime
}

func (s *GetOncallScheduleResponseBodyRotations) GetRotationName() *string {
	return s.RotationName
}

func (s *GetOncallScheduleResponseBodyRotations) GetRotationStartTime() *string {
	return s.RotationStartTime
}

func (s *GetOncallScheduleResponseBodyRotations) GetShiftLength() *int32 {
	return s.ShiftLength
}

func (s *GetOncallScheduleResponseBodyRotations) GetShiftRecurrenceFrequency() *string {
	return s.ShiftRecurrenceFrequency
}

func (s *GetOncallScheduleResponseBodyRotations) GetStartDate() *string {
	return s.StartDate
}

func (s *GetOncallScheduleResponseBodyRotations) GetTimeZone() *string {
	return s.TimeZone
}

func (s *GetOncallScheduleResponseBodyRotations) SetActiveDays(v []*int32) *GetOncallScheduleResponseBodyRotations {
	s.ActiveDays = v
	return s
}

func (s *GetOncallScheduleResponseBodyRotations) SetContacts(v []*string) *GetOncallScheduleResponseBodyRotations {
	s.Contacts = v
	return s
}

func (s *GetOncallScheduleResponseBodyRotations) SetRotationEndTime(v string) *GetOncallScheduleResponseBodyRotations {
	s.RotationEndTime = &v
	return s
}

func (s *GetOncallScheduleResponseBodyRotations) SetRotationName(v string) *GetOncallScheduleResponseBodyRotations {
	s.RotationName = &v
	return s
}

func (s *GetOncallScheduleResponseBodyRotations) SetRotationStartTime(v string) *GetOncallScheduleResponseBodyRotations {
	s.RotationStartTime = &v
	return s
}

func (s *GetOncallScheduleResponseBodyRotations) SetShiftLength(v int32) *GetOncallScheduleResponseBodyRotations {
	s.ShiftLength = &v
	return s
}

func (s *GetOncallScheduleResponseBodyRotations) SetShiftRecurrenceFrequency(v string) *GetOncallScheduleResponseBodyRotations {
	s.ShiftRecurrenceFrequency = &v
	return s
}

func (s *GetOncallScheduleResponseBodyRotations) SetStartDate(v string) *GetOncallScheduleResponseBodyRotations {
	s.StartDate = &v
	return s
}

func (s *GetOncallScheduleResponseBodyRotations) SetTimeZone(v string) *GetOncallScheduleResponseBodyRotations {
	s.TimeZone = &v
	return s
}

func (s *GetOncallScheduleResponseBodyRotations) Validate() error {
	return dara.Validate(s)
}

type GetOncallScheduleResponseBodySubstitudes struct {
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
	// The contact ID of the replaced contact.
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

func (s GetOncallScheduleResponseBodySubstitudes) String() string {
	return dara.Prettify(s)
}

func (s GetOncallScheduleResponseBodySubstitudes) GoString() string {
	return s.String()
}

func (s *GetOncallScheduleResponseBodySubstitudes) GetContactId() *string {
	return s.ContactId
}

func (s *GetOncallScheduleResponseBodySubstitudes) GetEndTime() *string {
	return s.EndTime
}

func (s *GetOncallScheduleResponseBodySubstitudes) GetReplacedContactId() *string {
	return s.ReplacedContactId
}

func (s *GetOncallScheduleResponseBodySubstitudes) GetStartTime() *string {
	return s.StartTime
}

func (s *GetOncallScheduleResponseBodySubstitudes) SetContactId(v string) *GetOncallScheduleResponseBodySubstitudes {
	s.ContactId = &v
	return s
}

func (s *GetOncallScheduleResponseBodySubstitudes) SetEndTime(v string) *GetOncallScheduleResponseBodySubstitudes {
	s.EndTime = &v
	return s
}

func (s *GetOncallScheduleResponseBodySubstitudes) SetReplacedContactId(v string) *GetOncallScheduleResponseBodySubstitudes {
	s.ReplacedContactId = &v
	return s
}

func (s *GetOncallScheduleResponseBodySubstitudes) SetStartTime(v string) *GetOncallScheduleResponseBodySubstitudes {
	s.StartTime = &v
	return s
}

func (s *GetOncallScheduleResponseBodySubstitudes) Validate() error {
	return dara.Validate(s)
}
