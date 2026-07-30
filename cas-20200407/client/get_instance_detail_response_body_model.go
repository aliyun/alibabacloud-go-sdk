// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoReissue(v string) *GetInstanceDetailResponseBody
	GetAutoReissue() *string
	SetAverageWaitingTime(v string) *GetInstanceDetailResponseBody
	GetAverageWaitingTime() *string
	SetBrand(v string) *GetInstanceDetailResponseBody
	GetBrand() *string
	SetCertIdentifier(v string) *GetInstanceDetailResponseBody
	GetCertIdentifier() *string
	SetCertificateId(v int32) *GetInstanceDetailResponseBody
	GetCertificateId() *int32
	SetCertificateName(v string) *GetInstanceDetailResponseBody
	GetCertificateName() *string
	SetCertificateNotAfter(v int64) *GetInstanceDetailResponseBody
	GetCertificateNotAfter() *int64
	SetCertificateNotBefore(v int64) *GetInstanceDetailResponseBody
	GetCertificateNotBefore() *int64
	SetCertificateRevokeTime(v int64) *GetInstanceDetailResponseBody
	GetCertificateRevokeTime() *int64
	SetCertificateStatus(v string) *GetInstanceDetailResponseBody
	GetCertificateStatus() *string
	SetCertificateType(v string) *GetInstanceDetailResponseBody
	GetCertificateType() *string
	SetCity(v string) *GetInstanceDetailResponseBody
	GetCity() *string
	SetCompanyId(v int64) *GetInstanceDetailResponseBody
	GetCompanyId() *int64
	SetContactIdList(v []*int64) *GetInstanceDetailResponseBody
	GetContactIdList() []*int64
	SetCountryCode(v string) *GetInstanceDetailResponseBody
	GetCountryCode() *string
	SetCsr(v string) *GetInstanceDetailResponseBody
	GetCsr() *string
	SetDingGroupList(v []*GetInstanceDetailResponseBodyDingGroupList) *GetInstanceDetailResponseBody
	GetDingGroupList() []*GetInstanceDetailResponseBodyDingGroupList
	SetDomain(v string) *GetInstanceDetailResponseBody
	GetDomain() *string
	SetDomainValidationList(v []*GetInstanceDetailResponseBodyDomainValidationList) *GetInstanceDetailResponseBody
	GetDomainValidationList() []*GetInstanceDetailResponseBodyDomainValidationList
	SetFullDomainCount(v int32) *GetInstanceDetailResponseBody
	GetFullDomainCount() *int32
	SetGenerateCsrMethod(v string) *GetInstanceDetailResponseBody
	GetGenerateCsrMethod() *string
	SetInstanceEndTime(v int64) *GetInstanceDetailResponseBody
	GetInstanceEndTime() *int64
	SetInstanceId(v string) *GetInstanceDetailResponseBody
	GetInstanceId() *string
	SetInstanceStartTime(v int64) *GetInstanceDetailResponseBody
	GetInstanceStartTime() *int64
	SetInstanceType(v string) *GetInstanceDetailResponseBody
	GetInstanceType() *string
	SetKeyAlgorithm(v string) *GetInstanceDetailResponseBody
	GetKeyAlgorithm() *string
	SetOrderEndTime(v int64) *GetInstanceDetailResponseBody
	GetOrderEndTime() *int64
	SetOrderStartTime(v int64) *GetInstanceDetailResponseBody
	GetOrderStartTime() *int64
	SetPendingResult(v string) *GetInstanceDetailResponseBody
	GetPendingResult() *string
	SetProvince(v string) *GetInstanceDetailResponseBody
	GetProvince() *string
	SetRequestId(v string) *GetInstanceDetailResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *GetInstanceDetailResponseBody
	GetResourceGroupId() *string
	SetSpec(v string) *GetInstanceDetailResponseBody
	GetSpec() *string
	SetStatus(v string) *GetInstanceDetailResponseBody
	GetStatus() *string
	SetTags(v []*GetInstanceDetailResponseBodyTags) *GetInstanceDetailResponseBody
	GetTags() []*GetInstanceDetailResponseBodyTags
	SetUpgradeStatus(v string) *GetInstanceDetailResponseBody
	GetUpgradeStatus() *string
	SetValidationMethod(v string) *GetInstanceDetailResponseBody
	GetValidationMethod() *string
	SetWildcardDomainCount(v int32) *GetInstanceDetailResponseBody
	GetWildcardDomainCount() *int32
}

