// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnCertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateListModel(v *DescribeDcdnCertificateListResponseBodyCertificateListModel) *DescribeDcdnCertificateListResponseBody
	GetCertificateListModel() *DescribeDcdnCertificateListResponseBodyCertificateListModel
	SetRequestId(v string) *DescribeDcdnCertificateListResponseBody
	GetRequestId() *string
}

type DescribeDcdnCertificateListResponseBody struct {
	// Details about certificates.
	CertificateListModel *DescribeDcdnCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// FC0E34AC-0239-44A7-AB0E-800DE522C8DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnCertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListResponseBody) GetCertificateListModel() *DescribeDcdnCertificateListResponseBodyCertificateListModel {
	return s.CertificateListModel
}

func (s *DescribeDcdnCertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnCertificateListResponseBody) SetCertificateListModel(v *DescribeDcdnCertificateListResponseBodyCertificateListModel) *DescribeDcdnCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeDcdnCertificateListResponseBody) SetRequestId(v string) *DescribeDcdnCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnCertificateListResponseBodyCertificateListModel struct {
	// Details about each certificate.
	CertList *DescribeDcdnCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
	// The number of certificates.
	//
	// example:
	//
	// 123
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModel) GetCertList() *DescribeDcdnCertificateListResponseBodyCertificateListModelCertList {
	return s.CertList
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModel) GetCount() *int32 {
	return s.Count
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeDcdnCertificateListResponseBodyCertificateListModelCertList) *DescribeDcdnCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeDcdnCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModel) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModelCertList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertList) GetCert() []*DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	return s.Cert
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertList) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 123
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// Certificate 2
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
	// 0151xxxx
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// DigiCert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The timestamp.
	//
	// example:
	//
	// 1548065550
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) GetCommon() *string {
	return s.Common
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetCertId(v int64) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetFingerprint(v string) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) SetLastTime(v int64) *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

func (s *DescribeDcdnCertificateListResponseBodyCertificateListModelCertListCert) Validate() error {
	return dara.Validate(s)
}
