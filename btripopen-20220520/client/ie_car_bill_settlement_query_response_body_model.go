// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIeCarBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *IeCarBillSettlementQueryResponseBody
	GetCode() *string
	SetMessage(v string) *IeCarBillSettlementQueryResponseBody
	GetMessage() *string
	SetModule(v *IeCarBillSettlementQueryResponseBodyModule) *IeCarBillSettlementQueryResponseBody
	GetModule() *IeCarBillSettlementQueryResponseBodyModule
	SetRequestId(v string) *IeCarBillSettlementQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *IeCarBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *IeCarBillSettlementQueryResponseBody
	GetTraceId() *string
}

type IeCarBillSettlementQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *string                                     `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                     `json:"message,omitempty" xml:"message,omitempty"`
	Module  *IeCarBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// A1984987-9C0D-5EEB-A2AC-0D5D76D41705
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 3b5287ed17606676287351344d0095
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s IeCarBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IeCarBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *IeCarBillSettlementQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *IeCarBillSettlementQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *IeCarBillSettlementQueryResponseBody) GetModule() *IeCarBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *IeCarBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IeCarBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *IeCarBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *IeCarBillSettlementQueryResponseBody) SetCode(v string) *IeCarBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBody) SetMessage(v string) *IeCarBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBody) SetModule(v *IeCarBillSettlementQueryResponseBodyModule) *IeCarBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *IeCarBillSettlementQueryResponseBody) SetRequestId(v string) *IeCarBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBody) SetSuccess(v bool) *IeCarBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBody) SetTraceId(v string) *IeCarBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type IeCarBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 14
	Category *int32                                             `json:"category,omitempty" xml:"category,omitempty"`
	CorpId   *string                                            `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Items    []*IeCarBillSettlementQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 2021-10-14
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2021-10-13
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	// example:
	//
	// CAESBgoEIgIIABgAIhkKFwMSAAAAMUw4MDAwMDAwMDA4YTU2NDMy
	ScrollId *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// 100
	TotalNum *int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}

func (s IeCarBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s IeCarBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *IeCarBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *IeCarBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *IeCarBillSettlementQueryResponseBodyModule) GetItems() []*IeCarBillSettlementQueryResponseBodyModuleItems {
	return s.Items
}

