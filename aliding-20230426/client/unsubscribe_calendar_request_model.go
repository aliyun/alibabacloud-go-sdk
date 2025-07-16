// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnsubscribeCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *UnsubscribeCalendarRequest
	GetCalendarId() *string
}

type UnsubscribeCalendarRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MzM5Mxxx
	CalendarId *string `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
}

func (s UnsubscribeCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s UnsubscribeCalendarRequest) GoString() string {
	return s.String()
}

func (s *UnsubscribeCalendarRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *UnsubscribeCalendarRequest) SetCalendarId(v string) *UnsubscribeCalendarRequest {
	s.CalendarId = &v
	return s
}

func (s *UnsubscribeCalendarRequest) Validate() error {
	return dara.Validate(s)
}
