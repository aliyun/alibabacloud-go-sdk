// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRuntime interface {
	dara.Model
	String() string
	GoString() string
	SetAccessibility(v string) *Runtime
	GetAccessibility() *string
	SetCreator(v string) *Runtime
	GetCreator() *string
	SetCredentialConfig(v *RuntimeCredentialConfig) *Runtime
	GetCredentialConfig() *RuntimeCredentialConfig
	SetDataSources(v []*RuntimeDataSources) *Runtime
	GetDataSources() []*RuntimeDataSources
	SetEcsSpec(v *RuntimeEcsSpec) *Runtime
	GetEcsSpec() *RuntimeEcsSpec
	SetEnvs(v []*RuntimeEnvs) *Runtime
	GetEnvs() []*RuntimeEnvs
	SetFlowId(v string) *Runtime
	GetFlowId() *string
	SetFlows(v []*RuntimeFlows) *Runtime
	GetFlows() []*RuntimeFlows
	SetGmtCreateTime(v string) *Runtime
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *Runtime
	GetGmtModifiedTime() *string
	SetIdleTimeout(v int32) *Runtime
	GetIdleTimeout() *int32
	SetLabels(v []*RuntimeLabels) *Runtime
	GetLabels() []*RuntimeLabels
	SetLatestVersion(v string) *Runtime
	GetLatestVersion() *string
	SetMCPConfig(v string) *Runtime
	GetMCPConfig() *string
	SetResourceId(v string) *Runtime
	GetResourceId() *string
	SetRunTimeout(v int32) *Runtime
	GetRunTimeout() *int32
	SetRuntimeId(v string) *Runtime
	GetRuntimeId() *string
	SetRuntimeLog(v string) *Runtime
	GetRuntimeLog() *string
	SetRuntimeName(v string) *Runtime
	GetRuntimeName() *string
	SetRuntimeStatus(v string) *Runtime
	GetRuntimeStatus() *string
	SetRuntimeType(v string) *Runtime
	GetRuntimeType() *string
	SetSupportSSEStream(v bool) *Runtime
	GetSupportSSEStream() *bool
	SetUserVpc(v *RuntimeUserVpc) *Runtime
	GetUserVpc() *RuntimeUserVpc
	SetVersion(v string) *Runtime
	GetVersion() *string
	SetWorkDir(v string) *Runtime
	GetWorkDir() *string
	SetWorkspaceId(v string) *Runtime
	GetWorkspaceId() *string
}

