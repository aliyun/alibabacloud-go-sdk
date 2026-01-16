// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateList(v []*ListClientCertificateResponseBodyCertificateList) *ListClientCertificateResponseBody
	GetCertificateList() []*ListClientCertificateResponseBodyCertificateList
	SetCurrentPage(v int32) *ListClientCertificateResponseBody
	GetCurrentPage() *int32
	SetMaxResults(v int32) *ListClientCertificateResponseBody
	GetMaxResults() *int32
	SetPageCount(v int32) *ListClientCertificateResponseBody
	GetPageCount() *int32
	SetRequestId(v string) *ListClientCertificateResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListClientCertificateResponseBody
	GetShowSize() *int32
	SetTotalCount(v int64) *ListClientCertificateResponseBody
	GetTotalCount() *int64
}

type ListClientCertificateResponseBody struct {
	// An array that consists of the details about all client certificates and server certificates.
	CertificateList []*ListClientCertificateResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// The page number of the current page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	MaxResults  *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
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
	// The number of certificates that are returned per page.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The number of client certificates and server certificates that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListClientCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListClientCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *ListClientCertificateResponseBody) GetCertificateList() []*ListClientCertificateResponseBodyCertificateList {
	return s.CertificateList
}

func (s *ListClientCertificateResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListClientCertificateResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListClientCertificateResponseBody) GetPageCount() *int32 {
	return s.PageCount
}

func (s *ListClientCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListClientCertificateResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListClientCertificateResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListClientCertificateResponseBody) SetCertificateList(v []*ListClientCertificateResponseBodyCertificateList) *ListClientCertificateResponseBody {
	s.CertificateList = v
	return s
}

func (s *ListClientCertificateResponseBody) SetCurrentPage(v int32) *ListClientCertificateResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListClientCertificateResponseBody) SetMaxResults(v int32) *ListClientCertificateResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListClientCertificateResponseBody) SetPageCount(v int32) *ListClientCertificateResponseBody {
	s.PageCount = &v
	return s
}

func (s *ListClientCertificateResponseBody) SetRequestId(v string) *ListClientCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListClientCertificateResponseBody) SetShowSize(v int32) *ListClientCertificateResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListClientCertificateResponseBody) SetTotalCount(v int64) *ListClientCertificateResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListClientCertificateResponseBody) Validate() error {
	if s.CertificateList != nil {
		for _, item := range s.CertificateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClientCertificateResponseBodyCertificateList struct {
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
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
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
	// aliyundoc.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The code of the country in which the organization is located. The organization is associated with the intermediate certificate from which the certificate is issued.
	//
	// For more information about country codes, see the **"Country codes"*	- section of the [Manage company profiles](https://help.aliyun.com/document_detail/198289.html) topic.
	//
	// example:
	//
	// CN
	CountryCode      *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The validity period of the certificate. Unit: days.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	Id   *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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

func (s ListClientCertificateResponseBodyCertificateList) String() string {
	return dara.Prettify(s)
}

func (s ListClientCertificateResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *ListClientCertificateResponseBodyCertificateList) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *ListClientCertificateResponseBodyCertificateList) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListClientCertificateResponseBodyCertificateList) GetAliasName() *string {
	return s.AliasName
}

func (s *ListClientCertificateResponseBodyCertificateList) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *ListClientCertificateResponseBodyCertificateList) GetCertificateType() *string {
	return s.CertificateType
}

func (s *ListClientCertificateResponseBodyCertificateList) GetCommonName() *string {
	return s.CommonName
}

func (s *ListClientCertificateResponseBodyCertificateList) GetCountryCode() *string {
	return s.CountryCode
}

func (s *ListClientCertificateResponseBodyCertificateList) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *ListClientCertificateResponseBodyCertificateList) GetDays() *int32 {
	return s.Days
}

func (s *ListClientCertificateResponseBodyCertificateList) GetId() *int64 {
	return s.Id
}

