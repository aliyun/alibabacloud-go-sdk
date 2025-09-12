// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLindormV2StorageUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetLindormV2StorageUsageResponseBody
	GetAccessDeniedDetail() *string
	SetCapacityByDiskCategory(v []map[string]interface{}) *GetLindormV2StorageUsageResponseBody
	GetCapacityByDiskCategory() []map[string]interface{}
	SetInstanceStorageZoneMap(v map[string]interface{}) *GetLindormV2StorageUsageResponseBody
	GetInstanceStorageZoneMap() map[string]interface{}
	SetRequestId(v string) *GetLindormV2StorageUsageResponseBody
	GetRequestId() *string
	SetUsageByDiskCategory(v []map[string]interface{}) *GetLindormV2StorageUsageResponseBody
	GetUsageByDiskCategory() []map[string]interface{}
}

type GetLindormV2StorageUsageResponseBody struct {
	AccessDeniedDetail     *string                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	CapacityByDiskCategory []map[string]interface{} `json:"CapacityByDiskCategory,omitempty" xml:"CapacityByDiskCategory,omitempty" type:"Repeated"`
	InstanceStorageZoneMap map[string]interface{}   `json:"InstanceStorageZoneMap,omitempty" xml:"InstanceStorageZoneMap,omitempty"`
	RequestId              *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageByDiskCategory    []map[string]interface{} `json:"UsageByDiskCategory,omitempty" xml:"UsageByDiskCategory,omitempty" type:"Repeated"`
}

func (s GetLindormV2StorageUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLindormV2StorageUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetLindormV2StorageUsageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetLindormV2StorageUsageResponseBody) GetCapacityByDiskCategory() []map[string]interface{} {
	return s.CapacityByDiskCategory
}

func (s *GetLindormV2StorageUsageResponseBody) GetInstanceStorageZoneMap() map[string]interface{} {
	return s.InstanceStorageZoneMap
}

func (s *GetLindormV2StorageUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLindormV2StorageUsageResponseBody) GetUsageByDiskCategory() []map[string]interface{} {
	return s.UsageByDiskCategory
}

func (s *GetLindormV2StorageUsageResponseBody) SetAccessDeniedDetail(v string) *GetLindormV2StorageUsageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetLindormV2StorageUsageResponseBody) SetCapacityByDiskCategory(v []map[string]interface{}) *GetLindormV2StorageUsageResponseBody {
	s.CapacityByDiskCategory = v
	return s
}

func (s *GetLindormV2StorageUsageResponseBody) SetInstanceStorageZoneMap(v map[string]interface{}) *GetLindormV2StorageUsageResponseBody {
	s.InstanceStorageZoneMap = v
	return s
}

func (s *GetLindormV2StorageUsageResponseBody) SetRequestId(v string) *GetLindormV2StorageUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLindormV2StorageUsageResponseBody) SetUsageByDiskCategory(v []map[string]interface{}) *GetLindormV2StorageUsageResponseBody {
	s.UsageByDiskCategory = v
	return s
}

func (s *GetLindormV2StorageUsageResponseBody) Validate() error {
	return dara.Validate(s)
}
