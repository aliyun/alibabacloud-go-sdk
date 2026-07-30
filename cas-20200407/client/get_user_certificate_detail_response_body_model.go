// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserCertificateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *GetUserCertificateDetailResponseBody
	GetAlgorithm() *string
	SetBuyInAliyun(v bool) *GetUserCertificateDetailResponseBody
	GetBuyInAliyun() *bool
	SetCert(v string) *GetUserCertificateDetailResponseBody
	GetCert() *string
	SetCertChain(v []*GetUserCertificateDetailResponseBodyCertChain) *GetUserCertificateDetailResponseBody
	GetCertChain() []*GetUserCertificateDetailResponseBodyCertChain
	SetCertIdentifier(v string) *GetUserCertificateDetailResponseBody
	GetCertIdentifier() *string
	SetCertSha2(v string) *GetUserCertificateDetailResponseBody
	GetCertSha2() *string
	SetCity(v string) *GetUserCertificateDetailResponseBody
	GetCity() *string
	SetCommon(v string) *GetUserCertificateDetailResponseBody
	GetCommon() *string
	SetCountry(v string) *GetUserCertificateDetailResponseBody
	GetCountry() *string
	SetEncryptCert(v string) *GetUserCertificateDetailResponseBody
	GetEncryptCert() *string
	SetEncryptPrivateKey(v string) *GetUserCertificateDetailResponseBody
	GetEncryptPrivateKey() *string
	SetEndDate(v string) *GetUserCertificateDetailResponseBody
	GetEndDate() *string
	SetExpired(v bool) *GetUserCertificateDetailResponseBody
	GetExpired() *bool
	SetFingerprint(v string) *GetUserCertificateDetailResponseBody
	GetFingerprint() *string
	SetId(v int64) *GetUserCertificateDetailResponseBody
	GetId() *int64
	SetInstanceId(v string) *GetUserCertificateDetailResponseBody
	GetInstanceId() *string
	SetIssuer(v string) *GetUserCertificateDetailResponseBody
	GetIssuer() *string
	SetKey(v string) *GetUserCertificateDetailResponseBody
	GetKey() *string
	SetName(v string) *GetUserCertificateDetailResponseBody
	GetName() *string
	SetNotAfter(v int64) *GetUserCertificateDetailResponseBody
	GetNotAfter() *int64
	SetNotBefore(v int64) *GetUserCertificateDetailResponseBody
	GetNotBefore() *int64
	SetOrderId(v int64) *GetUserCertificateDetailResponseBody
	GetOrderId() *int64
	SetOrgName(v string) *GetUserCertificateDetailResponseBody
	GetOrgName() *string
	SetProvince(v string) *GetUserCertificateDetailResponseBody
	GetProvince() *string
	SetRequestId(v string) *GetUserCertificateDetailResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetUserCertificateDetailResponseBody
	GetResourceGroupId() *string
	SetSans(v string) *GetUserCertificateDetailResponseBody
	GetSans() *string
	SetSerialNo(v string) *GetUserCertificateDetailResponseBody
	GetSerialNo() *string
	SetSha2(v string) *GetUserCertificateDetailResponseBody
	GetSha2() *string
	SetSignCert(v string) *GetUserCertificateDetailResponseBody
	GetSignCert() *string
	SetSignPrivateKey(v string) *GetUserCertificateDetailResponseBody
	GetSignPrivateKey() *string
	SetStartDate(v string) *GetUserCertificateDetailResponseBody
	GetStartDate() *string
	SetTags(v []*GetUserCertificateDetailResponseBodyTags) *GetUserCertificateDetailResponseBody
	GetTags() []*GetUserCertificateDetailResponseBodyTags
}