func (s *IeCarBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *IeCarBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *IeCarBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *IeCarBillSettlementQueryResponseBodyModule) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *IeCarBillSettlementQueryResponseBodyModule) SetCategory(v int32) *IeCarBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModule) SetCorpId(v string) *IeCarBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModule) SetItems(v []*IeCarBillSettlementQueryResponseBodyModuleItems) *IeCarBillSettlementQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *IeCarBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *IeCarBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModule) SetScrollId(v string) *IeCarBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *IeCarBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModule) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type IeCarBillSettlementQueryResponseBodyModuleItems struct {
	// example:
	//
	// 20251201
	AccountMonth *string `json:"account_month,omitempty" xml:"account_month,omitempty"`
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
	// 265741695
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 123@qq.com
	ApproverEmail *string `json:"approver_email,omitempty" xml:"approver_email,omitempty"`
	// example:
	//
	// 11
	ApproverId   *string `json:"approver_id,omitempty" xml:"approver_id,omitempty"`
	ApproverName *string `json:"approver_name,omitempty" xml:"approver_name,omitempty"`
	ArrCity      *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// 131000
	ArrCityCode *string `json:"arr_city_code,omitempty" xml:"arr_city_code,omitempty"`
	ArrCountry  *string `json:"arr_country,omitempty" xml:"arr_country,omitempty"`
	// example:
	//
	// 2025-12-02 10:00:00
	ArrDate *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// example:
	//
	// 2025-12-02 10:00:00
	ArrDateLocal *string `json:"arr_date_local,omitempty" xml:"arr_date_local,omitempty"`
	ArrLocation  *string `json:"arr_location,omitempty" xml:"arr_location,omitempty"`
	BaseLocation *string `json:"base_location,omitempty" xml:"base_location,omitempty"`
	// example:
	//
	// 2023-01-01 00:00:00
	BillRecordTime *string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	// example:
	//
	// 2025-11-01 00:00:00
	BillRecordTimeStr *string `json:"bill_record_time_str,omitempty" xml:"bill_record_time_str,omitempty"`
	BillingEntity     *string `json:"billing_entity,omitempty" xml:"billing_entity,omitempty"`
	BookChannel       *string `json:"book_channel,omitempty" xml:"book_channel,omitempty"`
	BookMode          *string `json:"book_mode,omitempty" xml:"book_mode,omitempty"`
	BookModel         *string `json:"book_model,omitempty" xml:"book_model,omitempty"`
	// example:
	//
	// 2023-01-01 00:00:00
	BookTime *string `json:"book_time,omitempty" xml:"book_time,omitempty"`
	// example:
	//
	// 2025-12-01 10:00:00
	BookTimeLocal *string `json:"book_time_local,omitempty" xml:"book_time_local,omitempty"`
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
	// 123
	BookerUseId       *string `json:"booker_use_id,omitempty" xml:"booker_use_id,omitempty"`
	BusinessCategory  *string `json:"business_category,omitempty" xml:"business_category,omitempty"`
	CapitalDirection  *string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CarLevel          *string `json:"car_level,omitempty" xml:"car_level,omitempty"`
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
	// example:
	//
	// 0.0
	Coupon *float64 `json:"coupon,omitempty" xml:"coupon,omitempty"`
	// example:
	//
	// 3.0
	CouponPrice *float64 `json:"coupon_price,omitempty" xml:"coupon_price,omitempty"`
	// example:
	//
	// "{\\"key1\\":\\"value1\\",\\"key2\\":\\"value2\\",\\"key3\\":\\"value3\\"}"
	CustomContent *string `json:"custom_content,omitempty" xml:"custom_content,omitempty"`
	// example:
	//
	// 0.11
	DeductibleTax *float64 `json:"deductible_tax,omitempty" xml:"deductible_tax,omitempty"`
	// example:
	//
	// 131000
	DepCityCode *string `json:"dep_city_code,omitempty" xml:"dep_city_code,omitempty"`
	DepCountry  *string `json:"dep_country,omitempty" xml:"dep_country,omitempty"`
	// example:
	//
	// 2025-12-01 10:00:00
	DepDateLocal *string `json:"dep_date_local,omitempty" xml:"dep_date_local,omitempty"`
	Department   *string `json:"department,omitempty" xml:"department,omitempty"`
	// example:
	//
	// 1112
	DepartmentId *string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	DeptCity     *string `json:"dept_city,omitempty" xml:"dept_city,omitempty"`
	// example:
	//
	// 2025-12-01 10:00:00
	DeptDate     *string `json:"dept_date,omitempty" xml:"dept_date,omitempty"`
	DeptLocation *string `json:"dept_location,omitempty" xml:"dept_location,omitempty"`
	// example:
	//
	// 2.0
	DriverAddDetail *string `json:"driver_add_detail,omitempty" xml:"driver_add_detail,omitempty"`
	// example:
	//
	// 1.0
	DriverAddFee *float64 `json:"driver_add_fee,omitempty" xml:"driver_add_fee,omitempty"`
	// example:
	//
	// 100
	EstimateDriveDistance *string `json:"estimate_drive_distance,omitempty" xml:"estimate_drive_distance,omitempty"`
	// example:
	//
	// 110.0
	EstimatePrice *float64 `json:"estimate_price,omitempty" xml:"estimate_price,omitempty"`
	// example:
	//
	// 140101
	FeeType          *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc      *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	FinancialSubject *string `json:"financial_subject,omitempty" xml:"financial_subject,omitempty"`
	ForeignersTag    *string `json:"foreigners_tag,omitempty" xml:"foreigners_tag,omitempty"`
	// example:
	//
	// 1
	Index          *string `json:"index,omitempty" xml:"index,omitempty"`
	InvoiceTitle   *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	LastDepartment *string `json:"last_department,omitempty" xml:"last_department,omitempty"`
	LevelName      *string `json:"level_name,omitempty" xml:"level_name,omitempty"`
	// Location
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// 112
	MainApplyId *string `json:"main_apply_id,omitempty" xml:"main_apply_id,omitempty"`
	// example:
	//
	// q1
	MappingCompanyCode *string `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	Memo               *string `json:"memo,omitempty" xml:"memo,omitempty"`
	// example:
	//
	// 1017034204020120044
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 100.0
	OrderPrice      *float64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	OrderStatusDesc *string  `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// example:
	//
	// 112
	OriginApplyId *string `json:"origin_apply_id,omitempty" xml:"origin_apply_id,omitempty"`
	// example:
	//
	// 123
	OverApplyId *string `json:"over_apply_id,omitempty" xml:"over_apply_id,omitempty"`
	// example:
	//
	// EN01002423
	PaymentDepartmentId   *string `json:"payment_department_id,omitempty" xml:"payment_department_id,omitempty"`
	PaymentDepartmentName *string `json:"payment_department_name,omitempty" xml:"payment_department_name,omitempty"`
	// example:
	//
	// 10.0
	PersonSettleFee *float64 `json:"person_settle_fee,omitempty" xml:"person_settle_fee,omitempty"`
	Position        *string  `json:"position,omitempty" xml:"position,omitempty"`
	PositionLevel   *string  `json:"position_level,omitempty" xml:"position_level,omitempty"`
	// example:
	//
	// 72328485
	PrimaryId *int64 `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	// example:
	//
	// www123
	ProcessorOaCode *string `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// acs
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// example:
	//
	// 10.0
	ProtocolDiscountFee *float64 `json:"protocol_discount_fee,omitempty" xml:"protocol_discount_fee,omitempty"`
	ProviderName        *string  `json:"provider_name,omitempty" xml:"provider_name,omitempty"`
	// example:
	//
	// 111224324
	PurchaseOrderId *string `json:"purchase_order_id,omitempty" xml:"purchase_order_id,omitempty"`
	// example:
	//
	// 120
	RealDriveDistance *string `json:"real_drive_distance,omitempty" xml:"real_drive_distance,omitempty"`
	RealFromAddr      *string `json:"real_from_addr,omitempty" xml:"real_from_addr,omitempty"`
	RealToAddr        *string `json:"real_to_addr,omitempty" xml:"real_to_addr,omitempty"`
	RecordStatus      *string `json:"record_status,omitempty" xml:"record_status,omitempty"`
	Remark            *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 1
	SceneId   *string `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	SceneName *string `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
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
	// example:
	//
	// 188385292
	ShowSubOrderId *string `json:"show_sub_order_id,omitempty" xml:"show_sub_order_id,omitempty"`
	// SIO
	//
	// example:
	//
	// SIO
	Sio           *string `json:"sio,omitempty" xml:"sio,omitempty"`
	SpecialOrder  *string `json:"special_order,omitempty" xml:"special_order,omitempty"`
	SpecialReason *string `json:"special_reason,omitempty" xml:"special_reason,omitempty"`
	// example:
	//
	// 2
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// 188385292
	SubOrderId *string `json:"sub_order_id,omitempty" xml:"sub_order_id,omitempty"`
	// example:
	//
	// 123456789
	SupplementApplyId *string `json:"supplement_apply_id,omitempty" xml:"supplement_apply_id,omitempty"`
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
	// 123
	ThirdPartBusinessId *string `json:"third_part_business_id,omitempty" xml:"third_part_business_id,omitempty"`
	// example:
	//
	// 7244-1968
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	TimeType         *string `json:"time_type,omitempty" xml:"time_type,omitempty"`
	TradeActionDesc  *string `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
	// example:
	//
	// 123@qq.com
	TravelerEmail *string `json:"traveler_email,omitempty" xml:"traveler_email,omitempty"`
	// example:
	//
	// 1234
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
	// 123
	TravelerUseId   *string `json:"traveler_use_id,omitempty" xml:"traveler_use_id,omitempty"`
	UserConfirmDesc *string `json:"user_confirm_desc,omitempty" xml:"user_confirm_desc,omitempty"`
	// example:
	//
	// 12
	VehicleSceneId   *string `json:"vehicle_scene_id,omitempty" xml:"vehicle_scene_id,omitempty"`
	VehicleSceneName *string `json:"vehicle_scene_name,omitempty" xml:"vehicle_scene_name,omitempty"`
	// example:
	//
	// 1
	VoucherType     *int32  `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	VoucherTypeDesc *string `json:"voucher_type_desc,omitempty" xml:"voucher_type_desc,omitempty"`
}

