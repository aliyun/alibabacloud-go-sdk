// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateId(v string) *GetCertificateRequest
	GetCertificateId() *string
}

type GetCertificateRequest struct {
	// The ID of the certificate. It is the globally unique identifier (GUID) of the certificate in Certificates Manager.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9a28de48-8d8b-484d-a766-dec4****
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
}

func (s GetCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateRequest) GoString() string {
	return s.String()
}

func (s *GetCertificateRequest) GetCertificateId() *string {
	return s.CertificateId
}

func (s *GetCertificateRequest) SetCertificateId(v string) *GetCertificateRequest {
	s.CertificateId = &v
	return s
}

func (s *GetCertificateRequest) Validate() error {
	return dara.Validate(s)
}
