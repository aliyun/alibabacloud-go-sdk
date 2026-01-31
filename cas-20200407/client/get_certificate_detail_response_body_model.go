// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCertificateDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *GetCertificateDetailResponseBody
	GetAlgorithm() *string
	SetCertIdentifier(v string) *GetCertificateDetailResponseBody
	GetCertIdentifier() *string
	SetCertificateChainList(v []*GetCertificateDetailResponseBodyCertificateChainList) *GetCertificateDetailResponseBody
	GetCertificateChainList() []*GetCertificateDetailResponseBodyCertificateChainList
	SetCertificateId(v int32) *GetCertificateDetailResponseBody
	GetCertificateId() *int32
	SetCertificateName(v string) *GetCertificateDetailResponseBody
	GetCertificateName() *string
	SetCertificateSource(v string) *GetCertificateDetailResponseBody
	GetCertificateSource() *string
	SetCertificateStatus(v string) *GetCertificateDetailResponseBody
	GetCertificateStatus() *string
	SetCommonName(v string) *GetCertificateDetailResponseBody
	GetCommonName() *string
	SetDomain(v string) *GetCertificateDetailResponseBody
	GetDomain() *string
	SetExistPrivateKey(v bool) *GetCertificateDetailResponseBody
	GetExistPrivateKey() *bool
	SetFingerPrint(v string) *GetCertificateDetailResponseBody
	GetFingerPrint() *string
	SetInstanceId(v string) *GetCertificateDetailResponseBody
	GetInstanceId() *string
	SetIssuer(v string) *GetCertificateDetailResponseBody
	GetIssuer() *string
	SetKeySize(v int32) *GetCertificateDetailResponseBody
	GetKeySize() *int32
	SetNotAfter(v int64) *GetCertificateDetailResponseBody
	GetNotAfter() *int64
	SetNotBefore(v int64) *GetCertificateDetailResponseBody
	GetNotBefore() *int64
	SetRequestId(v string) *GetCertificateDetailResponseBody
	GetRequestId() *string
	SetSerial(v string) *GetCertificateDetailResponseBody
	GetSerial() *string
	SetSubjectAlternativeNames(v []*string) *GetCertificateDetailResponseBody
	GetSubjectAlternativeNames() []*string
	SetTags(v []*GetCertificateDetailResponseBodyTags) *GetCertificateDetailResponseBody
	GetTags() []*GetCertificateDetailResponseBodyTags
	SetUsingProductList(v []*string) *GetCertificateDetailResponseBody
	GetUsingProductList() []*string
}

