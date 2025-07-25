// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressPoolReferenceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPoolId(v string) *DescribeCloudGtmAddressPoolReferenceResponseBody
	GetAddressPoolId() *string
	SetAddressPoolName(v string) *DescribeCloudGtmAddressPoolReferenceResponseBody
	GetAddressPoolName() *string
	SetInstanceConfigs(v *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs) *DescribeCloudGtmAddressPoolReferenceResponseBody
	GetInstanceConfigs() *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs
	SetRequestId(v string) *DescribeCloudGtmAddressPoolReferenceResponseBody
	GetRequestId() *string
}

type DescribeCloudGtmAddressPoolReferenceResponseBody struct {
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
	// app
	AddressPoolName *string `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	// The access domain names that reference the address pool.
	InstanceConfigs *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs `json:"InstanceConfigs,omitempty" xml:"InstanceConfigs,omitempty" type:"Struct"`
	// Unique request identification code.
	//
	// example:
	//
	// 853805EA-3D47-47D5-9A1A-A45C24313ABD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudGtmAddressPoolReferenceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolReferenceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBody) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBody) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBody) GetInstanceConfigs() *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs {
	return s.InstanceConfigs
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBody) SetAddressPoolId(v string) *DescribeCloudGtmAddressPoolReferenceResponseBody {
	s.AddressPoolId = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBody) SetAddressPoolName(v string) *DescribeCloudGtmAddressPoolReferenceResponseBody {
	s.AddressPoolName = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBody) SetInstanceConfigs(v *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs) *DescribeCloudGtmAddressPoolReferenceResponseBody {
	s.InstanceConfigs = v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBody) SetRequestId(v string) *DescribeCloudGtmAddressPoolReferenceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs struct {
	InstanceConfig []*DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig `json:"InstanceConfig,omitempty" xml:"InstanceConfig,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs) GetInstanceConfig() []*DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	return s.InstanceConfig
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs) SetInstanceConfig(v []*DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs {
	s.InstanceConfig = v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig struct {
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
	// 	- available: If the access domain name is **enabled*	- and the health state is **normal**, the access domain name is deemed **available**.
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
	// config-000**1
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
	// The ID of the Global Traffic Manager (GTM) 3.0 instance.
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
	// DNS record types for scheduling domains:
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
	// Global TTL, the TTL value for resolving the accessed domain name to addresses in the address pool, which affects the caching time of DNS records in the operator\\"s LocalDNS. Supports custom TTL values.
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

func (s DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetAddressPoolLbStrategy() *string {
	return s.AddressPoolLbStrategy
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetScheduleDomainName() *string {
	return s.ScheduleDomainName
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetScheduleHostname() *string {
	return s.ScheduleHostname
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetScheduleRrType() *string {
	return s.ScheduleRrType
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetScheduleZoneName() *string {
	return s.ScheduleZoneName
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) GetVersionCode() *string {
	return s.VersionCode
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetAddressPoolLbStrategy(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.AddressPoolLbStrategy = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetAvailableStatus(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.AvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetConfigId(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetEnableStatus(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.EnableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetHealthStatus(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.HealthStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetInstanceId(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.InstanceId = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetInstanceName(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.InstanceName = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetRemark(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.Remark = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetScheduleDomainName(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleDomainName = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetScheduleHostname(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleHostname = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetScheduleRrType(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleRrType = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetScheduleZoneName(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.ScheduleZoneName = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetSequenceLbStrategyMode(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetTtl(v int32) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.Ttl = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) SetVersionCode(v string) *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig {
	s.VersionCode = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig) Validate() error {
	return dara.Validate(s)
}
