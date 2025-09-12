// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCooperatorFlightBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CooperatorFlightBillSettlementQueryResponseBody
	GetCode() *string
	SetMessage(v string) *CooperatorFlightBillSettlementQueryResponseBody
	GetMessage() *string
	SetModule(v *CooperatorFlightBillSettlementQueryResponseBodyModule) *CooperatorFlightBillSettlementQueryResponseBody
	GetModule() *CooperatorFlightBillSettlementQueryResponseBodyModule
	SetRequestId(v string) *CooperatorFlightBillSettlementQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CooperatorFlightBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CooperatorFlightBillSettlementQueryResponseBody
	GetTraceId() *string
}

type CooperatorFlightBillSettlementQueryResponseBody struct {
	// example:
	//
	// 0
	Code    *string                                                `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  *CooperatorFlightBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// 2103ad1216872266815642815d7e03
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// trace_id
	//
	// example:
	//
	// 213e20c816937929648732715e16f1
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s CooperatorFlightBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CooperatorFlightBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) GetModule() *CooperatorFlightBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) SetCode(v string) *CooperatorFlightBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) SetMessage(v string) *CooperatorFlightBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) SetModule(v *CooperatorFlightBillSettlementQueryResponseBodyModule) *CooperatorFlightBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) SetRequestId(v string) *CooperatorFlightBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) SetSuccess(v bool) *CooperatorFlightBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) SetTraceId(v string) *CooperatorFlightBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CooperatorFlightBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 1
	Category *int32 `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 123
	CorpId *string                                                       `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Items  []*CooperatorFlightBillSettlementQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-11-02
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

func (s CooperatorFlightBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CooperatorFlightBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) GetItems() []*CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	return s.Items
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) SetCategory(v int32) *CooperatorFlightBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) SetCorpId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) SetItems(v []*CooperatorFlightBillSettlementQueryResponseBodyModuleItems) *CooperatorFlightBillSettlementQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *CooperatorFlightBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *CooperatorFlightBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) SetScrollId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) SetTotalSize(v int64) *CooperatorFlightBillSettlementQueryResponseBodyModule {
	s.TotalSize = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type CooperatorFlightBillSettlementQueryResponseBodyModuleItems struct {
	AdjustTime *string `json:"adjust_time,omitempty" xml:"adjust_time,omitempty"`
	// example:
	//
	// 1
	AdvanceDay *int32 `json:"advance_day,omitempty" xml:"advance_day,omitempty"`
	// example:
	//
	// AB
	AirlineCorpCode *string `json:"airline_corp_code,omitempty" xml:"airline_corp_code,omitempty"`
	AirlineCorpName *string `json:"airline_corp_name,omitempty" xml:"airline_corp_name,omitempty"`
	// example:
	//
	// 123aaa
	AlipayId *string `json:"alipay_id,omitempty" xml:"alipay_id,omitempty"`
	// example:
	//
	// a123
	AlipayTradeNo *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	// example:
	//
	// CD
	ApplyArrCityCode *string `json:"apply_arr_city_code,omitempty" xml:"apply_arr_city_code,omitempty"`
	ApplyArrCityName *string `json:"apply_arr_city_name,omitempty" xml:"apply_arr_city_name,omitempty"`
	// example:
	//
	// AB
	ApplyDepCityCode *string `json:"apply_dep_city_code,omitempty" xml:"apply_dep_city_code,omitempty"`
	ApplyDepCityName *string `json:"apply_dep_city_name,omitempty" xml:"apply_dep_city_name,omitempty"`
	ApplyExtendField *string `json:"apply_extend_field,omitempty" xml:"apply_extend_field,omitempty"`
	// example:
	//
	// 1004430880
	ApplyId       *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ApproverEmail *string `json:"approver_email,omitempty" xml:"approver_email,omitempty"`
	ApproverId    *string `json:"approver_id,omitempty" xml:"approver_id,omitempty"`
	ApproverName  *string `json:"approver_name,omitempty" xml:"approver_name,omitempty"`
	// example:
	//
	// CKG
	ArrAirportCode *string `json:"arr_airport_code,omitempty" xml:"arr_airport_code,omitempty"`
	ArrCity        *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// CKG
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// 2023-01-01
	ArrDate    *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	ArrStation *string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	// example:
	//
	// 12:00:00
	ArrTime      *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BaseLocation *string `json:"base_location,omitempty" xml:"base_location,omitempty"`
	// example:
	//
	// 2023-01-01 00:00:00
	BillRecordTime *string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BookChannel    *string `json:"book_channel,omitempty" xml:"book_channel,omitempty"`
	BookMode       *string `json:"book_mode,omitempty" xml:"book_mode,omitempty"`
	// example:
	//
	// 2023-01-01 00:00:00
	BookTime *string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// example:
	//
	// 1234
	BookerId *string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	// example:
	//
	// A1234
	BookerJobNo *string `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName  *string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	// example:
	//
	// 100.0
	BtripCouponFee *float64 `json:"btrip_coupon_fee,omitempty" xml:"btrip_coupon_fee,omitempty"`
	// example:
	//
	// 50.0
	BuildFee           *float64 `json:"build_fee,omitempty" xml:"build_fee,omitempty"`
	BusinessTripResult *string  `json:"business_trip_result,omitempty" xml:"business_trip_result,omitempty"`
	// example:
	//
	// A
	Cabin *string `json:"cabin,omitempty" xml:"cabin,omitempty"`
	// example:
	//
	// YS
	CabinClass        *string `json:"cabin_class,omitempty" xml:"cabin_class,omitempty"`
	CabinClassCode    *string `json:"cabin_class_code,omitempty" xml:"cabin_class_code,omitempty"`
	CapitalDirection  *string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment *string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CategoryDesc      *string `json:"category_desc,omitempty" xml:"category_desc,omitempty"`
	// example:
	//
	// 100.0
	ChangeFee    *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	ChangeResult *string  `json:"change_result,omitempty" xml:"change_result,omitempty"`
	// example:
	//
	// IN240102113438277278
	CooperatorBillCode *string `json:"cooperator_bill_code,omitempty" xml:"cooperator_bill_code,omitempty"`
	CooperatorName     *string `json:"cooperator_name,omitempty" xml:"cooperator_name,omitempty"`
	// example:
	//
	// DF24020163776907739
	CooperatorOrderId *string `json:"cooperator_order_id,omitempty" xml:"cooperator_order_id,omitempty"`
	// example:
	//
	// 100.0
	CorpPayOrderFee *float64 `json:"corp_pay_order_fee,omitempty" xml:"corp_pay_order_fee,omitempty"`
	// example:
	//
	// 100.0
	CorpSettlePrice *float64 `json:"corp_settle_price,omitempty" xml:"corp_settle_price,omitempty"`
	CostCenter      *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// example:
	//
	// cs1
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	CostDepartment   *string `json:"cost_department,omitempty" xml:"cost_department,omitempty"`
	// example:
	//
	// 0.0
	Coupon        *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	CustomContent *string  `json:"custom_content,omitempty" xml:"custom_content,omitempty"`
	// example:
	//
	// JHG
	DepAirportCode *string `json:"dep_airport_code,omitempty" xml:"dep_airport_code,omitempty"`
	// example:
	//
	// TAO
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	Department  *string `json:"department,omitempty" xml:"department,omitempty"`
	// example:
	//
	// 1112
	DepartmentId *string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	DeptCity     *string `json:"dept_city,omitempty" xml:"dept_city,omitempty"`
	// example:
	//
	// 2023-01-01
	DeptDate    *string `json:"dept_date,omitempty" xml:"dept_date,omitempty"`
	DeptStation *string `json:"dept_station,omitempty" xml:"dept_station,omitempty"`
	// example:
	//
	// 09:30:00
	DeptTime *string `json:"dept_time,omitempty" xml:"dept_time,omitempty"`
	// example:
	//
	// 1
	Discount     *string `json:"discount,omitempty" xml:"discount,omitempty"`
	ExceedReason *string `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	FeeType      *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc  *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	// example:
	//
	// CZ3590
	FlightNo       *string `json:"flight_no,omitempty" xml:"flight_no,omitempty"`
	FlightTripType *string `json:"flight_trip_type,omitempty" xml:"flight_trip_type,omitempty"`
	ForeignersTag  *string `json:"foreigners_tag,omitempty" xml:"foreigners_tag,omitempty"`
	// example:
	//
	// 1
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// example:
	//
	// 111
	InsOrderId *string `json:"ins_order_id,omitempty" xml:"ins_order_id,omitempty"`
	// example:
	//
	// 0.0
	InsuranceFee *float64 `json:"insurance_fee,omitempty" xml:"insurance_fee,omitempty"`
	// example:
	//
	// 1234A
	InsuranceNumber *string `json:"insurance_number,omitempty" xml:"insurance_number,omitempty"`
	InvoiceTitle    *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	ItemType        *string `json:"item_type,omitempty" xml:"item_type,omitempty"`
	// example:
	//
	// 6666666666
	ItineraryNum *string `json:"itinerary_num,omitempty" xml:"itinerary_num,omitempty"`
	// example:
	//
	// 100.0
	ItineraryPrice     *float64 `json:"itinerary_price,omitempty" xml:"itinerary_price,omitempty"`
	Location           *string  `json:"location,omitempty" xml:"location,omitempty"`
	MappingCompanyCode *string  `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	// example:
	//
	// 100
	Mileage *int32 `json:"mileage,omitempty" xml:"mileage,omitempty"`
	// example:
	//
	// 2023-01-01 00:00:00
	MostDifferenceDeptTime *string `json:"most_difference_dept_time,omitempty" xml:"most_difference_dept_time,omitempty"`
	// example:
	//
	// 1
	MostDifferenceDiscount *string `json:"most_difference_discount,omitempty" xml:"most_difference_discount,omitempty"`
	// example:
	//
	// 123
	MostDifferenceFlightNo *string `json:"most_difference_flight_no,omitempty" xml:"most_difference_flight_no,omitempty"`
	// example:
	//
	// 100.0
	MostDifferencePrice  *float64 `json:"most_difference_price,omitempty" xml:"most_difference_price,omitempty"`
	MostDifferenceReason *string  `json:"most_difference_reason,omitempty" xml:"most_difference_reason,omitempty"`
	// example:
	//
	// 100.0
	MostPrice *float64 `json:"most_price,omitempty" xml:"most_price,omitempty"`
	// example:
	//
	// 0.0
	NegotiationCouponFee *float64 `json:"negotiation_coupon_fee,omitempty" xml:"negotiation_coupon_fee,omitempty"`
	// example:
	//
	// 30.0
	OilFee *float64 `json:"oil_fee,omitempty" xml:"oil_fee,omitempty"`
	// example:
	//
	// 3137168772101111000
	OrderId         *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderStatusDesc *string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	OrderTicketNo   *string `json:"order_ticket_no,omitempty" xml:"order_ticket_no,omitempty"`
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
	// 100.0
	PersonSettlePrice *float64 `json:"person_settle_price,omitempty" xml:"person_settle_price,omitempty"`
	Position          *string  `json:"position,omitempty" xml:"position,omitempty"`
	PositionLevel     *string  `json:"position_level,omitempty" xml:"position_level,omitempty"`
	PreBookTip        *string  `json:"pre_book_tip,omitempty" xml:"pre_book_tip,omitempty"`
	// example:
	//
	// 60399513
	PrimaryId       *int64  `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProcessorOaCode *string `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// acs
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// example:
	//
	// 100.0
	RefundFee    *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	RefundResult *string  `json:"refund_result,omitempty" xml:"refund_result,omitempty"`
	// example:
	//
	// 20.0
	RefundUpgradeCost *float64 `json:"refund_upgrade_cost,omitempty" xml:"refund_upgrade_cost,omitempty"`
	Remark            *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	RepeatRefund      *string  `json:"repeat_refund,omitempty" xml:"repeat_refund,omitempty"`
	// example:
	//
	// 100.0
	SealPrice *float64 `json:"seal_price,omitempty" xml:"seal_price,omitempty"`
	// example:
	//
	// 0.0
	ServiceFee     *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettleTypeDesc *string  `json:"settle_type_desc,omitempty" xml:"settle_type_desc,omitempty"`
	// example:
	//
	// 200.0
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
	// example:
	//
	// 0
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// 169551103
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
	// example:
	//
	// 123-2345
	TicketId *string `json:"ticket_id,omitempty" xml:"ticket_id,omitempty"`
	// example:
	//
	// 1
	Trade           *string `json:"trade,omitempty" xml:"trade,omitempty"`
	TradeActionDesc *string `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
	TravelerEmail   *string `json:"traveler_email,omitempty" xml:"traveler_email,omitempty"`
	// example:
	//
	// A1234
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
	// 100.0
	UpgradeCost *float64 `json:"upgrade_cost,omitempty" xml:"upgrade_cost,omitempty"`
	// example:
	//
	// 11
	VoucherType     *int32  `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	VoucherTypeDesc *string `json:"voucher_type_desc,omitempty" xml:"voucher_type_desc,omitempty"`
	VoyageName      *string `json:"voyage_name,omitempty" xml:"voyage_name,omitempty"`
}

func (s CooperatorFlightBillSettlementQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetAdvanceDay() *int32 {
	return s.AdvanceDay
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetAirlineCorpCode() *string {
	return s.AirlineCorpCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetAirlineCorpName() *string {
	return s.AirlineCorpName
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetAlipayId() *string {
	return s.AlipayId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetApplyArrCityCode() *string {
	return s.ApplyArrCityCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetApplyArrCityName() *string {
	return s.ApplyArrCityName
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetApplyDepCityCode() *string {
	return s.ApplyDepCityCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetApplyDepCityName() *string {
	return s.ApplyDepCityName
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetApplyId() *string {
	return s.ApplyId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetApproverEmail() *string {
	return s.ApproverEmail
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetApproverId() *string {
	return s.ApproverId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetApproverName() *string {
	return s.ApproverName
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetArrAirportCode() *string {
	return s.ArrAirportCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetArrCity() *string {
	return s.ArrCity
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetArrDate() *string {
	return s.ArrDate
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetArrStation() *string {
	return s.ArrStation
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetArrTime() *string {
	return s.ArrTime
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetBaseLocation() *string {
	return s.BaseLocation
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetBookChannel() *string {
	return s.BookChannel
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetBookMode() *string {
	return s.BookMode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetBookTime() *string {
	return s.BookTime
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetBookerId() *string {
	return s.BookerId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetBookerName() *string {
	return s.BookerName
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetBtripCouponFee() *float64 {
	return s.BtripCouponFee
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetBuildFee() *float64 {
	return s.BuildFee
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetBusinessTripResult() *string {
	return s.BusinessTripResult
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCabin() *string {
	return s.Cabin
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCabinClass() *string {
	return s.CabinClass
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCabinClassCode() *string {
	return s.CabinClassCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetChangeFee() *float64 {
	return s.ChangeFee
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetChangeResult() *string {
	return s.ChangeResult
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCooperatorBillCode() *string {
	return s.CooperatorBillCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCooperatorName() *string {
	return s.CooperatorName
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCooperatorOrderId() *string {
	return s.CooperatorOrderId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCorpPayOrderFee() *float64 {
	return s.CorpPayOrderFee
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCorpSettlePrice() *float64 {
	return s.CorpSettlePrice
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCostCenter() *string {
	return s.CostCenter
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCoupon() *float64 {
	return s.Coupon
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetCustomContent() *string {
	return s.CustomContent
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetDepAirportCode() *string {
	return s.DepAirportCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetDepartment() *string {
	return s.Department
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetDeptCity() *string {
	return s.DeptCity
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetDeptDate() *string {
	return s.DeptDate
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetDeptStation() *string {
	return s.DeptStation
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetDeptTime() *string {
	return s.DeptTime
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetDiscount() *string {
	return s.Discount
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetExceedReason() *string {
	return s.ExceedReason
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetFeeType() *string {
	return s.FeeType
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetFlightNo() *string {
	return s.FlightNo
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetFlightTripType() *string {
	return s.FlightTripType
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetForeignersTag() *string {
	return s.ForeignersTag
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetIndex() *string {
	return s.Index
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetInsOrderId() *string {
	return s.InsOrderId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetInsuranceFee() *float64 {
	return s.InsuranceFee
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetInsuranceNumber() *string {
	return s.InsuranceNumber
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetItemType() *string {
	return s.ItemType
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetItineraryNum() *string {
	return s.ItineraryNum
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetItineraryPrice() *float64 {
	return s.ItineraryPrice
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetLocation() *string {
	return s.Location
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetMileage() *int32 {
	return s.Mileage
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetMostDifferenceDeptTime() *string {
	return s.MostDifferenceDeptTime
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetMostDifferenceDiscount() *string {
	return s.MostDifferenceDiscount
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetMostDifferenceFlightNo() *string {
	return s.MostDifferenceFlightNo
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetMostDifferencePrice() *float64 {
	return s.MostDifferencePrice
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetMostDifferenceReason() *string {
	return s.MostDifferenceReason
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetMostPrice() *float64 {
	return s.MostPrice
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetNegotiationCouponFee() *float64 {
	return s.NegotiationCouponFee
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetOilFee() *float64 {
	return s.OilFee
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetOrderId() *string {
	return s.OrderId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetOrderTicketNo() *string {
	return s.OrderTicketNo
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetOverApplyId() *string {
	return s.OverApplyId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetPersonSettlePrice() *float64 {
	return s.PersonSettlePrice
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetPosition() *string {
	return s.Position
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetPreBookTip() *string {
	return s.PreBookTip
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetProjectName() *string {
	return s.ProjectName
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetRefundResult() *string {
	return s.RefundResult
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetRefundUpgradeCost() *float64 {
	return s.RefundUpgradeCost
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetRemark() *string {
	return s.Remark
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetRepeatRefund() *string {
	return s.RepeatRefund
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetSealPrice() *float64 {
	return s.SealPrice
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetSettlementGrantFee() *float64 {
	return s.SettlementGrantFee
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetSettlementType() *string {
	return s.SettlementType
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetSio() *string {
	return s.Sio
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetStatus() *int32 {
	return s.Status
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetTaxRate() *string {
	return s.TaxRate
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetThirdInvoiceId() *string {
	return s.ThirdInvoiceId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetThirdItineraryId() *string {
	return s.ThirdItineraryId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetTicketId() *string {
	return s.TicketId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetTrade() *string {
	return s.Trade
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetTravelerEmail() *string {
	return s.TravelerEmail
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetTravelerId() *string {
	return s.TravelerId
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberTypeName() *string {
	return s.TravelerMemberTypeName
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetTravelerName() *string {
	return s.TravelerName
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetUpgradeCost() *float64 {
	return s.UpgradeCost
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) GetVoyageName() *string {
	return s.VoyageName
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetAdjustTime(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.AdjustTime = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetAdvanceDay(v int32) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.AdvanceDay = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetAirlineCorpCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.AirlineCorpCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetAirlineCorpName(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.AirlineCorpName = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetAlipayId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.AlipayId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetAlipayTradeNo(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.AlipayTradeNo = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetApplyArrCityCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ApplyArrCityCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetApplyArrCityName(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ApplyArrCityName = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetApplyDepCityCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ApplyDepCityCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetApplyDepCityName(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ApplyDepCityName = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetApplyExtendField(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ApplyExtendField = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetApplyId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ApplyId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetApproverEmail(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ApproverEmail = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetApproverId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ApproverId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetApproverName(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ApproverName = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetArrAirportCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ArrAirportCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetArrCity(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ArrCity = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetArrCityCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ArrCityCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetArrDate(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ArrDate = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetArrStation(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ArrStation = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetArrTime(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ArrTime = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetBaseLocation(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.BaseLocation = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetBillRecordTime(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.BillRecordTime = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetBookChannel(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.BookChannel = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetBookMode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.BookMode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetBookTime(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.BookTime = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetBookerId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.BookerId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetBookerJobNo(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.BookerJobNo = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetBookerName(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.BookerName = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetBtripCouponFee(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.BtripCouponFee = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetBuildFee(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.BuildFee = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetBusinessTripResult(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.BusinessTripResult = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCabin(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Cabin = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCabinClass(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CabinClass = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCabinClassCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CabinClassCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCapitalDirection(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CapitalDirection = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCascadeDepartment(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CascadeDepartment = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCategoryDesc(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CategoryDesc = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetChangeFee(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ChangeFee = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetChangeResult(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ChangeResult = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCooperatorBillCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CooperatorBillCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCooperatorName(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CooperatorName = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCooperatorOrderId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CooperatorOrderId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCorpPayOrderFee(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CorpPayOrderFee = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCorpSettlePrice(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CorpSettlePrice = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCostCenter(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CostCenter = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCostCenterNumber(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CostCenterNumber = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCostDepartment(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CostDepartment = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCoupon(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Coupon = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetCustomContent(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.CustomContent = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetDepAirportCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.DepAirportCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetDepCityCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.DepCityCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetDepartment(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Department = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetDepartmentId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.DepartmentId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetDeptCity(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.DeptCity = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetDeptDate(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.DeptDate = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetDeptStation(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.DeptStation = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetDeptTime(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.DeptTime = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetDiscount(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Discount = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetExceedReason(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ExceedReason = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetFeeType(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.FeeType = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetFeeTypeDesc(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.FeeTypeDesc = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetFlightNo(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.FlightNo = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetFlightTripType(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.FlightTripType = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetForeignersTag(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ForeignersTag = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetIndex(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Index = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetInsOrderId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.InsOrderId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetInsuranceFee(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.InsuranceFee = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetInsuranceNumber(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.InsuranceNumber = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetInvoiceTitle(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.InvoiceTitle = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetItemType(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ItemType = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetItineraryNum(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ItineraryNum = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetItineraryPrice(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ItineraryPrice = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetLocation(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Location = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetMappingCompanyCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.MappingCompanyCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetMileage(v int32) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Mileage = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetMostDifferenceDeptTime(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.MostDifferenceDeptTime = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetMostDifferenceDiscount(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.MostDifferenceDiscount = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetMostDifferenceFlightNo(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.MostDifferenceFlightNo = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetMostDifferencePrice(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.MostDifferencePrice = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetMostDifferenceReason(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.MostDifferenceReason = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetMostPrice(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.MostPrice = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetNegotiationCouponFee(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.NegotiationCouponFee = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetOilFee(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.OilFee = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetOrderId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.OrderId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetOrderStatusDesc(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.OrderStatusDesc = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetOrderTicketNo(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.OrderTicketNo = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetOverApplyId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.OverApplyId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetPaymentDepartmentId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.PaymentDepartmentId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetPaymentDepartmentName(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.PaymentDepartmentName = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetPersonSettlePrice(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.PersonSettlePrice = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetPosition(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Position = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetPositionLevel(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.PositionLevel = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetPreBookTip(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.PreBookTip = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetPrimaryId(v int64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.PrimaryId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetProcessorOaCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ProcessorOaCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetProjectCode(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ProjectCode = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetProjectName(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ProjectName = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetRefundFee(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.RefundFee = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetRefundResult(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.RefundResult = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetRefundUpgradeCost(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.RefundUpgradeCost = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetRemark(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Remark = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetRepeatRefund(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.RepeatRefund = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetSealPrice(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.SealPrice = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetServiceFee(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ServiceFee = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetSettleTypeDesc(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.SettleTypeDesc = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetSettlementFee(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.SettlementFee = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetSettlementGrantFee(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.SettlementGrantFee = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetSettlementTime(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.SettlementTime = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetSettlementType(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.SettlementType = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetSio(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Sio = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetStatus(v int32) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Status = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetStatusDesc(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.StatusDesc = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetSubOrderId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.SubOrderId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetTaxRate(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.TaxRate = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetThirdInvoiceId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ThirdInvoiceId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetThirdItineraryId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.ThirdItineraryId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetTicketId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.TicketId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetTrade(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.Trade = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetTradeActionDesc(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.TradeActionDesc = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetTravelerEmail(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.TravelerEmail = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetTravelerId(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.TravelerId = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetTravelerJobNo(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.TravelerJobNo = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberType(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberType = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberTypeName(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberTypeName = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetTravelerName(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.TravelerName = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetUpgradeCost(v float64) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.UpgradeCost = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetVoucherType(v int32) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.VoucherType = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetVoucherTypeDesc(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.VoucherTypeDesc = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) SetVoyageName(v string) *CooperatorFlightBillSettlementQueryResponseBodyModuleItems {
	s.VoyageName = &v
	return s
}

func (s *CooperatorFlightBillSettlementQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}
