// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarName(v string) *GetCalendarRequest
	GetCalendarName() *string
	SetClusterId(v string) *GetCalendarRequest
	GetClusterId() *string
	SetYear(v int32) *GetCalendarRequest
	GetYear() *int32
}

type GetCalendarRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// workday
	CalendarName *string `json:"CalendarName,omitempty" xml:"CalendarName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s GetCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCalendarRequest) GoString() string {
	return s.String()
}

func (s *GetCalendarRequest) GetCalendarName() *string {
	return s.CalendarName
}

func (s *GetCalendarRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetCalendarRequest) GetYear() *int32 {
	return s.Year
}

func (s *GetCalendarRequest) SetCalendarName(v string) *GetCalendarRequest {
	s.CalendarName = &v
	return s
}

func (s *GetCalendarRequest) SetClusterId(v string) *GetCalendarRequest {
	s.ClusterId = &v
	return s
}

func (s *GetCalendarRequest) SetYear(v int32) *GetCalendarRequest {
	s.Year = &v
	return s
}

func (s *GetCalendarRequest) Validate() error {
	return dara.Validate(s)
}
