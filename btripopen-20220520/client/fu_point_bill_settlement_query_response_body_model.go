// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFuPointBillSettlementQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *FuPointBillSettlementQueryResponseBody
	GetCode() *int32
	SetModule(v *FuPointBillSettlementQueryResponseBodyModule) *FuPointBillSettlementQueryResponseBody
	GetModule() *FuPointBillSettlementQueryResponseBodyModule
	SetMorePage(v bool) *FuPointBillSettlementQueryResponseBody
	GetMorePage() *bool
	SetRequestId(v string) *FuPointBillSettlementQueryResponseBody
	GetRequestId() *string
	SetResultMsg(v string) *FuPointBillSettlementQueryResponseBody
	GetResultMsg() *string
	SetSuccess(v bool) *FuPointBillSettlementQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *FuPointBillSettlementQueryResponseBody
	GetTraceId() *string
}

type FuPointBillSettlementQueryResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// module。
	Module *FuPointBillSettlementQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// example:
	//
	// true
	MorePage *bool `json:"more_page,omitempty" xml:"more_page,omitempty"`
	// example:
	//
	// 210bc22017109867047728291dd406
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// trace_id
	//
	// example:
	//
	// 213e382517240341253056547e41fc
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s FuPointBillSettlementQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FuPointBillSettlementQueryResponseBody) GoString() string {
	return s.String()
}

func (s *FuPointBillSettlementQueryResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *FuPointBillSettlementQueryResponseBody) GetModule() *FuPointBillSettlementQueryResponseBodyModule {
	return s.Module
}

func (s *FuPointBillSettlementQueryResponseBody) GetMorePage() *bool {
	return s.MorePage
}

func (s *FuPointBillSettlementQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FuPointBillSettlementQueryResponseBody) GetResultMsg() *string {
	return s.ResultMsg
}

func (s *FuPointBillSettlementQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FuPointBillSettlementQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *FuPointBillSettlementQueryResponseBody) SetCode(v int32) *FuPointBillSettlementQueryResponseBody {
	s.Code = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBody) SetModule(v *FuPointBillSettlementQueryResponseBodyModule) *FuPointBillSettlementQueryResponseBody {
	s.Module = v
	return s
}

func (s *FuPointBillSettlementQueryResponseBody) SetMorePage(v bool) *FuPointBillSettlementQueryResponseBody {
	s.MorePage = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBody) SetRequestId(v string) *FuPointBillSettlementQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBody) SetResultMsg(v string) *FuPointBillSettlementQueryResponseBody {
	s.ResultMsg = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBody) SetSuccess(v bool) *FuPointBillSettlementQueryResponseBody {
	s.Success = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBody) SetTraceId(v string) *FuPointBillSettlementQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type FuPointBillSettlementQueryResponseBodyModule struct {
	// example:
	//
	// 10
	Category *int32 `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// corpid
	CorpId *string                                              `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	Items  []*FuPointBillSettlementQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-07-02
	PeriodEnd *string `json:"period_end,omitempty" xml:"period_end,omitempty"`
	// example:
	//
	// 2021-10-13
	PeriodStart *string `json:"period_start,omitempty" xml:"period_start,omitempty"`
	// example:
	//
	// 1qwe
	ScrollId *string `json:"scroll_id,omitempty" xml:"scroll_id,omitempty"`
	// example:
	//
	// 30
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s FuPointBillSettlementQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s FuPointBillSettlementQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *FuPointBillSettlementQueryResponseBodyModule) GetCategory() *int32 {
	return s.Category
}

func (s *FuPointBillSettlementQueryResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *FuPointBillSettlementQueryResponseBodyModule) GetItems() []*FuPointBillSettlementQueryResponseBodyModuleItems {
	return s.Items
}

func (s *FuPointBillSettlementQueryResponseBodyModule) GetPeriodEnd() *string {
	return s.PeriodEnd
}

func (s *FuPointBillSettlementQueryResponseBodyModule) GetPeriodStart() *string {
	return s.PeriodStart
}

func (s *FuPointBillSettlementQueryResponseBodyModule) GetScrollId() *string {
	return s.ScrollId
}

func (s *FuPointBillSettlementQueryResponseBodyModule) GetTotalSize() *int64 {
	return s.TotalSize
}

