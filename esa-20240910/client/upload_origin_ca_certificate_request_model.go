// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadOriginCaCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *UploadOriginCaCertificateRequest
	GetCertificate() *string
	SetName(v string) *UploadOriginCaCertificateRequest
	GetName() *string
	SetSiteId(v int64) *UploadOriginCaCertificateRequest
	GetSiteId() *int64
}

type UploadOriginCaCertificateRequest struct {
	// This parameter is required.
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UploadOriginCaCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadOriginCaCertificateRequest) GoString() string {
	return s.String()
}

func (s *UploadOriginCaCertificateRequest) GetCertificate() *string {
	return s.Certificate
}

func (s *UploadOriginCaCertificateRequest) GetName() *string {
	return s.Name
}

func (s *UploadOriginCaCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UploadOriginCaCertificateRequest) SetCertificate(v string) *UploadOriginCaCertificateRequest {
	s.Certificate = &v
	return s
}

func (s *UploadOriginCaCertificateRequest) SetName(v string) *UploadOriginCaCertificateRequest {
	s.Name = &v
	return s
}

func (s *UploadOriginCaCertificateRequest) SetSiteId(v int64) *UploadOriginCaCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *UploadOriginCaCertificateRequest) Validate() error {
	return dara.Validate(s)
}
