// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDnsGtmAccessStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultAddrPool(v []*AddDnsGtmAccessStrategyRequestDefaultAddrPool) *AddDnsGtmAccessStrategyRequest
	GetDefaultAddrPool() []*AddDnsGtmAccessStrategyRequestDefaultAddrPool
	SetDefaultAddrPoolType(v string) *AddDnsGtmAccessStrategyRequest
	GetDefaultAddrPoolType() *string
	SetDefaultLatencyOptimization(v string) *AddDnsGtmAccessStrategyRequest
	GetDefaultLatencyOptimization() *string
	SetDefaultLbaStrategy(v string) *AddDnsGtmAccessStrategyRequest
	GetDefaultLbaStrategy() *string
	SetDefaultMaxReturnAddrNum(v int32) *AddDnsGtmAccessStrategyRequest
	GetDefaultMaxReturnAddrNum() *int32
	SetDefaultMinAvailableAddrNum(v int32) *AddDnsGtmAccessStrategyRequest
	GetDefaultMinAvailableAddrNum() *int32
	SetFailoverAddrPool(v []*AddDnsGtmAccessStrategyRequestFailoverAddrPool) *AddDnsGtmAccessStrategyRequest
	GetFailoverAddrPool() []*AddDnsGtmAccessStrategyRequestFailoverAddrPool
	SetFailoverAddrPoolType(v string) *AddDnsGtmAccessStrategyRequest
	GetFailoverAddrPoolType() *string
	SetFailoverLatencyOptimization(v string) *AddDnsGtmAccessStrategyRequest
	GetFailoverLatencyOptimization() *string
	SetFailoverLbaStrategy(v string) *AddDnsGtmAccessStrategyRequest
	GetFailoverLbaStrategy() *string
	SetFailoverMaxReturnAddrNum(v int32) *AddDnsGtmAccessStrategyRequest
	GetFailoverMaxReturnAddrNum() *int32
	SetFailoverMinAvailableAddrNum(v int32) *AddDnsGtmAccessStrategyRequest
	GetFailoverMinAvailableAddrNum() *int32
	SetInstanceId(v string) *AddDnsGtmAccessStrategyRequest
	GetInstanceId() *string
	SetLang(v string) *AddDnsGtmAccessStrategyRequest
	GetLang() *string
	SetLines(v string) *AddDnsGtmAccessStrategyRequest
	GetLines() *string
	SetStrategyMode(v string) *AddDnsGtmAccessStrategyRequest
	GetStrategyMode() *string
	SetStrategyName(v string) *AddDnsGtmAccessStrategyRequest
	GetStrategyName() *string
}

