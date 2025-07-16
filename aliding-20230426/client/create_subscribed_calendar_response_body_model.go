// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubscribedCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *CreateSubscribedCalendarResponseBody
	GetCalendarId() *string
	SetRequestId(v string) *CreateSubscribedCalendarResponseBody
	GetRequestId() *string
}

type CreateSubscribedCalendarResponseBody struct {
	// example:
	//
	// M5MjkxNDUxQHVzZXJzLmRpbmd0YWxrLmxxxxxxx
	CalendarId *string `json:"calendarId,omitempty" xml:"calendarId,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateSubscribedCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSubscribedCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubscribedCalendarResponseBody) GetCalendarId() *string {
	return s.CalendarId
}

func (s *CreateSubscribedCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSubscribedCalendarResponseBody) SetCalendarId(v string) *CreateSubscribedCalendarResponseBody {
	s.CalendarId = &v
	return s
}

func (s *CreateSubscribedCalendarResponseBody) SetRequestId(v string) *CreateSubscribedCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubscribedCalendarResponseBody) Validate() error {
	return dara.Validate(s)
}
