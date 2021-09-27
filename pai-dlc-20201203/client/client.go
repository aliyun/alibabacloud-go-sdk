// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type JobSettings struct {
	// 作业关联用户ID
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// 调用方
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// 自定义标签
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// 工作流ID
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s JobSettings) String() string {
	return tea.Prettify(s)
}

func (s JobSettings) GoString() string {
	return s.String()
}

func (s *JobSettings) SetBusinessUserId(v string) *JobSettings {
	s.BusinessUserId = &v
	return s
}

func (s *JobSettings) SetCaller(v string) *JobSettings {
	s.Caller = &v
	return s
}

func (s *JobSettings) SetTags(v map[string]*string) *JobSettings {
	s.Tags = v
	return s
}

func (s *JobSettings) SetPipelineId(v string) *JobSettings {
	s.PipelineId = &v
	return s
}

type QuotaDetail struct {
	// CPU核数
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// 内存容量
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// GPU卡数
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// GPU卡型
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// GPU卡型全名
	GPUTypeFullName *string `json:"GPUTypeFullName,omitempty" xml:"GPUTypeFullName,omitempty"`
	// GPU详情
	GPUDetails []*GPUDetail `json:"GPUDetails,omitempty" xml:"GPUDetails,omitempty" type:"Repeated"`
}

func (s QuotaDetail) String() string {
	return tea.Prettify(s)
}

func (s QuotaDetail) GoString() string {
	return s.String()
}

func (s *QuotaDetail) SetCPU(v string) *QuotaDetail {
	s.CPU = &v
	return s
}

func (s *QuotaDetail) SetMemory(v string) *QuotaDetail {
	s.Memory = &v
	return s
}

func (s *QuotaDetail) SetGPU(v string) *QuotaDetail {
	s.GPU = &v
	return s
}

func (s *QuotaDetail) SetGPUType(v string) *QuotaDetail {
	s.GPUType = &v
	return s
}

func (s *QuotaDetail) SetGPUTypeFullName(v string) *QuotaDetail {
	s.GPUTypeFullName = &v
	return s
}

func (s *QuotaDetail) SetGPUDetails(v []*GPUDetail) *QuotaDetail {
	s.GPUDetails = v
	return s
}

type GPUDetail struct {
	// GPU卡数
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// GPU卡型
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// GPU卡型全名
	GPUTypeFullName *string `json:"GPUTypeFullName,omitempty" xml:"GPUTypeFullName,omitempty"`
}

func (s GPUDetail) String() string {
	return tea.Prettify(s)
}

func (s GPUDetail) GoString() string {
	return s.String()
}

func (s *GPUDetail) SetGPU(v string) *GPUDetail {
	s.GPU = &v
	return s
}

func (s *GPUDetail) SetGPUType(v string) *GPUDetail {
	s.GPUType = &v
	return s
}

func (s *GPUDetail) SetGPUTypeFullName(v string) *GPUDetail {
	s.GPUTypeFullName = &v
	return s
}

type Tensorboard struct {
	// Tensorboard Id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	// Tensorboard URL
	TensorboardUrl *string `json:"TensorboardUrl,omitempty" xml:"TensorboardUrl,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 运行时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间（UTC）
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 展示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 文件路径
	SummaryPath *string `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	// 创建者
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 状态详情码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 状态详情
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s Tensorboard) String() string {
	return tea.Prettify(s)
}

func (s Tensorboard) GoString() string {
	return s.String()
}

func (s *Tensorboard) SetTensorboardId(v string) *Tensorboard {
	s.TensorboardId = &v
	return s
}

func (s *Tensorboard) SetTensorboardUrl(v string) *Tensorboard {
	s.TensorboardUrl = &v
	return s
}

func (s *Tensorboard) SetStatus(v string) *Tensorboard {
	s.Status = &v
	return s
}

func (s *Tensorboard) SetDuration(v string) *Tensorboard {
	s.Duration = &v
	return s
}

func (s *Tensorboard) SetGmtCreateTime(v string) *Tensorboard {
	s.GmtCreateTime = &v
	return s
}

func (s *Tensorboard) SetGmtModifyTime(v string) *Tensorboard {
	s.GmtModifyTime = &v
	return s
}

func (s *Tensorboard) SetRequestId(v string) *Tensorboard {
	s.RequestId = &v
	return s
}

func (s *Tensorboard) SetDisplayName(v string) *Tensorboard {
	s.DisplayName = &v
	return s
}

func (s *Tensorboard) SetDataSourceId(v string) *Tensorboard {
	s.DataSourceId = &v
	return s
}

func (s *Tensorboard) SetSummaryPath(v string) *Tensorboard {
	s.SummaryPath = &v
	return s
}

func (s *Tensorboard) SetUserId(v string) *Tensorboard {
	s.UserId = &v
	return s
}

func (s *Tensorboard) SetReasonCode(v string) *Tensorboard {
	s.ReasonCode = &v
	return s
}

func (s *Tensorboard) SetReasonMessage(v string) *Tensorboard {
	s.ReasonMessage = &v
	return s
}

func (s *Tensorboard) SetJobId(v string) *Tensorboard {
	s.JobId = &v
	return s
}

type AliyunAccounts struct {
	// Aliyun账号的UID
	AliyunUid *string `json:"AliyunUid,omitempty" xml:"AliyunUid,omitempty"`
	// 弹内用户的工号
	EmployeeId *string `json:"EmployeeId,omitempty" xml:"EmployeeId,omitempty"`
	// 这条记录的创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 这条记录的上次修改时间
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
}

func (s AliyunAccounts) String() string {
	return tea.Prettify(s)
}

func (s AliyunAccounts) GoString() string {
	return s.String()
}

func (s *AliyunAccounts) SetAliyunUid(v string) *AliyunAccounts {
	s.AliyunUid = &v
	return s
}

func (s *AliyunAccounts) SetEmployeeId(v string) *AliyunAccounts {
	s.EmployeeId = &v
	return s
}

func (s *AliyunAccounts) SetGmtCreateTime(v string) *AliyunAccounts {
	s.GmtCreateTime = &v
	return s
}

func (s *AliyunAccounts) SetGmtModifyTime(v string) *AliyunAccounts {
	s.GmtModifyTime = &v
	return s
}

type ContainerSpec struct {
	// 容器名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 容器镜像地址
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// 用户命令
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// 命令参数
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// 容器内工作目录
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
	// 环境变量
	Env []*EnvVar `json:"Env,omitempty" xml:"Env,omitempty" type:"Repeated"`
}

func (s ContainerSpec) String() string {
	return tea.Prettify(s)
}

func (s ContainerSpec) GoString() string {
	return s.String()
}

func (s *ContainerSpec) SetName(v string) *ContainerSpec {
	s.Name = &v
	return s
}

func (s *ContainerSpec) SetImage(v string) *ContainerSpec {
	s.Image = &v
	return s
}

func (s *ContainerSpec) SetCommand(v []*string) *ContainerSpec {
	s.Command = v
	return s
}

func (s *ContainerSpec) SetArgs(v []*string) *ContainerSpec {
	s.Args = v
	return s
}

func (s *ContainerSpec) SetWorkingDir(v string) *ContainerSpec {
	s.WorkingDir = &v
	return s
}

func (s *ContainerSpec) SetEnv(v []*EnvVar) *ContainerSpec {
	s.Env = v
	return s
}

type EcsSpec struct {
	// 规格类型
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// cpu数量
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// gpu数量
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// gpu类型
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// Memory数量
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s EcsSpec) String() string {
	return tea.Prettify(s)
}

func (s EcsSpec) GoString() string {
	return s.String()
}

func (s *EcsSpec) SetInstanceType(v string) *EcsSpec {
	s.InstanceType = &v
	return s
}

func (s *EcsSpec) SetCpu(v int32) *EcsSpec {
	s.Cpu = &v
	return s
}

func (s *EcsSpec) SetGpu(v int32) *EcsSpec {
	s.Gpu = &v
	return s
}

func (s *EcsSpec) SetGpuType(v string) *EcsSpec {
	s.GpuType = &v
	return s
}

func (s *EcsSpec) SetMemory(v int32) *EcsSpec {
	s.Memory = &v
	return s
}

type PodMetric struct {
	// Pod编号
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// 监控指标样本序列
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
}

func (s PodMetric) String() string {
	return tea.Prettify(s)
}

func (s PodMetric) GoString() string {
	return s.String()
}

func (s *PodMetric) SetPodId(v string) *PodMetric {
	s.PodId = &v
	return s
}

func (s *PodMetric) SetMetrics(v []*Metric) *PodMetric {
	s.Metrics = v
	return s
}

type JobItem struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 作业类型
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// 作业显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 作业提交人Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 作业状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 作业所属工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 作业所属工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	// 作业运行所在的资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 状态详情码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 状态详情
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 作业规格配置
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// 用户命令
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// 数据源配置列表
	DataSources []*JobItemDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// 代码源配置
	CodeSource *JobItemCodeSource `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	// 三方库配置列表
	ThirdpartyLibs []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// 三方库(requirements.txt)文件路径
	ThirdpartyLibDir *string `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	// 环境变量配置
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// 作业创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 作业结束时间（UTC）
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// 作业运行时长，单位：秒
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 作业额外参数
	Settings *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
}

func (s JobItem) String() string {
	return tea.Prettify(s)
}

func (s JobItem) GoString() string {
	return s.String()
}

func (s *JobItem) SetJobId(v string) *JobItem {
	s.JobId = &v
	return s
}

func (s *JobItem) SetJobType(v string) *JobItem {
	s.JobType = &v
	return s
}

func (s *JobItem) SetDisplayName(v string) *JobItem {
	s.DisplayName = &v
	return s
}

func (s *JobItem) SetUserId(v string) *JobItem {
	s.UserId = &v
	return s
}

