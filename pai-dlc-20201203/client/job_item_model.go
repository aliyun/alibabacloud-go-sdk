// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJobItem interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *JobItem
	GetAccessibility() *string
	SetClusterId(v string) *JobItem
	GetClusterId() *string
	SetCodeSource(v *JobItemCodeSource) *JobItem
	GetCodeSource() *JobItemCodeSource
	SetCredentialConfig(v *CredentialConfig) *JobItem
	GetCredentialConfig() *CredentialConfig
	SetDataSources(v []*JobItemDataSources) *JobItem
	GetDataSources() []*JobItemDataSources
	SetDisplayName(v string) *JobItem
	GetDisplayName() *string
	SetDuration(v int64) *JobItem
	GetDuration() *int64
	SetElasticSpec(v *JobElasticSpec) *JobItem
	GetElasticSpec() *JobElasticSpec
	SetEnablePreemptibleJob(v bool) *JobItem
	GetEnablePreemptibleJob() *bool
	SetEnabledDebugger(v bool) *JobItem
	GetEnabledDebugger() *bool
	SetEnvs(v map[string]*string) *JobItem
	GetEnvs() map[string]*string
	SetGmtCreateTime(v string) *JobItem
	GetGmtCreateTime() *string
	SetGmtFailedTime(v string) *JobItem
	GetGmtFailedTime() *string
	SetGmtFinishTime(v string) *JobItem
	GetGmtFinishTime() *string
	SetGmtModifiedTime(v string) *JobItem
	GetGmtModifiedTime() *string
	SetGmtRunningTime(v string) *JobItem
	GetGmtRunningTime() *string
	SetGmtStoppedTime(v string) *JobItem
	GetGmtStoppedTime() *string
	SetGmtSubmittedTime(v string) *JobItem
	GetGmtSubmittedTime() *string
	SetGmtSuccessedTime(v string) *JobItem
	GetGmtSuccessedTime() *string
	SetIsDeleted(v bool) *JobItem
	GetIsDeleted() *bool
	SetJobId(v string) *JobItem
	GetJobId() *string
	SetJobMaxRunningTimeMinutes(v int64) *JobItem
	GetJobMaxRunningTimeMinutes() *int64
	SetJobSpecs(v []*JobSpec) *JobItem
	GetJobSpecs() []*JobSpec
	SetJobType(v string) *JobItem
	GetJobType() *string
	SetNodeCount(v string) *JobItem
	GetNodeCount() *string
	SetNodeNames(v []*string) *JobItem
	GetNodeNames() []*string
	SetPods(v []*PodItem) *JobItem
	GetPods() []*PodItem
	SetPriority(v int32) *JobItem
	GetPriority() *int32
	SetReasonCode(v string) *JobItem
	GetReasonCode() *string
	SetReasonMessage(v string) *JobItem
	GetReasonMessage() *string
	SetRequestCPU(v int64) *JobItem
	GetRequestCPU() *int64
	SetRequestGPU(v string) *JobItem
	GetRequestGPU() *string
	SetRequestMemory(v string) *JobItem
	GetRequestMemory() *string
	SetResourceId(v string) *JobItem
	GetResourceId() *string
	SetResourceLevel(v string) *JobItem
	GetResourceLevel() *string
	SetResourceName(v string) *JobItem
	GetResourceName() *string
	SetResourceQuotaName(v string) *JobItem
	GetResourceQuotaName() *string
	SetResourceType(v string) *JobItem
	GetResourceType() *string
	SetRestartTimes(v string) *JobItem
	GetRestartTimes() *string
	SetSettings(v *JobSettings) *JobItem
	GetSettings() *JobSettings
	SetStatus(v string) *JobItem
	GetStatus() *string
	SetStatusHistory(v []*StatusTransitionItem) *JobItem
	GetStatusHistory() []*StatusTransitionItem
	SetSubStatus(v string) *JobItem
	GetSubStatus() *string
	SetSystemEnvs(v map[string]*string) *JobItem
	GetSystemEnvs() map[string]*string
	SetTenantId(v string) *JobItem
	GetTenantId() *string
	SetThirdpartyLibDir(v string) *JobItem
	GetThirdpartyLibDir() *string
	SetThirdpartyLibs(v []*string) *JobItem
	GetThirdpartyLibs() []*string
	SetUseOversoldResource(v bool) *JobItem
	GetUseOversoldResource() *bool
	SetUserCommand(v string) *JobItem
	GetUserCommand() *string
	SetUserId(v string) *JobItem
	GetUserId() *string
	SetUserScript(v string) *JobItem
	GetUserScript() *string
	SetUserVpc(v *JobItemUserVpc) *JobItem
	GetUserVpc() *JobItemUserVpc
	SetUsername(v string) *JobItem
	GetUsername() *string
	SetWorkingDir(v string) *JobItem
	GetWorkingDir() *string
	SetWorkspaceId(v string) *JobItem
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *JobItem
	GetWorkspaceName() *string
}

