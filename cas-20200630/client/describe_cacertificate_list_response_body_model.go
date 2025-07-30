// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCACertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateList(v []*DescribeCACertificateListResponseBodyCertificateList) *DescribeCACertificateListResponseBody
	GetCertificateList() []*DescribeCACertificateListResponseBodyCertificateList
	SetCurrentPage(v int32) *DescribeCACertificateListResponseBody
	GetCurrentPage() *int32
	SetPageCount(v int32) *DescribeCACertificateListResponseBody
	GetPageCount() *int32
	SetRequestId(v string) *DescribeCACertificateListResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *DescribeCACertificateListResponseBody
	GetShowSize() *int32
	SetTotalCount(v int32) *DescribeCACertificateListResponseBody
	GetTotalCount() *int32
}

type DescribeCACertificateListResponseBody struct {
	// The details about the CA certificates.
	CertificateList []*DescribeCACertificateListResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of returned pages.
	//
	// example:
	//
	// 1
	PageCount *int32 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of CA certificates returned per page.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of root CA certificates and intermediate CA certificates that are returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCACertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateListResponseBody) GetCertificateList() []*DescribeCACertificateListResponseBodyCertificateList {
	return s.CertificateList
}

func (s *DescribeCACertificateListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeCACertificateListResponseBody) GetPageCount() *int32 {
	return s.PageCount
}

func (s *DescribeCACertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCACertificateListResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *DescribeCACertificateListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCACertificateListResponseBody) SetCertificateList(v []*DescribeCACertificateListResponseBodyCertificateList) *DescribeCACertificateListResponseBody {
	s.CertificateList = v
	return s
}

func (s *DescribeCACertificateListResponseBody) SetCurrentPage(v int32) *DescribeCACertificateListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeCACertificateListResponseBody) SetPageCount(v int32) *DescribeCACertificateListResponseBody {
	s.PageCount = &v
	return s
}

func (s *DescribeCACertificateListResponseBody) SetRequestId(v string) *DescribeCACertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCACertificateListResponseBody) SetShowSize(v int32) *DescribeCACertificateListResponseBody {
	s.ShowSize = &v
	return s
}