type AddDnsGtmAccessStrategyRequest struct {
	// The address pools in the primary address pool set.
	//
	// This parameter is required.
	DefaultAddrPool []*AddDnsGtmAccessStrategyRequestDefaultAddrPool `json:"DefaultAddrPool,omitempty" xml:"DefaultAddrPool,omitempty" type:"Repeated"`
	// The type of the primary address pool. Valid values:
	//
	// 	- IPV4
	//
	// 	- IPV6
	//
	// 	- DOMAIN
	//
	// This parameter is required.
	//
	// example:
	//
	// ipv4
	DefaultAddrPoolType *string `json:"DefaultAddrPoolType,omitempty" xml:"DefaultAddrPoolType,omitempty"`
	// Specifies whether to enable DNS resolution with optimal latency for the primary address pool set. Valid values:
	//
	// 	- OPEN
	//
	// 	- CLOSE
	//
	// example:
	//
	// open
	DefaultLatencyOptimization *string `json:"DefaultLatencyOptimization,omitempty" xml:"DefaultLatencyOptimization,omitempty"`
	// The load balancing policy of the primary address pool set. Valid values:
	//
	// 	- ALL_RR: returns all addresses.
	//
	// 	- RATIO: returns addresses by weight.
	//
	// example:
	//
	// all_rr
	DefaultLbaStrategy *string `json:"DefaultLbaStrategy,omitempty" xml:"DefaultLbaStrategy,omitempty"`
	// The maximum number of addresses returned from the primary address pool set.
	//
	// example:
	//
	// 3
	DefaultMaxReturnAddrNum *int32 `json:"DefaultMaxReturnAddrNum,omitempty" xml:"DefaultMaxReturnAddrNum,omitempty"`
	// The minimum number of available addresses in the primary address pool set.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DefaultMinAvailableAddrNum *int32 `json:"DefaultMinAvailableAddrNum,omitempty" xml:"DefaultMinAvailableAddrNum,omitempty"`
	// The address pools in the secondary address pool set. If no address pool exists in the secondary address pool set, set this parameter to EMPTY.
	FailoverAddrPool []*AddDnsGtmAccessStrategyRequestFailoverAddrPool `json:"FailoverAddrPool,omitempty" xml:"FailoverAddrPool,omitempty" type:"Repeated"`
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
	// Specifies whether to enable DNS resolution with optimal latency for the secondary address pool set. Valid values:
	//
	// 	- OPEN
	//
	// 	- CLOSE
	//
	// example:
	//
	// open
	FailoverLatencyOptimization *string `json:"FailoverLatencyOptimization,omitempty" xml:"FailoverLatencyOptimization,omitempty"`
	// The load balancing policy of the secondary address pool set. Valid values:
	//
	// 	- ALL_RR: returns all addresses.
	//
	// 	- RATIO: returns addresses by weight.
	//
	// example:
	//
	// all_rr
	FailoverLbaStrategy *string `json:"FailoverLbaStrategy,omitempty" xml:"FailoverLbaStrategy,omitempty"`
	// The maximum number of addresses returned from the secondary address pool set.
	//
	// example:
	//
	// 1
	FailoverMaxReturnAddrNum *int32 `json:"FailoverMaxReturnAddrNum,omitempty" xml:"FailoverMaxReturnAddrNum,omitempty"`
	// The minimum number of available addresses in the secondary address pool set.
	//
	// example:
	//
	// 1
	FailoverMinAvailableAddrNum *int32 `json:"FailoverMinAvailableAddrNum,omitempty" xml:"FailoverMinAvailableAddrNum,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance1
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The language of the values for specific response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The Domain Name System (DNS) request source. For example: `["default", "drpeng"]` indicates Global and Dr. Peng Group.
	//
	// example:
	//
	// ["default", "drpeng"]
	Lines *string `json:"Lines,omitempty" xml:"Lines,omitempty"`
	// The type of the access policy. Valid values:
	//
	// 	- GEO: geographical location-based access policy
	//
	// 	- LATENCY: latency-based access policy
	//
	// This parameter is required.
	//
	// example:
	//
	// geo
	StrategyMode *string `json:"StrategyMode,omitempty" xml:"StrategyMode,omitempty"`
	// The name of the access policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// testStrategyName
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s AddDnsGtmAccessStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAccessStrategyRequest) GetDefaultAddrPool() []*AddDnsGtmAccessStrategyRequestDefaultAddrPool {
	return s.DefaultAddrPool
}

func (s *AddDnsGtmAccessStrategyRequest) GetDefaultAddrPoolType() *string {
	return s.DefaultAddrPoolType
}

func (s *AddDnsGtmAccessStrategyRequest) GetDefaultLatencyOptimization() *string {
	return s.DefaultLatencyOptimization
}

func (s *AddDnsGtmAccessStrategyRequest) GetDefaultLbaStrategy() *string {
	return s.DefaultLbaStrategy
}

func (s *AddDnsGtmAccessStrategyRequest) GetDefaultMaxReturnAddrNum() *int32 {
	return s.DefaultMaxReturnAddrNum
}

func (s *AddDnsGtmAccessStrategyRequest) GetDefaultMinAvailableAddrNum() *int32 {
	return s.DefaultMinAvailableAddrNum
}

func (s *AddDnsGtmAccessStrategyRequest) GetFailoverAddrPool() []*AddDnsGtmAccessStrategyRequestFailoverAddrPool {
	return s.FailoverAddrPool
}

func (s *AddDnsGtmAccessStrategyRequest) GetFailoverAddrPoolType() *string {
	return s.FailoverAddrPoolType
}

func (s *AddDnsGtmAccessStrategyRequest) GetFailoverLatencyOptimization() *string {
	return s.FailoverLatencyOptimization
}

func (s *AddDnsGtmAccessStrategyRequest) GetFailoverLbaStrategy() *string {
	return s.FailoverLbaStrategy
}

func (s *AddDnsGtmAccessStrategyRequest) GetFailoverMaxReturnAddrNum() *int32 {
	return s.FailoverMaxReturnAddrNum
}

func (s *AddDnsGtmAccessStrategyRequest) GetFailoverMinAvailableAddrNum() *int32 {
	return s.FailoverMinAvailableAddrNum
}

func (s *AddDnsGtmAccessStrategyRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *AddDnsGtmAccessStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *AddDnsGtmAccessStrategyRequest) GetLines() *string {
	return s.Lines
}

func (s *AddDnsGtmAccessStrategyRequest) GetStrategyMode() *string {
	return s.StrategyMode
}

func (s *AddDnsGtmAccessStrategyRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultAddrPool(v []*AddDnsGtmAccessStrategyRequestDefaultAddrPool) *AddDnsGtmAccessStrategyRequest {
	s.DefaultAddrPool = v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultAddrPoolType(v string) *AddDnsGtmAccessStrategyRequest {
	s.DefaultAddrPoolType = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultLatencyOptimization(v string) *AddDnsGtmAccessStrategyRequest {
	s.DefaultLatencyOptimization = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultLbaStrategy(v string) *AddDnsGtmAccessStrategyRequest {
	s.DefaultLbaStrategy = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultMaxReturnAddrNum(v int32) *AddDnsGtmAccessStrategyRequest {
	s.DefaultMaxReturnAddrNum = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetDefaultMinAvailableAddrNum(v int32) *AddDnsGtmAccessStrategyRequest {
	s.DefaultMinAvailableAddrNum = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverAddrPool(v []*AddDnsGtmAccessStrategyRequestFailoverAddrPool) *AddDnsGtmAccessStrategyRequest {
	s.FailoverAddrPool = v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverAddrPoolType(v string) *AddDnsGtmAccessStrategyRequest {
	s.FailoverAddrPoolType = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverLatencyOptimization(v string) *AddDnsGtmAccessStrategyRequest {
	s.FailoverLatencyOptimization = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverLbaStrategy(v string) *AddDnsGtmAccessStrategyRequest {
	s.FailoverLbaStrategy = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverMaxReturnAddrNum(v int32) *AddDnsGtmAccessStrategyRequest {
	s.FailoverMaxReturnAddrNum = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetFailoverMinAvailableAddrNum(v int32) *AddDnsGtmAccessStrategyRequest {
	s.FailoverMinAvailableAddrNum = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetInstanceId(v string) *AddDnsGtmAccessStrategyRequest {
	s.InstanceId = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetLang(v string) *AddDnsGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetLines(v string) *AddDnsGtmAccessStrategyRequest {
	s.Lines = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetStrategyMode(v string) *AddDnsGtmAccessStrategyRequest {
	s.StrategyMode = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) SetStrategyName(v string) *AddDnsGtmAccessStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequest) Validate() error {
	if s.DefaultAddrPool != nil {
		for _, item := range s.DefaultAddrPool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type AddDnsGtmAccessStrategyRequestDefaultAddrPool struct {
	// The ID of the address pool in the primary address pool set.
	//
	// example:
	//
	// pool1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The weight of the address pool in the primary address pool set.
	//
	// example:
	//
	// 1
	LbaWeight *int32 `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
}

