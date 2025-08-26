// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVATInvoiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeVATInvoiceResponseBodyData) *RecognizeVATInvoiceResponseBody
	GetData() *RecognizeVATInvoiceResponseBodyData
	SetRequestId(v string) *RecognizeVATInvoiceResponseBody
	GetRequestId() *string
}

type RecognizeVATInvoiceResponseBody struct {
	Data *RecognizeVATInvoiceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 56A10D65-ECE0-59DE-9775-F6494D2AF13B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVATInvoiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVATInvoiceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponseBody) GetData() *RecognizeVATInvoiceResponseBodyData {
	return s.Data
}

func (s *RecognizeVATInvoiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeVATInvoiceResponseBody) SetData(v *RecognizeVATInvoiceResponseBodyData) *RecognizeVATInvoiceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVATInvoiceResponseBody) SetRequestId(v string) *RecognizeVATInvoiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeVATInvoiceResponseBodyData struct {
	Box     *RecognizeVATInvoiceResponseBodyDataBox     `json:"Box,omitempty" xml:"Box,omitempty" type:"Struct"`
	Content *RecognizeVATInvoiceResponseBodyDataContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
}

func (s RecognizeVATInvoiceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVATInvoiceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponseBodyData) GetBox() *RecognizeVATInvoiceResponseBodyDataBox {
	return s.Box
}

func (s *RecognizeVATInvoiceResponseBodyData) GetContent() *RecognizeVATInvoiceResponseBodyDataContent {
	return s.Content
}

