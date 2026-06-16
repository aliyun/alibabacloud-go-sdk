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
	// The details of the client certificate or server-side certificate.
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
	if s.Certificate != nil {
		if err := s.Certificate.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClientCertificateResponseBodyCertificate struct {
	// The expiration date of the certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1665819958000
	AfterDate *int64 `json:"AfterDate,omitempty" xml:"AfterDate,omitempty"`
	// The type of the encryption algorithm. Valid values:
	//
	// - **RSA**: the RSA algorithm.
	//
	// - **ECC**: the ECC algorithm.
	//
	// - **SM2**: the SM2 algorithm.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The alias of the issued certificate.
	//
	// example:
	//
	// rsa_root_2048
	AliasName *string `json:"AliasName,omitempty" xml:"AliasName,omitempty"`
	// The issuance date of the certificate. This value is a UNIX timestamp. Unit: milliseconds.
	//
	// example:
	//
	// 1634283958000
	BeforeDate *int64 `json:"BeforeDate,omitempty" xml:"BeforeDate,omitempty"`
	// The complete certificate chain.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// cert
	//
	// -----END CERTIFICATE-----
	//
	// -----BEGIN CERTIFICATE-----
	//
	// subCA
	//
	// -----END CERTIFICATE-----
	//
	// -----BEGIN CERTIFICATE-----
	//
	// rootCA
	//
	// -----END CERTIFICATE-----
	CertChain *string `json:"CertChain,omitempty" xml:"CertChain,omitempty"`
	// The type of the certificate. Valid values:
	//
	// - **CLIENT**: a client certificate.
	//
	// - **SERVER**: a server-side certificate.
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
	// The country code of the subject organization.
	//
	// For more information about country codes, see the **International codes*	- section in [Manage company profiles](https://help.aliyun.com/document_detail/198289.html).
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The custom identifier, which is a unique key.
	//
	// example:
	//
	// ***3a32d96883a6650e672ea0276****
	CustomIdentifier *string `json:"CustomIdentifier,omitempty" xml:"CustomIdentifier,omitempty"`
	// The validity period of the certificate. Unit: days.
	//
	// example:
	//
	// 365
	Days *int32 `json:"Days,omitempty" xml:"Days,omitempty"`
	// The algorithm and its length.
	//
	// example:
	//
	// RSA_2048
	FullAlgorithm *string `json:"FullAlgorithm,omitempty" xml:"FullAlgorithm,omitempty"`
	// The ID of the data source to which the certificate order belongs.
	//
	// example:
	//
	// 1137354
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
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
	// The city where the subject organization is located.
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
	// The organization associated with the certificate of the issuing subordinate CA.
	//
	// example:
	//
	// Aliyun
	Organization *string `json:"Organization,omitempty" xml:"Organization,omitempty"`
	// The organizational unit of the certificate subject.
	//
	// example:
	//
	// Security
	OrganizationUnit *string `json:"OrganizationUnit,omitempty" xml:"OrganizationUnit,omitempty"`
	// The unique identifier of the subordinate CA certificate that issued the certificate.
	//
	// example:
	//
	// 160ae6bb538d538c70c01f81dcf2****
	ParentIdentifier *string `json:"ParentIdentifier,omitempty" xml:"ParentIdentifier,omitempty"`
	// The ID of the resource group to which the certificate belongs.
	//
	// example:
	//
	// rg-acfmxllajdpw3fi
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The Subject Alternative Name (SAN) extension of the certificate. The SAN extension indicates other domain names or IP addresses that are associated with the certificate.
	//
	// This parameter is a string that is converted from a JSON array. Each element in the JSON array is a struct that corresponds to a SAN extension. Each SAN extension struct contains the following parameters:
	//
	// - **Type**: The type of the extension. This parameter is of the Integer type. Valid values:
	//
	//   - **1**: an email address.
	//
	//   - **2**: a domain name.
	//
	//   - **6**: a Uniform Resource Identifier (URI).
	//
	//   - **7**: an IP address.
	//
	// - **Value**: The content of the extension. This parameter is of the String type.
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
	// The state or province where the subject organization is located.
	//
	// example:
	//
	// Zhejiang
	State *string `json:"State,omitempty" xml:"State,omitempty"`
	// The status of the certificate. Valid values:
	//
	// - **ISSUE**: The certificate is issued.
	//
	// - **REVOKE**: The certificate is revoked.
	//
	// example:
	//
	// ISSUE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The subject Distinguished Name (DN) of the certificate. This value is composed of the following fields:
	//
	// - **C**: Country.
	//
	// - **O**: Organization.
	//
	// - **OU**: Organizational unit.
	//
	// - **CN**: Common name.
	//
	// <props="china">
	//
	// - **ST**: The province, municipality, or autonomous region.
	//
	//
	//
	//
	// <props="intl">
	//
	// - **ST**: Province or state.
	//
	//
	//
	//
	// - **CN**: Common name.
	//
	// example:
	//
	// C=CN,O=Aliyun,OU=Security,L=Hangzhou,ST=Zhejiang,CN=Aliyun
	SubjectDN *string `json:"SubjectDN,omitempty" xml:"SubjectDN,omitempty"`
	// The list of tags.
	//
	// example:
	//
	// mtls
	Tags []*DescribeClientCertificateResponseBodyCertificateTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Indicates whether the certificate is synchronized to Digital Certificate Management Service.
	//
	// example:
	//
	// 1
	UploadFlag *int32 `json:"UploadFlag,omitempty" xml:"UploadFlag,omitempty"`
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

func (s *DescribeClientCertificateResponseBodyCertificate) GetAliasName() *string {
	return s.AliasName
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetBeforeDate() *int64 {
	return s.BeforeDate
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetCertChain() *string {
	return s.CertChain
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

func (s *DescribeClientCertificateResponseBodyCertificate) GetCustomIdentifier() *string {
	return s.CustomIdentifier
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetDays() *int32 {
	return s.Days
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetFullAlgorithm() *string {
	return s.FullAlgorithm
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetId() *int64 {
	return s.Id
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

func (s *DescribeClientCertificateResponseBodyCertificate) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *DescribeClientCertificateResponseBodyCertificate) GetTags() []*DescribeClientCertificateResponseBodyCertificateTags {
	return s.Tags
}

func (s *DescribeClientCertificateResponseBodyCertificate) GetUploadFlag() *int32 {
	return s.UploadFlag
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

func (s *DescribeClientCertificateResponseBodyCertificate) SetAliasName(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.AliasName = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetBeforeDate(v int64) *DescribeClientCertificateResponseBodyCertificate {
	s.BeforeDate = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetCertChain(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.CertChain = &v
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

func (s *DescribeClientCertificateResponseBodyCertificate) SetCustomIdentifier(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.CustomIdentifier = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetDays(v int32) *DescribeClientCertificateResponseBodyCertificate {
	s.Days = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetFullAlgorithm(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.FullAlgorithm = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetId(v int64) *DescribeClientCertificateResponseBodyCertificate {
	s.Id = &v
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

func (s *DescribeClientCertificateResponseBodyCertificate) SetResourceGroupId(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.ResourceGroupId = &v
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

func (s *DescribeClientCertificateResponseBodyCertificate) SetTags(v []*DescribeClientCertificateResponseBodyCertificateTags) *DescribeClientCertificateResponseBodyCertificate {
	s.Tags = v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetUploadFlag(v int32) *DescribeClientCertificateResponseBodyCertificate {
	s.UploadFlag = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) SetX509Certificate(v string) *DescribeClientCertificateResponseBodyCertificate {
	s.X509Certificate = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificate) Validate() error {
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

type DescribeClientCertificateResponseBodyCertificateTags struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// [{\\"tag\\":\\"PROPERTY_TYPE\\",\\"values\\":[]}]
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeClientCertificateResponseBodyCertificateTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeClientCertificateResponseBodyCertificateTags) GoString() string {
	return s.String()
}

func (s *DescribeClientCertificateResponseBodyCertificateTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeClientCertificateResponseBodyCertificateTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeClientCertificateResponseBodyCertificateTags) SetTagKey(v string) *DescribeClientCertificateResponseBodyCertificateTags {
	s.TagKey = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificateTags) SetTagValue(v string) *DescribeClientCertificateResponseBodyCertificateTags {
	s.TagValue = &v
	return s
}

func (s *DescribeClientCertificateResponseBodyCertificateTags) Validate() error {
	return dara.Validate(s)
}
