// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRootCACertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *CreateRootCACertificateResponseBody
	GetCertificate() *string
	SetCertificateChain(v string) *CreateRootCACertificateResponseBody
	GetCertificateChain() *string
	SetIdentifier(v string) *CreateRootCACertificateResponseBody
	GetIdentifier() *string
	SetRequestId(v string) *CreateRootCACertificateResponseBody
	GetRequestId() *string
}

type CreateRootCACertificateResponseBody struct {
	// The root CA certificate in the PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain of the root CA certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the root CA certificate.
	//
	// example:
	//
	// 1a83bcbb89e562885e40aa0108f5****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 6D9B4C5F-7140-5B41-924C-329181DC00C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRootCACertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRootCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRootCACertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *CreateRootCACertificateResponseBody) GetCertificateChain() *string {
	return s.CertificateChain
}

func (s *CreateRootCACertificateResponseBody) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateRootCACertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRootCACertificateResponseBody) SetCertificate(v string) *CreateRootCACertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *CreateRootCACertificateResponseBody) SetCertificateChain(v string) *CreateRootCACertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateRootCACertificateResponseBody) SetIdentifier(v string) *CreateRootCACertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateRootCACertificateResponseBody) SetRequestId(v string) *CreateRootCACertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRootCACertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
