// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceAddressPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddrPools(v *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) *DescribeGtmInstanceAddressPoolsResponseBody
	GetAddrPools() *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools
	SetPageNumber(v int32) *DescribeGtmInstanceAddressPoolsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGtmInstanceAddressPoolsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGtmInstanceAddressPoolsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeGtmInstanceAddressPoolsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeGtmInstanceAddressPoolsResponseBody
	GetTotalPages() *int32
}

type DescribeGtmInstanceAddressPoolsResponseBody struct {
	// The returned list of address pools of the GTM instance.
	AddrPools *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools `json:"AddrPools,omitempty" xml:"AddrPools,omitempty" type:"Struct"`
	// The number of the page returned.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned on all pages.
	//
	// example:
	//
	// 2
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) GetAddrPools() *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	return s.AddrPools
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetAddrPools(v *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.AddrPools = v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetPageNumber(v int32) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetPageSize(v int32) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetRequestId(v string) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetTotalItems(v int32) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) SetTotalPages(v int32) *DescribeGtmInstanceAddressPoolsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmInstanceAddressPoolsResponseBodyAddrPools struct {
	AddrPool []*DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool `json:"AddrPool,omitempty" xml:"AddrPool,omitempty" type:"Repeated"`
}

func (s DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) GetAddrPool() []*DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	return s.AddrPool
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) SetAddrPool(v []*DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.AddrPool = v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPools) Validate() error {
	return dara.Validate(s)
}

type DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool struct {
	// The number of addresses in the address pool.
	//
	// example:
	//
	// 2
	AddrCount *int32 `json:"AddrCount,omitempty" xml:"AddrCount,omitempty"`
	// The ID of the address pool.
	//
	// example:
	//
	// 1234abc
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The time when this address pool was created.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The minimum number of available addresses in the address pool.
	//
	// example:
	//
	// 2
	MinAvailableAddrNum *int32 `json:"MinAvailableAddrNum,omitempty" xml:"MinAvailableAddrNum,omitempty"`
	// The health check ID of the address pool.
	//
	// example:
	//
	// 100abc
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// Indicates whether health check was enabled for the address pool. Valid values:
	//
	// 	- **OPEN**: Enabled
	//
	// 	- **CLOSE**: Disabled
	//
	// 	- **UNCONFIGURED**: Not configured
	//
	// example:
	//
	// OPEN
	MonitorStatus *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	// The name of the address pool.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The availability status of the address pool. Valid values:
	//
	// 	- **AVAILABLE**: Available
	//
	// 	- **NOT_AVAILABLE**: Unavailable
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The type of the address pool. Valid values:
	//
	// 	- **IP**: IP address
	//
	// 	- **DOMAIN**: Domain name
	//
	// example:
	//
	// IP
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The last time when the address pool was updated.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// A timestamp that indicates the last time the address pool was updated.
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetAddrCount() *int32 {
	return s.AddrCount
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetMinAvailableAddrNum() *int32 {
	return s.MinAvailableAddrNum
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetName() *string {
	return s.Name
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetStatus() *string {
	return s.Status
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetType() *string {
	return s.Type
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetAddrCount(v int32) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.AddrCount = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetAddrPoolId(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetCreateTime(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.CreateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetCreateTimestamp(v int64) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetMinAvailableAddrNum(v int32) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.MinAvailableAddrNum = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetMonitorConfigId(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetMonitorStatus(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetName(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.Name = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetStatus(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.Status = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetType(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.Type = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetUpdateTime(v string) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.UpdateTime = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetUpdateTimestamp(v int64) *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) Validate() error {
	return dara.Validate(s)
}
