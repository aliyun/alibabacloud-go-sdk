// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAddressPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPools(v *ListCloudGtmAddressPoolsResponseBodyAddressPools) *ListCloudGtmAddressPoolsResponseBody
	GetAddressPools() *ListCloudGtmAddressPoolsResponseBodyAddressPools
	SetPageNumber(v int32) *ListCloudGtmAddressPoolsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCloudGtmAddressPoolsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCloudGtmAddressPoolsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *ListCloudGtmAddressPoolsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *ListCloudGtmAddressPoolsResponseBody
	GetTotalPages() *int32
}

type ListCloudGtmAddressPoolsResponseBody struct {
	// The address pools.
	AddressPools *ListCloudGtmAddressPoolsResponseBodyAddressPools `json:"AddressPools,omitempty" xml:"AddressPools,omitempty" type:"Struct"`
	// Current page number, starting at **1**, default is **1**.
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
	// Total number of entries in the address pool.
	//
	// example:
	//
	// 11
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s ListCloudGtmAddressPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressPoolsResponseBody) GetAddressPools() *ListCloudGtmAddressPoolsResponseBodyAddressPools {
	return s.AddressPools
}

func (s *ListCloudGtmAddressPoolsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCloudGtmAddressPoolsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudGtmAddressPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudGtmAddressPoolsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *ListCloudGtmAddressPoolsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *ListCloudGtmAddressPoolsResponseBody) SetAddressPools(v *ListCloudGtmAddressPoolsResponseBodyAddressPools) *ListCloudGtmAddressPoolsResponseBody {
	s.AddressPools = v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBody) SetPageNumber(v int32) *ListCloudGtmAddressPoolsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBody) SetPageSize(v int32) *ListCloudGtmAddressPoolsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBody) SetRequestId(v string) *ListCloudGtmAddressPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBody) SetTotalItems(v int32) *ListCloudGtmAddressPoolsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBody) SetTotalPages(v int32) *ListCloudGtmAddressPoolsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBody) Validate() error {
	if s.AddressPools != nil {
		if err := s.AddressPools.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudGtmAddressPoolsResponseBodyAddressPools struct {
	AddressPool []*ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool `json:"AddressPool,omitempty" xml:"AddressPool,omitempty" type:"Repeated"`
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPools) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPools) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPools) GetAddressPool() []*ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	return s.AddressPool
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPools) SetAddressPool(v []*ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) *ListCloudGtmAddressPoolsResponseBodyAddressPools {
	s.AddressPool = v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPools) Validate() error {
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

type ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool struct {
	// Load balancing policy among addresses in the address pool:
	//
	// - round_robin: Round-robin, for any source of DNS resolution requests, all addresses are returned. The order of all addresses is rotated each time.
	//
	// - sequence: Sequential, for any source of DNS resolution requests, the address with the smaller sequence number (the sequence number indicates the priority of address returns, with smaller numbers having higher priority) is returned. If the address with the smaller sequence number is unavailable, the next address with a smaller sequence number is returned.
	//
	// - weight: Weighted, supports setting different weight values for each address, realizing the return of addresses according to the ratio of weights in resolution queries.
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
	// The addresses.
	Addresses *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Struct"`
	// The availability state of the address pool. Valid values:
	//
	// 	- Available: The address pool is available.
	//
	// 	- unavailable: The address pool is unavailable.
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
	// The enabling state of the address pool. Valid values:
	//
	// 	- enable: The address pool is enabled.
	//
	// 	- disable: The address pool is disabled.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The condition for determining the health state of the address. Valid values:
	//
	// 	- any_ok: The health check results of at least one health check template are normal.
	//
	// 	- p30_ok: The health check results of at least 30% of health check templates are normal.
	//
	// 	- p50_ok: The health check results of at least 50% of health check templates are normal.
	//
	// 	- p70_ok: The health check results of at least 70% of health check templates are normal.
	//
	// 	- all_ok: The health check results of all health check templates are normal.
	//
	// example:
	//
	// any_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// The health state of the address pool. Valid values:
	//
	// 	- ok: The health state of the address pool is Normal and all addresses that are referenced by the address pool are available.
	//
	// 	- ok_alert: The health state of the address pool is Warning and some of the addresses that are referenced by the address pool are unavailable. However, the address pool is deemed normal. In this state, available address pools are normally used for DNS resolution, but unavailable address pools cannot be used for DNS resolution.
	//
	// 	- exceptional: The health state of the address pool is Abnormal and some or all of the addresses that are referenced by the address pool are unavailable. In this case, the address pool is deemed abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// Remark
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
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
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAddressLbStrategy() *string {
	return s.AddressLbStrategy
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAddressPoolType() *string {
	return s.AddressPoolType
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAddresses() *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses {
	return s.Addresses
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetRemark() *string {
	return s.Remark
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAddressLbStrategy(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.AddressLbStrategy = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAddressPoolId(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.AddressPoolId = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAddressPoolName(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.AddressPoolName = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAddressPoolType(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.AddressPoolType = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAddresses(v *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.Addresses = v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAvailableStatus(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.AvailableStatus = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetCreateTime(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.CreateTime = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetCreateTimestamp(v int64) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.CreateTimestamp = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetEnableStatus(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.EnableStatus = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetHealthJudgement(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.HealthJudgement = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetHealthStatus(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.HealthStatus = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetRemark(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.Remark = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetSequenceLbStrategyMode(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetUpdateTime(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetUpdateTimestamp(v int64) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) Validate() error {
	if s.Addresses != nil {
		if err := s.Addresses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses struct {
	Address []*ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) GetAddress() []*ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	return s.Address
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) SetAddress(v []*ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses {
	s.Address = v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) Validate() error {
	if s.Address != nil {
		for _, item := range s.Address {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress struct {
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
	// addr-895182181143688192
	AddressId *string `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	// Address ownership information, not supported in the current version.
	//
	// example:
	//
	// The current version does not support returning this parameter.
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The failover mode that is used when address exceptions are identified. Valid values:
	//
	// 	- auto: the automatic mode. The system determines whether to return an address based on the health check results. If the address fails health checks, the system does not return the address. If the address passes health checks, the system returns the address.
	//
	// 	- manual: the manual mode. If an address is in the unavailable state, the address is not returned for DNS requests even if the address passes health checks. If an address is in the available state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// example:
	//
	// auto
	AvailableMode *string `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	// The availability state of the address. Valid values:
	//
	// 	- available: The address is available.
	//
	// 	- unavailable: The address is unavailable.
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// Address creation time.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Address creation time (timestamp).
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// Address enable status:
	//
	// - enable: Enabled status
	//
	// - disable: Disabled status
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The condition for determining the health status of the address. Valid values:
	//
	// 	- any_ok: The health check results of at least one health check template are normal.
	//
	// 	- p30_ok: The health check results of at least 30% of health check templates are normal.
	//
	// 	- p50_ok: The health check results of at least 50% of health check templates are normal.
	//
	// 	- p70_ok: The health check results of at least 70% of health check templates are normal.
	//
	// 	- all_ok: The health check results of all health check templates are normal.
	//
	// example:
	//
	// any_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// The health check state of the address. Valid values:
	//
	// 	- ok: The address passes all health checks of the referenced health check templates.
	//
	// 	- ok_alert: The address fails some health checks of the referenced health check templates but the address is deemed normal.
	//
	// 	- ok_no_monitor: The address does not reference any health check template and is normal.
	//
	// 	- exceptional: The address fails some or all health checks of the referenced health check templates and the address is deemed abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The health check tasks.
	HealthTasks *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Struct"`
	// The availability state of the address when AvailableMode is set to manual for the address. Valid values:
	//
	// 	- available: The address is available. In this state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// 	- unavailable: The address is unavailable. In this state, the address is not returned for DNS requests even if the address passes health checks.
	//
	// example:
	//
	// available
	ManualAvailableStatus *string `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
	// Address name.
	//
	// example:
	//
	// Address-1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Address remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// List of request sources.
	RequestSource *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Struct"`
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
	// Sequence number, indicating the priority of address return, where smaller numbers have higher priority.
	//
	// example:
	//
	// 1
	SerialNumber *int32 `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// Address type:
	//
	// - IPv4: IPv4 address
	//
	// - IPv6: IPv6 address
	//
	// - domain: Domain name
	//
	// example:
	//
	// IPv4
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The last time the address was modified.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The last modification time of the address (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	// Weight value (integer between 1 and 100), supports setting different weight values for each address, enabling resolution queries to return addresses according to the weight ratio.
	//
	// example:
	//
	// 1
	WeightValue *int32 `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAddress() *string {
	return s.Address
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAddressId() *string {
	return s.AddressId
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAvailableMode() *string {
	return s.AvailableMode
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetHealthTasks() *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks {
	return s.HealthTasks
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetManualAvailableStatus() *string {
	return s.ManualAvailableStatus
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetName() *string {
	return s.Name
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetRemark() *string {
	return s.Remark
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetRequestSource() *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource {
	return s.RequestSource
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetSeqNonPreemptiveSchedule() *bool {
	return s.SeqNonPreemptiveSchedule
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetSerialNumber() *int32 {
	return s.SerialNumber
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetType() *string {
	return s.Type
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetWeightValue() *int32 {
	return s.WeightValue
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAddress(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Address = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAddressId(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AddressId = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAttributeInfo(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AttributeInfo = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAvailableMode(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AvailableMode = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAvailableStatus(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AvailableStatus = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetCreateTime(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.CreateTime = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetCreateTimestamp(v int64) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.CreateTimestamp = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetEnableStatus(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.EnableStatus = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetHealthJudgement(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.HealthJudgement = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetHealthStatus(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.HealthStatus = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetHealthTasks(v *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.HealthTasks = v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetManualAvailableStatus(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.ManualAvailableStatus = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetName(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Name = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetRemark(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Remark = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetRequestSource(v *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.RequestSource = v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetSeqNonPreemptiveSchedule(v bool) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.SeqNonPreemptiveSchedule = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetSerialNumber(v int32) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.SerialNumber = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetType(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Type = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetUpdateTime(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.UpdateTime = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetUpdateTimestamp(v int64) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetWeightValue(v int32) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.WeightValue = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) Validate() error {
	if s.HealthTasks != nil {
		if err := s.HealthTasks.Validate(); err != nil {
			return err
		}
	}
	if s.RequestSource != nil {
		if err := s.RequestSource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks struct {
	HealthTask []*ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask `json:"HealthTask,omitempty" xml:"HealthTask,omitempty" type:"Repeated"`
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) GetHealthTask() []*ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask {
	return s.HealthTask
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) SetHealthTask(v []*ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks {
	s.HealthTask = v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) Validate() error {
	if s.HealthTask != nil {
		for _, item := range s.HealthTask {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask struct {
	// The target service port for health checks. When the Ping protocol is selected for health checks, configuration of the service port is not supported.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the health check template.
	//
	// example:
	//
	// mtp-895180524251002880
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Health check template name.
	//
	// example:
	//
	// IPv4-Ping
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) GetPort() *int32 {
	return s.Port
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) GetTemplateName() *string {
	return s.TemplateName
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) SetPort(v int32) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask {
	s.Port = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) SetTemplateId(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask {
	s.TemplateId = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) SetTemplateName(v string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask {
	s.TemplateName = &v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) Validate() error {
	return dara.Validate(s)
}

type ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource struct {
	RequestSource []*string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Repeated"`
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) GetRequestSource() []*string {
	return s.RequestSource
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) SetRequestSource(v []*string) *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource {
	s.RequestSource = v
	return s
}

func (s *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource) Validate() error {
	return dara.Validate(s)
}
