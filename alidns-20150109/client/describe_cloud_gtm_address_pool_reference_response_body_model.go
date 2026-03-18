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
	AddressPoolName *string                                                          `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
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
	if s.InstanceConfigs != nil {
		if err := s.InstanceConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type DescribeCloudGtmAddressPoolReferenceResponseBodyInstanceConfigsInstanceConfig struct {
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
