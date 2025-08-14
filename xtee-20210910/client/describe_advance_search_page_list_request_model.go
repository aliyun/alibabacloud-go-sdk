// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvanceSearchPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeAdvanceSearchPageListRequest
	GetLang() *string
	SetCondition(v string) *DescribeAdvanceSearchPageListRequest
	GetCondition() *string
	SetCurrentPage(v int64) *DescribeAdvanceSearchPageListRequest
	GetCurrentPage() *int64
	SetEventBeginTime(v int64) *DescribeAdvanceSearchPageListRequest
	GetEventBeginTime() *int64
	SetEventCodes(v string) *DescribeAdvanceSearchPageListRequest
	GetEventCodes() *string
	SetEventEndTime(v int64) *DescribeAdvanceSearchPageListRequest
	GetEventEndTime() *int64
	SetFieldName(v string) *DescribeAdvanceSearchPageListRequest
	GetFieldName() *string
	SetFieldValue(v string) *DescribeAdvanceSearchPageListRequest
	GetFieldValue() *string
	SetPageSize(v int64) *DescribeAdvanceSearchPageListRequest
	GetPageSize() *int64
	SetRegId(v string) *DescribeAdvanceSearchPageListRequest
	GetRegId() *string
}

type DescribeAdvanceSearchPageListRequest struct {
	// Sets the language type for requests and received messages, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Condition value.
	//
	// example:
	//
	// {
	//
	//     "relationship": "and",
	//
	//     "list": [
	//
	//         {
	//
	//             "deepCount": 1,
	//
	//             "left": {
	//
	//                 "hasRightVariable": true,
	//
	//                 "fieldType": "STRING",
	//
	//                 "functionName": "",
	//
	//                 "leftVariableType": "NATIVE",
	//
	//                 "name": "accountId",
	//
	//                 "operatorCode": "equals"
	//
	//             },
	//
	//             "right": {
	//
	//                 "rightVariableType": "constant",
	//
	//                 "name": "10000",
	//
	//                 "functionName": ""
	//
	//             },
	//
	//             "operatorCode": "equals"
	//
	//         }
	//
	//     ]
	//
	// }
	Condition *string `json:"condition,omitempty" xml:"condition,omitempty"`
	// Current page number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Query start time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1753372800000
	EventBeginTime *int64 `json:"eventBeginTime,omitempty" xml:"eventBeginTime,omitempty"`
	// Event code.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["de_ahqhsw7665"]
	EventCodes *string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty"`
	// End time, accurate to milliseconds (ms).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1753459199059
	EventEndTime *int64 `json:"eventEndTime,omitempty" xml:"eventEndTime,omitempty"`
	// Field name
	//
	// example:
	//
	// age
	FieldName *string `json:"fieldName,omitempty" xml:"fieldName,omitempty"`
	// Field value
	//
	// example:
	//
	// 20
	FieldValue *string `json:"fieldValue,omitempty" xml:"fieldValue,omitempty"`
	// Page size, default value is 10
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
}

func (s DescribeAdvanceSearchPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvanceSearchPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeAdvanceSearchPageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeAdvanceSearchPageListRequest) GetCondition() *string {
	return s.Condition
}

func (s *DescribeAdvanceSearchPageListRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeAdvanceSearchPageListRequest) GetEventBeginTime() *int64 {
	return s.EventBeginTime
}

func (s *DescribeAdvanceSearchPageListRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeAdvanceSearchPageListRequest) GetEventEndTime() *int64 {
	return s.EventEndTime
}

func (s *DescribeAdvanceSearchPageListRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeAdvanceSearchPageListRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeAdvanceSearchPageListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeAdvanceSearchPageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeAdvanceSearchPageListRequest) SetLang(v string) *DescribeAdvanceSearchPageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeAdvanceSearchPageListRequest) SetCondition(v string) *DescribeAdvanceSearchPageListRequest {
	s.Condition = &v
	return s
}

func (s *DescribeAdvanceSearchPageListRequest) SetCurrentPage(v int64) *DescribeAdvanceSearchPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeAdvanceSearchPageListRequest) SetEventBeginTime(v int64) *DescribeAdvanceSearchPageListRequest {
	s.EventBeginTime = &v
	return s
}

func (s *DescribeAdvanceSearchPageListRequest) SetEventCodes(v string) *DescribeAdvanceSearchPageListRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeAdvanceSearchPageListRequest) SetEventEndTime(v int64) *DescribeAdvanceSearchPageListRequest {
	s.EventEndTime = &v
	return s
}

func (s *DescribeAdvanceSearchPageListRequest) SetFieldName(v string) *DescribeAdvanceSearchPageListRequest {
	s.FieldName = &v
	return s
}

func (s *DescribeAdvanceSearchPageListRequest) SetFieldValue(v string) *DescribeAdvanceSearchPageListRequest {
	s.FieldValue = &v
	return s
}

func (s *DescribeAdvanceSearchPageListRequest) SetPageSize(v int64) *DescribeAdvanceSearchPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAdvanceSearchPageListRequest) SetRegId(v string) *DescribeAdvanceSearchPageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeAdvanceSearchPageListRequest) Validate() error {
	return dara.Validate(s)
}
