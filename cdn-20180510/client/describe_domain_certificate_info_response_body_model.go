// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCertificateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeDomainCertificateInfoResponseBodyCertInfos) *DescribeDomainCertificateInfoResponseBody
	GetCertInfos() *DescribeDomainCertificateInfoResponseBodyCertInfos
	SetRequestId(v string) *DescribeDomainCertificateInfoResponseBody
	GetRequestId() *string
}

type DescribeDomainCertificateInfoResponseBody struct {
	// The information about the certificate.
	CertInfos *DescribeDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5C1E43DC-9E51-4771-82C0-7D5ECEB547A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDomainCertificateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoResponseBody) GetCertInfos() *DescribeDomainCertificateInfoResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeDomainCertificateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeDomainCertificateInfoResponseBodyCertInfos) *DescribeDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeDomainCertificateInfoResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfos) GetCertInfo() []*DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
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
	// The unit of the validity period of the certificate. Valid values:
	//
	// 	- **months**
	//
	// 	- **years**
	//
	// example:
	//
	// months
	CertLife *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// example.com
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The name of the certificate authority (CA) that issued the certificate.
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
	// The time when the certificate became effective.
	//
	// example:
	//
	// 2018-06-03T22:03:39Z
	CertStartTime *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The type of the certificate.
	//
	// 	- **free**: a free certificate
	//
	// 	- **cas**: a certificate that is purchased by using Certificate Management Service
	//
	// 	- **upload**: a custom certificate that you upload
	//
	// example:
	//
	// free
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The time at which the certificate was updated.
	//
	// example:
	//
	// 2018-06-03T22:03:39Z
	CertUpdateTime *string `json:"CertUpdateTime,omitempty" xml:"CertUpdateTime,omitempty"`
	// The CNAME status of the domain name.
	//
	// 	- **ok**: The domain name points to the CNAME assigned by Alibaba Cloud CDN.
	//
	// 	- **cname_error**: An error occurred and the domain name cannot point to the CNAME.
	//
	// 	- **op_domain_cname_error*	- : An error occurred to the CNAME of the top-level domain. The domain name cannot point to the CNAME.
	//
	// 	- **unsupport_wildcard**: The wildcard domain name is not supported.
	//
	// example:
	//
	// ok
	DomainCnameStatus *string `json:"DomainCnameStatus,omitempty" xml:"DomainCnameStatus,omitempty"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The public key of the certificate.
	//
	// example:
	//
	// asdadaxxxx
	ServerCertificate *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
	// The status of HTTPS.
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	ServerCertificateStatus *string `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	// The status of the certificate. Valid values:
	//
	// 	- **success**: The certificate has taken effect.
	//
	// 	- **checking**: The system is checking whether the domain name is using Alibaba Cloud CDN.
	//
	// 	- **cname_error**: No valid CNAME record has been added for the domain name.
	//
	// 	- **top_domain_cname_error**: No valid CNAME record has been added for the top-level domain.
	//
	// 	- **domain_invalid**: The domain name contains invalid characters.
	//
	// 	- **unsupport_wildcard**: The domain name is a wildcard domain name. Wildcard domain names are not supported.
	//
	// 	- **applying**: The certificate application is in progress.
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

func (s DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertDomainName() *string {
	return s.CertDomainName
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertId() *string {
	return s.CertId
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertLife() *string {
	return s.CertLife
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertOrg() *string {
	return s.CertOrg
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertUpdateTime() *string {
	return s.CertUpdateTime
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainCnameStatus() *string {
	return s.DomainCnameStatus
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetServerCertificate() *string {
	return s.ServerCertificate
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetServerCertificateStatus() *string {
	return s.ServerCertificateStatus
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertId(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertId = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertRegion(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertRegion = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertUpdateTime(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertUpdateTime = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainCnameStatus(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainCnameStatus = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificate(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificate = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificateStatus(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificateStatus = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

func (s *DescribeDomainCertificateInfoResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