func (s IeCarBillSettlementQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s IeCarBillSettlementQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetAccountMonth() *string {
	return s.AccountMonth
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetAlipayId() *string {
	return s.AlipayId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetApplyArrCityCode() *string {
	return s.ApplyArrCityCode
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetApplyArrCityName() *string {
	return s.ApplyArrCityName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetApplyDepCityCode() *string {
	return s.ApplyDepCityCode
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetApplyDepCityName() *string {
	return s.ApplyDepCityName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetApplyId() *string {
	return s.ApplyId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetApproverEmail() *string {
	return s.ApproverEmail
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetApproverId() *string {
	return s.ApproverId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetApproverName() *string {
	return s.ApproverName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetArrCity() *string {
	return s.ArrCity
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetArrCityCode() *string {
	return s.ArrCityCode
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetArrCountry() *string {
	return s.ArrCountry
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetArrDate() *string {
	return s.ArrDate
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetArrDateLocal() *string {
	return s.ArrDateLocal
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetArrLocation() *string {
	return s.ArrLocation
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBaseLocation() *string {
	return s.BaseLocation
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBillRecordTimeStr() *string {
	return s.BillRecordTimeStr
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBillingEntity() *string {
	return s.BillingEntity
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBookChannel() *string {
	return s.BookChannel
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBookMode() *string {
	return s.BookMode
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBookModel() *string {
	return s.BookModel
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBookTime() *string {
	return s.BookTime
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBookTimeLocal() *string {
	return s.BookTimeLocal
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBookerId() *string {
	return s.BookerId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBookerName() *string {
	return s.BookerName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBookerUseId() *string {
	return s.BookerUseId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetBusinessCategory() *string {
	return s.BusinessCategory
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetCarLevel() *string {
	return s.CarLevel
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetCostCenter() *string {
	return s.CostCenter
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetCoupon() *float64 {
	return s.Coupon
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetCouponPrice() *float64 {
	return s.CouponPrice
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetCustomContent() *string {
	return s.CustomContent
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetDeductibleTax() *float64 {
	return s.DeductibleTax
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetDepCityCode() *string {
	return s.DepCityCode
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetDepCountry() *string {
	return s.DepCountry
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetDepDateLocal() *string {
	return s.DepDateLocal
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetDepartment() *string {
	return s.Department
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetDeptCity() *string {
	return s.DeptCity
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetDeptDate() *string {
	return s.DeptDate
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetDeptLocation() *string {
	return s.DeptLocation
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetDriverAddDetail() *string {
	return s.DriverAddDetail
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetDriverAddFee() *float64 {
	return s.DriverAddFee
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetEstimateDriveDistance() *string {
	return s.EstimateDriveDistance
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetEstimatePrice() *float64 {
	return s.EstimatePrice
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetFeeType() *string {
	return s.FeeType
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetFinancialSubject() *string {
	return s.FinancialSubject
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetForeignersTag() *string {
	return s.ForeignersTag
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetIndex() *string {
	return s.Index
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetLastDepartment() *string {
	return s.LastDepartment
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetLevelName() *string {
	return s.LevelName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetLocation() *string {
	return s.Location
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetMainApplyId() *string {
	return s.MainApplyId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetMemo() *string {
	return s.Memo
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetOrderId() *string {
	return s.OrderId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetOrderPrice() *float64 {
	return s.OrderPrice
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetOriginApplyId() *string {
	return s.OriginApplyId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetOverApplyId() *string {
	return s.OverApplyId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetPersonSettleFee() *float64 {
	return s.PersonSettleFee
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetPosition() *string {
	return s.Position
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetPositionLevel() *string {
	return s.PositionLevel
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetProjectName() *string {
	return s.ProjectName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetProtocolDiscountFee() *float64 {
	return s.ProtocolDiscountFee
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetProviderName() *string {
	return s.ProviderName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetPurchaseOrderId() *string {
	return s.PurchaseOrderId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetRealDriveDistance() *string {
	return s.RealDriveDistance
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetRealFromAddr() *string {
	return s.RealFromAddr
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetRealToAddr() *string {
	return s.RealToAddr
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetRecordStatus() *string {
	return s.RecordStatus
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetRemark() *string {
	return s.Remark
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSceneId() *string {
	return s.SceneId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSceneName() *string {
	return s.SceneName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSettlementGrantFee() *float64 {
	return s.SettlementGrantFee
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSettlementType() *string {
	return s.SettlementType
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetShowSubOrderId() *string {
	return s.ShowSubOrderId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSio() *string {
	return s.Sio
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSpecialOrder() *string {
	return s.SpecialOrder
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSpecialReason() *string {
	return s.SpecialReason
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetStatus() *int32 {
	return s.Status
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetSupplementApplyId() *string {
	return s.SupplementApplyId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetTaxRate() *string {
	return s.TaxRate
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetThirdInvoiceId() *string {
	return s.ThirdInvoiceId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetThirdItineraryId() *string {
	return s.ThirdItineraryId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetThirdPartBusinessId() *string {
	return s.ThirdPartBusinessId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetTimeType() *string {
	return s.TimeType
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetTravelerEmail() *string {
	return s.TravelerEmail
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetTravelerId() *string {
	return s.TravelerId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberTypeName() *string {
	return s.TravelerMemberTypeName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetTravelerName() *string {
	return s.TravelerName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetTravelerUseId() *string {
	return s.TravelerUseId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetUserConfirmDesc() *string {
	return s.UserConfirmDesc
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetVehicleSceneId() *string {
	return s.VehicleSceneId
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetVehicleSceneName() *string {
	return s.VehicleSceneName
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetAccountMonth(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.AccountMonth = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetAdjustTime(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.AdjustTime = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetAlipayId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.AlipayId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetAlipayTradeNo(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.AlipayTradeNo = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetApplyArrCityCode(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ApplyArrCityCode = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetApplyArrCityName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ApplyArrCityName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetApplyDepCityCode(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ApplyDepCityCode = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetApplyDepCityName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ApplyDepCityName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetApplyExtendField(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ApplyExtendField = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetApplyId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ApplyId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetApproverEmail(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ApproverEmail = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetApproverId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ApproverId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetApproverName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ApproverName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetArrCity(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ArrCity = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetArrCityCode(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ArrCityCode = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetArrCountry(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ArrCountry = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetArrDate(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ArrDate = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetArrDateLocal(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ArrDateLocal = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetArrLocation(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ArrLocation = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBaseLocation(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BaseLocation = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBillRecordTime(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BillRecordTime = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBillRecordTimeStr(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BillRecordTimeStr = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBillingEntity(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BillingEntity = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBookChannel(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BookChannel = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBookMode(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BookMode = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBookModel(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BookModel = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBookTime(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BookTime = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBookTimeLocal(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BookTimeLocal = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBookerId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BookerId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBookerJobNo(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BookerJobNo = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBookerName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BookerName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBookerUseId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BookerUseId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetBusinessCategory(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.BusinessCategory = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetCapitalDirection(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.CapitalDirection = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetCarLevel(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.CarLevel = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetCascadeDepartment(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.CascadeDepartment = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetCategoryDesc(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.CategoryDesc = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetCostCenter(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.CostCenter = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetCostCenterNumber(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.CostCenterNumber = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetCostDepartment(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.CostDepartment = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetCoupon(v float64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.Coupon = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetCouponPrice(v float64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.CouponPrice = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetCustomContent(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.CustomContent = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetDeductibleTax(v float64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.DeductibleTax = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetDepCityCode(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.DepCityCode = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetDepCountry(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.DepCountry = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetDepDateLocal(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.DepDateLocal = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetDepartment(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.Department = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetDepartmentId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.DepartmentId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetDeptCity(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.DeptCity = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetDeptDate(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.DeptDate = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetDeptLocation(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.DeptLocation = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetDriverAddDetail(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.DriverAddDetail = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetDriverAddFee(v float64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.DriverAddFee = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetEstimateDriveDistance(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.EstimateDriveDistance = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetEstimatePrice(v float64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.EstimatePrice = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetFeeType(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.FeeType = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetFeeTypeDesc(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.FeeTypeDesc = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetFinancialSubject(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.FinancialSubject = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetForeignersTag(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ForeignersTag = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetIndex(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.Index = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetInvoiceTitle(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.InvoiceTitle = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetLastDepartment(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.LastDepartment = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetLevelName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.LevelName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetLocation(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.Location = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetMainApplyId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.MainApplyId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetMappingCompanyCode(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.MappingCompanyCode = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetMemo(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.Memo = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetOrderId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.OrderId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetOrderPrice(v float64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.OrderPrice = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetOrderStatusDesc(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.OrderStatusDesc = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetOriginApplyId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.OriginApplyId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetOverApplyId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.OverApplyId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetPaymentDepartmentId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.PaymentDepartmentId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetPaymentDepartmentName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.PaymentDepartmentName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetPersonSettleFee(v float64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.PersonSettleFee = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetPosition(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.Position = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetPositionLevel(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.PositionLevel = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetPrimaryId(v int64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.PrimaryId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetProcessorOaCode(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ProcessorOaCode = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetProjectCode(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ProjectCode = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetProjectName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ProjectName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetProtocolDiscountFee(v float64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ProtocolDiscountFee = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetProviderName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ProviderName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetPurchaseOrderId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.PurchaseOrderId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetRealDriveDistance(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.RealDriveDistance = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetRealFromAddr(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.RealFromAddr = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetRealToAddr(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.RealToAddr = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetRecordStatus(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.RecordStatus = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetRemark(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.Remark = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSceneId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.SceneId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSceneName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.SceneName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetServiceFee(v float64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ServiceFee = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSettleTypeDesc(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.SettleTypeDesc = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSettlementFee(v float64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.SettlementFee = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSettlementGrantFee(v float64) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.SettlementGrantFee = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSettlementTime(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.SettlementTime = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSettlementType(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.SettlementType = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetShowSubOrderId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ShowSubOrderId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSio(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.Sio = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSpecialOrder(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.SpecialOrder = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSpecialReason(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.SpecialReason = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetStatus(v int32) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.Status = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetStatusDesc(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.StatusDesc = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSubOrderId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.SubOrderId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetSupplementApplyId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.SupplementApplyId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetTaxRate(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.TaxRate = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetThirdInvoiceId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ThirdInvoiceId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetThirdItineraryId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ThirdItineraryId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetThirdPartBusinessId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ThirdPartBusinessId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetThirdpartApplyId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.ThirdpartApplyId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetTimeType(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.TimeType = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetTradeActionDesc(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.TradeActionDesc = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetTravelerEmail(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.TravelerEmail = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetTravelerId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.TravelerId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetTravelerJobNo(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.TravelerJobNo = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberType(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberType = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberTypeName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberTypeName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetTravelerName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.TravelerName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetTravelerUseId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.TravelerUseId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetUserConfirmDesc(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.UserConfirmDesc = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetVehicleSceneId(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.VehicleSceneId = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetVehicleSceneName(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.VehicleSceneName = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetVoucherType(v int32) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.VoucherType = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) SetVoucherTypeDesc(v string) *IeCarBillSettlementQueryResponseBodyModuleItems {
	s.VoucherTypeDesc = &v
	return s
}

func (s *IeCarBillSettlementQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}
