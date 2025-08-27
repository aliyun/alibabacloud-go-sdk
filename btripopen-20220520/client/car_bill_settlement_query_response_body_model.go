// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCarBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CarBillSettlementQueryResponseBody
	GetCode() *string
	SetMessage(v string) *CarBillSettlementQueryResponseBody
	GetMessage() *string
	SetModule(v *CarBillSettlementQueryResponseBodyModule) *CarBillSettlementQueryResponseBody
	GetModule() *CarBillSettlementQueryResponseBodyModule
	SetRequestId(v string) *CarBillSettlementQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CarBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CarBillSettlementQueryResponseBody
	GetTraceId() *string
}

type CarBillSettlementQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	Module  *CarBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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

func (s CarBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CarBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *CarBillSettlementQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CarBillSettlementQueryResponseBody) GetModule() *CarBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *CarBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CarBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CarBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CarBillSettlementQueryResponseBody) SetCode(v string) *CarBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *CarBillSettlementQueryResponseBody) SetMessage(v string) *CarBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *CarBillSettlementQueryResponseBody) SetModule(v *CarBillSettlementQueryResponseBodyModule) *CarBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *CarBillSettlementQueryResponseBody) SetRequestId(v string) *CarBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBody) SetSuccess(v bool) *CarBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *CarBillSettlementQueryResponseBody) SetTraceId(v string) *CarBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type CarBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 4
	Category *int32                                              `json:"category,omitempty" xml:"category,omitempty"`
	CorpId   *string                                             `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	DataList []*CarBillSettlementQueryResponseBodyModuleDataList `json:"data_list,omitempty" xml:"data_list,omitempty" type:"Repeated"`
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
	// 2695
	TotalNum *int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

func (s CarBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s CarBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *CarBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *CarBillSettlementQueryResponseBodyModule) GetDataList() []*CarBillSettlementQueryResponseBodyModuleDataList {
	return s.DataList
}

func (s *CarBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *CarBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *CarBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *CarBillSettlementQueryResponseBodyModule) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *CarBillSettlementQueryResponseBodyModule) SetCategory(v int32) *CarBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) SetCorpId(v string) *CarBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) SetDataList(v []*CarBillSettlementQueryResponseBodyModuleDataList) *CarBillSettlementQueryResponseBodyModule {
	s.DataList = v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *CarBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *CarBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) SetScrollId(v string) *CarBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *CarBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type CarBillSettlementQueryResponseBodyModuleDataList struct {
	AdjustTime *string `json:"adjust_time,omitempty" xml:"adjust_time,omitempty"`
	// example:
	//
	// 34534543545345
	AlipayTradeNo    *string `json:"alipay_trade_no,omitempty" xml:"alipay_trade_no,omitempty"`
	ApplyArrCityCode *string `json:"apply_arr_city_code,omitempty" xml:"apply_arr_city_code,omitempty"`
	ApplyArrCityName *string `json:"apply_arr_city_name,omitempty" xml:"apply_arr_city_name,omitempty"`
	ApplyDepCityCode *string `json:"apply_dep_city_code,omitempty" xml:"apply_dep_city_code,omitempty"`
	ApplyDepCityName *string `json:"apply_dep_city_name,omitempty" xml:"apply_dep_city_name,omitempty"`
	// 审批扩展自定义字段
	ApplyExtendField *string `json:"apply_extend_field,omitempty" xml:"apply_extend_field,omitempty"`
	ApplyId          *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	ArrCity          *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	ArrCityCode      *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	// example:
	//
	// 2022-07-02
	ArrDate     *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	ArrLocation *string `json:"arr_location,omitempty" xml:"arr_location,omitempty"`
	// example:
	//
	// 13:51:43
	ArrTime      *string `json:"arr_time,omitempty" xml:"arr_time,omitempty"`
	BaseLocation *string `json:"base_location,omitempty" xml:"base_location,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	BillRecordTime *string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BillingEntity  *string `json:"billing_entity,omitempty" xml:"billing_entity,omitempty"`
	BookModel      *string `json:"book_model,omitempty" xml:"book_model,omitempty"`
	// example:
	//
	// 2022-05-15 22:27:00
	BookTime *string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	BookerId *string `json:"booker_id,omitempty" xml:"booker_id,omitempty"`
	// example:
	//
	// 70022164
	BookerJobNo      *string `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName       *string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	BusinessCategory *string `json:"business_category,omitempty" xml:"business_category,omitempty"`
	// example:
	//
	// 1
	CapitalDirection  *string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CarLevel          *string `json:"car_level,omitempty" xml:"car_level,omitempty"`
	CascadeDepartment *string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CategoryDesc      *string `json:"category_desc,omitempty" xml:"category_desc,omitempty"`
	CostCenter        *string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// example:
	//
	// 2391-CN90.150
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	CostDepartment   *string `json:"cost_department,omitempty" xml:"cost_department,omitempty"`
	// example:
	//
	// 1
	Coupon *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	// example:
	//
	// 12.7
	CouponPrice   *float64 `json:"coupon_price,omitempty" xml:"coupon_price,omitempty"`
	CustomContent *string  `json:"custom_content,omitempty" xml:"custom_content,omitempty"`
	DeductibleTax *float64 `json:"deductible_tax,omitempty" xml:"deductible_tax,omitempty"`
	DepCityCode   *string  `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	Department    *string  `json:"department,omitempty" xml:"department,omitempty"`
	DepartmentId  *string  `json:"department_id,omitempty" xml:"department_id,omitempty"`
	DeptCity      *string  `json:"dept_city,omitempty" xml:"dept_city,omitempty"`
	// example:
	//
	// 2021-10-13
	DeptDate     *string `json:"dept_date,omitempty" xml:"dept_date,omitempty"`
	DeptLocation *string `json:"dept_location,omitempty" xml:"dept_location,omitempty"`
	// example:
	//
	// 13:46:05
	DeptTime        *string  `json:"dept_time,omitempty" xml:"dept_time,omitempty"`
	DriverAddDetail *string  `json:"driver_add_detail,omitempty" xml:"driver_add_detail,omitempty"`
	DriverAddFee    *float64 `json:"driver_add_fee,omitempty" xml:"driver_add_fee,omitempty"`
	// example:
	//
	// 29.07
	EstimateDriveDistance *string `json:"estimate_drive_distance,omitempty" xml:"estimate_drive_distance,omitempty"`
	// example:
	//
	// 69
	EstimatePrice *float64 `json:"estimate_price,omitempty" xml:"estimate_price,omitempty"`
	// example:
	//
	// 40107
	FeeType     *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	// example:
	//
	// 4988580
	Index              *string `json:"index,omitempty" xml:"index,omitempty"`
	InvoiceTitle       *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	LevelName          *string `json:"level_name,omitempty" xml:"level_name,omitempty"`
	MappingCompanyCode *string `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	Memo               *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// example:
	//
	// 110285961234324
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 12.7
	OrderPrice *float64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	// example:
	//
	// 34535465346
	OverApplyId           *string `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	PaymentDepartmentId   *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// 0
	PersonSettleFee *float64 `json:"person_settle_fee,omitempty" xml:"person_settle_fee,omitempty"`
	Position        *string  `json:"position,omitempty" xml:"position,omitempty"`
	PositionLevel   *string  `json:"position_level,omitempty" xml:"position_level,omitempty"`
	// example:
	//
	// 4988580
	PrimaryId       *int64  `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProcessorOaCode *string `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// 23423432423
	ProjectCode         *string  `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName         *string  `json:"project_name,omitempty" xml:"project_name,omitempty"`
	ProtocolDiscountFee *float64 `json:"protocol_discount_fee,omitempty" xml:"protocol_discount_fee,omitempty"`
	ProviderName        *string  `json:"provider_name,omitempty" xml:"provider_name,omitempty"`
	// example:
	//
	// 0.00
	RealDriveDistance *string `json:"real_drive_distance,omitempty" xml:"real_drive_distance,omitempty"`
	RealFromAddr      *string `json:"real_from_addr,omitempty" xml:"real_from_addr,omitempty"`
	RealToAddr        *string `json:"real_to_addr,omitempty" xml:"real_to_addr,omitempty"`
	Remark            *string `json:"remark,omitempty" xml:"remark,omitempty"`
	SceneId           *string `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	SceneName         *string `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	// example:
	//
	// 12.7
	ServiceFee     *float64 `json:"service_fee,omitempty" xml:"service_fee,omitempty"`
	SettleTypeDesc *string  `json:"settle_type_desc,omitempty" xml:"settle_type_desc,omitempty"`
	// example:
	//
	// 5
	SettlementFee *float64 `json:"settlement_fee,omitempty" xml:"settlement_fee,omitempty"`
	// example:
	//
	// 6.11
	SettlementGrantFee *float64 `json:"settlement_grant_fee,omitempty" xml:"settlement_grant_fee,omitempty"`
	// example:
	//
	// 2022-05-15 22:27:00
	SettlementTime *string `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	// example:
	//
	// 4
	SettlementType *string `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	SpecialOrder   *string `json:"special_order,omitempty" xml:"special_order,omitempty"`
	SpecialReason  *string `json:"special_reason,omitempty" xml:"special_reason,omitempty"`
	// example:
	//
	// 1
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// 123123232
	SubOrderId        *string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	SupplementApplyId *string `json:"supplement_apply_id,omitempty" xml:"supplement_apply_id,omitempty"`
	// 税率
	//
	// example:
	//
	// 6%
	TaxRate          *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	ThirdItineraryId *string `json:"third_itinerary_id,omitempty" xml:"third_itinerary_id,omitempty"`
	TimeType         *string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	TradeActionDesc  *string `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
	TravelerId       *string `json:"traveler_id,omitempty" xml:"traveler_id,omitempty"`
	// example:
	//
	// 70022164
	TravelerJobNo          *string `json:"traveler_job_no,omitempty" xml:"traveler_job_no,omitempty"`
	TravelerMemberType     *string `json:"traveler_member_type,omitempty" xml:"traveler_member_type,omitempty"`
	TravelerMemberTypeName *string `json:"traveler_member_type_name,omitempty" xml:"traveler_member_type_name,omitempty"`
	TravelerName           *string `json:"traveler_name,omitempty" xml:"traveler_name,omitempty"`
	UserConfirmDesc        *string `json:"user_confirm_desc,omitempty" xml:"user_confirm_desc,omitempty"`
	VehicleSceneId         *string `json:"vehicle_scene_id,omitempty" xml:"vehicle_scene_id,omitempty"`
	VehicleSceneName       *string `json:"vehicle_scene_name,omitempty" xml:"vehicle_scene_name,omitempty"`
	// example:
	//
	// 11
	VoucherType     *int32  `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	VoucherTypeDesc *string `json:"voucher_type_desc,omitempty" xml:"voucher_type_desc,omitempty"`
}

func (s CarBillSettlementQueryResponseBodyModuleDataList) String() string {
	return dara.Prettify(s)
}

func (s CarBillSettlementQueryResponseBodyModuleDataList) GoString() string {
	return s.String()
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityCode() *string {
	return s.ApplyArrCityCode
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetApplyArrCityName() *string {
	return s.ApplyArrCityName
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityCode() *string {
	return s.ApplyDepCityCode
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetApplyDepCityName() *string {
	return s.ApplyDepCityName
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetApplyId() *string {
	return s.ApplyId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetArrCity() *string {
	return s.ArrCity
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetArrDate() *string {
	return s.ArrDate
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetArrLocation() *string {
	return s.ArrLocation
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetArrTime() *string {
	return s.ArrTime
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetBaseLocation() *string {
	return s.BaseLocation
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetBillingEntity() *string {
	return s.BillingEntity
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetBookModel() *string {
	return s.BookModel
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetBookTime() *string {
	return s.BookTime
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetBookerId() *string {
	return s.BookerId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetBookerName() *string {
	return s.BookerName
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetBusinessCategory() *string {
	return s.BusinessCategory
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetCarLevel() *string {
	return s.CarLevel
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetCostCenter() *string {
	return s.CostCenter
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetCoupon() *float64 {
	return s.Coupon
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetCouponPrice() *float64 {
	return s.CouponPrice
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetCustomContent() *string {
	return s.CustomContent
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetDeductibleTax() *float64 {
	return s.DeductibleTax
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetDepartment() *string {
	return s.Department
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetDeptCity() *string {
	return s.DeptCity
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetDeptDate() *string {
	return s.DeptDate
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetDeptLocation() *string {
	return s.DeptLocation
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetDeptTime() *string {
	return s.DeptTime
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetDriverAddDetail() *string {
	return s.DriverAddDetail
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetDriverAddFee() *float64 {
	return s.DriverAddFee
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetEstimateDriveDistance() *string {
	return s.EstimateDriveDistance
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetEstimatePrice() *float64 {
	return s.EstimatePrice
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetFeeType() *string {
	return s.FeeType
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetIndex() *string {
	return s.Index
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetLevelName() *string {
	return s.LevelName
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetMemo() *string {
	return s.Memo
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetOrderId() *string {
	return s.OrderId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetOrderPrice() *float64 {
	return s.OrderPrice
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetOverApplyId() *string {
	return s.OverApplyId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetPersonSettleFee() *float64 {
	return s.PersonSettleFee
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetPosition() *string {
	return s.Position
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetProjectName() *string {
	return s.ProjectName
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetProtocolDiscountFee() *float64 {
	return s.ProtocolDiscountFee
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetProviderName() *string {
	return s.ProviderName
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetRealDriveDistance() *string {
	return s.RealDriveDistance
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetRealFromAddr() *string {
	return s.RealFromAddr
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetRealToAddr() *string {
	return s.RealToAddr
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetRemark() *string {
	return s.Remark
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetSceneId() *string {
	return s.SceneId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetSceneName() *string {
	return s.SceneName
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetSettlementGrantFee() *float64 {
	return s.SettlementGrantFee
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetSettlementType() *string {
	return s.SettlementType
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetSpecialOrder() *string {
	return s.SpecialOrder
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetSpecialReason() *string {
	return s.SpecialReason
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetStatus() *int32 {
	return s.Status
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetSupplementApplyId() *string {
	return s.SupplementApplyId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetTaxRate() *string {
	return s.TaxRate
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetThirdItineraryId() *string {
	return s.ThirdItineraryId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetTimeType() *string {
	return s.TimeType
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetTravelerId() *string {
	return s.TravelerId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetTravelerMemberTypeName() *string {
	return s.TravelerMemberTypeName
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetTravelerName() *string {
	return s.TravelerName
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetUserConfirmDesc() *string {
	return s.UserConfirmDesc
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetVehicleSceneId() *string {
	return s.VehicleSceneId
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetVehicleSceneName() *string {
	return s.VehicleSceneName
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetAdjustTime(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.AdjustTime = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetAlipayTradeNo(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.AlipayTradeNo = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityCode(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityCode = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetApplyArrCityName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyArrCityName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityCode(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityCode = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetApplyDepCityName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyDepCityName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetApplyExtendField(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyExtendField = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetApplyId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ApplyId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetArrCity(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCity = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetArrCityCode(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ArrCityCode = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetArrDate(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ArrDate = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetArrLocation(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ArrLocation = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetArrTime(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ArrTime = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBaseLocation(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BaseLocation = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBillRecordTime(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BillRecordTime = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBillingEntity(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BillingEntity = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBookModel(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BookModel = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBookTime(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BookTime = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBookerId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BookerId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBookerJobNo(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BookerJobNo = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBookerName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BookerName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetBusinessCategory(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.BusinessCategory = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCapitalDirection(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CapitalDirection = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCarLevel(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CarLevel = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCascadeDepartment(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CascadeDepartment = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCategoryDesc(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CategoryDesc = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCostCenter(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenter = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCostCenterNumber(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CostCenterNumber = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCostDepartment(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CostDepartment = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCoupon(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Coupon = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCouponPrice(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CouponPrice = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetCustomContent(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.CustomContent = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDeductibleTax(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DeductibleTax = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDepCityCode(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DepCityCode = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDepartment(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Department = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDepartmentId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DepartmentId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDeptCity(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DeptCity = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDeptDate(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DeptDate = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDeptLocation(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DeptLocation = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDeptTime(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DeptTime = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDriverAddDetail(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DriverAddDetail = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetDriverAddFee(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.DriverAddFee = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetEstimateDriveDistance(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.EstimateDriveDistance = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetEstimatePrice(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.EstimatePrice = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetFeeType(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.FeeType = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetFeeTypeDesc(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.FeeTypeDesc = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetIndex(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Index = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetInvoiceTitle(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.InvoiceTitle = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetLevelName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.LevelName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetMappingCompanyCode(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.MappingCompanyCode = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetMemo(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Memo = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetOrderId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.OrderId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetOrderPrice(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.OrderPrice = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetOverApplyId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.OverApplyId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetPaymentDepartmentName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.PaymentDepartmentName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetPersonSettleFee(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.PersonSettleFee = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetPosition(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Position = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetPositionLevel(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.PositionLevel = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetPrimaryId(v int64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.PrimaryId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetProcessorOaCode(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ProcessorOaCode = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetProjectCode(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectCode = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetProjectName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ProjectName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetProtocolDiscountFee(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ProtocolDiscountFee = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetProviderName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ProviderName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetRealDriveDistance(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.RealDriveDistance = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetRealFromAddr(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.RealFromAddr = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetRealToAddr(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.RealToAddr = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetRemark(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Remark = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSceneId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SceneId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSceneName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SceneName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetServiceFee(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ServiceFee = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSettleTypeDesc(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SettleTypeDesc = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSettlementFee(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementFee = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSettlementGrantFee(v float64) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementGrantFee = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSettlementTime(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementTime = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSettlementType(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SettlementType = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSpecialOrder(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SpecialOrder = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSpecialReason(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SpecialReason = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetStatus(v int32) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.Status = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetStatusDesc(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.StatusDesc = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSubOrderId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SubOrderId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetSupplementApplyId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.SupplementApplyId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetTaxRate(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.TaxRate = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetThirdItineraryId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.ThirdItineraryId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetTimeType(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.TimeType = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetTradeActionDesc(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.TradeActionDesc = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetTravelerId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetTravelerJobNo(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerJobNo = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetTravelerMemberType(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerMemberType = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetTravelerMemberTypeName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerMemberTypeName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetTravelerName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.TravelerName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetUserConfirmDesc(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.UserConfirmDesc = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetVehicleSceneId(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.VehicleSceneId = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetVehicleSceneName(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.VehicleSceneName = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetVoucherType(v int32) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherType = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) SetVoucherTypeDesc(v string) *CarBillSettlementQueryResponseBodyModuleDataList {
	s.VoucherTypeDesc = &v
	return s
}

func (s *CarBillSettlementQueryResponseBodyModuleDataList) Validate() error {
	return dara.Validate(s)
}
