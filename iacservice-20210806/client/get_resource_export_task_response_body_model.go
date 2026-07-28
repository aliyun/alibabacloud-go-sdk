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
	// The request ID.
	//
	// example:
	//
	// FC49AA8C-0A19-5556-8929-E7447F18D529
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The task information.
	Task *GetResourceExportTaskResponseBodyTask `json:"task,omitempty" xml:"task,omitempty" type:"Struct"`
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
	if s.Task != nil {
		if err := s.Task.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceExportTaskResponseBodyTask struct {
	// The time when the task was created.
	//
	// example:
	//
	// 2022-06-15T02:44:37Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The task description.
	//
	// example:
	//
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution duration.
	//
	// example:
	//
	// 4533
	ElapsedTime *int64 `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	// The ID of the resource export task.
	//
	// example:
	//
	// ex-al1111jlfh53i6mo4o94jj
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// Saves the exported template as a module. If this parameter is not set, the template is automatically saved in the registry.
	ExportToModule *GetResourceExportTaskResponseBodyTaskExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	// The resource export version.
	//
	// example:
	//
	// v2
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// The failure reason.
	//
	// example:
	//
	// Reason
	FailedReason *string `json:"failedReason,omitempty" xml:"failedReason,omitempty"`
	// The values of the include rules for resource export.
	IncludeRules []*GetResourceExportTaskResponseBodyTaskIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	// The module configuration for the exported resources.
	Modules []*GetResourceExportTaskResponseBodyTaskModules `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	// The task name.
	//
	// example:
	//
	// vpc_all
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The RAM role.
	//
	// example:
	//
	// role
	RamRole *string `json:"ramRole,omitempty" xml:"ramRole,omitempty"`
	// The task status. Valid values:
	//
	// - Available: the task is available and no job is running.
	//
	// - Running: a job is currently running.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The task output path.
	//
	// example:
	//
	// /
	TaskOutputPath *string `json:"taskOutputPath,omitempty" xml:"taskOutputPath,omitempty"`
	// The Terraform context.
	//
	// example:
	//
	// {}
	TerraformContext map[string]interface{} `json:"terraformContext,omitempty" xml:"terraformContext,omitempty"`
	// The Terraform provider version.
	//
	// example:
	//
	// 1.246.0
	TerraformProviderVersion *string `json:"terraformProviderVersion,omitempty" xml:"terraformProviderVersion,omitempty"`
	// The Terraform version.
	//
	// example:
	//
	// 1.5.7
	TerraformVersion *string `json:"terraformVersion,omitempty" xml:"terraformVersion,omitempty"`
	// The trigger strategy. Valid values:
	//
	// - Auto: triggered automatically when rules are modified or the trigger strategy is changed to Auto.
	//
	// - Manual: triggered manually.
	//
	// Default value: Manual.
	//
	// example:
	//
	// Manual
	TriggerStrategy *string `json:"triggerStrategy,omitempty" xml:"triggerStrategy,omitempty"`
	// The list of variables. Parameters in the exported resources are set as variables.
	Variables []*GetResourceExportTaskResponseBodyTaskVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
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
	if s.ExportToModule != nil {
		if err := s.ExportToModule.Validate(); err != nil {
			return err
		}
	}
	if s.IncludeRules != nil {
		for _, item := range s.IncludeRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Modules != nil {
		for _, item := range s.Modules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Variables != nil {
		for _, item := range s.Variables {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetResourceExportTaskResponseBodyTaskExportToModule struct {
	// The module type in which the exported template is saved. Valid values:
	//
	// - OSS: OSS
	//
	// - Registry: Terraform Registry.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The path where the template content is saved.
	//
	// - If Source is set to Registry, the format is: "cloudregistry::iacservice//"
	//
	// - If Source is set to OSS, the format is: "oss::https://.oss-ap-southeast-1.aliyuncs.com/xxx.zip".
	//
	// example:
	//
	// oss::https://.oss-cn-hangzhou.aliyuncs.com/xxx.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The path of the state file that corresponds to the module.
	//
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
	// The name of the include rule for resource export. Valid values:
	//
	// - ResourceType: required. The resource type. Example: ALIYUN::VPC::VPC.
	//
	// - RegionId: required. The region to which the resource belongs. Only one region is supported. Example: ap-southeast-1.
	//
	// - \\<ResourceType>:Id: the resource ID. Example: ALIYUN::VPC::VPC:Id.
	//
	// - ResourceGroupId: the resource group ID. Example: rg-1234.
	//
	// - ZoneId: the zone to which the resource belongs. Only one zone is supported. Example: ap-southeast-1a.
	//
	// By default, the relationship between multiple filter conditions is AND. A resource is considered matched only if all filter conditions are met.
	//
	// example:
	//
	// RegionId
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The values of the include rules for resource export.
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
	// The module type where the exported template is stored. Two formats are supported: CloudRegistry and OSS. If the ExportToModule parameter is specified, both formats are returned. Otherwise, only CloudRegistry is returned.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The download URL of the module where the exported template is stored.
	//
	// - If Source is set to CloudRegistry, the format is: "cloudregistry::iacservice//"
	//
	// - If Source is set to OSS, the format is: "oss::https://.oss-ap-southeast-1.aliyuncs.com/xxx.zip".
	//
	// example:
	//
	// oss::https://.oss-cn-hangzhou.aliyuncs.com/xxx.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The version of the module where the exported template is stored.
	//
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
	// The list of properties of the Terraform resource that corresponds to the resource type.
	Properties []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::VPC::VSwitch
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
