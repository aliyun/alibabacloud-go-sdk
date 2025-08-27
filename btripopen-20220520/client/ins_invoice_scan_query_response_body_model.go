// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsInvoiceScanQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InsInvoiceScanQueryResponseBody
	GetCode() *string
	SetMessage(v string) *InsInvoiceScanQueryResponseBody
	GetMessage() *string
	SetModule(v *InsInvoiceScanQueryResponseBodyModule) *InsInvoiceScanQueryResponseBody
	GetModule() *InsInvoiceScanQueryResponseBodyModule
	SetRequestId(v string) *InsInvoiceScanQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InsInvoiceScanQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *InsInvoiceScanQueryResponseBody
	GetTraceId() *string
}

type InsInvoiceScanQueryResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  *InsInvoiceScanQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
	// requestId
	//
	// example:
	//
	// C61ECFF6-606B-5F66-B81D-D77369043A5F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// traceId
	//
	// example:
	//
	// 210e847f16611516748613869de4f6
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s InsInvoiceScanQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsInvoiceScanQueryResponseBody) GoString() string {
	return s.String()
}

func (s *InsInvoiceScanQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *InsInvoiceScanQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsInvoiceScanQueryResponseBody) GetModule() *InsInvoiceScanQueryResponseBodyModule {
	return s.Module
}

func (s *InsInvoiceScanQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsInvoiceScanQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InsInvoiceScanQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *InsInvoiceScanQueryResponseBody) SetCode(v string) *InsInvoiceScanQueryResponseBody {
	s.Code = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBody) SetMessage(v string) *InsInvoiceScanQueryResponseBody {
	s.Message = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBody) SetModule(v *InsInvoiceScanQueryResponseBodyModule) *InsInvoiceScanQueryResponseBody {
	s.Module = v
	return s
}

