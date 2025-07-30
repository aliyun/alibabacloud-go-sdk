// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServerCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterTime(v int64) *CreateServerCertificateRequest
	GetAfterTime() *int64
	SetAlgorithm(v string) *CreateServerCertificateRequest
	GetAlgorithm() *string
	SetBeforeTime(v int64) *CreateServerCertificateRequest
	GetBeforeTime() *int64
	SetCommonName(v string) *CreateServerCertificateRequest
	GetCommonName() *string
	SetCountry(v string) *CreateServerCertificateRequest
	GetCountry() *string
	SetDays(v int32) *CreateServerCertificateRequest
	GetDays() *int32
	SetDomain(v string) *CreateServerCertificateRequest
	GetDomain() *string
	SetEnableCrl(v int64) *CreateServerCertificateRequest
	GetEnableCrl() *int64
	SetImmediately(v int32) *CreateServerCertificateRequest
	GetImmediately() *int32
	SetLocality(v string) *CreateServerCertificateRequest
	GetLocality() *string
	SetMonths(v int32) *CreateServerCertificateRequest
	GetMonths() *int32
	SetOrganization(v string) *CreateServerCertificateRequest
	GetOrganization() *string
	SetOrganizationUnit(v string) *CreateServerCertificateRequest
	GetOrganizationUnit() *string
	SetParentIdentifier(v string) *CreateServerCertificateRequest
	GetParentIdentifier() *string
	SetState(v string) *CreateServerCertificateRequest
	GetState() *string
	SetYears(v int32) *CreateServerCertificateRequest
	GetYears() *int32
}

type CreateServerCertificateRequest struct {
	// The expiration time of the server certificate. This value is a UNIX timestamp. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the server certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
	//
	// 	- **RSA_1024**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **RSA_2048**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **RSA_4096**: The signature algorithm is Sha256WithRSA.
	//
	// 	- **ECC_256**: The signature algorithm is Sha256WithECDSA.
	//
	// 	- **ECC_384**: The signature algorithm is Sha256WithECDSA.
	//
	// 	- **ECC_512**: The signature algorithm is Sha256WithECDSA.
	//
	// 	- **SM2_256**: The signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the server certificate must be the same as the encryption algorithm of the intermediate CA certificate. The key length can be different. For example, if the key algorithm of the intermediate CA certificate is RSA_2048, the key algorithm of the server certificate must be RSA_1024, RSA_2048, or RSA_4096.
	//
	// >  You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/328096.html) operation to query the key algorithm of an intermediate CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the server certificate. This value is a UNIX timestamp. The default value is the time when you call this operation. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The name of the certificate user. The user of a server certificate is a server. We recommend that you enter the domain name or IP address of the server.
	//
	// This parameter is required.
	//
	// example:
	//
	// www.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located, such as CN or US.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The validity period of the server certificate. Unit: days. You must specify at least one of the **Days**, **BeforeTime**, and **AfterTime*	- parameters. The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified. The following list describes how to specify these parameters:
	//
	// 	- If you specify the **Days*	- parameter, you can specify both the **BeforeTime*	- and **AfterTime*	- parameters or leave them both empty.
	//
	// 	- If you do not specify the **Days*	- parameter, you must specify both the **BeforeTime*	- and **AfterTime*	- parameters.
	//
	// >
	//
	// 	- If you specify the **Days**, **BeforeTime**, and **AfterTime*	- parameters together, the validity period of the server certificate is determined by the value of the **Days*	- parameter.
	//
	// 	- The validity period of the server certificate cannot exceed the validity period of the intermediate CA certificate. You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/328096.html) operation to query the validity period of an intermediate CA certificate.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The additional domain names and additional IP addresses of the server certificate. After you add additional domain names and additional IP addresses to a certificate, you can apply the certificate to the domain names and IP addresses.
	//
	// Separate multiple domain names and multiple IP addresses with commas (,).
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// include the CRL address.
	//
	// - 0- No
	//
	// - 1- Yes
	//
	// example:
	//
	// 1
	EnableCrl *int64 `json:"EnableCrl,omitempty" xml:"EnableCrl,omitempty"`
	// Specifies whether to return the certificate. Valid values:
	//
	// 	- **0**: does not return the certificate. This is the default value.
	//
	// 	- **1**: returns the certificate.
	//
	// 	- **2**: returns the certificate and the certificate chain of the certificate.
	//
	// example:
	//
	// 1
	Immediately *int32 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The name of the city in which the organization is located. The value can contain letters. The default value is the name of the city in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The validity period of the server certificate. Unit: months.
	//
	// example:
	//
	// 12
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	//
	// example:
	//
	// Alibaba Cloud
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Aliyun CDN.
	//
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate CA certificate from which the server certificate is issued.
	//
	// >  You can call the [DescribeCACertificateList](https://help.aliyun.com/document_detail/328095.html) operation to query the unique identifier of an intermediate CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 271ae6bb538d538c70c01f81dg3****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the server certificate. Unit: years.
	//
	// example:
	//
	// 1
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateServerCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServerCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateServerCertificateRequest) GetAfterTime() *int64 {
	return s.AfterTime
}

