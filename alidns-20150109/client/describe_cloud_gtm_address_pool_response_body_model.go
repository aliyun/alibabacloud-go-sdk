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
	// The load balancing policy for the addresses in the address pool. Valid values:
	//
	// - round_robin: For a DNS request from any source, all addresses are returned. The addresses are rotated in each response.
	//
	// - sequence: For a DNS request from any source, the address with the highest priority is returned. Priority is determined by the \\`SerialNumber\\`, where a smaller value indicates a higher priority. If the highest-priority address is unavailable, the address with the next highest priority is returned.
	//
	// - weight: A weight can be set for each address. DNS requests are resolved based on the specified weight ratio.
	//
	// - source_nearest: Global Traffic Manager (GTM) returns an address based on the source of the DNS request. This implements proximity-based access for users.
	//
	// example:
	//
	// round_robin
	AddressLbStrategy *string `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	// The unique ID of the address pool.
	//
	// example:
	//
	// pool-89564674533755****
	AddressPoolId *string `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// AddressPool-1
	AddressPoolName *string `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	// The type of the address pool. Valid values:
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
	AddressPoolType *string                                           `json:"AddressPoolType,omitempty" xml:"AddressPoolType,omitempty"`
	Addresses       *DescribeCloudGtmAddressPoolResponseBodyAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Struct"`
	// The availability status of the address pool. Valid values:
	//
	// - available: Available.
	//
	// - unavailable: Unavailable.
	//
	// example:
	//
	// available
	AvailableStatus *string `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	// The time when the address pool was created.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp that indicates when the address pool was created.
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The status of the address pool. Valid values:
	//
	// - enable: Enabled
	//
	// - disable: Disabled
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The condition used to determine the health status of the address pool. Valid values:
	//
	// - any_ok: At least one address is available.
	//
	// - p30_ok: At least 30% of the addresses are available.
	//
	// - p50_ok: At least 50% of the addresses are available.
	//
	// - p70_ok: At least 70% of the addresses are available.
	//
	// - all_ok: All addresses are available.
	//
	// example:
	//
	// any_ok
	HealthJudgement *string `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	// The health status of the address pool. Valid values:
	//
	// - ok: The address pool is healthy. All addresses in the address pool are available.
	//
	// - ok_alert: The address pool is in an alert state. Some addresses are unavailable, but the address pool is still considered healthy. In this state, DNS resolution is performed for available addresses, but not for unavailable addresses.
	//
	// - exceptional: The address pool is unhealthy. Some or all addresses are unavailable, and the address pool is considered unhealthy.
	//
	// example:
	//
	// ok
	HealthStatus *string `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	// The notes on the address.
	//
	// example:
	//
	// test
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The unique ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service recovery mode for a primary address that becomes available again when the load balancing policy is set to \\`sequence\\`. Valid values:
	//
	// - preemptive: The system preferentially uses the address with a smaller \\`SerialNumber\\`.
	//
	// - non_preemptive: The system continues to use the current address.
	//
	// example:
	//
	// preemptive
	SequenceLbStrategyMode *string `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
	// The time when the address pool was last modified.
	//
	// example:
	//
	// 2024-03-15T01:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The timestamp that indicates when the address pool was last modified.
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
	Address                  *string                                                               `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressId                *string                                                               `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	AttributeInfo            *string                                                               `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	AvailableMode            *string                                                               `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	AvailableStatus          *string                                                               `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	CreateTime               *string                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp          *int64                                                                `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EnableStatus             *string                                                               `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthJudgement          *string                                                               `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	HealthStatus             *string                                                               `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	HealthTasks              *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressHealthTasks   `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Struct"`
	ManualAvailableStatus    *string                                                               `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
	Name                     *string                                                               `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark                   *string                                                               `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RequestSource            *DescribeCloudGtmAddressPoolResponseBodyAddressesAddressRequestSource `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Struct"`
	SeqNonPreemptiveSchedule *bool                                                                 `json:"SeqNonPreemptiveSchedule,omitempty" xml:"SeqNonPreemptiveSchedule,omitempty"`
	SerialNumber             *int32                                                                `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Type                     *string                                                               `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime               *string                                                               `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp          *int64                                                                `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	WeightValue              *int32                                                                `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
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
	Port         *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
