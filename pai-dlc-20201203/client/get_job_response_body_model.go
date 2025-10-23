// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *GetJobResponseBody
	GetAccessibility() *string
	SetClusterId(v string) *GetJobResponseBody
	GetClusterId() *string
	SetCodeSource(v *GetJobResponseBodyCodeSource) *GetJobResponseBody
	GetCodeSource() *GetJobResponseBodyCodeSource
	SetCredentialConfig(v *CredentialConfig) *GetJobResponseBody
	GetCredentialConfig() *CredentialConfig
	SetDataSources(v []*GetJobResponseBodyDataSources) *GetJobResponseBody
	GetDataSources() []*GetJobResponseBodyDataSources
	SetDisplayName(v string) *GetJobResponseBody
	GetDisplayName() *string
	SetDuration(v int64) *GetJobResponseBody
	GetDuration() *int64
	SetElasticSpec(v *JobElasticSpec) *GetJobResponseBody
	GetElasticSpec() *JobElasticSpec
	SetEnabledDebugger(v bool) *GetJobResponseBody
	GetEnabledDebugger() *bool
	SetEnvs(v map[string]*string) *GetJobResponseBody
	GetEnvs() map[string]*string
	SetGmtCreateTime(v string) *GetJobResponseBody
	GetGmtCreateTime() *string
	SetGmtFailedTime(v string) *GetJobResponseBody
	GetGmtFailedTime() *string
	SetGmtFinishTime(v string) *GetJobResponseBody
	GetGmtFinishTime() *string
	SetGmtRunningTime(v string) *GetJobResponseBody
	GetGmtRunningTime() *string
	SetGmtStoppedTime(v string) *GetJobResponseBody
	GetGmtStoppedTime() *string
	SetGmtSubmittedTime(v string) *GetJobResponseBody
	GetGmtSubmittedTime() *string
	SetGmtSuccessedTime(v string) *GetJobResponseBody
	GetGmtSuccessedTime() *string
	SetJobId(v string) *GetJobResponseBody
	GetJobId() *string
	SetJobReplicaStatuses(v []*JobReplicaStatus) *GetJobResponseBody
	GetJobReplicaStatuses() []*JobReplicaStatus
	SetJobSpecs(v []*JobSpec) *GetJobResponseBody
	GetJobSpecs() []*JobSpec
	SetJobType(v string) *GetJobResponseBody
	GetJobType() *string
	SetPods(v []*GetJobResponseBodyPods) *GetJobResponseBody
	GetPods() []*GetJobResponseBodyPods
	SetPriority(v int32) *GetJobResponseBody
	GetPriority() *int32
	SetReasonCode(v string) *GetJobResponseBody
	GetReasonCode() *string
	SetReasonMessage(v string) *GetJobResponseBody
	GetReasonMessage() *string
	SetRequestId(v string) *GetJobResponseBody
	GetRequestId() *string
	SetResourceId(v string) *GetJobResponseBody
	GetResourceId() *string
	SetResourceLevel(v string) *GetJobResponseBody
	GetResourceLevel() *string
	SetResourceType(v string) *GetJobResponseBody
	GetResourceType() *string
	SetRestartRecord(v []*GetJobResponseBodyRestartRecord) *GetJobResponseBody
	GetRestartRecord() []*GetJobResponseBodyRestartRecord
	SetRestartTimes(v string) *GetJobResponseBody
	GetRestartTimes() *string
	SetSettings(v *JobSettings) *GetJobResponseBody
	GetSettings() *JobSettings
	SetStatus(v string) *GetJobResponseBody
	GetStatus() *string
	SetStatusHistory(v []*StatusTransitionItem) *GetJobResponseBody
	GetStatusHistory() []*StatusTransitionItem
	SetSubStatus(v string) *GetJobResponseBody
	GetSubStatus() *string
	SetTenantId(v string) *GetJobResponseBody
	GetTenantId() *string
	SetThirdpartyLibDir(v string) *GetJobResponseBody
	GetThirdpartyLibDir() *string
	SetThirdpartyLibs(v []*string) *GetJobResponseBody
	GetThirdpartyLibs() []*string
	SetUserCommand(v string) *GetJobResponseBody
	GetUserCommand() *string
	SetUserId(v string) *GetJobResponseBody
	GetUserId() *string
	SetUserVpc(v *GetJobResponseBodyUserVpc) *GetJobResponseBody
	GetUserVpc() *GetJobResponseBodyUserVpc
	SetWorkspaceId(v string) *GetJobResponseBody
	GetWorkspaceId() *string
	SetWorkspaceName(v string) *GetJobResponseBody
	GetWorkspaceName() *string
}

