// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadOriginClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *UploadOriginClientCertificateRequest
	GetCertificate() *string
	SetName(v string) *UploadOriginClientCertificateRequest
	GetName() *string
	SetPrivateKey(v string) *UploadOriginClientCertificateRequest
	GetPrivateKey() *string
	SetSiteId(v int64) *UploadOriginClientCertificateRequest
	GetSiteId() *int64
}

type UploadOriginClientCertificateRequest struct {
	// This parameter is required.
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	PrivateKey *string `json:"PrivateKey,omitempty" xml:"PrivateKey,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UploadOriginClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadOriginClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadOriginClientCertificateRequest) GetCertificate() *string {
	return s.Certificate
}

func (s *UploadOriginClientCertificateRequest) GetName() *string {
	return s.Name
}

func (s *UploadOriginClientCertificateRequest) GetPrivateKey() *string {
	return s.PrivateKey
}

func (s *UploadOriginClientCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UploadOriginClientCertificateRequest) SetCertificate(v string) *UploadOriginClientCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *UploadOriginClientCertificateRequest) SetName(v string) *UploadOriginClientCertificateRequest {
	s.Name = &v
	return s
}

func (s *UploadOriginClientCertificateRequest) SetPrivateKey(v string) *UploadOriginClientCertificateRequest {
	s.PrivateKey = &v
	return s
}

func (s *UploadOriginClientCertificateRequest) SetSiteId(v int64) *UploadOriginClientCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *UploadOriginClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
