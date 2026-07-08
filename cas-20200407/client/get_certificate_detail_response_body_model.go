// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *GetCertificateDetailResponseBody
	GetAlgorithm() *string
	SetCertIdentifier(v string) *GetCertificateDetailResponseBody
	GetCertIdentifier() *string
	SetCertificateChainList(v []*GetCertificateDetailResponseBodyCertificateChainList) *GetCertificateDetailResponseBody
	GetCertificateChainList() []*GetCertificateDetailResponseBodyCertificateChainList
	SetCertificateId(v int32) *GetCertificateDetailResponseBody
	GetCertificateId() *int32
	SetCertificateName(v string) *GetCertificateDetailResponseBody
	GetCertificateName() *string
	SetCertificateSource(v string) *GetCertificateDetailResponseBody
	GetCertificateSource() *string
	SetCertificateStatus(v string) *GetCertificateDetailResponseBody
	GetCertificateStatus() *string
	SetCommonName(v string) *GetCertificateDetailResponseBody
	GetCommonName() *string
	SetCompanyId(v int64) *GetCertificateDetailResponseBody
	GetCompanyId() *int64
	SetContactId(v int64) *GetCertificateDetailResponseBody
	GetContactId() *int64
	SetCsr(v string) *GetCertificateDetailResponseBody
	GetCsr() *string
	SetDomain(v string) *GetCertificateDetailResponseBody
	GetDomain() *string
	SetExistPrivateKey(v bool) *GetCertificateDetailResponseBody
	GetExistPrivateKey() *bool
	SetFingerPrint(v string) *GetCertificateDetailResponseBody
	GetFingerPrint() *string
	SetInstanceId(v string) *GetCertificateDetailResponseBody
	GetInstanceId() *string
	SetIssuer(v string) *GetCertificateDetailResponseBody
	GetIssuer() *string
	SetKeySize(v int32) *GetCertificateDetailResponseBody
	GetKeySize() *int32
	SetNotAfter(v int64) *GetCertificateDetailResponseBody
	GetNotAfter() *int64
	SetNotBefore(v int64) *GetCertificateDetailResponseBody
	GetNotBefore() *int64
	SetRequestId(v string) *GetCertificateDetailResponseBody
	GetRequestId() *string
	SetSerial(v string) *GetCertificateDetailResponseBody
	GetSerial() *string
	SetSubjectAlternativeNames(v []*string) *GetCertificateDetailResponseBody
	GetSubjectAlternativeNames() []*string
	SetTags(v []*GetCertificateDetailResponseBodyTags) *GetCertificateDetailResponseBody
	GetTags() []*GetCertificateDetailResponseBodyTags
	SetUsingProductList(v []*string) *GetCertificateDetailResponseBody
	GetUsingProductList() []*string
}

