// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmAccessStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMode(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetAccessMode() *string
	SetCreateTime(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeDnsGtmAccessStrategyResponseBody
	GetCreateTimestamp() *int64
	SetDefaultAddrPoolGroupStatus(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetDefaultAddrPoolGroupStatus() *string
	SetDefaultAddrPoolType(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetDefaultAddrPoolType() *string
	SetDefaultAddrPools(v *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) *DescribeDnsGtmAccessStrategyResponseBody
	GetDefaultAddrPools() *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools
	SetDefaultAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody
	GetDefaultAvailableAddrNum() *int32
	SetDefaultLatencyOptimization(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetDefaultLatencyOptimization() *string
	SetDefaultLbaStrategy(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetDefaultLbaStrategy() *string
	SetDefaultMaxReturnAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody
	GetDefaultMaxReturnAddrNum() *int32
	SetDefaultMinAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody
	GetDefaultMinAvailableAddrNum() *int32
	SetEffectiveAddrPoolGroupType(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetEffectiveAddrPoolGroupType() *string
	SetFailoverAddrPoolGroupStatus(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetFailoverAddrPoolGroupStatus() *string
	SetFailoverAddrPoolType(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetFailoverAddrPoolType() *string
	SetFailoverAddrPools(v *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) *DescribeDnsGtmAccessStrategyResponseBody
	GetFailoverAddrPools() *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools
	SetFailoverAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody
	GetFailoverAvailableAddrNum() *int32
	SetFailoverLatencyOptimization(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetFailoverLatencyOptimization() *string
	SetFailoverLbaStrategy(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetFailoverLbaStrategy() *string
	SetFailoverMaxReturnAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody
	GetFailoverMaxReturnAddrNum() *int32
	SetFailoverMinAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody
	GetFailoverMinAvailableAddrNum() *int32
	SetInstanceId(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetInstanceId() *string
	SetLines(v *DescribeDnsGtmAccessStrategyResponseBodyLines) *DescribeDnsGtmAccessStrategyResponseBody
	GetLines() *DescribeDnsGtmAccessStrategyResponseBodyLines
	SetRequestId(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetRequestId() *string
	SetStrategyId(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetStrategyId() *string
	SetStrategyMode(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetStrategyMode() *string
	SetStrategyName(v string) *DescribeDnsGtmAccessStrategyResponseBody
	GetStrategyName() *string
}

type DescribeDnsGtmAccessStrategyResponseBody struct {
	// The primary/secondary switchover policy for address pool groups. Valid values:
	//
	// 	- AUTO: performs automatic switchover between the primary and secondary address pool groups upon failures.
	//
	// 	- DEFAULT: uses the primary address pool group.
	//
	// 	- FAILOVER: uses the secondary address pool group.
	//
	// example:
	//
	// auto
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// The time when the access policy was created.
	//
	// example:
	//
	// 2018-08-09T00:10Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp that indicates when the access policy was created.
	//
	// example:
	//
	// 1533773400000
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The status of the primary address pool group. Valid values:
	//
	// 	- AVAILABLE: available
	//
	// 	- NOT_AVAILABLE: unavailable
	//
	// example:
	//
	// AVAILABLE
	DefaultAddrPoolGroupStatus *string `json:"DefaultAddrPoolGroupStatus,omitempty" xml:"DefaultAddrPoolGroupStatus,omitempty"`
	// The type of the primary address pool. Valid values:
	//
	// 	- IPV4
	//
	// 	- IPV6
	//
	// 	- DOMAIN
	//
	// example:
	//
	// ipv4
	DefaultAddrPoolType *string `json:"DefaultAddrPoolType,omitempty" xml:"DefaultAddrPoolType,omitempty"`
	// The address pools in the primary address pool group.
	DefaultAddrPools *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools `json:"DefaultAddrPools,omitempty" xml:"DefaultAddrPools,omitempty" type:"Struct"`
	// The number of available addresses in the primary address pool.
	//
	// example:
	//
	// 1
	DefaultAvailableAddrNum *int32 `json:"DefaultAvailableAddrNum,omitempty" xml:"DefaultAvailableAddrNum,omitempty"`
	// Indicates whether scheduling optimization for latency resolution was enabled for the primary address pool group. Valid values:
	//
	// 	- OPEN: enabled
	//
	// 	- CLOSE: disabled
	//
	// example:
	//
	// open
	DefaultLatencyOptimization *string `json:"DefaultLatencyOptimization,omitempty" xml:"DefaultLatencyOptimization,omitempty"`
	// The load balancing policy of the primary address pool group. Valid values:
	//
	// 	- ALL_RR: returns all addresses.
	//
	// 	- RATIO: returns addresses by weight.
	//
	// example:
	//
	// all_rr
	DefaultLbaStrategy *string `json:"DefaultLbaStrategy,omitempty" xml:"DefaultLbaStrategy,omitempty"`
	// The maximum number of addresses returned from the primary address pool group.
	//
	// example:
	//
	// 1
	DefaultMaxReturnAddrNum *int32 `json:"DefaultMaxReturnAddrNum,omitempty" xml:"DefaultMaxReturnAddrNum,omitempty"`
	// The minimum number of available addresses in the primary address pool group.
	//
	// example:
	//
	// 1
	DefaultMinAvailableAddrNum *int32 `json:"DefaultMinAvailableAddrNum,omitempty" xml:"DefaultMinAvailableAddrNum,omitempty"`
	// The type of the active address pool group. Valid values:
	//
	// 	- DEFAULT: the primary address pool group
	//
	// 	- FAILOVER: the secondary address pool group
	//
	// example:
	//
	// DEFAULT
	EffectiveAddrPoolGroupType *string `json:"EffectiveAddrPoolGroupType,omitempty" xml:"EffectiveAddrPoolGroupType,omitempty"`
	// The status of the secondary address pool group. Valid values:
	//
	// 	- AVAILABLE: available
	//
	// 	- NOT_AVAILABLE: unavailable
	//
	// example:
	//
	// AVAILABLE
	FailoverAddrPoolGroupStatus *string `json:"FailoverAddrPoolGroupStatus,omitempty" xml:"FailoverAddrPoolGroupStatus,omitempty"`
	// The type of the secondary address pool. Valid values:
	//
	// 	- IPV4
	//
	// 	- IPV6
	//
	// 	- DOMAIN
	//
	// example:
	//
	// ipv4
	FailoverAddrPoolType *string `json:"FailoverAddrPoolType,omitempty" xml:"FailoverAddrPoolType,omitempty"`
	// The address pools in the secondary address pool group.
	FailoverAddrPools *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools `json:"FailoverAddrPools,omitempty" xml:"FailoverAddrPools,omitempty" type:"Struct"`
	// The number of available addresses in the secondary address pool.
	//
	// example:
	//
	// 1
	FailoverAvailableAddrNum *int32 `json:"FailoverAvailableAddrNum,omitempty" xml:"FailoverAvailableAddrNum,omitempty"`
	// Indicates whether scheduling optimization for latency resolution was enabled for the secondary address pool group. Valid values:
	//
	// 	- OPEN: enabled
	//
	// 	- CLOSE: disabled
	//
	// example:
	//
	// open
	FailoverLatencyOptimization *string `json:"FailoverLatencyOptimization,omitempty" xml:"FailoverLatencyOptimization,omitempty"`
	// The load balancing policy of the secondary address pool group. Valid values:
	//
	// 	- ALL_RR: returns all addresses.
	//
	// 	- RATIO: returns addresses by weight.
	//
	// example:
	//
	// all_rr
	FailoverLbaStrategy *string `json:"FailoverLbaStrategy,omitempty" xml:"FailoverLbaStrategy,omitempty"`
	// The maximum number of addresses returned from the secondary address pool group.
	//
	// example:
	//
	// 1
	FailoverMaxReturnAddrNum *int32 `json:"FailoverMaxReturnAddrNum,omitempty" xml:"FailoverMaxReturnAddrNum,omitempty"`
	// The minimum number of available addresses in the secondary address pool group.
	//
	// example:
	//
	// 1
	FailoverMinAvailableAddrNum *int32 `json:"FailoverMinAvailableAddrNum,omitempty" xml:"FailoverMinAvailableAddrNum,omitempty"`
	// The ID of the associated instance.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The source regions.
	Lines *DescribeDnsGtmAccessStrategyResponseBodyLines `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// BA1608CA-834C-4E63-8682-8AF0B11ED72D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the access policy.
	//
	// example:
	//
	// strategyId1
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The type of the access policy. Valid values:
	//
	// 	- GEO: geographical location-based
	//
	// 	- LATENCY: latency-based
	//
	// example:
	//
	// geo
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	// The name of the access policy.
	//
	// example:
	//
	// strategyName1
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetAccessMode() *string {
	return s.AccessMode
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetDefaultAddrPoolGroupStatus() *string {
	return s.DefaultAddrPoolGroupStatus
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetDefaultAddrPoolType() *string {
	return s.DefaultAddrPoolType
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetDefaultAddrPools() *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools {
	return s.DefaultAddrPools
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetDefaultAvailableAddrNum() *int32 {
	return s.DefaultAvailableAddrNum
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetDefaultLatencyOptimization() *string {
	return s.DefaultLatencyOptimization
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetDefaultLbaStrategy() *string {
	return s.DefaultLbaStrategy
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetDefaultMaxReturnAddrNum() *int32 {
	return s.DefaultMaxReturnAddrNum
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetDefaultMinAvailableAddrNum() *int32 {
	return s.DefaultMinAvailableAddrNum
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetEffectiveAddrPoolGroupType() *string {
	return s.EffectiveAddrPoolGroupType
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetFailoverAddrPoolGroupStatus() *string {
	return s.FailoverAddrPoolGroupStatus
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetFailoverAddrPoolType() *string {
	return s.FailoverAddrPoolType
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetFailoverAddrPools() *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools {
	return s.FailoverAddrPools
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetFailoverAvailableAddrNum() *int32 {
	return s.FailoverAvailableAddrNum
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetFailoverLatencyOptimization() *string {
	return s.FailoverLatencyOptimization
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetFailoverLbaStrategy() *string {
	return s.FailoverLbaStrategy
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetFailoverMaxReturnAddrNum() *int32 {
	return s.FailoverMaxReturnAddrNum
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetFailoverMinAvailableAddrNum() *int32 {
	return s.FailoverMinAvailableAddrNum
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetLines() *DescribeDnsGtmAccessStrategyResponseBodyLines {
	return s.Lines
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetStrategyId() *string {
	return s.StrategyId
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetStrategyMode() *string {
	return s.StrategyMode
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) GetStrategyName() *string {
	return s.StrategyName
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetAccessMode(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.AccessMode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetCreateTime(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetCreateTimestamp(v int64) *DescribeDnsGtmAccessStrategyResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultAddrPoolGroupStatus(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultAddrPoolGroupStatus = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultAddrPoolType(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultAddrPoolType = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultAddrPools(v *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultAddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultAvailableAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultLatencyOptimization(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultLatencyOptimization = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultLbaStrategy(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultLbaStrategy = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultMaxReturnAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultMaxReturnAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetDefaultMinAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.DefaultMinAvailableAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetEffectiveAddrPoolGroupType(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.EffectiveAddrPoolGroupType = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverAddrPoolGroupStatus(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolGroupStatus = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverAddrPoolType(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverAddrPoolType = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverAddrPools(v *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverAddrPools = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverAvailableAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverLatencyOptimization(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverLatencyOptimization = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverLbaStrategy(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverLbaStrategy = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverMaxReturnAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverMaxReturnAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetFailoverMinAvailableAddrNum(v int32) *DescribeDnsGtmAccessStrategyResponseBody {
	s.FailoverMinAvailableAddrNum = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetInstanceId(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.InstanceId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetLines(v *DescribeDnsGtmAccessStrategyResponseBodyLines) *DescribeDnsGtmAccessStrategyResponseBody {
	s.Lines = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetRequestId(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetStrategyId(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.StrategyId = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetStrategyMode(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.StrategyMode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) SetStrategyName(v string) *DescribeDnsGtmAccessStrategyResponseBody {
	s.StrategyName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBody) Validate() error {
	if s.DefaultAddrPools != nil {
		if err := s.DefaultAddrPools.Validate(); err != nil {
			return err
		}
	}
	if s.FailoverAddrPools != nil {
		if err := s.FailoverAddrPools.Validate(); err != nil {
			return err
		}
	}
	if s.Lines != nil {
		if err := s.Lines.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools struct {
	DefaultAddrPool []*DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool `json:"DefaultAddrPool,omitempty" xml:"DefaultAddrPool,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) GetDefaultAddrPool() []*DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool {
	return s.DefaultAddrPool
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) SetDefaultAddrPool(v []*DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools {
	s.DefaultAddrPool = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPools) Validate() error {
	if s.DefaultAddrPool != nil {
		for _, item := range s.DefaultAddrPool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool struct {
	// The number of addresses in the address pool.
	//
	// example:
	//
	// 1
	AddrCount *int32 `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	// The ID of the address pool.
	//
	// example:
	//
	// pool1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The weight of the address pool.
	//
	// example:
	//
	// 1
	LbaWeight *int32 `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) GetAddrCount() *int32 {
	return s.AddrCount
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) GetId() *string {
	return s.Id
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) GetName() *string {
	return s.Name
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) SetId(v string) *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool {
	s.Id = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) SetLbaWeight(v int32) *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool {
	s.LbaWeight = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) SetName(v string) *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyDefaultAddrPoolsDefaultAddrPool) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools struct {
	FailoverAddrPool []*DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool `json:"FailoverAddrPool,omitempty" xml:"FailoverAddrPool,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) GetFailoverAddrPool() []*DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool {
	return s.FailoverAddrPool
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) SetFailoverAddrPool(v []*DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools {
	s.FailoverAddrPool = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPools) Validate() error {
	if s.FailoverAddrPool != nil {
		for _, item := range s.FailoverAddrPool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool struct {
	// The number of addresses in the address pool.
	//
	// example:
	//
	// 1
	AddrCount *int32 `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	// The ID of the address pool.
	//
	// example:
	//
	// pool1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The weight of the address pool.
	//
	// example:
	//
	// 1
	LbaWeight *int32 `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) GetAddrCount() *int32 {
	return s.AddrCount
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) GetId() *string {
	return s.Id
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) GetName() *string {
	return s.Name
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) SetAddrCount(v int32) *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) SetId(v string) *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool {
	s.Id = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) SetLbaWeight(v int32) *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool {
	s.LbaWeight = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) SetName(v string) *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyFailoverAddrPoolsFailoverAddrPool) Validate() error {
	return dara.Validate(s)
}

type DescribeDnsGtmAccessStrategyResponseBodyLines struct {
	Line []*DescribeDnsGtmAccessStrategyResponseBodyLinesLine `json:"Line,omitempty" xml:"Line,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmAccessStrategyResponseBodyLines) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponseBodyLines) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLines) GetLine() []*DescribeDnsGtmAccessStrategyResponseBodyLinesLine {
	return s.Line
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLines) SetLine(v []*DescribeDnsGtmAccessStrategyResponseBodyLinesLine) *DescribeDnsGtmAccessStrategyResponseBodyLines {
	s.Line = v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLines) Validate() error {
	if s.Line != nil {
		for _, item := range s.Line {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmAccessStrategyResponseBodyLinesLine struct {
	// The code of the source region group.
	//
	// example:
	//
	// default
	GroupCode *string `json:"GroupCode,omitempty" xml:"GroupCode,omitempty"`
	// The name of the source region group.
	//
	// example:
	//
	// global
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The line code of the source region.
	//
	// example:
	//
	// default
	LineCode *string `json:"LineCode,omitempty" xml:"LineCode,omitempty"`
	// The line name of the source region.
	//
	// example:
	//
	// global
	LineName *string `json:"LineName,omitempty" xml:"LineName,omitempty"`
}

func (s DescribeDnsGtmAccessStrategyResponseBodyLinesLine) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmAccessStrategyResponseBodyLinesLine) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLinesLine) GetGroupCode() *string {
	return s.GroupCode
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLinesLine) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLinesLine) GetLineCode() *string {
	return s.LineCode
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLinesLine) GetLineName() *string {
	return s.LineName
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLinesLine) SetGroupCode(v string) *DescribeDnsGtmAccessStrategyResponseBodyLinesLine {
	s.GroupCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLinesLine) SetGroupName(v string) *DescribeDnsGtmAccessStrategyResponseBodyLinesLine {
	s.GroupName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLinesLine) SetLineCode(v string) *DescribeDnsGtmAccessStrategyResponseBodyLinesLine {
	s.LineCode = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLinesLine) SetLineName(v string) *DescribeDnsGtmAccessStrategyResponseBodyLinesLine {
	s.LineName = &v
	return s
}

func (s *DescribeDnsGtmAccessStrategyResponseBodyLinesLine) Validate() error {
	return dara.Validate(s)
}
