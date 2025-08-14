// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListVariablePageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNameListVariablePageListRequest
	GetLang() *string
	SetCurrentPage(v int32) *DescribeNameListVariablePageListRequest
	GetCurrentPage() *int32
	SetName(v string) *DescribeNameListVariablePageListRequest
	GetName() *string
	SetNameListType(v string) *DescribeNameListVariablePageListRequest
	GetNameListType() *string
	SetPageSize(v int32) *DescribeNameListVariablePageListRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeNameListVariablePageListRequest
	GetRegId() *string
	SetValue(v string) *DescribeNameListVariablePageListRequest
	GetValue() *string
}

type DescribeNameListVariablePageListRequest struct {
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
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Variable name
	//
	// example:
	//
	// age
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Variable type
	//
	// example:
	//
	// accountId
	NameListType *string `json:"nameListType,omitempty" xml:"nameListType,omitempty"`
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
	// Name list value
	//
	// example:
	//
	// valuexxx
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeNameListVariablePageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListVariablePageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNameListVariablePageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNameListVariablePageListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeNameListVariablePageListRequest) GetName() *string {
	return s.Name
}

func (s *DescribeNameListVariablePageListRequest) GetNameListType() *string {
	return s.NameListType
}

func (s *DescribeNameListVariablePageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNameListVariablePageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeNameListVariablePageListRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeNameListVariablePageListRequest) SetLang(v string) *DescribeNameListVariablePageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNameListVariablePageListRequest) SetCurrentPage(v int32) *DescribeNameListVariablePageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNameListVariablePageListRequest) SetName(v string) *DescribeNameListVariablePageListRequest {
	s.Name = &v
	return s
}

func (s *DescribeNameListVariablePageListRequest) SetNameListType(v string) *DescribeNameListVariablePageListRequest {
	s.NameListType = &v
	return s
}

func (s *DescribeNameListVariablePageListRequest) SetPageSize(v int32) *DescribeNameListVariablePageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNameListVariablePageListRequest) SetRegId(v string) *DescribeNameListVariablePageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeNameListVariablePageListRequest) SetValue(v string) *DescribeNameListVariablePageListRequest {
	s.Value = &v
	return s
}

func (s *DescribeNameListVariablePageListRequest) Validate() error {
	return dara.Validate(s)
}
