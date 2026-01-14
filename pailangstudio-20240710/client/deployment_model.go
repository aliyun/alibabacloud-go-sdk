// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeployment interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Deployment
	GetAccessibility() *string
	SetAutoApproval(v bool) *Deployment
	GetAutoApproval() *bool
	SetChatHistoryConfig(v *DeploymentChatHistoryConfig) *Deployment
	GetChatHistoryConfig() *DeploymentChatHistoryConfig
	SetContentModerationConfig(v *DeploymentContentModerationConfig) *Deployment
	GetContentModerationConfig() *DeploymentContentModerationConfig
	SetCreator(v string) *Deployment
	GetCreator() *string
	SetCredentialConfig(v *DeploymentCredentialConfig) *Deployment
	GetCredentialConfig() *DeploymentCredentialConfig
	SetDataSources(v []*DeploymentDataSources) *Deployment
	GetDataSources() []*DeploymentDataSources
	SetDeploymentConfig(v string) *Deployment
	GetDeploymentConfig() *string
	SetDeploymentId(v string) *Deployment
	GetDeploymentId() *string
	SetDeploymentStages(v []*DeploymentDeploymentStages) *Deployment
	GetDeploymentStages() []*DeploymentDeploymentStages
	SetDeploymentStatus(v string) *Deployment
	GetDeploymentStatus() *string
	SetDescription(v string) *Deployment
	GetDescription() *string
	SetEcsSpec(v *DeploymentEcsSpec) *Deployment
	GetEcsSpec() *DeploymentEcsSpec
	SetEnableTrace(v bool) *Deployment
	GetEnableTrace() *bool
	SetEnvs(v []*DeploymentEnvs) *Deployment
	GetEnvs() []*DeploymentEnvs
	SetErrorMessage(v string) *Deployment
	GetErrorMessage() *string
	SetGmtCreateTime(v string) *Deployment
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *Deployment
	GetGmtModifiedTime() *string
	SetLabels(v []*DeploymentLabels) *Deployment
	GetLabels() []*DeploymentLabels
	SetOperationType(v string) *Deployment
	GetOperationType() *string
	SetResourceId(v string) *Deployment
	GetResourceId() *string
	SetResourceSnapshotId(v string) *Deployment
	GetResourceSnapshotId() *string
	SetResourceType(v string) *Deployment
	GetResourceType() *string
	SetServiceGroup(v string) *Deployment
	GetServiceGroup() *string
	SetServiceName(v string) *Deployment
	GetServiceName() *string
	SetUserVpc(v *DeploymentUserVpc) *Deployment
	GetUserVpc() *DeploymentUserVpc
	SetWorkDir(v string) *Deployment
	GetWorkDir() *string
	SetWorkspaceId(v string) *Deployment
	GetWorkspaceId() *string
}

