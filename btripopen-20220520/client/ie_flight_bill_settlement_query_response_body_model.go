// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeFlightBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IeFlightBillSettlementQueryResponseBody
	GetCode() *string
	SetMessage(v string) *IeFlightBillSettlementQueryResponseBody
	GetMessage() *string
	SetModule(v *IeFlightBillSettlementQueryResponseBodyModule) *IeFlightBillSettlementQueryResponseBody
	GetModule() *IeFlightBillSettlementQueryResponseBodyModule
	SetMorePage(v bool) *IeFlightBillSettlementQueryResponseBody
	GetMorePage() *bool
	SetRequestId(v string) *IeFlightBillSettlementQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IeFlightBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IeFlightBillSettlementQueryResponseBody
	GetTraceId() *string
}

type IeFlightBillSettlementQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                        `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                        `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IeFlightBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// true
	MorePage *bool `json:"more_page,omitempty" xml:"more_page,omitempty"`
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 210e842b16611337974412836dae27
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IeFlightBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IeFlightBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *IeFlightBillSettlementQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IeFlightBillSettlementQueryResponseBody) GetModule() *IeFlightBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *IeFlightBillSettlementQueryResponseBody) GetMorePage() *bool {
	return s.MorePage
}

func (s *IeFlightBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IeFlightBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IeFlightBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IeFlightBillSettlementQueryResponseBody) SetCode(v string) *IeFlightBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetMessage(v string) *IeFlightBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetModule(v *IeFlightBillSettlementQueryResponseBodyModule) *IeFlightBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetMorePage(v bool) *IeFlightBillSettlementQueryResponseBody {
	s.MorePage = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetRequestId(v string) *IeFlightBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetSuccess(v bool) *IeFlightBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) SetTraceId(v string) *IeFlightBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type IeFlightBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 11
	Category *int32 `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// corp1
	CorpId   *string                                                  `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	DataList []*IeFlightBillSettlementQueryResponseBodyModuleDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
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