func (s *DescribeCACertificateListResponseBody) SetTotalCount(v int32) *DescribeCACertificateListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCACertificateListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCACertificateListResponseBodyCertificateList struct {
	// The expiration date of the CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1665819958000
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The encryption algorithm of the CA certificate. Valid values:
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
	// The alias of the CA.
	//
	// example:
	//
	// Aliyun_CA
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The issuance date of the CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1634283958000
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the CA certificate. Valid values:
	//
	// 	- **ROOT**: a root CA certificate.
	//
	// 	- **SUB_ROOT**: an intermediate CA certificate.
	//
	// example:
	//
	// SUB_ROOT
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name or abbreviation of the organization that is associated with the CA certificate.
	//
	// example:
	//
	// Aliyun
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located.
	//
	// For more information about country codes, see the **"Country codes"*	- section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	Gift        *int32  `json:"Gift,omitempty" xml:"Gift,omitempty"`
	// The unique identifier of the CA certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// The key length of the CA certificate.
	//
	// example:
	//
	// 2048
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The name of the city in which the organization is located.
	//
	// example:
	//
	// Hangzhou
	Locality *string `json:"Locality,omitempty" xml:"Locality,omitempty"`
	// The MD5 fingerprint of the CA certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Md5 *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	// The name of the organization that is associated with the CA certificate.
	//
	// example:
	//
	// Alibaba Cloud Computing Co., Ltd.
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The name of the department or branch in the organization that is associated with the CA certificate.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the root CA certificate from which the CA certificate is issued.
	//
	// >  This parameter is returned only if the value of the **CertificateType*	- parameter is **SUB_ROOT**. The value SUB_ROOT indicates an intermediate CA certificate.
	//
	// example:
	//
	// 1a83bcbb89e562885e40aa0108f5****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// This parameter is deprecated.
	//
	// example:
	//
	// 1
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The serial number of the CA certificate.
	//
	// example:
	//
	// 70e3b2566d92805173767869727fb92e****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the CA certificate.
	//
	// example:
	//
	// 14dcc8afc7578e1fcec36d658f7e20de18f6957bbac42b373a66bc9de4e9****
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signature algorithm of the CA certificate.
	//
	// example:
	//
	// SHA256WITHRSA
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
	// The name of the province, municipality, or autonomous region in which the organization is located.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the CA certificate. Valid values:
	//
	// 	- **ISSUE**: The CA certificate is issued.
	//
	// 	- **REVOKE**: The CA certificate is revoked.
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The Distinguished Name (DN) attribute of the CA certificate, which indicates the user information of the certificate. The DN attribute contains the following information:
	//
	// 	- **C**: the code of the country in which the organization is located.
	//
	// 	- **O**: the name of the organization.
	//
	// 	- **OU**: the name of the department or branch in the organization.
	//
	// 	- **L**: the name of the city in which the organization is located.
	//
	// 	- **CN**: the common name or abbreviation of the organization.
	//
	// example:
	//
	// C=CN,O=Alibaba Cloud Computing Co., Ltd.,OU=Security,L=Hangzhou,ST=Zhejiang,CN=Aliyun
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	Trial     *int32  `json:"Trial,omitempty" xml:"Trial,omitempty"`
	// The content of the CA certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- …… -----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
	// The validity period of the CA certificate. Unit: years.
	//
	// example:
	//
	// 3
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s DescribeCACertificateListResponseBodyCertificateList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificateListResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetAlias() *string {
	return s.Alias
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetCertificateType() *string {
	return s.CertificateType
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetCountryCode() *string {
	return s.CountryCode
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetGift() *int32 {
	return s.Gift
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetKeySize() *int32 {
	return s.KeySize
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetLocality() *string {
	return s.Locality
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetMd5() *string {
	return s.Md5
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetOrganization() *string {
	return s.Organization
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetSans() *string {
	return s.Sans
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetSha2() *string {
	return s.Sha2
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetSignAlgorithm() *string {
	return s.SignAlgorithm
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetState() *string {
	return s.State
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetStatus() *string {
	return s.Status
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetSubjectDN() *string {
	return s.SubjectDN
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetTrial() *int32 {
	return s.Trial
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *DescribeCACertificateListResponseBodyCertificateList) GetYears() *int32 {
	return s.Years
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetAfterDate(v int64) *DescribeCACertificateListResponseBodyCertificateList {
	s.AfterDate = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetAlgorithm(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Algorithm = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetAlias(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Alias = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetBeforeDate(v int64) *DescribeCACertificateListResponseBodyCertificateList {
	s.BeforeDate = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetCertificateType(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.CertificateType = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetCommonName(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.CommonName = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetCountryCode(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.CountryCode = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetGift(v int32) *DescribeCACertificateListResponseBodyCertificateList {
	s.Gift = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetIdentifier(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Identifier = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetKeySize(v int32) *DescribeCACertificateListResponseBodyCertificateList {
	s.KeySize = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetLocality(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Locality = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetMd5(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Md5 = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetOrganization(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Organization = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetOrganizationUnit(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.OrganizationUnit = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetParentIdentifier(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.ParentIdentifier = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetSans(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Sans = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetSerialNumber(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.SerialNumber = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetSha2(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Sha2 = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetSignAlgorithm(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.SignAlgorithm = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetState(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.State = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetStatus(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.Status = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetSubjectDN(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.SubjectDN = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetTrial(v int32) *DescribeCACertificateListResponseBodyCertificateList {
	s.Trial = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetX509Certificate(v string) *DescribeCACertificateListResponseBodyCertificateList {
	s.X509Certificate = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) SetYears(v int32) *DescribeCACertificateListResponseBodyCertificateList {
	s.Years = &v
	return s
}

func (s *DescribeCACertificateListResponseBodyCertificateList) Validate() error {
	return dara.Validate(s)
}
