// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCSR(v string) *CreateClientCertificateRequest
	GetCSR() *string
	SetPkeyType(v string) *CreateClientCertificateRequest
	GetPkeyType() *string
	SetSiteId(v int64) *CreateClientCertificateRequest
	GetSiteId() *int64
	SetValidityDays(v int64) *CreateClientCertificateRequest
	GetValidityDays() *int64
}

type CreateClientCertificateRequest struct {
	// The certificate signing request (CSR).
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----
	CSR *string `json:"CSR,omitempty" xml:"CSR,omitempty"`
	// The type of the private key algorithm.
	//
	// example:
	//
	// RSA
	PkeyType *string `json:"PkeyType,omitempty" xml:"PkeyType,omitempty"`
	// The website ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	// The validity period of the certificate. Unit: day.
	//
	// This parameter is required.
	//
	// example:
	//
	// 365
	ValidityDays *int64 `json:"ValidityDays,omitempty" xml:"ValidityDays,omitempty"`
}

func (s CreateClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateRequest) GetCSR() *string {
	return s.CSR
}

func (s *CreateClientCertificateRequest) GetPkeyType() *string {
	return s.PkeyType
}

func (s *CreateClientCertificateRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateClientCertificateRequest) GetValidityDays() *int64 {
	return s.ValidityDays
}

func (s *CreateClientCertificateRequest) SetCSR(v string) *CreateClientCertificateRequest {
	s.CSR = &v
	return s
}

func (s *CreateClientCertificateRequest) SetPkeyType(v string) *CreateClientCertificateRequest {
	s.PkeyType = &v
	return s
}

func (s *CreateClientCertificateRequest) SetSiteId(v int64) *CreateClientCertificateRequest {
	s.SiteId = &v
	return s
}

func (s *CreateClientCertificateRequest) SetValidityDays(v int64) *CreateClientCertificateRequest {
	s.ValidityDays = &v
	return s
}

func (s *CreateClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
