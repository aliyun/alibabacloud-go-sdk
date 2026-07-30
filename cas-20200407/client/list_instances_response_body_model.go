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
	// The current page number in the paged query.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The list of instances.
	InstanceList []*ListInstancesResponseBodyInstanceList `json:"InstanceList,omitempty" xml:"InstanceList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 12345678-1234-1234-1234-123456789ABC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of records displayed per page.
	//
	// example:
	//
	// 20
	ShowSize *int32 `json:"ShowSize,omitempty" xml:"ShowSize,omitempty"`
	// The total number of instances.
	//
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
	// Indicates whether automatic managed renewal is enabled. Valid values:
	//
	// - enable: enabled.
	//
	// - disable: disabled.
	//
	// example:
	//
	// enable
	AutoReissue *string `json:"AutoReissue,omitempty" xml:"AutoReissue,omitempty"`
	// The CA brand. Valid values: WoSign, CFCA, DigiCert, GeoTrust, GlobalSign, vTrus, and Alibaba.
	//
	// example:
	//
	// DigiCert
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// The global certificate ID in the format of certificate ID + "-" + site region ID. This ID is commonly used across Alibaba Cloud services.
	//
	// - For the China site: certificate ID + "-cn-hangzhou"
	//
	// - For the China site: certificate ID + "-ap-southeast-1"
	//
	// For example, if the certificate ID is 123, the CertIdentifier for the China site is "123-cn-hangzhou", and the CertIdentifier for the International site is "123-ap-southeast-1".
	//
	// example:
	//
	// 21795675-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The domain name of the latest issued certificate.
	//
	// example:
	//
	// abc.com,www.abc.com
	CertificateDomain *string `json:"CertificateDomain,omitempty" xml:"CertificateDomain,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 18541349
	CertificateId *int64 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The certificate name.
	//
	// example:
	//
	// cert-13216408
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// The end time of the latest certificate. The value is a UNIX timestamp accurate to seconds. If no certificate has been issued, this field is empty.
	//
	// example:
	//
	// 1801324800000
	CertificateNotAfter *int64 `json:"CertificateNotAfter,omitempty" xml:"CertificateNotAfter,omitempty"`
	// The start time of the latest certificate. The value is a UNIX timestamp accurate to seconds. If no certificate has been issued, this field is empty.
	//
	// example:
	//
	// 1776988800000
	CertificateNotBefore *int64 `json:"CertificateNotBefore,omitempty" xml:"CertificateNotBefore,omitempty"`
	// The revocation time of the latest certificate. The value is a UNIX timestamp accurate to seconds.
	//
	// example:
	//
	// 1801324800000
	CertificateRevokeTime *int64 `json:"CertificateRevokeTime,omitempty" xml:"CertificateRevokeTime,omitempty"`
	// The status of the certificate. Valid values:
	//
	// - **issued**: Issued.
	//
	// - **revoked**: Revoked.
	//
	// - **willExpire**: About to expire.
	//
	// - **expired**: Expired.
	//
	// example:
	//
	// issued
	CertificateStatus *string `json:"CertificateStatus,omitempty" xml:"CertificateStatus,omitempty"`
	// The type of the certificate. Valid values: DV, OV, and EV.
	//
	// example:
	//
	// DV
	CertificateType *string `json:"CertificateType,omitempty" xml:"CertificateType,omitempty"`
	// The domain name bound to the certificate.
	//
	// example:
	//
	// test.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The number of exact-match domain names.
	//
	// example:
	//
	// 1
	FullDomainCount *int32 `json:"FullDomainCount,omitempty" xml:"FullDomainCount,omitempty"`
	// The expiration time of the instance. The value is a UNIX timestamp accurate to seconds. If no certificate has been issued, this field is empty.
	//
	// example:
	//
	// 1801324800000
	InstanceEndTime *int64 `json:"InstanceEndTime,omitempty" xml:"InstanceEndTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// cas_dv-cn-123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The start time of the instance. The value is a UNIX timestamp accurate to seconds. If no certificate has been issued, this field is empty.
	//
	// example:
	//
	// 1801324800000
	InstanceStartTime *int64 `json:"InstanceStartTime,omitempty" xml:"InstanceStartTime,omitempty"`
	// The instance type. Valid values:
	//
	// - BUY: official certificate.
	//
	// - TEST: test certificate.
	//
	// example:
	//
	// BUY
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The certificate algorithm. Default value: RSA_2048. Valid values:
	//
	// - **RSA_2048**
	//
	// - **RSA_3072**
	//
	// - **RSA_4096**
	//
	// - **ECC_256**
	//
	// - **SM2**
	//
	// example:
	//
	// RSA_2048
	KeyAlgorithm *string `json:"KeyAlgorithm,omitempty" xml:"KeyAlgorithm,omitempty"`
	// The end time of the instance purchase. The value is a UNIX timestamp accurate to seconds. Used to determine the purchase duration of the instance.
	//
	// example:
	//
	// 1801324800000
	OrderEndTime *int64 `json:"OrderEndTime,omitempty" xml:"OrderEndTime,omitempty"`
	// The start time of the instance purchase. The value is a UNIX timestamp accurate to seconds. Used to determine the refund time limit.
	//
	// example:
	//
	// 1801324800000
	OrderStartTime *int64 `json:"OrderStartTime,omitempty" xml:"OrderStartTime,omitempty"`
	// The result returned by the CA during the last certificate operation.
	//
	// example:
	//
	// pending
	PendingResult *string `json:"PendingResult,omitempty" xml:"PendingResult,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// 123
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The purchased instance specification.
	//
	// example:
	//
	// ss.dv.t
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The instance status. Valid values:
	//
	// - **inactive**: Pending use.
	//
	// - **pending**: Under review. The latest certificate is being reviewed.
	//
	// - **willExpire**: About to expire.
	//
	// - **expired**: Expired.
	//
	// - **refund**: Refunded.
	//
	// - **normal**: Normal.
	//
	// - **closed**: Closed and unavailable.
	//
	// example:
	//
	// inactive
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of cloud services to which the latest certificate is deployed.
	UsingProductList []*string `json:"UsingProductList,omitempty" xml:"UsingProductList,omitempty" type:"Repeated"`
	// The number of wildcard domain names.
	//
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