type Runtime struct {
	Accessibility    *string                  `json:"Accessibility,omitempty" xml:"Accessibility,omitempty"`
	Creator          *string                  `json:"Creator,omitempty" xml:"Creator,omitempty"`
	CredentialConfig *RuntimeCredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty" type:"Struct"`
	DataSources      []*RuntimeDataSources    `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	EcsSpec          *RuntimeEcsSpec          `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty" type:"Struct"`
	Envs             []*RuntimeEnvs           `json:"Envs,omitempty" xml:"Envs,omitempty" type:"Repeated"`
	FlowId           *string                  `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	Flows            []*RuntimeFlows          `json:"Flows,omitempty" xml:"Flows,omitempty" type:"Repeated"`
	GmtCreateTime    *string                  `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtModifiedTime  *string                  `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	IdleTimeout      *int32                   `json:"IdleTimeout,omitempty" xml:"IdleTimeout,omitempty"`
	Labels           []*RuntimeLabels         `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	LatestVersion    *string                  `json:"LatestVersion,omitempty" xml:"LatestVersion,omitempty"`
	MCPConfig        *string                  `json:"MCPConfig,omitempty" xml:"MCPConfig,omitempty"`
	ResourceId       *string                  `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	RunTimeout       *int32                   `json:"RunTimeout,omitempty" xml:"RunTimeout,omitempty"`
	RuntimeId        *string                  `json:"RuntimeId,omitempty" xml:"RuntimeId,omitempty"`
	RuntimeLog       *string                  `json:"RuntimeLog,omitempty" xml:"RuntimeLog,omitempty"`
	RuntimeName      *string                  `json:"RuntimeName,omitempty" xml:"RuntimeName,omitempty"`
	RuntimeStatus    *string                  `json:"RuntimeStatus,omitempty" xml:"RuntimeStatus,omitempty"`
	RuntimeType      *string                  `json:"RuntimeType,omitempty" xml:"RuntimeType,omitempty"`
	SupportSSEStream *bool                    `json:"SupportSSEStream,omitempty" xml:"SupportSSEStream,omitempty"`
	UserVpc          *RuntimeUserVpc          `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	Version          *string                  `json:"Version,omitempty" xml:"Version,omitempty"`
	WorkDir          *string                  `json:"WorkDir,omitempty" xml:"WorkDir,omitempty"`
	WorkspaceId      *string                  `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s Runtime) String() string {
	return dara.Prettify(s)
}

func (s Runtime) GoString() string {
	return s.String()
}

func (s *Runtime) GetAccessibility() *string {
	return s.Accessibility
}

func (s *Runtime) GetCreator() *string {
	return s.Creator
}

func (s *Runtime) GetCredentialConfig() *RuntimeCredentialConfig {
	return s.CredentialConfig
}

func (s *Runtime) GetDataSources() []*RuntimeDataSources {
	return s.DataSources
}

func (s *Runtime) GetEcsSpec() *RuntimeEcsSpec {
	return s.EcsSpec
}

func (s *Runtime) GetEnvs() []*RuntimeEnvs {
	return s.Envs
}

func (s *Runtime) GetFlowId() *string {
	return s.FlowId
}

func (s *Runtime) GetFlows() []*RuntimeFlows {
	return s.Flows
}

func (s *Runtime) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Runtime) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Runtime) GetIdleTimeout() *int32 {
	return s.IdleTimeout
}

func (s *Runtime) GetLabels() []*RuntimeLabels {
	return s.Labels
}

func (s *Runtime) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *Runtime) GetMCPConfig() *string {
	return s.MCPConfig
}

func (s *Runtime) GetResourceId() *string {
	return s.ResourceId
}

func (s *Runtime) GetRunTimeout() *int32 {
	return s.RunTimeout
}

func (s *Runtime) GetRuntimeId() *string {
	return s.RuntimeId
}

func (s *Runtime) GetRuntimeLog() *string {
	return s.RuntimeLog
}

func (s *Runtime) GetRuntimeName() *string {
	return s.RuntimeName
}

func (s *Runtime) GetRuntimeStatus() *string {
	return s.RuntimeStatus
}

func (s *Runtime) GetRuntimeType() *string {
	return s.RuntimeType
}

func (s *Runtime) GetSupportSSEStream() *bool {
	return s.SupportSSEStream
}

func (s *Runtime) GetUserVpc() *RuntimeUserVpc {
	return s.UserVpc
}

func (s *Runtime) GetVersion() *string {
	return s.Version
}

func (s *Runtime) GetWorkDir() *string {
	return s.WorkDir
}

func (s *Runtime) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *Runtime) SetAccessibility(v string) *Runtime {
	s.Accessibility = &v
	return s
}

func (s *Runtime) SetCreator(v string) *Runtime {
	s.Creator = &v
	return s
}

func (s *Runtime) SetCredentialConfig(v *RuntimeCredentialConfig) *Runtime {
	s.CredentialConfig = v
	return s
}

func (s *Runtime) SetDataSources(v []*RuntimeDataSources) *Runtime {
	s.DataSources = v
	return s
}

func (s *Runtime) SetEcsSpec(v *RuntimeEcsSpec) *Runtime {
	s.EcsSpec = v
	return s
}

func (s *Runtime) SetEnvs(v []*RuntimeEnvs) *Runtime {
	s.Envs = v
	return s
}

func (s *Runtime) SetFlowId(v string) *Runtime {
	s.FlowId = &v
	return s
}

func (s *Runtime) SetFlows(v []*RuntimeFlows) *Runtime {
	s.Flows = v
	return s
}

func (s *Runtime) SetGmtCreateTime(v string) *Runtime {
	s.GmtCreateTime = &v
	return s
}

func (s *Runtime) SetGmtModifiedTime(v string) *Runtime {
	s.GmtModifiedTime = &v
	return s
}

func (s *Runtime) SetIdleTimeout(v int32) *Runtime {
	s.IdleTimeout = &v
	return s
}

func (s *Runtime) SetLabels(v []*RuntimeLabels) *Runtime {
	s.Labels = v
	return s
}

func (s *Runtime) SetLatestVersion(v string) *Runtime {
	s.LatestVersion = &v
	return s
}

func (s *Runtime) SetMCPConfig(v string) *Runtime {
	s.MCPConfig = &v
	return s
}

func (s *Runtime) SetResourceId(v string) *Runtime {
	s.ResourceId = &v
	return s
}

func (s *Runtime) SetRunTimeout(v int32) *Runtime {
	s.RunTimeout = &v
	return s
}

func (s *Runtime) SetRuntimeId(v string) *Runtime {
	s.RuntimeId = &v
	return s
}

func (s *Runtime) SetRuntimeLog(v string) *Runtime {
	s.RuntimeLog = &v
	return s
}

func (s *Runtime) SetRuntimeName(v string) *Runtime {
	s.RuntimeName = &v
	return s
}

func (s *Runtime) SetRuntimeStatus(v string) *Runtime {
	s.RuntimeStatus = &v
	return s
}

func (s *Runtime) SetRuntimeType(v string) *Runtime {
	s.RuntimeType = &v
	return s
}

func (s *Runtime) SetSupportSSEStream(v bool) *Runtime {
	s.SupportSSEStream = &v
	return s
}

func (s *Runtime) SetUserVpc(v *RuntimeUserVpc) *Runtime {
	s.UserVpc = v
	return s
}

func (s *Runtime) SetVersion(v string) *Runtime {
	s.Version = &v
	return s
}

func (s *Runtime) SetWorkDir(v string) *Runtime {
	s.WorkDir = &v
	return s
}

func (s *Runtime) SetWorkspaceId(v string) *Runtime {
	s.WorkspaceId = &v
	return s
}

func (s *Runtime) Validate() error {
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
	if s.Flows != nil {
		for _, item := range s.Flows {
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

type RuntimeCredentialConfig struct {
	// AliyunEnvRoleKey
	AliyunEnvRoleKey *string `json:"AliyunEnvRoleKey,omitempty" xml:"AliyunEnvRoleKey,omitempty"`
	// Credential配置项列表
	CredentialConfigItems []*RuntimeCredentialConfigCredentialConfigItems `json:"CredentialConfigItems,omitempty" xml:"CredentialConfigItems,omitempty" type:"Repeated"`
	// 是否启用Credential注入
	EnableCredentialInject *bool `json:"EnableCredentialInject,omitempty" xml:"EnableCredentialInject,omitempty"`
}

func (s RuntimeCredentialConfig) String() string {
	return dara.Prettify(s)
}

func (s RuntimeCredentialConfig) GoString() string {
	return s.String()
}

func (s *RuntimeCredentialConfig) GetAliyunEnvRoleKey() *string {
	return s.AliyunEnvRoleKey
}

func (s *RuntimeCredentialConfig) GetCredentialConfigItems() []*RuntimeCredentialConfigCredentialConfigItems {
	return s.CredentialConfigItems
}

func (s *RuntimeCredentialConfig) GetEnableCredentialInject() *bool {
	return s.EnableCredentialInject
}

func (s *RuntimeCredentialConfig) SetAliyunEnvRoleKey(v string) *RuntimeCredentialConfig {
	s.AliyunEnvRoleKey = &v
	return s
}

func (s *RuntimeCredentialConfig) SetCredentialConfigItems(v []*RuntimeCredentialConfigCredentialConfigItems) *RuntimeCredentialConfig {
	s.CredentialConfigItems = v
	return s
}

func (s *RuntimeCredentialConfig) SetEnableCredentialInject(v bool) *RuntimeCredentialConfig {
	s.EnableCredentialInject = &v
	return s
}

func (s *RuntimeCredentialConfig) Validate() error {
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

type RuntimeCredentialConfigCredentialConfigItems struct {
	// Key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 角色列表
	Roles []*RuntimeCredentialConfigCredentialConfigItemsRoles `json:"Roles,omitempty" xml:"Roles,omitempty" type:"Repeated"`
	// Type
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RuntimeCredentialConfigCredentialConfigItems) String() string {
	return dara.Prettify(s)
}

func (s RuntimeCredentialConfigCredentialConfigItems) GoString() string {
	return s.String()
}

func (s *RuntimeCredentialConfigCredentialConfigItems) GetKey() *string {
	return s.Key
}

func (s *RuntimeCredentialConfigCredentialConfigItems) GetRoles() []*RuntimeCredentialConfigCredentialConfigItemsRoles {
	return s.Roles
}

func (s *RuntimeCredentialConfigCredentialConfigItems) GetType() *string {
	return s.Type
}

func (s *RuntimeCredentialConfigCredentialConfigItems) SetKey(v string) *RuntimeCredentialConfigCredentialConfigItems {
	s.Key = &v
	return s
}

func (s *RuntimeCredentialConfigCredentialConfigItems) SetRoles(v []*RuntimeCredentialConfigCredentialConfigItemsRoles) *RuntimeCredentialConfigCredentialConfigItems {
	s.Roles = v
	return s
}

func (s *RuntimeCredentialConfigCredentialConfigItems) SetType(v string) *RuntimeCredentialConfigCredentialConfigItems {
	s.Type = &v
	return s
}

func (s *RuntimeCredentialConfigCredentialConfigItems) Validate() error {
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

type RuntimeCredentialConfigCredentialConfigItemsRoles struct {
	// AssumeRoleFor
	AssumeRoleFor *string `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// 角色名称
	RoleArn *string `json:"RoleArn,omitempty" xml:"RoleArn,omitempty"`
	// 角色类型
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
}