type Deployment struct {
	Accessibility           *string                            `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	AutoApproval            *bool                              `json:"AutoApproval,omitempty" xml:"AutoApproval,omitempty"`
	ChatHistoryConfig       *DeploymentChatHistoryConfig       `json:"ChatHistoryConfig,omitempty" xml:"ChatHistoryConfig,omitempty" type:"Struct"`
	ContentModerationConfig *DeploymentContentModerationConfig `json:"ContentModerationConfig,omitempty" xml:"ContentModerationConfig,omitempty" type:"Struct"`
	Creator                 *string                            `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CredentialConfig        *DeploymentCredentialConfig        `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty" type:"Struct"`
	DataSources             []*DeploymentDataSources           `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	DeploymentConfig        *string                            `json:"DeploymentConfig,omitempty" xml:"DeploymentConfig,omitempty"`
	DeploymentId            *string                            `json:"DeploymentId,omitempty" xml:"DeploymentId,omitempty"`
	DeploymentStages        []*DeploymentDeploymentStages      `json:"DeploymentStages,omitempty" xml:"DeploymentStages,omitempty" type:"Repeated"`
	DeploymentStatus        *string                            `json:"DeploymentStatus,omitempty" xml:"DeploymentStatus,omitempty"`
	Description             *string                            `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsSpec                 *DeploymentEcsSpec                 `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty" type:"Struct"`
	EnableTrace             *bool                              `json:"EnableTrace,omitempty" xml:"EnableTrace,omitempty"`
	Envs                    []*DeploymentEnvs                  `json:"Envs,omitempty" xml:"Envs,omitempty" type:"Repeated"`
	ErrorMessage            *string                            `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	GmtCreateTime           *string                            `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime         *string                            `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	Labels                  []*DeploymentLabels                `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	OperationType           *string                            `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
	ResourceId              *string                            `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceSnapshotId      *string                            `json:"ResourceSnapshotId,omitempty" xml:"ResourceSnapshotId,omitempty"`
	ResourceType            *string                            `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	ServiceGroup            *string                            `json:"ServiceGroup,omitempty" xml:"ServiceGroup,omitempty"`
	ServiceName             *string                            `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	UserVpc                 *DeploymentUserVpc                 `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	WorkDir                 *string                            `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
	WorkspaceId             *string                            `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Deployment) String() string {
	return dara.Prettify(s)
}

func (s Deployment) GoString() string {
	return s.String()
}

func (s *Deployment) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Deployment) GetAutoApproval() *bool {
	return s.AutoApproval
}

func (s *Deployment) GetChatHistoryConfig() *DeploymentChatHistoryConfig {
	return s.ChatHistoryConfig
}

func (s *Deployment) GetContentModerationConfig() *DeploymentContentModerationConfig {
	return s.ContentModerationConfig
}

func (s *Deployment) GetCreator() *string {
	return s.Creator
}

func (s *Deployment) GetCredentialConfig() *DeploymentCredentialConfig {
	return s.CredentialConfig
}

func (s *Deployment) GetDataSources() []*DeploymentDataSources {
	return s.DataSources
}

func (s *Deployment) GetDeploymentConfig() *string {
	return s.DeploymentConfig
}

func (s *Deployment) GetDeploymentId() *string {
	return s.DeploymentId
}

func (s *Deployment) GetDeploymentStages() []*DeploymentDeploymentStages {
	return s.DeploymentStages
}

func (s *Deployment) GetDeploymentStatus() *string {
	return s.DeploymentStatus
}

func (s *Deployment) GetDescription() *string {
	return s.Description
}

func (s *Deployment) GetEcsSpec() *DeploymentEcsSpec {
	return s.EcsSpec
}

func (s *Deployment) GetEnableTrace() *bool {
	return s.EnableTrace
}

func (s *Deployment) GetEnvs() []*DeploymentEnvs {
	return s.Envs
}

func (s *Deployment) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *Deployment) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Deployment) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Deployment) GetLabels() []*DeploymentLabels {
	return s.Labels
}

func (s *Deployment) GetOperationType() *string {
	return s.OperationType
}

func (s *Deployment) GetResourceId() *string {
	return s.ResourceId
}

func (s *Deployment) GetResourceSnapshotId() *string {
	return s.ResourceSnapshotId
}

func (s *Deployment) GetResourceType() *string {
	return s.ResourceType
}

func (s *Deployment) GetServiceGroup() *string {
	return s.ServiceGroup
}

func (s *Deployment) GetServiceName() *string {
	return s.ServiceName
}

func (s *Deployment) GetUserVpc() *DeploymentUserVpc {
	return s.UserVpc
}

func (s *Deployment) GetWorkDir() *string {
	return s.WorkDir
}

func (s *Deployment) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Deployment) SetAccessibility(v string) *Deployment {
	s.Accessibility = &v
	return s
}

func (s *Deployment) SetAutoApproval(v bool) *Deployment {
	s.AutoApproval = &v
	return s
}

func (s *Deployment) SetChatHistoryConfig(v *DeploymentChatHistoryConfig) *Deployment {
	s.ChatHistoryConfig = v
	return s
}

func (s *Deployment) SetContentModerationConfig(v *DeploymentContentModerationConfig) *Deployment {
	s.ContentModerationConfig = v
	return s
}

func (s *Deployment) SetCreator(v string) *Deployment {
	s.Creator = &v
	return s
}

func (s *Deployment) SetCredentialConfig(v *DeploymentCredentialConfig) *Deployment {
	s.CredentialConfig = v
	return s
}

func (s *Deployment) SetDataSources(v []*DeploymentDataSources) *Deployment {
	s.DataSources = v
	return s
}

func (s *Deployment) SetDeploymentConfig(v string) *Deployment {
	s.DeploymentConfig = &v
	return s
}

func (s *Deployment) SetDeploymentId(v string) *Deployment {
	s.DeploymentId = &v
	return s
}

func (s *Deployment) SetDeploymentStages(v []*DeploymentDeploymentStages) *Deployment {
	s.DeploymentStages = v
	return s
}

func (s *Deployment) SetDeploymentStatus(v string) *Deployment {
	s.DeploymentStatus = &v
	return s
}

func (s *Deployment) SetDescription(v string) *Deployment {
	s.Description = &v
	return s
}

func (s *Deployment) SetEcsSpec(v *DeploymentEcsSpec) *Deployment {
	s.EcsSpec = v
	return s
}

func (s *Deployment) SetEnableTrace(v bool) *Deployment {
	s.EnableTrace = &v
	return s
}

func (s *Deployment) SetEnvs(v []*DeploymentEnvs) *Deployment {
	s.Envs = v
	return s
}

func (s *Deployment) SetErrorMessage(v string) *Deployment {
	s.ErrorMessage = &v
	return s
}

func (s *Deployment) SetGmtCreateTime(v string) *Deployment {
	s.GmtCreateTime = &v
	return s
}

func (s *Deployment) SetGmtModifiedTime(v string) *Deployment {
	s.GmtModifiedTime = &v
	return s
}

func (s *Deployment) SetLabels(v []*DeploymentLabels) *Deployment {
	s.Labels = v
	return s
}

func (s *Deployment) SetOperationType(v string) *Deployment {
	s.OperationType = &v
	return s
}

func (s *Deployment) SetResourceId(v string) *Deployment {
	s.ResourceId = &v
	return s
}

func (s *Deployment) SetResourceSnapshotId(v string) *Deployment {
	s.ResourceSnapshotId = &v
	return s
}

func (s *Deployment) SetResourceType(v string) *Deployment {
	s.ResourceType = &v
	return s
}

func (s *Deployment) SetServiceGroup(v string) *Deployment {
	s.ServiceGroup = &v
	return s
}

func (s *Deployment) SetServiceName(v string) *Deployment {
	s.ServiceName = &v
	return s
}

func (s *Deployment) SetUserVpc(v *DeploymentUserVpc) *Deployment {
	s.UserVpc = v
	return s
}

func (s *Deployment) SetWorkDir(v string) *Deployment {
	s.WorkDir = &v
	return s
}

func (s *Deployment) SetWorkspaceId(v string) *Deployment {
	s.WorkspaceId = &v
	return s
}

func (s *Deployment) Validate() error {
	if s.ChatHistoryConfig != nil {
		if err := s.ChatHistoryConfig.Validate(); err != nil {
			return err
		}
	}
	if s.ContentModerationConfig != nil {
		if err := s.ContentModerationConfig.Validate(); err != nil {
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
	if s.DeploymentStages != nil {
		for _, item := range s.DeploymentStages {
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
	if s.Labels != nil {
		for _, item := range s.Labels {
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

type DeploymentChatHistoryConfig struct {
	// 连接名称
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// 存储类型
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s DeploymentChatHistoryConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploymentChatHistoryConfig) GoString() string {
	return s.String()
}

func (s *DeploymentChatHistoryConfig) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *DeploymentChatHistoryConfig) GetStorageType() *string {
	return s.StorageType
}

func (s *DeploymentChatHistoryConfig) SetConnectionName(v string) *DeploymentChatHistoryConfig {
	s.ConnectionName = &v
	return s
}

func (s *DeploymentChatHistoryConfig) SetStorageType(v string) *DeploymentChatHistoryConfig {
	s.StorageType = &v
	return s
}

func (s *DeploymentChatHistoryConfig) Validate() error {
	return dara.Validate(s)
}

type DeploymentContentModerationConfig struct {
	// 启用输入内容审查
	EnableInputModeration *bool `json:"EnableInputModeration,omitempty" xml:"EnableInputModeration,omitempty"`
	// 启用输出内容审查
	EnableOutputModeration *bool `json:"EnableOutputModeration,omitempty" xml:"EnableOutputModeration,omitempty"`
	// 流式输出内容送审缓存大小
	StreamingModerationThreshold *int32 `json:"StreamingModerationThreshold,omitempty" xml:"StreamingModerationThreshold,omitempty"`
}

func (s DeploymentContentModerationConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploymentContentModerationConfig) GoString() string {
	return s.String()
}

func (s *DeploymentContentModerationConfig) GetEnableInputModeration() *bool {
	return s.EnableInputModeration
}

func (s *DeploymentContentModerationConfig) GetEnableOutputModeration() *bool {
	return s.EnableOutputModeration
}

func (s *DeploymentContentModerationConfig) GetStreamingModerationThreshold() *int32 {
	return s.StreamingModerationThreshold
}

func (s *DeploymentContentModerationConfig) SetEnableInputModeration(v bool) *DeploymentContentModerationConfig {
	s.EnableInputModeration = &v
	return s
}

func (s *DeploymentContentModerationConfig) SetEnableOutputModeration(v bool) *DeploymentContentModerationConfig {
	s.EnableOutputModeration = &v
	return s
}

func (s *DeploymentContentModerationConfig) SetStreamingModerationThreshold(v int32) *DeploymentContentModerationConfig {
	s.StreamingModerationThreshold = &v
	return s
}

func (s *DeploymentContentModerationConfig) Validate() error {
	return dara.Validate(s)
}

type DeploymentCredentialConfig struct {
	// AliyunEnvRoleKey
	AliyunEnvRoleKey *string `json:"AliyunEnvRoleKey,omitempty" xml:"AliyunEnvRoleKey,omitempty"`
	// Credential配置项列表
	CredentialConfigItems []*DeploymentCredentialConfigCredentialConfigItems `json:"CredentialConfigItems,omitempty" xml:"CredentialConfigItems,omitempty" type:"Repeated"`
	// 是否启用Credential注入
	EnableCredentialInject *bool `json:"EnableCredentialInject,omitempty" xml:"EnableCredentialInject,omitempty"`
}

func (s DeploymentCredentialConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploymentCredentialConfig) GoString() string {
	return s.String()
}

func (s *DeploymentCredentialConfig) GetAliyunEnvRoleKey() *string {
	return s.AliyunEnvRoleKey
}

func (s *DeploymentCredentialConfig) GetCredentialConfigItems() []*DeploymentCredentialConfigCredentialConfigItems {
	return s.CredentialConfigItems
}

func (s *DeploymentCredentialConfig) GetEnableCredentialInject() *bool {
	return s.EnableCredentialInject
}

func (s *DeploymentCredentialConfig) SetAliyunEnvRoleKey(v string) *DeploymentCredentialConfig {
	s.AliyunEnvRoleKey = &v
	return s
}

func (s *DeploymentCredentialConfig) SetCredentialConfigItems(v []*DeploymentCredentialConfigCredentialConfigItems) *DeploymentCredentialConfig {
	s.CredentialConfigItems = v
	return s
}

func (s *DeploymentCredentialConfig) SetEnableCredentialInject(v bool) *DeploymentCredentialConfig {
	s.EnableCredentialInject = &v
	return s
}

func (s *DeploymentCredentialConfig) Validate() error {
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

type DeploymentCredentialConfigCredentialConfigItems struct {
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 角色列表
	Roles []*DeploymentCredentialConfigCredentialConfigItemsRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// Type
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeploymentCredentialConfigCredentialConfigItems) String() string {
	return dara.Prettify(s)
}

func (s DeploymentCredentialConfigCredentialConfigItems) GoString() string {
	return s.String()
}

func (s *DeploymentCredentialConfigCredentialConfigItems) GetKey() *string {
	return s.Key
}

func (s *DeploymentCredentialConfigCredentialConfigItems) GetRoles() []*DeploymentCredentialConfigCredentialConfigItemsRoles {
	return s.Roles
}

func (s *DeploymentCredentialConfigCredentialConfigItems) GetType() *string {
	return s.Type
}

func (s *DeploymentCredentialConfigCredentialConfigItems) SetKey(v string) *DeploymentCredentialConfigCredentialConfigItems {
	s.Key = &v
	return s
}

func (s *DeploymentCredentialConfigCredentialConfigItems) SetRoles(v []*DeploymentCredentialConfigCredentialConfigItemsRoles) *DeploymentCredentialConfigCredentialConfigItems {
	s.Roles = v
	return s
}

func (s *DeploymentCredentialConfigCredentialConfigItems) SetType(v string) *DeploymentCredentialConfigCredentialConfigItems {
	s.Type = &v
	return s
}

func (s *DeploymentCredentialConfigCredentialConfigItems) Validate() error {
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

type DeploymentCredentialConfigCredentialConfigItemsRoles struct {
	// AssumeRoleFor
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// 角色名称
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// 角色类型
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s DeploymentCredentialConfigCredentialConfigItemsRoles) String() string {
	return dara.Prettify(s)
}

func (s DeploymentCredentialConfigCredentialConfigItemsRoles) GoString() string {
	return s.String()
}

func (s *DeploymentCredentialConfigCredentialConfigItemsRoles) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *DeploymentCredentialConfigCredentialConfigItemsRoles) GetRoleArn() *string {
	return s.RoleArn
}

func (s *DeploymentCredentialConfigCredentialConfigItemsRoles) GetRoleType() *string {
	return s.RoleType
}

func (s *DeploymentCredentialConfigCredentialConfigItemsRoles) SetAssumeRoleFor(v string) *DeploymentCredentialConfigCredentialConfigItemsRoles {
	s.AssumeRoleFor = &v
	return s
}

func (s *DeploymentCredentialConfigCredentialConfigItemsRoles) SetRoleArn(v string) *DeploymentCredentialConfigCredentialConfigItemsRoles {
	s.RoleArn = &v
	return s
}

func (s *DeploymentCredentialConfigCredentialConfigItemsRoles) SetRoleType(v string) *DeploymentCredentialConfigCredentialConfigItemsRoles {
	s.RoleType = &v
	return s
}

func (s *DeploymentCredentialConfigCredentialConfigItemsRoles) Validate() error {
	return dara.Validate(s)
}

type DeploymentDataSources struct {
	// 数据集ID
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 统一资源识别码
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DeploymentDataSources) String() string {
	return dara.Prettify(s)
}

func (s DeploymentDataSources) GoString() string {
	return s.String()
}

func (s *DeploymentDataSources) GetDatasetId() *string {
	return s.DatasetId
}

func (s *DeploymentDataSources) GetMountPath() *string {
	return s.MountPath
}

func (s *DeploymentDataSources) GetUri() *string {
	return s.Uri
}

func (s *DeploymentDataSources) SetDatasetId(v string) *DeploymentDataSources {
	s.DatasetId = &v
	return s
}

func (s *DeploymentDataSources) SetMountPath(v string) *DeploymentDataSources {
	s.MountPath = &v
	return s
}

func (s *DeploymentDataSources) SetUri(v string) *DeploymentDataSources {
	s.Uri = &v
	return s
}

func (s *DeploymentDataSources) Validate() error {
	return dara.Validate(s)
}

type DeploymentDeploymentStages struct {
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 错误信息
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// 结束时间
	GmtEndTime *string `json:"GmtEndTime,omitempty" xml:"GmtEndTime,omitempty"`
	// 开始时间
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// 阶段
	Stage *int32 `json:"Stage,omitempty" xml:"Stage,omitempty"`
	// 阶段信息
	StageInfo *string `json:"StageInfo,omitempty" xml:"StageInfo,omitempty"`
	// 阶段名称
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
	// 阶段状态
	StageStatus *string `json:"StageStatus,omitempty" xml:"StageStatus,omitempty"`
}

func (s DeploymentDeploymentStages) String() string {
	return dara.Prettify(s)
}

func (s DeploymentDeploymentStages) GoString() string {
	return s.String()
}

func (s *DeploymentDeploymentStages) GetDescription() *string {
	return s.Description
}

func (s *DeploymentDeploymentStages) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeploymentDeploymentStages) GetGmtEndTime() *string {
	return s.GmtEndTime
}

func (s *DeploymentDeploymentStages) GetGmtStartTime() *string {
	return s.GmtStartTime
}

func (s *DeploymentDeploymentStages) GetStage() *int32 {
	return s.Stage
}

func (s *DeploymentDeploymentStages) GetStageInfo() *string {
	return s.StageInfo
}

func (s *DeploymentDeploymentStages) GetStageName() *string {
	return s.StageName
}

func (s *DeploymentDeploymentStages) GetStageStatus() *string {
	return s.StageStatus
}

func (s *DeploymentDeploymentStages) SetDescription(v string) *DeploymentDeploymentStages {
	s.Description = &v
	return s
}

func (s *DeploymentDeploymentStages) SetErrorMessage(v string) *DeploymentDeploymentStages {
	s.ErrorMessage = &v
	return s
}

func (s *DeploymentDeploymentStages) SetGmtEndTime(v string) *DeploymentDeploymentStages {
	s.GmtEndTime = &v
	return s
}

func (s *DeploymentDeploymentStages) SetGmtStartTime(v string) *DeploymentDeploymentStages {
	s.GmtStartTime = &v
	return s
}

func (s *DeploymentDeploymentStages) SetStage(v int32) *DeploymentDeploymentStages {
	s.Stage = &v
	return s
}

func (s *DeploymentDeploymentStages) SetStageInfo(v string) *DeploymentDeploymentStages {
	s.StageInfo = &v
	return s
}

func (s *DeploymentDeploymentStages) SetStageName(v string) *DeploymentDeploymentStages {
	s.StageName = &v
	return s
}

func (s *DeploymentDeploymentStages) SetStageStatus(v string) *DeploymentDeploymentStages {
	s.StageStatus = &v
	return s
}

func (s *DeploymentDeploymentStages) Validate() error {
	return dara.Validate(s)
}

type DeploymentEcsSpec struct {
	// CPU信息
	CPU *int32 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// 竞价资源配置
	ComputingInstanceConfig *DeploymentEcsSpecComputingInstanceConfig `json:"ComputingInstanceConfig,omitempty" xml:"ComputingInstanceConfig,omitempty" type:"Struct"`
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
	// 是否启用弹性资源池
	ResourceBurstable *bool `json:"ResourceBurstable,omitempty" xml:"ResourceBurstable,omitempty"`
	// 资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 共享内存
	SharedMemory *int32 `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
}

