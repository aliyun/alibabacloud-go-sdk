// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscribedCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *DeleteSubscribedCalendarRequest
	GetCalendarId() *string
}

type DeleteSubscribedCalendarRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// MzM5Mxxx
	CalendarId *string `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
}

func (s DeleteSubscribedCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscribedCalendarRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubscribedCalendarRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *DeleteSubscribedCalendarRequest) SetCalendarId(v string) *DeleteSubscribedCalendarRequest {
	s.CalendarId = &v
	return s
}

func (s *DeleteSubscribedCalendarRequest) Validate() error {
	return dara.Validate(s)
}
