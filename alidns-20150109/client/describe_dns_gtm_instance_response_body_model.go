// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfig(v *DescribeDnsGtmInstanceResponseBodyConfig) *DescribeDnsGtmInstanceResponseBody
	GetConfig() *DescribeDnsGtmInstanceResponseBodyConfig
	SetCreateTime(v string) *DescribeDnsGtmInstanceResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeDnsGtmInstanceResponseBody
	GetCreateTimestamp() *int64
	SetExpireTime(v string) *DescribeDnsGtmInstanceResponseBody
	GetExpireTime() *string
	SetExpireTimestamp(v int64) *DescribeDnsGtmInstanceResponseBody
	GetExpireTimestamp() *int64
	SetInstanceId(v string) *DescribeDnsGtmInstanceResponseBody
	GetInstanceId() *string
	SetPaymentType(v string) *DescribeDnsGtmInstanceResponseBody
	GetPaymentType() *string
	SetRequestId(v string) *DescribeDnsGtmInstanceResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeDnsGtmInstanceResponseBody
	GetResourceGroupId() *string
	SetSmsQuota(v int32) *DescribeDnsGtmInstanceResponseBody
	GetSmsQuota() *int32
	SetTaskQuota(v int32) *DescribeDnsGtmInstanceResponseBody
	GetTaskQuota() *int32
	SetUsedQuota(v *DescribeDnsGtmInstanceResponseBodyUsedQuota) *DescribeDnsGtmInstanceResponseBody
	GetUsedQuota() *DescribeDnsGtmInstanceResponseBodyUsedQuota
	SetVersionCode(v string) *DescribeDnsGtmInstanceResponseBody
	GetVersionCode() *string
}

type DescribeDnsGtmInstanceResponseBody struct {
	// The configurations of the instance.
	Config *DescribeDnsGtmInstanceResponseBodyConfig `json:"Config,omitempty" xml:"Config,omitempty" type:"Struct"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2020-10-14T06:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The UNIX timestamp that indicates when the instance was created.
	//
	// example:
	//
	// 1602656937000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The time when the instance expires.
	//
	// example:
	//
	// 2020-10-14T06:58Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The UNIX timestamp that indicates when the instance expires.
	//
	// example:
	//
	// 1602656937000
	ExpireTimestamp *int64 `json:"ExpireTimestamp,omitempty" xml:"ExpireTimestamp,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// instanceid1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The billing method. Valid value:
	//
	// 	- Subscription: You can pay in advance for the use of resources.
	//
	// example:
	//
	// Subscription
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 84314904-D047-4176-A0EC-256D7F68C7F5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the instance belongs.
	//
	// example:
	//
	// resourcegroupid1
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The total number of SMS notifications.
	//
	// example:
	//
	// 100
	SmsQuota *int32 `json:"SmsQuota,omitempty" xml:"SmsQuota,omitempty"`
	// The total number of detection tasks.
	//
	// example:
	//
	// 100
	TaskQuota *int32 `json:"TaskQuota,omitempty" xml:"TaskQuota,omitempty"`
	// The used quota.
	UsedQuota *DescribeDnsGtmInstanceResponseBodyUsedQuota `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty" type:"Struct"`
	// The version of the instance.
	//
	// example:
	//
	// versioncode1
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s DescribeDnsGtmInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceResponseBody) GetConfig() *DescribeDnsGtmInstanceResponseBodyConfig {
	return s.Config
}

func (s *DescribeDnsGtmInstanceResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDnsGtmInstanceResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeDnsGtmInstanceResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeDnsGtmInstanceResponseBody) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *DescribeDnsGtmInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmInstanceResponseBody) GetPaymentType() *string {
	return s.PaymentType
}

func (s *DescribeDnsGtmInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmInstanceResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDnsGtmInstanceResponseBody) GetSmsQuota() *int32 {
	return s.SmsQuota
}

func (s *DescribeDnsGtmInstanceResponseBody) GetTaskQuota() *int32 {
	return s.TaskQuota
}

func (s *DescribeDnsGtmInstanceResponseBody) GetUsedQuota() *DescribeDnsGtmInstanceResponseBodyUsedQuota {
	return s.UsedQuota
}

