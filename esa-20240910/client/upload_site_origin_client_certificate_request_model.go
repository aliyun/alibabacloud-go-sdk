// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSiteOriginClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *UploadSiteOriginClientCertificateRequest
	GetCertificate() *string
	SetName(v string) *UploadSiteOriginClientCertificateRequest
	GetName() *string
	SetPrivateKey(v string) *UploadSiteOriginClientCertificateRequest
	GetPrivateKey() *string
	SetSiteId(v int64) *UploadSiteOriginClientCertificateRequest
	GetSiteId() *int64
}

type UploadSiteOriginClientCertificateRequest struct {
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
	// example
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The private key of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// -----BEGIN PRIVATE KEY-----
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// Site ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456789****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UploadSiteOriginClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadSiteOriginClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadSiteOriginClientCertificateRequest) GetCertificate() *string {
	return s.Certificate
}

func (s *UploadSiteOriginClientCertificateRequest) GetName() *string {
	return s.Name
}

func (s *UploadSiteOriginClientCertificateRequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *UploadSiteOriginClientCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UploadSiteOriginClientCertificateRequest) SetCertificate(v string) *UploadSiteOriginClientCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *UploadSiteOriginClientCertificateRequest) SetName(v string) *UploadSiteOriginClientCertificateRequest {
	s.Name = &v
	return s
}

func (s *UploadSiteOriginClientCertificateRequest) SetPrivateKey(v string) *UploadSiteOriginClientCertificateRequest {
	s.PrivateKey = &v
	return s
}

func (s *UploadSiteOriginClientCertificateRequest) SetSiteId(v int64) *UploadSiteOriginClientCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *UploadSiteOriginClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
