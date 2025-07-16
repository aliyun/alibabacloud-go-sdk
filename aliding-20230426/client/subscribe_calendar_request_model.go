// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscribeCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *SubscribeCalendarRequest
	GetCalendarId() *string
}

type SubscribeCalendarRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MzM5Mxxx
	CalendarId *string `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
}

func (s SubscribeCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s SubscribeCalendarRequest) GoString() string {
	return s.String()
}

func (s *SubscribeCalendarRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *SubscribeCalendarRequest) SetCalendarId(v string) *SubscribeCalendarRequest {
	s.CalendarId = &v
	return s
}

func (s *SubscribeCalendarRequest) Validate() error {
	return dara.Validate(s)
}