func (s *CreateServerCertificateRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateServerCertificateRequest) GetBeforeTime() *int64 {
	return s.BeforeTime
}

func (s *CreateServerCertificateRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateServerCertificateRequest) GetCountry() *string {
	return s.Country
}

func (s *CreateServerCertificateRequest) GetDays() *int32 {
	return s.Days
}

func (s *CreateServerCertificateRequest) GetDomain() *string {
	return s.Domain
}

func (s *CreateServerCertificateRequest) GetEnableCrl() *int64 {
	return s.EnableCrl
}

func (s *CreateServerCertificateRequest) GetImmediately() *int32 {
	return s.Immediately
}

func (s *CreateServerCertificateRequest) GetLocality() *string {
	return s.Locality
}

func (s *CreateServerCertificateRequest) GetMonths() *int32 {
	return s.Months
}

func (s *CreateServerCertificateRequest) GetOrganization() *string {
	return s.Organization
}

func (s *CreateServerCertificateRequest) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateServerCertificateRequest) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *CreateServerCertificateRequest) GetState() *string {
	return s.State
}

func (s *CreateServerCertificateRequest) GetYears() *int32 {
	return s.Years
}

func (s *CreateServerCertificateRequest) SetAfterTime(v int64) *CreateServerCertificateRequest {
	s.AfterTime = &v
	return s
}

func (s *CreateServerCertificateRequest) SetAlgorithm(v string) *CreateServerCertificateRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateServerCertificateRequest) SetBeforeTime(v int64) *CreateServerCertificateRequest {
	s.BeforeTime = &v
	return s
}

func (s *CreateServerCertificateRequest) SetCommonName(v string) *CreateServerCertificateRequest {
	s.CommonName = &v
	return s
}

func (s *CreateServerCertificateRequest) SetCountry(v string) *CreateServerCertificateRequest {
	s.Country = &v
	return s
}

func (s *CreateServerCertificateRequest) SetDays(v int32) *CreateServerCertificateRequest {
	s.Days = &v
	return s
}

func (s *CreateServerCertificateRequest) SetDomain(v string) *CreateServerCertificateRequest {
	s.Domain = &v
	return s
}

func (s *CreateServerCertificateRequest) SetEnableCrl(v int64) *CreateServerCertificateRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateServerCertificateRequest) SetImmediately(v int32) *CreateServerCertificateRequest {
	s.Immediately = &v
	return s
}

func (s *CreateServerCertificateRequest) SetLocality(v string) *CreateServerCertificateRequest {
	s.Locality = &v
	return s
}

func (s *CreateServerCertificateRequest) SetMonths(v int32) *CreateServerCertificateRequest {
	s.Months = &v
	return s
}

func (s *CreateServerCertificateRequest) SetOrganization(v string) *CreateServerCertificateRequest {
	s.Organization = &v
	return s
}

func (s *CreateServerCertificateRequest) SetOrganizationUnit(v string) *CreateServerCertificateRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateServerCertificateRequest) SetParentIdentifier(v string) *CreateServerCertificateRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateServerCertificateRequest) SetState(v string) *CreateServerCertificateRequest {
	s.State = &v
	return s
}

func (s *CreateServerCertificateRequest) SetYears(v int32) *CreateServerCertificateRequest {
	s.Years = &v
	return s
}

func (s *CreateServerCertificateRequest) Validate() error {
	return dara.Validate(s)
}
