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
	ExportTasks []*ListResourceExportTaskVersionsResponseBodyExportTasks `json:"exportTasks,omitempty" xml:"exportTasks,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 860FDEEE-1CA3-55F3-97F6-63FC40B7962D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	// example:
	//
	// 2025-05-11T02:18:50Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// demo
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// 4521
	ElapsedTime *int64 `json:"elapsedTime,omitempty" xml:"elapsedTime,omitempty"`
	// example:
	//
	// ex-al1711jl6hd8u5crggeq6v
	ExportTaskId   *string                                                              `json:"exportTaskId,omitempty" xml:"exportTaskId,omitempty"`
	ExportToModule *ListResourceExportTaskVersionsResponseBodyExportTasksExportToModule `json:"exportToModule,omitempty" xml:"exportToModule,omitempty" type:"Struct"`
	// example:
	//
	// v3
	ExportVersion *string `json:"exportVersion,omitempty" xml:"exportVersion,omitempty"`
	// example:
	//
	// Reason
	FailedReason *string                                                              `json:"failedReason,omitempty" xml:"failedReason,omitempty"`
	IncludeRules []*ListResourceExportTaskVersionsResponseBodyExportTasksIncludeRules `json:"includeRules,omitempty" xml:"includeRules,omitempty" type:"Repeated"`
	Modules      []*ListResourceExportTaskVersionsResponseBodyExportTasksModules      `json:"modules,omitempty" xml:"modules,omitempty" type:"Repeated"`
	// example:
	//
	// vpc_all
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// Running
	Status    *string                                                           `json:"status,omitempty" xml:"status,omitempty"`
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
	// example:
	//
	// ZoneId
	Key    *string   `json:"key,omitempty" xml:"key,omitempty"`
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
	Properties []*string `json:"properties,omitempty" xml:"properties,omitempty" type:"Repeated"`
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