func (s RuntimeCredentialConfigCredentialConfigItemsRoles) String() string {
	return dara.Prettify(s)
}

func (s RuntimeCredentialConfigCredentialConfigItemsRoles) GoString() string {
	return s.String()
}

func (s *RuntimeCredentialConfigCredentialConfigItemsRoles) GetAssumeRoleFor() *string {
	return s.AssumeRoleFor
}

func (s *RuntimeCredentialConfigCredentialConfigItemsRoles) GetRoleArn() *string {
	return s.RoleArn
}

func (s *RuntimeCredentialConfigCredentialConfigItemsRoles) GetRoleType() *string {
	return s.RoleType
}

func (s *RuntimeCredentialConfigCredentialConfigItemsRoles) SetAssumeRoleFor(v string) *RuntimeCredentialConfigCredentialConfigItemsRoles {
	s.AssumeRoleFor = &v
	return s
}

func (s *RuntimeCredentialConfigCredentialConfigItemsRoles) SetRoleArn(v string) *RuntimeCredentialConfigCredentialConfigItemsRoles {
	s.RoleArn = &v
	return s
}

func (s *RuntimeCredentialConfigCredentialConfigItemsRoles) SetRoleType(v string) *RuntimeCredentialConfigCredentialConfigItemsRoles {
	s.RoleType = &v
	return s
}

