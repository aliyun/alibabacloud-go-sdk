// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmInstanceConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceConfigs(v *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs) *ListCloudGtmInstanceConfigsResponseBody
	GetInstanceConfigs() *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs
	SetPageNumber(v int32) *ListCloudGtmInstanceConfigsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmInstanceConfigsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCloudGtmInstanceConfigsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *ListCloudGtmInstanceConfigsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListCloudGtmInstanceConfigsResponseBody
	GetTotalPages() *int32
}

type ListCloudGtmInstanceConfigsResponseBody struct {
	// The configurations of the instance.
	InstanceConfigs *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs `json:"InstanceConfigs,omitempty" xml:"InstanceConfigs,omitempty" type:"Struct"`
	// Current page number, starting from **1**, default is **1**.
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
	// Total number of entries for domain instance configurations.
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

func (s ListCloudGtmInstanceConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstanceConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstanceConfigsResponseBody) GetInstanceConfigs() *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs {
	return s.InstanceConfigs
}

func (s *ListCloudGtmInstanceConfigsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmInstanceConfigsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmInstanceConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudGtmInstanceConfigsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListCloudGtmInstanceConfigsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListCloudGtmInstanceConfigsResponseBody) SetInstanceConfigs(v *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs) *ListCloudGtmInstanceConfigsResponseBody {
	s.InstanceConfigs = v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBody) SetPageNumber(v int32) *ListCloudGtmInstanceConfigsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBody) SetPageSize(v int32) *ListCloudGtmInstanceConfigsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBody) SetRequestId(v string) *ListCloudGtmInstanceConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBody) SetTotalItems(v int32) *ListCloudGtmInstanceConfigsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBody) SetTotalPages(v int32) *ListCloudGtmInstanceConfigsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs struct {
	InstanceConfig []*ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig `json:"InstanceConfig,omitempty" xml:"InstanceConfig,omitempty" type:"Repeated"`
}

