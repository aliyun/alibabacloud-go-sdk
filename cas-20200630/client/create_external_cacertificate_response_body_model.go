// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalCACertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v string) *CreateExternalCACertificateResponseBody
	GetCertificate() *string
	SetCertificateChain(v string) *CreateExternalCACertificateResponseBody
	GetCertificateChain() *string
	SetIdentifier(v string) *CreateExternalCACertificateResponseBody
	GetIdentifier() *string
	SetRequestId(v string) *CreateExternalCACertificateResponseBody
	GetRequestId() *string
}

type CreateExternalCACertificateResponseBody struct {
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
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// ...
	//
	// ...
	//
	// -----END CERTIFICATE-----
	//
	// -----BEGIN CERTIFICATE-----
	//
	// ...
	//
	// ...
	//
	// -----END CERTIFICATE-----
	CertificateChain *string `json:"CertificateChain,omitempty" xml:"CertificateChain,omitempty"`
	// example:
	//
	// 1ed4068c-6f1b-6deb-8e32-3f8439a851cb
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateExternalCACertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExternalCACertificateResponseBody) GetCertificate() *string {
	return s.Certificate
}

func (s *CreateExternalCACertificateResponseBody) GetCertificateChain() *string {
	return s.CertificateChain
}

func (s *CreateExternalCACertificateResponseBody) GetIdentifier() *string {
	return s.Identifier
}

func (s *CreateExternalCACertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExternalCACertificateResponseBody) SetCertificate(v string) *CreateExternalCACertificateResponseBody {
	s.Certificate = &v
	return s
}

func (s *CreateExternalCACertificateResponseBody) SetCertificateChain(v string) *CreateExternalCACertificateResponseBody {
	s.CertificateChain = &v
	return s
}

func (s *CreateExternalCACertificateResponseBody) SetIdentifier(v string) *CreateExternalCACertificateResponseBody {
	s.Identifier = &v
	return s
}

func (s *CreateExternalCACertificateResponseBody) SetRequestId(v string) *CreateExternalCACertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExternalCACertificateResponseBody) Validate() error {
	return dara.Validate(s)
}
