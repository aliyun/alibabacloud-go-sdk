// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainByCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeDcdnDomainByCertificateResponseBodyCertInfos) *DescribeDcdnDomainByCertificateResponseBody
	GetCertInfos() *DescribeDcdnDomainByCertificateResponseBodyCertInfos
	SetRequestId(v string) *DescribeDcdnDomainByCertificateResponseBody
	GetRequestId() *string
}

type DescribeDcdnDomainByCertificateResponseBody struct {
	// The information about the certificate.
	CertInfos *DescribeDcdnDomainByCertificateResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// ASAF2FDS-12SADSA-DDSAE3D-DSADCD4C-CDADS2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnDomainByCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainByCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainByCertificateResponseBody) GetCertInfos() *DescribeDcdnDomainByCertificateResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeDcdnDomainByCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnDomainByCertificateResponseBody) SetCertInfos(v *DescribeDcdnDomainByCertificateResponseBodyCertInfos) *DescribeDcdnDomainByCertificateResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBody) SetRequestId(v string) *DescribeDcdnDomainByCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainByCertificateResponseBodyCertInfos struct {
	CertInfo []*DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeDcdnDomainByCertificateResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainByCertificateResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfos) GetCertInfo() []*DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfos) SetCertInfo(v []*DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) *DescribeDcdnDomainByCertificateResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo struct {
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
	// The time at which the certificate expires.
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
	// The time at which the certificate became effective.
	//
	// example:
	//
	// Nov 29 23:59:59 2017 GMT
	CertStartTime *string `json:"CertStartTime,omitempty" xml:"CertStartTime,omitempty"`
	// The name of the SSL certificate owner.
	//
	// example:
	//
	// example.aliyundoc.com
	CertSubjectCommonName *string `json:"CertSubjectCommonName,omitempty" xml:"CertSubjectCommonName,omitempty"`
	// The type of the certificate. Valid values: **RSA**, **DSA**, and **ECDSA**.
	//
	// example:
	//
	// RSA
	CertType *string `json:"CertType,omitempty" xml:"CertType,omitempty"`
	// The list of domain names that use the certificate.
	//
	// If one or more domain names are returned, the domain names are matched with the specified certificate. Multiple domain names are separated with commas (,).
	//
	// example:
	//
	// example.com,example.org
	DomainList *string `json:"DomainList,omitempty" xml:"DomainList,omitempty"`
	// The domain names (DNS fields) that match the certificate. Multiple domain names are separated with commas (,).
	//
	// example:
	//
	// *.example.com,example.org
	DomainNames *string `json:"DomainNames,omitempty" xml:"DomainNames,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// C=US, O=Symantec Corporation, OU=Symantec Trust Network, OU=Domain Validated SSL, CN=Symantec Basic DV SSL CA - G1
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
}

func (s DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertCaIsLegacy() *string {
	return s.CertCaIsLegacy
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertExpired() *string {
	return s.CertExpired
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertSubjectCommonName() *string {
	return s.CertSubjectCommonName
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) GetDomainList() *string {
	return s.DomainList
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) GetDomainNames() *string {
	return s.DomainNames
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertCaIsLegacy(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertCaIsLegacy = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertExpired(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertExpired = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertSubjectCommonName(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertSubjectCommonName = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetDomainList(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.DomainList = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetDomainNames(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.DomainNames = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) SetIssuer(v string) *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo {
	s.Issuer = &v
	return s
}

func (s *DescribeDcdnDomainByCertificateResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
