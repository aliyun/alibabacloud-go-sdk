// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCalendarName(v string) *UpdateCalendarRequest
	GetCalendarName() *string
	SetClientToken(v string) *UpdateCalendarRequest
	GetClientToken() *string
	SetClusterId(v string) *UpdateCalendarRequest
	GetClusterId() *string
	SetIncremental(v bool) *UpdateCalendarRequest
	GetIncremental() *bool
	SetMonths(v string) *UpdateCalendarRequest
	GetMonths() *string
	SetYear(v int32) *UpdateCalendarRequest
	GetYear() *int32
}

type UpdateCalendarRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// workday
	CalendarName *string `json:"CalendarName,omitempty" xml:"CalendarName,omitempty"`
	// example:
	//
	// D0DE9C33-992A-580B-89C4-B609A292748D
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// false
	Incremental *bool `json:"Incremental,omitempty" xml:"Incremental,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"month":1,"days":[3,4,5,6,9,10,11,12,13,16,17,18,19,20,28,29,30,31]},{"month":2,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28]},{"month":3,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30,31]},{"month":4,"days":[3,4,6,7,10,11,12,13,14,17,18,19,20,21,23,24,25,26,27,28]},{"month":5,"days":[4,5,6,8,9,10,11,12,15,16,17,18,19,22,23,24,25,26,29,30,31]},{"month":6,"days":[1,2,5,6,7,8,9,12,13,14,15,16,19,20,21,25,26,27,28,29,30]},{"month":7,"days":[3,4,5,6,7,10,11,12,13,14,17,18,19,20,21,24,25,26,27,28,31]},{"month":8,"days":[1,2,3,4,7,8,9,10,11,14,15,16,17,18,21,22,23,24,25,28,29,30,31]},{"month":9,"days":[1,4,5,6,7,8,11,12,13,14,15,18,19,20,21,22,25,26,27,28]},{"month":10,"days":[7,8,9,10,11,12,13,16,17,18,19,20,23,24,25,26,27,30,31]},{"month":11,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30]},{"month":12,"days":[1,4,5,6,7,8,11,12,13,14,15,18,19,20,21,22,25,26,27,28,29]}]
	Months *string `json:"Months,omitempty" xml:"Months,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s UpdateCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCalendarRequest) GoString() string {
	return s.String()
}

func (s *UpdateCalendarRequest) GetCalendarName() *string {
	return s.CalendarName
}

func (s *UpdateCalendarRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateCalendarRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateCalendarRequest) GetIncremental() *bool {
	return s.Incremental
}

func (s *UpdateCalendarRequest) GetMonths() *string {
	return s.Months
}

func (s *UpdateCalendarRequest) GetYear() *int32 {
	return s.Year
}

func (s *UpdateCalendarRequest) SetCalendarName(v string) *UpdateCalendarRequest {
	s.CalendarName = &v
	return s
}

func (s *UpdateCalendarRequest) SetClientToken(v string) *UpdateCalendarRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateCalendarRequest) SetClusterId(v string) *UpdateCalendarRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateCalendarRequest) SetIncremental(v bool) *UpdateCalendarRequest {
	s.Incremental = &v
	return s
}

func (s *UpdateCalendarRequest) SetMonths(v string) *UpdateCalendarRequest {
	s.Months = &v
	return s
}

func (s *UpdateCalendarRequest) SetYear(v int32) *UpdateCalendarRequest {
	s.Year = &v
	return s
}

func (s *UpdateCalendarRequest) Validate() error {
	return dara.Validate(s)
}
