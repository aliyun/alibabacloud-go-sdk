// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmAddressPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressPools(v *SearchCloudGtmAddressPoolsResponseBodyAddressPools) *SearchCloudGtmAddressPoolsResponseBody
	GetAddressPools() *SearchCloudGtmAddressPoolsResponseBodyAddressPools
	SetPageNumber(v int32) *SearchCloudGtmAddressPoolsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *SearchCloudGtmAddressPoolsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *SearchCloudGtmAddressPoolsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *SearchCloudGtmAddressPoolsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *SearchCloudGtmAddressPoolsResponseBody
	GetTotalPages() *int32
}

type SearchCloudGtmAddressPoolsResponseBody struct {
	// The address pools.
	AddressPools *SearchCloudGtmAddressPoolsResponseBodyAddressPools `json:"AddressPools,omitempty" xml:"AddressPools,omitempty" type:"Struct"`
	// Current page number, starting from 1, default is 1.
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
	// Total number of address pools matching the query conditions.
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

func (s SearchCloudGtmAddressPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressPoolsResponseBody) GetAddressPools() *SearchCloudGtmAddressPoolsResponseBodyAddressPools {
	return s.AddressPools
}

func (s *SearchCloudGtmAddressPoolsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *SearchCloudGtmAddressPoolsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *SearchCloudGtmAddressPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SearchCloudGtmAddressPoolsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *SearchCloudGtmAddressPoolsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *SearchCloudGtmAddressPoolsResponseBody) SetAddressPools(v *SearchCloudGtmAddressPoolsResponseBodyAddressPools) *SearchCloudGtmAddressPoolsResponseBody {
	s.AddressPools = v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBody) SetPageNumber(v int32) *SearchCloudGtmAddressPoolsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBody) SetPageSize(v int32) *SearchCloudGtmAddressPoolsResponseBody {
	s.PageSize = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBody) SetRequestId(v string) *SearchCloudGtmAddressPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBody) SetTotalItems(v int32) *SearchCloudGtmAddressPoolsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBody) SetTotalPages(v int32) *SearchCloudGtmAddressPoolsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBody) Validate() error {
	return dara.Validate(s)
}

