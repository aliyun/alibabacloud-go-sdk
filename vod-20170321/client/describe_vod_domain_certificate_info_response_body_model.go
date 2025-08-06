// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainCertificateInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeVodDomainCertificateInfoResponseBodyCertInfos) *DescribeVodDomainCertificateInfoResponseBody
	GetCertInfos() *DescribeVodDomainCertificateInfoResponseBodyCertInfos
	SetRequestId(v string) *DescribeVodDomainCertificateInfoResponseBody
	GetRequestId() *string
}

type DescribeVodDomainCertificateInfoResponseBody struct {
	// The certificate information.
	CertInfos *DescribeVodDomainCertificateInfoResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 5C1E43DC-9E51-4771-****-7D5ECEB547A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodDomainCertificateInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoResponseBody) GetCertInfos() *DescribeVodDomainCertificateInfoResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeVodDomainCertificateInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodDomainCertificateInfoResponseBody) SetCertInfos(v *DescribeVodDomainCertificateInfoResponseBodyCertInfos) *DescribeVodDomainCertificateInfoResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBody) SetRequestId(v string) *DescribeVodDomainCertificateInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainCertificateInfoResponseBodyCertInfos struct {
	CertInfo []*DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfos) GetCertInfo() []*DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfos) SetCertInfo(v []*DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) *DescribeVodDomainCertificateInfoResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo struct {
	// The domain name that matches the certificate.
	//
	// example:
	//
	// example.com
	CertDomainName *string `json:"CertDomainName,omitempty" xml:"CertDomainName,omitempty"`
	// The time at which the certificate expires. The time is in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-06-03T13:03:39Z
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 13227737-cn-hangzhou
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The validity period of the certificate. Unit: months or years.
	//
	// example:
	//
	// 3 months
	CertLife *string `json:"CertLife,omitempty" xml:"CertLife,omitempty"`
	// The certificate name.
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
	// The time when the certificate became effective.
	//
	// example:
	//
	// 2023-04-26T20:23:38Z
	CertStartTime *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **free**: a free certificate.
	//
	// 	- **cas**: a certificate that is purchased from Certificate Management Service.
	//
	// 	- **upload**: a user-uploaded certificate.
	//
	// example:
	//
	// free
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The time at which the certificate was updated.
	//
	// example:
	//
	// 2023-04-26T20:23:38Z
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
	// The accelerated domain name whose ICP filing status you want to update.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The public key of the certificate.
	//
	// example:
	//
	// ****
	ServerCertificate *string `json:"ServerCertificate,omitempty" xml:"ServerCertificate,omitempty"`
	// The status of the SSL certificate.
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// checking
	ServerCertificateStatus *string `json:"ServerCertificateStatus,omitempty" xml:"ServerCertificateStatus,omitempty"`
	// The status of the certificate.
	//
	// 	- **success**: The certificate is in effect.
	//
	// 	- **checking**: The system is checking whether the domain name is added to ApsaraVideo VOD.
	//
	// 	- **cname_error**: The domain name is not added to ApsaraVideo VOD.
	//
	// 	- **domain_invalid**: The domain name contains invalid characters.
	//
	// 	- **unsupport_wildcard**: The domain name is a wildcard domain name. Wildcard domain names are not supported.
	//
	// 	- **applying**: The certificate application is in progress.
	//
	// 	- **failed**: The certificate application failed.
	//
	// >  A value is returned for this parameter only when you set `CertType` to `free`. Otherwise, an empty value is returned for this parameter.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertDomainName() *string {
	return s.CertDomainName
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertId() *string {
	return s.CertId
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertLife() *string {
	return s.CertLife
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertOrg() *string {
	return s.CertOrg
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetCertUpdateTime() *string {
	return s.CertUpdateTime
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainCnameStatus() *string {
	return s.DomainCnameStatus
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetServerCertificate() *string {
	return s.ServerCertificate
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetServerCertificateStatus() *string {
	return s.ServerCertificateStatus
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertDomainName(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertDomainName = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertId(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertId = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertLife(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertLife = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertName(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertName = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertOrg(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertOrg = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertRegion(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertRegion = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetCertUpdateTime(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.CertUpdateTime = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainCnameStatus(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainCnameStatus = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetDomainName(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.DomainName = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificate(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificate = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetServerCertificateStatus(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.ServerCertificateStatus = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) SetStatus(v string) *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo {
	s.Status = &v
	return s
}

func (s *DescribeVodDomainCertificateInfoResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
