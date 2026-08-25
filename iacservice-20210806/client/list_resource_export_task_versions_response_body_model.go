// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceExportTaskVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExportTasks(v []*ListResourceExportTaskVersionsResponseBodyExportTasks) *ListResourceExportTaskVersionsResponseBody
	GetExportTasks() []*ListResourceExportTaskVersionsResponseBodyExportTasks
	SetPageNumber(v int32) *ListResourceExportTaskVersionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceExportTaskVersionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListResourceExportTaskVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListResourceExportTaskVersionsResponseBody
	GetTotalCount() *int32
}

type ListResourceExportTaskVersionsResponseBody struct {
	// The list of export task versions.
	ExportTasks []*ListResourceExportTaskVersionsResponseBodyExportTasks `json:"exportTasks,omitempty" xml:"exportTasks,omitempty" type:"Repeated"`
	// The page number. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of results per page. Default value: 20. Minimum value: 1. Maximum value: 100.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 860FDEEE-1CA3-55F3-97F6-63FC40B7962D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 72
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourceExportTaskVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBody) GetExportTasks() []*ListResourceExportTaskVersionsResponseBodyExportTasks {
	return s.ExportTasks
}

func (s *ListResourceExportTaskVersionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceExportTaskVersionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceExportTaskVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceExportTaskVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceExportTaskVersionsResponseBody) SetExportTasks(v []*ListResourceExportTaskVersionsResponseBodyExportTasks) *ListResourceExportTaskVersionsResponseBody {
	s.ExportTasks = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBody) SetPageNumber(v int32) *ListResourceExportTaskVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBody) SetPageSize(v int32) *ListResourceExportTaskVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBody) SetRequestId(v string) *ListResourceExportTaskVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBody) SetTotalCount(v int32) *ListResourceExportTaskVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBody) Validate() error {
	if s.ExportTasks != nil {
		for _, item := range s.ExportTasks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListResourceExportTaskVersionsResponseBodyExportTasks struct {
	// The creation time in UTC, in the ISO 8601 format of YYYY-MM-DDTHH:mm:ssZ.
	//
	// example:
	//
	// 2025-05-11T02:18:50Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description.
	//
	// example:
	//
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution duration.
	//
	// example:
	//
	// 4521
	ElapsedTime *int64 `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	// The ID of the resource export task.
	//
	// example:
	//
	// ex-al1711jl6hd8u5crggeq6v
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// The module to which the exported template is saved. If this parameter is not set, the template is automatically saved in the Registry.
	ExportToModule *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	// The resource export version.
	//
	// example:
	//
	// v3
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// The reason for the export failure.
	//
	// example:
	//
	// Reason
	FailedReason *string `json:"failedReason,omitempty" xml:"failedReason,omitempty"`
	// The list of include rules used when exporting resources.
	IncludeRules  []*ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	ManagedTaskId *string                                                              `json:"managedTaskId,omitempty" xml:"managedTaskId,omitempty"`
	// The module configuration of the exported resources.
	Modules []*ListResourceExportTaskVersionsResponseBodyExportTasksModules `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	// The name of the export task.
	//
	// example:
	//
	// TaskName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The version export status. Valid values:
	//
	// - Queue: queued
	//
	// - Pending: preparing to run
	//
	// - Success: succeeded
	//
	// - Errored: failed
	//
	// - Canceled: canceled
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of variables. Parameters of exported resources are set as variables.
	Variables []*ListResourceExportTaskVersionsResponseBodyExportTasksVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasks) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasks) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetDescription() *string {
	return s.Description
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetExportToModule() *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule {
	return s.ExportToModule
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetExportVersion() *string {
	return s.ExportVersion
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetFailedReason() *string {
	return s.FailedReason
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetIncludeRules() []*ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules {
	return s.IncludeRules
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetManagedTaskId() *string {
	return s.ManagedTaskId
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetModules() []*ListResourceExportTaskVersionsResponseBodyExportTasksModules {
	return s.Modules
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetName() *string {
	return s.Name
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetStatus() *string {
	return s.Status
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) GetVariables() []*ListResourceExportTaskVersionsResponseBodyExportTasksVariables {
	return s.Variables
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetCreateTime(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.CreateTime = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetDescription(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.Description = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetElapsedTime(v int64) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.ElapsedTime = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetExportTaskId(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.ExportTaskId = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetExportToModule(v *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.ExportToModule = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetExportVersion(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.ExportVersion = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetFailedReason(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.FailedReason = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetIncludeRules(v []*ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.IncludeRules = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetManagedTaskId(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.ManagedTaskId = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetModules(v []*ListResourceExportTaskVersionsResponseBodyExportTasksModules) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.Modules = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetName(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.Name = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetStatus(v string) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.Status = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) SetVariables(v []*ListResourceExportTaskVersionsResponseBodyExportTasksVariables) *ListResourceExportTaskVersionsResponseBodyExportTasks {
	s.Variables = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasks) Validate() error {
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

type ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule struct {
	// The module type to which the exported template is saved. Valid values:
	//
	// - OSS: OSS
	//
	// - Registry: Terraform Registry
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The path where the template content is saved.
	//
	// - If Source is set to Registry, the format is: "cloudregistry::iacservice//"
	//
	// - If Source is set to OSS, the format is: "oss::https://.oss-cn-hangzhou.aliyuncs.com/xxx.zip"
	//
	// example:
	//
	// oss::https://.oss-cn-hangzhou.aliyuncs.com/xxx.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The path of the State file corresponding to the module.
	//
	// example:
	//
	// /
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) GetSource() *string {
	return s.Source
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) GetSourcePath() *string {
	return s.SourcePath
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) GetStatePath() *string {
	return s.StatePath
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) SetSource(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule {
	s.Source = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) SetSourcePath(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule {
	s.SourcePath = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) SetStatePath(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule {
	s.StatePath = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule) Validate() error {
	return dara.Validate(s)
}

type ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules struct {
	// The name of the include rule for resource export. Valid values:
	//
	// - ResourceType: required. The resource type. Example: ALIYUN::VPC::VPC.
	//
	// - RegionId: required. The region to which the resource belongs. Only one region is supported. Example: cn-chengdu.
	//
	// - \\<ResourceType>:Id: the resource ID. Example: ALIYUN::VPC::VPC:Id.
	//
	// - ResourceGroupId: the resource group ID. Example: rg-1234.
	//
	// - ZoneId: the zone to which the resource belongs. Only one zone is supported. Example: cn-hangzhou-h.
	//
	// Multiple filter conditions have an AND relationship by default. A resource must meet all filter conditions to be considered a match.
	//
	// example:
	//
	// RegionId
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The values of the include rule for resource export.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) GetKey() *string {
	return s.Key
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) GetValues() []*string {
	return s.Values
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) SetKey(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules {
	s.Key = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) SetValues(v []*string) *ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules {
	s.Values = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules) Validate() error {
	return dara.Validate(s)
}

type ListResourceExportTaskVersionsResponseBodyExportTasksModules struct {
	// The module type where the exported template is located. Two formats are supported: CloudRegistry and OSS. If the ExportToModule parameter is specified, both formats are returned. Otherwise, only CloudRegistry is returned.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The download address of the exported template within the module.
	//
	// - If Source is CloudRegistry, the format is: "cloudregistry::iacservice//"
	//
	// - If Source is OSS, the format is: "oss::https://.oss-cn-hangzhou.aliyuncs.com/xxx.zip"
	//
	// example:
	//
	// oss::https://.oss-cn-hangzhou.aliyuncs.com/xxx.zip
	SourcePath *string `json:"sourcePath,omitempty" xml:"sourcePath,omitempty"`
	// The version of the module where the exported template is located.
	//
	// example:
	//
	// v1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksModules) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksModules) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksModules) GetSource() *string {
	return s.Source
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksModules) GetSourcePath() *string {
	return s.SourcePath
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksModules) GetVersion() *string {
	return s.Version
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksModules) SetSource(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksModules {
	s.Source = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksModules) SetSourcePath(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksModules {
	s.SourcePath = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksModules) SetVersion(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksModules {
	s.Version = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksModules) Validate() error {
	return dara.Validate(s)
}

type ListResourceExportTaskVersionsResponseBodyExportTasksVariables struct {
	// The list of Terraform resource properties corresponding to the resource type.
	Properties []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// Vswitch
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksVariables) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTaskVersionsResponseBodyExportTasksVariables) GoString() string {
	return s.String()
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksVariables) GetProperties() []*string {
	return s.Properties
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksVariables) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksVariables) SetProperties(v []*string) *ListResourceExportTaskVersionsResponseBodyExportTasksVariables {
	s.Properties = v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksVariables) SetResourceType(v string) *ListResourceExportTaskVersionsResponseBodyExportTasksVariables {
	s.ResourceType = &v
	return s
}

func (s *ListResourceExportTaskVersionsResponseBodyExportTasksVariables) Validate() error {
	return dara.Validate(s)
}
