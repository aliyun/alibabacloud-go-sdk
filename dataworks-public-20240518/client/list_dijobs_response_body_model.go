// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDIJobsResponseBodyPagingInfo) *ListDIJobsResponseBody
	GetPagingInfo() *ListDIJobsResponseBodyPagingInfo
	SetRequestId(v string) *ListDIJobsResponseBody
	GetRequestId() *string
}

type ListDIJobsResponseBody struct {
	// The pagination information.
	PagingInfo *ListDIJobsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 7263E4AC-9D2E-5B29-B8AF-7C5012E92A41
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBody) GetPagingInfo() *ListDIJobsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDIJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDIJobsResponseBody) SetPagingInfo(v *ListDIJobsResponseBodyPagingInfo) *ListDIJobsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIJobsResponseBody) SetRequestId(v string) *ListDIJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDIJobsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDIJobsResponseBodyPagingInfo struct {
	// The synchronization tasks returned.
	DIJobs []*ListDIJobsResponseBodyPagingInfoDIJobs `json:"DIJobs,omitempty" xml:"DIJobs,omitempty" type:"Repeated"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 12
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListDIJobsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBodyPagingInfo) GetDIJobs() []*ListDIJobsResponseBodyPagingInfoDIJobs {
	return s.DIJobs
}

func (s *ListDIJobsResponseBodyPagingInfo) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDIJobsResponseBodyPagingInfo) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDIJobsResponseBodyPagingInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListDIJobsResponseBodyPagingInfo) SetDIJobs(v []*ListDIJobsResponseBodyPagingInfoDIJobs) *ListDIJobsResponseBodyPagingInfo {
	s.DIJobs = v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfo) SetPageNumber(v int64) *ListDIJobsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfo) SetPageSize(v int64) *ListDIJobsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfo) SetTotalCount(v int64) *ListDIJobsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfo) Validate() error {
	if s.DIJobs != nil {
		for _, item := range s.DIJobs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDIJobsResponseBodyPagingInfoDIJobs struct {
	// Deprecated
	//
	// This parameter is deprecated. Use the Id parameter instead.
	//
	// example:
	//
	// 32599
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The destination type. Valid values: Hologres, OSS-HDFS, OSS, MaxCompute, Loghub, STARROCKS, DataHub, ANALYTICDB_FOR_MYSQL, Kafka, and Hive.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 32599
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the synchronization task.
	//
	// example:
	//
	// mysql_to_holo_sync_35197
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The status of the synchronization task. Valid values:
	//
	// 	- Finished
	//
	// 	- Initialized
	//
	// 	- Stopped
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- Stopping
	//
	// example:
	//
	// Running
	JobStatus *string `json:"JobStatus,omitempty" xml:"JobStatus,omitempty"`
	// The synchronization type. Valid values:
	//
	// 	- FullAndRealtimeIncremental: one-time full synchronization and real-time incremental synchronization
	//
	// 	- RealtimeIncremental: real-time incremental synchronization
	//
	// 	- Full: full synchronization
	//
	// 	- OfflineIncremental: batch incremental synchronization
	//
	// 	- FullAndOfflineIncremental: one-time full synchronization and batch incremental synchronization
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The ID of the DataWorks workspace to which the synchronization task belongs.
	//
	// example:
	//
	// 26442
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The source type. Valid values: PolarDB, MySQL, Kafka, Loghub, Hologres, Oracle, OceanBase, MongoDB, RedShift, Hive, SqlServer, Doris, and ClickHouse. If you do not configure this parameter, the API operation returns synchronization tasks that use all types of sources.
	//
	// example:
	//
	// Mysql
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
}

func (s ListDIJobsResponseBodyPagingInfoDIJobs) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobsResponseBodyPagingInfoDIJobs) GoString() string {
	return s.String()
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) GetDestinationDataSourceType() *string {
	return s.DestinationDataSourceType
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) GetId() *int64 {
	return s.Id
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) GetJobName() *string {
	return s.JobName
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) GetJobStatus() *string {
	return s.JobStatus
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) GetMigrationType() *string {
	return s.MigrationType
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) GetSourceDataSourceType() *string {
	return s.SourceDataSourceType
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetDIJobId(v int64) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.DIJobId = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetDestinationDataSourceType(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.DestinationDataSourceType = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetId(v int64) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.Id = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetJobName(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.JobName = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetJobStatus(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.JobStatus = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetMigrationType(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.MigrationType = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetProjectId(v int64) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.ProjectId = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) SetSourceDataSourceType(v string) *ListDIJobsResponseBodyPagingInfoDIJobs {
	s.SourceDataSourceType = &v
	return s
}

func (s *ListDIJobsResponseBodyPagingInfoDIJobs) Validate() error {
	return dara.Validate(s)
}
