// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkillAddMetaProfilingJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentId(v int64) *SkillAddMetaProfilingJobRequest
	GetComponentId() *int64
	SetJobDesc(v string) *SkillAddMetaProfilingJobRequest
	GetJobDesc() *string
	SetJobId(v int64) *SkillAddMetaProfilingJobRequest
	GetJobId() *int64
	SetJobName(v string) *SkillAddMetaProfilingJobRequest
	GetJobName() *string
	SetMetaIds(v []*string) *SkillAddMetaProfilingJobRequest
	GetMetaIds() []*string
	SetPluginId(v string) *SkillAddMetaProfilingJobRequest
	GetPluginId() *string
	SetProfilingFlag(v int32) *SkillAddMetaProfilingJobRequest
	GetProfilingFlag() *int32
	SetProfilingPermission(v int32) *SkillAddMetaProfilingJobRequest
	GetProfilingPermission() *int32
	SetProfilingRule(v string) *SkillAddMetaProfilingJobRequest
	GetProfilingRule() *string
	SetProfilingSchemaList(v []*string) *SkillAddMetaProfilingJobRequest
	GetProfilingSchemaList() []*string
	SetProfilingType(v int32) *SkillAddMetaProfilingJobRequest
	GetProfilingType() *int32
}

type SkillAddMetaProfilingJobRequest struct {
	// The component ID, which is the primary key identifier of the data source component.
	//
	// example:
	//
	// 2001
	ComponentId *int64 `json:"componentId,omitempty" xml:"componentId,omitempty"`
	// The description of the profiling task.
	//
	// example:
	//
	// Daily full metadata profiling task
	JobDesc *string `json:"jobDesc,omitempty" xml:"jobDesc,omitempty"`
	// The profiling task ID.
	//
	// example:
	//
	// 100256
	JobId *int64 `json:"jobId,omitempty" xml:"jobId,omitempty"`
	// The name of the profiling task.
	//
	// example:
	//
	// profiling_daily
	JobName *string `json:"jobName,omitempty" xml:"jobName,omitempty"`
	// The list of metadata object IDs to be profiled or operated on.
	MetaIds []*string `json:"metaIds,omitempty" xml:"metaIds,omitempty" type:"Repeated"`
	// The profiling plugin ID, which specifies the plugin used to perform the profiling.
	//
	// example:
	//
	// profiling-plugin-001
	PluginId *string `json:"pluginId,omitempty" xml:"pluginId,omitempty"`
	// Specifies whether to enable data profiling.
	//
	// example:
	//
	// 1
	ProfilingFlag *int32 `json:"profilingFlag,omitempty" xml:"profilingFlag,omitempty"`
	// The profiling permission configuration identifier.
	//
	// example:
	//
	// 1
	ProfilingPermission *int32 `json:"profilingPermission,omitempty" xml:"profilingPermission,omitempty"`
	// The profiling rule configuration content.
	//
	// example:
	//
	// basic
	ProfilingRule *string `json:"profilingRule,omitempty" xml:"profilingRule,omitempty"`
	// The list of schemas to be profiled.
	ProfilingSchemaList []*string `json:"profilingSchemaList,omitempty" xml:"profilingSchemaList,omitempty" type:"Repeated"`
	// The profiling type identifier.
	//
	// example:
	//
	// 1
	ProfilingType *int32 `json:"profilingType,omitempty" xml:"profilingType,omitempty"`
}

func (s SkillAddMetaProfilingJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SkillAddMetaProfilingJobRequest) GoString() string {
	return s.String()
}

func (s *SkillAddMetaProfilingJobRequest) GetComponentId() *int64 {
	return s.ComponentId
}

func (s *SkillAddMetaProfilingJobRequest) GetJobDesc() *string {
	return s.JobDesc
}

func (s *SkillAddMetaProfilingJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *SkillAddMetaProfilingJobRequest) GetJobName() *string {
	return s.JobName
}

func (s *SkillAddMetaProfilingJobRequest) GetMetaIds() []*string {
	return s.MetaIds
}

func (s *SkillAddMetaProfilingJobRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *SkillAddMetaProfilingJobRequest) GetProfilingFlag() *int32 {
	return s.ProfilingFlag
}

func (s *SkillAddMetaProfilingJobRequest) GetProfilingPermission() *int32 {
	return s.ProfilingPermission
}

func (s *SkillAddMetaProfilingJobRequest) GetProfilingRule() *string {
	return s.ProfilingRule
}

func (s *SkillAddMetaProfilingJobRequest) GetProfilingSchemaList() []*string {
	return s.ProfilingSchemaList
}

func (s *SkillAddMetaProfilingJobRequest) GetProfilingType() *int32 {
	return s.ProfilingType
}

func (s *SkillAddMetaProfilingJobRequest) SetComponentId(v int64) *SkillAddMetaProfilingJobRequest {
	s.ComponentId = &v
	return s
}

func (s *SkillAddMetaProfilingJobRequest) SetJobDesc(v string) *SkillAddMetaProfilingJobRequest {
	s.JobDesc = &v
	return s
}

func (s *SkillAddMetaProfilingJobRequest) SetJobId(v int64) *SkillAddMetaProfilingJobRequest {
	s.JobId = &v
	return s
}

func (s *SkillAddMetaProfilingJobRequest) SetJobName(v string) *SkillAddMetaProfilingJobRequest {
	s.JobName = &v
	return s
}

func (s *SkillAddMetaProfilingJobRequest) SetMetaIds(v []*string) *SkillAddMetaProfilingJobRequest {
	s.MetaIds = v
	return s
}

func (s *SkillAddMetaProfilingJobRequest) SetPluginId(v string) *SkillAddMetaProfilingJobRequest {
	s.PluginId = &v
	return s
}

func (s *SkillAddMetaProfilingJobRequest) SetProfilingFlag(v int32) *SkillAddMetaProfilingJobRequest {
	s.ProfilingFlag = &v
	return s
}

func (s *SkillAddMetaProfilingJobRequest) SetProfilingPermission(v int32) *SkillAddMetaProfilingJobRequest {
	s.ProfilingPermission = &v
	return s
}

func (s *SkillAddMetaProfilingJobRequest) SetProfilingRule(v string) *SkillAddMetaProfilingJobRequest {
	s.ProfilingRule = &v
	return s
}

func (s *SkillAddMetaProfilingJobRequest) SetProfilingSchemaList(v []*string) *SkillAddMetaProfilingJobRequest {
	s.ProfilingSchemaList = v
	return s
}

func (s *SkillAddMetaProfilingJobRequest) SetProfilingType(v int32) *SkillAddMetaProfilingJobRequest {
	s.ProfilingType = &v
	return s
}

func (s *SkillAddMetaProfilingJobRequest) Validate() error {
	return dara.Validate(s)
}
