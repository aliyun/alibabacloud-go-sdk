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
	if s.AddressPools != nil {
		if err := s.AddressPools.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPool struct {
	AddressLbStrategy      *string                                                                 `json:"AddressLbStrategy,omitempty" xml:"AddressLbStrategy,omitempty"`
	AddressPoolId          *string                                                                 `json:"AddressPoolId,omitempty" xml:"AddressPoolId,omitempty"`
	AddressPoolName        *string                                                                 `json:"AddressPoolName,omitempty" xml:"AddressPoolName,omitempty"`
	AddressPoolType        *string                                                                 `json:"AddressPoolType,omitempty" xml:"AddressPoolType,omitempty"`
	Addresses              *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddresses `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Struct"`
	AvailableStatus        *string                                                                 `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	CreateTime             *string                                                                 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp        *int64                                                                  `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EnableStatus           *string                                                                 `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthJudgement        *string                                                                 `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	HealthStatus           *string                                                                 `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	Remark                 *string                                                                 `json:"Remark,omitempty" xml:"Remark,omitempty"`
	SequenceLbStrategyMode *string                                                                 `json:"SequenceLbStrategyMode,omitempty" xml:"SequenceLbStrategyMode,omitempty"`
	UpdateTime             *string                                                                 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp        *int64                                                                  `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
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
	if s.Addresses != nil {
		if err := s.Addresses.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddress struct {
	Address               *string                                                                                   `json:"Address,omitempty" xml:"Address,omitempty"`
	AddressId             *string                                                                                   `json:"AddressId,omitempty" xml:"AddressId,omitempty"`
	AttributeInfo         *string                                                                                   `json:"AttributeInfo,omitempty" xml:"AttributeInfo,omitempty"`
	AvailableMode         *string                                                                                   `json:"AvailableMode,omitempty" xml:"AvailableMode,omitempty"`
	AvailableStatus       *string                                                                                   `json:"AvailableStatus,omitempty" xml:"AvailableStatus,omitempty"`
	CreateTime            *string                                                                                   `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CreateTimestamp       *int64                                                                                    `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	EnableStatus          *string                                                                                   `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	HealthJudgement       *string                                                                                   `json:"HealthJudgement,omitempty" xml:"HealthJudgement,omitempty"`
	HealthStatus          *string                                                                                   `json:"HealthStatus,omitempty" xml:"HealthStatus,omitempty"`
	HealthTasks           *SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasks `json:"HealthTasks,omitempty" xml:"HealthTasks,omitempty" type:"Struct"`
	ManualAvailableStatus *string                                                                                   `json:"ManualAvailableStatus,omitempty" xml:"ManualAvailableStatus,omitempty"`
	Name                  *string                                                                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	Remark                *string                                                                                   `json:"Remark,omitempty" xml:"Remark,omitempty"`
	RequestSource         *string                                                                                   `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
	SerialNumber          *int32                                                                                    `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Type                  *string                                                                                   `json:"Type,omitempty" xml:"Type,omitempty"`
	UpdateTime            *string                                                                                   `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UpdateTimestamp       *int64                                                                                    `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
	WeightValue           *int32                                                                                    `json:"WeightValue,omitempty" xml:"WeightValue,omitempty"`
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
	if s.HealthTasks != nil {
		if err := s.HealthTasks.Validate(); err != nil {
			return err
		}
	}
	return nil
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

type SearchCloudGtmAddressPoolsResponseBodyAddressPoolsAddressPoolAddressesAddressHealthTasksHealthTask struct {
	Port         *int32  `json:"Port,omitempty" xml:"Port,omitempty"`
	TemplateId   *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
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