type SearchCloudGtmAddressPoolsResponseBodyAddressPools struct {
	AddressPool []*SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool `json:"AddressPool,omitempty" xml:"AddressPool,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPools) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPools) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPools) GetAddressPool() []*SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	return s.AddressPool
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPools) SetAddressPool(v []*SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) *SearchCloudGtmAddressPoolsResponseBodyAddressPools {
	s.AddressPool = v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPools) Validate() error {
	return dara.Validate(s)
}

type SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool struct {
	// Load balancing policy among addresses in the address pool:
	//
	// - round_robin: Round-robin, for any source of DNS resolution requests, returns all addresses and rotates the order of all addresses each time.
	//
	// - sequence: Sequential, for any source of DNS resolution requests, returns the address with the smaller sequence number (the sequence number indicates the priority of the address return, the smaller the higher the priority). If the address with the smaller sequence number is unavailable, return the next address with a smaller sequence number.
	//
	// - weight: Weighted, supports setting different weight values for each address to realize returning addresses according to the weight ratio for resolution queries.
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
	// The IP addresses or domain names.
	Addresses *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Struct"`
	// Address pool availability status:
	//
	// - available: Available
	//
	// - unavailable: Unavailable
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
	// Address remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The mode used if the address with the smallest sequence number is recovered. This parameter is returned only when the policy for load balancing between addresses is sequence. Valid values:
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
	// 024-03-15T01:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Last modification time of the address pool (timestamp).
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAddressLbStrategy() *string {
	return s.AddressLbStrategy
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAddressPoolType() *string {
	return s.AddressPoolType
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAddresses() *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses {
	return s.Addresses
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetRemark() *string {
	return s.Remark
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAddressLbStrategy(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.AddressLbStrategy = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAddressPoolId(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.AddressPoolId = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAddressPoolName(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.AddressPoolName = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAddressPoolType(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.AddressPoolType = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAddresses(v *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.Addresses = v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetAvailableStatus(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.AvailableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetCreateTime(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.CreateTime = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetCreateTimestamp(v int64) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetEnableStatus(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.EnableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetHealthJudgement(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.HealthJudgement = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetHealthStatus(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.HealthStatus = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetRemark(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.Remark = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetSequenceLbStrategyMode(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetUpdateTime(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.UpdateTime = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) SetUpdateTimestamp(v int64) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool) Validate() error {
	return dara.Validate(s)
}

type SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses struct {
	Address []*SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) GetAddress() []*SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	return s.Address
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) SetAddress(v []*SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses {
	s.Address = v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses) Validate() error {
	return dara.Validate(s)
}

type SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress struct {
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
	// Not supported in the current version.
	AttributeInfo *string `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	// The failover method that is used if the address fails health checks. Valid values:
	//
	// 	- auto: the automatic mode. The system determines whether to return an address based on the health check results. If the address fails health checks, the system does not return the address. If the address passes health checks, the system returns the address.
	//
	// 	- manual: the manual mode. If an address is in the unavailable state, the address is not returned for DNS requests even if the address passes health checks. If an address is in the available state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// example:
	//
	// auto
	AvailableMode *string `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	// Address availability status:
	//
	// - available: Address is available
	//
	// - unavailable: Address is unavailable
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
	// Address enable status, indicating whether the address is currently available:
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
	// 	- ok_no_monitor: The address does not reference a health check template and is normal.
	//
	// 	- exceptional: The address fails some or all health checks of the referenced health check templates and the address is deemed abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The health check tasks.
	HealthTasks *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Struct"`
	// The availability state of the address when AvailableMode is set to manual. Valid values:
	//
	// 	- available: The address is normal. In this state, the address is returned for DNS requests even if an alert is triggered when the address fails health checks.
	//
	// 	- unavailable: The address is abnormal. In this state, the address is not returned for DNS requests even if the address passes health checks.
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
	// Remarks.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Request source, referring to the source of the request. GTM schedules based on the exit IP of the LocalDNS used by the terminal. If the LocalDNS supports ECS (edns-client-subnet), intelligent scheduling can also be performed based on the IP of the visiting terminal.
	//
	// example:
	//
	// default
	RequestSource *string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
	// Sequence number, indicating the priority of address return, where smaller numbers have higher priority.
	//
	// example:
	//
	// 1
	SerialNumber *int32 `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// Address type:
	//
	// - IPv4: ipv4 address - IPv6: ipv6 address - domain: domain name
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
	// Weight value (an integer between 1 and 100), allowing different weight values to be set for each address, enabling resolution queries to return addresses according to the weighted ratio.
	//
	// example:
	//
	// 1
	WeightValue *int32 `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAddress() *string {
	return s.Address
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAddressId() *string {
	return s.AddressId
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAvailableMode() *string {
	return s.AvailableMode
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetCreateTime() *string {
	return s.CreateTime
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetHealthTasks() *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks {
	return s.HealthTasks
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetManualAvailableStatus() *string {
	return s.ManualAvailableStatus
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetName() *string {
	return s.Name
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetRemark() *string {
	return s.Remark
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetRequestSource() *string {
	return s.RequestSource
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetSerialNumber() *int32 {
	return s.SerialNumber
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetType() *string {
	return s.Type
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) GetWeightValue() *int32 {
	return s.WeightValue
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAddress(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Address = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAddressId(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AddressId = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAttributeInfo(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AttributeInfo = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAvailableMode(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AvailableMode = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetAvailableStatus(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.AvailableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetCreateTime(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.CreateTime = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetCreateTimestamp(v int64) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.CreateTimestamp = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetEnableStatus(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.EnableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetHealthJudgement(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.HealthJudgement = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetHealthStatus(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.HealthStatus = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetHealthTasks(v *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.HealthTasks = v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetManualAvailableStatus(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.ManualAvailableStatus = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetName(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Name = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetRemark(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Remark = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetRequestSource(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.RequestSource = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetSerialNumber(v int32) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.SerialNumber = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetType(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.Type = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetUpdateTime(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.UpdateTime = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetUpdateTimestamp(v int64) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.UpdateTimestamp = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) SetWeightValue(v int32) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress {
	s.WeightValue = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress) Validate() error {
	return dara.Validate(s)
}

type SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks struct {
	HealthTask []*SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask `json:"HealthTask,omitempty" xml:"HealthTask,omitempty" type:"Repeated"`
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) GetHealthTask() []*SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask {
	return s.HealthTask
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) SetHealthTask(v []*SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks {
	s.HealthTask = v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks) Validate() error {
	return dara.Validate(s)
}

type SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask struct {
	// The target service port for health checks. When the Ping protocol is selected for health checks, configuration of the service port is not supported.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// The ID of the health check template associated with the address.
	//
	// example:
	//
	// mtp-89518052425100**80
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// Health check template name.
	//
	// example:
	//
	// IPv4-Ping
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) GetPort() *int32 {
	return s.Port
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) GetTemplateId() *string {
	return s.TemplateId
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) GetTemplateName() *string {
	return s.TemplateName
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) SetPort(v int32) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask {
	s.Port = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) SetTemplateId(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask {
	s.TemplateId = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) SetTemplateName(v string) *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask {
	s.TemplateName = &v
	return s
}

func (s *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask) Validate() error {
	return dara.Validate(s)
}
