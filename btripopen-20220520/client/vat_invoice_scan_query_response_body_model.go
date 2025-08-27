// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVatInvoiceScanQueryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VatInvoiceScanQueryResponseBody
	GetCode() *string
	SetMessage(v string) *VatInvoiceScanQueryResponseBody
	GetMessage() *string
	SetModule(v *VatInvoiceScanQueryResponseBodyModule) *VatInvoiceScanQueryResponseBody
	GetModule() *VatInvoiceScanQueryResponseBodyModule
	SetRequestId(v string) *VatInvoiceScanQueryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VatInvoiceScanQueryResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *VatInvoiceScanQueryResponseBody
	GetTraceId() *string
}

type VatInvoiceScanQueryResponseBody struct {
	// example:
	//
	// 200
	Code    *string                                `json:"code,omitempty" xml:"code,omitempty"`
	Message *string                                `json:"message,omitempty" xml:"message,omitempty"`
	Module  *VatInvoiceScanQueryResponseBodyModule `json:"module,omitempty" xml:"module,omitempty" type:"Struct"`
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
	// 210e842b16611337974412836dae27
	TraceId *string `json:"traceId,omitempty" xml:"traceId,omitempty"`
}

func (s VatInvoiceScanQueryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VatInvoiceScanQueryResponseBody) GoString() string {
	return s.String()
}

func (s *VatInvoiceScanQueryResponseBody) GetCode() *string {
	return s.Code
}

func (s *VatInvoiceScanQueryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VatInvoiceScanQueryResponseBody) GetModule() *VatInvoiceScanQueryResponseBodyModule {
	return s.Module
}

func (s *VatInvoiceScanQueryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VatInvoiceScanQueryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VatInvoiceScanQueryResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *VatInvoiceScanQueryResponseBody) SetCode(v string) *VatInvoiceScanQueryResponseBody {
	s.Code = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBody) SetMessage(v string) *VatInvoiceScanQueryResponseBody {
	s.Message = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBody) SetModule(v *VatInvoiceScanQueryResponseBodyModule) *VatInvoiceScanQueryResponseBody {
	s.Module = v
	return s
}

func (s *VatInvoiceScanQueryResponseBody) SetRequestId(v string) *VatInvoiceScanQueryResponseBody {
	s.RequestId = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBody) SetSuccess(v bool) *VatInvoiceScanQueryResponseBody {
	s.Success = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBody) SetTraceId(v string) *VatInvoiceScanQueryResponseBody {
	s.TraceId = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBody) Validate() error {
	return dara.Validate(s)
}

type VatInvoiceScanQueryResponseBodyModule struct {
	Items []*VatInvoiceScanQueryResponseBodyModuleItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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

func (s VatInvoiceScanQueryResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s VatInvoiceScanQueryResponseBodyModule) GoString() string {
	return s.String()
}

func (s *VatInvoiceScanQueryResponseBodyModule) GetItems() []*VatInvoiceScanQueryResponseBodyModuleItems {
	return s.Items
}

func (s *VatInvoiceScanQueryResponseBodyModule) GetPageNo() *int32 {
	return s.PageNo
}

func (s *VatInvoiceScanQueryResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *VatInvoiceScanQueryResponseBodyModule) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *VatInvoiceScanQueryResponseBodyModule) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *VatInvoiceScanQueryResponseBodyModule) SetItems(v []*VatInvoiceScanQueryResponseBodyModuleItems) *VatInvoiceScanQueryResponseBodyModule {
	s.Items = v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModule) SetPageNo(v int32) *VatInvoiceScanQueryResponseBodyModule {
	s.PageNo = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModule) SetPageSize(v int32) *VatInvoiceScanQueryResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModule) SetTotalPage(v int32) *VatInvoiceScanQueryResponseBodyModule {
	s.TotalPage = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModule) SetTotalSize(v int32) *VatInvoiceScanQueryResponseBodyModule {
	s.TotalSize = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModule) Validate() error {
	return dara.Validate(s)
}

