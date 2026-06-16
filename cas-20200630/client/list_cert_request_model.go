// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterDate(v string) *ListCertRequest
	GetAfterDate() *string
	SetBeforeDate(v string) *ListCertRequest
	GetBeforeDate() *string
	SetCurrentPage(v int32) *ListCertRequest
	GetCurrentPage() *int32
	SetInstanceUuid(v string) *ListCertRequest
	GetInstanceUuid() *string
	SetMaxResults(v int32) *ListCertRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListCertRequest
	GetNextToken() *string
	SetParentIdentifier(v string) *ListCertRequest
	GetParentIdentifier() *string
	SetShowSize(v int32) *ListCertRequest
	GetShowSize() *int32
	SetStatus(v string) *ListCertRequest
	GetStatus() *string
	SetType(v string) *ListCertRequest
	GetType() *string
}

type ListCertRequest struct {
	// Filters certificates modified after this date.
	//
	// example:
	//
	// 2024-05-13 12:59:45
	AfterDate *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// Filters certificates modified before this date.
	//
	// example:
	//
	// 2025-09-04
	BeforeDate *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The UUID of the instance.
	//
	// example:
	//
	// 1ef79512-569b-6a4e-9105-9b91473562f7
	InstanceUuid *string `json:"InstanceUuid,omitempty" xml:"InstanceUuid,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to retrieve the next page of results. This is the NextToken value from a previous response. If unspecified, the first page is returned.
	//
	// example:
	//
	// 1d2db86sca4384811e0b5e8707e68181f
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The identifier of the intermediate CA that issued the certificate.
	//
	// example:
	//
	// 273ae6bb538d538c70c01f81jh2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The page size. Default value: 50.
	//
	// example:
	//
	// 50
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The status of the certificate. Valid values:
	//
	// - ISSUE: Active
	//
	// - REVOKE: Revoked
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The certificate type. Valid values:
	//
	// - SERVER: Server certificate
	//
	// - CLIENT: Client certificate
	//
	// - END_ENTITY: End-entity certificate
	//
	// example:
	//
	// CLIENT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListCertRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCertRequest) GoString() string {
	return s.String()
}

func (s *ListCertRequest) GetAfterDate() *string {
	return s.AfterDate
}

func (s *ListCertRequest) GetBeforeDate() *string {
	return s.BeforeDate
}

func (s *ListCertRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCertRequest) GetInstanceUuid() *string {
	return s.InstanceUuid
}

func (s *ListCertRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListCertRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListCertRequest) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *ListCertRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCertRequest) GetStatus() *string {
	return s.Status
}

func (s *ListCertRequest) GetType() *string {
	return s.Type
}

func (s *ListCertRequest) SetAfterDate(v string) *ListCertRequest {
	s.AfterDate = &v
	return s
}

func (s *ListCertRequest) SetBeforeDate(v string) *ListCertRequest {
	s.BeforeDate = &v
	return s
}

func (s *ListCertRequest) SetCurrentPage(v int32) *ListCertRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCertRequest) SetInstanceUuid(v string) *ListCertRequest {
	s.InstanceUuid = &v
	return s
}

func (s *ListCertRequest) SetMaxResults(v int32) *ListCertRequest {
	s.MaxResults = &v
	return s
}

func (s *ListCertRequest) SetNextToken(v string) *ListCertRequest {
	s.NextToken = &v
	return s
}

func (s *ListCertRequest) SetParentIdentifier(v string) *ListCertRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *ListCertRequest) SetShowSize(v int32) *ListCertRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCertRequest) SetStatus(v string) *ListCertRequest {
	s.Status = &v
	return s
}

func (s *ListCertRequest) SetType(v string) *ListCertRequest {
	s.Type = &v
	return s
}

func (s *ListCertRequest) Validate() error {
	return dara.Validate(s)
}
