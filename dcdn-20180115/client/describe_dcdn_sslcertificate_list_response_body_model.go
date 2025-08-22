// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnSSLCertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateListModel(v *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) *DescribeDcdnSSLCertificateListResponseBody
	GetCertificateListModel() *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel
	SetRequestId(v string) *DescribeDcdnSSLCertificateListResponseBody
	GetRequestId() *string
}

type DescribeDcdnSSLCertificateListResponseBody struct {
	// The list of certificates.
	CertificateListModel *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-3C82-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDcdnSSLCertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSSLCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSSLCertificateListResponseBody) GetCertificateListModel() *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel {
	return s.CertificateListModel
}

func (s *DescribeDcdnSSLCertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDcdnSSLCertificateListResponseBody) SetCertificateListModel(v *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) *DescribeDcdnSSLCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBody) SetRequestId(v string) *DescribeDcdnSSLCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnSSLCertificateListResponseBodyCertificateListModel struct {
	// Details about each certificate.
	CertList *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
	// The number of certificates.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 2
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: an integer from 1 to 1000.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) GetCertList() *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertList {
	return s.CertList
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) GetCount() *int32 {
	return s.Count
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertList) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) SetPageNumber(v int64) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel {
	s.PageNumber = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) SetPageSize(v int64) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel {
	s.PageSize = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModel) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertList) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertList) GetCert() []*DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	return s.Cert
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertList) Validate() error {
	return dara.Validate(s)
}

type DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 7428244
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// yourCertName
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The region ID of the certificate. Valid values: **cn-hangzhou*	- and **ap-southeast-1**. Default value: **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
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
	// 4278e3b81ab5bc678d253e74c17ffb88
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// The certificate authority (CA) that issued the certificate.
	//
	// example:
	//
	// yourCertIssuer
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The time when the certificate was last modified. Unit: milliseconds.
	//
	// example:
	//
	// 1548065550
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
}

func (s DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCertName() *string {
	return s.CertName
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCommon() *string {
	return s.Common
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCertId(v int64) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCertRegion(v string) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertRegion = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetFingerprint(v string) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetLastTime(v int64) *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

func (s *DescribeDcdnSSLCertificateListResponseBodyCertificateListModelCertListCert) Validate() error {
	return dara.Validate(s)
}
