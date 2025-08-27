// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReimbursementOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryReimbursementOrderResponseBody
	GetCode() *string
	SetMessage(v string) *QueryReimbursementOrderResponseBody
	GetMessage() *string
	SetModule(v *QueryReimbursementOrderResponseBodyModule) *QueryReimbursementOrderResponseBody
	GetModule() *QueryReimbursementOrderResponseBodyModule
	SetRequestId(v string) *QueryReimbursementOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryReimbursementOrderResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *QueryReimbursementOrderResponseBody
	GetTraceId() *string
}

type QueryReimbursementOrderResponseBody struct {
	// example:
	//
	// PARAM_ERROR
	Code    *string                                    `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                    `json:"message,omitempty" xml:"message,omitempty"`
	Module  *QueryReimbursementOrderResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// B72B39C8-32DE-558D-AD1C-D53F11F6ADFE
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

func (s QueryReimbursementOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryReimbursementOrderResponseBody) GoString() string {
	return s.String()
}

func (s *QueryReimbursementOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryReimbursementOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryReimbursementOrderResponseBody) GetModule() *QueryReimbursementOrderResponseBodyModule {
	return s.Module
}

func (s *QueryReimbursementOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryReimbursementOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryReimbursementOrderResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *QueryReimbursementOrderResponseBody) SetCode(v string) *QueryReimbursementOrderResponseBody {
	s.Code = &v
	return s
}

func (s *QueryReimbursementOrderResponseBody) SetMessage(v string) *QueryReimbursementOrderResponseBody {
	s.Message = &v
	return s
}

func (s *QueryReimbursementOrderResponseBody) SetModule(v *QueryReimbursementOrderResponseBodyModule) *QueryReimbursementOrderResponseBody {
	s.Module = v
	return s
}

func (s *QueryReimbursementOrderResponseBody) SetRequestId(v string) *QueryReimbursementOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryReimbursementOrderResponseBody) SetSuccess(v bool) *QueryReimbursementOrderResponseBody {
	s.Success = &v
	return s
}

func (s *QueryReimbursementOrderResponseBody) SetTraceId(v string) *QueryReimbursementOrderResponseBody {
	s.TraceId = &v
	return s
}

func (s *QueryReimbursementOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryReimbursementOrderResponseBodyModule struct {
	// example:
	//
	// 10.00
	CompanyAmount *string `json:"company_amount,omitempty" xml:"company_amount,omitempty"`
	// example:
	//
	// 20.00
	CompanyPayAmount *string `json:"company_pay_amount,omitempty" xml:"company_pay_amount,omitempty"`
	// example:
	//
	// dinga809ed71b9201f35
	CorpId                    *string                                              `json:"corp_id,omitempty" xml:"corp_id,omitempty"`
	CostCenterCode            *string                                              `json:"cost_center_code,omitempty" xml:"cost_center_code,omitempty"`
	CostCenterName            *string                                              `json:"cost_center_name,omitempty" xml:"cost_center_name,omitempty"`
	Expenses                  []*QueryReimbursementOrderResponseBodyModuleExpenses `json:"expenses,omitempty" xml:"expenses,omitempty" type:"Repeated"`
	ExpensesCoverDeptId       *string                                              `json:"expenses_cover_dept_id,omitempty" xml:"expenses_cover_dept_id,omitempty"`
	ExpensesCoverDeptName     *string                                              `json:"expenses_cover_dept_name,omitempty" xml:"expenses_cover_dept_name,omitempty"`
	ExpensesCoverInvoiceTitle *string                                              `json:"expenses_cover_invoice_title,omitempty" xml:"expenses_cover_invoice_title,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	GmtCreate *string `json:"gmt_create,omitempty" xml:"gmt_create,omitempty"`
	// example:
	//
	// 2022-07-20T10:40Z
	GmtModified       *string                                                  `json:"gmt_modified,omitempty" xml:"gmt_modified,omitempty"`
	IsDeleted         *string                                                  `json:"is_deleted,omitempty" xml:"is_deleted,omitempty"`
	Itineraries       []*QueryReimbursementOrderResponseBodyModuleItineraries  `json:"itineraries,omitempty" xml:"itineraries,omitempty" type:"Repeated"`
	PaymentFinishTime *string                                                  `json:"payment_finish_time,omitempty" xml:"payment_finish_time,omitempty"`
	PaymentInfos      []*QueryReimbursementOrderResponseBodyModulePaymentInfos `json:"payment_infos,omitempty" xml:"payment_infos,omitempty" type:"Repeated"`
	// example:
	//
	// 10.00
	PersonalAmount *string `json:"personal_amount,omitempty" xml:"personal_amount,omitempty"`
	ProcessEndTime *string `json:"process_end_time,omitempty" xml:"process_end_time,omitempty"`
	ProjectCode    *string `json:"project_code,omitempty" xml:"project_code,omitempty"`
	ProjectName    *string `json:"project_name,omitempty" xml:"project_name,omitempty"`
	Reason         *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// example:
	//
	// RT203956
	ReimbursementNo    *string `json:"reimbursement_no,omitempty" xml:"reimbursement_no,omitempty"`
	Remark             *string `json:"remark,omitempty" xml:"remark,omitempty"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
	TravelThirdApplyId *string `json:"travel_third_apply_id,omitempty" xml:"travel_third_apply_id,omitempty"`
	// example:
	//
	// userId
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s QueryReimbursementOrderResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryReimbursementOrderResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryReimbursementOrderResponseBodyModule) GetCompanyAmount() *string {
	return s.CompanyAmount
}

func (s *QueryReimbursementOrderResponseBodyModule) GetCompanyPayAmount() *string {
	return s.CompanyPayAmount
}

func (s *QueryReimbursementOrderResponseBodyModule) GetCorpId() *string {
	return s.CorpId
}

func (s *QueryReimbursementOrderResponseBodyModule) GetCostCenterCode() *string {
	return s.CostCenterCode
}

func (s *QueryReimbursementOrderResponseBodyModule) GetCostCenterName() *string {
	return s.CostCenterName
}

func (s *QueryReimbursementOrderResponseBodyModule) GetExpenses() []*QueryReimbursementOrderResponseBodyModuleExpenses {
	return s.Expenses
}

func (s *QueryReimbursementOrderResponseBodyModule) GetExpensesCoverDeptId() *string {
	return s.ExpensesCoverDeptId
}

func (s *QueryReimbursementOrderResponseBodyModule) GetExpensesCoverDeptName() *string {
	return s.ExpensesCoverDeptName
}

func (s *QueryReimbursementOrderResponseBodyModule) GetExpensesCoverInvoiceTitle() *string {
	return s.ExpensesCoverInvoiceTitle
}

func (s *QueryReimbursementOrderResponseBodyModule) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryReimbursementOrderResponseBodyModule) GetGmtModified() *string {
	return s.GmtModified
}

func (s *QueryReimbursementOrderResponseBodyModule) GetIsDeleted() *string {
	return s.IsDeleted
}

func (s *QueryReimbursementOrderResponseBodyModule) GetItineraries() []*QueryReimbursementOrderResponseBodyModuleItineraries {
	return s.Itineraries
}

func (s *QueryReimbursementOrderResponseBodyModule) GetPaymentFinishTime() *string {
	return s.PaymentFinishTime
}

func (s *QueryReimbursementOrderResponseBodyModule) GetPaymentInfos() []*QueryReimbursementOrderResponseBodyModulePaymentInfos {
	return s.PaymentInfos
}

func (s *QueryReimbursementOrderResponseBodyModule) GetPersonalAmount() *string {
	return s.PersonalAmount
}

func (s *QueryReimbursementOrderResponseBodyModule) GetProcessEndTime() *string {
	return s.ProcessEndTime
}

func (s *QueryReimbursementOrderResponseBodyModule) GetProjectCode() *string {
	return s.ProjectCode
}

func (s *QueryReimbursementOrderResponseBodyModule) GetProjectName() *string {
	return s.ProjectName
}

func (s *QueryReimbursementOrderResponseBodyModule) GetReason() *string {
	return s.Reason
}

func (s *QueryReimbursementOrderResponseBodyModule) GetReimbursementNo() *string {
	return s.ReimbursementNo
}

func (s *QueryReimbursementOrderResponseBodyModule) GetRemark() *string {
	return s.Remark
}

func (s *QueryReimbursementOrderResponseBodyModule) GetStatus() *string {
	return s.Status
}

func (s *QueryReimbursementOrderResponseBodyModule) GetTravelThirdApplyId() *string {
	return s.TravelThirdApplyId
}

func (s *QueryReimbursementOrderResponseBodyModule) GetUserId() *string {
	return s.UserId
}

func (s *QueryReimbursementOrderResponseBodyModule) GetUserName() *string {
	return s.UserName
}

func (s *QueryReimbursementOrderResponseBodyModule) SetCompanyAmount(v string) *QueryReimbursementOrderResponseBodyModule {
	s.CompanyAmount = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetCompanyPayAmount(v string) *QueryReimbursementOrderResponseBodyModule {
	s.CompanyPayAmount = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetCorpId(v string) *QueryReimbursementOrderResponseBodyModule {
	s.CorpId = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetCostCenterCode(v string) *QueryReimbursementOrderResponseBodyModule {
	s.CostCenterCode = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetCostCenterName(v string) *QueryReimbursementOrderResponseBodyModule {
	s.CostCenterName = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetExpenses(v []*QueryReimbursementOrderResponseBodyModuleExpenses) *QueryReimbursementOrderResponseBodyModule {
	s.Expenses = v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetExpensesCoverDeptId(v string) *QueryReimbursementOrderResponseBodyModule {
	s.ExpensesCoverDeptId = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetExpensesCoverDeptName(v string) *QueryReimbursementOrderResponseBodyModule {
	s.ExpensesCoverDeptName = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetExpensesCoverInvoiceTitle(v string) *QueryReimbursementOrderResponseBodyModule {
	s.ExpensesCoverInvoiceTitle = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetGmtCreate(v string) *QueryReimbursementOrderResponseBodyModule {
	s.GmtCreate = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetGmtModified(v string) *QueryReimbursementOrderResponseBodyModule {
	s.GmtModified = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetIsDeleted(v string) *QueryReimbursementOrderResponseBodyModule {
	s.IsDeleted = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetItineraries(v []*QueryReimbursementOrderResponseBodyModuleItineraries) *QueryReimbursementOrderResponseBodyModule {
	s.Itineraries = v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetPaymentFinishTime(v string) *QueryReimbursementOrderResponseBodyModule {
	s.PaymentFinishTime = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetPaymentInfos(v []*QueryReimbursementOrderResponseBodyModulePaymentInfos) *QueryReimbursementOrderResponseBodyModule {
	s.PaymentInfos = v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetPersonalAmount(v string) *QueryReimbursementOrderResponseBodyModule {
	s.PersonalAmount = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetProcessEndTime(v string) *QueryReimbursementOrderResponseBodyModule {
	s.ProcessEndTime = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetProjectCode(v string) *QueryReimbursementOrderResponseBodyModule {
	s.ProjectCode = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetProjectName(v string) *QueryReimbursementOrderResponseBodyModule {
	s.ProjectName = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetReason(v string) *QueryReimbursementOrderResponseBodyModule {
	s.Reason = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetReimbursementNo(v string) *QueryReimbursementOrderResponseBodyModule {
	s.ReimbursementNo = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetRemark(v string) *QueryReimbursementOrderResponseBodyModule {
	s.Remark = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetStatus(v string) *QueryReimbursementOrderResponseBodyModule {
	s.Status = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetTravelThirdApplyId(v string) *QueryReimbursementOrderResponseBodyModule {
	s.TravelThirdApplyId = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetUserId(v string) *QueryReimbursementOrderResponseBodyModule {
	s.UserId = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) SetUserName(v string) *QueryReimbursementOrderResponseBodyModule {
	s.UserName = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type QueryReimbursementOrderResponseBodyModuleExpenses struct {
	// example:
	//
	// 20
	Amount              *string                                                                 `json:"amount,omitempty" xml:"amount,omitempty"`
	Currency            *string                                                                 `json:"currency,omitempty" xml:"currency,omitempty"`
	ExpenseCity         *string                                                                 `json:"expense_city,omitempty" xml:"expense_city,omitempty"`
	ExpenseCompositions []*QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions `json:"expense_compositions,omitempty" xml:"expense_compositions,omitempty" type:"Repeated"`
	// example:
	//
	// 2022-05-15T22:27Z
	ExpenseTime *string `json:"expense_time,omitempty" xml:"expense_time,omitempty"`
	ExpenseType *string `json:"expense_type,omitempty" xml:"expense_type,omitempty"`
	// example:
	//
	// code
	ExpenseTypeCode *string                                                          `json:"expense_type_code,omitempty" xml:"expense_type_code,omitempty"`
	InvoiceInfos    []*QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos `json:"invoice_infos,omitempty" xml:"invoice_infos,omitempty" type:"Repeated"`
	ReimbExpenseId  *int64                                                           `json:"reimb_expense_id,omitempty" xml:"reimb_expense_id,omitempty"`
	Remark          *string                                                          `json:"remark,omitempty" xml:"remark,omitempty"`
	SettlementType  *string                                                          `json:"settlement_type,omitempty" xml:"settlement_type,omitempty"`
}

func (s QueryReimbursementOrderResponseBodyModuleExpenses) String() string {
	return dara.Prettify(s)
}

func (s QueryReimbursementOrderResponseBodyModuleExpenses) GoString() string {
	return s.String()
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) GetAmount() *string {
	return s.Amount
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) GetCurrency() *string {
	return s.Currency
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) GetExpenseCity() *string {
	return s.ExpenseCity
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) GetExpenseCompositions() []*QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions {
	return s.ExpenseCompositions
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) GetExpenseTime() *string {
	return s.ExpenseTime
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) GetExpenseType() *string {
	return s.ExpenseType
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) GetExpenseTypeCode() *string {
	return s.ExpenseTypeCode
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) GetInvoiceInfos() []*QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos {
	return s.InvoiceInfos
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) GetReimbExpenseId() *int64 {
	return s.ReimbExpenseId
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) GetRemark() *string {
	return s.Remark
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) GetSettlementType() *string {
	return s.SettlementType
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) SetAmount(v string) *QueryReimbursementOrderResponseBodyModuleExpenses {
	s.Amount = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) SetCurrency(v string) *QueryReimbursementOrderResponseBodyModuleExpenses {
	s.Currency = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) SetExpenseCity(v string) *QueryReimbursementOrderResponseBodyModuleExpenses {
	s.ExpenseCity = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) SetExpenseCompositions(v []*QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) *QueryReimbursementOrderResponseBodyModuleExpenses {
	s.ExpenseCompositions = v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) SetExpenseTime(v string) *QueryReimbursementOrderResponseBodyModuleExpenses {
	s.ExpenseTime = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) SetExpenseType(v string) *QueryReimbursementOrderResponseBodyModuleExpenses {
	s.ExpenseType = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) SetExpenseTypeCode(v string) *QueryReimbursementOrderResponseBodyModuleExpenses {
	s.ExpenseTypeCode = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) SetInvoiceInfos(v []*QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) *QueryReimbursementOrderResponseBodyModuleExpenses {
	s.InvoiceInfos = v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) SetReimbExpenseId(v int64) *QueryReimbursementOrderResponseBodyModuleExpenses {
	s.ReimbExpenseId = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) SetRemark(v string) *QueryReimbursementOrderResponseBodyModuleExpenses {
	s.Remark = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) SetSettlementType(v string) *QueryReimbursementOrderResponseBodyModuleExpenses {
	s.SettlementType = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpenses) Validate() error {
	return dara.Validate(s)
}

type QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions struct {
	BillSettlementId *int64    `json:"bill_settlement_id,omitempty" xml:"bill_settlement_id,omitempty"`
	CapitalDirection *string   `json:"capital_direction,omitempty" xml:"capital_direction,omitempty"`
	FeeType          *string   `json:"fee_type,omitempty" xml:"fee_type,omitempty"`
	OrderId          *string   `json:"order_id,omitempty" xml:"order_id,omitempty"`
	Remark           *string   `json:"remark,omitempty" xml:"remark,omitempty"`
	RemindTagList    []*string `json:"remind_tag_list,omitempty" xml:"remind_tag_list,omitempty" type:"Repeated"`
	SettlementAmount *string   `json:"settlement_amount,omitempty" xml:"settlement_amount,omitempty"`
	SettlementTime   *string   `json:"settlement_time,omitempty" xml:"settlement_time,omitempty"`
	VoucherType      *int32    `json:"voucher_type,omitempty" xml:"voucher_type,omitempty"`
}

func (s QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) String() string {
	return dara.Prettify(s)
}

func (s QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) GoString() string {
	return s.String()
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) GetBillSettlementId() *int64 {
	return s.BillSettlementId
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) GetCapitalDirection() *string {
	return s.CapitalDirection
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) GetFeeType() *string {
	return s.FeeType
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) GetOrderId() *string {
	return s.OrderId
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) GetRemark() *string {
	return s.Remark
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) GetRemindTagList() []*string {
	return s.RemindTagList
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) GetSettlementAmount() *string {
	return s.SettlementAmount
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) GetSettlementTime() *string {
	return s.SettlementTime
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) GetVoucherType() *int32 {
	return s.VoucherType
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) SetBillSettlementId(v int64) *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions {
	s.BillSettlementId = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) SetCapitalDirection(v string) *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions {
	s.CapitalDirection = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) SetFeeType(v string) *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions {
	s.FeeType = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) SetOrderId(v string) *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions {
	s.OrderId = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) SetRemark(v string) *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions {
	s.Remark = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) SetRemindTagList(v []*string) *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions {
	s.RemindTagList = v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) SetSettlementAmount(v string) *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions {
	s.SettlementAmount = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) SetSettlementTime(v string) *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions {
	s.SettlementTime = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) SetVoucherType(v int32) *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions {
	s.VoucherType = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesExpenseCompositions) Validate() error {
	return dara.Validate(s)
}

type QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos struct {
	Amount        *string `json:"amount,omitempty" xml:"amount,omitempty"`
	InvoiceCode   *string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	InvoiceData   *string `json:"invoice_data,omitempty" xml:"invoice_data,omitempty"`
	InvoiceDate   *string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty"`
	InvoiceNumber *string `json:"invoice_number,omitempty" xml:"invoice_number,omitempty"`
	InvoiceType   *string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
}

