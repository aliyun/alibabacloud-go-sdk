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
	AddressId    *string                                                   `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
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
	if s.AddressPools != nil {
		if err := s.AddressPools.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPool struct {
	AddressLbStrategy      *string                                                                             `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	AddressPoolId          *string                                                                             `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	AddressPoolName        *string                                                                             `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	AddressPoolType        *string                                                                             `json:"AddressPoolType,omitempty" xml:"AddressPoolType,omitempty"`
	AvailableStatus        *string                                                                             `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	EnableStatus           *string                                                                             `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthJudgement        *string                                                                             `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	HealthStatus           *string                                                                             `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	InstanceConfigs        *DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigs `json:"InstanceConfigs,omitempty" xml:"InstanceConfigs,omitempty" type:"Struct"`
	Remark                 *string                                                                             `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SequenceLbStrategyMode *string                                                                             `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
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
	if s.InstanceConfigs != nil {
		if err := s.InstanceConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type DescribeCloudGtmAddressReferenceResponseBodyAddressPoolsAddressPoolInstanceConfigsInstanceConfig struct {
	AddressPoolLbStrategy  *string `json:"AddressPoolLbStrategy,omitempty" xml:"AddressPoolLbStrategy,omitempty"`
	AvailableStatus        *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	ConfigId               *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	EnableStatus           *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthStatus           *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	InstanceId             *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName           *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	Remark                 *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ScheduleDomainName     *string `json:"ScheduleDomainName,omitempty" xml:"ScheduleDomainName,omitempty"`
	ScheduleHostname       *string `json:"ScheduleHostname,omitempty" xml:"ScheduleHostname,omitempty"`
	ScheduleRrType         *string `json:"ScheduleRrType,omitempty" xml:"ScheduleRrType,omitempty"`
	ScheduleZoneName       *string `json:"ScheduleZoneName,omitempty" xml:"ScheduleZoneName,omitempty"`
	SequenceLbStrategyMode *string `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
	Ttl                    *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	VersionCode            *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
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
