// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotelBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *HotelBillSettlementQueryResponseBody
	GetCode() *string
	SetMessage(v string) *HotelBillSettlementQueryResponseBody
	GetMessage() *string
	SetModule(v *HotelBillSettlementQueryResponseBodyModule) *HotelBillSettlementQueryResponseBody
	GetModule() *HotelBillSettlementQueryResponseBodyModule
	SetRequestId(v string) *HotelBillSettlementQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *HotelBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *HotelBillSettlementQueryResponseBody
	GetTraceId() *string
}

type HotelBillSettlementQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                     `json:"message,omitempty" xml:"message,omitempty"`
	Module  *HotelBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s HotelBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s HotelBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *HotelBillSettlementQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *HotelBillSettlementQueryResponseBody) GetModule() *HotelBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *HotelBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *HotelBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *HotelBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *HotelBillSettlementQueryResponseBody) SetCode(v string) *HotelBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBody) SetMessage(v string) *HotelBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBody) SetModule(v *HotelBillSettlementQueryResponseBodyModule) *HotelBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *HotelBillSettlementQueryResponseBody) SetRequestId(v string) *HotelBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBody) SetSuccess(v bool) *HotelBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBody) SetTraceId(v string) *HotelBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type HotelBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 2
	Category *int32                                                `json:"category,omitempty" xml:"category,omitempty"`
	CorpId   *string                                               `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	DataList []*HotelBillSettlementQueryResponseBodyModuleDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-10-14
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2021-10-13
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	ScrollId    *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// 1402
	TotalNum *int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

func (s HotelBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s HotelBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *HotelBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *HotelBillSettlementQueryResponseBodyModule) GetDataList() []*HotelBillSettlementQueryResponseBodyModuleDataList {
	return s.DataList
}

func (s *HotelBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *HotelBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *HotelBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *HotelBillSettlementQueryResponseBodyModule) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetCategory(v int32) *HotelBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetCorpId(v string) *HotelBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetDataList(v []*HotelBillSettlementQueryResponseBodyModuleDataList) *HotelBillSettlementQueryResponseBodyModule {
	s.DataList = v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *HotelBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *HotelBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetScrollId(v string) *HotelBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *HotelBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type HotelBillSettlementQueryResponseBodyModuleDataList struct {
	AdjustTime            *string  `json:"adjust_time,omitempty" xml:"adjust_time,omitempty"`
	AgreementPromotionFee *float64 `json:"agreement_promotion_fee,omitempty" xml:"agreement_promotion_fee,omitempty"`
	// example:
	//
	// 234432432
	AlipayTradeNo    *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	ApplyArrCityCode *string `json:"apply_arr_city_code,omitempty" xml:"apply_arr_city_code,omitempty"`
	ApplyArrCityName *string `json:"apply_arr_city_name,omitempty" xml:"apply_arr_city_name,omitempty"`
	ApplyDepCityCode *string `json:"apply_dep_city_code,omitempty" xml:"apply_dep_city_code,omitempty"`
	ApplyDepCityName *string `json:"apply_dep_city_name,omitempty" xml:"apply_dep_city_name,omitempty"`
	// 审批扩展自定义字段
	ApplyExtendField *string `json:"apply_extend_field,omitempty" xml:"apply_extend_field,omitempty"`
	// example:
	//
	// 103208648
	ApplyId       *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	AverageNights *string `json:"average_nights,omitempty" xml:"average_nights,omitempty"`
	BaseLocation  *string `json:"base_location,omitempty" xml:"base_location,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	BillRecordTime *string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BillingEntity  *string `json:"billing_entity,omitempty" xml:"billing_entity,omitempty"`
	BookChannel    *string `json:"book_channel,omitempty" xml:"book_channel,omitempty"`
	BookMode       *string `json:"book_mode,omitempty" xml:"book_mode,omitempty"`
	BookReason     *string `json:"book_reason,omitempty" xml:"book_reason,omitempty"`
	// example:
	//
	// 2021-10-12 23:58:48
	BookTime             *string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	BookerId             *string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	BookerJobNo          *string `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName           *string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	BrandGroup           *string `json:"brand_group,omitempty" xml:"brand_group,omitempty"`
	BrandName            *string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	BusinessTripResult   *string `json:"business_trip_result,omitempty" xml:"business_trip_result,omitempty"`
	CancelOrModifyReason *string `json:"cancel_or_modify_reason,omitempty" xml:"cancel_or_modify_reason,omitempty"`
	CancelOrModifyScene  *string `json:"cancel_or_modify_scene,omitempty" xml:"cancel_or_modify_scene,omitempty"`
	// example:
	//
	// 1
	CapitalDirection  *string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment *string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CategoryDesc      *string `json:"category_desc,omitempty" xml:"category_desc,omitempty"`
	// example:
	//
	// 2021-10-14 00:00:00
	CheckInDate *string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// example:
	//
	// 2021-10-16 00:00:00
	CheckoutDate *string `json:"checkout_date,omitempty" xml:"checkout_date,omitempty"`
	City         *string `json:"city,omitempty" xml:"city,omitempty"`
	// example:
	//
	// 110100
	CityCode       *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityCounty     *string `json:"city_county,omitempty" xml:"city_county,omitempty"`
	CityCountyCode *string `json:"city_county_code,omitempty" xml:"city_county_code,omitempty"`
	// example:
	//
	// 12
	CorpRefundFee *float64 `json:"corp_refund_fee,omitempty" xml:"corp_refund_fee,omitempty"`
	// example:
	//
	// 1000
	CorpTotalFee *float64 `json:"corp_total_fee,omitempty" xml:"corp_total_fee,omitempty"`
	CostCenter   *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// example:
	//
	// 8b7f3cd-24324-097
	CostCenterNumber *string  `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	CostDepartment   *string  `json:"cost_department,omitempty" xml:"cost_department,omitempty"`
	CustomContent    *string  `json:"custom_content,omitempty" xml:"custom_content,omitempty"`
	DeductibleTax    *float64 `json:"deductible_tax,omitempty" xml:"deductible_tax,omitempty"`
	Department       *string  `json:"department,omitempty" xml:"department,omitempty"`
	DepartmentId     *string  `json:"department_id,omitempty" xml:"department_id,omitempty"`
	ExceedReason     *string  `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	// example:
	//
	// 20101
	FeeType     *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	// example:
	//
	// 0
	Fees  *float64 `json:"fees,omitempty" xml:"fees,omitempty"`
	Fines *float64 `json:"fines,omitempty" xml:"fines,omitempty"`
	// example:
	//
	// 12
	FuPointFee *float64 `json:"fu_point_fee,omitempty" xml:"fu_point_fee,omitempty"`
	HotelName  *string  `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// example:
	//
	// 5038018
	Index                *string  `json:"index,omitempty" xml:"index,omitempty"`
	InsOrderId           *string  `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	InsuranceNumber      *string  `json:"insurance_number,omitempty" xml:"insurance_number,omitempty"`
	InsurancePrice       *float64 `json:"insurance_price,omitempty" xml:"insurance_price,omitempty"`
	InsuranceProductName *string  `json:"insurance_product_name,omitempty" xml:"insurance_product_name,omitempty"`
	InvoiceTitle         *string  `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	IsEarlyDeparture     *string  `json:"is_early_departure,omitempty" xml:"is_early_departure,omitempty"`
	IsNegotiation        *string  `json:"is_negotiation,omitempty" xml:"is_negotiation,omitempty"`
	IsShareStr           *string  `json:"is_share_str,omitempty" xml:"is_share_str,omitempty"`
	MappingCompanyCode   *string  `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	// example:
	//
	// 2
	Nights                 *int32  `json:"nights,omitempty" xml:"nights,omitempty"`
	NoAdvanceBookingReason *string `json:"no_advance_booking_reason,omitempty" xml:"no_advance_booking_reason,omitempty"`
	// example:
	//
	// 223423423432422
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1088.96
	OrderPrice          *float64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	OrderType           *string  `json:"order_type,omitempty" xml:"order_type,omitempty"`
	OriginalReserveRule *string  `json:"original_reserve_rule,omitempty" xml:"original_reserve_rule,omitempty"`
	// example:
	//
	// 4234324
	OverApplyId           *string `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	PaymentDepartmentId   *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// 34
	PersonRefundFee *float64 `json:"person_refund_fee,omitempty" xml:"person_refund_fee,omitempty"`
	// example:
	//
	// 88.96
	PersonSettlePrice *float64 `json:"person_settle_price,omitempty" xml:"person_settle_price,omitempty"`
	Position          *string  `json:"position,omitempty" xml:"position,omitempty"`
	PositionLevel     *string  `json:"position_level,omitempty" xml:"position_level,omitempty"`
	// example:
	//
	// 5038018
	PrimaryId       *int64  `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProcessorOaCode *string `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// 223423432
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// example:
	//
	// 113.02
	PromotionFee           *float64 `json:"promotion_fee,omitempty" xml:"promotion_fee,omitempty"`
	RecoverMoneyReceiptAmt *float64 `json:"recover_money_receipt_amt,omitempty" xml:"recover_money_receipt_amt,omitempty"`
	Remark                 *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	ReserveRule            *string  `json:"reserve_rule,omitempty" xml:"reserve_rule,omitempty"`
	RoomNo                 *string  `json:"room_no,omitempty" xml:"room_no,omitempty"`
	// example:
	//
	// 1
	RoomNumber *int32 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// example:
	//
	// 1201.98
	RoomPrice *float64 `json:"room_price,omitempty" xml:"room_price,omitempty"`
	RoomType  *string  `json:"room_type,omitempty" xml:"room_type,omitempty"`
	SceneId   *string  `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	SceneName *string  `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	// example:
	//
	// 0
	ServiceFee     *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettleTypeDesc *string  `json:"settle_type_desc,omitempty" xml:"settle_type_desc,omitempty"`
	// example:
	//
	// 1000
	SettlementFee *float64 `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	// example:
	//
	// 5.68
	SettlementGrantFee *float64 `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	// example:
	//
	// 2021-10-12 23:58:56
	SettlementTime *string `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	// example:
	//
	// 4
	SettlementType *string `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	Star           *string `json:"star,omitempty" xml:"star,omitempty"`
	// example:
	//
	// 1
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// 税率
	//
	// example:
	//
	// 6%
	TaxRate          *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	ThirdItineraryId *string `json:"third_itinerary_id,omitempty" xml:"third_itinerary_id,omitempty"`
	// example:
	//
	// 2
	TotalNights     *int32  `json:"total_nights,omitempty" xml:"total_nights,omitempty"`
	TradeActionDesc *string `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
	TravelerId      *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// example:
	//
	// 326246
	TravelerJobNo             *string `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerMemberType        *string `json:"traveler_member_type,omitempty" xml:"traveler_member_type,omitempty"`
	TravelerMemberTypeName    *string `json:"traveler_member_type_name,omitempty" xml:"traveler_member_type_name,omitempty"`
	TravelerName              *string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	UnbookedLowestPriceReason *string `json:"unbooked_lowest_price_reason,omitempty" xml:"unbooked_lowest_price_reason,omitempty"`
	// example:
	//
	// 11
	VoucherType     *int32  `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	VoucherTypeDesc *string `json:"voucher_type_desc,omitempty" xml:"voucher_type_desc,omitempty"`
}

func (s HotelBillSettlementQueryResponseBodyModuleDataList) String() string {
	return dara.Prettify(s)
}

func (s HotelBillSettlementQueryResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetAgreementPromotionFee() *float64 {
	return s.AgreementPromotionFee
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityCode() *string {
	return s.ApplyArrCityCode
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityName() *string {
	return s.ApplyArrCityName
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityCode() *string {
	return s.ApplyDepCityCode
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityName() *string {
	return s.ApplyDepCityName
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetApplyId() *string {
	return s.ApplyId
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetAverageNights() *string {
	return s.AverageNights
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBaseLocation() *string {
	return s.BaseLocation
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBillingEntity() *string {
	return s.BillingEntity
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBookChannel() *string {
	return s.BookChannel
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBookMode() *string {
	return s.BookMode
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBookReason() *string {
	return s.BookReason
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBookTime() *string {
	return s.BookTime
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBookerId() *string {
	return s.BookerId
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBookerName() *string {
	return s.BookerName
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBrandGroup() *string {
	return s.BrandGroup
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBrandName() *string {
	return s.BrandName
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetBusinessTripResult() *string {
	return s.BusinessTripResult
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCancelOrModifyReason() *string {
	return s.CancelOrModifyReason
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCancelOrModifyScene() *string {
	return s.CancelOrModifyScene
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCheckoutDate() *string {
	return s.CheckoutDate
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCity() *string {
	return s.City
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCityCode() *string {
	return s.CityCode
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCityCounty() *string {
	return s.CityCounty
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCityCountyCode() *string {
	return s.CityCountyCode
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCorpRefundFee() *float64 {
	return s.CorpRefundFee
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCorpTotalFee() *float64 {
	return s.CorpTotalFee
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCostCenter() *string {
	return s.CostCenter
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetCustomContent() *string {
	return s.CustomContent
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetDeductibleTax() *float64 {
	return s.DeductibleTax
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetDepartment() *string {
	return s.Department
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetExceedReason() *string {
	return s.ExceedReason
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetFeeType() *string {
	return s.FeeType
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetFees() *float64 {
	return s.Fees
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetFines() *float64 {
	return s.Fines
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetFuPointFee() *float64 {
	return s.FuPointFee
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetHotelName() *string {
	return s.HotelName
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetIndex() *string {
	return s.Index
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetInsuranceNumber() *string {
	return s.InsuranceNumber
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetInsurancePrice() *float64 {
	return s.InsurancePrice
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetInsuranceProductName() *string {
	return s.InsuranceProductName
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetIsEarlyDeparture() *string {
	return s.IsEarlyDeparture
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetIsNegotiation() *string {
	return s.IsNegotiation
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetIsShareStr() *string {
	return s.IsShareStr
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetNights() *int32 {
	return s.Nights
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetNoAdvanceBookingReason() *string {
	return s.NoAdvanceBookingReason
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetOrderId() *string {
	return s.OrderId
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetOrderPrice() *float64 {
	return s.OrderPrice
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetOrderType() *string {
	return s.OrderType
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetOriginalReserveRule() *string {
	return s.OriginalReserveRule
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetOverApplyId() *string {
	return s.OverApplyId
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetPersonRefundFee() *float64 {
	return s.PersonRefundFee
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetPersonSettlePrice() *float64 {
	return s.PersonSettlePrice
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetPosition() *string {
	return s.Position
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetProjectName() *string {
	return s.ProjectName
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetPromotionFee() *float64 {
	return s.PromotionFee
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetRecoverMoneyReceiptAmt() *float64 {
	return s.RecoverMoneyReceiptAmt
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetRemark() *string {
	return s.Remark
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetReserveRule() *string {
	return s.ReserveRule
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetRoomNo() *string {
	return s.RoomNo
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetRoomNumber() *int32 {
	return s.RoomNumber
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetRoomPrice() *float64 {
	return s.RoomPrice
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetRoomType() *string {
	return s.RoomType
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetSceneId() *string {
	return s.SceneId
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetSceneName() *string {
	return s.SceneName
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetSettlementGrantFee() *float64 {
	return s.SettlementGrantFee
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetSettlementType() *string {
	return s.SettlementType
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetStar() *string {
	return s.Star
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetStatus() *int32 {
	return s.Status
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetTaxRate() *string {
	return s.TaxRate
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetThirdItineraryId() *string {
	return s.ThirdItineraryId
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetTotalNights() *int32 {
	return s.TotalNights
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetTravelerId() *string {
	return s.TravelerId
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetTravelerMemberTypeName() *string {
	return s.TravelerMemberTypeName
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetTravelerName() *string {
	return s.TravelerName
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetUnbookedLowestPriceReason() *string {
	return s.UnbookedLowestPriceReason
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetAdjustTime(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.AdjustTime = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetAgreementPromotionFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.AgreementPromotionFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetAlipayTradeNo(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityCode(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityCode = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityCode(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityCode = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetApplyExtendField(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyExtendField = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetApplyId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetAverageNights(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.AverageNights = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBaseLocation(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BaseLocation = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBillRecordTime(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BillRecordTime = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBillingEntity(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BillingEntity = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBookChannel(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookChannel = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBookMode(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookMode = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBookReason(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookReason = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBookTime(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBookerId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBookerJobNo(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBookerName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBrandGroup(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BrandGroup = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBrandName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BrandName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetBusinessTripResult(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.BusinessTripResult = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCancelOrModifyReason(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CancelOrModifyReason = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCancelOrModifyScene(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CancelOrModifyScene = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCapitalDirection(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCascadeDepartment(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCategoryDesc(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CategoryDesc = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCheckInDate(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CheckInDate = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCheckoutDate(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CheckoutDate = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCity(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.City = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCityCode(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CityCode = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCityCounty(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CityCounty = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCityCountyCode(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CityCountyCode = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCorpRefundFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CorpRefundFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCorpTotalFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CorpTotalFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCostCenter(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCostCenterNumber(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCostDepartment(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CostDepartment = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetCustomContent(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.CustomContent = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetDeductibleTax(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.DeductibleTax = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetDepartment(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetDepartmentId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetExceedReason(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ExceedReason = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetFeeType(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetFeeTypeDesc(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.FeeTypeDesc = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetFees(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Fees = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetFines(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Fines = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetFuPointFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.FuPointFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetHotelName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.HotelName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetIndex(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetInsOrderId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.InsOrderId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetInsuranceNumber(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.InsuranceNumber = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetInsurancePrice(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.InsurancePrice = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetInsuranceProductName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.InsuranceProductName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetInvoiceTitle(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetIsEarlyDeparture(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.IsEarlyDeparture = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetIsNegotiation(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.IsNegotiation = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetIsShareStr(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.IsShareStr = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetMappingCompanyCode(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.MappingCompanyCode = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetNights(v int32) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Nights = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetNoAdvanceBookingReason(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.NoAdvanceBookingReason = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetOrderId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetOrderPrice(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.OrderPrice = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetOrderType(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.OrderType = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetOriginalReserveRule(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.OriginalReserveRule = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetOverApplyId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPersonRefundFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.PersonRefundFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPersonSettlePrice(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.PersonSettlePrice = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPosition(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Position = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPositionLevel(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.PositionLevel = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPrimaryId(v int64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetProcessorOaCode(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ProcessorOaCode = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetProjectCode(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetProjectName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetPromotionFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.PromotionFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetRecoverMoneyReceiptAmt(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.RecoverMoneyReceiptAmt = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetRemark(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Remark = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetReserveRule(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ReserveRule = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetRoomNo(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.RoomNo = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetRoomNumber(v int32) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.RoomNumber = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetRoomPrice(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.RoomPrice = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetRoomType(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.RoomType = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetSceneId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.SceneId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetSceneName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.SceneName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetServiceFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetSettleTypeDesc(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettleTypeDesc = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementGrantFee(v float64) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementGrantFee = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementTime(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementType(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetStar(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Star = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetStatus(v int32) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetStatusDesc(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.StatusDesc = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTaxRate(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TaxRate = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetThirdItineraryId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.ThirdItineraryId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTotalNights(v int32) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TotalNights = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTradeActionDesc(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TradeActionDesc = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerId(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerJobNo(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerMemberType(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerMemberType = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerMemberTypeName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerMemberTypeName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerName(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetUnbookedLowestPriceReason(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.UnbookedLowestPriceReason = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetVoucherType(v int32) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) SetVoucherTypeDesc(v string) *HotelBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherTypeDesc = &v
	return s
}

func (s *HotelBillSettlementQueryResponseBodyModuleDataList) Validate() error {
	return dara.Validate(s)
}
