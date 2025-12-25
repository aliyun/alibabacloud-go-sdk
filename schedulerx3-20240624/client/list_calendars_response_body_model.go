// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCalendarsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListCalendarsResponseBody
	GetCode() *int32
	SetData(v *ListCalendarsResponseBodyData) *ListCalendarsResponseBody
	GetData() *ListCalendarsResponseBodyData
	SetMessage(v string) *ListCalendarsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListCalendarsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCalendarsResponseBody
	GetSuccess() *bool
}

type ListCalendarsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// -
	Data *ListCalendarsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3835AA29-2298-5434-BC53-9CC377CDFD2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCalendarsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListCalendarsResponseBody) GetData() *ListCalendarsResponseBodyData {
	return s.Data
}

func (s *ListCalendarsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListCalendarsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCalendarsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCalendarsResponseBody) SetCode(v int32) *ListCalendarsResponseBody {
	s.Code = &v
	return s
}

func (s *ListCalendarsResponseBody) SetData(v *ListCalendarsResponseBodyData) *ListCalendarsResponseBody {
	s.Data = v
	return s
}

func (s *ListCalendarsResponseBody) SetMessage(v string) *ListCalendarsResponseBody {
	s.Message = &v
	return s
}

func (s *ListCalendarsResponseBody) SetRequestId(v string) *ListCalendarsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCalendarsResponseBody) SetSuccess(v bool) *ListCalendarsResponseBody {
	s.Success = &v
	return s
}

func (s *ListCalendarsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCalendarsResponseBodyData struct {
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// eCKqVlS5FKF5EWGGOo8EgQ==
	NextToken *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	Records   []*ListCalendarsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListCalendarsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCalendarsResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCalendarsResponseBodyData) GetRecords() []*ListCalendarsResponseBodyDataRecords {
	return s.Records
}

func (s *ListCalendarsResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *ListCalendarsResponseBodyData) SetMaxResults(v int32) *ListCalendarsResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListCalendarsResponseBodyData) SetNextToken(v string) *ListCalendarsResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListCalendarsResponseBodyData) SetRecords(v []*ListCalendarsResponseBodyDataRecords) *ListCalendarsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListCalendarsResponseBodyData) SetTotal(v int64) *ListCalendarsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListCalendarsResponseBodyData) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCalendarsResponseBodyDataRecords struct {
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
	// 2025
	Year *int32 `json:"Year,omitempty" xml:"Year,omitempty"`
}

func (s ListCalendarsResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListCalendarsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListCalendarsResponseBodyDataRecords) GetCalendarName() *string {
	return s.CalendarName
}

func (s *ListCalendarsResponseBodyDataRecords) GetMonths() *string {
	return s.Months
}

func (s *ListCalendarsResponseBodyDataRecords) GetYear() *int32 {
	return s.Year
}

func (s *ListCalendarsResponseBodyDataRecords) SetCalendarName(v string) *ListCalendarsResponseBodyDataRecords {
	s.CalendarName = &v
	return s
}

func (s *ListCalendarsResponseBodyDataRecords) SetMonths(v string) *ListCalendarsResponseBodyDataRecords {
	s.Months = &v
	return s
}

func (s *ListCalendarsResponseBodyDataRecords) SetYear(v int32) *ListCalendarsResponseBodyDataRecords {
	s.Year = &v
	return s
}

func (s *ListCalendarsResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
