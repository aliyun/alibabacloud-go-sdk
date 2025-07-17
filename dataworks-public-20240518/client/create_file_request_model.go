// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedSettings(v string) *CreateFileRequest
	GetAdvancedSettings() *string
	SetApplyScheduleImmediately(v bool) *CreateFileRequest
	GetApplyScheduleImmediately() *bool
	SetAutoParsing(v bool) *CreateFileRequest
	GetAutoParsing() *bool
	SetAutoRerunIntervalMillis(v int32) *CreateFileRequest
	GetAutoRerunIntervalMillis() *int32
	SetAutoRerunTimes(v int32) *CreateFileRequest
	GetAutoRerunTimes() *int32
	SetConnectionName(v string) *CreateFileRequest
	GetConnectionName() *string
	SetContent(v string) *CreateFileRequest
	GetContent() *string
	SetCreateFolderIfNotExists(v bool) *CreateFileRequest
	GetCreateFolderIfNotExists() *bool
	SetCronExpress(v string) *CreateFileRequest
	GetCronExpress() *string
	SetCycleType(v string) *CreateFileRequest
	GetCycleType() *string
	SetDependentNodeIdList(v string) *CreateFileRequest
	GetDependentNodeIdList() *string
	SetDependentType(v string) *CreateFileRequest
	GetDependentType() *string
	SetEndEffectDate(v int64) *CreateFileRequest
	GetEndEffectDate() *int64
	SetFileDescription(v string) *CreateFileRequest
	GetFileDescription() *string
	SetFileFolderPath(v string) *CreateFileRequest
	GetFileFolderPath() *string
	SetFileName(v string) *CreateFileRequest
	GetFileName() *string
	SetFileType(v int32) *CreateFileRequest
	GetFileType() *int32
	SetIgnoreParentSkipRunningProperty(v bool) *CreateFileRequest
	GetIgnoreParentSkipRunningProperty() *bool
	SetImageId(v string) *CreateFileRequest
	GetImageId() *string
	SetInputList(v string) *CreateFileRequest
	GetInputList() *string
	SetInputParameters(v string) *CreateFileRequest
	GetInputParameters() *string
	SetOutputParameters(v string) *CreateFileRequest
	GetOutputParameters() *string
	SetOwner(v string) *CreateFileRequest
	GetOwner() *string
	SetParaValue(v string) *CreateFileRequest
	GetParaValue() *string
	SetProjectId(v int64) *CreateFileRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *CreateFileRequest
	GetProjectIdentifier() *string
	SetRerunMode(v string) *CreateFileRequest
	GetRerunMode() *string
	SetResourceGroupId(v int64) *CreateFileRequest
	GetResourceGroupId() *int64
	SetResourceGroupIdentifier(v string) *CreateFileRequest
	GetResourceGroupIdentifier() *string
	SetSchedulerType(v string) *CreateFileRequest
	GetSchedulerType() *string
	SetStartEffectDate(v int64) *CreateFileRequest
	GetStartEffectDate() *int64
	SetStartImmediately(v bool) *CreateFileRequest
	GetStartImmediately() *bool
	SetStop(v bool) *CreateFileRequest
	GetStop() *bool
	SetTimeout(v int32) *CreateFileRequest
	GetTimeout() *int32
}

