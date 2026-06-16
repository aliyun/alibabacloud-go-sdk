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
	SetOperationSummary(v string) *DescribeOperationLogPageListRequest
	GetOperationSummary() *string
	SetPageSize(v int32) *DescribeOperationLogPageListRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeOperationLogPageListRequest
	GetRegId() *string
	SetStartDate(v int64) *DescribeOperationLogPageListRequest
	GetStartDate() *int64
	SetUserNameSearch(v string) *DescribeOperationLogPageListRequest
	GetUserNameSearch() *string
}

type DescribeOperationLogPageListRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// The end time.
	//
	// example:
	//
	// 1733364850919
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// The change content to query.
	//
	// example:
	//
	// 修改变量
	OperationSummary *string `json:"operationSummary,omitempty" xml:"operationSummary,omitempty"`
	// The number of entries per page. Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The start time.
	//
	// example:
	//
	// 1733364850919
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// The operator name to query.
	//
	// example:
	//
	// root
	UserNameSearch *string `json:"userNameSearch,omitempty" xml:"userNameSearch,omitempty"`
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

func (s *DescribeOperationLogPageListRequest) GetOperationSummary() *string {
	return s.OperationSummary
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

func (s *DescribeOperationLogPageListRequest) GetUserNameSearch() *string {
	return s.UserNameSearch
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

func (s *DescribeOperationLogPageListRequest) SetOperationSummary(v string) *DescribeOperationLogPageListRequest {
	s.OperationSummary = &v
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

func (s *DescribeOperationLogPageListRequest) SetUserNameSearch(v string) *DescribeOperationLogPageListRequest {
	s.UserNameSearch = &v
	return s
}

func (s *DescribeOperationLogPageListRequest) Validate() error {
	return dara.Validate(s)
}
