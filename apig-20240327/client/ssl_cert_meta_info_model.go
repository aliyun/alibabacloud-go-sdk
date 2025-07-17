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
	Algorithm          *string `json:"algorithm,omitempty" xml:"algorithm,omitempty"`
	CertId             *int64  `json:"certId,omitempty" xml:"certId,omitempty"`
	CertIdentifier     *string `json:"certIdentifier,omitempty" xml:"certIdentifier,omitempty"`
	CertName           *string `json:"certName,omitempty" xml:"certName,omitempty"`
	CommonName         *string `json:"commonName,omitempty" xml:"commonName,omitempty"`
	Domain             *string `json:"domain,omitempty" xml:"domain,omitempty"`
	DomainMatchCert    *bool   `json:"domainMatchCert,omitempty" xml:"domainMatchCert,omitempty"`
	Fingerprint        *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
	InstanceId         *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsChainCompleted   *bool   `json:"isChainCompleted,omitempty" xml:"isChainCompleted,omitempty"`
	Issuer             *string `json:"issuer,omitempty" xml:"issuer,omitempty"`
	KeySize            *string `json:"keySize,omitempty" xml:"keySize,omitempty"`
	Md5                *string `json:"md5,omitempty" xml:"md5,omitempty"`
	NotAfterTimestamp  *int64  `json:"notAfterTimestamp,omitempty" xml:"notAfterTimestamp,omitempty"`
	NotBeforeTimestamp *int64  `json:"notBeforeTimestamp,omitempty" xml:"notBeforeTimestamp,omitempty"`
	Sans               *string `json:"sans,omitempty" xml:"sans,omitempty"`
	SerialNo           *string `json:"serialNo,omitempty" xml:"serialNo,omitempty"`
	Sha2               *string `json:"sha2,omitempty" xml:"sha2,omitempty"`
	SignAlgorithm      *string `json:"signAlgorithm,omitempty" xml:"signAlgorithm,omitempty"`
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