func (s IeFlightBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IeFlightBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) GetDataList() []*IeFlightBillSettlementQueryResponseBodyModuleDataList {
	return s.DataList
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetCategory(v int32) *IeFlightBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetCorpId(v string) *IeFlightBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetDataList(v []*IeFlightBillSettlementQueryResponseBodyModuleDataList) *IeFlightBillSettlementQueryResponseBodyModule {
	s.DataList = v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *IeFlightBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *IeFlightBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetScrollId(v string) *IeFlightBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *IeFlightBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type IeFlightBillSettlementQueryResponseBodyModuleDataList struct {
	AdjustTime *string `json:"adjust_time,omitempty" xml:"adjust_time,omitempty"`
	// example:
	//
	// 1
	AdvanceDay *int32 `json:"advance_day,omitempty" xml:"advance_day,omitempty"`
	// example:
	//
	// MU
	AirlineCorpCode *string `json:"airline_corp_code,omitempty" xml:"airline_corp_code,omitempty"`
	AirlineCorpName *string `json:"airline_corp_name,omitempty" xml:"airline_corp_name,omitempty"`
	// example:
	//
	// 2021123432260
	AlipayTradeNo    *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	ApplyArrCityCode *string `json:"apply_arr_city_code,omitempty" xml:"apply_arr_city_code,omitempty"`
	ApplyArrCityName *string `json:"apply_arr_city_name,omitempty" xml:"apply_arr_city_name,omitempty"`
	ApplyDepCityCode *string `json:"apply_dep_city_code,omitempty" xml:"apply_dep_city_code,omitempty"`
	ApplyDepCityName *string `json:"apply_dep_city_name,omitempty" xml:"apply_dep_city_name,omitempty"`
	// 审批扩展自定义字段
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
	// CAN
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrCity        *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityCode    *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCountry     *string `json:"arr_country,omitempty" xml:"arr_country,omitempty"`
	ArrCountryCode *string `json:"arr_country_code,omitempty" xml:"arr_country_code,omitempty"`
	// example:
	//
	// 2021-10-02
	ArrDate    *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	ArrStation *string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	// example:
	//
	// 13:30:00
	ArrTime      *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BaseLocation *string `json:"base_location,omitempty" xml:"base_location,omitempty"`
	// example:
	//
	// 2020-12-23T20:18Z
	BillRecordTime *string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BookMode       *string `json:"book_mode,omitempty" xml:"book_mode,omitempty"`
	// example:
	//
	// 2021-10-01 00:17:05
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
	// example:
	//
	// 23.9
	BtripCouponFee     *float64 `json:"btrip_coupon_fee,omitempty" xml:"btrip_coupon_fee,omitempty"`
	BusinessTripResult *string  `json:"business_trip_result,omitempty" xml:"business_trip_result,omitempty"`
	// example:
	//
	// R
	Cabin      *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	// example:
	//
	// 1
	CapitalDirection  *string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment *string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CategoryDesc      *string `json:"category_desc,omitempty" xml:"category_desc,omitempty"`
	// example:
	//
	// 23.0
	ChangeFee    *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	ChangeResult *string  `json:"change_result,omitempty" xml:"change_result,omitempty"`
	// example:
	//
	// 460
	CorpPayOrderFee *float64 `json:"corp_pay_order_fee,omitempty" xml:"corp_pay_order_fee,omitempty"`
	CostCenter      *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// example:
	//
	// 8b7f3cd-24324-097
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	CostDepartment   *string `json:"cost_department,omitempty" xml:"cost_department,omitempty"`
	// example:
	//
	// 1
	Coupon        *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	CustomContent *string  `json:"custom_content,omitempty" xml:"custom_content,omitempty"`
	DeductibleTax *float64 `json:"deductible_tax,omitempty" xml:"deductible_tax,omitempty"`
	// example:
	//
	// KHN
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	DepCityCode    *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCountry     *string `json:"dep_country,omitempty" xml:"dep_country,omitempty"`
	DepCountryCode *string `json:"dep_country_code,omitempty" xml:"dep_country_code,omitempty"`
	Department     *string `json:"department,omitempty" xml:"department,omitempty"`
	// example:
	//
	// 2345866
	DepartmentId *string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	DeptCity     *string `json:"dept_city,omitempty" xml:"dept_city,omitempty"`
	// example:
	//
	// 2021-10-02
	DeptDate    *string `json:"dept_date,omitempty" xml:"dept_date,omitempty"`
	DeptStation *string `json:"dept_station,omitempty" xml:"dept_station,omitempty"`
	// example:
	//
	// 12:00:00
	DeptTime *string `json:"dept_time,omitempty" xml:"dept_time,omitempty"`
	// example:
	//
	// 51%
	Discount     *string `json:"discount,omitempty" xml:"discount,omitempty"`
	ExceedReason *string `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	// example:
	//
	// 20101
	FeeType     *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	// example:
	//
	// MU9684
	FlightNo      *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	ForeignersTag *string `json:"foreigners_tag,omitempty" xml:"foreigners_tag,omitempty"`
	// example:
	//
	// 4564547
	Index      *string `json:"index,omitempty" xml:"index,omitempty"`
	InsOrderId *string `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	// example:
	//
	// 23.9
	InsuranceFee *float64 `json:"insurance_fee,omitempty" xml:"insurance_fee,omitempty"`
	// example:
	//
	// 15548214852
	InsuranceNumber      *string `json:"insurance_number,omitempty" xml:"insurance_number,omitempty"`
	InsuranceProductName *string `json:"insurance_product_name,omitempty" xml:"insurance_product_name,omitempty"`
	InvoiceTitle         *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	Location             *string `json:"location,omitempty" xml:"location,omitempty"`
	MappingCompanyCode   *string `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	// example:
	//
	// 2021-10-01 00:17:13
	MostDifferenceDeptTime *string `json:"most_difference_dept_time,omitempty" xml:"most_difference_dept_time,omitempty"`
	// example:
	//
	// 23%
	MostDifferenceDiscount *string `json:"most_difference_discount,omitempty" xml:"most_difference_discount,omitempty"`
	// example:
	//
	// MU9684
	MostDifferenceFlightNo *string `json:"most_difference_flight_no,omitempty" xml:"most_difference_flight_no,omitempty"`
	// example:
	//
	// 23.9
	MostDifferencePrice  *float64 `json:"most_difference_price,omitempty" xml:"most_difference_price,omitempty"`
	MostDifferenceReason *string  `json:"most_difference_reason,omitempty" xml:"most_difference_reason,omitempty"`
	// example:
	//
	// 23.9
	MostPrice *float64 `json:"most_price,omitempty" xml:"most_price,omitempty"`
	// example:
	//
	// 23.9
	NegotiationCouponFee *float64 `json:"negotiation_coupon_fee,omitempty" xml:"negotiation_coupon_fee,omitempty"`
	// example:
	//
	// 234223423423
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// null
	OrderStatusDesc *string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// example:
	//
	// 234324324423
	OverApplyId           *string `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	PaymentDepartmentId   *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	Position              *string `json:"position,omitempty" xml:"position,omitempty"`
	PositionLevel         *string `json:"position_level,omitempty" xml:"position_level,omitempty"`
	// example:
	//
	// 4564547
	PrimaryId       *int64  `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProcessorOaCode *string `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// 23423432423
	ProjectCode      *string  `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName      *string  `json:"project_name,omitempty" xml:"project_name,omitempty"`
	RefundChangeCost *float64 `json:"refund_change_cost,omitempty" xml:"refund_change_cost,omitempty"`
	// example:
	//
	// 23.9
	RefundFee    *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	RefundResult *string  `json:"refund_result,omitempty" xml:"refund_result,omitempty"`
	Remark       *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	RepeatRefund *string  `json:"repeat_refund,omitempty" xml:"repeat_refund,omitempty"`
	// example:
	//
	// 410
	SealPrice   *float64 `json:"seal_price,omitempty" xml:"seal_price,omitempty"`
	SegmentType *string  `json:"segment_type,omitempty" xml:"segment_type,omitempty"`
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
	// 2021-10-08 23:39:01
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
	// 125.6
	TaxFee *float64 `json:"tax_fee,omitempty" xml:"tax_fee,omitempty"`
	// 税率
	//
	// example:
	//
	// 6%
	TaxRate          *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	ThirdItineraryId *string `json:"third_itinerary_id,omitempty" xml:"third_itinerary_id,omitempty"`
	// example:
	//
	// 781-6586234234324
	TicketId *string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	// example:
	//
	// CAN-KUL-BKK
	Trade           *string `json:"trade,omitempty" xml:"trade,omitempty"`
	TradeActionDesc *string `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
	TravelerEmail   *string `json:"traveler_email,omitempty" xml:"traveler_email,omitempty"`
	// example:
	//
	// 54463464
	TravelerId *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// example:
	//
	// 326246
	TravelerJobNo          *string `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerMemberType     *string `json:"traveler_member_type,omitempty" xml:"traveler_member_type,omitempty"`
	TravelerMemberTypeName *string `json:"traveler_member_type_name,omitempty" xml:"traveler_member_type_name,omitempty"`
	TravelerName           *string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	TripType               *int32  `json:"trip_type,omitempty" xml:"trip_type,omitempty"`
	// example:
	//
	// 11
	VoucherType     *int32  `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	VoucherTypeDesc *string `json:"voucher_type_desc,omitempty" xml:"voucher_type_desc,omitempty"`
	VoyageName      *string `json:"voyage_name,omitempty" xml:"voyage_name,omitempty"`
}

func (s IeFlightBillSettlementQueryResponseBodyModuleDataList) String() string {
	return dara.Prettify(s)
}

func (s IeFlightBillSettlementQueryResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetAdvanceDay() *int32 {
	return s.AdvanceDay
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetAirlineCorpCode() *string {
	return s.AirlineCorpCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetAirlineCorpName() *string {
	return s.AirlineCorpName
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityCode() *string {
	return s.ApplyArrCityCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityName() *string {
	return s.ApplyArrCityName
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityCode() *string {
	return s.ApplyDepCityCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityName() *string {
	return s.ApplyDepCityName
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetApplyId() *string {
	return s.ApplyId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetApproverEmail() *string {
	return s.ApproverEmail
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetApproverId() *string {
	return s.ApproverId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetApproverName() *string {
	return s.ApproverName
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetArrCity() *string {
	return s.ArrCity
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetArrCountry() *string {
	return s.ArrCountry
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetArrCountryCode() *string {
	return s.ArrCountryCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetArrDate() *string {
	return s.ArrDate
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetArrStation() *string {
	return s.ArrStation
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetArrTime() *string {
	return s.ArrTime
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetBaseLocation() *string {
	return s.BaseLocation
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetBookMode() *string {
	return s.BookMode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetBookTime() *string {
	return s.BookTime
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetBookerId() *string {
	return s.BookerId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetBookerName() *string {
	return s.BookerName
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetBtripCouponFee() *float64 {
	return s.BtripCouponFee
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetBusinessTripResult() *string {
	return s.BusinessTripResult
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetCabin() *string {
	return s.Cabin
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetChangeFee() *float64 {
	return s.ChangeFee
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetChangeResult() *string {
	return s.ChangeResult
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetCorpPayOrderFee() *float64 {
	return s.CorpPayOrderFee
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetCostCenter() *string {
	return s.CostCenter
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetCoupon() *float64 {
	return s.Coupon
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetCustomContent() *string {
	return s.CustomContent
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDeductibleTax() *float64 {
	return s.DeductibleTax
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDepCountry() *string {
	return s.DepCountry
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDepCountryCode() *string {
	return s.DepCountryCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDepartment() *string {
	return s.Department
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDeptCity() *string {
	return s.DeptCity
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDeptDate() *string {
	return s.DeptDate
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDeptStation() *string {
	return s.DeptStation
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDeptTime() *string {
	return s.DeptTime
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetDiscount() *string {
	return s.Discount
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetExceedReason() *string {
	return s.ExceedReason
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetFeeType() *string {
	return s.FeeType
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetForeignersTag() *string {
	return s.ForeignersTag
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetIndex() *string {
	return s.Index
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetInsuranceFee() *float64 {
	return s.InsuranceFee
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetInsuranceNumber() *string {
	return s.InsuranceNumber
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetInsuranceProductName() *string {
	return s.InsuranceProductName
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetLocation() *string {
	return s.Location
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetMostDifferenceDeptTime() *string {
	return s.MostDifferenceDeptTime
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetMostDifferenceDiscount() *string {
	return s.MostDifferenceDiscount
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetMostDifferenceFlightNo() *string {
	return s.MostDifferenceFlightNo
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetMostDifferencePrice() *float64 {
	return s.MostDifferencePrice
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetMostDifferenceReason() *string {
	return s.MostDifferenceReason
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetMostPrice() *float64 {
	return s.MostPrice
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetNegotiationCouponFee() *float64 {
	return s.NegotiationCouponFee
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetOrderId() *string {
	return s.OrderId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetOverApplyId() *string {
	return s.OverApplyId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetPosition() *string {
	return s.Position
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetProjectName() *string {
	return s.ProjectName
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetRefundChangeCost() *float64 {
	return s.RefundChangeCost
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetRefundResult() *string {
	return s.RefundResult
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetRemark() *string {
	return s.Remark
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetRepeatRefund() *string {
	return s.RepeatRefund
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetSealPrice() *float64 {
	return s.SealPrice
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetSegmentType() *string {
	return s.SegmentType
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetSettlementGrantFee() *float64 {
	return s.SettlementGrantFee
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetSettlementType() *string {
	return s.SettlementType
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetSio() *string {
	return s.Sio
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetStatus() *int32 {
	return s.Status
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTaxFee() *float64 {
	return s.TaxFee
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTaxRate() *string {
	return s.TaxRate
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetThirdItineraryId() *string {
	return s.ThirdItineraryId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTicketId() *string {
	return s.TicketId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTrade() *string {
	return s.Trade
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerEmail() *string {
	return s.TravelerEmail
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerId() *string {
	return s.TravelerId
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerMemberTypeName() *string {
	return s.TravelerMemberTypeName
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerName() *string {
	return s.TravelerName
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetTripType() *int32 {
	return s.TripType
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) GetVoyageName() *string {
	return s.VoyageName
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetAdjustTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.AdjustTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetAdvanceDay(v int32) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.AdvanceDay = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetAirlineCorpCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.AirlineCorpCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetAirlineCorpName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.AirlineCorpName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetAlipayTradeNo(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetApplyExtendField(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyExtendField = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetApplyId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetApproverEmail(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverEmail = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetApproverId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetApproverName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrAirportCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrAirportCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrCity(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCity = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrCityCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCityCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrCountry(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCountry = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrCountryCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCountryCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrDate(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrDate = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrStation(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrStation = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetArrTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBaseLocation(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BaseLocation = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBillRecordTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BillRecordTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBookMode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookMode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBookTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBookerId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBookerJobNo(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBookerName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBtripCouponFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BtripCouponFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetBusinessTripResult(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.BusinessTripResult = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCabin(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Cabin = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCabinClass(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CabinClass = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCapitalDirection(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCascadeDepartment(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCategoryDesc(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CategoryDesc = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetChangeFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetChangeResult(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeResult = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCorpPayOrderFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CorpPayOrderFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCostCenter(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCostCenterNumber(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCostDepartment(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CostDepartment = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCoupon(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Coupon = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetCustomContent(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.CustomContent = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDeductibleTax(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeductibleTax = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDepAirportCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepAirportCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDepCityCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepCityCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDepCountry(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepCountry = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDepCountryCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepCountryCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDepartment(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDepartmentId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDeptCity(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptCity = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDeptDate(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptDate = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDeptStation(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptStation = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDeptTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetDiscount(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Discount = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetExceedReason(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ExceedReason = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetFeeType(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetFeeTypeDesc(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.FeeTypeDesc = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetFlightNo(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.FlightNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetForeignersTag(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ForeignersTag = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetIndex(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetInsOrderId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.InsOrderId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetInsuranceFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.InsuranceFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetInsuranceNumber(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.InsuranceNumber = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetInsuranceProductName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.InsuranceProductName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetInvoiceTitle(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetLocation(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Location = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMappingCompanyCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MappingCompanyCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceDeptTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceDeptTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceDiscount(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceDiscount = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceFlightNo(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceFlightNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferencePrice(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferencePrice = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceReason(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceReason = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetMostPrice(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostPrice = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetNegotiationCouponFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.NegotiationCouponFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetOrderId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetOrderStatusDesc(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.OrderStatusDesc = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetOverApplyId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetPosition(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Position = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetPositionLevel(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.PositionLevel = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetPrimaryId(v int64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetProcessorOaCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ProcessorOaCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetProjectCode(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetProjectName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetRefundChangeCost(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundChangeCost = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetRefundFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetRefundResult(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundResult = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetRemark(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Remark = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetRepeatRefund(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.RepeatRefund = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSealPrice(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SealPrice = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSegmentType(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SegmentType = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetServiceFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSettleTypeDesc(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettleTypeDesc = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementGrantFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementGrantFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementTime(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementType(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSio(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Sio = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetStatus(v int32) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetStatusDesc(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.StatusDesc = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetSubOrderId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.SubOrderId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTaxFee(v float64) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TaxFee = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTaxRate(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TaxRate = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetThirdItineraryId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.ThirdItineraryId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTicketId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TicketId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTrade(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.Trade = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTradeActionDesc(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TradeActionDesc = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerEmail(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerEmail = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerId(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerJobNo(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerMemberType(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerMemberType = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerMemberTypeName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerMemberTypeName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetTripType(v int32) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.TripType = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetVoucherType(v int32) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetVoucherTypeDesc(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherTypeDesc = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) SetVoyageName(v string) *IeFlightBillSettlementQueryResponseBodyModuleDataList {
	s.VoyageName = &v
	return s
}

func (s *IeFlightBillSettlementQueryResponseBodyModuleDataList) Validate() error {
	return dara.Validate(s)
}
