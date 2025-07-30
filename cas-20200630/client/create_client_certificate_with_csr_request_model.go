// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClientCertificateWithCsrRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterTime(v int64) *CreateClientCertificateWithCsrRequest
	GetAfterTime() *int64
	SetAlgorithm(v string) *CreateClientCertificateWithCsrRequest
	GetAlgorithm() *string
	SetBeforeTime(v int64) *CreateClientCertificateWithCsrRequest
	GetBeforeTime() *int64
	SetCommonName(v string) *CreateClientCertificateWithCsrRequest
	GetCommonName() *string
	SetCountry(v string) *CreateClientCertificateWithCsrRequest
	GetCountry() *string
	SetCsr(v string) *CreateClientCertificateWithCsrRequest
	GetCsr() *string
	SetDays(v int32) *CreateClientCertificateWithCsrRequest
	GetDays() *int32
	SetEnableCrl(v int64) *CreateClientCertificateWithCsrRequest
	GetEnableCrl() *int64
	SetImmediately(v int32) *CreateClientCertificateWithCsrRequest
	GetImmediately() *int32
	SetLocality(v string) *CreateClientCertificateWithCsrRequest
	GetLocality() *string
	SetMonths(v int32) *CreateClientCertificateWithCsrRequest
	GetMonths() *int32
	SetOrganization(v string) *CreateClientCertificateWithCsrRequest
	GetOrganization() *string
	SetOrganizationUnit(v string) *CreateClientCertificateWithCsrRequest
	GetOrganizationUnit() *string
	SetParentIdentifier(v string) *CreateClientCertificateWithCsrRequest
	GetParentIdentifier() *string
	SetSanType(v int32) *CreateClientCertificateWithCsrRequest
	GetSanType() *int32
	SetSanValue(v string) *CreateClientCertificateWithCsrRequest
	GetSanValue() *string
	SetState(v string) *CreateClientCertificateWithCsrRequest
	GetState() *string
	SetYears(v int32) *CreateClientCertificateWithCsrRequest
	GetYears() *int32
}

type CreateClientCertificateWithCsrRequest struct {
	// The expiration time of the client certificate. This value is a UNIX timestamp. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1665819958
	AfterTime *int64 `json:"AfterTime,omitempty" xml:"AfterTime,omitempty"`
	// The key algorithm of the client certificate. The key algorithm is in the `<Encryption algorithm>_<Key length>` format. Valid values:
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
	// The encryption algorithm of the client certificate must be the same with the encryption algorithm of the intermediate CA certificate. The key length can be different. For example, if the key algorithm of the intermediate CA certificate is RSA_2048, the key algorithm of the client certificate must be RSA_1024, RSA_2048, or RSA_4096.
	//
	// >  You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/328096.html) operation to query the key algorithm of an intermediate CA certificate.
	//
	// example:
	//
	// RSA_2048
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance time of the client certificate. This value is a UNIX timestamp. The default value is the time when you call this operation. Unit: seconds.
	//
	// >  The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified.
	//
	// example:
	//
	// 1634283958
	BeforeTime *int64 `json:"BeforeTime,omitempty" xml:"BeforeTime,omitempty"`
	// The common name of the certificate. The value can contain letters.
	//
	// >  If you specify the **CsrPemString*	- parameter, the value of the **CommonName*	- parameter is determined by the **CsrPemString*	- parameter.
	//
	// example:
	//
	// aliyundoc.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located, such as **CN*	- and **US**.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The content of the CSR file. You can generate a CSR file by using the OpenSSL tool or Keytool. For more information, see [How do I create a CSR file?](https://help.aliyun.com/document_detail/42218.html) You can also create a CSR file in the Certificate Management Service console. For more information, see [Create a CSR](https://help.aliyun.com/document_detail/313297.html).
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----   ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The validity period of the client certificate. Unit: days. You must specify at least one of the **Days**, **BeforeTime**, and **AfterTime*	- parameters. The **BeforeTime*	- and **AfterTime*	- parameters must be both empty or both specified. The following list describes how to specify these parameters:
	//
	// 	- If you specify the **Days*	- parameter, you can specify both the **BeforeTime*	- and **AfterTime*	- parameters or leave them both empty.********
	//
	// 	- If you do not specify the **Days*	- parameter, you must specify both the **BeforeTime*	- and **AfterTime*	- parameters.
	//
	// >
	//
	// 	- If you specify the **Days**, **BeforeTime**, and **AfterTime*	- parameters together, the validity period of the client certificate is determined by the value of the **Days*	- parameter.
	//
	// 	- The validity period of the client certificate cannot exceed the validity period of the intermediate CA certificate. You can call the [DescribeCACertificate](https://help.aliyun.com/document_detail/328096.html) operation to query the validity period of an intermediate CA certificate.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
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
	// The validity period of the client certificate. Unit: months.
	//
	// example:
	//
	// 12
	Months *int32 `json:"Months,omitempty" xml:"Months,omitempty"`
	// The name of the organization. Default value: Alibaba Inc.
	//
	// example:
	//
	// Alibaba Cloud Computing Co., Ltd.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department. Default value: Aliyun CDN.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate CA certificate from which the client certificate is issued.
	//
	// >  You can call the [DescribeCACertificateList](https://help.aliyun.com/document_detail/328095.html) operation to query the unique identifier of an intermediate CA certificate.
	//
	// example:
	//
	// 270ae6bb538d538c70c01f81fg3****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The type of the Subject Alternative Name (SAN) extension that is supported by the client certificate. Valid values:
	//
	// 	- **1**: an email address
	//
	// 	- **6**: a Uniform Resource Identifier (URI)
	//
	// example:
	//
	// 1
	SanType *int32 `json:"SanType,omitempty" xml:"SanType,omitempty"`
	// The content of the extension. You can specify multiple SAN extensions. If you want to specify multiple SAN extensions, separate them with commas (,).
	//
	// example:
	//
	// somebody@example.com
	SanValue *string `json:"SanValue,omitempty" xml:"SanValue,omitempty"`
	// The province, municipality, or autonomous region in which the organization is located. The value can contain letters. The default value is the name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate CA certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The validity period of the client certificate. Unit: years.
	//
	// example:
	//
	// 1
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s CreateClientCertificateWithCsrRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClientCertificateWithCsrRequest) GoString() string {
	return s.String()
}