func (s *JobItem) SetStatus(v string) *JobItem {
	s.Status = &v
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

func (s *JobItem) SetResourceId(v string) *JobItem {
	s.ResourceId = &v
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

func (s *JobItem) SetJobSpecs(v []*JobSpec) *JobItem {
	s.JobSpecs = v
	return s
}

func (s *JobItem) SetUserCommand(v string) *JobItem {
	s.UserCommand = &v
	return s
}

func (s *JobItem) SetDataSources(v []*JobItemDataSources) *JobItem {
	s.DataSources = v
	return s
}

func (s *JobItem) SetCodeSource(v *JobItemCodeSource) *JobItem {
	s.CodeSource = v
	return s
}

func (s *JobItem) SetThirdpartyLibs(v []*string) *JobItem {
	s.ThirdpartyLibs = v
	return s
}

func (s *JobItem) SetThirdpartyLibDir(v string) *JobItem {
	s.ThirdpartyLibDir = &v
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

func (s *JobItem) SetGmtFinishTime(v string) *JobItem {
	s.GmtFinishTime = &v
	return s
}

func (s *JobItem) SetDuration(v int64) *JobItem {
	s.Duration = &v
	return s
}

func (s *JobItem) SetSettings(v *JobSettings) *JobItem {
	s.Settings = v
	return s
}

type JobItemDataSources struct {
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 本地挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s JobItemDataSources) String() string {
	return tea.Prettify(s)
}

func (s JobItemDataSources) GoString() string {
	return s.String()
}

func (s *JobItemDataSources) SetDataSourceId(v string) *JobItemDataSources {
	s.DataSourceId = &v
	return s
}

func (s *JobItemDataSources) SetMountPath(v string) *JobItemDataSources {
	s.MountPath = &v
	return s
}

type JobItemCodeSource struct {
	// 代码源Id
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// 代码分支
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// 代码Commit
	Commit *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// 本地挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s JobItemCodeSource) String() string {
	return tea.Prettify(s)
}

func (s JobItemCodeSource) GoString() string {
	return s.String()
}

func (s *JobItemCodeSource) SetCodeSourceId(v string) *JobItemCodeSource {
	s.CodeSourceId = &v
	return s
}

func (s *JobItemCodeSource) SetBranch(v string) *JobItemCodeSource {
	s.Branch = &v
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

type JobElasticSpec struct {
	// 打开弹性训练
	EnableElasticTraining *bool `json:"EnableElasticTraining,omitempty" xml:"EnableElasticTraining,omitempty"`
	// 最小并行度
	MinParallelism *int32 `json:"MinParallelism,omitempty" xml:"MinParallelism,omitempty"`
	// 最大并行度
	MaxParallelism *int32 `json:"MaxParallelism,omitempty" xml:"MaxParallelism,omitempty"`
	// aimaster角色使用的资源规格
	AIMasterType *string `json:"AIMasterType,omitempty" xml:"AIMasterType,omitempty"`
}

func (s JobElasticSpec) String() string {
	return tea.Prettify(s)
}

func (s JobElasticSpec) GoString() string {
	return s.String()
}

func (s *JobElasticSpec) SetEnableElasticTraining(v bool) *JobElasticSpec {
	s.EnableElasticTraining = &v
	return s
}

func (s *JobElasticSpec) SetMinParallelism(v int32) *JobElasticSpec {
	s.MinParallelism = &v
	return s
}

func (s *JobElasticSpec) SetMaxParallelism(v int32) *JobElasticSpec {
	s.MaxParallelism = &v
	return s
}

func (s *JobElasticSpec) SetAIMasterType(v string) *JobElasticSpec {
	s.AIMasterType = &v
	return s
}

type NodeMetric struct {
	// 节点名称
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
	// 监控指标样本序列
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
}

func (s NodeMetric) String() string {
	return tea.Prettify(s)
}

func (s NodeMetric) GoString() string {
	return s.String()
}

func (s *NodeMetric) SetNodeName(v string) *NodeMetric {
	s.NodeName = &v
	return s
}

func (s *NodeMetric) SetMetrics(v []*Metric) *NodeMetric {
	s.Metrics = v
	return s
}

type CodeSourceItem struct {
	// 代码源ID
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// 代码源配置的名字
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 代码源详细描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代码仓库地址
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// 代码分支
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// 代码Commit ID
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// 访问代码仓库的用户名
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// 访问代码仓库所用的AccessToken
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// 代码源配置的用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
}

func (s CodeSourceItem) String() string {
	return tea.Prettify(s)
}

func (s CodeSourceItem) GoString() string {
	return s.String()
}

func (s *CodeSourceItem) SetCodeSourceId(v string) *CodeSourceItem {
	s.CodeSourceId = &v
	return s
}

func (s *CodeSourceItem) SetDisplayName(v string) *CodeSourceItem {
	s.DisplayName = &v
	return s
}

func (s *CodeSourceItem) SetDescription(v string) *CodeSourceItem {
	s.Description = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepo(v string) *CodeSourceItem {
	s.CodeRepo = &v
	return s
}

func (s *CodeSourceItem) SetCodeBranch(v string) *CodeSourceItem {
	s.CodeBranch = &v
	return s
}

func (s *CodeSourceItem) SetCodeCommit(v string) *CodeSourceItem {
	s.CodeCommit = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepoUserName(v string) *CodeSourceItem {
	s.CodeRepoUserName = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepoAccessToken(v string) *CodeSourceItem {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *CodeSourceItem) SetUserId(v string) *CodeSourceItem {
	s.UserId = &v
	return s
}

func (s *CodeSourceItem) SetGmtCreateTime(v string) *CodeSourceItem {
	s.GmtCreateTime = &v
	return s
}

func (s *CodeSourceItem) SetGmtModifyTime(v string) *CodeSourceItem {
	s.GmtModifyTime = &v
	return s
}

type Quota struct {
	// 资源配额id
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// 资源配额名称
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// 资源配额类型
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// 资源用量
	UsedQuota *QuotaDetail `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
	// 资源总量
	TotalQuota *QuotaDetail `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// 是否是独占Quota组
	IsExclusiveQuota *bool `json:"IsExclusiveQuota,omitempty" xml:"IsExclusiveQuota,omitempty"`
}

func (s Quota) String() string {
	return tea.Prettify(s)
}

func (s Quota) GoString() string {
	return s.String()
}

func (s *Quota) SetQuotaId(v string) *Quota {
	s.QuotaId = &v
	return s
}

func (s *Quota) SetQuotaName(v string) *Quota {
	s.QuotaName = &v
	return s
}

func (s *Quota) SetQuotaType(v string) *Quota {
	s.QuotaType = &v
	return s
}

func (s *Quota) SetUsedQuota(v *QuotaDetail) *Quota {
	s.UsedQuota = v
	return s
}

func (s *Quota) SetTotalQuota(v *QuotaDetail) *Quota {
	s.TotalQuota = v
	return s
}

func (s *Quota) SetClusterId(v string) *Quota {
	s.ClusterId = &v
	return s
}

func (s *Quota) SetClusterName(v string) *Quota {
	s.ClusterName = &v
	return s
}

func (s *Quota) SetIsExclusiveQuota(v bool) *Quota {
	s.IsExclusiveQuota = &v
	return s
}

type EnvVar struct {
	// 环境变量名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 环境变量值
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s EnvVar) String() string {
	return tea.Prettify(s)
}

func (s EnvVar) GoString() string {
	return s.String()
}

func (s *EnvVar) SetName(v string) *EnvVar {
	s.Name = &v
	return s
}

func (s *EnvVar) SetValue(v string) *EnvVar {
	s.Value = &v
	return s
}

type Metric struct {
	// 时间戳（毫秒）
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// 样本值
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s Metric) String() string {
	return tea.Prettify(s)
}

func (s Metric) GoString() string {
	return s.String()
}

func (s *Metric) SetTime(v int64) *Metric {
	s.Time = &v
	return s
}

func (s *Metric) SetValue(v float32) *Metric {
	s.Value = &v
	return s
}

type JobSpec struct {
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 镜像
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// 实例数量
	PodCount *int64 `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	// Ecs实例规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 额外的Pod配置
	ExtraPodSpec *ExtraPodSpec `json:"ExtraPodSpec,omitempty" xml:"ExtraPodSpec,omitempty"`
	// 资源配置
	ResourceConfig *ResourceConfig `json:"ResourceConfig,omitempty" xml:"ResourceConfig,omitempty"`
}

func (s JobSpec) String() string {
	return tea.Prettify(s)
}

func (s JobSpec) GoString() string {
	return s.String()
}

func (s *JobSpec) SetType(v string) *JobSpec {
	s.Type = &v
	return s
}

func (s *JobSpec) SetImage(v string) *JobSpec {
	s.Image = &v
	return s
}

func (s *JobSpec) SetPodCount(v int64) *JobSpec {
	s.PodCount = &v
	return s
}

func (s *JobSpec) SetEcsSpec(v string) *JobSpec {
	s.EcsSpec = &v
	return s
}

func (s *JobSpec) SetExtraPodSpec(v *ExtraPodSpec) *JobSpec {
	s.ExtraPodSpec = v
	return s
}

func (s *JobSpec) SetResourceConfig(v *ResourceConfig) *JobSpec {
	s.ResourceConfig = v
	return s
}

type DataSourceItem struct {
	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 数据源显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 数据源描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 阿里云NAS文件系统Id
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// 阿里云OSS文件系统路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 阿里云OSS文件系统服务端点
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// 阿里云OSS文件系统配置选项
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// 本地挂载目录
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 创建人Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间（UTC）
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
}

func (s DataSourceItem) String() string {
	return tea.Prettify(s)
}

func (s DataSourceItem) GoString() string {
	return s.String()
}

func (s *DataSourceItem) SetDataSourceType(v string) *DataSourceItem {
	s.DataSourceType = &v
	return s
}

func (s *DataSourceItem) SetDataSourceId(v string) *DataSourceItem {
	s.DataSourceId = &v
	return s
}

func (s *DataSourceItem) SetDisplayName(v string) *DataSourceItem {
	s.DisplayName = &v
	return s
}

func (s *DataSourceItem) SetDescription(v string) *DataSourceItem {
	s.Description = &v
	return s
}

func (s *DataSourceItem) SetFileSystemId(v string) *DataSourceItem {
	s.FileSystemId = &v
	return s
}

func (s *DataSourceItem) SetPath(v string) *DataSourceItem {
	s.Path = &v
	return s
}

func (s *DataSourceItem) SetEndpoint(v string) *DataSourceItem {
	s.Endpoint = &v
	return s
}

func (s *DataSourceItem) SetOptions(v string) *DataSourceItem {
	s.Options = &v
	return s
}

func (s *DataSourceItem) SetMountPath(v string) *DataSourceItem {
	s.MountPath = &v
	return s
}

func (s *DataSourceItem) SetUserId(v string) *DataSourceItem {
	s.UserId = &v
	return s
}

func (s *DataSourceItem) SetGmtCreateTime(v string) *DataSourceItem {
	s.GmtCreateTime = &v
	return s
}

func (s *DataSourceItem) SetGmtModifyTime(v string) *DataSourceItem {
	s.GmtModifyTime = &v
	return s
}

type ImageItem struct {
	// 镜像Tag
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 镜像vpc地址
	ImageUrlVpc *string `json:"ImageUrlVpc,omitempty" xml:"ImageUrlVpc,omitempty"`
	// 镜像类型
	ImageProviderType *string `json:"ImageProviderType,omitempty" xml:"ImageProviderType,omitempty"`
	// 加速器类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 镜像包含的框架类型
	Framework *string `json:"Framework,omitempty" xml:"Framework,omitempty"`
}

func (s ImageItem) String() string {
	return tea.Prettify(s)
}

func (s ImageItem) GoString() string {
	return s.String()
}

func (s *ImageItem) SetImageTag(v string) *ImageItem {
	s.ImageTag = &v
	return s
}

func (s *ImageItem) SetImageUrl(v string) *ImageItem {
	s.ImageUrl = &v
	return s
}

func (s *ImageItem) SetImageUrlVpc(v string) *ImageItem {
	s.ImageUrlVpc = &v
	return s
}

func (s *ImageItem) SetImageProviderType(v string) *ImageItem {
	s.ImageProviderType = &v
	return s
}

func (s *ImageItem) SetAcceleratorType(v string) *ImageItem {
	s.AcceleratorType = &v
	return s
}

func (s *ImageItem) SetFramework(v string) *ImageItem {
	s.Framework = &v
	return s
}

type ResourceConfig struct {
	// CPU核心数
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// GPU核心数
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// 内存容量
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 共享内存容量
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
	// 显卡类型
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
}

func (s ResourceConfig) String() string {
	return tea.Prettify(s)
}

func (s ResourceConfig) GoString() string {
	return s.String()
}

func (s *ResourceConfig) SetCPU(v string) *ResourceConfig {
	s.CPU = &v
	return s
}

func (s *ResourceConfig) SetGPU(v string) *ResourceConfig {
	s.GPU = &v
	return s
}

func (s *ResourceConfig) SetMemory(v string) *ResourceConfig {
	s.Memory = &v
	return s
}

func (s *ResourceConfig) SetSharedMemory(v string) *ResourceConfig {
	s.SharedMemory = &v
	return s
}

func (s *ResourceConfig) SetGPUType(v string) *ResourceConfig {
	s.GPUType = &v
	return s
}

type Resources struct {
	// CPU核心数
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// GPU卡数
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// 内存大小
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s Resources) String() string {
	return tea.Prettify(s)
}

func (s Resources) GoString() string {
	return s.String()
}

func (s *Resources) SetCPU(v string) *Resources {
	s.CPU = &v
	return s
}

func (s *Resources) SetGPU(v string) *Resources {
	s.GPU = &v
	return s
}

func (s *Resources) SetMemory(v string) *Resources {
	s.Memory = &v
	return s
}

type ExtraPodSpec struct {
	// 伴随容器
	SideCarContainers []*ContainerSpec `json:"SideCarContainers,omitempty" xml:"SideCarContainers,omitempty" type:"Repeated"`
	// 初始化容器
	InitContainers []*ContainerSpec `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	// Pod标签
	PodLabels map[string]*string `json:"PodLabels,omitempty" xml:"PodLabels,omitempty"`
	// Pod注解
	PodAnnotations map[string]*string `json:"PodAnnotations,omitempty" xml:"PodAnnotations,omitempty"`
	// 容器间共享的本地目录
	SharedVolumeMountPaths []*string `json:"SharedVolumeMountPaths,omitempty" xml:"SharedVolumeMountPaths,omitempty" type:"Repeated"`
}

func (s ExtraPodSpec) String() string {
	return tea.Prettify(s)
}

func (s ExtraPodSpec) GoString() string {
	return s.String()
}

func (s *ExtraPodSpec) SetSideCarContainers(v []*ContainerSpec) *ExtraPodSpec {
	s.SideCarContainers = v
	return s
}

func (s *ExtraPodSpec) SetInitContainers(v []*ContainerSpec) *ExtraPodSpec {
	s.InitContainers = v
	return s
}

func (s *ExtraPodSpec) SetPodLabels(v map[string]*string) *ExtraPodSpec {
	s.PodLabels = v
	return s
}

func (s *ExtraPodSpec) SetPodAnnotations(v map[string]*string) *ExtraPodSpec {
	s.PodAnnotations = v
	return s
}

func (s *ExtraPodSpec) SetSharedVolumeMountPaths(v []*string) *ExtraPodSpec {
	s.SharedVolumeMountPaths = v
	return s
}

type Member struct {
	// 成员id
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
	// 成员角色
	MemberType *string `json:"MemberType,omitempty" xml:"MemberType,omitempty"`
}

func (s Member) String() string {
	return tea.Prettify(s)
}

func (s Member) GoString() string {
	return s.String()
}

func (s *Member) SetMemberId(v string) *Member {
	s.MemberId = &v
	return s
}

func (s *Member) SetMemberType(v string) *Member {
	s.MemberType = &v
	return s
}

type Workspace struct {
	// 工作空间id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	// 创建者
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 最近修改时间
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 资源配额列表
	Quotas []*Quota `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// 资源总量
	TotalResources *Resources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty"`
	// 成员列表
	Members []*Member `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// 管理员列表
	WorkspaceAdmins []*Member `json:"WorkspaceAdmins,omitempty" xml:"WorkspaceAdmins,omitempty" type:"Repeated"`
}

func (s Workspace) String() string {
	return tea.Prettify(s)
}

func (s Workspace) GoString() string {
	return s.String()
}

func (s *Workspace) SetWorkspaceId(v string) *Workspace {
	s.WorkspaceId = &v
	return s
}

func (s *Workspace) SetWorkspaceName(v string) *Workspace {
	s.WorkspaceName = &v
	return s
}

func (s *Workspace) SetCreator(v string) *Workspace {
	s.Creator = &v
	return s
}

func (s *Workspace) SetGmtCreateTime(v string) *Workspace {
	s.GmtCreateTime = &v
	return s
}

func (s *Workspace) SetGmtModifyTime(v string) *Workspace {
	s.GmtModifyTime = &v
	return s
}

func (s *Workspace) SetQuotas(v []*Quota) *Workspace {
	s.Quotas = v
	return s
}

func (s *Workspace) SetTotalResources(v *Resources) *Workspace {
	s.TotalResources = v
	return s
}

func (s *Workspace) SetMembers(v []*Member) *Workspace {
	s.Members = v
	return s
}

func (s *Workspace) SetWorkspaceAdmins(v []*Member) *Workspace {
	s.WorkspaceAdmins = v
	return s
}

type ListDataSourcesRequest struct {
	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// 数据源显示名称，支持模糊查询
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 查询第几页数据
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 设置查询的分页大写
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListDataSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListDataSourcesRequest) SetDataSourceType(v string) *ListDataSourcesRequest {
	s.DataSourceType = &v
	return s
}

func (s *ListDataSourcesRequest) SetDisplayName(v string) *ListDataSourcesRequest {
	s.DisplayName = &v
	return s
}

func (s *ListDataSourcesRequest) SetPageNumber(v int32) *ListDataSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListDataSourcesRequest) SetPageSize(v int32) *ListDataSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListDataSourcesRequest) SetSortBy(v string) *ListDataSourcesRequest {
	s.SortBy = &v
	return s
}

func (s *ListDataSourcesRequest) SetOrder(v string) *ListDataSourcesRequest {
	s.Order = &v
	return s
}

type ListDataSourcesResponseBody struct {
	// 数据源配置列表
	DataSources []*DataSourceItem `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// 符合条件的数据源总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) SetDataSources(v []*DataSourceItem) *ListDataSourcesResponseBody {
	s.DataSources = v
	return s
}

func (s *ListDataSourcesResponseBody) SetTotalCount(v int64) *ListDataSourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListDataSourcesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListDataSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDataSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponse) SetHeaders(v map[string]*string) *ListDataSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListDataSourcesResponse) SetBody(v *ListDataSourcesResponseBody) *ListDataSourcesResponse {
	s.Body = v
	return s
}

type GetJobsStatisticsRequest struct {
	// 工作空间id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetJobsStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobsStatisticsRequest) GoString() string {
	return s.String()
}

func (s *GetJobsStatisticsRequest) SetWorkspaceId(v string) *GetJobsStatisticsRequest {
	s.WorkspaceId = &v
	return s
}

type GetJobsStatisticsResponseBody struct {
	// 当前的Workspace ID下面的JOB统计数据
	Statistics map[string]interface{} `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobsStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobsStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobsStatisticsResponseBody) SetStatistics(v map[string]interface{}) *GetJobsStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *GetJobsStatisticsResponseBody) SetRequestId(v string) *GetJobsStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type GetJobsStatisticsResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobsStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobsStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobsStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetJobsStatisticsResponse) SetHeaders(v map[string]*string) *GetJobsStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetJobsStatisticsResponse) SetBody(v *GetJobsStatisticsResponseBody) *GetJobsStatisticsResponse {
	s.Body = v
	return s
}

type GetDataSourceResponseBody struct {
	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 数据源显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 数据源描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 阿里云NAS文件系统Id
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// 阿里云OSS文件系统路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 阿里云OSS文件系统服务端点
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// 阿里云OSS文件系统配置选项
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// 本地挂载目录
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 创建人Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间（UTC）
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBody) SetDataSourceType(v string) *GetDataSourceResponseBody {
	s.DataSourceType = &v
	return s
}

func (s *GetDataSourceResponseBody) SetDataSourceId(v string) *GetDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *GetDataSourceResponseBody) SetDisplayName(v string) *GetDataSourceResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetDataSourceResponseBody) SetDescription(v string) *GetDataSourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetDataSourceResponseBody) SetFileSystemId(v string) *GetDataSourceResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *GetDataSourceResponseBody) SetPath(v string) *GetDataSourceResponseBody {
	s.Path = &v
	return s
}

func (s *GetDataSourceResponseBody) SetEndpoint(v string) *GetDataSourceResponseBody {
	s.Endpoint = &v
	return s
}

func (s *GetDataSourceResponseBody) SetOptions(v string) *GetDataSourceResponseBody {
	s.Options = &v
	return s
}

func (s *GetDataSourceResponseBody) SetMountPath(v string) *GetDataSourceResponseBody {
	s.MountPath = &v
	return s
}

func (s *GetDataSourceResponseBody) SetUserId(v string) *GetDataSourceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetDataSourceResponseBody) SetGmtCreateTime(v string) *GetDataSourceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetDataSourceResponseBody) SetGmtModifyTime(v string) *GetDataSourceResponseBody {
	s.GmtModifyTime = &v
	return s
}

func (s *GetDataSourceResponseBody) SetRequestId(v string) *GetDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type GetDataSourceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponse) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponse) SetHeaders(v map[string]*string) *GetDataSourceResponse {
	s.Headers = v
	return s
}

func (s *GetDataSourceResponse) SetBody(v *GetDataSourceResponseBody) *GetDataSourceResponse {
	s.Body = v
	return s
}

type CreateQuotaRequest struct {
	// 资源配额名称
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// 资源配额类型
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// 资源配额参数
	QuotaDetail *QuotaDetail `json:"QuotaDetail,omitempty" xml:"QuotaDetail,omitempty"`
}

func (s CreateQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaRequest) GoString() string {
	return s.String()
}

func (s *CreateQuotaRequest) SetQuotaName(v string) *CreateQuotaRequest {
	s.QuotaName = &v
	return s
}

func (s *CreateQuotaRequest) SetQuotaType(v string) *CreateQuotaRequest {
	s.QuotaType = &v
	return s
}

func (s *CreateQuotaRequest) SetQuotaDetail(v *QuotaDetail) *CreateQuotaRequest {
	s.QuotaDetail = v
	return s
}

type CreateQuotaResponseBody struct {
	// 资源配额id
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaResponseBody) SetQuotaId(v string) *CreateQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *CreateQuotaResponseBody) SetRequestId(v string) *CreateQuotaResponseBody {
	s.RequestId = &v
	return s
}

type CreateQuotaResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateQuotaResponse) GoString() string {
	return s.String()
}

func (s *CreateQuotaResponse) SetHeaders(v map[string]*string) *CreateQuotaResponse {
	s.Headers = v
	return s
}

func (s *CreateQuotaResponse) SetBody(v *CreateQuotaResponseBody) *CreateQuotaResponse {
	s.Body = v
	return s
}

type GetQuotaResponseBody struct {
	// 资源配额id
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// 资源配额名称
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// 资源配额类型
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// 资源用量
	UsedQuota *QuotaDetail `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
	// 资源总量
	TotalQuota *QuotaDetail `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBody) SetQuotaId(v string) *GetQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *GetQuotaResponseBody) SetQuotaName(v string) *GetQuotaResponseBody {
	s.QuotaName = &v
	return s
}

func (s *GetQuotaResponseBody) SetQuotaType(v string) *GetQuotaResponseBody {
	s.QuotaType = &v
	return s
}

func (s *GetQuotaResponseBody) SetUsedQuota(v *QuotaDetail) *GetQuotaResponseBody {
	s.UsedQuota = v
	return s
}

func (s *GetQuotaResponseBody) SetTotalQuota(v *QuotaDetail) *GetQuotaResponseBody {
	s.TotalQuota = v
	return s
}

func (s *GetQuotaResponseBody) SetClusterId(v string) *GetQuotaResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetQuotaResponseBody) SetClusterName(v string) *GetQuotaResponseBody {
	s.ClusterName = &v
	return s
}

func (s *GetQuotaResponseBody) SetRequestId(v string) *GetQuotaResponseBody {
	s.RequestId = &v
	return s
}

type GetQuotaResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaResponse) SetHeaders(v map[string]*string) *GetQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaResponse) SetBody(v *GetQuotaResponseBody) *GetQuotaResponse {
	s.Body = v
	return s
}

type GetCodeSourceResponseBody struct {
	// 代码源配置ID
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// 代码源配置名字
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 详细描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代码仓库地址
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// 代码仓库分支
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// 代码Commit
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// 代码仓库的用户名
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// 访问代码仓库的token
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// 代码本地挂载目录，默认挂载到/root/code/下
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 代码配置源的创建者ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCodeSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeSourceResponseBody) SetCodeSourceId(v string) *GetCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetDisplayName(v string) *GetCodeSourceResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetDescription(v string) *GetCodeSourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepo(v string) *GetCodeSourceResponseBody {
	s.CodeRepo = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeBranch(v string) *GetCodeSourceResponseBody {
	s.CodeBranch = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeCommit(v string) *GetCodeSourceResponseBody {
	s.CodeCommit = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepoUserName(v string) *GetCodeSourceResponseBody {
	s.CodeRepoUserName = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepoAccessToken(v string) *GetCodeSourceResponseBody {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetMountPath(v string) *GetCodeSourceResponseBody {
	s.MountPath = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetUserId(v string) *GetCodeSourceResponseBody {
	s.UserId = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetGmtCreateTime(v string) *GetCodeSourceResponseBody {
	s.GmtCreateTime = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetGmtModifyTime(v string) *GetCodeSourceResponseBody {
	s.GmtModifyTime = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetRequestId(v string) *GetCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

type GetCodeSourceResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetCodeSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *GetCodeSourceResponse) SetHeaders(v map[string]*string) *GetCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *GetCodeSourceResponse) SetBody(v *GetCodeSourceResponseBody) *GetCodeSourceResponse {
	s.Body = v
	return s
}

type GetSwitchResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 所属VPC的id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 交换机的id
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// 交换机的名称
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// 网段
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
}

func (s GetSwitchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *GetSwitchResponseBody) SetRequestId(v string) *GetSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSwitchResponseBody) SetVpcId(v string) *GetSwitchResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetSwitchResponseBody) SetVSwitchId(v string) *GetSwitchResponseBody {
	s.VSwitchId = &v
	return s
}

func (s *GetSwitchResponseBody) SetVSwitchName(v string) *GetSwitchResponseBody {
	s.VSwitchName = &v
	return s
}

func (s *GetSwitchResponseBody) SetCidrBlock(v string) *GetSwitchResponseBody {
	s.CidrBlock = &v
	return s
}

type GetSwitchResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSwitchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSwitchResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSwitchResponse) GoString() string {
	return s.String()
}

func (s *GetSwitchResponse) SetHeaders(v map[string]*string) *GetSwitchResponse {
	s.Headers = v
	return s
}

func (s *GetSwitchResponse) SetBody(v *GetSwitchResponseBody) *GetSwitchResponse {
	s.Body = v
	return s
}

type GetPodEventsRequest struct {
	// 返回的事件最大数量
	MaxEventsNum *int32 `json:"MaxEventsNum,omitempty" xml:"MaxEventsNum,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 截止时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s GetPodEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPodEventsRequest) GoString() string {
	return s.String()
}

func (s *GetPodEventsRequest) SetMaxEventsNum(v int32) *GetPodEventsRequest {
	s.MaxEventsNum = &v
	return s
}

func (s *GetPodEventsRequest) SetStartTime(v string) *GetPodEventsRequest {
	s.StartTime = &v
	return s
}

func (s *GetPodEventsRequest) SetEndTime(v string) *GetPodEventsRequest {
	s.EndTime = &v
	return s
}

type GetPodEventsResponseBody struct {
	// 作业ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 运行示例ID
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// 事件列表
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPodEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPodEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPodEventsResponseBody) SetJobId(v string) *GetPodEventsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetPodEventsResponseBody) SetPodId(v string) *GetPodEventsResponseBody {
	s.PodId = &v
	return s
}

func (s *GetPodEventsResponseBody) SetEvents(v []*string) *GetPodEventsResponseBody {
	s.Events = v
	return s
}

func (s *GetPodEventsResponseBody) SetRequestId(v string) *GetPodEventsResponseBody {
	s.RequestId = &v
	return s
}

type GetPodEventsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPodEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPodEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPodEventsResponse) GoString() string {
	return s.String()
}

func (s *GetPodEventsResponse) SetHeaders(v map[string]*string) *GetPodEventsResponse {
	s.Headers = v
	return s
}

func (s *GetPodEventsResponse) SetBody(v *GetPodEventsResponseBody) *GetPodEventsResponse {
	s.Body = v
	return s
}

type ListSwitchesRequest struct {
	// 所属VPC id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 取第几页的数据
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSwitchesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSwitchesRequest) GoString() string {
	return s.String()
}

func (s *ListSwitchesRequest) SetVpcId(v string) *ListSwitchesRequest {
	s.VpcId = &v
	return s
}

func (s *ListSwitchesRequest) SetPageNumber(v int32) *ListSwitchesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSwitchesRequest) SetPageSize(v int32) *ListSwitchesRequest {
	s.PageSize = &v
	return s
}

type ListSwitchesResponseBody struct {
	// 代码源配置列表
	Switches []*ListSwitchesResponseBodySwitches `json:"Switches,omitempty" xml:"Switches,omitempty" type:"Repeated"`
	// 符合过滤条件的代码源配置的总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSwitchesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSwitchesResponseBody) GoString() string {
	return s.String()
}

func (s *ListSwitchesResponseBody) SetSwitches(v []*ListSwitchesResponseBodySwitches) *ListSwitchesResponseBody {
	s.Switches = v
	return s
}

func (s *ListSwitchesResponseBody) SetTotalCount(v int64) *ListSwitchesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSwitchesResponseBody) SetRequestId(v string) *ListSwitchesResponseBody {
	s.RequestId = &v
	return s
}

type ListSwitchesResponseBodySwitches struct {
	// 交换机id
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// 交换机名称
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// 所属VPCid
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 网段
	CidrBlock *string `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
}

func (s ListSwitchesResponseBodySwitches) String() string {
	return tea.Prettify(s)
}

func (s ListSwitchesResponseBodySwitches) GoString() string {
	return s.String()
}

func (s *ListSwitchesResponseBodySwitches) SetVSwitchId(v string) *ListSwitchesResponseBodySwitches {
	s.VSwitchId = &v
	return s
}

func (s *ListSwitchesResponseBodySwitches) SetVSwitchName(v string) *ListSwitchesResponseBodySwitches {
	s.VSwitchName = &v
	return s
}

func (s *ListSwitchesResponseBodySwitches) SetVpcId(v string) *ListSwitchesResponseBodySwitches {
	s.VpcId = &v
	return s
}

func (s *ListSwitchesResponseBodySwitches) SetDescription(v string) *ListSwitchesResponseBodySwitches {
	s.Description = &v
	return s
}

func (s *ListSwitchesResponseBodySwitches) SetCidrBlock(v string) *ListSwitchesResponseBodySwitches {
	s.CidrBlock = &v
	return s
}

type ListSwitchesResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSwitchesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSwitchesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSwitchesResponse) GoString() string {
	return s.String()
}

func (s *ListSwitchesResponse) SetHeaders(v map[string]*string) *ListSwitchesResponse {
	s.Headers = v
	return s
}

func (s *ListSwitchesResponse) SetBody(v *ListSwitchesResponseBody) *ListSwitchesResponse {
	s.Body = v
	return s
}

type ListTensorboardsRequest struct {
	// 是否显示详情
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 展示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 根据状态过滤
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 截止时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 当前页
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的作业数
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// JobId
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// TensorboardId
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s ListTensorboardsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTensorboardsRequest) GoString() string {
	return s.String()
}

func (s *ListTensorboardsRequest) SetVerbose(v bool) *ListTensorboardsRequest {
	s.Verbose = &v
	return s
}

func (s *ListTensorboardsRequest) SetWorkspaceId(v string) *ListTensorboardsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListTensorboardsRequest) SetDisplayName(v string) *ListTensorboardsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListTensorboardsRequest) SetStatus(v string) *ListTensorboardsRequest {
	s.Status = &v
	return s
}

func (s *ListTensorboardsRequest) SetStartTime(v string) *ListTensorboardsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTensorboardsRequest) SetEndTime(v string) *ListTensorboardsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTensorboardsRequest) SetPageNumber(v int32) *ListTensorboardsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListTensorboardsRequest) SetPageSize(v int32) *ListTensorboardsRequest {
	s.PageSize = &v
	return s
}

func (s *ListTensorboardsRequest) SetSortBy(v string) *ListTensorboardsRequest {
	s.SortBy = &v
	return s
}

func (s *ListTensorboardsRequest) SetOrder(v string) *ListTensorboardsRequest {
	s.Order = &v
	return s
}

func (s *ListTensorboardsRequest) SetJobId(v string) *ListTensorboardsRequest {
	s.JobId = &v
	return s
}

func (s *ListTensorboardsRequest) SetTensorboardId(v string) *ListTensorboardsRequest {
	s.TensorboardId = &v
	return s
}

type ListTensorboardsResponseBody struct {
	// Tensorboard 列表
	Tensorboards []*Tensorboard `json:"Tensorboards,omitempty" xml:"Tensorboards,omitempty" type:"Repeated"`
	// 符合条件的数据源总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListTensorboardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTensorboardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTensorboardsResponseBody) SetTensorboards(v []*Tensorboard) *ListTensorboardsResponseBody {
	s.Tensorboards = v
	return s
}

func (s *ListTensorboardsResponseBody) SetTotalCount(v int64) *ListTensorboardsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTensorboardsResponseBody) SetRequestId(v string) *ListTensorboardsResponseBody {
	s.RequestId = &v
	return s
}

type ListTensorboardsResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTensorboardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTensorboardsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTensorboardsResponse) GoString() string {
	return s.String()
}

func (s *ListTensorboardsResponse) SetHeaders(v map[string]*string) *ListTensorboardsResponse {
	s.Headers = v
	return s
}

func (s *ListTensorboardsResponse) SetBody(v *ListTensorboardsResponseBody) *ListTensorboardsResponse {
	s.Body = v
	return s
}

type ListSecurityGroupsRequest struct {
	// 所属Vpc
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 取第几页的数据
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListSecurityGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsRequest) SetVpcId(v string) *ListSecurityGroupsRequest {
	s.VpcId = &v
	return s
}

func (s *ListSecurityGroupsRequest) SetPageNumber(v int32) *ListSecurityGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListSecurityGroupsRequest) SetPageSize(v int32) *ListSecurityGroupsRequest {
	s.PageSize = &v
	return s
}

type ListSecurityGroupsResponseBody struct {
	// 代码源配置列表
	SecurityGroups []*ListSecurityGroupsResponseBodySecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Repeated"`
	// 符合过滤条件的代码源配置的总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSecurityGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsResponseBody) SetSecurityGroups(v []*ListSecurityGroupsResponseBodySecurityGroups) *ListSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

func (s *ListSecurityGroupsResponseBody) SetTotalCount(v int64) *ListSecurityGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSecurityGroupsResponseBody) SetRequestId(v string) *ListSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListSecurityGroupsResponseBodySecurityGroups struct {
	// 安全组Id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	// 所属VPC ID
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ListSecurityGroupsResponseBodySecurityGroups) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsResponseBodySecurityGroups) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsResponseBodySecurityGroups) SetSecurityGroupId(v string) *ListSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroupId = &v
	return s
}

