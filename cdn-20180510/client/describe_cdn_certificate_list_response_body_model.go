// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnCertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateListModel(v *DescribeCdnCertificateListResponseBodyCertificateListModel) *DescribeCdnCertificateListResponseBody
	GetCertificateListModel() *DescribeCdnCertificateListResponseBodyCertificateListModel
	SetRequestId(v string) *DescribeCdnCertificateListResponseBody
	GetRequestId() *string
}

type DescribeCdnCertificateListResponseBody struct {
	// Details about certificates.
	CertificateListModel *DescribeCdnCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// FC0E34AC-0239-44A7-AB0E-800DE522C8DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnCertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListResponseBody) GetCertificateListModel() *DescribeCdnCertificateListResponseBodyCertificateListModel {
	return s.CertificateListModel
}

func (s *DescribeCdnCertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnCertificateListResponseBody) SetCertificateListModel(v *DescribeCdnCertificateListResponseBodyCertificateListModel) *DescribeCdnCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeCdnCertificateListResponseBody) SetRequestId(v string) *DescribeCdnCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnCertificateListResponseBodyCertificateListModel struct {
	// The list of certificates.
	CertList *DescribeCdnCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
	// The number of certificates that are returned.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModel) GetCertList() *DescribeCdnCertificateListResponseBodyCertificateListModelCertList {
	return s.CertList
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModel) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeCdnCertificateListResponseBodyCertificateListModelCertList) *DescribeCdnCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeCdnCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModel) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModelCertList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertList) GetCert() []*DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	return s.Cert
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) *DescribeCdnCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertList) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 1
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// Certificate1
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The Common Name (CN) attribute of the certificate. In most cases, the CN is a domain name.
	//
	// example:
	//
	// example.com
	Common *string `json:"Common,omitempty" xml:"Common,omitempty"`
	// The fingerprint of the certificate.
	//
	// example:
	//
	// 2ED68FD33786C5B42950D40A6C50353575BB****
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// CO****
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 1512388610
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) GetCommon() *string {
	return s.Common
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetCertId(v int64) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetFingerprint(v string) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) SetLastTime(v int64) *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

func (s *DescribeCdnCertificateListResponseBodyCertificateListModelCertListCert) Validate() error {
	return dara.Validate(s)
}
