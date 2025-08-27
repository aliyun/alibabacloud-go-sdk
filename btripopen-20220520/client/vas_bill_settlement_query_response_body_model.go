// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVasBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VasBillSettlementQueryResponseBody
	GetCode() *string
	SetMessage(v string) *VasBillSettlementQueryResponseBody
	GetMessage() *string
	SetModule(v *VasBillSettlementQueryResponseBodyModule) *VasBillSettlementQueryResponseBody
	GetModule() *VasBillSettlementQueryResponseBodyModule
	SetRequestId(v string) *VasBillSettlementQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VasBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *VasBillSettlementQueryResponseBody
	GetTraceId() *string
}

type VasBillSettlementQueryResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// ok
	Message *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module  *VasBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// A5009956-1077-52FB-B520-EA8C7E91D722
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// trace_id
	//
	// example:
	//
	// 3b52152017470153218107062d0096
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s VasBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VasBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *VasBillSettlementQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *VasBillSettlementQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VasBillSettlementQueryResponseBody) GetModule() *VasBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *VasBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VasBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VasBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *VasBillSettlementQueryResponseBody) SetCode(v string) *VasBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *VasBillSettlementQueryResponseBody) SetMessage(v string) *VasBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *VasBillSettlementQueryResponseBody) SetModule(v *VasBillSettlementQueryResponseBodyModule) *VasBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *VasBillSettlementQueryResponseBody) SetRequestId(v string) *VasBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBody) SetSuccess(v bool) *VasBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *VasBillSettlementQueryResponseBody) SetTraceId(v string) *VasBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type VasBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 21
	Category *int32                                           `json:"category,omitempty" xml:"category,omitempty"`
	CorpId   *string                                          `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Items    []*VasBillSettlementQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-07-02
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2022-07-01
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	// example:
	//
	// CAESBgoEIgIIABgAIhkKFwMSAAAAMUw4MDAwMDAwMDA2ZTFjZTNi
	ScrollId *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// 30
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s VasBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s VasBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *VasBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *VasBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *VasBillSettlementQueryResponseBodyModule) GetItems() []*VasBillSettlementQueryResponseBodyModuleItems {
	return s.Items
}

func (s *VasBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *VasBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *VasBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *VasBillSettlementQueryResponseBodyModule) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *VasBillSettlementQueryResponseBodyModule) SetCategory(v int32) *VasBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModule) SetCorpId(v string) *VasBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModule) SetItems(v []*VasBillSettlementQueryResponseBodyModuleItems) *VasBillSettlementQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *VasBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *VasBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModule) SetScrollId(v string) *VasBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModule) SetTotalSize(v int64) *VasBillSettlementQueryResponseBodyModule {
	s.TotalSize = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type VasBillSettlementQueryResponseBodyModuleItems struct {
	// example:
	//
	// 2025-01-01 00:00:00
	AdjustTime *string `json:"adjust_time,omitempty" xml:"adjust_time,omitempty"`
	// example:
	//
	// 123aaa
	AlipayId *string `json:"alipay_id,omitempty" xml:"alipay_id,omitempty"`
	// example:
	//
	// 124
	AlipayTradeNo *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// example:
	//
	// 123
	ApplyArrCityCode *string `json:"apply_arr_city_code,omitempty" xml:"apply_arr_city_code,omitempty"`
	ApplyArrCityName *string `json:"apply_arr_city_name,omitempty" xml:"apply_arr_city_name,omitempty"`
	// example:
	//
	// 123
	ApplyDepCityCode *string `json:"apply_dep_city_code,omitempty" xml:"apply_dep_city_code,omitempty"`
	ApplyDepCityName *string `json:"apply_dep_city_name,omitempty" xml:"apply_dep_city_name,omitempty"`
	ApplyExtendField *string `json:"apply_extend_field,omitempty" xml:"apply_extend_field,omitempty"`
	// example:
	//
	// 1424070910031252025
	ApplyId        *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	BelongBusiness *string `json:"belong_business,omitempty" xml:"belong_business,omitempty"`
	// example:
	//
	// 2023-01-01 00:00:00
	BillRecordTime *string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BillingEntity  *string `json:"billing_entity,omitempty" xml:"billing_entity,omitempty"`
	BookMode       *string `json:"book_mode,omitempty" xml:"book_mode,omitempty"`
	// example:
	//
	// 2023-01-01 00:00:00
	BookTime *string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// example:
	//
	// 123
	BookerId *string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	// example:
	//
	// zs123
	BookerJobNo       *string `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName        *string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	CapitalDirection  *string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment *string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CategoryDesc      *string `json:"category_desc,omitempty" xml:"category_desc,omitempty"`
	CostCenter        *string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// example:
	//
	// cs1
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	// example:
	//
	// code1
	CostDepartment *string `json:"cost_department,omitempty" xml:"cost_department,omitempty"`
	Department     *string `json:"department,omitempty" xml:"department,omitempty"`
	// example:
	//
	// 1112
	DepartmentId *string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	FeeType      *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc  *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	// example:
	//
	// 1
	Index        *string `json:"index,omitempty" xml:"index,omitempty"`
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// example:
	//
	// q1
	MappingCompanyCode *string `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	// example:
	//
	// 1007025201876066223
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 100.0
	OrderPrice      *float64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	OrderStatusDesc *string  `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// example:
	//
	// 123
	OverApplyId *string `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	// example:
	//
	// 123
	PaymentDepartmentId   *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// 87687788
	PrimaryId *int64 `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	// example:
	//
	// www123
	ProcessorOaCode *string `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// 1
	ProductCount *int32 `json:"product_count,omitempty" xml:"product_count,omitempty"`
	// example:
	//
	// 111
	ProductId   *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty"`
	// example:
	//
	// acs
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// example:
	//
	// 0.0
	PromotionFee *float64 `json:"promotion_fee,omitempty" xml:"promotion_fee,omitempty"`
	// example:
	//
	// 111224324
	PurchaseOrderId *string `json:"purchase_order_id,omitempty" xml:"purchase_order_id,omitempty"`
	Remark          *string `json:"remark,omitempty" xml:"remark,omitempty"`
	SettleTypeDesc  *string `json:"settle_type_desc,omitempty" xml:"settle_type_desc,omitempty"`
	// example:
	//
	// 110.0
	SettlementFee *float64 `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	// example:
	//
	// 0.0
	SettlementGrantFee *float64 `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	// example:
	//
	// 2023-01-01 00:00:00
	SettlementTime *string `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	SettlementType *string `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	Specification  *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// example:
	//
	// 2
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// 185025497
	SubOrderId *string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// example:
	//
	// 9%
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// example:
	//
	// cs2
	ThirdInvoiceId *string `json:"third_invoice_id,omitempty" xml:"third_invoice_id,omitempty"`
	// example:
	//
	// 11
	ThirdItineraryId *string `json:"third_itinerary_id,omitempty" xml:"third_itinerary_id,omitempty"`
	TradeActionDesc  *string `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
	TradeRemark      *string `json:"trade_remark,omitempty" xml:"trade_remark,omitempty"`
	// example:
	//
	// 254
	TravelerId *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// example:
	//
	// A1234
	TravelerJobNo          *string `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerMemberType     *string `json:"traveler_member_type,omitempty" xml:"traveler_member_type,omitempty"`
	TravelerMemberTypeName *string `json:"traveler_member_type_name,omitempty" xml:"traveler_member_type_name,omitempty"`
	TravelerName           *string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	// example:
	//
	// 1
	VoucherType     *int32  `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	VoucherTypeDesc *string `json:"voucher_type_desc,omitempty" xml:"voucher_type_desc,omitempty"`
}

func (s VasBillSettlementQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s VasBillSettlementQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetAlipayId() *string {
	return s.AlipayId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetApplyArrCityCode() *string {
	return s.ApplyArrCityCode
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetApplyArrCityName() *string {
	return s.ApplyArrCityName
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetApplyDepCityCode() *string {
	return s.ApplyDepCityCode
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetApplyDepCityName() *string {
	return s.ApplyDepCityName
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetApplyId() *string {
	return s.ApplyId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetBelongBusiness() *string {
	return s.BelongBusiness
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetBillingEntity() *string {
	return s.BillingEntity
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetBookMode() *string {
	return s.BookMode
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetBookTime() *string {
	return s.BookTime
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetBookerId() *string {
	return s.BookerId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetBookerName() *string {
	return s.BookerName
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetCostCenter() *string {
	return s.CostCenter
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetDepartment() *string {
	return s.Department
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetFeeType() *string {
	return s.FeeType
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetIndex() *string {
	return s.Index
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetOrderId() *string {
	return s.OrderId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetOrderPrice() *float64 {
	return s.OrderPrice
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetOverApplyId() *string {
	return s.OverApplyId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetProductCount() *int32 {
	return s.ProductCount
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetProductId() *string {
	return s.ProductId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetProductName() *string {
	return s.ProductName
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetProjectName() *string {
	return s.ProjectName
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetPromotionFee() *float64 {
	return s.PromotionFee
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetPurchaseOrderId() *string {
	return s.PurchaseOrderId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetRemark() *string {
	return s.Remark
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetSettlementGrantFee() *float64 {
	return s.SettlementGrantFee
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetSettlementType() *string {
	return s.SettlementType
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetSpecification() *string {
	return s.Specification
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetStatus() *int32 {
	return s.Status
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetTaxRate() *string {
	return s.TaxRate
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetThirdInvoiceId() *string {
	return s.ThirdInvoiceId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetThirdItineraryId() *string {
	return s.ThirdItineraryId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetTradeRemark() *string {
	return s.TradeRemark
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetTravelerId() *string {
	return s.TravelerId
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberTypeName() *string {
	return s.TravelerMemberTypeName
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetTravelerName() *string {
	return s.TravelerName
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetAdjustTime(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.AdjustTime = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetAlipayId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.AlipayId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetAlipayTradeNo(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.AlipayTradeNo = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetApplyArrCityCode(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ApplyArrCityCode = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetApplyArrCityName(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ApplyArrCityName = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetApplyDepCityCode(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ApplyDepCityCode = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetApplyDepCityName(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ApplyDepCityName = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetApplyExtendField(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ApplyExtendField = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetApplyId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ApplyId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetBelongBusiness(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.BelongBusiness = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetBillRecordTime(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.BillRecordTime = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetBillingEntity(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.BillingEntity = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetBookMode(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.BookMode = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetBookTime(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.BookTime = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetBookerId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.BookerId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetBookerJobNo(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.BookerJobNo = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetBookerName(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.BookerName = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetCapitalDirection(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.CapitalDirection = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetCascadeDepartment(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.CascadeDepartment = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetCategoryDesc(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.CategoryDesc = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetCostCenter(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.CostCenter = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetCostCenterNumber(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.CostCenterNumber = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetCostDepartment(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.CostDepartment = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetDepartment(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.Department = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetDepartmentId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.DepartmentId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetFeeType(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.FeeType = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetFeeTypeDesc(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.FeeTypeDesc = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetIndex(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.Index = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetInvoiceTitle(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.InvoiceTitle = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetMappingCompanyCode(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.MappingCompanyCode = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetOrderId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.OrderId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetOrderPrice(v float64) *VasBillSettlementQueryResponseBodyModuleItems {
	s.OrderPrice = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetOrderStatusDesc(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.OrderStatusDesc = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetOverApplyId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.OverApplyId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetPaymentDepartmentId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.PaymentDepartmentId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetPaymentDepartmentName(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.PaymentDepartmentName = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetPrimaryId(v int64) *VasBillSettlementQueryResponseBodyModuleItems {
	s.PrimaryId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetProcessorOaCode(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ProcessorOaCode = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetProductCount(v int32) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ProductCount = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetProductId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ProductId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetProductName(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ProductName = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetProjectCode(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ProjectCode = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetProjectName(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ProjectName = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetPromotionFee(v float64) *VasBillSettlementQueryResponseBodyModuleItems {
	s.PromotionFee = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetPurchaseOrderId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.PurchaseOrderId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetRemark(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.Remark = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetSettleTypeDesc(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.SettleTypeDesc = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetSettlementFee(v float64) *VasBillSettlementQueryResponseBodyModuleItems {
	s.SettlementFee = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetSettlementGrantFee(v float64) *VasBillSettlementQueryResponseBodyModuleItems {
	s.SettlementGrantFee = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetSettlementTime(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.SettlementTime = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetSettlementType(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.SettlementType = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetSpecification(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.Specification = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetStatus(v int32) *VasBillSettlementQueryResponseBodyModuleItems {
	s.Status = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetStatusDesc(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.StatusDesc = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetSubOrderId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.SubOrderId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetTaxRate(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.TaxRate = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetThirdInvoiceId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ThirdInvoiceId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetThirdItineraryId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.ThirdItineraryId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetTradeActionDesc(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.TradeActionDesc = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetTradeRemark(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.TradeRemark = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetTravelerId(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.TravelerId = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetTravelerJobNo(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.TravelerJobNo = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberType(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberType = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberTypeName(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberTypeName = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetTravelerName(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.TravelerName = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetVoucherType(v int32) *VasBillSettlementQueryResponseBodyModuleItems {
	s.VoucherType = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) SetVoucherTypeDesc(v string) *VasBillSettlementQueryResponseBodyModuleItems {
	s.VoucherTypeDesc = &v
	return s
}

func (s *VasBillSettlementQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}
