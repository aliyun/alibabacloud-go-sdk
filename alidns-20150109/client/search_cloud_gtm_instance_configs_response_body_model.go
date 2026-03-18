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
	AddressPoolLbStrategy *string                                                                             `json:"AddressPoolLbStrategy,omitempty" xml:"AddressPoolLbStrategy,omitempty"`
	AddressPools          *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPools `json:"AddressPools,omitempty" xml:"AddressPools,omitempty" type:"Struct"`
	AvailableStatus       *string                                                                             `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	CommodityCode         *string                                                                             `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	ConfigId              *string                                                                             `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
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

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) GetConfigLoggingSwitchStatus() *string {
	return s.ConfigLoggingSwitchStatus
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

func (s *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig) SetConfigLoggingSwitchStatus(v string) *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfig {
	s.ConfigLoggingSwitchStatus = &v
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
	AddressLbStrategy        *string                                                                                                     `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	AddressPoolId            *string                                                                                                     `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	AddressPoolName          *string                                                                                                     `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	AddressPoolType          *string                                                                                                     `json:"AddressPoolType,omitempty" xml:"AddressPoolType,omitempty"`
	AvailableStatus          *string                                                                                                     `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	CreateTime               *string                                                                                                     `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp          *int64                                                                                                      `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EnableStatus             *string                                                                                                     `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthJudgement          *string                                                                                                     `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	HealthStatus             *string                                                                                                     `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	RequestSource            *SearchCloudGtmInstanceConfigsResponseBodyInstanceConfigsInstanceConfigAddressPoolsAddressPoolRequestSource `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Struct"`
	SeqNonPreemptiveSchedule *bool                                                                                                       `json:"SeqNonPreemptiveSchedule,omitempty" xml:"SeqNonPreemptiveSchedule,omitempty"`
	SequenceLbStrategyMode   *string                                                                                                     `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
	SerialNumber             *int32                                                                                                      `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	UpdateTime               *string                                                                                                     `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp          *int64                                                                                                      `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	WeightValue              *int32                                                                                                      `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
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