func (s DeploymentEcsSpec) String() string {
	return dara.Prettify(s)
}

func (s DeploymentEcsSpec) GoString() string {
	return s.String()
}

func (s *DeploymentEcsSpec) GetCPU() *int32 {
	return s.CPU
}

func (s *DeploymentEcsSpec) GetComputingInstanceConfig() *DeploymentEcsSpecComputingInstanceConfig {
	return s.ComputingInstanceConfig
}

func (s *DeploymentEcsSpec) GetExtraEphemeralStorage() *int32 {
	return s.ExtraEphemeralStorage
}

func (s *DeploymentEcsSpec) GetGPU() *int32 {
	return s.GPU
}

func (s *DeploymentEcsSpec) GetGPUCorePercentage() *int32 {
	return s.GPUCorePercentage
}

func (s *DeploymentEcsSpec) GetGPUMemory() *int32 {
	return s.GPUMemory
}

func (s *DeploymentEcsSpec) GetGPUType() *string {
	return s.GPUType
}

func (s *DeploymentEcsSpec) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DeploymentEcsSpec) GetMemory() *int32 {
	return s.Memory
}

func (s *DeploymentEcsSpec) GetPodCount() *int32 {
	return s.PodCount
}

func (s *DeploymentEcsSpec) GetQuotaId() *string {
	return s.QuotaId
}