func (s *ListSecurityGroupsResponseBodySecurityGroups) SetSecurityGroupName(v string) *ListSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroupName = &v
	return s
}

func (s *ListSecurityGroupsResponseBodySecurityGroups) SetVpcId(v string) *ListSecurityGroupsResponseBodySecurityGroups {
	s.VpcId = &v
	return s
}

func (s *ListSecurityGroupsResponseBodySecurityGroups) SetDescription(v string) *ListSecurityGroupsResponseBodySecurityGroups {
	s.Description = &v
	return s
}

type ListSecurityGroupsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListSecurityGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListSecurityGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListSecurityGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListSecurityGroupsResponse) SetHeaders(v map[string]*string) *ListSecurityGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListSecurityGroupsResponse) SetBody(v *ListSecurityGroupsResponseBody) *ListSecurityGroupsResponse {
	s.Body = v
	return s
}

type GetPodLogsRequest struct {
	// 返回的日志的最大行数，默认值：2000。
	MaxLines *int32 `json:"MaxLines,omitempty" xml:"MaxLines,omitempty"`
	// 查询的起始时间，默认值：7天前。
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 查询的截止时间，默认值：当前。
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 是否下载日志文件，默认：false。
	DownloadToFile *bool `json:"DownloadToFile,omitempty" xml:"DownloadToFile,omitempty"`
}

func (s GetPodLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPodLogsRequest) GoString() string {
	return s.String()
}

func (s *GetPodLogsRequest) SetMaxLines(v int32) *GetPodLogsRequest {
	s.MaxLines = &v
	return s
}

func (s *GetPodLogsRequest) SetStartTime(v string) *GetPodLogsRequest {
	s.StartTime = &v
	return s
}

func (s *GetPodLogsRequest) SetEndTime(v string) *GetPodLogsRequest {
	s.EndTime = &v
	return s
}

func (s *GetPodLogsRequest) SetDownloadToFile(v bool) *GetPodLogsRequest {
	s.DownloadToFile = &v
	return s
}

