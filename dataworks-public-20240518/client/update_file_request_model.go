// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdvancedSettings(v string) *UpdateFileRequest
	GetAdvancedSettings() *string
	SetApplyScheduleImmediately(v bool) *UpdateFileRequest
	GetApplyScheduleImmediately() *bool
	SetAutoParsing(v bool) *UpdateFileRequest
	GetAutoParsing() *bool
	SetAutoRerunIntervalMillis(v int32) *UpdateFileRequest
	GetAutoRerunIntervalMillis() *int32
	SetAutoRerunTimes(v int32) *UpdateFileRequest
	GetAutoRerunTimes() *int32
	SetConnectionName(v string) *UpdateFileRequest
	GetConnectionName() *string
	SetContent(v string) *UpdateFileRequest
	GetContent() *string
	SetCronExpress(v string) *UpdateFileRequest
	GetCronExpress() *string
	SetCycleType(v string) *UpdateFileRequest
	GetCycleType() *string
	SetDependentNodeIdList(v string) *UpdateFileRequest
	GetDependentNodeIdList() *string
	SetDependentType(v string) *UpdateFileRequest
	GetDependentType() *string
	SetEndEffectDate(v int64) *UpdateFileRequest
	GetEndEffectDate() *int64
	SetFileDescription(v string) *UpdateFileRequest
	GetFileDescription() *string
	SetFileFolderPath(v string) *UpdateFileRequest
	GetFileFolderPath() *string
	SetFileId(v int64) *UpdateFileRequest
	GetFileId() *int64
	SetFileName(v string) *UpdateFileRequest
	GetFileName() *string
	SetIgnoreParentSkipRunningProperty(v bool) *UpdateFileRequest
	GetIgnoreParentSkipRunningProperty() *bool
	SetImageId(v string) *UpdateFileRequest
	GetImageId() *string
	SetInputList(v string) *UpdateFileRequest
	GetInputList() *string
	SetInputParameters(v string) *UpdateFileRequest
	GetInputParameters() *string
	SetOutputList(v string) *UpdateFileRequest
	GetOutputList() *string
	SetOutputParameters(v string) *UpdateFileRequest
	GetOutputParameters() *string
	SetOwner(v string) *UpdateFileRequest
	GetOwner() *string
	SetParaValue(v string) *UpdateFileRequest
	GetParaValue() *string
	SetProjectId(v int64) *UpdateFileRequest
	GetProjectId() *int64
	SetProjectIdentifier(v string) *UpdateFileRequest
	GetProjectIdentifier() *string
	SetRerunMode(v string) *UpdateFileRequest
	GetRerunMode() *string
	SetResourceGroupIdentifier(v string) *UpdateFileRequest
	GetResourceGroupIdentifier() *string
	SetSchedulerType(v string) *UpdateFileRequest
	GetSchedulerType() *string
	SetStartEffectDate(v int64) *UpdateFileRequest
	GetStartEffectDate() *int64
	SetStartImmediately(v bool) *UpdateFileRequest
	GetStartImmediately() *bool
	SetStop(v bool) *UpdateFileRequest
	GetStop() *bool
	SetTimeout(v int32) *UpdateFileRequest
	GetTimeout() *int32
}