func (s *DeploymentEcsSpec) GetQuotaType() *string {
	return s.QuotaType
}

func (s *DeploymentEcsSpec) GetResourceBurstable() *bool {
	return s.ResourceBurstable
}

func (s *DeploymentEcsSpec) GetResourceId() *string {
	return s.ResourceId
}

func (s *DeploymentEcsSpec) GetSharedMemory() *int32 {
	return s.SharedMemory
}

func (s *DeploymentEcsSpec) SetCPU(v int32) *DeploymentEcsSpec {
	s.CPU = &v
	return s
}

func (s *DeploymentEcsSpec) SetComputingInstanceConfig(v *DeploymentEcsSpecComputingInstanceConfig) *DeploymentEcsSpec {
	s.ComputingInstanceConfig = v
	return s
}

func (s *DeploymentEcsSpec) SetExtraEphemeralStorage(v int32) *DeploymentEcsSpec {
	s.ExtraEphemeralStorage = &v
	return s
}

func (s *DeploymentEcsSpec) SetGPU(v int32) *DeploymentEcsSpec {
	s.GPU = &v
	return s
}

func (s *DeploymentEcsSpec) SetGPUCorePercentage(v int32) *DeploymentEcsSpec {
	s.GPUCorePercentage = &v
	return s
}