type GetPodLogsResponseBody struct {
	// 作业ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 实例ID
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// 日志列表
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPodLogsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPodLogsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPodLogsResponseBody) SetJobId(v string) *GetPodLogsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetPodLogsResponseBody) SetPodId(v string) *GetPodLogsResponseBody {
	s.PodId = &v
	return s
}

func (s *GetPodLogsResponseBody) SetLogs(v []*string) *GetPodLogsResponseBody {
	s.Logs = v
	return s
}

func (s *GetPodLogsResponseBody) SetRequestId(v string) *GetPodLogsResponseBody {
	s.RequestId = &v
	return s
}

type GetPodLogsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPodLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPodLogsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPodLogsResponse) GoString() string {
	return s.String()
}

func (s *GetPodLogsResponse) SetHeaders(v map[string]*string) *GetPodLogsResponse {
	s.Headers = v
	return s
}

func (s *GetPodLogsResponse) SetBody(v *GetPodLogsResponseBody) *GetPodLogsResponse {
	s.Body = v
	return s
}

type ListVpcsRequest struct {
	// 取第几页的数据
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListVpcsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVpcsRequest) GoString() string {
	return s.String()
}

func (s *ListVpcsRequest) SetPageNumber(v int32) *ListVpcsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListVpcsRequest) SetPageSize(v int32) *ListVpcsRequest {
	s.PageSize = &v
	return s
}

type ListVpcsResponseBody struct {
	// 代码源配置列表
	Vpcs []*ListVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Repeated"`
	// 符合过滤条件的代码源配置的总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListVpcsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVpcsResponseBody) SetVpcs(v []*ListVpcsResponseBodyVpcs) *ListVpcsResponseBody {
	s.Vpcs = v
	return s
}

func (s *ListVpcsResponseBody) SetTotalCount(v int64) *ListVpcsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVpcsResponseBody) SetRequestId(v string) *ListVpcsResponseBody {
	s.RequestId = &v
	return s
}

type ListVpcsResponseBodyVpcs struct {
	// vpc的id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// VPC的名称。
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s ListVpcsResponseBodyVpcs) String() string {
	return tea.Prettify(s)
}

func (s ListVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *ListVpcsResponseBodyVpcs) SetVpcId(v string) *ListVpcsResponseBodyVpcs {
	s.VpcId = &v
	return s
}

func (s *ListVpcsResponseBodyVpcs) SetVpcName(v string) *ListVpcsResponseBodyVpcs {
	s.VpcName = &v
	return s
}

type ListVpcsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListVpcsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListVpcsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVpcsResponse) GoString() string {
	return s.String()
}

func (s *ListVpcsResponse) SetHeaders(v map[string]*string) *ListVpcsResponse {
	s.Headers = v
	return s
}

func (s *ListVpcsResponse) SetBody(v *ListVpcsResponseBody) *ListVpcsResponse {
	s.Body = v
	return s
}

type StopTensorboardRequest struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s StopTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s StopTensorboardRequest) GoString() string {
	return s.String()
}

func (s *StopTensorboardRequest) SetWorkspaceId(v string) *StopTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type StopTensorboardResponseBody struct {
	// Tensorboad Id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *StopTensorboardResponseBody) SetTensorboardId(v string) *StopTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

func (s *StopTensorboardResponseBody) SetRequestId(v string) *StopTensorboardResponseBody {
	s.RequestId = &v
	return s
}

type StopTensorboardResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s StopTensorboardResponse) GoString() string {
	return s.String()
}

func (s *StopTensorboardResponse) SetHeaders(v map[string]*string) *StopTensorboardResponse {
	s.Headers = v
	return s
}

func (s *StopTensorboardResponse) SetBody(v *StopTensorboardResponseBody) *StopTensorboardResponse {
	s.Body = v
	return s
}

type CreateJobRequest struct {
	// 作业显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 作业类型
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// 作业规格配置
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// 作业命令
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// 数据源配置列表
	DataSources []*CreateJobRequestDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// 代码源配置
	CodeSource *CreateJobRequestCodeSource `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	// 用户VPC
	UserVpc *CreateJobRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// 三方库配置列表
	ThirdpartyLibs []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// 三方库(requirements.txt)文件路径
	ThirdpartyLibDir *string `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	// 环境变量配置
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// 作业最大运行时间
	JobMaxRunningTimeMinutes *int64 `json:"JobMaxRunningTimeMinutes,omitempty" xml:"JobMaxRunningTimeMinutes,omitempty"`
	// 工作空间编号
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 资源组编号
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 作业优先级
	Priority    *int32          `json:"Priority,omitempty" xml:"Priority,omitempty"`
	Settings    *JobSettings    `json:"Settings,omitempty" xml:"Settings,omitempty"`
	ElasticSpec *JobElasticSpec `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetDisplayName(v string) *CreateJobRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateJobRequest) SetJobType(v string) *CreateJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateJobRequest) SetJobSpecs(v []*JobSpec) *CreateJobRequest {
	s.JobSpecs = v
	return s
}

func (s *CreateJobRequest) SetUserCommand(v string) *CreateJobRequest {
	s.UserCommand = &v
	return s
}

func (s *CreateJobRequest) SetDataSources(v []*CreateJobRequestDataSources) *CreateJobRequest {
	s.DataSources = v
	return s
}

func (s *CreateJobRequest) SetCodeSource(v *CreateJobRequestCodeSource) *CreateJobRequest {
	s.CodeSource = v
	return s
}

func (s *CreateJobRequest) SetUserVpc(v *CreateJobRequestUserVpc) *CreateJobRequest {
	s.UserVpc = v
	return s
}

func (s *CreateJobRequest) SetThirdpartyLibs(v []*string) *CreateJobRequest {
	s.ThirdpartyLibs = v
	return s
}

func (s *CreateJobRequest) SetThirdpartyLibDir(v string) *CreateJobRequest {
	s.ThirdpartyLibDir = &v
	return s
}

func (s *CreateJobRequest) SetEnvs(v map[string]*string) *CreateJobRequest {
	s.Envs = v
	return s
}

func (s *CreateJobRequest) SetJobMaxRunningTimeMinutes(v int64) *CreateJobRequest {
	s.JobMaxRunningTimeMinutes = &v
	return s
}

func (s *CreateJobRequest) SetWorkspaceId(v string) *CreateJobRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateJobRequest) SetResourceId(v string) *CreateJobRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateJobRequest) SetPriority(v int32) *CreateJobRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobRequest) SetSettings(v *JobSettings) *CreateJobRequest {
	s.Settings = v
	return s
}

func (s *CreateJobRequest) SetElasticSpec(v *JobElasticSpec) *CreateJobRequest {
	s.ElasticSpec = v
	return s
}

type CreateJobRequestDataSources struct {
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 本地挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s CreateJobRequestDataSources) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestDataSources) GoString() string {
	return s.String()
}

func (s *CreateJobRequestDataSources) SetDataSourceId(v string) *CreateJobRequestDataSources {
	s.DataSourceId = &v
	return s
}

func (s *CreateJobRequestDataSources) SetMountPath(v string) *CreateJobRequestDataSources {
	s.MountPath = &v
	return s
}

type CreateJobRequestCodeSource struct {
	// 代码源Id
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// 代码分支
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// 代码Commit
	Commit *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// 本地挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s CreateJobRequestCodeSource) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestCodeSource) GoString() string {
	return s.String()
}

func (s *CreateJobRequestCodeSource) SetCodeSourceId(v string) *CreateJobRequestCodeSource {
	s.CodeSourceId = &v
	return s
}

func (s *CreateJobRequestCodeSource) SetBranch(v string) *CreateJobRequestCodeSource {
	s.Branch = &v
	return s
}

func (s *CreateJobRequestCodeSource) SetCommit(v string) *CreateJobRequestCodeSource {
	s.Commit = &v
	return s
}

func (s *CreateJobRequestCodeSource) SetMountPath(v string) *CreateJobRequestCodeSource {
	s.MountPath = &v
	return s
}

type CreateJobRequestUserVpc struct {
	// 用户VPC的id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 用户交换机的id
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// 用户安全组的id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 扩展网段
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
}

func (s CreateJobRequestUserVpc) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateJobRequestUserVpc) SetVpcId(v string) *CreateJobRequestUserVpc {
	s.VpcId = &v
	return s
}

func (s *CreateJobRequestUserVpc) SetSwitchId(v string) *CreateJobRequestUserVpc {
	s.SwitchId = &v
	return s
}

func (s *CreateJobRequestUserVpc) SetSecurityGroupId(v string) *CreateJobRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateJobRequestUserVpc) SetExtendedCIDRs(v []*string) *CreateJobRequestUserVpc {
	s.ExtendedCIDRs = v
	return s
}

type CreateJobResponseBody struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateJobResponseBody) SetJobId(v string) *CreateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateJobResponseBody) SetRequestId(v string) *CreateJobResponseBody {
	s.RequestId = &v
	return s
}

type CreateJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateJobResponse) GoString() string {
	return s.String()
}

func (s *CreateJobResponse) SetHeaders(v map[string]*string) *CreateJobResponse {
	s.Headers = v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

type ListCodeSourcesRequest struct {
	// 代码源显示名称，支持模糊匹配
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 取第几页的数据
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 用于排序的字段名，可选字段名：'DisplayName' 'GmtCreateTime' 'GmtModifyTime'
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 排序顺序, 枚举值 desc 或者 asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListCodeSourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListCodeSourcesRequest) GoString() string {
	return s.String()
}

func (s *ListCodeSourcesRequest) SetDisplayName(v string) *ListCodeSourcesRequest {
	s.DisplayName = &v
	return s
}

func (s *ListCodeSourcesRequest) SetPageNumber(v int32) *ListCodeSourcesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListCodeSourcesRequest) SetPageSize(v int32) *ListCodeSourcesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCodeSourcesRequest) SetSortBy(v string) *ListCodeSourcesRequest {
	s.SortBy = &v
	return s
}

func (s *ListCodeSourcesRequest) SetOrder(v string) *ListCodeSourcesRequest {
	s.Order = &v
	return s
}

type ListCodeSourcesResponseBody struct {
	// 代码源配置列表
	CodeSources []*CodeSourceItem `json:"CodeSources,omitempty" xml:"CodeSources,omitempty" type:"Repeated"`
	// 符合过滤条件的代码源配置的总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListCodeSourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListCodeSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListCodeSourcesResponseBody) SetCodeSources(v []*CodeSourceItem) *ListCodeSourcesResponseBody {
	s.CodeSources = v
	return s
}

func (s *ListCodeSourcesResponseBody) SetTotalCount(v int64) *ListCodeSourcesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListCodeSourcesResponseBody) SetRequestId(v string) *ListCodeSourcesResponseBody {
	s.RequestId = &v
	return s
}

type ListCodeSourcesResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListCodeSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListCodeSourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListCodeSourcesResponse) GoString() string {
	return s.String()
}

func (s *ListCodeSourcesResponse) SetHeaders(v map[string]*string) *ListCodeSourcesResponse {
	s.Headers = v
	return s
}

func (s *ListCodeSourcesResponse) SetBody(v *ListCodeSourcesResponseBody) *ListCodeSourcesResponse {
	s.Body = v
	return s
}

type JobDispatchSubmitRequest struct {
	// PAI-Xflow名称
	AlgoName *string `json:"AlgoName,omitempty" xml:"AlgoName,omitempty"`
	// PAI-project名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// properties of pai command
	Properties map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// odps settings
	Settings map[string]*string `json:"Settings,omitempty" xml:"Settings,omitempty"`
}

func (s JobDispatchSubmitRequest) String() string {
	return tea.Prettify(s)
}

func (s JobDispatchSubmitRequest) GoString() string {
	return s.String()
}

func (s *JobDispatchSubmitRequest) SetAlgoName(v string) *JobDispatchSubmitRequest {
	s.AlgoName = &v
	return s
}

func (s *JobDispatchSubmitRequest) SetProjectName(v string) *JobDispatchSubmitRequest {
	s.ProjectName = &v
	return s
}

func (s *JobDispatchSubmitRequest) SetProperties(v map[string]*string) *JobDispatchSubmitRequest {
	s.Properties = v
	return s
}

func (s *JobDispatchSubmitRequest) SetSettings(v map[string]*string) *JobDispatchSubmitRequest {
	s.Settings = v
	return s
}

type JobDispatchSubmitResponseBody struct {
	// 作业Url
	JobUrl *string `json:"JobUrl,omitempty" xml:"JobUrl,omitempty"`
}

func (s JobDispatchSubmitResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JobDispatchSubmitResponseBody) GoString() string {
	return s.String()
}

func (s *JobDispatchSubmitResponseBody) SetJobUrl(v string) *JobDispatchSubmitResponseBody {
	s.JobUrl = &v
	return s
}

type JobDispatchSubmitResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *JobDispatchSubmitResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JobDispatchSubmitResponse) String() string {
	return tea.Prettify(s)
}

func (s JobDispatchSubmitResponse) GoString() string {
	return s.String()
}

func (s *JobDispatchSubmitResponse) SetHeaders(v map[string]*string) *JobDispatchSubmitResponse {
	s.Headers = v
	return s
}

func (s *JobDispatchSubmitResponse) SetBody(v *JobDispatchSubmitResponseBody) *JobDispatchSubmitResponse {
	s.Body = v
	return s
}

type ListEcsSpecsRequest struct {
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 查询第几页数据
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 设置查询的分页大写
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListEcsSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsRequest) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsRequest) SetSortBy(v string) *ListEcsSpecsRequest {
	s.SortBy = &v
	return s
}

func (s *ListEcsSpecsRequest) SetOrder(v string) *ListEcsSpecsRequest {
	s.Order = &v
	return s
}

func (s *ListEcsSpecsRequest) SetPageNumber(v int32) *ListEcsSpecsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListEcsSpecsRequest) SetPageSize(v int32) *ListEcsSpecsRequest {
	s.PageSize = &v
	return s
}

type ListEcsSpecsResponseBody struct {
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// ECS规格列表
	EcsSpecs []*EcsSpec `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// 符合过滤条件的总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEcsSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBody) SetRequestId(v string) *ListEcsSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetEcsSpecs(v []*EcsSpec) *ListEcsSpecsResponseBody {
	s.EcsSpecs = v
	return s
}

func (s *ListEcsSpecsResponseBody) SetTotalCount(v int64) *ListEcsSpecsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEcsSpecsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListEcsSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListEcsSpecsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponse) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponse) SetHeaders(v map[string]*string) *ListEcsSpecsResponse {
	s.Headers = v
	return s
}

func (s *ListEcsSpecsResponse) SetBody(v *ListEcsSpecsResponseBody) *ListEcsSpecsResponse {
	s.Body = v
	return s
}

type DeleteQuotaResponseBody struct {
	// 资源配额id
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQuotaResponseBody) SetQuotaId(v string) *DeleteQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *DeleteQuotaResponseBody) SetRequestId(v string) *DeleteQuotaResponseBody {
	s.RequestId = &v
	return s
}

type DeleteQuotaResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteQuotaResponse) GoString() string {
	return s.String()
}

func (s *DeleteQuotaResponse) SetHeaders(v map[string]*string) *DeleteQuotaResponse {
	s.Headers = v
	return s
}

func (s *DeleteQuotaResponse) SetBody(v *DeleteQuotaResponseBody) *DeleteQuotaResponse {
	s.Body = v
	return s
}

type StartTensorboardRequest struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s StartTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s StartTensorboardRequest) GoString() string {
	return s.String()
}

func (s *StartTensorboardRequest) SetWorkspaceId(v string) *StartTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type StartTensorboardResponseBody struct {
	// Tensorboad Id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *StartTensorboardResponseBody) SetTensorboardId(v string) *StartTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

func (s *StartTensorboardResponseBody) SetRequestId(v string) *StartTensorboardResponseBody {
	s.RequestId = &v
	return s
}

type StartTensorboardResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StartTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StartTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s StartTensorboardResponse) GoString() string {
	return s.String()
}

