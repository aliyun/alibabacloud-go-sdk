// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustVariablePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeCustVariablePageRequest
	GetLang() *string
	SetCreateType(v string) *DescribeCustVariablePageRequest
	GetCreateType() *string
	SetCurrentPage(v int32) *DescribeCustVariablePageRequest
	GetCurrentPage() *int32
	SetDescription(v string) *DescribeCustVariablePageRequest
	GetDescription() *string
	SetEventCode(v string) *DescribeCustVariablePageRequest
	GetEventCode() *string
	SetPageSize(v int32) *DescribeCustVariablePageRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeCustVariablePageRequest
	GetRegId() *string
	SetStatus(v string) *DescribeCustVariablePageRequest
	GetStatus() *string
}

type DescribeCustVariablePageRequest struct {
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
	// Creation type.
	//
	// example:
	//
	// NORMAL
	CreateType *string `json:"createType,omitempty" xml:"createType,omitempty"`
	// Pagination parameter, current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Description.
	//
	// example:
	//
	// 累计变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event code.
	//
	// example:
	//
	// de_aheldm3876
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Number of records per page, default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeCustVariablePageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustVariablePageRequest) GoString() string {
	return s.String()
}

func (s *DescribeCustVariablePageRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeCustVariablePageRequest) GetCreateType() *string {
	return s.CreateType
}

func (s *DescribeCustVariablePageRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCustVariablePageRequest) GetDescription() *string {
	return s.Description
}

func (s *DescribeCustVariablePageRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeCustVariablePageRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCustVariablePageRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeCustVariablePageRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeCustVariablePageRequest) SetLang(v string) *DescribeCustVariablePageRequest {
	s.Lang = &v
	return s
}

func (s *DescribeCustVariablePageRequest) SetCreateType(v string) *DescribeCustVariablePageRequest {
	s.CreateType = &v
	return s
}

func (s *DescribeCustVariablePageRequest) SetCurrentPage(v int32) *DescribeCustVariablePageRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCustVariablePageRequest) SetDescription(v string) *DescribeCustVariablePageRequest {
	s.Description = &v
	return s
}

func (s *DescribeCustVariablePageRequest) SetEventCode(v string) *DescribeCustVariablePageRequest {
	s.EventCode = &v
	return s
}

func (s *DescribeCustVariablePageRequest) SetPageSize(v int32) *DescribeCustVariablePageRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCustVariablePageRequest) SetRegId(v string) *DescribeCustVariablePageRequest {
	s.RegId = &v
	return s
}

func (s *DescribeCustVariablePageRequest) SetStatus(v string) *DescribeCustVariablePageRequest {
	s.Status = &v
	return s
}

func (s *DescribeCustVariablePageRequest) Validate() error {
	return dara.Validate(s)
}
