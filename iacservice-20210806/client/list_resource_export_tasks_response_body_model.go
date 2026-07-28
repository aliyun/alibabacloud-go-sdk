// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListResourceExportTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExportTasks(v []*ListResourceExportTasksResponseBodyExportTasks) *ListResourceExportTasksResponseBody
	GetExportTasks() []*ListResourceExportTasksResponseBodyExportTasks
	SetPageNumber(v int32) *ListResourceExportTasksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListResourceExportTasksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListResourceExportTasksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListResourceExportTasksResponseBody
	GetTotalCount() *int32
}

type ListResourceExportTasksResponseBody struct {
	// The list of export tasks.
	ExportTasks []*ListResourceExportTasksResponseBodyExportTasks `json:"exportTasks,omitempty" xml:"exportTasks,omitempty" type:"Repeated"`
	// The current page number.
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
	// 65287CB9-AC46-5FE7-B785-0106C159DA42
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 330
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListResourceExportTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTasksResponseBody) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBody) GetExportTasks() []*ListResourceExportTasksResponseBodyExportTasks {
	return s.ExportTasks
}

func (s *ListResourceExportTasksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListResourceExportTasksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListResourceExportTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListResourceExportTasksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListResourceExportTasksResponseBody) SetExportTasks(v []*ListResourceExportTasksResponseBodyExportTasks) *ListResourceExportTasksResponseBody {
	s.ExportTasks = v
	return s
}

