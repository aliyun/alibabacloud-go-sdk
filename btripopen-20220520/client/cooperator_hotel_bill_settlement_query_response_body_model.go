// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorHotelBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CooperatorHotelBillSettlementQueryResponseBody
	GetCode() *string
	SetMessage(v string) *CooperatorHotelBillSettlementQueryResponseBody
	GetMessage() *string
	SetModule(v *CooperatorHotelBillSettlementQueryResponseBodyModule) *CooperatorHotelBillSettlementQueryResponseBody
	GetModule() *CooperatorHotelBillSettlementQueryResponseBodyModule
	SetRequestId(v string) *CooperatorHotelBillSettlementQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CooperatorHotelBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CooperatorHotelBillSettlementQueryResponseBody
	GetTraceId() *string
}

type CooperatorHotelBillSettlementQueryResponseBody struct {
	// example:
	//
	// 0
	Code    *string                                               `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                               `json:"message,omitempty" xml:"message,omitempty"`
	Module  *CooperatorHotelBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// 407543AF-2BD9-5890-BD92-9D1AB7218B27
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 21041ce316577904808056433edbb2
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CooperatorHotelBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CooperatorHotelBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) GetModule() *CooperatorHotelBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) SetCode(v string) *CooperatorHotelBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) SetMessage(v string) *CooperatorHotelBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) SetModule(v *CooperatorHotelBillSettlementQueryResponseBodyModule) *CooperatorHotelBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) SetRequestId(v string) *CooperatorHotelBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) SetSuccess(v bool) *CooperatorHotelBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) SetTraceId(v string) *CooperatorHotelBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CooperatorHotelBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 2
	Category *int32 `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// open12ilgngll7us7v10Bm5UlMg700
	CorpId *string                                                      `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Items  []*CooperatorHotelBillSettlementQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	// 30
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s CooperatorHotelBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CooperatorHotelBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) GetItems() []*CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	return s.Items
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) SetCategory(v int32) *CooperatorHotelBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) SetCorpId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) SetItems(v []*CooperatorHotelBillSettlementQueryResponseBodyModuleItems) *CooperatorHotelBillSettlementQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *CooperatorHotelBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *CooperatorHotelBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) SetScrollId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) SetTotalSize(v int64) *CooperatorHotelBillSettlementQueryResponseBodyModule {
	s.TotalSize = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type CooperatorHotelBillSettlementQueryResponseBodyModuleItems struct {
	AdjustTime *string `json:"adjust_time,omitempty" xml:"adjust_time,omitempty"`
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
	// sdasdas123324
	ApplyId       *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ApproverEmail *string `json:"approver_email,omitempty" xml:"approver_email,omitempty"`
	ApproverId    *string `json:"approver_id,omitempty" xml:"approver_id,omitempty"`
	ApproverName  *string `json:"approver_name,omitempty" xml:"approver_name,omitempty"`
	// example:
	//
	// 1.0
	AverageNights *float64 `json:"average_nights,omitempty" xml:"average_nights,omitempty"`
	BaseLocation  *string  `json:"base_location,omitempty" xml:"base_location,omitempty"`
	// example:
	//
	// 2023-01-01 00:00:00
	BillRecordTime *string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BookChannel    *string `json:"book_channel,omitempty" xml:"book_channel,omitempty"`
	BookMode       *string `json:"book_mode,omitempty" xml:"book_mode,omitempty"`
	BookReason     *string `json:"book_reason,omitempty" xml:"book_reason,omitempty"`
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
	BookerJobNo        *string `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName         *string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	BrandGroup         *string `json:"brand_group,omitempty" xml:"brand_group,omitempty"`
	BrandName          *string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	BusinessTripResult *string `json:"business_trip_result,omitempty" xml:"business_trip_result,omitempty"`
	CapitalDirection   *string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment  *string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CategoryDesc       *string `json:"category_desc,omitempty" xml:"category_desc,omitempty"`
	// example:
	//
	// 2024-02-13
	CheckInDate *string `json:"check_in_date,omitempty" xml:"check_in_date,omitempty"`
	// example:
	//
	// 2024-02-15
	CheckoutDate *string `json:"checkout_date,omitempty" xml:"checkout_date,omitempty"`
	City         *string `json:"city,omitempty" xml:"city,omitempty"`
	// example:
	//
	// 330100
	CityCode   *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	CityCounty *string `json:"city_county,omitempty" xml:"city_county,omitempty"`
	// example:
	//
	// 330100
	CityCountyCode *int32 `json:"city_county_code,omitempty" xml:"city_county_code,omitempty"`
	// example:
	//
	// IN240102113438277278
	CooperatorBillCode *string `json:"cooperator_bill_code,omitempty" xml:"cooperator_bill_code,omitempty"`
	CooperatorName     *string `json:"cooperator_name,omitempty" xml:"cooperator_name,omitempty"`
	// example:
	//
	// HO20240125162800280928
	CooperatorOrderId *string `json:"cooperator_order_id,omitempty" xml:"cooperator_order_id,omitempty"`
	// example:
	//
	// 0.1
	CorpRefundFee *float64 `json:"corp_refund_fee,omitempty" xml:"corp_refund_fee,omitempty"`
	// example:
	//
	// 100.0
	CorpTotalFee *float64 `json:"corp_total_fee,omitempty" xml:"corp_total_fee,omitempty"`
	CostCenter   *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// example:
	//
	// 123
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	CostDepartment   *string `json:"cost_department,omitempty" xml:"cost_department,omitempty"`
	CustomContent    *string `json:"custom_content,omitempty" xml:"custom_content,omitempty"`
	Department       *string `json:"department,omitempty" xml:"department,omitempty"`
	// example:
	//
	// 123
	DepartmentId *string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	ExceedReason *string `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	FeeType      *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc  *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	// example:
	//
	// 10.0
	Fees *float64 `json:"fees,omitempty" xml:"fees,omitempty"`
	// example:
	//
	// 1.0
	Fines         *float64 `json:"fines,omitempty" xml:"fines,omitempty"`
	ForeignersTag *string  `json:"foreigners_tag,omitempty" xml:"foreigners_tag,omitempty"`
	// example:
	//
	// 10.0
	FuPointFee *float64 `json:"fu_point_fee,omitempty" xml:"fu_point_fee,omitempty"`
	HotelName  *string  `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// example:
	//
	// 1
	Index              *string `json:"index,omitempty" xml:"index,omitempty"`
	InvoiceTitle       *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	IsEarlyDeparture   *string `json:"is_early_departure,omitempty" xml:"is_early_departure,omitempty"`
	IsNegotiation      *string `json:"is_negotiation,omitempty" xml:"is_negotiation,omitempty"`
	IsShareStr         *string `json:"is_share_str,omitempty" xml:"is_share_str,omitempty"`
	Location           *string `json:"location,omitempty" xml:"location,omitempty"`
	MappingCompanyCode *string `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	// example:
	//
	// 2
	Nights *int32 `json:"nights,omitempty" xml:"nights,omitempty"`
	// example:
	//
	// 3137168772101111000
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 100.0
	OrderPrice      *float64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	OrderStatusDesc *string  `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	OrderType       *string  `json:"order_type,omitempty" xml:"order_type,omitempty"`
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
	// 0.1
	PersonRefundFee *float64 `json:"person_refund_fee,omitempty" xml:"person_refund_fee,omitempty"`
	// example:
	//
	// 10.0
	PersonSettlePrice *float64 `json:"person_settle_price,omitempty" xml:"person_settle_price,omitempty"`
	Position          *string  `json:"position,omitempty" xml:"position,omitempty"`
	PositionLevel     *string  `json:"position_level,omitempty" xml:"position_level,omitempty"`
	// example:
	//
	// 72328485
	PrimaryId       *int64  `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProcessorOaCode *string `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// acs
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// example:
	//
	// 0.0
	PromotionFee *float64 `json:"promotion_fee,omitempty" xml:"promotion_fee,omitempty"`
	Remark       *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 10
	ReserveRule *int32  `json:"reserve_rule,omitempty" xml:"reserve_rule,omitempty"`
	RoomNo      *string `json:"room_no,omitempty" xml:"room_no,omitempty"`
	// example:
	//
	// 1
	RoomNumber *int32 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// example:
	//
	// 105.0
	RoomPrice *float64 `json:"room_price,omitempty" xml:"room_price,omitempty"`
	RoomType  *string  `json:"room_type,omitempty" xml:"room_type,omitempty"`
	// example:
	//
	// 5.0
	ServiceFee     *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettleTypeDesc *string  `json:"settle_type_desc,omitempty" xml:"settle_type_desc,omitempty"`
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
	Sio            *string `json:"sio,omitempty" xml:"sio,omitempty"`
	Star           *string `json:"star,omitempty" xml:"star,omitempty"`
	// example:
	//
	// 2
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// 166564408
	SubOrderId *string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// example:
	//
	// 9%
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// example:
	//
	// 1
	ThirdInvoiceId *string `json:"third_invoice_id,omitempty" xml:"third_invoice_id,omitempty"`
	// example:
	//
	// 123
	ThirdItineraryId *string `json:"third_itinerary_id,omitempty" xml:"third_itinerary_id,omitempty"`
	// example:
	//
	// 2
	TotalNights     *int32  `json:"total_nights,omitempty" xml:"total_nights,omitempty"`
	TradeActionDesc *string `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
	TravelerEmail   *string `json:"traveler_email,omitempty" xml:"traveler_email,omitempty"`
	// example:
	//
	// 123
	TravelerId *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// example:
	//
	// zs123
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

func (s CooperatorHotelBillSettlementQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetApplyArrCityCode() *string {
	return s.ApplyArrCityCode
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetApplyArrCityName() *string {
	return s.ApplyArrCityName
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetApplyDepCityCode() *string {
	return s.ApplyDepCityCode
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetApplyDepCityName() *string {
	return s.ApplyDepCityName
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetApplyId() *string {
	return s.ApplyId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetApproverEmail() *string {
	return s.ApproverEmail
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetApproverId() *string {
	return s.ApproverId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetApproverName() *string {
	return s.ApproverName
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetAverageNights() *float64 {
	return s.AverageNights
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBaseLocation() *string {
	return s.BaseLocation
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBookChannel() *string {
	return s.BookChannel
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBookMode() *string {
	return s.BookMode
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBookReason() *string {
	return s.BookReason
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBookTime() *string {
	return s.BookTime
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBookerId() *string {
	return s.BookerId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBookerName() *string {
	return s.BookerName
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBrandGroup() *string {
	return s.BrandGroup
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBrandName() *string {
	return s.BrandName
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetBusinessTripResult() *string {
	return s.BusinessTripResult
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCheckoutDate() *string {
	return s.CheckoutDate
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCity() *string {
	return s.City
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCityCode() *string {
	return s.CityCode
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCityCounty() *string {
	return s.CityCounty
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCityCountyCode() *int32 {
	return s.CityCountyCode
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCooperatorBillCode() *string {
	return s.CooperatorBillCode
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCooperatorName() *string {
	return s.CooperatorName
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCooperatorOrderId() *string {
	return s.CooperatorOrderId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCorpRefundFee() *float64 {
	return s.CorpRefundFee
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCorpTotalFee() *float64 {
	return s.CorpTotalFee
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCostCenter() *string {
	return s.CostCenter
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetCustomContent() *string {
	return s.CustomContent
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetDepartment() *string {
	return s.Department
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetExceedReason() *string {
	return s.ExceedReason
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetFeeType() *string {
	return s.FeeType
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetFees() *float64 {
	return s.Fees
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetFines() *float64 {
	return s.Fines
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetForeignersTag() *string {
	return s.ForeignersTag
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetFuPointFee() *float64 {
	return s.FuPointFee
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetHotelName() *string {
	return s.HotelName
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetIndex() *string {
	return s.Index
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetIsEarlyDeparture() *string {
	return s.IsEarlyDeparture
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetIsNegotiation() *string {
	return s.IsNegotiation
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetIsShareStr() *string {
	return s.IsShareStr
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetLocation() *string {
	return s.Location
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetNights() *int32 {
	return s.Nights
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetOrderId() *string {
	return s.OrderId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetOrderPrice() *float64 {
	return s.OrderPrice
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetOrderType() *string {
	return s.OrderType
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetOverApplyId() *string {
	return s.OverApplyId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetPersonRefundFee() *float64 {
	return s.PersonRefundFee
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetPersonSettlePrice() *float64 {
	return s.PersonSettlePrice
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetPosition() *string {
	return s.Position
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetProjectName() *string {
	return s.ProjectName
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetPromotionFee() *float64 {
	return s.PromotionFee
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetRemark() *string {
	return s.Remark
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetReserveRule() *int32 {
	return s.ReserveRule
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetRoomNo() *string {
	return s.RoomNo
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetRoomNumber() *int32 {
	return s.RoomNumber
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetRoomPrice() *float64 {
	return s.RoomPrice
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetRoomType() *string {
	return s.RoomType
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetSettlementGrantFee() *float64 {
	return s.SettlementGrantFee
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetSettlementType() *string {
	return s.SettlementType
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetSio() *string {
	return s.Sio
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetStar() *string {
	return s.Star
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetStatus() *int32 {
	return s.Status
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetTaxRate() *string {
	return s.TaxRate
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetThirdInvoiceId() *string {
	return s.ThirdInvoiceId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetThirdItineraryId() *string {
	return s.ThirdItineraryId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetTotalNights() *int32 {
	return s.TotalNights
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetTravelerEmail() *string {
	return s.TravelerEmail
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetTravelerId() *string {
	return s.TravelerId
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberTypeName() *string {
	return s.TravelerMemberTypeName
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetTravelerName() *string {
	return s.TravelerName
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetAdjustTime(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.AdjustTime = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetAlipayTradeNo(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.AlipayTradeNo = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetApplyArrCityCode(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ApplyArrCityCode = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetApplyArrCityName(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ApplyArrCityName = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetApplyDepCityCode(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ApplyDepCityCode = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetApplyDepCityName(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ApplyDepCityName = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetApplyExtendField(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ApplyExtendField = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetApplyId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ApplyId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetApproverEmail(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ApproverEmail = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetApproverId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ApproverId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetApproverName(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ApproverName = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetAverageNights(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.AverageNights = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBaseLocation(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BaseLocation = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBillRecordTime(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BillRecordTime = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBookChannel(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BookChannel = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBookMode(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BookMode = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBookReason(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BookReason = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBookTime(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BookTime = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBookerId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BookerId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBookerJobNo(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BookerJobNo = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBookerName(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BookerName = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBrandGroup(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BrandGroup = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBrandName(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BrandName = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetBusinessTripResult(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.BusinessTripResult = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCapitalDirection(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CapitalDirection = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCascadeDepartment(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CascadeDepartment = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCategoryDesc(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CategoryDesc = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCheckInDate(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CheckInDate = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCheckoutDate(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CheckoutDate = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCity(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.City = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCityCode(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CityCode = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCityCounty(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CityCounty = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCityCountyCode(v int32) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CityCountyCode = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCooperatorBillCode(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CooperatorBillCode = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCooperatorName(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CooperatorName = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCooperatorOrderId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CooperatorOrderId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCorpRefundFee(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CorpRefundFee = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCorpTotalFee(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CorpTotalFee = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCostCenter(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CostCenter = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCostCenterNumber(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CostCenterNumber = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCostDepartment(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CostDepartment = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetCustomContent(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.CustomContent = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetDepartment(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.Department = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetDepartmentId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.DepartmentId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetExceedReason(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ExceedReason = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetFeeType(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.FeeType = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetFeeTypeDesc(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.FeeTypeDesc = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetFees(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.Fees = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetFines(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.Fines = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetForeignersTag(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ForeignersTag = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetFuPointFee(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.FuPointFee = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetHotelName(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.HotelName = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetIndex(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.Index = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetInvoiceTitle(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.InvoiceTitle = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetIsEarlyDeparture(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.IsEarlyDeparture = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetIsNegotiation(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.IsNegotiation = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetIsShareStr(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.IsShareStr = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetLocation(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.Location = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetMappingCompanyCode(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.MappingCompanyCode = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetNights(v int32) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.Nights = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetOrderId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.OrderId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetOrderPrice(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.OrderPrice = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetOrderStatusDesc(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.OrderStatusDesc = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetOrderType(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.OrderType = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetOverApplyId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.OverApplyId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetPaymentDepartmentId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.PaymentDepartmentId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetPaymentDepartmentName(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.PaymentDepartmentName = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetPersonRefundFee(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.PersonRefundFee = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetPersonSettlePrice(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.PersonSettlePrice = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetPosition(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.Position = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetPositionLevel(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.PositionLevel = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetPrimaryId(v int64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.PrimaryId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetProcessorOaCode(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ProcessorOaCode = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetProjectCode(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ProjectCode = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetProjectName(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ProjectName = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetPromotionFee(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.PromotionFee = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetRemark(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.Remark = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetReserveRule(v int32) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ReserveRule = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetRoomNo(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.RoomNo = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetRoomNumber(v int32) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.RoomNumber = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetRoomPrice(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.RoomPrice = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetRoomType(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.RoomType = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetServiceFee(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ServiceFee = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetSettleTypeDesc(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.SettleTypeDesc = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetSettlementFee(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.SettlementFee = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetSettlementGrantFee(v float64) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.SettlementGrantFee = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetSettlementTime(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.SettlementTime = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetSettlementType(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.SettlementType = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetSio(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.Sio = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetStar(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.Star = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetStatus(v int32) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.Status = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetStatusDesc(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.StatusDesc = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetSubOrderId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.SubOrderId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetTaxRate(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.TaxRate = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetThirdInvoiceId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ThirdInvoiceId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetThirdItineraryId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.ThirdItineraryId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetTotalNights(v int32) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.TotalNights = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetTradeActionDesc(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.TradeActionDesc = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetTravelerEmail(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.TravelerEmail = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetTravelerId(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.TravelerId = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetTravelerJobNo(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.TravelerJobNo = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberType(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberType = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberTypeName(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberTypeName = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetTravelerName(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.TravelerName = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetVoucherType(v int32) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.VoucherType = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) SetVoucherTypeDesc(v string) *CooperatorHotelBillSettlementQueryResponseBodyModuleItems {
	s.VoucherTypeDesc = &v
	return s
}

func (s *CooperatorHotelBillSettlementQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}
