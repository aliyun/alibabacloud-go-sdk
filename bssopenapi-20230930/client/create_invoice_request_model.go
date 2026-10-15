// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInvoiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAmount(v string) *CreateInvoiceRequest
	GetAmount() *string
	SetEcIdAccountIds(v []*CreateInvoiceRequestEcIdAccountIds) *CreateInvoiceRequest
	GetEcIdAccountIds() []*CreateInvoiceRequestEcIdAccountIds
	SetInvoiceCandidateIds(v []*string) *CreateInvoiceRequest
	GetInvoiceCandidateIds() []*string
	SetInvoiceMode(v int32) *CreateInvoiceRequest
	GetInvoiceMode() *int32
	SetInvoiceRemark(v string) *CreateInvoiceRequest
	GetInvoiceRemark() *string
	SetInvoiceTitleId(v string) *CreateInvoiceRequest
	GetInvoiceTitleId() *string
	SetInvoiceType(v int32) *CreateInvoiceRequest
	GetInvoiceType() *int32
	SetNbid(v string) *CreateInvoiceRequest
	GetNbid() *string
	SetRecipientEmails(v []*string) *CreateInvoiceRequest
	GetRecipientEmails() []*string
}

type CreateInvoiceRequest struct {
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
	EcIdAccountIds []*CreateInvoiceRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// The list of invoice candidate IDs.
	//
	// This parameter is required.
	InvoiceCandidateIds []*string `json:"InvoiceCandidateIds,omitempty" xml:"InvoiceCandidateIds,omitempty" type:"Repeated"`
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
	RecipientEmails []*string `json:"RecipientEmails,omitempty" xml:"RecipientEmails,omitempty" type:"Repeated"`
}

func (s CreateInvoiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInvoiceRequest) GoString() string {
	return s.String()
}

func (s *CreateInvoiceRequest) GetAmount() *string {
	return s.Amount
}

func (s *CreateInvoiceRequest) GetEcIdAccountIds() []*CreateInvoiceRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *CreateInvoiceRequest) GetInvoiceCandidateIds() []*string {
	return s.InvoiceCandidateIds
}

func (s *CreateInvoiceRequest) GetInvoiceMode() *int32 {
	return s.InvoiceMode
}

func (s *CreateInvoiceRequest) GetInvoiceRemark() *string {
	return s.InvoiceRemark
}

func (s *CreateInvoiceRequest) GetInvoiceTitleId() *string {
	return s.InvoiceTitleId
}

func (s *CreateInvoiceRequest) GetInvoiceType() *int32 {
	return s.InvoiceType
}

func (s *CreateInvoiceRequest) GetNbid() *string {
	return s.Nbid
}

func (s *CreateInvoiceRequest) GetRecipientEmails() []*string {
	return s.RecipientEmails
}

func (s *CreateInvoiceRequest) SetAmount(v string) *CreateInvoiceRequest {
	s.Amount = &v
	return s
}

func (s *CreateInvoiceRequest) SetEcIdAccountIds(v []*CreateInvoiceRequestEcIdAccountIds) *CreateInvoiceRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *CreateInvoiceRequest) SetInvoiceCandidateIds(v []*string) *CreateInvoiceRequest {
	s.InvoiceCandidateIds = v
	return s
}

func (s *CreateInvoiceRequest) SetInvoiceMode(v int32) *CreateInvoiceRequest {
	s.InvoiceMode = &v
	return s
}

func (s *CreateInvoiceRequest) SetInvoiceRemark(v string) *CreateInvoiceRequest {
	s.InvoiceRemark = &v
	return s
}

func (s *CreateInvoiceRequest) SetInvoiceTitleId(v string) *CreateInvoiceRequest {
	s.InvoiceTitleId = &v
	return s
}

func (s *CreateInvoiceRequest) SetInvoiceType(v int32) *CreateInvoiceRequest {
	s.InvoiceType = &v
	return s
}

func (s *CreateInvoiceRequest) SetNbid(v string) *CreateInvoiceRequest {
	s.Nbid = &v
	return s
}

func (s *CreateInvoiceRequest) SetRecipientEmails(v []*string) *CreateInvoiceRequest {
	s.RecipientEmails = v
	return s
}

func (s *CreateInvoiceRequest) Validate() error {
	if s.EcIdAccountIds != nil {
		for _, item := range s.EcIdAccountIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateInvoiceRequestEcIdAccountIds struct {
	// The list of accounts to access. If this parameter is left empty, all accounts under the current entity ID are selected.
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// The enterprise entity ID.
	//
	// example:
	//
	// 12345
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s CreateInvoiceRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s CreateInvoiceRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *CreateInvoiceRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *CreateInvoiceRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *CreateInvoiceRequestEcIdAccountIds) SetAccountIds(v []*int64) *CreateInvoiceRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *CreateInvoiceRequestEcIdAccountIds) SetEcId(v string) *CreateInvoiceRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *CreateInvoiceRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
