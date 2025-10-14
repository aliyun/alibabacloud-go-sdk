// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *SearchCloudGtmInstancesResponseBodyInstances) *SearchCloudGtmInstancesResponseBody
	GetInstances() *SearchCloudGtmInstancesResponseBodyInstances
	SetPageNumber(v int32) *SearchCloudGtmInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCloudGtmInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchCloudGtmInstancesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *SearchCloudGtmInstancesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *SearchCloudGtmInstancesResponseBody
	GetTotalPages() *int32
}

type SearchCloudGtmInstancesResponseBody struct {
	// The instances.
	Instances *SearchCloudGtmInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
	// Current page number, starting at **1**, default is **1**.
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
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of instances found from the search.
	//
	// example:
	//
	// 10
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s SearchCloudGtmInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstancesResponseBody) GetInstances() *SearchCloudGtmInstancesResponseBodyInstances {
	return s.Instances
}

func (s *SearchCloudGtmInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCloudGtmInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCloudGtmInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchCloudGtmInstancesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *SearchCloudGtmInstancesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *SearchCloudGtmInstancesResponseBody) SetInstances(v *SearchCloudGtmInstancesResponseBodyInstances) *SearchCloudGtmInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *SearchCloudGtmInstancesResponseBody) SetPageNumber(v int32) *SearchCloudGtmInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBody) SetPageSize(v int32) *SearchCloudGtmInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBody) SetRequestId(v string) *SearchCloudGtmInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBody) SetTotalItems(v int32) *SearchCloudGtmInstancesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBody) SetTotalPages(v int32) *SearchCloudGtmInstancesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchCloudGtmInstancesResponseBodyInstances struct {
	Instance []*SearchCloudGtmInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstancesResponseBodyInstances) GetInstance() []*SearchCloudGtmInstancesResponseBodyInstancesInstance {
	return s.Instance
}

func (s *SearchCloudGtmInstancesResponseBodyInstances) SetInstance(v []*SearchCloudGtmInstancesResponseBodyInstancesInstance) *SearchCloudGtmInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchCloudGtmInstancesResponseBodyInstancesInstance struct {
	// example:
	//
	// postpay / prepay
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
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
	// 1710467214858
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// Instance expiration time.
	//
	// example:
	//
	// 2024-09-05T16:00Z
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Instance expiration time (timestamp).
	//
	// example:
	//
	// 1725552000000
	ExpireTimestamp *string `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	// The ID of the GTM 3.0 instance.
	//
	// example:
	//
	// gtm-cn-wwo3a3hbz**
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Schedule instance name.
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
	// 200
	MonthlyEmailUsed *int32 `json:"MonthlyEmailUsed,omitempty" xml:"MonthlyEmailUsed,omitempty"`
	// SMS quota, only supported on the China site. International site does not support SMS.
	//
	// example:
	//
	// 2000
	MonthlySmsQuota *int32 `json:"MonthlySmsQuota,omitempty" xml:"MonthlySmsQuota,omitempty"`
	// Monthly SMS sending volume, only supported by the China site as international sites do not support SMS.
	//
	// example:
	//
	// 200
	MonthlySmsUsed *int32 `json:"MonthlySmsUsed,omitempty" xml:"MonthlySmsUsed,omitempty"`
	// Monthly webhook dispatch volume.
	//
	// example:
	//
	// 100
	MonthlyWebhookUsed *int32 `json:"MonthlyWebhookUsed,omitempty" xml:"MonthlyWebhookUsed,omitempty"`
	// The access domain name, which consists of a hostname and a zone or a subzone.
	//
	// example:
	//
	// www.example.com
	ScheduleDomainName *string `json:"ScheduleDomainName,omitempty" xml:"ScheduleDomainName,omitempty"`
	// The last modified time of the instance.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The last modified time of the instance (timestamp).
	//
	// example:
	//
	// 1710467214858
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// Global Traffic Management version 3.0 instance types:
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

func (s SearchCloudGtmInstancesResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetChargeType() *string {
	return s.ChargeType
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetExpireTime() *int64 {
	return s.ExpireTime
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetExpireTimestamp() *string {
	return s.ExpireTimestamp
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetInstanceName() *string {
	return s.InstanceName
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetMonitorTaskQuota() *int32 {
	return s.MonitorTaskQuota
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetMonthlyEmailUsed() *int32 {
	return s.MonthlyEmailUsed
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetMonthlySmsQuota() *int32 {
	return s.MonthlySmsQuota
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetMonthlySmsUsed() *int32 {
	return s.MonthlySmsUsed
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetMonthlyWebhookUsed() *int32 {
	return s.MonthlyWebhookUsed
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetScheduleDomainName() *string {
	return s.ScheduleDomainName
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) GetVersionCode() *string {
	return s.VersionCode
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetChargeType(v string) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.ChargeType = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetCommodityCode(v string) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.CommodityCode = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetCreateTime(v string) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.CreateTime = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetCreateTimestamp(v int64) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetExpireTime(v int64) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.ExpireTime = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetExpireTimestamp(v string) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.ExpireTimestamp = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetInstanceName(v string) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.InstanceName = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetMonitorTaskQuota(v int32) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.MonitorTaskQuota = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetMonthlyEmailUsed(v int32) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.MonthlyEmailUsed = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetMonthlySmsQuota(v int32) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.MonthlySmsQuota = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetMonthlySmsUsed(v int32) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.MonthlySmsUsed = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetMonthlyWebhookUsed(v int32) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.MonthlyWebhookUsed = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetScheduleDomainName(v string) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.ScheduleDomainName = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetUpdateTime(v string) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.UpdateTime = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetUpdateTimestamp(v int64) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) SetVersionCode(v string) *SearchCloudGtmInstancesResponseBodyInstancesInstance {
	s.VersionCode = &v
	return s
}

func (s *SearchCloudGtmInstancesResponseBodyInstancesInstance) Validate() error {
	return dara.Validate(s)
}
