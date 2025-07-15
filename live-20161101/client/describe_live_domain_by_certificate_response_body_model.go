// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainByCertificateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertInfos(v *DescribeLiveDomainByCertificateResponseBodyCertInfos) *DescribeLiveDomainByCertificateResponseBody
	GetCertInfos() *DescribeLiveDomainByCertificateResponseBodyCertInfos
	SetRequestId(v string) *DescribeLiveDomainByCertificateResponseBody
	GetRequestId() *string
}

type DescribeLiveDomainByCertificateResponseBody struct {
	// The information about the certificate.
	CertInfos *DescribeLiveDomainByCertificateResponseBodyCertInfos `json:"CertInfos,omitempty" xml:"CertInfos,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// ASAF2FDS-12SADSA-DDSAE3D-DSADCD4C-CDADS2D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveDomainByCertificateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainByCertificateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainByCertificateResponseBody) GetCertInfos() *DescribeLiveDomainByCertificateResponseBodyCertInfos {
	return s.CertInfos
}

func (s *DescribeLiveDomainByCertificateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveDomainByCertificateResponseBody) SetCertInfos(v *DescribeLiveDomainByCertificateResponseBodyCertInfos) *DescribeLiveDomainByCertificateResponseBody {
	s.CertInfos = v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBody) SetRequestId(v string) *DescribeLiveDomainByCertificateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainByCertificateResponseBodyCertInfos struct {
	CertInfo []*DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo `json:"CertInfo,omitempty" xml:"CertInfo,omitempty" type:"Repeated"`
}

func (s DescribeLiveDomainByCertificateResponseBodyCertInfos) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainByCertificateResponseBodyCertInfos) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfos) GetCertInfo() []*DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo {
	return s.CertInfo
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfos) SetCertInfo(v []*DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) *DescribeLiveDomainByCertificateResponseBodyCertInfos {
	s.CertInfo = v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfos) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo struct {
	// Indicates whether the SSL certificate is obsolete. Valid values:
	//
	// 	- **yes**
	//
	// 	- **no**
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
	// 	- **yes**
	//
	// 	- **no**
	//
	// example:
	//
	// yes
	CertExpired *string `json:"CertExpired,omitempty" xml:"CertExpired,omitempty"`
	// The effective time of the certificate.
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
	// The list of domain names. If a value is returned, the value matches the SSL certificate. Multiple domain names are separated by commas (,).
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
	// The certificate authority (CA) that issued the SSL certificate.
	//
	// example:
	//
	// C=US, O=Symantec Corporation, OU=Symantec Trust Network, OU=Domain Validated SSL, CN=Symantec Basic DV SSL CA - G1
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
}

func (s DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) GetCertCaIsLegacy() *string {
	return s.CertCaIsLegacy
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) GetCertExpireTime() *string {
	return s.CertExpireTime
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) GetCertExpired() *string {
	return s.CertExpired
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) GetCertStartTime() *string {
	return s.CertStartTime
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) GetCertSubjectCommonName() *string {
	return s.CertSubjectCommonName
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) GetCertType() *string {
	return s.CertType
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) GetDomainList() *string {
	return s.DomainList
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) GetDomainNames() *string {
	return s.DomainNames
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) SetCertCaIsLegacy(v string) *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertCaIsLegacy = &v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) SetCertExpireTime(v string) *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertExpireTime = &v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) SetCertExpired(v string) *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertExpired = &v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) SetCertStartTime(v string) *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertStartTime = &v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) SetCertSubjectCommonName(v string) *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertSubjectCommonName = &v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) SetCertType(v string) *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo {
	s.CertType = &v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) SetDomainList(v string) *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo {
	s.DomainList = &v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) SetDomainNames(v string) *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo {
	s.DomainNames = &v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) SetIssuer(v string) *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo {
	s.Issuer = &v
	return s
}

func (s *DescribeLiveDomainByCertificateResponseBodyCertInfosCertInfo) Validate() error {
	return dara.Validate(s)
}
