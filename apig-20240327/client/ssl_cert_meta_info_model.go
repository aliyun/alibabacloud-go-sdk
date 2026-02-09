// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSslCertMetaInfo interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *SslCertMetaInfo
	GetAlgorithm() *string
	SetCertId(v int64) *SslCertMetaInfo
	GetCertId() *int64
	SetCertIdentifier(v string) *SslCertMetaInfo
	GetCertIdentifier() *string
	SetCertName(v string) *SslCertMetaInfo
	GetCertName() *string
	SetCommonName(v string) *SslCertMetaInfo
	GetCommonName() *string
	SetDomain(v string) *SslCertMetaInfo
	GetDomain() *string
	SetDomainMatchCert(v bool) *SslCertMetaInfo
	GetDomainMatchCert() *bool
	SetFingerprint(v string) *SslCertMetaInfo
	GetFingerprint() *string
	SetInstanceId(v string) *SslCertMetaInfo
	GetInstanceId() *string
	SetIsChainCompleted(v bool) *SslCertMetaInfo
	GetIsChainCompleted() *bool
	SetIssuer(v string) *SslCertMetaInfo
	GetIssuer() *string
	SetKeySize(v string) *SslCertMetaInfo
	GetKeySize() *string
	SetMd5(v string) *SslCertMetaInfo
	GetMd5() *string
	SetNotAfterTimestamp(v int64) *SslCertMetaInfo
	GetNotAfterTimestamp() *int64
	SetNotBeforeTimestamp(v int64) *SslCertMetaInfo
	GetNotBeforeTimestamp() *int64
	SetSans(v string) *SslCertMetaInfo
	GetSans() *string
	SetSerialNo(v string) *SslCertMetaInfo
	GetSerialNo() *string
	SetSha2(v string) *SslCertMetaInfo
	GetSha2() *string
	SetSignAlgorithm(v string) *SslCertMetaInfo
	GetSignAlgorithm() *string
}

type SslCertMetaInfo struct {
	// The algorithm.
	//
	// example:
	//
	// RSA2048
	Algorithm *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 1234567890
	CertId *int64 `json:"certId,omitempty" xml:"certId,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// cert-123
	CertIdentifier *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// example.com
	CertName *string `json:"certName,omitempty" xml:"certName,omitempty"`
	// The name.
	//
	// example:
	//
	// example.com
	CommonName *string `json:"commonName,omitempty" xml:"commonName,omitempty"`
	// The domain name.
	//
	// example:
	//
	// api.example.com
	Domain *string `json:"domain,omitempty" xml:"domain,omitempty"`
	// The certificate matching the domain name.
	//
	// example:
	//
	// true
	DomainMatchCert *bool `json:"domainMatchCert,omitempty" xml:"domainMatchCert,omitempty"`
	// The certificate fingerprint.
	//
	// example:
	//
	// A1:B2:C3:D4:E5:F6:78:90:AB:CD:EF:12:34:56:78:90
	Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1234567890
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// isChainCompleted
	//
	// example:
	//
	// true
	IsChainCompleted *bool `json:"isChainCompleted,omitempty" xml:"isChainCompleted,omitempty"`
	// The certificate issuer.
	//
	// example:
	//
	// DigiCert Inc
	Issuer *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	// The key size.
	//
	// example:
	//
	// 2048
	KeySize *string `json:"keySize,omitempty" xml:"keySize,omitempty"`
	// The md5 value.
	//
	// example:
	//
	// A1B2C3D4E5F67890ABCDEF1234567890
	Md5 *string `json:"md5,omitempty" xml:"md5,omitempty"`
	// The time when the certificate expires.
	//
	// example:
	//
	// 1234567890000
	NotAfterTimestamp *int64 `json:"notAfterTimestamp,omitempty" xml:"notAfterTimestamp,omitempty"`
	// The time when the certificate starts to take effect.
	//
	// example:
	//
	// 1234567890000
	NotBeforeTimestamp *int64 `json:"notBeforeTimestamp,omitempty" xml:"notBeforeTimestamp,omitempty"`
	// sans
	//
	// example:
	//
	// *.example.com,api.example.com,www.example.com
	Sans *string `json:"sans,omitempty" xml:"sans,omitempty"`
	// The serial number.
	//
	// example:
	//
	// 03:A1:B2:C3:D4:E5:F6:78:90:AB:CD:EF:12:34:56:78:90
	SerialNo *string `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	// The sha2 value.
	//
	// example:
	//
	// A1B2C3D4E5F67890ABCDEF1234567890ABCDEF1234567890ABCDEF1234567890
	Sha2 *string `json:"sha2,omitempty" xml:"sha2,omitempty"`
	// The signature algorithm.
	//
	// example:
	//
	// sha256WithRSAEncryption
	SignAlgorithm *string `json:"signAlgorithm,omitempty" xml:"signAlgorithm,omitempty"`
}

