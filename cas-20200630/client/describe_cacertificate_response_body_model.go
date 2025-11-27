// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCACertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificate(v *DescribeCACertificateResponseBodyCertificate) *DescribeCACertificateResponseBody
	GetCertificate() *DescribeCACertificateResponseBodyCertificate
	SetRequestId(v string) *DescribeCACertificateResponseBody
	GetRequestId() *string
	SetYears(v int32) *DescribeCACertificateResponseBody
	GetYears() *int32
}

type DescribeCACertificateResponseBody struct {
	// The details about the CA certificate.
	Certificate *DescribeCACertificateResponseBodyCertificate `json:"Certificate,omitempty" xml:"Certificate,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The validity period of the CA certificate. Unit: years.
	//
	// example:
	//
	// 10
	Years *int32 `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s DescribeCACertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateResponseBody) GetCertificate() *DescribeCACertificateResponseBodyCertificate {
	return s.Certificate
}

func (s *DescribeCACertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCACertificateResponseBody) GetYears() *int32 {
	return s.Years
}

func (s *DescribeCACertificateResponseBody) SetCertificate(v *DescribeCACertificateResponseBodyCertificate) *DescribeCACertificateResponseBody {
	s.Certificate = v
	return s
}

func (s *DescribeCACertificateResponseBody) SetRequestId(v string) *DescribeCACertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCACertificateResponseBody) SetYears(v int32) *DescribeCACertificateResponseBody {
	s.Years = &v
	return s
}

