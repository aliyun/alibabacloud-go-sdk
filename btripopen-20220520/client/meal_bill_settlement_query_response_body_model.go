// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMealBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *MealBillSettlementQueryResponseBody
	GetCode() *int32
	SetMessage(v string) *MealBillSettlementQueryResponseBody
	GetMessage() *string
	SetModule(v *MealBillSettlementQueryResponseBodyModule) *MealBillSettlementQueryResponseBody
	GetModule() *MealBillSettlementQueryResponseBodyModule
	SetRequestId(v string) *MealBillSettlementQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *MealBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *MealBillSettlementQueryResponseBody
	GetTraceId() *string
}

type MealBillSettlementQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code    *int32                                     `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                    `json:"message,omitempty" xml:"message,omitempty"`
	Module  *MealBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 210f079e16603757182131635d866a
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s MealBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MealBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *MealBillSettlementQueryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *MealBillSettlementQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MealBillSettlementQueryResponseBody) GetModule() *MealBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *MealBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MealBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *MealBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *MealBillSettlementQueryResponseBody) SetCode(v int32) *MealBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *MealBillSettlementQueryResponseBody) SetMessage(v string) *MealBillSettlementQueryResponseBody {
	s.Message = &v
	return s
}

func (s *MealBillSettlementQueryResponseBody) SetModule(v *MealBillSettlementQueryResponseBodyModule) *MealBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *MealBillSettlementQueryResponseBody) SetRequestId(v string) *MealBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBody) SetSuccess(v bool) *MealBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *MealBillSettlementQueryResponseBody) SetTraceId(v string) *MealBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type MealBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 7
	Category *int32                                            `json:"category,omitempty" xml:"category,omitempty"`
	CorpId   *string                                           `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Items    []*MealBillSettlementQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1002039195025156700
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 2022-07-02
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2022-07-01
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	ScrollId    *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// 2695
	TotalNum *int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
	// example:
	//
	// 30
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s MealBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s MealBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *MealBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *MealBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *MealBillSettlementQueryResponseBodyModule) GetItems() []*MealBillSettlementQueryResponseBodyModuleItems {
	return s.Items
}

func (s *MealBillSettlementQueryResponseBodyModule) GetOrderId() *string {
	return s.OrderId
}