type UpdateFileRequest struct {
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
	// SELECT "1";
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 00 00-59/5 1-23 	- 	- ?
	CronExpress *string `json:"CronExpress,omitempty" xml:"CronExpress,omitempty"`
	// example:
	//
	// NOT_DAY
	CycleType *string `json:"CycleType,omitempty" xml:"CycleType,omitempty"`
	// example:
	//
	// 5,10,15,20
	DependentNodeIdList *string `json:"DependentNodeIdList,omitempty" xml:"DependentNodeIdList,omitempty"`
	// example:
	//
	// USER_DEFINE
	DependentType *string `json:"DependentType,omitempty" xml:"DependentType,omitempty"`
	// example:
	//
	// 4155787800000
	EndEffectDate *int64 `json:"EndEffectDate,omitempty" xml:"EndEffectDate,omitempty"`
	// example:
	//
	// Here is the file description
	FileDescription *string `json:"FileDescription,omitempty" xml:"FileDescription,omitempty"`
	// example:
	//
	// Business_process/First_Business_Process/data_integration/Folder_1/Folder_2
	FileFolderPath *string `json:"FileFolderPath,omitempty" xml:"FileFolderPath,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 100000001
	FileId *int64 `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// example:
	//
	// ods_user_info_d
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// true
	IgnoreParentSkipRunningProperty *bool `json:"IgnoreParentSkipRunningProperty,omitempty" xml:"IgnoreParentSkipRunningProperty,omitempty"`
	// example:
	//
	// m-uf6d7npxk1hhek8ng0cb
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
	// dw_project.ods_user_info_d
	OutputList *string `json:"OutputList,omitempty" xml:"OutputList,omitempty"`
	// example:
	//
	// [{"Type": 1,"Value": "${bizdate}","ParameterName": "bizdate_param"}]
	OutputParameters *string `json:"OutputParameters,omitempty" xml:"OutputParameters,omitempty"`
	// example:
	//
	// 18023848927592
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// x=a y=b z=c
	ParaValue *string `json:"ParaValue,omitempty" xml:"ParaValue,omitempty"`
	// example:
	//
	// 100001
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
	// default_group
	ResourceGroupIdentifier *string `json:"ResourceGroupIdentifier,omitempty" xml:"ResourceGroupIdentifier,omitempty"`
	// example:
	//
	// NORMAL
	SchedulerType *string `json:"SchedulerType,omitempty" xml:"SchedulerType,omitempty"`
	// example:
	//
	// 936923400000
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

func (s UpdateFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileRequest) GetAdvancedSettings() *string {
	return s.AdvancedSettings
}

func (s *UpdateFileRequest) GetApplyScheduleImmediately() *bool {
	return s.ApplyScheduleImmediately
}

func (s *UpdateFileRequest) GetAutoParsing() *bool {
	return s.AutoParsing
}

func (s *UpdateFileRequest) GetAutoRerunIntervalMillis() *int32 {
	return s.AutoRerunIntervalMillis
}

func (s *UpdateFileRequest) GetAutoRerunTimes() *int32 {
	return s.AutoRerunTimes
}

func (s *UpdateFileRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *UpdateFileRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateFileRequest) GetCronExpress() *string {
	return s.CronExpress
}

func (s *UpdateFileRequest) GetCycleType() *string {
	return s.CycleType
}

func (s *UpdateFileRequest) GetDependentNodeIdList() *string {
	return s.DependentNodeIdList
}

func (s *UpdateFileRequest) GetDependentType() *string {
	return s.DependentType
}

func (s *UpdateFileRequest) GetEndEffectDate() *int64 {
	return s.EndEffectDate
}

func (s *UpdateFileRequest) GetFileDescription() *string {
	return s.FileDescription
}

func (s *UpdateFileRequest) GetFileFolderPath() *string {
	return s.FileFolderPath
}

func (s *UpdateFileRequest) GetFileId() *int64 {
	return s.FileId
}

func (s *UpdateFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *UpdateFileRequest) GetIgnoreParentSkipRunningProperty() *bool {
	return s.IgnoreParentSkipRunningProperty
}

func (s *UpdateFileRequest) GetImageId() *string {
	return s.ImageId
}

func (s *UpdateFileRequest) GetInputList() *string {
	return s.InputList
}

func (s *UpdateFileRequest) GetInputParameters() *string {
	return s.InputParameters
}

func (s *UpdateFileRequest) GetOutputList() *string {
	return s.OutputList
}

func (s *UpdateFileRequest) GetOutputParameters() *string {
	return s.OutputParameters
}

func (s *UpdateFileRequest) GetOwner() *string {
	return s.Owner
}

func (s *UpdateFileRequest) GetParaValue() *string {
	return s.ParaValue
}

func (s *UpdateFileRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *UpdateFileRequest) GetProjectIdentifier() *string {
	return s.ProjectIdentifier
}

func (s *UpdateFileRequest) GetRerunMode() *string {
	return s.RerunMode
}

func (s *UpdateFileRequest) GetResourceGroupIdentifier() *string {
	return s.ResourceGroupIdentifier
}

func (s *UpdateFileRequest) GetSchedulerType() *string {
	return s.SchedulerType
}

func (s *UpdateFileRequest) GetStartEffectDate() *int64 {
	return s.StartEffectDate
}

func (s *UpdateFileRequest) GetStartImmediately() *bool {
	return s.StartImmediately
}

func (s *UpdateFileRequest) GetStop() *bool {
	return s.Stop
}

func (s *UpdateFileRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *UpdateFileRequest) SetAdvancedSettings(v string) *UpdateFileRequest {
	s.AdvancedSettings = &v
	return s
}

func (s *UpdateFileRequest) SetApplyScheduleImmediately(v bool) *UpdateFileRequest {
	s.ApplyScheduleImmediately = &v
	return s
}

func (s *UpdateFileRequest) SetAutoParsing(v bool) *UpdateFileRequest {
	s.AutoParsing = &v
	return s
}

func (s *UpdateFileRequest) SetAutoRerunIntervalMillis(v int32) *UpdateFileRequest {
	s.AutoRerunIntervalMillis = &v
	return s
}

func (s *UpdateFileRequest) SetAutoRerunTimes(v int32) *UpdateFileRequest {
	s.AutoRerunTimes = &v
	return s
}

func (s *UpdateFileRequest) SetConnectionName(v string) *UpdateFileRequest {
	s.ConnectionName = &v
	return s
}

func (s *UpdateFileRequest) SetContent(v string) *UpdateFileRequest {
	s.Content = &v
	return s
}

func (s *UpdateFileRequest) SetCronExpress(v string) *UpdateFileRequest {
	s.CronExpress = &v
	return s
}

func (s *UpdateFileRequest) SetCycleType(v string) *UpdateFileRequest {
	s.CycleType = &v
	return s
}

func (s *UpdateFileRequest) SetDependentNodeIdList(v string) *UpdateFileRequest {
	s.DependentNodeIdList = &v
	return s
}

func (s *UpdateFileRequest) SetDependentType(v string) *UpdateFileRequest {
	s.DependentType = &v
	return s
}

func (s *UpdateFileRequest) SetEndEffectDate(v int64) *UpdateFileRequest {
	s.EndEffectDate = &v
	return s
}

func (s *UpdateFileRequest) SetFileDescription(v string) *UpdateFileRequest {
	s.FileDescription = &v
	return s
}

func (s *UpdateFileRequest) SetFileFolderPath(v string) *UpdateFileRequest {
	s.FileFolderPath = &v
	return s
}

func (s *UpdateFileRequest) SetFileId(v int64) *UpdateFileRequest {
	s.FileId = &v
	return s
}

func (s *UpdateFileRequest) SetFileName(v string) *UpdateFileRequest {
	s.FileName = &v
	return s
}

func (s *UpdateFileRequest) SetIgnoreParentSkipRunningProperty(v bool) *UpdateFileRequest {
	s.IgnoreParentSkipRunningProperty = &v
	return s
}

func (s *UpdateFileRequest) SetImageId(v string) *UpdateFileRequest {
	s.ImageId = &v
	return s
}

func (s *UpdateFileRequest) SetInputList(v string) *UpdateFileRequest {
	s.InputList = &v
	return s
}

func (s *UpdateFileRequest) SetInputParameters(v string) *UpdateFileRequest {
	s.InputParameters = &v
	return s
}

func (s *UpdateFileRequest) SetOutputList(v string) *UpdateFileRequest {
	s.OutputList = &v
	return s
}

func (s *UpdateFileRequest) SetOutputParameters(v string) *UpdateFileRequest {
	s.OutputParameters = &v
	return s
}

func (s *UpdateFileRequest) SetOwner(v string) *UpdateFileRequest {
	s.Owner = &v
	return s
}

func (s *UpdateFileRequest) SetParaValue(v string) *UpdateFileRequest {
	s.ParaValue = &v
	return s
}

func (s *UpdateFileRequest) SetProjectId(v int64) *UpdateFileRequest {
	s.ProjectId = &v
	return s
}

func (s *UpdateFileRequest) SetProjectIdentifier(v string) *UpdateFileRequest {
	s.ProjectIdentifier = &v
	return s
}

func (s *UpdateFileRequest) SetRerunMode(v string) *UpdateFileRequest {
	s.RerunMode = &v
	return s
}

func (s *UpdateFileRequest) SetResourceGroupIdentifier(v string) *UpdateFileRequest {
	s.ResourceGroupIdentifier = &v
	return s
}

func (s *UpdateFileRequest) SetSchedulerType(v string) *UpdateFileRequest {
	s.SchedulerType = &v
	return s
}

func (s *UpdateFileRequest) SetStartEffectDate(v int64) *UpdateFileRequest {
	s.StartEffectDate = &v
	return s
}

func (s *UpdateFileRequest) SetStartImmediately(v bool) *UpdateFileRequest {
	s.StartImmediately = &v
	return s
}

func (s *UpdateFileRequest) SetStop(v bool) *UpdateFileRequest {
	s.Stop = &v
	return s
}

func (s *UpdateFileRequest) SetTimeout(v int32) *UpdateFileRequest {
	s.Timeout = &v
	return s
}

func (s *UpdateFileRequest) Validate() error {
	return dara.Validate(s)
}
