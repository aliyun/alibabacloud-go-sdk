// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRevokeCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateList(v []*ListRevokeCertificateResponseBodyCertificateList) *ListRevokeCertificateResponseBody
	GetCertificateList() []*ListRevokeCertificateResponseBodyCertificateList
	SetCurrentPage(v int32) *ListRevokeCertificateResponseBody
	GetCurrentPage() *int32
	SetPageCount(v int32) *ListRevokeCertificateResponseBody
	GetPageCount() *int32
	SetRequestId(v string) *ListRevokeCertificateResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListRevokeCertificateResponseBody
	GetShowSize() *int32
	SetTotalCount(v int64) *ListRevokeCertificateResponseBody
	GetTotalCount() *int64
}

type ListRevokeCertificateResponseBody struct {
	// An array that consists of the details about the revoked client certificates or server certificates.
	CertificateList []*ListRevokeCertificateResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// The page number of the current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The total number of pages returned.
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
	// The number of revoked certificates that are returned per page.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of revoked client certificates and server certificates that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRevokeCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRevokeCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ListRevokeCertificateResponseBody) GetCertificateList() []*ListRevokeCertificateResponseBodyCertificateList {
	return s.CertificateList
}

func (s *ListRevokeCertificateResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListRevokeCertificateResponseBody) GetPageCount() *int32 {
	return s.PageCount
}

func (s *ListRevokeCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRevokeCertificateResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListRevokeCertificateResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRevokeCertificateResponseBody) SetCertificateList(v []*ListRevokeCertificateResponseBodyCertificateList) *ListRevokeCertificateResponseBody {
	s.CertificateList = v
	return s
}

func (s *ListRevokeCertificateResponseBody) SetCurrentPage(v int32) *ListRevokeCertificateResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListRevokeCertificateResponseBody) SetPageCount(v int32) *ListRevokeCertificateResponseBody {
	s.PageCount = &v
	return s
}

func (s *ListRevokeCertificateResponseBody) SetRequestId(v string) *ListRevokeCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRevokeCertificateResponseBody) SetShowSize(v int32) *ListRevokeCertificateResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListRevokeCertificateResponseBody) SetTotalCount(v int64) *ListRevokeCertificateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRevokeCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListRevokeCertificateResponseBodyCertificateList struct {
	// The expiration date of the certificate. The date is in the `yyyy-MM-ddT00:00Z` format. For example, the value `2021-12-31T00:00Z` indicates December 31, 2021.
	//
	// example:
	//
	// 2021-12-31T00:00Z
	AfterDate *string `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
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
	// The issuance date of the certificate. The date is in the `yyyy-MM-ddT00:00Z` format. For example, the value `2021-01-01T00:00Z` indicates January 1, 2021.
	//
	// example:
	//
	// 2021-01-01T00:00Z
	BeforeDate *string `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The type of the certificate.
	//
	// example:
	//
	// SERVER
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// aliyundoc.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// For more information about country codes, see the **"Country codes"*	- section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The unique identifier of the certificate.
	//
	// example:
	//
	// 05e148d8d3ecc9976d9ecd2b2f25****
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
	// 05e148d8d3ecc9976d9ecd2b2f25****
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
	// The identifier of the root certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The date on which the certificate was revoked. The value is in the `yyyy-MM-ddT00:00Z` format. For example, the value `2021-09-01T00:00Z` indicates September 1, 2021.
	//
	// example:
	//
	// 2021-09-01T00:00Z
	RevokeDate *string `json:"RevokeDate,omitempty" xml:"RevokeDate,omitempty"`
	// The subject alternative name (SAN) extension of the certificate.
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
	// 168b12c42e62339f8d2340ff530f9365****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The SHA-256 fingerprint of the certificate.
	//
	// example:
	//
	// b60eff7e04323ff662f9ab5e6986f849f626a9c7bf2c59dcc752fa23779a****
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
	// The status.
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
	// C=CN,O=Alibaba Cloud Computing Co., Ltd.,OU=Security,L=ZheJiang,ST=HangZhou,CN=aliyundoc.com
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
}

func (s ListRevokeCertificateResponseBodyCertificateList) String() string {
	return dara.Prettify(s)
}

func (s ListRevokeCertificateResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetAfterDate() *string {
	return s.AfterDate
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetBeforeDate() *string {
	return s.BeforeDate
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetCertificateType() *string {
	return s.CertificateType
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetCommonName() *string {
	return s.CommonName
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetCountryCode() *string {
	return s.CountryCode
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetKeySize() *int32 {
	return s.KeySize
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetLocality() *string {
	return s.Locality
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetMd5() *string {
	return s.Md5
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetOrganization() *string {
	return s.Organization
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetRevokeDate() *string {
	return s.RevokeDate
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetSans() *string {
	return s.Sans
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetSha2() *string {
	return s.Sha2
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetSignAlgorithm() *string {
	return s.SignAlgorithm
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetState() *string {
	return s.State
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetStatus() *string {
	return s.Status
}

func (s *ListRevokeCertificateResponseBodyCertificateList) GetSubjectDN() *string {
	return s.SubjectDN
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetAfterDate(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.AfterDate = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetAlgorithm(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Algorithm = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetBeforeDate(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.BeforeDate = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetCertificateType(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.CertificateType = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetCommonName(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.CommonName = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetCountryCode(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.CountryCode = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetIdentifier(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Identifier = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetKeySize(v int32) *ListRevokeCertificateResponseBodyCertificateList {
	s.KeySize = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetLocality(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Locality = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetMd5(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Md5 = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetOrganization(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Organization = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetOrganizationUnit(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.OrganizationUnit = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetParentIdentifier(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.ParentIdentifier = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetRevokeDate(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.RevokeDate = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetSans(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Sans = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetSerialNumber(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.SerialNumber = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetSha2(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Sha2 = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetSignAlgorithm(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.SignAlgorithm = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetState(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.State = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetStatus(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.Status = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) SetSubjectDN(v string) *ListRevokeCertificateResponseBodyCertificateList {
	s.SubjectDN = &v
	return s
}

func (s *ListRevokeCertificateResponseBodyCertificateList) Validate() error {
	return dara.Validate(s)
}