type GetInstanceDetailResponseBody struct {
	// Indicates whether automatic managed renewal is enabled. Valid values:
	//
	// - enable: Enabled.
	//
	// - disable: Disabled.
	//
	// example:
	//
	// enable
	AutoReissue *string `json:"AutoReissue,omitempty" xml:"AutoReissue,omitempty"`
	// The average waiting time for issuing a certificate of this specification. Unit: seconds.
	//
	// example:
	//
	// 120
	AverageWaitingTime *string `json:"AverageWaitingTime,omitempty" xml:"AverageWaitingTime,omitempty"`
	// The CA brand. Valid values: WoSign, CFCA, DigiCert, GeoTrust, GlobalSign, vTrus, and Alibaba.
	//
	// example:
	//
	// DigiCert
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// The global certificate ID, in the format of certificate ID + "-" + site region ID. This ID is commonly used across Alibaba Cloud services.
	//
	//   --For the China site, the format is certificate ID + "-cn-hangzhou".
	//
	// For the China site, the format is certificate ID + "-ap-southeast-1".
	//
	// For example, if the certificate ID is 123, the CertIdentifier on the China site is "123-cn-hangzhou", and the CertIdentifier on the China site is "123-ap-southeast-1".
	//
	// example:
	//
	// 22783111-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// The certificate ID.
	//
	// example:
	//
	// 1234567890
	CertificateId *int32 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// The name of the instance. When a certificate is issued, this name is used as the default certificate name.
	//
	// example:
	//
	// 123
	CertificateName *string `json:"CertificateName,omitempty" xml:"CertificateName,omitempty"`
	// The end time of the latest certificate, in UNIX timestamp format. This value is empty if no certificate has been issued. The value is accurate to the second.
	//
	// example:
	//
	// 1801324800000
	CertificateNotAfter *int64 `json:"CertificateNotAfter,omitempty" xml:"CertificateNotAfter,omitempty"`
	// The start time of the latest certificate, in UNIX timestamp format. This value is empty if no certificate has been issued. The value is accurate to the second.
	//
	// example:
	//
	// 1781568000000
	CertificateNotBefore *int64 `json:"CertificateNotBefore,omitempty" xml:"CertificateNotBefore,omitempty"`
	// The revocation time of the latest certificate, in UNIX timestamp format. The value is accurate to the second.
	//
	// example:
	//
	// 1801324800000
	CertificateRevokeTime *int64 `json:"CertificateRevokeTime,omitempty" xml:"CertificateRevokeTime,omitempty"`
	// The status of the certificate. Valid values:
	//
	// - **issued**: issued.
	//
	// - **revoked**: revoked.
	//
	// - **willExpire**: about to expire.
	//
	// - **expired**: expired.
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
	// The city where the company or organization of the certificate purchaser is located. This field is required when generating a certificate signing request. Default value: Beijing.
	//
	// example:
	//
	// Beijing
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// The company information ID.
	//
	// example:
	//
	// 47305
	CompanyId *int64 `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	// The list of contact IDs.
	ContactIdList []*int64 `json:"ContactIdList,omitempty" xml:"ContactIdList,omitempty" type:"Repeated"`
	// The country or region code of the certificate organization. For example, CN indicates China, and US indicates the United States. This field is required when generating a certificate signing request. Default value: CN.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The certificate signing request in PEM format.
	//
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----   ...... -----END CERTIFICATE REQUEST-----
	Csr *string `json:"Csr,omitempty" xml:"Csr,omitempty"`
	// The list of associated expert service DingTalk groups.
	DingGroupList []*GetInstanceDetailResponseBodyDingGroupList `json:"DingGroupList,omitempty" xml:"DingGroupList,omitempty" type:"Repeated"`
	// The domain name bound to the certificate.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The list of domain validations.
	DomainValidationList []*GetInstanceDetailResponseBodyDomainValidationList `json:"DomainValidationList,omitempty" xml:"DomainValidationList,omitempty" type:"Repeated"`
	// The number of exact-match domain names.
	//
	// example:
	//
	// 1
	FullDomainCount *int32 `json:"FullDomainCount,omitempty" xml:"FullDomainCount,omitempty"`
	// The CSR generation method. Valid values:
	//
	// - online: system-generated. The Csr field is ignored.
	//
	// - upload: user-uploaded. The Csr field is required.
	//
	// example:
	//
	// online
	GenerateCsrMethod *string `json:"GenerateCsrMethod,omitempty" xml:"GenerateCsrMethod,omitempty"`
	// The expiration time of the instance, in UNIX timestamp format. This value is empty if no certificate has been issued. The value is accurate to the second.
	//
	// example:
	//
	// 1801324800000
	InstanceEndTime *int64 `json:"InstanceEndTime,omitempty" xml:"InstanceEndTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// cas_dv-cn-123
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The start time of the instance, in UNIX timestamp format. This value is empty if no certificate has been issued. The value is accurate to the second.
	//
	// example:
	//
	// 1801324800000
	InstanceStartTime *int64 `json:"InstanceStartTime,omitempty" xml:"InstanceStartTime,omitempty"`
	// The instance type. Valid values:
	//
	// - **BUY**: formal certificate.
	//
	// - **TEST**: test certificate.
	//
	// example:
	//
	// TEST
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The certificate algorithm. Valid values:
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
	// The end time of the instance purchase, in UNIX timestamp format. This value is used to determine the purchase duration of the instance.
	//
	// example:
	//
	// 1801324800000
	OrderEndTime *int64 `json:"OrderEndTime,omitempty" xml:"OrderEndTime,omitempty"`
	// The start time of the instance purchase, in UNIX timestamp format. This value is used to determine the refund time limit. The value is accurate to the second.
	//
	// example:
	//
	// 1801324800000
	OrderStartTime *int64 `json:"OrderStartTime,omitempty" xml:"OrderStartTime,omitempty"`
	// The result returned by the certification authority (CA) during the last certificate operation.
	//
	// example:
	//
	// pending
	PendingResult *string `json:"PendingResult,omitempty" xml:"PendingResult,omitempty"`
	// The province or region where the company is located. This field is required when generating a certificate signing request. Default value: Beijing.
	//
	// example:
	//
	// Beijing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// The request ID. Alibaba Cloud generates a unique identifier for each request. You can use the request ID to troubleshoot issues.
	//
	// example:
	//
	// B2CE1D02-6D5E-56E5-A9BD-EE288255C7F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek****wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The purchased instance specification.
	//
	// example:
	//
	// ss.dv.t
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// The instance status. Valid values:
	//
	// - **inactive**: pending use.
	//
	// - **pending**: under review. The latest certificate is being reviewed.
	//
	// - **willExpire**: the instance is about to expire.
	//
	// - **expired**: the instance has expired.
	//
	// - **refund**: refunded.
	//
	// - **normal**: normal.
	//
	// - **closed**: closed and unavailable.
	//
	// example:
	//
	// inactive
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags.
	Tags []*GetInstanceDetailResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The upgrade status of the instance. Valid values:
	//
	// - none: the instance has not been upgraded.
	//
	// - payed: the instance upgrade has been paid.
	//
	// - issued: the latest certificate has been issued after the instance upgrade.
	//
	// example:
	//
	// none
	UpgradeStatus *string `json:"UpgradeStatus,omitempty" xml:"UpgradeStatus,omitempty"`
	// The certificate validation method. Valid values:
	//
	// - DNS: DNS validation, using TXT or CNAME.
	//
	// - HTTP: file-based validation.
	//
	// example:
	//
	// DNS
	ValidationMethod *string `json:"ValidationMethod,omitempty" xml:"ValidationMethod,omitempty"`
	// The number of wildcard domain names.
	//
	// example:
	//
	// 0
	WildcardDomainCount *int32 `json:"WildcardDomainCount,omitempty" xml:"WildcardDomainCount,omitempty"`
}

