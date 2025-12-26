// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFlowRun interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *FlowRun
	GetAccessibility() *string
	SetChildRuns(v []*FlowRunChildRuns) *FlowRun
	GetChildRuns() []*FlowRunChildRuns
	SetCreator(v string) *FlowRun
	GetCreator() *string
	SetCredentialConfig(v *FlowRunCredentialConfig) *FlowRun
	GetCredentialConfig() *FlowRunCredentialConfig
	SetDataColumnMapping(v map[string]*string) *FlowRun
	GetDataColumnMapping() map[string]*string
	SetDataSources(v []*FlowRunDataSources) *FlowRun
	GetDataSources() []*FlowRunDataSources
	SetDuration(v int32) *FlowRun
	GetDuration() *int32
	SetEcsSpec(v *FlowRunEcsSpec) *FlowRun
	GetEcsSpec() *FlowRunEcsSpec
	SetEnvs(v []*FlowRunEnvs) *FlowRun
	GetEnvs() []*FlowRunEnvs
	SetEvaluationConfigs(v []*FlowRunEvaluationConfigs) *FlowRun
	GetEvaluationConfigs() []*FlowRunEvaluationConfigs
	SetEvaluationWorkerCount(v int32) *FlowRun
	GetEvaluationWorkerCount() *int32
	SetException(v string) *FlowRun
	GetException() *string
	SetFlowId(v string) *FlowRun
	GetFlowId() *string
	SetFlowName(v string) *FlowRun
	GetFlowName() *string
	SetFlowRunId(v string) *FlowRun
	GetFlowRunId() *string
	SetGmtCreateTime(v string) *FlowRun
	GetGmtCreateTime() *string
	SetGmtFinishTime(v string) *FlowRun
	GetGmtFinishTime() *string
	SetGmtModifiedTime(v string) *FlowRun
	GetGmtModifiedTime() *string
	SetGmtStartTime(v string) *FlowRun
	GetGmtStartTime() *string
	SetLabels(v []*FlowRunLabels) *FlowRun
	GetLabels() []*FlowRunLabels
	SetNodeName(v string) *FlowRun
	GetNodeName() *string
	SetNodeRunInfos(v []*FlowRunNodeRunInfos) *FlowRun
	GetNodeRunInfos() []*FlowRunNodeRunInfos
	SetResourceId(v string) *FlowRun
	GetResourceId() *string
	SetRunMode(v string) *FlowRun
	GetRunMode() *string
	SetRunName(v string) *FlowRun
	GetRunName() *string
	SetRunResult(v string) *FlowRun
	GetRunResult() *string
	SetRunStatus(v string) *FlowRun
	GetRunStatus() *string
	SetRunTimeout(v int32) *FlowRun
	GetRunTimeout() *int32
	SetUserVpc(v *FlowRunUserVpc) *FlowRun
	GetUserVpc() *FlowRunUserVpc
	SetVariant(v string) *FlowRun
	GetVariant() *string
	SetWorkspaceId(v string) *FlowRun
	GetWorkspaceId() *string
}

type FlowRun struct {
	Accessibility         *string                     `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	ChildRuns             []*FlowRunChildRuns         `json:"ChildRuns,omitempty" xml:"ChildRuns,omitempty" type:"Repeated"`
	Creator               *string                     `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CredentialConfig      *FlowRunCredentialConfig    `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty" type:"Struct"`
	DataColumnMapping     map[string]*string          `json:"DataColumnMapping,omitempty" xml:"DataColumnMapping,omitempty"`
	DataSources           []*FlowRunDataSources       `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	Duration              *int32                      `json:"Duration,omitempty" xml:"Duration,omitempty"`
	EcsSpec               *FlowRunEcsSpec             `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty" type:"Struct"`
	Envs                  []*FlowRunEnvs              `json:"Envs,omitempty" xml:"Envs,omitempty" type:"Repeated"`
	EvaluationConfigs     []*FlowRunEvaluationConfigs `json:"EvaluationConfigs,omitempty" xml:"EvaluationConfigs,omitempty" type:"Repeated"`
	EvaluationWorkerCount *int32                      `json:"EvaluationWorkerCount,omitempty" xml:"EvaluationWorkerCount,omitempty"`
	Exception             *string                     `json:"Exception,omitempty" xml:"Exception,omitempty"`
	FlowId                *string                     `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	FlowName              *string                     `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
	FlowRunId             *string                     `json:"FlowRunId,omitempty" xml:"FlowRunId,omitempty"`
	GmtCreateTime         *string                     `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtFinishTime         *string                     `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	GmtModifiedTime       *string                     `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	GmtStartTime          *string                     `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	Labels                []*FlowRunLabels            `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	NodeName              *string                     `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	NodeRunInfos          []*FlowRunNodeRunInfos      `json:"NodeRunInfos,omitempty" xml:"NodeRunInfos,omitempty" type:"Repeated"`
	ResourceId            *string                     `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	RunMode               *string                     `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	RunName               *string                     `json:"RunName,omitempty" xml:"RunName,omitempty"`
	RunResult             *string                     `json:"RunResult,omitempty" xml:"RunResult,omitempty"`
	RunStatus             *string                     `json:"RunStatus,omitempty" xml:"RunStatus,omitempty"`
	RunTimeout            *int32                      `json:"RunTimeout,omitempty" xml:"RunTimeout,omitempty"`
	UserVpc               *FlowRunUserVpc             `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	Variant               *string                     `json:"Variant,omitempty" xml:"Variant,omitempty"`
	WorkspaceId           *string                     `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s FlowRun) String() string {
	return dara.Prettify(s)
}

