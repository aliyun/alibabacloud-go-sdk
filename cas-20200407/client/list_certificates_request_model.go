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
	// example:
	//
	// BUY
	CertificateSource *string `json:"CertificateSource,omitempty" xml:"CertificateSource,omitempty"`
	// example:
	//
	// issued
	CertificateStatus *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// cas-ivauto-hqito6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// test
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// example:
	//
	// rg-aek****wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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
