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
	if s.InstanceConfigs != nil {
		if err := s.InstanceConfigs.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig struct {
	AddressPoolLbStrategy *string                                                                           `json:"AddressPoolLbStrategy,omitempty" xml:"AddressPoolLbStrategy,omitempty"`
	AddressPools          *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools `json:"AddressPools,omitempty" xml:"AddressPools,omitempty" type:"Struct"`
	AvailableStatus       *string                                                                           `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	CommodityCode         *string                                                                           `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConfigId              *string                                                                           `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// example:
	//
	// ENABLE
	ConfigLoggingSwitchStatus *string `json:"ConfigLoggingSwitchStatus,omitempty" xml:"ConfigLoggingSwitchStatus,omitempty"`
	CreateTime                *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp           *int64  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EnableStatus              *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthStatus              *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	InstanceId                *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Remark                    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	ScheduleDomainName        *string `json:"ScheduleDomainName,omitempty" xml:"ScheduleDomainName,omitempty"`
	ScheduleHostname          *string `json:"ScheduleHostname,omitempty" xml:"ScheduleHostname,omitempty"`
	ScheduleRrType            *string `json:"ScheduleRrType,omitempty" xml:"ScheduleRrType,omitempty"`
	ScheduleZoneMode          *string `json:"ScheduleZoneMode,omitempty" xml:"ScheduleZoneMode,omitempty"`
	ScheduleZoneName          *string `json:"ScheduleZoneName,omitempty" xml:"ScheduleZoneName,omitempty"`
	SequenceLbStrategyMode    *string `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
	Ttl                       *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	UpdateTime                *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp           *int64  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	VersionCode               *string `json:"VersionCode,omitempty" xml:"VersionCode,omitempty"`
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

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetConfigLoggingSwitchStatus() *string {
	return s.ConfigLoggingSwitchStatus
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

func (s *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetConfigLoggingSwitchStatus(v string) *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ConfigLoggingSwitchStatus = &v
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
	if s.AddressPools != nil {
		if err := s.AddressPools.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPool struct {
	AddressLbStrategy        *string                                                                                                   `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	AddressPoolId            *string                                                                                                   `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	AddressPoolName          *string                                                                                                   `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	AddressPoolType          *string                                                                                                   `json:"AddressPoolType,omitempty" xml:"AddressPoolType,omitempty"`
	AvailableStatus          *string                                                                                                   `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	CreateTime               *string                                                                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp          *int64                                                                                                    `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EnableStatus             *string                                                                                                   `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthJudgement          *string                                                                                                   `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	HealthStatus             *string                                                                                                   `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	RequestSource            *ListCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Struct"`
	SeqNonPreemptiveSchedule *bool                                                                                                     `json:"SeqNonPreemptiveSchedule,omitempty" xml:"SeqNonPreemptiveSchedule,omitempty"`
	SequenceLbStrategyMode   *string                                                                                                   `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
	SerialNumber             *int32                                                                                                    `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	UpdateTime               *string                                                                                                   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp          *int64                                                                                                    `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	WeightValue              *int32                                                                                                    `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
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
	if s.RequestSource != nil {
		if err := s.RequestSource.Validate(); err != nil {
			return err
		}
	}
	return nil
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