type CreateFileRequest struct {
	// example:
	//
	// {"queue":"default","SPARK_CONF":"--conf spark.driver.memory=2g"}
	AdvancedSettings *string `json:"AdvancedSettings,omitempty" xml:"AdvancedSettings,omitempty"`
	// example:
	//
	// true
	ApplyScheduleImmediately *bool `json:"ApplyScheduleImmediately,omitempty" xml:"ApplyScheduleImmediately,omitempty"`
	// example:
	//
	// true
	AutoParsing *bool `json:"AutoParsing,omitempty" xml:"AutoParsing,omitempty"`
	// example:
	//
	// 120000
	AutoRerunIntervalMillis *int32 `json:"AutoRerunIntervalMillis,omitempty" xml:"AutoRerunIntervalMillis,omitempty"`
	// example:
	//
	// 3
	AutoRerunTimes *int32 `json:"AutoRerunTimes,omitempty" xml:"AutoRerunTimes,omitempty"`
	// example:
	//
	// odps_source
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// example:
	//
	// SHOW TABLES;
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// false
	CreateFolderIfNotExists *bool `json:"CreateFolderIfNotExists,omitempty" xml:"CreateFolderIfNotExists,omitempty"`
	// example:
	//
	// 00 05 00 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// example:
	//
	// DAY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// example:
	//
	// abc
	DependentNodeIdList *string `json:"DependentNodeIdList,omitempty" xml:"DependentNodeIdList,omitempty"`
	// example:
	//
	// NONE
	DependentType *string `json:"DependentType,omitempty" xml:"DependentType,omitempty"`
	// example:
	//
	// 1671694850000
	EndEffectDate   *int64  `json:"EndEffectDate,omitempty" xml:"EndEffectDate,omitempty"`
	FileDescription *string `json:"FileDescription,omitempty" xml:"FileDescription,omitempty"`
	// example:
	//
	// Business_process/First_Business_Process/MaxCompute/Folder_1/Folder_2
	FileFolderPath *string `json:"FileFolderPath,omitempty" xml:"FileFolderPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// File name
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	FileType *int32 `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// example:
	//
	// false
	IgnoreParentSkipRunningProperty *bool `json:"IgnoreParentSkipRunningProperty,omitempty" xml:"IgnoreParentSkipRunningProperty,omitempty"`
	// example:
	//
	// m-bp1h4b5a8ogkbll2f3tr
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// project_root,project.file1,project.001_out
	InputList *string `json:"InputList,omitempty" xml:"InputList,omitempty"`
	// example:
	//
	// [{"ValueSource": "project_001.first_node:bizdate_param","ParameterName": "bizdate_input"}]
	InputParameters *string `json:"InputParameters,omitempty" xml:"InputParameters,omitempty"`
	// example:
	//
	// [{"Type": 1,"Value": "${bizdate}","ParameterName": "bizdate_param"}]
	OutputParameters *string `json:"OutputParameters,omitempty" xml:"OutputParameters,omitempty"`
	// example:
	//
	// 1000000000001
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// a=x b=y
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// example:
	//
	// dw_project
	ProjectIdentifier *string `json:"ProjectIdentifier,omitempty" xml:"ProjectIdentifier,omitempty"`
	// example:
	//
	// ALL_ALLOWED
	RerunMode *string `json:"RerunMode,omitempty" xml:"RerunMode,omitempty"`
	// example:
	//
	// 375827434852437
	ResourceGroupId *int64 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// S_res_group_559_1613715566828
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// example:
	//
	// 1671608450000
	StartEffectDate *int64 `json:"StartEffectDate,omitempty" xml:"StartEffectDate,omitempty"`
	// example:
	//
	// true
	StartImmediately *bool `json:"StartImmediately,omitempty" xml:"StartImmediately,omitempty"`
	// example:
	//
	// false
	Stop *bool `json:"Stop,omitempty" xml:"Stop,omitempty"`
	// example:
	//
	// 1
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s CreateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileRequest) GoString() string {
	return s.String()
}

func (s *CreateFileRequest) GetAdvancedSettings() *string {
	return s.AdvancedSettings
}

func (s *CreateFileRequest) GetApplyScheduleImmediately() *bool {
	return s.ApplyScheduleImmediately
}

func (s *CreateFileRequest) GetAutoParsing() *bool {
	return s.AutoParsing
}

func (s *CreateFileRequest) GetAutoRerunIntervalMillis() *int32 {
	return s.AutoRerunIntervalMillis
}

func (s *CreateFileRequest) GetAutoRerunTimes() *int32 {
	return s.AutoRerunTimes
}

func (s *CreateFileRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *CreateFileRequest) GetContent() *string {
	return s.Content
}

func (s *CreateFileRequest) GetCreateFolderIfNotExists() *bool {
	return s.CreateFolderIfNotExists
}

func (s *CreateFileRequest) GetCronExpress() *string {
	return s.CronExpress
}

func (s *CreateFileRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *CreateFileRequest) GetDependentNodeIdList() *string {
	return s.DependentNodeIdList
}

func (s *CreateFileRequest) GetDependentType() *string {
	return s.DependentType
}

func (s *CreateFileRequest) GetEndEffectDate() *int64 {
	return s.EndEffectDate
}

func (s *CreateFileRequest) GetFileDescription() *string {
	return s.FileDescription
}

func (s *CreateFileRequest) GetFileFolderPath() *string {
	return s.FileFolderPath
}

func (s *CreateFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *CreateFileRequest) GetFileType() *int32 {
	return s.FileType
}

func (s *CreateFileRequest) GetIgnoreParentSkipRunningProperty() *bool {
	return s.IgnoreParentSkipRunningProperty
}

func (s *CreateFileRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateFileRequest) GetInputList() *string {
	return s.InputList
}

func (s *CreateFileRequest) GetInputParameters() *string {
	return s.InputParameters
}

func (s *CreateFileRequest) GetOutputParameters() *string {
	return s.OutputParameters
}

func (s *CreateFileRequest) GetOwner() *string {
	return s.Owner
}

func (s *CreateFileRequest) GetParaValue() *string {
	return s.ParaValue
}

func (s *CreateFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *CreateFileRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *CreateFileRequest) GetRerunMode() *string {
	return s.RerunMode
}

func (s *CreateFileRequest) GetResourceGroupId() *int64 {
	return s.ResourceGroupId
}

func (s *CreateFileRequest) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *CreateFileRequest) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *CreateFileRequest) GetStartEffectDate() *int64 {
	return s.StartEffectDate
}

