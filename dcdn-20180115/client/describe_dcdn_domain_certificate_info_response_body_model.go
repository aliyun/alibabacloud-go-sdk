// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainCertificateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos) *DescribeDcdnDomainCertificateInfoResponseBody
	GetCertInfos() *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos
	SetRequestId(v string) *DescribeDcdnDomainCertificateInfoResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainCertificateInfoResponseBody struct {
	// The information about the certificate.
	CertInfos *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5C1E43DC-9E51-4771-82C0-7D5ECEB547A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainCertificateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCertificateInfoResponseBody) GetCertInfos() *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeDcdnDomainCertificateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos) *DescribeDcdnDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeDcdnDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainCertificateInfoResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos) GetCertInfo() []*DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
	// The domain name that matches the certificate.
	//
	// example:
	//
	// example.com
	CertDomainName *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	// The time at which the certificate expires.
	//
	// example:
	//
	// 2018-06-03T22:03:39Z
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 9002448
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The validity period of the certificate. Unit: **months*	- or **years**.
	//
	// example:
	//
	// 3 months
	CertLife *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// cert-example.com
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// Let\\"s Encrypt
	CertOrg *string `json:"CertOrg,omitempty" xml:"CertOrg,omitempty"`
	// The region where the certificate is used.
	//
	// example:
	//
	// cn-hangzhou
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// The type of the certificate.
	//
	// 	- **cas**: a certificate that is purchased by using Certificates Management Service
	//
	// 	- **upload**: a custom certificate that you upload
	//
	// example:
	//
	// cas
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The status of HTTPS. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	SSLProtocol *string `json:"SSLProtocol,omitempty" xml:"SSLProtocol,omitempty"`
	// The public key of the certificate.
	//
	// example:
	//
	// xxxx
	SSLPub *string `json:"SSLPub,omitempty" xml:"SSLPub,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- **success**: The certificate has taken effect.
	//
	// 	- **checking**: The system is checking whether the domain name is using Dynamic Route for CDN (DCDN).
	//
	// 	- **cname_error**: The domain name is not using DCDN.
	//
	// 	- **domain_invalid**: The domain name contains invalid characters.
	//
	// 	- **unsupport_wildcard**: The wildcard domain name is not supported.
	//
	// 	- **applying**: Certificate application is in progress.
	//
	// 	- **get_token_timeout**: The certificate application request has timed out.
	//
	// 	- **check_token_timeout**: The verification has timed out.
	//
	// 	- **get_cert_timeout**: The request to obtain the certificate has timed out.
	//
	// 	- **failed**: The certificate application request failed.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertDomainName() *string {
	return s.CertDomainName
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertId() *string {
	return s.CertId
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertLife() *string {
	return s.CertLife
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertOrg() *string {
	return s.CertOrg
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetSSLProtocol() *string {
	return s.SSLProtocol
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetSSLPub() *string {
	return s.SSLPub
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertId(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertId = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertRegion(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertRegion = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLProtocol(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLProtocol = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetSSLPub(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.SSLPub = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

func (s *DescribeDcdnDomainCertificateInfoResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
