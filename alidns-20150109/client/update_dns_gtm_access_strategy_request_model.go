// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDnsGtmAccessStrategyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessMode(v string) *UpdateDnsGtmAccessStrategyRequest
	GetAccessMode() *string
	SetDefaultAddrPool(v []*UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) *UpdateDnsGtmAccessStrategyRequest
	GetDefaultAddrPool() []*UpdateDnsGtmAccessStrategyRequestDefaultAddrPool
	SetDefaultAddrPoolType(v string) *UpdateDnsGtmAccessStrategyRequest
	GetDefaultAddrPoolType() *string
	SetDefaultLatencyOptimization(v string) *UpdateDnsGtmAccessStrategyRequest
	GetDefaultLatencyOptimization() *string
	SetDefaultLbaStrategy(v string) *UpdateDnsGtmAccessStrategyRequest
	GetDefaultLbaStrategy() *string
	SetDefaultMaxReturnAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest
	GetDefaultMaxReturnAddrNum() *int32
	SetDefaultMinAvailableAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest
	GetDefaultMinAvailableAddrNum() *int32
	SetFailoverAddrPool(v []*UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) *UpdateDnsGtmAccessStrategyRequest
	GetFailoverAddrPool() []*UpdateDnsGtmAccessStrategyRequestFailoverAddrPool
	SetFailoverAddrPoolType(v string) *UpdateDnsGtmAccessStrategyRequest
	GetFailoverAddrPoolType() *string
	SetFailoverLatencyOptimization(v string) *UpdateDnsGtmAccessStrategyRequest
	GetFailoverLatencyOptimization() *string
	SetFailoverLbaStrategy(v string) *UpdateDnsGtmAccessStrategyRequest
	GetFailoverLbaStrategy() *string
	SetFailoverMaxReturnAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest
	GetFailoverMaxReturnAddrNum() *int32
	SetFailoverMinAvailableAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest
	GetFailoverMinAvailableAddrNum() *int32
	SetLang(v string) *UpdateDnsGtmAccessStrategyRequest
	GetLang() *string
	SetLines(v string) *UpdateDnsGtmAccessStrategyRequest
	GetLines() *string
	SetStrategyId(v string) *UpdateDnsGtmAccessStrategyRequest
	GetStrategyId() *string
	SetStrategyName(v string) *UpdateDnsGtmAccessStrategyRequest
	GetStrategyName() *string
}

type UpdateDnsGtmAccessStrategyRequest struct {
	// The primary/secondary switchover policy for address pool sets. Valid values:
	//
	// 	- AUTO: performs automatic switchover between the primary and secondary address pool sets upon failures.
	//
	// 	- DEFAULT: the primary address pool set
	//
	// 	- FAILOVER: the secondary address pool set
	//
	// example:
	//
	// DEFAULT
	AccessMode *string `json:"AccessMode,omitempty" xml:"AccessMode,omitempty"`
	// The address pools in the primary address pool set.
	//
	// This parameter is required.
	DefaultAddrPool []*UpdateDnsGtmAccessStrategyRequestDefaultAddrPool `json:"DefaultAddrPool,omitempty" xml:"DefaultAddrPool,omitempty" type:"Repeated"`
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
	// Specifies whether to enable Domain Name System (DNS) resolution with optimal latency for the primary address pool set. Valid values:
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
	// 1
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
	FailoverAddrPool []*UpdateDnsGtmAccessStrategyRequestFailoverAddrPool `json:"FailoverAddrPool,omitempty" xml:"FailoverAddrPool,omitempty" type:"Repeated"`
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
	// The language of the values for specific response parameters. Default value: en. Valid values: en, zh, and ja.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The line codes of the source regions. Example: `["default", "drpeng"]`, which indicates the global line and Dr. Peng Group line.
	//
	// example:
	//
	// ["default", "drpeng"]
	Lines *string `json:"Lines,omitempty" xml:"Lines,omitempty"`
	// The ID of the access policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// StrategyId1
	StrategyId *string `json:"StrategyId,omitempty" xml:"StrategyId,omitempty"`
	// The name of the access policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// StrategyName1
	StrategyName *string `json:"StrategyName,omitempty" xml:"StrategyName,omitempty"`
}

func (s UpdateDnsGtmAccessStrategyRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmAccessStrategyRequest) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetAccessMode() *string {
	return s.AccessMode
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetDefaultAddrPool() []*UpdateDnsGtmAccessStrategyRequestDefaultAddrPool {
	return s.DefaultAddrPool
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetDefaultAddrPoolType() *string {
	return s.DefaultAddrPoolType
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetDefaultLatencyOptimization() *string {
	return s.DefaultLatencyOptimization
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetDefaultLbaStrategy() *string {
	return s.DefaultLbaStrategy
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetDefaultMaxReturnAddrNum() *int32 {
	return s.DefaultMaxReturnAddrNum
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetDefaultMinAvailableAddrNum() *int32 {
	return s.DefaultMinAvailableAddrNum
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetFailoverAddrPool() []*UpdateDnsGtmAccessStrategyRequestFailoverAddrPool {
	return s.FailoverAddrPool
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetFailoverAddrPoolType() *string {
	return s.FailoverAddrPoolType
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetFailoverLatencyOptimization() *string {
	return s.FailoverLatencyOptimization
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetFailoverLbaStrategy() *string {
	return s.FailoverLbaStrategy
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetFailoverMaxReturnAddrNum() *int32 {
	return s.FailoverMaxReturnAddrNum
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetFailoverMinAvailableAddrNum() *int32 {
	return s.FailoverMinAvailableAddrNum
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetLines() *string {
	return s.Lines
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetStrategyId() *string {
	return s.StrategyId
}

func (s *UpdateDnsGtmAccessStrategyRequest) GetStrategyName() *string {
	return s.StrategyName
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetAccessMode(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.AccessMode = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultAddrPool(v []*UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultAddrPool = v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultAddrPoolType(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultAddrPoolType = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultLatencyOptimization(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultLatencyOptimization = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultLbaStrategy(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultLbaStrategy = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultMaxReturnAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultMaxReturnAddrNum = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetDefaultMinAvailableAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest {
	s.DefaultMinAvailableAddrNum = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverAddrPool(v []*UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverAddrPool = v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverAddrPoolType(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverAddrPoolType = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverLatencyOptimization(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverLatencyOptimization = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverLbaStrategy(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverLbaStrategy = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverMaxReturnAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverMaxReturnAddrNum = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetFailoverMinAvailableAddrNum(v int32) *UpdateDnsGtmAccessStrategyRequest {
	s.FailoverMinAvailableAddrNum = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetLang(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetLines(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.Lines = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetStrategyId(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.StrategyId = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) SetStrategyName(v string) *UpdateDnsGtmAccessStrategyRequest {
	s.StrategyName = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequest) Validate() error {
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

type UpdateDnsGtmAccessStrategyRequestDefaultAddrPool struct {
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

func (s UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) GetId() *string {
	return s.Id
}

func (s *UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) SetId(v string) *UpdateDnsGtmAccessStrategyRequestDefaultAddrPool {
	s.Id = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) SetLbaWeight(v int32) *UpdateDnsGtmAccessStrategyRequestDefaultAddrPool {
	s.LbaWeight = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequestDefaultAddrPool) Validate() error {
	return dara.Validate(s)
}

type UpdateDnsGtmAccessStrategyRequestFailoverAddrPool struct {
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

func (s UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) String() string {
	return dara.Prettify(s)
}

func (s UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) GoString() string {
	return s.String()
}

func (s *UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) GetId() *string {
	return s.Id
}

func (s *UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) GetLbaWeight() *int32 {
	return s.LbaWeight
}

func (s *UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) SetId(v string) *UpdateDnsGtmAccessStrategyRequestFailoverAddrPool {
	s.Id = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) SetLbaWeight(v int32) *UpdateDnsGtmAccessStrategyRequestFailoverAddrPool {
	s.LbaWeight = &v
	return s
}

func (s *UpdateDnsGtmAccessStrategyRequestFailoverAddrPool) Validate() error {
	return dara.Validate(s)
}