func (s FlowRun) GoString() string {
	return s.String()
}

func (s *FlowRun) GetAccessibility() *string {
	return s.Accessibility
}

func (s *FlowRun) GetChildRuns() []*FlowRunChildRuns {
	return s.ChildRuns
}

func (s *FlowRun) GetCreator() *string {
	return s.Creator
}

func (s *FlowRun) GetCredentialConfig() *FlowRunCredentialConfig {
	return s.CredentialConfig
}

func (s *FlowRun) GetDataColumnMapping() map[string]*string {
	return s.DataColumnMapping
}

func (s *FlowRun) GetDataSources() []*FlowRunDataSources {
	return s.DataSources
}

func (s *FlowRun) GetDuration() *int32 {
	return s.Duration
}

func (s *FlowRun) GetEcsSpec() *FlowRunEcsSpec {
	return s.EcsSpec
}

func (s *FlowRun) GetEnvs() []*FlowRunEnvs {
	return s.Envs
}

func (s *FlowRun) GetEvaluationConfigs() []*FlowRunEvaluationConfigs {
	return s.EvaluationConfigs
}

func (s *FlowRun) GetEvaluationWorkerCount() *int32 {
	return s.EvaluationWorkerCount
}

func (s *FlowRun) GetException() *string {
	return s.Exception
}

func (s *FlowRun) GetFlowId() *string {
	return s.FlowId
}

func (s *FlowRun) GetFlowName() *string {
	return s.FlowName
}

func (s *FlowRun) GetFlowRunId() *string {
	return s.FlowRunId
}

func (s *FlowRun) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *FlowRun) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *FlowRun) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *FlowRun) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *FlowRun) GetLabels() []*FlowRunLabels {
	return s.Labels
}

func (s *FlowRun) GetNodeName() *string {
	return s.NodeName
}

func (s *FlowRun) GetNodeRunInfos() []*FlowRunNodeRunInfos {
	return s.NodeRunInfos
}

func (s *FlowRun) GetResourceId() *string {
	return s.ResourceId
}

func (s *FlowRun) GetRunMode() *string {
	return s.RunMode
}

func (s *FlowRun) GetRunName() *string {
	return s.RunName
}

func (s *FlowRun) GetRunResult() *string {
	return s.RunResult
}

func (s *FlowRun) GetRunStatus() *string {
	return s.RunStatus
}

func (s *FlowRun) GetRunTimeout() *int32 {
	return s.RunTimeout
}

func (s *FlowRun) GetUserVpc() *FlowRunUserVpc {
	return s.UserVpc
}

func (s *FlowRun) GetVariant() *string {
	return s.Variant
}

func (s *FlowRun) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *FlowRun) SetAccessibility(v string) *FlowRun {
	s.Accessibility = &v
	return s
}

func (s *FlowRun) SetChildRuns(v []*FlowRunChildRuns) *FlowRun {
	s.ChildRuns = v
	return s
}

func (s *FlowRun) SetCreator(v string) *FlowRun {
	s.Creator = &v
	return s
}

func (s *FlowRun) SetCredentialConfig(v *FlowRunCredentialConfig) *FlowRun {
	s.CredentialConfig = v
	return s
}

func (s *FlowRun) SetDataColumnMapping(v map[string]*string) *FlowRun {
	s.DataColumnMapping = v
	return s
}

func (s *FlowRun) SetDataSources(v []*FlowRunDataSources) *FlowRun {
	s.DataSources = v
	return s
}

