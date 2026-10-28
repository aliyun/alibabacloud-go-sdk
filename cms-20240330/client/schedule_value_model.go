// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduleValue interface {
	dara.Model
	String() string
	GoString() string
	SetContactId(v string) *ScheduleValue
	GetContactId() *string
	SetContactName(v string) *ScheduleValue
	GetContactName() *string
	SetEnd(v string) *ScheduleValue
	GetEnd() *string
	SetStart(v string) *ScheduleValue
	GetStart() *string
}

type ScheduleValue struct {
	// example:
	//
	// testId
	ContactId *string `json:"contactId,omitempty" xml:"contactId,omitempty"`
	// example:
	//
	// testName
	ContactName *string `json:"contactName,omitempty" xml:"contactName,omitempty"`
	// example:
	//
	// 2024-10-01 02:00:00
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 2024-10-01 22:00:00
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ScheduleValue) String() string {
	return dara.Prettify(s)
}

func (s ScheduleValue) GoString() string {
	return s.String()
}

func (s *ScheduleValue) GetContactId() *string {
	return s.ContactId
}

func (s *ScheduleValue) GetContactName() *string {
	return s.ContactName
}

func (s *ScheduleValue) GetEnd() *string {
	return s.End
}

func (s *ScheduleValue) GetStart() *string {
	return s.Start
}

func (s *ScheduleValue) SetContactId(v string) *ScheduleValue {
	s.ContactId = &v
	return s
}

func (s *ScheduleValue) SetContactName(v string) *ScheduleValue {
	s.ContactName = &v
	return s
}

func (s *ScheduleValue) SetEnd(v string) *ScheduleValue {
	s.End = &v
	return s
}

func (s *ScheduleValue) SetStart(v string) *ScheduleValue {
	s.Start = &v
	return s
}

func (s *ScheduleValue) Validate() error {
	return dara.Validate(s)
}
