// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataSourcePageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDataSourcePageListRequest
	GetLang() *string
	SetSourceIp(v string) *DescribeDataSourcePageListRequest
	GetSourceIp() *string
	SetCurrentPage(v int64) *DescribeDataSourcePageListRequest
	GetCurrentPage() *int64
	SetName(v string) *DescribeDataSourcePageListRequest
	GetName() *string
	SetPageSize(v int64) *DescribeDataSourcePageListRequest
	GetPageSize() *int64
	SetRegId(v string) *DescribeDataSourcePageListRequest
	GetRegId() *string
	SetType(v string) *DescribeDataSourcePageListRequest
	GetType() *string
}

type DescribeDataSourcePageListRequest struct {
	// Set the language type for request and response, default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Request source IP.
	//
	// example:
	//
	// 220.250.21.83
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Variable name
	//
	// example:
	//
	// data_source
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Page size, default value is 10
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Type
	//
	// example:
	//
	// FILE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeDataSourcePageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataSourcePageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeDataSourcePageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDataSourcePageListRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeDataSourcePageListRequest) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeDataSourcePageListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeDataSourcePageListRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDataSourcePageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeDataSourcePageListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDataSourcePageListRequest) SetLang(v string) *DescribeDataSourcePageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDataSourcePageListRequest) SetSourceIp(v string) *DescribeDataSourcePageListRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeDataSourcePageListRequest) SetCurrentPage(v int64) *DescribeDataSourcePageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeDataSourcePageListRequest) SetName(v string) *DescribeDataSourcePageListRequest {
	s.Name = &v
	return s
}

func (s *DescribeDataSourcePageListRequest) SetPageSize(v int64) *DescribeDataSourcePageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDataSourcePageListRequest) SetRegId(v string) *DescribeDataSourcePageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeDataSourcePageListRequest) SetType(v string) *DescribeDataSourcePageListRequest {
	s.Type = &v
	return s
}

func (s *DescribeDataSourcePageListRequest) Validate() error {
	return dara.Validate(s)
}
