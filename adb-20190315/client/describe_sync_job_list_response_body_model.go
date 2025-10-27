// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncJobListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v int32) *DescribeSyncJobListResponseBody
	GetDBClusterId() *int32
	SetPageNumber(v int32) *DescribeSyncJobListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSyncJobListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeSyncJobListResponseBody
	GetRequestId() *string
	SetSyncJobs(v []*DescribeSyncJobListResponseBodySyncJobs) *DescribeSyncJobListResponseBody
	GetSyncJobs() []*DescribeSyncJobListResponseBodySyncJobs
	SetTotalCount(v int32) *DescribeSyncJobListResponseBody
	GetTotalCount() *int32
}

type DescribeSyncJobListResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// am-j6c22lubl8d9l3989
	DBClusterId *int32 `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 482B8BAE-6EC0-5C0E-B2AF-FD42A3FC5B67
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The queried synchronization jobs.
	SyncJobs []*DescribeSyncJobListResponseBodySyncJobs `json:"SyncJobs,omitempty" xml:"SyncJobs,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSyncJobListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncJobListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSyncJobListResponseBody) GetDBClusterId() *int32 {
	return s.DBClusterId
}

func (s *DescribeSyncJobListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSyncJobListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSyncJobListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSyncJobListResponseBody) GetSyncJobs() []*DescribeSyncJobListResponseBodySyncJobs {
	return s.SyncJobs
}

func (s *DescribeSyncJobListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSyncJobListResponseBody) SetDBClusterId(v int32) *DescribeSyncJobListResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeSyncJobListResponseBody) SetPageNumber(v int32) *DescribeSyncJobListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSyncJobListResponseBody) SetPageSize(v int32) *DescribeSyncJobListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSyncJobListResponseBody) SetRequestId(v string) *DescribeSyncJobListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSyncJobListResponseBody) SetSyncJobs(v []*DescribeSyncJobListResponseBodySyncJobs) *DescribeSyncJobListResponseBody {
	s.SyncJobs = v
	return s
}

func (s *DescribeSyncJobListResponseBody) SetTotalCount(v int32) *DescribeSyncJobListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSyncJobListResponseBody) Validate() error {
	if s.SyncJobs != nil {
		for _, item := range s.SyncJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSyncJobListResponseBodySyncJobs struct {
	// The ID of the Spark job.
	//
	// example:
	//
	// dts-xxxxx
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The description of the source instance or cluster.
	//
	// example:
	//
	// test
	SourceDBClusterDescription *string `json:"SourceDBClusterDescription,omitempty" xml:"SourceDBClusterDescription,omitempty"`
	// The ID of the source cluster. You can call the [DescribeDBClusters](https://help.aliyun.com/document_detail/170879.html) operation to query backup set IDs.
	//
	// >  If you want to restore the data of an ApsaraDB for ClickHouse cluster, this parameter is required.
	//
	// example:
	//
	// pc-t4n766v2llx852n81
	SourceDBClusterId *string `json:"SourceDBClusterId,omitempty" xml:"SourceDBClusterId,omitempty"`
	// The database type of the source instance or cluster.
	//
	// example:
	//
	// polardb
	SourceDBType *string `json:"SourceDBType,omitempty" xml:"SourceDBType,omitempty"`
	// The storage size of the source instance or cluster.
	//
	// example:
	//
	// 100
	SourceStorageSize *int32 `json:"SourceStorageSize,omitempty" xml:"SourceStorageSize,omitempty"`
	// The number of tables in the source instance or cluster.
	//
	// example:
	//
	// 100
	SourceTableNumber *int32 `json:"SourceTableNumber,omitempty" xml:"SourceTableNumber,omitempty"`
	// The synchronization platform.
	//
	// example:
	//
	// ADB-CDC
	SyncPlatform *string `json:"SyncPlatform,omitempty" xml:"SyncPlatform,omitempty"`
}

func (s DescribeSyncJobListResponseBodySyncJobs) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncJobListResponseBodySyncJobs) GoString() string {
	return s.String()
}

func (s *DescribeSyncJobListResponseBodySyncJobs) GetJobId() *string {
	return s.JobId
}

func (s *DescribeSyncJobListResponseBodySyncJobs) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSyncJobListResponseBodySyncJobs) GetSourceDBClusterDescription() *string {
	return s.SourceDBClusterDescription
}

func (s *DescribeSyncJobListResponseBodySyncJobs) GetSourceDBClusterId() *string {
	return s.SourceDBClusterId
}

func (s *DescribeSyncJobListResponseBodySyncJobs) GetSourceDBType() *string {
	return s.SourceDBType
}

func (s *DescribeSyncJobListResponseBodySyncJobs) GetSourceStorageSize() *int32 {
	return s.SourceStorageSize
}

func (s *DescribeSyncJobListResponseBodySyncJobs) GetSourceTableNumber() *int32 {
	return s.SourceTableNumber
}

func (s *DescribeSyncJobListResponseBodySyncJobs) GetSyncPlatform() *string {
	return s.SyncPlatform
}

func (s *DescribeSyncJobListResponseBodySyncJobs) SetJobId(v string) *DescribeSyncJobListResponseBodySyncJobs {
	s.JobId = &v
	return s
}

func (s *DescribeSyncJobListResponseBodySyncJobs) SetRegionId(v string) *DescribeSyncJobListResponseBodySyncJobs {
	s.RegionId = &v
	return s
}

func (s *DescribeSyncJobListResponseBodySyncJobs) SetSourceDBClusterDescription(v string) *DescribeSyncJobListResponseBodySyncJobs {
	s.SourceDBClusterDescription = &v
	return s
}

func (s *DescribeSyncJobListResponseBodySyncJobs) SetSourceDBClusterId(v string) *DescribeSyncJobListResponseBodySyncJobs {
	s.SourceDBClusterId = &v
	return s
}

func (s *DescribeSyncJobListResponseBodySyncJobs) SetSourceDBType(v string) *DescribeSyncJobListResponseBodySyncJobs {
	s.SourceDBType = &v
	return s
}

func (s *DescribeSyncJobListResponseBodySyncJobs) SetSourceStorageSize(v int32) *DescribeSyncJobListResponseBodySyncJobs {
	s.SourceStorageSize = &v
	return s
}

func (s *DescribeSyncJobListResponseBodySyncJobs) SetSourceTableNumber(v int32) *DescribeSyncJobListResponseBodySyncJobs {
	s.SourceTableNumber = &v
	return s
}

func (s *DescribeSyncJobListResponseBodySyncJobs) SetSyncPlatform(v string) *DescribeSyncJobListResponseBodySyncJobs {
	s.SyncPlatform = &v
	return s
}

func (s *DescribeSyncJobListResponseBodySyncJobs) Validate() error {
	return dara.Validate(s)
}
