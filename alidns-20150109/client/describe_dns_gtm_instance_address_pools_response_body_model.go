// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDnsGtmInstanceAddressPoolsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddrPools(v *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) *DescribeDnsGtmInstanceAddressPoolsResponseBody
	GetAddrPools() *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools
	SetPageNumber(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBody
	GetRequestId() *string
	SetTotalItems(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody
	GetTotalItems() *int32
	SetTotalPages(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody
	GetTotalPages() *int32
}

type DescribeDnsGtmInstanceAddressPoolsResponseBody struct {
	// The returned address pools.
	AddrPools *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools `json:"AddrPools,omitempty" xml:"AddrPools,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 1
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
	// 1
	TotalItems *int32 `json:"TotalItems,omitempty" xml:"TotalItems,omitempty"`
	// The total number of pages returned.
	//
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) GetAddrPools() *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	return s.AddrPools
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) GetTotalItems() *int32 {
	return s.TotalItems
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetAddrPools(v *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.AddrPools = v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetPageNumber(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetPageSize(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetRequestId(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetTotalItems(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.TotalItems = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) SetTotalPages(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBody {
	s.TotalPages = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBody) Validate() error {
	if s.AddrPools != nil {
		if err := s.AddrPools.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools struct {
	AddrPool []*DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool `json:"AddrPool,omitempty" xml:"AddrPool,omitempty" type:"Repeated"`
}

func (s DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) GetAddrPool() []*DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	return s.AddrPool
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) SetAddrPool(v []*DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools {
	s.AddrPool = v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPools) Validate() error {
	if s.AddrPool != nil {
		for _, item := range s.AddrPool {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool struct {
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
	// pool-1
	AddrPoolId *string `json:"AddrPoolId,omitempty" xml:"AddrPoolId,omitempty"`
	// The time when the address pool was created.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The timestamp that indicates when the address pool was created.
	//
	// example:
	//
	// 1527690629357
	CreateTimestamp *int64 `json:"CreateTimestamp,omitempty" xml:"CreateTimestamp,omitempty"`
	// The load balancing policy of the address pool. Valid values:
	//
	// 	- ALL_RR: returns all addresses.
	//
	// 	- RATIO: returns addresses by weight.
	//
	// example:
	//
	// all_rr
	LbaStrategy *string `json:"LbaStrategy,omitempty" xml:"LbaStrategy,omitempty"`
	// The ID of the health check task.
	//
	// example:
	//
	// abc123
	MonitorConfigId *string `json:"MonitorConfigId,omitempty" xml:"MonitorConfigId,omitempty"`
	// Indicates whether health checks are configured. Valid values:
	//
	// 	- OPEN: enabled
	//
	// 	- CLOSE: disabled
	//
	// 	- UNCONFIGURED: not configured
	//
	// example:
	//
	// open
	MonitorStatus *string `json:"MonitorStatus,omitempty" xml:"MonitorStatus,omitempty"`
	// The name of the address pool.
	//
	// example:
	//
	// testpool
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the address pool. Valid values:
	//
	// 	- IPV4: IPv4 address
	//
	// 	- IPV6: IPv6 address
	//
	// 	- DOMAIN: domain name
	//
	// example:
	//
	// ipv4
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The time when the address pool was updated.
	//
	// example:
	//
	// 2017-12-28T13:08Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The timestamp that indicates when the address pool was updated.
	//
	// example:
	//
	// 1527690629357
	UpdateTimestamp *int64 `json:"UpdateTimestamp,omitempty" xml:"UpdateTimestamp,omitempty"`
}

func (s DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) String() string {
	return dara.Prettify(s)
}

func (s DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GoString() string {
	return s.String()
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetAddrCount() *int32 {
	return s.AddrCount
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetAddrPoolId() *string {
	return s.AddrPoolId
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetLbaStrategy() *string {
	return s.LbaStrategy
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetMonitorConfigId() *string {
	return s.MonitorConfigId
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetMonitorStatus() *string {
	return s.MonitorStatus
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetName() *string {
	return s.Name
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetType() *string {
	return s.Type
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetAddrCount(v int32) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.AddrCount = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetAddrPoolId(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.AddrPoolId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetCreateTime(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.CreateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetCreateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.CreateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetLbaStrategy(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.LbaStrategy = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetMonitorConfigId(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.MonitorConfigId = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetMonitorStatus(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.MonitorStatus = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetName(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.Name = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetType(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.Type = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetUpdateTime(v string) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) SetUpdateTimestamp(v int64) *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool {
	s.UpdateTimestamp = &v
	return s
}

func (s *DescribeDnsGtmInstanceAddressPoolsResponseBodyAddrPoolsAddrPool) Validate() error {
	return dara.Validate(s)
}