func (s *StartTensorboardResponse) SetHeaders(v map[string]*string) *StartTensorboardResponse {
	s.Headers = v
	return s
}

func (s *StartTensorboardResponse) SetBody(v *StartTensorboardResponseBody) *StartTensorboardResponse {
	s.Body = v
	return s
}

type GetTensorboardRequest struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// JodId
	JodId *string `json:"JodId,omitempty" xml:"JodId,omitempty"`
}

func (s GetTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTensorboardRequest) GoString() string {
	return s.String()
}

func (s *GetTensorboardRequest) SetWorkspaceId(v string) *GetTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

func (s *GetTensorboardRequest) SetJodId(v string) *GetTensorboardRequest {
	s.JodId = &v
	return s
}

type GetTensorboardResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *Tensorboard       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTensorboardResponse) GoString() string {
	return s.String()
}

func (s *GetTensorboardResponse) SetHeaders(v map[string]*string) *GetTensorboardResponse {
	s.Headers = v
	return s
}

func (s *GetTensorboardResponse) SetBody(v *Tensorboard) *GetTensorboardResponse {
	s.Body = v
	return s
}

type GetWorkspaceResponseBody struct {
	// 工作空间
	Workspace *Workspace `json:"Workspace,omitempty" xml:"Workspace,omitempty"`
	// 是否是当前工作空间的管理员
	IsWorkspaceAdmin *bool `json:"IsWorkspaceAdmin,omitempty" xml:"IsWorkspaceAdmin,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetWorkspaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponseBody) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponseBody) SetWorkspace(v *Workspace) *GetWorkspaceResponseBody {
	s.Workspace = v
	return s
}

func (s *GetWorkspaceResponseBody) SetIsWorkspaceAdmin(v bool) *GetWorkspaceResponseBody {
	s.IsWorkspaceAdmin = &v
	return s
}

func (s *GetWorkspaceResponseBody) SetRequestId(v string) *GetWorkspaceResponseBody {
	s.RequestId = &v
	return s
}

type GetWorkspaceResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetWorkspaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetWorkspaceResponse) String() string {
	return tea.Prettify(s)
}

func (s GetWorkspaceResponse) GoString() string {
	return s.String()
}

func (s *GetWorkspaceResponse) SetHeaders(v map[string]*string) *GetWorkspaceResponse {
	s.Headers = v
	return s
}

func (s *GetWorkspaceResponse) SetBody(v *GetWorkspaceResponseBody) *GetWorkspaceResponse {
	s.Body = v
	return s
}

type JobDispatchQueryRequest struct {
	// PAI-Xflow名称
	AlgoName *string `json:"AlgoName,omitempty" xml:"AlgoName,omitempty"`
	// PAI-project名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// properties of pai command
	Properties map[string]*string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// odps settings
	Settings map[string]*string `json:"Settings,omitempty" xml:"Settings,omitempty"`
}

func (s JobDispatchQueryRequest) String() string {
	return tea.Prettify(s)
}

func (s JobDispatchQueryRequest) GoString() string {
	return s.String()
}

func (s *JobDispatchQueryRequest) SetAlgoName(v string) *JobDispatchQueryRequest {
	s.AlgoName = &v
	return s
}

func (s *JobDispatchQueryRequest) SetProjectName(v string) *JobDispatchQueryRequest {
	s.ProjectName = &v
	return s
}

func (s *JobDispatchQueryRequest) SetProperties(v map[string]*string) *JobDispatchQueryRequest {
	s.Properties = v
	return s
}

func (s *JobDispatchQueryRequest) SetSettings(v map[string]*string) *JobDispatchQueryRequest {
	s.Settings = v
	return s
}

type JobDispatchQueryShrinkRequest struct {
	// PAI-Xflow名称
	AlgoName *string `json:"AlgoName,omitempty" xml:"AlgoName,omitempty"`
	// PAI-project名称
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// properties of pai command
	PropertiesShrink *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// odps settings
	SettingsShrink *string `json:"Settings,omitempty" xml:"Settings,omitempty"`
}

func (s JobDispatchQueryShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s JobDispatchQueryShrinkRequest) GoString() string {
	return s.String()
}

func (s *JobDispatchQueryShrinkRequest) SetAlgoName(v string) *JobDispatchQueryShrinkRequest {
	s.AlgoName = &v
	return s
}

func (s *JobDispatchQueryShrinkRequest) SetProjectName(v string) *JobDispatchQueryShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *JobDispatchQueryShrinkRequest) SetPropertiesShrink(v string) *JobDispatchQueryShrinkRequest {
	s.PropertiesShrink = &v
	return s
}

func (s *JobDispatchQueryShrinkRequest) SetSettingsShrink(v string) *JobDispatchQueryShrinkRequest {
	s.SettingsShrink = &v
	return s
}

type JobDispatchQueryResponseBody struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s JobDispatchQueryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s JobDispatchQueryResponseBody) GoString() string {
	return s.String()
}

func (s *JobDispatchQueryResponseBody) SetJobId(v string) *JobDispatchQueryResponseBody {
	s.JobId = &v
	return s
}

func (s *JobDispatchQueryResponseBody) SetRequestId(v string) *JobDispatchQueryResponseBody {
	s.RequestId = &v
	return s
}

type JobDispatchQueryResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *JobDispatchQueryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s JobDispatchQueryResponse) String() string {
	return tea.Prettify(s)
}

func (s JobDispatchQueryResponse) GoString() string {
	return s.String()
}

func (s *JobDispatchQueryResponse) SetHeaders(v map[string]*string) *JobDispatchQueryResponse {
	s.Headers = v
	return s
}

func (s *JobDispatchQueryResponse) SetBody(v *JobDispatchQueryResponseBody) *JobDispatchQueryResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	// 作业显示名称，支持模糊查询
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 作业类型
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// 作业状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 截止时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 当前页
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的作业数
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 是否只返回当前登录者所提交的作业
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 作业关联用户ID
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// 调用方
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// 自定义标签
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// 工作流ID
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetDisplayName(v string) *ListJobsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListJobsRequest) SetJobType(v string) *ListJobsRequest {
	s.JobType = &v
	return s
}

func (s *ListJobsRequest) SetStatus(v string) *ListJobsRequest {
	s.Status = &v
	return s
}

func (s *ListJobsRequest) SetStartTime(v string) *ListJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobsRequest) SetEndTime(v string) *ListJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobsRequest) SetPageNumber(v int32) *ListJobsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsRequest) SetPageSize(v int32) *ListJobsRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsRequest) SetSortBy(v string) *ListJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListJobsRequest) SetOrder(v string) *ListJobsRequest {
	s.Order = &v
	return s
}

func (s *ListJobsRequest) SetShowOwn(v bool) *ListJobsRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListJobsRequest) SetWorkspaceId(v string) *ListJobsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListJobsRequest) SetResourceId(v string) *ListJobsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListJobsRequest) SetBusinessUserId(v string) *ListJobsRequest {
	s.BusinessUserId = &v
	return s
}

func (s *ListJobsRequest) SetCaller(v string) *ListJobsRequest {
	s.Caller = &v
	return s
}

func (s *ListJobsRequest) SetTags(v map[string]*string) *ListJobsRequest {
	s.Tags = v
	return s
}

func (s *ListJobsRequest) SetPipelineId(v string) *ListJobsRequest {
	s.PipelineId = &v
	return s
}

type ListJobsShrinkRequest struct {
	// 作业显示名称，支持模糊查询
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 作业类型
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// 作业状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 截止时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 当前页
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的作业数
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 是否只返回当前登录者所提交的作业
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 作业关联用户ID
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// 调用方
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// 自定义标签
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// 工作流ID
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
}

func (s ListJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobsShrinkRequest) SetDisplayName(v string) *ListJobsShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *ListJobsShrinkRequest) SetJobType(v string) *ListJobsShrinkRequest {
	s.JobType = &v
	return s
}

func (s *ListJobsShrinkRequest) SetStatus(v string) *ListJobsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListJobsShrinkRequest) SetStartTime(v string) *ListJobsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobsShrinkRequest) SetEndTime(v string) *ListJobsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageNumber(v int32) *ListJobsShrinkRequest {
	s.PageNumber = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPageSize(v int32) *ListJobsShrinkRequest {
	s.PageSize = &v
	return s
}

func (s *ListJobsShrinkRequest) SetSortBy(v string) *ListJobsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListJobsShrinkRequest) SetOrder(v string) *ListJobsShrinkRequest {
	s.Order = &v
	return s
}

func (s *ListJobsShrinkRequest) SetShowOwn(v bool) *ListJobsShrinkRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListJobsShrinkRequest) SetWorkspaceId(v string) *ListJobsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetResourceId(v string) *ListJobsShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetBusinessUserId(v string) *ListJobsShrinkRequest {
	s.BusinessUserId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetCaller(v string) *ListJobsShrinkRequest {
	s.Caller = &v
	return s
}

func (s *ListJobsShrinkRequest) SetTagsShrink(v string) *ListJobsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListJobsShrinkRequest) SetPipelineId(v string) *ListJobsShrinkRequest {
	s.PipelineId = &v
	return s
}

type ListJobsResponseBody struct {
	// 作业列表
	Jobs []*JobItem `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// 符合过滤条件的总作业数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListJobsResponseBody) SetJobs(v []*JobItem) *ListJobsResponseBody {
	s.Jobs = v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int64) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

type ListJobsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListJobsResponse) GoString() string {
	return s.String()
}

func (s *ListJobsResponse) SetHeaders(v map[string]*string) *ListJobsResponse {
	s.Headers = v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type GetVpcResponseBody struct {
	// Vpc的ID
	VpcId *int32 `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Vpc名称
	VpcName *string `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVpcResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVpcResponseBody) GoString() string {
	return s.String()
}

func (s *GetVpcResponseBody) SetVpcId(v int32) *GetVpcResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetVpcResponseBody) SetVpcName(v string) *GetVpcResponseBody {
	s.VpcName = &v
	return s
}

func (s *GetVpcResponseBody) SetRequestId(v string) *GetVpcResponseBody {
	s.RequestId = &v
	return s
}

type GetVpcResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVpcResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVpcResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVpcResponse) GoString() string {
	return s.String()
}

func (s *GetVpcResponse) SetHeaders(v map[string]*string) *GetVpcResponse {
	s.Headers = v
	return s
}

func (s *GetVpcResponse) SetBody(v *GetVpcResponseBody) *GetVpcResponse {
	s.Body = v
	return s
}

type CreateCodeSourceRequest struct {
	// 代码源配置名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 代码源详细描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代码仓库地址
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// 代码分支
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// 代码本地挂载目录，默认挂载到/root/code/下
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 用户名
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// 密码
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
}

func (s CreateCodeSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCodeSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateCodeSourceRequest) SetDisplayName(v string) *CreateCodeSourceRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateCodeSourceRequest) SetDescription(v string) *CreateCodeSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepo(v string) *CreateCodeSourceRequest {
	s.CodeRepo = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeBranch(v string) *CreateCodeSourceRequest {
	s.CodeBranch = &v
	return s
}

func (s *CreateCodeSourceRequest) SetMountPath(v string) *CreateCodeSourceRequest {
	s.MountPath = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepoUserName(v string) *CreateCodeSourceRequest {
	s.CodeRepoUserName = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepoAccessToken(v string) *CreateCodeSourceRequest {
	s.CodeRepoAccessToken = &v
	return s
}

type CreateCodeSourceResponseBody struct {
	// 创建的代码源配置的ID
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCodeSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCodeSourceResponseBody) SetCodeSourceId(v string) *CreateCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *CreateCodeSourceResponseBody) SetRequestId(v string) *CreateCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateCodeSourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateCodeSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateCodeSourceResponse) SetHeaders(v map[string]*string) *CreateCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateCodeSourceResponse) SetBody(v *CreateCodeSourceResponseBody) *CreateCodeSourceResponse {
	s.Body = v
	return s
}

type GetJobMetricsRequest struct {
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 截止时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 时间间隔
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
	// 指标类型
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// Token
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetJobMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetJobMetricsRequest) SetStartTime(v string) *GetJobMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *GetJobMetricsRequest) SetEndTime(v string) *GetJobMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJobMetricsRequest) SetTimeStep(v string) *GetJobMetricsRequest {
	s.TimeStep = &v
	return s
}

func (s *GetJobMetricsRequest) SetMetricType(v string) *GetJobMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *GetJobMetricsRequest) SetToken(v string) *GetJobMetricsRequest {
	s.Token = &v
	return s
}

