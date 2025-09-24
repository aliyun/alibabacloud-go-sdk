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
	// The ID of the address to which the invoice is delivered. This parameter is required if the invoice is a paper invoice. Set the ID to the value of the AddressId parameter returned by calling the QueryCustomerAddressList operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 237958367
	AddressId *int64 `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The nickname of the applicant. The system does not verify the nickname.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	ApplyUserNick *string `json:"ApplyUserNick,omitempty" xml:"ApplyUserNick,omitempty"`
	// The ID of the customer. Set the ID to the value of the CustomerId parameter returned by calling the QueryInvoicingCustomerList operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 124132423
	CustomerId *int64 `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// The amount of the invoice. Unit: Cent.
	//
	// This parameter is required.
	//
	// example:
	//
	// 124132
	InvoiceAmount *int64 `json:"InvoiceAmount,omitempty" xml:"InvoiceAmount,omitempty"`
	// Specifies whether to invoice by amount. A value of true indicates that the user applies for the invoice based on the InvoiceAmount parameter. A value of false indicates that the user applies for the invoice based on the total amount of the invoicing items.
	//
	// example:
	//
	// true
	InvoiceByAmount *bool `json:"InvoiceByAmount,omitempty" xml:"InvoiceByAmount,omitempty"`
	// The type of the invoice. Valid values:
	//
	// 	- 0: paper invoice
	//
	// 	- 1: electronic invoice
	//
	// example:
	//
	// 1
	InvoicingType *int32 `json:"InvoicingType,omitempty" xml:"InvoicingType,omitempty"`
	OwnerId       *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The channel that is used to process the invoice. A value of 0 indicates that the invoice is processed by Alibaba Cloud. A value of 1 indicates that the invoice is processed by the tax platform. Set the value to 1.
	//
	// example:
	//
	// 1
	ProcessWay *int32 `json:"ProcessWay,omitempty" xml:"ProcessWay,omitempty"`
	// The IDs of the selected invoicing items. Set the IDs to the IDs returned by calling the QueryEvaluateList operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 384752367
	SelectedIds []*int64 `json:"SelectedIds,omitempty" xml:"SelectedIds,omitempty" type:"Repeated"`
	// The remarks made by the user.
	//
	// example:
	//
	// test
	UserRemark *string `json:"UserRemark,omitempty" xml:"UserRemark,omitempty"`
	Emails     *string `json:"emails,omitempty" xml:"emails,omitempty"`
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
