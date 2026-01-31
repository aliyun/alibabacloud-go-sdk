// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *ListInstancesResponseBody
	GetCurrentPage() *int32
	SetInstanceList(v []*ListInstancesResponseBodyInstanceList) *ListInstancesResponseBody
	GetInstanceList() []*ListInstancesResponseBodyInstanceList
	SetRequestId(v string) *ListInstancesResponseBody
	GetRequestId() *string
	SetShowSize(v int32) *ListInstancesResponseBody
	GetShowSize() *int32
	SetTotalCount(v int64) *ListInstancesResponseBody
	GetTotalCount() *int64
}

type ListInstancesResponseBody struct {
	// example:
	//
	// 1
	CurrentPage  *int32                                   `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	InstanceList []*ListInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListInstancesResponseBody) GetInstanceList() []*ListInstancesResponseBodyInstanceList {
	return s.InstanceList
}

func (s *ListInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstancesResponseBody) GetShowSize() *int32 {
	return s.ShowSize
}

func (s *ListInstancesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstancesResponseBody) SetCurrentPage(v int32) *ListInstancesResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *ListInstancesResponseBody) SetInstanceList(v []*ListInstancesResponseBodyInstanceList) *ListInstancesResponseBody {
	s.InstanceList = v
	return s
}

func (s *ListInstancesResponseBody) SetRequestId(v string) *ListInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstancesResponseBody) SetShowSize(v int32) *ListInstancesResponseBody {
	s.ShowSize = &v
	return s
}

func (s *ListInstancesResponseBody) SetTotalCount(v int64) *ListInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstancesResponseBody) Validate() error {
	if s.InstanceList != nil {
		for _, item := range s.InstanceList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstancesResponseBodyInstanceList struct {
	// example:
	//
	// enable
	AutoReissue *string `json:"AutoReissue,omitempty" xml:"AutoReissue,omitempty"`
	// example:
	//
	// DigiCert
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// example:
	//
	// 21795675-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// example:
	//
	// 18541349
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// example:
	//
	// cert-13216408
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// example:
	//
	// 1801324800000
	CertificateNotAfter *int64 `json:"CertificateNotAfter,omitempty" xml:"CertificateNotAfter,omitempty"`
	// example:
	//
	// 1801324800000
	CertificateRevokeTime *int64 `json:"CertificateRevokeTime,omitempty" xml:"CertificateRevokeTime,omitempty"`
	// example:
	//
	// issued
	CertificateStatus *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	// example:
	//
	// DV
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// example:
	//
	// test.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 1
	FullDomainCount *int32 `json:"FullDomainCount,omitempty" xml:"FullDomainCount,omitempty"`
	// example:
	//
	// 1801324800000
	InstanceEndTime *int64 `json:"InstanceEndTime,omitempty" xml:"InstanceEndTime,omitempty"`
	// example:
	//
	// cas_dv-cn-123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 1801324800000
	InstanceStartTime *int64 `json:"InstanceStartTime,omitempty" xml:"InstanceStartTime,omitempty"`
	// example:
	//
	// BUY
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// example:
	//
	// RSA_2048
	KeyAlgorithm *string `json:"KeyAlgorithm,omitempty" xml:"KeyAlgorithm,omitempty"`
	// example:
	//
	// 1801324800000
	OrderEndTime *int64 `json:"OrderEndTime,omitempty" xml:"OrderEndTime,omitempty"`
	// example:
	//
	// 1801324800000
	OrderStartTime *int64 `json:"OrderStartTime,omitempty" xml:"OrderStartTime,omitempty"`
	// example:
	//
	// pending
	PendingResult *string `json:"PendingResult,omitempty" xml:"PendingResult,omitempty"`
	// example:
	//
	// ss.dv.t
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// example:
	//
	// inactive
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// 0
	WildcardDomainCount *int32 `json:"WildcardDomainCount,omitempty" xml:"WildcardDomainCount,omitempty"`
}

func (s ListInstancesResponseBodyInstanceList) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesResponseBodyInstanceList) GoString() string {
	return s.String()
}

func (s *ListInstancesResponseBodyInstanceList) GetAutoReissue() *string {
	return s.AutoReissue
}

func (s *ListInstancesResponseBodyInstanceList) GetBrand() *string {
	return s.Brand
}

func (s *ListInstancesResponseBodyInstanceList) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *ListInstancesResponseBodyInstanceList) GetCertificateId() *int64 {
	return s.CertificateId
}

func (s *ListInstancesResponseBodyInstanceList) GetCertificateName() *string {
	return s.CertificateName
}

func (s *ListInstancesResponseBodyInstanceList) GetCertificateNotAfter() *int64 {
	return s.CertificateNotAfter
}

func (s *ListInstancesResponseBodyInstanceList) GetCertificateRevokeTime() *int64 {
	return s.CertificateRevokeTime
}

func (s *ListInstancesResponseBodyInstanceList) GetCertificateStatus() *string {
	return s.CertificateStatus
}

func (s *ListInstancesResponseBodyInstanceList) GetCertificateType() *string {
	return s.CertificateType
}

func (s *ListInstancesResponseBodyInstanceList) GetDomain() *string {
	return s.Domain
}

func (s *ListInstancesResponseBodyInstanceList) GetFullDomainCount() *int32 {
	return s.FullDomainCount
}

func (s *ListInstancesResponseBodyInstanceList) GetInstanceEndTime() *int64 {
	return s.InstanceEndTime
}

func (s *ListInstancesResponseBodyInstanceList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstancesResponseBodyInstanceList) GetInstanceStartTime() *int64 {
	return s.InstanceStartTime
}

func (s *ListInstancesResponseBodyInstanceList) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ListInstancesResponseBodyInstanceList) GetKeyAlgorithm() *string {
	return s.KeyAlgorithm
}

func (s *ListInstancesResponseBodyInstanceList) GetOrderEndTime() *int64 {
	return s.OrderEndTime
}

func (s *ListInstancesResponseBodyInstanceList) GetOrderStartTime() *int64 {
	return s.OrderStartTime
}

func (s *ListInstancesResponseBodyInstanceList) GetPendingResult() *string {
	return s.PendingResult
}

func (s *ListInstancesResponseBodyInstanceList) GetSpec() *string {
	return s.Spec
}

func (s *ListInstancesResponseBodyInstanceList) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstanceList) GetWildcardDomainCount() *int32 {
	return s.WildcardDomainCount
}

func (s *ListInstancesResponseBodyInstanceList) SetAutoReissue(v string) *ListInstancesResponseBodyInstanceList {
	s.AutoReissue = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetBrand(v string) *ListInstancesResponseBodyInstanceList {
	s.Brand = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetCertIdentifier(v string) *ListInstancesResponseBodyInstanceList {
	s.CertIdentifier = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetCertificateId(v int64) *ListInstancesResponseBodyInstanceList {
	s.CertificateId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetCertificateName(v string) *ListInstancesResponseBodyInstanceList {
	s.CertificateName = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetCertificateNotAfter(v int64) *ListInstancesResponseBodyInstanceList {
	s.CertificateNotAfter = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetCertificateRevokeTime(v int64) *ListInstancesResponseBodyInstanceList {
	s.CertificateRevokeTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetCertificateStatus(v string) *ListInstancesResponseBodyInstanceList {
	s.CertificateStatus = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetCertificateType(v string) *ListInstancesResponseBodyInstanceList {
	s.CertificateType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetDomain(v string) *ListInstancesResponseBodyInstanceList {
	s.Domain = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetFullDomainCount(v int32) *ListInstancesResponseBodyInstanceList {
	s.FullDomainCount = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceEndTime(v int64) *ListInstancesResponseBodyInstanceList {
	s.InstanceEndTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceId(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceId = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceStartTime(v int64) *ListInstancesResponseBodyInstanceList {
	s.InstanceStartTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetInstanceType(v string) *ListInstancesResponseBodyInstanceList {
	s.InstanceType = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetKeyAlgorithm(v string) *ListInstancesResponseBodyInstanceList {
	s.KeyAlgorithm = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetOrderEndTime(v int64) *ListInstancesResponseBodyInstanceList {
	s.OrderEndTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetOrderStartTime(v int64) *ListInstancesResponseBodyInstanceList {
	s.OrderStartTime = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetPendingResult(v string) *ListInstancesResponseBodyInstanceList {
	s.PendingResult = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetSpec(v string) *ListInstancesResponseBodyInstanceList {
	s.Spec = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetStatus(v string) *ListInstancesResponseBodyInstanceList {
	s.Status = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetWildcardDomainCount(v int32) *ListInstancesResponseBodyInstanceList {
	s.WildcardDomainCount = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) Validate() error {
	return dara.Validate(s)
}
