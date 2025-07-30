// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *CreateCustomCertificateResponseBody
	GetCertificate() *string
	SetCertificateChain(v string) *CreateCustomCertificateResponseBody
	GetCertificateChain() *string
	SetIdentifier(v string) *CreateCustomCertificateResponseBody
	GetIdentifier() *string
	SetRequestId(v string) *CreateCustomCertificateResponseBody
	GetRequestId() *string
	SetSerialNumber(v string) *CreateCustomCertificateResponseBody
	GetSerialNumber() *string
}

type CreateCustomCertificateResponseBody struct {
	// The content of the certificate. This parameter is returned only if Immediately is set to 1 or 2.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIEkjCCA3qgAwIBAgIQCgFBQgAAAVOFc2oLheynCDANBgkqhkiG9w0BAQsFADA/
	//
	// ...
	//
	// ...
	//
	// ...
	//
	// KOqkqm57TH2H3eDJAkSnh6/DNFu0Qg==
	//
	// -----END CERTIFICATE-----
	Certificate *string `json:"Certificate,omitempty" xml:"Certificate,omitempty"`
	// The certificate chain of the certificate. This parameter is returned only if Immediately is set to 2.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIBfzCCATGgAwIBAgIUfI5kSdcO2S0+LkpdL3b2VUJG10YwBQYDK2VwMDUxCzAJ
	//
	// ...
	//
	// ...
	//
	// ...
	//
	// ZYYG
	//
	// -----END CERTIFICATE-----
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIIBczCCARgCAQAwgYoxFDASBgNVBAMMC2FsaXl1bi50ZXN0MQ0wCwYDVQQ
	//
	// ...
	//
	// ...
	//
	// ...
	//
	// KL5cUmF
	//
	// -----END CERTIFICATE-----
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the certificate. This parameter is returned only if Immediately is set to 1 or 2.
	//
	// example:
	//
	// 084bde9cd233f0ddae33adc438cfbbbd****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s CreateCustomCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomCertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *CreateCustomCertificateResponseBody) GetCertificateChain() *string {
	return s.CertificateChain
}

func (s *CreateCustomCertificateResponseBody) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateCustomCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomCertificateResponseBody) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *CreateCustomCertificateResponseBody) SetCertificate(v string) *CreateCustomCertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *CreateCustomCertificateResponseBody) SetCertificateChain(v string) *CreateCustomCertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateCustomCertificateResponseBody) SetIdentifier(v string) *CreateCustomCertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateCustomCertificateResponseBody) SetRequestId(v string) *CreateCustomCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomCertificateResponseBody) SetSerialNumber(v string) *CreateCustomCertificateResponseBody {
	s.SerialNumber = &v
	return s
}

func (s *CreateCustomCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
