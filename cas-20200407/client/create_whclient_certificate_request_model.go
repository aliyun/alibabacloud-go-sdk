// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWHClientCertificateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterTime(v int64) *CreateWHClientCertificateRequest
	GetAfterTime() *int64
	SetAlgorithm(v string) *CreateWHClientCertificateRequest
	GetAlgorithm() *string
	SetBeforeTime(v int64) *CreateWHClientCertificateRequest
	GetBeforeTime() *int64
	SetCommonName(v string) *CreateWHClientCertificateRequest
	GetCommonName() *string
	SetCountry(v string) *CreateWHClientCertificateRequest
	GetCountry() *string
	SetCsr(v string) *CreateWHClientCertificateRequest
	GetCsr() *string
	SetDays(v int64) *CreateWHClientCertificateRequest
	GetDays() *int64
	SetImmediately(v int64) *CreateWHClientCertificateRequest
	GetImmediately() *int64
	SetLocality(v string) *CreateWHClientCertificateRequest
	GetLocality() *string
	SetMonths(v int64) *CreateWHClientCertificateRequest
	GetMonths() *int64
	SetOrganization(v string) *CreateWHClientCertificateRequest
	GetOrganization() *string
	SetOrganizationUnit(v string) *CreateWHClientCertificateRequest
	GetOrganizationUnit() *string
	SetParentIdentifier(v string) *CreateWHClientCertificateRequest
	GetParentIdentifier() *string
	SetSanType(v int64) *CreateWHClientCertificateRequest
	GetSanType() *int64
	SetSanValue(v string) *CreateWHClientCertificateRequest
	GetSanValue() *string
	SetState(v string) *CreateWHClientCertificateRequest
	GetState() *string
	SetYears(v int64) *CreateWHClientCertificateRequest
	GetYears() *int64
}

type CreateWHClientCertificateRequest struct {
	// The expiration time of the client certificate, specified as a Unix timestamp in seconds.
	//
	// > The `BeforeTime` and `AfterTime` parameters must be specified together or not at all.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm for the client certificate. The format is `<encryption_algorithm>_<key_length>`. Valid values:
	//
	// - **RSA_1024**: The corresponding signature algorithm is Sha256WithRSA.
	//
	// - **RSA_2048**: The corresponding signature algorithm is Sha256WithRSA.
	//
	// - **RSA_4096**: The corresponding signature algorithm is Sha256WithRSA.
	//
	// - **ECC_256**: The corresponding signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_384**: The corresponding signature algorithm is Sha256WithECDSA.
	//
	// - **ECC_512**: The corresponding signature algorithm is Sha256WithECDSA.
	//
	// - **SM2_256**: The corresponding signature algorithm is SM3WithSM2.
	//
	// The encryption algorithm of the client certificate must match that of the issuing subordinate CA certificate, but the key lengths can differ. For example, if the key algorithm of the subordinate CA certificate is RSA_2048, the key algorithm for the client certificate must be one of RSA_1024, RSA_2048, or RSA_4096.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the client certificate, as a Unix timestamp in seconds. If omitted, this defaults to the time of the API call.
	//
	// > The `BeforeTime` and `AfterTime` parameters must be specified together or not at all.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The common name of the client certificate. Supports Chinese, English, and other characters.
	//
	// example:
	//
	// aliyun
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The country where the organization is located.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The content of the certificate signing request (CSR). You can generate a CSR with tools like OpenSSL or Keytool.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The validity period of the client certificate, in days.
	//
	// You cannot leave the `Days`, `BeforeTime`, and `AfterTime` parameters all empty. The `BeforeTime` and `AfterTime` parameters must be specified together or not at all.
	//
	// - If you specify the `Days` parameter, specifying `BeforeTime` and `AfterTime` is optional.
	//
	// - If you do not specify the `Days` parameter, you must specify both `BeforeTime` and `AfterTime`.
	//
	// > If you specify `Days`, `BeforeTime`, and `AfterTime` simultaneously, the `Days` parameter takes precedence in determining the validity period.
	//
	// example:
	//
	// 365
	Days *int64 `json:"Days,omitempty" xml:"Days,omitempty"`
	// Specifies which certificate content to return in the response.
	//
	// - **0**: Does not return the certificate (default).
	//
	// - **1**: Returns the certificate.
	//
	// - **2**: Returns the certificate and its certificate chain.
	//
	// example:
	//
	// 1
	Immediately *int64 `json:"Immediately,omitempty" xml:"Immediately,omitempty"`
	// The city where the organization is located. Chinese, English, and other characters are supported.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The validity period of the certificate, in months.
	//
	// example:
	//
	// 12
	Months *int64 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The organization name associated with the root CA certificate, typically your company or enterprise name. Supports Chinese, English, and other characters.
	//
	// example:
	//
	// 阿里巴巴网络技术有限公司
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or business unit within the organization.
	//
	// example:
	//
	// IT
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the issuing subordinate CA certificate.
	//
	// This parameter is required.
	//
	// example:
	//
	// 273ae6bb538d538c70c01f81jh2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The type of the subject alternative name (SAN) for the client certificate. Valid values:
	//
	// - **1**: email address.
	//
	// - **2**: domain name.
	//
	// - **6**: Uniform Resource Identifier (URI).
	//
	// - **7**: IP address.
	//
	// example:
	//
	// 2
	SanType *int64 `json:"SanType,omitempty" xml:"SanType,omitempty"`
	// The value of the SAN extension. To specify multiple values, separate them with commas (,).
	//
	// example:
	//
	// example.com
	SanValue *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
	// The province, municipality, or autonomous region where the organization is located. Chinese, English, and other characters are supported.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the certificate, in years.
	//
	// example:
	//
	// 1
	Years *int64 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateWHClientCertificateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWHClientCertificateRequest) GoString() string {
	return s.String()
}

