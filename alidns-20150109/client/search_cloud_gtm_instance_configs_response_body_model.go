// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmInstanceConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceConfigs(v *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs) *SearchCloudGtmInstanceConfigsResponseBody
	GetInstanceConfigs() *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs
	SetPageNumber(v int32) *SearchCloudGtmInstanceConfigsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCloudGtmInstanceConfigsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchCloudGtmInstanceConfigsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *SearchCloudGtmInstanceConfigsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *SearchCloudGtmInstanceConfigsResponseBody
	GetTotalPages() *int32
}

type SearchCloudGtmInstanceConfigsResponseBody struct {
	// The instances list.
	InstanceConfigs *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs `json:"InstanceConfigs,omitempty" xml:"InstanceConfigs,omitempty" type:"Struct"`
	// Current page number, starting from 1, default is 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of rows per page when paginating queries, with a maximum value of **100**, and a default of **20**.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 6AEC7A64-3CB1-4C49-8B35-0B901F1E26BF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of instance configuration entries.
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

func (s SearchCloudGtmInstanceConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstanceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) GetInstanceConfigs() *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs {
	return s.InstanceConfigs
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) SetInstanceConfigs(v *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs) *SearchCloudGtmInstanceConfigsResponseBody {
	s.InstanceConfigs = v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) SetPageNumber(v int32) *SearchCloudGtmInstanceConfigsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) SetPageSize(v int32) *SearchCloudGtmInstanceConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) SetRequestId(v string) *SearchCloudGtmInstanceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) SetTotalItems(v int32) *SearchCloudGtmInstanceConfigsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) SetTotalPages(v int32) *SearchCloudGtmInstanceConfigsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBody) Validate() error {
	if s.InstanceConfigs != nil {
		if err := s.InstanceConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs struct {
	InstanceConfig []*SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig `json:"InstanceConfig,omitempty" xml:"InstanceConfig,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs) GetInstanceConfig() []*SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	return s.InstanceConfig
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs) SetInstanceConfig(v []*SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs {
	s.InstanceConfig = v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigs) Validate() error {
	if s.InstanceConfig != nil {
		for _, item := range s.InstanceConfig {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig struct {
	// The policy for load balancing between address pools. Valid values:
	//
	// 	- round_robin: All address pools are returned for DNS requests from any source. All address pools are sorted in round-robin mode each time they are returned.
	//
	// 	- sequence: The address pool with the smallest sequence number is preferentially returned for DNS requests from any source. The sequence number indicates the priority for returning the address pool. A smaller sequence number indicates a higher priority. If the address pool with the smallest sequence number is unavailable, the address pool with the second smallest sequence number is returned.
	//
	// 	- weight: You can set a different weight value for each address pool. This way, address pools are returned based on the weight values.
	//
	// 	- source_nearest: Different address pools are returned based on the sources of DNS requests. This way, users can access nearby address pools.
	//
	// example:
	//
	// round_robin
	AddressPoolLbStrategy *string `json:"AddressPoolLbStrategy,omitempty" xml:"AddressPoolLbStrategy,omitempty"`
	// The address pools.
	AddressPools *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools `json:"AddressPools,omitempty" xml:"AddressPools,omitempty" type:"Struct"`
	// The availability state of the access domain name. Valid values:
	//
	// 	- available: If the access domain name is **enabled*	- and the health state is **normal**, the access domain name is deemed **available**.
	//
	// 	- unavailable: If the access domain name is **disabled*	- or the health state is **abnormal**, the access domain name is deemed **unavailable**.
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
	// Domain instance creation time.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Domain instance creation time (timestamp).
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
	// 	- ok: The health state of the access domain name is normal and all address pools that are referenced by the access domain name are available.
	//
	// 	- ok_alert: The health state of the access domain name is warning and some of the address pools that are referenced by the access domain name are unavailable. In this case, only the available address pools are returned for DNS requests.
	//
	// 	- exceptional: The health state of the access domain name is abnormal and all address pools that are referenced by the access domain name are unavailable. In this case, addresses in the non-empty address pool with the smallest sequence number are preferentially used for fallback resolution. This returns DNS results for clients as much as possible.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The ID of the GTM 3.0 instance.
	//
	// example:
	//
	// gtm-cn-x0r38e0**03
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Remarks for the domain instance.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The access domain name. The value of this parameter is composed of the value of ScheduleHostname and the value of ScheduleZoneName.
	//
	// example:
	//
	// www.example.com
	ScheduleDomainName *string `json:"ScheduleDomainName,omitempty" xml:"ScheduleDomainName,omitempty"`
	// Host record of the domain accessed by GTM.
	//
	// example:
	//
	// www
	ScheduleHostname *string `json:"ScheduleHostname,omitempty" xml:"ScheduleHostname,omitempty"`
	// DNS record types for the scheduling domain:
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
	// 	- custom: custom allocation. You must specify a custom hostname and associate the hostname with a zone that is hosted by the Public Authoritative DNS module within the account to which the GTM instance belongs to generate an access domain name.
	//
	// 	- sys_assign: system allocation. This mode is not supported. Do not set ScheduleZoneMode to sys_assign.
	//
	// example:
	//
	// custom
	ScheduleZoneMode *string `json:"ScheduleZoneMode,omitempty" xml:"ScheduleZoneMode,omitempty"`
	// The zone such as example.com or subzone such as a.example.com of the access domain name. In most cases, the zone or subzone is hosted by the Public Authoritative DNS module of Alibaba Cloud DNS. This zone belongs to the account to which the GTM instance belongs.
	//
	// example:
	//
	// example.com
	ScheduleZoneName *string `json:"ScheduleZoneName,omitempty" xml:"ScheduleZoneName,omitempty"`
	// The mode used if the address pool with the smallest sequence number is recovered. This parameter is returned when AddressPoolLbStrategy is set to sequence. Valid values:
	//
	// 	- preemptive: The address pool with the smallest sequence number is preferentially used if this address pool is recovered.
	//
	// 	- non_preemptive: The current address pool is still used even if the address pool with the smallest sequence number is recovered.
	//
	// example:
	//
	// preemptive
	SequenceLbStrategyMode *string `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
	// Global TTL (in seconds), the TTL value for resolving the access domain name to the address pool, which affects the caching time of DNS records in the operator\\"s LocalDNS. Supports custom TTL values.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The last modification time of the domain instance.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The last modification time of the domain instance (timestamp).
	//
	// example:
	//
	// 1527690629357
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

func (s SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetAddressPoolLbStrategy() *string {
	return s.AddressPoolLbStrategy
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetAddressPools() *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools {
	return s.AddressPools
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetConfigId() *string {
	return s.ConfigId
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetRemark() *string {
	return s.Remark
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetScheduleDomainName() *string {
	return s.ScheduleDomainName
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetScheduleHostname() *string {
	return s.ScheduleHostname
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetScheduleRrType() *string {
	return s.ScheduleRrType
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetScheduleZoneMode() *string {
	return s.ScheduleZoneMode
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetScheduleZoneName() *string {
	return s.ScheduleZoneName
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetTtl() *int32 {
	return s.Ttl
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetVersionCode() *string {
	return s.VersionCode
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetAddressPoolLbStrategy(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.AddressPoolLbStrategy = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetAddressPools(v *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.AddressPools = v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetAvailableStatus(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.AvailableStatus = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetCommodityCode(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.CommodityCode = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetConfigId(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ConfigId = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetCreateTime(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.CreateTime = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetCreateTimestamp(v int64) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetEnableStatus(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.EnableStatus = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetHealthStatus(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.HealthStatus = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetInstanceId(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.InstanceId = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetRemark(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.Remark = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetScheduleDomainName(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleDomainName = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetScheduleHostname(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleHostname = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetScheduleRrType(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleRrType = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetScheduleZoneMode(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleZoneMode = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetScheduleZoneName(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleZoneName = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetSequenceLbStrategyMode(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetTtl(v int32) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.Ttl = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetUpdateTime(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.UpdateTime = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetUpdateTimestamp(v int64) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetVersionCode(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.VersionCode = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) Validate() error {
	if s.AddressPools != nil {
		if err := s.AddressPools.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools struct {
	AddressPool []*SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool `json:"AddressPool,omitempty" xml:"AddressPool,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) GetAddressPool() []*SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	return s.AddressPool
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) SetAddressPool(v []*SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools {
	s.AddressPool = v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) Validate() error {
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

type SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool struct {
	// Load balancing policy among addresses in the address pool:
	//
	// - round_robin: Round-robin, where for any source of DNS resolution requests, all addresses are returned, with a rotation of the order for every request.
	//
	// - sequence: Sequential, where for any source of DNS resolution requests, the address with the lower sequence number (indicating a higher priority, the smaller the number, the higher the priority) is returned. If the address with the lower sequence number is unavailable, the next address with a lower sequence number is returned.
	//
	// - weight: Weighted, supporting the setting of different weight values for each address to realize returning addresses according to the ratio of weights in DNS query resolutions.
	//
	// - source_nearest: Source-nearest, referring to the intelligent resolution feature, where GTM can return different addresses based on the source of different DNS resolution requests, achieving the effect of users accessing the nearest server.
	//
	// example:
	//
	// round_robin
	AddressLbStrategy *string `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	// Address pool ID, uniquely identifying the address pool.
	//
	// example:
	//
	// pool-89564504435014**60
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
	// Address pool availability status:
	//
	// - available: Available
	//
	// - unavailable: Unavailable
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
	// Address pool status:
	//
	// - enable: Enabled status
	//
	// - disable: Disabled status
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The condition for determining the health status of the address pool. Valid values:
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
	// The health state of the address pool. Valid values:
	//
	// 	- ok: The health state of the address pool is normal and all addresses that are referenced by the address pool are available.
	//
	// 	- ok_alert: The health state of the address pool is warning and some of the addresses that are referenced by the address pool are unavailable. However, the address pool is deemed normal. In this case, only the available addresses are returned for DNS requests.
	//
	// 	- exceptional: The health state of the address pool is abnormal and some or all of the addresses that are referenced by the address pool are unavailable. In this case, the address pool is deemed abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// Parse the request source list.
	RequestSource *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Struct"`
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
	// The mode used if the address with the smallest sequence number is recovered. This parameter is required only when the policy for load balancing between addresses is sequence. Valid values:
	//
	// 	- preemptive: The address with the smallest sequence number is preferentially used if this address is recovered.
	//
	// 	- non_preemptive: The current address is still used even if the address with the smallest sequence number is recovered.
	//
	// example:
	//
	// preemptive
	SequenceLbStrategyMode *string `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
	// Sequence number. For any parsing request, the address pool with the smaller sequence number (indicating the priority of the address pool returned, with smaller numbers having higher priority) will be returned.
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
	// Update time (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// Weight value (an integer between 1 and 100, inclusive), allowing different weight values to be set for each address pool, implementing the return of address pools according to weight ratios in resolution queries.
	//
	// example:
	//
	// 1
	WeightValue *int32 `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
}

func (s SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetAddressLbStrategy() *string {
	return s.AddressLbStrategy
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetAddressPoolType() *string {
	return s.AddressPoolType
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetRequestSource() *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource {
	return s.RequestSource
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetSeqNonPreemptiveSchedule() *bool {
	return s.SeqNonPreemptiveSchedule
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetSerialNumber() *int32 {
	return s.SerialNumber
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetWeightValue() *int32 {
	return s.WeightValue
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetAddressLbStrategy(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.AddressLbStrategy = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetAddressPoolId(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.AddressPoolId = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetAddressPoolName(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.AddressPoolName = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetAddressPoolType(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.AddressPoolType = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetAvailableStatus(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.AvailableStatus = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetCreateTime(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.CreateTime = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetCreateTimestamp(v int64) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetEnableStatus(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.EnableStatus = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetHealthJudgement(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.HealthJudgement = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetHealthStatus(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.HealthStatus = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetRequestSource(v *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.RequestSource = v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetSeqNonPreemptiveSchedule(v bool) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.SeqNonPreemptiveSchedule = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetSequenceLbStrategyMode(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetSerialNumber(v int32) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.SerialNumber = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetUpdateTime(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.UpdateTime = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetUpdateTimestamp(v int64) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetWeightValue(v int32) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.WeightValue = &v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) Validate() error {
	if s.RequestSource != nil {
		if err := s.RequestSource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource struct {
	RequestSource []*string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) GetRequestSource() []*string {
	return s.RequestSource
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) SetRequestSource(v []*string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource {
	s.RequestSource = v
	return s
}

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) Validate() error {
	return dara.Validate(s)
}