func (s GetInstanceDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBody) GetAutoReissue() *string {
	return s.AutoReissue
}

func (s *GetInstanceDetailResponseBody) GetAverageWaitingTime() *string {
	return s.AverageWaitingTime
}

func (s *GetInstanceDetailResponseBody) GetBrand() *string {
	return s.Brand
}

func (s *GetInstanceDetailResponseBody) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *GetInstanceDetailResponseBody) GetCertificateId() *int32 {
	return s.CertificateId
}

func (s *GetInstanceDetailResponseBody) GetCertificateName() *string {
	return s.CertificateName
}

func (s *GetInstanceDetailResponseBody) GetCertificateNotAfter() *int64 {
	return s.CertificateNotAfter
}

func (s *GetInstanceDetailResponseBody) GetCertificateNotBefore() *int64 {
	return s.CertificateNotBefore
}

func (s *GetInstanceDetailResponseBody) GetCertificateRevokeTime() *int64 {
	return s.CertificateRevokeTime
}

func (s *GetInstanceDetailResponseBody) GetCertificateStatus() *string {
	return s.CertificateStatus
}

func (s *GetInstanceDetailResponseBody) GetCertificateType() *string {
	return s.CertificateType
}

func (s *GetInstanceDetailResponseBody) GetCity() *string {
	return s.City
}

