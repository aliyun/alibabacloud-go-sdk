// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKVCacheStoreResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKvCacheStore(v *GetKVCacheStoreResponseBodyKvCacheStore) *GetKVCacheStoreResponseBody
	GetKvCacheStore() *GetKVCacheStoreResponseBodyKvCacheStore
	SetRequestId(v string) *GetKVCacheStoreResponseBody
	GetRequestId() *string
}

type GetKVCacheStoreResponseBody struct {
	// The details of the KvCacheStore instance.
	KvCacheStore *GetKVCacheStoreResponseBodyKvCacheStore `json:"KvCacheStore,omitempty" xml:"KvCacheStore,omitempty" type:"Struct"`
	// The ID of the request. The request ID is returned regardless of whether the call is successful.
	//
	// example:
	//
	// 019FB5E9-F9E8-52F5-9C56-2CDF479CBEB2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetKVCacheStoreResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKVCacheStoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetKVCacheStoreResponseBody) GetKvCacheStore() *GetKVCacheStoreResponseBodyKvCacheStore {
	return s.KvCacheStore
}

func (s *GetKVCacheStoreResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKVCacheStoreResponseBody) SetKvCacheStore(v *GetKVCacheStoreResponseBodyKvCacheStore) *GetKVCacheStoreResponseBody {
	s.KvCacheStore = v
	return s
}

func (s *GetKVCacheStoreResponseBody) SetRequestId(v string) *GetKVCacheStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKVCacheStoreResponseBody) Validate() error {
	if s.KvCacheStore != nil {
		if err := s.KvCacheStore.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKVCacheStoreResponseBodyKvCacheStore struct {
	// The storage capacity, in GiB.
	//
	// example:
	//
	// 100
	Capacity *int64 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The time when the instance was created. The time follows the ISO 8601 standard.
	//
	// example:
	//
	// 2026-06-18T10:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the instance.
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
	// The cluster number.
	//
	// example:
	//
	// default
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// kvcs-xxxxx
	KvcsId *string `json:"KvcsId,omitempty" xml:"KvcsId,omitempty"`
	// The ID of the file system-level mount point. The mount point is shared by instances in the same file system. For more information, call the ListKVCacheInstanceAttachInfo operation.
	//
	// example:
	//
	// mp-xxxxx
	MountPointId *string `json:"MountPointId,omitempty" xml:"MountPointId,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// production-instance
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The billing method. Valid values: PREPAY and POSTPAY.
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
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzrwkxbdvkctq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The instance status. Valid values: Creating, Available, InUse, Stopping, Stopped, and Deleting.
	//
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the resource.
	Tags []*GetKVCacheStoreResponseBodyKvCacheStoreTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The instance type, such as kvcs (KVCacheStore, CPFS).
	//
	// example:
	//
	// kvcs
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The zone ID.
	//
	// example:
	//
	// cn-hangzhou-a
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetKVCacheStoreResponseBodyKvCacheStore) String() string {
	return dara.Prettify(s)
}

func (s GetKVCacheStoreResponseBodyKvCacheStore) GoString() string {
	return s.String()
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetCapacity() *int64 {
	return s.Capacity
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetDescription() *string {
	return s.Description
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetExtraStatus() *string {
	return s.ExtraStatus
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetHpnZone() *string {
	return s.HpnZone
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetKvcsId() *string {
	return s.KvcsId
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetMountPointId() *string {
	return s.MountPointId
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetName() *string {
	return s.Name
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetPaymentType() *string {
	return s.PaymentType
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetRegionId() *string {
	return s.RegionId
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetStatus() *string {
	return s.Status
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetTags() []*GetKVCacheStoreResponseBodyKvCacheStoreTags {
	return s.Tags
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetType() *string {
	return s.Type
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetCapacity(v int64) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.Capacity = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetCreateTime(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.CreateTime = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetDescription(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.Description = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetExtraStatus(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.ExtraStatus = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetHpnZone(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.HpnZone = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetKvcsId(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.KvcsId = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetMountPointId(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.MountPointId = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetName(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.Name = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetPaymentType(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.PaymentType = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetRegionId(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.RegionId = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetResourceGroupId(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.ResourceGroupId = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetStatus(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.Status = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetTags(v []*GetKVCacheStoreResponseBodyKvCacheStoreTags) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.Tags = v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetType(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.Type = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) SetZoneId(v string) *GetKVCacheStoreResponseBodyKvCacheStore {
	s.ZoneId = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStore) Validate() error {
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

type GetKVCacheStoreResponseBodyKvCacheStoreTags struct {
	// The key of the tag.
	//
	// example:
	//
	// ac-cus-tag-6
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// advanced
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetKVCacheStoreResponseBodyKvCacheStoreTags) String() string {
	return dara.Prettify(s)
}

func (s GetKVCacheStoreResponseBodyKvCacheStoreTags) GoString() string {
	return s.String()
}

func (s *GetKVCacheStoreResponseBodyKvCacheStoreTags) GetTagKey() *string {
	return s.TagKey
}

func (s *GetKVCacheStoreResponseBodyKvCacheStoreTags) GetTagValue() *string {
	return s.TagValue
}

func (s *GetKVCacheStoreResponseBodyKvCacheStoreTags) SetTagKey(v string) *GetKVCacheStoreResponseBodyKvCacheStoreTags {
	s.TagKey = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStoreTags) SetTagValue(v string) *GetKVCacheStoreResponseBodyKvCacheStoreTags {
	s.TagValue = &v
	return s
}

func (s *GetKVCacheStoreResponseBodyKvCacheStoreTags) Validate() error {
	return dara.Validate(s)
}