type GetJobResponseBody struct {
	// The visibility of the job. Valid values:
	//
	// 	- PUBLIC: The code is public in the workspace.
	//
	// 	- PRIVATE: The workspace is visible only to you and the administrator of the workspace. This is the default value.
	//
	// example:
	//
	// PRIVATE
	Accessibility *string `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// a*****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The code source.
	CodeSource *GetJobResponseBodyCodeSource `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	// The access credential configurations.
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The data sources.
	DataSources []*GetJobResponseBodyDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// The job name.
	//
	// example:
	//
	// tf-mnist-test
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The duration of the job (seconds).
	//
	// example:
	//
	// 3602
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The elastic job parameters.
	ElasticSpec *JobElasticSpec `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
	// Specifies whether to enable the debugger job.
	//
	// example:
	//
	// false
	EnabledDebugger *bool `json:"EnabledDebugger,omitempty" xml:"EnabledDebugger,omitempty"`
	// The configurations of environment variables.
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// The time when the job was created (UTC).
	//
	// example:
	//
	// 2021-01-12T14:35:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The time of the job failed (UTC).
	//
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtFailedTime *string `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	// The time when the job ended (UTC).
	//
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// The start time of the job (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:21Z
	GmtRunningTime *string `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	// The time when the job stopped (UTC).
	//
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtStoppedTime *string `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	// The time when the job was submitted to the cluster (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtSubmittedTime *string `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	// The time when the job succeeded (UTC).
	//
	// example:
	//
	// 2021-01-12T15:36:08Z
	GmtSuccessedTime *string `json:"GmtSuccessedTime,omitempty" xml:"GmtSuccessedTime,omitempty"`
	// The job ID.
	//
	// example:
	//
	// dlc*******
	JobId              *string             `json:"JobId,omitempty" xml:"JobId,omitempty"`
	JobReplicaStatuses []*JobReplicaStatus `json:"JobReplicaStatuses,omitempty" xml:"JobReplicaStatuses,omitempty" type:"Repeated"`
	// The node configuration of the job, which is **JobSpecs*	- in the CreateJob operation.
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// The job type. Specified by the JobType parameter of the [CreateJob](https://help.aliyun.com/document_detail/459672.html) operation.
	//
	// example:
	//
	// TFJob
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// All running nodes of the job.
	Pods []*GetJobResponseBodyPods `json:"Pods,omitempty" xml:"Pods,omitempty" type:"Repeated"`
	// The priority of the job. Valid values: 1 to 9.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The status detail code, which is a sub-status under the current status.
	//
	// example:
	//
	// JobStoppedByUser
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// The description of the status detail code.
	//
	// example:
	//
	// Job is stopped by user.
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// The request ID, which can be used for troubleshooting.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the resource group to which the job belongs.
	//
	// example:
	//
	// r******
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The resource level that the job uses.
	//
	// example:
	//
	// L0
	ResourceLevel *string `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
	// The resource type. Valid values: ECS, Lingjun, and ACS.
	//
	// example:
	//
	// ECS
	ResourceType  *string                            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	RestartRecord []*GetJobResponseBodyRestartRecord `json:"RestartRecord,omitempty" xml:"RestartRecord,omitempty" type:"Repeated"`
	// The number of retries and the maximum number of retries used by the job.
	//
	// example:
	//
	// 0/10
	RestartTimes *string `json:"RestartTimes,omitempty" xml:"RestartTimes,omitempty"`
	// The additional parameter configurations of the job.
	Settings *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// The status of the job. Valid values:
	//
	// 	- Creating
	//
	// 	- Queuing
	//
	// 	- Bidding (Only for Lingjun preemptible jobs)
	//
	// 	- EnvPreparing
	//
	// 	- SanityChecking
	//
	// 	- Running
	//
	// 	- Restarting
	//
	// 	- Stopping
	//
	// 	- SucceededReserving
	//
	// 	- FailedReserving
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// 	- Stopped
	//
	// example:
	//
	// Stopped
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The status history.
	StatusHistory []*StatusTransitionItem `json:"StatusHistory,omitempty" xml:"StatusHistory,omitempty" type:"Repeated"`
	// The sub-status of the job, such as its preemption status.
	//
	// example:
	//
	// Restarting
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// GAR***W134
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// The directory that contains requirements.txt.
	//
	// example:
	//
	// /root/code/
	ThirdpartyLibDir *string `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	// The third-party Python libraries to be installed.
	ThirdpartyLibs []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// The command that is run to start each node.
	//
	// example:
	//
	// python /root/code/mnist.py
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// The UID of the Alibaba Cloud account who submitted the job.
	//
	// example:
	//
	// 12*********
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The VPC of the user.
	UserVpc *GetJobResponseBodyUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// The ID of the workspace to which the job belongs.
	//
	// example:
	//
	// 268
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// The name of the workspace to which the job belongs.
	//
	// example:
	//
	// dlc-workspace
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) GetAccessibility() *string {
	return s.Accessibility
}

