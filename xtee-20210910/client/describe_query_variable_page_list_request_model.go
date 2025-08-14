// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeQueryVariablePageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeQueryVariablePageListRequest
	GetLang() *string
	SetCurrentPage(v int64) *DescribeQueryVariablePageListRequest
	GetCurrentPage() *int64
	SetDataSourceCode(v string) *DescribeQueryVariablePageListRequest
	GetDataSourceCode() *string
	SetEventCode(v string) *DescribeQueryVariablePageListRequest
	GetEventCode() *string
	SetName(v string) *DescribeQueryVariablePageListRequest
	GetName() *string
	SetPageSize(v int64) *DescribeQueryVariablePageListRequest
	GetPageSize() *int64
	SetRegId(v string) *DescribeQueryVariablePageListRequest
	GetRegId() *string
}

type DescribeQueryVariablePageListRequest struct {
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
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Data source code
	//
	// example:
	//
	// ds_vcaoii1697
	DataSourceCode *string `json:"dataSourceCode,omitempty" xml:"dataSourceCode,omitempty"`
	// Event code
	//
	// example:
	//
	// de_ahpayh4121
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Query variable name
	//
	// example:
	//
	// 名称
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Page size, default value is 10
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

func (s DescribeQueryVariablePageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeQueryVariablePageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeQueryVariablePageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeQueryVariablePageListRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeQueryVariablePageListRequest) GetDataSourceCode() *string {
	return s.DataSourceCode
}

func (s *DescribeQueryVariablePageListRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeQueryVariablePageListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeQueryVariablePageListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeQueryVariablePageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeQueryVariablePageListRequest) SetLang(v string) *DescribeQueryVariablePageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeQueryVariablePageListRequest) SetCurrentPage(v int64) *DescribeQueryVariablePageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeQueryVariablePageListRequest) SetDataSourceCode(v string) *DescribeQueryVariablePageListRequest {
	s.DataSourceCode = &v
	return s
}

func (s *DescribeQueryVariablePageListRequest) SetEventCode(v string) *DescribeQueryVariablePageListRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeQueryVariablePageListRequest) SetName(v string) *DescribeQueryVariablePageListRequest {
	s.Name = &v
	return s
}

func (s *DescribeQueryVariablePageListRequest) SetPageSize(v int64) *DescribeQueryVariablePageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeQueryVariablePageListRequest) SetRegId(v string) *DescribeQueryVariablePageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeQueryVariablePageListRequest) Validate() error {
	return dara.Validate(s)
}
