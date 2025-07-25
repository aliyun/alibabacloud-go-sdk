// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *ListCloudGtmInstancesResponseBodyInstances) *ListCloudGtmInstancesResponseBody
	GetInstances() *ListCloudGtmInstancesResponseBodyInstances
	SetPageNumber(v int32) *ListCloudGtmInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCloudGtmInstancesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *ListCloudGtmInstancesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListCloudGtmInstancesResponseBody
	GetTotalPages() *int32
}

type ListCloudGtmInstancesResponseBody struct {
	// The instances.
	Instances *ListCloudGtmInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// Current page number, starting with **1**, default is **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of 100 and a default of 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of instance entries.
	//
	// example:
	//
	// 15
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListCloudGtmInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstancesResponseBody) GetInstances() *ListCloudGtmInstancesResponseBodyInstances {
	return s.Instances
}

func (s *ListCloudGtmInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudGtmInstancesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListCloudGtmInstancesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListCloudGtmInstancesResponseBody) SetInstances(v *ListCloudGtmInstancesResponseBodyInstances) *ListCloudGtmInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *ListCloudGtmInstancesResponseBody) SetPageNumber(v int32) *ListCloudGtmInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBody) SetPageSize(v int32) *ListCloudGtmInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBody) SetRequestId(v string) *ListCloudGtmInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBody) SetTotalItems(v int32) *ListCloudGtmInstancesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBody) SetTotalPages(v int32) *ListCloudGtmInstancesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmInstancesResponseBodyInstances struct {
	Instance []*ListCloudGtmInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s ListCloudGtmInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstancesResponseBodyInstances) GetInstance() []*ListCloudGtmInstancesResponseBodyInstancesInstance {
	return s.Instance
}

func (s *ListCloudGtmInstancesResponseBodyInstances) SetInstance(v []*ListCloudGtmInstancesResponseBodyInstancesInstance) *ListCloudGtmInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstances) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmInstancesResponseBodyInstancesInstance struct {
	// The commodity code. Valid values:
	//
	// 	- dns_gtm_public_cn: commodity code on the China site (aliyun.com)
	//
	// 	- dns_gtm_public_intl: commodity code on the international site (alibabacloud.com)
	//
	// example:
	//
	// dns_gtm_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// Instance creation time.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Instance creation time (timestamp).
	//
	// example:
	//
	// 1231298343343
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// Instance expiration time.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Instance expiration time (timestamp).
	//
	// example:
	//
	// 1231298343343
	ExpireTimestamp *int64 `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	// The ID of the GTM instance.
	//
	// example:
	//
	// gtm-cn-jmp3qnw**03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Monitor probe task quota.
	//
	// example:
	//
	// 100
	MonitorTaskQuota *int32 `json:"MonitorTaskQuota,omitempty" xml:"MonitorTaskQuota,omitempty"`
	// Monthly email sending volume.
	//
	// example:
	//
	// 50
	MonthlyEmailUsed *int32 `json:"MonthlyEmailUsed,omitempty" xml:"MonthlyEmailUsed,omitempty"`
	// SMS quota, only supported on the China site as international sites do not support SMS.
	//
	// example:
	//
	// 1000
	MonthlySmsQuota *int32 `json:"MonthlySmsQuota,omitempty" xml:"MonthlySmsQuota,omitempty"`
	// Monthly SMS sending volume, only supported by the China site as international sites do not support SMS.
	//
	// example:
	//
	// 100
	MonthlySmsUsed *int32 `json:"MonthlySmsUsed,omitempty" xml:"MonthlySmsUsed,omitempty"`
	// Monthly webhook send volume.
	//
	// example:
	//
	// 80
	MonthlyWebhookUsed *int32 `json:"MonthlyWebhookUsed,omitempty" xml:"MonthlyWebhookUsed,omitempty"`
	// The access domain name, which consists of a hostname and a zone or a subzone.
	//
	// example:
	//
	// www.example.com
	ScheduleDomainName *string `json:"ScheduleDomainName,omitempty" xml:"ScheduleDomainName,omitempty"`
	// The last time the instance was modified.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The last modification time of the instance (timestamp).
	//
	// example:
	//
	// 1231298343343
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// GTM instance version:
	//
	// - standard: Standard Edition
	//
	// - ultimate: Ultimate Edition
	//
	// example:
	//
	// ultimate
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s ListCloudGtmInstancesResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetMonitorTaskQuota() *int32 {
	return s.MonitorTaskQuota
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetMonthlyEmailUsed() *int32 {
	return s.MonthlyEmailUsed
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetMonthlySmsQuota() *int32 {
	return s.MonthlySmsQuota
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetMonthlySmsUsed() *int32 {
	return s.MonthlySmsUsed
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetMonthlyWebhookUsed() *int32 {
	return s.MonthlyWebhookUsed
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetScheduleDomainName() *string {
	return s.ScheduleDomainName
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) GetVersionCode() *string {
	return s.VersionCode
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetCommodityCode(v string) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.CommodityCode = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetCreateTime(v string) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.CreateTime = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetCreateTimestamp(v int64) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.CreateTimestamp = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetExpireTime(v string) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.ExpireTime = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetExpireTimestamp(v int64) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.ExpireTimestamp = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetMonitorTaskQuota(v int32) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.MonitorTaskQuota = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetMonthlyEmailUsed(v int32) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.MonthlyEmailUsed = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetMonthlySmsQuota(v int32) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.MonthlySmsQuota = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetMonthlySmsUsed(v int32) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.MonthlySmsUsed = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetMonthlyWebhookUsed(v int32) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.MonthlyWebhookUsed = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetScheduleDomainName(v string) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.ScheduleDomainName = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetUpdateTime(v string) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetUpdateTimestamp(v int64) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) SetVersionCode(v string) *ListCloudGtmInstancesResponseBodyInstancesInstance {
	s.VersionCode = &v
	return s
}

func (s *ListCloudGtmInstancesResponseBodyInstancesInstance) Validate() error {
	return dara.Validate(s)
}
