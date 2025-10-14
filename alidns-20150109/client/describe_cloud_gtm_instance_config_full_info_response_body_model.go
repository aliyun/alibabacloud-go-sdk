// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmInstanceConfigFullInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPoolLbStrategy(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetAddressPoolLbStrategy() *string
	SetAddressPools(v *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetAddressPools() *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools
	SetAlertConfig(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetAlertConfig() *string
	SetAlertGroup(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetAlertGroup() *string
	SetAvailableStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetAvailableStatus() *string
	SetCommodityCode(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetCommodityCode() *string
	SetConfigId(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetConfigId() *string
	SetCreateTime(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetCreateTimestamp() *int64
	SetEnableStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetEnableStatus() *string
	SetHealthStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetHealthStatus() *string
	SetInstanceId(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetInstanceId() *string
	SetInstanceName(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetInstanceName() *string
	SetRemark(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetRequestId() *string
	SetScheduleDomainName(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetScheduleDomainName() *string
	SetScheduleHostname(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetScheduleHostname() *string
	SetScheduleRrType(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetScheduleRrType() *string
	SetScheduleZoneMode(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetScheduleZoneMode() *string
	SetScheduleZoneName(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetScheduleZoneName() *string
	SetSequenceLbStrategyMode(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetSequenceLbStrategyMode() *string
	SetTtl(v int32) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetTtl() *int32
	SetUpdateTime(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetUpdateTimestamp() *int64
	SetVersionCode(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody
	GetVersionCode() *string
}

type DescribeCloudGtmInstanceConfigFullInfoResponseBody struct {
	// The policy for load balancing between address pools. Valid values:
	//
	// 	- round_robin: All address pools are returned for DNS requests from any source. All address pools are sorted in round-robin mode each time they are returned.
	//
	// 	- sequence: The address pool with the smallest sequence number is preferentially returned for DNS requests from any source. The sequence number indicates the priority for returning the address pool. A smaller sequence number indicates a higher priority. If the address pool with the smallest sequence number is unavailable, the address pool with the second smallest sequence number is returned.
	//
	// 	- weight: You can set a different weight value for each address pool. This way, address pools are returned based on the weight values.
	//
	// 	- source_nearest: GTM returns different addresses based on the sources of DNS requests. This way, users can access nearby addresses.
	//
	// example:
	//
	// round_robin
	AddressPoolLbStrategy *string `json:"AddressPoolLbStrategy,omitempty" xml:"AddressPoolLbStrategy,omitempty"`
	// The address pools.
	AddressPools *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools `json:"AddressPools,omitempty" xml:"AddressPools,omitempty" type:"Struct"`
	// Alert notification configuration.
	//
	// example:
	//
	// [{\\"NoticeType\\":\\"addr_alert\\",\\"SmsNotice\\":true,\\"EmailNotice\\":true,\\"DingtalkNotice\\":true},{\\"NoticeType\\":\\"addr_resume\\",\\"SmsNotice\\":true,\\"EmailNotice\\":true,\\"DingtalkNotice\\":true},{\\"NoticeType\\":\\"addr_pool_unavailable\\",\\"SmsNotice\\":true,\\"EmailNotice\\":true,\\"DingtalkNotice\\":true},{\\"NoticeType\\":\\"addr_pool_available\\",\\"SmsNotice\\":true,\\"EmailNotice\\":true,\\"DingtalkNotice\\":true}]"
	AlertConfig *string `json:"AlertConfig,omitempty" xml:"AlertConfig,omitempty"`
	// Alert notification group.
	//
	// example:
	//
	// [\\"Default Contact Group\\"]
	AlertGroup *string `json:"AlertGroup,omitempty" xml:"AlertGroup,omitempty"`
	// The availability state of the access domain name. Valid values:
	//
	// 	- available: If the access domain name is **enabled*	- and the health state of the access domain name is **Normal**, the access domain name is deemed **available**.
	//
	// 	- unavailable: If the access domain name is **disabled*	- or the health state of the access domain name is **Abnormal**, the access domain name is deemed **unavailable**.
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// The commodity code. Valid values:
	//
	// 	- dns_gtm_public_cn: the commodity code on the China site (aliyun.com)
	//
	// 	- dns_gtm_public_intl: the commodity code on the international site (alibabacloud.com)
	//
	// example:
	//
	// dns_gtm_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The configuration ID of the access domain name. Two configuration IDs exist when the access domain name is bound to the same GTM instance but an A record and an AAAA record are configured for the access domain name. The configuration ID uniquely identifies a configuration.
	//
	// example:
	//
	// Config-000**11
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
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
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The enabling state of the access domain name. Valid values:
	//
	// 	- enable: The access domain name is enabled and the intelligent scheduling policy of the corresponding GTM instance takes effect.
	//
	// 	- disable: The access domain name is disabled and the intelligent scheduling policy of the corresponding GTM instance does not take effect.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health state of the access domain name. Valid values:
	//
	// 	- ok: The health state of the access domain name is Normal and all address pools that are referenced by the access domain name are available.
	//
	// 	- ok_alert: The health state of the access domain name is Warning and some of the address pools that are referenced by the access domain name are unavailable. In this case, the available address pools are normally used for DNS resolution, but the unavailable address pools cannot be used for DNS resolution.
	//
	// 	- exceptional: The health state of the access domain name is Abnormal and all address pools that are referenced by the access domain name are unavailable. In this case, addresses in the non-empty address pool with the smallest sequence number are preferentially used for fallback resolution. This returns DNS results for clients as much as possible.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
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
	// Remarks of the configuration of domain instance.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 29D0F8F8-5499-4F6C-9FDC-1EE13BF55925
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The GTM access domain name. The value of this parameter is composed of the value of ScheduleHostname and the value of ScheduleZoneName.
	//
	// example:
	//
	// www.example.com
	ScheduleDomainName *string `json:"ScheduleDomainName,omitempty" xml:"ScheduleDomainName,omitempty"`
	// Host name of the domain accessed by GTM.
	//
	// example:
	//
	// www
	ScheduleHostname *string `json:"ScheduleHostname,omitempty" xml:"ScheduleHostname,omitempty"`
	// DNS record types for the ScheduleDomainName:
	//
	// - A: IPv4 address
	//
	// - AAAA: IPv6 address
	//
	// - CNAME: Domain name
	//
	// example:
	//
	// A
	ScheduleRrType *string `json:"ScheduleRrType,omitempty" xml:"ScheduleRrType,omitempty"`
	// The allocation mode of the access domain name. Valid values:
	//
	// 	- custom: custom allocation. You must specify a custom hostname and associate the hostname with a zone or subzone within the account to which the GTM instance belongs to generate an access domain name.
	//
	// 	- sys_assign: The system assigns an access domain name by default. This mode is no longer supported. Do not choose this mode.
	//
	// example:
	//
	// custom
	ScheduleZoneMode *string `json:"ScheduleZoneMode,omitempty" xml:"ScheduleZoneMode,omitempty"`
	// The zone (such as example.com) or subzone (such as a.example.com) associated with the GTM access domain name. In most cases, the zone or subzone is hosted in Authoritative DNS Resolution of the Alibaba Cloud DNS console within the account to which the GTM instance belongs.
	//
	// example:
	//
	// example.com
	ScheduleZoneName *string `json:"ScheduleZoneName,omitempty" xml:"ScheduleZoneName,omitempty"`
	// The mode used if the address pool with the smallest sequence number is recovered. This parameter is required when AddressPoolLbStrategy is set to sequence. Valid values:
	//
	// 	- preemptive: The address pool with the smallest sequence number is preferentially used if this address pool is recovered.
	//
	// 	- non_preemptive: The current address pool is still used even if the address pool with the smallest sequence number is recovered.
	//
	// example:
	//
	// preemptive
	SequenceLbStrategyMode *string `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
	// Global TTL (in seconds), the TTL value for resolving the access domain to addresses in the address pool, which affects the caching time of DNS records in the ISP\\"s LocalDNS. Custom TTL values are supported.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// Last modified time.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Last modified time (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// Global Traffic Management version 3.0 instances:
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

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetAddressPoolLbStrategy() *string {
	return s.AddressPoolLbStrategy
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetAddressPools() *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools {
	return s.AddressPools
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetAlertConfig() *string {
	return s.AlertConfig
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetAlertGroup() *string {
	return s.AlertGroup
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetScheduleDomainName() *string {
	return s.ScheduleDomainName
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetScheduleHostname() *string {
	return s.ScheduleHostname
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetScheduleRrType() *string {
	return s.ScheduleRrType
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetScheduleZoneMode() *string {
	return s.ScheduleZoneMode
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetScheduleZoneName() *string {
	return s.ScheduleZoneName
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetAddressPoolLbStrategy(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.AddressPoolLbStrategy = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetAddressPools(v *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.AddressPools = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetAlertConfig(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.AlertConfig = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetAlertGroup(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.AlertGroup = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetAvailableStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.AvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetCommodityCode(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetConfigId(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.ConfigId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetCreateTime(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetCreateTimestamp(v int64) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetEnableStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.EnableStatus = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetHealthStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.HealthStatus = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetInstanceId(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetInstanceName(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.InstanceName = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetRemark(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetRequestId(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetScheduleDomainName(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.ScheduleDomainName = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetScheduleHostname(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.ScheduleHostname = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetScheduleRrType(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.ScheduleRrType = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetScheduleZoneMode(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.ScheduleZoneMode = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetScheduleZoneName(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.ScheduleZoneName = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetSequenceLbStrategyMode(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetTtl(v int32) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.Ttl = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetUpdateTime(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetUpdateTimestamp(v int64) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) SetVersionCode(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBody {
	s.VersionCode = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBody) Validate() error {
	if s.AddressPools != nil {
		if err := s.AddressPools.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools struct {
	AddressPool []*DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool `json:"AddressPool,omitempty" xml:"AddressPool,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools) GetAddressPool() []*DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	return s.AddressPool
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools) SetAddressPool(v []*DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools {
	s.AddressPool = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPools) Validate() error {
	if s.AddressPool != nil {
		for _, item := range s.AddressPool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool struct {
	// Load balancing policy among addresses in the address pool:
	//
	// - round_robin: Round-robin, for any source of DNS resolution requests, returns all addresses and rotates their order for each request.
	//
	// - sequence: Sequential, for any source of DNS resolution requests, returns the address with the smaller sequence number (the sequence number indicates the priority of the address return, with smaller numbers having higher priority). If the address with the smaller sequence number is unavailable, the next address with a smaller sequence number is returned.
	//
	// - weight: Weighted, supports setting different weight values for each address to realize returning addresses according to the weight ratio for resolution queries.
	//
	// - source_nearest: Source-nearest, i.e., intelligent resolution function, where GTM can return different addresses based on the source of different DNS resolution requests, achieving the effect of users accessing nearby.
	//
	// example:
	//
	// round_robin
	AddressLbStrategy *string `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	// The ID of the address pool. This ID uniquely identifies the address pool.
	//
	// example:
	//
	// pool-89528023225442**16
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// Address pool name.
	//
	// example:
	//
	// AddressPool-1
	AddressPoolName *string `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	// Address pool type:
	//
	// - IPv4
	//
	// - IPv6
	//
	// - domain
	//
	// example:
	//
	// IPv4
	AddressPoolType *string `json:"AddressPoolType,omitempty" xml:"AddressPoolType,omitempty"`
	// The addresses.
	Addresses *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Struct"`
	// The availability state of the address pool. Valid values:
	//
	// 	- Available
	//
	// 	- unavailable
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// Address pool creation time.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Address pool creation time (timestamp).
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The enabling state of the address pool. Valid values:
	//
	// 	- enable
	//
	// 	- disable
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The condition for determining the health state of the address pool. Valid values:
	//
	// 	- any_ok: At least one address in the address pool is available.
	//
	// 	- p30_ok: At least 30% of the addresses in the address pool are available.
	//
	// 	- p50_ok: At least 50% of the addresses in the address pool are available.
	//
	// 	- p70_ok: At least 70% of the addresses in the address pool are available.
	//
	// 	- all_ok: All addresses in the address pool are available.
	//
	// example:
	//
	// any_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// Address pool health status:
	//
	// - ok: Normal, all addresses referenced by the address pool are available.
	//
	// - ok_alert: Warning, some addresses referenced by the address pool are unavailable, but the address pool status is deemed normal. In the warning state, available address pools are resolved normally, while unavailable ones stop resolving.
	//
	// - exceptional: Abnormal, some or all of the addresses referenced by the address pool are unavailable, and the address pool status is determined to be abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// Parse the list of request sources.
	RequestSource *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolRequestSource `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Struct"`
	// Indicates whether it is a sequential (non-preemptive) scheduling object for hybrid cloud management scenarios:
	//
	// - true: yes
	//
	// - false: no
	//
	// example:
	//
	// false
	SeqNonPreemptiveSchedule *bool `json:"SeqNonPreemptiveSchedule,omitempty" xml:"SeqNonPreemptiveSchedule,omitempty"`
	// The mode used if the address with the smallest sequence number is recovered. This parameter is required only when AddressLbStrategy is set to sequence. Valid values:
	//
	// 	- preemptive: The address with the smallest sequence number is preferentially used if this address is recovered.
	//
	// 	- non_preemptive: The current address is still used even if the address with the smallest sequence number is recovered.
	//
	// example:
	//
	// preemptive
	SequenceLbStrategyMode *string `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
	// Sequence number. For any parsing request from any source, the address pool with the smaller sequence number is returned (the sequence number indicates the priority of the address pool returned, with smaller numbers having higher priority).
	//
	// example:
	//
	// 1
	SerialNumber *int32 `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// Last modification time of the address pool.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Last modification time of the address pool (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// Weight value (an integer between 1 and 100, inclusive), allowing different weight values to be set for each address pool, enabling resolution queries to return address pools according to the weighted ratio.
	//
	// example:
	//
	// 1
	WeightValue *int32 `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetAddressLbStrategy() *string {
	return s.AddressLbStrategy
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetAddressPoolType() *string {
	return s.AddressPoolType
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetAddresses() *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddresses {
	return s.Addresses
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetRequestSource() *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolRequestSource {
	return s.RequestSource
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetSeqNonPreemptiveSchedule() *bool {
	return s.SeqNonPreemptiveSchedule
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetSerialNumber() *int32 {
	return s.SerialNumber
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) GetWeightValue() *int32 {
	return s.WeightValue
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetAddressLbStrategy(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.AddressLbStrategy = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetAddressPoolId(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.AddressPoolId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetAddressPoolName(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.AddressPoolName = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetAddressPoolType(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.AddressPoolType = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetAddresses(v *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddresses) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.Addresses = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetAvailableStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.AvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetCreateTime(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetCreateTimestamp(v int64) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetEnableStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.EnableStatus = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetHealthJudgement(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.HealthJudgement = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetHealthStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.HealthStatus = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetRequestSource(v *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolRequestSource) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.RequestSource = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetSeqNonPreemptiveSchedule(v bool) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.SeqNonPreemptiveSchedule = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetSequenceLbStrategyMode(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetSerialNumber(v int32) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.SerialNumber = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetUpdateTime(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetUpdateTimestamp(v int64) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) SetWeightValue(v int32) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool {
	s.WeightValue = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPool) Validate() error {
	if s.Addresses != nil {
		if err := s.Addresses.Validate(); err != nil {
			return err
		}
	}
	if s.RequestSource != nil {
		if err := s.RequestSource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddresses struct {
	Address []*DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddresses) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddresses) GetAddress() []*DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	return s.Address
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddresses) SetAddress(v []*DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddresses {
	s.Address = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddresses) Validate() error {
	if s.Address != nil {
		for _, item := range s.Address {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress struct {
	// IP address or domain name.
	//
	// example:
	//
	// 223.5.XX.XX
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// The address ID. This ID uniquely identifies the address.
	//
	// example:
	//
	// addr-89564712295703**96
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// Address ownership information, not supported in the current version.
	//
	// example:
	//
	// The current version does not support returning this parameter.
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The failover mode that is used when address exceptions are identified. Valid values:
	//
	// 	- auto: the automatic mode. The system determines whether to return an address based on the health check results. If the address fails health checks, the system does not return the address. If the address passes health checks, the system returns the address.
	//
	// 	- manual: the manual mode. If an address is in the unavailable state, the address is not returned for DNS requests even if the address passes health checks. If an address is in the available state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// example:
	//
	// auto
	AvailableMode *string `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	// The availability state of the address. Valid values:
	//
	// 	- available
	//
	// 	- unavailable
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// Address creation time.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Address creation time (timestamp).
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The enabling state of the address. Valid values:
	//
	// 	- enable
	//
	// 	- disable
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The condition for determining the health state of the address. Valid values:
	//
	// 	- any_ok: The health check results of at least one health check template are normal.
	//
	// 	- p30_ok: The health check results of at least 30% of health check templates are normal.
	//
	// 	- p50_ok: The health check results of at least 50% of health check templates are normal.
	//
	// 	- p70_ok: The health check results of at least 70% of health check templates are normal.
	//
	// 	- all_ok: The health check results of all health check templates are normal.
	//
	// example:
	//
	// any_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// The health check state of the address. Valid values:
	//
	// 	- ok: The address passes all health checks of the referenced health check templates.
	//
	// 	- ok_alert: The address fails some health checks of the referenced health check templates but the address is deemed normal.
	//
	// 	- ok_no_monitor: The address does not reference any health check template and is normal.
	//
	// 	- exceptional: The address fails some or all health checks of the referenced health check templates and the address is deemed abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The availability state of the address when AvailableMode is set to manual. Valid values:
	//
	// 	- available: The address is normal. In this state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// 	- unavailable: The address is abnormal. In this state, the address is not returned for DNS requests even if the address passes health checks.
	//
	// example:
	//
	// available
	ManualAvailableStatus *string `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
	// Address name.
	//
	// example:
	//
	// Address-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The remark of the address.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Request source list.
	RequestSource *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Struct"`
	// Indicates whether it is a sequential (non-preemptive) mode scheduling object, applicable to hybrid cloud management scenarios:
	//
	// - true: yes
	//
	// - false: no
	//
	// example:
	//
	// false
	SeqNonPreemptiveSchedule *bool `json:"SeqNonPreemptiveSchedule,omitempty" xml:"SeqNonPreemptiveSchedule,omitempty"`
	// Sequence number, indicating the priority of address return, where smaller numbers have higher priority.
	//
	// example:
	//
	// 1
	SerialNumber *int32 `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The type of the address. Valid values:
	//
	// 	- IPV4: the IPv4 address
	//
	// 	- IPv6: the IPv6 address
	//
	// 	- domain: the domain name
	//
	// example:
	//
	// IPv4
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Last modified time of the address.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The last modification time of the address (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// Weight value (an integer between 1 and 100, inclusive), allowing different weight values to be set for each address, enabling resolution queries to return addresses in proportion to their weights.
	//
	// example:
	//
	// 1
	WeightValue *int32 `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAddress() *string {
	return s.Address
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAddressId() *string {
	return s.AddressId
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAvailableMode() *string {
	return s.AvailableMode
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetManualAvailableStatus() *string {
	return s.ManualAvailableStatus
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetName() *string {
	return s.Name
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetRequestSource() *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource {
	return s.RequestSource
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetSeqNonPreemptiveSchedule() *bool {
	return s.SeqNonPreemptiveSchedule
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetSerialNumber() *int32 {
	return s.SerialNumber
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetType() *string {
	return s.Type
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) GetWeightValue() *int32 {
	return s.WeightValue
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAddress(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Address = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAddressId(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AddressId = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAttributeInfo(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AttributeInfo = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAvailableMode(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AvailableMode = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAvailableStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetCreateTime(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetCreateTimestamp(v int64) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetEnableStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.EnableStatus = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetHealthJudgement(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.HealthJudgement = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetHealthStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.HealthStatus = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetManualAvailableStatus(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.ManualAvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetName(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Name = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetRemark(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Remark = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetRequestSource(v *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.RequestSource = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetSeqNonPreemptiveSchedule(v bool) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.SeqNonPreemptiveSchedule = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetSerialNumber(v int32) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.SerialNumber = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetType(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Type = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetUpdateTime(v string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetUpdateTimestamp(v int64) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) SetWeightValue(v int32) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.WeightValue = &v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddress) Validate() error {
	if s.RequestSource != nil {
		if err := s.RequestSource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource struct {
	RequestSource []*string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) GetRequestSource() []*string {
	return s.RequestSource
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) SetRequestSource(v []*string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource {
	s.RequestSource = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolRequestSource struct {
	RequestSource []*string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolRequestSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolRequestSource) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolRequestSource) GetRequestSource() []*string {
	return s.RequestSource
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolRequestSource) SetRequestSource(v []*string) *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolRequestSource {
	s.RequestSource = v
	return s
}

func (s *DescribeCloudGtmInstanceConfigFullInfoResponseBodyAddressPoolsAddressPoolRequestSource) Validate() error {
	return dara.Validate(s)
}
