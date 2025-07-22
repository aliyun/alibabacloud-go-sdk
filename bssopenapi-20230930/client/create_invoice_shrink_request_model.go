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
	// example:
	//
	// 0.01
	Amount               *string `json:"Amount,omitempty" xml:"Amount,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// This parameter is required.
	InvoiceCandidateIdsShrink *string `json:"InvoiceCandidateIds,omitempty" xml:"InvoiceCandidateIds,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	InvoiceMode   *int32  `json:"InvoiceMode,omitempty" xml:"InvoiceMode,omitempty"`
	InvoiceRemark *string `json:"InvoiceRemark,omitempty" xml:"InvoiceRemark,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345
	InvoiceTitleId *string `json:"InvoiceTitleId,omitempty" xml:"InvoiceTitleId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 0
	InvoiceType *int32 `json:"InvoiceType,omitempty" xml:"InvoiceType,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
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