type GetCertificateDetailResponseBody struct {
	// The certificate algorithm. Valid values:
	//
	// - **RSA**: The RSA algorithm.
	//
	// - **ECC**: The ECC algorithm.
	//
	// - **SM2**: The SM2 algorithm.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The global ID of the certificate, which is used in various Alibaba Cloud services. The format of the ID is `<Certificate ID>-<Region ID>`. The region ID is `cn-hangzhou` for the China site and `ap-southeast-1` for the International site. For example, if a certificate ID is `123`, its `CertIdentifier` is `123-cn-hangzhou` for the China site and `123-ap-southeast-1` for the International site.
	//
	// example:
	//
	// 21912069-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The information about the certificate chain.
	CertificateChainList []*GetCertificateDetailResponseBodyCertificateChainList `json:"CertificateChainList,omitempty" xml:"CertificateChainList,omitempty" type:"Repeated"`
	// The certificate ID.
	//
	// example:
	//
	// 22559621
	CertificateId *int32 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// 123
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// The source of the certificate. Valid values:
	//
	// - **BUY**: a purchased certificate.
	//
	// - **TEST**: a test certificate.
	//
	// - Upload the certificate.
	//
	// example:
	//
	// BUY
	CertificateSource *string `json:"CertificateSource,omitempty" xml:"CertificateSource,omitempty"`
	// The status of the certificate. Valid values:
	//
	// - **issued**: The certificate is issued.
	//
	// - **revoked**: The certificate is revoked.
	//
	// - **willExpire**: The certificate is about to expire.
	//
	// - **expired**: The certificate has expired.
	//
	// example:
	//
	// issued
	CertificateStatus *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	// The common name.
	//
	// example:
	//
	// www.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The ID of the company profile that is associated with the certificate application. This parameter is empty for DV certificates.
	//
	// example:
	//
	// 44211
	CompanyId *int64 `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	// The ID of the contact.
	//
	// example:
	//
	// 304066
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST----- ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The domain names that are bound to the certificate. Multiple domain names are separated by commas (,).
	//
	// example:
	//
	// aliyundoc.com,example.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// Indicates whether a private key is available. Valid values:
	//
	// - **true**
	//
	// - **false**
	//
	// example:
	//
	// true
	ExistPrivateKey *bool `json:"ExistPrivateKey,omitempty" xml:"ExistPrivateKey,omitempty"`
	// The fingerprint of the public key.
	//
	// example:
	//
	// 123
	FingerPrint *string `json:"FingerPrint,omitempty" xml:"FingerPrint,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cas_dv-cn-123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The issuer of the certificate.
	//
	// example:
	//
	// Digicert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The key size.
	//
	// - For RSA algorithms, the key size is typically 2,048, 3,072, or 4,096 bits.
	//
	// - For ECC and SM2 algorithms, the key size is typically 256 bits.
	//
	// example:
	//
	// 2048
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// The end of the validity period of the certificate.
	//
	// example:
	//
	// 17326613180000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The beginning of the validity period of the certificate.
	//
	// example:
	//
	// 17321613180000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The serial number of the certificate.
	//
	// example:
	//
	// 123
	Serial *string `json:"Serial,omitempty" xml:"Serial,omitempty"`
	// The subject alternative names (SANs) of the certificate.
	SubjectAlternativeNames []*string `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Repeated"`
	// The list of tags.
	Tags []*GetCertificateDetailResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The list of cloud services in which the certificate is deployed.
	UsingProductList []*string `json:"UsingProductList,omitempty" xml:"UsingProductList,omitempty" type:"Repeated"`
}

func (s GetCertificateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertificateDetailResponseBody) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *GetCertificateDetailResponseBody) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *GetCertificateDetailResponseBody) GetCertificateChainList() []*GetCertificateDetailResponseBodyCertificateChainList {
	return s.CertificateChainList
}

func (s *GetCertificateDetailResponseBody) GetCertificateId() *int32 {
	return s.CertificateId
}

func (s *GetCertificateDetailResponseBody) GetCertificateName() *string {
	return s.CertificateName
}

func (s *GetCertificateDetailResponseBody) GetCertificateSource() *string {
	return s.CertificateSource
}

func (s *GetCertificateDetailResponseBody) GetCertificateStatus() *string {
	return s.CertificateStatus
}

func (s *GetCertificateDetailResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *GetCertificateDetailResponseBody) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *GetCertificateDetailResponseBody) GetContactId() *int64 {
	return s.ContactId
}

func (s *GetCertificateDetailResponseBody) GetCsr() *string {
	return s.Csr
}

func (s *GetCertificateDetailResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *GetCertificateDetailResponseBody) GetExistPrivateKey() *bool {
	return s.ExistPrivateKey
}

func (s *GetCertificateDetailResponseBody) GetFingerPrint() *string {
	return s.FingerPrint
}

func (s *GetCertificateDetailResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCertificateDetailResponseBody) GetIssuer() *string {
	return s.Issuer
}

func (s *GetCertificateDetailResponseBody) GetKeySize() *int32 {
	return s.KeySize
}

func (s *GetCertificateDetailResponseBody) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *GetCertificateDetailResponseBody) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *GetCertificateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCertificateDetailResponseBody) GetSerial() *string {
	return s.Serial
}

func (s *GetCertificateDetailResponseBody) GetSubjectAlternativeNames() []*string {
	return s.SubjectAlternativeNames
}

func (s *GetCertificateDetailResponseBody) GetTags() []*GetCertificateDetailResponseBodyTags {
	return s.Tags
}

func (s *GetCertificateDetailResponseBody) GetUsingProductList() []*string {
	return s.UsingProductList
}