type GetJobMetricsResponseBody struct {
	// 作业ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 任务监控指标序列集合
	PodMetrics []*PodMetric `json:"PodMetrics,omitempty" xml:"PodMetrics,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobMetricsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobMetricsResponseBody) SetJobId(v string) *GetJobMetricsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobMetricsResponseBody) SetPodMetrics(v []*PodMetric) *GetJobMetricsResponseBody {
	s.PodMetrics = v
	return s
}

func (s *GetJobMetricsResponseBody) SetRequestId(v string) *GetJobMetricsResponseBody {
	s.RequestId = &v
	return s
}

type GetJobMetricsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobMetricsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetJobMetricsResponse) SetHeaders(v map[string]*string) *GetJobMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetJobMetricsResponse) SetBody(v *GetJobMetricsResponseBody) *GetJobMetricsResponse {
	s.Body = v
	return s
}

type GetJobResponseBody struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 作业类型
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// 作业显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 作业提交人Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 作业状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 作业所属工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 作业所属工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
	// 作业运行所在的资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 状态详情码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 状态详情
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 作业规格配置
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// 用户命令
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// 数据源配置列表
	DataSources []*GetJobResponseBodyDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// 代码源配置
	CodeSource *GetJobResponseBodyCodeSource `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	// 三方库配置列表
	ThirdpartyLibs []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// 三方库(requirements.txt)文件路径
	ThirdpartyLibDir *string `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	// 环境变量配置
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// 作业创建时间（UTC）
	GmtCreateTime    *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtSubmittedTime *string `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	GmtRunningTime   *string `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	GmtSuccessedTime *string `json:"GmtSuccessedTime,omitempty" xml:"GmtSuccessedTime,omitempty"`
	GmtStoppedTime   *string `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	GmtFailedTime    *string `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	// 作业结束时间（UTC）
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// 作业运行时长（s）
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 作业所以运行Pod列表
	Pods []*GetJobResponseBodyPods `json:"Pods,omitempty" xml:"Pods,omitempty" type:"Repeated"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 作业额外参数配置
	Settings *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 弹性任务参数
	ElasticSpec *JobElasticSpec `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetJobId(v string) *GetJobResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobResponseBody) SetJobType(v string) *GetJobResponseBody {
	s.JobType = &v
	return s
}

func (s *GetJobResponseBody) SetDisplayName(v string) *GetJobResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetJobResponseBody) SetUserId(v string) *GetJobResponseBody {
	s.UserId = &v
	return s
}

func (s *GetJobResponseBody) SetStatus(v string) *GetJobResponseBody {
	s.Status = &v
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

func (s *GetJobResponseBody) SetResourceId(v string) *GetJobResponseBody {
	s.ResourceId = &v
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

func (s *GetJobResponseBody) SetJobSpecs(v []*JobSpec) *GetJobResponseBody {
	s.JobSpecs = v
	return s
}

func (s *GetJobResponseBody) SetUserCommand(v string) *GetJobResponseBody {
	s.UserCommand = &v
	return s
}

func (s *GetJobResponseBody) SetDataSources(v []*GetJobResponseBodyDataSources) *GetJobResponseBody {
	s.DataSources = v
	return s
}

func (s *GetJobResponseBody) SetCodeSource(v *GetJobResponseBodyCodeSource) *GetJobResponseBody {
	s.CodeSource = v
	return s
}

func (s *GetJobResponseBody) SetThirdpartyLibs(v []*string) *GetJobResponseBody {
	s.ThirdpartyLibs = v
	return s
}

func (s *GetJobResponseBody) SetThirdpartyLibDir(v string) *GetJobResponseBody {
	s.ThirdpartyLibDir = &v
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

func (s *GetJobResponseBody) SetGmtSubmittedTime(v string) *GetJobResponseBody {
	s.GmtSubmittedTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtRunningTime(v string) *GetJobResponseBody {
	s.GmtRunningTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtSuccessedTime(v string) *GetJobResponseBody {
	s.GmtSuccessedTime = &v
	return s
}

func (s *GetJobResponseBody) SetGmtStoppedTime(v string) *GetJobResponseBody {
	s.GmtStoppedTime = &v
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

func (s *GetJobResponseBody) SetDuration(v int64) *GetJobResponseBody {
	s.Duration = &v
	return s
}

func (s *GetJobResponseBody) SetPods(v []*GetJobResponseBodyPods) *GetJobResponseBody {
	s.Pods = v
	return s
}

func (s *GetJobResponseBody) SetRequestId(v string) *GetJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJobResponseBody) SetSettings(v *JobSettings) *GetJobResponseBody {
	s.Settings = v
	return s
}

func (s *GetJobResponseBody) SetClusterId(v string) *GetJobResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetJobResponseBody) SetElasticSpec(v *JobElasticSpec) *GetJobResponseBody {
	s.ElasticSpec = v
	return s
}

type GetJobResponseBodyDataSources struct {
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 本地挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s GetJobResponseBodyDataSources) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyDataSources) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyDataSources) SetDataSourceId(v string) *GetJobResponseBodyDataSources {
	s.DataSourceId = &v
	return s
}

func (s *GetJobResponseBodyDataSources) SetMountPath(v string) *GetJobResponseBodyDataSources {
	s.MountPath = &v
	return s
}

type GetJobResponseBodyCodeSource struct {
	// 代码源Id
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// 代码分支
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// 代码Commit
	Commit *string `json:"Commit,omitempty" xml:"Commit,omitempty"`
	// 本地挂载路径
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s GetJobResponseBodyCodeSource) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyCodeSource) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyCodeSource) SetCodeSourceId(v string) *GetJobResponseBodyCodeSource {
	s.CodeSourceId = &v
	return s
}

func (s *GetJobResponseBodyCodeSource) SetBranch(v string) *GetJobResponseBodyCodeSource {
	s.Branch = &v
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

type GetJobResponseBodyPods struct {
	// Pod类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// Pod Id
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// Pod状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Pod Ip
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// Pod创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Pod启动时间（UTC）
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// Pod结束时间（UTC）
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
}

func (s GetJobResponseBodyPods) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyPods) GoString() string {
	return s.String()
}

func (s *GetJobResponseBodyPods) SetType(v string) *GetJobResponseBodyPods {
	s.Type = &v
	return s
}

func (s *GetJobResponseBodyPods) SetPodId(v string) *GetJobResponseBodyPods {
	s.PodId = &v
	return s
}

func (s *GetJobResponseBodyPods) SetStatus(v string) *GetJobResponseBodyPods {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyPods) SetIp(v string) *GetJobResponseBodyPods {
	s.Ip = &v
	return s
}

func (s *GetJobResponseBodyPods) SetGmtCreateTime(v string) *GetJobResponseBodyPods {
	s.GmtCreateTime = &v
	return s
}

func (s *GetJobResponseBodyPods) SetGmtStartTime(v string) *GetJobResponseBodyPods {
	s.GmtStartTime = &v
	return s
}

func (s *GetJobResponseBodyPods) SetGmtFinishTime(v string) *GetJobResponseBodyPods {
	s.GmtFinishTime = &v
	return s
}

type GetJobResponse struct {
	Headers map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponse) GoString() string {
	return s.String()
}

func (s *GetJobResponse) SetHeaders(v map[string]*string) *GetJobResponse {
	s.Headers = v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type BatchGetJobsStatisticsRequest struct {
	// 工作空间id列表
	WorkspaceIds *string `json:"WorkspaceIds,omitempty" xml:"WorkspaceIds,omitempty"`
}

func (s BatchGetJobsStatisticsRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchGetJobsStatisticsRequest) GoString() string {
	return s.String()
}

func (s *BatchGetJobsStatisticsRequest) SetWorkspaceIds(v string) *BatchGetJobsStatisticsRequest {
	s.WorkspaceIds = &v
	return s
}

type BatchGetJobsStatisticsResponseBody struct {
	// 每一个工作空间id下面的Job按照状态的分类统计信息
	Statistics map[string]interface{} `json:"Statistics,omitempty" xml:"Statistics,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BatchGetJobsStatisticsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchGetJobsStatisticsResponseBody) GoString() string {
	return s.String()
}

func (s *BatchGetJobsStatisticsResponseBody) SetStatistics(v map[string]interface{}) *BatchGetJobsStatisticsResponseBody {
	s.Statistics = v
	return s
}

func (s *BatchGetJobsStatisticsResponseBody) SetRequestId(v string) *BatchGetJobsStatisticsResponseBody {
	s.RequestId = &v
	return s
}

type BatchGetJobsStatisticsResponse struct {
	Headers map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *BatchGetJobsStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s BatchGetJobsStatisticsResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchGetJobsStatisticsResponse) GoString() string {
	return s.String()
}

func (s *BatchGetJobsStatisticsResponse) SetHeaders(v map[string]*string) *BatchGetJobsStatisticsResponse {
	s.Headers = v
	return s
}

func (s *BatchGetJobsStatisticsResponse) SetBody(v *BatchGetJobsStatisticsResponseBody) *BatchGetJobsStatisticsResponse {
	s.Body = v
	return s
}

type CreateDataSourceRequest struct {
	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// 数据源显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 数据源描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 阿里云NAS文件系统Id
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// 阿里云OSS文件系统路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 阿里云OSS文件系统服务端点
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// 阿里云OSS文件系统配置选项
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// 本地挂载目录
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s CreateDataSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDataSourceRequest) SetDataSourceType(v string) *CreateDataSourceRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateDataSourceRequest) SetDisplayName(v string) *CreateDataSourceRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateDataSourceRequest) SetDescription(v string) *CreateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateDataSourceRequest) SetFileSystemId(v string) *CreateDataSourceRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDataSourceRequest) SetPath(v string) *CreateDataSourceRequest {
	s.Path = &v
	return s
}

func (s *CreateDataSourceRequest) SetEndpoint(v string) *CreateDataSourceRequest {
	s.Endpoint = &v
	return s
}

func (s *CreateDataSourceRequest) SetOptions(v string) *CreateDataSourceRequest {
	s.Options = &v
	return s
}

func (s *CreateDataSourceRequest) SetMountPath(v string) *CreateDataSourceRequest {
	s.MountPath = &v
	return s
}

type CreateDataSourceResponseBody struct {
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponseBody) SetDataSourceId(v string) *CreateDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *CreateDataSourceResponseBody) SetRequestId(v string) *CreateDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type CreateDataSourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDataSourceResponse) GoString() string {
	return s.String()
}

func (s *CreateDataSourceResponse) SetHeaders(v map[string]*string) *CreateDataSourceResponse {
	s.Headers = v
	return s
}

func (s *CreateDataSourceResponse) SetBody(v *CreateDataSourceResponseBody) *CreateDataSourceResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	// 镜像类型
	ImageProviderType *string `json:"ImageProviderType,omitempty" xml:"ImageProviderType,omitempty"`
	// 加速器类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 镜像包含的框架类型
	Framework *string `json:"Framework,omitempty" xml:"Framework,omitempty"`
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetImageProviderType(v string) *ListImagesRequest {
	s.ImageProviderType = &v
	return s
}

func (s *ListImagesRequest) SetAcceleratorType(v string) *ListImagesRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListImagesRequest) SetFramework(v string) *ListImagesRequest {
	s.Framework = &v
	return s
}

func (s *ListImagesRequest) SetSortBy(v string) *ListImagesRequest {
	s.SortBy = &v
	return s
}

func (s *ListImagesRequest) SetOrder(v string) *ListImagesRequest {
	s.Order = &v
	return s
}

type ListImagesResponseBody struct {
	// 镜像详情列表
	Images []*ImageItem `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetImages(v []*ImageItem) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetTotalCount(v int64) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListImagesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type GetSecurityGroupResponseBody struct {
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 所属vpc的id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// 安全组id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 安全组名称
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
}

func (s GetSecurityGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityGroupResponseBody) GoString() string {
	return s.String()
}

func (s *GetSecurityGroupResponseBody) SetRequestId(v string) *GetSecurityGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSecurityGroupResponseBody) SetVpcId(v string) *GetSecurityGroupResponseBody {
	s.VpcId = &v
	return s
}

func (s *GetSecurityGroupResponseBody) SetSecurityGroupId(v string) *GetSecurityGroupResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *GetSecurityGroupResponseBody) SetSecurityGroupName(v string) *GetSecurityGroupResponseBody {
	s.SecurityGroupName = &v
	return s
}

type GetSecurityGroupResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSecurityGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSecurityGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s *GetSecurityGroupResponse) SetHeaders(v map[string]*string) *GetSecurityGroupResponse {
	s.Headers = v
	return s
}

func (s *GetSecurityGroupResponse) SetBody(v *GetSecurityGroupResponseBody) *GetSecurityGroupResponse {
	s.Body = v
	return s
}

type GetUserAuthorizationResponseBody struct {
	// 是否通过鉴权
	IsPassed *bool `json:"IsPassed,omitempty" xml:"IsPassed,omitempty"`
	// 请求ID
	RequestId *int32 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUserAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUserAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserAuthorizationResponseBody) SetIsPassed(v bool) *GetUserAuthorizationResponseBody {
	s.IsPassed = &v
	return s
}

func (s *GetUserAuthorizationResponseBody) SetRequestId(v int32) *GetUserAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

type GetUserAuthorizationResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetUserAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetUserAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *GetUserAuthorizationResponse) SetHeaders(v map[string]*string) *GetUserAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *GetUserAuthorizationResponse) SetBody(v *GetUserAuthorizationResponseBody) *GetUserAuthorizationResponse {
	s.Body = v
	return s
}

type DeleteTensorboardRequest struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s DeleteTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteTensorboardRequest) GoString() string {
	return s.String()
}

func (s *DeleteTensorboardRequest) SetWorkspaceId(v string) *DeleteTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type DeleteTensorboardResponseBody struct {
	// Tensorboad Id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTensorboardResponseBody) SetTensorboardId(v string) *DeleteTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

func (s *DeleteTensorboardResponseBody) SetRequestId(v string) *DeleteTensorboardResponseBody {
	s.RequestId = &v
	return s
}

type DeleteTensorboardResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteTensorboardResponse) GoString() string {
	return s.String()
}

func (s *DeleteTensorboardResponse) SetHeaders(v map[string]*string) *DeleteTensorboardResponse {
	s.Headers = v
	return s
}

func (s *DeleteTensorboardResponse) SetBody(v *DeleteTensorboardResponseBody) *DeleteTensorboardResponse {
	s.Body = v
	return s
}

type StopJobResponseBody struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopJobResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobResponseBody) SetJobId(v string) *StopJobResponseBody {
	s.JobId = &v
	return s
}

func (s *StopJobResponseBody) SetRequestId(v string) *StopJobResponseBody {
	s.RequestId = &v
	return s
}

type StopJobResponse struct {
	Headers map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *StopJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopJobResponse) String() string {
	return tea.Prettify(s)
}

func (s StopJobResponse) GoString() string {
	return s.String()
}

func (s *StopJobResponse) SetHeaders(v map[string]*string) *StopJobResponse {
	s.Headers = v
	return s
}

func (s *StopJobResponse) SetBody(v *StopJobResponseBody) *StopJobResponse {
	s.Body = v
	return s
}

type GetJobEventsRequest struct {
	// 获取事件的最大数目，默认值：2000
	MaxEventsNum *int32 `json:"MaxEventsNum,omitempty" xml:"MaxEventsNum,omitempty"`
	// 查询事件的时间区间的起始时间，默认值是7天前。
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 查询事件的时间区间的截止时间，默认值是当前。
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
}

func (s GetJobEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobEventsRequest) GoString() string {
	return s.String()
}

func (s *GetJobEventsRequest) SetMaxEventsNum(v int32) *GetJobEventsRequest {
	s.MaxEventsNum = &v
	return s
}

func (s *GetJobEventsRequest) SetStartTime(v string) *GetJobEventsRequest {
	s.StartTime = &v
	return s
}

func (s *GetJobEventsRequest) SetEndTime(v string) *GetJobEventsRequest {
	s.EndTime = &v
	return s
}

type GetJobEventsResponseBody struct {
	// 作业ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 事件
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *int32 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobEventsResponseBody) SetJobId(v string) *GetJobEventsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobEventsResponseBody) SetEvents(v []*string) *GetJobEventsResponseBody {
	s.Events = v
	return s
}

func (s *GetJobEventsResponseBody) SetRequestId(v int32) *GetJobEventsResponseBody {
	s.RequestId = &v
	return s
}

type GetJobEventsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetJobEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetJobEventsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetJobEventsResponse) GoString() string {
	return s.String()
}

func (s *GetJobEventsResponse) SetHeaders(v map[string]*string) *GetJobEventsResponse {
	s.Headers = v
	return s
}

func (s *GetJobEventsResponse) SetBody(v *GetJobEventsResponseBody) *GetJobEventsResponse {
	s.Body = v
	return s
}

type DeleteJobResponseBody struct {
	// 作业Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobResponseBody) SetJobId(v string) *DeleteJobResponseBody {
	s.JobId = &v
	return s
}

func (s *DeleteJobResponseBody) SetRequestId(v string) *DeleteJobResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobResponse) SetHeaders(v map[string]*string) *DeleteJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobResponse) SetBody(v *DeleteJobResponseBody) *DeleteJobResponse {
	s.Body = v
	return s
}

type DeleteCodeSourceResponseBody struct {
	// 被删除的代码源配置ID
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCodeSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCodeSourceResponseBody) SetCodeSourceId(v string) *DeleteCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *DeleteCodeSourceResponseBody) SetRequestId(v string) *DeleteCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteCodeSourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteCodeSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteCodeSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteCodeSourceResponse) SetHeaders(v map[string]*string) *DeleteCodeSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteCodeSourceResponse) SetBody(v *DeleteCodeSourceResponseBody) *DeleteCodeSourceResponse {
	s.Body = v
	return s
}

type DeleteDataSourceResponseBody struct {
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponseBody) SetDataSourceId(v string) *DeleteDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *DeleteDataSourceResponseBody) SetRequestId(v string) *DeleteDataSourceResponseBody {
	s.RequestId = &v
	return s
}

type DeleteDataSourceResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteDataSourceResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteDataSourceResponse) SetHeaders(v map[string]*string) *DeleteDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteDataSourceResponse) SetBody(v *DeleteDataSourceResponseBody) *DeleteDataSourceResponse {
	s.Body = v
	return s
}

type ListWorkspacesRequest struct {
	// 是否返回详情(Quotas, Members)
	NeedDetail *bool `json:"NeedDetail,omitempty" xml:"NeedDetail,omitempty"`
	// 查询第几页数据,最小值为1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 设置查询的分页大小,最小值为1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 返回结果的排序字段名，枚举值
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 排列顺序: desc 或者 asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListWorkspacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesRequest) GoString() string {
	return s.String()
}

func (s *ListWorkspacesRequest) SetNeedDetail(v bool) *ListWorkspacesRequest {
	s.NeedDetail = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageNumber(v int32) *ListWorkspacesRequest {
	s.PageNumber = &v
	return s
}

func (s *ListWorkspacesRequest) SetPageSize(v int32) *ListWorkspacesRequest {
	s.PageSize = &v
	return s
}

func (s *ListWorkspacesRequest) SetSortBy(v string) *ListWorkspacesRequest {
	s.SortBy = &v
	return s
}

func (s *ListWorkspacesRequest) SetOrder(v string) *ListWorkspacesRequest {
	s.Order = &v
	return s
}

type ListWorkspacesResponseBody struct {
	// 工作空间列表
	Workspaces []*Workspace `json:"Workspaces,omitempty" xml:"Workspaces,omitempty" type:"Repeated"`
	// 结果数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListWorkspacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponseBody) SetWorkspaces(v []*Workspace) *ListWorkspacesResponseBody {
	s.Workspaces = v
	return s
}

func (s *ListWorkspacesResponseBody) SetTotalCount(v int64) *ListWorkspacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListWorkspacesResponseBody) SetRequestId(v string) *ListWorkspacesResponseBody {
	s.RequestId = &v
	return s
}

type ListWorkspacesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListWorkspacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListWorkspacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListWorkspacesResponse) GoString() string {
	return s.String()
}

func (s *ListWorkspacesResponse) SetHeaders(v map[string]*string) *ListWorkspacesResponse {
	s.Headers = v
	return s
}

func (s *ListWorkspacesResponse) SetBody(v *ListWorkspacesResponseBody) *ListWorkspacesResponse {
	s.Body = v
	return s
}

type GetTokenRequest struct {
	// TargetType
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// TargetId
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// ExpireTime
	ExpireTime *int64 `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
}

