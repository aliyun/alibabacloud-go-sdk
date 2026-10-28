// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommissionableProductsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCodeList(v string) *GetCommissionableProductsShrinkRequest
	GetCommodityCodeList() *string
	SetFiscalYear(v string) *GetCommissionableProductsShrinkRequest
	GetFiscalYear() *string
	SetListShowStatusListShrink(v string) *GetCommissionableProductsShrinkRequest
	GetListShowStatusListShrink() *string
	SetPageNo(v int32) *GetCommissionableProductsShrinkRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetCommissionableProductsShrinkRequest
	GetPageSize() *int32
	SetPipCodeList(v string) *GetCommissionableProductsShrinkRequest
	GetPipCodeList() *string
	SetRealEndMonth(v string) *GetCommissionableProductsShrinkRequest
	GetRealEndMonth() *string
	SetRealStartMonth(v string) *GetCommissionableProductsShrinkRequest
	GetRealStartMonth() *string
}

type GetCommissionableProductsShrinkRequest struct {
	// The commodity codes. Separate multiple codes with commas.
	//
	// example:
	//
	// "oceanbase_obpre_public_intl,savingplan_common_public_intl"
	CommodityCodeList *string `json:"CommodityCodeList,omitempty" xml:"CommodityCodeList,omitempty"`
	// The fiscal year. If this parameter is left empty, the current fiscal year is used by default.
	//
	// example:
	//
	// “FY26”
	FiscalYear *string `json:"FiscalYear,omitempty" xml:"FiscalYear,omitempty"`
	// The list of display statuses for commission-eligible commodities.
	ListShowStatusListShrink *string `json:"ListShowStatusList,omitempty" xml:"ListShowStatusList,omitempty"`
	// The current page number. Pages start from 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 100.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product codes. Separate multiple codes with commas.
	//
	// example:
	//
	// "oceanbase"
	PipCodeList *string `json:"PipCodeList,omitempty" xml:"PipCodeList,omitempty"`
	// The actual end month in the yyyymm format.
	//
	// example:
	//
	// “202509”
	RealEndMonth *string `json:"RealEndMonth,omitempty" xml:"RealEndMonth,omitempty"`
	// The actual start month in the yyyymm format.
	//
	// example:
	//
	// “202502”
	RealStartMonth *string `json:"RealStartMonth,omitempty" xml:"RealStartMonth,omitempty"`
}

func (s GetCommissionableProductsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCommissionableProductsShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetCommissionableProductsShrinkRequest) GetCommodityCodeList() *string {
	return s.CommodityCodeList
}

func (s *GetCommissionableProductsShrinkRequest) GetFiscalYear() *string {
	return s.FiscalYear
}

func (s *GetCommissionableProductsShrinkRequest) GetListShowStatusListShrink() *string {
	return s.ListShowStatusListShrink
}

func (s *GetCommissionableProductsShrinkRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetCommissionableProductsShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCommissionableProductsShrinkRequest) GetPipCodeList() *string {
	return s.PipCodeList
}

func (s *GetCommissionableProductsShrinkRequest) GetRealEndMonth() *string {
	return s.RealEndMonth
}

func (s *GetCommissionableProductsShrinkRequest) GetRealStartMonth() *string {
	return s.RealStartMonth
}

func (s *GetCommissionableProductsShrinkRequest) SetCommodityCodeList(v string) *GetCommissionableProductsShrinkRequest {
	s.CommodityCodeList = &v
	return s
}

func (s *GetCommissionableProductsShrinkRequest) SetFiscalYear(v string) *GetCommissionableProductsShrinkRequest {
	s.FiscalYear = &v
	return s
}

func (s *GetCommissionableProductsShrinkRequest) SetListShowStatusListShrink(v string) *GetCommissionableProductsShrinkRequest {
	s.ListShowStatusListShrink = &v
	return s
}

func (s *GetCommissionableProductsShrinkRequest) SetPageNo(v int32) *GetCommissionableProductsShrinkRequest {
	s.PageNo = &v
	return s
}

func (s *GetCommissionableProductsShrinkRequest) SetPageSize(v int32) *GetCommissionableProductsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *GetCommissionableProductsShrinkRequest) SetPipCodeList(v string) *GetCommissionableProductsShrinkRequest {
	s.PipCodeList = &v
	return s
}

func (s *GetCommissionableProductsShrinkRequest) SetRealEndMonth(v string) *GetCommissionableProductsShrinkRequest {
	s.RealEndMonth = &v
	return s
}

func (s *GetCommissionableProductsShrinkRequest) SetRealStartMonth(v string) *GetCommissionableProductsShrinkRequest {
	s.RealStartMonth = &v
	return s
}

func (s *GetCommissionableProductsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
