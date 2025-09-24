// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInvoicingCustomerListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryInvoicingCustomerListResponseBody
	GetCode() *string
	SetData(v *QueryInvoicingCustomerListResponseBodyData) *QueryInvoicingCustomerListResponseBody
	GetData() *QueryInvoicingCustomerListResponseBodyData
	SetMessage(v string) *QueryInvoicingCustomerListResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryInvoicingCustomerListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryInvoicingCustomerListResponseBody
	GetSuccess() *bool
}

type QueryInvoicingCustomerListResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *QueryInvoicingCustomerListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BBEF51A3-E933-4F40-A534-C673CBDB9C80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryInvoicingCustomerListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryInvoicingCustomerListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryInvoicingCustomerListResponseBody) GetData() *QueryInvoicingCustomerListResponseBodyData {
	return s.Data
}

func (s *QueryInvoicingCustomerListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryInvoicingCustomerListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryInvoicingCustomerListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryInvoicingCustomerListResponseBody) SetCode(v string) *QueryInvoicingCustomerListResponseBody {
	s.Code = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBody) SetData(v *QueryInvoicingCustomerListResponseBodyData) *QueryInvoicingCustomerListResponseBody {
	s.Data = v
	return s
}

func (s *QueryInvoicingCustomerListResponseBody) SetMessage(v string) *QueryInvoicingCustomerListResponseBody {
	s.Message = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBody) SetRequestId(v string) *QueryInvoicingCustomerListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBody) SetSuccess(v bool) *QueryInvoicingCustomerListResponseBody {
	s.Success = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryInvoicingCustomerListResponseBodyData struct {
	// The information about the invoice.
	CustomerInvoiceList *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList `json:"CustomerInvoiceList,omitempty" xml:"CustomerInvoiceList,omitempty" type:"Struct"`
}

func (s QueryInvoicingCustomerListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryInvoicingCustomerListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListResponseBodyData) GetCustomerInvoiceList() *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList {
	return s.CustomerInvoiceList
}

func (s *QueryInvoicingCustomerListResponseBodyData) SetCustomerInvoiceList(v *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList) *QueryInvoicingCustomerListResponseBodyData {
	s.CustomerInvoiceList = v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList struct {
	CustomerInvoice []*QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice `json:"CustomerInvoice,omitempty" xml:"CustomerInvoice,omitempty" type:"Repeated"`
}

func (s QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList) String() string {
	return dara.Prettify(s)
}

func (s QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList) GetCustomerInvoice() []*QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	return s.CustomerInvoice
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList) SetCustomerInvoice(v []*QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList {
	s.CustomerInvoice = v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceList) Validate() error {
	return dara.Validate(s)
}

type QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice struct {
	// The type of invoice that was changed to.
	//
	// example:
	//
	// 1
	AdjustType *int64 `json:"AdjustType,omitempty" xml:"AdjustType,omitempty"`
	// The bank that issues the invoice.
	//
	// example:
	//
	// Test Bank
	Bank *string `json:"Bank,omitempty" xml:"Bank,omitempty"`
	// The bank account number.
	//
	// example:
	//
	// 389576348573296349853476
	BankNo *string `json:"BankNo,omitempty" xml:"BankNo,omitempty"`
	// The authentication type of Alipay. Valid values:
	//
	// 	- 1: individual
	//
	// 	- 2: company
	//
	// example:
	//
	// 1
	CustomerType *int64 `json:"CustomerType,omitempty" xml:"CustomerType,omitempty"`
	// The default note that is attached when the title is specified.
	//
	// example:
	//
	// PO Code: 12984554
	DefaultRemark *string `json:"DefaultRemark,omitempty" xml:"DefaultRemark,omitempty"`
	// The time when the payment ended.
	//
	// example:
	//
	// 202003
	EndCycle *int64 `json:"EndCycle,omitempty" xml:"EndCycle,omitempty"`
	// The time when the invoice was created. The time was in the yyyy-mm-dd hh:mm:ss format.
	//
	// example:
	//
	// 2018-09-07 15:26:20
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The ID of the invoice.
	//
	// example:
	//
	// 239875502738
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The company name in the invoice title.
	//
	// example:
	//
	// Test Company
	InvoiceTitle *string `json:"InvoiceTitle,omitempty" xml:"InvoiceTitle,omitempty"`
	// The type of issue.
	//
	// example:
	//
	// 1
	IssueType *int64 `json:"IssueType,omitempty" xml:"IssueType,omitempty"`
	// The address of the business license.
	//
	// example:
	//
	// XXX, XXX district, XXX city, XXX province
	OperatingLicenseAddress *string `json:"OperatingLicenseAddress,omitempty" xml:"OperatingLicenseAddress,omitempty"`
	// The phone number of the business license.
	//
	// example:
	//
	// 138xxxxxxxx
	OperatingLicensePhone *string `json:"OperatingLicensePhone,omitempty" xml:"OperatingLicensePhone,omitempty"`
	// The tax registration number.
	//
	// example:
	//
	// 21343245342534
	RegisterNo *string `json:"RegisterNo,omitempty" xml:"RegisterNo,omitempty"`
	// The time when the payment started.
	//
	// example:
	//
	// 202002
	StartCycle *int64 `json:"StartCycle,omitempty" xml:"StartCycle,omitempty"`
	// The status of the invoice title.
	//
	// example:
	//
	// 2
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The path and file name of the scanned copy of the tax registration certificate.
	//
	// example:
	//
	// taxationLicense.jpg
	TaxationLicense *string `json:"TaxationLicense,omitempty" xml:"TaxationLicense,omitempty"`
	// The type of the taxpayer. Valid values:
	//
	// 	- 1: general taxpayer
	//
	// 	- 2: special taxpayer
	//
	// example:
	//
	// 1
	TaxpayerType *int64 `json:"TaxpayerType,omitempty" xml:"TaxpayerType,omitempty"`
	// The instruction document of the invoice title change.
	//
	// example:
	//
	// instruction.doc
	TitleChangeInstructions *string `json:"TitleChangeInstructions,omitempty" xml:"TitleChangeInstructions,omitempty"`
	// The type of the invoice. Valid values:
	//
	// 	- 0: plain value-added tax (VAT) invoice
	//
	// 	- 1: special VAT invoice
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 34565465675
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// testNick
	UserNick *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
}

func (s QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) String() string {
	return dara.Prettify(s)
}

func (s QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GoString() string {
	return s.String()
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetAdjustType() *int64 {
	return s.AdjustType
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetBank() *string {
	return s.Bank
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetBankNo() *string {
	return s.BankNo
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetCustomerType() *int64 {
	return s.CustomerType
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetDefaultRemark() *string {
	return s.DefaultRemark
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetEndCycle() *int64 {
	return s.EndCycle
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetId() *int64 {
	return s.Id
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetIssueType() *int64 {
	return s.IssueType
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetOperatingLicenseAddress() *string {
	return s.OperatingLicenseAddress
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetOperatingLicensePhone() *string {
	return s.OperatingLicensePhone
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetRegisterNo() *string {
	return s.RegisterNo
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetStartCycle() *int64 {
	return s.StartCycle
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetStatus() *int64 {
	return s.Status
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetTaxationLicense() *string {
	return s.TaxationLicense
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetTaxpayerType() *int64 {
	return s.TaxpayerType
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetTitleChangeInstructions() *string {
	return s.TitleChangeInstructions
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetType() *int64 {
	return s.Type
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetUserId() *int64 {
	return s.UserId
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) GetUserNick() *string {
	return s.UserNick
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetAdjustType(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.AdjustType = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetBank(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.Bank = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetBankNo(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.BankNo = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetCustomerType(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.CustomerType = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetDefaultRemark(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.DefaultRemark = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetEndCycle(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.EndCycle = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetGmtCreate(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.GmtCreate = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetId(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.Id = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetInvoiceTitle(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.InvoiceTitle = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetIssueType(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.IssueType = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetOperatingLicenseAddress(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.OperatingLicenseAddress = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetOperatingLicensePhone(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.OperatingLicensePhone = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetRegisterNo(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.RegisterNo = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetStartCycle(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.StartCycle = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetStatus(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.Status = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetTaxationLicense(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.TaxationLicense = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetTaxpayerType(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.TaxpayerType = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetTitleChangeInstructions(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.TitleChangeInstructions = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetType(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.Type = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetUserId(v int64) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.UserId = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) SetUserNick(v string) *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice {
	s.UserNick = &v
	return s
}

func (s *QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice) Validate() error {
	return dara.Validate(s)
}
