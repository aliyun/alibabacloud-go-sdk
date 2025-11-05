// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCommissionableProductsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetCommissionableProductsResponseBody
	GetCode() *string
	SetData(v []*GetCommissionableProductsResponseBodyData) *GetCommissionableProductsResponseBody
	GetData() []*GetCommissionableProductsResponseBodyData
	SetPageNo(v int32) *GetCommissionableProductsResponseBody
	GetPageNo() *int32
	SetPageSize(v int32) *GetCommissionableProductsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetCommissionableProductsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetCommissionableProductsResponseBody
	GetSuccess() *bool
	SetTotal(v int32) *GetCommissionableProductsResponseBody
	GetTotal() *int32
}

type GetCommissionableProductsResponseBody struct {
	// example:
	//
	// 200
	Code *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetCommissionableProductsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2103a30617045934095083027d88c5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 10
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetCommissionableProductsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCommissionableProductsResponseBody) GoString() string {
	return s.String()
}

func (s *GetCommissionableProductsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetCommissionableProductsResponseBody) GetData() []*GetCommissionableProductsResponseBodyData {
	return s.Data
}

func (s *GetCommissionableProductsResponseBody) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetCommissionableProductsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetCommissionableProductsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCommissionableProductsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetCommissionableProductsResponseBody) GetTotal() *int32 {
	return s.Total
}

func (s *GetCommissionableProductsResponseBody) SetCode(v string) *GetCommissionableProductsResponseBody {
	s.Code = &v
	return s
}

func (s *GetCommissionableProductsResponseBody) SetData(v []*GetCommissionableProductsResponseBodyData) *GetCommissionableProductsResponseBody {
	s.Data = v
	return s
}

func (s *GetCommissionableProductsResponseBody) SetPageNo(v int32) *GetCommissionableProductsResponseBody {
	s.PageNo = &v
	return s
}

func (s *GetCommissionableProductsResponseBody) SetPageSize(v int32) *GetCommissionableProductsResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetCommissionableProductsResponseBody) SetRequestId(v string) *GetCommissionableProductsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCommissionableProductsResponseBody) SetSuccess(v bool) *GetCommissionableProductsResponseBody {
	s.Success = &v
	return s
}

func (s *GetCommissionableProductsResponseBody) SetTotal(v int32) *GetCommissionableProductsResponseBody {
	s.Total = &v
	return s
}

func (s *GetCommissionableProductsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCommissionableProductsResponseBodyData struct {
	// example:
	//
	// “202502”
	ActualStartMonth *string `json:"ActualStartMonth,omitempty" xml:"ActualStartMonth,omitempty"`
	// example:
	//
	// “oceanbase_obpre_public_intl”
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// “ApsaraDB for OceanBase Pre”
	CommodityName *string `json:"CommodityName,omitempty" xml:"CommodityName,omitempty"`
	// example:
	//
	// Yes
	EligibleForCommission *string `json:"EligibleForCommission,omitempty" xml:"EligibleForCommission,omitempty"`
	// example:
	//
	// Yes
	EligibleForDiscount *string `json:"EligibleForDiscount,omitempty" xml:"EligibleForDiscount,omitempty"`
	// example:
	//
	// “202509”
	EndMonth *string `json:"EndMonth,omitempty" xml:"EndMonth,omitempty"`
	// example:
	//
	// Yes
	ProductCampaign *string `json:"ProductCampaign,omitempty" xml:"ProductCampaign,omitempty"`
	// example:
	//
	// “oceanbase”
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// “ApsaraDB for OceanBase”
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// Yes
	SevenCoreProducts *string `json:"SevenCoreProducts,omitempty" xml:"SevenCoreProducts,omitempty"`
	// example:
	//
	// “ONLINE”
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetCommissionableProductsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetCommissionableProductsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetCommissionableProductsResponseBodyData) GetActualStartMonth() *string {
	return s.ActualStartMonth
}

func (s *GetCommissionableProductsResponseBodyData) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetCommissionableProductsResponseBodyData) GetCommodityName() *string {
	return s.CommodityName
}

func (s *GetCommissionableProductsResponseBodyData) GetEligibleForCommission() *string {
	return s.EligibleForCommission
}

func (s *GetCommissionableProductsResponseBodyData) GetEligibleForDiscount() *string {
	return s.EligibleForDiscount
}

func (s *GetCommissionableProductsResponseBodyData) GetEndMonth() *string {
	return s.EndMonth
}

func (s *GetCommissionableProductsResponseBodyData) GetProductCampaign() *string {
	return s.ProductCampaign
}

func (s *GetCommissionableProductsResponseBodyData) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetCommissionableProductsResponseBodyData) GetProductName() *string {
	return s.ProductName
}

func (s *GetCommissionableProductsResponseBodyData) GetSevenCoreProducts() *string {
	return s.SevenCoreProducts
}

func (s *GetCommissionableProductsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetCommissionableProductsResponseBodyData) SetActualStartMonth(v string) *GetCommissionableProductsResponseBodyData {
	s.ActualStartMonth = &v
	return s
}

func (s *GetCommissionableProductsResponseBodyData) SetCommodityCode(v string) *GetCommissionableProductsResponseBodyData {
	s.CommodityCode = &v
	return s
}

func (s *GetCommissionableProductsResponseBodyData) SetCommodityName(v string) *GetCommissionableProductsResponseBodyData {
	s.CommodityName = &v
	return s
}

func (s *GetCommissionableProductsResponseBodyData) SetEligibleForCommission(v string) *GetCommissionableProductsResponseBodyData {
	s.EligibleForCommission = &v
	return s
}

func (s *GetCommissionableProductsResponseBodyData) SetEligibleForDiscount(v string) *GetCommissionableProductsResponseBodyData {
	s.EligibleForDiscount = &v
	return s
}

func (s *GetCommissionableProductsResponseBodyData) SetEndMonth(v string) *GetCommissionableProductsResponseBodyData {
	s.EndMonth = &v
	return s
}

func (s *GetCommissionableProductsResponseBodyData) SetProductCampaign(v string) *GetCommissionableProductsResponseBodyData {
	s.ProductCampaign = &v
	return s
}

func (s *GetCommissionableProductsResponseBodyData) SetProductCode(v string) *GetCommissionableProductsResponseBodyData {
	s.ProductCode = &v
	return s
}

func (s *GetCommissionableProductsResponseBodyData) SetProductName(v string) *GetCommissionableProductsResponseBodyData {
	s.ProductName = &v
	return s
}

func (s *GetCommissionableProductsResponseBodyData) SetSevenCoreProducts(v string) *GetCommissionableProductsResponseBodyData {
	s.SevenCoreProducts = &v
	return s
}

func (s *GetCommissionableProductsResponseBodyData) SetStatus(v string) *GetCommissionableProductsResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetCommissionableProductsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
