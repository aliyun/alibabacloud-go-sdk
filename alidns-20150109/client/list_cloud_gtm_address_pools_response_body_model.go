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
	AddressLbStrategy      *string                                                               `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	AddressPoolId          *string                                                               `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	AddressPoolName        *string                                                               `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	AddressPoolType        *string                                                               `json:"AddressPoolType,omitempty" xml:"AddressPoolType,omitempty"`
	Addresses              *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Struct"`
	AvailableStatus        *string                                                               `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	CreateTime             *string                                                               `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp        *int64                                                                `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EnableStatus           *string                                                               `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthJudgement        *string                                                               `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	HealthStatus           *string                                                               `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	Remark                 *string                                                               `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SequenceLbStrategyMode *string                                                               `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
	UpdateTime             *string                                                               `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp        *int64                                                                `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
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
	Address                  *string                                                                                   `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressId                *string                                                                                   `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	AttributeInfo            *string                                                                                   `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	AvailableMode            *string                                                                                   `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	AvailableStatus          *string                                                                                   `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	CreateTime               *string                                                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp          *int64                                                                                    `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EnableStatus             *string                                                                                   `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthJudgement          *string                                                                                   `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	HealthStatus             *string                                                                                   `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	HealthTasks              *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks   `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Struct"`
	ManualAvailableStatus    *string                                                                                   `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
	Name                     *string                                                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark                   *string                                                                                   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RequestSource            *ListCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressRequestSource `json:"RequestSource,omitempty" xml:"RequestSource,omitempty" type:"Struct"`
	SeqNonPreemptiveSchedule *bool                                                                                     `json:"SeqNonPreemptiveSchedule,omitempty" xml:"SeqNonPreemptiveSchedule,omitempty"`
	SerialNumber             *int32                                                                                    `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Type                     *string                                                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime               *string                                                                                   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp          *int64                                                                                    `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	WeightValue              *int32                                                                                    `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
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
	Port         *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