func (s GetTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) SetTargetType(v string) *GetTokenRequest {
	s.TargetType = &v
	return s
}

func (s *GetTokenRequest) SetTargetId(v string) *GetTokenRequest {
	s.TargetId = &v
	return s
}

func (s *GetTokenRequest) SetExpireTime(v int64) *GetTokenRequest {
	s.ExpireTime = &v
	return s
}

type GetTokenResponseBody struct {
	// RequestId
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Token
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) SetToken(v string) *GetTokenResponseBody {
	s.Token = &v
	return s
}

type GetTokenResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTokenResponse) GoString() string {
	return s.String()
}

func (s *GetTokenResponse) SetHeaders(v map[string]*string) *GetTokenResponse {
	s.Headers = v
	return s
}

func (s *GetTokenResponse) SetBody(v *GetTokenResponseBody) *GetTokenResponse {
	s.Body = v
	return s
}

type GetAuthorizationResponseBody struct {
	// 是否授权
	Authorized *bool `json:"Authorized,omitempty" xml:"Authorized,omitempty"`
	// 授权失败码
	AuthorizationFailedCode *string `json:"AuthorizationFailedCode,omitempty" xml:"AuthorizationFailedCode,omitempty"`
	// 授权失败消息
	AuthorizationFailedMessage *string `json:"AuthorizationFailedMessage,omitempty" xml:"AuthorizationFailedMessage,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAuthorizationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *GetAuthorizationResponseBody) SetAuthorized(v bool) *GetAuthorizationResponseBody {
	s.Authorized = &v
	return s
}

func (s *GetAuthorizationResponseBody) SetAuthorizationFailedCode(v string) *GetAuthorizationResponseBody {
	s.AuthorizationFailedCode = &v
	return s
}

func (s *GetAuthorizationResponseBody) SetAuthorizationFailedMessage(v string) *GetAuthorizationResponseBody {
	s.AuthorizationFailedMessage = &v
	return s
}

func (s *GetAuthorizationResponseBody) SetRequestId(v string) *GetAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

type GetAuthorizationResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAuthorizationResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *GetAuthorizationResponse) SetHeaders(v map[string]*string) *GetAuthorizationResponse {
	s.Headers = v
	return s
}

func (s *GetAuthorizationResponse) SetBody(v *GetAuthorizationResponseBody) *GetAuthorizationResponse {
	s.Body = v
	return s
}

type ListQuotasRequest struct {
	// 当前页
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的作业数
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
}

func (s ListQuotasRequest) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasRequest) GoString() string {
	return s.String()
}

func (s *ListQuotasRequest) SetPageNumber(v int32) *ListQuotasRequest {
	s.PageNumber = &v
	return s
}

func (s *ListQuotasRequest) SetPageSize(v int32) *ListQuotasRequest {
	s.PageSize = &v
	return s
}

func (s *ListQuotasRequest) SetSortBy(v string) *ListQuotasRequest {
	s.SortBy = &v
	return s
}

func (s *ListQuotasRequest) SetOrder(v string) *ListQuotasRequest {
	s.Order = &v
	return s
}

type ListQuotasResponseBody struct {
	// 资源配额列表
	Quotas []*Quota `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListQuotasResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBody) SetQuotas(v []*Quota) *ListQuotasResponseBody {
	s.Quotas = v
	return s
}

func (s *ListQuotasResponseBody) SetRequestId(v string) *ListQuotasResponseBody {
	s.RequestId = &v
	return s
}

type ListQuotasResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListQuotasResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListQuotasResponse) String() string {
	return tea.Prettify(s)
}

func (s ListQuotasResponse) GoString() string {
	return s.String()
}

func (s *ListQuotasResponse) SetHeaders(v map[string]*string) *ListQuotasResponse {
	s.Headers = v
	return s
}

func (s *ListQuotasResponse) SetBody(v *ListQuotasResponseBody) *ListQuotasResponse {
	s.Body = v
	return s
}

type CreateTensorboardRequest struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// DataSource Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Tensorboard名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// Summary 目录
	SummaryPath *string           `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	DataSources []*DataSourceItem `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
}

func (s CreateTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTensorboardRequest) GoString() string {
	return s.String()
}

func (s *CreateTensorboardRequest) SetWorkspaceId(v string) *CreateTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateTensorboardRequest) SetJobId(v string) *CreateTensorboardRequest {
	s.JobId = &v
	return s
}

func (s *CreateTensorboardRequest) SetDataSourceId(v string) *CreateTensorboardRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateTensorboardRequest) SetDisplayName(v string) *CreateTensorboardRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateTensorboardRequest) SetSummaryPath(v string) *CreateTensorboardRequest {
	s.SummaryPath = &v
	return s
}

func (s *CreateTensorboardRequest) SetDataSources(v []*DataSourceItem) *CreateTensorboardRequest {
	s.DataSources = v
	return s
}

type CreateTensorboardResponseBody struct {
	// 任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// DataSourceId
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// Tensorboard id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTensorboardResponseBody) SetJobId(v string) *CreateTensorboardResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetDataSourceId(v string) *CreateTensorboardResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetTensorboardId(v string) *CreateTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetRequestId(v string) *CreateTensorboardResponseBody {
	s.RequestId = &v
	return s
}

type CreateTensorboardResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTensorboardResponse) GoString() string {
	return s.String()
}

func (s *CreateTensorboardResponse) SetHeaders(v map[string]*string) *CreateTensorboardResponse {
	s.Headers = v
	return s
}

func (s *CreateTensorboardResponse) SetBody(v *CreateTensorboardResponseBody) *CreateTensorboardResponse {
	s.Body = v
	return s
}

