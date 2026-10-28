// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommissionableProductsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommodityCodeList(v string) *GetCommissionableProductsRequest
	GetCommodityCodeList() *string
	SetFiscalYear(v string) *GetCommissionableProductsRequest
	GetFiscalYear() *string
	SetListShowStatusList(v []*string) *GetCommissionableProductsRequest
	GetListShowStatusList() []*string
	SetPageNo(v int32) *GetCommissionableProductsRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetCommissionableProductsRequest
	GetPageSize() *int32
	SetPipCodeList(v string) *GetCommissionableProductsRequest
	GetPipCodeList() *string
	SetRealEndMonth(v string) *GetCommissionableProductsRequest
	GetRealEndMonth() *string
	SetRealStartMonth(v string) *GetCommissionableProductsRequest
	GetRealStartMonth() *string
}

type GetCommissionableProductsRequest struct {
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
	ListShowStatusList []*string `json:"ListShowStatusList,omitempty" xml:"ListShowStatusList,omitempty" type:"Repeated"`
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

func (s GetCommissionableProductsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCommissionableProductsRequest) GoString() string {
	return s.String()
}

func (s *GetCommissionableProductsRequest) GetCommodityCodeList() *string {
	return s.CommodityCodeList
}

func (s *GetCommissionableProductsRequest) GetFiscalYear() *string {
	return s.FiscalYear
}

func (s *GetCommissionableProductsRequest) GetListShowStatusList() []*string {
	return s.ListShowStatusList
}

func (s *GetCommissionableProductsRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetCommissionableProductsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCommissionableProductsRequest) GetPipCodeList() *string {
	return s.PipCodeList
}

func (s *GetCommissionableProductsRequest) GetRealEndMonth() *string {
	return s.RealEndMonth
}

func (s *GetCommissionableProductsRequest) GetRealStartMonth() *string {
	return s.RealStartMonth
}

func (s *GetCommissionableProductsRequest) SetCommodityCodeList(v string) *GetCommissionableProductsRequest {
	s.CommodityCodeList = &v
	return s
}

func (s *GetCommissionableProductsRequest) SetFiscalYear(v string) *GetCommissionableProductsRequest {
	s.FiscalYear = &v
	return s
}

func (s *GetCommissionableProductsRequest) SetListShowStatusList(v []*string) *GetCommissionableProductsRequest {
	s.ListShowStatusList = v
	return s
}

func (s *GetCommissionableProductsRequest) SetPageNo(v int32) *GetCommissionableProductsRequest {
	s.PageNo = &v
	return s
}

func (s *GetCommissionableProductsRequest) SetPageSize(v int32) *GetCommissionableProductsRequest {
	s.PageSize = &v
	return s
}

func (s *GetCommissionableProductsRequest) SetPipCodeList(v string) *GetCommissionableProductsRequest {
	s.PipCodeList = &v
	return s
}

func (s *GetCommissionableProductsRequest) SetRealEndMonth(v string) *GetCommissionableProductsRequest {
	s.RealEndMonth = &v
	return s
}

func (s *GetCommissionableProductsRequest) SetRealStartMonth(v string) *GetCommissionableProductsRequest {
	s.RealStartMonth = &v
	return s
}

func (s *GetCommissionableProductsRequest) Validate() error {
	return dara.Validate(s)
}