func (s *GetCertificateDetailResponseBody) SetAlgorithm(v string) *GetCertificateDetailResponseBody {
	s.Algorithm = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertIdentifier(v string) *GetCertificateDetailResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertificateChainList(v []*GetCertificateDetailResponseBodyCertificateChainList) *GetCertificateDetailResponseBody {
	s.CertificateChainList = v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertificateId(v int32) *GetCertificateDetailResponseBody {
	s.CertificateId = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertificateName(v string) *GetCertificateDetailResponseBody {
	s.CertificateName = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertificateSource(v string) *GetCertificateDetailResponseBody {
	s.CertificateSource = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertificateStatus(v string) *GetCertificateDetailResponseBody {
	s.CertificateStatus = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCommonName(v string) *GetCertificateDetailResponseBody {
	s.CommonName = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCompanyId(v int64) *GetCertificateDetailResponseBody {
	s.CompanyId = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetContactId(v int64) *GetCertificateDetailResponseBody {
	s.ContactId = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCsr(v string) *GetCertificateDetailResponseBody {
	s.Csr = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetDomain(v string) *GetCertificateDetailResponseBody {
	s.Domain = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetExistPrivateKey(v bool) *GetCertificateDetailResponseBody {
	s.ExistPrivateKey = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetFingerPrint(v string) *GetCertificateDetailResponseBody {
	s.FingerPrint = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetInstanceId(v string) *GetCertificateDetailResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetIssuer(v string) *GetCertificateDetailResponseBody {
	s.Issuer = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetKeySize(v int32) *GetCertificateDetailResponseBody {
	s.KeySize = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetNotAfter(v int64) *GetCertificateDetailResponseBody {
	s.NotAfter = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetNotBefore(v int64) *GetCertificateDetailResponseBody {
	s.NotBefore = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetRequestId(v string) *GetCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetSerial(v string) *GetCertificateDetailResponseBody {
	s.Serial = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetSubjectAlternativeNames(v []*string) *GetCertificateDetailResponseBody {
	s.SubjectAlternativeNames = v
	return s
}

func (s *GetCertificateDetailResponseBody) SetTags(v []*GetCertificateDetailResponseBodyTags) *GetCertificateDetailResponseBody {
	s.Tags = v
	return s
}

func (s *GetCertificateDetailResponseBody) SetUsingProductList(v []*string) *GetCertificateDetailResponseBody {
	s.UsingProductList = v
	return s
}

func (s *GetCertificateDetailResponseBody) Validate() error {
	if s.CertificateChainList != nil {
		for _, item := range s.CertificateChainList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type GetCertificateDetailResponseBodyCertificateChainList struct {
	// The issuer of the certificate chain.
	//
	// example:
	//
	// Digicert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The end of the validity period.
	//
	// example:
	//
	// 17326613180000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The beginning of the validity period.
	//
	// example:
	//
	// 17321613180000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The remaining validity period of the certificate chain.
	//
	// example:
	//
	// 10
	RemainDay *int32 `json:"RemainDay,omitempty" xml:"RemainDay,omitempty"`
	// The common name of the certificate chain.
	//
	// example:
	//
	// Digicert
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s GetCertificateDetailResponseBodyCertificateChainList) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateDetailResponseBodyCertificateChainList) GoString() string {
	return s.String()
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) GetIssuer() *string {
	return s.Issuer
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) GetRemainDay() *int32 {
	return s.RemainDay
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) GetSubject() *string {
	return s.Subject
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) SetIssuer(v string) *GetCertificateDetailResponseBodyCertificateChainList {
	s.Issuer = &v
	return s
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) SetNotAfter(v int64) *GetCertificateDetailResponseBodyCertificateChainList {
	s.NotAfter = &v
	return s
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) SetNotBefore(v int64) *GetCertificateDetailResponseBodyCertificateChainList {
	s.NotBefore = &v
	return s
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) SetRemainDay(v int32) *GetCertificateDetailResponseBodyCertificateChainList {
	s.RemainDay = &v
	return s
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) SetSubject(v string) *GetCertificateDetailResponseBodyCertificateChainList {
	s.Subject = &v
	return s
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) Validate() error {
	return dara.Validate(s)
}

type GetCertificateDetailResponseBodyTags struct {
	// The tag key of the instance. You can specify 1 to 20 tag keys. The value cannot be an empty string.
	//
	// The value can be up to 64 characters in length, cannot start with `aliyun` or `acs:`, and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetCertificateDetailResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateDetailResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetCertificateDetailResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetCertificateDetailResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetCertificateDetailResponseBodyTags) SetTagKey(v string) *GetCertificateDetailResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetCertificateDetailResponseBodyTags) SetTagValue(v string) *GetCertificateDetailResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetCertificateDetailResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
