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

type CodeSourceItem struct {
	// 代码分支
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// 代码Commit ID
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// 代码仓库地址
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// 访问代码仓库所用的AccessToken
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// 访问代码仓库的用户名
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// 代码源ID
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// 代码源详细描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代码源配置的名字
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 代码源配置的用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s CodeSourceItem) String() string {
	return tea.Prettify(s)
}

func (s CodeSourceItem) GoString() string {
	return s.String()
}

func (s *CodeSourceItem) SetCodeBranch(v string) *CodeSourceItem {
	s.CodeBranch = &v
	return s
}

func (s *CodeSourceItem) SetCodeCommit(v string) *CodeSourceItem {
	s.CodeCommit = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepo(v string) *CodeSourceItem {
	s.CodeRepo = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepoAccessToken(v string) *CodeSourceItem {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *CodeSourceItem) SetCodeRepoUserName(v string) *CodeSourceItem {
	s.CodeRepoUserName = &v
	return s
}

func (s *CodeSourceItem) SetCodeSourceId(v string) *CodeSourceItem {
	s.CodeSourceId = &v
	return s
}

func (s *CodeSourceItem) SetDescription(v string) *CodeSourceItem {
	s.Description = &v
	return s
}

func (s *CodeSourceItem) SetDisplayName(v string) *CodeSourceItem {
	s.DisplayName = &v
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

func (s *CodeSourceItem) SetUserId(v string) *CodeSourceItem {
	s.UserId = &v
	return s
}

type ContainerSpec struct {
	// 命令参数
	Args []*string `json:"Args,omitempty" xml:"Args,omitempty" type:"Repeated"`
	// 用户命令
	Command []*string `json:"Command,omitempty" xml:"Command,omitempty" type:"Repeated"`
	// 环境变量
	Env []*EnvVar `json:"Env,omitempty" xml:"Env,omitempty" type:"Repeated"`
	// 容器镜像地址
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// 容器名称
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 容器资源
	Resources *ResourceRequirements `json:"Resources,omitempty" xml:"Resources,omitempty"`
	// 容器内工作目录
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s ContainerSpec) String() string {
	return tea.Prettify(s)
}

func (s ContainerSpec) GoString() string {
	return s.String()
}

func (s *ContainerSpec) SetArgs(v []*string) *ContainerSpec {
	s.Args = v
	return s
}

func (s *ContainerSpec) SetCommand(v []*string) *ContainerSpec {
	s.Command = v
	return s
}

func (s *ContainerSpec) SetEnv(v []*EnvVar) *ContainerSpec {
	s.Env = v
	return s
}

func (s *ContainerSpec) SetImage(v string) *ContainerSpec {
	s.Image = &v
	return s
}

func (s *ContainerSpec) SetName(v string) *ContainerSpec {
	s.Name = &v
	return s
}

func (s *ContainerSpec) SetResources(v *ResourceRequirements) *ContainerSpec {
	s.Resources = v
	return s
}

func (s *ContainerSpec) SetWorkingDir(v string) *ContainerSpec {
	s.WorkingDir = &v
	return s
}

type DataSourceItem struct {
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// 数据源描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 数据源显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 阿里云OSS文件系统服务端点
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// 阿里云NAS文件系统Id
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// 创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间（UTC）
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 本地挂载目录
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 阿里云OSS文件系统配置选项
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// 阿里云OSS文件系统路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 创建人Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DataSourceItem) String() string {
	return tea.Prettify(s)
}

func (s DataSourceItem) GoString() string {
	return s.String()
}

func (s *DataSourceItem) SetDataSourceId(v string) *DataSourceItem {
	s.DataSourceId = &v
	return s
}

func (s *DataSourceItem) SetDataSourceType(v string) *DataSourceItem {
	s.DataSourceType = &v
	return s
}

func (s *DataSourceItem) SetDescription(v string) *DataSourceItem {
	s.Description = &v
	return s
}

func (s *DataSourceItem) SetDisplayName(v string) *DataSourceItem {
	s.DisplayName = &v
	return s
}

func (s *DataSourceItem) SetEndpoint(v string) *DataSourceItem {
	s.Endpoint = &v
	return s
}

func (s *DataSourceItem) SetFileSystemId(v string) *DataSourceItem {
	s.FileSystemId = &v
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

func (s *DataSourceItem) SetMountPath(v string) *DataSourceItem {
	s.MountPath = &v
	return s
}

func (s *DataSourceItem) SetOptions(v string) *DataSourceItem {
	s.Options = &v
	return s
}

func (s *DataSourceItem) SetPath(v string) *DataSourceItem {
	s.Path = &v
	return s
}

func (s *DataSourceItem) SetUserId(v string) *DataSourceItem {
	s.UserId = &v
	return s
}

type DebuggerConfig struct {
	// 配置项细节，json结构
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// 系统生成的debug config唯一ID
	DebuggerConfigId *string `json:"DebuggerConfigId,omitempty" xml:"DebuggerConfigId,omitempty"`
	// 配置项描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 模板配置项名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间（UTC）
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
}

func (s DebuggerConfig) String() string {
	return tea.Prettify(s)
}

func (s DebuggerConfig) GoString() string {
	return s.String()
}

func (s *DebuggerConfig) SetContent(v string) *DebuggerConfig {
	s.Content = &v
	return s
}

func (s *DebuggerConfig) SetDebuggerConfigId(v string) *DebuggerConfig {
	s.DebuggerConfigId = &v
	return s
}

func (s *DebuggerConfig) SetDescription(v string) *DebuggerConfig {
	s.Description = &v
	return s
}

func (s *DebuggerConfig) SetDisplayName(v string) *DebuggerConfig {
	s.DisplayName = &v
	return s
}

func (s *DebuggerConfig) SetGmtCreateTime(v string) *DebuggerConfig {
	s.GmtCreateTime = &v
	return s
}

func (s *DebuggerConfig) SetGmtModifyTime(v string) *DebuggerConfig {
	s.GmtModifyTime = &v
	return s
}

type DebuggerJob struct {
	// debugger分析作业ID
	DebuggerJobId *string `json:"DebuggerJobId,omitempty" xml:"DebuggerJobId,omitempty"`
	// 作业显示名
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 作业运行的时长（单位秒）
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 失败时间（UTC）
	GmtFailedTime *string `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	// 任务完成时间（UTC）
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// 运行开始时间（UTC）
	GmtRunningTime *string `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	// 结束时间（UTC）
	GmtStoppedTime *string `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	// 提交时间（UTC）
	GmtSubmittedTime *string `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	// 成功结束时间（UTC）
	GmtSucceedTime *string `json:"GmtSucceedTime,omitempty" xml:"GmtSucceedTime,omitempty"`
	// 作业运行状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 用户ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 作业所属的运行工作空间
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 作业所属的运行工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s DebuggerJob) String() string {
	return tea.Prettify(s)
}

func (s DebuggerJob) GoString() string {
	return s.String()
}

func (s *DebuggerJob) SetDebuggerJobId(v string) *DebuggerJob {
	s.DebuggerJobId = &v
	return s
}

func (s *DebuggerJob) SetDisplayName(v string) *DebuggerJob {
	s.DisplayName = &v
	return s
}

func (s *DebuggerJob) SetDuration(v string) *DebuggerJob {
	s.Duration = &v
	return s
}

func (s *DebuggerJob) SetGmtCreateTime(v string) *DebuggerJob {
	s.GmtCreateTime = &v
	return s
}

func (s *DebuggerJob) SetGmtFailedTime(v string) *DebuggerJob {
	s.GmtFailedTime = &v
	return s
}

func (s *DebuggerJob) SetGmtFinishTime(v string) *DebuggerJob {
	s.GmtFinishTime = &v
	return s
}

func (s *DebuggerJob) SetGmtRunningTime(v string) *DebuggerJob {
	s.GmtRunningTime = &v
	return s
}

func (s *DebuggerJob) SetGmtStoppedTime(v string) *DebuggerJob {
	s.GmtStoppedTime = &v
	return s
}

func (s *DebuggerJob) SetGmtSubmittedTime(v string) *DebuggerJob {
	s.GmtSubmittedTime = &v
	return s
}

func (s *DebuggerJob) SetGmtSucceedTime(v string) *DebuggerJob {
	s.GmtSucceedTime = &v
	return s
}

func (s *DebuggerJob) SetStatus(v string) *DebuggerJob {
	s.Status = &v
	return s
}

func (s *DebuggerJob) SetUserId(v string) *DebuggerJob {
	s.UserId = &v
	return s
}

func (s *DebuggerJob) SetWorkspaceId(v string) *DebuggerJob {
	s.WorkspaceId = &v
	return s
}

func (s *DebuggerJob) SetWorkspaceName(v string) *DebuggerJob {
	s.WorkspaceName = &v
	return s
}

type DebuggerJobIssue struct {
	// 规则报告的具体json描述
	DebuggerJobIssue *string `json:"DebuggerJobIssue,omitempty" xml:"DebuggerJobIssue,omitempty"`
	// 创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// event的全局唯一ID
	JobDebuggerIssueId *string `json:"JobDebuggerIssueId,omitempty" xml:"JobDebuggerIssueId,omitempty"`
	// 作业ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 规则触发原因的编码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 规则触发的原因
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 规则名称
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
}

func (s DebuggerJobIssue) String() string {
	return tea.Prettify(s)
}

func (s DebuggerJobIssue) GoString() string {
	return s.String()
}

func (s *DebuggerJobIssue) SetDebuggerJobIssue(v string) *DebuggerJobIssue {
	s.DebuggerJobIssue = &v
	return s
}

func (s *DebuggerJobIssue) SetGmtCreateTime(v string) *DebuggerJobIssue {
	s.GmtCreateTime = &v
	return s
}

func (s *DebuggerJobIssue) SetJobDebuggerIssueId(v string) *DebuggerJobIssue {
	s.JobDebuggerIssueId = &v
	return s
}

func (s *DebuggerJobIssue) SetJobId(v string) *DebuggerJobIssue {
	s.JobId = &v
	return s
}

func (s *DebuggerJobIssue) SetReasonCode(v string) *DebuggerJobIssue {
	s.ReasonCode = &v
	return s
}

func (s *DebuggerJobIssue) SetReasonMessage(v string) *DebuggerJobIssue {
	s.ReasonMessage = &v
	return s
}

func (s *DebuggerJobIssue) SetRuleName(v string) *DebuggerJobIssue {
	s.RuleName = &v
	return s
}

type DebuggerResult struct {
	// 配置项细节，json结构
	DebuggerConfigContent *string `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	// 规则报告细节信息
	DebuggerJobIssues *string `json:"DebuggerJobIssues,omitempty" xml:"DebuggerJobIssues,omitempty"`
	// debugger job运行状态描述
	DebuggerJobStatus *string `json:"DebuggerJobStatus,omitempty" xml:"DebuggerJobStatus,omitempty"`
	// 报告文件下载地址
	DebuggerReportURL *string `json:"DebuggerReportURL,omitempty" xml:"DebuggerReportURL,omitempty"`
	// 作业显示名
	JobDisplayName *string `json:"JobDisplayName,omitempty" xml:"JobDisplayName,omitempty"`
	// 作业ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 用户ID
	JobUserId *string `json:"JobUserId,omitempty" xml:"JobUserId,omitempty"`
}

func (s DebuggerResult) String() string {
	return tea.Prettify(s)
}

func (s DebuggerResult) GoString() string {
	return s.String()
}

func (s *DebuggerResult) SetDebuggerConfigContent(v string) *DebuggerResult {
	s.DebuggerConfigContent = &v
	return s
}

func (s *DebuggerResult) SetDebuggerJobIssues(v string) *DebuggerResult {
	s.DebuggerJobIssues = &v
	return s
}

func (s *DebuggerResult) SetDebuggerJobStatus(v string) *DebuggerResult {
	s.DebuggerJobStatus = &v
	return s
}

func (s *DebuggerResult) SetDebuggerReportURL(v string) *DebuggerResult {
	s.DebuggerReportURL = &v
	return s
}

func (s *DebuggerResult) SetJobDisplayName(v string) *DebuggerResult {
	s.JobDisplayName = &v
	return s
}

func (s *DebuggerResult) SetJobId(v string) *DebuggerResult {
	s.JobId = &v
	return s
}

func (s *DebuggerResult) SetJobUserId(v string) *DebuggerResult {
	s.JobUserId = &v
	return s
}

type EcsSpec struct {
	// 加速器类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// cpu数量
	Cpu *int32 `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	// gpu数量
	Gpu *int32 `json:"Gpu,omitempty" xml:"Gpu,omitempty"`
	// gpu类型
	GpuType *string `json:"GpuType,omitempty" xml:"GpuType,omitempty"`
	// 规格类型
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// 是否有库存
	IsAvailable *bool `json:"IsAvailable,omitempty" xml:"IsAvailable,omitempty"`
	// Memory数量
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s EcsSpec) String() string {
	return tea.Prettify(s)
}

func (s EcsSpec) GoString() string {
	return s.String()
}

func (s *EcsSpec) SetAcceleratorType(v string) *EcsSpec {
	s.AcceleratorType = &v
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

func (s *EcsSpec) SetInstanceType(v string) *EcsSpec {
	s.InstanceType = &v
	return s
}

func (s *EcsSpec) SetIsAvailable(v bool) *EcsSpec {
	s.IsAvailable = &v
	return s
}

func (s *EcsSpec) SetMemory(v int32) *EcsSpec {
	s.Memory = &v
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

type ExtraPodSpec struct {
	// 初始化容器
	InitContainers []*ContainerSpec `json:"InitContainers,omitempty" xml:"InitContainers,omitempty" type:"Repeated"`
	// Pod注解
	PodAnnotations map[string]*string `json:"PodAnnotations,omitempty" xml:"PodAnnotations,omitempty"`
	// Pod标签
	PodLabels map[string]*string `json:"PodLabels,omitempty" xml:"PodLabels,omitempty"`
	// 容器间共享的本地目录
	SharedVolumeMountPaths []*string `json:"SharedVolumeMountPaths,omitempty" xml:"SharedVolumeMountPaths,omitempty" type:"Repeated"`
	// 伴随容器
	SideCarContainers []*ContainerSpec `json:"SideCarContainers,omitempty" xml:"SideCarContainers,omitempty" type:"Repeated"`
}

func (s ExtraPodSpec) String() string {
	return tea.Prettify(s)
}

func (s ExtraPodSpec) GoString() string {
	return s.String()
}

func (s *ExtraPodSpec) SetInitContainers(v []*ContainerSpec) *ExtraPodSpec {
	s.InitContainers = v
	return s
}

func (s *ExtraPodSpec) SetPodAnnotations(v map[string]*string) *ExtraPodSpec {
	s.PodAnnotations = v
	return s
}

func (s *ExtraPodSpec) SetPodLabels(v map[string]*string) *ExtraPodSpec {
	s.PodLabels = v
	return s
}

func (s *ExtraPodSpec) SetSharedVolumeMountPaths(v []*string) *ExtraPodSpec {
	s.SharedVolumeMountPaths = v
	return s
}

func (s *ExtraPodSpec) SetSideCarContainers(v []*ContainerSpec) *ExtraPodSpec {
	s.SideCarContainers = v
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

type ImageItem struct {
	// 加速器类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 镜像作者
	AuthorId *string `json:"AuthorId,omitempty" xml:"AuthorId,omitempty"`
	// 镜像包含的框架类型
	Framework *string `json:"Framework,omitempty" xml:"Framework,omitempty"`
	// 镜像类型
	ImageProviderType *string `json:"ImageProviderType,omitempty" xml:"ImageProviderType,omitempty"`
	// 镜像Tag
	ImageTag *string `json:"ImageTag,omitempty" xml:"ImageTag,omitempty"`
	// 镜像地址
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// 镜像vpc地址
	ImageUrlVpc *string `json:"ImageUrlVpc,omitempty" xml:"ImageUrlVpc,omitempty"`
}

func (s ImageItem) String() string {
	return tea.Prettify(s)
}

func (s ImageItem) GoString() string {
	return s.String()
}

func (s *ImageItem) SetAcceleratorType(v string) *ImageItem {
	s.AcceleratorType = &v
	return s
}

func (s *ImageItem) SetAuthorId(v string) *ImageItem {
	s.AuthorId = &v
	return s
}

func (s *ImageItem) SetFramework(v string) *ImageItem {
	s.Framework = &v
	return s
}

func (s *ImageItem) SetImageProviderType(v string) *ImageItem {
	s.ImageProviderType = &v
	return s
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

type JobDebuggerConfig struct {
	// debugger配置信息
	DebuggerConfigContent *string `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	// debugger配置项ID
	DebuggerConfigId *string `json:"DebuggerConfigId,omitempty" xml:"DebuggerConfigId,omitempty"`
	// debugger配置创建的时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 作业ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s JobDebuggerConfig) String() string {
	return tea.Prettify(s)
}

func (s JobDebuggerConfig) GoString() string {
	return s.String()
}

func (s *JobDebuggerConfig) SetDebuggerConfigContent(v string) *JobDebuggerConfig {
	s.DebuggerConfigContent = &v
	return s
}

func (s *JobDebuggerConfig) SetDebuggerConfigId(v string) *JobDebuggerConfig {
	s.DebuggerConfigId = &v
	return s
}

func (s *JobDebuggerConfig) SetGmtCreateTime(v string) *JobDebuggerConfig {
	s.GmtCreateTime = &v
	return s
}

func (s *JobDebuggerConfig) SetJobId(v string) *JobDebuggerConfig {
	s.JobId = &v
	return s
}

type JobElasticSpec struct {
	// aimaster角色使用的资源规格
	AIMasterType *string `json:"AIMasterType,omitempty" xml:"AIMasterType,omitempty"`
	// 打开弹性训练
	EnableElasticTraining *bool `json:"EnableElasticTraining,omitempty" xml:"EnableElasticTraining,omitempty"`
	// 最大并行度
	MaxParallelism *int32 `json:"MaxParallelism,omitempty" xml:"MaxParallelism,omitempty"`
	// 最小并行度
	MinParallelism *int32 `json:"MinParallelism,omitempty" xml:"MinParallelism,omitempty"`
}

func (s JobElasticSpec) String() string {
	return tea.Prettify(s)
}

func (s JobElasticSpec) GoString() string {
	return s.String()
}

func (s *JobElasticSpec) SetAIMasterType(v string) *JobElasticSpec {
	s.AIMasterType = &v
	return s
}

func (s *JobElasticSpec) SetEnableElasticTraining(v bool) *JobElasticSpec {
	s.EnableElasticTraining = &v
	return s
}

func (s *JobElasticSpec) SetMaxParallelism(v int32) *JobElasticSpec {
	s.MaxParallelism = &v
	return s
}

func (s *JobElasticSpec) SetMinParallelism(v int32) *JobElasticSpec {
	s.MinParallelism = &v
	return s
}

type JobItem struct {
	// 代码源配置
	CodeSource *JobItemCodeSource `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	// 数据源配置列表
	DataSources []*JobItemDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// 作业显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 作业运行时长，单位：秒
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 是否开启了debugger分析
	EnabledDebugger *bool `json:"EnabledDebugger,omitempty" xml:"EnabledDebugger,omitempty"`
	// 环境变量配置
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// 作业创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 作业结束时间（UTC）
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// 作业Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 作业规格配置
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// 作业类型
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// 状态详情码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 状态详情
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 作业运行所在的资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 作业运行时的资源级别
	ResourceLevel *string `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
	// 作业运行的资源名称
	ResourceName *string `json:"ResourceName,omitempty" xml:"ResourceName,omitempty"`
	// 作业额外参数
	Settings *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// 作业状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 三方库(requirements.txt)文件路径
	ThirdpartyLibDir *string `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	// 三方库配置列表
	ThirdpartyLibs []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// 用户命令
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// 作业提交人Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 作业所属工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 作业所属工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s JobItem) String() string {
	return tea.Prettify(s)
}

func (s JobItem) GoString() string {
	return s.String()
}

func (s *JobItem) SetCodeSource(v *JobItemCodeSource) *JobItem {
	s.CodeSource = v
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

func (s *JobItem) SetGmtFinishTime(v string) *JobItem {
	s.GmtFinishTime = &v
	return s
}

func (s *JobItem) SetJobId(v string) *JobItem {
	s.JobId = &v
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

func (s *JobItem) SetReasonCode(v string) *JobItem {
	s.ReasonCode = &v
	return s
}

func (s *JobItem) SetReasonMessage(v string) *JobItem {
	s.ReasonMessage = &v
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

func (s *JobItem) SetSettings(v *JobSettings) *JobItem {
	s.Settings = v
	return s
}

func (s *JobItem) SetStatus(v string) *JobItem {
	s.Status = &v
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

func (s *JobItem) SetUserCommand(v string) *JobItem {
	s.UserCommand = &v
	return s
}

func (s *JobItem) SetUserId(v string) *JobItem {
	s.UserId = &v
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

type JobItemCodeSource struct {
	// 代码分支
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// 代码源Id
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
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

type JobSettings struct {
	// 作业关联用户ID
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// 调用方
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// 允许打开作业容错监控
	EnableErrorMonitoringInAIMaster *bool `json:"EnableErrorMonitoringInAIMaster,omitempty" xml:"EnableErrorMonitoringInAIMaster,omitempty"`
	// 允许作业使用RDMA
	EnableRDMA *bool `json:"EnableRDMA,omitempty" xml:"EnableRDMA,omitempty"`
	// 允许作业使用潮汐资源
	EnableTideResource *bool `json:"EnableTideResource,omitempty" xml:"EnableTideResource,omitempty"`
	// 用户指定容错监控的配置参数，比如指定是否启动基于log hang的检测
	ErrorMonitoringArgs *string `json:"ErrorMonitoringArgs,omitempty" xml:"ErrorMonitoringArgs,omitempty"`
	// 工作流ID
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// 自定义标签
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s *JobSettings) SetEnableErrorMonitoringInAIMaster(v bool) *JobSettings {
	s.EnableErrorMonitoringInAIMaster = &v
	return s
}

func (s *JobSettings) SetEnableRDMA(v bool) *JobSettings {
	s.EnableRDMA = &v
	return s
}

func (s *JobSettings) SetEnableTideResource(v bool) *JobSettings {
	s.EnableTideResource = &v
	return s
}

func (s *JobSettings) SetErrorMonitoringArgs(v string) *JobSettings {
	s.ErrorMonitoringArgs = &v
	return s
}

func (s *JobSettings) SetPipelineId(v string) *JobSettings {
	s.PipelineId = &v
	return s
}

func (s *JobSettings) SetTags(v map[string]*string) *JobSettings {
	s.Tags = v
	return s
}

type JobSpec struct {
	// Ecs实例规格
	EcsSpec *string `json:"EcsSpec,omitempty" xml:"EcsSpec,omitempty"`
	// 额外的Pod配置
	ExtraPodSpec *ExtraPodSpec `json:"ExtraPodSpec,omitempty" xml:"ExtraPodSpec,omitempty"`
	// 镜像
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// 实例数量
	PodCount *int64 `json:"PodCount,omitempty" xml:"PodCount,omitempty"`
	// 资源配置
	ResourceConfig *ResourceConfig `json:"ResourceConfig,omitempty" xml:"ResourceConfig,omitempty"`
	// 类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 是否使用竞价实例
	UseSpotInstance *bool `json:"UseSpotInstance,omitempty" xml:"UseSpotInstance,omitempty"`
}

func (s JobSpec) String() string {
	return tea.Prettify(s)
}

func (s JobSpec) GoString() string {
	return s.String()
}

func (s *JobSpec) SetEcsSpec(v string) *JobSpec {
	s.EcsSpec = &v
	return s
}

func (s *JobSpec) SetExtraPodSpec(v *ExtraPodSpec) *JobSpec {
	s.ExtraPodSpec = v
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

func (s *JobSpec) SetResourceConfig(v *ResourceConfig) *JobSpec {
	s.ResourceConfig = v
	return s
}

func (s *JobSpec) SetType(v string) *JobSpec {
	s.Type = &v
	return s
}

func (s *JobSpec) SetUseSpotInstance(v bool) *JobSpec {
	s.UseSpotInstance = &v
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

type NodeMetric struct {
	// 监控指标样本序列
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// 节点名称
	NodeName *string `json:"NodeName,omitempty" xml:"NodeName,omitempty"`
}

func (s NodeMetric) String() string {
	return tea.Prettify(s)
}

func (s NodeMetric) GoString() string {
	return s.String()
}

func (s *NodeMetric) SetMetrics(v []*Metric) *NodeMetric {
	s.Metrics = v
	return s
}

func (s *NodeMetric) SetNodeName(v string) *NodeMetric {
	s.NodeName = &v
	return s
}

type PodMetric struct {
	// 监控指标样本序列
	Metrics []*Metric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// Pod编号
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
}

func (s PodMetric) String() string {
	return tea.Prettify(s)
}

func (s PodMetric) GoString() string {
	return s.String()
}

func (s *PodMetric) SetMetrics(v []*Metric) *PodMetric {
	s.Metrics = v
	return s
}

func (s *PodMetric) SetPodId(v string) *PodMetric {
	s.PodId = &v
	return s
}

type Quota struct {
	// 集群id
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 集群名称
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// 是否允许使用潮汐资源
	EnableTideResource *bool `json:"EnableTideResource,omitempty" xml:"EnableTideResource,omitempty"`
	// 是否是独占Quota组
	IsExclusiveQuota *bool `json:"IsExclusiveQuota,omitempty" xml:"IsExclusiveQuota,omitempty"`
	// 资源配额id
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// 资源配额名称
	QuotaName *string `json:"QuotaName,omitempty" xml:"QuotaName,omitempty"`
	// 资源配额类型
	QuotaType *string `json:"QuotaType,omitempty" xml:"QuotaType,omitempty"`
	// 资源组允许使用的潮汐资源级别
	ResourceLevel *string `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
	// 资源总量
	TotalQuota *QuotaDetail `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// 潮汐资源总量
	TotalTideQuota *QuotaDetail `json:"TotalTideQuota,omitempty" xml:"TotalTideQuota,omitempty"`
	// 资源用量
	UsedQuota *QuotaDetail `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
	// 潮汐资源用量
	UsedTideQuota *QuotaDetail `json:"UsedTideQuota,omitempty" xml:"UsedTideQuota,omitempty"`
}

func (s Quota) String() string {
	return tea.Prettify(s)
}

func (s Quota) GoString() string {
	return s.String()
}

func (s *Quota) SetClusterId(v string) *Quota {
	s.ClusterId = &v
	return s
}

func (s *Quota) SetClusterName(v string) *Quota {
	s.ClusterName = &v
	return s
}

func (s *Quota) SetEnableTideResource(v bool) *Quota {
	s.EnableTideResource = &v
	return s
}

func (s *Quota) SetIsExclusiveQuota(v bool) *Quota {
	s.IsExclusiveQuota = &v
	return s
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

func (s *Quota) SetResourceLevel(v string) *Quota {
	s.ResourceLevel = &v
	return s
}

func (s *Quota) SetTotalQuota(v *QuotaDetail) *Quota {
	s.TotalQuota = v
	return s
}

func (s *Quota) SetTotalTideQuota(v *QuotaDetail) *Quota {
	s.TotalTideQuota = v
	return s
}

func (s *Quota) SetUsedQuota(v *QuotaDetail) *Quota {
	s.UsedQuota = v
	return s
}

func (s *Quota) SetUsedTideQuota(v *QuotaDetail) *Quota {
	s.UsedTideQuota = v
	return s
}

type QuotaDetail struct {
	// CPU核数
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// GPU卡数
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// GPU详情
	GPUDetails []*GPUDetail `json:"GPUDetails,omitempty" xml:"GPUDetails,omitempty" type:"Repeated"`
	// GPU卡型
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// GPU卡型全名
	GPUTypeFullName *string `json:"GPUTypeFullName,omitempty" xml:"GPUTypeFullName,omitempty"`
	// 内存容量
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
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

func (s *QuotaDetail) SetGPU(v string) *QuotaDetail {
	s.GPU = &v
	return s
}

func (s *QuotaDetail) SetGPUDetails(v []*GPUDetail) *QuotaDetail {
	s.GPUDetails = v
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

func (s *QuotaDetail) SetMemory(v string) *QuotaDetail {
	s.Memory = &v
	return s
}

type ResourceConfig struct {
	// CPU核心数
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// GPU核心数
	GPU *string `json:"GPU,omitempty" xml:"GPU,omitempty"`
	// 显卡类型
	GPUType *string `json:"GPUType,omitempty" xml:"GPUType,omitempty"`
	// 内存容量
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// 共享内存容量
	SharedMemory *string `json:"SharedMemory,omitempty" xml:"SharedMemory,omitempty"`
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

func (s *ResourceConfig) SetGPUType(v string) *ResourceConfig {
	s.GPUType = &v
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

type ResourceRequirements struct {
	// 资源限制
	Limits map[string]*string `json:"Limits,omitempty" xml:"Limits,omitempty"`
	// 资源需求
	Requests map[string]*string `json:"Requests,omitempty" xml:"Requests,omitempty"`
}

func (s ResourceRequirements) String() string {
	return tea.Prettify(s)
}

func (s ResourceRequirements) GoString() string {
	return s.String()
}

func (s *ResourceRequirements) SetLimits(v map[string]*string) *ResourceRequirements {
	s.Limits = v
	return s
}

func (s *ResourceRequirements) SetRequests(v map[string]*string) *ResourceRequirements {
	s.Requests = v
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

type SmartCache struct {
	// cache worker的数量单位是个
	CacheWorkerNum *int64 `json:"CacheWorkerNum,omitempty" xml:"CacheWorkerNum,omitempty"`
	// 每个cache worker的cache大小单位是GB
	CacheWorkerSize *int64 `json:"CacheWorkerSize,omitempty" xml:"CacheWorkerSize,omitempty"`
	// SmartCache 描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// SmartCache 名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 已运行时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// OSS Endpoint
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// 阿里云的NAS文件系统ID
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// 创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间（UTC）
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 数据源本地挂载目录
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 可选的超参数
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// OSS数据源路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// SmartCacheId
	SmartCacheId *string `json:"SmartCacheId,omitempty" xml:"SmartCacheId,omitempty"`
	// 运行状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 挂载的数据类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// 创建者Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s SmartCache) String() string {
	return tea.Prettify(s)
}

func (s SmartCache) GoString() string {
	return s.String()
}

func (s *SmartCache) SetCacheWorkerNum(v int64) *SmartCache {
	s.CacheWorkerNum = &v
	return s
}

func (s *SmartCache) SetCacheWorkerSize(v int64) *SmartCache {
	s.CacheWorkerSize = &v
	return s
}

func (s *SmartCache) SetDescription(v string) *SmartCache {
	s.Description = &v
	return s
}

func (s *SmartCache) SetDisplayName(v string) *SmartCache {
	s.DisplayName = &v
	return s
}

func (s *SmartCache) SetDuration(v string) *SmartCache {
	s.Duration = &v
	return s
}

func (s *SmartCache) SetEndpoint(v string) *SmartCache {
	s.Endpoint = &v
	return s
}

func (s *SmartCache) SetFileSystemId(v string) *SmartCache {
	s.FileSystemId = &v
	return s
}

func (s *SmartCache) SetGmtCreateTime(v string) *SmartCache {
	s.GmtCreateTime = &v
	return s
}

func (s *SmartCache) SetGmtModifyTime(v string) *SmartCache {
	s.GmtModifyTime = &v
	return s
}

func (s *SmartCache) SetMountPath(v string) *SmartCache {
	s.MountPath = &v
	return s
}

func (s *SmartCache) SetOptions(v string) *SmartCache {
	s.Options = &v
	return s
}

func (s *SmartCache) SetPath(v string) *SmartCache {
	s.Path = &v
	return s
}

func (s *SmartCache) SetSmartCacheId(v string) *SmartCache {
	s.SmartCacheId = &v
	return s
}

func (s *SmartCache) SetStatus(v string) *SmartCache {
	s.Status = &v
	return s
}

func (s *SmartCache) SetType(v string) *SmartCache {
	s.Type = &v
	return s
}

func (s *SmartCache) SetUserId(v string) *SmartCache {
	s.UserId = &v
	return s
}

type Tensorboard struct {
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 展示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 运行时长
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间（UTC）
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 状态详情码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 状态详情
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 文件路径
	SummaryPath *string `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	// Tensorboard Id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	// Tensorboard URL
	TensorboardUrl *string `json:"TensorboardUrl,omitempty" xml:"TensorboardUrl,omitempty"`
	// 创建者
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s Tensorboard) String() string {
	return tea.Prettify(s)
}

func (s Tensorboard) GoString() string {
	return s.String()
}

func (s *Tensorboard) SetDataSourceId(v string) *Tensorboard {
	s.DataSourceId = &v
	return s
}

func (s *Tensorboard) SetDisplayName(v string) *Tensorboard {
	s.DisplayName = &v
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

func (s *Tensorboard) SetJobId(v string) *Tensorboard {
	s.JobId = &v
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

func (s *Tensorboard) SetRequestId(v string) *Tensorboard {
	s.RequestId = &v
	return s
}

func (s *Tensorboard) SetStatus(v string) *Tensorboard {
	s.Status = &v
	return s
}

func (s *Tensorboard) SetSummaryPath(v string) *Tensorboard {
	s.SummaryPath = &v
	return s
}

func (s *Tensorboard) SetTensorboardId(v string) *Tensorboard {
	s.TensorboardId = &v
	return s
}

func (s *Tensorboard) SetTensorboardUrl(v string) *Tensorboard {
	s.TensorboardUrl = &v
	return s
}

func (s *Tensorboard) SetUserId(v string) *Tensorboard {
	s.UserId = &v
	return s
}

type Workspace struct {
	// 创建者
	Creator *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 最近修改时间
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 成员列表
	Members []*Member `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// 资源配额列表
	Quotas []*Quota `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// 资源总量
	TotalResources *Resources `json:"TotalResources,omitempty" xml:"TotalResources,omitempty"`
	// 管理员列表
	WorkspaceAdmins []*Member `json:"WorkspaceAdmins,omitempty" xml:"WorkspaceAdmins,omitempty" type:"Repeated"`
	// 工作空间id
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s Workspace) String() string {
	return tea.Prettify(s)
}

func (s Workspace) GoString() string {
	return s.String()
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

func (s *Workspace) SetMembers(v []*Member) *Workspace {
	s.Members = v
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

func (s *Workspace) SetWorkspaceAdmins(v []*Member) *Workspace {
	s.WorkspaceAdmins = v
	return s
}

func (s *Workspace) SetWorkspaceId(v string) *Workspace {
	s.WorkspaceId = &v
	return s
}

func (s *Workspace) SetWorkspaceName(v string) *Workspace {
	s.WorkspaceName = &v
	return s
}

type CreateCodeSourceRequest struct {
	// 代码分支
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// 代码仓库地址
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// 密码
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// 用户名
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// 代码源详细描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代码源配置名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 代码本地挂载目录，默认挂载到/root/code/下
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
}

func (s CreateCodeSourceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCodeSourceRequest) GoString() string {
	return s.String()
}

func (s *CreateCodeSourceRequest) SetCodeBranch(v string) *CreateCodeSourceRequest {
	s.CodeBranch = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepo(v string) *CreateCodeSourceRequest {
	s.CodeRepo = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepoAccessToken(v string) *CreateCodeSourceRequest {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *CreateCodeSourceRequest) SetCodeRepoUserName(v string) *CreateCodeSourceRequest {
	s.CodeRepoUserName = &v
	return s
}

func (s *CreateCodeSourceRequest) SetDescription(v string) *CreateCodeSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateCodeSourceRequest) SetDisplayName(v string) *CreateCodeSourceRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateCodeSourceRequest) SetMountPath(v string) *CreateCodeSourceRequest {
	s.MountPath = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateCodeSourceResponse) SetStatusCode(v int32) *CreateCodeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCodeSourceResponse) SetBody(v *CreateCodeSourceResponseBody) *CreateCodeSourceResponse {
	s.Body = v
	return s
}

type CreateDataSourceRequest struct {
	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// 数据源描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 数据源显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 阿里云OSS文件系统服务端点
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// 阿里云NAS文件系统Id
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// 本地挂载目录
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 阿里云OSS文件系统配置选项
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// 阿里云OSS文件系统路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
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

func (s *CreateDataSourceRequest) SetDescription(v string) *CreateDataSourceRequest {
	s.Description = &v
	return s
}

func (s *CreateDataSourceRequest) SetDisplayName(v string) *CreateDataSourceRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateDataSourceRequest) SetEndpoint(v string) *CreateDataSourceRequest {
	s.Endpoint = &v
	return s
}

func (s *CreateDataSourceRequest) SetFileSystemId(v string) *CreateDataSourceRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDataSourceRequest) SetMountPath(v string) *CreateDataSourceRequest {
	s.MountPath = &v
	return s
}

func (s *CreateDataSourceRequest) SetOptions(v string) *CreateDataSourceRequest {
	s.Options = &v
	return s
}

func (s *CreateDataSourceRequest) SetPath(v string) *CreateDataSourceRequest {
	s.Path = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateDataSourceResponse) SetStatusCode(v int32) *CreateDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataSourceResponse) SetBody(v *CreateDataSourceResponseBody) *CreateDataSourceResponse {
	s.Body = v
	return s
}

type CreateJobRequest struct {
	// 代码源配置
	CodeSource *CreateJobRequestCodeSource `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	// 数据源配置列表
	DataSources []*CreateJobRequestDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// debugger参数
	DebuggerConfigContent *string `json:"DebuggerConfigContent,omitempty" xml:"DebuggerConfigContent,omitempty"`
	// 作业显示名称
	DisplayName *string         `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	ElasticSpec *JobElasticSpec `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
	// 环境变量配置
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// 作业最大运行时间
	JobMaxRunningTimeMinutes *int64 `json:"JobMaxRunningTimeMinutes,omitempty" xml:"JobMaxRunningTimeMinutes,omitempty"`
	// 作业规格配置
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// 作业类型
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// 作业优先级
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// 资源组编号
	ResourceId *string      `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Settings   *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// 三方库(requirements.txt)文件路径
	ThirdpartyLibDir *string `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	// 三方库配置列表
	ThirdpartyLibs []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// 作业命令
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// 用户VPC
	UserVpc *CreateJobRequestUserVpc `json:"UserVpc,omitempty" xml:"UserVpc,omitempty" type:"Struct"`
	// 工作空间编号
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequest) GoString() string {
	return s.String()
}

func (s *CreateJobRequest) SetCodeSource(v *CreateJobRequestCodeSource) *CreateJobRequest {
	s.CodeSource = v
	return s
}

func (s *CreateJobRequest) SetDataSources(v []*CreateJobRequestDataSources) *CreateJobRequest {
	s.DataSources = v
	return s
}

func (s *CreateJobRequest) SetDebuggerConfigContent(v string) *CreateJobRequest {
	s.DebuggerConfigContent = &v
	return s
}

func (s *CreateJobRequest) SetDisplayName(v string) *CreateJobRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateJobRequest) SetElasticSpec(v *JobElasticSpec) *CreateJobRequest {
	s.ElasticSpec = v
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

func (s *CreateJobRequest) SetJobSpecs(v []*JobSpec) *CreateJobRequest {
	s.JobSpecs = v
	return s
}

func (s *CreateJobRequest) SetJobType(v string) *CreateJobRequest {
	s.JobType = &v
	return s
}

func (s *CreateJobRequest) SetOptions(v string) *CreateJobRequest {
	s.Options = &v
	return s
}

func (s *CreateJobRequest) SetPriority(v int32) *CreateJobRequest {
	s.Priority = &v
	return s
}

func (s *CreateJobRequest) SetResourceId(v string) *CreateJobRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateJobRequest) SetSettings(v *JobSettings) *CreateJobRequest {
	s.Settings = v
	return s
}

func (s *CreateJobRequest) SetThirdpartyLibDir(v string) *CreateJobRequest {
	s.ThirdpartyLibDir = &v
	return s
}

func (s *CreateJobRequest) SetThirdpartyLibs(v []*string) *CreateJobRequest {
	s.ThirdpartyLibs = v
	return s
}

func (s *CreateJobRequest) SetUserCommand(v string) *CreateJobRequest {
	s.UserCommand = &v
	return s
}

func (s *CreateJobRequest) SetUserVpc(v *CreateJobRequestUserVpc) *CreateJobRequest {
	s.UserVpc = v
	return s
}

func (s *CreateJobRequest) SetWorkspaceId(v string) *CreateJobRequest {
	s.WorkspaceId = &v
	return s
}

type CreateJobRequestCodeSource struct {
	// 代码分支
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// 代码源Id
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
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

func (s *CreateJobRequestCodeSource) SetBranch(v string) *CreateJobRequestCodeSource {
	s.Branch = &v
	return s
}

func (s *CreateJobRequestCodeSource) SetCodeSourceId(v string) *CreateJobRequestCodeSource {
	s.CodeSourceId = &v
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

type CreateJobRequestUserVpc struct {
	// 扩展网段
	ExtendedCIDRs []*string `json:"ExtendedCIDRs,omitempty" xml:"ExtendedCIDRs,omitempty" type:"Repeated"`
	// 用户安全组的id
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// 用户交换机的id
	SwitchId *string `json:"SwitchId,omitempty" xml:"SwitchId,omitempty"`
	// 用户VPC的id
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateJobRequestUserVpc) String() string {
	return tea.Prettify(s)
}

func (s CreateJobRequestUserVpc) GoString() string {
	return s.String()
}

func (s *CreateJobRequestUserVpc) SetExtendedCIDRs(v []*string) *CreateJobRequestUserVpc {
	s.ExtendedCIDRs = v
	return s
}

func (s *CreateJobRequestUserVpc) SetSecurityGroupId(v string) *CreateJobRequestUserVpc {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateJobRequestUserVpc) SetSwitchId(v string) *CreateJobRequestUserVpc {
	s.SwitchId = &v
	return s
}

func (s *CreateJobRequestUserVpc) SetVpcId(v string) *CreateJobRequestUserVpc {
	s.VpcId = &v
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateJobResponse) SetStatusCode(v int32) *CreateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateJobResponse) SetBody(v *CreateJobResponseBody) *CreateJobResponse {
	s.Body = v
	return s
}

type CreateTensorboardRequest struct {
	// DataSource Id
	DataSourceId   *string           `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	DataSourceType *string           `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	DataSources    []*DataSourceItem `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// Tensorboard名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 最长运行时长
	MaxRunningTimeMinutes *int64  `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
	Options               *string `json:"Options,omitempty" xml:"Options,omitempty"`
	SourceId              *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType            *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// Summary 目录
	SummaryPath         *string `json:"SummaryPath,omitempty" xml:"SummaryPath,omitempty"`
	SummaryRelativePath *string `json:"SummaryRelativePath,omitempty" xml:"SummaryRelativePath,omitempty"`
	Uri                 *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
	WorkspaceId         *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTensorboardRequest) GoString() string {
	return s.String()
}

func (s *CreateTensorboardRequest) SetDataSourceId(v string) *CreateTensorboardRequest {
	s.DataSourceId = &v
	return s
}

func (s *CreateTensorboardRequest) SetDataSourceType(v string) *CreateTensorboardRequest {
	s.DataSourceType = &v
	return s
}

func (s *CreateTensorboardRequest) SetDataSources(v []*DataSourceItem) *CreateTensorboardRequest {
	s.DataSources = v
	return s
}

func (s *CreateTensorboardRequest) SetDisplayName(v string) *CreateTensorboardRequest {
	s.DisplayName = &v
	return s
}

func (s *CreateTensorboardRequest) SetJobId(v string) *CreateTensorboardRequest {
	s.JobId = &v
	return s
}

func (s *CreateTensorboardRequest) SetMaxRunningTimeMinutes(v int64) *CreateTensorboardRequest {
	s.MaxRunningTimeMinutes = &v
	return s
}

func (s *CreateTensorboardRequest) SetOptions(v string) *CreateTensorboardRequest {
	s.Options = &v
	return s
}

func (s *CreateTensorboardRequest) SetSourceId(v string) *CreateTensorboardRequest {
	s.SourceId = &v
	return s
}

func (s *CreateTensorboardRequest) SetSourceType(v string) *CreateTensorboardRequest {
	s.SourceType = &v
	return s
}

func (s *CreateTensorboardRequest) SetSummaryPath(v string) *CreateTensorboardRequest {
	s.SummaryPath = &v
	return s
}

func (s *CreateTensorboardRequest) SetSummaryRelativePath(v string) *CreateTensorboardRequest {
	s.SummaryRelativePath = &v
	return s
}

func (s *CreateTensorboardRequest) SetUri(v string) *CreateTensorboardRequest {
	s.Uri = &v
	return s
}

func (s *CreateTensorboardRequest) SetWorkspaceId(v string) *CreateTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type CreateTensorboardResponseBody struct {
	// DataSourceId
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 任务Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Tensorboard id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s CreateTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTensorboardResponseBody) SetDataSourceId(v string) *CreateTensorboardResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetJobId(v string) *CreateTensorboardResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetRequestId(v string) *CreateTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTensorboardResponseBody) SetTensorboardId(v string) *CreateTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

type CreateTensorboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *CreateTensorboardResponse) SetStatusCode(v int32) *CreateTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTensorboardResponse) SetBody(v *CreateTensorboardResponseBody) *CreateTensorboardResponse {
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteCodeSourceResponse) SetStatusCode(v int32) *DeleteCodeSourceResponse {
	s.StatusCode = &v
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
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteDataSourceResponse) SetStatusCode(v int32) *DeleteDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDataSourceResponse) SetBody(v *DeleteDataSourceResponseBody) *DeleteDataSourceResponse {
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteJobResponse) SetStatusCode(v int32) *DeleteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobResponse) SetBody(v *DeleteJobResponseBody) *DeleteJobResponse {
	s.Body = v
	return s
}

type DeleteJobsRequest struct {
	// 作业ID列表
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
}

func (s DeleteJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsRequest) GoString() string {
	return s.String()
}

func (s *DeleteJobsRequest) SetJobIds(v []*string) *DeleteJobsRequest {
	s.JobIds = v
	return s
}

type DeleteJobsResponseBody struct {
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponseBody) SetJobIds(v []*string) *DeleteJobsResponseBody {
	s.JobIds = v
	return s
}

func (s *DeleteJobsResponseBody) SetRequestId(v string) *DeleteJobsResponseBody {
	s.RequestId = &v
	return s
}

type DeleteJobsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteJobsResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobsResponse) SetHeaders(v map[string]*string) *DeleteJobsResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobsResponse) SetStatusCode(v int32) *DeleteJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobsResponse) SetBody(v *DeleteJobsResponseBody) *DeleteJobsResponse {
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
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Tensorboad Id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s DeleteTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTensorboardResponseBody) SetRequestId(v string) *DeleteTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTensorboardResponseBody) SetTensorboardId(v string) *DeleteTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

type DeleteTensorboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *DeleteTensorboardResponse) SetStatusCode(v int32) *DeleteTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTensorboardResponse) SetBody(v *DeleteTensorboardResponseBody) *DeleteTensorboardResponse {
	s.Body = v
	return s
}

type GetCodeSourceResponseBody struct {
	// 代码仓库分支
	CodeBranch *string `json:"CodeBranch,omitempty" xml:"CodeBranch,omitempty"`
	// 代码Commit
	CodeCommit *string `json:"CodeCommit,omitempty" xml:"CodeCommit,omitempty"`
	// 代码仓库地址
	CodeRepo *string `json:"CodeRepo,omitempty" xml:"CodeRepo,omitempty"`
	// 访问代码仓库的token
	CodeRepoAccessToken *string `json:"CodeRepoAccessToken,omitempty" xml:"CodeRepoAccessToken,omitempty"`
	// 代码仓库的用户名
	CodeRepoUserName *string `json:"CodeRepoUserName,omitempty" xml:"CodeRepoUserName,omitempty"`
	// 代码源配置ID
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
	// 详细描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 代码源配置名字
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 创建时间
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 代码本地挂载目录，默认挂载到/root/code/下
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 代码配置源的创建者ID
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetCodeSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetCodeSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetCodeSourceResponseBody) SetCodeBranch(v string) *GetCodeSourceResponseBody {
	s.CodeBranch = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeCommit(v string) *GetCodeSourceResponseBody {
	s.CodeCommit = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepo(v string) *GetCodeSourceResponseBody {
	s.CodeRepo = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepoAccessToken(v string) *GetCodeSourceResponseBody {
	s.CodeRepoAccessToken = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeRepoUserName(v string) *GetCodeSourceResponseBody {
	s.CodeRepoUserName = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetCodeSourceId(v string) *GetCodeSourceResponseBody {
	s.CodeSourceId = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetDescription(v string) *GetCodeSourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetDisplayName(v string) *GetCodeSourceResponseBody {
	s.DisplayName = &v
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

func (s *GetCodeSourceResponseBody) SetMountPath(v string) *GetCodeSourceResponseBody {
	s.MountPath = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetRequestId(v string) *GetCodeSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCodeSourceResponseBody) SetUserId(v string) *GetCodeSourceResponseBody {
	s.UserId = &v
	return s
}

type GetCodeSourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetCodeSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetCodeSourceResponse) SetStatusCode(v int32) *GetCodeSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCodeSourceResponse) SetBody(v *GetCodeSourceResponseBody) *GetCodeSourceResponse {
	s.Body = v
	return s
}

type GetDataSourceResponseBody struct {
	// 数据源Id
	DataSourceId *string `json:"DataSourceId,omitempty" xml:"DataSourceId,omitempty"`
	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// 数据源描述
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// 数据源显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 阿里云OSS文件系统服务端点
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// 阿里云NAS文件系统Id
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// 创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// 修改时间（UTC）
	GmtModifyTime *string `json:"GmtModifyTime,omitempty" xml:"GmtModifyTime,omitempty"`
	// 本地挂载目录
	MountPath *string `json:"MountPath,omitempty" xml:"MountPath,omitempty"`
	// 阿里云OSS文件系统配置选项
	Options *string `json:"Options,omitempty" xml:"Options,omitempty"`
	// 阿里云OSS文件系统路径
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 创建人Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetDataSourceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataSourceResponseBody) SetDataSourceId(v string) *GetDataSourceResponseBody {
	s.DataSourceId = &v
	return s
}

func (s *GetDataSourceResponseBody) SetDataSourceType(v string) *GetDataSourceResponseBody {
	s.DataSourceType = &v
	return s
}

func (s *GetDataSourceResponseBody) SetDescription(v string) *GetDataSourceResponseBody {
	s.Description = &v
	return s
}

func (s *GetDataSourceResponseBody) SetDisplayName(v string) *GetDataSourceResponseBody {
	s.DisplayName = &v
	return s
}

func (s *GetDataSourceResponseBody) SetEndpoint(v string) *GetDataSourceResponseBody {
	s.Endpoint = &v
	return s
}

func (s *GetDataSourceResponseBody) SetFileSystemId(v string) *GetDataSourceResponseBody {
	s.FileSystemId = &v
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

func (s *GetDataSourceResponseBody) SetMountPath(v string) *GetDataSourceResponseBody {
	s.MountPath = &v
	return s
}

func (s *GetDataSourceResponseBody) SetOptions(v string) *GetDataSourceResponseBody {
	s.Options = &v
	return s
}

func (s *GetDataSourceResponseBody) SetPath(v string) *GetDataSourceResponseBody {
	s.Path = &v
	return s
}

func (s *GetDataSourceResponseBody) SetRequestId(v string) *GetDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataSourceResponseBody) SetUserId(v string) *GetDataSourceResponseBody {
	s.UserId = &v
	return s
}

type GetDataSourceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDataSourceResponse) SetStatusCode(v int32) *GetDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDataSourceResponse) SetBody(v *GetDataSourceResponseBody) *GetDataSourceResponse {
	s.Body = v
	return s
}

type GetJobResponseBody struct {
	// 集群ID
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// 代码源配置
	CodeSource *GetJobResponseBodyCodeSource `json:"CodeSource,omitempty" xml:"CodeSource,omitempty" type:"Struct"`
	// 数据源配置列表
	DataSources []*GetJobResponseBodyDataSources `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// 作业显示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 作业运行时长（s）
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 弹性任务参数
	ElasticSpec *JobElasticSpec `json:"ElasticSpec,omitempty" xml:"ElasticSpec,omitempty"`
	// 是否开启debugger任务
	EnabledDebugger *bool `json:"EnabledDebugger,omitempty" xml:"EnabledDebugger,omitempty"`
	// 环境变量配置
	Envs map[string]*string `json:"Envs,omitempty" xml:"Envs,omitempty"`
	// 作业创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	GmtFailedTime *string `json:"GmtFailedTime,omitempty" xml:"GmtFailedTime,omitempty"`
	// 作业结束时间（UTC）
	GmtFinishTime    *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	GmtRunningTime   *string `json:"GmtRunningTime,omitempty" xml:"GmtRunningTime,omitempty"`
	GmtStoppedTime   *string `json:"GmtStoppedTime,omitempty" xml:"GmtStoppedTime,omitempty"`
	GmtSubmittedTime *string `json:"GmtSubmittedTime,omitempty" xml:"GmtSubmittedTime,omitempty"`
	GmtSuccessedTime *string `json:"GmtSuccessedTime,omitempty" xml:"GmtSuccessedTime,omitempty"`
	// 作业Id
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 作业规格配置
	JobSpecs []*JobSpec `json:"JobSpecs,omitempty" xml:"JobSpecs,omitempty" type:"Repeated"`
	// 作业类型
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// 作业所以运行Pod列表
	Pods []*GetJobResponseBodyPods `json:"Pods,omitempty" xml:"Pods,omitempty" type:"Repeated"`
	// 状态详情码
	ReasonCode *string `json:"ReasonCode,omitempty" xml:"ReasonCode,omitempty"`
	// 状态详情
	ReasonMessage *string `json:"ReasonMessage,omitempty" xml:"ReasonMessage,omitempty"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 作业运行所在的资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 作业运行时使用的资源级别
	ResourceLevel *string `json:"ResourceLevel,omitempty" xml:"ResourceLevel,omitempty"`
	// 作业额外参数配置
	Settings *JobSettings `json:"Settings,omitempty" xml:"Settings,omitempty"`
	// 作业状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 三方库(requirements.txt)文件路径
	ThirdpartyLibDir *string `json:"ThirdpartyLibDir,omitempty" xml:"ThirdpartyLibDir,omitempty"`
	// 三方库配置列表
	ThirdpartyLibs []*string `json:"ThirdpartyLibs,omitempty" xml:"ThirdpartyLibs,omitempty" type:"Repeated"`
	// 用户命令
	UserCommand *string `json:"UserCommand,omitempty" xml:"UserCommand,omitempty"`
	// 作业提交人Id
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// 作业所属工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	// 作业所属工作空间名称
	WorkspaceName *string `json:"WorkspaceName,omitempty" xml:"WorkspaceName,omitempty"`
}

func (s GetJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobResponseBody) SetClusterId(v string) *GetJobResponseBody {
	s.ClusterId = &v
	return s
}

func (s *GetJobResponseBody) SetCodeSource(v *GetJobResponseBodyCodeSource) *GetJobResponseBody {
	s.CodeSource = v
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

func (s *GetJobResponseBody) SetSettings(v *JobSettings) *GetJobResponseBody {
	s.Settings = v
	return s
}

func (s *GetJobResponseBody) SetStatus(v string) *GetJobResponseBody {
	s.Status = &v
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

func (s *GetJobResponseBody) SetWorkspaceId(v string) *GetJobResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetJobResponseBody) SetWorkspaceName(v string) *GetJobResponseBody {
	s.WorkspaceName = &v
	return s
}

type GetJobResponseBodyCodeSource struct {
	// 代码分支
	Branch *string `json:"Branch,omitempty" xml:"Branch,omitempty"`
	// 代码源Id
	CodeSourceId *string `json:"CodeSourceId,omitempty" xml:"CodeSourceId,omitempty"`
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

type GetJobResponseBodyPods struct {
	// Pod创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Pod结束时间（UTC）
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// Pod启动时间（UTC）
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// 历史Pods
	HistoryPods []*GetJobResponseBodyPodsHistoryPods `json:"HistoryPods,omitempty" xml:"HistoryPods,omitempty" type:"Repeated"`
	// Pod Ip
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// Pod Id
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// Pod UId
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// Pod状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Pod类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetJobResponseBodyPods) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyPods) GoString() string {
	return s.String()
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

func (s *GetJobResponseBodyPods) SetStatus(v string) *GetJobResponseBodyPods {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyPods) SetType(v string) *GetJobResponseBodyPods {
	s.Type = &v
	return s
}

type GetJobResponseBodyPodsHistoryPods struct {
	// Pod创建时间（UTC）
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// Pod结束时间（UTC）
	GmtFinishTime *string `json:"GmtFinishTime,omitempty" xml:"GmtFinishTime,omitempty"`
	// Pod启动时间（UTC）
	GmtStartTime *string `json:"GmtStartTime,omitempty" xml:"GmtStartTime,omitempty"`
	// Pod Ip
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// Pod Id
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// Pod UId
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// Pod状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Pod类型
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetJobResponseBodyPodsHistoryPods) String() string {
	return tea.Prettify(s)
}

func (s GetJobResponseBodyPodsHistoryPods) GoString() string {
	return s.String()
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

func (s *GetJobResponseBodyPodsHistoryPods) SetStatus(v string) *GetJobResponseBodyPodsHistoryPods {
	s.Status = &v
	return s
}

func (s *GetJobResponseBodyPodsHistoryPods) SetType(v string) *GetJobResponseBodyPodsHistoryPods {
	s.Type = &v
	return s
}

type GetJobResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetJobResponse) SetStatusCode(v int32) *GetJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobResponse) SetBody(v *GetJobResponseBody) *GetJobResponse {
	s.Body = v
	return s
}

type GetJobEventsRequest struct {
	// 查询事件的时间区间的截止时间，默认值是当前。
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 获取事件的最大数目，默认值：2000
	MaxEventsNum *int32 `json:"MaxEventsNum,omitempty" xml:"MaxEventsNum,omitempty"`
	// 查询事件的时间区间的起始时间，默认值是7天前。
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetJobEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobEventsRequest) GoString() string {
	return s.String()
}

func (s *GetJobEventsRequest) SetEndTime(v string) *GetJobEventsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJobEventsRequest) SetMaxEventsNum(v int32) *GetJobEventsRequest {
	s.MaxEventsNum = &v
	return s
}

func (s *GetJobEventsRequest) SetStartTime(v string) *GetJobEventsRequest {
	s.StartTime = &v
	return s
}

type GetJobEventsResponseBody struct {
	// 事件
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// 作业ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 请求ID
	RequestId *int32 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetJobEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetJobEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJobEventsResponseBody) SetEvents(v []*string) *GetJobEventsResponseBody {
	s.Events = v
	return s
}

func (s *GetJobEventsResponseBody) SetJobId(v string) *GetJobEventsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetJobEventsResponseBody) SetRequestId(v int32) *GetJobEventsResponseBody {
	s.RequestId = &v
	return s
}

type GetJobEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetJobEventsResponse) SetStatusCode(v int32) *GetJobEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobEventsResponse) SetBody(v *GetJobEventsResponseBody) *GetJobEventsResponse {
	s.Body = v
	return s
}

type GetJobMetricsRequest struct {
	// 截止时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 指标类型
	MetricType *string `json:"MetricType,omitempty" xml:"MetricType,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 时间间隔
	TimeStep *string `json:"TimeStep,omitempty" xml:"TimeStep,omitempty"`
	// Token
	Token *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s GetJobMetricsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetJobMetricsRequest) GoString() string {
	return s.String()
}

func (s *GetJobMetricsRequest) SetEndTime(v string) *GetJobMetricsRequest {
	s.EndTime = &v
	return s
}

func (s *GetJobMetricsRequest) SetMetricType(v string) *GetJobMetricsRequest {
	s.MetricType = &v
	return s
}

func (s *GetJobMetricsRequest) SetStartTime(v string) *GetJobMetricsRequest {
	s.StartTime = &v
	return s
}

func (s *GetJobMetricsRequest) SetTimeStep(v string) *GetJobMetricsRequest {
	s.TimeStep = &v
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetJobMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetJobMetricsResponse) SetStatusCode(v int32) *GetJobMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetJobMetricsResponse) SetBody(v *GetJobMetricsResponseBody) *GetJobMetricsResponse {
	s.Body = v
	return s
}

type GetPodEventsRequest struct {
	// 截止时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 返回的事件最大数量
	MaxEventsNum *int32 `json:"MaxEventsNum,omitempty" xml:"MaxEventsNum,omitempty"`
	// 运行实例UID
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetPodEventsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPodEventsRequest) GoString() string {
	return s.String()
}

func (s *GetPodEventsRequest) SetEndTime(v string) *GetPodEventsRequest {
	s.EndTime = &v
	return s
}

func (s *GetPodEventsRequest) SetMaxEventsNum(v int32) *GetPodEventsRequest {
	s.MaxEventsNum = &v
	return s
}

func (s *GetPodEventsRequest) SetPodUid(v string) *GetPodEventsRequest {
	s.PodUid = &v
	return s
}

func (s *GetPodEventsRequest) SetStartTime(v string) *GetPodEventsRequest {
	s.StartTime = &v
	return s
}

type GetPodEventsResponseBody struct {
	// 事件列表
	Events []*string `json:"Events,omitempty" xml:"Events,omitempty" type:"Repeated"`
	// 作业ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 运行示例ID
	PodId  *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetPodEventsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPodEventsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPodEventsResponseBody) SetEvents(v []*string) *GetPodEventsResponseBody {
	s.Events = v
	return s
}

func (s *GetPodEventsResponseBody) SetJobId(v string) *GetPodEventsResponseBody {
	s.JobId = &v
	return s
}

func (s *GetPodEventsResponseBody) SetPodId(v string) *GetPodEventsResponseBody {
	s.PodId = &v
	return s
}

func (s *GetPodEventsResponseBody) SetPodUid(v string) *GetPodEventsResponseBody {
	s.PodUid = &v
	return s
}

func (s *GetPodEventsResponseBody) SetRequestId(v string) *GetPodEventsResponseBody {
	s.RequestId = &v
	return s
}

type GetPodEventsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPodEventsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPodEventsResponse) SetStatusCode(v int32) *GetPodEventsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPodEventsResponse) SetBody(v *GetPodEventsResponseBody) *GetPodEventsResponse {
	s.Body = v
	return s
}

type GetPodLogsRequest struct {
	// 是否下载日志文件，默认：false。
	DownloadToFile *bool `json:"DownloadToFile,omitempty" xml:"DownloadToFile,omitempty"`
	// 查询的截止时间，默认值：当前。
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// 返回的日志的最大行数，默认值：2000。
	MaxLines *int32  `json:"MaxLines,omitempty" xml:"MaxLines,omitempty"`
	PodUid   *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
	// 查询的起始时间，默认值：7天前。
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetPodLogsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPodLogsRequest) GoString() string {
	return s.String()
}

func (s *GetPodLogsRequest) SetDownloadToFile(v bool) *GetPodLogsRequest {
	s.DownloadToFile = &v
	return s
}

func (s *GetPodLogsRequest) SetEndTime(v string) *GetPodLogsRequest {
	s.EndTime = &v
	return s
}

func (s *GetPodLogsRequest) SetMaxLines(v int32) *GetPodLogsRequest {
	s.MaxLines = &v
	return s
}

func (s *GetPodLogsRequest) SetPodUid(v string) *GetPodLogsRequest {
	s.PodUid = &v
	return s
}

func (s *GetPodLogsRequest) SetStartTime(v string) *GetPodLogsRequest {
	s.StartTime = &v
	return s
}

type GetPodLogsResponseBody struct {
	// 作业ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 日志列表
	Logs []*string `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// 实例ID
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
	// 实例UID
	PodUid *string `json:"PodUid,omitempty" xml:"PodUid,omitempty"`
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

func (s *GetPodLogsResponseBody) SetLogs(v []*string) *GetPodLogsResponseBody {
	s.Logs = v
	return s
}

func (s *GetPodLogsResponseBody) SetPodId(v string) *GetPodLogsResponseBody {
	s.PodId = &v
	return s
}

func (s *GetPodLogsResponseBody) SetPodUid(v string) *GetPodLogsResponseBody {
	s.PodUid = &v
	return s
}

func (s *GetPodLogsResponseBody) SetRequestId(v string) *GetPodLogsResponseBody {
	s.RequestId = &v
	return s
}

type GetPodLogsResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetPodLogsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetPodLogsResponse) SetStatusCode(v int32) *GetPodLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPodLogsResponse) SetBody(v *GetPodLogsResponseBody) *GetPodLogsResponse {
	s.Body = v
	return s
}

type GetTensorboardRequest struct {
	// JodId
	JodId *string `json:"JodId,omitempty" xml:"JodId,omitempty"`
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTensorboardRequest) GoString() string {
	return s.String()
}

func (s *GetTensorboardRequest) SetJodId(v string) *GetTensorboardRequest {
	s.JodId = &v
	return s
}

func (s *GetTensorboardRequest) SetWorkspaceId(v string) *GetTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type GetTensorboardResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Tensorboard       `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetTensorboardResponse) SetStatusCode(v int32) *GetTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTensorboardResponse) SetBody(v *Tensorboard) *GetTensorboardResponse {
	s.Body = v
	return s
}

type ListCodeSourcesRequest struct {
	// 代码源显示名称，支持模糊匹配
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 排序顺序, 枚举值 desc 或者 asc
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 取第几页的数据
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 分页大小
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 用于排序的字段名，可选字段名：DisplayName、GmtCreateTime、 GmtModifyTime
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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

func (s *ListCodeSourcesRequest) SetOrder(v string) *ListCodeSourcesRequest {
	s.Order = &v
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

type ListCodeSourcesResponseBody struct {
	// 代码源配置列表
	CodeSources []*CodeSourceItem `json:"CodeSources,omitempty" xml:"CodeSources,omitempty" type:"Repeated"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合过滤条件的代码源配置的总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListCodeSourcesResponseBody) SetRequestId(v string) *ListCodeSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCodeSourcesResponseBody) SetTotalCount(v int64) *ListCodeSourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListCodeSourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListCodeSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListCodeSourcesResponse) SetStatusCode(v int32) *ListCodeSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCodeSourcesResponse) SetBody(v *ListCodeSourcesResponseBody) *ListCodeSourcesResponse {
	s.Body = v
	return s
}

type ListDataSourcesRequest struct {
	// 数据源类型
	DataSourceType *string `json:"DataSourceType,omitempty" xml:"DataSourceType,omitempty"`
	// 数据源显示名称，支持模糊查询
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 查询第几页数据
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 设置查询的分页大写
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
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

func (s *ListDataSourcesRequest) SetOrder(v string) *ListDataSourcesRequest {
	s.Order = &v
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

type ListDataSourcesResponseBody struct {
	// 数据源配置列表
	DataSources []*DataSourceItem `json:"DataSources,omitempty" xml:"DataSources,omitempty" type:"Repeated"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合条件的数据源总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourcesResponseBody) SetTotalCount(v int64) *ListDataSourcesResponseBody {
	s.TotalCount = &v
	return s
}

type ListDataSourcesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDataSourcesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListDataSourcesResponse) SetStatusCode(v int32) *ListDataSourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataSourcesResponse) SetBody(v *ListDataSourcesResponseBody) *ListDataSourcesResponse {
	s.Body = v
	return s
}

type ListEcsSpecsRequest struct {
	// 按加速器类型过滤
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 查询第几页数据
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 设置查询的分页大写
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
}

func (s ListEcsSpecsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsRequest) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsRequest) SetAcceleratorType(v string) *ListEcsSpecsRequest {
	s.AcceleratorType = &v
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

func (s *ListEcsSpecsRequest) SetSortBy(v string) *ListEcsSpecsRequest {
	s.SortBy = &v
	return s
}

type ListEcsSpecsResponseBody struct {
	// ECS规格列表
	EcsSpecs []*EcsSpec `json:"EcsSpecs,omitempty" xml:"EcsSpecs,omitempty" type:"Repeated"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合过滤条件的总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListEcsSpecsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListEcsSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEcsSpecsResponseBody) SetEcsSpecs(v []*EcsSpec) *ListEcsSpecsResponseBody {
	s.EcsSpecs = v
	return s
}

func (s *ListEcsSpecsResponseBody) SetRequestId(v string) *ListEcsSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEcsSpecsResponseBody) SetTotalCount(v int64) *ListEcsSpecsResponseBody {
	s.TotalCount = &v
	return s
}

type ListEcsSpecsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListEcsSpecsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListEcsSpecsResponse) SetStatusCode(v int32) *ListEcsSpecsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEcsSpecsResponse) SetBody(v *ListEcsSpecsResponseBody) *ListEcsSpecsResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	// 加速器类型
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// 镜像包含的框架类型
	Framework *string `json:"Framework,omitempty" xml:"Framework,omitempty"`
	// 镜像类型
	ImageProviderType *string `json:"ImageProviderType,omitempty" xml:"ImageProviderType,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetAcceleratorType(v string) *ListImagesRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListImagesRequest) SetFramework(v string) *ListImagesRequest {
	s.Framework = &v
	return s
}

func (s *ListImagesRequest) SetImageProviderType(v string) *ListImagesRequest {
	s.ImageProviderType = &v
	return s
}

func (s *ListImagesRequest) SetOrder(v string) *ListImagesRequest {
	s.Order = &v
	return s
}

func (s *ListImagesRequest) SetSortBy(v string) *ListImagesRequest {
	s.SortBy = &v
	return s
}

func (s *ListImagesRequest) SetWorkspaceId(v string) *ListImagesRequest {
	s.WorkspaceId = &v
	return s
}

type ListImagesResponseBody struct {
	// 镜像详情列表
	Images []*ImageItem `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 2
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImagesResponseBody) SetTotalCount(v int64) *ListImagesResponseBody {
	s.TotalCount = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListJobsRequest struct {
	// 作业关联用户ID
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// 调用方
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// 作业显示名称，支持模糊查询
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 截止时间
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FromAllWorkspaces *bool   `json:"FromAllWorkspaces,omitempty" xml:"FromAllWorkspaces,omitempty"`
	// 作业类型
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 当前页
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的作业数
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 工作流ID
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// 资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 是否只返回当前登录者所提交的作业
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 作业状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 自定义标签
	Tags map[string]*string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsRequest) GoString() string {
	return s.String()
}

func (s *ListJobsRequest) SetBusinessUserId(v string) *ListJobsRequest {
	s.BusinessUserId = &v
	return s
}

func (s *ListJobsRequest) SetCaller(v string) *ListJobsRequest {
	s.Caller = &v
	return s
}

func (s *ListJobsRequest) SetDisplayName(v string) *ListJobsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListJobsRequest) SetEndTime(v string) *ListJobsRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobsRequest) SetFromAllWorkspaces(v bool) *ListJobsRequest {
	s.FromAllWorkspaces = &v
	return s
}

func (s *ListJobsRequest) SetJobType(v string) *ListJobsRequest {
	s.JobType = &v
	return s
}

func (s *ListJobsRequest) SetOrder(v string) *ListJobsRequest {
	s.Order = &v
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

func (s *ListJobsRequest) SetPipelineId(v string) *ListJobsRequest {
	s.PipelineId = &v
	return s
}

func (s *ListJobsRequest) SetResourceId(v string) *ListJobsRequest {
	s.ResourceId = &v
	return s
}

func (s *ListJobsRequest) SetShowOwn(v bool) *ListJobsRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListJobsRequest) SetSortBy(v string) *ListJobsRequest {
	s.SortBy = &v
	return s
}

func (s *ListJobsRequest) SetStartTime(v string) *ListJobsRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobsRequest) SetStatus(v string) *ListJobsRequest {
	s.Status = &v
	return s
}

func (s *ListJobsRequest) SetTags(v map[string]*string) *ListJobsRequest {
	s.Tags = v
	return s
}

func (s *ListJobsRequest) SetWorkspaceId(v string) *ListJobsRequest {
	s.WorkspaceId = &v
	return s
}

type ListJobsShrinkRequest struct {
	// 作业关联用户ID
	BusinessUserId *string `json:"BusinessUserId,omitempty" xml:"BusinessUserId,omitempty"`
	// 调用方
	Caller *string `json:"Caller,omitempty" xml:"Caller,omitempty"`
	// 作业显示名称，支持模糊查询
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 截止时间
	EndTime           *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FromAllWorkspaces *bool   `json:"FromAllWorkspaces,omitempty" xml:"FromAllWorkspaces,omitempty"`
	// 作业类型
	JobType *string `json:"JobType,omitempty" xml:"JobType,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 当前页
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的作业数
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 工作流ID
	PipelineId *string `json:"PipelineId,omitempty" xml:"PipelineId,omitempty"`
	// 资源组ID
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 是否只返回当前登录者所提交的作业
	ShowOwn *bool `json:"ShowOwn,omitempty" xml:"ShowOwn,omitempty"`
	// 按返回字段排序
	SortBy *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 作业状态
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// 自定义标签
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListJobsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListJobsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListJobsShrinkRequest) SetBusinessUserId(v string) *ListJobsShrinkRequest {
	s.BusinessUserId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetCaller(v string) *ListJobsShrinkRequest {
	s.Caller = &v
	return s
}

func (s *ListJobsShrinkRequest) SetDisplayName(v string) *ListJobsShrinkRequest {
	s.DisplayName = &v
	return s
}

func (s *ListJobsShrinkRequest) SetEndTime(v string) *ListJobsShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ListJobsShrinkRequest) SetFromAllWorkspaces(v bool) *ListJobsShrinkRequest {
	s.FromAllWorkspaces = &v
	return s
}

func (s *ListJobsShrinkRequest) SetJobType(v string) *ListJobsShrinkRequest {
	s.JobType = &v
	return s
}

func (s *ListJobsShrinkRequest) SetOrder(v string) *ListJobsShrinkRequest {
	s.Order = &v
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

func (s *ListJobsShrinkRequest) SetPipelineId(v string) *ListJobsShrinkRequest {
	s.PipelineId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetResourceId(v string) *ListJobsShrinkRequest {
	s.ResourceId = &v
	return s
}

func (s *ListJobsShrinkRequest) SetShowOwn(v bool) *ListJobsShrinkRequest {
	s.ShowOwn = &v
	return s
}

func (s *ListJobsShrinkRequest) SetSortBy(v string) *ListJobsShrinkRequest {
	s.SortBy = &v
	return s
}

func (s *ListJobsShrinkRequest) SetStartTime(v string) *ListJobsShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ListJobsShrinkRequest) SetStatus(v string) *ListJobsShrinkRequest {
	s.Status = &v
	return s
}

func (s *ListJobsShrinkRequest) SetTagsShrink(v string) *ListJobsShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *ListJobsShrinkRequest) SetWorkspaceId(v string) *ListJobsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

type ListJobsResponseBody struct {
	// 作业列表
	Jobs []*JobItem `json:"Jobs,omitempty" xml:"Jobs,omitempty" type:"Repeated"`
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 符合过滤条件的总作业数
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
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

func (s *ListJobsResponseBody) SetRequestId(v string) *ListJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListJobsResponseBody) SetTotalCount(v int64) *ListJobsResponseBody {
	s.TotalCount = &v
	return s
}

type ListJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListJobsResponse) SetStatusCode(v int32) *ListJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListJobsResponse) SetBody(v *ListJobsResponseBody) *ListJobsResponse {
	s.Body = v
	return s
}

type ListTensorboardsRequest struct {
	// 展示名称
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// 截止时间
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// JobId
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 排序顺序
	Order *string `json:"Order,omitempty" xml:"Order,omitempty"`
	// 当前页
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页返回的作业数
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 按返回字段排序
	SortBy     *string `json:"SortBy,omitempty" xml:"SortBy,omitempty"`
	SourceId   *string `json:"SourceId,omitempty" xml:"SourceId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// 起始时间
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// 根据状态过滤
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// TensorboardId
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
	// 是否显示详情
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s ListTensorboardsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTensorboardsRequest) GoString() string {
	return s.String()
}

func (s *ListTensorboardsRequest) SetDisplayName(v string) *ListTensorboardsRequest {
	s.DisplayName = &v
	return s
}

func (s *ListTensorboardsRequest) SetEndTime(v string) *ListTensorboardsRequest {
	s.EndTime = &v
	return s
}

func (s *ListTensorboardsRequest) SetJobId(v string) *ListTensorboardsRequest {
	s.JobId = &v
	return s
}

func (s *ListTensorboardsRequest) SetOrder(v string) *ListTensorboardsRequest {
	s.Order = &v
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

func (s *ListTensorboardsRequest) SetSourceId(v string) *ListTensorboardsRequest {
	s.SourceId = &v
	return s
}

func (s *ListTensorboardsRequest) SetSourceType(v string) *ListTensorboardsRequest {
	s.SourceType = &v
	return s
}

func (s *ListTensorboardsRequest) SetStartTime(v string) *ListTensorboardsRequest {
	s.StartTime = &v
	return s
}

func (s *ListTensorboardsRequest) SetStatus(v string) *ListTensorboardsRequest {
	s.Status = &v
	return s
}

func (s *ListTensorboardsRequest) SetTensorboardId(v string) *ListTensorboardsRequest {
	s.TensorboardId = &v
	return s
}

func (s *ListTensorboardsRequest) SetVerbose(v bool) *ListTensorboardsRequest {
	s.Verbose = &v
	return s
}

func (s *ListTensorboardsRequest) SetWorkspaceId(v string) *ListTensorboardsRequest {
	s.WorkspaceId = &v
	return s
}

type ListTensorboardsResponseBody struct {
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Tensorboard 列表
	Tensorboards []*Tensorboard `json:"Tensorboards,omitempty" xml:"Tensorboards,omitempty" type:"Repeated"`
	// 符合条件的数据源总数量
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTensorboardsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTensorboardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTensorboardsResponseBody) SetRequestId(v string) *ListTensorboardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTensorboardsResponseBody) SetTensorboards(v []*Tensorboard) *ListTensorboardsResponseBody {
	s.Tensorboards = v
	return s
}

func (s *ListTensorboardsResponseBody) SetTotalCount(v int64) *ListTensorboardsResponseBody {
	s.TotalCount = &v
	return s
}

type ListTensorboardsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTensorboardsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTensorboardsResponse) SetStatusCode(v int32) *ListTensorboardsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTensorboardsResponse) SetBody(v *ListTensorboardsResponseBody) *ListTensorboardsResponse {
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
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Tensorboad Id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s StartTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StartTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *StartTensorboardResponseBody) SetRequestId(v string) *StartTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartTensorboardResponseBody) SetTensorboardId(v string) *StartTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

type StartTensorboardResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StartTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StartTensorboardResponse) SetStatusCode(v int32) *StartTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTensorboardResponse) SetBody(v *StartTensorboardResponseBody) *StartTensorboardResponse {
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
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopJobResponse) SetStatusCode(v int32) *StopJobResponse {
	s.StatusCode = &v
	return s
}

func (s *StopJobResponse) SetBody(v *StopJobResponseBody) *StopJobResponse {
	s.Body = v
	return s
}

type StopJobsRequest struct {
	// 作业ID列表
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
}

func (s StopJobsRequest) String() string {
	return tea.Prettify(s)
}

func (s StopJobsRequest) GoString() string {
	return s.String()
}

func (s *StopJobsRequest) SetJobIds(v []*string) *StopJobsRequest {
	s.JobIds = v
	return s
}

type StopJobsResponseBody struct {
	// 作业ID列表
	JobIds []*string `json:"JobIds,omitempty" xml:"JobIds,omitempty" type:"Repeated"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopJobsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopJobsResponseBody) GoString() string {
	return s.String()
}

func (s *StopJobsResponseBody) SetJobIds(v []*string) *StopJobsResponseBody {
	s.JobIds = v
	return s
}

func (s *StopJobsResponseBody) SetRequestId(v string) *StopJobsResponseBody {
	s.RequestId = &v
	return s
}

type StopJobsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopJobsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s StopJobsResponse) String() string {
	return tea.Prettify(s)
}

func (s StopJobsResponse) GoString() string {
	return s.String()
}

func (s *StopJobsResponse) SetHeaders(v map[string]*string) *StopJobsResponse {
	s.Headers = v
	return s
}

func (s *StopJobsResponse) SetStatusCode(v int32) *StopJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *StopJobsResponse) SetBody(v *StopJobsResponseBody) *StopJobsResponse {
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
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Tensorboad Id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s StopTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *StopTensorboardResponseBody) SetRequestId(v string) *StopTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopTensorboardResponseBody) SetTensorboardId(v string) *StopTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

type StopTensorboardResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *StopTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *StopTensorboardResponse) SetStatusCode(v int32) *StopTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTensorboardResponse) SetBody(v *StopTensorboardResponseBody) *StopTensorboardResponse {
	s.Body = v
	return s
}

type UpdateJobRequest struct {
	// 设置优先级
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
}

func (s UpdateJobRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobRequest) GoString() string {
	return s.String()
}

func (s *UpdateJobRequest) SetPriority(v int32) *UpdateJobRequest {
	s.Priority = &v
	return s
}

type UpdateJobResponseBody struct {
	// 作业ID
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// 请求ID
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateJobResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateJobResponseBody) SetJobId(v string) *UpdateJobResponseBody {
	s.JobId = &v
	return s
}

func (s *UpdateJobResponseBody) SetRequestId(v string) *UpdateJobResponseBody {
	s.RequestId = &v
	return s
}

type UpdateJobResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateJobResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s UpdateJobResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateJobResponse) GoString() string {
	return s.String()
}

func (s *UpdateJobResponse) SetHeaders(v map[string]*string) *UpdateJobResponse {
	s.Headers = v
	return s
}

func (s *UpdateJobResponse) SetStatusCode(v int32) *UpdateJobResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateJobResponse) SetBody(v *UpdateJobResponseBody) *UpdateJobResponse {
	s.Body = v
	return s
}

type UpdateTensorboardRequest struct {
	// MaxRunningTimeMinutes
	MaxRunningTimeMinutes *int64 `json:"MaxRunningTimeMinutes,omitempty" xml:"MaxRunningTimeMinutes,omitempty"`
	// 工作空间ID
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s UpdateTensorboardRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTensorboardRequest) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardRequest) SetMaxRunningTimeMinutes(v int64) *UpdateTensorboardRequest {
	s.MaxRunningTimeMinutes = &v
	return s
}

func (s *UpdateTensorboardRequest) SetWorkspaceId(v string) *UpdateTensorboardRequest {
	s.WorkspaceId = &v
	return s
}

type UpdateTensorboardResponseBody struct {
	// 请求Id
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Tensorboad Id
	TensorboardId *string `json:"TensorboardId,omitempty" xml:"TensorboardId,omitempty"`
}

func (s UpdateTensorboardResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateTensorboardResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateTensorboardResponseBody) SetRequestId(v string) *UpdateTensorboardResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateTensorboardResponseBody) SetTensorboardId(v string) *UpdateTensorboardResponseBody {
	s.TensorboardId = &v
	return s
}

type UpdateTensorboardResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *UpdateTensorboardResponse) SetStatusCode(v int32) *UpdateTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTensorboardResponse) SetBody(v *UpdateTensorboardResponseBody) *UpdateTensorboardResponse {
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
	client.SignatureAlgorithm = tea.String("v2")
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
	if !tea.BoolValue(util.IsUnset(request.CodeBranch)) {
		body["CodeBranch"] = request.CodeBranch
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepo)) {
		body["CodeRepo"] = request.CodeRepo
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoAccessToken)) {
		body["CodeRepoAccessToken"] = request.CodeRepoAccessToken
	}

	if !tea.BoolValue(util.IsUnset(request.CodeRepoUserName)) {
		body["CodeRepoUserName"] = request.CodeRepoUserName
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.MountPath)) {
		body["MountPath"] = request.MountPath
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCodeSource"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.Endpoint)) {
		body["Endpoint"] = request.Endpoint
	}

	if !tea.BoolValue(util.IsUnset(request.FileSystemId)) {
		body["FileSystemId"] = request.FileSystemId
	}

	if !tea.BoolValue(util.IsUnset(request.MountPath)) {
		body["MountPath"] = request.MountPath
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.Path)) {
		body["Path"] = request.Path
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDataSource"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasources"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.CodeSource))) {
		body["CodeSource"] = request.CodeSource
	}

	if !tea.BoolValue(util.IsUnset(request.DataSources)) {
		body["DataSources"] = request.DataSources
	}

	if !tea.BoolValue(util.IsUnset(request.DebuggerConfigContent)) {
		body["DebuggerConfigContent"] = request.DebuggerConfigContent
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ElasticSpec))) {
		body["ElasticSpec"] = request.ElasticSpec
	}

	if !tea.BoolValue(util.IsUnset(request.Envs)) {
		body["Envs"] = request.Envs
	}

	if !tea.BoolValue(util.IsUnset(request.JobMaxRunningTimeMinutes)) {
		body["JobMaxRunningTimeMinutes"] = request.JobMaxRunningTimeMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.JobSpecs)) {
		body["JobSpecs"] = request.JobSpecs
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		body["JobType"] = request.JobType
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		body["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.Settings))) {
		body["Settings"] = request.Settings
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartyLibDir)) {
		body["ThirdpartyLibDir"] = request.ThirdpartyLibDir
	}

	if !tea.BoolValue(util.IsUnset(request.ThirdpartyLibs)) {
		body["ThirdpartyLibs"] = request.ThirdpartyLibs
	}

	if !tea.BoolValue(util.IsUnset(request.UserCommand)) {
		body["UserCommand"] = request.UserCommand
	}

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.UserVpc))) {
		body["UserVpc"] = request.UserVpc
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateJob"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DataSourceId)) {
		body["DataSourceId"] = request.DataSourceId
	}

	if !tea.BoolValue(util.IsUnset(request.DataSourceType)) {
		body["DataSourceType"] = request.DataSourceType
	}

	if !tea.BoolValue(util.IsUnset(request.DataSources)) {
		body["DataSources"] = request.DataSources
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		body["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		body["JobId"] = request.JobId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxRunningTimeMinutes)) {
		body["MaxRunningTimeMinutes"] = request.MaxRunningTimeMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.Options)) {
		body["Options"] = request.Options
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		body["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		body["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.SummaryPath)) {
		body["SummaryPath"] = request.SummaryPath
	}

	if !tea.BoolValue(util.IsUnset(request.SummaryRelativePath)) {
		body["SummaryRelativePath"] = request.SummaryRelativePath
	}

	if !tea.BoolValue(util.IsUnset(request.Uri)) {
		body["Uri"] = request.Uri
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		body["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DeleteCodeSource"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources/" + tea.StringValue(CodeSourceId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DeleteDataSource"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasources/" + tea.StringValue(DataSourceId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DeleteJob"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(JobId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteJobs(request *DeleteJobsRequest) (_result *DeleteJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteJobsResponse{}
	_body, _err := client.DeleteJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteJobsWithOptions(request *DeleteJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobIds)) {
		body["JobIds"] = request.JobIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteJobs"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/batch/jobs/delete"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("DeleteTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards/" + tea.StringValue(TensorboardId)),
		Method:      tea.String("DELETE"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetCodeSource"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources/" + tea.StringValue(CodeSourceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetCodeSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetDataSource"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasources/" + tea.StringValue(DataSourceId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDataSourceResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("GetJob"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(JobId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxEventsNum)) {
		query["MaxEventsNum"] = request.MaxEventsNum
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobEvents"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(JobId) + "/events"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MetricType)) {
		query["MetricType"] = request.MetricType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.TimeStep)) {
		query["TimeStep"] = request.TimeStep
	}

	if !tea.BoolValue(util.IsUnset(request.Token)) {
		query["Token"] = request.Token
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetJobMetrics"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(JobId) + "/metrics"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetJobMetricsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxEventsNum)) {
		query["MaxEventsNum"] = request.MaxEventsNum
	}

	if !tea.BoolValue(util.IsUnset(request.PodUid)) {
		query["PodUid"] = request.PodUid
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPodEvents"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(JobId) + "/pods/" + tea.StringValue(PodId) + "/events"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPodEventsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.DownloadToFile)) {
		query["DownloadToFile"] = request.DownloadToFile
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.MaxLines)) {
		query["MaxLines"] = request.MaxLines
	}

	if !tea.BoolValue(util.IsUnset(request.PodUid)) {
		query["PodUid"] = request.PodUid
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetPodLogs"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(JobId) + "/pods/" + tea.StringValue(PodId) + "/logs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetPodLogsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.JodId)) {
		query["JodId"] = request.JodId
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards/" + tea.StringValue(TensorboardId)),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListCodeSources"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/codesources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListCodeSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
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

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
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

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDataSources"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/datasources"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDataSourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AcceleratorType)) {
		query["AcceleratorType"] = request.AcceleratorType
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

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListEcsSpecs"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/ecsspecs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListEcsSpecsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.AcceleratorType)) {
		query["AcceleratorType"] = request.AcceleratorType
	}

	if !tea.BoolValue(util.IsUnset(request.Framework)) {
		query["Framework"] = request.Framework
	}

	if !tea.BoolValue(util.IsUnset(request.ImageProviderType)) {
		query["ImageProviderType"] = request.ImageProviderType
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		query["Order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/images"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.BusinessUserId)) {
		query["BusinessUserId"] = request.BusinessUserId
	}

	if !tea.BoolValue(util.IsUnset(request.Caller)) {
		query["Caller"] = request.Caller
	}

	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.FromAllWorkspaces)) {
		query["FromAllWorkspaces"] = request.FromAllWorkspaces
	}

	if !tea.BoolValue(util.IsUnset(request.JobType)) {
		query["JobType"] = request.JobType
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

	if !tea.BoolValue(util.IsUnset(request.PipelineId)) {
		query["PipelineId"] = request.PipelineId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ShowOwn)) {
		query["ShowOwn"] = request.ShowOwn
	}

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TagsShrink)) {
		query["Tags"] = request.TagsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListJobs"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.DisplayName)) {
		query["DisplayName"] = request.DisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.EndTime)) {
		query["EndTime"] = request.EndTime
	}

	if !tea.BoolValue(util.IsUnset(request.JobId)) {
		query["JobId"] = request.JobId
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

	if !tea.BoolValue(util.IsUnset(request.SortBy)) {
		query["SortBy"] = request.SortBy
	}

	if !tea.BoolValue(util.IsUnset(request.SourceId)) {
		query["SourceId"] = request.SourceId
	}

	if !tea.BoolValue(util.IsUnset(request.SourceType)) {
		query["SourceType"] = request.SourceType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		query["StartTime"] = request.StartTime
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		query["Status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TensorboardId)) {
		query["TensorboardId"] = request.TensorboardId
	}

	if !tea.BoolValue(util.IsUnset(request.Verbose)) {
		query["Verbose"] = request.Verbose
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTensorboards"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTensorboardsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("StartTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards/" + tea.StringValue(TensorboardId) + "/start"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StartTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("StopJob"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(JobId) + "/stop"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) StopJobs(request *StopJobsRequest) (_result *StopJobsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StopJobsResponse{}
	_body, _err := client.StopJobsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) StopJobsWithOptions(request *StopJobsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StopJobsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.JobIds)) {
		body["JobIds"] = request.JobIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopJobs"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/batch/jobs/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopJobsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	params := &openapi.Params{
		Action:      tea.String("StopTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards/" + tea.StringValue(TensorboardId) + "/stop"),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &StopTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) UpdateJob(JobId *string, request *UpdateJobRequest) (_result *UpdateJobResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateJobResponse{}
	_body, _err := client.UpdateJobWithOptions(JobId, request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) UpdateJobWithOptions(JobId *string, request *UpdateJobRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateJobResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	JobId = openapiutil.GetEncodeParam(JobId)
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Priority)) {
		body["Priority"] = request.Priority
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateJob"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/jobs/" + tea.StringValue(JobId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateJobResponse{}
	_body, _err := client.CallApi(params, req, runtime)
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
	if !tea.BoolValue(util.IsUnset(request.MaxRunningTimeMinutes)) {
		query["MaxRunningTimeMinutes"] = request.MaxRunningTimeMinutes
	}

	if !tea.BoolValue(util.IsUnset(request.WorkspaceId)) {
		query["WorkspaceId"] = request.WorkspaceId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateTensorboard"),
		Version:     tea.String("2020-12-03"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/api/v1/tensorboards/" + tea.StringValue(TensorboardId)),
		Method:      tea.String("PUT"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateTensorboardResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
