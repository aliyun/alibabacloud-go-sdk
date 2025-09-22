// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindGuestTicketRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v string) *FindGuestTicketRecordRequest
	GetActivityId() *string
	SetDateTimeString(v string) *FindGuestTicketRecordRequest
	GetDateTimeString() *string
	SetEndDateTime(v string) *FindGuestTicketRecordRequest
	GetEndDateTime() *string
	SetStartDateTime(v string) *FindGuestTicketRecordRequest
	GetStartDateTime() *string
}

type FindGuestTicketRecordRequest struct {
	// example:
	//
	// 34434
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 2023-09-04 15:14:00
	DateTimeString *string `json:"DateTimeString,omitempty" xml:"DateTimeString,omitempty"`
	EndDateTime    *string `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	StartDateTime  *string `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
}

func (s FindGuestTicketRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s FindGuestTicketRecordRequest) GoString() string {
	return s.String()
}

func (s *FindGuestTicketRecordRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *FindGuestTicketRecordRequest) GetDateTimeString() *string {
	return s.DateTimeString
}

func (s *FindGuestTicketRecordRequest) GetEndDateTime() *string {
	return s.EndDateTime
}

func (s *FindGuestTicketRecordRequest) GetStartDateTime() *string {
	return s.StartDateTime
}

func (s *FindGuestTicketRecordRequest) SetActivityId(v string) *FindGuestTicketRecordRequest {
	s.ActivityId = &v
	return s
}

func (s *FindGuestTicketRecordRequest) SetDateTimeString(v string) *FindGuestTicketRecordRequest {
	s.DateTimeString = &v
	return s
}

func (s *FindGuestTicketRecordRequest) SetEndDateTime(v string) *FindGuestTicketRecordRequest {
	s.EndDateTime = &v
	return s
}

func (s *FindGuestTicketRecordRequest) SetStartDateTime(v string) *FindGuestTicketRecordRequest {
	s.StartDateTime = &v
	return s
}

func (s *FindGuestTicketRecordRequest) Validate() error {
	return dara.Validate(s)
}