func (s *InsInvoiceScanQueryResponseBody) SetRequestId(v string) *InsInvoiceScanQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBody) SetSuccess(v bool) *InsInvoiceScanQueryResponseBody {
	s.Success = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBody) SetTraceId(v string) *InsInvoiceScanQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type InsInvoiceScanQueryResponseBodyModule struct {
	Items []*InsInvoiceScanQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int32 `json:"page_no,omitempty" xml:"page_no,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// example:
	//
	// 2
	TotalPage *int32 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// example:
	//
	// 30
	TotalSize *int32 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s InsInvoiceScanQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s InsInvoiceScanQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *InsInvoiceScanQueryResponseBodyModule) GetItems() []*InsInvoiceScanQueryResponseBodyModuleItems {
	return s.Items
}

func (s *InsInvoiceScanQueryResponseBodyModule) GetPageNo() *int32 {
	return s.PageNo
}

func (s *InsInvoiceScanQueryResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *InsInvoiceScanQueryResponseBodyModule) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *InsInvoiceScanQueryResponseBodyModule) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *InsInvoiceScanQueryResponseBodyModule) SetItems(v []*InsInvoiceScanQueryResponseBodyModuleItems) *InsInvoiceScanQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModule) SetPageNo(v int32) *InsInvoiceScanQueryResponseBodyModule {
	s.PageNo = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModule) SetPageSize(v int32) *InsInvoiceScanQueryResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModule) SetTotalPage(v int32) *InsInvoiceScanQueryResponseBodyModule {
	s.TotalPage = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModule) SetTotalSize(v int32) *InsInvoiceScanQueryResponseBodyModule {
	s.TotalSize = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type InsInvoiceScanQueryResponseBodyModuleItems struct {
	// example:
	//
	// 20
	AmountWithTax *string `json:"amount_with_tax,omitempty" xml:"amount_with_tax,omitempty"`
	// example:
	//
	// 18.87
	AmountWithoutTax *string `json:"amount_without_tax,omitempty" xml:"amount_without_tax,omitempty"`
	// example:
	//
	// 2022-12-01
	BillDate *string `json:"bill_date,omitempty" xml:"bill_date,omitempty"`
	Category *int32  `json:"category,omitempty" xml:"category,omitempty"`
	// 校验码
	//
	// example:
	//
	// 07122942791187744475
	CheckCode  *string `json:"check_code,omitempty" xml:"check_code,omitempty"`
	CostCenter *string `json:"cost_center,omitempty" xml:"cost_center,omitempty"`
	Department *string `json:"department,omitempty" xml:"department,omitempty"`
	// 开票人
	Drawer *string `json:"drawer,omitempty" xml:"drawer,omitempty"`
	// 应用ID
	//
	// example:
	//
	// 566
	Id               *string `json:"id,omitempty" xml:"id,omitempty"`
	InsuranceCompany *string `json:"insurance_company,omitempty" xml:"insurance_company,omitempty"`
	// example:
	//
	// T123343234242323232
	InsuranceOrderId *string `json:"insurance_order_id,omitempty" xml:"insurance_order_id,omitempty"`
	InsuranceType    *string `json:"insurance_type,omitempty" xml:"insurance_type,omitempty"`
	// example:
	//
	// 3300111303
	InvoiceCode *string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// example:
	//
	// 2022-12-01
	InvoiceDay *string `json:"invoice_day,omitempty" xml:"invoice_day,omitempty"`
	// 发票明细
	InvoiceDetails []*InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails `json:"invoice_details,omitempty" xml:"invoice_details,omitempty" type:"Repeated"`
	// 发票地区
	InvoiceLocation *string `json:"invoice_location,omitempty" xml:"invoice_location,omitempty"`
	// example:
	//
	// 24021111
	InvoiceNo *string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// 发票抬头
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	InvoiceType  *int32  `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 机器码
	//
	// example:
	//
	// 661619906841
	MachineCode *string `json:"machine_code,omitempty" xml:"machine_code,omitempty"`
	OfdOssUrl   *string `json:"ofd_oss_url,omitempty" xml:"ofd_oss_url,omitempty"`
	// example:
	//
	// 3137168772101111000
	OrderId *int64 `json:"order_id,omitempty" xml:"order_id,omitempty"`
	// example:
	//
	// https://www.testurl.com
	OssUrl    *string `json:"oss_url,omitempty" xml:"oss_url,omitempty"`
	Passenger *string `json:"passenger,omitempty" xml:"passenger,omitempty"`
	// 密码区
	//
	// example:
	//
	// <87*>>53>5023>-446>/4+83/5	- *>5/81<75/1931>4>>
	PasswordArea *string `json:"password_area,omitempty" xml:"password_area,omitempty"`
	Project      *string `json:"project,omitempty" xml:"project,omitempty"`
	// 购方银行信息
	PurchaserBankAccountInfo *string `json:"purchaser_bank_account_info,omitempty" xml:"purchaser_bank_account_info,omitempty"`
	// 购方联系方式
	PurchaserContactInfo *string `json:"purchaser_contact_info,omitempty" xml:"purchaser_contact_info,omitempty"`
	PurchaserName        *string `json:"purchaser_name,omitempty" xml:"purchaser_name,omitempty"`
	// example:
	//
	// 91441111111111111S
	PurchaserTaxNo *string `json:"purchaser_tax_no,omitempty" xml:"purchaser_tax_no,omitempty"`
	// 收款人
	Recipient *string `json:"recipient,omitempty" xml:"recipient,omitempty"`
	// 备注
	Remarks *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
	// 复核人
	Reviewer *string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	// 销售方银行信息
	SellerBankAccountInfo *string `json:"seller_bank_account_info,omitempty" xml:"seller_bank_account_info,omitempty"`
	// 销售方联系方式
	SellerContactInfo *string `json:"seller_contact_info,omitempty" xml:"seller_contact_info,omitempty"`
	SellerName        *string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// example:
	//
	// 91441111111111111N
	SellerTaxNo *string `json:"seller_tax_no,omitempty" xml:"seller_tax_no,omitempty"`
	// 校验码
	//
	// example:
	//
	// 4475
	SmartCheckCode *string `json:"smart_check_code,omitempty" xml:"smart_check_code,omitempty"`
	// example:
	//
	// 1.13
	TaxAmount *string `json:"tax_amount,omitempty" xml:"tax_amount,omitempty"`
	// example:
	//
	// 6%
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 大写金额
	TotalAmountInWords *string `json:"total_amount_in_words,omitempty" xml:"total_amount_in_words,omitempty"`
	XmlOssUrl          *string `json:"xml_oss_url,omitempty" xml:"xml_oss_url,omitempty"`
}