func (s ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs) GetInstanceConfig() []*ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	return s.InstanceConfig
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs) SetInstanceConfig(v []*ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs {
	s.InstanceConfig = v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigs) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig struct {
	// The policy for load balancing between address pools. Valid values:
	//
	// 	- round_robin: All address pools are returned for DNS requests from any source. All address pools are sorted in round-robin mode each time they are returned.
	//
	// 	- sequence: The address pool with the smallest sequence number is preferentially returned for DNS requests from any source. The sequence number indicates the priority for returning the address pool. A smaller sequence number indicates a higher priority. If the address pool with the smallest sequence number is unavailable, the address pool with the second smallest sequence number is returned.
	//
	// 	- weight: You can set a different weight value for each address pool. This way, address pools are returned based on the weight values.
	//
	// 	- source_nearest: GTM returns different address pools based on the sources of DNS requests. This way, users can access nearby addresses.
	//
	// example:
	//
	// round_robin
	AddressPoolLbStrategy *string `json:"AddressPoolLbStrategy,omitempty" xml:"AddressPoolLbStrategy,omitempty"`
	// The address pools.
	AddressPools *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools `json:"AddressPools,omitempty" xml:"AddressPools,omitempty" type:"Struct"`
	// The availability state of the access domain name. Valid values:
	//
	// 	- available: If the access domain name is **enabled*	- and the health state of the access domain name is **Normal**, the access domain name is deemed **Available**.
	//
	// 	- unavailable: If the access domain name is **disabled*	- or the health state of the access domain name is **Abnormal**, the access domain name is deemed **Unavailable**.
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
	// Instance configuration creation time.
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
	// 	- enable: The access domain name is enabled and the intelligent scheduling policy of the GTM instance takes effect.
	//
	// 	- disable: The access domain name is disabled and the intelligent scheduling policy of the GTM instance does not take effect.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The health state of the access domain name. Valid values:
	//
	// 	- ok: The health state of the access domain name is Normal and all address pools that are referenced by the access domain name are available.
	//
	// 	- ok_alert: The health state of the access domain name is Warning and some of the address pools that are referenced by the access domain name are unavailable. In this case, available address pools are normally used for DNS resolution, but unavailable address pools cannot be used for DNS resolution.
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
	// Remarks on the configuration of domain instance.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The GTM access domain name. The value of this parameter is composed of the value of ScheduleHostname and the value of ScheduleZoneName.
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
	// 	- custom: custom allocation. You must specify a custom hostname and associate the hostname with a zone within the account to which the GTM instance belongs to generate an access domain name.
	//
	// 	- sys_assign: system allocation. This mode is not supported. Do not set ScheduleZoneMode to sys_assign.
	//
	// example:
	//
	// custom
	ScheduleZoneMode *string `json:"ScheduleZoneMode,omitempty" xml:"ScheduleZoneMode,omitempty"`
	// The zone (such as example.com) or subzone (such as a.example.com) of the GTM access domain name. In most cases, the zone or subzone is hosted in Authoritative DNS Resolution of the Alibaba Cloud DNS console within the account to which the GTM instance belongs.
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
	// Global TTL (in seconds), the TTL value for domain resolution to addresses in the address pool, affecting the caching time of DNS records in the ISP\\"s LocalDNS. Supports custom TTL values.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The last modified time of the instance configuration.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The last modification time of the instance configuration (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// The edition of the GTM 3.0 instance. Valid values:
	//
	// 	- standard: Standard Edition
	//
	// 	- ultimate: Ultimate Edition
	//
	// example:
	//
	// ultimate
	VersionCode *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
}

func (s ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetAddressPoolLbStrategy() *string {
	return s.AddressPoolLbStrategy
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetAddressPools() *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools {
	return s.AddressPools
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetConfigId() *string {
	return s.ConfigId
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetRemark() *string {
	return s.Remark
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetScheduleDomainName() *string {
	return s.ScheduleDomainName
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetScheduleHostname() *string {
	return s.ScheduleHostname
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetScheduleRrType() *string {
	return s.ScheduleRrType
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetScheduleZoneMode() *string {
	return s.ScheduleZoneMode
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetScheduleZoneName() *string {
	return s.ScheduleZoneName
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetTtl() *int32 {
	return s.Ttl
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetVersionCode() *string {
	return s.VersionCode
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetAddressPoolLbStrategy(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.AddressPoolLbStrategy = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetAddressPools(v *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.AddressPools = v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetAvailableStatus(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.AvailableStatus = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetCommodityCode(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.CommodityCode = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetConfigId(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ConfigId = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetCreateTime(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.CreateTime = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetCreateTimestamp(v int64) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.CreateTimestamp = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetEnableStatus(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.EnableStatus = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetHealthStatus(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.HealthStatus = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetInstanceId(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.InstanceId = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetRemark(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.Remark = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetScheduleDomainName(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleDomainName = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetScheduleHostname(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleHostname = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetScheduleRrType(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleRrType = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetScheduleZoneMode(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleZoneMode = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetScheduleZoneName(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleZoneName = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetSequenceLbStrategyMode(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetTtl(v int32) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.Ttl = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetUpdateTime(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetUpdateTimestamp(v int64) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetVersionCode(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.VersionCode = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools struct {
	AddressPool []*ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool `json:"AddressPool,omitempty" xml:"AddressPool,omitempty" type:"Repeated"`
}

func (s ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) GetAddressPool() []*ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	return s.AddressPool
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) SetAddressPool(v []*ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools {
	s.AddressPool = v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool struct {
	// Load balancing policy among addresses in the address pool:
	//
	// - round_robin: Round-robin, for any source of DNS resolution requests, returns all addresses and rotates the order of all addresses each time.
	//
	// - sequence: Sequential, for any source of DNS resolution requests, returns the address with the smaller sequence number (the sequence number indicates the priority of the address return, the smaller the higher the priority). If the address with the smaller sequence number is unavailable, return the next address with a smaller sequence number.
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
	// The health state of the address pool. Valid values:
	//
	// 	- ok: The health state of the address pool is Normal and all addresses that are referenced by the address pool are available.
	//
	// 	- ok_alert: The health state of the address pool is Warning and some of the addresses that are referenced by the address pool are unavailable. However, the address pool is deemed normal. In this case, available address pools are normally used for DNS resolution, but unavailable address pools cannot be used for DNS resolution.
	//
	// 	- exceptional: The health state of the address pool is Abnormal and some or all of the addresses that are referenced by the address pool are unavailable. In this case, the address pool is deemed abnormal.
	//
	// example:
	//
	// ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// The health state of the address pool. Valid values:
	//
	// 	- ok: The health state of the address pool is Normal and all addresses that are referenced by the address pool are available.
	//
	// 	- ok_alert: The health state of the address pool is Warning and some of the addresses that are referenced by the address pool are unavailable. However, the address pool is deemed normal. In this case, available address pools are normally used for DNS resolution, but unavailable address pools cannot be used for DNS resolution.
	//
	// 	- exceptional: The health state of the address pool is Abnormal and some or all of the addresses that are referenced by the address pool are unavailable. In this case, the address pool is deemed abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// Parse the request source list.
	RequestSource *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Struct"`
	// Indicates whether the mode of the sequence policy for load balancing between address pools is non-preemptive. This parameter is available only for the multicloud integration scenario. Valid values:
	//
	// 	- true
	//
	// 	- false
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
	// Sequence number. For any parsing request, the address pool with the smaller sequence number (indicating the priority of the address pool returned, with smaller numbers having higher priority) is returned.
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
	// Weight value (an integer between 1 and 100, including both 1 and 100), which supports setting different weight values for each address pool, enabling the resolution query to return address pools according to the weighted ratio.
	//
	// example:
	//
	// 1
	WeightValue *int32 `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
}

func (s ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetAddressLbStrategy() *string {
	return s.AddressLbStrategy
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetAddressPoolType() *string {
	return s.AddressPoolType
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetRequestSource() *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource {
	return s.RequestSource
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetSeqNonPreemptiveSchedule() *bool {
	return s.SeqNonPreemptiveSchedule
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetSerialNumber() *int32 {
	return s.SerialNumber
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) GetWeightValue() *int32 {
	return s.WeightValue
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetAddressLbStrategy(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.AddressLbStrategy = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetAddressPoolId(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.AddressPoolId = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetAddressPoolName(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.AddressPoolName = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetAddressPoolType(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.AddressPoolType = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetAvailableStatus(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.AvailableStatus = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetCreateTime(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.CreateTime = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetCreateTimestamp(v int64) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.CreateTimestamp = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetEnableStatus(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.EnableStatus = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetHealthJudgement(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.HealthJudgement = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetHealthStatus(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.HealthStatus = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetRequestSource(v *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.RequestSource = v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetSeqNonPreemptiveSchedule(v bool) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.SeqNonPreemptiveSchedule = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetSequenceLbStrategyMode(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetSerialNumber(v int32) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.SerialNumber = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetUpdateTime(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetUpdateTimestamp(v int64) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) SetWeightValue(v int32) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool {
	s.WeightValue = &v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource struct {
	RequestSource []*string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Repeated"`
}

func (s ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) GoString() string {
	return s.String()
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) GetRequestSource() []*string {
	return s.RequestSource
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) SetRequestSource(v []*string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource {
	s.RequestSource = v
	return s
}

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource) Validate() error {
	return dara.Validate(s)
}
