// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDistributeCacheResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalDistributeCaches(v []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) *DescribeGlobalDistributeCacheResponseBody
	GetGlobalDistributeCaches() []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches
	SetPageNumber(v int32) *DescribeGlobalDistributeCacheResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeGlobalDistributeCacheResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeGlobalDistributeCacheResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeGlobalDistributeCacheResponseBody
	GetTotalRecordCount() *int32
}

type DescribeGlobalDistributeCacheResponseBody struct {
	// The list of distributed instances.
	GlobalDistributeCaches []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches `json:"GlobalDistributeCaches,omitempty" xml:"GlobalDistributeCaches,omitempty" type:"Repeated"`
	// The page number.
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
	// The request ID.
	//
	// example:
	//
	// F3F44BE3-5419-4B61-9BAC-E66E295A****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeGlobalDistributeCacheResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDistributeCacheResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDistributeCacheResponseBody) GetGlobalDistributeCaches() []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches {
	return s.GlobalDistributeCaches
}

func (s *DescribeGlobalDistributeCacheResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeGlobalDistributeCacheResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeGlobalDistributeCacheResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeGlobalDistributeCacheResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeGlobalDistributeCacheResponseBody) SetGlobalDistributeCaches(v []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) *DescribeGlobalDistributeCacheResponseBody {
	s.GlobalDistributeCaches = v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBody) SetPageNumber(v int32) *DescribeGlobalDistributeCacheResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBody) SetPageSize(v int32) *DescribeGlobalDistributeCacheResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBody) SetRequestId(v string) *DescribeGlobalDistributeCacheResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBody) SetTotalRecordCount(v int32) *DescribeGlobalDistributeCacheResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBody) Validate() error {
	if s.GlobalDistributeCaches != nil {
		for _, item := range s.GlobalDistributeCaches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches struct {
	// The instance ID of the distributed instance.
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// The instance status of the distributed instance. Valid values:
	//
	// 	- **Normal**: The instance is running.
	//
	// 	- **Changing**: The instance is being modified.
	//
	// 	- **Creating**: The instance is being created.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of child instances.
	SubInstances []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances `json:"SubInstances,omitempty" xml:"SubInstances,omitempty" type:"Repeated"`
}

func (s DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) GetGlobalInstanceId() *string {
	return s.GlobalInstanceId
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) GetStatus() *string {
	return s.Status
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) GetSubInstances() []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances {
	return s.SubInstances
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) SetGlobalInstanceId(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches {
	s.GlobalInstanceId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) SetStatus(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches {
	s.Status = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) SetSubInstances(v []*DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches {
	s.SubInstances = v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCaches) Validate() error {
	if s.SubInstances != nil {
		for _, item := range s.SubInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances struct {
	// The instance ID of the distributed instance.
	//
	// example:
	//
	// gr-bp14rkqrhac****
	GlobalInstanceId *string `json:"GlobalInstanceId,omitempty" xml:"GlobalInstanceId,omitempty"`
	// The instance type of the child instance. For more information, see:
	//
	// 	- [Standard memory-optimized instance types](https://help.aliyun.com/document_detail/145228.html).
	//
	// 	- [Cluster memory-optimized instance types](https://help.aliyun.com/document_detail/150458.html).
	//
	// 	- [Read/write splitting memory-optimized instance types](https://help.aliyun.com/document_detail/150459.html).
	//
	// example:
	//
	// redis.amber.logic.sharding.2g.2db.0rodb.6proxy.multithread
	InstanceClass *string `json:"InstanceClass,omitempty" xml:"InstanceClass,omitempty"`
	// The instance ID of the child instance.
	//
	// example:
	//
	// r-bp1zxszhcgatnx****
	InstanceID *string `json:"InstanceID,omitempty" xml:"InstanceID,omitempty"`
	// The instance status of the child instance. Valid values:
	//
	// 	- **Normal**: The instance is in a normal state.
	//
	// 	- **Creating**: The instance is being created.
	//
	// 	- **Changing**: The instance is being modified.
	//
	// 	- **Inactive**: The instance is disabled.
	//
	// 	- **Flushing**: The instance is being purged.
	//
	// 	- **Released**: The instance is released.
	//
	// 	- **Transforming**: The instance is being transformed.
	//
	// 	- **Unavailable**: The service is stopped.
	//
	// 	- **Error**: The instance failed to be created.
	//
	// 	- **Migrating**: The instance is being migrated.
	//
	// 	- **BackupRecovering**: The instance is recovering from a backup.
	//
	// 	- **MinorVersionUpgrading**: The minor version of the instance is being upgraded.
	//
	// 	- **NetworkModifying**: The network of the instance is being modified.
	//
	// 	- **SSLModifying**: The SSL certificate of the instance is being modified.
	//
	// 	- **MajorVersionUpgrading**: A major engine version upgrade is in progress. The instance remains accessible.
	//
	// > For more information about instance statuses and their impacts, see [Instance statuses and impacts](https://help.aliyun.com/document_detail/200740.html).
	//
	// example:
	//
	// Normal
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) GetGlobalInstanceId() *string {
	return s.GlobalInstanceId
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) GetInstanceClass() *string {
	return s.InstanceClass
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) GetInstanceID() *string {
	return s.InstanceID
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) SetGlobalInstanceId(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances {
	s.GlobalInstanceId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) SetInstanceClass(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances {
	s.InstanceClass = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) SetInstanceID(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances {
	s.InstanceID = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) SetInstanceStatus(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) SetRegionId(v string) *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances {
	s.RegionId = &v
	return s
}

func (s *DescribeGlobalDistributeCacheResponseBodyGlobalDistributeCachesSubInstances) Validate() error {
	return dara.Validate(s)
}
