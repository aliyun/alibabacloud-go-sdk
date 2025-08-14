// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNameListRequest
	GetLang() *string
	SetCreateType(v string) *DescribeNameListRequest
	GetCreateType() *string
	SetCurrentPage(v string) *DescribeNameListRequest
	GetCurrentPage() *string
	SetPageSize(v string) *DescribeNameListRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeNameListRequest
	GetRegId() *string
	SetValue(v string) *DescribeNameListRequest
	GetValue() *string
	SetVariableId(v string) *DescribeNameListRequest
	GetVariableId() *string
}

type DescribeNameListRequest struct {
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
	// Creation type.
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, with a default value of 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Search value.
	//
	// example:
	//
	// 白名单
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// Variable ID.
	//
	// example:
	//
	// 393314
	VariableId *string `json:"variableId,omitempty" xml:"variableId,omitempty"`
}

func (s DescribeNameListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNameListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNameListRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeNameListRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeNameListRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeNameListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeNameListRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeNameListRequest) GetVariableId() *string {
	return s.VariableId
}

func (s *DescribeNameListRequest) SetLang(v string) *DescribeNameListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNameListRequest) SetCreateType(v string) *DescribeNameListRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeNameListRequest) SetCurrentPage(v string) *DescribeNameListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNameListRequest) SetPageSize(v string) *DescribeNameListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNameListRequest) SetRegId(v string) *DescribeNameListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeNameListRequest) SetValue(v string) *DescribeNameListRequest {
	s.Value = &v
	return s
}

func (s *DescribeNameListRequest) SetVariableId(v string) *DescribeNameListRequest {
	s.VariableId = &v
	return s
}

func (s *DescribeNameListRequest) Validate() error {
	return dara.Validate(s)
}
