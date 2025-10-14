// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGtmInstances(v []*DescribeDnsGtmInstancesResponseBodyGtmInstances) *DescribeDnsGtmInstancesResponseBody
	GetGtmInstances() []*DescribeDnsGtmInstancesResponseBodyGtmInstances
	SetPageNumber(v int32) *DescribeDnsGtmInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDnsGtmInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDnsGtmInstancesResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeDnsGtmInstancesResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeDnsGtmInstancesResponseBody
	GetTotalPages() *int32
}

type DescribeDnsGtmInstancesResponseBody struct {
	// The Global Traffic Manager (GTM) instances.
	GtmInstances []*DescribeDnsGtmInstancesResponseBodyGtmInstances `json:"GtmInstances,omitempty" xml:"GtmInstances,omitempty" type:"Repeated"`
	// The page number. Pages start from page **1**. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 84314904-D047-4176-A0EC-256D7F68C7F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 123
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeDnsGtmInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponseBody) GetGtmInstances() []*DescribeDnsGtmInstancesResponseBodyGtmInstances {
	return s.GtmInstances
}

func (s *DescribeDnsGtmInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDnsGtmInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDnsGtmInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmInstancesResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeDnsGtmInstancesResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeDnsGtmInstancesResponseBody) SetGtmInstances(v []*DescribeDnsGtmInstancesResponseBodyGtmInstances) *DescribeDnsGtmInstancesResponseBody {
	s.GtmInstances = v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBody) SetPageNumber(v int32) *DescribeDnsGtmInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBody) SetPageSize(v int32) *DescribeDnsGtmInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBody) SetRequestId(v string) *DescribeDnsGtmInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBody) SetTotalItems(v int32) *DescribeDnsGtmInstancesResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBody) SetTotalPages(v int32) *DescribeDnsGtmInstancesResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBody) Validate() error {
	if s.GtmInstances != nil {
		for _, item := range s.GtmInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmInstancesResponseBodyGtmInstances struct {
	// The configurations of the instance.
	Config *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The time when the instance was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-14T06:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the instance was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1602658709000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The time when the instance expires. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-10-14T06:58Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The time when the instance expires. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1602658709000
	ExpireTimestamp *int64 `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The billing method of the GTM instance. Valid value:
	//
	// 	- Subscription.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// resourceGroupid123
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The total number of Short Message Service (SMS) notifications.
	//
	// example:
	//
	// 1
	SmsQuota *int32 `json:"SmsQuota,omitempty" xml:"SmsQuota,omitempty"`
	// The total number of detection tasks.
	//
	// example:
	//
	// 1
	TaskQuota *int32 `json:"TaskQuota,omitempty" xml:"TaskQuota,omitempty"`
	// The used quota.
	UsedQuota *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty" type:"Struct"`
	// The version of the instance.
	//
	// example:
	//
	// testVersion1
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstances) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetConfig() *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	return s.Config
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetPaymentType() *string {
	return s.PaymentType
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetSmsQuota() *int32 {
	return s.SmsQuota
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetTaskQuota() *int32 {
	return s.TaskQuota
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetUsedQuota() *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota {
	return s.UsedQuota
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetConfig(v *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.Config = v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetCreateTime(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetCreateTimestamp(v int64) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetExpireTime(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetExpireTimestamp(v int64) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetInstanceId(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetPaymentType(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.PaymentType = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetResourceGroupId(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetSmsQuota(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.SmsQuota = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetTaskQuota(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.TaskQuota = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetUsedQuota(v *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.UsedQuota = v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) SetVersionCode(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstances {
	s.VersionCode = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstances) Validate() error {
	if s.Config != nil {
		if err := s.Config.Validate(); err != nil {
			return err
		}
	}
	if s.UsedQuota != nil {
		if err := s.UsedQuota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig struct {
	// The alert notification method.
	AlertConfig []*DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
	// The alert contact groups. The value is in the JSON format.
	//
	// example:
	//
	// testgroup
	AlertGroup *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The type of the CNAME. Valid value:
	//
	// 	- PUBLIC
	//
	// example:
	//
	// public
	CnameType *string `json:"CnameType,omitempty" xml:"CnameType,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// instanceTest
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Specifies whether to use a custom CNAME or a system-assigned CNAME to access GTM over the Internet. Valid values:
	//
	// 	- CUSTOM: a custom CNAME
	//
	// 	- SYSTEM_ASSIGN: a system-assigned CNAME. You cannot set PublicCnameMode to this value.
	//
	// example:
	//
	// custom
	PublicCnameMode *string `json:"PublicCnameMode,omitempty" xml:"PublicCnameMode,omitempty"`
	// The hostname of the domain name that is used to access GTM over the Internet.
	//
	// example:
	//
	// test.rr
	PublicRr *string `json:"PublicRr,omitempty" xml:"PublicRr,omitempty"`
	// The domain name that is used to access GTM over the Internet.
	//
	// example:
	//
	// example.com
	PublicUserDomainName *string `json:"PublicUserDomainName,omitempty" xml:"PublicUserDomainName,omitempty"`
	// The canonical name (CNAME) that is used to access GTM over the Internet.
	//
	// example:
	//
	// test.rr.gtm-003.com
	PublicZoneName *string `json:"PublicZoneName,omitempty" xml:"PublicZoneName,omitempty"`
	// The type of the access policy. Valid values:
	//
	// 	- LATENCY: latency-based access policy
	//
	// 	- GEO: geographical location-based access policy
	//
	// example:
	//
	// geo
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	// The global time to live (TTL).
	//
	// example:
	//
	// 1
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GetAlertConfig() []*DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig {
	return s.AlertConfig
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GetAlertGroup() *string {
	return s.AlertGroup
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GetCnameType() *string {
	return s.CnameType
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GetPublicCnameMode() *string {
	return s.PublicCnameMode
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GetPublicRr() *string {
	return s.PublicRr
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GetPublicUserDomainName() *string {
	return s.PublicUserDomainName
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GetPublicZoneName() *string {
	return s.PublicZoneName
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GetStrategyMode() *string {
	return s.StrategyMode
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetAlertConfig(v []*DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.AlertConfig = v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetAlertGroup(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.AlertGroup = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetCnameType(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.CnameType = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetInstanceName(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.InstanceName = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetPublicCnameMode(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.PublicCnameMode = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetPublicRr(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.PublicRr = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetPublicUserDomainName(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.PublicUserDomainName = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetPublicZoneName(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.PublicZoneName = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetStrategyMode(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.StrategyMode = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) SetTtl(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig {
	s.Ttl = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfig) Validate() error {
	if s.AlertConfig != nil {
		for _, item := range s.AlertConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig struct {
	// Indicates whether DingTalk alert notifications are configured. Valid values:
	//
	// 	- true
	//
	// 	- false | null
	//
	// example:
	//
	// true
	DingtalkNotice *string `json:"DingtalkNotice,omitempty" xml:"DingtalkNotice,omitempty"`
	// Indicates whether email notifications are configured. Valid values:
	//
	// 	- true
	//
	// 	- false | null
	//
	// example:
	//
	// true
	EmailNotice *string `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
	// The type of the alert event. Valid values:
	//
	// 	- ADDR_ALERT: The address is unavailable.
	//
	// 	- ADDR_RESUME: The address becomes available.
	//
	// 	- ADDR_POOL_GROUP_UNAVAILABLE: The address pool set is unavailable.
	//
	// 	- ADDR_POOL_GROUP_AVAILABLE: The address pool set becomes available.
	//
	// 	- ACCESS_STRATEGY_POOL_GROUP_SWITCH: Switchover is triggered between the primary and secondary address pools.
	//
	// example:
	//
	// ADDR_ALERT
	NoticeType *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// Indicates whether SMS notifications are configured. Valid values:
	//
	// 	- true
	//
	// 	- false | null
	//
	// example:
	//
	// true
	SmsNotice *string `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) GetDingtalkNotice() *string {
	return s.DingtalkNotice
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) GetEmailNotice() *string {
	return s.EmailNotice
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) GetNoticeType() *string {
	return s.NoticeType
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) GetSmsNotice() *string {
	return s.SmsNotice
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) SetDingtalkNotice(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig {
	s.DingtalkNotice = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) SetEmailNotice(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig {
	s.EmailNotice = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) SetNoticeType(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig {
	s.NoticeType = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) SetSmsNotice(v string) *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesConfigAlertConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota struct {
	// The total number of sent DingTalk notifications.
	//
	// example:
	//
	// 100
	DingtalkUsedCount *int32 `json:"DingtalkUsedCount,omitempty" xml:"DingtalkUsedCount,omitempty"`
	// The total number of sent email notifications.
	//
	// example:
	//
	// 100
	EmailUsedCount *int32 `json:"EmailUsedCount,omitempty" xml:"EmailUsedCount,omitempty"`
	// The total number of sent SMS notifications.
	//
	// example:
	//
	// 100
	SmsUsedCount *int32 `json:"SmsUsedCount,omitempty" xml:"SmsUsedCount,omitempty"`
	// The number of created detection tasks.
	//
	// example:
	//
	// 100
	TaskUsedCount *int32 `json:"TaskUsedCount,omitempty" xml:"TaskUsedCount,omitempty"`
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) GetDingtalkUsedCount() *int32 {
	return s.DingtalkUsedCount
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) GetEmailUsedCount() *int32 {
	return s.EmailUsedCount
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) GetSmsUsedCount() *int32 {
	return s.SmsUsedCount
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) GetTaskUsedCount() *int32 {
	return s.TaskUsedCount
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) SetDingtalkUsedCount(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota {
	s.DingtalkUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) SetEmailUsedCount(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota {
	s.EmailUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) SetSmsUsedCount(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota {
	s.SmsUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) SetTaskUsedCount(v int32) *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota {
	s.TaskUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstancesResponseBodyGtmInstancesUsedQuota) Validate() error {
	return dara.Validate(s)
}