func (s *CreateClientCertificateWithCsrRequest) GetAfterTime() *int64 {
	return s.AfterTime
}

func (s *CreateClientCertificateWithCsrRequest) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *CreateClientCertificateWithCsrRequest) GetBeforeTime() *int64 {
	return s.BeforeTime
}

func (s *CreateClientCertificateWithCsrRequest) GetCommonName() *string {
	return s.CommonName
}

func (s *CreateClientCertificateWithCsrRequest) GetCountry() *string {
	return s.Country
}

func (s *CreateClientCertificateWithCsrRequest) GetCsr() *string {
	return s.Csr
}

func (s *CreateClientCertificateWithCsrRequest) GetDays() *int32 {
	return s.Days
}

func (s *CreateClientCertificateWithCsrRequest) GetEnableCrl() *int64 {
	return s.EnableCrl
}

func (s *CreateClientCertificateWithCsrRequest) GetImmediately() *int32 {
	return s.Immediately
}

func (s *CreateClientCertificateWithCsrRequest) GetLocality() *string {
	return s.Locality
}

func (s *CreateClientCertificateWithCsrRequest) GetMonths() *int32 {
	return s.Months
}

func (s *CreateClientCertificateWithCsrRequest) GetOrganization() *string {
	return s.Organization
}

func (s *CreateClientCertificateWithCsrRequest) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *CreateClientCertificateWithCsrRequest) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *CreateClientCertificateWithCsrRequest) GetSanType() *int32 {
	return s.SanType
}

func (s *CreateClientCertificateWithCsrRequest) GetSanValue() *string {
	return s.SanValue
}

func (s *CreateClientCertificateWithCsrRequest) GetState() *string {
	return s.State
}

func (s *CreateClientCertificateWithCsrRequest) GetYears() *int32 {
	return s.Years
}

func (s *CreateClientCertificateWithCsrRequest) SetAfterTime(v int64) *CreateClientCertificateWithCsrRequest {
	s.AfterTime = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetAlgorithm(v string) *CreateClientCertificateWithCsrRequest {
	s.Algorithm = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetBeforeTime(v int64) *CreateClientCertificateWithCsrRequest {
	s.BeforeTime = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetCommonName(v string) *CreateClientCertificateWithCsrRequest {
	s.CommonName = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetCountry(v string) *CreateClientCertificateWithCsrRequest {
	s.Country = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetCsr(v string) *CreateClientCertificateWithCsrRequest {
	s.Csr = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetDays(v int32) *CreateClientCertificateWithCsrRequest {
	s.Days = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetEnableCrl(v int64) *CreateClientCertificateWithCsrRequest {
	s.EnableCrl = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetImmediately(v int32) *CreateClientCertificateWithCsrRequest {
	s.Immediately = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetLocality(v string) *CreateClientCertificateWithCsrRequest {
	s.Locality = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetMonths(v int32) *CreateClientCertificateWithCsrRequest {
	s.Months = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetOrganization(v string) *CreateClientCertificateWithCsrRequest {
	s.Organization = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetOrganizationUnit(v string) *CreateClientCertificateWithCsrRequest {
	s.OrganizationUnit = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetParentIdentifier(v string) *CreateClientCertificateWithCsrRequest {
	s.ParentIdentifier = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetSanType(v int32) *CreateClientCertificateWithCsrRequest {
	s.SanType = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetSanValue(v string) *CreateClientCertificateWithCsrRequest {
	s.SanValue = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetState(v string) *CreateClientCertificateWithCsrRequest {
	s.State = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) SetYears(v int32) *CreateClientCertificateWithCsrRequest {
	s.Years = &v
	return s
}

func (s *CreateClientCertificateWithCsrRequest) Validate() error {
	return dara.Validate(s)
}