func (s *ListResourceExportTasksResponseBody) SetPageNumber(v int32) *ListResourceExportTasksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListResourceExportTasksResponseBody) SetPageSize(v int32) *ListResourceExportTasksResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListResourceExportTasksResponseBody) SetRequestId(v string) *ListResourceExportTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListResourceExportTasksResponseBody) SetTotalCount(v int32) *ListResourceExportTasksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListResourceExportTasksResponseBody) Validate() error {
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

type ListResourceExportTasksResponseBodyExportTasks struct {
	// The creation time.
	//
	// example:
	//
	// 2025-02-20T02:10:06Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description of the export task.
	//
	// example:
	//
	// this is description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The execution duration.
	//
	// example:
	//
	// 4243
	ElapsedTime *int64 `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	// The export status. Valid values:
	//
	// - Queue: queued
	//
	// - Pending: preparing to run
	//
	// - Success: succeeded
	//
	// - Errored: failed
	//
	// - Canceled: canceled.
	//
	// example:
	//
	// Success
	ExportStatus *string `json:"exportStatus,omitempty" xml:"exportStatus,omitempty"`
	// The ID of the resource export task.
	//
	// example:
	//
	// ex-kw1a1ol8c0pngjav17q8eri
	ExportTaskId *string `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	// The module to which the exported template is saved. If this parameter is not set, the template is automatically saved in the Registry.
	ExportToModule *ListResourceExportTasksResponseBodyExportTasksExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	// The export version.
	//
	// example:
	//
	// v2
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// The values of the include rules for resource export.
	IncludeRules []*ListResourceExportTasksResponseBodyExportTasksIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	// The module configuration of the exported resources.
	Modules []*ListResourceExportTasksResponseBodyExportTasksModules `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	// The name of the export task.
	//
	// example:
	//
	// TaskName
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The task status. Valid values:
	//
	// - Available: available
	//
	// - Running: running.
	//
	// example:
	//
	// Available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The list of variables. The parameters of the exported resources are set as variables.
	Variables []*ListResourceExportTasksResponseBodyExportTasksVariables `json:"variables,omitempty" xml:"variables,omitempty" type:"Repeated"`
}

func (s ListResourceExportTasksResponseBodyExportTasks) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTasksResponseBodyExportTasks) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetDescription() *string {
	return s.Description
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetElapsedTime() *int64 {
	return s.ElapsedTime
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetExportStatus() *string {
	return s.ExportStatus
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetExportTaskId() *string {
	return s.ExportTaskId
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetExportToModule() *ListResourceExportTasksResponseBodyExportTasksExportToModule {
	return s.ExportToModule
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetExportVersion() *string {
	return s.ExportVersion
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetIncludeRules() []*ListResourceExportTasksResponseBodyExportTasksIncludeRules {
	return s.IncludeRules
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetModules() []*ListResourceExportTasksResponseBodyExportTasksModules {
	return s.Modules
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetName() *string {
	return s.Name
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetStatus() *string {
	return s.Status
}

func (s *ListResourceExportTasksResponseBodyExportTasks) GetVariables() []*ListResourceExportTasksResponseBodyExportTasksVariables {
	return s.Variables
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetCreateTime(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.CreateTime = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetDescription(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.Description = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetElapsedTime(v int64) *ListResourceExportTasksResponseBodyExportTasks {
	s.ElapsedTime = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetExportStatus(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.ExportStatus = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetExportTaskId(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.ExportTaskId = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetExportToModule(v *ListResourceExportTasksResponseBodyExportTasksExportToModule) *ListResourceExportTasksResponseBodyExportTasks {
	s.ExportToModule = v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetExportVersion(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.ExportVersion = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetIncludeRules(v []*ListResourceExportTasksResponseBodyExportTasksIncludeRules) *ListResourceExportTasksResponseBodyExportTasks {
	s.IncludeRules = v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetModules(v []*ListResourceExportTasksResponseBodyExportTasksModules) *ListResourceExportTasksResponseBodyExportTasks {
	s.Modules = v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetName(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.Name = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetStatus(v string) *ListResourceExportTasksResponseBodyExportTasks {
	s.Status = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) SetVariables(v []*ListResourceExportTasksResponseBodyExportTasksVariables) *ListResourceExportTasksResponseBodyExportTasks {
	s.Variables = v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasks) Validate() error {
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

type ListResourceExportTasksResponseBodyExportTasksExportToModule struct {
	// The module type to which the exported template is saved. Valid values:
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
	// The path where the template state file is saved.
	//
	// example:
	//
	// /
	StatePath *string `json:"statePath,omitempty" xml:"statePath,omitempty"`
}

func (s ListResourceExportTasksResponseBodyExportTasksExportToModule) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTasksResponseBodyExportTasksExportToModule) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBodyExportTasksExportToModule) GetSource() *string {
	return s.Source
}

func (s *ListResourceExportTasksResponseBodyExportTasksExportToModule) GetSourcePath() *string {
	return s.SourcePath
}

func (s *ListResourceExportTasksResponseBodyExportTasksExportToModule) GetStatePath() *string {
	return s.StatePath
}

func (s *ListResourceExportTasksResponseBodyExportTasksExportToModule) SetSource(v string) *ListResourceExportTasksResponseBodyExportTasksExportToModule {
	s.Source = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksExportToModule) SetSourcePath(v string) *ListResourceExportTasksResponseBodyExportTasksExportToModule {
	s.SourcePath = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksExportToModule) SetStatePath(v string) *ListResourceExportTasksResponseBodyExportTasksExportToModule {
	s.StatePath = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksExportToModule) Validate() error {
	return dara.Validate(s)
}

type ListResourceExportTasksResponseBodyExportTasksIncludeRules struct {
	// The name of the include rule for resource export. Valid values:
	//
	// - ResourceType: required. The resource type, such as ALIYUN::VPC::VPC.
	//
	// - RegionId: required. The region to which the resource belongs. Only one region is supported, such as ap-southeast-1.
	//
	// - \\<ResourceType>:Id: the resource ID, such as ALIYUN::VPC::VPC:Id.
	//
	// - ResourceGroupId: the resource group ID, such as rg-1234.
	//
	// - ZoneId: the zone to which the resource belongs. Only one zone is supported, such as ap-southeast-1h.
	//
	// By default, the relationship between multiple filter conditions is AND. A resource is considered matched only when all filter conditions are met.
	//
	// example:
	//
	// RegionId
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The values of the include rules for resource export.
	Values []*string `json:"values,omitempty" xml:"values,omitempty" type:"Repeated"`
}

func (s ListResourceExportTasksResponseBodyExportTasksIncludeRules) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTasksResponseBodyExportTasksIncludeRules) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBodyExportTasksIncludeRules) GetKey() *string {
	return s.Key
}

func (s *ListResourceExportTasksResponseBodyExportTasksIncludeRules) GetValues() []*string {
	return s.Values
}

func (s *ListResourceExportTasksResponseBodyExportTasksIncludeRules) SetKey(v string) *ListResourceExportTasksResponseBodyExportTasksIncludeRules {
	s.Key = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksIncludeRules) SetValues(v []*string) *ListResourceExportTasksResponseBodyExportTasksIncludeRules {
	s.Values = v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksIncludeRules) Validate() error {
	return dara.Validate(s)
}

type ListResourceExportTasksResponseBodyExportTasksModules struct {
	// The module type where the exported template is stored. Two formats are supported: CloudRegistry and OSS. If the ExportToModule parameter is specified, both formats are returned. Otherwise, only CloudRegistry is returned.
	//
	// example:
	//
	// OSS
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The download URL of the module where the exported template is stored.
	//
	// - If Source is set to CloudRegistry, the format is: "cloudregistry::iacservice/<exportTaskId>/<Provider Name>"
	//
	// - If Source is set to OSS, the format is: "oss::https://<BucketName>.oss-ap-southeast-1.aliyuncs.com/xxx.zip".
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

func (s ListResourceExportTasksResponseBodyExportTasksModules) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTasksResponseBodyExportTasksModules) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBodyExportTasksModules) GetSource() *string {
	return s.Source
}

func (s *ListResourceExportTasksResponseBodyExportTasksModules) GetSourcePath() *string {
	return s.SourcePath
}

func (s *ListResourceExportTasksResponseBodyExportTasksModules) GetVersion() *string {
	return s.Version
}

func (s *ListResourceExportTasksResponseBodyExportTasksModules) SetSource(v string) *ListResourceExportTasksResponseBodyExportTasksModules {
	s.Source = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksModules) SetSourcePath(v string) *ListResourceExportTasksResponseBodyExportTasksModules {
	s.SourcePath = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksModules) SetVersion(v string) *ListResourceExportTasksResponseBodyExportTasksModules {
	s.Version = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksModules) Validate() error {
	return dara.Validate(s)
}

type ListResourceExportTasksResponseBodyExportTasksVariables struct {
	// The list of Terraform resource properties corresponding to the resource type.
	Properties []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
	// The resource type.
	//
	// example:
	//
	// ALIYUN::VPC::VSwitch
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
}

func (s ListResourceExportTasksResponseBodyExportTasksVariables) String() string {
	return dara.Prettify(s)
}

func (s ListResourceExportTasksResponseBodyExportTasksVariables) GoString() string {
	return s.String()
}

func (s *ListResourceExportTasksResponseBodyExportTasksVariables) GetProperties() []*string {
	return s.Properties
}

func (s *ListResourceExportTasksResponseBodyExportTasksVariables) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListResourceExportTasksResponseBodyExportTasksVariables) SetProperties(v []*string) *ListResourceExportTasksResponseBodyExportTasksVariables {
	s.Properties = v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksVariables) SetResourceType(v string) *ListResourceExportTasksResponseBodyExportTasksVariables {
	s.ResourceType = &v
	return s
}

func (s *ListResourceExportTasksResponseBodyExportTasksVariables) Validate() error {
	return dara.Validate(s)
}