func (s *DescribeCACertificateResponseBody) Validate() error {
	if s.Certificate != nil {
		if err := s.Certificate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCACertificateResponseBodyCertificate struct {
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
	// The issuance date of the CA certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1634283958000
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// CA certificate chain.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// 用户证书
	//
	// -----END CERTIFICATE-----
	//
	// -----BEGIN CERTIFICATE-----
	//
	// 中间证书
	//
	// -----END CERTIFICATE-----
	//
	// -----BEGIN CERTIFICATE-----
	//
	// 根证书
	//
	// -----END CERTIFICATE-----
	CaCertChain *string `json:"CaCertChain,omitempty" xml:"CaCertChain,omitempty"`
	// The number of certificates issued by private CA instances.
	//
	// example:
	//
	// 10
	CertIssuedCount *int64 `json:"CertIssuedCount,omitempty" xml:"CertIssuedCount,omitempty"`
	// The remaining number of assignable certificate quotas.
	//
	// example:
	//
	// 30
	CertRemainingCount *int64 `json:"CertRemainingCount,omitempty" xml:"CertRemainingCount,omitempty"`
	// The total number of purchased certificate quotas.
	//
	// example:
	//
	// 40
	CertTotalCount *int64 `json:"CertTotalCount,omitempty" xml:"CertTotalCount,omitempty"`
	// The type of the CA certificate. Valid values:
	//
	// 	- **ROOT**: root CA certificate
	//
	// 	- **SUB_ROOT**: intermediate CA certificate
	//
	// example:
	//
	// SUB_ROOT
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	ClusterId       *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
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
	// CRL validity period: 1-365 days.
	//
	// example:
	//
	// 90
	CrlDay *int32 `json:"CrlDay,omitempty" xml:"CrlDay,omitempty"`
	// The status of the certificate revocation list (CRL) feature.
	//
	// example:
	//
	// ACTIVE
	CrlStatus *string `json:"CrlStatus,omitempty" xml:"CrlStatus,omitempty"`
	// The address of the CRL.
	//
	// example:
	//
	// https://crl-cn-publish.oss-cn-hangzhou.aliyuncs.com/pca/crl/1925647866611395/1ed40789-483f-6023-b6b8-29ddd3bb0a9a.crl
	CrlUrl        *string `json:"CrlUrl,omitempty" xml:"CrlUrl,omitempty"`
	FullAlgorithm *string `json:"FullAlgorithm,omitempty" xml:"FullAlgorithm,omitempty"`
	// The unique identifier of the CA certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	IssuerType *string `json:"IssuerType,omitempty" xml:"IssuerType,omitempty"`
	KeyIndex   *int32  `json:"KeyIndex,omitempty" xml:"KeyIndex,omitempty"`
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
	ResourceGroupId  *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
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
	// The user attribute of the CA certificate, which contains the following information:
	//
	// 	- **C**: the country code in which the organization is located
	//
	// 	- **O**: the name of the organization
	//
	// 	- **OU**: the name of the department or branch in the organization
	//
	// 	- **L**: the name of the city in which the organization is located
	//
	// 	- **ST**: the name of the province, municipality, or autonomous region in which the organization is located
	//
	// 	- **CN**: the common name or abbreviation of the organization
	//
	// example:
	//
	// C=CN,O=Alibaba Cloud Computing Co., Ltd.,OU=Security,L=Hangzhou,ST=Zhejiang,CN=Aliyun
	SubjectDN *string                                             `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	Tags      []*DescribeCACertificateResponseBodyCertificateTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The content of the CA certificate.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE----- …… -----END CERTIFICATE-----
	X509Certificate *string `json:"X509Certificate,omitempty" xml:"X509Certificate,omitempty"`
	Years           *int32  `json:"Years,omitempty" xml:"Years,omitempty"`
}

func (s DescribeCACertificateResponseBodyCertificate) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificateResponseBodyCertificate) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateResponseBodyCertificate) GetAfterDate() *int64 {
	return s.AfterDate
}

func (s *DescribeCACertificateResponseBodyCertificate) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DescribeCACertificateResponseBodyCertificate) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *DescribeCACertificateResponseBodyCertificate) GetCaCertChain() *string {
	return s.CaCertChain
}

func (s *DescribeCACertificateResponseBodyCertificate) GetCertIssuedCount() *int64 {
	return s.CertIssuedCount
}

func (s *DescribeCACertificateResponseBodyCertificate) GetCertRemainingCount() *int64 {
	return s.CertRemainingCount
}

func (s *DescribeCACertificateResponseBodyCertificate) GetCertTotalCount() *int64 {
	return s.CertTotalCount
}

func (s *DescribeCACertificateResponseBodyCertificate) GetCertificateType() *string {
	return s.CertificateType
}

func (s *DescribeCACertificateResponseBodyCertificate) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeCACertificateResponseBodyCertificate) GetCommonName() *string {
	return s.CommonName
}

func (s *DescribeCACertificateResponseBodyCertificate) GetCountryCode() *string {
	return s.CountryCode
}

func (s *DescribeCACertificateResponseBodyCertificate) GetCrlDay() *int32 {
	return s.CrlDay
}

func (s *DescribeCACertificateResponseBodyCertificate) GetCrlStatus() *string {
	return s.CrlStatus
}

func (s *DescribeCACertificateResponseBodyCertificate) GetCrlUrl() *string {
	return s.CrlUrl
}

func (s *DescribeCACertificateResponseBodyCertificate) GetFullAlgorithm() *string {
	return s.FullAlgorithm
}

func (s *DescribeCACertificateResponseBodyCertificate) GetIdentifier() *string {
	return s.Identifier
}

func (s *DescribeCACertificateResponseBodyCertificate) GetIssuerType() *string {
	return s.IssuerType
}

func (s *DescribeCACertificateResponseBodyCertificate) GetKeyIndex() *int32 {
	return s.KeyIndex
}

func (s *DescribeCACertificateResponseBodyCertificate) GetKeySize() *int32 {
	return s.KeySize
}

func (s *DescribeCACertificateResponseBodyCertificate) GetLocality() *string {
	return s.Locality
}

func (s *DescribeCACertificateResponseBodyCertificate) GetMd5() *string {
	return s.Md5
}

func (s *DescribeCACertificateResponseBodyCertificate) GetOrganization() *string {
	return s.Organization
}

func (s *DescribeCACertificateResponseBodyCertificate) GetOrganizationUnit() *string {
	return s.OrganizationUnit
}

func (s *DescribeCACertificateResponseBodyCertificate) GetParentIdentifier() *string {
	return s.ParentIdentifier
}

func (s *DescribeCACertificateResponseBodyCertificate) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeCACertificateResponseBodyCertificate) GetSans() *string {
	return s.Sans
}

func (s *DescribeCACertificateResponseBodyCertificate) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeCACertificateResponseBodyCertificate) GetSha2() *string {
	return s.Sha2
}

func (s *DescribeCACertificateResponseBodyCertificate) GetSignAlgorithm() *string {
	return s.SignAlgorithm
}

func (s *DescribeCACertificateResponseBodyCertificate) GetState() *string {
	return s.State
}

func (s *DescribeCACertificateResponseBodyCertificate) GetStatus() *string {
	return s.Status
}

func (s *DescribeCACertificateResponseBodyCertificate) GetSubjectDN() *string {
	return s.SubjectDN
}

func (s *DescribeCACertificateResponseBodyCertificate) GetTags() []*DescribeCACertificateResponseBodyCertificateTags {
	return s.Tags
}

func (s *DescribeCACertificateResponseBodyCertificate) GetX509Certificate() *string {
	return s.X509Certificate
}

func (s *DescribeCACertificateResponseBodyCertificate) GetYears() *int32 {
	return s.Years
}

func (s *DescribeCACertificateResponseBodyCertificate) SetAfterDate(v int64) *DescribeCACertificateResponseBodyCertificate {
	s.AfterDate = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetAlgorithm(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Algorithm = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetBeforeDate(v int64) *DescribeCACertificateResponseBodyCertificate {
	s.BeforeDate = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCaCertChain(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CaCertChain = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCertIssuedCount(v int64) *DescribeCACertificateResponseBodyCertificate {
	s.CertIssuedCount = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCertRemainingCount(v int64) *DescribeCACertificateResponseBodyCertificate {
	s.CertRemainingCount = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCertTotalCount(v int64) *DescribeCACertificateResponseBodyCertificate {
	s.CertTotalCount = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCertificateType(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CertificateType = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetClusterId(v string) *DescribeCACertificateResponseBodyCertificate {
	s.ClusterId = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCommonName(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CommonName = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCountryCode(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CountryCode = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCrlDay(v int32) *DescribeCACertificateResponseBodyCertificate {
	s.CrlDay = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCrlStatus(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CrlStatus = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetCrlUrl(v string) *DescribeCACertificateResponseBodyCertificate {
	s.CrlUrl = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetFullAlgorithm(v string) *DescribeCACertificateResponseBodyCertificate {
	s.FullAlgorithm = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetIdentifier(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Identifier = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetIssuerType(v string) *DescribeCACertificateResponseBodyCertificate {
	s.IssuerType = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetKeyIndex(v int32) *DescribeCACertificateResponseBodyCertificate {
	s.KeyIndex = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetKeySize(v int32) *DescribeCACertificateResponseBodyCertificate {
	s.KeySize = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetLocality(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Locality = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetMd5(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Md5 = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetOrganization(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Organization = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetOrganizationUnit(v string) *DescribeCACertificateResponseBodyCertificate {
	s.OrganizationUnit = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetParentIdentifier(v string) *DescribeCACertificateResponseBodyCertificate {
	s.ParentIdentifier = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetResourceGroupId(v string) *DescribeCACertificateResponseBodyCertificate {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetSans(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Sans = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetSerialNumber(v string) *DescribeCACertificateResponseBodyCertificate {
	s.SerialNumber = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetSha2(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Sha2 = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetSignAlgorithm(v string) *DescribeCACertificateResponseBodyCertificate {
	s.SignAlgorithm = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetState(v string) *DescribeCACertificateResponseBodyCertificate {
	s.State = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetStatus(v string) *DescribeCACertificateResponseBodyCertificate {
	s.Status = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetSubjectDN(v string) *DescribeCACertificateResponseBodyCertificate {
	s.SubjectDN = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetTags(v []*DescribeCACertificateResponseBodyCertificateTags) *DescribeCACertificateResponseBodyCertificate {
	s.Tags = v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetX509Certificate(v string) *DescribeCACertificateResponseBodyCertificate {
	s.X509Certificate = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) SetYears(v int32) *DescribeCACertificateResponseBodyCertificate {
	s.Years = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificate) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCACertificateResponseBodyCertificateTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeCACertificateResponseBodyCertificateTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeCACertificateResponseBodyCertificateTags) GoString() string {
	return s.String()
}

func (s *DescribeCACertificateResponseBodyCertificateTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeCACertificateResponseBodyCertificateTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeCACertificateResponseBodyCertificateTags) SetTagKey(v string) *DescribeCACertificateResponseBodyCertificateTags {
	s.TagKey = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificateTags) SetTagValue(v string) *DescribeCACertificateResponseBodyCertificateTags {
	s.TagValue = &v
	return s
}

func (s *DescribeCACertificateResponseBodyCertificateTags) Validate() error {
	return dara.Validate(s)
}