func (s *GetJobResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetJobResponseBody) GetCodeSource() *GetJobResponseBodyCodeSource {
	return s.CodeSource
}

func (s *GetJobResponseBody) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *GetJobResponseBody) GetDataSources() []*GetJobResponseBodyDataSources {
	return s.DataSources
}

func (s *GetJobResponseBody) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetJobResponseBody) GetDuration() *int64 {
	return s.Duration
}

func (s *GetJobResponseBody) GetElasticSpec() *JobElasticSpec {
	return s.ElasticSpec
}

func (s *GetJobResponseBody) GetEnabledDebugger() *bool {
	return s.EnabledDebugger
}

func (s *GetJobResponseBody) GetEnvs() map[string]*string {
	return s.Envs
}

func (s *GetJobResponseBody) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetJobResponseBody) GetGmtFailedTime() *string {
	return s.GmtFailedTime
}

func (s *GetJobResponseBody) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *GetJobResponseBody) GetGmtRunningTime() *string {
	return s.GmtRunningTime
}

func (s *GetJobResponseBody) GetGmtStoppedTime() *string {
	return s.GmtStoppedTime
}

func (s *GetJobResponseBody) GetGmtSubmittedTime() *string {
	return s.GmtSubmittedTime
}

func (s *GetJobResponseBody) GetGmtSuccessedTime() *string {
	return s.GmtSuccessedTime
}

func (s *GetJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *GetJobResponseBody) GetJobReplicaStatuses() []*JobReplicaStatus {
	return s.JobReplicaStatuses
}

func (s *GetJobResponseBody) GetJobSpecs() []*JobSpec {
	return s.JobSpecs
}

func (s *GetJobResponseBody) GetJobType() *string {
	return s.JobType
}

func (s *GetJobResponseBody) GetPods() []*GetJobResponseBodyPods {
	return s.Pods
}

func (s *GetJobResponseBody) GetPriority() *int32 {
	return s.Priority
}

func (s *GetJobResponseBody) GetReasonCode() *string {
	return s.ReasonCode
}

func (s *GetJobResponseBody) GetReasonMessage() *string {
	return s.ReasonMessage
}

func (s *GetJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJobResponseBody) GetResourceId() *string {
	return s.ResourceId
}

func (s *GetJobResponseBody) GetResourceLevel() *string {
	return s.ResourceLevel
}

func (s *GetJobResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetJobResponseBody) GetRestartRecord() []*GetJobResponseBodyRestartRecord {
	return s.RestartRecord
}

func (s *GetJobResponseBody) GetRestartTimes() *string {
	return s.RestartTimes
}

func (s *GetJobResponseBody) GetSettings() *JobSettings {
	return s.Settings
}

func (s *GetJobResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetJobResponseBody) GetStatusHistory() []*StatusTransitionItem {
	return s.StatusHistory
}

func (s *GetJobResponseBody) GetSubStatus() *string {
	return s.SubStatus
}

func (s *GetJobResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *GetJobResponseBody) GetThirdpartyLibDir() *string {
	return s.ThirdpartyLibDir
}

func (s *GetJobResponseBody) GetThirdpartyLibs() []*string {
	return s.ThirdpartyLibs
}

func (s *GetJobResponseBody) GetUserCommand() *string {
	return s.UserCommand
}

func (s *GetJobResponseBody) GetUserId() *string {
	return s.UserId
}

func (s *GetJobResponseBody) GetUserVpc() *GetJobResponseBodyUserVpc {
	return s.UserVpc
}

func (s *GetJobResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetJobResponseBody) GetWorkspaceName() *string {
	return s.WorkspaceName
}

func (s *GetJobResponseBody) SetAccessibility(v string) *GetJobResponseBody {
	s.Accessibility = &v
	return s
}

func (s *GetJobResponseBody) SetClusterId(v string) *GetJobResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetJobResponseBody) SetCodeSource(v *GetJobResponseBodyCodeSource) *GetJobResponseBody {
	s.CodeSource = v
	return s
}

func (s *GetJobResponseBody) SetCredentialConfig(v *CredentialConfig) *GetJobResponseBody {
	s.CredentialConfig = v
	return s
}

func (s *GetJobResponseBody) SetDataSources(v []*GetJobResponseBodyDataSources) *GetJobResponseBody {
	s.DataSources = v
	return s
}

func (s *GetJobResponseBody) SetDisplayName(v string) *GetJobResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetJobResponseBody) SetDuration(v int64) *GetJobResponseBody {
	s.Duration = &v
	return s
}

