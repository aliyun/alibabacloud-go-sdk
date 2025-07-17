// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationDataSourceType(v string) *ListDIJobsRequest
	GetDestinationDataSourceType() *string
	SetMigrationType(v string) *ListDIJobsRequest
	GetMigrationType() *string
	SetName(v string) *ListDIJobsRequest
	GetName() *string
	SetPageNumber(v int64) *ListDIJobsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListDIJobsRequest
	GetPageSize() *int64
	SetProjectId(v int64) *ListDIJobsRequest
	GetProjectId() *int64
	SetSourceDataSourceType(v string) *ListDIJobsRequest
	GetSourceDataSourceType() *string
}

type ListDIJobsRequest struct {
	// The destination type. Valid values: Hologres, OSS-HDFS, OSS, MaxCompute, Loghub, STARROCKS, Datahub, ANALYTICDB_FOR_MYSQL, Kafka, and Hive. If you do not configure this parameter, the API operation queries synchronization tasks that use all type of destinations.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
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
	// The name of the export task.
	//
	// The name of each export task must be unique. You must make sure that the names of the export tasks in the current workspace are unique.
	//
	// example:
	//
	// test_export_01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Pages start from page 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default value: 10. Maximum value: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1967
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The source type. Valid values: PolarDB, MySQL, Kafka, Loghub, Hologres, Oracle, OceanBase, MongoDB, RedShift, Hive, SqlServer, Doris, and ClickHouse. If you do not configure this parameter, the API operation queries synchronization tasks that use all types of sources.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
}

func (s ListDIJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobsRequest) GoString() string {
	return s.String()
}

func (s *ListDIJobsRequest) GetDestinationDataSourceType() *string {
	return s.DestinationDataSourceType
}

func (s *ListDIJobsRequest) GetMigrationType() *string {
	return s.MigrationType
}

func (s *ListDIJobsRequest) GetName() *string {
	return s.Name
}

func (s *ListDIJobsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListDIJobsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListDIJobsRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListDIJobsRequest) GetSourceDataSourceType() *string {
	return s.SourceDataSourceType
}

func (s *ListDIJobsRequest) SetDestinationDataSourceType(v string) *ListDIJobsRequest {
	s.DestinationDataSourceType = &v
	return s
}

func (s *ListDIJobsRequest) SetMigrationType(v string) *ListDIJobsRequest {
	s.MigrationType = &v
	return s
}

func (s *ListDIJobsRequest) SetName(v string) *ListDIJobsRequest {
	s.Name = &v
	return s
}

func (s *ListDIJobsRequest) SetPageNumber(v int64) *ListDIJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDIJobsRequest) SetPageSize(v int64) *ListDIJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListDIJobsRequest) SetProjectId(v int64) *ListDIJobsRequest {
	s.ProjectId = &v
	return s
}

func (s *ListDIJobsRequest) SetSourceDataSourceType(v string) *ListDIJobsRequest {
	s.SourceDataSourceType = &v
	return s
}

func (s *ListDIJobsRequest) Validate() error {
	return dara.Validate(s)
}
