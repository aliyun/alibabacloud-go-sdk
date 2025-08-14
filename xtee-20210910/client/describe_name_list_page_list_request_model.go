// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListPageListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeNameListPageListRequest
	GetLang() *string
	SetCurrentPage(v int32) *DescribeNameListPageListRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeNameListPageListRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeNameListPageListRequest
	GetRegId() *string
	SetUpdateBeginTime(v int64) *DescribeNameListPageListRequest
	GetUpdateBeginTime() *int64
	SetUpdateEndTime(v int64) *DescribeNameListPageListRequest
	GetUpdateEndTime() *int64
	SetValue(v string) *DescribeNameListPageListRequest
	GetValue() *string
	SetVariableId(v int64) *DescribeNameListPageListRequest
	GetVariableId() *int64
}

type DescribeNameListPageListRequest struct {
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
	// 3
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
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
	// Update start time.
	//
	// example:
	//
	// 1753372800000
	UpdateBeginTime *int64 `json:"updateBeginTime,omitempty" xml:"updateBeginTime,omitempty"`
	// Update end time.
	//
	// example:
	//
	// 1753459199059
	UpdateEndTime *int64 `json:"updateEndTime,omitempty" xml:"updateEndTime,omitempty"`
	// Variable name/description
	//
	// example:
	//
	// 白名单
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
	// Variable ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 762
	VariableId *int64 `json:"variableId,omitempty" xml:"variableId,omitempty"`
}

func (s DescribeNameListPageListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListPageListRequest) GoString() string {
	return s.String()
}

func (s *DescribeNameListPageListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNameListPageListRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeNameListPageListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNameListPageListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeNameListPageListRequest) GetUpdateBeginTime() *int64 {
	return s.UpdateBeginTime
}

func (s *DescribeNameListPageListRequest) GetUpdateEndTime() *int64 {
	return s.UpdateEndTime
}

func (s *DescribeNameListPageListRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeNameListPageListRequest) GetVariableId() *int64 {
	return s.VariableId
}

func (s *DescribeNameListPageListRequest) SetLang(v string) *DescribeNameListPageListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNameListPageListRequest) SetCurrentPage(v int32) *DescribeNameListPageListRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeNameListPageListRequest) SetPageSize(v int32) *DescribeNameListPageListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNameListPageListRequest) SetRegId(v string) *DescribeNameListPageListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeNameListPageListRequest) SetUpdateBeginTime(v int64) *DescribeNameListPageListRequest {
	s.UpdateBeginTime = &v
	return s
}

func (s *DescribeNameListPageListRequest) SetUpdateEndTime(v int64) *DescribeNameListPageListRequest {
	s.UpdateEndTime = &v
	return s
}

func (s *DescribeNameListPageListRequest) SetValue(v string) *DescribeNameListPageListRequest {
	s.Value = &v
	return s
}

func (s *DescribeNameListPageListRequest) SetVariableId(v int64) *DescribeNameListPageListRequest {
	s.VariableId = &v
	return s
}

func (s *DescribeNameListPageListRequest) Validate() error {
	return dara.Validate(s)
}
