// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperationLogPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeOperationLogPageListRequest
	GetLang() *string
	SetCurrentPage(v int32) *DescribeOperationLogPageListRequest
	GetCurrentPage() *int32
	SetEndDate(v int64) *DescribeOperationLogPageListRequest
	GetEndDate() *int64
	SetPageSize(v int32) *DescribeOperationLogPageListRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeOperationLogPageListRequest
	GetRegId() *string
	SetStartDate(v int64) *DescribeOperationLogPageListRequest
	GetStartDate() *int64
}

type DescribeOperationLogPageListRequest struct {
	// Set the language type for request and response messages, default value is **zh**. Values:
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
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// End time.
	//
	// example:
	//
	// 1733364850919
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Page size, default value is 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Start time.
	//
	// example:
	//
	// 1733364850919
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s DescribeOperationLogPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperationLogPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperationLogPageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOperationLogPageListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeOperationLogPageListRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *DescribeOperationLogPageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeOperationLogPageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeOperationLogPageListRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *DescribeOperationLogPageListRequest) SetLang(v string) *DescribeOperationLogPageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOperationLogPageListRequest) SetCurrentPage(v int32) *DescribeOperationLogPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeOperationLogPageListRequest) SetEndDate(v int64) *DescribeOperationLogPageListRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeOperationLogPageListRequest) SetPageSize(v int32) *DescribeOperationLogPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeOperationLogPageListRequest) SetRegId(v string) *DescribeOperationLogPageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeOperationLogPageListRequest) SetStartDate(v int64) *DescribeOperationLogPageListRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeOperationLogPageListRequest) Validate() error {
	return dara.Validate(s)
}
