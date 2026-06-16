// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSubCACertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *CreateSubCACertificateResponseBody
	GetCertificate() *string
	SetCertificateChain(v string) *CreateSubCACertificateResponseBody
	GetCertificateChain() *string
	SetIdentifier(v string) *CreateSubCACertificateResponseBody
	GetIdentifier() *string
	SetRequestId(v string) *CreateSubCACertificateResponseBody
	GetRequestId() *string
}

type CreateSubCACertificateResponseBody struct {
	// The certificate returned by this call, in PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The CA certificate chain of the certificate that is returned by the call.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the intermediate CA certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSubCACertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSubCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSubCACertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *CreateSubCACertificateResponseBody) GetCertificateChain() *string {
	return s.CertificateChain
}

func (s *CreateSubCACertificateResponseBody) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateSubCACertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSubCACertificateResponseBody) SetCertificate(v string) *CreateSubCACertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *CreateSubCACertificateResponseBody) SetCertificateChain(v string) *CreateSubCACertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateSubCACertificateResponseBody) SetIdentifier(v string) *CreateSubCACertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateSubCACertificateResponseBody) SetRequestId(v string) *CreateSubCACertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSubCACertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