func (s QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) String() string {
	return dara.Prettify(s)
}

func (s QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) GoString() string {
	return s.String()
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) GetAmount() *string {
	return s.Amount
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) GetInvoiceCode() *string {
	return s.InvoiceCode
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) GetInvoiceData() *string {
	return s.InvoiceData
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) GetInvoiceDate() *string {
	return s.InvoiceDate
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) GetInvoiceNumber() *string {
	return s.InvoiceNumber
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) GetInvoiceType() *string {
	return s.InvoiceType
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) SetAmount(v string) *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos {
	s.Amount = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) SetInvoiceCode(v string) *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos {
	s.InvoiceCode = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) SetInvoiceData(v string) *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos {
	s.InvoiceData = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) SetInvoiceDate(v string) *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos {
	s.InvoiceDate = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) SetInvoiceNumber(v string) *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos {
	s.InvoiceNumber = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) SetInvoiceType(v string) *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos {
	s.InvoiceType = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleExpensesInvoiceInfos) Validate() error {
	return dara.Validate(s)
}

type QueryReimbursementOrderResponseBodyModuleItineraries struct {
	ArrCity *string `json:"arr_city,omitempty" xml:"arr_city,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	ArrDate *string `json:"arr_date,omitempty" xml:"arr_date,omitempty"`
	// example:
	//
	// CTU
	DepCity *string `json:"dep_city,omitempty" xml:"dep_city,omitempty"`
	// example:
	//
	// 2022-05-15T22:27Z
	DepDate    *string `json:"dep_date,omitempty" xml:"dep_date,omitempty"`
	TrafficWay *string `json:"traffic_way,omitempty" xml:"traffic_way,omitempty"`
	TripWay    *string `json:"trip_way,omitempty" xml:"trip_way,omitempty"`
}

func (s QueryReimbursementOrderResponseBodyModuleItineraries) String() string {
	return dara.Prettify(s)
}

func (s QueryReimbursementOrderResponseBodyModuleItineraries) GoString() string {
	return s.String()
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) GetArrCity() *string {
	return s.ArrCity
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) GetArrDate() *string {
	return s.ArrDate
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) GetDepCity() *string {
	return s.DepCity
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) GetDepDate() *string {
	return s.DepDate
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) GetTrafficWay() *string {
	return s.TrafficWay
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) GetTripWay() *string {
	return s.TripWay
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) SetArrCity(v string) *QueryReimbursementOrderResponseBodyModuleItineraries {
	s.ArrCity = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) SetArrDate(v string) *QueryReimbursementOrderResponseBodyModuleItineraries {
	s.ArrDate = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) SetDepCity(v string) *QueryReimbursementOrderResponseBodyModuleItineraries {
	s.DepCity = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) SetDepDate(v string) *QueryReimbursementOrderResponseBodyModuleItineraries {
	s.DepDate = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) SetTrafficWay(v string) *QueryReimbursementOrderResponseBodyModuleItineraries {
	s.TrafficWay = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) SetTripWay(v string) *QueryReimbursementOrderResponseBodyModuleItineraries {
	s.TripWay = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModuleItineraries) Validate() error {
	return dara.Validate(s)
}

type QueryReimbursementOrderResponseBodyModulePaymentInfos struct {
	// example:
	//
	// 20.00
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// example:
	//
	// userId
	PayeeUserId *string `json:"payee_user_id,omitempty" xml:"payee_user_id,omitempty"`
}

func (s QueryReimbursementOrderResponseBodyModulePaymentInfos) String() string {
	return dara.Prettify(s)
}

func (s QueryReimbursementOrderResponseBodyModulePaymentInfos) GoString() string {
	return s.String()
}

func (s *QueryReimbursementOrderResponseBodyModulePaymentInfos) GetAmount() *string {
	return s.Amount
}

func (s *QueryReimbursementOrderResponseBodyModulePaymentInfos) GetPayeeUserId() *string {
	return s.PayeeUserId
}

func (s *QueryReimbursementOrderResponseBodyModulePaymentInfos) SetAmount(v string) *QueryReimbursementOrderResponseBodyModulePaymentInfos {
	s.Amount = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModulePaymentInfos) SetPayeeUserId(v string) *QueryReimbursementOrderResponseBodyModulePaymentInfos {
	s.PayeeUserId = &v
	return s
}

func (s *QueryReimbursementOrderResponseBodyModulePaymentInfos) Validate() error {
	return dara.Validate(s)
}
