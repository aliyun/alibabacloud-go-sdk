// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeHotelBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IeHotelBillSettlementQueryResponseBody
	GetCode() *string
	SetMessage(v string) *IeHotelBillSettlementQueryResponseBody
	GetMessage() *string
	SetModule(v *IeHotelBillSettlementQueryResponseBodyModule) *IeHotelBillSettlementQueryResponseBody
	GetModule() *IeHotelBillSettlementQueryResponseBodyModule
	SetRequestId(v string) *IeHotelBillSettlementQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IeHotelBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IeHotelBillSettlementQueryResponseBody
	GetTraceId() *string
}

type IeHotelBillSettlementQueryResponseBody struct {
	// example:
	//
	// 200
	Code    *string `json:"code,omitempty" xml:"code,omitempty"`
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// module
	Module *IeHotelBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 210e877f16763560074236874d5268
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 2103a08a16861217249785276d5a87
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IeHotelBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IeHotelBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *IeHotelBillSettlementQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *IeHotelBillSettlementQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IeHotelBillSettlementQueryResponseBody) GetModule() *IeHotelBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *IeHotelBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IeHotelBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IeHotelBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IeHotelBillSettlementQueryResponseBody) SetCode(v string) *IeHotelBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBody) SetMessage(v string) *IeHotelBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBody) SetModule(v *IeHotelBillSettlementQueryResponseBodyModule) *IeHotelBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBody) SetRequestId(v string) *IeHotelBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBody) SetSuccess(v bool) *IeHotelBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBody) SetTraceId(v string) *IeHotelBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type IeHotelBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 12
	Category *int32 `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// corp1
	CorpId   *string                                                 `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	DataList []*IeHotelBillSettlementQueryResponseBodyModuleDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
	// example:
	//
	// 1012039195340093034
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 2022-11-02
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2022-11-01
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	ScrollId    *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// 30
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s IeHotelBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IeHotelBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) GetDataList() []*IeHotelBillSettlementQueryResponseBodyModuleDataList {
	return s.DataList
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) SetCategory(v int32) *IeHotelBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) SetCorpId(v string) *IeHotelBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) SetDataList(v []*IeHotelBillSettlementQueryResponseBodyModuleDataList) *IeHotelBillSettlementQueryResponseBodyModule {
	s.DataList = v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) SetOrderId(v string) *IeHotelBillSettlementQueryResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *IeHotelBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *IeHotelBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) SetScrollId(v string) *IeHotelBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) SetTotalSize(v int64) *IeHotelBillSettlementQueryResponseBodyModule {
	s.TotalSize = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IeHotelBillSettlementQueryResponseBodyModuleDataList struct {
	AdjustTime *string `json:"adjust_time,omitempty" xml:"adjust_time,omitempty"`
	// example:
	//
	// 2021123432260
	AlipayTradeNo *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// example:
	//
	// AUD
	AmountCurrency *string `json:"amount_currency,omitempty" xml:"amount_currency,omitempty"`
	// example:
	//
	// MDG
	ApplyArrCityCode *string `json:"apply_arr_city_code,omitempty" xml:"apply_arr_city_code,omitempty"`
	ApplyArrCityName *string `json:"apply_arr_city_name,omitempty" xml:"apply_arr_city_name,omitempty"`
	// example:
	//
	// HRB
	ApplyDepCityCode *string `json:"apply_dep_city_code,omitempty" xml:"apply_dep_city_code,omitempty"`
	ApplyDepCityName *string `json:"apply_dep_city_name,omitempty" xml:"apply_dep_city_name,omitempty"`
	ApplyExtendField *string `json:"apply_extend_field,omitempty" xml:"apply_extend_field,omitempty"`
	// example:
	//
	// 103189557
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
	// 2022-07-20T10:40Z
	BillRecordTime *string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BookMode       *string `json:"book_mode,omitempty" xml:"book_mode,omitempty"`
	BookReason     *string `json:"book_reason,omitempty" xml:"book_reason,omitempty"`
	// example:
	//
	// 2021-10-08 23:38:55
	BookTime *string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// example:
	//
	// al_xinuan.zsy
	BookerId *string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	// example:
	//
	// 70022164
	BookerJobNo *string `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName  *string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	BrandGroup  *string `json:"brand_group,omitempty" xml:"brand_group,omitempty"`
	// example:
	//
	// XXX
	BrandName *string `json:"brand_name,omitempty" xml:"brand_name,omitempty"`
	// example:
	//
	// 1345
	BusinessExpense    *int64  `json:"business_expense,omitempty" xml:"business_expense,omitempty"`
	BusinessTripResult *string `json:"business_trip_result,omitempty" xml:"business_trip_result,omitempty"`
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
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
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
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	CostDepartment   *string `json:"cost_department,omitempty" xml:"cost_department,omitempty"`
	Country          *string `json:"country,omitempty" xml:"country,omitempty"`
	// example:
	//
	// 1454567
	CountryCode   *string  `json:"country_code,omitempty" xml:"country_code,omitempty"`
	CustomContent *string  `json:"custom_content,omitempty" xml:"custom_content,omitempty"`
	DeductibleTax *float64 `json:"deductible_tax,omitempty" xml:"deductible_tax,omitempty"`
	Department    *string  `json:"department,omitempty" xml:"department,omitempty"`
	DepartmentId  *string  `json:"department_id,omitempty" xml:"department_id,omitempty"`
	ExceedReason  *string  `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	// example:
	//
	// 20101
	FeeType     *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	// example:
	//
	// 10.1
	Fines *float64 `json:"fines,omitempty" xml:"fines,omitempty"`
	// example:
	//
	// 345
	ForeignBusinessExpense *int64  `json:"foreign_business_expense,omitempty" xml:"foreign_business_expense,omitempty"`
	ForeignersTag          *string `json:"foreigners_tag,omitempty" xml:"foreigners_tag,omitempty"`
	HotelName              *string `json:"hotel_name,omitempty" xml:"hotel_name,omitempty"`
	// example:
	//
	// 5
	HotelStar *string `json:"hotel_star,omitempty" xml:"hotel_star,omitempty"`
	// example:
	//
	// 4564547
	Index            *string `json:"index,omitempty" xml:"index,omitempty"`
	InvoiceTitle     *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	IsEarlyDeparture *string `json:"is_early_departure,omitempty" xml:"is_early_departure,omitempty"`
	IsNegotiation    *string `json:"is_negotiation,omitempty" xml:"is_negotiation,omitempty"`
	IsShareStr       *string `json:"is_share_str,omitempty" xml:"is_share_str,omitempty"`
	Location         *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// 132143534543
	MainApplyId        *string `json:"main_apply_id,omitempty" xml:"main_apply_id,omitempty"`
	MappingCompanyCode *string `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	// example:
	//
	// 2
	Nights *int32 `json:"nights,omitempty" xml:"nights,omitempty"`
	// example:
	//
	// 110285961234324
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 1088.96
	OrderPrice *float64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	// example:
	//
	// null
	OrderStatusDesc     *string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	OrderType           *string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	OriginalReserveRule *string `json:"original_reserve_rule,omitempty" xml:"original_reserve_rule,omitempty"`
	// example:
	//
	// 534545345
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
	// 4564547
	PrimaryId       *int64  `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProcessorOaCode *string `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// 2345235435
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// example:
	//
	// 113.02
	PromotionFee *float64 `json:"promotion_fee,omitempty" xml:"promotion_fee,omitempty"`
	// example:
	//
	// 1.0d
	Rate   *string `json:"rate,omitempty" xml:"rate,omitempty"`
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 1
	ReserveRule *int32  `json:"reserve_rule,omitempty" xml:"reserve_rule,omitempty"`
	RoomNo      *string `json:"room_no,omitempty" xml:"room_no,omitempty"`
	// example:
	//
	// 1
	RoomNumber *int32 `json:"room_number,omitempty" xml:"room_number,omitempty"`
	// example:
	//
	// 1201.98
	RoomPrice *float64 `json:"room_price,omitempty" xml:"room_price,omitempty"`
	RoomType  *string  `json:"room_type,omitempty" xml:"room_type,omitempty"`
	// example:
	//
	// 23.9
	ServiceFee     *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettleTypeDesc *string  `json:"settle_type_desc,omitempty" xml:"settle_type_desc,omitempty"`
	// example:
	//
	// 350
	SettlementFee *float64 `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	// example:
	//
	// 6.11
	SettlementGrantFee *float64 `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	SettlementTime *string `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	// example:
	//
	// 4
	SettlementType *string `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	Sio            *string `json:"sio,omitempty" xml:"sio,omitempty"`
	// example:
	//
	// 1
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// 123123232
	SubOrderId *string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// example:
	//
	// 6%
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// example:
	//
	// 1321445511345
	ThirdInvoiceId *string `json:"third_invoice_id,omitempty" xml:"third_invoice_id,omitempty"`
	// example:
	//
	// AB0-CDE-1-F-1234567891011
	ThirdItineraryId *string `json:"third_itinerary_id,omitempty" xml:"third_itinerary_id,omitempty"`
	// example:
	//
	// 202311081011000348578
	ThirdPartBusinessId *string `json:"third_part_business_id,omitempty" xml:"third_part_business_id,omitempty"`
	// example:
	//
	// MGI18000230221072483
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	// example:
	//
	// 2
	TotalNights     *int32  `json:"total_nights,omitempty" xml:"total_nights,omitempty"`
	TradeActionDesc *string `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
	TravelerEmail   *string `json:"traveler_email,omitempty" xml:"traveler_email,omitempty"`
	// example:
	//
	// al_xinuan.zsy
	TravelerId *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// example:
	//
	// 345345
	TravelerJobNo      *string `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerMemberType *string `json:"traveler_member_type,omitempty" xml:"traveler_member_type,omitempty"`
	TravelerName       *string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	// example:
	//
	// 11
	VoucherType     *int32  `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	VoucherTypeDesc *string `json:"voucher_type_desc,omitempty" xml:"voucher_type_desc,omitempty"`
}

func (s IeHotelBillSettlementQueryResponseBodyModuleDataList) String() string {
	return dara.Prettify(s)
}

func (s IeHotelBillSettlementQueryResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetAmountCurrency() *string {
	return s.AmountCurrency
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityCode() *string {
	return s.ApplyArrCityCode
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityName() *string {
	return s.ApplyArrCityName
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityCode() *string {
	return s.ApplyDepCityCode
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityName() *string {
	return s.ApplyDepCityName
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetApplyId() *string {
	return s.ApplyId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetApproverEmail() *string {
	return s.ApproverEmail
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetApproverId() *string {
	return s.ApproverId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetApproverName() *string {
	return s.ApproverName
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetAverageNights() *float64 {
	return s.AverageNights
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBaseLocation() *string {
	return s.BaseLocation
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBookMode() *string {
	return s.BookMode
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBookReason() *string {
	return s.BookReason
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBookTime() *string {
	return s.BookTime
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBookerId() *string {
	return s.BookerId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBookerName() *string {
	return s.BookerName
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBrandGroup() *string {
	return s.BrandGroup
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBrandName() *string {
	return s.BrandName
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBusinessExpense() *int64 {
	return s.BusinessExpense
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetBusinessTripResult() *string {
	return s.BusinessTripResult
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCheckInDate() *string {
	return s.CheckInDate
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCheckoutDate() *string {
	return s.CheckoutDate
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCity() *string {
	return s.City
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCityCode() *string {
	return s.CityCode
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCorpRefundFee() *float64 {
	return s.CorpRefundFee
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCorpTotalFee() *float64 {
	return s.CorpTotalFee
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCostCenter() *string {
	return s.CostCenter
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCountry() *string {
	return s.Country
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCountryCode() *string {
	return s.CountryCode
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetCustomContent() *string {
	return s.CustomContent
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetDeductibleTax() *float64 {
	return s.DeductibleTax
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetDepartment() *string {
	return s.Department
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetExceedReason() *string {
	return s.ExceedReason
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetFeeType() *string {
	return s.FeeType
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetFines() *float64 {
	return s.Fines
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetForeignBusinessExpense() *int64 {
	return s.ForeignBusinessExpense
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetForeignersTag() *string {
	return s.ForeignersTag
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetHotelName() *string {
	return s.HotelName
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetHotelStar() *string {
	return s.HotelStar
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetIndex() *string {
	return s.Index
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetIsEarlyDeparture() *string {
	return s.IsEarlyDeparture
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetIsNegotiation() *string {
	return s.IsNegotiation
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetIsShareStr() *string {
	return s.IsShareStr
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetLocation() *string {
	return s.Location
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetMainApplyId() *string {
	return s.MainApplyId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetNights() *int32 {
	return s.Nights
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetOrderId() *string {
	return s.OrderId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetOrderPrice() *float64 {
	return s.OrderPrice
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetOrderType() *string {
	return s.OrderType
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetOriginalReserveRule() *string {
	return s.OriginalReserveRule
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetOverApplyId() *string {
	return s.OverApplyId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetPersonRefundFee() *float64 {
	return s.PersonRefundFee
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetPersonSettlePrice() *float64 {
	return s.PersonSettlePrice
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetPosition() *string {
	return s.Position
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetProjectName() *string {
	return s.ProjectName
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetPromotionFee() *float64 {
	return s.PromotionFee
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetRate() *string {
	return s.Rate
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetRemark() *string {
	return s.Remark
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetReserveRule() *int32 {
	return s.ReserveRule
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetRoomNo() *string {
	return s.RoomNo
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetRoomNumber() *int32 {
	return s.RoomNumber
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetRoomPrice() *float64 {
	return s.RoomPrice
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetRoomType() *string {
	return s.RoomType
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetSettlementGrantFee() *float64 {
	return s.SettlementGrantFee
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetSettlementType() *string {
	return s.SettlementType
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetSio() *string {
	return s.Sio
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetStatus() *int32 {
	return s.Status
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetTaxRate() *string {
	return s.TaxRate
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetThirdInvoiceId() *string {
	return s.ThirdInvoiceId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetThirdItineraryId() *string {
	return s.ThirdItineraryId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetThirdPartBusinessId() *string {
	return s.ThirdPartBusinessId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetTotalNights() *int32 {
	return s.TotalNights
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetTravelerEmail() *string {
	return s.TravelerEmail
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetTravelerId() *string {
	return s.TravelerId
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetTravelerName() *string {
	return s.TravelerName
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetAdjustTime(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.AdjustTime = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetAlipayTradeNo(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetAmountCurrency(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.AmountCurrency = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityCode(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityCode = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityName(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityName = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityCode(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityCode = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityName(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityName = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetApplyExtendField(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyExtendField = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetApplyId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetApproverEmail(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverEmail = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetApproverId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetApproverName(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverName = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetAverageNights(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.AverageNights = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBaseLocation(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BaseLocation = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBillRecordTime(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BillRecordTime = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBookMode(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookMode = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBookReason(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookReason = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBookTime(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBookerId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBookerJobNo(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBookerName(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBrandGroup(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BrandGroup = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBrandName(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BrandName = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBusinessExpense(v int64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BusinessExpense = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetBusinessTripResult(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.BusinessTripResult = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCapitalDirection(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCascadeDepartment(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCategoryDesc(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CategoryDesc = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCheckInDate(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CheckInDate = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCheckoutDate(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CheckoutDate = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCity(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.City = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCityCode(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CityCode = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCorpRefundFee(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CorpRefundFee = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCorpTotalFee(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CorpTotalFee = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCostCenter(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCostCenterNumber(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCostDepartment(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CostDepartment = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCountry(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.Country = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCountryCode(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CountryCode = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetCustomContent(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.CustomContent = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetDeductibleTax(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.DeductibleTax = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetDepartment(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetDepartmentId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetExceedReason(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ExceedReason = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetFeeType(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetFeeTypeDesc(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.FeeTypeDesc = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetFines(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.Fines = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetForeignBusinessExpense(v int64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ForeignBusinessExpense = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetForeignersTag(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ForeignersTag = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetHotelName(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.HotelName = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetHotelStar(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.HotelStar = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetIndex(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetInvoiceTitle(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetIsEarlyDeparture(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.IsEarlyDeparture = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetIsNegotiation(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.IsNegotiation = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetIsShareStr(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.IsShareStr = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetLocation(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.Location = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetMainApplyId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.MainApplyId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetMappingCompanyCode(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.MappingCompanyCode = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetNights(v int32) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.Nights = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetOrderId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetOrderPrice(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.OrderPrice = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetOrderStatusDesc(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.OrderStatusDesc = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetOrderType(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.OrderType = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetOriginalReserveRule(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.OriginalReserveRule = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetOverApplyId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentName(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetPersonRefundFee(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.PersonRefundFee = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetPersonSettlePrice(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.PersonSettlePrice = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetPosition(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.Position = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetPositionLevel(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.PositionLevel = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetPrimaryId(v int64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetProcessorOaCode(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ProcessorOaCode = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetProjectCode(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetProjectName(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetPromotionFee(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.PromotionFee = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetRate(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.Rate = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetRemark(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.Remark = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetReserveRule(v int32) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ReserveRule = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetRoomNo(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.RoomNo = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetRoomNumber(v int32) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.RoomNumber = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetRoomPrice(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.RoomPrice = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetRoomType(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.RoomType = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetServiceFee(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetSettleTypeDesc(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettleTypeDesc = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementFee(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementGrantFee(v float64) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementGrantFee = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementTime(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetSettlementType(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetSio(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.Sio = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetStatus(v int32) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetStatusDesc(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.StatusDesc = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetSubOrderId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.SubOrderId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetTaxRate(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.TaxRate = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetThirdInvoiceId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ThirdInvoiceId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetThirdItineraryId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ThirdItineraryId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetThirdPartBusinessId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ThirdPartBusinessId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetThirdpartApplyId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.ThirdpartApplyId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetTotalNights(v int32) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.TotalNights = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetTradeActionDesc(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.TradeActionDesc = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerEmail(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerEmail = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerId(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerJobNo(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerMemberType(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerMemberType = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetTravelerName(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetVoucherType(v int32) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) SetVoucherTypeDesc(v string) *IeHotelBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherTypeDesc = &v
	return s
}

func (s *IeHotelBillSettlementQueryResponseBodyModuleDataList) Validate() error {
	return dara.Validate(s)
}