func (s *DeploymentEcsSpec) SetGPUMemory(v int32) *DeploymentEcsSpec {
	s.GPUMemory = &v
	return s
}

func (s *DeploymentEcsSpec) SetGPUType(v string) *DeploymentEcsSpec {
	s.GPUType = &v
	return s
}

func (s *DeploymentEcsSpec) SetInstanceType(v string) *DeploymentEcsSpec {
	s.InstanceType = &v
	return s
}

func (s *DeploymentEcsSpec) SetMemory(v int32) *DeploymentEcsSpec {
	s.Memory = &v
	return s
}

func (s *DeploymentEcsSpec) SetPodCount(v int32) *DeploymentEcsSpec {
	s.PodCount = &v
	return s
}

func (s *DeploymentEcsSpec) SetQuotaId(v string) *DeploymentEcsSpec {
	s.QuotaId = &v
	return s
}

func (s *DeploymentEcsSpec) SetQuotaType(v string) *DeploymentEcsSpec {
	s.QuotaType = &v
	return s
}

func (s *DeploymentEcsSpec) SetResourceBurstable(v bool) *DeploymentEcsSpec {
	s.ResourceBurstable = &v
	return s
}

func (s *DeploymentEcsSpec) SetResourceId(v string) *DeploymentEcsSpec {
	s.ResourceId = &v
	return s
}

