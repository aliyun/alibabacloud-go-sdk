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
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail     *string                  `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	CapacityByDiskCategory []map[string]interface{} `json:"CapacityByDiskCategory,omitempty" xml:"CapacityByDiskCategory,omitempty" type:"Repeated"`
	// example:
	//
	// {
	//
	//     "cn-hangzhou-i": {
	//
	//       "diskTypeCapacity": [
	//
	//         {
	//
	//           "mode": "CLOUD_STORAGE",
	//
	//           "perfLevel": "PL1",
	//
	//           "usedCapacity": 0,
	//
	//           "category": "PERF_CLOUD_ESSD_PL1",
	//
	//           "capacity": 4000
	//
	//         }
	//
	//       ],
	//
	//       "diskTypeUsage": [
	//
	//         {
	//
	//           "usedLindormColumn3": 688935,
	//
	//           "usedLindormTable": 1086288931872,
	//
	//           "usedLindormTsdb": 0,
	//
	//           "usedOther": 0,
	//
	//           "usedLindormMessage3": 0,
	//
	//           "diskType": "PerformanceCloudStorage",
	//
	//           "used": 1719816329046,
	//
	//           "usedLindormSearch3": 36339905446,
	//
	//           "usedLindormSpark": 2131936938,
	//
	//           "capacity": 4294967296000,
	//
	//           "usedLindormSearch": 0,
	//
	//           "usedLindormVector3": 595054865855
	//
	//         }
	//
	//       ]
	//
	//     }
	//
	//   }
	InstanceStorageZoneMap map[string]interface{} `json:"InstanceStorageZoneMap,omitempty" xml:"InstanceStorageZoneMap,omitempty"`
	// example:
	//
	// BDDB1954-002B-4249-B2DF-2CDDA0259668
	RequestId           *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UsageByDiskCategory []map[string]interface{} `json:"UsageByDiskCategory,omitempty" xml:"UsageByDiskCategory,omitempty" type:"Repeated"`
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
