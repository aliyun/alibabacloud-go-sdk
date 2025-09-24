// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPayAsYouGoPriceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetPayAsYouGoPriceResponseBody
	GetCode() *string
	SetData(v *GetPayAsYouGoPriceResponseBodyData) *GetPayAsYouGoPriceResponseBody
	GetData() *GetPayAsYouGoPriceResponseBodyData
	SetMessage(v string) *GetPayAsYouGoPriceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetPayAsYouGoPriceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetPayAsYouGoPriceResponseBody
	GetSuccess() *bool
}

type GetPayAsYouGoPriceResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetPayAsYouGoPriceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1537A007-72D7-4165-8A26-8694A38E219A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetPayAsYouGoPriceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBody) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetPayAsYouGoPriceResponseBody) GetData() *GetPayAsYouGoPriceResponseBodyData {
	return s.Data
}

func (s *GetPayAsYouGoPriceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetPayAsYouGoPriceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetPayAsYouGoPriceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetPayAsYouGoPriceResponseBody) SetCode(v string) *GetPayAsYouGoPriceResponseBody {
	s.Code = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBody) SetData(v *GetPayAsYouGoPriceResponseBodyData) *GetPayAsYouGoPriceResponseBody {
	s.Data = v
	return s
}

func (s *GetPayAsYouGoPriceResponseBody) SetMessage(v string) *GetPayAsYouGoPriceResponseBody {
	s.Message = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBody) SetRequestId(v string) *GetPayAsYouGoPriceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBody) SetSuccess(v bool) *GetPayAsYouGoPriceResponseBody {
	s.Success = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetPayAsYouGoPriceResponseBodyData struct {
	// The type of the currency. Valid values:
	//
	// 	- CNY: Chinese Yuan
	//
	// 	- USD: US dollar
	//
	// 	- JPY: Japanese Yen
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The price details of the pricing module.
	ModuleDetails *GetPayAsYouGoPriceResponseBodyDataModuleDetails `json:"ModuleDetails,omitempty" xml:"ModuleDetails,omitempty" type:"Struct"`
	// The details of the discount.
	PromotionDetails *GetPayAsYouGoPriceResponseBodyDataPromotionDetails `json:"PromotionDetails,omitempty" xml:"PromotionDetails,omitempty" type:"Struct"`
}

func (s GetPayAsYouGoPriceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBodyData) GetCurrency() *string {
	return s.Currency
}

func (s *GetPayAsYouGoPriceResponseBodyData) GetModuleDetails() *GetPayAsYouGoPriceResponseBodyDataModuleDetails {
	return s.ModuleDetails
}

func (s *GetPayAsYouGoPriceResponseBodyData) GetPromotionDetails() *GetPayAsYouGoPriceResponseBodyDataPromotionDetails {
	return s.PromotionDetails
}

func (s *GetPayAsYouGoPriceResponseBodyData) SetCurrency(v string) *GetPayAsYouGoPriceResponseBodyData {
	s.Currency = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyData) SetModuleDetails(v *GetPayAsYouGoPriceResponseBodyDataModuleDetails) *GetPayAsYouGoPriceResponseBodyData {
	s.ModuleDetails = v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyData) SetPromotionDetails(v *GetPayAsYouGoPriceResponseBodyDataPromotionDetails) *GetPayAsYouGoPriceResponseBodyData {
	s.PromotionDetails = v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetPayAsYouGoPriceResponseBodyDataModuleDetails struct {
	ModuleDetail []*GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail `json:"ModuleDetail,omitempty" xml:"ModuleDetail,omitempty" type:"Repeated"`
}

func (s GetPayAsYouGoPriceResponseBodyDataModuleDetails) String() string {
	return dara.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBodyDataModuleDetails) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetails) GetModuleDetail() []*GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail {
	return s.ModuleDetail
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetails) SetModuleDetail(v []*GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) *GetPayAsYouGoPriceResponseBodyDataModuleDetails {
	s.ModuleDetail = v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetails) Validate() error {
	return dara.Validate(s)
}

type GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail struct {
	// The discount price.
	//
	// example:
	//
	// 100
	CostAfterDiscount *float32 `json:"CostAfterDiscount,omitempty" xml:"CostAfterDiscount,omitempty"`
	// The discount that was applied.
	//
	// example:
	//
	// 20
	InvoiceDiscount *float32 `json:"InvoiceDiscount,omitempty" xml:"InvoiceDiscount,omitempty"`
	// The code of the pricing module.
	//
	// example:
	//
	// InstanceType
	ModuleCode *string `json:"ModuleCode,omitempty" xml:"ModuleCode,omitempty"`
	// The original price.
	//
	// example:
	//
	// 1.77
	OriginalCost *float32 `json:"OriginalCost,omitempty" xml:"OriginalCost,omitempty"`
	// The unit price.
	//
	// example:
	//
	// 0
	UnitPrice *float32 `json:"UnitPrice,omitempty" xml:"UnitPrice,omitempty"`
}

func (s GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) String() string {
	return dara.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) GetCostAfterDiscount() *float32 {
	return s.CostAfterDiscount
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) GetInvoiceDiscount() *float32 {
	return s.InvoiceDiscount
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) GetModuleCode() *string {
	return s.ModuleCode
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) GetOriginalCost() *float32 {
	return s.OriginalCost
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) GetUnitPrice() *float32 {
	return s.UnitPrice
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) SetCostAfterDiscount(v float32) *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail {
	s.CostAfterDiscount = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) SetInvoiceDiscount(v float32) *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail {
	s.InvoiceDiscount = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) SetModuleCode(v string) *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail {
	s.ModuleCode = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) SetOriginalCost(v float32) *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail {
	s.OriginalCost = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) SetUnitPrice(v float32) *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail {
	s.UnitPrice = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataModuleDetailsModuleDetail) Validate() error {
	return dara.Validate(s)
}

type GetPayAsYouGoPriceResponseBodyDataPromotionDetails struct {
	PromotionDetail []*GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail `json:"PromotionDetail,omitempty" xml:"PromotionDetail,omitempty" type:"Repeated"`
}

func (s GetPayAsYouGoPriceResponseBodyDataPromotionDetails) String() string {
	return dara.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBodyDataPromotionDetails) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetails) GetPromotionDetail() []*GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail {
	return s.PromotionDetail
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetails) SetPromotionDetail(v []*GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) *GetPayAsYouGoPriceResponseBodyDataPromotionDetails {
	s.PromotionDetail = v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetails) Validate() error {
	return dara.Validate(s)
}

type GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail struct {
	// The description of the discount.
	//
	// example:
	//
	// This discount allows you to use a service at the minimum price and is provided for testing purposes only.
	PromotionDesc *string `json:"PromotionDesc,omitempty" xml:"PromotionDesc,omitempty"`
	// The ID of the discount.
	//
	// example:
	//
	// 10200210
	PromotionId *int64 `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The name of the discount.
	//
	// example:
	//
	// This discount allows you to use a service at the minimum price and is provided for testing purposes only.
	PromotionName *string `json:"PromotionName,omitempty" xml:"PromotionName,omitempty"`
}

func (s GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) String() string {
	return dara.Prettify(s)
}

func (s GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) GoString() string {
	return s.String()
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) GetPromotionDesc() *string {
	return s.PromotionDesc
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) GetPromotionId() *int64 {
	return s.PromotionId
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) GetPromotionName() *string {
	return s.PromotionName
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionDesc(v string) *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionDesc = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionId(v int64) *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionId = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) SetPromotionName(v string) *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail {
	s.PromotionName = &v
	return s
}

func (s *GetPayAsYouGoPriceResponseBodyDataPromotionDetailsPromotionDetail) Validate() error {
	return dara.Validate(s)
}