type VatInvoiceScanQueryResponseBodyModuleItems struct {
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
	// example:
	//
	// 07122942791187744475
	CheckCode *string `json:"check_code,omitempty" xml:"check_code,omitempty"`
	Drawer    *string `json:"drawer,omitempty" xml:"drawer,omitempty"`
	// example:
	//
	// 60
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 3300111303
	InvoiceCode *string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty"`
	// example:
	//
	// 2022-12-01
	InvoiceDay      *string                                                     `json:"invoice_day,omitempty" xml:"invoice_day,omitempty"`
	InvoiceDetail   *string                                                     `json:"invoice_detail,omitempty" xml:"invoice_detail,omitempty"`
	InvoiceDetails  []*VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails `json:"invoice_details,omitempty" xml:"invoice_details,omitempty" type:"Repeated"`
	InvoiceLocation *string                                                     `json:"invoice_location,omitempty" xml:"invoice_location,omitempty"`
	// example:
	//
	// 24021111
	InvoiceNo *string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty"`
	// example:
	//
	// 123
	InvoiceSubTaskId *int64 `json:"invoice_sub_task_id,omitempty" xml:"invoice_sub_task_id,omitempty"`
	// example:
	//
	// 2
	InvoiceType     *int32  `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	InvoiceTypeDesc *string `json:"invoice_type_desc,omitempty" xml:"invoice_type_desc,omitempty"`
	// 机器码
	//
	// example:
	//
	// 661619906841
	MachineCode *string `json:"machine_code,omitempty" xml:"machine_code,omitempty"`
	OfdOssUrl   *string `json:"ofd_oss_url,omitempty" xml:"ofd_oss_url,omitempty"`
	// example:
	//
	// https://www.testurl.com
	OssUrl *string `json:"oss_url,omitempty" xml:"oss_url,omitempty"`
	// example:
	//
	// <87*>>53>5023>-446>/4+83/5	- *>5/81<75/1931>4>>
	PasswordArea             *string `json:"password_area,omitempty" xml:"password_area,omitempty"`
	PdfOssUrl                *string `json:"pdf_oss_url,omitempty" xml:"pdf_oss_url,omitempty"`
	PurchaserBankAccountInfo *string `json:"purchaser_bank_account_info,omitempty" xml:"purchaser_bank_account_info,omitempty"`
	PurchaserContactInfo     *string `json:"purchaser_contact_info,omitempty" xml:"purchaser_contact_info,omitempty"`
	PurchaserName            *string `json:"purchaser_name,omitempty" xml:"purchaser_name,omitempty"`
	// example:
	//
	// 91441111111111111S
	PurchaserTaxNo        *string `json:"purchaser_tax_no,omitempty" xml:"purchaser_tax_no,omitempty"`
	Recipient             *string `json:"recipient,omitempty" xml:"recipient,omitempty"`
	Remarks               *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
	Reviewer              *string `json:"reviewer,omitempty" xml:"reviewer,omitempty"`
	SellerBankAccountInfo *string `json:"seller_bank_account_info,omitempty" xml:"seller_bank_account_info,omitempty"`
	SellerContactInfo     *string `json:"seller_contact_info,omitempty" xml:"seller_contact_info,omitempty"`
	SellerName            *string `json:"seller_name,omitempty" xml:"seller_name,omitempty"`
	// example:
	//
	// 91441111111111111N
	SellerTaxNo *string `json:"seller_tax_no,omitempty" xml:"seller_tax_no,omitempty"`
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
	TaxRate            *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	TotalAmountInWords *string `json:"total_amount_in_words,omitempty" xml:"total_amount_in_words,omitempty"`
	XmlOssUrl          *string `json:"xml_oss_url,omitempty" xml:"xml_oss_url,omitempty"`
}

func (s VatInvoiceScanQueryResponseBodyModuleItems) String() string {
	return dara.Prettify(s)
}

func (s VatInvoiceScanQueryResponseBodyModuleItems) GoString() string {
	return s.String()
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetAmountWithTax() *string {
	return s.AmountWithTax
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetAmountWithoutTax() *string {
	return s.AmountWithoutTax
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetBillDate() *string {
	return s.BillDate
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetCheckCode() *string {
	return s.CheckCode
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetDrawer() *string {
	return s.Drawer
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetId() *string {
	return s.Id
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetInvoiceCode() *string {
	return s.InvoiceCode
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetInvoiceDay() *string {
	return s.InvoiceDay
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetInvoiceDetail() *string {
	return s.InvoiceDetail
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetInvoiceDetails() []*VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	return s.InvoiceDetails
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetInvoiceLocation() *string {
	return s.InvoiceLocation
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetInvoiceNo() *string {
	return s.InvoiceNo
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetInvoiceSubTaskId() *int64 {
	return s.InvoiceSubTaskId
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetInvoiceTypeDesc() *string {
	return s.InvoiceTypeDesc
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetMachineCode() *string {
	return s.MachineCode
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetOfdOssUrl() *string {
	return s.OfdOssUrl
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetOssUrl() *string {
	return s.OssUrl
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetPasswordArea() *string {
	return s.PasswordArea
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetPdfOssUrl() *string {
	return s.PdfOssUrl
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetPurchaserBankAccountInfo() *string {
	return s.PurchaserBankAccountInfo
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetPurchaserContactInfo() *string {
	return s.PurchaserContactInfo
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetPurchaserName() *string {
	return s.PurchaserName
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetPurchaserTaxNo() *string {
	return s.PurchaserTaxNo
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetRecipient() *string {
	return s.Recipient
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetRemarks() *string {
	return s.Remarks
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetReviewer() *string {
	return s.Reviewer
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetSellerBankAccountInfo() *string {
	return s.SellerBankAccountInfo
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetSellerContactInfo() *string {
	return s.SellerContactInfo
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetSellerName() *string {
	return s.SellerName
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetSellerTaxNo() *string {
	return s.SellerTaxNo
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetSmartCheckCode() *string {
	return s.SmartCheckCode
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetTaxRate() *string {
	return s.TaxRate
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetTotalAmountInWords() *string {
	return s.TotalAmountInWords
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) GetXmlOssUrl() *string {
	return s.XmlOssUrl
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetAmountWithTax(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.AmountWithTax = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetAmountWithoutTax(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.AmountWithoutTax = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetBillDate(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.BillDate = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetCheckCode(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.CheckCode = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetDrawer(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.Drawer = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetId(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.Id = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetInvoiceCode(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceCode = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetInvoiceDay(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceDay = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetInvoiceDetail(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceDetail = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetInvoiceDetails(v []*VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceDetails = v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetInvoiceLocation(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceLocation = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetInvoiceNo(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceNo = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetInvoiceSubTaskId(v int64) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceSubTaskId = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetInvoiceType(v int32) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceType = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetInvoiceTypeDesc(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.InvoiceTypeDesc = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetMachineCode(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.MachineCode = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetOfdOssUrl(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.OfdOssUrl = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetOssUrl(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.OssUrl = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetPasswordArea(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.PasswordArea = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetPdfOssUrl(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.PdfOssUrl = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetPurchaserBankAccountInfo(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.PurchaserBankAccountInfo = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetPurchaserContactInfo(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.PurchaserContactInfo = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetPurchaserName(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.PurchaserName = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetPurchaserTaxNo(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.PurchaserTaxNo = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetRecipient(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.Recipient = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetRemarks(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.Remarks = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetReviewer(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.Reviewer = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetSellerBankAccountInfo(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.SellerBankAccountInfo = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetSellerContactInfo(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.SellerContactInfo = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetSellerName(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.SellerName = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetSellerTaxNo(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.SellerTaxNo = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetSmartCheckCode(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.SmartCheckCode = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetTaxAmount(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.TaxAmount = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetTaxRate(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.TaxRate = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetTotalAmountInWords(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.TotalAmountInWords = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) SetXmlOssUrl(v string) *VatInvoiceScanQueryResponseBodyModuleItems {
	s.XmlOssUrl = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItems) Validate() error {
	return dara.Validate(s)
}

type VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails struct {
	// example:
	//
	// 75.21
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 行号
	//
	// example:
	//
	// 0
	Index    *string `json:"index,omitempty" xml:"index,omitempty"`
	ItemName *string `json:"item_name,omitempty" xml:"item_name,omitempty"`
	// example:
	//
	// 1
	Quantity *string `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// example:
	//
	// 66PT11230069
	Specification *string `json:"specification,omitempty" xml:"specification,omitempty"`
	// example:
	//
	// 12.79
	Tax *string `json:"tax,omitempty" xml:"tax,omitempty"`
	// example:
	//
	// 17%
	TaxRate *string `json:"tax_rate,omitempty" xml:"tax_rate,omitempty"`
	Unit    *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// 75.21
	UnitPrice *string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
}

func (s VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) String() string {
	return dara.Prettify(s)
}

func (s VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GoString() string {
	return s.String()
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetAmount() *string {
	return s.Amount
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetIndex() *string {
	return s.Index
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetItemName() *string {
	return s.ItemName
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetQuantity() *string {
	return s.Quantity
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetSpecification() *string {
	return s.Specification
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetTax() *string {
	return s.Tax
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetTaxRate() *string {
	return s.TaxRate
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetUnit() *string {
	return s.Unit
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) GetUnitPrice() *string {
	return s.UnitPrice
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetAmount(v string) *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Amount = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetIndex(v string) *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Index = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetItemName(v string) *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.ItemName = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetQuantity(v string) *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Quantity = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetSpecification(v string) *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Specification = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetTax(v string) *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Tax = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetTaxRate(v string) *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.TaxRate = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetUnit(v string) *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.Unit = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) SetUnitPrice(v string) *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails {
	s.UnitPrice = &v
	return s
}

func (s *VatInvoiceScanQueryResponseBodyModuleItemsInvoiceDetails) Validate() error {
	return dara.Validate(s)
}
