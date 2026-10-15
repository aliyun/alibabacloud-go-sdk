// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyInvoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressId(v int64) *ApplyInvoiceRequest
	GetAddressId() *int64
	SetApplyUserNick(v string) *ApplyInvoiceRequest
	GetApplyUserNick() *string
	SetCustomerId(v int64) *ApplyInvoiceRequest
	GetCustomerId() *int64
	SetInvoiceAmount(v int64) *ApplyInvoiceRequest
	GetInvoiceAmount() *int64
	SetInvoiceByAmount(v bool) *ApplyInvoiceRequest
	GetInvoiceByAmount() *bool
	SetInvoicingType(v int32) *ApplyInvoiceRequest
	GetInvoicingType() *int32
	SetOwnerId(v int64) *ApplyInvoiceRequest
	GetOwnerId() *int64
	SetProcessWay(v int32) *ApplyInvoiceRequest
	GetProcessWay() *int32
	SetSelectedIds(v []*int64) *ApplyInvoiceRequest
	GetSelectedIds() []*int64
	SetUserRemark(v string) *ApplyInvoiceRequest
	GetUserRemark() *string
	SetEmails(v string) *ApplyInvoiceRequest
	GetEmails() *string
}

type ApplyInvoiceRequest struct {
	// The ID of the mailing address, which corresponds to the ID parameter returned by the QueryCustomerAddressList operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 237958367
	AddressId *int64 `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The nickname of the applicant. The system does not verify this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ApplyUserNick *string `json:"ApplyUserNick,omitempty" xml:"ApplyUserNick,omitempty"`
	// The ID of the invoice title information, which corresponds to the ID parameter returned by the QueryInvoicingCustomerList operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 124132423
	CustomerId *int64 `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// The requested invoice amount, in cents.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	InvoiceAmount *int64 `json:"InvoiceAmount,omitempty" xml:"InvoiceAmount,omitempty"`
	// Specifies whether to issue the invoice based on a specified amount. If this parameter is set to true, the invoice is applied for based on the specified InvoiceAmount. Otherwise, the invoice is applied for based on the total amount of the selected objects to be invoiced.
	//
	// example:
	//
	// true
	InvoiceByAmount *bool `json:"InvoiceByAmount,omitempty" xml:"InvoiceByAmount,omitempty"`
	// The invoice type. Valid values:
	//
	// - 0: VAT special invoice.
	//
	// - 1: VAT general invoice.
	//
	// example:
	//
	// 1
	InvoicingType *int32 `json:"InvoicingType,omitempty" xml:"InvoicingType,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The invoicing channel. Set this parameter to 1. Valid values:
	//
	// - 0: Processed by Alibaba Cloud.
	//
	// - 1: Processed by the tax platform.
	//
	// example:
	//
	// 1
	ProcessWay *int32 `json:"ProcessWay,omitempty" xml:"ProcessWay,omitempty"`
	// The list of IDs of the selected objects to be invoiced, which correspond to the ID parameters returned by the QueryEvaluateList operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 384752367
	SelectedIds []*int64 `json:"SelectedIds,omitempty" xml:"SelectedIds,omitempty" type:"Repeated"`
	// The user remarks.
	//
	// example:
	//
	// test
	UserRemark *string `json:"UserRemark,omitempty" xml:"UserRemark,omitempty"`
	// The email addresses for receiving the invoice. The maximum length of this field is 200 characters. Separate multiple email addresses with commas (,).
	//
	// example:
	//
	// aaa@aa.com,bbb@bb.com
	Emails *string `json:"emails,omitempty" xml:"emails,omitempty"`
}

func (s ApplyInvoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyInvoiceRequest) GoString() string {
	return s.String()
}

func (s *ApplyInvoiceRequest) GetAddressId() *int64 {
	return s.AddressId
}

func (s *ApplyInvoiceRequest) GetApplyUserNick() *string {
	return s.ApplyUserNick
}

func (s *ApplyInvoiceRequest) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *ApplyInvoiceRequest) GetInvoiceAmount() *int64 {
	return s.InvoiceAmount
}

func (s *ApplyInvoiceRequest) GetInvoiceByAmount() *bool {
	return s.InvoiceByAmount
}

func (s *ApplyInvoiceRequest) GetInvoicingType() *int32 {
	return s.InvoicingType
}

func (s *ApplyInvoiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ApplyInvoiceRequest) GetProcessWay() *int32 {
	return s.ProcessWay
}

func (s *ApplyInvoiceRequest) GetSelectedIds() []*int64 {
	return s.SelectedIds
}

func (s *ApplyInvoiceRequest) GetUserRemark() *string {
	return s.UserRemark
}

func (s *ApplyInvoiceRequest) GetEmails() *string {
	return s.Emails
}

func (s *ApplyInvoiceRequest) SetAddressId(v int64) *ApplyInvoiceRequest {
	s.AddressId = &v
	return s
}

func (s *ApplyInvoiceRequest) SetApplyUserNick(v string) *ApplyInvoiceRequest {
	s.ApplyUserNick = &v
	return s
}

func (s *ApplyInvoiceRequest) SetCustomerId(v int64) *ApplyInvoiceRequest {
	s.CustomerId = &v
	return s
}

func (s *ApplyInvoiceRequest) SetInvoiceAmount(v int64) *ApplyInvoiceRequest {
	s.InvoiceAmount = &v
	return s
}

func (s *ApplyInvoiceRequest) SetInvoiceByAmount(v bool) *ApplyInvoiceRequest {
	s.InvoiceByAmount = &v
	return s
}

func (s *ApplyInvoiceRequest) SetInvoicingType(v int32) *ApplyInvoiceRequest {
	s.InvoicingType = &v
	return s
}

func (s *ApplyInvoiceRequest) SetOwnerId(v int64) *ApplyInvoiceRequest {
	s.OwnerId = &v
	return s
}

func (s *ApplyInvoiceRequest) SetProcessWay(v int32) *ApplyInvoiceRequest {
	s.ProcessWay = &v
	return s
}

func (s *ApplyInvoiceRequest) SetSelectedIds(v []*int64) *ApplyInvoiceRequest {
	s.SelectedIds = v
	return s
}

func (s *ApplyInvoiceRequest) SetUserRemark(v string) *ApplyInvoiceRequest {
	s.UserRemark = &v
	return s
}

func (s *ApplyInvoiceRequest) SetEmails(v string) *ApplyInvoiceRequest {
	s.Emails = &v
	return s
}

func (s *ApplyInvoiceRequest) Validate() error {
	return dara.Validate(s)
}