func (s InsInvoiceScanQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s InsInvoiceScanQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetAmountWithTax() *string {
	return s.AmountWithTax
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetAmountWithoutTax() *string {
	return s.AmountWithoutTax
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetBillDate() *string {
	return s.BillDate
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetCategory() *int32 {
	return s.Category
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetCheckCode() *string {
	return s.CheckCode
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetCostCenter() *string {
	return s.CostCenter
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetDepartment() *string {
	return s.Department
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetDrawer() *string {
	return s.Drawer
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetId() *string {
	return s.Id
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetInsuranceCompany() *string {
	return s.InsuranceCompany
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetInsuranceOrderId() *string {
	return s.InsuranceOrderId
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetInsuranceType() *string {
	return s.InsuranceType
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetInvoiceCode() *string {
	return s.InvoiceCode
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetInvoiceDay() *string {
	return s.InvoiceDay
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetInvoiceDetails() []*InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	return s.InvoiceDetails
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetInvoiceLocation() *string {
	return s.InvoiceLocation
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetInvoiceNo() *string {
	return s.InvoiceNo
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetMachineCode() *string {
	return s.MachineCode
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetOfdOssUrl() *string {
	return s.OfdOssUrl
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetOrderId() *int64 {
	return s.OrderId
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetOssUrl() *string {
	return s.OssUrl
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetPassenger() *string {
	return s.Passenger
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetPasswordArea() *string {
	return s.PasswordArea
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetProject() *string {
	return s.Project
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetPurchaserBankAccountInfo() *string {
	return s.PurchaserBankAccountInfo
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetPurchaserContactInfo() *string {
	return s.PurchaserContactInfo
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetPurchaserName() *string {
	return s.PurchaserName
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetPurchaserTaxNo() *string {
	return s.PurchaserTaxNo
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetRecipient() *string {
	return s.Recipient
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetRemarks() *string {
	return s.Remarks
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetReviewer() *string {
	return s.Reviewer
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetSellerBankAccountInfo() *string {
	return s.SellerBankAccountInfo
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetSellerContactInfo() *string {
	return s.SellerContactInfo
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetSellerName() *string {
	return s.SellerName
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetSellerTaxNo() *string {
	return s.SellerTaxNo
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetSmartCheckCode() *string {
	return s.SmartCheckCode
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetTaxRate() *string {
	return s.TaxRate
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetTotalAmountInWords() *string {
	return s.TotalAmountInWords
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) GetXmlOssUrl() *string {
	return s.XmlOssUrl
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetAmountWithTax(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.AmountWithTax = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetAmountWithoutTax(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.AmountWithoutTax = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetBillDate(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.BillDate = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetCategory(v int32) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.Category = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetCheckCode(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.CheckCode = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetCostCenter(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.CostCenter = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetDepartment(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.Department = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetDrawer(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.Drawer = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetId(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.Id = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetInsuranceCompany(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.InsuranceCompany = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetInsuranceOrderId(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.InsuranceOrderId = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetInsuranceType(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.InsuranceType = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetInvoiceCode(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceCode = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetInvoiceDay(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceDay = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetInvoiceDetails(v []*InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceDetails = v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetInvoiceLocation(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceLocation = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetInvoiceNo(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceNo = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetInvoiceTitle(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceTitle = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetInvoiceType(v int32) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceType = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetMachineCode(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.MachineCode = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetOfdOssUrl(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.OfdOssUrl = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetOrderId(v int64) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.OrderId = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetOssUrl(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.OssUrl = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetPassenger(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.Passenger = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetPasswordArea(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.PasswordArea = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetProject(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.Project = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetPurchaserBankAccountInfo(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.PurchaserBankAccountInfo = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetPurchaserContactInfo(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.PurchaserContactInfo = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetPurchaserName(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.PurchaserName = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetPurchaserTaxNo(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.PurchaserTaxNo = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetRecipient(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.Recipient = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetRemarks(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.Remarks = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetReviewer(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.Reviewer = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetSellerBankAccountInfo(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.SellerBankAccountInfo = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetSellerContactInfo(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.SellerContactInfo = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetSellerName(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.SellerName = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetSellerTaxNo(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.SellerTaxNo = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetSmartCheckCode(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.SmartCheckCode = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetTaxAmount(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.TaxAmount = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetTaxRate(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.TaxRate = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetTotalAmountInWords(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.TotalAmountInWords = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) SetXmlOssUrl(v string) *InsInvoiceScanQueryResponseBodyModuleItems {
	s.XmlOssUrl = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}

type InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails struct {
	// 金额
	//
	// example:
	//
	// 75.21
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 行号
	//
	// example:
	//
	// 0
	Index *string `json:"index,omitempty" xml:"index,omitempty"`
	// 货物或应税劳务、服务名称
	ItemName *string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// 数量
	//
	// example:
	//
	// 1
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// 规格型号
	//
	// example:
	//
	// 66PT11230069
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// 税额
	//
	// example:
	//
	// 12.79
	Tax *string `json:"tax,omitempty" xml:"tax,omitempty"`
	// 税率
	//
	// example:
	//
	// 17%
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	// 单位
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// 单价
	//
	// example:
	//
	// 75.21
	UnitPrice *string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
}

func (s InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) String() string {
	return dara.Prettify(s)
}

func (s InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GoString() string {
	return s.String()
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetAmount() *string {
	return s.Amount
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetIndex() *string {
	return s.Index
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetItemName() *string {
	return s.ItemName
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetQuantity() *string {
	return s.Quantity
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetSpecification() *string {
	return s.Specification
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetTax() *string {
	return s.Tax
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetTaxRate() *string {
	return s.TaxRate
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetUnit() *string {
	return s.Unit
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetUnitPrice() *string {
	return s.UnitPrice
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetAmount(v string) *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Amount = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetIndex(v string) *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Index = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetItemName(v string) *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.ItemName = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetQuantity(v string) *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Quantity = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetSpecification(v string) *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Specification = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetTax(v string) *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Tax = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetTaxRate(v string) *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.TaxRate = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetUnit(v string) *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Unit = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetUnitPrice(v string) *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.UnitPrice = &v
	return s
}

func (s *InsInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) Validate() error {
	return dara.Validate(s)
}