func (s *RecognizeVATInvoiceResponseBodyData) SetBox(v *RecognizeVATInvoiceResponseBodyDataBox) *RecognizeVATInvoiceResponseBodyData {
	s.Box = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyData) SetContent(v *RecognizeVATInvoiceResponseBodyDataContent) *RecognizeVATInvoiceResponseBodyData {
	s.Content = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeVATInvoiceResponseBodyDataBox struct {
	// 1
	Checkers []*float32 `json:"Checkers,omitempty" xml:"Checkers,omitempty" type:"Repeated"`
	// 1
	Clerks []*float32 `json:"Clerks,omitempty" xml:"Clerks,omitempty" type:"Repeated"`
	// 1
	InvoiceAmounts []*float32 `json:"InvoiceAmounts,omitempty" xml:"InvoiceAmounts,omitempty" type:"Repeated"`
	// 1
	InvoiceCodes []*float32 `json:"InvoiceCodes,omitempty" xml:"InvoiceCodes,omitempty" type:"Repeated"`
	// 1
	InvoiceDates []*float32 `json:"InvoiceDates,omitempty" xml:"InvoiceDates,omitempty" type:"Repeated"`
	// 1
	InvoiceFakeCodes []*float32 `json:"InvoiceFakeCodes,omitempty" xml:"InvoiceFakeCodes,omitempty" type:"Repeated"`
	// 1
	InvoiceNoes []*float32 `json:"InvoiceNoes,omitempty" xml:"InvoiceNoes,omitempty" type:"Repeated"`
	// 1
	ItemNames []*int32 `json:"ItemNames,omitempty" xml:"ItemNames,omitempty" type:"Repeated"`
	// 1
	PayeeAddresses []*float32 `json:"PayeeAddresses,omitempty" xml:"PayeeAddresses,omitempty" type:"Repeated"`
	// 1
	PayeeBankNames []*float32 `json:"PayeeBankNames,omitempty" xml:"PayeeBankNames,omitempty" type:"Repeated"`
	// 1
	PayeeNames []*float32 `json:"PayeeNames,omitempty" xml:"PayeeNames,omitempty" type:"Repeated"`
	// 1
	PayeeRegisterNoes []*float32 `json:"PayeeRegisterNoes,omitempty" xml:"PayeeRegisterNoes,omitempty" type:"Repeated"`
	// 1
	Payees []*float32 `json:"Payees,omitempty" xml:"Payees,omitempty" type:"Repeated"`
	// 1
	PayerAddresses []*float32 `json:"PayerAddresses,omitempty" xml:"PayerAddresses,omitempty" type:"Repeated"`
	// 1
	PayerBankNames []*float32 `json:"PayerBankNames,omitempty" xml:"PayerBankNames,omitempty" type:"Repeated"`
	// 1
	PayerNames []*float32 `json:"PayerNames,omitempty" xml:"PayerNames,omitempty" type:"Repeated"`
	// 1
	PayerRegisterNoes []*float32 `json:"PayerRegisterNoes,omitempty" xml:"PayerRegisterNoes,omitempty" type:"Repeated"`
	// 1
	SumAmounts []*float32 `json:"SumAmounts,omitempty" xml:"SumAmounts,omitempty" type:"Repeated"`
	// 1
	TaxAmounts []*float32 `json:"TaxAmounts,omitempty" xml:"TaxAmounts,omitempty" type:"Repeated"`
	// 1
	WithoutTaxAmounts []*float32 `json:"WithoutTaxAmounts,omitempty" xml:"WithoutTaxAmounts,omitempty" type:"Repeated"`
}

func (s RecognizeVATInvoiceResponseBodyDataBox) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVATInvoiceResponseBodyDataBox) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetCheckers() []*float32 {
	return s.Checkers
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetClerks() []*float32 {
	return s.Clerks
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetInvoiceAmounts() []*float32 {
	return s.InvoiceAmounts
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetInvoiceCodes() []*float32 {
	return s.InvoiceCodes
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetInvoiceDates() []*float32 {
	return s.InvoiceDates
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetInvoiceFakeCodes() []*float32 {
	return s.InvoiceFakeCodes
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetInvoiceNoes() []*float32 {
	return s.InvoiceNoes
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetItemNames() []*int32 {
	return s.ItemNames
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetPayeeAddresses() []*float32 {
	return s.PayeeAddresses
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetPayeeBankNames() []*float32 {
	return s.PayeeBankNames
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetPayeeNames() []*float32 {
	return s.PayeeNames
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetPayeeRegisterNoes() []*float32 {
	return s.PayeeRegisterNoes
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetPayees() []*float32 {
	return s.Payees
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetPayerAddresses() []*float32 {
	return s.PayerAddresses
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetPayerBankNames() []*float32 {
	return s.PayerBankNames
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetPayerNames() []*float32 {
	return s.PayerNames
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetPayerRegisterNoes() []*float32 {
	return s.PayerRegisterNoes
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetSumAmounts() []*float32 {
	return s.SumAmounts
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetTaxAmounts() []*float32 {
	return s.TaxAmounts
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) GetWithoutTaxAmounts() []*float32 {
	return s.WithoutTaxAmounts
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetCheckers(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.Checkers = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetClerks(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.Clerks = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceAmounts = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceCodes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceCodes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceDates(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceDates = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceFakeCodes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceFakeCodes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetInvoiceNoes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.InvoiceNoes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetItemNames(v []*int32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.ItemNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayeeAddresses(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayeeAddresses = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayeeBankNames(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayeeBankNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayeeNames(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayeeNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayeeRegisterNoes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayeeRegisterNoes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayees(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.Payees = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerAddresses(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerAddresses = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerBankNames(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerBankNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerNames(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerNames = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetPayerRegisterNoes(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.PayerRegisterNoes = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetSumAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.SumAmounts = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetTaxAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.TaxAmounts = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) SetWithoutTaxAmounts(v []*float32) *RecognizeVATInvoiceResponseBodyDataBox {
	s.WithoutTaxAmounts = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataBox) Validate() error {
	return dara.Validate(s)
}

type RecognizeVATInvoiceResponseBodyDataContent struct {
	// example:
	//
	// 02702870934284730434
	AntiFakeCode *string `json:"AntiFakeCode,omitempty" xml:"AntiFakeCode,omitempty"`
	Checker      *string `json:"Checker,omitempty" xml:"Checker,omitempty"`
	Clerk        *string `json:"Clerk,omitempty" xml:"Clerk,omitempty"`
	// example:
	//
	// 200.00
	InvoiceAmount *string `json:"InvoiceAmount,omitempty" xml:"InvoiceAmount,omitempty"`
	// example:
	//
	// 031001600311
	InvoiceCode *string `json:"InvoiceCode,omitempty" xml:"InvoiceCode,omitempty"`
	// example:
	//
	// 20190415
	InvoiceDate *string `json:"InvoiceDate,omitempty" xml:"InvoiceDate,omitempty"`
	// example:
	//
	// 03753869
	InvoiceNo *string `json:"InvoiceNo,omitempty" xml:"InvoiceNo,omitempty"`
	// 1
	ItemName      []*string `json:"ItemName,omitempty" xml:"ItemName,omitempty" type:"Repeated"`
	Payee         *string   `json:"Payee,omitempty" xml:"Payee,omitempty"`
	PayeeAddress  *string   `json:"PayeeAddress,omitempty" xml:"PayeeAddress,omitempty"`
	PayeeBankName *string   `json:"PayeeBankName,omitempty" xml:"PayeeBankName,omitempty"`
	PayeeName     *string   `json:"PayeeName,omitempty" xml:"PayeeName,omitempty"`
	// example:
	//
	// 91420200000123403
	PayeeRegisterNo *string `json:"PayeeRegisterNo,omitempty" xml:"PayeeRegisterNo,omitempty"`
	PayerAddress    *string `json:"PayerAddress,omitempty" xml:"PayerAddress,omitempty"`
	// example:
	//
	// 6221************1234
	PayerBankName *string `json:"PayerBankName,omitempty" xml:"PayerBankName,omitempty"`
	PayerName     *string `json:"PayerName,omitempty" xml:"PayerName,omitempty"`
	// example:
	//
	// 91420200000123403
	PayerRegisterNo *string `json:"PayerRegisterNo,omitempty" xml:"PayerRegisterNo,omitempty"`
	// example:
	//
	// 87
	SumAmount *string `json:"SumAmount,omitempty" xml:"SumAmount,omitempty"`
	// example:
	//
	// 9.52
	TaxAmount *string `json:"TaxAmount,omitempty" xml:"TaxAmount,omitempty"`
	// example:
	//
	// 190.48
	WithoutTaxAmount *string `json:"WithoutTaxAmount,omitempty" xml:"WithoutTaxAmount,omitempty"`
}

func (s RecognizeVATInvoiceResponseBodyDataContent) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVATInvoiceResponseBodyDataContent) GoString() string {
	return s.String()
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetAntiFakeCode() *string {
	return s.AntiFakeCode
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetChecker() *string {
	return s.Checker
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetClerk() *string {
	return s.Clerk
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetInvoiceAmount() *string {
	return s.InvoiceAmount
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetInvoiceCode() *string {
	return s.InvoiceCode
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetInvoiceDate() *string {
	return s.InvoiceDate
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetInvoiceNo() *string {
	return s.InvoiceNo
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetItemName() []*string {
	return s.ItemName
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetPayee() *string {
	return s.Payee
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetPayeeAddress() *string {
	return s.PayeeAddress
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetPayeeBankName() *string {
	return s.PayeeBankName
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetPayeeName() *string {
	return s.PayeeName
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetPayeeRegisterNo() *string {
	return s.PayeeRegisterNo
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetPayerAddress() *string {
	return s.PayerAddress
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetPayerBankName() *string {
	return s.PayerBankName
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetPayerName() *string {
	return s.PayerName
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetPayerRegisterNo() *string {
	return s.PayerRegisterNo
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetSumAmount() *string {
	return s.SumAmount
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetTaxAmount() *string {
	return s.TaxAmount
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) GetWithoutTaxAmount() *string {
	return s.WithoutTaxAmount
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetAntiFakeCode(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.AntiFakeCode = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetChecker(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.Checker = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetClerk(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.Clerk = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceAmount = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceCode(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceCode = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceDate(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceDate = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetInvoiceNo(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.InvoiceNo = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetItemName(v []*string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.ItemName = v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayee(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.Payee = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeAddress(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeAddress = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeBankName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeBankName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayeeRegisterNo(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayeeRegisterNo = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerAddress(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerAddress = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerBankName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerBankName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerName(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerName = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetPayerRegisterNo(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.PayerRegisterNo = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetSumAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.SumAmount = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetTaxAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.TaxAmount = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) SetWithoutTaxAmount(v string) *RecognizeVATInvoiceResponseBodyDataContent {
	s.WithoutTaxAmount = &v
	return s
}

func (s *RecognizeVATInvoiceResponseBodyDataContent) Validate() error {
	return dara.Validate(s)
}
