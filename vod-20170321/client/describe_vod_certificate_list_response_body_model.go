// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodCertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateListModel(v *DescribeVodCertificateListResponseBodyCertificateListModel) *DescribeVodCertificateListResponseBody
	GetCertificateListModel() *DescribeVodCertificateListResponseBodyCertificateListModel
	SetRequestId(v string) *DescribeVodCertificateListResponseBody
	GetRequestId() *string
}

type DescribeVodCertificateListResponseBody struct {
	// The information about each certificate.
	CertificateListModel *DescribeVodCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// FC0E34AC-0239-44A7-****-800DE522C8DA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodCertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListResponseBody) GetCertificateListModel() *DescribeVodCertificateListResponseBodyCertificateListModel {
	return s.CertificateListModel
}

func (s *DescribeVodCertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodCertificateListResponseBody) SetCertificateListModel(v *DescribeVodCertificateListResponseBodyCertificateListModel) *DescribeVodCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeVodCertificateListResponseBody) SetRequestId(v string) *DescribeVodCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodCertificateListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodCertificateListResponseBodyCertificateListModel struct {
	// The list of certificates.
	CertList *DescribeVodCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
	// The number of certificates that are returned.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeVodCertificateListResponseBodyCertificateListModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModel) GetCertList() *DescribeVodCertificateListResponseBodyCertificateListModelCertList {
	return s.CertList
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModel) GetCount() *int32 {
	return s.Count
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeVodCertificateListResponseBodyCertificateListModelCertList) *DescribeVodCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeVodCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModel) Validate() error {
	return dara.Validate(s)
}

type DescribeVodCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeVodCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeVodCertificateListResponseBodyCertificateListModelCertList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertList) GetCert() []*DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	return s.Cert
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) *DescribeVodCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertList) Validate() error {
	return dara.Validate(s)
}

type DescribeVodCertificateListResponseBodyCertificateListModelCertListCert struct {
	// The algorithm.
	//
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 235437
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The ID of the certificate.
	//
	// example:
	//
	// 14173772-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// certificate
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The common name of the certificate.
	//
	// example:
	//
	// test
	Common *string `json:"Common,omitempty" xml:"Common,omitempty"`
	// The time when the certificate was created.
	//
	// example:
	//
	// 1725206400000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// DomainMatchCert.
	//
	// example:
	//
	// false
	DomainMatchCert *bool `json:"DomainMatchCert,omitempty" xml:"DomainMatchCert,omitempty"`
	// The time when the certificate expired.
	//
	// example:
	//
	// 1759507200000
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The fingerprint of the certificate.
	//
	// example:
	//
	// ****
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cert-cn-cd806ae0fdfbfa60
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// ****
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The time when the certificate was issued. Unit: seconds.
	//
	// example:
	//
	// 1512388610
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
	// The signAlgorithm.
	//
	// example:
	//
	// sha256withrsa
	SignAlgorithm *string `json:"SignAlgorithm,omitempty" xml:"SignAlgorithm,omitempty"`
}

func (s DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetCommon() *string {
	return s.Common
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetDomainMatchCert() *bool {
	return s.DomainMatchCert
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) GetSignAlgorithm() *string {
	return s.SignAlgorithm
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetAlgorithm(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.Algorithm = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetCertId(v int64) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetCertIdentifier(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertIdentifier = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetCreateTime(v int64) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.CreateTime = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetDomainMatchCert(v bool) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.DomainMatchCert = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetEndTime(v int64) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.EndTime = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetFingerprint(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetInstanceId(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.InstanceId = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetLastTime(v int64) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) SetSignAlgorithm(v string) *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert {
	s.SignAlgorithm = &v
	return s
}

func (s *DescribeVodCertificateListResponseBodyCertificateListModelCertListCert) Validate() error {
	return dara.Validate(s)
}