func (s *GetJobResponseBody) SetElasticSpec(v *JobElasticSpec) *GetJobResponseBody {
	s.ElasticSpec = v
	return s
}

func (s *GetJobResponseBody) SetEnabledDebugger(v bool) *GetJobResponseBody {
	s.EnabledDebugger = &v
	return s
}

func (s *GetJobResponseBody) SetEnvs(v map[string]*string) *GetJobResponseBody {
	s.Envs = v
	return s
}

func (s *GetJobResponseBody) SetGmtCreateTime(v string) *GetJobResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtFailedTime(v string) *GetJobResponseBody {
	s.GmtFailedTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtFinishTime(v string) *GetJobResponseBody {
	s.GmtFinishTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtRunningTime(v string) *GetJobResponseBody {
	s.GmtRunningTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtStoppedTime(v string) *GetJobResponseBody {
	s.GmtStoppedTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtSubmittedTime(v string) *GetJobResponseBody {
	s.GmtSubmittedTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtSuccessedTime(v string) *GetJobResponseBody {
	s.GmtSuccessedTime = &v
	return s
}

func (s *GetJobResponseBody) SetJobId(v string) *GetJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBody) SetJobReplicaStatuses(v []*JobReplicaStatus) *GetJobResponseBody {
	s.JobReplicaStatuses = v
	return s
}

func (s *GetJobResponseBody) SetJobSpecs(v []*JobSpec) *GetJobResponseBody {
	s.JobSpecs = v
	return s
}

func (s *GetJobResponseBody) SetJobType(v string) *GetJobResponseBody {
	s.JobType = &v
	return s
}

func (s *GetJobResponseBody) SetPods(v []*GetJobResponseBodyPods) *GetJobResponseBody {
	s.Pods = v
	return s
}

func (s *GetJobResponseBody) SetPriority(v int32) *GetJobResponseBody {
	s.Priority = &v
	return s
}

func (s *GetJobResponseBody) SetReasonCode(v string) *GetJobResponseBody {
	s.ReasonCode = &v
	return s
}

func (s *GetJobResponseBody) SetReasonMessage(v string) *GetJobResponseBody {
	s.ReasonMessage = &v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetResourceId(v string) *GetJobResponseBody {
	s.ResourceId = &v
	return s
}

func (s *GetJobResponseBody) SetResourceLevel(v string) *GetJobResponseBody {
	s.ResourceLevel = &v
	return s
}

func (s *GetJobResponseBody) SetResourceType(v string) *GetJobResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetJobResponseBody) SetRestartRecord(v []*GetJobResponseBodyRestartRecord) *GetJobResponseBody {
	s.RestartRecord = v
	return s
}

func (s *GetJobResponseBody) SetRestartTimes(v string) *GetJobResponseBody {
	s.RestartTimes = &v
	return s
}

func (s *GetJobResponseBody) SetSettings(v *JobSettings) *GetJobResponseBody {
	s.Settings = v
	return s
}

func (s *GetJobResponseBody) SetStatus(v string) *GetJobResponseBody {
	s.Status = &v
	return s
}

func (s *GetJobResponseBody) SetStatusHistory(v []*StatusTransitionItem) *GetJobResponseBody {
	s.StatusHistory = v
	return s
}

func (s *GetJobResponseBody) SetSubStatus(v string) *GetJobResponseBody {
	s.SubStatus = &v
	return s
}

func (s *GetJobResponseBody) SetTenantId(v string) *GetJobResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetJobResponseBody) SetThirdpartyLibDir(v string) *GetJobResponseBody {
	s.ThirdpartyLibDir = &v
	return s
}

func (s *GetJobResponseBody) SetThirdpartyLibs(v []*string) *GetJobResponseBody {
	s.ThirdpartyLibs = v
	return s
}

func (s *GetJobResponseBody) SetUserCommand(v string) *GetJobResponseBody {
	s.UserCommand = &v
	return s
}

func (s *GetJobResponseBody) SetUserId(v string) *GetJobResponseBody {
	s.UserId = &v
	return s
}

func (s *GetJobResponseBody) SetUserVpc(v *GetJobResponseBodyUserVpc) *GetJobResponseBody {
	s.UserVpc = v
	return s
}

func (s *GetJobResponseBody) SetWorkspaceId(v string) *GetJobResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetJobResponseBody) SetWorkspaceName(v string) *GetJobResponseBody {
	s.WorkspaceName = &v
	return s
}

func (s *GetJobResponseBody) Validate() error {
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
	if s.JobReplicaStatuses != nil {
		for _, item := range s.JobReplicaStatuses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	if s.RestartRecord != nil {
		for _, item := range s.RestartRecord {
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

type GetJobResponseBodyCodeSource struct {
	// The code branch.
	//
	// example:
	//
	// master
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// The code source ID.
	//
	// example:
	//
	// code******
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// The code commit ID
	//
	// example:
	//
	// 44da109b59f8596152987eaa8f3b2487xxxxxx
	Commit *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// The local mount path.
	//
	// example:
	//
	// /mnt/data
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s GetJobResponseBodyCodeSource) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyCodeSource) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyCodeSource) GetBranch() *string {
	return s.Branch
}

func (s *GetJobResponseBodyCodeSource) GetCodeSourceId() *string {
	return s.CodeSourceId
}

func (s *GetJobResponseBodyCodeSource) GetCommit() *string {
	return s.Commit
}

func (s *GetJobResponseBodyCodeSource) GetMountPath() *string {
	return s.MountPath
}

func (s *GetJobResponseBodyCodeSource) SetBranch(v string) *GetJobResponseBodyCodeSource {
	s.Branch = &v
	return s
}

func (s *GetJobResponseBodyCodeSource) SetCodeSourceId(v string) *GetJobResponseBodyCodeSource {
	s.CodeSourceId = &v
	return s
}

func (s *GetJobResponseBodyCodeSource) SetCommit(v string) *GetJobResponseBodyCodeSource {
	s.Commit = &v
	return s
}

func (s *GetJobResponseBodyCodeSource) SetMountPath(v string) *GetJobResponseBodyCodeSource {
	s.MountPath = &v
	return s
}

func (s *GetJobResponseBodyCodeSource) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyDataSources struct {
	// The data source ID.
	//
	// example:
	//
	// d*******
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// The local mount path. This parameter is optional. The default value is empty, which specifies that the mount path in the data source is used.
	//
	// example:
	//
	// /mnt/data/
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// The data source URL.
	//
	// example:
	//
	// oss://bucket.oss-cn-hangzhou-internal.aliyuncs.com/path/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s GetJobResponseBodyDataSources) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyDataSources) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyDataSources) GetDataSourceId() *string {
	return s.DataSourceId
}

func (s *GetJobResponseBodyDataSources) GetMountPath() *string {
	return s.MountPath
}

func (s *GetJobResponseBodyDataSources) GetUri() *string {
	return s.Uri
}

func (s *GetJobResponseBodyDataSources) SetDataSourceId(v string) *GetJobResponseBodyDataSources {
	s.DataSourceId = &v
	return s
}

func (s *GetJobResponseBodyDataSources) SetMountPath(v string) *GetJobResponseBodyDataSources {
	s.MountPath = &v
	return s
}

func (s *GetJobResponseBodyDataSources) SetUri(v string) *GetJobResponseBodyDataSources {
	s.Uri = &v
	return s
}

func (s *GetJobResponseBodyDataSources) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyPods struct {
	// The time when the node was created (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The end time of the node (UTC).
	//
	// example:
	//
	// 2021-01-12T15:36:05Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// The start time of the node (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// The historical nodes.
	HistoryPods []*GetJobResponseBodyPodsHistoryPods `json:"HistoryPods,omitempty" xml:"HistoryPods,omitempty" type:"Repeated"`
	// The IP address of the node.
	//
	// example:
	//
	// 10.0.1.2
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The node ID. It can be used in the GetPodLogs and GetPodEvents operations to obtain the detailed logs and events of the node.
	//
	// example:
	//
	// Worker
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// The UID of the node.
	//
	// example:
	//
	// fe846462-af2c-4521-bd6f-96787a57591d
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The resource type of the node.
	//
	// example:
	//
	// Normal
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the node. Valid values:
	//
	// 	- Pending
	//
	// 	- Running
	//
	// 	- Succeeded
	//
	// 	- Failed
	//
	// 	- Unknown
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub-status of the node, such as its preemption status. Valid values:
	//
	// 	- Normal
	//
	// 	- Evicted
	//
	// example:
	//
	// Normal
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// The node type, which corresponds to a specific JobSpec in JobSpecs of the CreateJob operation.
	//
	// example:
	//
	// Worker
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetJobResponseBodyPods) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyPods) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyPods) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetJobResponseBodyPods) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *GetJobResponseBodyPods) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *GetJobResponseBodyPods) GetHistoryPods() []*GetJobResponseBodyPodsHistoryPods {
	return s.HistoryPods
}

