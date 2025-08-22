// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCalendarRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ImportCalendarRequest
	GetClusterId() *string
	SetMonths(v string) *ImportCalendarRequest
	GetMonths() *string
	SetName(v string) *ImportCalendarRequest
	GetName() *string
	SetYear(v int32) *ImportCalendarRequest
	GetYear() *int32
}

type ImportCalendarRequest struct {
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
	// [{"month":1,"days":[3,4,5,6,9,10,11,12,13,16,17,18,19,20,28,29,30,31]},{"month":2,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28]},{"month":3,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30,31]},{"month":4,"days":[3,4,6,7,10,11,12,13,14,17,18,19,20,21,23,24,25,26,27,28]},{"month":5,"days":[4,5,6,8,9,10,11,12,15,16,17,18,19,22,23,24,25,26,29,30,31]},{"month":6,"days":[1,2,5,6,7,8,9,12,13,14,15,16,19,20,21,25,26,27,28,29,30]},{"month":7,"days":[3,4,5,6,7,10,11,12,13,14,17,18,19,20,21,24,25,26,27,28,31]},{"month":8,"days":[1,2,3,4,7,8,9,10,11,14,15,16,17,18,21,22,23,24,25,28,29,30,31]},{"month":9,"days":[1,4,5,6,7,8,11,12,13,14,15,18,19,20,21,22,25,26,27,28]},{"month":10,"days":[7,8,9,10,11,12,13,16,17,18,19,20,23,24,25,26,27,30,31]},{"month":11,"days":[1,2,3,6,7,8,9,10,13,14,15,16,17,20,21,22,23,24,27,28,29,30]},{"month":12,"days":[1,4,5,6,7,8,11,12,13,14,15,18,19,20,21,22,25,26,27,28,29]}]
	Months *string `json:"Months,omitempty" xml:"Months,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// workday
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 2024
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ImportCalendarRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportCalendarRequest) GoString() string {
	return s.String()
}

func (s *ImportCalendarRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ImportCalendarRequest) GetMonths() *string {
	return s.Months
}

func (s *ImportCalendarRequest) GetName() *string {
	return s.Name
}

func (s *ImportCalendarRequest) GetYear() *int32 {
	return s.Year
}

func (s *ImportCalendarRequest) SetClusterId(v string) *ImportCalendarRequest {
	s.ClusterId = &v
	return s
}

func (s *ImportCalendarRequest) SetMonths(v string) *ImportCalendarRequest {
	s.Months = &v
	return s
}

func (s *ImportCalendarRequest) SetName(v string) *ImportCalendarRequest {
	s.Name = &v
	return s
}

func (s *ImportCalendarRequest) SetYear(v int32) *ImportCalendarRequest {
	s.Year = &v
	return s
}

func (s *ImportCalendarRequest) Validate() error {
	return dara.Validate(s)
}
