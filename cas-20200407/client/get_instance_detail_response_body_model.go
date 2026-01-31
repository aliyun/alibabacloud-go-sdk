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
	SetValidationMethod(v string) *GetInstanceDetailResponseBody
	GetValidationMethod() *string
	SetWildcardDomainCount(v int32) *GetInstanceDetailResponseBody
	GetWildcardDomainCount() *int32
}

type GetInstanceDetailResponseBody struct {
	// example:
	//
	// enable
	AutoReissue *string `json:"AutoReissue,omitempty" xml:"AutoReissue,omitempty"`
	// example:
	//
	// 120
	AverageWaitingTime *string `json:"AverageWaitingTime,omitempty" xml:"AverageWaitingTime,omitempty"`
	// example:
	//
	// DigiCert
	Brand *string `json:"Brand,omitempty" xml:"Brand,omitempty"`
	// example:
	//
	// 22783111-cn-hangzhou
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// example:
	//
	// 1234567890
	CertificateId *int32 `json:"CertificateId,omitempty" xml:"CertificateId,omitempty"`
	// example:
	//
	// 123
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
	// Beijing
	City *string `json:"City,omitempty" xml:"City,omitempty"`
	// example:
	//
	// 47305
	CompanyId     *int64   `json:"CompanyId,omitempty" xml:"CompanyId,omitempty"`
	ContactIdList []*int64 `json:"ContactIdList,omitempty" xml:"ContactIdList,omitempty" type:"Repeated"`
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// example:
	//
	// -----BEGIN CERTIFICATE REQUEST-----   ...... -----END CERTIFICATE REQUEST-----
	Csr           *string                                       `json:"Csr,omitempty" xml:"Csr,omitempty"`
	DingGroupList []*GetInstanceDetailResponseBodyDingGroupList `json:"DingGroupList,omitempty" xml:"DingGroupList,omitempty" type:"Repeated"`
	// example:
	//
	// example.com
	Domain               *string                                              `json:"Domain,omitempty" xml:"Domain,omitempty"`
	DomainValidationList []*GetInstanceDetailResponseBodyDomainValidationList `json:"DomainValidationList,omitempty" xml:"DomainValidationList,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	FullDomainCount *int32 `json:"FullDomainCount,omitempty" xml:"FullDomainCount,omitempty"`
	// example:
	//
	// online
	GenerateCsrMethod *string `json:"GenerateCsrMethod,omitempty" xml:"GenerateCsrMethod,omitempty"`
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
	// TEST
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
	// Beijing
	Province *string `json:"Province,omitempty" xml:"Province,omitempty"`
	// example:
	//
	// B2CE1D02-6D5E-56E5-A9BD-EE288255C7F9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-aek****wia
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// ss.dv.t
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// example:
	//
	// inactive
	Status *string                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*GetInstanceDetailResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// DNS
	ValidationMethod *string `json:"ValidationMethod,omitempty" xml:"ValidationMethod,omitempty"`
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
	// example:
	//
	// 123
	DingGroupInstanceId *string `json:"DingGroupInstanceId,omitempty" xml:"DingGroupInstanceId,omitempty"`
	// example:
	//
	// 123
	DingGroupName *string `json:"DingGroupName,omitempty" xml:"DingGroupName,omitempty"`
	// example:
	//
	// remote
	DingGroupType *string `json:"DingGroupType,omitempty" xml:"DingGroupType,omitempty"`
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
	// example:
	//
	// 123.com
	Cname *string `json:"Cname,omitempty" xml:"Cname,omitempty"`
	// example:
	//
	// example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// example.com
	RootDomain *string `json:"RootDomain,omitempty" xml:"RootDomain,omitempty"`
	// example:
	//
	// @
	ValidationKey *string `json:"ValidationKey,omitempty" xml:"ValidationKey,omitempty"`
	// example:
	//
	// TXT
	ValidationType *string `json:"ValidationType,omitempty" xml:"ValidationType,omitempty"`
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
	// example:
	//
	// test
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
