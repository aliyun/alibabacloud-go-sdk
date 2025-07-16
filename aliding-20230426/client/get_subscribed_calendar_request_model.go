// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscribedCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarId(v string) *GetSubscribedCalendarRequest
	GetCalendarId() *string
}

type GetSubscribedCalendarRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// M5MjkxNDUxQHVzZXJzLmRpbmd0YWxrLmNv
	CalendarId *string `json:"CalendarId,omitempty" xml:"CalendarId,omitempty"`
}

func (s GetSubscribedCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubscribedCalendarRequest) GoString() string {
	return s.String()
}

func (s *GetSubscribedCalendarRequest) GetCalendarId() *string {
	return s.CalendarId
}

func (s *GetSubscribedCalendarRequest) SetCalendarId(v string) *GetSubscribedCalendarRequest {
	s.CalendarId = &v
	return s
}

func (s *GetSubscribedCalendarRequest) Validate() error {
	return dara.Validate(s)
}
