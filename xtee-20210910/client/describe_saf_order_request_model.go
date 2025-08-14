// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafOrderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeSafOrderRequest
	GetLang() *string
	SetCurrentPage(v int32) *DescribeSafOrderRequest
	GetCurrentPage() *int32
	SetEndDate(v string) *DescribeSafOrderRequest
	GetEndDate() *string
	SetExactProductCode(v string) *DescribeSafOrderRequest
	GetExactProductCode() *string
	SetPageSize(v int32) *DescribeSafOrderRequest
	GetPageSize() *int32
	SetRegId(v string) *DescribeSafOrderRequest
	GetRegId() *string
	SetStartDate(v string) *DescribeSafOrderRequest
	GetStartDate() *string
}

type DescribeSafOrderRequest struct {
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
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// End date.
	//
	// example:
	//
	// 1755076800000
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Product code.
	//
	// example:
	//
	// saf_de
	ExactProductCode *string `json:"exactProductCode,omitempty" xml:"exactProductCode,omitempty"`
	// Page size, default value is 10.
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
	// Start time.
	//
	// example:
	//
	// 1752076800000
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s DescribeSafOrderRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafOrderRequest) GoString() string {
	return s.String()
}

func (s *DescribeSafOrderRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSafOrderRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSafOrderRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeSafOrderRequest) GetExactProductCode() *string {
	return s.ExactProductCode
}

func (s *DescribeSafOrderRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSafOrderRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeSafOrderRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeSafOrderRequest) SetLang(v string) *DescribeSafOrderRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSafOrderRequest) SetCurrentPage(v int32) *DescribeSafOrderRequest {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSafOrderRequest) SetEndDate(v string) *DescribeSafOrderRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeSafOrderRequest) SetExactProductCode(v string) *DescribeSafOrderRequest {
	s.ExactProductCode = &v
	return s
}

func (s *DescribeSafOrderRequest) SetPageSize(v int32) *DescribeSafOrderRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSafOrderRequest) SetRegId(v string) *DescribeSafOrderRequest {
	s.RegId = &v
	return s
}

func (s *DescribeSafOrderRequest) SetStartDate(v string) *DescribeSafOrderRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeSafOrderRequest) Validate() error {
	return dara.Validate(s)
}
