// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientCertificateWithCsrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertKmcRep1(v string) *CreateClientCertificateWithCsrResponseBody
	GetCertKmcRep1() *string
	SetCertSignBufKmc(v string) *CreateClientCertificateWithCsrResponseBody
	GetCertSignBufKmc() *string
	SetCertificateChain(v string) *CreateClientCertificateWithCsrResponseBody
	GetCertificateChain() *string
	SetIdentifier(v string) *CreateClientCertificateWithCsrResponseBody
	GetIdentifier() *string
	SetRequestId(v string) *CreateClientCertificateWithCsrResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *CreateClientCertificateWithCsrResponseBody
	GetSerialNumber() *string
	SetX509Certificate(v string) *CreateClientCertificateWithCsrResponseBody
	GetX509Certificate() *string
}

type CreateClientCertificateWithCsrResponseBody struct {
	// CertKmcRep1.
	//
	// example:
	//
	// userSeal=MHkCIEu94PQAahFWuFk%
	//
	// ***
	//
	// EtFw%2FkMMBjw8i5bFfSkV%2FIUrcOJD
	CertKmcRep1 *string `json:"CertKmcRep1,omitempty" xml:"CertKmcRep1,omitempty"`
	// Cert Sign Buf Kmc.
	//
	// example:
	//
	// userSeal=MHkCIEu94PQAahFWuFk%
	//
	// ***
	//
	// EtFw%2FkMMBjw8i5bFfSkV%2FIUrcOJD
	CertSignBufKmc *string `json:"CertSignBufKmc,omitempty" xml:"CertSignBufKmc,omitempty"`
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
	// 200ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 31C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the server certificate.
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

func (s CreateClientCertificateWithCsrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateClientCertificateWithCsrResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateWithCsrResponseBody) GetCertKmcRep1() *string {
	return s.CertKmcRep1
}

func (s *CreateClientCertificateWithCsrResponseBody) GetCertSignBufKmc() *string {
	return s.CertSignBufKmc
}

func (s *CreateClientCertificateWithCsrResponseBody) GetCertificateChain() *string {
	return s.CertificateChain
}

func (s *CreateClientCertificateWithCsrResponseBody) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateClientCertificateWithCsrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClientCertificateWithCsrResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *CreateClientCertificateWithCsrResponseBody) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *CreateClientCertificateWithCsrResponseBody) SetCertKmcRep1(v string) *CreateClientCertificateWithCsrResponseBody {
	s.CertKmcRep1 = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetCertSignBufKmc(v string) *CreateClientCertificateWithCsrResponseBody {
	s.CertSignBufKmc = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetCertificateChain(v string) *CreateClientCertificateWithCsrResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetIdentifier(v string) *CreateClientCertificateWithCsrResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetRequestId(v string) *CreateClientCertificateWithCsrResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetSerialNumber(v string) *CreateClientCertificateWithCsrResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) SetX509Certificate(v string) *CreateClientCertificateWithCsrResponseBody {
	s.X509Certificate = &v
	return s
}

func (s *CreateClientCertificateWithCsrResponseBody) Validate() error {
	return dara.Validate(s)
}