func (s *GetInstanceDetailResponseBody) GetCompanyId() *int64 {
	return s.CompanyId
}

func (s *GetInstanceDetailResponseBody) GetContactIdList() []*int64 {
	return s.ContactIdList
}

func (s *GetInstanceDetailResponseBody) GetCountryCode() *string {
	return s.CountryCode
}

func (s *GetInstanceDetailResponseBody) GetCsr() *string {
	return s.Csr
}

func (s *GetInstanceDetailResponseBody) GetDingGroupList() []*GetInstanceDetailResponseBodyDingGroupList {
	return s.DingGroupList
}

func (s *GetInstanceDetailResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *GetInstanceDetailResponseBody) GetDomainValidationList() []*GetInstanceDetailResponseBodyDomainValidationList {
	return s.DomainValidationList
}

func (s *GetInstanceDetailResponseBody) GetFullDomainCount() *int32 {
	return s.FullDomainCount
}

func (s *GetInstanceDetailResponseBody) GetGenerateCsrMethod() *string {
	return s.GenerateCsrMethod
}

func (s *GetInstanceDetailResponseBody) GetInstanceEndTime() *int64 {
	return s.InstanceEndTime
}

func (s *GetInstanceDetailResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceDetailResponseBody) GetInstanceStartTime() *int64 {
	return s.InstanceStartTime
}

func (s *GetInstanceDetailResponseBody) GetInstanceType() *string {
	return s.InstanceType
}

func (s *GetInstanceDetailResponseBody) GetKeyAlgorithm() *string {
	return s.KeyAlgorithm
}

func (s *GetInstanceDetailResponseBody) GetOrderEndTime() *int64 {
	return s.OrderEndTime
}

func (s *GetInstanceDetailResponseBody) GetOrderStartTime() *int64 {
	return s.OrderStartTime
}

func (s *GetInstanceDetailResponseBody) GetPendingResult() *string {
	return s.PendingResult
}

func (s *GetInstanceDetailResponseBody) GetProvince() *string {
	return s.Province
}

func (s *GetInstanceDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceDetailResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetInstanceDetailResponseBody) GetSpec() *string {
	return s.Spec
}

func (s *GetInstanceDetailResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceDetailResponseBody) GetTags() []*GetInstanceDetailResponseBodyTags {
	return s.Tags
}

func (s *GetInstanceDetailResponseBody) GetUpgradeStatus() *string {
	return s.UpgradeStatus
}

func (s *GetInstanceDetailResponseBody) GetValidationMethod() *string {
	return s.ValidationMethod
}

func (s *GetInstanceDetailResponseBody) GetWildcardDomainCount() *int32 {
	return s.WildcardDomainCount
}

