// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInvoiceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v string) *CreateInvoiceShrinkRequest
	GetAmount() *string
	SetEcIdAccountIdsShrink(v string) *CreateInvoiceShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetInvoiceCandidateIdsShrink(v string) *CreateInvoiceShrinkRequest
	GetInvoiceCandidateIdsShrink() *string
	SetInvoiceMode(v int32) *CreateInvoiceShrinkRequest
	GetInvoiceMode() *int32
	SetInvoiceRemark(v string) *CreateInvoiceShrinkRequest
	GetInvoiceRemark() *string
	SetInvoiceTitleId(v string) *CreateInvoiceShrinkRequest
	GetInvoiceTitleId() *string
	SetInvoiceType(v int32) *CreateInvoiceShrinkRequest
	GetInvoiceType() *int32
	SetNbid(v string) *CreateInvoiceShrinkRequest
	GetNbid() *string
	SetRecipientEmailsShrink(v string) *CreateInvoiceShrinkRequest
	GetRecipientEmailsShrink() *string
}

type CreateInvoiceShrinkRequest struct {
	// Specifies the invoice amount. Two decimal places are supported.
	//
	// - If left empty, the invoice is issued for the total invoiceable amount of all invoice candidate IDs.
	//
	// - If specified, the invoice is issued for the specified amount. The specified amount cannot exceed the total invoiceable amount of all invoice candidate IDs.
	//
	// example:
	//
	// 0.01
	Amount *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// The list of enterprises and accounts. If this parameter is left empty, the current account is queried.
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// The list of invoice candidate IDs.
	//
	// This parameter is required.
	InvoiceCandidateIdsShrink *string `json:"InvoiceCandidateIds,omitempty" xml:"InvoiceCandidateIds,omitempty"`
	// The invoicing mode. Valid values:
	//
	// - 0: Independent invoicing. The consumption of multiple accounts under an enterprise is invoiced independently for each account.
	//
	// - 1: Consolidated invoicing. The consumption of multiple accounts under an enterprise is consolidated and invoiced based on the invoicing entity.
	//
	// If only one account is specified in the AccountIds request parameter, independent invoicing is applied.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	InvoiceMode *int32 `json:"InvoiceMode,omitempty" xml:"InvoiceMode,omitempty"`
	// The invoice remark.
	//
	// example:
	//
	// Remark information
	InvoiceRemark *string `json:"InvoiceRemark,omitempty" xml:"InvoiceRemark,omitempty"`
	// The invoice title ID.
	//
	// - The ID returned by the ListInvoiceTitle operation for the current logon account.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	InvoiceTitleId *string `json:"InvoiceTitleId,omitempty" xml:"InvoiceTitleId,omitempty"`
	// The invoice type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0
	InvoiceType *int32 `json:"InvoiceType,omitempty" xml:"InvoiceType,omitempty"`
	// The level-1 marketplace ID. If this parameter is left empty, the marketplace ID of the current user is used by default.
	//
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// The list of email addresses that receive the invoice. You can specify up to three email addresses.
	//
	// This parameter is required.
	RecipientEmailsShrink *string `json:"RecipientEmails,omitempty" xml:"RecipientEmails,omitempty"`
}

func (s CreateInvoiceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInvoiceShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateInvoiceShrinkRequest) GetAmount() *string {
	return s.Amount
}

func (s *CreateInvoiceShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *CreateInvoiceShrinkRequest) GetInvoiceCandidateIdsShrink() *string {
	return s.InvoiceCandidateIdsShrink
}

func (s *CreateInvoiceShrinkRequest) GetInvoiceMode() *int32 {
	return s.InvoiceMode
}

func (s *CreateInvoiceShrinkRequest) GetInvoiceRemark() *string {
	return s.InvoiceRemark
}

func (s *CreateInvoiceShrinkRequest) GetInvoiceTitleId() *string {
	return s.InvoiceTitleId
}

func (s *CreateInvoiceShrinkRequest) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *CreateInvoiceShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateInvoiceShrinkRequest) GetRecipientEmailsShrink() *string {
	return s.RecipientEmailsShrink
}

func (s *CreateInvoiceShrinkRequest) SetAmount(v string) *CreateInvoiceShrinkRequest {
	s.Amount = &v
	return s
}

func (s *CreateInvoiceShrinkRequest) SetEcIdAccountIdsShrink(v string) *CreateInvoiceShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *CreateInvoiceShrinkRequest) SetInvoiceCandidateIdsShrink(v string) *CreateInvoiceShrinkRequest {
	s.InvoiceCandidateIdsShrink = &v
	return s
}

func (s *CreateInvoiceShrinkRequest) SetInvoiceMode(v int32) *CreateInvoiceShrinkRequest {
	s.InvoiceMode = &v
	return s
}

func (s *CreateInvoiceShrinkRequest) SetInvoiceRemark(v string) *CreateInvoiceShrinkRequest {
	s.InvoiceRemark = &v
	return s
}

func (s *CreateInvoiceShrinkRequest) SetInvoiceTitleId(v string) *CreateInvoiceShrinkRequest {
	s.InvoiceTitleId = &v
	return s
}

func (s *CreateInvoiceShrinkRequest) SetInvoiceType(v int32) *CreateInvoiceShrinkRequest {
	s.InvoiceType = &v
	return s
}

func (s *CreateInvoiceShrinkRequest) SetNbid(v string) *CreateInvoiceShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *CreateInvoiceShrinkRequest) SetRecipientEmailsShrink(v string) *CreateInvoiceShrinkRequest {
	s.RecipientEmailsShrink = &v
	return s
}

func (s *CreateInvoiceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
