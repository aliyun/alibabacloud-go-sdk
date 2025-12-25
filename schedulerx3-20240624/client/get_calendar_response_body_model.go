// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCalendarResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetCalendarResponseBody
	GetCode() *int32
	SetData(v *GetCalendarResponseBodyData) *GetCalendarResponseBody
	GetData() *GetCalendarResponseBodyData
	SetMessage(v string) *GetCalendarResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetCalendarResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCalendarResponseBody
	GetSuccess() *bool
}

type GetCalendarResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *GetCalendarResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2C3E52FF-CBE9-5C0E-8252-37ACFF1F5EFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetCalendarResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCalendarResponseBody) GoString() string {
	return s.String()
}

func (s *GetCalendarResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetCalendarResponseBody) GetData() *GetCalendarResponseBodyData {
	return s.Data
}

func (s *GetCalendarResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetCalendarResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCalendarResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCalendarResponseBody) SetCode(v int32) *GetCalendarResponseBody {
	s.Code = &v
	return s
}

func (s *GetCalendarResponseBody) SetData(v *GetCalendarResponseBodyData) *GetCalendarResponseBody {
	s.Data = v
	return s
}

func (s *GetCalendarResponseBody) SetMessage(v string) *GetCalendarResponseBody {
	s.Message = &v
	return s
}

func (s *GetCalendarResponseBody) SetRequestId(v string) *GetCalendarResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCalendarResponseBody) SetSuccess(v bool) *GetCalendarResponseBody {
	s.Success = &v
	return s
}

func (s *GetCalendarResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCalendarResponseBodyData struct {
	// example:
	//
	// workday
	CalendarName *string `json:"CalendarName,omitempty" xml:"CalendarName,omitempty"`
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
	// example:
	//
	// 2030
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s GetCalendarResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCalendarResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCalendarResponseBodyData) GetCalendarName() *string {
	return s.CalendarName
}

func (s *GetCalendarResponseBodyData) GetMonths() *string {
	return s.Months
}

func (s *GetCalendarResponseBodyData) GetYear() *int32 {
	return s.Year
}

func (s *GetCalendarResponseBodyData) SetCalendarName(v string) *GetCalendarResponseBodyData {
	s.CalendarName = &v
	return s
}

func (s *GetCalendarResponseBodyData) SetMonths(v string) *GetCalendarResponseBodyData {
	s.Months = &v
	return s
}

func (s *GetCalendarResponseBodyData) SetYear(v int32) *GetCalendarResponseBodyData {
	s.Year = &v
	return s
}

func (s *GetCalendarResponseBodyData) Validate() error {
	return dara.Validate(s)
}
