// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFindGuestCredentialsRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivityId(v string) *FindGuestCredentialsRecordRequest
	GetActivityId() *string
	SetDateTimeString(v string) *FindGuestCredentialsRecordRequest
	GetDateTimeString() *string
	SetEndDateTime(v string) *FindGuestCredentialsRecordRequest
	GetEndDateTime() *string
	SetStartDateTime(v string) *FindGuestCredentialsRecordRequest
	GetStartDateTime() *string
}

type FindGuestCredentialsRecordRequest struct {
	// example:
	//
	// 34429
	ActivityId *string `json:"ActivityId,omitempty" xml:"ActivityId,omitempty"`
	// example:
	//
	// 2023-08-07 12:00:00
	DateTimeString *string `json:"DateTimeString,omitempty" xml:"DateTimeString,omitempty"`
	EndDateTime    *string `json:"EndDateTime,omitempty" xml:"EndDateTime,omitempty"`
	StartDateTime  *string `json:"StartDateTime,omitempty" xml:"StartDateTime,omitempty"`
}

func (s FindGuestCredentialsRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s FindGuestCredentialsRecordRequest) GoString() string {
	return s.String()
}

func (s *FindGuestCredentialsRecordRequest) GetActivityId() *string {
	return s.ActivityId
}

func (s *FindGuestCredentialsRecordRequest) GetDateTimeString() *string {
	return s.DateTimeString
}

func (s *FindGuestCredentialsRecordRequest) GetEndDateTime() *string {
	return s.EndDateTime
}

func (s *FindGuestCredentialsRecordRequest) GetStartDateTime() *string {
	return s.StartDateTime
}

func (s *FindGuestCredentialsRecordRequest) SetActivityId(v string) *FindGuestCredentialsRecordRequest {
	s.ActivityId = &v
	return s
}

func (s *FindGuestCredentialsRecordRequest) SetDateTimeString(v string) *FindGuestCredentialsRecordRequest {
	s.DateTimeString = &v
	return s
}

func (s *FindGuestCredentialsRecordRequest) SetEndDateTime(v string) *FindGuestCredentialsRecordRequest {
	s.EndDateTime = &v
	return s
}

func (s *FindGuestCredentialsRecordRequest) SetStartDateTime(v string) *FindGuestCredentialsRecordRequest {
	s.StartDateTime = &v
	return s
}

func (s *FindGuestCredentialsRecordRequest) Validate() error {
	return dara.Validate(s)
}
