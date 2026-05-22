// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadClientCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *UploadClientCaCertificateRequest
	GetCertificate() *string
	SetName(v string) *UploadClientCaCertificateRequest
	GetName() *string
	SetSiteId(v int64) *UploadClientCaCertificateRequest
	GetSiteId() *int64
}

type UploadClientCaCertificateRequest struct {
	// The certificate content.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// yourCertName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The website ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UploadClientCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadClientCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadClientCaCertificateRequest) GetCertificate() *string {
	return s.Certificate
}

func (s *UploadClientCaCertificateRequest) GetName() *string {
	return s.Name
}

func (s *UploadClientCaCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UploadClientCaCertificateRequest) SetCertificate(v string) *UploadClientCaCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *UploadClientCaCertificateRequest) SetName(v string) *UploadClientCaCertificateRequest {
	s.Name = &v
	return s
}

func (s *UploadClientCaCertificateRequest) SetSiteId(v int64) *UploadClientCaCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *UploadClientCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
