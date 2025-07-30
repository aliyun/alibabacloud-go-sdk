// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerCertificateWithCsrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateChain(v string) *CreateServerCertificateWithCsrResponseBody
	GetCertificateChain() *string
	SetIdentifier(v string) *CreateServerCertificateWithCsrResponseBody
	GetIdentifier() *string
	SetRequestId(v string) *CreateServerCertificateWithCsrResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *CreateServerCertificateWithCsrResponseBody
	GetSerialNumber() *string
	SetX509Certificate(v string) *CreateServerCertificateWithCsrResponseBody
	GetX509Certificate() *string
}

type CreateServerCertificateWithCsrResponseBody struct {
	// The certificate chain of the server certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n-----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----\\n
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the server certificate.
	//
	// example:
	//
	// 180ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 55C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the server certificate.
	//
	// example:
	//
	// 084bde9cd233f0ddae33adc438cfbbbd****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The content of the server certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s CreateServerCertificateWithCsrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServerCertificateWithCsrResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateWithCsrResponseBody) GetCertificateChain() *string {
	return s.CertificateChain
}

func (s *CreateServerCertificateWithCsrResponseBody) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateServerCertificateWithCsrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServerCertificateWithCsrResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *CreateServerCertificateWithCsrResponseBody) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *CreateServerCertificateWithCsrResponseBody) SetCertificateChain(v string) *CreateServerCertificateWithCsrResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) SetIdentifier(v string) *CreateServerCertificateWithCsrResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) SetRequestId(v string) *CreateServerCertificateWithCsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) SetSerialNumber(v string) *CreateServerCertificateWithCsrResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) SetX509Certificate(v string) *CreateServerCertificateWithCsrResponseBody {
	s.X509Certificate = &v
	return s
}

func (s *CreateServerCertificateWithCsrResponseBody) Validate() error {
	return dara.Validate(s)
}