func (s *CreateWHClientCertificateRequest) GetAfterTime() *int64 {
	return s.AfterTime
}

func (s *CreateWHClientCertificateRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateWHClientCertificateRequest) GetBeforeTime() *int64 {
	return s.BeforeTime
}

func (s *CreateWHClientCertificateRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateWHClientCertificateRequest) GetCountry() *string {
	return s.Country
}

func (s *CreateWHClientCertificateRequest) GetCsr() *string {
	return s.Csr
}

func (s *CreateWHClientCertificateRequest) GetDays() *int64 {
	return s.Days
}

func (s *CreateWHClientCertificateRequest) GetImmediately() *int64 {
	return s.Immediately
}

func (s *CreateWHClientCertificateRequest) GetLocality() *string {
	return s.Locality
}

func (s *CreateWHClientCertificateRequest) GetMonths() *int64 {
	return s.Months
}

func (s *CreateWHClientCertificateRequest) GetOrganization() *string {
	return s.Organization
}

func (s *CreateWHClientCertificateRequest) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateWHClientCertificateRequest) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *CreateWHClientCertificateRequest) GetSanType() *int64 {
	return s.SanType
}

func (s *CreateWHClientCertificateRequest) GetSanValue() *string {
	return s.SanValue
}

func (s *CreateWHClientCertificateRequest) GetState() *string {
	return s.State
}

func (s *CreateWHClientCertificateRequest) GetYears() *int64 {
	return s.Years
}

func (s *CreateWHClientCertificateRequest) SetAfterTime(v int64) *CreateWHClientCertificateRequest {
	s.AfterTime = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetAlgorithm(v string) *CreateWHClientCertificateRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetBeforeTime(v int64) *CreateWHClientCertificateRequest {
	s.BeforeTime = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetCommonName(v string) *CreateWHClientCertificateRequest {
	s.CommonName = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetCountry(v string) *CreateWHClientCertificateRequest {
	s.Country = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetCsr(v string) *CreateWHClientCertificateRequest {
	s.Csr = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetDays(v int64) *CreateWHClientCertificateRequest {
	s.Days = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetImmediately(v int64) *CreateWHClientCertificateRequest {
	s.Immediately = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetLocality(v string) *CreateWHClientCertificateRequest {
	s.Locality = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetMonths(v int64) *CreateWHClientCertificateRequest {
	s.Months = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetOrganization(v string) *CreateWHClientCertificateRequest {
	s.Organization = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetOrganizationUnit(v string) *CreateWHClientCertificateRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetParentIdentifier(v string) *CreateWHClientCertificateRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetSanType(v int64) *CreateWHClientCertificateRequest {
	s.SanType = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetSanValue(v string) *CreateWHClientCertificateRequest {
	s.SanValue = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetState(v string) *CreateWHClientCertificateRequest {
	s.State = &v
	return s
}

func (s *CreateWHClientCertificateRequest) SetYears(v int64) *CreateWHClientCertificateRequest {
	s.Years = &v
	return s
}

func (s *CreateWHClientCertificateRequest) Validate() error {
	return dara.Validate(s)
}
