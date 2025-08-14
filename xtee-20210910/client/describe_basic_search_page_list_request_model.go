// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBasicSearchPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeBasicSearchPageListRequest
	GetLang() *string
	SetCurrentPage(v int64) *DescribeBasicSearchPageListRequest
	GetCurrentPage() *int64
	SetEventBeginTime(v int64) *DescribeBasicSearchPageListRequest
	GetEventBeginTime() *int64
	SetEventCodes(v string) *DescribeBasicSearchPageListRequest
	GetEventCodes() *string
	SetEventEndTime(v int64) *DescribeBasicSearchPageListRequest
	GetEventEndTime() *int64
	SetFieldName(v string) *DescribeBasicSearchPageListRequest
	GetFieldName() *string
	SetFieldValue(v string) *DescribeBasicSearchPageListRequest
	GetFieldValue() *string
	SetPageSize(v int64) *DescribeBasicSearchPageListRequest
	GetPageSize() *int64
	SetRegId(v string) *DescribeBasicSearchPageListRequest
	GetRegId() *string
}

type DescribeBasicSearchPageListRequest struct {
	// Sets the language type for requests and received messages, with a default value of **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
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
	// Page size, with a default value of 10
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

func (s DescribeBasicSearchPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBasicSearchPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeBasicSearchPageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeBasicSearchPageListRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeBasicSearchPageListRequest) GetEventBeginTime() *int64 {
	return s.EventBeginTime
}

func (s *DescribeBasicSearchPageListRequest) GetEventCodes() *string {
	return s.EventCodes
}

func (s *DescribeBasicSearchPageListRequest) GetEventEndTime() *int64 {
	return s.EventEndTime
}

func (s *DescribeBasicSearchPageListRequest) GetFieldName() *string {
	return s.FieldName
}

func (s *DescribeBasicSearchPageListRequest) GetFieldValue() *string {
	return s.FieldValue
}

func (s *DescribeBasicSearchPageListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeBasicSearchPageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeBasicSearchPageListRequest) SetLang(v string) *DescribeBasicSearchPageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeBasicSearchPageListRequest) SetCurrentPage(v int64) *DescribeBasicSearchPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeBasicSearchPageListRequest) SetEventBeginTime(v int64) *DescribeBasicSearchPageListRequest {
	s.EventBeginTime = &v
	return s
}

func (s *DescribeBasicSearchPageListRequest) SetEventCodes(v string) *DescribeBasicSearchPageListRequest {
	s.EventCodes = &v
	return s
}

func (s *DescribeBasicSearchPageListRequest) SetEventEndTime(v int64) *DescribeBasicSearchPageListRequest {
	s.EventEndTime = &v
	return s
}

func (s *DescribeBasicSearchPageListRequest) SetFieldName(v string) *DescribeBasicSearchPageListRequest {
	s.FieldName = &v
	return s
}

func (s *DescribeBasicSearchPageListRequest) SetFieldValue(v string) *DescribeBasicSearchPageListRequest {
	s.FieldValue = &v
	return s
}

func (s *DescribeBasicSearchPageListRequest) SetPageSize(v int64) *DescribeBasicSearchPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeBasicSearchPageListRequest) SetRegId(v string) *DescribeBasicSearchPageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeBasicSearchPageListRequest) Validate() error {
	return dara.Validate(s)
}
