// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressReferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *DescribeCloudGtmAddressReferenceResponseBody
	GetAddress() *string
	SetAddressId(v string) *DescribeCloudGtmAddressReferenceResponseBody
	GetAddressId() *string
	SetAddressPools(v *DescribeCloudGtmAddressReferenceResponseBodyAddressPools) *DescribeCloudGtmAddressReferenceResponseBody
	GetAddressPools() *DescribeCloudGtmAddressReferenceResponseBodyAddressPools
	SetName(v string) *DescribeCloudGtmAddressReferenceResponseBody
	GetName() *string
	SetRequestId(v string) *DescribeCloudGtmAddressReferenceResponseBody
	GetRequestId() *string
}

type DescribeCloudGtmAddressReferenceResponseBody struct {
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
	// addr-89564584963974**40
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// The address pools.
	AddressPools *DescribeCloudGtmAddressReferenceResponseBodyAddressPools `json:"AddressPools,omitempty" xml:"AddressPools,omitempty" type:"Struct"`
	// Address name.
	//
	// example:
	//
	// Address-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudGtmAddressReferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressReferenceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressReferenceResponseBody) GetAddress() *string {
	return s.Address
}

func (s *DescribeCloudGtmAddressReferenceResponseBody) GetAddressId() *string {
	return s.AddressId
}

func (s *DescribeCloudGtmAddressReferenceResponseBody) GetAddressPools() *DescribeCloudGtmAddressReferenceResponseBodyAddressPools {
	return s.AddressPools
}

func (s *DescribeCloudGtmAddressReferenceResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeCloudGtmAddressReferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudGtmAddressReferenceResponseBody) SetAddress(v string) *DescribeCloudGtmAddressReferenceResponseBody {
	s.Address = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBody) SetAddressId(v string) *DescribeCloudGtmAddressReferenceResponseBody {
	s.AddressId = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBody) SetAddressPools(v *DescribeCloudGtmAddressReferenceResponseBodyAddressPools) *DescribeCloudGtmAddressReferenceResponseBody {
	s.AddressPools = v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBody) SetName(v string) *DescribeCloudGtmAddressReferenceResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBody) SetRequestId(v string) *DescribeCloudGtmAddressReferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmAddressReferenceResponseBodyAddressPools struct {
	AddressPool []*DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool `json:"AddressPool,omitempty" xml:"AddressPool,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmAddressReferenceResponseBodyAddressPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressReferenceResponseBodyAddressPools) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPools) GetAddressPool() []*DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	return s.AddressPool
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPools) SetAddressPool(v []*DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) *DescribeCloudGtmAddressReferenceResponseBodyAddressPools {
	s.AddressPool = v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPools) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool struct {
	// Load balancing policy among addresses in the address pool:
	//
	// - round_robin: Round-robin, for any source of DNS resolution requests, all addresses are returned, with a rotation sort applied to all addresses each time.
	//
	// - sequence: Sequential, for any source of DNS resolution requests, returns the address with the smaller sequence number (the sequence number indicates the priority of address return, with smaller numbers having higher priority). If the address with the smaller sequence number is unavailable, the next address with a smaller sequence number is returned.
	//
	// - weight: Weighted, supports setting different weight values for each address, realizing the return of addresses according to the ratio of weight for DNS query resolutions.
	//
	// - source_nearest: Source-nearest, i.e., intelligent resolution function, where GTM can return different addresses based on the source of different DNS resolution requests, achieving the effect of users accessing nearby.
	//
	// example:
	//
	// round_robin
	AddressLbStrategy *string `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	// The ID of the address pool.
	//
	// example:
	//
	// pool-895280232254422016
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
	// - available
	//
	// - unavailable
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
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
	// The instances that reference the address pool.
	InstanceConfigs *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigs `json:"InstanceConfigs,omitempty" xml:"InstanceConfigs,omitempty" type:"Struct"`
	// Remarks for the address pool.
	//
	// example:
	//
	// pool-1
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Load balancing policy between addresses in sequential mode during the recovery of preceding resources service mode:
	//
	// - preemptive: Preemption mode, where upon recovery of preceding resources, priority is given to using addresses with smaller sequence numbers;
	//
	// - non_preemptive: Non-preemption mode, where upon recovery of preceding resources, the current address continues to be used;
	//
	// example:
	//
	// preemptive
	SequenceLbStrategyMode *string `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
}

func (s DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GetAddressLbStrategy() *string {
	return s.AddressLbStrategy
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GetAddressPoolType() *string {
	return s.AddressPoolType
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GetInstanceConfigs() *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigs {
	return s.InstanceConfigs
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) SetAddressLbStrategy(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	s.AddressLbStrategy = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) SetAddressPoolId(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	s.AddressPoolId = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) SetAddressPoolName(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	s.AddressPoolName = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) SetAddressPoolType(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	s.AddressPoolType = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) SetAvailableStatus(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	s.AvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) SetEnableStatus(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	s.EnableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) SetHealthJudgement(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	s.HealthJudgement = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) SetHealthStatus(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	s.HealthStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) SetInstanceConfigs(v *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigs) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	s.InstanceConfigs = v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) SetRemark(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	s.Remark = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) SetSequenceLbStrategyMode(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigs struct {
	InstanceConfig []*DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig `json:"InstanceConfig,omitempty" xml:"InstanceConfig,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigs) GetInstanceConfig() []*DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	return s.InstanceConfig
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigs) SetInstanceConfig(v []*DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigs {
	s.InstanceConfig = v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig struct {
	// The policy for load balancing between address pools. Valid values:
	//
	// 	- round_robin: All address pools are returned for Domain Name System (DNS) requests from any source. All address pools are sorted in round-robin mode each time they are returned.
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
	// The availability state of the access domain name. Valid values:
	//
	// 	- available: If the access domain name is **enabled*	- and the health state is normal, the access domain name is deemed **available**.
	//
	// 	- unavailable: If the access domain name is **disabled*	- or the health state is **abnormal**, the access domain name is deemed **unavailable**.
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// The configuration ID of the access domain name. Two configuration IDs exist when the access domain name is bound to the same GTM instance but an A record and an AAAA record are configured for the access domain name. The configuration ID uniquely identifies a configuration.
	//
	// example:
	//
	// config-00**01
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
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
	// gtm-cn-zz11t58**0k
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Instance name.
	//
	// example:
	//
	// test
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// Remarks.
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
	// Global TTL (in seconds), the TTL value for domain name resolution to addresses in the address pool, which affects the caching time of DNS records in the ISP\\"s LocalDNS. Custom TTL values are supported.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
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

func (s DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetAddressPoolLbStrategy() *string {
	return s.AddressPoolLbStrategy
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetScheduleDomainName() *string {
	return s.ScheduleDomainName
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetScheduleHostname() *string {
	return s.ScheduleHostname
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetScheduleRrType() *string {
	return s.ScheduleRrType
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetScheduleZoneName() *string {
	return s.ScheduleZoneName
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetAddressPoolLbStrategy(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.AddressPoolLbStrategy = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetAvailableStatus(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.AvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetConfigId(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetEnableStatus(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.EnableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetHealthStatus(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.HealthStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetInstanceId(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetInstanceName(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.InstanceName = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetRemark(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.Remark = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetScheduleDomainName(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.ScheduleDomainName = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetScheduleHostname(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.ScheduleHostname = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetScheduleRrType(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.ScheduleRrType = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetScheduleZoneName(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.ScheduleZoneName = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetSequenceLbStrategyMode(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetTtl(v int32) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.Ttl = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) SetVersionCode(v string) *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig {
	s.VersionCode = &v
	return s
}

func (s *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig) Validate() error {
	return dara.Validate(s)
}
