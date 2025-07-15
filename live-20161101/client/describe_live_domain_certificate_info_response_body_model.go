// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainCertificateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeLiveDomainCertificateInfoResponseBodyCertInfos) *DescribeLiveDomainCertificateInfoResponseBody
	GetCertInfos() *DescribeLiveDomainCertificateInfoResponseBodyCertInfos
	SetRequestId(v string) *DescribeLiveDomainCertificateInfoResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainCertificateInfoResponseBody struct {
	// The certificate information.
	CertInfos *DescribeLiveDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 5C1E43DC-9E51-4771-82C0-7D5ECEB547A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainCertificateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoResponseBody) GetCertInfos() *DescribeLiveDomainCertificateInfoResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeLiveDomainCertificateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeLiveDomainCertificateInfoResponseBodyCertInfos) *DescribeLiveDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeLiveDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainCertificateInfoResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfos) GetCertInfo() []*DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeLiveDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
	// The streaming domain or ingest domain that matches the certificate.
	//
	// example:
	//
	// example.com
	CertDomainName *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	// The expiration time of the certificate. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-06-03T22:03:39Z
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	// The validity period of the certificate.
	//
	// 	- If the validity period is greater than 12 months, XX years XX months is displayed. For example, 2 years 3 months indicates that the validity period is 2 years and 3 months.
	//
	// 	- If the validity period is less than 12 months, XX months is displayed. For example, 3 months indicates that the validity period is 3 months.
	//
	// example:
	//
	// 3 months
	CertLife *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// Cert-****
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// Let\\"s Encrypt
	CertOrg *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **free**: a free certificate (for testing)
	//
	// 	- **cas**: a certificate that is purchased from Certificate Management Service
	//
	// 	- **upload**: a custom certificate that you uploaded
	//
	// example:
	//
	// cas
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The streaming domain or ingest domain.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of HTTPS. Valid values:
	//
	// 	- **on**: HTTPS is enabled.
	//
	// 	- **off**: HTTPS is disabled.
	//
	// example:
	//
	// on
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The public key of the certificate.
	//
	// example:
	//
	// yourSSLPub
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	// The status of the free certificate that is used for testing. Valid values:
	//
	// 	- **success**: The certificate is effective.
	//
	// 	- **checking**: The system is checking whether the domain name is mapped to the CNAME that is assigned by ApsaraVideo Live.
	//
	// 	- **cname_error**: The domain name is not mapped to the CNAME that is assigned by ApsaraVideo Live.
	//
	// 	- **domain_invalid**: The domain name contains invalid characters.
	//
	// 	- **unsupport_wildcard**: The domain name is a wildcard domain name, which is not supported.
	//
	// 	- **applying**: The certificate is in the application progress.
	//
	// 	- **failed**: The application for the certificate failed.
	//
	// >  The **Status*	- parameter is valid only if the certificate is a free certificate for testing. If the certificate is not a free certificate for testing, an empty value is returned.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertDomainName() *string {
	return s.CertDomainName
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertLife() *string {
	return s.CertLife
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertOrg() *string {
	return s.CertOrg
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLProtocol(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLPub(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLPub = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

func (s *DescribeLiveDomainCertificateInfoResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
