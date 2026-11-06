// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKVCacheInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeKVCacheInstancesResponseBodyItems) *DescribeKVCacheInstancesResponseBody
	GetItems() *DescribeKVCacheInstancesResponseBodyItems
	SetPageNumber(v int32) *DescribeKVCacheInstancesResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeKVCacheInstancesResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeKVCacheInstancesResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeKVCacheInstancesResponseBody
	GetTotalRecordCount() *int32
}

type DescribeKVCacheInstancesResponseBody struct {
	Items *DescribeKVCacheInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of instances returned on the current page.
	//
	// example:
	//
	// 1
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2921D843-433A-5FB3-A03B-4EC093B219F8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances that match the filter conditions.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeKVCacheInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKVCacheInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKVCacheInstancesResponseBody) GetItems() *DescribeKVCacheInstancesResponseBodyItems {
	return s.Items
}

func (s *DescribeKVCacheInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeKVCacheInstancesResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeKVCacheInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKVCacheInstancesResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeKVCacheInstancesResponseBody) SetItems(v *DescribeKVCacheInstancesResponseBodyItems) *DescribeKVCacheInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeKVCacheInstancesResponseBody) SetPageNumber(v int32) *DescribeKVCacheInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBody) SetPageRecordCount(v int32) *DescribeKVCacheInstancesResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBody) SetRequestId(v string) *DescribeKVCacheInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBody) SetTotalRecordCount(v int32) *DescribeKVCacheInstancesResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeKVCacheInstancesResponseBodyItems struct {
	KVCacheItem []*DescribeKVCacheInstancesResponseBodyItemsKVCacheItem `json:"KVCacheItem,omitempty" xml:"KVCacheItem,omitempty" type:"Repeated"`
}

func (s DescribeKVCacheInstancesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeKVCacheInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeKVCacheInstancesResponseBodyItems) GetKVCacheItem() []*DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	return s.KVCacheItem
}

func (s *DescribeKVCacheInstancesResponseBodyItems) SetKVCacheItem(v []*DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) *DescribeKVCacheInstancesResponseBodyItems {
	s.KVCacheItem = v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItems) Validate() error {
	if s.KVCacheItem != nil {
		for _, item := range s.KVCacheItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKVCacheInstancesResponseBodyItemsKVCacheItem struct {
	AIDBClusterId         *string `json:"AIDBClusterId,omitempty" xml:"AIDBClusterId,omitempty"`
	CreationTime          *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DBClusterId           *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	ExpireTime            *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	Expired               *string `json:"Expired,omitempty" xml:"Expired,omitempty"`
	KVCacheInstanceId     *string `json:"KVCacheInstanceId,omitempty" xml:"KVCacheInstanceId,omitempty"`
	KVCacheInstanceStatus *string `json:"KVCacheInstanceStatus,omitempty" xml:"KVCacheInstanceStatus,omitempty"`
	KVCacheSpace          *int32  `json:"KVCacheSpace,omitempty" xml:"KVCacheSpace,omitempty"`
	PayType               *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	RegionId              *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	VPCId                 *string `json:"VPCId,omitempty" xml:"VPCId,omitempty"`
	VSwitchId             *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	ZoneId                *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GoString() string {
	return s.String()
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetAIDBClusterId() *string {
	return s.AIDBClusterId
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetExpired() *string {
	return s.Expired
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetKVCacheInstanceId() *string {
	return s.KVCacheInstanceId
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetKVCacheInstanceStatus() *string {
	return s.KVCacheInstanceStatus
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetKVCacheSpace() *int32 {
	return s.KVCacheSpace
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetPayType() *string {
	return s.PayType
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetVPCId() *string {
	return s.VPCId
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetAIDBClusterId(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.AIDBClusterId = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetCreationTime(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.CreationTime = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetDBClusterId(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.DBClusterId = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetExpireTime(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.ExpireTime = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetExpired(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.Expired = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetKVCacheInstanceId(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.KVCacheInstanceId = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetKVCacheInstanceStatus(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.KVCacheInstanceStatus = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetKVCacheSpace(v int32) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.KVCacheSpace = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetPayType(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.PayType = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetRegionId(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.RegionId = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetVPCId(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.VPCId = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetVSwitchId(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.VSwitchId = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) SetZoneId(v string) *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem {
	s.ZoneId = &v
	return s
}

func (s *DescribeKVCacheInstancesResponseBodyItemsKVCacheItem) Validate() error {
	return dara.Validate(s)
}
