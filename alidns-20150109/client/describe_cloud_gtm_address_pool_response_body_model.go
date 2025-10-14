// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudGtmAddressPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddressLbStrategy(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetAddressLbStrategy() *string
	SetAddressPoolId(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetAddressPoolId() *string
	SetAddressPoolName(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetAddressPoolName() *string
	SetAddressPoolType(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetAddressPoolType() *string
	SetAddresses(v *DescribeCloudGtmAddressPoolResponseBodyAddresses) *DescribeCloudGtmAddressPoolResponseBody
	GetAddresses() *DescribeCloudGtmAddressPoolResponseBodyAddresses
	SetAvailableStatus(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetAvailableStatus() *string
	SetCreateTime(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetCreateTime() *string
	SetCreateTimestamp(v int64) *DescribeCloudGtmAddressPoolResponseBody
	GetCreateTimestamp() *int64
	SetEnableStatus(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetEnableStatus() *string
	SetHealthJudgement(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetHealthJudgement() *string
	SetHealthStatus(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetHealthStatus() *string
	SetRemark(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetRemark() *string
	SetRequestId(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetRequestId() *string
	SetSequenceLbStrategyMode(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetSequenceLbStrategyMode() *string
	SetUpdateTime(v string) *DescribeCloudGtmAddressPoolResponseBody
	GetUpdateTime() *string
	SetUpdateTimestamp(v int64) *DescribeCloudGtmAddressPoolResponseBody
	GetUpdateTimestamp() *int64
}

type DescribeCloudGtmAddressPoolResponseBody struct {
	// Load balancing policy among addresses in the address pool:
	//
	// - round_robin: Round-robin, where for any source of DNS resolution requests, all addresses are returned, with a rotation of the order for every request.
	//
	// - sequence: Sequential, where for any source of DNS resolution requests, the address with the lower sequence number (indicating a higher priority, the smaller the number, the higher the priority) is returned. If the address with the lower sequence number is unavailable, the next address with a lower sequence number is returned.
	//
	// - weight: Weighted, supporting the setting of different weight values for each address to realize returning addresses according to the weight ratio of query resolutions.
	//
	// - source_nearest: Source-nearest, also known as intelligent resolution, where GTM can return different addresses based on the source of different DNS resolution requests, achieving the effect of users accessing nearby servers.
	//
	// example:
	//
	// round_robin
	AddressLbStrategy *string `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	// The ID of the address pool. This ID uniquely identifies the address pool.
	//
	// example:
	//
	// pool-89564674533755**96
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
	Addresses *DescribeCloudGtmAddressPoolResponseBodyAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Struct"`
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
	// 	- ok_alert: The health state of the address pool is warning and some of the addresses that are referenced by the address pool are unavailable. However, the address pool is deemed normal. In this case, only the available addresses are returned for Domain Name System (DNS) requests.
	//
	// 	- exceptional: The health state of the address pool is abnormal and some or all of the addresses that are referenced by the address pool are unavailable. In this case, the address pool is deemed abnormal.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// Remarks for the address.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// Unique request identification code.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	// The last modification time of the address pool.
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

func (s DescribeCloudGtmAddressPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetAddressLbStrategy() *string {
	return s.AddressLbStrategy
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetAddressPoolId() *string {
	return s.AddressPoolId
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetAddressPoolName() *string {
	return s.AddressPoolName
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetAddressPoolType() *string {
	return s.AddressPoolType
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetAddresses() *DescribeCloudGtmAddressPoolResponseBodyAddresses {
	return s.Addresses
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetSequenceLbStrategyMode() *string {
	return s.SequenceLbStrategyMode
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeCloudGtmAddressPoolResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetAddressLbStrategy(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.AddressLbStrategy = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetAddressPoolId(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.AddressPoolId = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetAddressPoolName(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.AddressPoolName = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetAddressPoolType(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.AddressPoolType = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetAddresses(v *DescribeCloudGtmAddressPoolResponseBodyAddresses) *DescribeCloudGtmAddressPoolResponseBody {
	s.Addresses = v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetAvailableStatus(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.AvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetCreateTime(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetCreateTimestamp(v int64) *DescribeCloudGtmAddressPoolResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetEnableStatus(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.EnableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetHealthJudgement(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.HealthJudgement = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetHealthStatus(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.HealthStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetRemark(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.Remark = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetRequestId(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetSequenceLbStrategyMode(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.SequenceLbStrategyMode = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetUpdateTime(v string) *DescribeCloudGtmAddressPoolResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) SetUpdateTimestamp(v int64) *DescribeCloudGtmAddressPoolResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBody) Validate() error {
	if s.Addresses != nil {
		if err := s.Addresses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCloudGtmAddressPoolResponseBodyAddresses struct {
	Address []*DescribeCloudGtmAddressPoolResponseBodyAddressesAddress `json:"Address,omitempty" xml:"Address,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmAddressPoolResponseBodyAddresses) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolResponseBodyAddresses) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddresses) GetAddress() []*DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	return s.Address
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddresses) SetAddress(v []*DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) *DescribeCloudGtmAddressPoolResponseBodyAddresses {
	s.Address = v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddresses) Validate() error {
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

type DescribeCloudGtmAddressPoolResponseBodyAddressesAddress struct {
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
	// addr-89518218114368**92
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
	// - available: Address available
	//
	// - unavailable: Address unavailable
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
	// Health check task list.
	HealthTasks *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Struct"`
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
	// Remarks for the address.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The request source list.
	RequestSource *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressRequestSource `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Struct"`
	// Indicates whether it is a sequential (non-preemptive) scheduling object for hybrid cloud management scenarios: - true: yes - false: no
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
	// The address type. Valid values:
	//
	// 	- IPv4: IPv4 address
	//
	// 	- IPv6: IPv6 address
	//
	// 	- domain: domain name
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
	// Weight value (an integer between 1 and 100, inclusive), allowing different weight values to be set for each address, enabling resolution queries to return addresses according to the weighted ratio.
	//
	// example:
	//
	// 1
	WeightValue *int32 `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
}

func (s DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetAddress() *string {
	return s.Address
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetAddressId() *string {
	return s.AddressId
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetAttributeInfo() *string {
	return s.AttributeInfo
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetAvailableMode() *string {
	return s.AvailableMode
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetAvailableStatus() *string {
	return s.AvailableStatus
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetHealthJudgement() *string {
	return s.HealthJudgement
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetHealthStatus() *string {
	return s.HealthStatus
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetHealthTasks() *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasks {
	return s.HealthTasks
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetManualAvailableStatus() *string {
	return s.ManualAvailableStatus
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetName() *string {
	return s.Name
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetRemark() *string {
	return s.Remark
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetRequestSource() *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressRequestSource {
	return s.RequestSource
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetSeqNonPreemptiveSchedule() *bool {
	return s.SeqNonPreemptiveSchedule
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetSerialNumber() *int32 {
	return s.SerialNumber
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetType() *string {
	return s.Type
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) GetWeightValue() *int32 {
	return s.WeightValue
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetAddress(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.Address = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetAddressId(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.AddressId = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetAttributeInfo(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.AttributeInfo = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetAvailableMode(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.AvailableMode = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetAvailableStatus(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.AvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetCreateTime(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.CreateTime = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetCreateTimestamp(v int64) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetEnableStatus(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.EnableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetHealthJudgement(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.HealthJudgement = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetHealthStatus(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.HealthStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetHealthTasks(v *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasks) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.HealthTasks = v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetManualAvailableStatus(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.ManualAvailableStatus = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetName(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.Name = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetRemark(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.Remark = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetRequestSource(v *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressRequestSource) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.RequestSource = v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetSeqNonPreemptiveSchedule(v bool) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.SeqNonPreemptiveSchedule = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetSerialNumber(v int32) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.SerialNumber = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetType(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.Type = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetUpdateTime(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.UpdateTime = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetUpdateTimestamp(v int64) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) SetWeightValue(v int32) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress {
	s.WeightValue = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddress) Validate() error {
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

type DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasks struct {
	HealthTask []*DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask `json:"HealthTask,omitempty" xml:"HealthTask,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasks) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasks) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasks) GetHealthTask() []*DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask {
	return s.HealthTask
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasks) SetHealthTask(v []*DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasks {
	s.HealthTask = v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasks) Validate() error {
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

type DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask struct {
	// The target service port for health check probes. When the health check protocol is set to Ping, configuration of the service port is not supported.
	//
	// example:
	//
	// 80
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// ID of the health check template associated with the address.
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

func (s DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask) GetPort() *int32 {
	return s.Port
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask) SetPort(v int32) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask {
	s.Port = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask) SetTemplateId(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask {
	s.TemplateId = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask) SetTemplateName(v string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask {
	s.TemplateName = &v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasksHealthTask) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudGtmAddressPoolResponseBodyAddressesAddressRequestSource struct {
	RequestSource []*string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Repeated"`
}

func (s DescribeCloudGtmAddressPoolResponseBodyAddressesAddressRequestSource) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudGtmAddressPoolResponseBodyAddressesAddressRequestSource) GoString() string {
	return s.String()
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressRequestSource) GetRequestSource() []*string {
	return s.RequestSource
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressRequestSource) SetRequestSource(v []*string) *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressRequestSource {
	s.RequestSource = v
	return s
}

func (s *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressRequestSource) Validate() error {
	return dara.Validate(s)
}
