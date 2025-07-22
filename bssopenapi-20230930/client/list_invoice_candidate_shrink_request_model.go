// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInvoiceCandidateShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingCyclesShrink(v string) *ListInvoiceCandidateShrinkRequest
	GetBillingCyclesShrink() *string
	SetBusinessIdsShrink(v string) *ListInvoiceCandidateShrinkRequest
	GetBusinessIdsShrink() *string
	SetCurrentPage(v int32) *ListInvoiceCandidateShrinkRequest
	GetCurrentPage() *int32
	SetEcIdAccountIdsShrink(v string) *ListInvoiceCandidateShrinkRequest
	GetEcIdAccountIdsShrink() *string
	SetEndTime(v string) *ListInvoiceCandidateShrinkRequest
	GetEndTime() *string
	SetInvoiceIssuersShrink(v string) *ListInvoiceCandidateShrinkRequest
	GetInvoiceIssuersShrink() *string
	SetNbid(v string) *ListInvoiceCandidateShrinkRequest
	GetNbid() *string
	SetPageSize(v int32) *ListInvoiceCandidateShrinkRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListInvoiceCandidateShrinkRequest
	GetStartTime() *string
	SetStatusShrink(v string) *ListInvoiceCandidateShrinkRequest
	GetStatusShrink() *string
	SetTypesShrink(v string) *ListInvoiceCandidateShrinkRequest
	GetTypesShrink() *string
}

type ListInvoiceCandidateShrinkRequest struct {
	BillingCyclesShrink *string `json:"BillingCycles,omitempty" xml:"BillingCycles,omitempty"`
	BusinessIdsShrink   *string `json:"BusinessIds,omitempty" xml:"BusinessIds,omitempty"`
	// example:
	//
	// 1
	CurrentPage          *int32  `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIdsShrink *string `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty"`
	// example:
	//
	// 2025-07-01 00:00:00
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InvoiceIssuersShrink *string `json:"InvoiceIssuers,omitempty" xml:"InvoiceIssuers,omitempty"`
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 2025-06-01 00:00:00
	StartTime    *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StatusShrink *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TypesShrink  *string `json:"Types,omitempty" xml:"Types,omitempty"`
}

func (s ListInvoiceCandidateShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInvoiceCandidateShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListInvoiceCandidateShrinkRequest) GetBillingCyclesShrink() *string {
	return s.BillingCyclesShrink
}

func (s *ListInvoiceCandidateShrinkRequest) GetBusinessIdsShrink() *string {
	return s.BusinessIdsShrink
}

func (s *ListInvoiceCandidateShrinkRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInvoiceCandidateShrinkRequest) GetEcIdAccountIdsShrink() *string {
	return s.EcIdAccountIdsShrink
}

func (s *ListInvoiceCandidateShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListInvoiceCandidateShrinkRequest) GetInvoiceIssuersShrink() *string {
	return s.InvoiceIssuersShrink
}

func (s *ListInvoiceCandidateShrinkRequest) GetNbid() *string {
	return s.Nbid
}

func (s *ListInvoiceCandidateShrinkRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInvoiceCandidateShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListInvoiceCandidateShrinkRequest) GetStatusShrink() *string {
	return s.StatusShrink
}

func (s *ListInvoiceCandidateShrinkRequest) GetTypesShrink() *string {
	return s.TypesShrink
}

func (s *ListInvoiceCandidateShrinkRequest) SetBillingCyclesShrink(v string) *ListInvoiceCandidateShrinkRequest {
	s.BillingCyclesShrink = &v
	return s
}

func (s *ListInvoiceCandidateShrinkRequest) SetBusinessIdsShrink(v string) *ListInvoiceCandidateShrinkRequest {
	s.BusinessIdsShrink = &v
	return s
}

func (s *ListInvoiceCandidateShrinkRequest) SetCurrentPage(v int32) *ListInvoiceCandidateShrinkRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListInvoiceCandidateShrinkRequest) SetEcIdAccountIdsShrink(v string) *ListInvoiceCandidateShrinkRequest {
	s.EcIdAccountIdsShrink = &v
	return s
}

func (s *ListInvoiceCandidateShrinkRequest) SetEndTime(v string) *ListInvoiceCandidateShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListInvoiceCandidateShrinkRequest) SetInvoiceIssuersShrink(v string) *ListInvoiceCandidateShrinkRequest {
	s.InvoiceIssuersShrink = &v
	return s
}

func (s *ListInvoiceCandidateShrinkRequest) SetNbid(v string) *ListInvoiceCandidateShrinkRequest {
	s.Nbid = &v
	return s
}

func (s *ListInvoiceCandidateShrinkRequest) SetPageSize(v int32) *ListInvoiceCandidateShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListInvoiceCandidateShrinkRequest) SetStartTime(v string) *ListInvoiceCandidateShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListInvoiceCandidateShrinkRequest) SetStatusShrink(v string) *ListInvoiceCandidateShrinkRequest {
	s.StatusShrink = &v
	return s
}

func (s *ListInvoiceCandidateShrinkRequest) SetTypesShrink(v string) *ListInvoiceCandidateShrinkRequest {
	s.TypesShrink = &v
	return s
}

func (s *ListInvoiceCandidateShrinkRequest) Validate() error {
	return dara.Validate(s)
}
