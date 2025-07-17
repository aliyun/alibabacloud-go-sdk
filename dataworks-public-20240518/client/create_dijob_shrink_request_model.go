// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDIJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDIJobShrinkRequest
	GetDescription() *string
	SetDestinationDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest
	GetDestinationDataSourceSettingsShrink() *string
	SetDestinationDataSourceType(v string) *CreateDIJobShrinkRequest
	GetDestinationDataSourceType() *string
	SetJobName(v string) *CreateDIJobShrinkRequest
	GetJobName() *string
	SetJobSettingsShrink(v string) *CreateDIJobShrinkRequest
	GetJobSettingsShrink() *string
	SetJobType(v string) *CreateDIJobShrinkRequest
	GetJobType() *string
	SetMigrationType(v string) *CreateDIJobShrinkRequest
	GetMigrationType() *string
	SetName(v string) *CreateDIJobShrinkRequest
	GetName() *string
	SetProjectId(v int64) *CreateDIJobShrinkRequest
	GetProjectId() *int64
	SetResourceSettingsShrink(v string) *CreateDIJobShrinkRequest
	GetResourceSettingsShrink() *string
	SetSourceDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest
	GetSourceDataSourceSettingsShrink() *string
	SetSourceDataSourceType(v string) *CreateDIJobShrinkRequest
	GetSourceDataSourceType() *string
	SetTableMappingsShrink(v string) *CreateDIJobShrinkRequest
	GetTableMappingsShrink() *string
	SetTransformationRulesShrink(v string) *CreateDIJobShrinkRequest
	GetTransformationRulesShrink() *string
}

type CreateDIJobShrinkRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	DestinationDataSourceSettingsShrink *string `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty"`
	// The destination type. Valid values: Hologres, OSS-HDFS, OSS, MaxCompute, LogHub, StarRocks, DataHub, AnalyticDB for MySQL, Kafka, and Hive.
	//
	// This parameter is required.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated and is replaced by the Name parameter.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	JobName           *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	JobSettingsShrink *string `json:"JobSettings,omitempty" xml:"JobSettings,omitempty"`
	// The type of the synchronization task. Valid values:
	//
	// 	- DatabaseRealtimeMigration: A real-time synchronization task used to synchronize only full data, only incremental data, or full and incremental data in multiple tables of multiple databases in the source.
	//
	// 	- DatabaseOfflineMigration: A batch synchronization task used to synchronize only full data, only incremental data, or full and incremental data in multiple tables of multiple databases in the source.
	//
	// 	- SingleTableRealtimeMigration: A real-time synchronization task used to synchronize data only in a single table in the source.
	//
	// example:
	//
	// DatabaseRealtimeMigration
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// The synchronization type. Valid values:
	//
	// 	- FullAndRealtimeIncremental
	//
	// 	- RealtimeIncremental
	//
	// 	- Full
	//
	// 	- OfflineIncremental
	//
	// 	- FullAndOfflineIncremental
	//
	// This parameter is required.
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The name of the synchronization task.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The DataWorks workspace ID. You can log on to the [DataWorks console](https://workbench.data.aliyun.com/console) and go to the Workspace page to obtain the ID.
	//
	// You must configure this parameter to specify the DataWorks workspace to which the API operation is applied.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// This parameter is required.
	ResourceSettingsShrink *string `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty"`
	// This parameter is required.
	SourceDataSourceSettingsShrink *string `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty"`
	// The source type. Valid values: PolarDB, MySQL, Kafka, LogHub, Hologres, Oracle, OceanBase, MongoDB, Redshift, Hive, SQL Server, Doris, and ClickHouse.
	//
	// This parameter is required.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// This parameter is required.
	TableMappingsShrink       *string `json:"TableMappings,omitempty" xml:"TableMappings,omitempty"`
	TransformationRulesShrink *string `json:"TransformationRules,omitempty" xml:"TransformationRules,omitempty"`
}

func (s CreateDIJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDIJobShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDIJobShrinkRequest) GetDestinationDataSourceSettingsShrink() *string {
	return s.DestinationDataSourceSettingsShrink
}

func (s *CreateDIJobShrinkRequest) GetDestinationDataSourceType() *string {
	return s.DestinationDataSourceType
}

func (s *CreateDIJobShrinkRequest) GetJobName() *string {
	return s.JobName
}

func (s *CreateDIJobShrinkRequest) GetJobSettingsShrink() *string {
	return s.JobSettingsShrink
}

func (s *CreateDIJobShrinkRequest) GetJobType() *string {
	return s.JobType
}

func (s *CreateDIJobShrinkRequest) GetMigrationType() *string {
	return s.MigrationType
}

func (s *CreateDIJobShrinkRequest) GetName() *string {
	return s.Name
}

func (s *CreateDIJobShrinkRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateDIJobShrinkRequest) GetResourceSettingsShrink() *string {
	return s.ResourceSettingsShrink
}

func (s *CreateDIJobShrinkRequest) GetSourceDataSourceSettingsShrink() *string {
	return s.SourceDataSourceSettingsShrink
}

func (s *CreateDIJobShrinkRequest) GetSourceDataSourceType() *string {
	return s.SourceDataSourceType
}

func (s *CreateDIJobShrinkRequest) GetTableMappingsShrink() *string {
	return s.TableMappingsShrink
}

func (s *CreateDIJobShrinkRequest) GetTransformationRulesShrink() *string {
	return s.TransformationRulesShrink
}

func (s *CreateDIJobShrinkRequest) SetDescription(v string) *CreateDIJobShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetDestinationDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.DestinationDataSourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetDestinationDataSourceType(v string) *CreateDIJobShrinkRequest {
	s.DestinationDataSourceType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetJobName(v string) *CreateDIJobShrinkRequest {
	s.JobName = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetJobSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.JobSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetJobType(v string) *CreateDIJobShrinkRequest {
	s.JobType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetMigrationType(v string) *CreateDIJobShrinkRequest {
	s.MigrationType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetName(v string) *CreateDIJobShrinkRequest {
	s.Name = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetProjectId(v int64) *CreateDIJobShrinkRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetResourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.ResourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetSourceDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest {
	s.SourceDataSourceSettingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetSourceDataSourceType(v string) *CreateDIJobShrinkRequest {
	s.SourceDataSourceType = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetTableMappingsShrink(v string) *CreateDIJobShrinkRequest {
	s.TableMappingsShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) SetTransformationRulesShrink(v string) *CreateDIJobShrinkRequest {
	s.TransformationRulesShrink = &v
	return s
}

func (s *CreateDIJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