func (s SslCertMetaInfo) String() string {
	return dara.Prettify(s)
}

func (s SslCertMetaInfo) GoString() string {
	return s.String()
}

func (s *SslCertMetaInfo) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *SslCertMetaInfo) GetCertId() *int64 {
	return s.CertId
}

func (s *SslCertMetaInfo) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *SslCertMetaInfo) GetCertName() *string {
	return s.CertName
}

func (s *SslCertMetaInfo) GetCommonName() *string {
	return s.CommonName
}

func (s *SslCertMetaInfo) GetDomain() *string {
	return s.Domain
}

func (s *SslCertMetaInfo) GetDomainMatchCert() *bool {
	return s.DomainMatchCert
}

func (s *SslCertMetaInfo) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *SslCertMetaInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SslCertMetaInfo) GetIsChainCompleted() *bool {
	return s.IsChainCompleted
}

func (s *SslCertMetaInfo) GetIssuer() *string {
	return s.Issuer
}

func (s *SslCertMetaInfo) GetKeySize() *string {
	return s.KeySize
}

func (s *SslCertMetaInfo) GetMd5() *string {
	return s.Md5
}

func (s *SslCertMetaInfo) GetNotAfterTimestamp() *int64 {
	return s.NotAfterTimestamp
}

func (s *SslCertMetaInfo) GetNotBeforeTimestamp() *int64 {
	return s.NotBeforeTimestamp
}

func (s *SslCertMetaInfo) GetSans() *string {
	return s.Sans
}

func (s *SslCertMetaInfo) GetSerialNo() *string {
	return s.SerialNo
}

func (s *SslCertMetaInfo) GetSha2() *string {
	return s.Sha2
}

func (s *SslCertMetaInfo) GetSignAlgorithm() *string {
	return s.SignAlgorithm
}

func (s *SslCertMetaInfo) SetAlgorithm(v string) *SslCertMetaInfo {
	s.Algorithm = &v
	return s
}

func (s *SslCertMetaInfo) SetCertId(v int64) *SslCertMetaInfo {
	s.CertId = &v
	return s
}

func (s *SslCertMetaInfo) SetCertIdentifier(v string) *SslCertMetaInfo {
	s.CertIdentifier = &v
	return s
}

func (s *SslCertMetaInfo) SetCertName(v string) *SslCertMetaInfo {
	s.CertName = &v
	return s
}

func (s *SslCertMetaInfo) SetCommonName(v string) *SslCertMetaInfo {
	s.CommonName = &v
	return s
}

func (s *SslCertMetaInfo) SetDomain(v string) *SslCertMetaInfo {
	s.Domain = &v
	return s
}

func (s *SslCertMetaInfo) SetDomainMatchCert(v bool) *SslCertMetaInfo {
	s.DomainMatchCert = &v
	return s
}

func (s *SslCertMetaInfo) SetFingerprint(v string) *SslCertMetaInfo {
	s.Fingerprint = &v
	return s
}

func (s *SslCertMetaInfo) SetInstanceId(v string) *SslCertMetaInfo {
	s.InstanceId = &v
	return s
}

func (s *SslCertMetaInfo) SetIsChainCompleted(v bool) *SslCertMetaInfo {
	s.IsChainCompleted = &v
	return s
}

func (s *SslCertMetaInfo) SetIssuer(v string) *SslCertMetaInfo {
	s.Issuer = &v
	return s
}

func (s *SslCertMetaInfo) SetKeySize(v string) *SslCertMetaInfo {
	s.KeySize = &v
	return s
}

func (s *SslCertMetaInfo) SetMd5(v string) *SslCertMetaInfo {
	s.Md5 = &v
	return s
}

func (s *SslCertMetaInfo) SetNotAfterTimestamp(v int64) *SslCertMetaInfo {
	s.NotAfterTimestamp = &v
	return s
}

func (s *SslCertMetaInfo) SetNotBeforeTimestamp(v int64) *SslCertMetaInfo {
	s.NotBeforeTimestamp = &v
	return s
}

func (s *SslCertMetaInfo) SetSans(v string) *SslCertMetaInfo {
	s.Sans = &v
	return s
}

func (s *SslCertMetaInfo) SetSerialNo(v string) *SslCertMetaInfo {
	s.SerialNo = &v
	return s
}

func (s *SslCertMetaInfo) SetSha2(v string) *SslCertMetaInfo {
	s.Sha2 = &v
	return s
}

func (s *SslCertMetaInfo) SetSignAlgorithm(v string) *SslCertMetaInfo {
	s.SignAlgorithm = &v
	return s
}

func (s *SslCertMetaInfo) Validate() error {
	return dara.Validate(s)
}