func (s *RuntimeCredentialConfigCredentialConfigItemsRoles) Validate() error {
	return dara.Validate(s)
}

type RuntimeDataSources struct {
	// 数据集ID
	DatasetId *string `json:"DatasetId,omitempty" xml:"DatasetId,omitempty"`
	// 挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 统一资源识别码
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s RuntimeDataSources) String() string {
	return dara.Prettify(s)
}

func (s RuntimeDataSources) GoString() string {
	return s.String()
}

func (s *RuntimeDataSources) GetDatasetId() *string {
	return s.DatasetId
}

func (s *RuntimeDataSources) GetMountPath() *string {
	return s.MountPath
}

func (s *RuntimeDataSources) GetUri() *string {
	return s.Uri
}

func (s *RuntimeDataSources) SetDatasetId(v string) *RuntimeDataSources {
	s.DatasetId = &v
	return s
}

func (s *RuntimeDataSources) SetMountPath(v string) *RuntimeDataSources {
	s.MountPath = &v
	return s
}

func (s *RuntimeDataSources) SetUri(v string) *RuntimeDataSources {
	s.Uri = &v
	return s
}

func (s *RuntimeDataSources) Validate() error {
	return dara.Validate(s)
}

type RuntimeEcsSpec struct {
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

func (s RuntimeEcsSpec) String() string {
	return dara.Prettify(s)
}

func (s RuntimeEcsSpec) GoString() string {
	return s.String()
}

func (s *RuntimeEcsSpec) GetCPU() *int32 {
	return s.CPU
}

func (s *RuntimeEcsSpec) GetExtraEphemeralStorage() *int32 {
	return s.ExtraEphemeralStorage
}

func (s *RuntimeEcsSpec) GetGPU() *int32 {
	return s.GPU
}

func (s *RuntimeEcsSpec) GetGPUCorePercentage() *int32 {
	return s.GPUCorePercentage
}

func (s *RuntimeEcsSpec) GetGPUMemory() *int32 {
	return s.GPUMemory
}

func (s *RuntimeEcsSpec) GetGPUType() *string {
	return s.GPUType
}

func (s *RuntimeEcsSpec) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RuntimeEcsSpec) GetMemory() *int32 {
	return s.Memory
}

