// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInvoiceCandidateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBillingCycles(v []*int32) *ListInvoiceCandidateRequest
	GetBillingCycles() []*int32
	SetBusinessIds(v []*string) *ListInvoiceCandidateRequest
	GetBusinessIds() []*string
	SetCurrentPage(v int32) *ListInvoiceCandidateRequest
	GetCurrentPage() *int32
	SetEcIdAccountIds(v []*ListInvoiceCandidateRequestEcIdAccountIds) *ListInvoiceCandidateRequest
	GetEcIdAccountIds() []*ListInvoiceCandidateRequestEcIdAccountIds
	SetEndTime(v string) *ListInvoiceCandidateRequest
	GetEndTime() *string
	SetInvoiceIssuers(v []*string) *ListInvoiceCandidateRequest
	GetInvoiceIssuers() []*string
	SetNbid(v string) *ListInvoiceCandidateRequest
	GetNbid() *string
	SetPageSize(v int32) *ListInvoiceCandidateRequest
	GetPageSize() *int32
	SetStartTime(v string) *ListInvoiceCandidateRequest
	GetStartTime() *string
	SetStatus(v []*int32) *ListInvoiceCandidateRequest
	GetStatus() []*int32
	SetTypes(v []*int32) *ListInvoiceCandidateRequest
	GetTypes() []*int32
}

type ListInvoiceCandidateRequest struct {
	BillingCycles []*int32  `json:"BillingCycles,omitempty" xml:"BillingCycles,omitempty" type:"Repeated"`
	BusinessIds   []*string `json:"BusinessIds,omitempty" xml:"BusinessIds,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage    *int32                                       `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	EcIdAccountIds []*ListInvoiceCandidateRequestEcIdAccountIds `json:"EcIdAccountIds,omitempty" xml:"EcIdAccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 2025-07-01 00:00:00
	EndTime        *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	InvoiceIssuers []*string `json:"InvoiceIssuers,omitempty" xml:"InvoiceIssuers,omitempty" type:"Repeated"`
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
	StartTime *string  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Status    []*int32 `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	Types     []*int32 `json:"Types,omitempty" xml:"Types,omitempty" type:"Repeated"`
}

func (s ListInvoiceCandidateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInvoiceCandidateRequest) GoString() string {
	return s.String()
}

func (s *ListInvoiceCandidateRequest) GetBillingCycles() []*int32 {
	return s.BillingCycles
}

func (s *ListInvoiceCandidateRequest) GetBusinessIds() []*string {
	return s.BusinessIds
}

func (s *ListInvoiceCandidateRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInvoiceCandidateRequest) GetEcIdAccountIds() []*ListInvoiceCandidateRequestEcIdAccountIds {
	return s.EcIdAccountIds
}

func (s *ListInvoiceCandidateRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListInvoiceCandidateRequest) GetInvoiceIssuers() []*string {
	return s.InvoiceIssuers
}

func (s *ListInvoiceCandidateRequest) GetNbid() *string {
	return s.Nbid
}

func (s *ListInvoiceCandidateRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListInvoiceCandidateRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListInvoiceCandidateRequest) GetStatus() []*int32 {
	return s.Status
}

func (s *ListInvoiceCandidateRequest) GetTypes() []*int32 {
	return s.Types
}

func (s *ListInvoiceCandidateRequest) SetBillingCycles(v []*int32) *ListInvoiceCandidateRequest {
	s.BillingCycles = v
	return s
}

func (s *ListInvoiceCandidateRequest) SetBusinessIds(v []*string) *ListInvoiceCandidateRequest {
	s.BusinessIds = v
	return s
}

func (s *ListInvoiceCandidateRequest) SetCurrentPage(v int32) *ListInvoiceCandidateRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListInvoiceCandidateRequest) SetEcIdAccountIds(v []*ListInvoiceCandidateRequestEcIdAccountIds) *ListInvoiceCandidateRequest {
	s.EcIdAccountIds = v
	return s
}

func (s *ListInvoiceCandidateRequest) SetEndTime(v string) *ListInvoiceCandidateRequest {
	s.EndTime = &v
	return s
}

func (s *ListInvoiceCandidateRequest) SetInvoiceIssuers(v []*string) *ListInvoiceCandidateRequest {
	s.InvoiceIssuers = v
	return s
}

func (s *ListInvoiceCandidateRequest) SetNbid(v string) *ListInvoiceCandidateRequest {
	s.Nbid = &v
	return s
}

func (s *ListInvoiceCandidateRequest) SetPageSize(v int32) *ListInvoiceCandidateRequest {
	s.PageSize = &v
	return s
}

func (s *ListInvoiceCandidateRequest) SetStartTime(v string) *ListInvoiceCandidateRequest {
	s.StartTime = &v
	return s
}

func (s *ListInvoiceCandidateRequest) SetStatus(v []*int32) *ListInvoiceCandidateRequest {
	s.Status = v
	return s
}

func (s *ListInvoiceCandidateRequest) SetTypes(v []*int32) *ListInvoiceCandidateRequest {
	s.Types = v
	return s
}

func (s *ListInvoiceCandidateRequest) Validate() error {
	return dara.Validate(s)
}

type ListInvoiceCandidateRequestEcIdAccountIds struct {
	AccountIds []*int64 `json:"AccountIds,omitempty" xml:"AccountIds,omitempty" type:"Repeated"`
	// example:
	//
	// 12345
	EcId *string `json:"EcId,omitempty" xml:"EcId,omitempty"`
}

func (s ListInvoiceCandidateRequestEcIdAccountIds) String() string {
	return dara.Prettify(s)
}

func (s ListInvoiceCandidateRequestEcIdAccountIds) GoString() string {
	return s.String()
}

func (s *ListInvoiceCandidateRequestEcIdAccountIds) GetAccountIds() []*int64 {
	return s.AccountIds
}

func (s *ListInvoiceCandidateRequestEcIdAccountIds) GetEcId() *string {
	return s.EcId
}

func (s *ListInvoiceCandidateRequestEcIdAccountIds) SetAccountIds(v []*int64) *ListInvoiceCandidateRequestEcIdAccountIds {
	s.AccountIds = v
	return s
}

func (s *ListInvoiceCandidateRequestEcIdAccountIds) SetEcId(v string) *ListInvoiceCandidateRequestEcIdAccountIds {
	s.EcId = &v
	return s
}

func (s *ListInvoiceCandidateRequestEcIdAccountIds) Validate() error {
	return dara.Validate(s)
}