func (s *GetJobResponseBodyPods) GetIp() *string {
	return s.Ip
}

func (s *GetJobResponseBodyPods) GetPodId() *string {
	return s.PodId
}

func (s *GetJobResponseBodyPods) GetPodUid() *string {
	return s.PodUid
}

func (s *GetJobResponseBodyPods) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetJobResponseBodyPods) GetStatus() *string {
	return s.Status
}

func (s *GetJobResponseBodyPods) GetSubStatus() *string {
	return s.SubStatus
}

func (s *GetJobResponseBodyPods) GetType() *string {
	return s.Type
}

func (s *GetJobResponseBodyPods) SetGmtCreateTime(v string) *GetJobResponseBodyPods {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobResponseBodyPods) SetGmtFinishTime(v string) *GetJobResponseBodyPods {
	s.GmtFinishTime = &v
	return s
}

func (s *GetJobResponseBodyPods) SetGmtStartTime(v string) *GetJobResponseBodyPods {
	s.GmtStartTime = &v
	return s
}

func (s *GetJobResponseBodyPods) SetHistoryPods(v []*GetJobResponseBodyPodsHistoryPods) *GetJobResponseBodyPods {
	s.HistoryPods = v
	return s
}

func (s *GetJobResponseBodyPods) SetIp(v string) *GetJobResponseBodyPods {
	s.Ip = &v
	return s
}