type GetCertificateDetailResponseBody struct {
	// example:
	//
	// RSA
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// example:
	//
	// 21912069-cn-hangzhou
	CertIdentifier       *string                                                 `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	CertificateChainList []*GetCertificateDetailResponseBodyCertificateChainList `json:"CertificateChainList,omitempty" xml:"CertificateChainList,omitempty" type:"Repeated"`
	// example:
	//
	// 22559621
	CertificateId *int32 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// example:
	//
	// 123
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
	// www.example.com
	CommonName *string `json:"CommonName,omitempty" xml:"CommonName,omitempty"`
	// example:
	//
	// aliyundoc.com,example.aliyundoc.com
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
	// cas_dv-cn-123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// Digicert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// 2048
	KeySize *int32 `json:"KeySize,omitempty" xml:"KeySize,omitempty"`
	// example:
	//
	// 17326613180000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// example:
	//
	// 17321613180000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// example:
	//
	// 5979d897-d69f-4fc9-87dd-f3bb73c40b80
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 123
	Serial                  *string                                 `json:"Serial,omitempty" xml:"Serial,omitempty"`
	SubjectAlternativeNames []*string                               `json:"SubjectAlternativeNames,omitempty" xml:"SubjectAlternativeNames,omitempty" type:"Repeated"`
	Tags                    []*GetCertificateDetailResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	UsingProductList        []*string                               `json:"UsingProductList,omitempty" xml:"UsingProductList,omitempty" type:"Repeated"`
}

func (s GetCertificateDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetCertificateDetailResponseBody) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *GetCertificateDetailResponseBody) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *GetCertificateDetailResponseBody) GetCertificateChainList() []*GetCertificateDetailResponseBodyCertificateChainList {
	return s.CertificateChainList
}

func (s *GetCertificateDetailResponseBody) GetCertificateId() *int32 {
	return s.CertificateId
}

func (s *GetCertificateDetailResponseBody) GetCertificateName() *string {
	return s.CertificateName
}

func (s *GetCertificateDetailResponseBody) GetCertificateSource() *string {
	return s.CertificateSource
}

func (s *GetCertificateDetailResponseBody) GetCertificateStatus() *string {
	return s.CertificateStatus
}

func (s *GetCertificateDetailResponseBody) GetCommonName() *string {
	return s.CommonName
}

func (s *GetCertificateDetailResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *GetCertificateDetailResponseBody) GetExistPrivateKey() *bool {
	return s.ExistPrivateKey
}

func (s *GetCertificateDetailResponseBody) GetFingerPrint() *string {
	return s.FingerPrint
}

func (s *GetCertificateDetailResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCertificateDetailResponseBody) GetIssuer() *string {
	return s.Issuer
}

func (s *GetCertificateDetailResponseBody) GetKeySize() *int32 {
	return s.KeySize
}

func (s *GetCertificateDetailResponseBody) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *GetCertificateDetailResponseBody) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *GetCertificateDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCertificateDetailResponseBody) GetSerial() *string {
	return s.Serial
}

func (s *GetCertificateDetailResponseBody) GetSubjectAlternativeNames() []*string {
	return s.SubjectAlternativeNames
}

func (s *GetCertificateDetailResponseBody) GetTags() []*GetCertificateDetailResponseBodyTags {
	return s.Tags
}

func (s *GetCertificateDetailResponseBody) GetUsingProductList() []*string {
	return s.UsingProductList
}

func (s *GetCertificateDetailResponseBody) SetAlgorithm(v string) *GetCertificateDetailResponseBody {
	s.Algorithm = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertIdentifier(v string) *GetCertificateDetailResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertificateChainList(v []*GetCertificateDetailResponseBodyCertificateChainList) *GetCertificateDetailResponseBody {
	s.CertificateChainList = v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertificateId(v int32) *GetCertificateDetailResponseBody {
	s.CertificateId = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertificateName(v string) *GetCertificateDetailResponseBody {
	s.CertificateName = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertificateSource(v string) *GetCertificateDetailResponseBody {
	s.CertificateSource = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCertificateStatus(v string) *GetCertificateDetailResponseBody {
	s.CertificateStatus = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetCommonName(v string) *GetCertificateDetailResponseBody {
	s.CommonName = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetDomain(v string) *GetCertificateDetailResponseBody {
	s.Domain = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetExistPrivateKey(v bool) *GetCertificateDetailResponseBody {
	s.ExistPrivateKey = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetFingerPrint(v string) *GetCertificateDetailResponseBody {
	s.FingerPrint = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetInstanceId(v string) *GetCertificateDetailResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetIssuer(v string) *GetCertificateDetailResponseBody {
	s.Issuer = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetKeySize(v int32) *GetCertificateDetailResponseBody {
	s.KeySize = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetNotAfter(v int64) *GetCertificateDetailResponseBody {
	s.NotAfter = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetNotBefore(v int64) *GetCertificateDetailResponseBody {
	s.NotBefore = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetRequestId(v string) *GetCertificateDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetSerial(v string) *GetCertificateDetailResponseBody {
	s.Serial = &v
	return s
}

func (s *GetCertificateDetailResponseBody) SetSubjectAlternativeNames(v []*string) *GetCertificateDetailResponseBody {
	s.SubjectAlternativeNames = v
	return s
}

func (s *GetCertificateDetailResponseBody) SetTags(v []*GetCertificateDetailResponseBodyTags) *GetCertificateDetailResponseBody {
	s.Tags = v
	return s
}

func (s *GetCertificateDetailResponseBody) SetUsingProductList(v []*string) *GetCertificateDetailResponseBody {
	s.UsingProductList = v
	return s
}

func (s *GetCertificateDetailResponseBody) Validate() error {
	if s.CertificateChainList != nil {
		for _, item := range s.CertificateChainList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCertificateDetailResponseBodyCertificateChainList struct {
	// example:
	//
	// Digicert
	Issuer *string `json:"Issuer,omitempty" xml:"Issuer,omitempty"`
	// example:
	//
	// 17326613180000
	NotAfter *int64 `json:"NotAfter,omitempty" xml:"NotAfter,omitempty"`
	// example:
	//
	// 17321613180000
	NotBefore *int64 `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
	// example:
	//
	// 10
	RemainDay *int32 `json:"RemainDay,omitempty" xml:"RemainDay,omitempty"`
	// example:
	//
	// Digicert
	Subject *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
}

func (s GetCertificateDetailResponseBodyCertificateChainList) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateDetailResponseBodyCertificateChainList) GoString() string {
	return s.String()
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) GetIssuer() *string {
	return s.Issuer
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) GetNotAfter() *int64 {
	return s.NotAfter
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) GetNotBefore() *int64 {
	return s.NotBefore
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) GetRemainDay() *int32 {
	return s.RemainDay
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) GetSubject() *string {
	return s.Subject
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) SetIssuer(v string) *GetCertificateDetailResponseBodyCertificateChainList {
	s.Issuer = &v
	return s
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) SetNotAfter(v int64) *GetCertificateDetailResponseBodyCertificateChainList {
	s.NotAfter = &v
	return s
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) SetNotBefore(v int64) *GetCertificateDetailResponseBodyCertificateChainList {
	s.NotBefore = &v
	return s
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) SetRemainDay(v int32) *GetCertificateDetailResponseBodyCertificateChainList {
	s.RemainDay = &v
	return s
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) SetSubject(v string) *GetCertificateDetailResponseBodyCertificateChainList {
	s.Subject = &v
	return s
}

func (s *GetCertificateDetailResponseBodyCertificateChainList) Validate() error {
	return dara.Validate(s)
}

type GetCertificateDetailResponseBodyTags struct {
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetCertificateDetailResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetCertificateDetailResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetCertificateDetailResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetCertificateDetailResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetCertificateDetailResponseBodyTags) SetTagKey(v string) *GetCertificateDetailResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetCertificateDetailResponseBodyTags) SetTagValue(v string) *GetCertificateDetailResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetCertificateDetailResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
