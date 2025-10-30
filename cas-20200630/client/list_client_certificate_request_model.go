// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListClientCertificateRequest
	GetCurrentPage() *int32
	SetIdentifier(v string) *ListClientCertificateRequest
	GetIdentifier() *string
	SetResourceGroupId(v string) *ListClientCertificateRequest
	GetResourceGroupId() *string
	SetShowSize(v int32) *ListClientCertificateRequest
	GetShowSize() *int32
}

type ListClientCertificateRequest struct {
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The unique identifier of the client certificate or the server certificate that you want to query.
	//
	// >  You can call the [ListClientCertificate](https://help.aliyun.com/document_detail/330884.html) operation to query the unique identifiers of all client certificates and server certificates.
	//
	// example:
	//
	// 190ae6bb538d538c70c01f81dcf2****
	Identifier      *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The number of certificates to return on each page. Default value: **20**.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
}

func (s ListClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *ListClientCertificateRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClientCertificateRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListClientCertificateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClientCertificateRequest) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListClientCertificateRequest) SetCurrentPage(v int32) *ListClientCertificateRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListClientCertificateRequest) SetIdentifier(v string) *ListClientCertificateRequest {
	s.Identifier = &v
	return s
}

func (s *ListClientCertificateRequest) SetResourceGroupId(v string) *ListClientCertificateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClientCertificateRequest) SetShowSize(v int32) *ListClientCertificateRequest {
	s.ShowSize = &v
	return s
}

func (s *ListClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
