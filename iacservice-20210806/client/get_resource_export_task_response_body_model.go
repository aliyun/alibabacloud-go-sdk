// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceExportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceExportTaskResponseBody
	GetRequestId() *string
	SetTask(v *GetResourceExportTaskResponseBodyTask) *GetResourceExportTaskResponseBody
	GetTask() *GetResourceExportTaskResponseBodyTask
}

type GetResourceExportTaskResponseBody struct {
	// example:
	//
	// FC49AA8C-0A19-5556-8929-E7447F18D529
	RequestId *string                                `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Task      *GetResourceExportTaskResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
}

func (s GetResourceExportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceExportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceExportTaskResponseBody) GetTask() *GetResourceExportTaskResponseBodyTask {
	return s.Task
}

func (s *GetResourceExportTaskResponseBody) SetRequestId(v string) *GetResourceExportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceExportTaskResponseBody) SetTask(v *GetResourceExportTaskResponseBodyTask) *GetResourceExportTaskResponseBody {
	s.Task = v
	return s
}

func (s *GetResourceExportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetResourceExportTaskResponseBodyTask struct {
	// example:
	//
	// 2022-06-15T02:44:37Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// demo
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 4533
	ElapsedTime *int64 `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	// example:
	//
	// ex-al1111jlfh53i6mo4o94jj
	ExportTaskId   *string                                              `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	ExportToModule *GetResourceExportTaskResponseBodyTaskExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	// example:
	//
	// v2
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// example:
	//
	// Reason
	FailedReason *string                                              `json:"failedReason,omitempty" xml:"failedReason,omitempty"`
	IncludeRules []*GetResourceExportTaskResponseBodyTaskIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	Modules      []*GetResourceExportTaskResponseBodyTaskModules      `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	// example:
	//
	// vpc_all
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// role
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// /
	TaskOutputPath *string `json:"taskOutputPath,omitempty" xml:"taskOutputPath,omitempty"`
	// example:
	//
	// {}
	TerraformContext map[string]interface{} `json:"terraformContext,omitempty" xml:"terraformContext,omitempty"`
	// example:
	//
	// 1.246.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// example:
	//
	// 1.2.6
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// example:
	//
	// Manual
	TriggerStrategy *string                                           `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	Variables       []*GetResourceExportTaskResponseBodyTaskVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s GetResourceExportTaskResponseBodyTask) String() string {
	return dara.Prettify(s)
}

func (s GetResourceExportTaskResponseBodyTask) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBodyTask) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetResourceExportTaskResponseBodyTask) GetDescription() *string {
	return s.Description
}

func (s *GetResourceExportTaskResponseBodyTask) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *GetResourceExportTaskResponseBodyTask) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *GetResourceExportTaskResponseBodyTask) GetExportToModule() *GetResourceExportTaskResponseBodyTaskExportToModule {
	return s.ExportToModule
}

func (s *GetResourceExportTaskResponseBodyTask) GetExportVersion() *string {
	return s.ExportVersion
}

func (s *GetResourceExportTaskResponseBodyTask) GetFailedReason() *string {
	return s.FailedReason
}

func (s *GetResourceExportTaskResponseBodyTask) GetIncludeRules() []*GetResourceExportTaskResponseBodyTaskIncludeRules {
	return s.IncludeRules
}

func (s *GetResourceExportTaskResponseBodyTask) GetModules() []*GetResourceExportTaskResponseBodyTaskModules {
	return s.Modules
}

func (s *GetResourceExportTaskResponseBodyTask) GetName() *string {
	return s.Name
}

func (s *GetResourceExportTaskResponseBodyTask) GetRamRole() *string {
	return s.RamRole
}

func (s *GetResourceExportTaskResponseBodyTask) GetStatus() *string {
	return s.Status
}

func (s *GetResourceExportTaskResponseBodyTask) GetTaskOutputPath() *string {
	return s.TaskOutputPath
}

func (s *GetResourceExportTaskResponseBodyTask) GetTerraformContext() map[string]interface{} {
	return s.TerraformContext
}

func (s *GetResourceExportTaskResponseBodyTask) GetTerraformProviderVersion() *string {
	return s.TerraformProviderVersion
}

func (s *GetResourceExportTaskResponseBodyTask) GetTerraformVersion() *string {
	return s.TerraformVersion
}

func (s *GetResourceExportTaskResponseBodyTask) GetTriggerStrategy() *string {
	return s.TriggerStrategy
}

func (s *GetResourceExportTaskResponseBodyTask) GetVariables() []*GetResourceExportTaskResponseBodyTaskVariables {
	return s.Variables
}

func (s *GetResourceExportTaskResponseBodyTask) SetCreateTime(v string) *GetResourceExportTaskResponseBodyTask {
	s.CreateTime = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetDescription(v string) *GetResourceExportTaskResponseBodyTask {
	s.Description = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetElapsedTime(v int64) *GetResourceExportTaskResponseBodyTask {
	s.ElapsedTime = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetExportTaskId(v string) *GetResourceExportTaskResponseBodyTask {
	s.ExportTaskId = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetExportToModule(v *GetResourceExportTaskResponseBodyTaskExportToModule) *GetResourceExportTaskResponseBodyTask {
	s.ExportToModule = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetExportVersion(v string) *GetResourceExportTaskResponseBodyTask {
	s.ExportVersion = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetFailedReason(v string) *GetResourceExportTaskResponseBodyTask {
	s.FailedReason = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetIncludeRules(v []*GetResourceExportTaskResponseBodyTaskIncludeRules) *GetResourceExportTaskResponseBodyTask {
	s.IncludeRules = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetModules(v []*GetResourceExportTaskResponseBodyTaskModules) *GetResourceExportTaskResponseBodyTask {
	s.Modules = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetName(v string) *GetResourceExportTaskResponseBodyTask {
	s.Name = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetRamRole(v string) *GetResourceExportTaskResponseBodyTask {
	s.RamRole = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetStatus(v string) *GetResourceExportTaskResponseBodyTask {
	s.Status = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetTaskOutputPath(v string) *GetResourceExportTaskResponseBodyTask {
	s.TaskOutputPath = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetTerraformContext(v map[string]interface{}) *GetResourceExportTaskResponseBodyTask {
	s.TerraformContext = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetTerraformProviderVersion(v string) *GetResourceExportTaskResponseBodyTask {
	s.TerraformProviderVersion = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetTerraformVersion(v string) *GetResourceExportTaskResponseBodyTask {
	s.TerraformVersion = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetTriggerStrategy(v string) *GetResourceExportTaskResponseBodyTask {
	s.TriggerStrategy = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) SetVariables(v []*GetResourceExportTaskResponseBodyTaskVariables) *GetResourceExportTaskResponseBodyTask {
	s.Variables = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTask) Validate() error {
	return dara.Validate(s)
}

type GetResourceExportTaskResponseBodyTaskExportToModule struct {
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// alibaba/security-group/alicloud
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// example:
	//
	// /
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
}

func (s GetResourceExportTaskResponseBodyTaskExportToModule) String() string {
	return dara.Prettify(s)
}

func (s GetResourceExportTaskResponseBodyTaskExportToModule) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBodyTaskExportToModule) GetSource() *string {
	return s.Source
}

func (s *GetResourceExportTaskResponseBodyTaskExportToModule) GetSourcePath() *string {
	return s.SourcePath
}

func (s *GetResourceExportTaskResponseBodyTaskExportToModule) GetStatePath() *string {
	return s.StatePath
}

func (s *GetResourceExportTaskResponseBodyTaskExportToModule) SetSource(v string) *GetResourceExportTaskResponseBodyTaskExportToModule {
	s.Source = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskExportToModule) SetSourcePath(v string) *GetResourceExportTaskResponseBodyTaskExportToModule {
	s.SourcePath = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskExportToModule) SetStatePath(v string) *GetResourceExportTaskResponseBodyTaskExportToModule {
	s.StatePath = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskExportToModule) Validate() error {
	return dara.Validate(s)
}

type GetResourceExportTaskResponseBodyTaskIncludeRules struct {
	// example:
	//
	// ZoneId
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s GetResourceExportTaskResponseBodyTaskIncludeRules) String() string {
	return dara.Prettify(s)
}

func (s GetResourceExportTaskResponseBodyTaskIncludeRules) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBodyTaskIncludeRules) GetKey() *string {
	return s.Key
}

func (s *GetResourceExportTaskResponseBodyTaskIncludeRules) GetValues() []*string {
	return s.Values
}

func (s *GetResourceExportTaskResponseBodyTaskIncludeRules) SetKey(v string) *GetResourceExportTaskResponseBodyTaskIncludeRules {
	s.Key = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskIncludeRules) SetValues(v []*string) *GetResourceExportTaskResponseBodyTaskIncludeRules {
	s.Values = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskIncludeRules) Validate() error {
	return dara.Validate(s)
}

type GetResourceExportTaskResponseBodyTaskModules struct {
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// example:
	//
	// oss::https://.oss-cn-hangzhou.aliyuncs.com/xxx.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// example:
	//
	// v3
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetResourceExportTaskResponseBodyTaskModules) String() string {
	return dara.Prettify(s)
}

func (s GetResourceExportTaskResponseBodyTaskModules) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBodyTaskModules) GetSource() *string {
	return s.Source
}

func (s *GetResourceExportTaskResponseBodyTaskModules) GetSourcePath() *string {
	return s.SourcePath
}

func (s *GetResourceExportTaskResponseBodyTaskModules) GetVersion() *string {
	return s.Version
}

func (s *GetResourceExportTaskResponseBodyTaskModules) SetSource(v string) *GetResourceExportTaskResponseBodyTaskModules {
	s.Source = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskModules) SetSourcePath(v string) *GetResourceExportTaskResponseBodyTaskModules {
	s.SourcePath = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskModules) SetVersion(v string) *GetResourceExportTaskResponseBodyTaskModules {
	s.Version = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskModules) Validate() error {
	return dara.Validate(s)
}

type GetResourceExportTaskResponseBodyTaskVariables struct {
	Properties []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	// example:
	//
	// ALIYUN::Bastionhost::Instance
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s GetResourceExportTaskResponseBodyTaskVariables) String() string {
	return dara.Prettify(s)
}

func (s GetResourceExportTaskResponseBodyTaskVariables) GoString() string {
	return s.String()
}

func (s *GetResourceExportTaskResponseBodyTaskVariables) GetProperties() []*string {
	return s.Properties
}

func (s *GetResourceExportTaskResponseBodyTaskVariables) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetResourceExportTaskResponseBodyTaskVariables) SetProperties(v []*string) *GetResourceExportTaskResponseBodyTaskVariables {
	s.Properties = v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskVariables) SetResourceType(v string) *GetResourceExportTaskResponseBodyTaskVariables {
	s.ResourceType = &v
	return s
}

func (s *GetResourceExportTaskResponseBodyTaskVariables) Validate() error {
	return dara.Validate(s)
}
