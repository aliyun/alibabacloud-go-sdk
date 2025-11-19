// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodSSLCertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateListModel(v *DescribeVodSSLCertificateListResponseBodyCertificateListModel) *DescribeVodSSLCertificateListResponseBody
	GetCertificateListModel() *DescribeVodSSLCertificateListResponseBodyCertificateListModel
	SetRequestId(v string) *DescribeVodSSLCertificateListResponseBody
	GetRequestId() *string
}

type DescribeVodSSLCertificateListResponseBody struct {
	// The information about certificates.
	CertificateListModel *DescribeVodSSLCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 25818875-5F78-4AF6-D7393642CA58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVodSSLCertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodSSLCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodSSLCertificateListResponseBody) GetCertificateListModel() *DescribeVodSSLCertificateListResponseBodyCertificateListModel {
	return s.CertificateListModel
}

func (s *DescribeVodSSLCertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodSSLCertificateListResponseBody) SetCertificateListModel(v *DescribeVodSSLCertificateListResponseBodyCertificateListModel) *DescribeVodSSLCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBody) SetRequestId(v string) *DescribeVodSSLCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBody) Validate() error {
	if s.CertificateListModel != nil {
		if err := s.CertificateListModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodSSLCertificateListResponseBodyCertificateListModel struct {
	// The list of certificates.
	CertList *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Struct"`
	// The number of certificates that are returned.
	//
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: integers from 1 to 1000.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s DescribeVodSSLCertificateListResponseBodyCertificateListModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodSSLCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModel) GetCertList() *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertList {
	return s.CertList
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModel) GetCount() *int32 {
	return s.Count
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModel) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModel) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModel) SetCertList(v *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertList) *DescribeVodSSLCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeVodSSLCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModel) SetPageNumber(v int64) *DescribeVodSSLCertificateListResponseBodyCertificateListModel {
	s.PageNumber = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModel) SetPageSize(v int64) *DescribeVodSSLCertificateListResponseBodyCertificateListModel {
	s.PageSize = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModel) Validate() error {
	if s.CertList != nil {
		if err := s.CertList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVodSSLCertificateListResponseBodyCertificateListModelCertList struct {
	Cert []*DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert `json:"Cert,omitempty" xml:"Cert,omitempty" type:"Repeated"`
}

func (s DescribeVodSSLCertificateListResponseBodyCertificateListModelCertList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodSSLCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertList) GetCert() []*DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert {
	return s.Cert
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertList) SetCert(v []*DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertList {
	s.Cert = v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertList) Validate() error {
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

type DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert struct {
	// The ID of the certificate.
	//
	// example:
	//
	// 235437
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// The name of the certificate.
	//
	// example:
	//
	// video-ssl
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// The region of the certificate. Valid values: **cn-hangzhou*	- and **ap-southeast-1**. Default value: **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	CertRegion *string `json:"CertRegion,omitempty" xml:"CertRegion,omitempty"`
	// The Common Name (CN) attribute of the certificate. In most cases, the CN is a domain name.
	//
	// example:
	//
	// test
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
	// ****
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// The time when the certificate was last modified. Unit: milliseconds.
	//
	// example:
	//
	// 1512388610
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
}

func (s DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) GoString() string {
	return s.String()
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCertRegion() *string {
	return s.CertRegion
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) GetCommon() *string {
	return s.Common
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCertId(v int64) *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertId = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCertName(v string) *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertName = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCertRegion(v string) *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.CertRegion = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) SetCommon(v string) *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.Common = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) SetFingerprint(v string) *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.Fingerprint = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) SetIssuer(v string) *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.Issuer = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) SetLastTime(v int64) *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert {
	s.LastTime = &v
	return s
}

func (s *DescribeVodSSLCertificateListResponseBodyCertificateListModelCertListCert) Validate() error {
	return dara.Validate(s)
}