func (s *ListClientCertificateResponseBodyCertificateList) GetIdentifier() *string {
	return s.Identifier
}

func (s *ListClientCertificateResponseBodyCertificateList) GetKeySize() *int32 {
	return s.KeySize
}

func (s *ListClientCertificateResponseBodyCertificateList) GetLocality() *string {
	return s.Locality
}

func (s *ListClientCertificateResponseBodyCertificateList) GetMd5() *string {
	return s.Md5
}

func (s *ListClientCertificateResponseBodyCertificateList) GetOrganization() *string {
	return s.Organization
}

func (s *ListClientCertificateResponseBodyCertificateList) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *ListClientCertificateResponseBodyCertificateList) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *ListClientCertificateResponseBodyCertificateList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClientCertificateResponseBodyCertificateList) GetSans() *string {
	return s.Sans
}

func (s *ListClientCertificateResponseBodyCertificateList) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *ListClientCertificateResponseBodyCertificateList) GetSha2() *string {
	return s.Sha2
}

func (s *ListClientCertificateResponseBodyCertificateList) GetSignAlgorithm() *string {
	return s.SignAlgorithm
}

func (s *ListClientCertificateResponseBodyCertificateList) GetState() *string {
	return s.State
}

func (s *ListClientCertificateResponseBodyCertificateList) GetStatus() *string {
	return s.Status
}

func (s *ListClientCertificateResponseBodyCertificateList) GetSubjectDN() *string {
	return s.SubjectDN
}

func (s *ListClientCertificateResponseBodyCertificateList) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *ListClientCertificateResponseBodyCertificateList) SetAfterDate(v int64) *ListClientCertificateResponseBodyCertificateList {
	s.AfterDate = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetAlgorithm(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Algorithm = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetAliasName(v string) *ListClientCertificateResponseBodyCertificateList {
	s.AliasName = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetBeforeDate(v int64) *ListClientCertificateResponseBodyCertificateList {
	s.BeforeDate = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetCertificateType(v string) *ListClientCertificateResponseBodyCertificateList {
	s.CertificateType = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetCommonName(v string) *ListClientCertificateResponseBodyCertificateList {
	s.CommonName = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetCountryCode(v string) *ListClientCertificateResponseBodyCertificateList {
	s.CountryCode = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetCustomIdentifier(v string) *ListClientCertificateResponseBodyCertificateList {
	s.CustomIdentifier = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetDays(v int32) *ListClientCertificateResponseBodyCertificateList {
	s.Days = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetId(v int64) *ListClientCertificateResponseBodyCertificateList {
	s.Id = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetIdentifier(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Identifier = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetKeySize(v int32) *ListClientCertificateResponseBodyCertificateList {
	s.KeySize = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetLocality(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Locality = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetMd5(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Md5 = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetOrganization(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Organization = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetOrganizationUnit(v string) *ListClientCertificateResponseBodyCertificateList {
	s.OrganizationUnit = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetParentIdentifier(v string) *ListClientCertificateResponseBodyCertificateList {
	s.ParentIdentifier = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetResourceGroupId(v string) *ListClientCertificateResponseBodyCertificateList {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetSans(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Sans = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetSerialNumber(v string) *ListClientCertificateResponseBodyCertificateList {
	s.SerialNumber = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetSha2(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Sha2 = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetSignAlgorithm(v string) *ListClientCertificateResponseBodyCertificateList {
	s.SignAlgorithm = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetState(v string) *ListClientCertificateResponseBodyCertificateList {
	s.State = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetStatus(v string) *ListClientCertificateResponseBodyCertificateList {
	s.Status = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetSubjectDN(v string) *ListClientCertificateResponseBodyCertificateList {
	s.SubjectDN = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) SetX509Certificate(v string) *ListClientCertificateResponseBodyCertificateList {
	s.X509Certificate = &v
	return s
}

func (s *ListClientCertificateResponseBodyCertificateList) Validate() error {
	return dara.Validate(s)
}
