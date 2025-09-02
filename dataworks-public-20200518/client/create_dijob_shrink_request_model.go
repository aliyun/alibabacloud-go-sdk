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
	SetMigrationType(v string) *CreateDIJobShrinkRequest
	GetMigrationType() *string
	SetProjectId(v int64) *CreateDIJobShrinkRequest
	GetProjectId() *int64
	SetResourceSettingsShrink(v string) *CreateDIJobShrinkRequest
	GetResourceSettingsShrink() *string
	SetSourceDataSourceSettingsShrink(v string) *CreateDIJobShrinkRequest
	GetSourceDataSourceSettingsShrink() *string
	SetSourceDataSourceType(v string) *CreateDIJobShrinkRequest
	GetSourceDataSourceType() *string
	SetSystemDebug(v string) *CreateDIJobShrinkRequest
	GetSystemDebug() *string
	SetTableMappingsShrink(v string) *CreateDIJobShrinkRequest
	GetTableMappingsShrink() *string
	SetTransformationRulesShrink(v string) *CreateDIJobShrinkRequest
	GetTransformationRulesShrink() *string
}

type CreateDIJobShrinkRequest struct {
	// The description of the synchronization task.
	//
	// example:
	//
	// Synchronize mysql to hologres
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The settings of the destination. Only a single destination is supported.
	DestinationDataSourceSettingsShrink *string `json:"DestinationDataSourceSettings,omitempty" xml:"DestinationDataSourceSettings,omitempty"`
	// The destination type. Valid values: Hologres and Hive.
	//
	// example:
	//
	// Hologres
	DestinationDataSourceType *string `json:"DestinationDataSourceType,omitempty" xml:"DestinationDataSourceType,omitempty"`
	// The name of the synchronization task.
	//
	// example:
	//
	// mysql_to_holo_sync_8772
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The settings for the dimension of the synchronization task. The settings include processing policies for DDL messages, policies for data type mappings between source fields and destination fields, and runtime parameters of the synchronization task.
	JobSettingsShrink *string `json:"JobSettings,omitempty" xml:"JobSettings,omitempty"`
	// The synchronization type. Valid values:
	//
	// 	- FullAndRealtimeIncremental (one-time full synchronization and real-time incremental synchronization)
	//
	// 	- RealtimeIncremental (real-time incremental synchronization)
	//
	// 	- Full (full synchronization)
	//
	// 	- OfflineIncremental (batch incremental synchronization)
	//
	// 	- FullAndOfflineIncremental (one-time full synchronization and batch incremental synchronization)
	//
	// example:
	//
	// FullAndRealtimeIncremental
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	// The DataWorks workspace ID. You can call the [ListProjects](https://help.aliyun.com/document_detail/178393.html) operation to obtain the ID.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The resource settings.
	ResourceSettingsShrink *string `json:"ResourceSettings,omitempty" xml:"ResourceSettings,omitempty"`
	// The settings of the source. Only a single source is supported.
	SourceDataSourceSettingsShrink *string `json:"SourceDataSourceSettings,omitempty" xml:"SourceDataSourceSettings,omitempty"`
	// The source type. Set this parameter to MySQL.
	//
	// example:
	//
	// MySQL
	SourceDataSourceType *string `json:"SourceDataSourceType,omitempty" xml:"SourceDataSourceType,omitempty"`
	// Specifies whether to perform system debugging. Valid values: true and false. Default value: false.
	//
	// example:
	//
	// false
	SystemDebug *string `json:"SystemDebug,omitempty" xml:"SystemDebug,omitempty"`
	// The list of mappings between rules used to select synchronization objects in the source and transformation rules applied to the selected synchronization objects. Each entry in the list displays a mapping between a rule used to select synchronization objects and a transformation rule applied to the selected synchronization objects.
	TableMappingsShrink *string `json:"TableMappings,omitempty" xml:"TableMappings,omitempty"`
	// The list of transformation rules that you want to apply to the synchronization objects selected from the source. Each entry in the list defines a transformation rule.
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

func (s *CreateDIJobShrinkRequest) GetMigrationType() *string {
	return s.MigrationType
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

func (s *CreateDIJobShrinkRequest) GetSystemDebug() *string {
	return s.SystemDebug
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

func (s *CreateDIJobShrinkRequest) SetMigrationType(v string) *CreateDIJobShrinkRequest {
	s.MigrationType = &v
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

func (s *CreateDIJobShrinkRequest) SetSystemDebug(v string) *CreateDIJobShrinkRequest {
	s.SystemDebug = &v
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
