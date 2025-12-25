// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarName(v string) *CreateCalendarRequest
	GetCalendarName() *string
	SetClientToken(v string) *CreateCalendarRequest
	GetClientToken() *string
	SetClusterId(v string) *CreateCalendarRequest
	GetClusterId() *string
	SetMonths(v string) *CreateCalendarRequest
	GetMonths() *string
	SetYear(v int32) *CreateCalendarRequest
	GetYear() *int32
}

type CreateCalendarRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// workday
	CalendarName *string `json:"CalendarName,omitempty" xml:"CalendarName,omitempty"`
	// example:
	//
	// 123456789
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
	// [
	//
	//   {"month":1,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30,31]},
	//
	//   {"month":2,"days":[3,4,5,6,7,10,11,12,13,14,17,18,19,20,21,24,25,26,27,28]},
	//
	//   {"month":3,"days":[3,4,5,6,7,10,11,12,13,14,17,18,19,20,21,24,25,26,27,28,31]},
	//
	//   {"month":4,"days":[1,2,3,4,7,8,9,10,11,14,15,16,17,18,21,22,23,24,25,28,29,30]},
	//
	//   {"month":5,"days":[1,2,5,6,7,8,9,12,13,14,15,16,19,20,21,22,23,26,27,28,29,30]},
	//
	//   {"month":6,"days":[2,3,4,5,6,9,10,11,12,13,16,17,18,19,20,23,24,25,26,27,30]},
	//
	//   {"month":7,"days":[1,2,3,4,7,8,9,10,11,14,15,16,17,18,21,22,23,24,25,28,29,30,31]},
	//
	//   {"month":8,"days":[1,4,5,6,7,8,11,12,13,14,15,18,19,20,21,22,25,26,27,28,29]},
	//
	//   {"month":9,"days":[1,2,3,4,5,8,9,10,11,12,15,16,17,18,19,22,23,24,25,26,29,30]},
	//
	//   {"month":10,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30,31]},
	//
	//   {"month":11,"days":[3,4,5,6,7,10,11,12,13,14,17,18,19,20,21,24,25,26,27,28]},
	//
	//   {"month":12,"days":[1,2,3,4,5,8,9,10,11,12,15,16,17,18,19,22,23,24,25,26,29,30,31]}
	//
	// ]
	Months *string `json:"Months,omitempty" xml:"Months,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2026
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s CreateCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCalendarRequest) GoString() string {
	return s.String()
}

func (s *CreateCalendarRequest) GetCalendarName() *string {
	return s.CalendarName
}

func (s *CreateCalendarRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCalendarRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateCalendarRequest) GetMonths() *string {
	return s.Months
}

func (s *CreateCalendarRequest) GetYear() *int32 {
	return s.Year
}

func (s *CreateCalendarRequest) SetCalendarName(v string) *CreateCalendarRequest {
	s.CalendarName = &v
	return s
}

func (s *CreateCalendarRequest) SetClientToken(v string) *CreateCalendarRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCalendarRequest) SetClusterId(v string) *CreateCalendarRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateCalendarRequest) SetMonths(v string) *CreateCalendarRequest {
	s.Months = &v
	return s
}

func (s *CreateCalendarRequest) SetYear(v int32) *CreateCalendarRequest {
	s.Year = &v
	return s
}

func (s *CreateCalendarRequest) Validate() error {
	return dara.Validate(s)
}