func (s *FlowRun) SetDuration(v int32) *FlowRun {
	s.Duration = &v
	return s
}

func (s *FlowRun) SetEcsSpec(v *FlowRunEcsSpec) *FlowRun {
	s.EcsSpec = v
	return s
}

func (s *FlowRun) SetEnvs(v []*FlowRunEnvs) *FlowRun {
	s.Envs = v
	return s
}

func (s *FlowRun) SetEvaluationConfigs(v []*FlowRunEvaluationConfigs) *FlowRun {
	s.EvaluationConfigs = v
	return s
}

func (s *FlowRun) SetEvaluationWorkerCount(v int32) *FlowRun {
	s.EvaluationWorkerCount = &v
	return s
}

func (s *FlowRun) SetException(v string) *FlowRun {
	s.Exception = &v
	return s
}

func (s *FlowRun) SetFlowId(v string) *FlowRun {
	s.FlowId = &v
	return s
}

func (s *FlowRun) SetFlowName(v string) *FlowRun {
	s.FlowName = &v
	return s
}

func (s *FlowRun) SetFlowRunId(v string) *FlowRun {
	s.FlowRunId = &v
	return s
}

func (s *FlowRun) SetGmtCreateTime(v string) *FlowRun {
	s.GmtCreateTime = &v
	return s
}

func (s *FlowRun) SetGmtFinishTime(v string) *FlowRun {
	s.GmtFinishTime = &v
	return s
}

func (s *FlowRun) SetGmtModifiedTime(v string) *FlowRun {
	s.GmtModifiedTime = &v
	return s
}

func (s *FlowRun) SetGmtStartTime(v string) *FlowRun {
	s.GmtStartTime = &v
	return s
}

func (s *FlowRun) SetLabels(v []*FlowRunLabels) *FlowRun {
	s.Labels = v
	return s
}

func (s *FlowRun) SetNodeName(v string) *FlowRun {
	s.NodeName = &v
	return s
}

func (s *FlowRun) SetNodeRunInfos(v []*FlowRunNodeRunInfos) *FlowRun {
	s.NodeRunInfos = v
	return s
}

func (s *FlowRun) SetResourceId(v string) *FlowRun {
	s.ResourceId = &v
	return s
}

func (s *FlowRun) SetRunMode(v string) *FlowRun {
	s.RunMode = &v
	return s
}

func (s *FlowRun) SetRunName(v string) *FlowRun {
	s.RunName = &v
	return s
}

func (s *FlowRun) SetRunResult(v string) *FlowRun {
	s.RunResult = &v
	return s
}

func (s *FlowRun) SetRunStatus(v string) *FlowRun {
	s.RunStatus = &v
	return s
}

func (s *FlowRun) SetRunTimeout(v int32) *FlowRun {
	s.RunTimeout = &v
	return s
}

func (s *FlowRun) SetUserVpc(v *FlowRunUserVpc) *FlowRun {
	s.UserVpc = v
	return s
}

func (s *FlowRun) SetVariant(v string) *FlowRun {
	s.Variant = &v
	return s
}

func (s *FlowRun) SetWorkspaceId(v string) *FlowRun {
	s.WorkspaceId = &v
	return s
}

