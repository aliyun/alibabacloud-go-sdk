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
	// The returned data.
	Data *QueryInvoicingCustomerListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
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
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryInvoicingCustomerListResponseBodyData struct {
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
	if s.CustomerInvoiceList != nil {
		if err := s.CustomerInvoiceList.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.CustomerInvoice != nil {
		for _, item := range s.CustomerInvoice {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryInvoicingCustomerListResponseBodyDataCustomerInvoiceListCustomerInvoice struct {
	AdjustType              *int64  `json:"AdjustType,omitempty" xml:"AdjustType,omitempty"`
	Bank                    *string `json:"Bank,omitempty" xml:"Bank,omitempty"`
	BankNo                  *string `json:"BankNo,omitempty" xml:"BankNo,omitempty"`
	CustomerType            *int64  `json:"CustomerType,omitempty" xml:"CustomerType,omitempty"`
	DefaultRemark           *string `json:"DefaultRemark,omitempty" xml:"DefaultRemark,omitempty"`
	EndCycle                *int64  `json:"EndCycle,omitempty" xml:"EndCycle,omitempty"`
	GmtCreate               *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	Id                      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	InvoiceTitle            *string `json:"InvoiceTitle,omitempty" xml:"InvoiceTitle,omitempty"`
	IssueType               *int64  `json:"IssueType,omitempty" xml:"IssueType,omitempty"`
	OperatingLicenseAddress *string `json:"OperatingLicenseAddress,omitempty" xml:"OperatingLicenseAddress,omitempty"`
	OperatingLicensePhone   *string `json:"OperatingLicensePhone,omitempty" xml:"OperatingLicensePhone,omitempty"`
	RegisterNo              *string `json:"RegisterNo,omitempty" xml:"RegisterNo,omitempty"`
	StartCycle              *int64  `json:"StartCycle,omitempty" xml:"StartCycle,omitempty"`
	Status                  *int64  `json:"Status,omitempty" xml:"Status,omitempty"`
	TaxationLicense         *string `json:"TaxationLicense,omitempty" xml:"TaxationLicense,omitempty"`
	TaxpayerType            *int64  `json:"TaxpayerType,omitempty" xml:"TaxpayerType,omitempty"`
	TitleChangeInstructions *string `json:"TitleChangeInstructions,omitempty" xml:"TitleChangeInstructions,omitempty"`
	Type                    *int64  `json:"Type,omitempty" xml:"Type,omitempty"`
	UserId                  *int64  `json:"UserId,omitempty" xml:"UserId,omitempty"`
	UserNick                *string `json:"UserNick,omitempty" xml:"UserNick,omitempty"`
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
