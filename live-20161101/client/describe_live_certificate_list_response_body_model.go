// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveCertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateListModel(v *DescribeLiveCertificateListResponseBodyCertificateListModel) *DescribeLiveCertificateListResponseBody
	GetCertificateListModel() *DescribeLiveCertificateListResponseBodyCertificateListModel
	SetRequestId(v string) *DescribeLiveCertificateListResponseBody
	GetRequestId() *string
}

type DescribeLiveCertificateListResponseBody struct {
	// The details.
	CertificateListModel *DescribeLiveCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLiveCertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateListResponseBody) GetCertificateListModel() *DescribeLiveCertificateListResponseBodyCertificateListModel {
	return s.CertificateListModel
}

func (s *DescribeLiveCertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveCertificateListResponseBody) SetCertificateListModel(v *DescribeLiveCertificateListResponseBodyCertificateListModel) *DescribeLiveCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeLiveCertificateListResponseBody) SetRequestId(v string) *DescribeLiveCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveCertificateListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveCertificateListResponseBodyCertificateListModel struct {
	// The certificates.
	CertList *DescribeLiveCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
	// The number of certificates.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeLiveCertificateListResponseBodyCertificateListModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModel) GetCertList() *DescribeLiveCertificateListResponseBodyCertificateListModelCertList {
	return s.CertList
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModel) GetCount() *int32 {
	return s.Count
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeLiveCertificateListResponseBodyCertificateListModelCertList) *DescribeLiveCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeLiveCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModel) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeLiveCertificateListResponseBodyCertificateListModelCertList) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertList) GetCert() []*DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert {
	return s.Cert
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) *DescribeLiveCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertList) Validate() error {
	return dara.Validate(s)
}

type DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 123456
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// Cert-****
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The Common Name (CN) attribute of the certificate.
	//
	// example:
	//
	// example.net
	Common *string `json:"Common,omitempty" xml:"Common,omitempty"`
	// The fingerprint of the certificate.
	//
	// example:
	//
	// ****
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// *****
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The time when the certificate was issued. Unit: seconds.
	//
	// example:
	//
	// 1512388659
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
}

func (s DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) GetCertName() *string {
	return s.CertName
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) GetCommon() *string {
	return s.Common
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) SetCertId(v int64) *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) SetFingerprint(v string) *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) SetLastTime(v int64) *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

func (s *DescribeLiveCertificateListResponseBodyCertificateListModelCertListCert) Validate() error {
	return dara.Validate(s)
}