func (s *FlowRun) Validate() error {
	if s.ChildRuns != nil {
		for _, item := range s.ChildRuns {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EcsSpec != nil {
		if err := s.EcsSpec.Validate(); err != nil {
			return err
		}
	}
	if s.Envs != nil {
		for _, item := range s.Envs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.EvaluationConfigs != nil {
		for _, item := range s.EvaluationConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NodeRunInfos != nil {
		for _, item := range s.NodeRunInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.UserVpc != nil {
		if err := s.UserVpc.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FlowRunChildRuns struct {
	// 运行时长
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 应用流运行ID
	FlowRunId *string `json:"FlowRunId,omitempty" xml:"FlowRunId,omitempty"`
	// 应用流来源
	FlowSource *FlowRunChildRunsFlowSource `json:"FlowSource,omitempty" xml:"FlowSource,omitempty" type:"Struct"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 结束时间
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// 修改时间
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// 开始时间
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// 任务信息
	JobInfo *FlowRunChildRunsJobInfo `json:"JobInfo,omitempty" xml:"JobInfo,omitempty" type:"Struct"`
	// 运行模式
	RunMode *string `json:"RunMode,omitempty" xml:"RunMode,omitempty"`
	// 运行名称
	RunName *string `json:"RunName,omitempty" xml:"RunName,omitempty"`
	// 运行结果
	RunResult *string `json:"RunResult,omitempty" xml:"RunResult,omitempty"`
	// 运行状态
	RunStatus *string `json:"RunStatus,omitempty" xml:"RunStatus,omitempty"`
}

func (s FlowRunChildRuns) String() string {
	return dara.Prettify(s)
}

func (s FlowRunChildRuns) GoString() string {
	return s.String()
}

func (s *FlowRunChildRuns) GetDuration() *int32 {
	return s.Duration
}

func (s *FlowRunChildRuns) GetFlowRunId() *string {
	return s.FlowRunId
}

func (s *FlowRunChildRuns) GetFlowSource() *FlowRunChildRunsFlowSource {
	return s.FlowSource
}

func (s *FlowRunChildRuns) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *FlowRunChildRuns) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *FlowRunChildRuns) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *FlowRunChildRuns) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *FlowRunChildRuns) GetJobInfo() *FlowRunChildRunsJobInfo {
	return s.JobInfo
}

func (s *FlowRunChildRuns) GetRunMode() *string {
	return s.RunMode
}

func (s *FlowRunChildRuns) GetRunName() *string {
	return s.RunName
}

func (s *FlowRunChildRuns) GetRunResult() *string {
	return s.RunResult
}

func (s *FlowRunChildRuns) GetRunStatus() *string {
	return s.RunStatus
}

func (s *FlowRunChildRuns) SetDuration(v int32) *FlowRunChildRuns {
	s.Duration = &v
	return s
}

func (s *FlowRunChildRuns) SetFlowRunId(v string) *FlowRunChildRuns {
	s.FlowRunId = &v
	return s
}

func (s *FlowRunChildRuns) SetFlowSource(v *FlowRunChildRunsFlowSource) *FlowRunChildRuns {
	s.FlowSource = v
	return s
}

func (s *FlowRunChildRuns) SetGmtCreateTime(v string) *FlowRunChildRuns {
	s.GmtCreateTime = &v
	return s
}

func (s *FlowRunChildRuns) SetGmtFinishTime(v string) *FlowRunChildRuns {
	s.GmtFinishTime = &v
	return s
}

func (s *FlowRunChildRuns) SetGmtModifiedTime(v string) *FlowRunChildRuns {
	s.GmtModifiedTime = &v
	return s
}

func (s *FlowRunChildRuns) SetGmtStartTime(v string) *FlowRunChildRuns {
	s.GmtStartTime = &v
	return s
}

func (s *FlowRunChildRuns) SetJobInfo(v *FlowRunChildRunsJobInfo) *FlowRunChildRuns {
	s.JobInfo = v
	return s
}

func (s *FlowRunChildRuns) SetRunMode(v string) *FlowRunChildRuns {
	s.RunMode = &v
	return s
}

func (s *FlowRunChildRuns) SetRunName(v string) *FlowRunChildRuns {
	s.RunName = &v
	return s
}

func (s *FlowRunChildRuns) SetRunResult(v string) *FlowRunChildRuns {
	s.RunResult = &v
	return s
}

func (s *FlowRunChildRuns) SetRunStatus(v string) *FlowRunChildRuns {
	s.RunStatus = &v
	return s
}

func (s *FlowRunChildRuns) Validate() error {
	if s.FlowSource != nil {
		if err := s.FlowSource.Validate(); err != nil {
			return err
		}
	}
	if s.JobInfo != nil {
		if err := s.JobInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FlowRunChildRunsFlowSource struct {
	// ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s FlowRunChildRunsFlowSource) String() string {
	return dara.Prettify(s)
}

func (s FlowRunChildRunsFlowSource) GoString() string {
	return s.String()
}

func (s *FlowRunChildRunsFlowSource) GetId() *string {
	return s.Id
}

func (s *FlowRunChildRunsFlowSource) GetName() *string {
	return s.Name
}

func (s *FlowRunChildRunsFlowSource) GetType() *string {
	return s.Type
}

func (s *FlowRunChildRunsFlowSource) SetId(v string) *FlowRunChildRunsFlowSource {
	s.Id = &v
	return s
}

func (s *FlowRunChildRunsFlowSource) SetName(v string) *FlowRunChildRunsFlowSource {
	s.Name = &v
	return s
}

func (s *FlowRunChildRunsFlowSource) SetType(v string) *FlowRunChildRunsFlowSource {
	s.Type = &v
	return s
}

func (s *FlowRunChildRunsFlowSource) Validate() error {
	return dara.Validate(s)
}

type FlowRunChildRunsJobInfo struct {
	// 任务ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s FlowRunChildRunsJobInfo) String() string {
	return dara.Prettify(s)
}

func (s FlowRunChildRunsJobInfo) GoString() string {
	return s.String()
}

func (s *FlowRunChildRunsJobInfo) GetJobId() *string {
	return s.JobId
}

func (s *FlowRunChildRunsJobInfo) SetJobId(v string) *FlowRunChildRunsJobInfo {
	s.JobId = &v
	return s
}

func (s *FlowRunChildRunsJobInfo) Validate() error {
	return dara.Validate(s)
}

type FlowRunCredentialConfig struct {
	// AliyunEnvRoleKey
	AliyunEnvRoleKey *string `json:"AliyunEnvRoleKey,omitempty" xml:"AliyunEnvRoleKey,omitempty"`
	// Credential配置项列表
	CredentialConfigItems []*FlowRunCredentialConfigCredentialConfigItems `json:"CredentialConfigItems,omitempty" xml:"CredentialConfigItems,omitempty" type:"Repeated"`
	// 是否启用Credential注入
	EnableCredentialInject *bool `json:"EnableCredentialInject,omitempty" xml:"EnableCredentialInject,omitempty"`
}

func (s FlowRunCredentialConfig) String() string {
	return dara.Prettify(s)
}

func (s FlowRunCredentialConfig) GoString() string {
	return s.String()
}

func (s *FlowRunCredentialConfig) GetAliyunEnvRoleKey() *string {
	return s.AliyunEnvRoleKey
}

func (s *FlowRunCredentialConfig) GetCredentialConfigItems() []*FlowRunCredentialConfigCredentialConfigItems {
	return s.CredentialConfigItems
}

func (s *FlowRunCredentialConfig) GetEnableCredentialInject() *bool {
	return s.EnableCredentialInject
}

func (s *FlowRunCredentialConfig) SetAliyunEnvRoleKey(v string) *FlowRunCredentialConfig {
	s.AliyunEnvRoleKey = &v
	return s
}

func (s *FlowRunCredentialConfig) SetCredentialConfigItems(v []*FlowRunCredentialConfigCredentialConfigItems) *FlowRunCredentialConfig {
	s.CredentialConfigItems = v
	return s
}

func (s *FlowRunCredentialConfig) SetEnableCredentialInject(v bool) *FlowRunCredentialConfig {
	s.EnableCredentialInject = &v
	return s
}

func (s *FlowRunCredentialConfig) Validate() error {
	if s.CredentialConfigItems != nil {
		for _, item := range s.CredentialConfigItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FlowRunCredentialConfigCredentialConfigItems struct {
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 角色列表
	Roles []*FlowRunCredentialConfigCredentialConfigItemsRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// Type
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s FlowRunCredentialConfigCredentialConfigItems) String() string {
	return dara.Prettify(s)
}

func (s FlowRunCredentialConfigCredentialConfigItems) GoString() string {
	return s.String()
}

func (s *FlowRunCredentialConfigCredentialConfigItems) GetKey() *string {
	return s.Key
}

func (s *FlowRunCredentialConfigCredentialConfigItems) GetRoles() []*FlowRunCredentialConfigCredentialConfigItemsRoles {
	return s.Roles
}

func (s *FlowRunCredentialConfigCredentialConfigItems) GetType() *string {
	return s.Type
}

func (s *FlowRunCredentialConfigCredentialConfigItems) SetKey(v string) *FlowRunCredentialConfigCredentialConfigItems {
	s.Key = &v
	return s
}

func (s *FlowRunCredentialConfigCredentialConfigItems) SetRoles(v []*FlowRunCredentialConfigCredentialConfigItemsRoles) *FlowRunCredentialConfigCredentialConfigItems {
	s.Roles = v
	return s
}

func (s *FlowRunCredentialConfigCredentialConfigItems) SetType(v string) *FlowRunCredentialConfigCredentialConfigItems {
	s.Type = &v
	return s
}

func (s *FlowRunCredentialConfigCredentialConfigItems) Validate() error {
	if s.Roles != nil {
		for _, item := range s.Roles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type FlowRunCredentialConfigCredentialConfigItemsRoles struct {
	// AssumeRoleFor
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// 角色名称
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// 角色类型
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s FlowRunCredentialConfigCredentialConfigItemsRoles) String() string {
	return dara.Prettify(s)
}

func (s FlowRunCredentialConfigCredentialConfigItemsRoles) GoString() string {
	return s.String()
}

func (s *FlowRunCredentialConfigCredentialConfigItemsRoles) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *FlowRunCredentialConfigCredentialConfigItemsRoles) GetRoleArn() *string {
	return s.RoleArn
}

func (s *FlowRunCredentialConfigCredentialConfigItemsRoles) GetRoleType() *string {
	return s.RoleType
}

func (s *FlowRunCredentialConfigCredentialConfigItemsRoles) SetAssumeRoleFor(v string) *FlowRunCredentialConfigCredentialConfigItemsRoles {
	s.AssumeRoleFor = &v
	return s
}

func (s *FlowRunCredentialConfigCredentialConfigItemsRoles) SetRoleArn(v string) *FlowRunCredentialConfigCredentialConfigItemsRoles {
	s.RoleArn = &v
	return s
}

func (s *FlowRunCredentialConfigCredentialConfigItemsRoles) SetRoleType(v string) *FlowRunCredentialConfigCredentialConfigItemsRoles {
	s.RoleType = &v
	return s
}

func (s *FlowRunCredentialConfigCredentialConfigItemsRoles) Validate() error {
	return dara.Validate(s)
}

type FlowRunDataSources struct {
	// 数据集ID
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 统一资源识别码
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s FlowRunDataSources) String() string {
	return dara.Prettify(s)
}

func (s FlowRunDataSources) GoString() string {
	return s.String()
}

func (s *FlowRunDataSources) GetDatasetId() *string {
	return s.DatasetId
}

func (s *FlowRunDataSources) GetMountPath() *string {
	return s.MountPath
}

func (s *FlowRunDataSources) GetUri() *string {
	return s.Uri
}

func (s *FlowRunDataSources) SetDatasetId(v string) *FlowRunDataSources {
	s.DatasetId = &v
	return s
}

func (s *FlowRunDataSources) SetMountPath(v string) *FlowRunDataSources {
	s.MountPath = &v
	return s
}

func (s *FlowRunDataSources) SetUri(v string) *FlowRunDataSources {
	s.Uri = &v
	return s
}

func (s *FlowRunDataSources) Validate() error {
	return dara.Validate(s)
}

type FlowRunEcsSpec struct {
	// CPU信息
	CPU *int32 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// 额外系统盘
	ExtraEphemeralStorage *int32 `json:"ExtraEphemeralStorage,omitempty" xml:"ExtraEphemeralStorage,omitempty"`
	// GPU信息
	GPU *int32 `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// GPU算力占比
	GPUCorePercentage *int32 `json:"GPUCorePercentage,omitempty" xml:"GPUCorePercentage,omitempty"`
	// GPU显存
	GPUMemory *int32 `json:"GPUMemory,omitempty" xml:"GPUMemory,omitempty"`
	// GPU类型
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// 实例类型
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 内存信息
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// Pod数量
	PodCount *int32 `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	// 资源配额ID
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// 资源配额类型
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// 共享内存
	SharedMemory *int32 `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s FlowRunEcsSpec) String() string {
	return dara.Prettify(s)
}

func (s FlowRunEcsSpec) GoString() string {
	return s.String()
}

func (s *FlowRunEcsSpec) GetCPU() *int32 {
	return s.CPU
}

func (s *FlowRunEcsSpec) GetExtraEphemeralStorage() *int32 {
	return s.ExtraEphemeralStorage
}

func (s *FlowRunEcsSpec) GetGPU() *int32 {
	return s.GPU
}

func (s *FlowRunEcsSpec) GetGPUCorePercentage() *int32 {
	return s.GPUCorePercentage
}

func (s *FlowRunEcsSpec) GetGPUMemory() *int32 {
	return s.GPUMemory
}

func (s *FlowRunEcsSpec) GetGPUType() *string {
	return s.GPUType
}

func (s *FlowRunEcsSpec) GetInstanceType() *string {
	return s.InstanceType
}

func (s *FlowRunEcsSpec) GetMemory() *int32 {
	return s.Memory
}

func (s *FlowRunEcsSpec) GetPodCount() *int32 {
	return s.PodCount
}

func (s *FlowRunEcsSpec) GetQuotaId() *string {
	return s.QuotaId
}

func (s *FlowRunEcsSpec) GetQuotaType() *string {
	return s.QuotaType
}

func (s *FlowRunEcsSpec) GetSharedMemory() *int32 {
	return s.SharedMemory
}

func (s *FlowRunEcsSpec) SetCPU(v int32) *FlowRunEcsSpec {
	s.CPU = &v
	return s
}

func (s *FlowRunEcsSpec) SetExtraEphemeralStorage(v int32) *FlowRunEcsSpec {
	s.ExtraEphemeralStorage = &v
	return s
}

func (s *FlowRunEcsSpec) SetGPU(v int32) *FlowRunEcsSpec {
	s.GPU = &v
	return s
}

func (s *FlowRunEcsSpec) SetGPUCorePercentage(v int32) *FlowRunEcsSpec {
	s.GPUCorePercentage = &v
	return s
}

func (s *FlowRunEcsSpec) SetGPUMemory(v int32) *FlowRunEcsSpec {
	s.GPUMemory = &v
	return s
}

func (s *FlowRunEcsSpec) SetGPUType(v string) *FlowRunEcsSpec {
	s.GPUType = &v
	return s
}

func (s *FlowRunEcsSpec) SetInstanceType(v string) *FlowRunEcsSpec {
	s.InstanceType = &v
	return s
}

func (s *FlowRunEcsSpec) SetMemory(v int32) *FlowRunEcsSpec {
	s.Memory = &v
	return s
}

func (s *FlowRunEcsSpec) SetPodCount(v int32) *FlowRunEcsSpec {
	s.PodCount = &v
	return s
}

func (s *FlowRunEcsSpec) SetQuotaId(v string) *FlowRunEcsSpec {
	s.QuotaId = &v
	return s
}

func (s *FlowRunEcsSpec) SetQuotaType(v string) *FlowRunEcsSpec {
	s.QuotaType = &v
	return s
}

func (s *FlowRunEcsSpec) SetSharedMemory(v int32) *FlowRunEcsSpec {
	s.SharedMemory = &v
	return s
}

func (s *FlowRunEcsSpec) Validate() error {
	return dara.Validate(s)
}

type FlowRunEnvs struct {
	// 环境键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 环境值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s FlowRunEnvs) String() string {
	return dara.Prettify(s)
}

func (s FlowRunEnvs) GoString() string {
	return s.String()
}

func (s *FlowRunEnvs) GetKey() *string {
	return s.Key
}

func (s *FlowRunEnvs) GetValue() *string {
	return s.Value
}

func (s *FlowRunEnvs) SetKey(v string) *FlowRunEnvs {
	s.Key = &v
	return s
}

func (s *FlowRunEnvs) SetValue(v string) *FlowRunEnvs {
	s.Value = &v
	return s
}

func (s *FlowRunEnvs) Validate() error {
	return dara.Validate(s)
}

type FlowRunEvaluationConfigs struct {
	// 映射配置
	DataColumnMapping map[string]*string `json:"DataColumnMapping,omitempty" xml:"DataColumnMapping,omitempty"`
	// 应用流来源
	FlowSource *FlowRunEvaluationConfigsFlowSource `json:"FlowSource,omitempty" xml:"FlowSource,omitempty" type:"Struct"`
	// 输入配置
	InputsOverrideConfig *string `json:"InputsOverrideConfig,omitempty" xml:"InputsOverrideConfig,omitempty"`
}

func (s FlowRunEvaluationConfigs) String() string {
	return dara.Prettify(s)
}

func (s FlowRunEvaluationConfigs) GoString() string {
	return s.String()
}

func (s *FlowRunEvaluationConfigs) GetDataColumnMapping() map[string]*string {
	return s.DataColumnMapping
}

func (s *FlowRunEvaluationConfigs) GetFlowSource() *FlowRunEvaluationConfigsFlowSource {
	return s.FlowSource
}

func (s *FlowRunEvaluationConfigs) GetInputsOverrideConfig() *string {
	return s.InputsOverrideConfig
}

func (s *FlowRunEvaluationConfigs) SetDataColumnMapping(v map[string]*string) *FlowRunEvaluationConfigs {
	s.DataColumnMapping = v
	return s
}

func (s *FlowRunEvaluationConfigs) SetFlowSource(v *FlowRunEvaluationConfigsFlowSource) *FlowRunEvaluationConfigs {
	s.FlowSource = v
	return s
}

func (s *FlowRunEvaluationConfigs) SetInputsOverrideConfig(v string) *FlowRunEvaluationConfigs {
	s.InputsOverrideConfig = &v
	return s
}

func (s *FlowRunEvaluationConfigs) Validate() error {
	if s.FlowSource != nil {
		if err := s.FlowSource.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FlowRunEvaluationConfigsFlowSource struct {
	// ID
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// 名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s FlowRunEvaluationConfigsFlowSource) String() string {
	return dara.Prettify(s)
}

func (s FlowRunEvaluationConfigsFlowSource) GoString() string {
	return s.String()
}

func (s *FlowRunEvaluationConfigsFlowSource) GetId() *string {
	return s.Id
}

func (s *FlowRunEvaluationConfigsFlowSource) GetName() *string {
	return s.Name
}

func (s *FlowRunEvaluationConfigsFlowSource) GetType() *string {
	return s.Type
}

func (s *FlowRunEvaluationConfigsFlowSource) SetId(v string) *FlowRunEvaluationConfigsFlowSource {
	s.Id = &v
	return s
}

func (s *FlowRunEvaluationConfigsFlowSource) SetName(v string) *FlowRunEvaluationConfigsFlowSource {
	s.Name = &v
	return s
}

func (s *FlowRunEvaluationConfigsFlowSource) SetType(v string) *FlowRunEvaluationConfigsFlowSource {
	s.Type = &v
	return s
}

func (s *FlowRunEvaluationConfigsFlowSource) Validate() error {
	return dara.Validate(s)
}

type FlowRunLabels struct {
	// 标签键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s FlowRunLabels) String() string {
	return dara.Prettify(s)
}

func (s FlowRunLabels) GoString() string {
	return s.String()
}

func (s *FlowRunLabels) GetKey() *string {
	return s.Key
}

func (s *FlowRunLabels) GetValue() *string {
	return s.Value
}

func (s *FlowRunLabels) SetKey(v string) *FlowRunLabels {
	s.Key = &v
	return s
}

func (s *FlowRunLabels) SetValue(v string) *FlowRunLabels {
	s.Value = &v
	return s
}

func (s *FlowRunLabels) Validate() error {
	return dara.Validate(s)
}

type FlowRunNodeRunInfos struct {
	// 耗时
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 输入
	Inputs *string `json:"Inputs,omitempty" xml:"Inputs,omitempty"`
	// 节点名称
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// 输出
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// 节点状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 日志信息
	Stdout *string `json:"Stdout,omitempty" xml:"Stdout,omitempty"`
}

func (s FlowRunNodeRunInfos) String() string {
	return dara.Prettify(s)
}

func (s FlowRunNodeRunInfos) GoString() string {
	return s.String()
}

func (s *FlowRunNodeRunInfos) GetDuration() *string {
	return s.Duration
}

func (s *FlowRunNodeRunInfos) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *FlowRunNodeRunInfos) GetInputs() *string {
	return s.Inputs
}

func (s *FlowRunNodeRunInfos) GetNodeName() *string {
	return s.NodeName
}

func (s *FlowRunNodeRunInfos) GetOutput() *string {
	return s.Output
}

func (s *FlowRunNodeRunInfos) GetStatus() *string {
	return s.Status
}

func (s *FlowRunNodeRunInfos) GetStdout() *string {
	return s.Stdout
}

func (s *FlowRunNodeRunInfos) SetDuration(v string) *FlowRunNodeRunInfos {
	s.Duration = &v
	return s
}

func (s *FlowRunNodeRunInfos) SetErrorMessage(v string) *FlowRunNodeRunInfos {
	s.ErrorMessage = &v
	return s
}

func (s *FlowRunNodeRunInfos) SetInputs(v string) *FlowRunNodeRunInfos {
	s.Inputs = &v
	return s
}

func (s *FlowRunNodeRunInfos) SetNodeName(v string) *FlowRunNodeRunInfos {
	s.NodeName = &v
	return s
}

func (s *FlowRunNodeRunInfos) SetOutput(v string) *FlowRunNodeRunInfos {
	s.Output = &v
	return s
}

func (s *FlowRunNodeRunInfos) SetStatus(v string) *FlowRunNodeRunInfos {
	s.Status = &v
	return s
}

func (s *FlowRunNodeRunInfos) SetStdout(v string) *FlowRunNodeRunInfos {
	s.Stdout = &v
	return s
}

func (s *FlowRunNodeRunInfos) Validate() error {
	return dara.Validate(s)
}

type FlowRunUserVpc struct {
	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 交换机ID
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s FlowRunUserVpc) String() string {
	return dara.Prettify(s)
}

func (s FlowRunUserVpc) GoString() string {
	return s.String()
}

func (s *FlowRunUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *FlowRunUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *FlowRunUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *FlowRunUserVpc) SetSecurityGroupId(v string) *FlowRunUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *FlowRunUserVpc) SetVSwitchId(v string) *FlowRunUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *FlowRunUserVpc) SetVpcId(v string) *FlowRunUserVpc {
	s.VpcId = &v
	return s
}

func (s *FlowRunUserVpc) Validate() error {
	return dara.Validate(s)
}