func (s *DeploymentEcsSpec) SetSharedMemory(v int32) *DeploymentEcsSpec {
	s.SharedMemory = &v
	return s
}

func (s *DeploymentEcsSpec) Validate() error {
	if s.ComputingInstanceConfig != nil {
		if err := s.ComputingInstanceConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeploymentEcsSpecComputingInstanceConfig struct {
	// 机器资源配置
	ComputingInstances []*DeploymentEcsSpecComputingInstanceConfigComputingInstances `json:"ComputingInstances,omitempty" xml:"ComputingInstances,omitempty" type:"Repeated"`
	// 是否启用竞价保留时长
	DisableSpotProtectionPeriod *bool `json:"DisableSpotProtectionPeriod,omitempty" xml:"DisableSpotProtectionPeriod,omitempty"`
}

func (s DeploymentEcsSpecComputingInstanceConfig) String() string {
	return dara.Prettify(s)
}

func (s DeploymentEcsSpecComputingInstanceConfig) GoString() string {
	return s.String()
}

func (s *DeploymentEcsSpecComputingInstanceConfig) GetComputingInstances() []*DeploymentEcsSpecComputingInstanceConfigComputingInstances {
	return s.ComputingInstances
}

func (s *DeploymentEcsSpecComputingInstanceConfig) GetDisableSpotProtectionPeriod() *bool {
	return s.DisableSpotProtectionPeriod
}

func (s *DeploymentEcsSpecComputingInstanceConfig) SetComputingInstances(v []*DeploymentEcsSpecComputingInstanceConfigComputingInstances) *DeploymentEcsSpecComputingInstanceConfig {
	s.ComputingInstances = v
	return s
}

func (s *DeploymentEcsSpecComputingInstanceConfig) SetDisableSpotProtectionPeriod(v bool) *DeploymentEcsSpecComputingInstanceConfig {
	s.DisableSpotProtectionPeriod = &v
	return s
}

func (s *DeploymentEcsSpecComputingInstanceConfig) Validate() error {
	if s.ComputingInstances != nil {
		for _, item := range s.ComputingInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeploymentEcsSpecComputingInstanceConfigComputingInstances struct {
	// 竞价出价
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// 机型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DeploymentEcsSpecComputingInstanceConfigComputingInstances) String() string {
	return dara.Prettify(s)
}

func (s DeploymentEcsSpecComputingInstanceConfigComputingInstances) GoString() string {
	return s.String()
}

func (s *DeploymentEcsSpecComputingInstanceConfigComputingInstances) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *DeploymentEcsSpecComputingInstanceConfigComputingInstances) GetType() *string {
	return s.Type
}

func (s *DeploymentEcsSpecComputingInstanceConfigComputingInstances) SetSpotPriceLimit(v float32) *DeploymentEcsSpecComputingInstanceConfigComputingInstances {
	s.SpotPriceLimit = &v
	return s
}

func (s *DeploymentEcsSpecComputingInstanceConfigComputingInstances) SetType(v string) *DeploymentEcsSpecComputingInstanceConfigComputingInstances {
	s.Type = &v
	return s
}

func (s *DeploymentEcsSpecComputingInstanceConfigComputingInstances) Validate() error {
	return dara.Validate(s)
}

type DeploymentEnvs struct {
	// 环境键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 环境值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeploymentEnvs) String() string {
	return dara.Prettify(s)
}

func (s DeploymentEnvs) GoString() string {
	return s.String()
}

func (s *DeploymentEnvs) GetKey() *string {
	return s.Key
}

func (s *DeploymentEnvs) GetValue() *string {
	return s.Value
}

func (s *DeploymentEnvs) SetKey(v string) *DeploymentEnvs {
	s.Key = &v
	return s
}

func (s *DeploymentEnvs) SetValue(v string) *DeploymentEnvs {
	s.Value = &v
	return s
}

func (s *DeploymentEnvs) Validate() error {
	return dara.Validate(s)
}

type DeploymentLabels struct {
	// 标签键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DeploymentLabels) String() string {
	return dara.Prettify(s)
}

func (s DeploymentLabels) GoString() string {
	return s.String()
}

func (s *DeploymentLabels) GetKey() *string {
	return s.Key
}

func (s *DeploymentLabels) GetValue() *string {
	return s.Value
}

func (s *DeploymentLabels) SetKey(v string) *DeploymentLabels {
	s.Key = &v
	return s
}

func (s *DeploymentLabels) SetValue(v string) *DeploymentLabels {
	s.Value = &v
	return s
}

func (s *DeploymentLabels) Validate() error {
	return dara.Validate(s)
}

type DeploymentUserVpc struct {
	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 交换机ID
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DeploymentUserVpc) String() string {
	return dara.Prettify(s)
}

func (s DeploymentUserVpc) GoString() string {
	return s.String()
}

func (s *DeploymentUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DeploymentUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DeploymentUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DeploymentUserVpc) SetSecurityGroupId(v string) *DeploymentUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *DeploymentUserVpc) SetVSwitchId(v string) *DeploymentUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *DeploymentUserVpc) SetVpcId(v string) *DeploymentUserVpc {
	s.VpcId = &v
	return s
}

func (s *DeploymentUserVpc) Validate() error {
	return dara.Validate(s)
}
