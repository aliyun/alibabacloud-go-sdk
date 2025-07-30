// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateChain(v string) *CreateClientCertificateResponseBody
	GetCertificateChain() *string
	SetIdentifier(v string) *CreateClientCertificateResponseBody
	GetIdentifier() *string
	SetRequestId(v string) *CreateClientCertificateResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *CreateClientCertificateResponseBody
	GetSerialNumber() *string
	SetX509Certificate(v string) *CreateClientCertificateResponseBody
	GetX509Certificate() *string
}

type CreateClientCertificateResponseBody struct {
	// The certificate chain of the client certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the client certificate.
	//
	// example:
	//
	// 190ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 8C467B38-3910-447D-87BC-AC049166F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// 0f29522da2dae7a1c4b6ab7132ad3c06
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The content of the client certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s CreateClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateResponseBody) GetCertificateChain() *string {
	return s.CertificateChain
}

func (s *CreateClientCertificateResponseBody) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClientCertificateResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *CreateClientCertificateResponseBody) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *CreateClientCertificateResponseBody) SetCertificateChain(v string) *CreateClientCertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetIdentifier(v string) *CreateClientCertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetRequestId(v string) *CreateClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetSerialNumber(v string) *CreateClientCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *CreateClientCertificateResponseBody) SetX509Certificate(v string) *CreateClientCertificateResponseBody {
	s.X509Certificate = &v
	return s
}

func (s *CreateClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
