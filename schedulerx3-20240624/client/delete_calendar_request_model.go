// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarName(v string) *DeleteCalendarRequest
	GetCalendarName() *string
	SetClusterId(v string) *DeleteCalendarRequest
	GetClusterId() *string
	SetYear(v int32) *DeleteCalendarRequest
	GetYear() *int32
}

type DeleteCalendarRequest struct {
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

func (s DeleteCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCalendarRequest) GoString() string {
	return s.String()
}

func (s *DeleteCalendarRequest) GetCalendarName() *string {
	return s.CalendarName
}

func (s *DeleteCalendarRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteCalendarRequest) GetYear() *int32 {
	return s.Year
}

func (s *DeleteCalendarRequest) SetCalendarName(v string) *DeleteCalendarRequest {
	s.CalendarName = &v
	return s
}

func (s *DeleteCalendarRequest) SetClusterId(v string) *DeleteCalendarRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteCalendarRequest) SetYear(v int32) *DeleteCalendarRequest {
	s.Year = &v
	return s
}

func (s *DeleteCalendarRequest) Validate() error {
	return dara.Validate(s)
}
