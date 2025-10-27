// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlsInspectCertificateDownloadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCaCertId(v string) *GetTlsInspectCertificateDownloadUrlRequest
	GetCaCertId() *string
}

type GetTlsInspectCertificateDownloadUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// C3E91391-16CD-1BFC-A133-******D429
	CaCertId *string `json:"CaCertId,omitempty" xml:"CaCertId,omitempty"`
}

func (s GetTlsInspectCertificateDownloadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTlsInspectCertificateDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetTlsInspectCertificateDownloadUrlRequest) GetCaCertId() *string {
	return s.CaCertId
}

func (s *GetTlsInspectCertificateDownloadUrlRequest) SetCaCertId(v string) *GetTlsInspectCertificateDownloadUrlRequest {
	s.CaCertId = &v
	return s
}

func (s *GetTlsInspectCertificateDownloadUrlRequest) Validate() error {
	return dara.Validate(s)
}