type GetUserCertificateDetailResponseBody struct {
	// The algorithm.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// Indicates whether the certificate was purchased from Alibaba Cloud. Valid values:
	//
	// - **true**: Yes.
	//
	// - **false**: No.
	//
	// example:
	//
	// true
	BuyInAliyun *bool `json:"BuyInAliyun,omitempty" xml:"BuyInAliyun,omitempty"`
	// The certificate content that uses a non-Chinese cryptographic algorithm. This field is returned when the CertFilter parameter is set to false, and is not returned when set to true.
	//
	// example:
	//
	// ---BEGIN CERTIFICATE----- MIIF...... -----END CERTIFICATE-----
	Cert *string `json:"Cert,omitempty" xml:"Cert,omitempty"`
	// The certificate chain information.
	CertChain []*GetUserCertificateDetailResponseBodyCertChain `json:"CertChain,omitempty" xml:"CertChain,omitempty" type:"Repeated"`
	// The certificate ID appended with "-cn-hangzhou". For example, if the certificate ID is 123, the CertIdentifier is "123-cn-hangzhou".
	//
	// example:
	//
	// 13781326-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The SHA-2 algorithm value of the certificate.
	//
	// example:
	//
	// E696C2DB574DF99467A90BEF62B341294A9B0C731D228A0BD38265E1467D8EF0
	CertSha2 *string `json:"CertSha2,omitempty" xml:"CertSha2,omitempty"`
	// The city where the company or organization of the certificate purchaser is located.
	//
	// example:
	//
	// hangzhou
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The primary domain name bound to the certificate.
	//
	// example:
	//
	// *.com
	Common *string `json:"Common,omitempty" xml:"Common,omitempty"`
	// The country where the company or organization of the certificate purchaser is located.
	//
	// example:
	//
	// CN
	Country *string `json:"Country,omitempty" xml:"Country,omitempty"`
	// The encryption certificate content in PEM format that uses a Chinese cryptographic algorithm. This field is returned when the CertFilter parameter is set to false, and is not returned when set to true.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICDzCCA***
	//
	// -----END CERTIFICATE-----
	EncryptCert *string `json:"EncryptCert,omitempty" xml:"EncryptCert,omitempty"`
	// The private key of the encryption certificate in PEM format that uses a Chinese cryptographic algorithm. This field is returned when the CertFilter parameter is set to false, and is not returned when set to true.
	//
	// example:
	//
	// -----BEGIN EC PRIVATE KEY-----
	//
	// MHcCAQEEI****
	//
	// -----END EC PRIVATE KEY-----
	EncryptPrivateKey *string `json:"EncryptPrivateKey,omitempty" xml:"EncryptPrivateKey,omitempty"`
	// The certificate expiration date in the format of YYYY-MM-DD.
	//
	// example:
	//
	// 2023-10-25
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Indicates whether the certificate has expired. Valid values:
	//
	// - **true**: The certificate has expired.
	//
	// - **false**: The certificate has not expired.
	//
	// example:
	//
	// true
	Expired *bool `json:"Expired,omitempty" xml:"Expired,omitempty"`
	// The certificate fingerprint.
	//
	// example:
	//
	// 1D7801BBE772D5DE55CBF1F88AEB41A42402DA07
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 121345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The resource instance ID.
	//
	// example:
	//
	// cas-ivauto-hqito6
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The certification authority.
	//
	// example:
	//
	// Digicert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The private key that uses a non-Chinese cryptographic algorithm. This field is returned when the CertFilter parameter is set to false, and is not returned when set to true.
	//
	// example:
	//
	// -----BEGIN RSA PRIVATE KEY----- MII.... -----END RSA PRIVATE KEY-----
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// cert_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The end time of the certificate validity period.
	//
	// example:
	//
	// 17326613180000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The start time of the certificate validity period.
	//
	// example:
	//
	// 17321613180000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The ID of the certificate application order.
	//
	// example:
	//
	// 123456
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The name of the company or organization to which the certificate purchaser belongs.
	//
	// example:
	//
	// Alibaba
	OrgName *string `json:"OrgName,omitempty" xml:"OrgName,omitempty"`
	// The province where the company or organization of the certificate purchaser is located.
	//
	// example:
	//
	// zhejiang
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The request ID. Alibaba Cloud generates a unique identifier for each request, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek****wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// All domain names bound to the certificate.
	//
	// example:
	//
	// *.com
	Sans *string `json:"Sans,omitempty" xml:"Sans,omitempty"`
	// The certificate serial number.
	//
	// example:
	//
	// 033cd852608689ef5e368fde89e0961769e8
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	// The SHA-2 value of the certificate.
	//
	// example:
	//
	// 573415B23243066AD345AE5A57BD0FAE94F598BDD06D906278B5FF318F090FC8
	Sha2 *string `json:"Sha2,omitempty" xml:"Sha2,omitempty"`
	// The signing certificate content in PEM format that uses a Chinese cryptographic algorithm. This field is returned when the CertFilter parameter is set to false, and is not returned when set to true.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE-----
	//
	// MIICDzCCAbagAw****
	//
	// -----END CERTIFICATE-----
	SignCert *string `json:"SignCert,omitempty" xml:"SignCert,omitempty"`
	// The private key of the signing certificate in PEM format that uses a Chinese cryptographic algorithm. This field is returned when the CertFilter parameter is set to false, and is not returned when set to true.
	//
	// example:
	//
	// -----BEGIN EC PRIVATE KEY-----
	//
	// MHcCAQEEILR****
	//
	// -----END EC PRIVATE KEY-----
	SignPrivateKey *string `json:"SignPrivateKey,omitempty" xml:"SignPrivateKey,omitempty"`
	// The certificate issuance date in the format of YYYY-MM-DD.
	//
	// example:
	//
	// 2018-07-13
	StartDate *string `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// The tag list.
	Tags []*GetUserCertificateDetailResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s GetUserCertificateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserCertificateDetailResponseBody) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *GetUserCertificateDetailResponseBody) GetBuyInAliyun() *bool {
	return s.BuyInAliyun
}

func (s *GetUserCertificateDetailResponseBody) GetCert() *string {
	return s.Cert
}

func (s *GetUserCertificateDetailResponseBody) GetCertChain() []*GetUserCertificateDetailResponseBodyCertChain {
	return s.CertChain
}

func (s *GetUserCertificateDetailResponseBody) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *GetUserCertificateDetailResponseBody) GetCertSha2() *string {
	return s.CertSha2
}

func (s *GetUserCertificateDetailResponseBody) GetCity() *string {
	return s.City
}

func (s *GetUserCertificateDetailResponseBody) GetCommon() *string {
	return s.Common
}

func (s *GetUserCertificateDetailResponseBody) GetCountry() *string {
	return s.Country
}

func (s *GetUserCertificateDetailResponseBody) GetEncryptCert() *string {
	return s.EncryptCert
}

func (s *GetUserCertificateDetailResponseBody) GetEncryptPrivateKey() *string {
	return s.EncryptPrivateKey
}

func (s *GetUserCertificateDetailResponseBody) GetEndDate() *string {
	return s.EndDate
}

func (s *GetUserCertificateDetailResponseBody) GetExpired() *bool {
	return s.Expired
}

func (s *GetUserCertificateDetailResponseBody) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *GetUserCertificateDetailResponseBody) GetId() *int64 {
	return s.Id
}

func (s *GetUserCertificateDetailResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserCertificateDetailResponseBody) GetIssuer() *string {
	return s.Issuer
}

func (s *GetUserCertificateDetailResponseBody) GetKey() *string {
	return s.Key
}

func (s *GetUserCertificateDetailResponseBody) GetName() *string {
	return s.Name
}

func (s *GetUserCertificateDetailResponseBody) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *GetUserCertificateDetailResponseBody) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *GetUserCertificateDetailResponseBody) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetUserCertificateDetailResponseBody) GetOrgName() *string {
	return s.OrgName
}

func (s *GetUserCertificateDetailResponseBody) GetProvince() *string {
	return s.Province
}

func (s *GetUserCertificateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserCertificateDetailResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetUserCertificateDetailResponseBody) GetSans() *string {
	return s.Sans
}

func (s *GetUserCertificateDetailResponseBody) GetSerialNo() *string {
	return s.SerialNo
}

func (s *GetUserCertificateDetailResponseBody) GetSha2() *string {
	return s.Sha2
}

func (s *GetUserCertificateDetailResponseBody) GetSignCert() *string {
	return s.SignCert
}

func (s *GetUserCertificateDetailResponseBody) GetSignPrivateKey() *string {
	return s.SignPrivateKey
}

func (s *GetUserCertificateDetailResponseBody) GetStartDate() *string {
	return s.StartDate
}

func (s *GetUserCertificateDetailResponseBody) GetTags() []*GetUserCertificateDetailResponseBodyTags {
	return s.Tags
}

func (s *GetUserCertificateDetailResponseBody) SetAlgorithm(v string) *GetUserCertificateDetailResponseBody {
	s.Algorithm = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetBuyInAliyun(v bool) *GetUserCertificateDetailResponseBody {
	s.BuyInAliyun = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetCert(v string) *GetUserCertificateDetailResponseBody {
	s.Cert = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetCertChain(v []*GetUserCertificateDetailResponseBodyCertChain) *GetUserCertificateDetailResponseBody {
	s.CertChain = v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetCertIdentifier(v string) *GetUserCertificateDetailResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetCertSha2(v string) *GetUserCertificateDetailResponseBody {
	s.CertSha2 = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetCity(v string) *GetUserCertificateDetailResponseBody {
	s.City = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetCommon(v string) *GetUserCertificateDetailResponseBody {
	s.Common = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetCountry(v string) *GetUserCertificateDetailResponseBody {
	s.Country = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetEncryptCert(v string) *GetUserCertificateDetailResponseBody {
	s.EncryptCert = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetEncryptPrivateKey(v string) *GetUserCertificateDetailResponseBody {
	s.EncryptPrivateKey = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetEndDate(v string) *GetUserCertificateDetailResponseBody {
	s.EndDate = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetExpired(v bool) *GetUserCertificateDetailResponseBody {
	s.Expired = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetFingerprint(v string) *GetUserCertificateDetailResponseBody {
	s.Fingerprint = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetId(v int64) *GetUserCertificateDetailResponseBody {
	s.Id = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetInstanceId(v string) *GetUserCertificateDetailResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetIssuer(v string) *GetUserCertificateDetailResponseBody {
	s.Issuer = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetKey(v string) *GetUserCertificateDetailResponseBody {
	s.Key = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetName(v string) *GetUserCertificateDetailResponseBody {
	s.Name = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetNotAfter(v int64) *GetUserCertificateDetailResponseBody {
	s.NotAfter = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetNotBefore(v int64) *GetUserCertificateDetailResponseBody {
	s.NotBefore = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetOrderId(v int64) *GetUserCertificateDetailResponseBody {
	s.OrderId = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetOrgName(v string) *GetUserCertificateDetailResponseBody {
	s.OrgName = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetProvince(v string) *GetUserCertificateDetailResponseBody {
	s.Province = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetRequestId(v string) *GetUserCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetResourceGroupId(v string) *GetUserCertificateDetailResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetSans(v string) *GetUserCertificateDetailResponseBody {
	s.Sans = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetSerialNo(v string) *GetUserCertificateDetailResponseBody {
	s.SerialNo = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetSha2(v string) *GetUserCertificateDetailResponseBody {
	s.Sha2 = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetSignCert(v string) *GetUserCertificateDetailResponseBody {
	s.SignCert = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetSignPrivateKey(v string) *GetUserCertificateDetailResponseBody {
	s.SignPrivateKey = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetStartDate(v string) *GetUserCertificateDetailResponseBody {
	s.StartDate = &v
	return s
}

func (s *GetUserCertificateDetailResponseBody) SetTags(v []*GetUserCertificateDetailResponseBodyTags) *GetUserCertificateDetailResponseBody {
	s.Tags = v
	return s
}

func (s *GetUserCertificateDetailResponseBody) Validate() error {
	if s.CertChain != nil {
		for _, item := range s.CertChain {
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

type GetUserCertificateDetailResponseBodyCertChain struct {
	// The common name in the certificate chain.
	//
	// example:
	//
	// test
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// The common name of the issuer.
	//
	// example:
	//
	// Encryption Everywhere DV TLS CA - G2
	IssuerCommonName *string `json:"IssuerCommonName,omitempty" xml:"IssuerCommonName,omitempty"`
	// The end time of the certificate chain validity period.
	//
	// example:
	//
	// 17352613180000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// The start time of the certificate chain validity period.
	//
	// example:
	//
	// 17322633180000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// The remaining days of the certificate chain validity period.
	//
	// example:
	//
	// 1000
	RemainDay *int32 `json:"RemainDay,omitempty" xml:"RemainDay,omitempty"`
}

func (s GetUserCertificateDetailResponseBodyCertChain) String() string {
	return dara.Prettify(s)
}

func (s GetUserCertificateDetailResponseBodyCertChain) GoString() string {
	return s.String()
}

func (s *GetUserCertificateDetailResponseBodyCertChain) GetCommonName() *string {
	return s.CommonName
}

func (s *GetUserCertificateDetailResponseBodyCertChain) GetIssuerCommonName() *string {
	return s.IssuerCommonName
}

func (s *GetUserCertificateDetailResponseBodyCertChain) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *GetUserCertificateDetailResponseBodyCertChain) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *GetUserCertificateDetailResponseBodyCertChain) GetRemainDay() *int32 {
	return s.RemainDay
}

func (s *GetUserCertificateDetailResponseBodyCertChain) SetCommonName(v string) *GetUserCertificateDetailResponseBodyCertChain {
	s.CommonName = &v
	return s
}

func (s *GetUserCertificateDetailResponseBodyCertChain) SetIssuerCommonName(v string) *GetUserCertificateDetailResponseBodyCertChain {
	s.IssuerCommonName = &v
	return s
}

func (s *GetUserCertificateDetailResponseBodyCertChain) SetNotAfter(v int64) *GetUserCertificateDetailResponseBodyCertChain {
	s.NotAfter = &v
	return s
}

func (s *GetUserCertificateDetailResponseBodyCertChain) SetNotBefore(v int64) *GetUserCertificateDetailResponseBodyCertChain {
	s.NotBefore = &v
	return s
}

func (s *GetUserCertificateDetailResponseBodyCertChain) SetRemainDay(v int32) *GetUserCertificateDetailResponseBodyCertChain {
	s.RemainDay = &v
	return s
}

func (s *GetUserCertificateDetailResponseBodyCertChain) Validate() error {
	return dara.Validate(s)
}

type GetUserCertificateDetailResponseBodyTags struct {
	// The tag key.
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

func (s GetUserCertificateDetailResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetUserCertificateDetailResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetUserCertificateDetailResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetUserCertificateDetailResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetUserCertificateDetailResponseBodyTags) SetTagKey(v string) *GetUserCertificateDetailResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetUserCertificateDetailResponseBodyTags) SetTagValue(v string) *GetUserCertificateDetailResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetUserCertificateDetailResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
