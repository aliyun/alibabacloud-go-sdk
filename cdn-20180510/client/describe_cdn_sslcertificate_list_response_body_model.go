// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnSSLCertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateListModel(v *DescribeCdnSSLCertificateListResponseBodyCertificateListModel) *DescribeCdnSSLCertificateListResponseBody
	GetCertificateListModel() *DescribeCdnSSLCertificateListResponseBodyCertificateListModel
	SetRequestId(v string) *DescribeCdnSSLCertificateListResponseBody
	GetRequestId() *string
}

type DescribeCdnSSLCertificateListResponseBody struct {
	// The list of certificates.
	CertificateListModel *DescribeCdnSSLCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E9D3257A-1B7C-414C-90C1-8D07AC47BCAC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnSSLCertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSSLCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnSSLCertificateListResponseBody) GetCertificateListModel() *DescribeCdnSSLCertificateListResponseBodyCertificateListModel {
	return s.CertificateListModel
}

func (s *DescribeCdnSSLCertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnSSLCertificateListResponseBody) SetCertificateListModel(v *DescribeCdnSSLCertificateListResponseBodyCertificateListModel) *DescribeCdnSSLCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBody) SetRequestId(v string) *DescribeCdnSSLCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBody) Validate() error {
	if s.CertificateListModel != nil {
		if err := s.CertificateListModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnSSLCertificateListResponseBodyCertificateListModel struct {
	// Details about each certificate.
	CertList *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
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

func (s DescribeCdnSSLCertificateListResponseBodyCertificateListModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSSLCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModel) GetCertList() *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertList {
	return s.CertList
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModel) GetCount() *int32 {
	return s.Count
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModel) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertList) *DescribeCdnSSLCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeCdnSSLCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModel) SetPageNumber(v int64) *DescribeCdnSSLCertificateListResponseBodyCertificateListModel {
	s.PageNumber = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModel) SetPageSize(v int64) *DescribeCdnSSLCertificateListResponseBodyCertificateListModel {
	s.PageSize = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModel) Validate() error {
	if s.CertList != nil {
		if err := s.CertList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertList) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertList) GetCert() []*DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	return s.Cert
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertList) Validate() error {
	if s.Cert != nil {
		for _, item := range s.Cert {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 9128192
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
	// 933c6ddee95c9c41a40f9f50493d82be03ad87bf
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
	// 1679896965
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
}

func (s DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCertName() *string {
	return s.CertName
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCommon() *string {
	return s.Common
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCertId(v int64) *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCertRegion(v string) *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertRegion = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetFingerprint(v string) *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) SetLastTime(v int64) *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

func (s *DescribeCdnSSLCertificateListResponseBodyCertificateListModelCertListCert) Validate() error {
	return dara.Validate(s)
}
