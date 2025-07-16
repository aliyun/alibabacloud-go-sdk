// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainByCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeCdnDomainByCertificateResponseBodyCertInfos) *DescribeCdnDomainByCertificateResponseBody
	GetCertInfos() *DescribeCdnDomainByCertificateResponseBodyCertInfos
	SetRequestId(v string) *DescribeCdnDomainByCertificateResponseBody
	GetRequestId() *string
}

type DescribeCdnDomainByCertificateResponseBody struct {
	// The information about the certificate.
	CertInfos *DescribeCdnDomainByCertificateResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// ASAF2FDS-12SADSA-DDSAE3D-DSADCD4C-CDADS2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDomainByCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainByCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainByCertificateResponseBody) GetCertInfos() *DescribeCdnDomainByCertificateResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeCdnDomainByCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDomainByCertificateResponseBody) SetCertInfos(v *DescribeCdnDomainByCertificateResponseBodyCertInfos) *DescribeCdnDomainByCertificateResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBody) SetRequestId(v string) *DescribeCdnDomainByCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainByCertificateResponseBodyCertInfos struct {
	CertInfo []*DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainByCertificateResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainByCertificateResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfos) GetCertInfo() []*DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfos) SetCertInfo(v []*DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) *DescribeCdnDomainByCertificateResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo struct {
	// Indicates whether the SSL certificate is obsolete. Valid values:
	//
	// 	- **yes**: The SSL certificate is obsolete.
	//
	// 	- **no**: The SSL certificate is working as expected.
	//
	// example:
	//
	// yes
	CertCaIsLegacy *string `json:"CertCaIsLegacy,omitempty" xml:"CertCaIsLegacy,omitempty"`
	// The expiration time of the certificate.
	//
	// example:
	//
	// Nov 29 00:00:00 2016 GMT
	CertExpireTime *string `json:"CertExpireTime,omitempty" xml:"CertExpireTime,omitempty"`
	// Indicates whether the SSL certificate is expired. Valid values:
	//
	// 	- **yes**: The SSL certificate is expired.
	//
	// 	- **no**: The SSL certificate is not expired.
	//
	// example:
	//
	// yes
	CertExpired *string `json:"CertExpired,omitempty" xml:"CertExpired,omitempty"`
	// The time when the certificate became effective.
	//
	// example:
	//
	// Nov 29 23:59:59 2017 GMT
	CertStartTime *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The name of the SSL certificate owner.
	//
	// example:
	//
	// owner
	CertSubjectCommonName *string `json:"CertSubjectCommonName,omitempty" xml:"CertSubjectCommonName,omitempty"`
	// The type of the certificate. Valid values: **RSA**, **DSA**, and **ECDSA**.
	//
	// example:
	//
	// RSA
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// If a value is returned, the value matches the SSL certificate. Multiple domain names are separated by commas (,).
	//
	// example:
	//
	// example.com,aliyundoc.com
	DomainList *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	// The domain names (DNS fields) that match the SSL certificate. Multiple domain names are separated by commas (,).
	//
	// example:
	//
	// *.example.com,aliyundoc.com
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// C=US, O=Symantec Corporation, OU=Symantec Trust Network, OU=Domain Validated SSL, CN=Symantec Basic DV SSL CA - G1
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
}

func (s DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertCaIsLegacy() *string {
	return s.CertCaIsLegacy
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertExpired() *string {
	return s.CertExpired
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertSubjectCommonName() *string {
	return s.CertSubjectCommonName
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) GetDomainList() *string {
	return s.DomainList
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) GetDomainNames() *string {
	return s.DomainNames
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertCaIsLegacy(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertCaIsLegacy = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertExpired(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertExpired = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertSubjectCommonName(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertSubjectCommonName = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetDomainList(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.DomainList = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetDomainNames(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.DomainNames = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) SetIssuer(v string) *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.Issuer = &v
	return s
}

func (s *DescribeCdnDomainByCertificateResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
