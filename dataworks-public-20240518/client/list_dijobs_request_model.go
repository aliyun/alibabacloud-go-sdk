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
	SetSpecType(v string) *ListDIJobsRequest
	GetSpecType() *string
}

type ListDIJobsRequest struct {
	// The type of the destination data source. If you do not specify this parameter, jobs are not filtered by this criterion. Valid values: `Hologres`, `OSS-HDFS`, `OSS`, `MaxCompute`, `LogHub`, `StarRocks`, `DataHub`, `AnalyticDB_For_MySQL`, `Kafka`, and `Hive`.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The synchronization type. Valid values:
	//
	// - `FullAndRealtimeIncremental`: full and real-time incremental synchronization
	//
	// - `RealtimeIncremental`: real-time incremental synchronization
	//
	// - `Full`: full synchronization
	//
	// - `OfflineIncremental`: offline incremental synchronization
	//
	// - `FullAndOfflineIncremental`: full and offline incremental synchronization
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The name of the Data Integration job.
	//
	// The name must be unique within the DataWorks workspace.
	//
	// example:
	//
	// test_export_01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Pages are numbered starting from 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Default: 10. Maximum: 100.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the DataWorks workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1967
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The type of the source data source. If you do not specify this parameter, jobs are not filtered by this criterion. Valid values: `PolarDB`, `MySQL`, `Kafka`, `LogHub`, `Hologres`, `Oracle`, `OceanBase`, `MongoDB`, `RedShift`, `Hive`, `SQLServer`, `Doris`, and `ClickHouse`.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// The configuration type of the job. Valid values: `FILESPEC`, `CLASSIC`, and `ALL`. `FILESPEC` indicates a new job type configured based on a structured file specification. `CLASSIC` indicates a job configured in the traditional mode. If you set this parameter to `ALL`, jobs of both types are returned.
	SpecType *string `json:"SpecType,omitempty" xml:"SpecType,omitempty"`
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

func (s *ListDIJobsRequest) GetSpecType() *string {
	return s.SpecType
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

func (s *ListDIJobsRequest) SetSpecType(v string) *ListDIJobsRequest {
	s.SpecType = &v
	return s
}

func (s *ListDIJobsRequest) Validate() error {
	return dara.Validate(s)
}