func (s AddDnsGtmAccessStrategyRequestDefaultAddrPool) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmAccessStrategyRequestDefaultAddrPool) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAccessStrategyRequestDefaultAddrPool) GetId() *string {
	return s.Id
}

func (s *AddDnsGtmAccessStrategyRequestDefaultAddrPool) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *AddDnsGtmAccessStrategyRequestDefaultAddrPool) SetId(v string) *AddDnsGtmAccessStrategyRequestDefaultAddrPool {
	s.Id = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequestDefaultAddrPool) SetLbaWeight(v int32) *AddDnsGtmAccessStrategyRequestDefaultAddrPool {
	s.LbaWeight = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequestDefaultAddrPool) Validate() error {
	return dara.Validate(s)
}

type AddDnsGtmAccessStrategyRequestFailoverAddrPool struct {
	// The ID of the address pool in the secondary address pool set.
	//
	// example:
	//
	// pool1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The weight of the address pool in the secondary address pool set.
	//
	// example:
	//
	// 1
	LbaWeight *int32 `json:"LbaWeight,omitempty" xml:"LbaWeight,omitempty"`
}

func (s AddDnsGtmAccessStrategyRequestFailoverAddrPool) String() string {
	return dara.Prettify(s)
}

func (s AddDnsGtmAccessStrategyRequestFailoverAddrPool) GoString() string {
	return s.String()
}

func (s *AddDnsGtmAccessStrategyRequestFailoverAddrPool) GetId() *string {
	return s.Id
}

func (s *AddDnsGtmAccessStrategyRequestFailoverAddrPool) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *AddDnsGtmAccessStrategyRequestFailoverAddrPool) SetId(v string) *AddDnsGtmAccessStrategyRequestFailoverAddrPool {
	s.Id = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequestFailoverAddrPool) SetLbaWeight(v int32) *AddDnsGtmAccessStrategyRequestFailoverAddrPool {
	s.LbaWeight = &v
	return s
}

func (s *AddDnsGtmAccessStrategyRequestFailoverAddrPool) Validate() error {
	return dara.Validate(s)
}
