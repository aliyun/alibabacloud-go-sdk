// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCacheReserveInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceInfo(v []*ListCacheReserveInstancesResponseBodyInstanceInfo) *ListCacheReserveInstancesResponseBody
	GetInstanceInfo() []*ListCacheReserveInstancesResponseBodyInstanceInfo
	SetPageNumber(v int32) *ListCacheReserveInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListCacheReserveInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListCacheReserveInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListCacheReserveInstancesResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListCacheReserveInstancesResponseBody
	GetTotalPage() *int32
}

type ListCacheReserveInstancesResponseBody struct {
	// The cache reserve instances.
	InstanceInfo []*ListCacheReserveInstancesResponseBodyInstanceInfo `json:"InstanceInfo,omitempty" xml:"InstanceInfo,omitempty" type:"Repeated"`
	// Page number. Default value: **1**.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// Page size, default **500**, range: **1~500**.
	//
	// example:
	//
	// 500
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 65C66B7B-671A-8297-9187-2R5477247B76
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total count.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// Total pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListCacheReserveInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCacheReserveInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCacheReserveInstancesResponseBody) GetInstanceInfo() []*ListCacheReserveInstancesResponseBodyInstanceInfo {
	return s.InstanceInfo
}

func (s *ListCacheReserveInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListCacheReserveInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCacheReserveInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCacheReserveInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCacheReserveInstancesResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListCacheReserveInstancesResponseBody) SetInstanceInfo(v []*ListCacheReserveInstancesResponseBodyInstanceInfo) *ListCacheReserveInstancesResponseBody {
	s.InstanceInfo = v
	return s
}

func (s *ListCacheReserveInstancesResponseBody) SetPageNumber(v int32) *ListCacheReserveInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBody) SetPageSize(v int32) *ListCacheReserveInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBody) SetRequestId(v string) *ListCacheReserveInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBody) SetTotalCount(v int32) *ListCacheReserveInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBody) SetTotalPage(v int32) *ListCacheReserveInstancesResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBody) Validate() error {
	if s.InstanceInfo != nil {
		for _, item := range s.InstanceInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCacheReserveInstancesResponseBodyInstanceInfo struct {
	// Cache reserve capacity. Unit: GB.
	//
	// example:
	//
	// 512000
	CacheReserveCapacity *int64 `json:"CacheReserveCapacity,omitempty" xml:"CacheReserveCapacity,omitempty"`
	// Cache reserve usage region.
	//
	// example:
	//
	// HK
	CacheReserveRegion *string `json:"CacheReserveRegion,omitempty" xml:"CacheReserveRegion,omitempty"`
	ChargeType         *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// Instance purchase time.
	//
	// example:
	//
	// 2024-04-12T05:41:51Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Duration of the instance purchase, unit: months.
	//
	// example:
	//
	// 3
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Instance expiration time.
	//
	// example:
	//
	// 2024-10-05T16:00:00Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// Instance ID.
	//
	// example:
	//
	// sp-xcdn-96wblslz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The status of the cache reserve instance. Valid values:
	//
	// 	- **online**: The instance is in service.
	//
	// 	- **offline**: The instance has expired within an allowable period. In this state, it is unavailable.
	//
	// 	- **disable**: The instance has been released.
	//
	// 	- **overdue**: The instance has been stopped due to overdue payments.
	//
	// example:
	//
	// online
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListCacheReserveInstancesResponseBodyInstanceInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCacheReserveInstancesResponseBodyInstanceInfo) GoString() string {
	return s.String()
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) GetCacheReserveCapacity() *int64 {
	return s.CacheReserveCapacity
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) GetCacheReserveRegion() *string {
	return s.CacheReserveRegion
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) GetChargeType() *string {
	return s.ChargeType
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) GetDuration() *int32 {
	return s.Duration
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) GetStatus() *string {
	return s.Status
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) SetCacheReserveCapacity(v int64) *ListCacheReserveInstancesResponseBodyInstanceInfo {
	s.CacheReserveCapacity = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) SetCacheReserveRegion(v string) *ListCacheReserveInstancesResponseBodyInstanceInfo {
	s.CacheReserveRegion = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) SetChargeType(v string) *ListCacheReserveInstancesResponseBodyInstanceInfo {
	s.ChargeType = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) SetCreateTime(v string) *ListCacheReserveInstancesResponseBodyInstanceInfo {
	s.CreateTime = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) SetDuration(v int32) *ListCacheReserveInstancesResponseBodyInstanceInfo {
	s.Duration = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) SetExpireTime(v string) *ListCacheReserveInstancesResponseBodyInstanceInfo {
	s.ExpireTime = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) SetInstanceId(v string) *ListCacheReserveInstancesResponseBodyInstanceInfo {
	s.InstanceId = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) SetStatus(v string) *ListCacheReserveInstancesResponseBodyInstanceInfo {
	s.Status = &v
	return s
}

func (s *ListCacheReserveInstancesResponseBodyInstanceInfo) Validate() error {
	return dara.Validate(s)
}
