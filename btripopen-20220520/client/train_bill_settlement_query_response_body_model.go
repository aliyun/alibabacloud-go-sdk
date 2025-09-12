// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrainBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TrainBillSettlementQueryResponseBody
	GetCode() *string
	SetMessage(v string) *TrainBillSettlementQueryResponseBody
	GetMessage() *string
	SetModule(v *TrainBillSettlementQueryResponseBodyModule) *TrainBillSettlementQueryResponseBody
	GetModule() *TrainBillSettlementQueryResponseBodyModule
	SetRequestId(v string) *TrainBillSettlementQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *TrainBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *TrainBillSettlementQueryResponseBody
	GetTraceId() *string
}

type TrainBillSettlementQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                     `json:"message,omitempty" xml:"message,omitempty"`
	Module  *TrainBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s TrainBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TrainBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *TrainBillSettlementQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TrainBillSettlementQueryResponseBody) GetModule() *TrainBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *TrainBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TrainBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *TrainBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *TrainBillSettlementQueryResponseBody) SetCode(v string) *TrainBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBody) SetMessage(v string) *TrainBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBody) SetModule(v *TrainBillSettlementQueryResponseBodyModule) *TrainBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *TrainBillSettlementQueryResponseBody) SetRequestId(v string) *TrainBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBody) SetSuccess(v bool) *TrainBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBody) SetTraceId(v string) *TrainBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type TrainBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 6
	Category *int32                                                `json:"category,omitempty" xml:"category,omitempty"`
	CorpId   *string                                               `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	DataList []*TrainBillSettlementQueryResponseBodyModuleDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
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
	// 2694
	TotalNum *int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

func (s TrainBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s TrainBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *TrainBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *TrainBillSettlementQueryResponseBodyModule) GetDataList() []*TrainBillSettlementQueryResponseBodyModuleDataList {
	return s.DataList
}

func (s *TrainBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *TrainBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *TrainBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *TrainBillSettlementQueryResponseBodyModule) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetCategory(v int32) *TrainBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetCorpId(v string) *TrainBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetDataList(v []*TrainBillSettlementQueryResponseBodyModuleDataList) *TrainBillSettlementQueryResponseBodyModule {
	s.DataList = v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *TrainBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *TrainBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetScrollId(v string) *TrainBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *TrainBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type TrainBillSettlementQueryResponseBodyModuleDataList struct {
	AdjustTime *string `json:"adjust_time,omitempty" xml:"adjust_time,omitempty"`
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
	ArrCityCode   *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCityName   *string `json:"arr_city_name,omitempty" xml:"arr_city_name,omitempty"`
	// example:
	//
	// 2021-10-13
	ArrDate                *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	ArrStation             *string `json:"arr_station,omitempty" xml:"arr_station,omitempty"`
	ArrStationLocation     *string `json:"arr_station_location,omitempty" xml:"arr_station_location,omitempty"`
	ArrStationLocationCode *string `json:"arr_station_location_code,omitempty" xml:"arr_station_location_code,omitempty"`
	// example:
	//
	// 12:30
	ArrTime      *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BaseLocation *string `json:"base_location,omitempty" xml:"base_location,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	BillRecordTime *string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	// example:
	//
	// 2021-10-08 23:38:55
	BookTime *string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// example:
	//
	// al_xinuan.zsy
	BookerId           *string  `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	BookerJobNo        *string  `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName         *string  `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	BusinessTripResult *string  `json:"business_trip_result,omitempty" xml:"business_trip_result,omitempty"`
	CabinMaxPrice      *float64 `json:"cabin_max_price,omitempty" xml:"cabin_max_price,omitempty"`
	// example:
	//
	// 1
	CapitalDirection  *string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment *string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CategoryDesc      *string `json:"category_desc,omitempty" xml:"category_desc,omitempty"`
	ChangeAffiliateNo *string `json:"change_affiliate_no,omitempty" xml:"change_affiliate_no,omitempty"`
	ChangeApplyId     *string `json:"change_apply_id,omitempty" xml:"change_apply_id,omitempty"`
	// example:
	//
	// 23.0
	ChangeFee    *float64 `json:"change_fee,omitempty" xml:"change_fee,omitempty"`
	ChangeResult *string  `json:"change_result,omitempty" xml:"change_result,omitempty"`
	CoachNo      *string  `json:"coach_no,omitempty" xml:"coach_no,omitempty"`
	CostCenter   *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// example:
	//
	// T85
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	CostDepartment   *string `json:"cost_department,omitempty" xml:"cost_department,omitempty"`
	// example:
	//
	// 0
	Coupon                 *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	CustomContent          *string  `json:"custom_content,omitempty" xml:"custom_content,omitempty"`
	DeductibleTax          *float64 `json:"deductible_tax,omitempty" xml:"deductible_tax,omitempty"`
	DepCityCode            *string  `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCityName            *string  `json:"dep_city_name,omitempty" xml:"dep_city_name,omitempty"`
	DepStationLocation     *string  `json:"dep_station_location,omitempty" xml:"dep_station_location,omitempty"`
	DepStationLocationCode *string  `json:"dep_station_location_code,omitempty" xml:"dep_station_location_code,omitempty"`
	Department             *string  `json:"department,omitempty" xml:"department,omitempty"`
	DepartmentId           *string  `json:"department_id,omitempty" xml:"department_id,omitempty"`
	// example:
	//
	// 2021-10-14
	DeptDate    *string `json:"dept_date,omitempty" xml:"dept_date,omitempty"`
	DeptStation *string `json:"dept_station,omitempty" xml:"dept_station,omitempty"`
	// example:
	//
	// 09:44
	DeptTime     *string `json:"dept_time,omitempty" xml:"dept_time,omitempty"`
	ExceedReason *string `json:"exceed_reason,omitempty" xml:"exceed_reason,omitempty"`
	// example:
	//
	// 6001
	FeeType       *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc   *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	ForeignersTag *string `json:"foreigners_tag,omitempty" xml:"foreigners_tag,omitempty"`
	// example:
	//
	// 4740293
	Index              *string `json:"index,omitempty" xml:"index,omitempty"`
	InvoiceTitle       *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	IsTransferOrder    *string `json:"is_transfer_order,omitempty" xml:"is_transfer_order,omitempty"`
	Location           *string `json:"location,omitempty" xml:"location,omitempty"`
	LongTicketNo       *string `json:"long_ticket_no,omitempty" xml:"long_ticket_no,omitempty"`
	MappingCompanyCode *string `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	MaxCabin           *string `json:"max_cabin,omitempty" xml:"max_cabin,omitempty"`
	// example:
	//
	// 23432692343243432
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 350
	OrderPrice    *float64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	OrderTicketNo *string  `json:"order_ticket_no,omitempty" xml:"order_ticket_no,omitempty"`
	// example:
	//
	// 534545345
	OverApplyId           *string `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	PaymentDepartmentId   *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	Position              *string `json:"position,omitempty" xml:"position,omitempty"`
	PositionLevel         *string `json:"position_level,omitempty" xml:"position_level,omitempty"`
	// example:
	//
	// 4740293
	PrimaryId        *int64   `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	PrintTicketPrice *float64 `json:"print_ticket_price,omitempty" xml:"print_ticket_price,omitempty"`
	ProcessorOaCode  *string  `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// 2345235435
	ProjectCode       *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName       *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	RefundAffiliateNo *string `json:"refund_affiliate_no,omitempty" xml:"refund_affiliate_no,omitempty"`
	RefundApplyId     *string `json:"refund_apply_id,omitempty" xml:"refund_apply_id,omitempty"`
	// example:
	//
	// 0
	RefundFee    *float64 `json:"refund_fee,omitempty" xml:"refund_fee,omitempty"`
	RefundReason *string  `json:"refund_reason,omitempty" xml:"refund_reason,omitempty"`
	Remark       *string  `json:"remark,omitempty" xml:"remark,omitempty"`
	ReserveMode  *string  `json:"reserve_mode,omitempty" xml:"reserve_mode,omitempty"`
	RunTime      *string  `json:"run_time,omitempty" xml:"run_time,omitempty"`
	SceneId      *string  `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	SceneName    *string  `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	// example:
	//
	// 004F
	SeatNo   *string `json:"seat_no,omitempty" xml:"seat_no,omitempty"`
	SeatType *string `json:"seat_type,omitempty" xml:"seat_type,omitempty"`
	// example:
	//
	// 23.0
	ServiceFee     *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettleTypeDesc *string  `json:"settle_type_desc,omitempty" xml:"settle_type_desc,omitempty"`
	// example:
	//
	// 350
	SettlementFee *float64 `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	// example:
	//
	// 4.56
	SettlementGrantFee *float64 `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	// example:
	//
	// 2021-10-08 23:39:01
	SettlementTime *string `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	// example:
	//
	// 2
	SettlementType  *string  `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	ShortTicketNo   *string  `json:"short_ticket_no,omitempty" xml:"short_ticket_no,omitempty"`
	Sio             *string  `json:"sio,omitempty" xml:"sio,omitempty"`
	SpeedPackageFee *float64 `json:"speed_package_fee,omitempty" xml:"speed_package_fee,omitempty"`
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
	TaxRate            *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	ThirdItineraryId   *string `json:"third_itinerary_id,omitempty" xml:"third_itinerary_id,omitempty"`
	TicketCorpPayPrice *string `json:"ticket_corp_pay_price,omitempty" xml:"ticket_corp_pay_price,omitempty"`
	// example:
	//
	// 2115242342342424067354
	TicketNo             *string `json:"ticket_no,omitempty" xml:"ticket_no,omitempty"`
	TicketPersonPayPrice *string `json:"ticket_person_pay_price,omitempty" xml:"ticket_person_pay_price,omitempty"`
	// example:
	//
	// 350
	TicketPrice     *float64 `json:"ticket_price,omitempty" xml:"ticket_price,omitempty"`
	TradeActionDesc *string  `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
	// example:
	//
	// G906
	TrainNo                *string `json:"train_no,omitempty" xml:"train_no,omitempty"`
	TrainType              *string `json:"train_type,omitempty" xml:"train_type,omitempty"`
	TravelerEmail          *string `json:"traveler_email,omitempty" xml:"traveler_email,omitempty"`
	TravelerId             *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	TravelerJobNo          *string `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerMemberType     *string `json:"traveler_member_type,omitempty" xml:"traveler_member_type,omitempty"`
	TravelerMemberTypeName *string `json:"traveler_member_type_name,omitempty" xml:"traveler_member_type_name,omitempty"`
	TravelerName           *string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	VerifyFailedReason     *string `json:"verify_failed_reason,omitempty" xml:"verify_failed_reason,omitempty"`
	VerifyStatus           *int32  `json:"verify_status,omitempty" xml:"verify_status,omitempty"`
	// example:
	//
	// 11
	VoucherType     *int32  `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	VoucherTypeDesc *string `json:"voucher_type_desc,omitempty" xml:"voucher_type_desc,omitempty"`
}

func (s TrainBillSettlementQueryResponseBodyModuleDataList) String() string {
	return dara.Prettify(s)
}

func (s TrainBillSettlementQueryResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityCode() *string {
	return s.ApplyArrCityCode
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityName() *string {
	return s.ApplyArrCityName
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityCode() *string {
	return s.ApplyDepCityCode
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityName() *string {
	return s.ApplyDepCityName
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetApplyId() *string {
	return s.ApplyId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetApproverEmail() *string {
	return s.ApproverEmail
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetApproverId() *string {
	return s.ApproverId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetApproverName() *string {
	return s.ApproverName
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetArrCityName() *string {
	return s.ArrCityName
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetArrDate() *string {
	return s.ArrDate
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetArrStation() *string {
	return s.ArrStation
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetArrStationLocation() *string {
	return s.ArrStationLocation
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetArrStationLocationCode() *string {
	return s.ArrStationLocationCode
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetArrTime() *string {
	return s.ArrTime
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetBaseLocation() *string {
	return s.BaseLocation
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetBookTime() *string {
	return s.BookTime
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetBookerId() *string {
	return s.BookerId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetBookerName() *string {
	return s.BookerName
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetBusinessTripResult() *string {
	return s.BusinessTripResult
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetCabinMaxPrice() *float64 {
	return s.CabinMaxPrice
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetChangeAffiliateNo() *string {
	return s.ChangeAffiliateNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetChangeApplyId() *string {
	return s.ChangeApplyId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetChangeFee() *float64 {
	return s.ChangeFee
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetChangeResult() *string {
	return s.ChangeResult
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetCoachNo() *string {
	return s.CoachNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetCostCenter() *string {
	return s.CostCenter
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetCoupon() *float64 {
	return s.Coupon
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetCustomContent() *string {
	return s.CustomContent
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetDeductibleTax() *float64 {
	return s.DeductibleTax
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetDepCityName() *string {
	return s.DepCityName
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetDepStationLocation() *string {
	return s.DepStationLocation
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetDepStationLocationCode() *string {
	return s.DepStationLocationCode
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetDepartment() *string {
	return s.Department
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetDeptDate() *string {
	return s.DeptDate
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetDeptStation() *string {
	return s.DeptStation
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetDeptTime() *string {
	return s.DeptTime
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetExceedReason() *string {
	return s.ExceedReason
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetFeeType() *string {
	return s.FeeType
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetForeignersTag() *string {
	return s.ForeignersTag
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetIndex() *string {
	return s.Index
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetIsTransferOrder() *string {
	return s.IsTransferOrder
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetLocation() *string {
	return s.Location
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetLongTicketNo() *string {
	return s.LongTicketNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetMaxCabin() *string {
	return s.MaxCabin
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetOrderId() *string {
	return s.OrderId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetOrderPrice() *float64 {
	return s.OrderPrice
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetOrderTicketNo() *string {
	return s.OrderTicketNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetOverApplyId() *string {
	return s.OverApplyId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetPosition() *string {
	return s.Position
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetPrintTicketPrice() *float64 {
	return s.PrintTicketPrice
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetProjectName() *string {
	return s.ProjectName
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetRefundAffiliateNo() *string {
	return s.RefundAffiliateNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetRefundApplyId() *string {
	return s.RefundApplyId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetRefundFee() *float64 {
	return s.RefundFee
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetRefundReason() *string {
	return s.RefundReason
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetRemark() *string {
	return s.Remark
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetReserveMode() *string {
	return s.ReserveMode
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetRunTime() *string {
	return s.RunTime
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetSceneId() *string {
	return s.SceneId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetSceneName() *string {
	return s.SceneName
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetSeatNo() *string {
	return s.SeatNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetSeatType() *string {
	return s.SeatType
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetSettlementGrantFee() *float64 {
	return s.SettlementGrantFee
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetSettlementType() *string {
	return s.SettlementType
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetShortTicketNo() *string {
	return s.ShortTicketNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetSio() *string {
	return s.Sio
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetSpeedPackageFee() *float64 {
	return s.SpeedPackageFee
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetStatus() *int32 {
	return s.Status
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTaxRate() *string {
	return s.TaxRate
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetThirdItineraryId() *string {
	return s.ThirdItineraryId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTicketCorpPayPrice() *string {
	return s.TicketCorpPayPrice
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTicketNo() *string {
	return s.TicketNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTicketPersonPayPrice() *string {
	return s.TicketPersonPayPrice
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTicketPrice() *float64 {
	return s.TicketPrice
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTrainNo() *string {
	return s.TrainNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTrainType() *string {
	return s.TrainType
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTravelerEmail() *string {
	return s.TravelerEmail
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTravelerId() *string {
	return s.TravelerId
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTravelerMemberTypeName() *string {
	return s.TravelerMemberTypeName
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetTravelerName() *string {
	return s.TravelerName
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetVerifyFailedReason() *string {
	return s.VerifyFailedReason
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetVerifyStatus() *int32 {
	return s.VerifyStatus
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetAdjustTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.AdjustTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetAlipayTradeNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityCode(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityCode(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetApplyExtendField(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyExtendField = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetApplyId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetApproverEmail(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverEmail = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetApproverId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetApproverName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ApproverName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetArrCityCode(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCityCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetArrCityName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCityName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetArrDate(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ArrDate = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetArrStation(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ArrStation = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetArrStationLocation(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ArrStationLocation = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetArrStationLocationCode(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ArrStationLocationCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetArrTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ArrTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBaseLocation(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BaseLocation = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBillRecordTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BillRecordTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBookTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBookerId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBookerJobNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBookerName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetBusinessTripResult(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.BusinessTripResult = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCabinMaxPrice(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CabinMaxPrice = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCapitalDirection(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCascadeDepartment(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCategoryDesc(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CategoryDesc = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetChangeAffiliateNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeAffiliateNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetChangeApplyId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeApplyId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetChangeFee(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeFee = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetChangeResult(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ChangeResult = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCoachNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CoachNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCostCenter(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCostCenterNumber(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCostDepartment(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CostDepartment = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCoupon(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Coupon = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetCustomContent(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.CustomContent = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDeductibleTax(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DeductibleTax = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDepCityCode(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DepCityCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDepCityName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DepCityName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDepStationLocation(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DepStationLocation = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDepStationLocationCode(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DepStationLocationCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDepartment(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDepartmentId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDeptDate(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DeptDate = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDeptStation(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DeptStation = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetDeptTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.DeptTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetExceedReason(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ExceedReason = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetFeeType(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetFeeTypeDesc(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.FeeTypeDesc = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetForeignersTag(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ForeignersTag = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetIndex(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetInvoiceTitle(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetIsTransferOrder(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.IsTransferOrder = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetLocation(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Location = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetLongTicketNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.LongTicketNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetMappingCompanyCode(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.MappingCompanyCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetMaxCabin(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.MaxCabin = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetOrderId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetOrderPrice(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.OrderPrice = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetOrderTicketNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.OrderTicketNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetOverApplyId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetPosition(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Position = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetPositionLevel(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.PositionLevel = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetPrimaryId(v int64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetPrintTicketPrice(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.PrintTicketPrice = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetProcessorOaCode(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ProcessorOaCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetProjectCode(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetProjectName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetRefundAffiliateNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.RefundAffiliateNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetRefundApplyId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.RefundApplyId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetRefundFee(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.RefundFee = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetRefundReason(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.RefundReason = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetRemark(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Remark = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetReserveMode(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ReserveMode = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetRunTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.RunTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSceneId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SceneId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSceneName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SceneName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSeatNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SeatNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSeatType(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SeatType = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetServiceFee(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSettleTypeDesc(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SettleTypeDesc = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSettlementFee(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSettlementGrantFee(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementGrantFee = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSettlementTime(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSettlementType(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetShortTicketNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ShortTicketNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSio(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Sio = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetSpeedPackageFee(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.SpeedPackageFee = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetStatus(v int32) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetStatusDesc(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.StatusDesc = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTaxRate(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TaxRate = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetThirdItineraryId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.ThirdItineraryId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTicketCorpPayPrice(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TicketCorpPayPrice = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTicketNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TicketNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTicketPersonPayPrice(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TicketPersonPayPrice = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTicketPrice(v float64) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TicketPrice = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTradeActionDesc(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TradeActionDesc = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTrainNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TrainNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTrainType(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TrainType = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTravelerEmail(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerEmail = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTravelerId(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTravelerJobNo(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTravelerMemberType(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerMemberType = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTravelerMemberTypeName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerMemberTypeName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetTravelerName(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetVerifyFailedReason(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.VerifyFailedReason = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetVerifyStatus(v int32) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.VerifyStatus = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetVoucherType(v int32) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) SetVoucherTypeDesc(v string) *TrainBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherTypeDesc = &v
	return s
}

func (s *TrainBillSettlementQueryResponseBodyModuleDataList) Validate() error {
	return dara.Validate(s)
}