func (s *GetJobResponseBodyPods) SetPodId(v string) *GetJobResponseBodyPods {
	s.PodId = &v
	return s
}

func (s *GetJobResponseBodyPods) SetPodUid(v string) *GetJobResponseBodyPods {
	s.PodUid = &v
	return s
}

func (s *GetJobResponseBodyPods) SetResourceType(v string) *GetJobResponseBodyPods {
	s.ResourceType = &v
	return s
}

func (s *GetJobResponseBodyPods) SetStatus(v string) *GetJobResponseBodyPods {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyPods) SetSubStatus(v string) *GetJobResponseBodyPods {
	s.SubStatus = &v
	return s
}

func (s *GetJobResponseBodyPods) SetType(v string) *GetJobResponseBodyPods {
	s.Type = &v
	return s
}

func (s *GetJobResponseBodyPods) Validate() error {
	if s.HistoryPods != nil {
		for _, item := range s.HistoryPods {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetJobResponseBodyPodsHistoryPods struct {
	// The time when the node was created (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// The end time of the node (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// The start time of the node (UTC).
	//
	// example:
	//
	// 2021-01-12T14:36:01Z
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// The IP address of the node.
	//
	// example:
	//
	// 10.0.1.3
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// Worker
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// The UID of the node.
	//
	// example:
	//
	// fe846462-af2c-4521-bd6f-96787a57591d
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// The resource type of the node.
	//
	// example:
	//
	// Normal
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The status of the node.
	//
	// example:
	//
	// Failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The sub-status of the node, such as its preemption status. Valid values:
	//
	// 	- Normal
	//
	// 	- Evicted
	//
	// example:
	//
	// Normal
	SubStatus *string `json:"SubStatus,omitempty" xml:"SubStatus,omitempty"`
	// The type of the node.
	//
	// example:
	//
	// Worker
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetJobResponseBodyPodsHistoryPods) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyPodsHistoryPods) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyPodsHistoryPods) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *GetJobResponseBodyPodsHistoryPods) GetGmtFinishTime() *string {
	return s.GmtFinishTime
}

func (s *GetJobResponseBodyPodsHistoryPods) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *GetJobResponseBodyPodsHistoryPods) GetIp() *string {
	return s.Ip
}

func (s *GetJobResponseBodyPodsHistoryPods) GetPodId() *string {
	return s.PodId
}

func (s *GetJobResponseBodyPodsHistoryPods) GetPodUid() *string {
	return s.PodUid
}

func (s *GetJobResponseBodyPodsHistoryPods) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetJobResponseBodyPodsHistoryPods) GetStatus() *string {
	return s.Status
}

func (s *GetJobResponseBodyPodsHistoryPods) GetSubStatus() *string {
	return s.SubStatus
}

func (s *GetJobResponseBodyPodsHistoryPods) GetType() *string {
	return s.Type
}

