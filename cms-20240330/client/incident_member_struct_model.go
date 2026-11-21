// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncidentMemberStruct interface {
	dara.Model
	String() string
	GoString() string
	SetAcknowledge(v *IncidentMemberStructAcknowledge) *IncidentMemberStruct
	GetAcknowledge() *IncidentMemberStructAcknowledge
	SetContactId(v string) *IncidentMemberStruct
	GetContactId() *string
	SetContacts(v []*IncidentMemberStructContacts) *IncidentMemberStruct
	GetContacts() []*IncidentMemberStructContacts
	SetEscalation(v *IncidentMemberStructEscalation) *IncidentMemberStruct
	GetEscalation() *IncidentMemberStructEscalation
	SetIncidentId(v string) *IncidentMemberStruct
	GetIncidentId() *string
	SetIncidentMemberId(v string) *IncidentMemberStruct
	GetIncidentMemberId() *string
	SetScheduleGroup(v *IncidentMemberStructScheduleGroup) *IncidentMemberStruct
	GetScheduleGroup() *IncidentMemberStructScheduleGroup
	SetTime(v int64) *IncidentMemberStruct
	GetTime() *int64
	SetUserId(v int64) *IncidentMemberStruct
	GetUserId() *int64
}

type IncidentMemberStruct struct {
	// The acknowledgment information struct.
	Acknowledge *IncidentMemberStructAcknowledge `json:"acknowledge,omitempty" xml:"acknowledge,omitempty" type:"Struct"`
	// The contact ID.
	//
	// example:
	//
	// user-12345
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	// The contact list.
	Contacts []*IncidentMemberStructContacts `json:"contacts,omitempty" xml:"contacts,omitempty" type:"Repeated"`
	// The escalation stage information.
	Escalation *IncidentMemberStructEscalation `json:"escalation,omitempty" xml:"escalation,omitempty" type:"Struct"`
	// The incident ID.
	//
	// example:
	//
	// incident-001
	IncidentId *string `json:"incidentId,omitempty" xml:"incidentId,omitempty"`
	// The incident member ID.
	//
	// example:
	//
	// member-001
	IncidentMemberId *string `json:"incidentMemberId,omitempty" xml:"incidentMemberId,omitempty"`
	// The scheduling group.
	ScheduleGroup *IncidentMemberStructScheduleGroup `json:"scheduleGroup,omitempty" xml:"scheduleGroup,omitempty" type:"Struct"`
	// The time.
	//
	// example:
	//
	// 1794984221885
	Time *int64 `json:"time,omitempty" xml:"time,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 4123456
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s IncidentMemberStruct) String() string {
	return dara.Prettify(s)
}

func (s IncidentMemberStruct) GoString() string {
	return s.String()
}

func (s *IncidentMemberStruct) GetAcknowledge() *IncidentMemberStructAcknowledge {
	return s.Acknowledge
}

func (s *IncidentMemberStruct) GetContactId() *string {
	return s.ContactId
}

func (s *IncidentMemberStruct) GetContacts() []*IncidentMemberStructContacts {
	return s.Contacts
}

func (s *IncidentMemberStruct) GetEscalation() *IncidentMemberStructEscalation {
	return s.Escalation
}

func (s *IncidentMemberStruct) GetIncidentId() *string {
	return s.IncidentId
}

func (s *IncidentMemberStruct) GetIncidentMemberId() *string {
	return s.IncidentMemberId
}

func (s *IncidentMemberStruct) GetScheduleGroup() *IncidentMemberStructScheduleGroup {
	return s.ScheduleGroup
}

func (s *IncidentMemberStruct) GetTime() *int64 {
	return s.Time
}

func (s *IncidentMemberStruct) GetUserId() *int64 {
	return s.UserId
}

func (s *IncidentMemberStruct) SetAcknowledge(v *IncidentMemberStructAcknowledge) *IncidentMemberStruct {
	s.Acknowledge = v
	return s
}

func (s *IncidentMemberStruct) SetContactId(v string) *IncidentMemberStruct {
	s.ContactId = &v
	return s
}

func (s *IncidentMemberStruct) SetContacts(v []*IncidentMemberStructContacts) *IncidentMemberStruct {
	s.Contacts = v
	return s
}

func (s *IncidentMemberStruct) SetEscalation(v *IncidentMemberStructEscalation) *IncidentMemberStruct {
	s.Escalation = v
	return s
}

func (s *IncidentMemberStruct) SetIncidentId(v string) *IncidentMemberStruct {
	s.IncidentId = &v
	return s
}

func (s *IncidentMemberStruct) SetIncidentMemberId(v string) *IncidentMemberStruct {
	s.IncidentMemberId = &v
	return s
}

func (s *IncidentMemberStruct) SetScheduleGroup(v *IncidentMemberStructScheduleGroup) *IncidentMemberStruct {
	s.ScheduleGroup = v
	return s
}

func (s *IncidentMemberStruct) SetTime(v int64) *IncidentMemberStruct {
	s.Time = &v
	return s
}

func (s *IncidentMemberStruct) SetUserId(v int64) *IncidentMemberStruct {
	s.UserId = &v
	return s
}

func (s *IncidentMemberStruct) Validate() error {
	if s.Acknowledge != nil {
		if err := s.Acknowledge.Validate(); err != nil {
			return err
		}
	}
	if s.Contacts != nil {
		for _, item := range s.Contacts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Escalation != nil {
		if err := s.Escalation.Validate(); err != nil {
			return err
		}
	}
	if s.ScheduleGroup != nil {
		if err := s.ScheduleGroup.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IncidentMemberStructAcknowledge struct {
	// The acknowledgment level.
	//
	// example:
	//
	// 2
	BreakLevel *string `json:"breakLevel,omitempty" xml:"breakLevel,omitempty"`
	// The verification time.
	//
	// example:
	//
	// 1741234567890
	VerifyTime *int64 `json:"verifyTime,omitempty" xml:"verifyTime,omitempty"`
}

func (s IncidentMemberStructAcknowledge) String() string {
	return dara.Prettify(s)
}

func (s IncidentMemberStructAcknowledge) GoString() string {
	return s.String()
}

func (s *IncidentMemberStructAcknowledge) GetBreakLevel() *string {
	return s.BreakLevel
}

func (s *IncidentMemberStructAcknowledge) GetVerifyTime() *int64 {
	return s.VerifyTime
}

func (s *IncidentMemberStructAcknowledge) SetBreakLevel(v string) *IncidentMemberStructAcknowledge {
	s.BreakLevel = &v
	return s
}

func (s *IncidentMemberStructAcknowledge) SetVerifyTime(v int64) *IncidentMemberStructAcknowledge {
	s.VerifyTime = &v
	return s
}

func (s *IncidentMemberStructAcknowledge) Validate() error {
	return dara.Validate(s)
}

type IncidentMemberStructContacts struct {
	// The channel.
	//
	// example:
	//
	// dingtalk
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// The contact ID (masked).
	//
	// example:
	//
	// user-12345
	ContactMask *string `json:"contactMask,omitempty" xml:"contactMask,omitempty"`
}

func (s IncidentMemberStructContacts) String() string {
	return dara.Prettify(s)
}

func (s IncidentMemberStructContacts) GoString() string {
	return s.String()
}

func (s *IncidentMemberStructContacts) GetChannel() *string {
	return s.Channel
}

func (s *IncidentMemberStructContacts) GetContactMask() *string {
	return s.ContactMask
}

func (s *IncidentMemberStructContacts) SetChannel(v string) *IncidentMemberStructContacts {
	s.Channel = &v
	return s
}

func (s *IncidentMemberStructContacts) SetContactMask(v string) *IncidentMemberStructContacts {
	s.ContactMask = &v
	return s
}

func (s *IncidentMemberStructContacts) Validate() error {
	return dara.Validate(s)
}

type IncidentMemberStructEscalation struct {
	// The description.
	//
	// example:
	//
	// Notify the O&M team
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The incident escalation ID.
	//
	// example:
	//
	// escalation-001
	IncidentEscalationId *string `json:"incidentEscalationId,omitempty" xml:"incidentEscalationId,omitempty"`
	// The name.
	//
	// example:
	//
	// Stage 1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The escalation stage index.
	//
	// example:
	//
	// 1
	StageIndex *string `json:"stageIndex,omitempty" xml:"stageIndex,omitempty"`
	// The title.
	//
	// example:
	//
	// Alert escalated to the O&M team
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s IncidentMemberStructEscalation) String() string {
	return dara.Prettify(s)
}

func (s IncidentMemberStructEscalation) GoString() string {
	return s.String()
}

func (s *IncidentMemberStructEscalation) GetDescription() *string {
	return s.Description
}

func (s *IncidentMemberStructEscalation) GetIncidentEscalationId() *string {
	return s.IncidentEscalationId
}

func (s *IncidentMemberStructEscalation) GetName() *string {
	return s.Name
}

func (s *IncidentMemberStructEscalation) GetStageIndex() *string {
	return s.StageIndex
}

func (s *IncidentMemberStructEscalation) GetTitle() *string {
	return s.Title
}

func (s *IncidentMemberStructEscalation) SetDescription(v string) *IncidentMemberStructEscalation {
	s.Description = &v
	return s
}

func (s *IncidentMemberStructEscalation) SetIncidentEscalationId(v string) *IncidentMemberStructEscalation {
	s.IncidentEscalationId = &v
	return s
}

func (s *IncidentMemberStructEscalation) SetName(v string) *IncidentMemberStructEscalation {
	s.Name = &v
	return s
}

func (s *IncidentMemberStructEscalation) SetStageIndex(v string) *IncidentMemberStructEscalation {
	s.StageIndex = &v
	return s
}

func (s *IncidentMemberStructEscalation) SetTitle(v string) *IncidentMemberStructEscalation {
	s.Title = &v
	return s
}

func (s *IncidentMemberStructEscalation) Validate() error {
	return dara.Validate(s)
}

type IncidentMemberStructScheduleGroup struct {
	// The contact ID.
	//
	// example:
	//
	// group-001
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	// The name.
	//
	// example:
	//
	// Night on-call group
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s IncidentMemberStructScheduleGroup) String() string {
	return dara.Prettify(s)
}

func (s IncidentMemberStructScheduleGroup) GoString() string {
	return s.String()
}

func (s *IncidentMemberStructScheduleGroup) GetContactId() *string {
	return s.ContactId
}

func (s *IncidentMemberStructScheduleGroup) GetName() *string {
	return s.Name
}

func (s *IncidentMemberStructScheduleGroup) SetContactId(v string) *IncidentMemberStructScheduleGroup {
	s.ContactId = &v
	return s
}

func (s *IncidentMemberStructScheduleGroup) SetName(v string) *IncidentMemberStructScheduleGroup {
	s.Name = &v
	return s
}

func (s *IncidentMemberStructScheduleGroup) Validate() error {
	return dara.Validate(s)
}
