// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlightBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *FlightBillSettlementQueryResponseBody
	GetCode() *string
	SetMessage(v string) *FlightBillSettlementQueryResponseBody
	GetMessage() *string
	SetModule(v *FlightBillSettlementQueryResponseBodyModule) *FlightBillSettlementQueryResponseBody
	GetModule() *FlightBillSettlementQueryResponseBodyModule
	SetRequestId(v string) *FlightBillSettlementQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FlightBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FlightBillSettlementQueryResponseBody
	GetTraceId() *string
}

type FlightBillSettlementQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                      `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                      `json:"message,omitempty" xml:"message,omitempty"`
	Module  *FlightBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s FlightBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FlightBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *FlightBillSettlementQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *FlightBillSettlementQueryResponseBody) GetModule() *FlightBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *FlightBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FlightBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FlightBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FlightBillSettlementQueryResponseBody) SetCode(v string) *FlightBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBody) SetMessage(v string) *FlightBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBody) SetModule(v *FlightBillSettlementQueryResponseBodyModule) *FlightBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *FlightBillSettlementQueryResponseBody) SetRequestId(v string) *FlightBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBody) SetSuccess(v bool) *FlightBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBody) SetTraceId(v string) *FlightBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type FlightBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 1
	Category *int32                                                 `json:"category,omitempty" xml:"category,omitempty"`
	CorpId   *string                                                `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	DataList []*FlightBillSettlementQueryResponseBodyModuleDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
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
	// 5180
	TotalNum *int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

func (s FlightBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FlightBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *FlightBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *FlightBillSettlementQueryResponseBodyModule) GetDataList() []*FlightBillSettlementQueryResponseBodyModuleDataList {
	return s.DataList
}

