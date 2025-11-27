// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsCertificateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateListModel(v *DescribeVsCertificateListResponseBodyCertificateListModel) *DescribeVsCertificateListResponseBody
	GetCertificateListModel() *DescribeVsCertificateListResponseBodyCertificateListModel
	SetRequestId(v string) *DescribeVsCertificateListResponseBody
	GetRequestId() *string
}

type DescribeVsCertificateListResponseBody struct {
	CertificateListModel *DescribeVsCertificateListResponseBodyCertificateListModel `json:"CertificateListModel,omitempty" xml:"CertificateListModel,omitempty" type:"Struct"`
	// example:
	//
	// 6E310519-E035-51AB-80D4-C1CBECD39EB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeVsCertificateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsCertificateListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateListResponseBody) GetCertificateListModel() *DescribeVsCertificateListResponseBodyCertificateListModel {
	return s.CertificateListModel
}

func (s *DescribeVsCertificateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVsCertificateListResponseBody) SetCertificateListModel(v *DescribeVsCertificateListResponseBodyCertificateListModel) *DescribeVsCertificateListResponseBody {
	s.CertificateListModel = v
	return s
}

func (s *DescribeVsCertificateListResponseBody) SetRequestId(v string) *DescribeVsCertificateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVsCertificateListResponseBody) Validate() error {
	if s.CertificateListModel != nil {
		if err := s.CertificateListModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVsCertificateListResponseBodyCertificateListModel struct {
	CertList []*DescribeVsCertificateListResponseBodyCertificateListModelCertList `json:"CertList,omitempty" xml:"CertList,omitempty" type:"Repeated"`
	// example:
	//
	// 2
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
}

func (s DescribeVsCertificateListResponseBodyCertificateListModel) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsCertificateListResponseBodyCertificateListModel) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModel) GetCertList() []*DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	return s.CertList
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModel) GetCount() *int32 {
	return s.Count
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModel) SetCertList(v []*DescribeVsCertificateListResponseBodyCertificateListModelCertList) *DescribeVsCertificateListResponseBodyCertificateListModel {
	s.CertList = v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModel) SetCount(v int32) *DescribeVsCertificateListResponseBodyCertificateListModel {
	s.Count = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModel) Validate() error {
	if s.CertList != nil {
		for _, item := range s.CertList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVsCertificateListResponseBodyCertificateListModelCertList struct {
	// example:
	//
	// 6338888
	CertId *int64 `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// example:
	//
	// cert-5391062
	CertName *string `json:"CertName,omitempty" xml:"CertName,omitempty"`
	// example:
	//
	// example.aliyundoc.com
	Common *string `json:"Common,omitempty" xml:"Common,omitempty"`
	// example:
	//
	// 3EB2585309AE5C8F369****7CDA6A8F5CEC8B2D4
	Fingerprint *string `json:"Fingerprint,omitempty" xml:"Fingerprint,omitempty"`
	// example:
	//
	// xxxxCert Inc
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// 1632462708
	LastTime *int64 `json:"LastTime,omitempty" xml:"LastTime,omitempty"`
}

func (s DescribeVsCertificateListResponseBodyCertificateListModelCertList) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsCertificateListResponseBodyCertificateListModelCertList) GoString() string {
	return s.String()
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) GetCertId() *int64 {
	return s.CertId
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) GetCertName() *string {
	return s.CertName
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) GetCommon() *string {
	return s.Common
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) GetIssuer() *string {
	return s.Issuer
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) GetLastTime() *int64 {
	return s.LastTime
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetCertId(v int64) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.CertId = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetCertName(v string) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.CertName = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetCommon(v string) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.Common = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetFingerprint(v string) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.Fingerprint = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetIssuer(v string) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.Issuer = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) SetLastTime(v int64) *DescribeVsCertificateListResponseBodyCertificateListModelCertList {
	s.LastTime = &v
	return s
}

func (s *DescribeVsCertificateListResponseBodyCertificateListModelCertList) Validate() error {
	return dara.Validate(s)
}