func (s *CreateFileRequest) GetStartImmediately() *bool {
	return s.StartImmediately
}

func (s *CreateFileRequest) GetStop() *bool {
	return s.Stop
}

func (s *CreateFileRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *CreateFileRequest) SetAdvancedSettings(v string) *CreateFileRequest {
	s.AdvancedSettings = &v
	return s
}

func (s *CreateFileRequest) SetApplyScheduleImmediately(v bool) *CreateFileRequest {
	s.ApplyScheduleImmediately = &v
	return s
}

func (s *CreateFileRequest) SetAutoParsing(v bool) *CreateFileRequest {
	s.AutoParsing = &v
	return s
}

func (s *CreateFileRequest) SetAutoRerunIntervalMillis(v int32) *CreateFileRequest {
	s.AutoRerunIntervalMillis = &v
	return s
}

func (s *CreateFileRequest) SetAutoRerunTimes(v int32) *CreateFileRequest {
	s.AutoRerunTimes = &v
	return s
}

func (s *CreateFileRequest) SetConnectionName(v string) *CreateFileRequest {
	s.ConnectionName = &v
	return s
}

func (s *CreateFileRequest) SetContent(v string) *CreateFileRequest {
	s.Content = &v
	return s
}

func (s *CreateFileRequest) SetCreateFolderIfNotExists(v bool) *CreateFileRequest {
	s.CreateFolderIfNotExists = &v
	return s
}

func (s *CreateFileRequest) SetCronExpress(v string) *CreateFileRequest {
	s.CronExpress = &v
	return s
}

func (s *CreateFileRequest) SetCycleType(v string) *CreateFileRequest {
	s.CycleType = &v
	return s
}

func (s *CreateFileRequest) SetDependentNodeIdList(v string) *CreateFileRequest {
	s.DependentNodeIdList = &v
	return s
}

func (s *CreateFileRequest) SetDependentType(v string) *CreateFileRequest {
	s.DependentType = &v
	return s
}

func (s *CreateFileRequest) SetEndEffectDate(v int64) *CreateFileRequest {
	s.EndEffectDate = &v
	return s
}

func (s *CreateFileRequest) SetFileDescription(v string) *CreateFileRequest {
	s.FileDescription = &v
	return s
}

func (s *CreateFileRequest) SetFileFolderPath(v string) *CreateFileRequest {
	s.FileFolderPath = &v
	return s
}

func (s *CreateFileRequest) SetFileName(v string) *CreateFileRequest {
	s.FileName = &v
	return s
}

func (s *CreateFileRequest) SetFileType(v int32) *CreateFileRequest {
	s.FileType = &v
	return s
}

func (s *CreateFileRequest) SetIgnoreParentSkipRunningProperty(v bool) *CreateFileRequest {
	s.IgnoreParentSkipRunningProperty = &v
	return s
}

func (s *CreateFileRequest) SetImageId(v string) *CreateFileRequest {
	s.ImageId = &v
	return s
}

func (s *CreateFileRequest) SetInputList(v string) *CreateFileRequest {
	s.InputList = &v
	return s
}

func (s *CreateFileRequest) SetInputParameters(v string) *CreateFileRequest {
	s.InputParameters = &v
	return s
}

func (s *CreateFileRequest) SetOutputParameters(v string) *CreateFileRequest {
	s.OutputParameters = &v
	return s
}

func (s *CreateFileRequest) SetOwner(v string) *CreateFileRequest {
	s.Owner = &v
	return s
}

func (s *CreateFileRequest) SetParaValue(v string) *CreateFileRequest {
	s.ParaValue = &v
	return s
}

func (s *CreateFileRequest) SetProjectId(v int64) *CreateFileRequest {
	s.ProjectId = &v
	return s
}

func (s *CreateFileRequest) SetProjectIdentifier(v string) *CreateFileRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *CreateFileRequest) SetRerunMode(v string) *CreateFileRequest {
	s.RerunMode = &v
	return s
}

func (s *CreateFileRequest) SetResourceGroupId(v int64) *CreateFileRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateFileRequest) SetResourceGroupIdentifier(v string) *CreateFileRequest {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *CreateFileRequest) SetSchedulerType(v string) *CreateFileRequest {
	s.SchedulerType = &v
	return s
}

func (s *CreateFileRequest) SetStartEffectDate(v int64) *CreateFileRequest {
	s.StartEffectDate = &v
	return s
}

func (s *CreateFileRequest) SetStartImmediately(v bool) *CreateFileRequest {
	s.StartImmediately = &v
	return s
}

func (s *CreateFileRequest) SetStop(v bool) *CreateFileRequest {
	s.Stop = &v
	return s
}

func (s *CreateFileRequest) SetTimeout(v int32) *CreateFileRequest {
	s.Timeout = &v
	return s
}

func (s *CreateFileRequest) Validate() error {
	return dara.Validate(s)
}
