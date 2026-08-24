// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKVCacheStoresResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKVCacheStores(v []*ListKVCacheStoresResponseBodyKVCacheStores) *ListKVCacheStoresResponseBody
	GetKVCacheStores() []*ListKVCacheStoresResponseBodyKVCacheStores
	SetMaxResults(v int32) *ListKVCacheStoresResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListKVCacheStoresResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *ListKVCacheStoresResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKVCacheStoresResponseBody
	GetPageSize() *int32
	SetPageTotal(v int32) *ListKVCacheStoresResponseBody
	GetPageTotal() *int32
	SetRequestId(v string) *ListKVCacheStoresResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListKVCacheStoresResponseBody
	GetTotalCount() *int32
}

type ListKVCacheStoresResponseBody struct {
	// The list of KVCacheStore instances. Each element contains the following fields: KvcsId, Name, Status, ExtraStatus, RegionId, ZoneId, HpnZone, Type, Capacity, PaymentType, MountPointId, CreateTime, and Description.
	KVCacheStores []*ListKVCacheStoresResponseBodyKVCacheStores `json:"KVCacheStores,omitempty" xml:"KVCacheStores,omitempty" type:"Repeated"`
	// The maximum number of entries returned per pagination request.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token for the next page. This value is empty when no more data is available. This parameter is valid only for cursor-based pagination.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The current page number. This parameter is valid only for page number-based pagination.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. This parameter is valid only for page number-based pagination.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of pages. This value is returned only for page number-based pagination.
	//
	// example:
	//
	// 1
	PageTotal *int32 `json:"PageTotal,omitempty" xml:"PageTotal,omitempty"`
	// The request ID. A request ID is returned regardless of whether the API call succeeds.
	//
	// example:
	//
	// 56AC37CD-388E-5D21-951B-C50D16D8E812
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of instances. This value is returned only for page number-based pagination. For cursor-based pagination, the value is -1.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKVCacheStoresResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoresResponseBody) GetKVCacheStores() []*ListKVCacheStoresResponseBodyKVCacheStores {
	return s.KVCacheStores
}

func (s *ListKVCacheStoresResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListKVCacheStoresResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListKVCacheStoresResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKVCacheStoresResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKVCacheStoresResponseBody) GetPageTotal() *int32 {
	return s.PageTotal
}

func (s *ListKVCacheStoresResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKVCacheStoresResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKVCacheStoresResponseBody) SetKVCacheStores(v []*ListKVCacheStoresResponseBodyKVCacheStores) *ListKVCacheStoresResponseBody {
	s.KVCacheStores = v
	return s
}

func (s *ListKVCacheStoresResponseBody) SetMaxResults(v int32) *ListKVCacheStoresResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListKVCacheStoresResponseBody) SetNextToken(v string) *ListKVCacheStoresResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListKVCacheStoresResponseBody) SetPageNumber(v int32) *ListKVCacheStoresResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKVCacheStoresResponseBody) SetPageSize(v int32) *ListKVCacheStoresResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKVCacheStoresResponseBody) SetPageTotal(v int32) *ListKVCacheStoresResponseBody {
	s.PageTotal = &v
	return s
}

func (s *ListKVCacheStoresResponseBody) SetRequestId(v string) *ListKVCacheStoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKVCacheStoresResponseBody) SetTotalCount(v int32) *ListKVCacheStoresResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKVCacheStoresResponseBody) Validate() error {
	if s.KVCacheStores != nil {
		for _, item := range s.KVCacheStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKVCacheStoresResponseBodyKVCacheStores struct {
	// The storage capacity. Unit: GiB.
	//
	// example:
	//
	// 100
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The creation time in ISO 8601 format.
	//
	// example:
	//
	// 2026-06-18T10:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The instance description.
	//
	// example:
	//
	// production kvcachestore
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extra status information. Valid values: CapacityExpanding, CapacityExpandSuccess, and CapacityExpandFail.
	//
	// example:
	//
	// CapacityExpanding
	ExtraStatus *string `json:"ExtraStatus,omitempty" xml:"ExtraStatus,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// default
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// kvcs-xxxxx
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// The file system-level mount point ID. Instances under the same file system share this mount point. For more information, call ListKVCacheInstanceAttachInfo.
	//
	// example:
	//
	// mp-xxxxx
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	// The instance name.
	//
	// example:
	//
	// production-instance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The payment type. Valid values: PREPAY and POSTPAY.
	//
	// example:
	//
	// POSTPAY
	PaymentType *string `json:"PaymentType,omitempty" xml:"PaymentType,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek3dnrvdxj2dvq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance status. Valid values: Creating, Available, InUse, Stopping, Stopped, and Deleting.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of resource tags.
	Tags []*ListKVCacheStoresResponseBodyKVCacheStoresTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The instance type. Valid values: kvcs (KVCacheStore, CPFS).
	//
	// example:
	//
	// preview
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListKVCacheStoresResponseBodyKVCacheStores) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoresResponseBodyKVCacheStores) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetCapacity() *int64 {
	return s.Capacity
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetDescription() *string {
	return s.Description
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetExtraStatus() *string {
	return s.ExtraStatus
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetKvcsId() *string {
	return s.KvcsId
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetMountPointId() *string {
	return s.MountPointId
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetName() *string {
	return s.Name
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetPaymentType() *string {
	return s.PaymentType
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetStatus() *string {
	return s.Status
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetTags() []*ListKVCacheStoresResponseBodyKVCacheStoresTags {
	return s.Tags
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetType() *string {
	return s.Type
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetCapacity(v int64) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.Capacity = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetCreateTime(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.CreateTime = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetDescription(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.Description = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetExtraStatus(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.ExtraStatus = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetHpnZone(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.HpnZone = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetKvcsId(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.KvcsId = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetMountPointId(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.MountPointId = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetName(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.Name = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetPaymentType(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.PaymentType = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetRegionId(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.RegionId = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetResourceGroupId(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.ResourceGroupId = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetStatus(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.Status = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetTags(v []*ListKVCacheStoresResponseBodyKVCacheStoresTags) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.Tags = v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetType(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.Type = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) SetZoneId(v string) *ListKVCacheStoresResponseBodyKVCacheStores {
	s.ZoneId = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStores) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKVCacheStoresResponseBodyKVCacheStoresTags struct {
	// The tag key of the resource.
	//
	// example:
	//
	// chapter
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the resource.
	//
	// example:
	//
	// test-value-1766542011
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListKVCacheStoresResponseBodyKVCacheStoresTags) String() string {
	return dara.Prettify(s)
}

func (s ListKVCacheStoresResponseBodyKVCacheStoresTags) GoString() string {
	return s.String()
}

func (s *ListKVCacheStoresResponseBodyKVCacheStoresTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListKVCacheStoresResponseBodyKVCacheStoresTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListKVCacheStoresResponseBodyKVCacheStoresTags) SetTagKey(v string) *ListKVCacheStoresResponseBodyKVCacheStoresTags {
	s.TagKey = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStoresTags) SetTagValue(v string) *ListKVCacheStoresResponseBodyKVCacheStoresTags {
	s.TagValue = &v
	return s
}

func (s *ListKVCacheStoresResponseBodyKVCacheStoresTags) Validate() error {
	return dara.Validate(s)
}