func (s *MealBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *MealBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *MealBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *MealBillSettlementQueryResponseBodyModule) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *MealBillSettlementQueryResponseBodyModule) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *MealBillSettlementQueryResponseBodyModule) SetCategory(v int32) *MealBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModule) SetCorpId(v string) *MealBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModule) SetItems(v []*MealBillSettlementQueryResponseBodyModuleItems) *MealBillSettlementQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModule) SetOrderId(v string) *MealBillSettlementQueryResponseBodyModule {
	s.OrderId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *MealBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *MealBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModule) SetScrollId(v string) *MealBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModule) SetTotalNum(v int64) *MealBillSettlementQueryResponseBodyModule {
	s.TotalNum = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModule) SetTotalSize(v int64) *MealBillSettlementQueryResponseBodyModule {
	s.TotalSize = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type MealBillSettlementQueryResponseBodyModuleItems struct {
	AdjustTime       *string `json:"adjust_time,omitempty" xml:"adjust_time,omitempty"`
	ApplyExtendField *string `json:"apply_extend_field,omitempty" xml:"apply_extend_field,omitempty"`
	// example:
	//
	// 1004430880
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// example:
	//
	// 2023-01-01 00:00:00
	BillRecordTime *string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
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
	// 1
	CapitalDirection      *string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment     *string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CategoryDesc          *string `json:"category_desc,omitempty" xml:"category_desc,omitempty"`
	ConsumeReportAddress  *string `json:"consume_report_address,omitempty" xml:"consume_report_address,omitempty"`
	ConsumeReportCity     *string `json:"consume_report_city,omitempty" xml:"consume_report_city,omitempty"`
	ConsumeReportCityCode *string `json:"consume_report_city_code,omitempty" xml:"consume_report_city_code,omitempty"`
	ConsumerScene         *string `json:"consumer_scene,omitempty" xml:"consumer_scene,omitempty"`
	// example:
	//
	// 100.0
	CorpSettleFee *float64 `json:"corp_settle_fee,omitempty" xml:"corp_settle_fee,omitempty"`
	CostCenter    *string  `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// example:
	//
	// cs1
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	CostDepartment   *string `json:"cost_department,omitempty" xml:"cost_department,omitempty"`
	Department       *string `json:"department,omitempty" xml:"department,omitempty"`
	// example:
	//
	// 1112
	DepartmentId *string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	// example:
	//
	// 70101
	FeeType     *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	// example:
	//
	// 1
	Index        *string `json:"index,omitempty" xml:"index,omitempty"`
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// example:
	//
	// 123
	MainApplyId        *string `json:"main_apply_id,omitempty" xml:"main_apply_id,omitempty"`
	MappingCompanyCode *string `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	// example:
	//
	// XXXX
	MealAddress *string `json:"meal_address,omitempty" xml:"meal_address,omitempty"`
	MealCity    *string `json:"meal_city,omitempty" xml:"meal_city,omitempty"`
	// example:
	//
	// XXX
	MealCityCode *string `json:"meal_city_code,omitempty" xml:"meal_city_code,omitempty"`
	MealReason   *string `json:"meal_reason,omitempty" xml:"meal_reason,omitempty"`
	MealRule     *string `json:"meal_rule,omitempty" xml:"meal_rule,omitempty"`
	MealScene    *string `json:"meal_scene,omitempty" xml:"meal_scene,omitempty"`
	// example:
	//
	// MEAL
	MerchantCategory *string `json:"merchant_category,omitempty" xml:"merchant_category,omitempty"`
	// example:
	//
	// XXXX
	MerchantName *string `json:"merchant_name,omitempty" xml:"merchant_name,omitempty"`
	// example:
	//
	// 1002039196909288346
	OrderId *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// 100.0
	OrderPrice      *float64 `json:"order_price,omitempty" xml:"order_price,omitempty"`
	OrderStatusDesc *string  `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	// example:
	//
	// 0.0
	PersonSettlePrice *float64 `json:"person_settle_price,omitempty" xml:"person_settle_price,omitempty"`
	// example:
	//
	// 60698599
	PrimaryId       *int64  `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProcessorOaCode *string `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// acs
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	Remark      *string `json:"remark,omitempty" xml:"remark,omitempty"`
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
	// 2023-01-01 00:00:00
	SettlementTime *string `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	// example:
	//
	// 4
	SettlementType *string `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
	ShareDinner    *string `json:"share_dinner,omitempty" xml:"share_dinner,omitempty"`
	// example:
	//
	// 0
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// XXXX
	StoreAddress *string `json:"store_address,omitempty" xml:"store_address,omitempty"`
	// example:
	//
	// 6%
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// example:
	//
	// cs2
	ThirdInvoiceId *string `json:"third_invoice_id,omitempty" xml:"third_invoice_id,omitempty"`
	// example:
	//
	// 123
	ThirdPartBusinessId *string `json:"third_part_business_id,omitempty" xml:"third_part_business_id,omitempty"`
	// example:
	//
	// 7244-1968
	ThirdpartApplyId *string `json:"thirdpart_apply_id,omitempty" xml:"thirdpart_apply_id,omitempty"`
	TradeActionDesc  *string `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
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
	// 1
	VoucherType     *int32  `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	VoucherTypeDesc *string `json:"voucher_type_desc,omitempty" xml:"voucher_type_desc,omitempty"`
}

func (s MealBillSettlementQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s MealBillSettlementQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetApplyId() *string {
	return s.ApplyId
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetBookTime() *string {
	return s.BookTime
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetBookerId() *string {
	return s.BookerId
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetBookerName() *string {
	return s.BookerName
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetConsumeReportAddress() *string {
	return s.ConsumeReportAddress
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetConsumeReportCity() *string {
	return s.ConsumeReportCity
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetConsumeReportCityCode() *string {
	return s.ConsumeReportCityCode
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetConsumerScene() *string {
	return s.ConsumerScene
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetCorpSettleFee() *float64 {
	return s.CorpSettleFee
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetCostCenter() *string {
	return s.CostCenter
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetDepartment() *string {
	return s.Department
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetFeeType() *string {
	return s.FeeType
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetIndex() *string {
	return s.Index
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetMainApplyId() *string {
	return s.MainApplyId
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetMealAddress() *string {
	return s.MealAddress
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetMealCity() *string {
	return s.MealCity
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetMealCityCode() *string {
	return s.MealCityCode
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetMealReason() *string {
	return s.MealReason
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetMealRule() *string {
	return s.MealRule
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetMealScene() *string {
	return s.MealScene
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetMerchantCategory() *string {
	return s.MerchantCategory
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetMerchantName() *string {
	return s.MerchantName
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetOrderId() *string {
	return s.OrderId
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetOrderPrice() *float64 {
	return s.OrderPrice
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetPersonSettlePrice() *float64 {
	return s.PersonSettlePrice
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetProjectName() *string {
	return s.ProjectName
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetRemark() *string {
	return s.Remark
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetServiceFee() *float64 {
	return s.ServiceFee
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetSettlementType() *string {
	return s.SettlementType
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetShareDinner() *string {
	return s.ShareDinner
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetStatus() *int32 {
	return s.Status
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetStoreAddress() *string {
	return s.StoreAddress
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetTaxRate() *string {
	return s.TaxRate
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetThirdInvoiceId() *string {
	return s.ThirdInvoiceId
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetThirdPartBusinessId() *string {
	return s.ThirdPartBusinessId
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetThirdpartApplyId() *string {
	return s.ThirdpartApplyId
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetTravelerId() *string {
	return s.TravelerId
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberTypeName() *string {
	return s.TravelerMemberTypeName
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetTravelerName() *string {
	return s.TravelerName
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetAdjustTime(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.AdjustTime = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetApplyExtendField(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ApplyExtendField = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetApplyId(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ApplyId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetBillRecordTime(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.BillRecordTime = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetBookTime(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.BookTime = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetBookerId(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.BookerId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetBookerJobNo(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.BookerJobNo = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetBookerName(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.BookerName = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetCapitalDirection(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.CapitalDirection = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetCascadeDepartment(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.CascadeDepartment = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetCategoryDesc(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.CategoryDesc = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetConsumeReportAddress(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ConsumeReportAddress = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetConsumeReportCity(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ConsumeReportCity = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetConsumeReportCityCode(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ConsumeReportCityCode = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetConsumerScene(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ConsumerScene = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetCorpSettleFee(v float64) *MealBillSettlementQueryResponseBodyModuleItems {
	s.CorpSettleFee = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetCostCenter(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.CostCenter = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetCostCenterNumber(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.CostCenterNumber = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetCostDepartment(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.CostDepartment = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetDepartment(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.Department = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetDepartmentId(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.DepartmentId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetFeeType(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.FeeType = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetFeeTypeDesc(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.FeeTypeDesc = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetIndex(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.Index = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetInvoiceTitle(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.InvoiceTitle = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetMainApplyId(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.MainApplyId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetMappingCompanyCode(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.MappingCompanyCode = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetMealAddress(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.MealAddress = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetMealCity(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.MealCity = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetMealCityCode(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.MealCityCode = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetMealReason(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.MealReason = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetMealRule(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.MealRule = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetMealScene(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.MealScene = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetMerchantCategory(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.MerchantCategory = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetMerchantName(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.MerchantName = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetOrderId(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.OrderId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetOrderPrice(v float64) *MealBillSettlementQueryResponseBodyModuleItems {
	s.OrderPrice = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetOrderStatusDesc(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.OrderStatusDesc = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetPersonSettlePrice(v float64) *MealBillSettlementQueryResponseBodyModuleItems {
	s.PersonSettlePrice = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetPrimaryId(v int64) *MealBillSettlementQueryResponseBodyModuleItems {
	s.PrimaryId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetProcessorOaCode(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ProcessorOaCode = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetProjectCode(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ProjectCode = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetProjectName(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ProjectName = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetRemark(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.Remark = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetServiceFee(v float64) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ServiceFee = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetSettleTypeDesc(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.SettleTypeDesc = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetSettlementFee(v float64) *MealBillSettlementQueryResponseBodyModuleItems {
	s.SettlementFee = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetSettlementTime(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.SettlementTime = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetSettlementType(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.SettlementType = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetShareDinner(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ShareDinner = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetStatus(v int32) *MealBillSettlementQueryResponseBodyModuleItems {
	s.Status = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetStatusDesc(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.StatusDesc = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetStoreAddress(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.StoreAddress = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetTaxRate(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.TaxRate = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetThirdInvoiceId(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ThirdInvoiceId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetThirdPartBusinessId(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ThirdPartBusinessId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetThirdpartApplyId(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.ThirdpartApplyId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetTradeActionDesc(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.TradeActionDesc = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetTravelerId(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.TravelerId = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetTravelerJobNo(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.TravelerJobNo = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberType(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberType = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberTypeName(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberTypeName = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetTravelerName(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.TravelerName = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetVoucherType(v int32) *MealBillSettlementQueryResponseBodyModuleItems {
	s.VoucherType = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) SetVoucherTypeDesc(v string) *MealBillSettlementQueryResponseBodyModuleItems {
	s.VoucherTypeDesc = &v
	return s
}

func (s *MealBillSettlementQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}
