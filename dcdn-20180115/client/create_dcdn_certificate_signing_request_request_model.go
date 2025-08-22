// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnCertificateSigningRequestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCity(v string) *CreateDcdnCertificateSigningRequestRequest
	GetCity() *string
	SetCommonName(v string) *CreateDcdnCertificateSigningRequestRequest
	GetCommonName() *string
	SetCountry(v string) *CreateDcdnCertificateSigningRequestRequest
	GetCountry() *string
	SetEmail(v string) *CreateDcdnCertificateSigningRequestRequest
	GetEmail() *string
	SetOrganization(v string) *CreateDcdnCertificateSigningRequestRequest
	GetOrganization() *string
	SetOrganizationUnit(v string) *CreateDcdnCertificateSigningRequestRequest
	GetOrganizationUnit() *string
	SetSANs(v string) *CreateDcdnCertificateSigningRequestRequest
	GetSANs() *string
	SetState(v string) *CreateDcdnCertificateSigningRequestRequest
	GetState() *string
}

type CreateDcdnCertificateSigningRequestRequest struct {
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
	// testName
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
	// test@aliyundoc.com
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
	// The Subject Alternative Name (SAN) extension that allows multiple domain names to be associated with the certificate. Separate multiple domain names with commas (,).
	//
	// example:
	//
	// example.com,example.org
	SANs *string `json:"SANs,omitempty" xml:"SANs,omitempty"`
	// The provincial district. Default value: Zhejiang.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CreateDcdnCertificateSigningRequestRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnCertificateSigningRequestRequest) GoString() string {
	return s.String()
}

func (s *CreateDcdnCertificateSigningRequestRequest) GetCity() *string {
	return s.City
}

func (s *CreateDcdnCertificateSigningRequestRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateDcdnCertificateSigningRequestRequest) GetCountry() *string {
	return s.Country
}

func (s *CreateDcdnCertificateSigningRequestRequest) GetEmail() *string {
	return s.Email
}

func (s *CreateDcdnCertificateSigningRequestRequest) GetOrganization() *string {
	return s.Organization
}

func (s *CreateDcdnCertificateSigningRequestRequest) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateDcdnCertificateSigningRequestRequest) GetSANs() *string {
	return s.SANs
}

func (s *CreateDcdnCertificateSigningRequestRequest) GetState() *string {
	return s.State
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetCity(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.City = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetCommonName(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.CommonName = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetCountry(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.Country = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetEmail(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.Email = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetOrganization(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.Organization = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetOrganizationUnit(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetSANs(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.SANs = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) SetState(v string) *CreateDcdnCertificateSigningRequestRequest {
	s.State = &v
	return s
}

func (s *CreateDcdnCertificateSigningRequestRequest) Validate() error {
	return dara.Validate(s)
}
