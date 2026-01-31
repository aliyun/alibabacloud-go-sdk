// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCertificatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertificateList(v []*ListCertificatesResponseBodyCertificateList) *ListCertificatesResponseBody
	GetCertificateList() []*ListCertificatesResponseBodyCertificateList
	SetCurrentPage(v int32) *ListCertificatesResponseBody
	GetCurrentPage() *int32
	SetRequestId(v string) *ListCertificatesResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListCertificatesResponseBody
	GetShowSize() *int32
	SetTotalCount(v int64) *ListCertificatesResponseBody
	GetTotalCount() *int64
}

type ListCertificatesResponseBody struct {
	CertificateList []*ListCertificatesResponseBodyCertificateList `json:"CertificateList,omitempty" xml:"CertificateList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCertificatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBody) GetCertificateList() []*ListCertificatesResponseBodyCertificateList {
	return s.CertificateList
}

func (s *ListCertificatesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCertificatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCertificatesResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListCertificatesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListCertificatesResponseBody) SetCertificateList(v []*ListCertificatesResponseBodyCertificateList) *ListCertificatesResponseBody {
	s.CertificateList = v
	return s
}

func (s *ListCertificatesResponseBody) SetCurrentPage(v int32) *ListCertificatesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListCertificatesResponseBody) SetRequestId(v string) *ListCertificatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCertificatesResponseBody) SetShowSize(v int32) *ListCertificatesResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListCertificatesResponseBody) SetTotalCount(v int64) *ListCertificatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCertificatesResponseBody) Validate() error {
	if s.CertificateList != nil {
		for _, item := range s.CertificateList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCertificatesResponseBodyCertificateList struct {
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// 21589515-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// example:
	//
	// 17281539
	CertificateId *string `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// example:
	//
	// test
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// example:
	//
	// BUY
	CertificateSource *string `json:"CertificateSource,omitempty" xml:"CertificateSource,omitempty"`
	// example:
	//
	// issued
	CertificateStatus *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	// example:
	//
	// aliyun.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// test.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// true
	ExistPrivateKey *bool `json:"ExistPrivateKey,omitempty" xml:"ExistPrivateKey,omitempty"`
	// example:
	//
	// 123
	FingerPrint *string `json:"FingerPrint,omitempty" xml:"FingerPrint,omitempty"`
	// example:
	//
	// cas-cn-v***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// DigiCert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// 2048
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// example:
	//
	// 1749580567000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// example:
	//
	// 1760745600000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// example:
	//
	// 123
	Serial                  *string   `json:"Serial,omitempty" xml:"Serial,omitempty"`
	SubjectAlternativeNames []*string `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Repeated"`
	UsingProductList        []*string `json:"UsingProductList,omitempty" xml:"UsingProductList,omitempty" type:"Repeated"`
}

func (s ListCertificatesResponseBodyCertificateList) String() string {
	return dara.Prettify(s)
}

func (s ListCertificatesResponseBodyCertificateList) GoString() string {
	return s.String()
}

func (s *ListCertificatesResponseBodyCertificateList) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *ListCertificatesResponseBodyCertificateList) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *ListCertificatesResponseBodyCertificateList) GetCertificateId() *string {
	return s.CertificateId
}

func (s *ListCertificatesResponseBodyCertificateList) GetCertificateName() *string {
	return s.CertificateName
}

func (s *ListCertificatesResponseBodyCertificateList) GetCertificateSource() *string {
	return s.CertificateSource
}

func (s *ListCertificatesResponseBodyCertificateList) GetCertificateStatus() *string {
	return s.CertificateStatus
}

func (s *ListCertificatesResponseBodyCertificateList) GetCommonName() *string {
	return s.CommonName
}

func (s *ListCertificatesResponseBodyCertificateList) GetDomain() *string {
	return s.Domain
}

func (s *ListCertificatesResponseBodyCertificateList) GetExistPrivateKey() *bool {
	return s.ExistPrivateKey
}

func (s *ListCertificatesResponseBodyCertificateList) GetFingerPrint() *string {
	return s.FingerPrint
}

func (s *ListCertificatesResponseBodyCertificateList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCertificatesResponseBodyCertificateList) GetIssuer() *string {
	return s.Issuer
}

func (s *ListCertificatesResponseBodyCertificateList) GetKeySize() *int32 {
	return s.KeySize
}

func (s *ListCertificatesResponseBodyCertificateList) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *ListCertificatesResponseBodyCertificateList) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *ListCertificatesResponseBodyCertificateList) GetSerial() *string {
	return s.Serial
}

func (s *ListCertificatesResponseBodyCertificateList) GetSubjectAlternativeNames() []*string {
	return s.SubjectAlternativeNames
}

func (s *ListCertificatesResponseBodyCertificateList) GetUsingProductList() []*string {
	return s.UsingProductList
}

func (s *ListCertificatesResponseBodyCertificateList) SetAlgorithm(v string) *ListCertificatesResponseBodyCertificateList {
	s.Algorithm = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCertIdentifier(v string) *ListCertificatesResponseBodyCertificateList {
	s.CertIdentifier = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCertificateId(v string) *ListCertificatesResponseBodyCertificateList {
	s.CertificateId = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCertificateName(v string) *ListCertificatesResponseBodyCertificateList {
	s.CertificateName = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCertificateSource(v string) *ListCertificatesResponseBodyCertificateList {
	s.CertificateSource = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCertificateStatus(v string) *ListCertificatesResponseBodyCertificateList {
	s.CertificateStatus = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetCommonName(v string) *ListCertificatesResponseBodyCertificateList {
	s.CommonName = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetDomain(v string) *ListCertificatesResponseBodyCertificateList {
	s.Domain = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetExistPrivateKey(v bool) *ListCertificatesResponseBodyCertificateList {
	s.ExistPrivateKey = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetFingerPrint(v string) *ListCertificatesResponseBodyCertificateList {
	s.FingerPrint = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetInstanceId(v string) *ListCertificatesResponseBodyCertificateList {
	s.InstanceId = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetIssuer(v string) *ListCertificatesResponseBodyCertificateList {
	s.Issuer = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetKeySize(v int32) *ListCertificatesResponseBodyCertificateList {
	s.KeySize = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetNotAfter(v int64) *ListCertificatesResponseBodyCertificateList {
	s.NotAfter = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetNotBefore(v int64) *ListCertificatesResponseBodyCertificateList {
	s.NotBefore = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetSerial(v string) *ListCertificatesResponseBodyCertificateList {
	s.Serial = &v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetSubjectAlternativeNames(v []*string) *ListCertificatesResponseBodyCertificateList {
	s.SubjectAlternativeNames = v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) SetUsingProductList(v []*string) *ListCertificatesResponseBodyCertificateList {
	s.UsingProductList = v
	return s
}

func (s *ListCertificatesResponseBodyCertificateList) Validate() error {
	return dara.Validate(s)
}
