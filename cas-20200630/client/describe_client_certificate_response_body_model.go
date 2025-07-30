// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v *DescribeClientCertificateResponseBodyCertificate) *DescribeClientCertificateResponseBody
	GetCertificate() *DescribeClientCertificateResponseBodyCertificate
	SetRequestId(v string) *DescribeClientCertificateResponseBody
	GetRequestId() *string
}

type DescribeClientCertificateResponseBody struct {
	// The details about the client certificate or the server certificate.
	Certificate *DescribeClientCertificateResponseBodyCertificate `json:"Certificate,omitempty" xml:"Certificate,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateResponseBody) GetCertificate() *DescribeClientCertificateResponseBodyCertificate {
	return s.Certificate
}

func (s *DescribeClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClientCertificateResponseBody) SetCertificate(v *DescribeClientCertificateResponseBodyCertificate) *DescribeClientCertificateResponseBody {
	s.Certificate = v
	return s
}

func (s *DescribeClientCertificateResponseBody) SetRequestId(v string) *DescribeClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClientCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClientCertificateResponseBodyCertificate struct {
	// The expiration date of the certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1665819958000
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The type of the encryption algorithm of the certificate. Valid values:
	//
	// 	- **RSA**: the Rivest-Shamir-Adleman (RSA) algorithm.
	//
	// 	- **ECC**: the elliptic curve cryptography (ECC) algorithm.
	//
	// 	- **SM2**: the SM2 algorithm, which is developed and approved by the State Cryptography Administration of China.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The issuance date of the certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1634283958000
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **CLIENT**: client certificate
	//
	// 	- **SERVER**: server certificate
	//
	// example:
	//
	// SERVER
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// aliyun.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// For more information about country codes, see the **"Country codes"*	- section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The validity period of the certificate. Unit: days.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The unique identifier of the certificate.
	//
	// example:
	//
	// d3b95700998e47afc4d95f886579****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the certificate.
	//
	// example:
	//
	// 4096
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the certificate.
	//
	// example:
	//
	// d3b95700998e47afc4d95f886579****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Alibaba Cloud Computing Co., Ltd.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department in the organization. The organization is associated with the intermediate certificate authority (CA) certificate from which the certificate is issued.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the intermediate certificate from which the client certificate is issued.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The subject alternative name (SAN) extension of the certificate. The value indicates additional information, including the additional domain names or IP addresses that are associated with the certificate.
	//
	// The value is a string that consists of JSON arrays. Each element in a JSON array is a JSON struct that corresponds to a SAN extension. A SAN extension struct contains the following parameters:
	//
	// 	- **Type**: the type of the extension. Data type: integer. Valid values:
	//
	//     	- **1**: an email address
	//
	//     	- **2**: a domain name
	//
	//     	- **6**: a Uniform Resource Identifier (URI)
	//
	//     	- **7**: an IP address
	//
	// 	- **Value**: the value of the extension. Data type: string.
	//
	// example:
	//
	// [ {"Type": 7, "Value": "192.0.XX.XX"}, {"Type": 2, "Value": "www.aliyundoc.com"}, ]
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// 62b2b943a32d96883a6650e672ea0276****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the certificate.
	//
	// example:
	//
	// 14dcc8afc7578e1fcec36d658f7e20de18f6957bbac42b373a66bc9de4e9****
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the certificate.
	//
	// example:
	//
	// SHA256WITHRSA
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- **ISSUE**: issued
	//
	// 	- **REVOKE**: revoked
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The distinguished name (DN) extension of the certificate, which indicates the user of the certificate. The DN extension includes the following information:
	//
	// 	- **C**: the country
	//
	// 	- **O**: the organization
	//
	// 	- **OU**: the department
	//
	// 	- **L**: the city
	//
	// 	- **ST**: the province, municipality, or autonomous region
	//
	// 	- **CN**: the common name
	//
	// example:
	//
	// C=CN,O=Alibaba Cloud Computing Co., Ltd.,OU=Security,L=Hangzhou,ST=Zhejiang,CN=Aliyun
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	// The content of the certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----  ...... -----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
}

func (s DescribeClientCertificateResponseBodyCertificate) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateResponseBodyCertificate) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetCertificateType() *string {
	return s.CertificateType
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetCountryCode() *string {
	return s.CountryCode
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetDays() *int32 {
	return s.Days
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetKeySize() *int32 {
	return s.KeySize
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetLocality() *string {
	return s.Locality
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetMd5() *string {
	return s.Md5
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetOrganization() *string {
	return s.Organization
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetSans() *string {
	return s.Sans
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetSha2() *string {
	return s.Sha2
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetSignAlgorithm() *string {
	return s.SignAlgorithm
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetState() *string {
	return s.State
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetStatus() *string {
	return s.Status
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetSubjectDN() *string {
	return s.SubjectDN
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetAfterDate(v int64) *DescribeClientCertificateResponseBodyCertificate {
	s.AfterDate = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetAlgorithm(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Algorithm = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetBeforeDate(v int64) *DescribeClientCertificateResponseBodyCertificate {
	s.BeforeDate = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetCertificateType(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.CertificateType = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetCommonName(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.CommonName = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetCountryCode(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.CountryCode = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetDays(v int32) *DescribeClientCertificateResponseBodyCertificate {
	s.Days = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetIdentifier(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Identifier = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetKeySize(v int32) *DescribeClientCertificateResponseBodyCertificate {
	s.KeySize = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetLocality(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Locality = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetMd5(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Md5 = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetOrganization(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Organization = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetOrganizationUnit(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.OrganizationUnit = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetParentIdentifier(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.ParentIdentifier = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetSans(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Sans = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetSerialNumber(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.SerialNumber = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetSha2(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Sha2 = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetSignAlgorithm(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.SignAlgorithm = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetState(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.State = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetStatus(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.Status = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetSubjectDN(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.SubjectDN = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetX509Certificate(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.X509Certificate = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) Validate() error {
	return dara.Validate(s)
}
