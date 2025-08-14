// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpressionVariablePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeExpressionVariablePageRequest
	GetLang() *string
	SetCurrentPage(v string) *DescribeExpressionVariablePageRequest
	GetCurrentPage() *string
	SetEventCode(v string) *DescribeExpressionVariablePageRequest
	GetEventCode() *string
	SetOutputs(v string) *DescribeExpressionVariablePageRequest
	GetOutputs() *string
	SetPageSize(v string) *DescribeExpressionVariablePageRequest
	GetPageSize() *string
	SetRegId(v string) *DescribeExpressionVariablePageRequest
	GetRegId() *string
	SetStatus(v string) *DescribeExpressionVariablePageRequest
	GetStatus() *string
	SetValue(v string) *DescribeExpressionVariablePageRequest
	GetValue() *string
}

type DescribeExpressionVariablePageRequest struct {
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
	CurrentPage *string `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aamexg3015,de_aamexg3342
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Variable return type
	//
	// example:
	//
	// STRING
	Outputs *string `json:"outputs,omitempty" xml:"outputs,omitempty"`
	// Page size, with a default value of 10
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *string `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Variable name/description
	//
	// example:
	//
	// 自定义变量
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s DescribeExpressionVariablePageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpressionVariablePageRequest) GoString() string {
	return s.String()
}

func (s *DescribeExpressionVariablePageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeExpressionVariablePageRequest) GetCurrentPage() *string {
	return s.CurrentPage
}

func (s *DescribeExpressionVariablePageRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeExpressionVariablePageRequest) GetOutputs() *string {
	return s.Outputs
}

func (s *DescribeExpressionVariablePageRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeExpressionVariablePageRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeExpressionVariablePageRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeExpressionVariablePageRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeExpressionVariablePageRequest) SetLang(v string) *DescribeExpressionVariablePageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeExpressionVariablePageRequest) SetCurrentPage(v string) *DescribeExpressionVariablePageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeExpressionVariablePageRequest) SetEventCode(v string) *DescribeExpressionVariablePageRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeExpressionVariablePageRequest) SetOutputs(v string) *DescribeExpressionVariablePageRequest {
	s.Outputs = &v
	return s
}

func (s *DescribeExpressionVariablePageRequest) SetPageSize(v string) *DescribeExpressionVariablePageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeExpressionVariablePageRequest) SetRegId(v string) *DescribeExpressionVariablePageRequest {
	s.RegId = &v
	return s
}

func (s *DescribeExpressionVariablePageRequest) SetStatus(v string) *DescribeExpressionVariablePageRequest {
	s.Status = &v
	return s
}

func (s *DescribeExpressionVariablePageRequest) SetValue(v string) *DescribeExpressionVariablePageRequest {
	s.Value = &v
	return s
}

func (s *DescribeExpressionVariablePageRequest) Validate() error {
	return dara.Validate(s)
}