func (s *GetInstanceDetailResponseBody) SetAutoReissue(v string) *GetInstanceDetailResponseBody {
	s.AutoReissue = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetAverageWaitingTime(v string) *GetInstanceDetailResponseBody {
	s.AverageWaitingTime = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetBrand(v string) *GetInstanceDetailResponseBody {
	s.Brand = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCertIdentifier(v string) *GetInstanceDetailResponseBody {
	s.CertIdentifier = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCertificateId(v int32) *GetInstanceDetailResponseBody {
	s.CertificateId = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCertificateName(v string) *GetInstanceDetailResponseBody {
	s.CertificateName = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCertificateNotAfter(v int64) *GetInstanceDetailResponseBody {
	s.CertificateNotAfter = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCertificateNotBefore(v int64) *GetInstanceDetailResponseBody {
	s.CertificateNotBefore = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCertificateRevokeTime(v int64) *GetInstanceDetailResponseBody {
	s.CertificateRevokeTime = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCertificateStatus(v string) *GetInstanceDetailResponseBody {
	s.CertificateStatus = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCertificateType(v string) *GetInstanceDetailResponseBody {
	s.CertificateType = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCity(v string) *GetInstanceDetailResponseBody {
	s.City = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCompanyId(v int64) *GetInstanceDetailResponseBody {
	s.CompanyId = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetContactIdList(v []*int64) *GetInstanceDetailResponseBody {
	s.ContactIdList = v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCountryCode(v string) *GetInstanceDetailResponseBody {
	s.CountryCode = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetCsr(v string) *GetInstanceDetailResponseBody {
	s.Csr = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetDingGroupList(v []*GetInstanceDetailResponseBodyDingGroupList) *GetInstanceDetailResponseBody {
	s.DingGroupList = v
	return s
}

func (s *GetInstanceDetailResponseBody) SetDomain(v string) *GetInstanceDetailResponseBody {
	s.Domain = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetDomainValidationList(v []*GetInstanceDetailResponseBodyDomainValidationList) *GetInstanceDetailResponseBody {
	s.DomainValidationList = v
	return s
}

func (s *GetInstanceDetailResponseBody) SetFullDomainCount(v int32) *GetInstanceDetailResponseBody {
	s.FullDomainCount = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetGenerateCsrMethod(v string) *GetInstanceDetailResponseBody {
	s.GenerateCsrMethod = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetInstanceEndTime(v int64) *GetInstanceDetailResponseBody {
	s.InstanceEndTime = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetInstanceId(v string) *GetInstanceDetailResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetInstanceStartTime(v int64) *GetInstanceDetailResponseBody {
	s.InstanceStartTime = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetInstanceType(v string) *GetInstanceDetailResponseBody {
	s.InstanceType = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetKeyAlgorithm(v string) *GetInstanceDetailResponseBody {
	s.KeyAlgorithm = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetOrderEndTime(v int64) *GetInstanceDetailResponseBody {
	s.OrderEndTime = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetOrderStartTime(v int64) *GetInstanceDetailResponseBody {
	s.OrderStartTime = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetPendingResult(v string) *GetInstanceDetailResponseBody {
	s.PendingResult = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetProvince(v string) *GetInstanceDetailResponseBody {
	s.Province = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetRequestId(v string) *GetInstanceDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetResourceGroupId(v string) *GetInstanceDetailResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetSpec(v string) *GetInstanceDetailResponseBody {
	s.Spec = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetStatus(v string) *GetInstanceDetailResponseBody {
	s.Status = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetTags(v []*GetInstanceDetailResponseBodyTags) *GetInstanceDetailResponseBody {
	s.Tags = v
	return s
}

func (s *GetInstanceDetailResponseBody) SetUpgradeStatus(v string) *GetInstanceDetailResponseBody {
	s.UpgradeStatus = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetValidationMethod(v string) *GetInstanceDetailResponseBody {
	s.ValidationMethod = &v
	return s
}

func (s *GetInstanceDetailResponseBody) SetWildcardDomainCount(v int32) *GetInstanceDetailResponseBody {
	s.WildcardDomainCount = &v
	return s
}

func (s *GetInstanceDetailResponseBody) Validate() error {
	if s.DingGroupList != nil {
		for _, item := range s.DingGroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DomainValidationList != nil {
		for _, item := range s.DomainValidationList {
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

type GetInstanceDetailResponseBodyDingGroupList struct {
	// The instance ID of the expert service DingTalk group.
	//
	// example:
	//
	// 123
	DingGroupInstanceId *string `json:"DingGroupInstanceId,omitempty" xml:"DingGroupInstanceId,omitempty"`
	// The name of the expert service DingTalk group.
	//
	// example:
	//
	// 123
	DingGroupName *string `json:"DingGroupName,omitempty" xml:"DingGroupName,omitempty"`
	// The type of the expert service DingTalk group. Valid values:
	//
	// - expedite: application assistance.
	//
	// - remote: offline deployment.
	//
	// example:
	//
	// remote
	DingGroupType *string `json:"DingGroupType,omitempty" xml:"DingGroupType,omitempty"`
	// The URL for joining the expert service DingTalk group.
	//
	// example:
	//
	// https://123.com
	DingGroupUrl *string `json:"DingGroupUrl,omitempty" xml:"DingGroupUrl,omitempty"`
}

func (s GetInstanceDetailResponseBodyDingGroupList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponseBodyDingGroupList) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBodyDingGroupList) GetDingGroupInstanceId() *string {
	return s.DingGroupInstanceId
}

func (s *GetInstanceDetailResponseBodyDingGroupList) GetDingGroupName() *string {
	return s.DingGroupName
}

func (s *GetInstanceDetailResponseBodyDingGroupList) GetDingGroupType() *string {
	return s.DingGroupType
}

func (s *GetInstanceDetailResponseBodyDingGroupList) GetDingGroupUrl() *string {
	return s.DingGroupUrl
}

func (s *GetInstanceDetailResponseBodyDingGroupList) SetDingGroupInstanceId(v string) *GetInstanceDetailResponseBodyDingGroupList {
	s.DingGroupInstanceId = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDingGroupList) SetDingGroupName(v string) *GetInstanceDetailResponseBodyDingGroupList {
	s.DingGroupName = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDingGroupList) SetDingGroupType(v string) *GetInstanceDetailResponseBodyDingGroupList {
	s.DingGroupType = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDingGroupList) SetDingGroupUrl(v string) *GetInstanceDetailResponseBodyDingGroupList {
	s.DingGroupUrl = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDingGroupList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceDetailResponseBodyDomainValidationList struct {
	// The CNAME record value for verification-free authorization. This value may be empty.
	//
	// example:
	//
	// 123.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// The prefix for CNAME validation.
	//
	// example:
	//
	// abc
	CnameKey *string `json:"CnameKey,omitempty" xml:"CnameKey,omitempty"`
	// The domain name to be validated.
	//
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The root domain name.
	//
	// example:
	//
	// example.com
	RootDomain *string `json:"RootDomain,omitempty" xml:"RootDomain,omitempty"`
	// The host record.
	//
	// example:
	//
	// @
	ValidationKey *string `json:"ValidationKey,omitempty" xml:"ValidationKey,omitempty"`
	// The validation type. Valid values: TXT, HTTP, and CNAME.
	//
	// example:
	//
	// TXT
	ValidationType *string `json:"ValidationType,omitempty" xml:"ValidationType,omitempty"`
	// The validation host record value.
	//
	// example:
	//
	// 123
	ValidationValue *string `json:"ValidationValue,omitempty" xml:"ValidationValue,omitempty"`
}

func (s GetInstanceDetailResponseBodyDomainValidationList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponseBodyDomainValidationList) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) GetCname() *string {
	return s.Cname
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) GetCnameKey() *string {
	return s.CnameKey
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) GetDomain() *string {
	return s.Domain
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) GetRootDomain() *string {
	return s.RootDomain
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) GetValidationKey() *string {
	return s.ValidationKey
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) GetValidationType() *string {
	return s.ValidationType
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) GetValidationValue() *string {
	return s.ValidationValue
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) SetCname(v string) *GetInstanceDetailResponseBodyDomainValidationList {
	s.Cname = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) SetCnameKey(v string) *GetInstanceDetailResponseBodyDomainValidationList {
	s.CnameKey = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) SetDomain(v string) *GetInstanceDetailResponseBodyDomainValidationList {
	s.Domain = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) SetRootDomain(v string) *GetInstanceDetailResponseBodyDomainValidationList {
	s.RootDomain = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) SetValidationKey(v string) *GetInstanceDetailResponseBodyDomainValidationList {
	s.ValidationKey = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) SetValidationType(v string) *GetInstanceDetailResponseBodyDomainValidationList {
	s.ValidationType = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) SetValidationValue(v string) *GetInstanceDetailResponseBodyDomainValidationList {
	s.ValidationValue = &v
	return s
}

func (s *GetInstanceDetailResponseBodyDomainValidationList) Validate() error {
	return dara.Validate(s)
}

type GetInstanceDetailResponseBodyTags struct {
	// The tag key.
	//
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value.
	//
	// example:
	//
	// test
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetInstanceDetailResponseBodyTags) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailResponseBodyTags) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailResponseBodyTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetInstanceDetailResponseBodyTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetInstanceDetailResponseBodyTags) SetTagKey(v string) *GetInstanceDetailResponseBodyTags {
	s.TagKey = &v
	return s
}

func (s *GetInstanceDetailResponseBodyTags) SetTagValue(v string) *GetInstanceDetailResponseBodyTags {
	s.TagValue = &v
	return s
}

func (s *GetInstanceDetailResponseBodyTags) Validate() error {
	return dara.Validate(s)
}