func (s *GetJobResponseBodyPodsHistoryPods) SetGmtCreateTime(v string) *GetJobResponseBodyPodsHistoryPods {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetGmtFinishTime(v string) *GetJobResponseBodyPodsHistoryPods {
	s.GmtFinishTime = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetGmtStartTime(v string) *GetJobResponseBodyPodsHistoryPods {
	s.GmtStartTime = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetIp(v string) *GetJobResponseBodyPodsHistoryPods {
	s.Ip = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetPodId(v string) *GetJobResponseBodyPodsHistoryPods {
	s.PodId = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetPodUid(v string) *GetJobResponseBodyPodsHistoryPods {
	s.PodUid = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetResourceType(v string) *GetJobResponseBodyPodsHistoryPods {
	s.ResourceType = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetStatus(v string) *GetJobResponseBodyPodsHistoryPods {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetSubStatus(v string) *GetJobResponseBodyPodsHistoryPods {
	s.SubStatus = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetType(v string) *GetJobResponseBodyPodsHistoryPods {
	s.Type = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyRestartRecord struct {
	DetailErrorInfoList  []*GetJobResponseBodyRestartRecordDetailErrorInfoList `json:"DetailErrorInfoList,omitempty" xml:"DetailErrorInfoList,omitempty" type:"Repeated"`
	JobRestartCount      *int64                                                `json:"JobRestartCount,omitempty" xml:"JobRestartCount,omitempty"`
	OccurPhase           *string                                               `json:"OccurPhase,omitempty" xml:"OccurPhase,omitempty"`
	OccurTime            *string                                               `json:"OccurTime,omitempty" xml:"OccurTime,omitempty"`
	Reason               *string                                               `json:"Reason,omitempty" xml:"Reason,omitempty"`
	RestartDurationInSec *int64                                                `json:"RestartDurationInSec,omitempty" xml:"RestartDurationInSec,omitempty"`
	RestartFailReason    *string                                               `json:"RestartFailReason,omitempty" xml:"RestartFailReason,omitempty"`
	RestartStatus        *string                                               `json:"RestartStatus,omitempty" xml:"RestartStatus,omitempty"`
	TriggerID            *string                                               `json:"TriggerID,omitempty" xml:"TriggerID,omitempty"`
}

func (s GetJobResponseBodyRestartRecord) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyRestartRecord) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyRestartRecord) GetDetailErrorInfoList() []*GetJobResponseBodyRestartRecordDetailErrorInfoList {
	return s.DetailErrorInfoList
}

func (s *GetJobResponseBodyRestartRecord) GetJobRestartCount() *int64 {
	return s.JobRestartCount
}

func (s *GetJobResponseBodyRestartRecord) GetOccurPhase() *string {
	return s.OccurPhase
}

func (s *GetJobResponseBodyRestartRecord) GetOccurTime() *string {
	return s.OccurTime
}

func (s *GetJobResponseBodyRestartRecord) GetReason() *string {
	return s.Reason
}

func (s *GetJobResponseBodyRestartRecord) GetRestartDurationInSec() *int64 {
	return s.RestartDurationInSec
}

func (s *GetJobResponseBodyRestartRecord) GetRestartFailReason() *string {
	return s.RestartFailReason
}

func (s *GetJobResponseBodyRestartRecord) GetRestartStatus() *string {
	return s.RestartStatus
}

func (s *GetJobResponseBodyRestartRecord) GetTriggerID() *string {
	return s.TriggerID
}

func (s *GetJobResponseBodyRestartRecord) SetDetailErrorInfoList(v []*GetJobResponseBodyRestartRecordDetailErrorInfoList) *GetJobResponseBodyRestartRecord {
	s.DetailErrorInfoList = v
	return s
}

func (s *GetJobResponseBodyRestartRecord) SetJobRestartCount(v int64) *GetJobResponseBodyRestartRecord {
	s.JobRestartCount = &v
	return s
}

func (s *GetJobResponseBodyRestartRecord) SetOccurPhase(v string) *GetJobResponseBodyRestartRecord {
	s.OccurPhase = &v
	return s
}

func (s *GetJobResponseBodyRestartRecord) SetOccurTime(v string) *GetJobResponseBodyRestartRecord {
	s.OccurTime = &v
	return s
}

func (s *GetJobResponseBodyRestartRecord) SetReason(v string) *GetJobResponseBodyRestartRecord {
	s.Reason = &v
	return s
}

func (s *GetJobResponseBodyRestartRecord) SetRestartDurationInSec(v int64) *GetJobResponseBodyRestartRecord {
	s.RestartDurationInSec = &v
	return s
}

func (s *GetJobResponseBodyRestartRecord) SetRestartFailReason(v string) *GetJobResponseBodyRestartRecord {
	s.RestartFailReason = &v
	return s
}

func (s *GetJobResponseBodyRestartRecord) SetRestartStatus(v string) *GetJobResponseBodyRestartRecord {
	s.RestartStatus = &v
	return s
}

func (s *GetJobResponseBodyRestartRecord) SetTriggerID(v string) *GetJobResponseBodyRestartRecord {
	s.TriggerID = &v
	return s
}

func (s *GetJobResponseBodyRestartRecord) Validate() error {
	if s.DetailErrorInfoList != nil {
		for _, item := range s.DetailErrorInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetJobResponseBodyRestartRecordDetailErrorInfoList struct {
	AddJobLevelBlacklist *bool   `json:"AddJobLevelBlacklist,omitempty" xml:"AddJobLevelBlacklist,omitempty"`
	AddNodeToBlacklist   *bool   `json:"AddNodeToBlacklist,omitempty" xml:"AddNodeToBlacklist,omitempty"`
	DetailErrorMsg       *string `json:"DetailErrorMsg,omitempty" xml:"DetailErrorMsg,omitempty"`
	ErrorCode            *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	ErrorMsg             *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	ErrorSource          *string `json:"ErrorSource,omitempty" xml:"ErrorSource,omitempty"`
	Node                 *string `json:"Node,omitempty" xml:"Node,omitempty"`
	Pod                  *string `json:"Pod,omitempty" xml:"Pod,omitempty"`
	TriggerRestart       *bool   `json:"TriggerRestart,omitempty" xml:"TriggerRestart,omitempty"`
}

func (s GetJobResponseBodyRestartRecordDetailErrorInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyRestartRecordDetailErrorInfoList) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) GetAddJobLevelBlacklist() *bool {
	return s.AddJobLevelBlacklist
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) GetAddNodeToBlacklist() *bool {
	return s.AddNodeToBlacklist
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) GetDetailErrorMsg() *string {
	return s.DetailErrorMsg
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) GetErrorSource() *string {
	return s.ErrorSource
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) GetNode() *string {
	return s.Node
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) GetPod() *string {
	return s.Pod
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) GetTriggerRestart() *bool {
	return s.TriggerRestart
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) SetAddJobLevelBlacklist(v bool) *GetJobResponseBodyRestartRecordDetailErrorInfoList {
	s.AddJobLevelBlacklist = &v
	return s
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) SetAddNodeToBlacklist(v bool) *GetJobResponseBodyRestartRecordDetailErrorInfoList {
	s.AddNodeToBlacklist = &v
	return s
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) SetDetailErrorMsg(v string) *GetJobResponseBodyRestartRecordDetailErrorInfoList {
	s.DetailErrorMsg = &v
	return s
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) SetErrorCode(v string) *GetJobResponseBodyRestartRecordDetailErrorInfoList {
	s.ErrorCode = &v
	return s
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) SetErrorMsg(v string) *GetJobResponseBodyRestartRecordDetailErrorInfoList {
	s.ErrorMsg = &v
	return s
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) SetErrorSource(v string) *GetJobResponseBodyRestartRecordDetailErrorInfoList {
	s.ErrorSource = &v
	return s
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) SetNode(v string) *GetJobResponseBodyRestartRecordDetailErrorInfoList {
	s.Node = &v
	return s
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) SetPod(v string) *GetJobResponseBodyRestartRecordDetailErrorInfoList {
	s.Pod = &v
	return s
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) SetTriggerRestart(v bool) *GetJobResponseBodyRestartRecordDetailErrorInfoList {
	s.TriggerRestart = &v
	return s
}

func (s *GetJobResponseBodyRestartRecordDetailErrorInfoList) Validate() error {
	return dara.Validate(s)
}

type GetJobResponseBodyUserVpc struct {
	// The default router. This parameter is valid only for general-purpose computing resources. Valid values:
	//
	// eth0: The default network interface is used to access the Internet through the public gateway. eth1: The user\\"s Elastic Network Interface is used to access the Internet through the private gateway.
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// The extended CIDR block. Example: 192.168.0.1/24.
	ExtendedCidrs []*string `json:"ExtendedCidrs,omitempty" xml:"ExtendedCidrs,omitempty" type:"Repeated"`
	// The security group ID.
	//
	// example:
	//
	// sg-abcdef****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vs-abcdef****
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-abcdef****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetJobResponseBodyUserVpc) String() string {
	return dara.Prettify(s)
}

func (s GetJobResponseBodyUserVpc) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyUserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *GetJobResponseBodyUserVpc) GetExtendedCidrs() []*string {
	return s.ExtendedCidrs
}

func (s *GetJobResponseBodyUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetJobResponseBodyUserVpc) GetSwitchId() *string {
	return s.SwitchId
}

func (s *GetJobResponseBodyUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *GetJobResponseBodyUserVpc) SetDefaultRoute(v string) *GetJobResponseBodyUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *GetJobResponseBodyUserVpc) SetExtendedCidrs(v []*string) *GetJobResponseBodyUserVpc {
	s.ExtendedCidrs = v
	return s
}

func (s *GetJobResponseBodyUserVpc) SetSecurityGroupId(v string) *GetJobResponseBodyUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *GetJobResponseBodyUserVpc) SetSwitchId(v string) *GetJobResponseBodyUserVpc {
	s.SwitchId = &v
	return s
}

func (s *GetJobResponseBodyUserVpc) SetVpcId(v string) *GetJobResponseBodyUserVpc {
	s.VpcId = &v
	return s
}

func (s *GetJobResponseBodyUserVpc) Validate() error {
	return dara.Validate(s)
}