func (s *FlightBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *FlightBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *FlightBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *FlightBillSettlementQueryResponseBodyModule) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetCategory(v int32) *FlightBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetCorpId(v string) *FlightBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetDataList(v []*FlightBillSettlementQueryResponseBodyModuleDataList) *FlightBillSettlementQueryResponseBodyModule {
	s.DataList = v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *FlightBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *FlightBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetScrollId(v string) *FlightBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *FlightBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FlightBillSettlementQueryResponseBodyModuleDataList struct {
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
	AlipayId        *string `json:"alipay_id,omitempty" xml:"alipay_id,omitempty"`
	// example:
	//
	// 2021100122001138061456080520
	AlipayTradeNo    *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	ApplyArrCityCode *string `json:"apply_arr_city_code,omitempty" xml:"apply_arr_city_code,omitempty"`
	ApplyArrCityName *string `json:"apply_arr_city_name,omitempty" xml:"apply_arr_city_name,omitempty"`
	ApplyDepCityCode *string `json:"apply_dep_city_code,omitempty" xml:"apply_dep_city_code,omitempty"`
	ApplyDepCityName *string `json:"apply_dep_city_name,omitempty" xml:"apply_dep_city_name,omitempty"`
	// 审批扩展自定义字段
	ApplyExtendField *string `json:"apply_extend_field,omitempty" xml:"apply_extend_field,omitempty"`
	// example:
	//
	// 103177854
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
	BookChannel    *string `json:"book_channel,omitempty" xml:"book_channel,omitempty"`
	BookMode       *string `json:"book_mode,omitempty" xml:"book_mode,omitempty"`
	// example:
	//
	// 2021-10-01 00:17:05
	BookTime *string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	BookerId *string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	// example:
	//
	// 2342432
	BookerJobNo *string `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName  *string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	// example:
	//
	// 23.9
	BtripCouponFee *float64 `json:"btrip_coupon_fee,omitempty" xml:"btrip_coupon_fee,omitempty"`
	// example:
	//
	// 50
	BuildFee           *float64 `json:"build_fee,omitempty" xml:"build_fee,omitempty"`
	BusinessTripResult *string  `json:"business_trip_result,omitempty" xml:"business_trip_result,omitempty"`
	// example:
	//
	// R
	Cabin          *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	CabinClass     *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassCode *string `json:"cabin_class_code,omitempty" xml:"cabin_class_code,omitempty"`
	// example:
	//
	// 1
	CapitalDirection  *string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment *string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CategoryDesc      *string `json:"category_desc,omitempty" xml:"category_desc,omitempty"`
	// example:
	//
	// 23.9
	ChangeFee        *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	ChangeReasonCode *string  `json:"change_reason_code,omitempty" xml:"change_reason_code,omitempty"`
	ChangeResult     *string  `json:"change_result,omitempty" xml:"change_result,omitempty"`
	// example:
	//
	// 460
	CorpPayOrderFee *float64 `json:"corp_pay_order_fee,omitempty" xml:"corp_pay_order_fee,omitempty"`
	CorpSettlePrice *float64 `json:"corp_settle_price,omitempty" xml:"corp_settle_price,omitempty"`
	CostCenter      *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// example:
	//
	// 48b7f3cd-8a4f-4df9-ae2c-883e008cf097
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
	Department     *string `json:"department,omitempty" xml:"department,omitempty"`
	DepartmentId   *string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	DeptCity       *string `json:"dept_city,omitempty" xml:"dept_city,omitempty"`
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
	// 10101
	FeeType     *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	// example:
	//
	// MU9684
	FlightNo       *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightTripType *string `json:"flight_trip_type,omitempty" xml:"flight_trip_type,omitempty"`
	ForeignersTag  *string `json:"foreigners_tag,omitempty" xml:"foreigners_tag,omitempty"`
	// example:
	//
	// 4564547
	Index      *string `json:"index,omitempty" xml:"index,omitempty"`
	InsOrderId *string `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	// example:
	//
	// 23.9
	InsuranceFee    *float64 `json:"insurance_fee,omitempty" xml:"insurance_fee,omitempty"`
	InsuranceNumber *string  `json:"insurance_number,omitempty" xml:"insurance_number,omitempty"`
	InvoiceTitle    *string  `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	ItemType        *string  `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// example:
	//
	// 5334916421
	ItineraryNum *string `json:"itinerary_num,omitempty" xml:"itinerary_num,omitempty"`
	// example:
	//
	// 460
	ItineraryPrice     *float64 `json:"itinerary_price,omitempty" xml:"itinerary_price,omitempty"`
	Location           *string  `json:"location,omitempty" xml:"location,omitempty"`
	MappingCompanyCode *string  `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	Mileage            *string  `json:"mileage,omitempty" xml:"mileage,omitempty"`
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
	OfflineStatus        *string  `json:"offline_status,omitempty" xml:"offline_status,omitempty"`
	// example:
	//
	// 0
	OilFee *float64 `json:"oil_fee,omitempty" xml:"oil_fee,omitempty"`
	// example:
	//
	// 234223423423
	OrderId       *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderTicketNo *string `json:"order_ticket_no,omitempty" xml:"order_ticket_no,omitempty"`
	// example:
	//
	// 234324324423
	OverApplyId           *string  `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	PaymentDepartmentId   *string  `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string  `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	PersonSettlePrice     *float64 `json:"person_settle_price,omitempty" xml:"person_settle_price,omitempty"`
	Position              *string  `json:"position,omitempty" xml:"position,omitempty"`
	PositionLevel         *string  `json:"position_level,omitempty" xml:"position_level,omitempty"`
	PreBookTip            *string  `json:"pre_book_tip,omitempty" xml:"pre_book_tip,omitempty"`
	// example:
	//
	// 4564547
	PrimaryId       *int64  `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProcessorOaCode *string `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// 45623234
	ProjectCode       *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName       *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	RefundAffiliateNo *string `json:"refund_affiliate_no,omitempty" xml:"refund_affiliate_no,omitempty"`
	RefundApplyId     *string `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
	// example:
	//
	// 23.9
	RefundFee        *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	RefundReasonCode *string  `json:"refund_reason_code,omitempty" xml:"refund_reason_code,omitempty"`
	RefundResult     *string  `json:"refund_result,omitempty" xml:"refund_result,omitempty"`
	// example:
	//
	// 23.9
	RefundUpgradeCost *float64 `json:"refund_upgrade_cost,omitempty" xml:"refund_upgrade_cost,omitempty"`
	Remark            *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	RepeatRefund      *string  `json:"repeat_refund,omitempty" xml:"repeat_refund,omitempty"`
	SceneId           *string  `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	SceneName         *string  `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	// example:
	//
	// 410
	SealPrice *float64 `json:"seal_price,omitempty" xml:"seal_price,omitempty"`
	// example:
	//
	// 23.9
	ServiceFee     *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettleTypeDesc *string  `json:"settle_type_desc,omitempty" xml:"settle_type_desc,omitempty"`
	// example:
	//
	// 460
	SettlementFee *float64 `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	// example:
	//
	// 5.67
	SettlementGrantFee *float64 `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	// example:
	//
	// 2021-10-01 00:17:13
	SettlementTime *string `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	// example:
	//
	// 2
	SettlementType *string `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	Sio            *string `json:"sio,omitempty" xml:"sio,omitempty"`
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
	// 781-6586234234324
	TicketId        *string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	Trade           *string `json:"trade,omitempty" xml:"trade,omitempty"`
	TradeActionDesc *string `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
	TravelerEmail   *string `json:"traveler_email,omitempty" xml:"traveler_email,omitempty"`
	TravelerId      *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// example:
	//
	// 345345
	TravelerJobNo          *string `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerMemberType     *string `json:"traveler_member_type,omitempty" xml:"traveler_member_type,omitempty"`
	TravelerMemberTypeName *string `json:"traveler_member_type_name,omitempty" xml:"traveler_member_type_name,omitempty"`
	TravelerName           *string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	// example:
	//
	// 23.9
	UpgradeCost *float64 `json:"upgrade_cost,omitempty" xml:"upgrade_cost,omitempty"`
	// example:
	//
	// 11
	VoucherType     *int32  `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	VoucherTypeDesc *string `json:"voucher_type_desc,omitempty" xml:"voucher_type_desc,omitempty"`
	VoyageName      *string `json:"voyage_name,omitempty" xml:"voyage_name,omitempty"`
}

func (s FlightBillSettlementQueryResponseBodyModuleDataList) String() string {
	return dara.Prettify(s)
}

func (s FlightBillSettlementQueryResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetAdvanceDay() *int32 {
	return s.AdvanceDay
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetAirlineCorpCode() *string {
	return s.AirlineCorpCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetAirlineCorpName() *string {
	return s.AirlineCorpName
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetAlipayId() *string {
	return s.AlipayId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityCode() *string {
	return s.ApplyArrCityCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityName() *string {
	return s.ApplyArrCityName
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityCode() *string {
	return s.ApplyDepCityCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityName() *string {
	return s.ApplyDepCityName
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetApplyId() *string {
	return s.ApplyId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetApproverEmail() *string {
	return s.ApproverEmail
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetApproverId() *string {
	return s.ApproverId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetApproverName() *string {
	return s.ApproverName
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetArrCity() *string {
	return s.ArrCity
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetArrDate() *string {
	return s.ArrDate
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetArrStation() *string {
	return s.ArrStation
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetArrTime() *string {
	return s.ArrTime
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetBaseLocation() *string {
	return s.BaseLocation
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetBookChannel() *string {
	return s.BookChannel
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetBookMode() *string {
	return s.BookMode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetBookTime() *string {
	return s.BookTime
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetBookerId() *string {
	return s.BookerId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetBookerName() *string {
	return s.BookerName
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetBtripCouponFee() *float64 {
	return s.BtripCouponFee
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetBuildFee() *float64 {
	return s.BuildFee
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetBusinessTripResult() *string {
	return s.BusinessTripResult
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCabin() *string {
	return s.Cabin
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCabinClass() *string {
	return s.CabinClass
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCabinClassCode() *string {
	return s.CabinClassCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetChangeFee() *float64 {
	return s.ChangeFee
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetChangeReasonCode() *string {
	return s.ChangeReasonCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetChangeResult() *string {
	return s.ChangeResult
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCorpPayOrderFee() *float64 {
	return s.CorpPayOrderFee
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCorpSettlePrice() *float64 {
	return s.CorpSettlePrice
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCostCenter() *string {
	return s.CostCenter
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCoupon() *float64 {
	return s.Coupon
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetCustomContent() *string {
	return s.CustomContent
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetDeductibleTax() *float64 {
	return s.DeductibleTax
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetDepartment() *string {
	return s.Department
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetDeptCity() *string {
	return s.DeptCity
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetDeptDate() *string {
	return s.DeptDate
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetDeptStation() *string {
	return s.DeptStation
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetDeptTime() *string {
	return s.DeptTime
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetDiscount() *string {
	return s.Discount
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetExceedReason() *string {
	return s.ExceedReason
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetFeeType() *string {
	return s.FeeType
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetFlightNo() *string {
	return s.FlightNo
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetFlightTripType() *string {
	return s.FlightTripType
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetForeignersTag() *string {
	return s.ForeignersTag
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetIndex() *string {
	return s.Index
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetInsuranceFee() *float64 {
	return s.InsuranceFee
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetInsuranceNumber() *string {
	return s.InsuranceNumber
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetItemType() *string {
	return s.ItemType
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetItineraryNum() *string {
	return s.ItineraryNum
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetItineraryPrice() *float64 {
	return s.ItineraryPrice
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetLocation() *string {
	return s.Location
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetMileage() *string {
	return s.Mileage
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetMostDifferenceDeptTime() *string {
	return s.MostDifferenceDeptTime
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetMostDifferenceDiscount() *string {
	return s.MostDifferenceDiscount
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetMostDifferenceFlightNo() *string {
	return s.MostDifferenceFlightNo
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetMostDifferencePrice() *float64 {
	return s.MostDifferencePrice
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetMostDifferenceReason() *string {
	return s.MostDifferenceReason
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetMostPrice() *float64 {
	return s.MostPrice
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetNegotiationCouponFee() *float64 {
	return s.NegotiationCouponFee
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetOfflineStatus() *string {
	return s.OfflineStatus
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetOilFee() *float64 {
	return s.OilFee
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetOrderId() *string {
	return s.OrderId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetOrderTicketNo() *string {
	return s.OrderTicketNo
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetOverApplyId() *string {
	return s.OverApplyId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetPersonSettlePrice() *float64 {
	return s.PersonSettlePrice
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetPosition() *string {
	return s.Position
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetPreBookTip() *string {
	return s.PreBookTip
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetProjectName() *string {
	return s.ProjectName
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetRefundAffiliateNo() *string {
	return s.RefundAffiliateNo
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetRefundApplyId() *string {
	return s.RefundApplyId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetRefundReasonCode() *string {
	return s.RefundReasonCode
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetRefundResult() *string {
	return s.RefundResult
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetRefundUpgradeCost() *float64 {
	return s.RefundUpgradeCost
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetRemark() *string {
	return s.Remark
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetRepeatRefund() *string {
	return s.RepeatRefund
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetSceneId() *string {
	return s.SceneId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetSceneName() *string {
	return s.SceneName
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetSealPrice() *float64 {
	return s.SealPrice
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetSettlementGrantFee() *float64 {
	return s.SettlementGrantFee
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetSettlementType() *string {
	return s.SettlementType
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetSio() *string {
	return s.Sio
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetStatus() *int32 {
	return s.Status
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetTaxRate() *string {
	return s.TaxRate
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetThirdItineraryId() *string {
	return s.ThirdItineraryId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetTicketId() *string {
	return s.TicketId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetTrade() *string {
	return s.Trade
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerEmail() *string {
	return s.TravelerEmail
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerId() *string {
	return s.TravelerId
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerMemberTypeName() *string {
	return s.TravelerMemberTypeName
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetTravelerName() *string {
	return s.TravelerName
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetUpgradeCost() *float64 {
	return s.UpgradeCost
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) GetVoyageName() *string {
	return s.VoyageName
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetAdjustTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.AdjustTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetAdvanceDay(v int32) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.AdvanceDay = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetAirlineCorpCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.AirlineCorpCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetAirlineCorpName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.AirlineCorpName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetAlipayId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetAlipayTradeNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetApplyExtendField(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyExtendField = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetApplyId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetApproverEmail(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverEmail = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetApproverId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetApproverName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetArrAirportCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrAirportCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetArrCity(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCity = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetArrCityCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCityCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetArrDate(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrDate = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetArrStation(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrStation = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetArrTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ArrTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBaseLocation(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BaseLocation = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBillRecordTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BillRecordTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBookChannel(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookChannel = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBookMode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookMode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBookTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBookerId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBookerJobNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBookerName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBtripCouponFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BtripCouponFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBuildFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BuildFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetBusinessTripResult(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.BusinessTripResult = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCabin(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Cabin = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCabinClass(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CabinClass = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCabinClassCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CabinClassCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCapitalDirection(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCascadeDepartment(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCategoryDesc(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CategoryDesc = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetChangeFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetChangeReasonCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeReasonCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetChangeResult(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeResult = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCorpPayOrderFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CorpPayOrderFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCorpSettlePrice(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CorpSettlePrice = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCostCenter(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCostCenterNumber(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCostDepartment(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CostDepartment = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCoupon(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Coupon = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetCustomContent(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.CustomContent = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDeductibleTax(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeductibleTax = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDepAirportCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepAirportCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDepCityCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepCityCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDepartment(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDepartmentId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDeptCity(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptCity = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDeptDate(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptDate = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDeptStation(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptStation = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDeptTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.DeptTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetDiscount(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Discount = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetExceedReason(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ExceedReason = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetFeeType(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetFeeTypeDesc(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.FeeTypeDesc = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetFlightNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.FlightNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetFlightTripType(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.FlightTripType = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetForeignersTag(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ForeignersTag = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetIndex(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetInsOrderId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.InsOrderId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetInsuranceFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.InsuranceFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetInsuranceNumber(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.InsuranceNumber = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetInvoiceTitle(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetItemType(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ItemType = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetItineraryNum(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ItineraryNum = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetItineraryPrice(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ItineraryPrice = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetLocation(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Location = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMappingCompanyCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MappingCompanyCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMileage(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Mileage = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceDeptTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceDeptTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceDiscount(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceDiscount = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceFlightNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceFlightNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferencePrice(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferencePrice = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostDifferenceReason(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostDifferenceReason = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetMostPrice(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.MostPrice = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetNegotiationCouponFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.NegotiationCouponFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetOfflineStatus(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.OfflineStatus = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetOilFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.OilFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetOrderId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetOrderTicketNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.OrderTicketNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetOverApplyId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetPersonSettlePrice(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.PersonSettlePrice = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetPosition(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Position = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetPositionLevel(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.PositionLevel = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetPreBookTip(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.PreBookTip = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetPrimaryId(v int64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetProcessorOaCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ProcessorOaCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetProjectCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetProjectName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRefundAffiliateNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundAffiliateNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRefundApplyId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundApplyId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRefundFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRefundReasonCode(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundReasonCode = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRefundResult(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundResult = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRefundUpgradeCost(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.RefundUpgradeCost = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRemark(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Remark = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetRepeatRefund(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.RepeatRefund = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSceneId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SceneId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSceneName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SceneName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSealPrice(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SealPrice = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetServiceFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSettleTypeDesc(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettleTypeDesc = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementGrantFee(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementGrantFee = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementTime(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSettlementType(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetSio(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Sio = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetStatus(v int32) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetStatusDesc(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.StatusDesc = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTaxRate(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TaxRate = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetThirdItineraryId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.ThirdItineraryId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTicketId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TicketId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTrade(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.Trade = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTradeActionDesc(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TradeActionDesc = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerEmail(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerEmail = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerId(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerJobNo(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerMemberType(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerMemberType = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerMemberTypeName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerMemberTypeName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetTravelerName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetUpgradeCost(v float64) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.UpgradeCost = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetVoucherType(v int32) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetVoucherTypeDesc(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherTypeDesc = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) SetVoyageName(v string) *FlightBillSettlementQueryResponseBodyModuleDataList {
	s.VoyageName = &v
	return s
}

func (s *FlightBillSettlementQueryResponseBodyModuleDataList) Validate() error {
	return dara.Validate(s)
}
