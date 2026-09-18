// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateSource(v string) *ListCertificatesRequest
	GetCertificateSource() *string
	SetCertificateStatus(v string) *ListCertificatesRequest
	GetCertificateStatus() *string
	SetCurrentPage(v int32) *ListCertificatesRequest
	GetCurrentPage() *int32
	SetInstanceId(v string) *ListCertificatesRequest
	GetInstanceId() *string
	SetKeyword(v string) *ListCertificatesRequest
	GetKeyword() *string
	SetResourceGroupId(v string) *ListCertificatesRequest
	GetResourceGroupId() *string
	SetShowSize(v int32) *ListCertificatesRequest
	GetShowSize() *int32
}

type ListCertificatesRequest struct {
	// The source of the certificate. Valid values:
	//
	// - BUY: a formal certificate.
	//
	// - TEST: a test certificate.
	//
	// - UPLOAD: an uploaded certificate.
	//
	// example:
	//
	// BUY
	CertificateSource *string `json:"CertificateSource,omitempty" xml:"CertificateSource,omitempty"`
	// The status of the certificate. Valid values:
	//
	// - **issued**: Issued.
	//
	// - **revoked**: Revoked.
	//
	// - **willExpire**: About to expire.
	//
	// - **expired**: Expired.
	//
	// example:
	//
	// issued
	CertificateStatus *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	// The page number of the current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cas-ivauto-hqito6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The keyword for fuzzy match. The keyword is matched against domain names, names, and corresponding resource IDs.
	//
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aek****wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListCertificatesRequest) GetCertificateSource() *string {
	return s.CertificateSource
}

func (s *ListCertificatesRequest) GetCertificateStatus() *string {
	return s.CertificateStatus
}

func (s *ListCertificatesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCertificatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCertificatesRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *ListCertificatesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListCertificatesRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCertificatesRequest) SetCertificateSource(v string) *ListCertificatesRequest {
	s.CertificateSource = &v
	return s
}

func (s *ListCertificatesRequest) SetCertificateStatus(v string) *ListCertificatesRequest {
	s.CertificateStatus = &v
	return s
}

func (s *ListCertificatesRequest) SetCurrentPage(v int32) *ListCertificatesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCertificatesRequest) SetInstanceId(v string) *ListCertificatesRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCertificatesRequest) SetKeyword(v string) *ListCertificatesRequest {
	s.Keyword = &v
	return s
}

func (s *ListCertificatesRequest) SetResourceGroupId(v string) *ListCertificatesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListCertificatesRequest) SetShowSize(v int32) *ListCertificatesRequest {
	s.ShowSize = &v
	return s
}

func (s *ListCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