func (s *FuPointBillSettlementQueryResponseBodyModule) SetCategory(v int32) *FuPointBillSettlementQueryResponseBodyModule {
	s.Category = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModule) SetCorpId(v string) *FuPointBillSettlementQueryResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModule) SetItems(v []*FuPointBillSettlementQueryResponseBodyModuleItems) *FuPointBillSettlementQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModule) SetPeriodEnd(v string) *FuPointBillSettlementQueryResponseBodyModule {
	s.PeriodEnd = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModule) SetPeriodStart(v string) *FuPointBillSettlementQueryResponseBodyModule {
	s.PeriodStart = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModule) SetScrollId(v string) *FuPointBillSettlementQueryResponseBodyModule {
	s.ScrollId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModule) SetTotalSize(v int64) *FuPointBillSettlementQueryResponseBodyModule {
	s.TotalSize = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type FuPointBillSettlementQueryResponseBodyModuleItems struct {
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
	// CD
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
	// 1424041616244499302
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
	// example:
	//
	// 10.5
	AwardNum *float64 `json:"award_num,omitempty" xml:"award_num,omitempty"`
	// example:
	//
	// 100.12
	BasisAmount *string `json:"basis_amount,omitempty" xml:"basis_amount,omitempty"`
	// example:
	//
	// 2023-01-01 00:00:00
	BillRecordTime *string `json:"bill_record_time,omitempty" xml:"bill_record_time,omitempty"`
	BillingEntity  *string `json:"billing_entity,omitempty" xml:"billing_entity,omitempty"`
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
	// zs123
	BookerJobNo       *string `json:"booker_job_no,omitempty" xml:"booker_job_no,omitempty"`
	BookerName        *string `json:"booker_name,omitempty" xml:"booker_name,omitempty"`
	CapitalDirection  *string `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	CascadeDepartment *string `json:"cascade_department,omitempty" xml:"cascade_department,omitempty"`
	CategoryDesc      *string `json:"category_desc,omitempty" xml:"category_desc,omitempty"`
	// example:
	//
	// 100.32
	CategoryOrderSettlePrice *string `json:"category_order_settle_price,omitempty" xml:"category_order_settle_price,omitempty"`
	CostCenter               *string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	// example:
	//
	// cs1
	CostCenterNumber *string `json:"cost_center_number,omitempty" xml:"cost_center_number,omitempty"`
	CostDepartment   *string `json:"cost_department,omitempty" xml:"cost_department,omitempty"`
	// example:
	//
	// 0.11
	DeductibleTax *float64 `json:"deductible_tax,omitempty" xml:"deductible_tax,omitempty"`
	Department    *string  `json:"department,omitempty" xml:"department,omitempty"`
	// example:
	//
	// 1112
	DepartmentId  *string `json:"department_id,omitempty" xml:"department_id,omitempty"`
	FeeType       *string `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	FeeTypeDesc   *string `json:"fee_type_desc,omitempty" xml:"fee_type_desc,omitempty"`
	ForeignersTag *string `json:"foreigners_tag,omitempty" xml:"foreigners_tag,omitempty"`
	// example:
	//
	// 10
	GrantNum *float64 `json:"grant_num,omitempty" xml:"grant_num,omitempty"`
	// example:
	//
	// 1
	Index        *string `json:"index,omitempty" xml:"index,omitempty"`
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// Location
	Location           *string `json:"location,omitempty" xml:"location,omitempty"`
	MakeInvoice        *string `json:"make_invoice,omitempty" xml:"make_invoice,omitempty"`
	MappingCompanyCode *string `json:"mapping_company_code,omitempty" xml:"mapping_company_code,omitempty"`
	// example:
	//
	// 4801105714092
	OrderId         *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
	OrderStatusDesc *string `json:"order_status_desc,omitempty" xml:"order_status_desc,omitempty"`
	OriginCategory  *string `json:"origin_category,omitempty" xml:"origin_category,omitempty"`
	// example:
	//
	// 111234
	OriginOrderId *string `json:"origin_order_id,omitempty" xml:"origin_order_id,omitempty"`
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
	// 93746933
	PrimaryId       *int64  `json:"primary_id,omitempty" xml:"primary_id,omitempty"`
	ProcessorOaCode *string `json:"processor_oa_code,omitempty" xml:"processor_oa_code,omitempty"`
	// example:
	//
	// acs
	ProjectCode *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	// example:
	//
	// 111224324
	PurchaseOrderId *string `json:"purchase_order_id,omitempty" xml:"purchase_order_id,omitempty"`
	Remark          *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// example:
	//
	// 10.45
	SaveAmount *float64 `json:"save_amount,omitempty" xml:"save_amount,omitempty"`
	// example:
	//
	// 1
	SceneId        *string `json:"scene_id,omitempty" xml:"scene_id,omitempty"`
	SceneName      *string `json:"scene_name,omitempty" xml:"scene_name,omitempty"`
	SettleTypeDesc *string `json:"settle_type_desc,omitempty" xml:"settle_type_desc,omitempty"`
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
	// 166564408
	ShowSubOrderId *string `json:"show_sub_order_id,omitempty" xml:"show_sub_order_id,omitempty"`
	// SIO
	//
	// example:
	//
	// SIO
	Sio *string `json:"sio,omitempty" xml:"sio,omitempty"`
	// example:
	//
	// 2
	Status     *int32  `json:"status,omitempty" xml:"status,omitempty"`
	StatusDesc *string `json:"status_desc,omitempty" xml:"status_desc,omitempty"`
	// example:
	//
	// 1019199938284381
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
	TradeActionDesc  *string `json:"trade_action_desc,omitempty" xml:"trade_action_desc,omitempty"`
	TradeReason      *string `json:"trade_reason,omitempty" xml:"trade_reason,omitempty"`
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
	// 11
	VoucherType     *int32  `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
	VoucherTypeDesc *string `json:"voucher_type_desc,omitempty" xml:"voucher_type_desc,omitempty"`
}

func (s FuPointBillSettlementQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s FuPointBillSettlementQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetAdjustTime() *string {
	return s.AdjustTime
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetAlipayId() *string {
	return s.AlipayId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetAlipayTradeNo() *string {
	return s.AlipayTradeNo
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetApplyArrCityCode() *string {
	return s.ApplyArrCityCode
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetApplyArrCityName() *string {
	return s.ApplyArrCityName
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetApplyDepCityCode() *string {
	return s.ApplyDepCityCode
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetApplyDepCityName() *string {
	return s.ApplyDepCityName
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetApplyExtendField() *string {
	return s.ApplyExtendField
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetApplyId() *string {
	return s.ApplyId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetApproverEmail() *string {
	return s.ApproverEmail
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetApproverId() *string {
	return s.ApproverId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetApproverName() *string {
	return s.ApproverName
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetAwardNum() *float64 {
	return s.AwardNum
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetBasisAmount() *string {
	return s.BasisAmount
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetBillRecordTime() *string {
	return s.BillRecordTime
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetBillingEntity() *string {
	return s.BillingEntity
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetBookMode() *string {
	return s.BookMode
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetBookTime() *string {
	return s.BookTime
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetBookerId() *string {
	return s.BookerId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetBookerJobNo() *string {
	return s.BookerJobNo
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetBookerName() *string {
	return s.BookerName
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetCascadeDepartment() *string {
	return s.CascadeDepartment
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetCategoryDesc() *string {
	return s.CategoryDesc
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetCategoryOrderSettlePrice() *string {
	return s.CategoryOrderSettlePrice
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetCostCenter() *string {
	return s.CostCenter
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetCostCenterNumber() *string {
	return s.CostCenterNumber
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetCostDepartment() *string {
	return s.CostDepartment
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetDeductibleTax() *float64 {
	return s.DeductibleTax
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetDepartment() *string {
	return s.Department
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetDepartmentId() *string {
	return s.DepartmentId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetFeeType() *string {
	return s.FeeType
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetFeeTypeDesc() *string {
	return s.FeeTypeDesc
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetForeignersTag() *string {
	return s.ForeignersTag
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetGrantNum() *float64 {
	return s.GrantNum
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetIndex() *string {
	return s.Index
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetLocation() *string {
	return s.Location
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetMakeInvoice() *string {
	return s.MakeInvoice
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetMappingCompanyCode() *string {
	return s.MappingCompanyCode
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetOrderId() *string {
	return s.OrderId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetOrderStatusDesc() *string {
	return s.OrderStatusDesc
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetOriginCategory() *string {
	return s.OriginCategory
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetOriginOrderId() *string {
	return s.OriginOrderId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetOverApplyId() *string {
	return s.OverApplyId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetPaymentDepartmentId() *string {
	return s.PaymentDepartmentId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetPaymentDepartmentName() *string {
	return s.PaymentDepartmentName
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetPrimaryId() *int64 {
	return s.PrimaryId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetProcessorOaCode() *string {
	return s.ProcessorOaCode
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetProjectName() *string {
	return s.ProjectName
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetPurchaseOrderId() *string {
	return s.PurchaseOrderId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetRemark() *string {
	return s.Remark
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetSaveAmount() *float64 {
	return s.SaveAmount
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetSceneId() *string {
	return s.SceneId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetSceneName() *string {
	return s.SceneName
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetSettleTypeDesc() *string {
	return s.SettleTypeDesc
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetSettlementFee() *float64 {
	return s.SettlementFee
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetSettlementGrantFee() *float64 {
	return s.SettlementGrantFee
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetSettlementType() *string {
	return s.SettlementType
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetShowSubOrderId() *string {
	return s.ShowSubOrderId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetSio() *string {
	return s.Sio
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetStatus() *int32 {
	return s.Status
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetSubOrderId() *string {
	return s.SubOrderId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetTaxRate() *string {
	return s.TaxRate
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetThirdInvoiceId() *string {
	return s.ThirdInvoiceId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetThirdItineraryId() *string {
	return s.ThirdItineraryId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetTradeActionDesc() *string {
	return s.TradeActionDesc
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetTradeReason() *string {
	return s.TradeReason
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetTravelerEmail() *string {
	return s.TravelerEmail
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetTravelerId() *string {
	return s.TravelerId
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetTravelerJobNo() *string {
	return s.TravelerJobNo
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberType() *string {
	return s.TravelerMemberType
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetTravelerMemberTypeName() *string {
	return s.TravelerMemberTypeName
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetTravelerName() *string {
	return s.TravelerName
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) GetVoucherTypeDesc() *string {
	return s.VoucherTypeDesc
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetAdjustTime(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.AdjustTime = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetAlipayId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.AlipayId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetAlipayTradeNo(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.AlipayTradeNo = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetApplyArrCityCode(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ApplyArrCityCode = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetApplyArrCityName(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ApplyArrCityName = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetApplyDepCityCode(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ApplyDepCityCode = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetApplyDepCityName(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ApplyDepCityName = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetApplyExtendField(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ApplyExtendField = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetApplyId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ApplyId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetApproverEmail(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ApproverEmail = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetApproverId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ApproverId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetApproverName(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ApproverName = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetAwardNum(v float64) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.AwardNum = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetBasisAmount(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.BasisAmount = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetBillRecordTime(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.BillRecordTime = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetBillingEntity(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.BillingEntity = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetBookMode(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.BookMode = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetBookTime(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.BookTime = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetBookerId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.BookerId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetBookerJobNo(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.BookerJobNo = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetBookerName(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.BookerName = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetCapitalDirection(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.CapitalDirection = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetCascadeDepartment(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.CascadeDepartment = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetCategoryDesc(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.CategoryDesc = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetCategoryOrderSettlePrice(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.CategoryOrderSettlePrice = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetCostCenter(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.CostCenter = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetCostCenterNumber(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.CostCenterNumber = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetCostDepartment(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.CostDepartment = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetDeductibleTax(v float64) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.DeductibleTax = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetDepartment(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.Department = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetDepartmentId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.DepartmentId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetFeeType(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.FeeType = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetFeeTypeDesc(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.FeeTypeDesc = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetForeignersTag(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ForeignersTag = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetGrantNum(v float64) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.GrantNum = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetIndex(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.Index = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetInvoiceTitle(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.InvoiceTitle = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetLocation(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.Location = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetMakeInvoice(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.MakeInvoice = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetMappingCompanyCode(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.MappingCompanyCode = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetOrderId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.OrderId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetOrderStatusDesc(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.OrderStatusDesc = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetOriginCategory(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.OriginCategory = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetOriginOrderId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.OriginOrderId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetOverApplyId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.OverApplyId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetPaymentDepartmentId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.PaymentDepartmentId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetPaymentDepartmentName(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.PaymentDepartmentName = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetPrimaryId(v int64) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.PrimaryId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetProcessorOaCode(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ProcessorOaCode = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetProjectCode(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ProjectCode = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetProjectName(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ProjectName = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetPurchaseOrderId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.PurchaseOrderId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetRemark(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.Remark = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetSaveAmount(v float64) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.SaveAmount = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetSceneId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.SceneId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetSceneName(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.SceneName = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetSettleTypeDesc(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.SettleTypeDesc = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetSettlementFee(v float64) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.SettlementFee = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetSettlementGrantFee(v float64) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.SettlementGrantFee = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetSettlementTime(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.SettlementTime = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetSettlementType(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.SettlementType = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetShowSubOrderId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ShowSubOrderId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetSio(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.Sio = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetStatus(v int32) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.Status = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetStatusDesc(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.StatusDesc = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetSubOrderId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.SubOrderId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetTaxRate(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.TaxRate = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetThirdInvoiceId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ThirdInvoiceId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetThirdItineraryId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.ThirdItineraryId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetTradeActionDesc(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.TradeActionDesc = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetTradeReason(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.TradeReason = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetTravelerEmail(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.TravelerEmail = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetTravelerId(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.TravelerId = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetTravelerJobNo(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.TravelerJobNo = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberType(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberType = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetTravelerMemberTypeName(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.TravelerMemberTypeName = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetTravelerName(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.TravelerName = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetVoucherType(v int32) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.VoucherType = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) SetVoucherTypeDesc(v string) *FuPointBillSettlementQueryResponseBodyModuleItems {
	s.VoucherTypeDesc = &v
	return s
}

func (s *FuPointBillSettlementQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}
