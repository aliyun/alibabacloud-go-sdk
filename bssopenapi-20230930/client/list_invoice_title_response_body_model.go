// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInvoiceTitleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListInvoiceTitleResponseBodyData) *ListInvoiceTitleResponseBody
	GetData() []*ListInvoiceTitleResponseBodyData
	SetMetadata(v interface{}) *ListInvoiceTitleResponseBody
	GetMetadata() interface{}
	SetRequestId(v string) *ListInvoiceTitleResponseBody
	GetRequestId() *string
}

type ListInvoiceTitleResponseBody struct {
	Data []*ListInvoiceTitleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// {}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// example:
	//
	// 6000EE23-274B-4E07-A697-FF2E999520A4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInvoiceTitleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInvoiceTitleResponseBody) GoString() string {
	return s.String()
}

func (s *ListInvoiceTitleResponseBody) GetData() []*ListInvoiceTitleResponseBodyData {
	return s.Data
}

func (s *ListInvoiceTitleResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *ListInvoiceTitleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInvoiceTitleResponseBody) SetData(v []*ListInvoiceTitleResponseBodyData) *ListInvoiceTitleResponseBody {
	s.Data = v
	return s
}

func (s *ListInvoiceTitleResponseBody) SetMetadata(v interface{}) *ListInvoiceTitleResponseBody {
	s.Metadata = v
	return s
}

func (s *ListInvoiceTitleResponseBody) SetRequestId(v string) *ListInvoiceTitleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInvoiceTitleResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInvoiceTitleResponseBodyData struct {
	AccountBankName *string `json:"AccountBankName,omitempty" xml:"AccountBankName,omitempty"`
	// example:
	//
	// 1990699401005016
	AccountId *int64 `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	// example:
	//
	// 1234567890
	BankAccountNumber *string `json:"BankAccountNumber,omitempty" xml:"BankAccountNumber,omitempty"`
	// example:
	//
	// 2025-06-01 00:00:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 123456
	Id                *string `json:"Id,omitempty" xml:"Id,omitempty"`
	InvoiceTitle      *string `json:"InvoiceTitle,omitempty" xml:"InvoiceTitle,omitempty"`
	RegisteredAddress *string `json:"RegisteredAddress,omitempty" xml:"RegisteredAddress,omitempty"`
	// example:
	//
	// 010-12345678
	RegisteredLandline *string `json:"RegisteredLandline,omitempty" xml:"RegisteredLandline,omitempty"`
	// example:
	//
	// 111111111111111
	UnifiedSocialCreditCode *string `json:"UnifiedSocialCreditCode,omitempty" xml:"UnifiedSocialCreditCode,omitempty"`
}

func (s ListInvoiceTitleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListInvoiceTitleResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListInvoiceTitleResponseBodyData) GetAccountBankName() *string {
	return s.AccountBankName
}

func (s *ListInvoiceTitleResponseBodyData) GetAccountId() *int64 {
	return s.AccountId
}

func (s *ListInvoiceTitleResponseBodyData) GetBankAccountNumber() *string {
	return s.BankAccountNumber
}

func (s *ListInvoiceTitleResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListInvoiceTitleResponseBodyData) GetId() *string {
	return s.Id
}

func (s *ListInvoiceTitleResponseBodyData) GetInvoiceTitle() *string {
	return s.InvoiceTitle
}

func (s *ListInvoiceTitleResponseBodyData) GetRegisteredAddress() *string {
	return s.RegisteredAddress
}

func (s *ListInvoiceTitleResponseBodyData) GetRegisteredLandline() *string {
	return s.RegisteredLandline
}

func (s *ListInvoiceTitleResponseBodyData) GetUnifiedSocialCreditCode() *string {
	return s.UnifiedSocialCreditCode
}

func (s *ListInvoiceTitleResponseBodyData) SetAccountBankName(v string) *ListInvoiceTitleResponseBodyData {
	s.AccountBankName = &v
	return s
}

func (s *ListInvoiceTitleResponseBodyData) SetAccountId(v int64) *ListInvoiceTitleResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *ListInvoiceTitleResponseBodyData) SetBankAccountNumber(v string) *ListInvoiceTitleResponseBodyData {
	s.BankAccountNumber = &v
	return s
}

func (s *ListInvoiceTitleResponseBodyData) SetCreateTime(v string) *ListInvoiceTitleResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListInvoiceTitleResponseBodyData) SetId(v string) *ListInvoiceTitleResponseBodyData {
	s.Id = &v
	return s
}

func (s *ListInvoiceTitleResponseBodyData) SetInvoiceTitle(v string) *ListInvoiceTitleResponseBodyData {
	s.InvoiceTitle = &v
	return s
}

func (s *ListInvoiceTitleResponseBodyData) SetRegisteredAddress(v string) *ListInvoiceTitleResponseBodyData {
	s.RegisteredAddress = &v
	return s
}

func (s *ListInvoiceTitleResponseBodyData) SetRegisteredLandline(v string) *ListInvoiceTitleResponseBodyData {
	s.RegisteredLandline = &v
	return s
}

func (s *ListInvoiceTitleResponseBodyData) SetUnifiedSocialCreditCode(v string) *ListInvoiceTitleResponseBodyData {
	s.UnifiedSocialCreditCode = &v
	return s
}

func (s *ListInvoiceTitleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