func (s *ListInstancesResponseBodyInstanceList) GetCertificateDomain() *string {
	return s.CertificateDomain
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

func (s *ListInstancesResponseBodyInstanceList) GetCertificateNotBefore() *int64 {
	return s.CertificateNotBefore
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

func (s *ListInstancesResponseBodyInstanceList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListInstancesResponseBodyInstanceList) GetSpec() *string {
	return s.Spec
}

func (s *ListInstancesResponseBodyInstanceList) GetStatus() *string {
	return s.Status
}

func (s *ListInstancesResponseBodyInstanceList) GetUsingProductList() []*string {
	return s.UsingProductList
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

func (s *ListInstancesResponseBodyInstanceList) SetCertificateDomain(v string) *ListInstancesResponseBodyInstanceList {
	s.CertificateDomain = &v
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

func (s *ListInstancesResponseBodyInstanceList) SetCertificateNotBefore(v int64) *ListInstancesResponseBodyInstanceList {
	s.CertificateNotBefore = &v
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

func (s *ListInstancesResponseBodyInstanceList) SetResourceGroupId(v string) *ListInstancesResponseBodyInstanceList {
	s.ResourceGroupId = &v
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

func (s *ListInstancesResponseBodyInstanceList) SetUsingProductList(v []*string) *ListInstancesResponseBodyInstanceList {
	s.UsingProductList = v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) SetWildcardDomainCount(v int32) *ListInstancesResponseBodyInstanceList {
	s.WildcardDomainCount = &v
	return s
}

func (s *ListInstancesResponseBodyInstanceList) Validate() error {
	return dara.Validate(s)
}
