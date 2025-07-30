// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateChain(v string) *CreateServerCertificateResponseBody
	GetCertificateChain() *string
	SetIdentifier(v string) *CreateServerCertificateResponseBody
	GetIdentifier() *string
	SetRequestId(v string) *CreateServerCertificateResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *CreateServerCertificateResponseBody
	GetSerialNumber() *string
	SetX509Certificate(v string) *CreateServerCertificateResponseBody
	GetX509Certificate() *string
}

type CreateServerCertificateResponseBody struct {
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
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the server certificate.
	//
	// example:
	//
	// 0f29522da2dae7a1c4b6ab7132ad3c06
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The content of the server certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----\\n......\\n-----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s CreateServerCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServerCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateResponseBody) GetCertificateChain() *string {
	return s.CertificateChain
}

func (s *CreateServerCertificateResponseBody) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateServerCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServerCertificateResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *CreateServerCertificateResponseBody) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *CreateServerCertificateResponseBody) SetCertificateChain(v string) *CreateServerCertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateServerCertificateResponseBody) SetIdentifier(v string) *CreateServerCertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateServerCertificateResponseBody) SetRequestId(v string) *CreateServerCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServerCertificateResponseBody) SetSerialNumber(v string) *CreateServerCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *CreateServerCertificateResponseBody) SetX509Certificate(v string) *CreateServerCertificateResponseBody {
	s.X509Certificate = &v
	return s
}

func (s *CreateServerCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
