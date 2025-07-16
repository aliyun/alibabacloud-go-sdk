// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCdnCertificateSigningRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *CreateCdnCertificateSigningRequestRequest
	GetCity() *string
	SetCommonName(v string) *CreateCdnCertificateSigningRequestRequest
	GetCommonName() *string
	SetCountry(v string) *CreateCdnCertificateSigningRequestRequest
	GetCountry() *string
	SetEmail(v string) *CreateCdnCertificateSigningRequestRequest
	GetEmail() *string
	SetOrganization(v string) *CreateCdnCertificateSigningRequestRequest
	GetOrganization() *string
	SetOrganizationUnit(v string) *CreateCdnCertificateSigningRequestRequest
	GetOrganizationUnit() *string
	SetSANs(v string) *CreateCdnCertificateSigningRequestRequest
	GetSANs() *string
	SetState(v string) *CreateCdnCertificateSigningRequestRequest
	GetState() *string
}

type CreateCdnCertificateSigningRequestRequest struct {
	// The city. Default value: Hangzhou.
	//
	// example:
	//
	// Hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The Common Name of the certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// CommonName
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The country or region in which the organization is located. Default value: CN.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The email address.
	//
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	//
	// example:
	//
	// Alibaba Inc
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Aliyun CDN.
	//
	// example:
	//
	// Aliyun CDN
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The Subject Alternative Name (SAN) extension of the SSL certificate. This extension is used to add domain names to the certificate. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// example.com
	SANs *string `json:"SANs,omitempty" xml:"SANs,omitempty"`
	// The provincial district. Default value: Zhejiang.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateCdnCertificateSigningRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCdnCertificateSigningRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateCdnCertificateSigningRequestRequest) GetCity() *string {
	return s.City
}

func (s *CreateCdnCertificateSigningRequestRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateCdnCertificateSigningRequestRequest) GetCountry() *string {
	return s.Country
}

func (s *CreateCdnCertificateSigningRequestRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateCdnCertificateSigningRequestRequest) GetOrganization() *string {
	return s.Organization
}

func (s *CreateCdnCertificateSigningRequestRequest) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateCdnCertificateSigningRequestRequest) GetSANs() *string {
	return s.SANs
}

func (s *CreateCdnCertificateSigningRequestRequest) GetState() *string {
	return s.State
}

func (s *CreateCdnCertificateSigningRequestRequest) SetCity(v string) *CreateCdnCertificateSigningRequestRequest {
	s.City = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetCommonName(v string) *CreateCdnCertificateSigningRequestRequest {
	s.CommonName = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetCountry(v string) *CreateCdnCertificateSigningRequestRequest {
	s.Country = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetEmail(v string) *CreateCdnCertificateSigningRequestRequest {
	s.Email = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetOrganization(v string) *CreateCdnCertificateSigningRequestRequest {
	s.Organization = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetOrganizationUnit(v string) *CreateCdnCertificateSigningRequestRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetSANs(v string) *CreateCdnCertificateSigningRequestRequest {
	s.SANs = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) SetState(v string) *CreateCdnCertificateSigningRequestRequest {
	s.State = &v
	return s
}

func (s *CreateCdnCertificateSigningRequestRequest) Validate() error {
	return dara.Validate(s)
}