type UpdateTensorboardRequest struct {
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// MaxRunningTimeMinutes
	MaxRunningTimeMinutes *string `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
}

func (s UpdateTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTensorboardRequest) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardRequest) SetWorkspaceId(v string) *UpdateTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

func (s *UpdateTensorboardRequest) SetMaxRunningTimeMinutes(v string) *UpdateTensorboardRequest {
	s.MaxRunningTimeMinutes = &v
	return s
}

type UpdateTensorboardResponseBody struct {
	// Tensorboad Id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardResponseBody) SetTensorboardId(v string) *UpdateTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

func (s *UpdateTensorboardResponseBody) SetRequestId(v string) *UpdateTensorboardResponseBody {
	s.RequestId = &v
	return s
}

type UpdateTensorboardResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateTensorboardResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTensorboardResponse) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardResponse) SetHeaders(v map[string]*string) *UpdateTensorboardResponse {
	s.Headers = v
	return s
}

func (s *UpdateTensorboardResponse) SetBody(v *UpdateTensorboardResponseBody) *UpdateTensorboardResponse {
	s.Body = v
	return s
}

type UpdateQuotaRequest struct {
	// 资源配额名称
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// 资源配额类型
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// 资源配额参数
	QuotaDetail *QuotaDetail `json:"QuotaDetail,omitempty" xml:"QuotaDetail,omitempty"`
}

func (s UpdateQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaRequest) GoString() string {
	return s.String()
}

func (s *UpdateQuotaRequest) SetQuotaName(v string) *UpdateQuotaRequest {
	s.QuotaName = &v
	return s
}

func (s *UpdateQuotaRequest) SetQuotaType(v string) *UpdateQuotaRequest {
	s.QuotaType = &v
	return s
}

func (s *UpdateQuotaRequest) SetQuotaDetail(v *QuotaDetail) *UpdateQuotaRequest {
	s.QuotaDetail = v
	return s
}

type UpdateQuotaResponseBody struct {
	// 资源配额id
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponseBody) SetQuotaId(v string) *UpdateQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *UpdateQuotaResponseBody) SetRequestId(v string) *UpdateQuotaResponseBody {
	s.RequestId = &v
	return s
}

type UpdateQuotaResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *UpdateQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateQuotaResponse) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponse) SetHeaders(v map[string]*string) *UpdateQuotaResponse {
	s.Headers = v
	return s
}

func (s *UpdateQuotaResponse) SetBody(v *UpdateQuotaResponseBody) *UpdateQuotaResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("pai-dlc"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSources(request *ListDataSourcesRequest) (_result *ListDataSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDataSourcesResponse{}
	_body, _err := client.ListDataSourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListDataSourcesWithOptions(request *ListDataSourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDataSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		query["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListDataSources"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/datasources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobsStatistics(request *GetJobsStatisticsRequest) (_result *GetJobsStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobsStatisticsResponse{}
	_body, _err := client.GetJobsStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobsStatisticsWithOptions(request *GetJobsStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobsStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetJobsStatisticsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetJobsStatistics"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/statistics/jobs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDataSource(DataSourceId *string) (_result *GetDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDataSourceResponse{}
	_body, _err := client.GetDataSourceWithOptions(DataSourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDataSourceWithOptions(DataSourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDataSourceResponse, _err error) {
	DataSourceId = openapiutil.GetEncodeParam(DataSourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetDataSourceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetDataSource"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/datasources/"+tea.StringValue(DataSourceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateQuota(request *CreateQuotaRequest) (_result *CreateQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateQuotaResponse{}
	_body, _err := client.CreateQuotaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateQuotaWithOptions(request *CreateQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QuotaName)) {
		body["QuotaName"] = request.QuotaName
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaType)) {
		body["QuotaType"] = request.QuotaType
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.QuotaDetail))) {
		body["QuotaDetail"] = request.QuotaDetail
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateQuotaResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateQuota"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/quotas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQuota(QuotaId *string) (_result *GetQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetQuotaResponse{}
	_body, _err := client.GetQuotaWithOptions(QuotaId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQuotaWithOptions(QuotaId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetQuotaResponse, _err error) {
	QuotaId = openapiutil.GetEncodeParam(QuotaId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetQuotaResponse{}
	_body, _err := client.DoROARequest(tea.String("GetQuota"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/quotas/"+tea.StringValue(QuotaId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetCodeSource(CodeSourceId *string) (_result *GetCodeSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCodeSourceResponse{}
	_body, _err := client.GetCodeSourceWithOptions(CodeSourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetCodeSourceWithOptions(CodeSourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCodeSourceResponse, _err error) {
	CodeSourceId = openapiutil.GetEncodeParam(CodeSourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetCodeSourceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetCodeSource"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/codesources/"+tea.StringValue(CodeSourceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSwitch(SwitchId *string) (_result *GetSwitchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSwitchResponse{}
	_body, _err := client.GetSwitchWithOptions(SwitchId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSwitchWithOptions(SwitchId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSwitchResponse, _err error) {
	SwitchId = openapiutil.GetEncodeParam(SwitchId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetSwitchResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSwitch"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/switches/"+tea.StringValue(SwitchId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPodEvents(JobId *string, PodId *string, request *GetPodEventsRequest) (_result *GetPodEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPodEventsResponse{}
	_body, _err := client.GetPodEventsWithOptions(JobId, PodId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPodEventsWithOptions(JobId *string, PodId *string, request *GetPodEventsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPodEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	JobId = openapiutil.GetEncodeParam(JobId)
	PodId = openapiutil.GetEncodeParam(PodId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxEventsNum)) {
		query["MaxEventsNum"] = request.MaxEventsNum
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetPodEventsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPodEvents"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/jobs/"+tea.StringValue(JobId)+"/pods/"+tea.StringValue(PodId)+"/events"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSwitches(request *ListSwitchesRequest) (_result *ListSwitchesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSwitchesResponse{}
	_body, _err := client.ListSwitchesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSwitchesWithOptions(request *ListSwitchesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSwitchesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListSwitchesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSwitches"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/switches"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTensorboards(request *ListTensorboardsRequest) (_result *ListTensorboardsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTensorboardsResponse{}
	_body, _err := client.ListTensorboardsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTensorboardsWithOptions(request *ListTensorboardsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTensorboardsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.TensorboardId)) {
		query["TensorboardId"] = request.TensorboardId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListTensorboardsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListTensorboards"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/tensorboards"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListSecurityGroups(request *ListSecurityGroupsRequest) (_result *ListSecurityGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListSecurityGroupsResponse{}
	_body, _err := client.ListSecurityGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListSecurityGroupsWithOptions(request *ListSecurityGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListSecurityGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VpcId)) {
		query["VpcId"] = request.VpcId
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListSecurityGroupsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListSecurityGroups"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/securitygroups"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPodLogs(JobId *string, PodId *string, request *GetPodLogsRequest) (_result *GetPodLogsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetPodLogsResponse{}
	_body, _err := client.GetPodLogsWithOptions(JobId, PodId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPodLogsWithOptions(JobId *string, PodId *string, request *GetPodLogsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetPodLogsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	JobId = openapiutil.GetEncodeParam(JobId)
	PodId = openapiutil.GetEncodeParam(PodId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxLines)) {
		query["MaxLines"] = request.MaxLines
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.DownloadToFile)) {
		query["DownloadToFile"] = request.DownloadToFile
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetPodLogsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetPodLogs"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/jobs/"+tea.StringValue(JobId)+"/pods/"+tea.StringValue(PodId)+"/logs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListVpcs(request *ListVpcsRequest) (_result *ListVpcsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListVpcsResponse{}
	_body, _err := client.ListVpcsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListVpcsWithOptions(request *ListVpcsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListVpcsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListVpcsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListVpcs"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/vpcs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopTensorboard(TensorboardId *string, request *StopTensorboardRequest) (_result *StopTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopTensorboardResponse{}
	_body, _err := client.StopTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopTensorboardWithOptions(TensorboardId *string, request *StopTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopTensorboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	TensorboardId = openapiutil.GetEncodeParam(TensorboardId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &StopTensorboardResponse{}
	_body, _err := client.DoROARequest(tea.String("StopTensorboard"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/tensorboards/"+tea.StringValue(TensorboardId)+"/stop"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateJob(request *CreateJobRequest) (_result *CreateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateJobResponse{}
	_body, _err := client.CreateJobWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateJobWithOptions(request *CreateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		body["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.JobSpecs)) {
		body["JobSpecs"] = request.JobSpecs
	}

	if !tea.BoolValue(util.IsUnset(request.UserCommand)) {
		body["UserCommand"] = request.UserCommand
	}

	if !tea.BoolValue(util.IsUnset(request.DataSources)) {
		body["DataSources"] = request.DataSources
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CodeSource))) {
		body["CodeSource"] = request.CodeSource
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.UserVpc))) {
		body["UserVpc"] = request.UserVpc
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartyLibs)) {
		body["ThirdpartyLibs"] = request.ThirdpartyLibs
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartyLibDir)) {
		body["ThirdpartyLibDir"] = request.ThirdpartyLibDir
	}

	if !tea.BoolValue(util.IsUnset(request.Envs)) {
		body["Envs"] = request.Envs
	}

	if !tea.BoolValue(util.IsUnset(request.JobMaxRunningTimeMinutes)) {
		body["JobMaxRunningTimeMinutes"] = request.JobMaxRunningTimeMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Settings))) {
		body["Settings"] = request.Settings
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ElasticSpec))) {
		body["ElasticSpec"] = request.ElasticSpec
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateJob"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/jobs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListCodeSources(request *ListCodeSourcesRequest) (_result *ListCodeSourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListCodeSourcesResponse{}
	_body, _err := client.ListCodeSourcesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListCodeSourcesWithOptions(request *ListCodeSourcesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListCodeSourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListCodeSourcesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListCodeSources"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/codesources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JobDispatchSubmit(request *JobDispatchSubmitRequest) (_result *JobDispatchSubmitResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &JobDispatchSubmitResponse{}
	_body, _err := client.JobDispatchSubmitWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JobDispatchSubmitWithOptions(request *JobDispatchSubmitRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *JobDispatchSubmitResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgoName)) {
		body["AlgoName"] = request.AlgoName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		body["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.Properties)) {
		body["Properties"] = request.Properties
	}

	if !tea.BoolValue(util.IsUnset(request.Settings)) {
		body["Settings"] = request.Settings
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &JobDispatchSubmitResponse{}
	_body, _err := client.DoROARequest(tea.String("JobDispatchSubmit"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/jobdispatch"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListEcsSpecs(request *ListEcsSpecsRequest) (_result *ListEcsSpecsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.ListEcsSpecsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListEcsSpecsWithOptions(request *ListEcsSpecsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListEcsSpecsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListEcsSpecs"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/ecsspecs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteQuota(QuotaId *string) (_result *DeleteQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteQuotaResponse{}
	_body, _err := client.DeleteQuotaWithOptions(QuotaId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteQuotaWithOptions(QuotaId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteQuotaResponse, _err error) {
	QuotaId = openapiutil.GetEncodeParam(QuotaId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteQuotaResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteQuota"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/quotas/"+tea.StringValue(QuotaId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StartTensorboard(TensorboardId *string, request *StartTensorboardRequest) (_result *StartTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartTensorboardResponse{}
	_body, _err := client.StartTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StartTensorboardWithOptions(TensorboardId *string, request *StartTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartTensorboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	TensorboardId = openapiutil.GetEncodeParam(TensorboardId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &StartTensorboardResponse{}
	_body, _err := client.DoROARequest(tea.String("StartTensorboard"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/tensorboards/"+tea.StringValue(TensorboardId)+"/start"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetTensorboard(TensorboardId *string, request *GetTensorboardRequest) (_result *GetTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTensorboardResponse{}
	_body, _err := client.GetTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTensorboardWithOptions(TensorboardId *string, request *GetTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTensorboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	TensorboardId = openapiutil.GetEncodeParam(TensorboardId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.JodId)) {
		query["JodId"] = request.JodId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetTensorboardResponse{}
	_body, _err := client.DoROARequest(tea.String("GetTensorboard"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/tensorboards/"+tea.StringValue(TensorboardId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetWorkspace(WorkspaceId *string) (_result *GetWorkspaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetWorkspaceResponse{}
	_body, _err := client.GetWorkspaceWithOptions(WorkspaceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetWorkspaceWithOptions(WorkspaceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetWorkspaceResponse, _err error) {
	WorkspaceId = openapiutil.GetEncodeParam(WorkspaceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetWorkspaceResponse{}
	_body, _err := client.DoROARequest(tea.String("GetWorkspace"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/workspaces/"+tea.StringValue(WorkspaceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) JobDispatchQuery(request *JobDispatchQueryRequest) (_result *JobDispatchQueryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &JobDispatchQueryResponse{}
	_body, _err := client.JobDispatchQueryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) JobDispatchQueryWithOptions(tmpReq *JobDispatchQueryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *JobDispatchQueryResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &JobDispatchQueryShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Properties)) {
		request.PropertiesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Properties, tea.String("Properties"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Settings)) {
		request.SettingsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Settings, tea.String("Settings"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AlgoName)) {
		query["AlgoName"] = request.AlgoName
	}

	if !tea.BoolValue(util.IsUnset(request.ProjectName)) {
		query["ProjectName"] = request.ProjectName
	}

	if !tea.BoolValue(util.IsUnset(request.PropertiesShrink)) {
		query["Properties"] = request.PropertiesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.SettingsShrink)) {
		query["Settings"] = request.SettingsShrink
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &JobDispatchQueryResponse{}
	_body, _err := client.DoROARequest(tea.String("JobDispatchQuery"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/jobdispatch"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListJobs(request *ListJobsRequest) (_result *ListJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListJobsResponse{}
	_body, _err := client.ListJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListJobsWithOptions(tmpReq *ListJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListJobsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListJobsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Tags)) {
		request.TagsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Tags, tea.String("Tags"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.ShowOwn)) {
		query["ShowOwn"] = request.ShowOwn
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.BusinessUserId)) {
		query["BusinessUserId"] = request.BusinessUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.DoROARequest(tea.String("ListJobs"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/jobs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVpc(VpcId *string) (_result *GetVpcResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVpcResponse{}
	_body, _err := client.GetVpcWithOptions(VpcId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVpcWithOptions(VpcId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetVpcResponse, _err error) {
	VpcId = openapiutil.GetEncodeParam(VpcId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetVpcResponse{}
	_body, _err := client.DoROARequest(tea.String("GetVpc"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/vpcs/"+tea.StringValue(VpcId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateCodeSource(request *CreateCodeSourceRequest) (_result *CreateCodeSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCodeSourceResponse{}
	_body, _err := client.CreateCodeSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateCodeSourceWithOptions(request *CreateCodeSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCodeSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepo)) {
		body["CodeRepo"] = request.CodeRepo
	}

	if !tea.BoolValue(util.IsUnset(request.CodeBranch)) {
		body["CodeBranch"] = request.CodeBranch
	}

	if !tea.BoolValue(util.IsUnset(request.MountPath)) {
		body["MountPath"] = request.MountPath
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoUserName)) {
		body["CodeRepoUserName"] = request.CodeRepoUserName
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoAccessToken)) {
		body["CodeRepoAccessToken"] = request.CodeRepoAccessToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateCodeSourceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateCodeSource"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/codesources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobMetrics(JobId *string, request *GetJobMetricsRequest) (_result *GetJobMetricsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobMetricsResponse{}
	_body, _err := client.GetJobMetricsWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobMetricsWithOptions(JobId *string, request *GetJobMetricsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobMetricsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	JobId = openapiutil.GetEncodeParam(JobId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStep)) {
		query["TimeStep"] = request.TimeStep
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetJobMetricsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetJobMetrics"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/jobs/"+tea.StringValue(JobId)+"/metrics"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJob(JobId *string) (_result *GetJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobResponse{}
	_body, _err := client.GetJobWithOptions(JobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobWithOptions(JobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobResponse, _err error) {
	JobId = openapiutil.GetEncodeParam(JobId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetJobResponse{}
	_body, _err := client.DoROARequest(tea.String("GetJob"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/jobs/"+tea.StringValue(JobId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) BatchGetJobsStatistics(request *BatchGetJobsStatisticsRequest) (_result *BatchGetJobsStatisticsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchGetJobsStatisticsResponse{}
	_body, _err := client.BatchGetJobsStatisticsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) BatchGetJobsStatisticsWithOptions(request *BatchGetJobsStatisticsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchGetJobsStatisticsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceIds)) {
		query["WorkspaceIds"] = request.WorkspaceIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &BatchGetJobsStatisticsResponse{}
	_body, _err := client.DoROARequest(tea.String("BatchGetJobsStatistics"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/batch/statistics/jobs"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateDataSource(request *CreateDataSourceRequest) (_result *CreateDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CreateDataSourceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateDataSourceWithOptions(request *CreateDataSourceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDataSourceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		body["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		body["Endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.MountPath)) {
		body["MountPath"] = request.MountPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateDataSource"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/datasources"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListImagesWithOptions(request *ListImagesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ImageProviderType)) {
		query["ImageProviderType"] = request.ImageProviderType
	}

	if !tea.BoolValue(util.IsUnset(request.AcceleratorType)) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	if !tea.BoolValue(util.IsUnset(request.Framework)) {
		query["Framework"] = request.Framework
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListImages"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/images"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSecurityGroup(SecurityGroupId *string) (_result *GetSecurityGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetSecurityGroupResponse{}
	_body, _err := client.GetSecurityGroupWithOptions(SecurityGroupId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSecurityGroupWithOptions(SecurityGroupId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetSecurityGroupResponse, _err error) {
	SecurityGroupId = openapiutil.GetEncodeParam(SecurityGroupId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetSecurityGroupResponse{}
	_body, _err := client.DoROARequest(tea.String("GetSecurityGroup"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/securitygroups/"+tea.StringValue(SecurityGroupId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetUserAuthorization(UserId *string) (_result *GetUserAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserAuthorizationResponse{}
	_body, _err := client.GetUserAuthorizationWithOptions(UserId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetUserAuthorizationWithOptions(UserId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserAuthorizationResponse, _err error) {
	UserId = openapiutil.GetEncodeParam(UserId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetUserAuthorizationResponse{}
	_body, _err := client.DoROARequest(tea.String("GetUserAuthorization"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/users/"+tea.StringValue(UserId)+"/authorization"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteTensorboard(TensorboardId *string, request *DeleteTensorboardRequest) (_result *DeleteTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteTensorboardResponse{}
	_body, _err := client.DeleteTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteTensorboardWithOptions(TensorboardId *string, request *DeleteTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteTensorboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	TensorboardId = openapiutil.GetEncodeParam(TensorboardId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &DeleteTensorboardResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteTensorboard"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/tensorboards/"+tea.StringValue(TensorboardId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopJob(JobId *string) (_result *StopJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopJobResponse{}
	_body, _err := client.StopJobWithOptions(JobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopJobWithOptions(JobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopJobResponse, _err error) {
	JobId = openapiutil.GetEncodeParam(JobId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &StopJobResponse{}
	_body, _err := client.DoROARequest(tea.String("StopJob"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/jobs/"+tea.StringValue(JobId)+"/stop"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetJobEvents(JobId *string, request *GetJobEventsRequest) (_result *GetJobEventsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetJobEventsResponse{}
	_body, _err := client.GetJobEventsWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetJobEventsWithOptions(JobId *string, request *GetJobEventsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetJobEventsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	JobId = openapiutil.GetEncodeParam(JobId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxEventsNum)) {
		query["MaxEventsNum"] = request.MaxEventsNum
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetJobEventsResponse{}
	_body, _err := client.DoROARequest(tea.String("GetJobEvents"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/jobs/"+tea.StringValue(JobId)+"/events"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJob(JobId *string) (_result *DeleteJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteJobResponse{}
	_body, _err := client.DeleteJobWithOptions(JobId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobWithOptions(JobId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteJobResponse, _err error) {
	JobId = openapiutil.GetEncodeParam(JobId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteJob"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/jobs/"+tea.StringValue(JobId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteCodeSource(CodeSourceId *string) (_result *DeleteCodeSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteCodeSourceResponse{}
	_body, _err := client.DeleteCodeSourceWithOptions(CodeSourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteCodeSourceWithOptions(CodeSourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteCodeSourceResponse, _err error) {
	CodeSourceId = openapiutil.GetEncodeParam(CodeSourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteCodeSourceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteCodeSource"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/codesources/"+tea.StringValue(CodeSourceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteDataSource(DataSourceId *string) (_result *DeleteDataSourceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DeleteDataSourceWithOptions(DataSourceId, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteDataSourceWithOptions(DataSourceId *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDataSourceResponse, _err error) {
	DataSourceId = openapiutil.GetEncodeParam(DataSourceId)
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.DoROARequest(tea.String("DeleteDataSource"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("DELETE"), tea.String("AK"), tea.String("/api/v1/datasources/"+tea.StringValue(DataSourceId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListWorkspaces(request *ListWorkspacesRequest) (_result *ListWorkspacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListWorkspacesResponse{}
	_body, _err := client.ListWorkspacesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListWorkspacesWithOptions(request *ListWorkspacesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListWorkspacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NeedDetail)) {
		query["NeedDetail"] = request.NeedDetail
	}

	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListWorkspacesResponse{}
	_body, _err := client.DoROARequest(tea.String("ListWorkspaces"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/workspaces"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetToken(request *GetTokenRequest) (_result *GetTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTokenResponse{}
	_body, _err := client.GetTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetTokenWithOptions(request *GetTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TargetType)) {
		query["TargetType"] = request.TargetType
	}

	if !tea.BoolValue(util.IsUnset(request.TargetId)) {
		query["TargetId"] = request.TargetId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		query["ExpireTime"] = request.ExpireTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &GetTokenResponse{}
	_body, _err := client.DoROARequest(tea.String("GetToken"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/tokens"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAuthorization() (_result *GetAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAuthorizationResponse{}
	_body, _err := client.GetAuthorizationWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAuthorizationWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAuthorizationResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &GetAuthorizationResponse{}
	_body, _err := client.DoROARequest(tea.String("GetAuthorization"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/authorization"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListQuotas(request *ListQuotasRequest) (_result *ListQuotasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListQuotasResponse{}
	_body, _err := client.ListQuotasWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListQuotasWithOptions(request *ListQuotasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListQuotasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.PageNumber)) {
		query["PageNumber"] = request.PageNumber
	}

	if !tea.BoolValue(util.IsUnset(request.PageSize)) {
		query["PageSize"] = request.PageSize
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &ListQuotasResponse{}
	_body, _err := client.DoROARequest(tea.String("ListQuotas"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("GET"), tea.String("AK"), tea.String("/api/v1/quotas"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTensorboard(request *CreateTensorboardRequest) (_result *CreateTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTensorboardResponse{}
	_body, _err := client.CreateTensorboardWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTensorboardWithOptions(request *CreateTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTensorboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.SummaryPath)) {
		body["SummaryPath"] = request.SummaryPath
	}

	if !tea.BoolValue(util.IsUnset(request.DataSources)) {
		body["DataSources"] = request.DataSources
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &CreateTensorboardResponse{}
	_body, _err := client.DoROARequest(tea.String("CreateTensorboard"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/api/v1/tensorboards"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateTensorboard(TensorboardId *string, request *UpdateTensorboardRequest) (_result *UpdateTensorboardResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTensorboardResponse{}
	_body, _err := client.UpdateTensorboardWithOptions(TensorboardId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateTensorboardWithOptions(TensorboardId *string, request *UpdateTensorboardRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTensorboardResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	TensorboardId = openapiutil.GetEncodeParam(TensorboardId)
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxRunningTimeMinutes)) {
		query["MaxRunningTimeMinutes"] = request.MaxRunningTimeMinutes
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	_result = &UpdateTensorboardResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateTensorboard"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/tensorboards/"+tea.StringValue(TensorboardId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateQuota(QuotaId *string, request *UpdateQuotaRequest) (_result *UpdateQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateQuotaResponse{}
	_body, _err := client.UpdateQuotaWithOptions(QuotaId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateQuotaWithOptions(QuotaId *string, request *UpdateQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	QuotaId = openapiutil.GetEncodeParam(QuotaId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.QuotaName)) {
		body["QuotaName"] = request.QuotaName
	}

	if !tea.BoolValue(util.IsUnset(request.QuotaType)) {
		body["QuotaType"] = request.QuotaType
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.QuotaDetail))) {
		body["QuotaDetail"] = request.QuotaDetail
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	_result = &UpdateQuotaResponse{}
	_body, _err := client.DoROARequest(tea.String("UpdateQuota"), tea.String("2020-12-03"), tea.String("HTTPS"), tea.String("PUT"), tea.String("AK"), tea.String("/api/v1/quotas/"+tea.StringValue(QuotaId)), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