func (s *DescribeDnsGtmInstanceResponseBody) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeDnsGtmInstanceResponseBody) SetConfig(v *DescribeDnsGtmInstanceResponseBodyConfig) *DescribeDnsGtmInstanceResponseBody {
	s.Config = v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetCreateTime(v string) *DescribeDnsGtmInstanceResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetCreateTimestamp(v int64) *DescribeDnsGtmInstanceResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetExpireTime(v string) *DescribeDnsGtmInstanceResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetExpireTimestamp(v int64) *DescribeDnsGtmInstanceResponseBody {
	s.ExpireTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetInstanceId(v string) *DescribeDnsGtmInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetPaymentType(v string) *DescribeDnsGtmInstanceResponseBody {
	s.PaymentType = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetRequestId(v string) *DescribeDnsGtmInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetResourceGroupId(v string) *DescribeDnsGtmInstanceResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetSmsQuota(v int32) *DescribeDnsGtmInstanceResponseBody {
	s.SmsQuota = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetTaskQuota(v int32) *DescribeDnsGtmInstanceResponseBody {
	s.TaskQuota = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetUsedQuota(v *DescribeDnsGtmInstanceResponseBodyUsedQuota) *DescribeDnsGtmInstanceResponseBody {
	s.UsedQuota = v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) SetVersionCode(v string) *DescribeDnsGtmInstanceResponseBody {
	s.VersionCode = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmInstanceResponseBodyConfig struct {
	// The alert notification method.
	AlertConfig *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Struct"`
	// The name of the alert group.
	//
	// example:
	//
	// alertgroup1
	AlertGroup *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The type of the CNAME domain name that is used to access the instance. Valid value:
	//
	// 	- PUBLIC: The CNAME domain name is used to access the instance over the Internet.
	//
	// example:
	//
	// public
	CnameType *string `json:"CnameType,omitempty" xml:"CnameType,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// instancetest1
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The domain name that is used to access the instance over the Internet.
	//
	// example:
	//
	// test.rr.gtm-003.com
	PubicZoneName *string `json:"PubicZoneName,omitempty" xml:"PubicZoneName,omitempty"`
	// Indicates whether a custom CNAME domain name or a CNAME domain name assigned by the system is used to access the instance over the Internet. Valid values:
	//
	// 	- CUSTOM: A custom CNAME domain name is used.
	//
	// 	- SYSTEM_ASSIGN: A CNAME domain name assigned by the system is used.
	//
	// example:
	//
	// custom
	PublicCnameMode *string `json:"PublicCnameMode,omitempty" xml:"PublicCnameMode,omitempty"`
	// The hostname corresponding to the CNAME domain name that is used to access the instance over the Internet.
	//
	// example:
	//
	// test.rr
	PublicRr *string `json:"PublicRr,omitempty" xml:"PublicRr,omitempty"`
	// The service domain name that is used over the Internet.
	//
	// example:
	//
	// example.com
	PublicUserDomainName *string `json:"PublicUserDomainName,omitempty" xml:"PublicUserDomainName,omitempty"`
	// The type of the access policy. Valid values:
	//
	// 	- LATENCY: Latency-based
	//
	// 	- GEO: Geographical location-based
	//
	// example:
	//
	// GEO
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	// The global time to live (TTL).
	//
	// example:
	//
	// 1
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s DescribeDnsGtmInstanceResponseBodyConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceResponseBodyConfig) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) GetAlertConfig() *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig {
	return s.AlertConfig
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) GetAlertGroup() *string {
	return s.AlertGroup
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) GetCnameType() *string {
	return s.CnameType
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) GetPubicZoneName() *string {
	return s.PubicZoneName
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) GetPublicCnameMode() *string {
	return s.PublicCnameMode
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) GetPublicRr() *string {
	return s.PublicRr
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) GetPublicUserDomainName() *string {
	return s.PublicUserDomainName
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) GetStrategyMode() *string {
	return s.StrategyMode
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetAlertConfig(v *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.AlertConfig = v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetAlertGroup(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.AlertGroup = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetCnameType(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.CnameType = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetInstanceName(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.InstanceName = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetPubicZoneName(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.PubicZoneName = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetPublicCnameMode(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.PublicCnameMode = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetPublicRr(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.PublicRr = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetPublicUserDomainName(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.PublicUserDomainName = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetStrategyMode(v string) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.StrategyMode = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) SetTtl(v int32) *DescribeDnsGtmInstanceResponseBodyConfig {
	s.Ttl = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmInstanceResponseBodyConfigAlertConfig struct {
	AlertConfig []*DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) GetAlertConfig() []*DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig {
	return s.AlertConfig
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) SetAlertConfig(v []*DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig {
	s.AlertConfig = v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig struct {
	DingtalkNotice *bool `json:"DingtalkNotice,omitempty" xml:"DingtalkNotice,omitempty"`
	// Indicates whether email notification is configured. Valid values:
	//
	// 	- true: Email notification is configured.
	//
	// 	- false: Email notification is not configured. null: Email notification is not configured.
	//
	// example:
	//
	// true
	EmailNotice *bool `json:"EmailNotice,omitempty" xml:"EmailNotice,omitempty"`
	// The type of the alert event. Valid values:
	//
	// 	- ADDR_ALERT: The address is unavailable.
	//
	// 	- ADDR_RESUME: The address is restored and becomes available.
	//
	// 	- ADDR_POOL_GROUP_UNAVAILABLE: The address pool group is unavailable.
	//
	// 	- ADDR_POOL_GROUP_AVAILABLE: The address pool group is restored and becomes available.
	//
	// 	- ACCESS_STRATEGY_POOL_GROUP_SWITCH: Switchover is triggered between the primary and secondary address pools.
	//
	// 	- MONITOR_NODE_IP_CHANGE: The IP address of the monitoring node has changed.
	//
	// example:
	//
	// ADDR_ALERT
	NoticeType *string `json:"NoticeType,omitempty" xml:"NoticeType,omitempty"`
	// Indicates whether SMS notification is configured. Valid values:
	//
	// 	- true: SMS notification is configured.
	//
	// 	- false: SMS notification is not configured. null: SMS notification is not configured.
	//
	// example:
	//
	// true
	SmsNotice *bool `json:"SmsNotice,omitempty" xml:"SmsNotice,omitempty"`
}

func (s DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) GetDingtalkNotice() *bool {
	return s.DingtalkNotice
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) GetEmailNotice() *bool {
	return s.EmailNotice
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) GetNoticeType() *string {
	return s.NoticeType
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) GetSmsNotice() *bool {
	return s.SmsNotice
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) SetDingtalkNotice(v bool) *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig {
	s.DingtalkNotice = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) SetEmailNotice(v bool) *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig {
	s.EmailNotice = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) SetNoticeType(v string) *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig {
	s.NoticeType = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) SetSmsNotice(v bool) *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig {
	s.SmsNotice = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyConfigAlertConfigAlertConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmInstanceResponseBodyUsedQuota struct {
	DingtalkUsedCount *int32 `json:"DingtalkUsedCount,omitempty" xml:"DingtalkUsedCount,omitempty"`
	// The total number of emails that were sent.
	//
	// example:
	//
	// 123
	EmailUsedCount *int32 `json:"EmailUsedCount,omitempty" xml:"EmailUsedCount,omitempty"`
	// The total number of short messages that were sent.
	//
	// example:
	//
	// 123
	SmsUsedCount *int32 `json:"SmsUsedCount,omitempty" xml:"SmsUsedCount,omitempty"`
	// The number of detection tasks that were created.
	//
	// example:
	//
	// 123
	TaskUsedCount *int32 `json:"TaskUsedCount,omitempty" xml:"TaskUsedCount,omitempty"`
}

func (s DescribeDnsGtmInstanceResponseBodyUsedQuota) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceResponseBodyUsedQuota) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) GetDingtalkUsedCount() *int32 {
	return s.DingtalkUsedCount
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) GetEmailUsedCount() *int32 {
	return s.EmailUsedCount
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) GetSmsUsedCount() *int32 {
	return s.SmsUsedCount
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) GetTaskUsedCount() *int32 {
	return s.TaskUsedCount
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) SetDingtalkUsedCount(v int32) *DescribeDnsGtmInstanceResponseBodyUsedQuota {
	s.DingtalkUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) SetEmailUsedCount(v int32) *DescribeDnsGtmInstanceResponseBodyUsedQuota {
	s.EmailUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) SetSmsUsedCount(v int32) *DescribeDnsGtmInstanceResponseBodyUsedQuota {
	s.SmsUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) SetTaskUsedCount(v int32) *DescribeDnsGtmInstanceResponseBodyUsedQuota {
	s.TaskUsedCount = &v
	return s
}

func (s *DescribeDnsGtmInstanceResponseBodyUsedQuota) Validate() error {
	return dara.Validate(s)
}
