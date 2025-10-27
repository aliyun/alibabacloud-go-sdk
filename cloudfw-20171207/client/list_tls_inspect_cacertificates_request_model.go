// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTlsInspectCACertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaCertId(v string) *ListTlsInspectCACertificatesRequest
	GetCaCertId() *string
	SetCurrentPage(v int32) *ListTlsInspectCACertificatesRequest
	GetCurrentPage() *int32
	SetPageSize(v int32) *ListTlsInspectCACertificatesRequest
	GetPageSize() *int32
}

type ListTlsInspectCACertificatesRequest struct {
	// example:
	//
	// C3E91391-16CD-1BFC-A133-******D429
	CaCertId *string `json:"CaCertId,omitempty" xml:"CaCertId,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListTlsInspectCACertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTlsInspectCACertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListTlsInspectCACertificatesRequest) GetCaCertId() *string {
	return s.CaCertId
}

func (s *ListTlsInspectCACertificatesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListTlsInspectCACertificatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListTlsInspectCACertificatesRequest) SetCaCertId(v string) *ListTlsInspectCACertificatesRequest {
	s.CaCertId = &v
	return s
}

func (s *ListTlsInspectCACertificatesRequest) SetCurrentPage(v int32) *ListTlsInspectCACertificatesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListTlsInspectCACertificatesRequest) SetPageSize(v int32) *ListTlsInspectCACertificatesRequest {
	s.PageSize = &v
	return s
}

func (s *ListTlsInspectCACertificatesRequest) Validate() error {
	return dara.Validate(s)
}