func (s *RuntimeEcsSpec) GetPodCount() *int32 {
	return s.PodCount
}

func (s *RuntimeEcsSpec) GetQuotaId() *string {
	return s.QuotaId
}

func (s *RuntimeEcsSpec) GetQuotaType() *string {
	return s.QuotaType
}

func (s *RuntimeEcsSpec) GetSharedMemory() *int32 {
	return s.SharedMemory
}

func (s *RuntimeEcsSpec) SetCPU(v int32) *RuntimeEcsSpec {
	s.CPU = &v
	return s
}

func (s *RuntimeEcsSpec) SetExtraEphemeralStorage(v int32) *RuntimeEcsSpec {
	s.ExtraEphemeralStorage = &v
	return s
}

func (s *RuntimeEcsSpec) SetGPU(v int32) *RuntimeEcsSpec {
	s.GPU = &v
	return s
}

func (s *RuntimeEcsSpec) SetGPUCorePercentage(v int32) *RuntimeEcsSpec {
	s.GPUCorePercentage = &v
	return s
}

func (s *RuntimeEcsSpec) SetGPUMemory(v int32) *RuntimeEcsSpec {
	s.GPUMemory = &v
	return s
}

func (s *RuntimeEcsSpec) SetGPUType(v string) *RuntimeEcsSpec {
	s.GPUType = &v
	return s
}

func (s *RuntimeEcsSpec) SetInstanceType(v string) *RuntimeEcsSpec {
	s.InstanceType = &v
	return s
}

func (s *RuntimeEcsSpec) SetMemory(v int32) *RuntimeEcsSpec {
	s.Memory = &v
	return s
}

func (s *RuntimeEcsSpec) SetPodCount(v int32) *RuntimeEcsSpec {
	s.PodCount = &v
	return s
}

func (s *RuntimeEcsSpec) SetQuotaId(v string) *RuntimeEcsSpec {
	s.QuotaId = &v
	return s
}

func (s *RuntimeEcsSpec) SetQuotaType(v string) *RuntimeEcsSpec {
	s.QuotaType = &v
	return s
}