type JobItem struct {
	// example:
	//
	// PUBLIC
	Accessibility    *string               `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	ClusterId        *string               `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	CodeSource       *JobItemCodeSource    `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	CredentialConfig *CredentialConfig     `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	DataSources      []*JobItemDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// example:
	//
	// tf-mnist-test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 3602
	Duration    *int64          `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ElasticSpec *JobElasticSpec `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
	// example:
	//
	// false
	EnablePreemptibleJob *bool `json:"EnablePreemptibleJob,omitempty" xml:"EnablePreemptibleJob,omitempty"`
	// example:
	//
	// false
	EnabledDebugger *bool              `json:"EnabledDebugger,omitempty" xml:"EnabledDebugger,omitempty"`
	Envs            map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtFailedTime *string `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtRunningTime *string `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtStoppedTime *string `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtSubmittedTime *string `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtSuccessedTime *string `json:"GmtSuccessedTime,omitempty" xml:"GmtSuccessedTime,omitempty"`
	// example:
	//
	// false
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// example:
	//
	// dlc-20210126170216-mtl37ge7gkvdz
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 1
	JobMaxRunningTimeMinutes *int64     `json:"JobMaxRunningTimeMinutes,omitempty" xml:"JobMaxRunningTimeMinutes,omitempty"`
	JobSpecs                 []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// example:
	//
	// TFJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// example:
	//
	// 1
	NodeCount *string    `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeNames []*string  `json:"NodeNames,omitempty" xml:"NodeNames,omitempty" type:"Repeated"`
	Pods      []*PodItem `json:"Pods,omitempty" xml:"Pods,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// JobStoppedByUser
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// example:
	//
	// Job is stopped by user.
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// example:
	//
	// 1
	RequestCPU *int64 `json:"RequestCPU,omitempty" xml:"RequestCPU,omitempty"`
	// example:
	//
	// 1
	RequestGPU *string `json:"RequestGPU,omitempty" xml:"RequestGPU,omitempty"`
	// example:
	//
	// 1Gi
	RequestMemory *string `json:"RequestMemory,omitempty" xml:"RequestMemory,omitempty"`
	// example:
	//
	// dlc-quota
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// L0
	ResourceLevel *string `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
	// example:
	//
	// my_resource_group
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// example:
	//
	// test
	ResourceQuotaName *string `json:"ResourceQuotaName,omitempty" xml:"ResourceQuotaName,omitempty"`
	// example:
	//
	// ECS
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// 1
	RestartTimes *string      `json:"RestartTimes,omitempty" xml:"RestartTimes,omitempty"`
	Settings     *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// example:
	//
	// Stopped
	Status        *string                 `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusHistory []*StatusTransitionItem `json:"StatusHistory,omitempty" xml:"StatusHistory,omitempty" type:"Repeated"`
	// example:
	//
	// Restarting
	SubStatus  *string            `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	SystemEnvs map[string]*string `json:"SystemEnvs,omitempty" xml:"SystemEnvs,omitempty"`
	TenantId   *string            `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// example:
	//
	// /root/code/
	ThirdpartyLibDir *string   `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	ThirdpartyLibs   []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// example:
	//
	// false
	UseOversoldResource *bool `json:"UseOversoldResource,omitempty" xml:"UseOversoldResource,omitempty"`
	// example:
	//
	// python /root/code/mnist.py
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// example:
	//
	// 123456789
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// ls
	UserScript *string `json:"UserScript,omitempty" xml:"UserScript,omitempty"`
	// example:
	//
	// vpc-1
	UserVpc *JobItemUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// example:
	//
	// pai-dlc-role
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// example:
	//
	// /mnt/data
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	// example:
	//
	// 268
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// example:
	//
	// dlc-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s JobItem) String() string {
	return dara.Prettify(s)
}

func (s JobItem) GoString() string {
	return s.String()
}

func (s *JobItem) GetAccessibility() *string {
	return s.Accessibility
}

func (s *JobItem) GetClusterId() *string {
	return s.ClusterId
}

func (s *JobItem) GetCodeSource() *JobItemCodeSource {
	return s.CodeSource
}

func (s *JobItem) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *JobItem) GetDataSources() []*JobItemDataSources {
	return s.DataSources
}

func (s *JobItem) GetDisplayName() *string {
	return s.DisplayName
}

func (s *JobItem) GetDuration() *int64 {
	return s.Duration
}

func (s *JobItem) GetElasticSpec() *JobElasticSpec {
	return s.ElasticSpec
}

func (s *JobItem) GetEnablePreemptibleJob() *bool {
	return s.EnablePreemptibleJob
}

func (s *JobItem) GetEnabledDebugger() *bool {
	return s.EnabledDebugger
}

func (s *JobItem) GetEnvs() map[string]*string {
	return s.Envs
}

func (s *JobItem) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *JobItem) GetGmtFailedTime() *string {
	return s.GmtFailedTime
}

func (s *JobItem) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *JobItem) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *JobItem) GetGmtRunningTime() *string {
	return s.GmtRunningTime
}

func (s *JobItem) GetGmtStoppedTime() *string {
	return s.GmtStoppedTime
}

func (s *JobItem) GetGmtSubmittedTime() *string {
	return s.GmtSubmittedTime
}

func (s *JobItem) GetGmtSuccessedTime() *string {
	return s.GmtSuccessedTime
}

func (s *JobItem) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *JobItem) GetJobId() *string {
	return s.JobId
}

func (s *JobItem) GetJobMaxRunningTimeMinutes() *int64 {
	return s.JobMaxRunningTimeMinutes
}

func (s *JobItem) GetJobSpecs() []*JobSpec {
	return s.JobSpecs
}

func (s *JobItem) GetJobType() *string {
	return s.JobType
}

func (s *JobItem) GetNodeCount() *string {
	return s.NodeCount
}

func (s *JobItem) GetNodeNames() []*string {
	return s.NodeNames
}

func (s *JobItem) GetPods() []*PodItem {
	return s.Pods
}

func (s *JobItem) GetPriority() *int32 {
	return s.Priority
}

func (s *JobItem) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *JobItem) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *JobItem) GetRequestCPU() *int64 {
	return s.RequestCPU
}

func (s *JobItem) GetRequestGPU() *string {
	return s.RequestGPU
}

func (s *JobItem) GetRequestMemory() *string {
	return s.RequestMemory
}

func (s *JobItem) GetResourceId() *string {
	return s.ResourceId
}

func (s *JobItem) GetResourceLevel() *string {
	return s.ResourceLevel
}

func (s *JobItem) GetResourceName() *string {
	return s.ResourceName
}

func (s *JobItem) GetResourceQuotaName() *string {
	return s.ResourceQuotaName
}

func (s *JobItem) GetResourceType() *string {
	return s.ResourceType
}

func (s *JobItem) GetRestartTimes() *string {
	return s.RestartTimes
}

func (s *JobItem) GetSettings() *JobSettings {
	return s.Settings
}

func (s *JobItem) GetStatus() *string {
	return s.Status
}

func (s *JobItem) GetStatusHistory() []*StatusTransitionItem {
	return s.StatusHistory
}

func (s *JobItem) GetSubStatus() *string {
	return s.SubStatus
}

func (s *JobItem) GetSystemEnvs() map[string]*string {
	return s.SystemEnvs
}

func (s *JobItem) GetTenantId() *string {
	return s.TenantId
}

func (s *JobItem) GetThirdpartyLibDir() *string {
	return s.ThirdpartyLibDir
}

func (s *JobItem) GetThirdpartyLibs() []*string {
	return s.ThirdpartyLibs
}

func (s *JobItem) GetUseOversoldResource() *bool {
	return s.UseOversoldResource
}

func (s *JobItem) GetUserCommand() *string {
	return s.UserCommand
}

func (s *JobItem) GetUserId() *string {
	return s.UserId
}

func (s *JobItem) GetUserScript() *string {
	return s.UserScript
}

func (s *JobItem) GetUserVpc() *JobItemUserVpc {
	return s.UserVpc
}

func (s *JobItem) GetUsername() *string {
	return s.Username
}

func (s *JobItem) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *JobItem) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *JobItem) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *JobItem) SetAccessibility(v string) *JobItem {
	s.Accessibility = &v
	return s
}

func (s *JobItem) SetClusterId(v string) *JobItem {
	s.ClusterId = &v
	return s
}

func (s *JobItem) SetCodeSource(v *JobItemCodeSource) *JobItem {
	s.CodeSource = v
	return s
}

func (s *JobItem) SetCredentialConfig(v *CredentialConfig) *JobItem {
	s.CredentialConfig = v
	return s
}

func (s *JobItem) SetDataSources(v []*JobItemDataSources) *JobItem {
	s.DataSources = v
	return s
}

func (s *JobItem) SetDisplayName(v string) *JobItem {
	s.DisplayName = &v
	return s
}

func (s *JobItem) SetDuration(v int64) *JobItem {
	s.Duration = &v
	return s
}

func (s *JobItem) SetElasticSpec(v *JobElasticSpec) *JobItem {
	s.ElasticSpec = v
	return s
}

func (s *JobItem) SetEnablePreemptibleJob(v bool) *JobItem {
	s.EnablePreemptibleJob = &v
	return s
}

func (s *JobItem) SetEnabledDebugger(v bool) *JobItem {
	s.EnabledDebugger = &v
	return s
}

func (s *JobItem) SetEnvs(v map[string]*string) *JobItem {
	s.Envs = v
	return s
}

func (s *JobItem) SetGmtCreateTime(v string) *JobItem {
	s.GmtCreateTime = &v
	return s
}

func (s *JobItem) SetGmtFailedTime(v string) *JobItem {
	s.GmtFailedTime = &v
	return s
}

func (s *JobItem) SetGmtFinishTime(v string) *JobItem {
	s.GmtFinishTime = &v
	return s
}

func (s *JobItem) SetGmtModifiedTime(v string) *JobItem {
	s.GmtModifiedTime = &v
	return s
}

func (s *JobItem) SetGmtRunningTime(v string) *JobItem {
	s.GmtRunningTime = &v
	return s
}

func (s *JobItem) SetGmtStoppedTime(v string) *JobItem {
	s.GmtStoppedTime = &v
	return s
}

func (s *JobItem) SetGmtSubmittedTime(v string) *JobItem {
	s.GmtSubmittedTime = &v
	return s
}

func (s *JobItem) SetGmtSuccessedTime(v string) *JobItem {
	s.GmtSuccessedTime = &v
	return s
}

func (s *JobItem) SetIsDeleted(v bool) *JobItem {
	s.IsDeleted = &v
	return s
}

func (s *JobItem) SetJobId(v string) *JobItem {
	s.JobId = &v
	return s
}

func (s *JobItem) SetJobMaxRunningTimeMinutes(v int64) *JobItem {
	s.JobMaxRunningTimeMinutes = &v
	return s
}

func (s *JobItem) SetJobSpecs(v []*JobSpec) *JobItem {
	s.JobSpecs = v
	return s
}

func (s *JobItem) SetJobType(v string) *JobItem {
	s.JobType = &v
	return s
}

func (s *JobItem) SetNodeCount(v string) *JobItem {
	s.NodeCount = &v
	return s
}

func (s *JobItem) SetNodeNames(v []*string) *JobItem {
	s.NodeNames = v
	return s
}

func (s *JobItem) SetPods(v []*PodItem) *JobItem {
	s.Pods = v
	return s
}

func (s *JobItem) SetPriority(v int32) *JobItem {
	s.Priority = &v
	return s
}

func (s *JobItem) SetReasonCode(v string) *JobItem {
	s.ReasonCode = &v
	return s
}

func (s *JobItem) SetReasonMessage(v string) *JobItem {
	s.ReasonMessage = &v
	return s
}

func (s *JobItem) SetRequestCPU(v int64) *JobItem {
	s.RequestCPU = &v
	return s
}

func (s *JobItem) SetRequestGPU(v string) *JobItem {
	s.RequestGPU = &v
	return s
}

func (s *JobItem) SetRequestMemory(v string) *JobItem {
	s.RequestMemory = &v
	return s
}

func (s *JobItem) SetResourceId(v string) *JobItem {
	s.ResourceId = &v
	return s
}

func (s *JobItem) SetResourceLevel(v string) *JobItem {
	s.ResourceLevel = &v
	return s
}

func (s *JobItem) SetResourceName(v string) *JobItem {
	s.ResourceName = &v
	return s
}

func (s *JobItem) SetResourceQuotaName(v string) *JobItem {
	s.ResourceQuotaName = &v
	return s
}

func (s *JobItem) SetResourceType(v string) *JobItem {
	s.ResourceType = &v
	return s
}

func (s *JobItem) SetRestartTimes(v string) *JobItem {
	s.RestartTimes = &v
	return s
}

func (s *JobItem) SetSettings(v *JobSettings) *JobItem {
	s.Settings = v
	return s
}

func (s *JobItem) SetStatus(v string) *JobItem {
	s.Status = &v
	return s
}

func (s *JobItem) SetStatusHistory(v []*StatusTransitionItem) *JobItem {
	s.StatusHistory = v
	return s
}

func (s *JobItem) SetSubStatus(v string) *JobItem {
	s.SubStatus = &v
	return s
}

func (s *JobItem) SetSystemEnvs(v map[string]*string) *JobItem {
	s.SystemEnvs = v
	return s
}

func (s *JobItem) SetTenantId(v string) *JobItem {
	s.TenantId = &v
	return s
}

func (s *JobItem) SetThirdpartyLibDir(v string) *JobItem {
	s.ThirdpartyLibDir = &v
	return s
}

func (s *JobItem) SetThirdpartyLibs(v []*string) *JobItem {
	s.ThirdpartyLibs = v
	return s
}

func (s *JobItem) SetUseOversoldResource(v bool) *JobItem {
	s.UseOversoldResource = &v
	return s
}

func (s *JobItem) SetUserCommand(v string) *JobItem {
	s.UserCommand = &v
	return s
}

func (s *JobItem) SetUserId(v string) *JobItem {
	s.UserId = &v
	return s
}

func (s *JobItem) SetUserScript(v string) *JobItem {
	s.UserScript = &v
	return s
}

func (s *JobItem) SetUserVpc(v *JobItemUserVpc) *JobItem {
	s.UserVpc = v
	return s
}

func (s *JobItem) SetUsername(v string) *JobItem {
	s.Username = &v
	return s
}

func (s *JobItem) SetWorkingDir(v string) *JobItem {
	s.WorkingDir = &v
	return s
}

func (s *JobItem) SetWorkspaceId(v string) *JobItem {
	s.WorkspaceId = &v
	return s
}

func (s *JobItem) SetWorkspaceName(v string) *JobItem {
	s.WorkspaceName = &v
	return s
}

func (s *JobItem) Validate() error {
	if s.CodeSource != nil {
		if err := s.CodeSource.Validate(); err != nil {
			return err
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
	if s.ElasticSpec != nil {
		if err := s.ElasticSpec.Validate(); err != nil {
			return err
		}
	}
	if s.JobSpecs != nil {
		for _, item := range s.JobSpecs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Pods != nil {
		for _, item := range s.Pods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Settings != nil {
		if err := s.Settings.Validate(); err != nil {
			return err
		}
	}
	if s.StatusHistory != nil {
		for _, item := range s.StatusHistory {
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

type JobItemCodeSource struct {
	// example:
	//
	// master
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// example:
	//
	// code-20210111103721-85qz78ia96lu
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// example:
	//
	// 44da109b59f8596152987eaa8f3b2487bb72ea63
	Commit *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s JobItemCodeSource) String() string {
	return dara.Prettify(s)
}

func (s JobItemCodeSource) GoString() string {
	return s.String()
}

func (s *JobItemCodeSource) GetBranch() *string {
	return s.Branch
}

func (s *JobItemCodeSource) GetCodeSourceId() *string {
	return s.CodeSourceId
}

func (s *JobItemCodeSource) GetCommit() *string {
	return s.Commit
}

func (s *JobItemCodeSource) GetMountPath() *string {
	return s.MountPath
}

func (s *JobItemCodeSource) SetBranch(v string) *JobItemCodeSource {
	s.Branch = &v
	return s
}

func (s *JobItemCodeSource) SetCodeSourceId(v string) *JobItemCodeSource {
	s.CodeSourceId = &v
	return s
}

func (s *JobItemCodeSource) SetCommit(v string) *JobItemCodeSource {
	s.Commit = &v
	return s
}

func (s *JobItemCodeSource) SetMountPath(v string) *JobItemCodeSource {
	s.MountPath = &v
	return s
}

func (s *JobItemCodeSource) Validate() error {
	return dara.Validate(s)
}

type JobItemDataSources struct {
	// example:
	//
	// data-20210114104214-vf9lowjt3pso
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s JobItemDataSources) String() string {
	return dara.Prettify(s)
}

func (s JobItemDataSources) GoString() string {
	return s.String()
}

func (s *JobItemDataSources) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *JobItemDataSources) GetMountPath() *string {
	return s.MountPath
}

func (s *JobItemDataSources) SetDataSourceId(v string) *JobItemDataSources {
	s.DataSourceId = &v
	return s
}

func (s *JobItemDataSources) SetMountPath(v string) *JobItemDataSources {
	s.MountPath = &v
	return s
}

func (s *JobItemDataSources) Validate() error {
	return dara.Validate(s)
}

type JobItemUserVpc struct {
	DefaultRoute    *string   `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	ExtendedCidrs   []*string `json:"ExtendedCidrs,omitempty" xml:"ExtendedCidrs,omitempty" type:"Repeated"`
	SecurityGroupId *string   `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SwitchId        *string   `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	VpcId           *string   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s JobItemUserVpc) String() string {
	return dara.Prettify(s)
}

func (s JobItemUserVpc) GoString() string {
	return s.String()
}

func (s *JobItemUserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *JobItemUserVpc) GetExtendedCidrs() []*string {
	return s.ExtendedCidrs
}

func (s *JobItemUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *JobItemUserVpc) GetSwitchId() *string {
	return s.SwitchId
}

func (s *JobItemUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *JobItemUserVpc) SetDefaultRoute(v string) *JobItemUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *JobItemUserVpc) SetExtendedCidrs(v []*string) *JobItemUserVpc {
	s.ExtendedCidrs = v
	return s
}

func (s *JobItemUserVpc) SetSecurityGroupId(v string) *JobItemUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *JobItemUserVpc) SetSwitchId(v string) *JobItemUserVpc {
	s.SwitchId = &v
	return s
}

func (s *JobItemUserVpc) SetVpcId(v string) *JobItemUserVpc {
	s.VpcId = &v
	return s
}

func (s *JobItemUserVpc) Validate() error {
	return dara.Validate(s)
}