func (s *RuntimeEcsSpec) SetSharedMemory(v int32) *RuntimeEcsSpec {
	s.SharedMemory = &v
	return s
}

func (s *RuntimeEcsSpec) Validate() error {
	return dara.Validate(s)
}

type RuntimeEnvs struct {
	// 环境键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 环境值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RuntimeEnvs) String() string {
	return dara.Prettify(s)
}

func (s RuntimeEnvs) GoString() string {
	return s.String()
}

func (s *RuntimeEnvs) GetKey() *string {
	return s.Key
}

func (s *RuntimeEnvs) GetValue() *string {
	return s.Value
}

func (s *RuntimeEnvs) SetKey(v string) *RuntimeEnvs {
	s.Key = &v
	return s
}

func (s *RuntimeEnvs) SetValue(v string) *RuntimeEnvs {
	s.Value = &v
	return s
}

func (s *RuntimeEnvs) Validate() error {
	return dara.Validate(s)
}

type RuntimeFlows struct {
	// 应用流ID
	FlowId *string `json:"FlowId,omitempty" xml:"FlowId,omitempty"`
	// 应用流名称
	FlowName *string `json:"FlowName,omitempty" xml:"FlowName,omitempty"`
}

func (s RuntimeFlows) String() string {
	return dara.Prettify(s)
}

func (s RuntimeFlows) GoString() string {
	return s.String()
}

func (s *RuntimeFlows) GetFlowId() *string {
	return s.FlowId
}

func (s *RuntimeFlows) GetFlowName() *string {
	return s.FlowName
}

func (s *RuntimeFlows) SetFlowId(v string) *RuntimeFlows {
	s.FlowId = &v
	return s
}

func (s *RuntimeFlows) SetFlowName(v string) *RuntimeFlows {
	s.FlowName = &v
	return s
}

func (s *RuntimeFlows) Validate() error {
	return dara.Validate(s)
}

type RuntimeLabels struct {
	// 标签键
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// 标签值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RuntimeLabels) String() string {
	return dara.Prettify(s)
}

func (s RuntimeLabels) GoString() string {
	return s.String()
}

func (s *RuntimeLabels) GetKey() *string {
	return s.Key
}

func (s *RuntimeLabels) GetValue() *string {
	return s.Value
}

func (s *RuntimeLabels) SetKey(v string) *RuntimeLabels {
	s.Key = &v
	return s
}

func (s *RuntimeLabels) SetValue(v string) *RuntimeLabels {
	s.Value = &v
	return s
}

func (s *RuntimeLabels) Validate() error {
	return dara.Validate(s)
}

type RuntimeUserVpc struct {
	// 默认路由
	DefaultRoute *string `json:"DefaultRoute,omitempty" xml:"DefaultRoute,omitempty"`
	// 扩展网段
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// 安全组ID
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 交换机ID
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s RuntimeUserVpc) String() string {
	return dara.Prettify(s)
}

func (s RuntimeUserVpc) GoString() string {
	return s.String()
}

func (s *RuntimeUserVpc) GetDefaultRoute() *string {
	return s.DefaultRoute
}

func (s *RuntimeUserVpc) GetExtendedCIDRs() []*string {
	return s.ExtendedCIDRs
}

func (s *RuntimeUserVpc) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RuntimeUserVpc) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RuntimeUserVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *RuntimeUserVpc) SetDefaultRoute(v string) *RuntimeUserVpc {
	s.DefaultRoute = &v
	return s
}

func (s *RuntimeUserVpc) SetExtendedCIDRs(v []*string) *RuntimeUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *RuntimeUserVpc) SetSecurityGroupId(v string) *RuntimeUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *RuntimeUserVpc) SetVSwitchId(v string) *RuntimeUserVpc {
	s.VSwitchId = &v
	return s
}

func (s *RuntimeUserVpc) SetVpcId(v string) *RuntimeUserVpc {
	s.VpcId = &v
	return s
}

func (s *RuntimeUserVpc) Validate() error {
	return dara.Validate(s)
}
